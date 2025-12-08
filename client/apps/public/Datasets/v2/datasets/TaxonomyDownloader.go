package datasets

import (
	openapi "datasets/openapi/v2"
	cmdflags "datasets_cli/v2/datasets/flags"
	_nethttp "net/http"

	"context"
	"fmt"
	"strconv"
)

type TaxonomyBase struct {
	cli *openapi.APIClient
}

type TaxonomyDownloader struct {
	TaxonomyBase
	request *openapi.V2TaxonomyDatasetRequest
}

type taxonRelatedIdsApi struct {
	taxonApi *openapi.TaxonomyAPIService
}

func (apiService *taxonRelatedIdsApi) GetPage(request *openapi.V2TaxonomyRelatedIdRequest) (*openapi.V2reportsTaxonomyDataReportPage, *_nethttp.Response, error) {
	apiRequest := openapi.ApiTaxonomyRelatedIdsPostRequest{
		ApiService: apiService.taxonApi,
	}
	reportPage := openapi.NewV2reportsTaxonomyDataReportPage()
	taxIdsPage, http_resp, err := apiRequest.V2TaxonomyRelatedIdRequest(*request).Execute()
	if err == nil {
		// Switch page result types since the simpler TaxonomyTaxIdsPage does not have the fields required by the ProcessPages api
		taxCount := len(taxIdsPage.GetTaxIds())
		reportPage.SetNextPageToken(taxIdsPage.GetNextPageToken())
		reportPage.SetTotalCount(int32(taxCount))
		// Yes, golang requires a two-step assignment here
		reports := make([]openapi.V2reportsTaxonomyReportMatch, taxCount)
		reportPage.SetReports(reports)
		for i := 0; i < taxCount; i++ {
			reportMatch := openapi.NewV2reportsTaxonomyReportMatch()
			reportMatch.SetTaxonomy(*openapi.NewV2reportsTaxonomyNode())
			reportMatch.Taxonomy.SetTaxId(taxIdsPage.GetTaxIds()[i])
			reportPage.Reports[i] = *reportMatch
		}
	}
	return reportPage, http_resp, err
}

func (apiService *taxonRelatedIdsApi) GetPagePtr(page openapi.V2reportsTaxonomyDataReportPage) *openapi.V2reportsTaxonomyDataReportPage {
	return &page
}

func (td *TaxonomyBase) getFilteredTaxids(taxIds []int32, taxidRetriever *TaxonomyIdRetriever) error {
	request := openapi.NewV2TaxonomyMetadataRequest()
	request.SetTaxons(GetTaxonsFromTaxIds(taxIds))
	request.SetRanks(taxidRetriever.GetRanks())
	// Just gathering taxids here - minimize amount returned for performance
	request.SetReturnedContent(openapi.V2TAXONOMYMETADATAREQUESTCONTENTTYPE_TAXIDS)

	api := TaxonApi{taxonApi: td.cli.TaxonomyAPI}

	_, err := ProcessAllPages[*openapi.V2TaxonomyMetadataRequest,
		openapi.V2reportsTaxonomyDataReportPage,
		openapi.V2reportsTaxonomyReportMatch,
		*openapi.V2reportsTaxonomyDataReportPage](NewDefaultRequestIterator(request), &api, taxidRetriever)

	return err
}

func (td *TaxonomyBase) getChildTaxIds(taxId int32, taxidRetriever *TaxonomyIdRetriever, getParents bool, getChildren bool) error {
	related_ids_request := openapi.NewV2TaxonomyRelatedIdRequest()
	related_ids_request.SetTaxId(taxId)
	related_ids_request.SetIncludeSubtree(getChildren)
	related_ids_request.SetIncludeLineage(getParents)
	related_ids_request.SetRanks(taxidRetriever.GetRanks())

	api := taxonRelatedIdsApi{taxonApi: td.cli.TaxonomyAPI}
	_, err := ProcessAllPages[*openapi.V2TaxonomyRelatedIdRequest,
		openapi.V2reportsTaxonomyDataReportPage,
		openapi.V2reportsTaxonomyReportMatch,
		*openapi.V2reportsTaxonomyDataReportPage](NewDefaultRequestIterator(related_ids_request), &api, taxidRetriever)

	return err
}

