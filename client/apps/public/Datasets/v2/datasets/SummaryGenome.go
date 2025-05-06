package datasets

import (
	openapi "datasets/openapi/v2"
	cmdflags "datasets_cli/v2/datasets/flags"
	"errors"
	_nethttp "net/http"

	"github.com/antihax/optional"
	"github.com/spf13/cobra"
	"gitlab.com/metakeule/fmtdate"
)

func InitializedV2SortField(fld string, dir openapi.V2SortDirection) openapi.V2SortField {
	this := openapi.V2SortField{}
	this.SetDirection(dir)
	this.SetField(fld)
	return this
}

func updateAssemblyReportRequestOption(
	request *openapi.V2AssemblyDatasetReportsRequest,
	assemblyRequestFlag AssemblyRequestFlag) (err error) {
	filter := openapi.NewV2AssemblyDatasetDescriptorsFilter()
	filter.SetReferenceOnly(assemblyRequestFlag.referenceOnlyFlag.ReferenceOnly())
	filter.SetSearchText(assemblyRequestFlag.searchStringFlag.GetSearchText())
	filter.SetHasAnnotation(assemblyRequestFlag.annotatedOnlyFlag.AnnotatedOnly())
	filter.SetExcludeAtypical(assemblyRequestFlag.excludeAtypical.ExcludeAtypical())
	filter.SetIsTypeMaterial(assemblyRequestFlag.typeMaterialFlag.TypeMaterial())
	filter.SetExcludeMultiIsolate(assemblyRequestFlag.excludeMultiIsolateFlag.ExcludeMultiIsolate())
	filter.SetIsMetagenomeDerived(assemblyRequestFlag.metaGenomeFlag.GetMetaGenomeFilter())
	filter.SetAssemblySource(assemblyRequestFlag.assemblySourceFlag.GetAssemblySourceFilter())
	filter.SetAssemblyLevel(assemblyRequestFlag.genomeAssemblyLevelFlag.GetAssemblyLevels())
	filter.SetAssemblyVersion(assemblyRequestFlag.assemblyVersionFlag.GetAssemblyVersion())

	dateAfter, dateAfterErr := getDate(assemblyRequestFlag.releasedAfterFlag.ReleasedAfterFlagDate())
	if dateAfterErr != nil {
		return dateAfterErr
	}
	if dateAfter.IsSet() {
		filter.SetFirstReleaseDate(dateAfter.Value())
	}

	dateBefore, dateBeforeErr := getDate(assemblyRequestFlag.relBeforeFlag.ReleasedBeforeFlagDate())
	if dateBeforeErr != nil {
		return dateBeforeErr
	}
	if dateBefore.IsSet() {
		filter.SetLastReleaseDate(dateBefore.Value())
	}

	request.SetFilters(*filter)
	request.SetPageSize(1000)

	sort_rules := []openapi.V2SortField{
		InitializedV2SortField("organismName", openapi.V2SORTDIRECTION_ASCENDING),
		InitializedV2SortField("isRefGenome", openapi.V2SORTDIRECTION_DESCENDING),
		InitializedV2SortField("isRepGenome", openapi.V2SORTDIRECTION_DESCENDING),
		InitializedV2SortField("isRefseq", openapi.V2SORTDIRECTION_DESCENDING),
		InitializedV2SortField("accession", openapi.V2SORTDIRECTION_ASCENDING),
	}

	request.SetSort(sort_rules)

	return nil
}

type GenomeApi struct {
	genomeApi *openapi.GenomeAPIService
}

func (apiService *GenomeApi) GetPage(request *openapi.V2AssemblyDatasetReportsRequest) (*openapi.V2reportsAssemblyDataReportPage, *_nethttp.Response, error) {
	apiRequest := openapi.ApiGenomeDatasetReportByPostRequest{
		ApiService: apiService.genomeApi,
	}
	return apiRequest.V2AssemblyDatasetReportsRequest(*request).Execute()
}

func (apiService *GenomeApi) GetPagePtr(page openapi.V2reportsAssemblyDataReportPage) *openapi.V2reportsAssemblyDataReportPage {
	return &page
}

type GenomeSequenceApi struct {
	genomeApi *openapi.GenomeAPIService
}

func (apiService *GenomeSequenceApi) GetPage(request *openapi.V2AssemblySequenceReportsRequest) (*openapi.V2SequenceReportPage, *_nethttp.Response, error) {
	apiRequest := openapi.ApiGenomeSequenceReportByPostRequest{
		ApiService: apiService.genomeApi,
	}
	return apiRequest.V2AssemblySequenceReportsRequest(*request).Execute()
}

func (apiService *GenomeSequenceApi) GetPagePtr(page openapi.V2SequenceReportPage) *openapi.V2SequenceReportPage {
	return &page
}

func getDate(argDate string) (optional.Time, error) {
	if argDate != "" {
		formattedDate, err := fmtdate.Parse(dateFormat, argDate)
		if err != nil {
			return optional.EmptyTime(), err
		}
		return optional.NewTime(formattedDate), nil
	}
	return optional.EmptyTime(), nil
}

