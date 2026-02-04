package datasets

import (
	"fmt"
	"math"
	_nethttp "net/http"

	"github.com/gosuri/uiprogress"
	"golang.org/x/text/language"
	"golang.org/x/text/message"

	openapi "datasets/openapi/v2"
)

const (
	MAX_PAGE_SIZE       int = 1000
	PAGE_ITER_THRESHOLD int = 2000
)

type PagedRequest interface {
	*openapi.V2AssemblyDatasetReportsRequest |
		*openapi.V2GeneDatasetReportsRequest |
		*openapi.V2OrthologRequest |
		*openapi.V2VirusDataReportRequest |
		*openapi.V2VirusAnnotationReportRequest |
		*openapi.V2AssemblySequenceReportsRequest |
		*openapi.V2TaxonomyMetadataRequest |
		*openapi.V2TaxonomyRelatedIdRequest
	SetPageToken(string)
	GetPageToken() string
	SetPageSize(int32)
}

type PagedResponse interface {
	openapi.V2reportsAssemblyDataReportPage |
		openapi.V2reportsGeneDataReportPage |
		openapi.V2reportsVirusDataReportPage |
		openapi.V2reportsVirusAnnotationReportPage |
		openapi.V2SequenceReportPage |
		openapi.V2reportsTaxonomyDataReportPage |
		openapi.V2reportsTaxonomyNamesDataReportPage
	MarshalJSON() ([]byte, error)
}

type PPagedResponse[REP DatasetReport] interface {
	*openapi.V2reportsAssemblyDataReportPage |
		*openapi.V2reportsGeneDataReportPage |
		*openapi.V2reportsVirusDataReportPage |
		*openapi.V2reportsVirusAnnotationReportPage |
		*openapi.V2SequenceReportPage |
		*openapi.V2reportsTaxonomyDataReportPage |
		*openapi.V2reportsTaxonomyNamesDataReportPage
	GetTotalCount() int32
	SetTotalCount(int32)
	GetReports() []REP
	GetNextPageToken() string
}

type DatasetReport interface {
	openapi.V2reportsAssemblyDataReport |
		openapi.V2reportsVirusAssembly |
		openapi.V2reportsVirusAnnotationReport |
		openapi.V2reportsGeneReportMatch |
		openapi.V2reportsSequenceInfo |
		openapi.V2reportsTaxonomyReportMatch |
		openapi.V2reportsTaxonomyNamesReportMatch
}

type PageRetriever[R PagedRequest, P PagedResponse, REP DatasetReport, PP PPagedResponse[REP]] interface {
	GetPage(R) (*P, *_nethttp.Response, error)
	GetPagePtr(P) PP
}

type PageProcessor[REP DatasetReport, PP PPagedResponse[REP]] interface {
	ProcessPage(PP)
	Finish(PP)
	SetLimit(int)
	ReportName() string
	RetrievalCount(int) int
}

// If you have an implementation of PageProcessor that doesn't need any of the
// Finsih(), SetLimit() or ReportName() functions you can include this to get the defaults, e.g.
// type AssmAccRetriever struct { DefaultPageProcessorFuncs[openapi.V2reportsAssemblyDataReport, *openapi.V2reportsAssemblyDataReportPage]}
type DefaultPageProcessorFuncs[REP DatasetReport, PP PPagedResponse[REP]] struct {
}

func (skipper *DefaultPageProcessorFuncs[REP, PP]) SetLimit(_ int) {
}

func (skipper *DefaultPageProcessorFuncs[REP, PP]) ReportName() string {
	return ""
}

func (skipper *DefaultPageProcessorFuncs[REP, PP]) Finish(_ PP) {
}

func (skipper *DefaultPageProcessorFuncs[REP, PP]) RetrievalCount(totalCount int) int {
	return totalCount
}

type PagePrinter[REP DatasetReport, PP PPagedResponse[REP]] struct {
	recCount      int
	headerPrinted bool
	limit         int
	reportName    string
	skipFinish    bool
	totalCount    int
	asJsonLines   bool

	// Some services, e.g. gene, may return the same record for different ids (e.g. merged genes)
	duplicatesMap        map[string]bool
	printJsonResult      func(REP)
	printJsonLinesResult func(REP)
	isDuplicateCheck     func(REP, *map[string]bool) bool
	// Print messages allows the cli to add messages to the json output that were not returned by the api
	printMessages func(bool, bool) bool
}

