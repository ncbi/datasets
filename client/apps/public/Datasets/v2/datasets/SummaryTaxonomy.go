package datasets

import (
	openapi "datasets/openapi/v2"
	_nethttp "net/http"

	"github.com/spf13/cobra"
	"strconv"
)

type TaxonApi struct {
	taxonApi *openapi.TaxonomyAPIService
}

func (apiService *TaxonApi) GetPage(request *openapi.V2TaxonomyMetadataRequest) (*openapi.V2reportsTaxonomyDataReportPage, *_nethttp.Response, error) {
	apiRequest := openapi.ApiTaxonomyDataReportPostRequest{
		ApiService: apiService.taxonApi,
	}
	return apiRequest.V2TaxonomyMetadataRequest(*request).Execute()
}

func (apiService *TaxonApi) GetPagePtr(page openapi.V2reportsTaxonomyDataReportPage) *openapi.V2reportsTaxonomyDataReportPage {
	return &page
}

type taxonNamesApi struct {
	taxonApi *openapi.TaxonomyAPIService
}

func (apiService *taxonNamesApi) GetPage(request *openapi.V2TaxonomyMetadataRequest) (*openapi.V2reportsTaxonomyNamesDataReportPage, *_nethttp.Response, error) {
	apiRequest := openapi.ApiTaxonomyNamesPostRequest{
		ApiService: apiService.taxonApi,
	}
	return apiRequest.V2TaxonomyMetadataRequest(*request).Execute()
}

func (apiService *taxonNamesApi) GetPagePtr(page openapi.V2reportsTaxonomyNamesDataReportPage) *openapi.V2reportsTaxonomyNamesDataReportPage {
	return &page
}

func getTaxonApi() (api *TaxonApi, err error) {
	cli, cliErr := createOAClient()
	if cliErr != nil {
		return nil, cliErr
	}
	api = &TaxonApi{taxonApi: cli.TaxonomyAPI}
	return api, nil
}

func getTaxonNamesApi() (api *taxonNamesApi, err error) {
	cli, cliErr := createOAClient()
	if cliErr != nil {
		return nil, cliErr
	}
	api = &taxonNamesApi{taxonApi: cli.TaxonomyAPI}
	return api, nil
}

// It should be possible to remove some of the duplicate code below through generics or interfaces,
// but (at least for generics), there are several stumbling blocks (e.g. needing both ptr and non-ptr types)
func createCheckTaxonomyNamesDuplicates() func(openapi.V2reportsTaxonomyNamesReportMatch, *map[string]bool) bool {
	return func(report openapi.V2reportsTaxonomyNamesReportMatch, duplicatesMap *map[string]bool) bool {
		if report.HasTaxonomy() {
			taxId := report.GetTaxonomy().TaxId
			if (*duplicatesMap)[*taxId] {
				return true
			}
			(*duplicatesMap)[*taxId] = true
		}
		return false
	}
}

func createCheckTaxonomyDuplicates() func(openapi.V2reportsTaxonomyReportMatch, *map[string]bool) bool {
	return func(report openapi.V2reportsTaxonomyReportMatch, duplicatesMap *map[string]bool) bool {
		if report.HasTaxonomy() {
			taxId := strconv.Itoa(int((report.Taxonomy).GetTaxId()))
			if (*duplicatesMap)[taxId] {
				return true
			}
			(*duplicatesMap)[taxId] = true
		}
		return false
	}
}

func updateTaxonNamesReportQuery(report *openapi.V2reportsTaxonomyNamesReportMatch, taxIdsMap map[string][]string) {
	if report.HasTaxonomy() {
		taxId := report.GetTaxonomy().TaxId
		taxons := make([]string, len(taxIdsMap[*taxId]))
		copy(taxons, taxIdsMap[*taxId])
		report.SetQuery(taxons)
	}
}

func updateTaxonReportQuery(report *openapi.V2reportsTaxonomyReportMatch, taxIdsMap map[string][]string) {
	if report.HasTaxonomy() {
		// If there is only 1 taxid in the map, all results will be assigned the same taxon query
		// This will capture the cases where lineage or children are returned
		if len(taxIdsMap) == 1 {
			for _, query := range taxIdsMap {
				taxons := make([]string, len(query))
				copy(taxons, query)
				report.SetQuery(taxons)
				return
			}
		}
		taxId := strconv.Itoa(int((report.Taxonomy).GetTaxId()))
		taxons := make([]string, len(taxIdsMap[taxId]))
		copy(taxons, taxIdsMap[taxId])
		report.SetQuery(taxons)
	}
}

