package datasets

import (
	"errors"
	"fmt"
	_nethttp "net/http"
	"os"
	"time"

	"github.com/gosuri/uiprogress"
	"github.com/spf13/cobra"
	"golang.org/x/text/language"
	"golang.org/x/text/message"

	openapi "datasets_cli/v1/openapi"
	datasets_util "datasets_cli/v1/util"
)

func updateAssemblyMetadataRequestOption(request *openapi.V1AssemblyMetadataRequest) (err error) {
	filter := openapi.NewV1AssemblyDatasetDescriptorsFilter()
	filter.SetReferenceOnly(argReferenceOnly)
	filter.SetSearchText(argSearchText)
	filter.SetHasAnnotation(argAnnotatedGenomesOnly)
	AssemblySource, source_err := getAssemblySource()
	if source_err != nil {
		return source_err
	} else {
		filter.SetAssemblySource(openapi.V1AssemblyDatasetDescriptorsFilterAssemblySource(AssemblySource))
	}
	assembly_level := []openapi.V1AssemblyDatasetDescriptorsFilterAssemblyLevel{}
	assembly_level, err = getAssemblyLevels()
	if err != nil {
		return
	}
	filter.SetAssemblyLevel(assembly_level)
	date_since, date_since_err := getDate(argGenomeReleasedSince)
	if date_since_err != nil {
		return date_since_err
	} else {
		if date_since.IsSet() {
			filter.SetFirstReleaseDate(date_since.Value())
		}
	}
	date_until, date_until_err := getDate(argGenomeReleasedUntil)
	if date_until_err != nil {
		return date_until_err
	} else {
		if date_until.IsSet() {
			filter.SetLastReleaseDate(date_until.Value())
		}
	}
	request.SetFilters(*filter)
	request.SetPageSize(1000)

	request.SetTaxExactMatch(argTaxExactMatch)

	if argAssmAccsOnly {
		request.SetReturnedContent(openapi.V1ASSEMBLYMETADATAREQUESTCONTENTTYPE_ASSM_ACC)
	}

	return nil
}

func getAssemblyMetadataPage(cli *openapi.APIClient, request *openapi.V1AssemblyMetadataRequest) (openapi.V1AssemblyMetadata, *_nethttp.Response, error) {
	return cli.GenomeApi.GenomeMetadataByPost(nil, request).Execute()
}

func SetPageToken(request *openapi.V1AssemblyMetadataRequest, token string) {
	request.SetPageToken(token)
}

func getAssemblyAccessionsWithPost(request *openapi.V1AssemblyMetadataRequest) ([]string, error) {
	var accessions []string
	getAccessions := func(page openapi.V1AssemblyMatch) {
		accessions = append(accessions, *page.Assembly.AssemblyAccession)
	}

	_, err := getAssemblyMetadataWithPost(request, getAccessions)
	return accessions, err
}

func printAssemblyMetadataWithPost(request *openapi.V1AssemblyMetadataRequest) error {
	var assm_count int
	var header_printed bool
	var err error
	var total_count int

	printAssembly := func(page openapi.V1AssemblyMatch) {
		if !argJsonLinesFormat {
			if !header_printed {
				fmt.Printf("\"assemblies\": [")
				header_printed = true
			}
			if assm_count > 0 {
				fmt.Printf(",")
			}
			fmt.Printf("{\"assembly\": ")
			assm_count++
			defer func() {
				fmt.Printf("}")
			}()
			printResultsNoNewline(page.GetAssembly())
		} else {
			printResults(page.GetAssembly())
		}
	}

	if !argJsonLinesFormat {
		fmt.Printf("{")
		defer func() {
			if header_printed {
				fmt.Printf("],")
			}
			fmt.Printf("\"total_count\": %d}", total_count)
		}()
	}
	total_count, err = getAssemblyMetadataWithPost(request, printAssembly)
	return err
}

func getAssemblyMetadataWithPost(request *openapi.V1AssemblyMetadataRequest, callbackFn func(openapi.V1AssemblyMatch)) (int, error) {
	cli, err := createOAClient()
	if err != nil {
		return 0, err
	}

	countOnly := false
	if argLimit != "" {
		if argLimit == "none" {
			request.SetPageSize(1)
			countOnly = true
		} else if argLimit != "all" {
			// arg_limit, err := strconv.Atoi(argLimit)
			// if err != nil {
			// 	err = fmt.Errorf("Invalid 'limit' value %s. Must be 'all', 'none', or a number.", argLimit)
			// 	return 0, err
			// }
			request.SetPageSize(1000)
		}
	}

	var bar *uiprogress.Bar
	p := message.NewPrinter(language.English)

	retrievalCount := 0
	for {
		var retry_delay = time.Duration(10)

		const retry_count = 3
		var result_page *openapi.V1AssemblyMetadata
		for i := 1; i <= retry_count; i++ {
			page, resp, api_err := getAssemblyMetadataPage(cli, request)
			err = handleHTTPResponse(resp, api_err)
			if err == nil {
				result_page = &page
				break
			}
			if argDebug {
				fmt.Fprintf(os.Stderr, "Error: '%s' in genome summary paging. Retrying: %d of %d\n", err.Error(), i, retry_count)
			}
			time.Sleep(retry_delay * time.Millisecond)
			retry_delay *= 2
		}

		if result_page.HasMessages() {
			err = datasets_util.MessagesToError(result_page.GetMessages())
		}
		if err != nil {
			return 0, err
		}

		if countOnly {
			retrievalCount = int(*result_page.TotalCount)
			return retrievalCount, nil
		}

		retrievalCount += len(result_page.GetAssemblies())

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

		for _, pg := range result_page.GetAssemblies() {
			callbackFn(pg)
		}
		if result_page.GetNextPageToken() == "" {
			break
		}
		SetPageToken(request, result_page.GetNextPageToken())
	}
	return retrievalCount, err
}

var summaryGenomeAccessionCmd = &cobra.Command{
	Use:   "accession <accession ...>",
	Short: "print a summary of a genome dataset by assembly accession",
	Long: `
Print a summary of a genome dataset by NCBI Assembly or BioProject accession. The summary is returned in JSON format.

Refer to NCBI's [command line quickstart](https://www.ncbi.nlm.nih.gov/datasets/docs/quickstarts/command-line-tools/) documentation for information about getting started with the command-line tools.`,
	Example: `  datasets summary genome accession GCF_000001405.39
  datasets summary genome accession GCA_003774525.2 GCA_000001635
  datasets summary genome accession PRJNA31257`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := processAssmAccArgs(args)
		if err != nil {
			return err
		}

		cli, err := createOAClient()
		if err != nil {
			return err
		}

		accs, err := gencollAccessionsForMixedAccessions(cli, argIDArgs)
		if err != nil {
			return err
		}
		if len(accs) == 0 {
			return errors.New("no assemblies found for specified accessions")
		}
		request := openapi.NewV1AssemblyMetadataRequest()
		accessions := openapi.NewV1Accessions()
		accessions.SetAccessions(accs)
		request.SetAccessions(*accessions)

		err = updateAssemblyMetadataRequestOption(request)
		if err != nil {
			return err
		}

		return printAssemblyMetadataWithPost(request)
	},
}

func init() {
	registerHiddenStringPair(
		summaryGenomeAccessionCmd.PersistentFlags(),
		&argInputFile,
		"inputfile",
		"i",
		"",
		"read a list of NCBI Assembly or BioProject accessions from a file to use as input",
	)
}
