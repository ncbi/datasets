package datasets

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	datasets_util "datasets_cli/v1/util"

	"github.com/gosuri/uiprogress"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

var (
	argPackageDir   string
	argMatchPattern string
	argListOnly     bool
	argNumWorkers   int
	afs             afero.Fs
)

type fetchLine struct {
	url  string
	size string
	path string
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// downloadCmd represents the download command
var rehydrateCmd = &cobra.Command{
	Use:   "rehydrate [flags] --directory <directory_name>",
	Short: "rehydrate a downloaded, dehydrated dataset",
	Long: `
Retrieve data files for an [unzipped, dehydrated zip archive](https://www.ncbi.nlm.nih.gov/datasets/docs/how-tos/genomes/rehydrate-package/).  Data files specified in fetch.txt will be downloaded from NCBI.
`,
	Args:         cobra.MaximumNArgs(0),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		fi, e := afs.Stat(argPackageDir)
		if e != nil {
			return fmt.Errorf("opening '%s' for rehydration: %s", argPackageDir, e.Error())
		}

		var file fetchLine
		var fileList []fetchLine

		switch mode := fi.Mode(); {
		case mode.IsDir():
			// do directory stuff
			fetchFile := path.Join(argPackageDir, "ncbi_dataset", "fetch.txt")
			fp, e := afs.Open(fetchFile)
			if e != nil {
				return fmt.Errorf("opening file '%s' for rehydration: %s", fetchFile, e.Error())
			}
			defer fp.Close()
			reader := csv.NewReader(fp)
			reader.Comma = '\t'
			lines, e := reader.ReadAll()
			if e != nil {
				return e
			}
			var proxyURL *url.URL
			if argProxyURL != "" { // parse the proxy url
				var err error
				proxyURL, err = url.Parse(argProxyURL)
				if err != nil {
					return err
				}
			}
			for _, line := range lines {
				// Filter here
				if argMatchPattern != "" {
					matched, e := regexp.MatchString(argMatchPattern, line[2])
					if e != nil {
						return e
					} else if !matched {
						continue
					}
				}

				// Update fetch url if using a non-production proxy
				file.url = patchURL(line[0], proxyURL)
				file.size = line[1]
				file.path = path.Join(argPackageDir, "ncbi_dataset", line[2])
				fileList = append(fileList, file)
				if argListOnly {
					fmt.Println(line[2])
				}
			}
		case mode.IsRegular():
			// do zip file stuff
			return errors.New("zip file not supported.  please extract")
		}

		fmt.Printf("Found %d files for rehydration\n", len(fileList))
		if !argListOnly {
			return downloadMultipleFiles(fileList)
		}
		return nil
	},
}

func patchURL(fileURL string, proxy *url.URL) string {
	if proxy == nil {
		return fileURL
	}
	// parse the fetchable url
	parsedURL, err := url.Parse(fileURL)
	if err != nil {
		log.Fatal(err)
	}

	// only change if this points to the production website
	if parsedURL.Host != "api.ncbi.nlm.nih.gov" {
		return fileURL
	}

	path := parsedURL.Path
	if !strings.HasPrefix(proxy.Path, "/datasets") {
		path = strings.Replace(parsedURL.Path, "/datasets", "", 1)
	}

	return fmt.Sprintf("%s://%s%s", proxy.Scheme, proxy.Host, path)
}

func processHTTPRequest(client *http.Client, request *http.Request) (*http.Response, error) {
	if argDebug {
		dumpPre, errPre := httputil.DumpRequestOut(request, true)
		if errPre != nil {
			return nil, errPre
		}
		fmt.Printf("\n%s\n", string(dumpPre))
	}
	resp, err := client.Do(request)
	if resp != nil && argDebug {
		dumpPost, errPost := httputil.DumpResponse(resp, false)
		if errPost != nil {
			return resp, errPost
		}
		fmt.Printf("\n%s\n", string(dumpPost))
	}
	if err = datasets_util.HandleHttpResponseWithCustomErr(resp, err, "%s"); err != nil {
		return nil, err
	}
	return resp, err
}

