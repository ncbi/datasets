package datasets

import (
	openapi "datasets/openapi/v2"
	cmdflags "datasets_cli/v2/datasets/flags"
	_nethttp "net/http"

	"github.com/spf13/cobra"
	"github.com/thediveo/enumflag/v2"
)

type SummaryGeneFlag struct {
	jsonLinesLimitFlag *cmdflags.JsonLinesAndLimitFlag
	cmdFlagSet         []cmdflags.FlagInterface
}

func initSummaryGeneFlag() *SummaryGeneFlag {

	jlf := cmdflags.NewJsonLineAndLimitFlag("gene")
	sGeneFlag := &SummaryGeneFlag{
		jsonLinesLimitFlag: jlf,
		cmdFlagSet:         []cmdflags.FlagInterface{jlf},
	}
	return sGeneFlag
}

func NewGeneAccessionRequestIter(request *openapi.V2GeneDatasetReportsRequest) *DefaultPagedRequestIterator[*openapi.V2GeneDatasetReportsRequest, string] {
	accRequester := NewDefaultPagedRequestIterator[*openapi.V2GeneDatasetReportsRequest, string](request,
		func(request *openapi.V2GeneDatasetReportsRequest, accessions []string) {
			request.SetAccessions(accessions)
		},
	)

	if len(request.GetAccessions()) > PAGE_ITER_THRESHOLD {
		accRequester.ids = request.GetAccessions()
		request.SetAccessions([]string{})
	}

	return accRequester
}

func NewGeneLocusTagRequestIter(request *openapi.V2GeneDatasetReportsRequest) *DefaultPagedRequestIterator[*openapi.V2GeneDatasetReportsRequest, string] {
	locusTagRequester := NewDefaultPagedRequestIterator[*openapi.V2GeneDatasetReportsRequest, string](request,
		func(request *openapi.V2GeneDatasetReportsRequest, locusTags []string) {
			request.SetLocusTags(locusTags)
		},
	)

	if len(request.GetLocusTags()) > PAGE_ITER_THRESHOLD {
		locusTagRequester.ids = request.GetLocusTags()
		request.SetLocusTags([]string{})
	}

	return locusTagRequester
}

func NewGeneIdRequestIter(request *openapi.V2GeneDatasetReportsRequest) *DefaultPagedRequestIterator[*openapi.V2GeneDatasetReportsRequest, int32] {
	geneIdRequester := NewDefaultPagedRequestIterator[*openapi.V2GeneDatasetReportsRequest, int32](request,
		func(request *openapi.V2GeneDatasetReportsRequest, geneIds []int32) { request.SetGeneIds(geneIds) },
	)

	if len(request.GetGeneIds()) > PAGE_ITER_THRESHOLD {
		geneIdRequester.ids = request.GetGeneIds()
		request.SetGeneIds([]int32{})
	}

	return geneIdRequester
}

func NewGeneSymbolRequestIter(request *openapi.V2GeneDatasetReportsRequest) *DefaultPagedRequestIterator[*openapi.V2GeneDatasetReportsRequest, string] {
	symbolRequester := NewDefaultPagedRequestIterator[*openapi.V2GeneDatasetReportsRequest, string](request,
		func(request *openapi.V2GeneDatasetReportsRequest, symbols []string) {
			request.SymbolsForTaxon.SetSymbols(symbols)
		},
	)

	if len(request.SymbolsForTaxon.GetSymbols()) > PAGE_ITER_THRESHOLD {
		symbolRequester.ids = request.SymbolsForTaxon.GetSymbols()
		request.SymbolsForTaxon.SetSymbols([]string{})
	}

	return symbolRequester
}

func createCheckGeneIdDuplicates() func(openapi.V2reportsGeneReportMatch, *map[string]bool) bool {
	return func(report openapi.V2reportsGeneReportMatch, duplicatesMap *map[string]bool) bool {
		geneId := ""
		if report.HasGene() {
			geneId = report.Gene.GetGeneId()
		} else if report.HasProduct() {
			geneId = report.Product.GetGeneId()
		}
		if (*duplicatesMap)[geneId] {
			return true
		}
		(*duplicatesMap)[geneId] = true
		return false
	}
}

type GeneDatasetApi struct {
	geneApi *openapi.GeneAPIService
}

func (apiService *GeneDatasetApi) GetPage(request *openapi.V2GeneDatasetReportsRequest) (*openapi.V2reportsGeneDataReportPage, *_nethttp.Response, error) {
	apiRequest := openapi.ApiGeneDatasetReportRequest{
		ApiService: apiService.geneApi,
	}
	return apiRequest.V2GeneDatasetReportsRequest(*request).Execute()
}

func (apiService *GeneDatasetApi) GetPagePtr(page openapi.V2reportsGeneDataReportPage) *openapi.V2reportsGeneDataReportPage {
	return &page
}

type GeneProductApi struct {
	geneApi *openapi.GeneAPIService
}

func (apiService *GeneProductApi) GetPage(request *openapi.V2GeneDatasetReportsRequest) (*openapi.V2reportsGeneDataReportPage, *_nethttp.Response, error) {
	apiRequest := openapi.ApiGeneProductReportRequest{
		ApiService: apiService.geneApi,
	}
	return apiRequest.V2GeneDatasetReportsRequest(*request).Execute()
}

func (apiService *GeneProductApi) GetPagePtr(page openapi.V2reportsGeneDataReportPage) *openapi.V2reportsGeneDataReportPage {
	return &page
}

type GeneOrthologApi struct {
	geneApi *openapi.GeneAPIService
}

