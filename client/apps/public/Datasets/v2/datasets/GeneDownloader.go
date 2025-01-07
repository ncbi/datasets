package datasets

import (
	"context"
	"fmt"

	cmdflags "datasets_cli/v2/datasets/flags"
	openapi "datasets/openapi/v2"
)

type GeneDownloaderBase struct {
	cli *openapi.APIClient
}

type GeneDownloader struct {
	GeneDownloaderBase
	isPreview  bool
	geneCounts GeneCounts
	request    *openapi.V2GeneDatasetRequest
}

type GeneDownloaderIdOption func(f *GeneDownloader) error

func (gd *GeneDownloader) setGeneIdsForRequestIter(requestIter RequestIterator[*openapi.V2GeneDatasetReportsRequest]) error {
	api := GeneApi{geneApi: gd.cli.GeneAPI}

	if gd.isPreview {
		geneCountRetriever := NewGeneCountRetriever()
		requestIter.GetRequest().SetReturnedContent(openapi.V2GENEDATASETREPORTSREQUESTCONTENTTYPE_COUNTS_ONLY)
		if _, err := ProcessAllPages[*openapi.V2GeneDatasetReportsRequest,
			openapi.V2reportsGeneDataReportPage,
			openapi.V2reportsGeneReportMatch,
			*openapi.V2reportsGeneDataReportPage](
			requestIter,
			&api,
			&geneCountRetriever,
		); err != nil {
			return err
		}
		gd.geneCounts = geneCountRetriever.geneCounts
	} else {
		geneIdRetriever := NewGeneIdRetriever()
		requestIter.GetRequest().SetReturnedContent(openapi.V2GENEDATASETREPORTSREQUESTCONTENTTYPE_IDS_ONLY)
		if _, err := ProcessAllPages[*openapi.V2GeneDatasetReportsRequest,
			openapi.V2reportsGeneDataReportPage,
			openapi.V2reportsGeneReportMatch,
			*openapi.V2reportsGeneDataReportPage](
			requestIter,
			&api,
			&geneIdRetriever,
		); err != nil {
			return err
		}
		gd.request.SetGeneIds(geneIdRetriever.GetGeneIds())
	}

	return nil
}

func (gd *GeneDownloader) setGeneIdsForRequest(request *openapi.V2GeneDatasetReportsRequest) error {
	if request.HasAccessions() && len(request.GetAccessions()) > 0 {
		return gd.setGeneIdsForRequestIter(NewGeneAccessionRequestIter(request))
	} else if request.HasSymbolsForTaxon() && len(request.SymbolsForTaxon.GetSymbols()) > 0 {
		return gd.setGeneIdsForRequestIter(NewGeneSymbolRequestIter(request))
	}
	return gd.setGeneIdsForRequestIter(NewGeneIdRequestIter(request))
}

func WithSymbolAndTaxon(iff *cmdflags.InputFileFlag, otf *cmdflags.OrthologTaxonFilterFlag, taxId string) GeneDownloaderIdOption {
	return func(gd *GeneDownloader) error {
		// This may return gene-ids or symbols, based on whether it's an ortholog request
		request, err := GeneDatasetReportRequestForSymbolAndTaxon(gd.cli, iff, otf, taxId)
		if err != nil {
			return err
		}
		return gd.setGeneIdsForRequest(request)
	}
}

func WithTaxon(taxon string) GeneDownloaderIdOption {
	return func(gd *GeneDownloader) error {
		request := openapi.NewV2GeneDatasetReportsRequest()
		request.SetTaxon(taxon)
		return gd.setGeneIdsForRequest(request)
	}
}

func WithAccessions(iff *cmdflags.InputFileFlag, otf *cmdflags.OrthologTaxonFilterFlag) GeneDownloaderIdOption {
	return func(gd *GeneDownloader) error {
		// This may return gene-ids or accession, based on whether it's an ortholog request
		request, err := GeneDatasetReportRequestForAccessions(gd.cli, iff, otf)
		if err != nil {
			return err
		}
		return gd.setGeneIdsForRequest(request)
	}
}

func WithGeneIds(iff *cmdflags.InputFileFlag, otf *cmdflags.OrthologTaxonFilterFlag) GeneDownloaderIdOption {
	return func(gd *GeneDownloader) error {
		geneInts, err := GeneIdsAsIntsForInputs(gd.cli, iff, otf)
		if err != nil {
			return err
		}

		request := openapi.NewV2GeneDatasetReportsRequest()
		request.SetGeneIds(geneInts)
		return gd.setGeneIdsForRequest(request)
	}
}

