package datasets

import (
	"archive/zip"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	_nethttp "net/http"
	"os"

	openapi "datasets_cli/v1/openapi"
	"github.com/gosuri/uiprogress"

	"github.com/spf13/cobra"
)

var (
	argDownloadFilename  string
	argJsonInputFilename string
	argChromosomes       []string
	argExcludeSeq        bool
	argExcludeRna        bool
	argExcludeGene       bool
	argExcludeProtein    bool
	argExcludeGff3       bool
	argExcludeCds        bool
	argIncludeGbff       bool
	argIncludeGtf        bool
	argDehydrated        bool
	argIncludeCds        bool
	argInclude5putr      bool
	argInclude3putr      bool
)

func downloadData(f *os.File, resp *_nethttp.Response, inError error, filename string, length int64) (err error) {
	if inError != nil {
		err = fmt.Errorf("Error connecting to service: %s", inError)
		return
	}
	if resp.StatusCode == 200 {
		defer resp.Body.Close()
		progressBar := &copyProgressBar{}
		progressBar.filename = filename
		if _, e := progressBar.Copy(f, resp.Body); e != nil {
			progressBar.status = "error"
			err = fmt.Errorf("Download error: %s", e)
			return
		}
		if !isValidZip(filename) {
			os.Remove(filename)
			err = errors.New("Internal error (invalid zip archive). Please try again")
			return
		}
	} else if resp.StatusCode == 404 {
		err = fmt.Errorf("request does not match any items in our database")
	} else if resp.StatusCode == 429 {
		msg := `
Selected items are too large for direct download. Please add '--dehydrated' to the
command to download a zip archive with links to the required data. The full dataset can then
be retrieved by unzipping that file and then executing the command:
datasets rehydrate ncbi_dataset`
		err = fmt.Errorf(msg)
	} else {
		err = fmt.Errorf("Unexpected Error: %s", resp.Status)
	}
	return
}

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "download a gene, genome or coronavirus dataset as a zip file",
	Long: `
Download genome, gene and coronavirus data packages, including sequence, annotation, and metadata, as a zip file.

Refer to NCBI's [command line quickstart](https://www.ncbi.nlm.nih.gov/datasets/docs/quickstarts/command-line-tools/) documentation for information about getting started with the command-line tools.`,
	Example: `  datasets download genome accession GCF_000001405.39 --chromosomes X,Y --exclude-gff3 --exclude-rna
  datasets download genome taxon "bos taurus"
  datasets download gene gene-id 672
  datasets download gene symbol brca1 --taxon mouse
  datasets download gene accession NP_000483.3
  datasets download virus genome taxon sars-cov-2 --host dog
  datasets download virus protein S --host dog --filename SARS2-spike-dog.zip
  datasets download --input-json request_file.json --filename output.zip`,
	Args: cobra.ExactArgs(0),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
		argIDArgs, err = getArgsFromListOrFile(args, argInputFile)
		if err != nil {
			return
		}
		if !argNoProgress {
			uiprogress.Start()
		}
		return
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		if !argNoProgress {
			uiprogress.Stop()
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
		req := openapi.NewV1DatasetRequest()
		err = json.Unmarshal([]byte(content), &req)
		if err != nil {
			return fmt.Errorf("Parsing JSON request file %s: \"%s\"", argJsonInputFilename, err)
		}

		if req.GenomeV1 != nil {
			return downloadAssembly(req.GenomeV1, argDownloadFilename)
		} else if req.GeneV1 != nil {
			return downloadGeneForRequest(req.GeneV1)
		} else {
			return errors.New("request did not have a valid gene or genome request object")
		}
	},
}

func isValidZip(filename string) bool {
	_, err := zip.OpenReader(filename)
	if err != nil {
		return false
	}
	return true
}

func init() {
	downloadCmd.AddCommand(downloadGeneCmd)
	downloadCmd.AddCommand(genomeCmd)
	downloadCmd.AddCommand(virusCmd)
	downloadCmd.AddCommand(assemblyCmd)
	downloadCmd.AddCommand(downloadOrthologCmd)

	pflags := downloadCmd.PersistentFlags()
	lflags := downloadCmd.Flags()
	registerHiddenStringPair(pflags, &argDownloadFilename, "filename", "f", "ncbi_dataset.zip", "specify a custom file name for the downloaded dataset")
	registerHiddenStringPair(lflags, &argJsonInputFilename, "input-json", "j", "", "a file that contains a valid json request object for genome or gene queries")
}