func (apiService *GeneOrthologApi) GetPage(request *openapi.V2OrthologRequest) (*openapi.V2reportsGeneDataReportPage, *_nethttp.Response, error) {
	apiRequest := openapi.ApiGeneOrthologsByPostRequest{
		ApiService: apiService.geneApi,
	}
	return apiRequest.V2OrthologRequest(*request).Execute()
}

func (apiService *GeneOrthologApi) GetPagePtr(page openapi.V2reportsGeneDataReportPage) *openapi.V2reportsGeneDataReportPage {
	return &page
}

func getGeneApi(cli *openapi.APIClient) (api PageRetriever[*openapi.V2GeneDatasetReportsRequest, openapi.V2reportsGeneDataReportPage, openapi.V2reportsGeneReportMatch, *openapi.V2reportsGeneDataReportPage]) {
	if argGeneReportMode == Product {
		return &GeneProductApi{geneApi: cli.GeneAPI}
	}
	return &GeneDatasetApi{geneApi: cli.GeneAPI}
}

func geneSummaryPagePrinter(
	sGeneFlag *SummaryGeneFlag,
	requestIter RequestIterator[*openapi.V2GeneDatasetReportsRequest],
	api PageRetriever[*openapi.V2GeneDatasetReportsRequest, openapi.V2reportsGeneDataReportPage, openapi.V2reportsGeneReportMatch, *openapi.V2reportsGeneDataReportPage],
) (err error) {
	requestIter.GetRequest().SetReturnedContent(openapi.V2GeneDatasetReportsRequestContentType(GeneReportModeOpenapi[argGeneReportMode]))
	if sGeneFlag.jsonLinesLimitFlag.CountOnly() {
		requestIter.GetRequest().SetReturnedContent(openapi.V2GENEDATASETREPORTSREQUESTCONTENTTYPE_IDS_ONLY)
	}

	// Look for duplicates in gene paging, since multiple gene ids can map to the same gene (used here in all cases, but not needed for queries
	// that do not do client-side paging)
	pagePrinter := NewPagePrinter[openapi.V2reportsGeneReportMatch, *openapi.V2reportsGeneDataReportPage](
		"gene",
		sGeneFlag.jsonLinesLimitFlag.JsonLines(),
		WithPrintJsonLinesResult[openapi.V2reportsGeneReportMatch, *openapi.V2reportsGeneDataReportPage](createJsonLinesPrintDescriptor(argGeneReportMode)),
		WithCheckDuplicates[openapi.V2reportsGeneReportMatch, *openapi.V2reportsGeneDataReportPage](createCheckGeneIdDuplicates()),
	)

	_, err = ProcessPages[*openapi.V2GeneDatasetReportsRequest,
		openapi.V2reportsGeneDataReportPage,
		openapi.V2reportsGeneReportMatch,
		*openapi.V2reportsGeneDataReportPage](requestIter, api, &pagePrinter, sGeneFlag.jsonLinesLimitFlag.RetrievalCount(), sGeneFlag.jsonLinesLimitFlag.CountOnly())

	return err
}

func createSummaryGeneCmd() *cobra.Command {

	sGeneFlag := initSummaryGeneFlag()

	cmd := &cobra.Command{
		Use:   "gene",
		Short: "Print a summary of a gene dataset",
		Long: `
Print a data report containing gene metadata.  The data report is returned in JSON format.`,
		Example: `  datasets summary gene gene-id 672
  datasets summary gene symbol brca1 --taxon "mus musculus"
  datasets summary gene accession NP_000483.3`,
		Args:              cobra.NoArgs,
		PersistentPreRunE: cmdflags.ExecutePreRunEFor(sGeneFlag.cmdFlagSet),
		RunE:              ParentCommandRunE,
	}

	cmdflags.RegisterAllFlags(sGeneFlag.cmdFlagSet, cmd.PersistentFlags())

	cmd.AddCommand(createSummaryGeneIdCmd(sGeneFlag))
	cmd.AddCommand(createSummaryGeneSymbolCmd(sGeneFlag))
	cmd.AddCommand(createSummaryGeneAccessionCmd(sGeneFlag))
	cmd.AddCommand(createSummaryGeneTaxonCmd(sGeneFlag))
	cmd.AddCommand(createSummaryGeneLocusTagCmd(sGeneFlag))

	pflags := cmd.PersistentFlags()

	pflags.Var(
		enumflag.New(&argGeneReportMode, "string", GeneReportModeIds, enumflag.EnumCaseInsensitive),
		"report",
		`Choose the output type:
  * gene:     Retrieve the primary gene report
  * product:  Retrieve product data report
  * ids_only: Only retrieve gene-ids
    `)

	return cmd
}

type GeneReportMode enumflag.Flag

// Define the enumeration values for ReportMode.
const (
	Data GeneReportMode = iota
	Product
	IdsOnly
)

// Map enumeration values to their textual representations (value
// identifiers).
var GeneReportModeIds = map[GeneReportMode][]string{
	Data:    {"gene", "complete"},
	Product: {"product"},
	IdsOnly: {"ids_only"},
}

// Map enums to their open-api defined string values to be passed to the api
var GeneReportModeOpenapi = map[GeneReportMode]openapi.V2GeneDatasetRequestContentType{
	IdsOnly: openapi.V2GENEDATASETREQUESTCONTENTTYPE_IDS_ONLY,
	Data:    openapi.V2GENEDATASETREQUESTCONTENTTYPE_COMPLETE,
	Product: openapi.V2GENEDATASETREQUESTCONTENTTYPE_COMPLETE,
}

var argGeneReportMode GeneReportMode