func defaultPrintJsonLinesReport[REP DatasetReport](datasetReport REP) {
	printResults(datasetReport)
}

func defaultPrintJsonReport[REP DatasetReport](datasetReport REP) {
	printResultsNoNewline(datasetReport)
}

func defaultIsDuplicateCheck[REP DatasetReport](datasetReport REP, duplicatesMap *map[string]bool) bool {
	return false
}

func defaultPrintMessages(printArrayToo bool, jsonFormat bool) bool {
	return false
}

type Option[REP DatasetReport, PP PPagedResponse[REP]] func(f *PagePrinter[REP, PP])

func WithSkipFinish[REP DatasetReport, PP PPagedResponse[REP]]() Option[REP, PP] {
	return func(f *PagePrinter[REP, PP]) {
		f.skipFinish = true
	}
}

func WithPrintJsonResult[REP DatasetReport, PP PPagedResponse[REP]](printJsonResult func(REP)) Option[REP, PP] {
	return func(f *PagePrinter[REP, PP]) {
		f.printJsonResult = printJsonResult
	}
}

func WithPrintJsonLinesResult[REP DatasetReport, PP PPagedResponse[REP]](printJsonLinesResult func(REP)) Option[REP, PP] {
	return func(f *PagePrinter[REP, PP]) {
		f.printJsonLinesResult = printJsonLinesResult
	}
}

func WithCheckDuplicates[REP DatasetReport, PP PPagedResponse[REP]](isDuplicateCheck func(REP, *map[string]bool) bool) Option[REP, PP] {
	return func(f *PagePrinter[REP, PP]) {
		f.isDuplicateCheck = isDuplicateCheck
	}
}

func WithPrintMessages[REP DatasetReport, PP PPagedResponse[REP]](printMessages func(bool, bool) bool) Option[REP, PP] {
	return func(f *PagePrinter[REP, PP]) {
		f.printMessages = printMessages
	}
}

func NewPagePrinter[REP DatasetReport, PP PPagedResponse[REP]](reportName string, asJsonLines bool, opts ...Option[REP, PP]) PagePrinter[REP, PP] {
	pp := PagePrinter[REP, PP]{
		reportName:           reportName,
		asJsonLines:          asJsonLines,
		printJsonResult:      defaultPrintJsonReport[REP],
		printJsonLinesResult: defaultPrintJsonLinesReport[REP],
		isDuplicateCheck:     defaultIsDuplicateCheck[REP],
		duplicatesMap:        make(map[string]bool),
		printMessages:        defaultPrintMessages,
	}

	for _, applyOpt := range opts {
		applyOpt(&pp)
	}

	return pp
}

func (pagePrinter *PagePrinter[REP, PP]) SetLimit(l int) {
	pagePrinter.limit = l
}

func (pagePrinter *PagePrinter[REP, PP]) ReportName() string {
	return pagePrinter.reportName
}

func (pagePrinter *PagePrinter[REP, PP]) ProcessPage(ppage PP) {
	for _, report := range ppage.GetReports() {
		if pagePrinter.limit == 0 {
			break
		}
		if pagePrinter.recCount >= pagePrinter.limit {
			break
		}
		if !pagePrinter.isDuplicateCheck(report, &pagePrinter.duplicatesMap) {
			if !pagePrinter.asJsonLines {
				if !pagePrinter.headerPrinted {
					fmt.Print("{\"reports\": [")
					pagePrinter.headerPrinted = true
					// If cli has any of its own messages to print, add them at beginning of the output
					hasMessages := pagePrinter.printMessages(false, pagePrinter.asJsonLines)
					if hasMessages {
						fmt.Printf(",")
					}
				} else {
					fmt.Printf(",")
				}
				pagePrinter.printJsonResult(report)
			} else {
				if pagePrinter.recCount == 0 {
					pagePrinter.printMessages(false, pagePrinter.asJsonLines)
				}
				pagePrinter.printJsonLinesResult(report)
			}
			pagePrinter.recCount++
		}
	}
}

func (pagePrinter *PagePrinter[REP, PP]) RetrievalCount(totalCount int) int {
	return pagePrinter.recCount
}

