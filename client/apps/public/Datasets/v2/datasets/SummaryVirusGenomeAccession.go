package datasets

import (
	_nethttp "net/http"

	"github.com/spf13/cobra"

	openapi "datasets/openapi/v2"
	cmdflags "datasets_cli/v2/datasets/flags"
)

type VirusDatasetApi struct {
	virusApi *openapi.VirusAPIService
}

func (apiService *VirusDatasetApi) GetPage(request *openapi.V2VirusDataReportRequest) (*openapi.V2reportsVirusDataReportPage, *_nethttp.Response, error) {
	apiRequest := openapi.ApiVirusReportsByPostRequest{
		ApiService: apiService.virusApi,
	}
	return apiRequest.V2VirusDataReportRequest(*request).Execute()

}

func (apiService *VirusDatasetApi) GetPagePtr(page openapi.V2reportsVirusDataReportPage) *openapi.V2reportsVirusDataReportPage {
	return &page
}

type VirusAnnotationApi struct {
	virusApi *openapi.VirusAPIService
}

func (apiService *VirusAnnotationApi) GetPage(request *openapi.V2VirusAnnotationReportRequest) (*openapi.V2reportsVirusAnnotationReportPage, *_nethttp.Response, error) {
	apiRequest := openapi.ApiVirusAnnotationReportsByPostRequest{
		ApiService: apiService.virusApi,
	}
	return apiRequest.V2VirusAnnotationReportRequest(*request).Execute()
}

func (apiService *VirusAnnotationApi) GetPagePtr(page openapi.V2reportsVirusAnnotationReportPage) *openapi.V2reportsVirusAnnotationReportPage {
	return &page
}

func NewVirusGenomeAccessionRequestIterator(request *openapi.V2VirusDataReportRequest) *DefaultPagedRequestIterator[*openapi.V2VirusDataReportRequest, string] {
	accRequester := NewDefaultPagedRequestIterator[*openapi.V2VirusDataReportRequest, string](request,
		func(request *openapi.V2VirusDataReportRequest, accessions []string) {
			request.Filter.SetAccessions(accessions)
		},
	)

	if request.Filter.HasAccessions() && len(request.Filter.GetAccessions()) > PAGE_ITER_THRESHOLD {
		accRequester.ids = request.Filter.GetAccessions()
		request.Filter.SetAccessions([]string{})
	}

	return accRequester
}

func NewVirusAnnotAccessionRequestIterator(request *openapi.V2VirusAnnotationReportRequest) *DefaultPagedRequestIterator[*openapi.V2VirusAnnotationReportRequest, string] {
	accRequester := NewDefaultPagedRequestIterator[*openapi.V2VirusAnnotationReportRequest, string](request,
		func(request *openapi.V2VirusAnnotationReportRequest, accessions []string) {
			request.Filter.SetAccessions(accessions)
		},
	)

	if request.Filter.HasAccessions() && len(request.Filter.GetAccessions()) > PAGE_ITER_THRESHOLD {
		accRequester.ids = request.Filter.GetAccessions()
		request.Filter.SetAccessions([]string{})
	}

	return accRequester
}

func executeSummaryVirusGenomeCmd(taxons []string, svf SummaryVirusFlag, iff *cmdflags.InputFileFlag) error {

	var accs []string
	if iff != nil {
		accs = iff.InputIDArgs
	}

	cli, cliErr := createOAClient()
	if cliErr != nil {
		return cliErr
	}

	if cmdflags.ArgsVirusSummaryReportMode == cmdflags.ANNOTATION {
		api := VirusAnnotationApi{virusApi: cli.VirusAPI}
		pagePrinter := NewPagePrinter[openapi.V2reportsVirusAnnotationReport, *openapi.V2reportsVirusAnnotationReportPage](
			"virus",
			svf.jsonLinesLimitFlag.JsonLines(),
		)
		request := svf.PrepareAnnotationReportRequest(accs, taxons)
		_, err := ProcessPages[*openapi.V2VirusAnnotationReportRequest,
			openapi.V2reportsVirusAnnotationReportPage,
			openapi.V2reportsVirusAnnotationReport,
			*openapi.V2reportsVirusAnnotationReportPage](NewVirusAnnotAccessionRequestIterator(request), &api, &pagePrinter, svf.jsonLinesLimitFlag.RetrievalCount(), svf.jsonLinesLimitFlag.CountOnly())
		return err
	} else {
		api := VirusDatasetApi{virusApi: cli.VirusAPI}
		pagePrinter := NewPagePrinter[openapi.V2reportsVirusAssembly, *openapi.V2reportsVirusDataReportPage](
			"virus",
			svf.jsonLinesLimitFlag.JsonLines(),
		)
		request := svf.PrepareDatasetReportRequest(accs, taxons)
		_, err := ProcessPages[*openapi.V2VirusDataReportRequest,
			openapi.V2reportsVirusDataReportPage,
			openapi.V2reportsVirusAssembly,
			*openapi.V2reportsVirusDataReportPage](NewVirusGenomeAccessionRequestIterator(request), &api, &pagePrinter, svf.jsonLinesLimitFlag.RetrievalCount(), svf.jsonLinesLimitFlag.CountOnly())
		return err
	}

}

func initInputFlagSummaryVirusAccession() (iff *cmdflags.InputFileFlag, flagSets []cmdflags.FlagInterface) {
	iff = cmdflags.NewInputFileFlag(cmdflags.InputFileListTypeVirusAcc, cmdflags.AsIntegerFalse)
	flagSets = []cmdflags.FlagInterface{iff}

	return iff, flagSets
}

func createSummaryVirusGenomeAccCmd(vsf SummaryVirusFlag) *cobra.Command {
	iff, flagSets := initInputFlagSummaryVirusAccession()
	cmd := &cobra.Command{
		Use:   "accession",
		Short: "Print a data report containing virus genome metadata by accession",
		Long: `
Print a data report containing virus genome metadata by nucleotide accession. The data report is returned in JSON format.

Refer to NCBI's [download and install](https://www.ncbi.nlm.nih.gov/datasets/docs/v2/download-and-install/) documentation for information about getting started with the command-line tools.`,
		Example: "  datasets summary virus genome accession NC_045512.2",
		PreRunE: cmdflags.ExecutePreRunEFor(flagSets),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			var NoTaxons []string
			err = executeSummaryVirusGenomeCmd(NoTaxons, vsf, iff)

			return err
		},
	}
	cmdflags.RegisterAllFlags(flagSets, cmd.Flags())
	return cmd
}