func getGenomeApi() (api *GenomeApi, err error) {
	cli, cliErr := createOAClient()
	if cliErr != nil {
		return nil, cliErr
	}
	api = &GenomeApi{genomeApi: cli.GenomeAPI}
	return api, nil
}

func getGenomeSequenceApi() (api *GenomeSequenceApi, err error) {
	cli, cliErr := createOAClient()
	if cliErr != nil {
		return nil, cliErr
	}
	api = &GenomeSequenceApi{genomeApi: cli.GenomeAPI}
	return api, nil
}

func NewGenomeAccessionRequestIter(request *openapi.V2AssemblyDatasetReportsRequest) *DefaultPagedRequestIterator[*openapi.V2AssemblyDatasetReportsRequest, string] {
	accRequester := NewDefaultPagedRequestIterator[*openapi.V2AssemblyDatasetReportsRequest, string](request,
		func(request *openapi.V2AssemblyDatasetReportsRequest, accessions []string) {
			request.SetAccessions(accessions)
		},
	)

	if len(request.GetAccessions()) > PAGE_ITER_THRESHOLD {
		accRequester.ids = request.GetAccessions()
		request.SetAccessions([]string{})
	}

	return accRequester
}

func getGenomeReport(requestIter RequestIterator[*openapi.V2AssemblyDatasetReportsRequest], sgf SummaryGenomeFlag) (err error) {
	api, apiErr := getGenomeApi()
	if apiErr != nil {
		return apiErr
	}
	pagePrinter := NewPagePrinter[openapi.V2reportsAssemblyDataReport, *openapi.V2reportsAssemblyDataReportPage](
		"genome",
		sgf.jsonLinesLimitFlag.JsonLines(),
	)
	_, err = ProcessPages[*openapi.V2AssemblyDatasetReportsRequest,
		openapi.V2reportsAssemblyDataReportPage,
		openapi.V2reportsAssemblyDataReport,
		*openapi.V2reportsAssemblyDataReportPage](requestIter, api, &pagePrinter, sgf.jsonLinesLimitFlag.RetrievalCount(), sgf.jsonLinesLimitFlag.CountOnly())

	return
}

func getSequenceReport(
	requestIter RequestIterator[*openapi.V2AssemblyDatasetReportsRequest],
	sgf SummaryGenomeFlag,
) (err error) {
	api, apiErr := getGenomeApi()
	if apiErr != nil {
		return apiErr
	}

	// Gather accessions using the standard report interface (which has the full set of filter parameters)
	requestIter.GetRequest().SetReturnedContent(openapi.V2ASSEMBLYDATASETREPORTSREQUESTCONTENTTYPE_ASSM_ACC)
	accRetriever := GenomeAccRetriever{}
	_, metaErr := ProcessAllPages[*openapi.V2AssemblyDatasetReportsRequest,
		openapi.V2reportsAssemblyDataReportPage,
		openapi.V2reportsAssemblyDataReport,
		*openapi.V2reportsAssemblyDataReportPage](requestIter, api, &accRetriever)
	if metaErr != nil {
		return metaErr
	}
	if len(accRetriever.assemblyAccessions) == 0 {
		return errors.New("No accessions match current filters")
	}

	// Print sequence reports
	seqReportRequest := openapi.NewV2AssemblySequenceReportsRequest()
	seqReportRequest.SetPageSize(100)

	sequenceApi, sequenceApiErr := getGenomeSequenceApi()
	if sequenceApiErr != nil {
		return sequenceApiErr
	}
	// Iterate over accessions, printing results for one accession at a time
	pagePrinter := NewPagePrinter[openapi.V2reportsSequenceInfo, *openapi.V2SequenceReportPage](
		"genome",
		sgf.jsonLinesLimitFlag.JsonLines(),
		WithSkipFinish[openapi.V2reportsSequenceInfo, *openapi.V2SequenceReportPage](),
	)

	for idx, acc := range accRetriever.assemblyAccessions {
		seqReportRequest.SetAccession(acc)
		seqReportRequest.SetPageToken("")
		if idx == len(accRetriever.assemblyAccessions)-1 {
			pagePrinter.skipFinish = false
		}
		_, err = ProcessPagesRequest[*openapi.V2AssemblySequenceReportsRequest,
			openapi.V2SequenceReportPage,
			openapi.V2reportsSequenceInfo,
			*openapi.V2SequenceReportPage](seqReportRequest, sequenceApi, &pagePrinter, sgf.jsonLinesLimitFlag.RetrievalCount(), sgf.jsonLinesLimitFlag.CountOnly())
		if err != nil {
			return err
		}
	}
	return nil
}

func getGenomeSummary(
	requestIter RequestIterator[*openapi.V2AssemblyDatasetReportsRequest],
	sgf SummaryGenomeFlag,
	assemblyRequestFlag AssemblyRequestFlag) (err error) {
	if sgf.rptFlag.IdsOnly() {
		requestIter.GetRequest().SetReturnedContent(openapi.V2ASSEMBLYDATASETREPORTSREQUESTCONTENTTYPE_ASSM_ACC)
	}

	err = updateAssemblyReportRequestOption(requestIter.GetRequest(), assemblyRequestFlag)
	if err != nil {
		return
	}

	if sgf.rptFlag.IdsOnly() || sgf.rptFlag.GenomeReport() {
		return getGenomeReport(requestIter, sgf)
	}
	return getSequenceReport(requestIter, sgf)
}

