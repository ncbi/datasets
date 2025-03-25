package datasets

import (
	"archive/zip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	_nethttp "net/http"
	"os"

	openapi "datasets/openapi/v2"
	"github.com/gosuri/uiprogress"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

var (
	argDownloadFilename  string
	argJsonInputFilename string
)

func downloadDataForFile(resp *_nethttp.Response, inError error, filename string, length int64, argSkipZipVal bool) (err error) {
	f, e := afs.Create(filename)
	if e != nil {
		err = fmt.Errorf("'%s' opening output file: %s", e, filename)
		return
	}
	defer f.Close()
	return downloadData(&f, resp, err, filename, length, argSkipZipVal)
}

func downloadData(f *afero.File, resp *_nethttp.Response, inError error, filename string, length int64, argSkipZipVal bool) (err error) {
	if inError != nil {
		err = fmt.Errorf("Error connecting to service: %s", inError)
		return
	}
	if resp.StatusCode == 200 {
		defer resp.Body.Close()
		progressBar := &copyProgressBar{}
		progressBar.filename = filename
		if _, e := progressBar.Copy(*f, resp.Body); e != nil {
			progressBar.status = "error"
			err = fmt.Errorf("Download error: %s", e)
			return
		}
		if !isValidZip(filename, argSkipZipVal) {
			afs.Remove(filename) //nolint:errcheck
			err = errors.New("Internal error (invalid zip archive). Please try again")
			if !argNoProgress {
				progressBar.status = "invalid zip archive"
			}
			return
		}
		if !argNoProgress {
			if argSkipZipVal {
				progressBar.status = "valid zip structure -- files not checked"
			} else {
				progressBar.status = "valid data package"
			}
		}
	} else if resp.StatusCode == 404 {
		err = fmt.Errorf("request does not match any items in our database")
	} else if resp.StatusCode == 429 {
		msg := `
Selected items are too large for direct download. Please add '--dehydrated' to the
command to download a zip archive with links to the required data. The full dataset can then
be retrieved by unzipping that file and then executing the command:
datasets rehydrate ncbi_dataset`
		err = errors.New(msg)
	} else {
		err = fmt.Errorf("Unexpected Error: %s", resp.Status)
	}
	return
}

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download a gene, genome or virus dataset as a zip file",
	Long: `
Download genome, gene and virus data packages, including sequence, annotation, and metadata, as a zip file.

Refer to NCBI's [download and install](https://www.ncbi.nlm.nih.gov/datasets/docs/v2/download-and-install/) documentation for information about getting started with the command-line tools.`,
	Example: `  datasets download genome accession GCF_000001405.40 --chromosomes X,Y --exclude-gff3 --exclude-rna
  datasets download genome taxon "bos taurus"
  datasets download gene gene-id 672
  datasets download gene symbol brca1 --taxon "mus musculus"
  datasets download gene accession NP_000483.3
  datasets download taxonomy taxon human,sars-cov-2
  datasets download virus genome taxon sars-cov-2 --host dog
  datasets download virus protein S --host dog --filename SARS2-spike-dog.zip`,
	Args: cobra.NoArgs,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
		if !argNoProgress {
			progress.Start()
		}
		return nil
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		if !argNoProgress {
			progress.Stop()
		}
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if argJsonInputFilename == "" {
			return errors.New("Must provide a valid json input file")
		}

		content, err := ioutil.ReadFile(argJsonInputFilename)
		if err != nil {
			return fmt.Errorf("Opening json request file: %s: \"%s\"", argJsonInputFilename, err)
		}

		// Convert json string to request structure
		req := openapi.NewV2DatasetRequest()
		err = json.Unmarshal([]byte(content), &req)
		if err != nil {
			return fmt.Errorf("Parsing JSON request file %s: \"%s\"", argJsonInputFilename, err)
		}

		if req.GenomeV2 != nil {
			downloader, err := NewGenomeRequestDownloader(req.GenomeV2, initAssemblyRequestFlag())
			if err != nil {
				return err
			}
			return downloader.Download(false)
		} else if req.GeneV2 != nil {
			downloader, err := NewGeneRequestDownloader(req.GeneV2)
			if err != nil {
				return err
			}
			return downloader.Download(false)
		} else {
			return errors.New("request did not have a valid gene or genome request object")
		}
	},
}

func isValidZip(filename string, argSkipZipVal bool) bool {
	var progressBar *uiprogress.Bar
	progressStatus := "Validating package"
	if !argNoProgress {
		progressBar = progress.AddBar(1)
		progressBar.Width = 2
		progressBar.PrependFunc(func(b *uiprogress.Bar) string { return progressStatus })
	}
	fileInfo, err := afs.Stat(filename)
	if err != nil {
		return false
	}

	file, err := afs.OpenFile(filename, os.O_RDONLY, os.FileMode(0644))
	if err != nil {
		return false
	}

	zipfile, err := zip.NewReader(file, fileInfo.Size())
	if err != nil {
		return false
	}

	if !argNoProgress && progressBar != nil && len(zipfile.File) > 0 {
		progressBar.Total = len(zipfile.File)
		progressBar.Width = 50
		progressBar.AppendCompleted()
		progressBar.AppendFunc(func(b *uiprogress.Bar) string { return fmt.Sprintf("%d/%d", b.Current(), b.Total) })
	}
	for _, zippedfile := range zipfile.File {
		// fmt.Printf("Name=%s, size=%d, crc=%d\n", zippedfile.Name, zippedfile.UncompressedSize64, zippedfile.CRC32)
		r, err := zippedfile.Open()
		if err != nil {
			return false
		}
		if !argSkipZipVal {
			progressStatus = "Validating package files"
			if _, err := io.Copy(io.Discard, r); err != nil {
				return false
			}
		}
		if !argNoProgress && progressBar != nil {
			progressBar.Incr()
		}
	}

	return true
}

func init() {
	downloadCmd.AddCommand(createGeneCmd())
	downloadCmd.AddCommand(createGenomeCmd())
	downloadCmd.AddCommand(createTaxonomyCmd())
	downloadCmd.AddCommand(createVirusCmd())

	pflags := downloadCmd.PersistentFlags()
	lflags := downloadCmd.Flags()
	pflags.StringVar(&argDownloadFilename, "filename", "ncbi_dataset.zip", "Specify a custom file name for the downloaded data package")
	pflags.BoolVar(&argNoProgress, "no-progressbar", false, "Hide progress bar")
	lflags.StringVar(&argJsonInputFilename, "input-json", "", "a file that contains a valid json request object for genome or gene queries")
	if err := lflags.MarkHidden("input-json"); err != nil {
		defaultLogger.Fatalln("Invalid attempt to create hidden flag")
	}
}
