package datasets

import (
	"errors"

	"github.com/spf13/cobra"
)

var downloadOrthologGeneIDCmd = &cobra.Command{
	Use:     "gene-id <gene-id>",
	Short:   "download an ortholog dataset by NCBI Gene ID",
	Example: `  datasets download ortholog gene-id 672`,
	Long: `
Download an ortholog dataset by NCBI Gene ID. Ortholog data is calculated by NCBI for vertebrates and insects. Ortholog datasets include gene, transcript and protein sequence, a data table and a data report. Datasets are downloaded as a zip file.

The default ortholog dataset includes the following files:
 * gene.fna (gene sequences)
 * rna.fna (transcript sequences)
 * protein.faa (protein sequences)
 * data_report.jsonl (data report with gene metadata)
 * data_table.tsv (data table with gene metadata, one transcript per row)
 * dataset_catalog.json (a list of files and file types included in the dataset)

Refer to NCBI's [command line quickstart](https://www.ncbi.nlm.nih.gov/datasets/docs/quickstarts/command-line-tools/) documentation for information about getting started with the command-line tools. `,

	RunE: func(cmd *cobra.Command, args []string) error {
		geneInts, err := strToInt32ListErr(argIDArgs)
		if err != nil {
			return err
		}
		if len(geneInts) < 1 {
			return errors.New("Input gene-ids not specified")
		}
		return downloadOrthologsByGeneId(geneInts)
	},
}

func init() {
	downloadOrthologGeneIDCmd.PersistentFlags().StringVarP(&argInputFile, "inputfile", "i", "", "read a list of NCBI Gene IDs from a file to use as input")
}
