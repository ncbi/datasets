package datasets

import (
	"archive/zip"
	"errors"
	"fmt"
	_nethttp "net/http"
	"os"

	"github.com/gosuri/uiprogress"
	"github.com/spf13/cobra"
)

var (
	argAssmFilename   string
	argChromosomes    []string
	argExcludeSeq     bool
	argExcludeRna     bool
	argExcludeGene    bool
	argExcludeProtein bool
	argExcludeGff3    bool
	argIncludeGbff    bool
	argIncludeGtf     bool
	argDehydrated     bool
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "download a gene, genome or coronavirus dataset as a zip file",
	Long: `
Download genome, gene and coronavirus datasets, including sequence, annotation, and metadata, as a zip file.

Refer to NCBI's [command line quickstart](https://www.ncbi.nlm.nih.gov/datasets/docs/quickstarts/command-line-tools/) documentation for information about getting started with the command-line tools.`,
	Example: `  datasets download genome accession GCF_000001405.39 --chromosomes X,Y --exclude-gff3 --exclude-rna
  datasets download genome taxon "bos taurus"
  datasets download gene gene-id 672
  datasets download gene symbol brca1 --taxon mouse
  datasets download gene accession NP_000483.3
  datasets download virus genome taxon sars-cov-2 --released-since 05/05/2021
  datasets download virus protein S --released-since 05/05/2021 --filename SARS2-spike-all.zip`,
	Args: cobra.MaximumNArgs(0),
}

func downloadData(f *os.File, resp *_nethttp.Response, inError error, filename string, length int64) (err error) {
	if inError != nil {
		err = fmt.Errorf("Error connecting to service: %s", inError)
		return
	}
	if resp.StatusCode == 200 {
		defer resp.Body.Close()
		progressBar := &copyProgressBar{}
		uiprogress.Start()
		progressBar.filename = filename
		if _, e := progressBar.Copy(f, resp.Body); e != nil {
			progressBar.status = "error"
			uiprogress.Stop()
			err = fmt.Errorf("Download error: %s", e)
			return
		}
		uiprogress.Stop()
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
}
