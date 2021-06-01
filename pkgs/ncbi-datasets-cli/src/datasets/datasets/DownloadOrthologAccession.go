package datasets

import (
	"github.com/spf13/cobra"

	openapi "main/openapi_client"
)

var downloadOrthologAccessionCmd = &cobra.Command{
	Use:   "accession <refseq-accession>",
	Short: "download an ortholog dataset by RefSeq nucleotide or protein accession",
	Example: `  datasets download ortholog accession NP_000483.3
  datasets download ortholog accession NM_000546.6`,
	Long: `
Download an ortholog dataset by RefSeq nucleotide or protein accession. Ortholog data is calculated by NCBI for vertebrates and insects. Ortholog datasets include gene, transcript and protein sequence, a data table and a data report. Datasets are downloaded as a zip file.

The default ortholog dataset includes the following files:
 * gene.fna (gene sequences)
 * rna.fna (transcript sequences)
 * protein.faa (protein sequences)
 * data_report.jsonl (data report with gene metadata)
 * data_table.tsv (data table with gene metadata, one transcript per row)
 * dataset_catalog.json (a list of files and file types included in the dataset)

Refer to NCBI's [command line quickstart](https://www.ncbi.nlm.nih.gov/datasets/docs/quickstarts/command-line-tools/) documentation for information about getting started with the command-line tools. `,

	RunE: func(cmd *cobra.Command, args []string) error {
		request := new(openapi.V1alpha1GeneDatasetRequest)
		request.Accessions = argIDArgs
		return downloadOrthologByGeneReq(request)
	},
}

func init() {
	registerHiddenStringPair(
		downloadOrthologAccessionCmd.Flags(),
		&argInputFile,
		"inputfile",
		"i",
		"",
		"read a list of RefSeq nucleotide or protein accessions from a file to use as input",
	)
}
