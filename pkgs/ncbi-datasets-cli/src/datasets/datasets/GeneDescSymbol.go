package datasets

import (
	"github.com/spf13/cobra"
)

var geneDescrSymbolCmd = &cobra.Command{
	Deprecated: "please use \"summary gene symbol\"",
	Use:        "symbol <symbol ...>",
	Short:      "Gene descriptions by gene symbol",
	Example: `  datasets gene-descriptors symbol tp53 --taxon dog
  datasets gene-descriptors symbol brca1 brca2 --taxon 10090`,
	Long: `
Get gene descriptors (metadata) by gene symbol. Specify a gene symbol and species name or species-level Taxonomy ID.
If no species is provided, gene descriptors will be provided for human genes.
`,
	RunE: cmdRunSummaryGeneSymbol,
}

func init() {
	geneDescrSymbolCmd.PersistentFlags().StringVarP(&argTaxon, "taxon", "t", "human", "Specify a taxonomy ID, species name or NCBI Taxonomy ID")
	registerHiddenStringPair(geneDescrSymbolCmd.Flags(), &argInputFile, "inputfile", "i", "", "read a list of gene symbols from a file to use as input")
}