// Flags for assembly requests - shared with DownloadGenome.
type AssemblyRequestFlag struct {
	excludeAtypical         *cmdflags.ExcludeAtypicalFlag
	typeMaterialFlag        *cmdflags.TypeMaterialFlag
	referenceOnlyFlag       *cmdflags.ReferenceOnlyFlag
	annotatedOnlyFlag       *cmdflags.AnnotatedOnlyFlag
	metaGenomeFlag          *cmdflags.MetaGenomeDerivedFlag
	assemblySourceFlag      *cmdflags.AssemblySourceFlag
	genomeAssemblyLevelFlag *cmdflags.GenomeAssemblyLevelFlag
	searchStringFlag        *cmdflags.SearchStringFlag
	assemblyVersionFlag     *cmdflags.AssemblyVersionFlag
	relBeforeFlag           *cmdflags.ReleasedBeforeFlag
	releasedAfterFlag       *cmdflags.ReleasedAfterFlag
	excludeMultiIsolateFlag *cmdflags.ExcludeMultiIsolateFlag

	cmdFlagSet []cmdflags.FlagInterface
}

func initAssemblyRequestFlag() AssemblyRequestFlag {
	ea := cmdflags.NewExcludeAtypicalFlag()
	tmf := cmdflags.NewTypeMaterialFlag()
	rof := cmdflags.NewReferenceOnlyFlag()
	aof := cmdflags.NewAnnotatedOnlyFlag()
	mgf := cmdflags.NewMetaGenomeDerivedFlag()
	asf := cmdflags.NewAssemblySourceFlag()
	alf := cmdflags.NewGenomeAssemblyLevelFlag()
	ssf := cmdflags.NewSearchStringFlag()
	avf := cmdflags.NewAssemblyVersionFlag(cmdflags.Default)
	rbf := cmdflags.NewReleasedBeforeFlag(cmdflags.GenomeReleasedBeforeDesc)
	raf := cmdflags.NewReleasedAfterFlag(cmdflags.GenomeReleasedAfterDesc)
	emi := cmdflags.NewExcludeMultiIsolateFlag()

	sgf := AssemblyRequestFlag{
		excludeAtypical:         ea,
		typeMaterialFlag:        tmf,
		referenceOnlyFlag:       rof,
		annotatedOnlyFlag:       aof,
		metaGenomeFlag:          mgf,
		assemblySourceFlag:      asf,
		genomeAssemblyLevelFlag: alf,
		searchStringFlag:        ssf,
		assemblyVersionFlag:     avf,
		relBeforeFlag:           rbf,
		releasedAfterFlag:       raf,
		excludeMultiIsolateFlag: emi,

		cmdFlagSet: []cmdflags.FlagInterface{ea, tmf, rof, aof, mgf, asf, alf, ssf, avf, rbf, raf, emi},
	}

	return sgf
}

// Flags set used only be the genome summary command
type SummaryGenomeFlag struct {
	rptFlag            *cmdflags.GenomeReportFlag
	jsonLinesLimitFlag *cmdflags.JsonLinesAndLimitFlag
	cmdFlagSet         []cmdflags.FlagInterface
}

func initSummaryGenomeFlag() SummaryGenomeFlag {
	rf := cmdflags.NewGenomeReportFlag()
	jll := cmdflags.NewJsonLineAndLimitFlag("genome")

	sgf := SummaryGenomeFlag{
		rptFlag:            rf,
		jsonLinesLimitFlag: jll,
		cmdFlagSet:         []cmdflags.FlagInterface{rf, jll},
	}

	return sgf
}

func createSummaryGenomeCmd() *cobra.Command {
	summaryGenomeFlag := initSummaryGenomeFlag()
	assemblyRequestFlag := initAssemblyRequestFlag()
	allFlags := append(summaryGenomeFlag.cmdFlagSet, assemblyRequestFlag.cmdFlagSet...)

	cmd := &cobra.Command{
		Use:   "genome",
		Short: "Print a data report containing genome metadata",
		Long: `
Print a data report containing genome metadata. The data report is returned in JSON format.`,
		Example: `  datasets summary genome accession GCF_000001405.40
  datasets summary genome taxon "mus musculus"
  datasets summary genome taxon human --assembly-level chromosome,complete
  datasets summary genome taxon "mus musculus" --search C57BL/6J --search "Broad Institute"`,
		Args:              cobra.NoArgs,
		PersistentPreRunE: cmdflags.PersistentPreRunEFor(allFlags, rootCmd),
		RunE:              ParentCommandRunE,
	}

	cmdflags.RegisterAllFlags(allFlags, cmd.PersistentFlags())

	cmd.AddCommand(createSummaryGenomeAccessionCmd(summaryGenomeFlag, assemblyRequestFlag))
	cmd.AddCommand(createSummaryGenomeTaxonCmd(summaryGenomeFlag, assemblyRequestFlag))

	return cmd
}