func GetTaxonsFromTaxIds(taxIds []int32) []string {
	taxons := make([]string, len(taxIds))
	for idx, taxId := range taxIds {
		taxons[idx] = strconv.Itoa(int(taxId))
	}

	return taxons
}

func (td *TaxonomyBase) retrieveChildAndParentTaxIds(taxIds []int32, ranks []openapi.V2reportsRankType, getParents bool, getChildren bool) (error, []int32) {
	// taxid(s) passed in may not be of requested rank, but if lineage or children are requested, we still get
	// the parents/children and test Those against the rank(s)

	// We only need to filter taxids here if ranks is provided (there are no other filters)
	taxidRetriever := NewTaxonomyIdRetriever()
	if len(ranks) > 0 {
		taxidRetriever.SetRanks(ranks)
		err := td.getFilteredTaxids(taxIds, &taxidRetriever)
		if err != nil {
			return err, taxidRetriever.GetTaxIds()
		}
	} else {
		taxidRetriever.AddTaxIds(taxIds)
	}

	// Get all children. Note the caller should have already checked that taxIds had length==1
	if getChildren || getParents {
		err := td.getChildTaxIds(taxIds[0], &taxidRetriever, getParents, getChildren)
		if err != nil {
			return err, taxidRetriever.GetTaxIds()
		}
	}

	return nil, taxidRetriever.GetTaxIds()
}

func (td *TaxonomyDownloader) setAllTaxidsForRequest(taxIds []int32, downloadFlags DownloadTaxonomyFlag) error {
	td.request.SetAuxReports([]openapi.V2TaxonomyDatasetRequestTaxonomyReportType{openapi.V2TAXONOMYDATASETREQUESTTAXONOMYREPORTTYPE_TAXONOMY_SUMMARY})
	for _, rpt := range downloadFlags.rptFlag.IncludeReports {
		td.request.SetAuxReports(append(td.request.GetAuxReports(), cmdflags.TaxonomyDownloadIncludeFlagOpenapi[rpt]))
	}

	// If the user supplied ranks, the retrieval of children and parents is automatically set to true -
	// otherwise it would just be filtering the command line taxons against the rank which would have little value
	if len(downloadFlags.rankFlag.GetRanks()) > 0 {
		downloadFlags.childrenFlag.SetChildren(true)
		downloadFlags.parentsFlag.SetParents(true)
	}

	err, taxIds := td.retrieveChildAndParentTaxIds(taxIds, downloadFlags.rankFlag.GetRanks(), downloadFlags.parentsFlag.GetParents(), downloadFlags.childrenFlag.GetChildren())
	if err != nil {
		return err
	}
	td.request.SetTaxIds(taxIds)

	return nil
}

func NewTaxonomyDownloader(taxIds []int32, downloadFlags DownloadTaxonomyFlag) (*TaxonomyDownloader, string, error) {
	cli, err := createOAClient()
	if err != nil {
		return nil, "", err
	}
	td := &TaxonomyDownloader{
		TaxonomyBase: TaxonomyBase{
			cli: cli,
		},
		request: openapi.NewV2TaxonomyDatasetRequest(),
	}
	optErr := td.setAllTaxidsForRequest(taxIds, downloadFlags)
	if optErr != nil {
		return nil, "", optErr
	}
	if len(td.request.GetTaxIds()) == 0 {
		return nil, "", fmt.Errorf("No taxons found that match selection\n")
	}

	return td, "", nil
}

func (td *TaxonomyDownloader) Download(argSkipZipVal bool) (err error) {
	_, resp, err := td.cli.TaxonomyAPI.DownloadTaxonomyPackageByPost(context.TODO()).V2TaxonomyDatasetRequest(*td.request).Execute()
	if err = handleHTTPResponse(resp, err); err != nil {
		return
	}
	length := int64(-1) // unknown length
	err = downloadDataForFile(resp, err, argDownloadFilename, length, argSkipZipVal)
	return
}
