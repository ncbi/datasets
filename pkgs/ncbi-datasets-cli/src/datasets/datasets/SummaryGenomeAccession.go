package datasets

import (
	"errors"
	"fmt"
	"math"
	_nethttp "net/http"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"

	openapi "main/openapi_client"
)

func updateAssemblyMetadataRequestOption(request *openapi.V1alpha1AssemblyMetadataRequest) (err error) {
	request.Filters.ReferenceOnly = argReferenceOnly
	request.Filters.SearchText = argSearchText
	request.Filters.HasAnnotation = argAnnotatedGenomesOnly
	AssemblySource, source_err := getAssemblySource()
	if source_err != nil {
		return source_err
	} else {
		request.Filters.AssemblySource = openapi.AssemblyDatasetDescriptorsFilterAssemblySource(AssemblySource)
	}
	request.Filters.AssemblyLevel, err = getAssemblyLevels()
	if err != nil {
		return
	}
	date_since, date_since_err := getDate(argGenomeReleasedSince)
	if date_since_err != nil {
		return date_since_err
	} else {
		if date_since.IsSet() {
			request.Filters.FirstReleaseDate = date_since.Value()
		}
	}
	date_until, date_until_err := getDate(argGenomeReleasedUntil)
	if date_until_err != nil {
		return date_until_err
	} else {
		if date_until.IsSet() {
			request.Filters.LastReleaseDate = date_until.Value()
		}
	}
	request.PageSize = 1000

	request.TaxExactMatch = argTaxExactMatch

	if argAssmAccsOnly {
		request.ReturnedContent = openapi.V1ALPHA1ASSEMBLYMETADATAREQUESTCONTENTTYPE_ASSM_ACC
	}

	return nil

}

func getAssemblyMetadataPage(cli *openapi.APIClient, request *openapi.V1alpha1AssemblyMetadataRequest, page_size int32) (openapi.V1alpha1AssemblyMetadata, *_nethttp.Response, error) {
	request.PageSize = page_size
	return cli.GenomeApi.GenomeMetadataByPost(nil, *request)
}

func SetPageToken(request *openapi.V1alpha1AssemblyMetadataRequest, token string) {
	request.PageToken = token
}

func getAssemblyMetadataWithPost(request *openapi.V1alpha1AssemblyMetadataRequest, ignore_limit bool) (result openapi.V1alpha1AssemblyMetadata, err error) {
	cli, cli_err := createOAClient()
	if cli_err != nil {
		return result, cli_err
	}

	maxRetrieval := math.MaxInt32
	countOnly := false
	if !ignore_limit {
		if argLimit != "" {
			if argLimit == "none" {
				maxRetrieval = 1
				countOnly = true
			} else if argLimit != "all" {
				maxRetrieval, err = strconv.Atoi(argLimit)
				if err != nil {
					err = fmt.Errorf("Invalid 'limit' value %s. Must be 'all', 'none', or a number.", argLimit)
					return
				}
			}
		}
	}

	retrievalCount := 0
	for {
		page_size := int32(minOf(MAX_PAGE_SIZE, maxRetrieval-retrievalCount))

		var retry_delay = time.Duration(10)

		const retry_count = 3
		var result_page *openapi.V1alpha1AssemblyMetadata
		for i := 1; i <= retry_count; i++ {
			page, resp, api_err := getAssemblyMetadataPage(cli, request, page_size)
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

		if err != nil {
			result = openapi.V1alpha1AssemblyMetadata{}
			return
		}
		if result.TotalCount == 0 {
			result.TotalCount = result_page.TotalCount
		}
		if countOnly {
			return
		}

		retrievalCount += len(result_page.Assemblies)
		result.Assemblies = append(result.Assemblies, result_page.Assemblies...)
		result.Messages = append(result.Messages, result_page.Messages...)
		if result_page.NextPageToken == "" || retrievalCount >= maxRetrieval {
			break
		}
		SetPageToken(request, result_page.NextPageToken)
	}
	return
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
		request := new(openapi.V1alpha1AssemblyMetadataRequest)
		request.Accessions.Accessions = accs
		err = updateAssemblyMetadataRequestOption(request)
		if err != nil {
			return err
		}

		result, err := getAssemblyMetadataWithPost(request, false)
		if err != nil {
			return err
		}
		printResults(&result)
		return nil
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