func createTaxonomyNamesJsonLinesPrint(taxIdsMap map[string][]string) func(openapi.V2reportsTaxonomyNamesReportMatch) {
	return func(report openapi.V2reportsTaxonomyNamesReportMatch) {
		updateTaxonNamesReportQuery(&report, taxIdsMap)
		printResults(report)
	}
}

func createTaxonomyNamesJsonPrint(taxIdsMap map[string][]string) func(openapi.V2reportsTaxonomyNamesReportMatch) {
	return func(report openapi.V2reportsTaxonomyNamesReportMatch) {
		updateTaxonNamesReportQuery(&report, taxIdsMap)
		printResultsNoNewline(report)
	}
}

func createTaxonomyJsonLinesPrint(taxIdsMap map[string][]string) func(openapi.V2reportsTaxonomyReportMatch) {
	return func(report openapi.V2reportsTaxonomyReportMatch) {
		updateTaxonReportQuery(&report, taxIdsMap)
		printResults(report)
	}
}

func createTaxonomyJsonPrint(taxIdsMap map[string][]string) func(openapi.V2reportsTaxonomyReportMatch) {
	return func(report openapi.V2reportsTaxonomyReportMatch) {
		updateTaxonReportQuery(&report, taxIdsMap)
		printResultsNoNewline(report)
	}
}

func getTaxonomyReport(requestIter RequestIterator[*openapi.V2TaxonomyMetadataRequest], taxIdsMap map[string][]string, flags SummaryTaxonomyFlag) (err error) {
	api, apiErr := getTaxonApi()
	if apiErr != nil {
		return apiErr
	}

	pagePrinter := NewPagePrinter[openapi.V2reportsTaxonomyReportMatch, *openapi.V2reportsTaxonomyDataReportPage](
		"taxonomy",
		flags.jsonLinesLimitFlag.JsonLines(),
		WithPrintJsonLinesResult[openapi.V2reportsTaxonomyReportMatch, *openapi.V2reportsTaxonomyDataReportPage](createTaxonomyJsonLinesPrint(taxIdsMap)),
		WithPrintJsonResult[openapi.V2reportsTaxonomyReportMatch, *openapi.V2reportsTaxonomyDataReportPage](createTaxonomyJsonPrint(taxIdsMap)),
		WithCheckDuplicates[openapi.V2reportsTaxonomyReportMatch, *openapi.V2reportsTaxonomyDataReportPage](createCheckTaxonomyDuplicates()),
	)

	_, err = ProcessPages[*openapi.V2TaxonomyMetadataRequest,
		openapi.V2reportsTaxonomyDataReportPage,
		openapi.V2reportsTaxonomyReportMatch,
		*openapi.V2reportsTaxonomyDataReportPage](requestIter, api, &pagePrinter, flags.jsonLinesLimitFlag.RetrievalCount(), flags.jsonLinesLimitFlag.CountOnly())

	return
}

func getTaxonomyNamesReport(requestIter RequestIterator[*openapi.V2TaxonomyMetadataRequest], taxIdsMap map[string][]string, flags SummaryTaxonomyFlag) (err error) {
	api, apiErr := getTaxonNamesApi()
	if apiErr != nil {
		return apiErr
	}

	pagePrinter := NewPagePrinter[openapi.V2reportsTaxonomyNamesReportMatch, *openapi.V2reportsTaxonomyNamesDataReportPage](
		"taxonomy names",
		flags.jsonLinesLimitFlag.JsonLines(),
		WithPrintJsonLinesResult[openapi.V2reportsTaxonomyNamesReportMatch, *openapi.V2reportsTaxonomyNamesDataReportPage](createTaxonomyNamesJsonLinesPrint(taxIdsMap)),
		WithPrintJsonResult[openapi.V2reportsTaxonomyNamesReportMatch, *openapi.V2reportsTaxonomyNamesDataReportPage](createTaxonomyNamesJsonPrint(taxIdsMap)),
		WithCheckDuplicates[openapi.V2reportsTaxonomyNamesReportMatch, *openapi.V2reportsTaxonomyNamesDataReportPage](createCheckTaxonomyNamesDuplicates()),
	)

	_, err = ProcessPages[*openapi.V2TaxonomyMetadataRequest,
		openapi.V2reportsTaxonomyNamesDataReportPage,
		openapi.V2reportsTaxonomyNamesReportMatch,
		*openapi.V2reportsTaxonomyNamesDataReportPage](requestIter, api, &pagePrinter, flags.jsonLinesLimitFlag.RetrievalCount(), flags.jsonLinesLimitFlag.CountOnly())
	return
}

