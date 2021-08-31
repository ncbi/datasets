package datasets

import (
	"github.com/spf13/cobra"

	openapi "main/openapi_client"
)

var downloadGeneTaxonCmd = &cobra.Command{
	Use:     "taxon <taxon>",
	Short:   "download a gene dataset by taxon",
	Example: `  datasets download gene taxon human --exclude-gene --exclude-protein --exclude-rna`,
	Long: `
Download a gene dataset by taxon. Gene datasets include gene, transcript and protein sequence, a data table and a data report. Datasets are downloaded as a zip file.

The default gene dataset includes the following files:
 * gene.fna (gene sequences)
 * rna.fna (transcript sequences)
 * protein.faa (protein sequences)
 * data_report.jsonl (data report with gene metadata)
 * data_table.tsv (data table with gene metadata, one transcript per row)
 * dataset_catalog.json (a list of files and file types included in the dataset)

Refer to [NCBI Datasets Gene Package](https://www.ncbi.nlm.nih.gov/datasets/docs/data-packages/gene/) documentation for more information about the gene package. `,
	Args: cobra.MaximumNArgs(1),
	RunE: cmdDownloadGeneTaxon,
}

func cmdDownloadGeneTaxon(cmd *cobra.Command, args []string) error {
	req := openapi.NewV1GeneDatasetRequest()
	req.SetTaxon(argIDArgs[0])
	return downloadGeneForRequest(req)
}