func (pagePrinter *PagePrinter[REP, PP]) finishPrint() {
	if !pagePrinter.asJsonLines {
		if pagePrinter.headerPrinted {
			fmt.Printf("],")
		} else {
			if !pagePrinter.printMessages(true, pagePrinter.asJsonLines) {
				fmt.Printf("{")
			}
		}
		fmt.Printf("\"total_count\": %d}\n", pagePrinter.recCount)
		return
	}
	if pagePrinter.totalCount == 0 {
		pagePrinter.printMessages(false, pagePrinter.asJsonLines)
	}
}

func (pagePrinter *PagePrinter[REP, PP]) Finish(ppage PP) {
	// totalCount allows caller to call process-pages multiple times and sum up results
	pagePrinter.totalCount += int(ppage.GetTotalCount())
	if !pagePrinter.skipFinish {
		// If limit is 0, only the count is being returned, so recCount will be 0.
		if pagePrinter.limit == 0 {
			pagePrinter.recCount = pagePrinter.totalCount
		}
		pagePrinter.finishPrint()
	}
}

func pagePtrWithRetryFor[R PagedRequest, P PagedResponse, REP DatasetReport, PP PPagedResponse[REP]](request R, api PageRetriever[R, P, REP, PP]) (response PP, err error) {
	page, resp, getPageErr := api.GetPage(request)
	err = handleHTTPResponse(resp, getPageErr)

	if page != nil {
		response = api.GetPagePtr(*page)
		return
	}
	return nil, err
}

// RequestIterator will incrementally move request ids (accessions, gene ids, etc) from an 'ids' array
// to the request. This handles the case where the request has a (potentially) huge number of ids that
// can't be processed (even with server-side paging) in a single query.
type RequestIterator[R PagedRequest] interface {
	UpdateRequest()
	GetRequest() R
	Finished() bool
	SetRequestSize(int)
	HasIds() bool
	RemainingIds() int
}

type IdRequestIterator[R PagedRequest, T any] struct {
	request     R
	requestSize int
	ids         []T
	setIds      func(R, []T)
}

func (requestIter *IdRequestIterator[R, T]) GetRequest() R {
	return requestIter.request
}

func (requestIter *IdRequestIterator[R, T]) HasIds() bool {
	return len(requestIter.ids) > 0
}

func (requestIter *IdRequestIterator[R, T]) RemainingIds() int {
	return len(requestIter.ids)
}

func (requestIter *IdRequestIterator[R, T]) SetRequestSize(rsize int) {
	requestIter.requestSize = rsize
}

func (requestIter *IdRequestIterator[R, T]) Finished() bool {
	// return true if requestIter.ids is exhausted and we've paged though current request
	return len(requestIter.ids) == 0 && requestIter.request.GetPageToken() == ""
}

func (requestIter *IdRequestIterator[R, T]) GetIdUpdateCount() int {
	if requestIter.request.GetPageToken() != "" {
		return 0
	}

	if len(requestIter.ids) < requestIter.requestSize {
		return len(requestIter.ids)
	}
	return requestIter.requestSize
}

// DefaultRequestIterator is for requests that never need client-side paging, e.g. taxon requests
type DefaultRequestIterator[R PagedRequest] struct {
	IdRequestIterator[R, int32]
}

func (requestIter *DefaultRequestIterator[R]) UpdateRequest() {
}

func NewDefaultRequestIterator[R PagedRequest](request R) *DefaultRequestIterator[R] {
	return &DefaultRequestIterator[R]{
		IdRequestIterator: IdRequestIterator[R, int32]{
			request:     request,
			requestSize: MAX_PAGE_SIZE,
		},
	}
}

// DefaultPagedRequestIterator can be instantiated for any request type that may need client side paging, e.g.
// gene ids, accessions etc. The second generic parameter is the id type, e.g. string for accessions.
type DefaultPagedRequestIterator[R PagedRequest, T any] struct {
	IdRequestIterator[R, T]
}

func (requestIter *DefaultPagedRequestIterator[R, T]) UpdateRequest() {
	sliceSize := requestIter.GetIdUpdateCount()

	if sliceSize > 0 {
		requestIter.setIds(requestIter.request, requestIter.ids[:sliceSize])
		requestIter.ids = requestIter.ids[sliceSize:]
	}
}

func NewDefaultPagedRequestIterator[R PagedRequest, T any](request R, setRequestIds func(R, []T)) *DefaultPagedRequestIterator[R, T] {
	return &DefaultPagedRequestIterator[R, T]{
		IdRequestIterator: IdRequestIterator[R, T]{
			request:     request,
			requestSize: MAX_PAGE_SIZE,
			setIds:      setRequestIds,
		},
	}
}

