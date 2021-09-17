package datasets

import (
	"github.com/spf13/cobra"

	openapi "datasets_cli/v1/openapi"
)

var downloadOrthologSymbolCmd = &cobra.Command{
	Use:   "symbol <gene_symbol>",
	Short: "download an ortholog dataset by gene symbol",
	Example: `  datasets download ortholog symbol tp53
  datasets download ortholog symbol brca1 --taxon mouse`,
	Long: `
Download an ortholog dataset by gene symbol and taxon (species name or species-level NCBI Taxonomy ID). If no taxon is specified, data will be returned for human. Ortholog data is calculated by NCBI for vertebrates and insects. Ortholog datasets include gene, transcript and protein sequence, a data table and a data report. Datasets are downloaded as a zip file.

The default ortholog dataset includes the following files:
 * gene.fna (gene sequences)
 * rna.fna (transcript sequences)
 * protein.faa (protein sequences)
 * data_report.jsonl (data report with gene metadata)
 * data_table.tsv (data table with gene metadata, one transcript per row)
 * dataset_catalog.json (a list of files and file types included in the dataset)

Refer to NCBI's [command line quickstart](https://www.ncbi.nlm.nih.gov/datasets/docs/quickstarts/command-line-tools/) documentation for information about getting started with the command-line tools. `,

	RunE: func(cmd *cobra.Command, args []string) error {
		req := openapi.NewV1GeneDatasetRequest()
		symbols_taxon := openapi.NewV1GeneDatasetRequestSymbolsForTaxon()
		symbols_taxon.SetSymbols(argIDArgs)
		symbols_taxon.SetTaxon(argTaxon)
		req.SetSymbolsForTaxon(*symbols_taxon)
		return downloadOrthologByGeneReq(req)
	},
}

func init() {
	flags := downloadOrthologSymbolCmd.PersistentFlags()
	registerHiddenStringPair(flags, &argTaxon, "taxon", "t", "human", "specify a species name (common or scientific) or species-level NCBI Taxonomy ID")
	registerHiddenStringPair(flags, &argInputFile, "inputfile", "i", "", "read a list of gene symbols from a file to use as input")
}
