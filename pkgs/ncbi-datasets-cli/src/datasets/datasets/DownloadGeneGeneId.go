package datasets

import (
	"github.com/spf13/cobra"
)

var downloadGeneGeneIDCmd = &cobra.Command{
	Use:   "gene-id <gene-id ...>",
	Short: "download a gene dataset by NCBI Gene ID",
	Example: `  datasets download gene gene-id 672
  datasets download gene gene-id 2597 14433`,
	Long: `
Download a gene dataset by NCBI Gene ID. Gene datasets include gene, transcript and protein sequence, a data table and a data report. Datasets are downloaded as a zip file.

The default gene dataset includes the following files:
 * gene.fna (gene sequences)
 * rna.fna (transcript sequences)
 * protein.faa (protein sequences)
 * data_report.jsonl (data report with gene metadata)
 * data_table.tsv (data table with gene metadata, one transcript per row)
 * dataset_catalog.json (a list of files and file types included in the dataset)

Refer to [NCBI Datasets Gene Package](https://www.ncbi.nlm.nih.gov/datasets/docs/data-packages/gene/) documentation for more information about the gene package. `,

	RunE: cmdDownloadGeneGeneID,
}

func init() {
	downloadGeneGeneIDCmd.PersistentFlags().StringVarP(&argInputFile, "inputfile", "i", "", "read a list of NCBI Gene IDs from a file to use as input")
}
