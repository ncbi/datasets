package datasets

import (
	"context"
	"errors"
	"fmt"
	_nethttp "net/http"
	"strings"

	cmdflags "datasets_cli/v2/datasets/flags"

	openapi "datasets/openapi/v2"
)

type VirusDownloader struct {
	request *openapi.V2VirusDatasetRequest
	cli     *openapi.APIClient
}

const VIRAL_TAXID = "10239"
const VIRAL_TAXID_COMP = int32(10239)
const SARS2_TAXID = "2697049"

func (vd *VirusDownloader) Download(argSkipZipVal bool) (err error) {

	_, resp, err := vd.cli.VirusAPI.VirusGenomeDownloadPost(context.TODO()).V2VirusDatasetRequest(*vd.request).Execute()
	if err != nil {
		return err
	}
	if err = handleHTTPResponse(resp, err); err != nil {
		return err
	}
	length := int64(-1) // unknown length
	err = downloadDataForFile(resp, err, argDownloadFilename, length, argSkipZipVal)
	return

}

func (vd *VirusDownloader) checkTaxonWithinScope(taxons []string) bool {
	contains := func(s []int32, e int32) bool {
		for _, a := range s {
			if a == e {
				return true
			}
		}
		return false
	}

	// Shortcut. Also needed in clouldbuild, which does not load taxonomy service
	if len(taxons) == 1 && cmdflags.Contains(KNOWN_VIRUS_TAXONS, strings.TrimSpace(taxons[0])) {
		return true
	}
	org, resp, err := vd.cli.TaxonomyAPI.TaxonomyMetadataPost(context.TODO()).V2TaxonomyMetadataRequest(
		openapi.V2TaxonomyMetadataRequest{
			Taxons: taxons,
		},
	).Execute()
	if err = handleHTTPResponse(resp, err); err == nil {
		taxNodes := org.GetTaxonomyNodes()
		if len(taxNodes) < 1 {
			return false
		}
		for _, taxNode := range taxNodes {
			taxonomy := taxNode.GetTaxonomy()
			lineage := taxonomy.GetLineage()
			if !contains(lineage, VIRAL_TAXID_COMP) || taxonomy.GetTaxId() == VIRAL_TAXID_COMP {
				return false
			}
		}
	}
	return true
}

func (vd *VirusDownloader) getVirusAvailability(argIDs []string) (*openapi.V2VirusAvailability, *_nethttp.Response, error) {
	virusAvailabilityRequest := openapi.NewV2VirusAvailabilityRequest()
	virusAvailabilityRequest.SetAccessions(argIDs)
	return vd.cli.VirusAPI.VirusAccessionAvailabilityPost(context.TODO()).V2VirusAvailabilityRequest(*virusAvailabilityRequest).Execute()
}

type VirusDownloaderIdOption func(vd *VirusDownloader) (string, error)

func VirusDownloadWithAccession(accessions []string, dvf DownloadVirusFlag) VirusDownloaderIdOption {
	return func(vd *VirusDownloader) (string, error) {
		var warning = ""
		virusAvailability, _, err := vd.getVirusAvailability(accessions)
		if err != nil {
			warning = "There was a problem validating your requested accessions"
			return warning, err
		}
		request, invAccessionWarning, err := GetVirusAccessionRequest(*virusAvailability)
		if invAccessionWarning != "" {
			warning += invAccessionWarning
		}
		if err != nil {
			return "", err
		}
		dvf.addFiltersAndOptionsTo(request)
		vd.request = request
		return warning, err
	}
}

func VirusDownloadWithTaxon(taxons []string, dvf DownloadVirusFlag) VirusDownloaderIdOption {
	return func(vd *VirusDownloader) (string, error) {
		if !vd.checkTaxonWithinScope(taxons) {
			var warning = "The download virus genome taxon command only supports virus taxa.\nFor data on other organisms, please use the download genome taxon command.\n"
			return warning, fmt.Errorf("taxa %s is out of scope", taxons)
		}
		request := GetVirusTaxonRequest(taxons)
		/// Imagine we send in the vff here, and it produces the filter request for us.
		dvf.addFiltersAndOptionsTo(request)
		vd.request = request
		return "", nil
	}
}

func GetVirusAccessionRequest(virusAvailability openapi.V2VirusAvailability) (request *openapi.V2VirusDatasetRequest, warning string, err error) {
	if len(virusAvailability.GetInvalidAccessions()) > 0 {
		warning += virusAvailability.GetMessage()
	}
	validAccessions := virusAvailability.GetValidAccessions()
	if len(validAccessions) == 0 {
		err := errors.New("No valid accessions were passed")
		return nil, "", err
	}
	request = openapi.NewV2VirusDatasetRequest()
	request.SetAccessions(validAccessions)

	return request, warning, err
}

func GetVirusTaxonRequest(taxons []string) *openapi.V2VirusDatasetRequest {

	request := openapi.NewV2VirusDatasetRequest()
	request.SetTaxons(taxons)

	return request
}

func NewVirusDownloader(setIdOption VirusDownloaderIdOption) (vd *VirusDownloader, warning string, err error) {
	cli, err := createOAClient()
	if err != nil {
		return nil, "", err
	}
	vd = &VirusDownloader{
		request: openapi.NewV2VirusDatasetRequest(),
		cli:     cli,
	}

	warning, idErr := setIdOption(vd)
	if idErr != nil {
		return nil, warning, idErr
	}

	return vd, warning, err

}
