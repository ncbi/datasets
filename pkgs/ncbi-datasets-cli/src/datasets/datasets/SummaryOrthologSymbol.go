package datasets

import (
	"errors"

	"github.com/spf13/cobra"

	openapi "datasets_cli/v1/openapi"
)

func cmdRunSummaryOrthologSymbol(cmd *cobra.Command, args []string) (err error) {
	req := openapi.NewV1GeneDatasetRequest()
	symbols_taxon := openapi.NewV1GeneDatasetRequestSymbolsForTaxon()
	symbols_taxon.SetSymbols(argIDArgs)
	symbols_taxon.SetTaxon(argTaxon)
	req.SetSymbolsForTaxon(*symbols_taxon)

	var geneInts []int32
	geneInts, err = allGeneIdForRequest(req)
	if err != nil {
		return
	}
	if len(geneInts) == 0 {
		err = errors.New("No genes found for search term")
		return
	}
	if !argJsonFormat {
		err = streamOrthologs(geneInts, openapi.V1ORTHOLOGREQUESTCONTENTTYPE_COMPLETE, &JsonLinesStreamProcessor{})
	} else {
		err = streamOrthologs(geneInts, openapi.V1ORTHOLOGREQUESTCONTENTTYPE_COMPLETE, &JsonStreamProcessor{wrapperName: "orthologs"})
	}

	return
}

var summaryOrthologSymbolCmd = &cobra.Command{
	Use:   "symbol <symbol ...>",
	Short: "print a summary of an ortholog dataset by NCBI Gene symbol",
	Long:  `Print a summary of an ortholog dataset by NCBI Gene symbol. Ortholog data is calculated by NCBI for vertebrates and insects. The summary is returned in JSON format.`,
	Example: `  datasets summary ortholog symbol brca1 --taxon 9606
  datasets summary ortholog symbol gapdh --taxon 10090
  datasets summary ortholog symbol gnas --taxon 9606 --taxon-filter 9606,10090`,
	RunE: cmdRunSummaryOrthologSymbol,
}

func init() {
	pflags := summaryOrthologSymbolCmd.PersistentFlags()
	pflags.StringVarP(&argTaxon, "taxon", "t", "human", "Specify a taxonomy ID, species name or NCBI Taxonomy ID")
	registerHiddenStringSlicePair(pflags, &taxonFilter, "taxon-filter", "", []string{}, "limit results to ortholog data for a specified taxonomic group")
	registerHiddenStringPair(summaryOrthologSymbolCmd.Flags(), &argInputFile, "inputfile", "i", "", "read a list of gene symbols from a file to use as input")
}