type TaxonomySummary struct {
	TaxonomyBase
	request *openapi.V2TaxonomyMetadataRequest
}

func (ts *TaxonomySummary) setAllTaxidsForRequest(taxIds []int32, flags SummaryTaxonomyFlag) error {
	if flags.rptFlag.IdsOnly() {
		ts.request.SetReturnedContent(openapi.V2TAXONOMYMETADATAREQUESTCONTENTTYPE_TAXIDS)
	} else {
		ts.request.SetReturnedContent(openapi.V2TAXONOMYMETADATAREQUESTCONTENTTYPE_COMPLETE)
	}

	// If the user supplied ranks, the retrieval of children and parents is automatically set to true -
	// otherwise it would just be filtering the command line taxons against the rank which would have little value
	if len(flags.rankFlag.GetRanks()) > 0 {
		flags.childrenFlag.SetChildren(true)
		flags.parentsFlag.SetParents(true)
	}
	ts.request.SetRanks(flags.rankFlag.GetRanks())
	ts.request.SetPageSize(1000)

	err, taxIds := ts.retrieveChildAndParentTaxIds(taxIds, flags.rankFlag.GetRanks(), flags.parentsFlag.GetParents(), flags.childrenFlag.GetChildren())
	if err != nil {
		return err
	}
	ts.request.SetTaxons(GetTaxonsFromTaxIds(taxIds))

	return nil
}

func NewTaxonomyTaxonRequestIter(taxons []string, flags SummaryTaxonomyFlag) (*DefaultPagedRequestIterator[*openapi.V2TaxonomyMetadataRequest, string], error) {
	cli, err := createOAClient()
	if err != nil {
		return nil, err
	}
	ts := &TaxonomySummary{
		TaxonomyBase: TaxonomyBase{
			cli: cli,
		},
		request: openapi.NewV2TaxonomyMetadataRequest(),
	}
	if err := ts.setAllTaxidsForRequest(strToInt32List(taxons), flags); err != nil {
		return nil, err
	}

	taxonRequester := NewDefaultPagedRequestIterator[*openapi.V2TaxonomyMetadataRequest, string](ts.request,
		func(request *openapi.V2TaxonomyMetadataRequest, taxons []string) {
			request.SetTaxons(taxons)
		},
	)

	// Since taxonomy doesn't support total count, we set all queries as id queries which prevents us from
	// doing an initial query to get total count
	taxonRequester.ids = ts.request.GetTaxons()
	ts.request.SetTaxons([]string{})

	return taxonRequester, nil
}

func getTaxonomySummary(taxons []string, flags SummaryTaxonomyFlag, taxIdsMap map[string][]string) (err error) {
	requestIter, err := NewTaxonomyTaxonRequestIter(taxons, flags)
	if err != nil {
		return err
	}

	if flags.rptFlag.IdsOnly() || flags.rptFlag.TaxonomyReport() {
		return getTaxonomyReport(requestIter, taxIdsMap, flags)
	}
	return getTaxonomyNamesReport(requestIter, taxIdsMap, flags)
}

func createSummaryTaxonomyCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "taxonomy",
		Short: "Print a data report containing taxonomy metadata",
		Long: `
Print a data report containing taxonomy metadata. The data report is returned in JSON format.`,
		Example: `  datasets summary taxonomy taxon human
  datasets summary taxonomy taxon human "mus musculus"
  datasets summary taxonomy taxon human --report names
  datasets summary taxonomy taxon "mus musculus" --report ids_only`,
		Args: cobra.NoArgs,
		RunE: ParentCommandRunE,
	}

	cmd.AddCommand(createSummaryTaxonomyTaxonCmd())

	return cmd
}