func ProcessAllPages[R PagedRequest, P PagedResponse, REP DatasetReport, PP PPagedResponse[REP]](requestIter RequestIterator[R], api PageRetriever[R, P, REP, PP], pageProcessor PageProcessor[REP, PP]) (int, error) {
	return ProcessPages[R, P, REP, PP](requestIter, api, pageProcessor, math.MaxInt, false)
}

func ProcessAllPagesRequest[R PagedRequest, P PagedResponse, REP DatasetReport, PP PPagedResponse[REP]](request R, api PageRetriever[R, P, REP, PP], pageProcessor PageProcessor[REP, PP]) (int, error) {
	return ProcessPages[R, P, REP, PP](NewDefaultRequestIterator[R](request), api, pageProcessor, math.MaxInt, false)
}

func ProcessPagesRequest[R PagedRequest, P PagedResponse, REP DatasetReport, PP PPagedResponse[REP]](request R, api PageRetriever[R, P, REP, PP], pageProcessor PageProcessor[REP, PP], maxRetrieval int, countOnly bool) (int, error) {
	return ProcessPages[R, P, REP, PP](NewDefaultRequestIterator[R](request), api, pageProcessor, maxRetrieval, countOnly)
}

func ProcessPages[R PagedRequest, P PagedResponse, REP DatasetReport, PP PPagedResponse[REP]](requestIter RequestIterator[R], api PageRetriever[R, P, REP, PP], pageProcessor PageProcessor[REP, PP], maxRetrieval int, countOnly bool) (int, error) {
	if countOnly {
		pageProcessor.SetLimit(0)
	} else {
		pageProcessor.SetLimit(maxRetrieval)
	}

	var bar *uiprogress.Bar

	// Can't do an initial query to get total count if there are a large number of ids (they need to be paged by client)
	bar_count := 0
	if !requestIter.HasIds() {
		requestIter.GetRequest().SetPageSize(int32(1))
		ppage_first, err_first := pagePtrWithRetryFor(requestIter.GetRequest(), api)
		if err_first != nil {
			return 0, err_first
		}
		bar_count = int(ppage_first.GetTotalCount())

		// We check len(ppage_first.GetReports()) too because the taxonomy endpoint has deprecated the total_count field (SEQUI-8537)
		if countOnly || ((len(ppage_first.GetReports()) == 0) && ppage_first.GetTotalCount() == 0) {
			pageProcessor.Finish(ppage_first)
			return int(ppage_first.GetTotalCount()), nil
		}
	} else {
		bar_count = requestIter.RemainingIds()
	}

	if bar_count > 0 {
		bar = progress.AddBar(bar_count).AppendCompleted()
		p := message.NewPrinter(language.English)
		bar.PrependFunc(func(b *uiprogress.Bar) string {
			plural := ""
			if b.Total > 1 {
				plural = "s"
			}
			reportName := pageProcessor.ReportName()
			if reportName != "" {
				reportName += " "
			}
			return p.Sprintf("Collecting %d %srecord%s", b.Total, reportName, plural)
		})
		bar.AppendFunc(func(b *uiprogress.Bar) string {
			return fmt.Sprintf("%d/%d", b.Current(), b.Total)
		})
		bar.Width = 50
	}

	var retrievalCount int = 0
	pageSize := minOf(MAX_PAGE_SIZE, maxRetrieval)
	requestIter.GetRequest().SetPageSize(int32(pageSize))
	for {
		requestIter.UpdateRequest()
		ppage, err := pagePtrWithRetryFor(requestIter.GetRequest(), api)
		if err != nil {
			return 0, err
		}

		retrievalCount += len(ppage.GetReports())
		if bar != nil {
			bar.Set(int(retrievalCount)) //nolint:errcheck
		}
		if !countOnly && len(ppage.GetReports()) > 0 {
			pageProcessor.ProcessPage(ppage)
		}

		// ppage.GetNextPageToken() may be blank, in which case UpdateRequest will refresh the ids if needed
		requestIter.GetRequest().SetPageToken(ppage.GetNextPageToken())
		if requestIter.Finished() || (pageProcessor.RetrievalCount(retrievalCount) >= maxRetrieval && !countOnly) {
			ppage.SetTotalCount(int32(retrievalCount))
			pageProcessor.Finish(ppage)
			break
		}
	}
	return retrievalCount, nil
}
