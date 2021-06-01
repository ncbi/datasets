package datasets

import (
	"errors"

	"github.com/spf13/cobra"

	openapi "main/openapi_client"
)

var downloadGeneSymbolCmd = &cobra.Command{
	Use:   "symbol <gene_symbol ...>",
	Short: "download a gene dataset by gene symbol",
	Example: `  datasets download gene symbol tp53
  datasets download gene symbol brca1 --taxon mouse`,
	Long: `
Download a gene dataset by gene symbol and taxon (species name or species-level NCBI Taxonomy ID). If no taxon is specified, data will be returned for human. Gene datasets include gene, transcript and protein sequence, a data table and a data report. Datasets are downloaded as a zip file.

The default gene dataset includes the following files:
 * gene.fna (gene sequences)
 * rna.fna (transcript sequences)
 * protein.faa (protein sequences)
 * data_report.jsonl (data report with gene metadata)
 * data_table.tsv (data table with gene metadata, one transcript per row)
 * dataset_catalog.json (a list of files and file types included in the dataset)

Refer to [NCBI Datasets Gene Package](https://www.ncbi.nlm.nih.gov/datasets/docs/data-packages/gene/) documentation for more information about the gene package. `,

	RunE: func(cmd *cobra.Command, args []string) error {
		if len(argIDArgs) == 0 {
			return errors.New("Input symbols not specified")
		}
		req := new(openapi.V1alpha1GeneDatasetRequest)
		req.SymbolsForTaxon.Symbols = argIDArgs
		req.SymbolsForTaxon.Taxon = argTaxon
		return downloadGeneForRequest(req)
	},
}

func init() {
	flags := downloadGeneSymbolCmd.PersistentFlags()
	registerHiddenStringPair(flags, &argTaxon, "taxon", "t", "human", "specify a species name (common or scientific) or species-level NCBI Taxonomy ID")
	registerHiddenStringPair(flags, &argInputFile, "inputfile", "i", "", "read a list of gene symbols from a file to use as input")
}
