package datasets

import (
	openapi "datasets/openapi/v2"
	_nethttp "net/http"
	"sort"

	cmdflags "datasets_cli/v2/datasets/flags"

	"github.com/spf13/cobra"
)

// Move these two types to GeneServices.go ?
type GeneApi struct {
	geneApi *openapi.GeneAPIService
}

func (apiService *GeneApi) GetPage(request *openapi.V2GeneDatasetReportsRequest) (*openapi.V2reportsGeneDataReportPage, *_nethttp.Response, error) {
	apiRequest := openapi.ApiGeneMetadataByPostRequest{
		ApiService: apiService.geneApi,
	}
	return apiRequest.V2GeneDatasetReportsRequest(*request).Execute()
}

func (apiService *GeneApi) GetPagePtr(page openapi.V2reportsGeneDataReportPage) *openapi.V2reportsGeneDataReportPage {
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

func createJsonLinesPrintDescriptor(reportMode GeneReportMode) func(openapi.V2reportsGeneReportMatch) {
	return func(report openapi.V2reportsGeneReportMatch) {
		if reportMode == Product {
			printResults(report.GetProduct())
			return
		}
		printResults(report.GetGene())
	}
}

func GeneIdsAsIntsForInputs(cli *openapi.APIClient, iff *cmdflags.InputFileFlag, otf *cmdflags.OrthologTaxonFilterFlag) (geneInts []int32, err error) {

	if !otf.IsOrthologRequested() {
		return iff.AsInt32List, nil
	}
	return RetrieveOrthologGeneIdsFor(cli, otf, iff.AsInt32List)
}

func RetrieveOrthologGeneIdsFor(cli *openapi.APIClient, otf *cmdflags.OrthologTaxonFilterFlag, inputGeneInts []int32) (geneInts []int32, err error) {

	geneOrthologApi := GeneOrthologApi{geneApi: cli.GeneAPI}

	geneIdRetriever := NewGeneIdRetriever()

	for _, geneId := range inputGeneInts {
		request := openapi.NewV2OrthologRequest()
		request.SetGeneId(geneId)
		request.SetReturnedContent(openapi.V2ORTHOLOGREQUESTCONTENTTYPE_IDS_ONLY)
		if !otf.RequestAllOrthologs() {
			request.SetTaxonFilter(otf.OrthologTaxonValue())
		}

		_, err = ProcessAllPagesRequest[
			*openapi.V2OrthologRequest,
			openapi.V2reportsGeneDataReportPage,
			openapi.V2reportsGeneReportMatch,
			*openapi.V2reportsGeneDataReportPage](request, &geneOrthologApi, &geneIdRetriever)
		if err != nil {
			return nil, err
		}
	}
	return geneIdRetriever.GetGeneIds(), nil
}

func createSummaryGeneIdCmd(sGeneFlag *SummaryGeneFlag) *cobra.Command {
	iff := cmdflags.NewInputFileFlag(cmdflags.InputFileListTypeGeneId, cmdflags.AsIntegerTrue)
	otf := cmdflags.NewOrthologTaxonFilterFlag()
	flagSets := []cmdflags.FlagInterface{iff, otf}

	cmd := &cobra.Command{
		Use:   "gene-id <gene-id ...> [flags]",
		Short: "Print a data report containing gene metadata by NCBI Gene ID",
		Long: `
Print a data report containing gene metadata by NCBI Gene ID. The data report is returned in JSON format.`,
		Example: `  datasets summary gene gene-id 672
  datasets summary gene gene-id 2597 14433`,
		PreRunE: cmdflags.ExecutePreRunEFor(flagSets),
		RunE: func(cmd *cobra.Command, args []string) error {

			// This will go in a future PreRunE someday ...
			cli, cliErr := createOAClient()
			if cliErr != nil {
				return cliErr
			}

			geneInts, err := GeneIdsAsIntsForInputs(cli, iff, otf)
			if err != nil {
				return err
			}
			if len(geneInts) == 0 {
				return nil
			}
			sort.Slice(geneInts, func(i, j int) bool { return geneInts[i] < geneInts[j] })

			geneApi := GeneApi{geneApi: cli.GeneAPI}

			request := openapi.NewV2GeneDatasetReportsRequest()
			request.SetGeneIds(geneInts)
			return geneSummaryPagePrinter(sGeneFlag, NewGeneIdRequestIter(request), geneApi)
		},
	}

	cmdflags.RegisterAllFlags(flagSets, cmd.PersistentFlags())

	return cmd
}
