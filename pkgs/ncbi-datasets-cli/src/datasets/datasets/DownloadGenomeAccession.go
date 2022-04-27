package datasets

import (
	"errors"

	"github.com/spf13/cobra"
)

var downloadGenomeAccessionCmd = &cobra.Command{
	Use:   "accession <accession ...>",
	Short: "download a genome dataset by NCBI Assembly or BioProject accession",
	Long: `
Download a genome dataset by NCBI Assembly or BioProject accession. Genome data packages include genome, transcript and protein sequence, annotation and a detailed data report. Datasets are downloaded as a zip file.

The default genome dataset includes the following files (if available):
  * genomic.fna (genomic sequences)
  * rna.fna (transcript sequences)
  * protein.faa (protein sequences)
  * genomic.gff (genome annotation in gff3 format)
  * data_report.jsonl (data report with genome assembly and annotation metadata)
  * dataset_catalog.json (a list of files and file types included in the dataset)

Refer to NCBI's [command line quickstart](https://www.ncbi.nlm.nih.gov/datasets/docs/quickstarts/command-line-tools/) documentation for information about getting started with the command-line tools.`,

	Example: `  datasets download genome accession GCF_000001405.40 --chromosomes X,Y --exclude-gff3 --exclude-rna
  datasets download genome accession GCA_003774525.2 GCA_000001635 --chromosomes X,Y,Un.9
  datasets download genome accession PRJNA289059 --exclude-seq`,

	RunE: func(cmd *cobra.Command, args []string) error {
		validAssmAccs, warning, err := checkAssemblyAvailability(argIDArgs)
		if err != nil {
			return err
		}
		if warning != "" {
			cmd.Println(warning)
		}
		if len(validAssmAccs) == 0 {
			return errors.New("Input accessions not specified")
		}
		downloadOpts := getDownloadRequest(validAssmAccs)
		return downloadAssembly(downloadOpts, argDownloadFilename)
	},
}

func init() {
	pflags := downloadGenomeAccessionCmd.PersistentFlags()
	registerHiddenStringPair(pflags, &argInputFile, "inputfile", "i", "", "read a list of NCBI Assembly accessions from a file to use as input")
	pflags.BoolVar(&argDehydrated, "dehydrated", false, "download a dehydrated bag file including the data report and locations of data files. Use the rehydrate command to retrieve data files when needed")
}