func WithGeneRequest(request *openapi.V2GeneDatasetRequest) GeneDownloaderIdOption {
	return func(gd *GeneDownloader) error {
		gd.request = request
		return nil
	}
}

func NewGeneDownloader(isPreview bool, geneIncludeFlag *cmdflags.GeneIncludeFlag, fastaFilterFlag *cmdflags.FastaFilterFlag, setIdOption ...GeneDownloaderIdOption) (*GeneDownloader, error) {
	cli, err := createOAClient()
	if err != nil {
		return nil, err
	}
	gd := &GeneDownloader{
		GeneDownloaderBase: GeneDownloaderBase{
			cli: cli,
		},
		isPreview: isPreview,
		request:   openapi.NewV2GeneDatasetRequest(),
	}
	for _, opt := range setIdOption {
		optErr := opt(gd)
		if optErr != nil {
			return nil, optErr
		}
	}
	if !isPreview {
		if len(gd.request.GetGeneIds()) == 0 {
			return nil, fmt.Errorf("No genes found that match selection\n")
		}
		gd.request.SetFastaFilter(fastaFilterFlag.Filters())
		geneIncludeFlag.SetGeneDownloadFlags(gd.request)
	}
	return gd, nil
}

func (gd *GeneDownloader) Download(argSkipZipVal bool) (err error) {
	if gd.isPreview {
		if !argNoProgress {
			progress.Stop()
		}
		printResults(gd.geneCounts)
		return nil
	}
	_, resp, err := gd.cli.GeneAPI.DownloadGenePackagePost(context.TODO()).V2GeneDatasetRequest(*gd.request).Execute()
	if err = handleHTTPResponse(resp, err); err != nil {
		return
	}
	length := int64(-1) // unknown length
	err = downloadDataForFile(resp, err, argDownloadFilename, length, argSkipZipVal)
	return
}

type GeneProkDownloader struct {
	request *openapi.V2ProkaryoteGeneRequest
	GeneDownloaderBase
}

type ProkGeneDownloaderIdOption func(f *GeneProkDownloader) error

func WithProkAccessions(iff *cmdflags.InputFileFlag) ProkGeneDownloaderIdOption {
	return func(gd *GeneProkDownloader) error {
		gd.request = openapi.NewV2ProkaryoteGeneRequest()
		gd.request.SetAccessions(iff.AsStringList())
		return nil
	}
}

func NewGeneProkDownloader(isPreview bool, taxId string, geneIncludeFlag *cmdflags.GeneIncludeFlag, wdf *cmdflags.WpDownloadFlags, setIdOption ...ProkGeneDownloaderIdOption) (*GeneProkDownloader, error) {
	if isPreview {
		return nil, fmt.Errorf("Preview is not currently available for prokaryotic (WP_) genes\n")
	}

	cli, err := createOAClient()
	if err != nil {
		return nil, err
	}
	gd := &GeneProkDownloader{
		GeneDownloaderBase: GeneDownloaderBase{
			cli: cli,
		},
		request: openapi.NewV2ProkaryoteGeneRequest(),
	}
	for _, opt := range setIdOption {
		optErr := opt(gd)
		if optErr != nil {
			return nil, optErr
		}
	}
	if taxId != "" {
		gd.request.SetTaxon(taxId)
	}

	includeErr := geneIncludeFlag.SetProkDownloadFlags(gd.request)
	if includeErr != nil {
		return nil, includeErr
	}
	if wdf.IncludeFlankBp() > 0 {
		argFlank32 := int32(wdf.IncludeFlankBp())
		gd.request.SetGeneFlankConfig(openapi.V2ProkaryoteGeneRequestGeneFlankConfig{Length: &argFlank32})
		gd.request.SetIncludeAnnotationType(append(gd.request.GetIncludeAnnotationType(), openapi.V2FASTA_GENE_FLANK))
	}
	return gd, nil
}

func (gd *GeneProkDownloader) Download(argSkipZipVal bool) (err error) {
	_, resp, err := gd.cli.ProkaryoteAPI.DownloadProkaryoteGenePackagePost(context.TODO()).V2ProkaryoteGeneRequest(*gd.request).Execute()
	if err = handleHTTPResponse(resp, err); err != nil {
		return
	}
	length := int64(-1) // unknown length
	err = downloadDataForFile(resp, err, argDownloadFilename, length, argSkipZipVal)
	return
}
