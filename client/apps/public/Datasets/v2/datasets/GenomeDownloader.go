package datasets

import (
	"context"
	openapi "datasets/openapi/v2"
	cmdflags "datasets_cli/v2/datasets/flags"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type GenomeDownloader struct {
	request *openapi.V2AssemblyDatasetRequest
	flags   AssemblyRequestFlag
	cli     *openapi.APIClient
}

type GenomeDownloaderAccOption func(f *GenomeDownloader) error

func (gd *GenomeDownloader) setAssemblyAccsForRequest(request *openapi.V2AssemblyDatasetReportsRequest) error {
	reqErr := updateAssemblyReportRequestOption(request, gd.flags)
	if reqErr != nil {
		return reqErr
	}

	request.SetReturnedContent(openapi.V2ASSEMBLYDATASETREPORTSREQUESTCONTENTTYPE_ASSM_ACC)
	api := GenomeApi{genomeApi: gd.cli.GenomeAPI}
	accRetriever := GenomeAccRetriever{}

	_, err := ProcessPages[*openapi.V2AssemblyDatasetReportsRequest,
		openapi.V2reportsAssemblyDataReportPage,
		openapi.V2reportsAssemblyDataReport,
		*openapi.V2reportsAssemblyDataReportPage](NewGenomeAccessionRequestIter(request), &api, &accRetriever, math.MaxInt, false)

	if err == nil {
		gd.request.SetAccessions(append(gd.request.GetAccessions(), accRetriever.assemblyAccessions...))
	}

	return err
}

func isValidAssemblyAcc(acc string) bool {
	isAccVerSuffix := func(accver_suffix string) bool {
		_, err := strconv.ParseFloat(accver_suffix, 32)
		return (err == nil)
	}

	validPrefixes := []string{
		"GCA_",
		"GCF_",
		"PRJNA",
		"PRJDA",
		"PRJDB",
		"PRJEA",
		"PRJEB",
	}
	for _, prefix := range validPrefixes {
		if strings.HasPrefix(acc, prefix) && isAccVerSuffix(acc[len(prefix):]) {
			return true
		}
	}
	return false
}

func separateAccessions(accs []string) (retAccs []string, bioprojectAccs []string, err error) {
	accNum := 0
	for _, acc := range accs {
		acc = strings.ToUpper(acc)
		if isValidAssemblyAcc(acc) {
			if strings.HasPrefix(acc, "PRJ") {
				bioprojectAccs = append(bioprojectAccs, acc)
			} else {
				accs[accNum] = acc
				accNum++
			}
		} else {
			err = fmt.Errorf("invalid or unsupported assembly accession: %s", acc)
			return
		}
	}
	retAccs = accs[:accNum]
	return
}

func GetGenomeReportsTaxonRequest(taxons []string, taxExactMatch bool) *openapi.V2AssemblyDatasetReportsRequest {
	request := openapi.NewV2AssemblyDatasetReportsRequest()
	request.SetTaxons(taxons)
	request.SetTaxExactMatch(taxExactMatch)
	return request
}

func GenomeWithTaxon(taxons []string, taxExactMatch bool) GenomeDownloaderAccOption {
	return func(gd *GenomeDownloader) error {
		return gd.setAssemblyAccsForRequest(GetGenomeReportsTaxonRequest(taxons, taxExactMatch))
	}
}

func GetGenomeReportsAccessionRequest(accessions []string) (*openapi.V2AssemblyDatasetReportsRequest, error) {
	accs, bioprojectAccs, err := separateAccessions(accessions)
	if err != nil {
		return nil, err
	}
	request := openapi.NewV2AssemblyDatasetReportsRequest()
	if len(bioprojectAccs) > 0 {
		request.SetBioprojects(bioprojectAccs)
	}
	if len(accs) > 0 {
		request.SetAccessions(accs)
	}
	return request, nil
}

func GenomeWithAccessions(accessions []string) GenomeDownloaderAccOption {
	return func(gd *GenomeDownloader) error {
		request, err := GetGenomeReportsAccessionRequest(accessions)
		if err != nil {
			return err
		}
		return gd.setAssemblyAccsForRequest(request)
	}
}

func WithGenomeRequest(request *openapi.V2AssemblyDatasetRequest) GenomeDownloaderAccOption {
	return func(gd *GenomeDownloader) error {
		gd.request = request
		return nil
	}
}

func SetRequestArgs(request *openapi.V2AssemblyDatasetRequest, downloadGenomeFlag DownloadGenomeFlag) (warning string, err error) {
	annotations := make([]openapi.V2AnnotationForAssemblyType, 0)

	for _, fl := range downloadGenomeFlag.genomeIncludeAnnotationFlag.IncludeAnnotation {
		if fl == cmdflags.GenomeIncludeNoneFlag {
			if len(downloadGenomeFlag.genomeIncludeAnnotationFlag.IncludeAnnotation) != 1 {
				err = fmt.Errorf("If file type 'none' is requested, no other file types may be included")
				return
			}
			break
		}

		annotations = append(annotations, cmdflags.GenomeIncludeAnnotationOpenapi[fl])
		annotation_file := true
		if fl == cmdflags.GenomeSeq || fl == cmdflags.GenomeSequenceReport {
			annotation_file = false
		}
		if (len(downloadGenomeFlag.chromosomesFlag.GetChromosomes()) > 0 && warning == "") && annotation_file {
			warning = "Warning - included annotation files cover whole genome even though chromosome fasta was selected"
		}
	}
	request.SetIncludeAnnotationType(annotations)

	if len(downloadGenomeFlag.chromosomesFlag.GetChromosomes()) > 0 {
		request.SetChromosomes(downloadGenomeFlag.chromosomesFlag.GetChromosomes())
	}
	if downloadGenomeFlag.dehydratedFlag.Dehydrated() {
		request.SetHydrated(openapi.V2ASSEMBLYDATASETREQUESTRESOLUTION_DATA_REPORT_ONLY)
	}
	return
}

func (gd *GenomeDownloader) setRequestArgs(downloadGenomeFlag DownloadGenomeFlag) (warning string, err error) {
	return SetRequestArgs(gd.request, downloadGenomeFlag)
}

func NewGenomeDownloader(setIdOption GenomeDownloaderAccOption, downloadGenomeFlag DownloadGenomeFlag, assemblyRequestFlag AssemblyRequestFlag) (*GenomeDownloader, string, error) {
	cli, err := createOAClient()
	if err != nil {
		return nil, "", err
	}
	gd := &GenomeDownloader{
		request: openapi.NewV2AssemblyDatasetRequest(),
		flags:   assemblyRequestFlag,
		cli:     cli,
	}
	optErr := setIdOption(gd)
	if optErr != nil {
		return nil, "", optErr
	}
	if len(gd.request.GetAccessions()) == 0 {
		return nil, "", fmt.Errorf("NCBI Datasets provides access to assembled genome data by BioProject accession and no genome assemblies were found for the given BioProject.\nRetrieval of SRA data by BioProject accession is not supported by NCBI Datasets.\n")
	}

	warning, fastaErr := gd.setRequestArgs(downloadGenomeFlag)
	if fastaErr != nil {
		return nil, warning, fastaErr
	}

	return gd, warning, nil
}

func NewGenomeRequestDownloader(request *openapi.V2AssemblyDatasetRequest, assemblyRequestFlag AssemblyRequestFlag) (*GenomeDownloader, error) {
	cli, err := createOAClient()
	if err != nil {
		return nil, err
	}
	gd := &GenomeDownloader{
		request: request,
		flags:   assemblyRequestFlag,
		cli:     cli,
	}

	if len(gd.request.GetAccessions()) == 0 {
		return nil, fmt.Errorf("Request must include one or more genome accessions\n")
	}

	return gd, nil
}

func (gd *GenomeDownloader) Download(argSkipZipVal bool) (err error) {
	_, resp, err := gd.cli.GenomeAPI.DownloadAssemblyPackagePost(context.TODO()).V2AssemblyDatasetRequest(*gd.request).Execute()
	if err = handleHTTPResponse(resp, err); err != nil {
		return
	}
	length := int64(-1) // unknown length
	err = downloadDataForFile(resp, err, argDownloadFilename, length, argSkipZipVal)
	return
}
