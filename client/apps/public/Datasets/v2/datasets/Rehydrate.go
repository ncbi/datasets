package datasets

import (
	"compress/gzip"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"syscall"

	"github.com/gosuri/uiprogress"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

var (
	argPackageDir      string
	argMatchPattern    string
	argListOnly        bool
	argNumWorkers      int
	afs                afero.Fs
	argRehydrateAsGzip bool
	tempfileList       sync.Map
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

var rehydrateCmd = &cobra.Command{
	Use:   "rehydrate [flags] --directory <directory_name>",
	Short: "Rehydrate a downloaded, dehydrated dataset",
	Long: `
Download data files for an unzipped, dehydrated genome data package. Data files specified in fetch.txt will be downloaded from NCBI. Read more about how rehydration can help with large genome downloads: https://www.ncbi.nlm.nih.gov/datasets/docs/v2/how-tos/genomes/large-download/`,
	Args: cobra.NoArgs,
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
			if argGatewayURL != "" { // parse the proxy url
				var err error
				proxyURL, err = url.Parse(argGatewayURL)
				if err != nil {
					return err
				}
			}
			for rowIdx, line := range lines {
				// Skip empty lines (all fields empty or whitespace) and comments
				isEmpty := true
				for _, field := range line {
					if strings.TrimSpace(field) != "" {
						isEmpty = false
						break
					}
				}
				if len(line) == 0 || isEmpty || strings.HasPrefix(strings.TrimSpace(line[0]), "#") {
					continue
				}
				if len(line) < 3 {
					fmt.Fprintf(os.Stderr, "invalid fetch.txt row %d: expected at least 3 tab-delimited columns, got %d: %v\n", rowIdx+1, len(line), line)
					continue
				}
				// Only use the first 3 fields, ignore extras
				urlField := line[0]
				sizeField := line[1]
				pathField := line[2]
				if strings.TrimSpace(urlField) == "" {
					fmt.Fprintf(os.Stderr, "invalid fetch.txt row %d: url field is empty: %v\n", rowIdx+1, line)
					continue
				}
				if strings.TrimSpace(pathField) == "" {
					fmt.Fprintf(os.Stderr, "invalid fetch.txt row %d: path field is empty: %v\n", rowIdx+1, line)
					continue
				}
				// Filter here
				if argMatchPattern != "" {
					matched, e := regexp.MatchString(argMatchPattern, pathField)
					if e != nil {
						fmt.Fprintf(os.Stderr, "error matching pattern in fetch.txt row: %v: %v\n", line, e)
						continue
					} else if !matched {
						continue
					}
				}

				// Update fetch url if using a non-production proxy
				file.url = patchURL(urlField, proxyURL)
				file.size = sizeField
				file.path = path.Join(argPackageDir, "ncbi_dataset", filepath.Clean(pathField))
				fileList = append(fileList, file)
				if argListOnly {
					fmt.Println(pathField)
				}
			}
		case mode.IsRegular():
			// do zip file stuff
			return errors.New("zip file not supported.  please extract")
		}

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
	if globalDebugFlag.Debug() {
		dumpPre, errPre := httputil.DumpRequestOut(request, true)
		if errPre != nil {
			return nil, errPre
		}
		fmt.Printf("\n%s\n", string(dumpPre))
	}
	resp, err := client.Do(request)
	if resp != nil && globalDebugFlag.Debug() {
		dumpPost, errPost := httputil.DumpResponse(resp, false)
		if errPost != nil {
			return resp, errPost
		}
		fmt.Printf("\n%s\n", string(dumpPost))
	}
	if err = handleHTTPResponseWithCustomErr(resp, err, "%s"); err != nil {
		return nil, err
	}
	return resp, err
}

func downloadFileWorker(client *http.Client, bar *uiprogress.Bar, files <-chan fetchLine, errch chan<- error, wg *sync.WaitGroup) {
	progressBar := &copyProgressBar{}

	hasError := func(err error, location string, path string) bool {
		if err != nil {
			progressBar.status = "error"
			progressBar.bar = nil
			if perr, ok := err.(*fs.PathError); ok {
				fmt.Printf("Critical Error (%s, %s): %s\n", location, path, perr)
				sigChan <- syscall.SIGTERM
				return false
			}
			errch <- fmt.Errorf("Error (%s, %s): %w", location, path, err)
			return true
		}
		return false
	}

	for file := range files {
		func() {
			wg.Add(1)
			defer wg.Done()

			// Get the data
			progressBar.filename = file.path

			req, err := http.NewRequest("GET", file.url, nil)
			if hasError(err, "NewRequest", file.path) {
				return
			}

			for k, v := range clientHeaders {
				req.Header.Set(k, v)
			}

			// grpc-gateway says to add this to return trailing metadata headers (but didn't work for me)
			// req.Header.Set("TE", "trailers")  # see https://github.com/grpc-ecosystem/grpc-gateway/pull/2124/files

			// Always request the gzipped version of the file for transfer.
			req.Header.Set("X-Datasets-Gzip-Request", "true")
			resp, err := processHTTPRequest(client, req)
			if err != nil {
				// Custom logging to push error
				newError := errors.New("File unavailable (" + getGatewayRuntimeError(err) + "): " + file.path)
				hasError(newError, "processRequest", file.path)
				return
			}
			defer resp.Body.Close()

			// Some files (sequence-reports) have md5s that are computed on-the-fly and so they
			// won't already be in the md5 file. If we can get trailing metadata to return, we could
			// retrieve from the headers, e.g.:
			// fmt.Printf("md5 header: %v\n", resp.Header.Get("Grpc-Metadata-Content-Md5"))

			// Success is indicated with 2xx status codes:
			statusOK := resp.StatusCode >= 200 && resp.StatusCode < 300
			if !statusOK {
				hasError(fmt.Errorf("%s", resp.Status), "Status", file.path)
				return
			}
			// If the directory in which the file resides does not yet exist, create it
			var fileDir = filepath.Dir(file.path)
			if fileDir != "." {
				err := afs.MkdirAll(fileDir, 0755)
				if hasError(err, "MakeDirs", file.path) {
					return
				}
			}

			tempfile := filepath.Clean(file.path) + ".temp"
			if argRehydrateAsGzip && resp.Header.Get("Content-Type") == "application/x-gzip" {
				file.path += ".gz"
			}
			defer afs.Rename(tempfile, file.path) //nolint:errcheck

			// If tempfile already exists, error out and notify user
			if exists, err := afero.Exists(afs, tempfile); err == nil && exists {
				afs.Remove(tempfile) //nolint:errcheck
				if exists, err = afero.Exists(afs, tempfile); err == nil && exists {
					newError := errors.New(tempfile + " may be in use by another rehydration job")
					hasError(newError, "Remove", file.path)
					return
				}
			}

			// Create temporary file at current working directory
			tempfileList.Store(tempfile, nil)
			temp, err := afs.Create(tempfile)
			if hasError(err, "Create", file.path) {
				deleteFile(file.path)
				return
			}
			defer temp.Close()

			w := temp.(io.Writer)

			// Write the body to file
			if argRehydrateAsGzip || resp.Header.Get("Content-Type") != "application/x-gzip" {
				// Json and Jsonl files are transfered as plain/text
				_, err = progressBar.Copy(w, resp.Body)
				if hasError(err, "Copy", file.path) {
					deleteFile(file.path)
					return
				}
			} else {
				gzip_reader, err := gzip.NewReader(resp.Body)
				if hasError(err, "gzRead", file.path) {
					deleteFile(file.path)
					return
				}
				_, err = progressBar.Copy(w, gzip_reader)
				if hasError(err, "Copy decompress", file.path) {
					deleteFile(file.path)
					return
				}
			}

			if !argNoProgress {
				bar.Incr()
			}

			errch <- nil
		}()
	}
}

func deleteFile(f string) {
	if f != "" {
		if exists, err := afero.Exists(afs, f); err == nil && exists {
			err_rm := afs.Remove(f)
			if err_rm != nil && !strings.Contains(err_rm.Error(), "no such file") {
				fmt.Printf("Unable to remove corrupted file %s: %s\n", f, err_rm.Error())
			}
		}
	}
}

func cleanup_tempfiles() {
	tempfileList.Range(func(k, v interface{}) bool {
		path := k.(string)
		if exists, err := afero.Exists(afs, path); err == nil && exists {
			afs.Remove(path) //nolint:errcheck
		}
		return true
	})
}

var sigChan = make(chan os.Signal)

func downloadMultipleFiles(fileList []fetchLine) error {
	if argNumWorkers < 1 || argNumWorkers > 30 {
		return errors.New("number of workers cannot be less than 1 or greater than 30")
	}

	if len(fileList) == 0 {
		fmt.Printf("Found no files for rehydration\n")
		return nil
	}

	var downloadFileList []fetchLine
	for _, file := range fileList {
		if argRehydrateAsGzip && !strings.HasSuffix(file.path, ".jsonl") && !strings.HasSuffix(file.path, ".json") {
			if exists, err := afero.Exists(afs, file.path+".gz"); err != nil || !exists {
				downloadFileList = append(downloadFileList, file)
			}
		} else if exists, err := afero.Exists(afs, file.path); err != nil || !exists {
			downloadFileList = append(downloadFileList, file)
		}
	}

	if len(downloadFileList) == 0 {
		fmt.Printf("All %d files already rehydrated\n", len(fileList))
		return nil
	} else {
		fmt.Printf("Found %d of %d files for rehydration\n", len(downloadFileList), len(fileList))
	}

	// Initialize the progress bar
	totalFiles := len(downloadFileList)
	var bar *uiprogress.Bar

	if !argNoProgress {
		bar = uiprogress.AddBar(len(fileList)).AppendCompleted()
		bar.Set(len(fileList) - len(downloadFileList)) //nolint:errcheck
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
	wg := sync.WaitGroup{}
	argNumWorkers = min(totalFiles, argNumWorkers)

	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT) //nolint:govet
	go func() {
		<-sigChan
		cleanup_tempfiles()
		os.Exit(1)
	}()

	for w := 1; w <= argNumWorkers; w++ {
		go downloadFileWorker(client, bar, files, errch, &wg)
	}
	for _, file := range downloadFileList {
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
	} else if err == nil {
		fmt.Println("Done")
	}
	wg.Wait()
	return err
}

func init() {
	afs = afero.NewOsFs()
	pflags := rehydrateCmd.PersistentFlags()
	pflags.StringVar(&argPackageDir, "directory", "", "Specify the directory containing the unzipped dehydrated bag")
	pflags.BoolVar(&argListOnly, "list", false, "List files that would be downloaded during rehydration")
	pflags.StringVar(&argMatchPattern, "match", "", "Specify substring that matches files for rehydration")
	pflags.IntVar(&argNumWorkers, "max-workers", 10, "Limit the maximum number of concurrent download workers (allowed range is 1-30)")
	pflags.BoolVar(&argRehydrateAsGzip, "gzip", false, "rehydrate files to gzip format")
	pflags.BoolVar(&argNoProgress, "no-progressbar", false, "Hide progress bar")
}