func downloadFileWorker(client *http.Client, bar *uiprogress.Bar, files <-chan fetchLine, errch chan<- error) {
	progressBar := &copyProgressBar{}

	hasError := func(err error) bool {
		if err != nil {
			errch <- err
			progressBar.status = "error"
			progressBar.bar = nil
			if perr, ok := err.(*fs.PathError); ok {
				fmt.Printf("Critical Error: %s\n", perr)
				os.Exit(-1)
			}
			return true
		}
		return false
	}

	for file := range files {
		func() {
			// Get the data
			progressBar.filename = file.path

			req, err := http.NewRequest("GET", file.url, nil)
			if hasError(err) {
				return
			}

			for k, v := range clientHeaders {
				req.Header.Set(k, v)
			}
			resp, err := processHTTPRequest(client, req)
			if err != nil {
				// Custom logging to push error
				newError := errors.New("File unavailable (" + getGatewayRuntimeError(err) + "): " + file.path)
				hasError(newError)
				return
			}
			defer resp.Body.Close()

			// If the directory in which the file resides does not yet exist, create it
			var fileDir = filepath.Dir(file.path)
			if fileDir != "." {
				err := afs.MkdirAll(fileDir, 0755)
				if hasError(err) {
					return
				}
			}

			// Create the file
			out, err := afs.Create(file.path)
			if hasError(err) {
				return
			}
			defer out.Close()

			// Write the body to file
			_, err = progressBar.Copy(out, resp.Body)
			if hasError(err) {
				return
			}
			if !argNoProgress {
				bar.Incr()
			}
			errch <- nil
		}()
	}
}

func downloadMultipleFiles(fileList []fetchLine) error {
	if argNumWorkers < 1 || argNumWorkers > 30 {
		return errors.New("number of workers cannot be less than 1 or greater than 30")
	}

	// Initialize the progress bar
	totalFiles := len(fileList)
	var bar *uiprogress.Bar
	if !argNoProgress {
		bar = uiprogress.AddBar(totalFiles).AppendCompleted()
		bar.Width = 50
		bar.PrependFunc(func(b *uiprogress.Bar) string {
			return fmt.Sprintf("Completed %d of %d", b.Current(), b.Total)
		})
		uiprogress.Start()
	}

	// Create channel to hold all required downloads
	files := make(chan fetchLine, totalFiles)
	// Create error channel for each download
	errch := make(chan error, totalFiles)

	client := initRetryableClient()
	// Create workers to process all fetches
	argNumWorkers = min(totalFiles, argNumWorkers)
	for w := 1; w <= argNumWorkers; w++ {
		go downloadFileWorker(client, bar, files, errch)
	}
	for _, file := range fileList {
		files <- file // add file onto job queue
	}
	close(files)

	var errStr string
	for i := 0; i < totalFiles; i++ {
		if err := <-errch; err != nil {
			errStr = errStr + "\n" + err.Error()
		}
	}
	var err error
	if errStr != "" {
		err = errors.New(errStr)
	}
	if !argNoProgress {
		uiprogress.Stop()
	}
	return err
}

func init() {
	afs = afero.NewOsFs()
	pflags := rehydrateCmd.PersistentFlags()
	registerHiddenStringPair(pflags, &argPackageDir, "directory", "d", "", "specify the directory containing the unzipped dehydrated bag")
	registerHiddenBoolPair(pflags, &argListOnly, "list", "l", false, "list files that would be downloaded during rehydration")
	pflags.StringVar(&argMatchPattern, "match", "", "specify substring that matches files for rehydration")
	registerHiddenIntPair(pflags, &argNumWorkers, "max-workers", "w", 10, "limit the maximum number of concurrent download workers (allowed range is 1-30)")
}
