package datasets

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gosuri/uiprogress"
	"github.com/spf13/cobra"
	"golang.org/x/text/language"
	"golang.org/x/text/message"

	openapi "datasets_cli/v1/openapi"
)

func summaryVirusGenomeRequestWithAcc(args []string) (request *openapi.V1VirusDataReportRequest, err error) {

	filter, err := setBaseFilter()
	filter.SetAccessions(args)
	request = openapi.NewV1VirusDataReportRequest()
	request.SetFilter(filter)
	request.SetPageSize(1000)
	return
}

func printViralMetadataWithPost(request *openapi.V1VirusDataReportRequest) error {
	var assm_count int
	var header_printed bool
	var err error
	var total_count int

	printVirus := func(report openapi.V1reportsVirusAssembly) {
		if !argJsonLinesFormat {
			if !header_printed {
				fmt.Printf("\"reports\": [")
				header_printed = true
			}
			if assm_count > 0 {
				fmt.Printf(",")
			}
			assm_count++
			printResultsNoNewline(report)
		} else {
			printResults(report)
		}
	}

	if !argJsonLinesFormat {
		fmt.Printf("{")
		defer func() {
			if header_printed {
				fmt.Printf("],")
			}
			fmt.Printf("\"total_count\": %d}\n", total_count)
		}()
	}
	total_count, err = getSummaryVirusGenomeByAcc(request, printVirus)
	return err
}

func getSummaryVirusGenomeByAcc(request *openapi.V1VirusDataReportRequest, callbackFn func(openapi.V1reportsVirusAssembly)) (int, error) {
	cli, err := createOAClient()
	if err != nil {
		return 0, err
	}

	var arg_limit int
	limit := false
	countOnly := false
	if argLimit != "" {
		if argLimit == "none" || argLimit == "0" {
			request.SetPageSize(1)
			countOnly = true
		} else if argLimit != "all" {
			arg_limit, err = strconv.Atoi(argLimit)
			request.SetPageSize(mint32(int32(arg_limit), int32(1000)))
			if err != nil {
				err = fmt.Errorf("Invalid 'limit' value %s. Must be 'all', 'none', or a number.", argLimit)
				return 0, err
			}
			limit = true
		}
	}
	var bar *uiprogress.Bar
	p := message.NewPrinter(language.English)

	retrievalCount := 0
	limitCount := 0
	cont := true
	for cont {
		const retry_count = 3

		retry_delay := time.Duration(10)
		var result_page *openapi.V1reportsVirusDataReportPage
		for i := 1; i <= retry_count; i++ {
			var apiRequest openapi.ApiVirusReportsByPostRequest
			apiRequest.ApiService = cli.VirusApi
			apiRequest.V1VirusDataReportRequest(*request)
			//temporary medium testing workaround
			for key, value := range jsonAcceptHeader {
				apiRequest.Headers = make(map[string]string)
				apiRequest.Headers[key] = value
			}
			page, resp, api_err := apiRequest.Execute()
			err = handleHTTPResponse(resp, api_err)
			if err == nil {
				result_page = &page
				break
			}
			if argDebug {
				fmt.Fprintf(os.Stderr, "Error: '%s' in virus summary paging. Retrying: %d of %d\n", err.Error(), i, retry_count)
			}
			time.Sleep(retry_delay * time.Millisecond)
			retry_delay *= 2
		}
		if countOnly {
			retrievalCount = int(*result_page.TotalCount)
			return retrievalCount, nil
		}

		retrievalCount += len(result_page.GetReports())

		if bar == nil {
			bar = uiprogress.AddBar(int(result_page.GetTotalCount())).AppendCompleted()
			bar.PrependFunc(func(b *uiprogress.Bar) string {
				return p.Sprintf("Collecting %d genome accessions", b.Total)

			})
			bar.AppendFunc(func(b *uiprogress.Bar) string {
				return fmt.Sprintf("%d/%d", b.Current(), b.Total)
			})
			bar.Width = 50
		}
		bar.Set(int(retrievalCount))

		for _, report := range result_page.GetReports() {
			if limit {
				if limitCount >= arg_limit {
					retrievalCount = limitCount
					cont = false
					break
				}
				limitCount += 1
			}
			callbackFn(report)
		}
		if result_page.GetNextPageToken() == "" {
			break
		}
		request.SetPageToken(result_page.GetNextPageToken())

	}

	return retrievalCount, err

}

var summaryVirusGenomeAccCmd = &cobra.Command{
	Use:   "accession",
	Short: "Print coronavirus genome metadata by accession",
	Long: `
Print coronavirus genome metadata (data report) by nucleotide accession. The data report is returned in JSON format.

Refer to NCBI's [command line quickstart](https://www.ncbi.nlm.nih.gov/datasets/docs/quickstarts/command-line-tools/) documentation for information about getting started with the command-line tools.`,
	Example: "  datasets summary virus genome accession NC_045512.2",

	RunE: func(cmd *cobra.Command, args []string) error {
		argIDArgs, err := getArgsFromListOrFile(args, argInputFile)
		err = processAssmAccArgs(args)
		if err != nil {
			return err
		}
		request, err := summaryVirusGenomeRequestWithAcc(argIDArgs)
		if err != nil {
			return err
		}

		return printViralMetadataWithPost(request)
	},
}

func init() {

}
