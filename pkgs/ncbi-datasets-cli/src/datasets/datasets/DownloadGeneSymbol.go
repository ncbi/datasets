package datasets

import (
	"errors"

	"github.com/spf13/cobra"

	openapi "datasets_cli/v1/openapi"
)

var downloadGeneSymbolCmd = &cobra.Command{
	Use:   "symbol <gene_symbol ...>",
	Short: "download a gene dataset by gene symbol",
	Example: `  datasets download gene symbol tp53
  datasets download gene symbol brca1 --taxon mouse`,
	Long: `
Download a gene dataset by gene symbol and taxon (species name or species-level NCBI Taxonomy ID). If no taxon is specified, data will be returned for human. Gene data packages include gene, transcript and protein sequence, a data table and a data report. Datasets are downloaded as a zip file.

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
		req := openapi.NewV1GeneDatasetRequest()
		symbols_taxon := openapi.NewV1GeneDatasetRequestSymbolsForTaxon()
		symbols_taxon.SetSymbols(argIDArgs)
		symbols_taxon.SetTaxon(argTaxon)
		req.SetSymbolsForTaxon(*symbols_taxon)
		return downloadGeneForRequest(req)
	},
}

func init() {
	flags := downloadGeneSymbolCmd.PersistentFlags()
	registerHiddenStringPair(flags, &argTaxon, "taxon", "t", "human", "specify a species name (common or scientific) or species-level NCBI Taxonomy ID")
	registerHiddenStringPair(flags, &argInputFile, "inputfile", "i", "", "read a list of gene symbols from a file to use as input")
}
