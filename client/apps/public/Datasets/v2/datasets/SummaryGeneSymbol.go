package datasets

import (
	openapi "datasets/openapi/v2"
	cmdflags "datasets_cli/v2/datasets/flags"

	"github.com/spf13/cobra"
)

func addSymbolAndTaxonTo(request *openapi.V2GeneDatasetReportsRequest, symbols []string, taxon string) {
	symbols_taxon := openapi.NewV2GeneDatasetReportsRequestSymbolsForTaxon()
	symbols_taxon.SetSymbols(symbols)
	symbols_taxon.SetTaxon(taxon)
	request.SetSymbolsForTaxon(*symbols_taxon)
}

func RetrieveGeneIdsForSymbolsAndTaxon(cli *openapi.APIClient, symbols []string, taxon string) (geneInts []int32, err error) {

	api := GeneDatasetApi{geneApi: cli.GeneAPI}
	geneIdRetriever := NewGeneIdRetriever()

	request := openapi.NewV2GeneDatasetReportsRequest()
	addSymbolAndTaxonTo(request, symbols, taxon)
	request.SetReturnedContent(openapi.V2GENEDATASETREPORTSREQUESTCONTENTTYPE_IDS_ONLY)
	_, err = ProcessAllPages[
		*openapi.V2GeneDatasetReportsRequest,
		openapi.V2reportsGeneDataReportPage,
		openapi.V2reportsGeneReportMatch,
		*openapi.V2reportsGeneDataReportPage](NewGeneSymbolRequestIter(request), &api, &geneIdRetriever)
	if err != nil {
		return nil, err
	}

	return geneIdRetriever.GetGeneIds(), nil
}

func GeneDatasetReportRequestForSymbolAndTaxon(cli *openapi.APIClient, iff *cmdflags.InputFileFlag, otf *cmdflags.OrthologTaxonFilterFlag, taxId string) (*openapi.V2GeneDatasetReportsRequest, error) {
	request := openapi.NewV2GeneDatasetReportsRequest()

	geneSymbols := iff.AsStringList()
	if !otf.IsOrthologRequested() {
		addSymbolAndTaxonTo(request, geneSymbols, taxId)
		return request, nil
	}
	// Compute the gene-id list via ortholog
	// For each symbol, -> gene_id
	// then for each gene_id -> ortholog gene_ids
	inputGeneInts, errSymbol := RetrieveGeneIdsForSymbolsAndTaxon(cli, geneSymbols, taxId)

	if errSymbol != nil {
		return nil, errSymbol
	}

	geneInts, err := RetrieveOrthologGeneIdsFor(cli, otf, inputGeneInts)
	if err != nil {
		return nil, err
	}
	if len(geneInts) == 0 {
		return nil, nil
	}

	request.SetGeneIds(geneInts)
	return request, nil
}

func createSummaryGeneSymbolCmd(sGeneFlag *SummaryGeneFlag) *cobra.Command {
	iff := cmdflags.NewInputFileFlag(cmdflags.InputFileListTypeGeneSymbol, cmdflags.AsIntegerFalse)
	otf := cmdflags.NewOrthologTaxonFilterFlag()
	gtff := cmdflags.NewGeneTaxonFilterFlag()
	flagSets := []cmdflags.FlagInterface{iff, otf, gtff}

	cmd := &cobra.Command{
		Use:   "symbol <gene symbol ...> [flags]",
		Short: "Print a data report containing gene metadata by gene symbol",
		Long: `
Print a data report containing gene metadata by gene symbol and taxon (NCBI Taxonomy ID, scientific or common name for a species). If no taxon is specified, data will be returned for human. The data report is returned in JSON format.`,
		Example: `  datasets summary gene symbol tp53
  datasets summary gene symbol brca1 --taxon "mus musculus"`,
		PreRunE: cmdflags.ExecutePreRunEFor(flagSets),
		RunE: func(cmd *cobra.Command, args []string) error {
			cli, cliErr := createOAClient()
			if cliErr != nil {
				return cliErr
			}

			taxId, taxError := RetrieveTaxIdForTaxon(
				gtff.Taxon,
				false,
				openapi.V2ORGANISMQUERYREQUESTTAXONRESOURCEFILTER_GENE,
				"gene",
			)
			if taxError != nil {
				return taxError
			}
			request, err := GeneDatasetReportRequestForSymbolAndTaxon(cli, iff, otf, taxId)

			if err != nil {
				return err
			}
			if request == nil {
				return nil
			}

			return geneSummaryPagePrinter(sGeneFlag, NewGeneSymbolRequestIter(request), getGeneApi(cli))
		},
	}

	cmdflags.RegisterAllFlags(flagSets, cmd.PersistentFlags())
	return cmd
}
