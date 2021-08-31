package datasets

import (
	"errors"

	"github.com/spf13/cobra"

	openapi "main/openapi_client"
)

type streamProcessorOrtholog interface {
	insertRecordBegin()
	insertRecordEnd()
	insertRecordDelimiter()
}

func cmdRunSummaryOrthologID(cmd *cobra.Command, args []string) (err error) {
	geneInts, err := strToInt32ListErr(argIDArgs)
	if err != nil {
		return
	}
	if len(geneInts) == 0 {
		err = errors.New("No genes found for search term")
		return
	}
	if !argJsonFormat {
		err = streamOrthologs(geneInts, openapi.V1ORTHOLOGREQUESTCONTENTTYPE_COMPLETE, &JsonLinesStreamProcessor{})
	} else {
		err = streamOrthologs(geneInts, openapi.V1ORTHOLOGREQUESTCONTENTTYPE_COMPLETE, &JsonStreamProcessor{wrapperName: "orthologs"})
	}

	return
}

var summaryOrthologIDCmd = &cobra.Command{
	Use:     "gene-id <gene-id>",
	Short:   "print a summary of an ortholog dataset by NCBI Gene ID",
	Long:    `Print a summary of an ortholog dataset by NCBI Gene ID. Ortholog data is calculated by NCBI for vertebrates and insects. The summary is returned in JSON format.`,
	Example: `  datasets summary ortholog gene-id 672`,
	RunE:    cmdRunSummaryOrthologID,
}

func init() {
	flags := summaryOrthologIDCmd.Flags()
	registerHiddenStringPair(flags, &argInputFile, "inputfile", "i", "", "read a list of NCBI Gene IDs from a file to use as input")
	pflags := summaryOrthologIDCmd.PersistentFlags()
	registerHiddenStringSlicePair(pflags, &taxonFilter, "taxon-filter", "", []string{}, "limit results to ortholog data for a specified taxonomic group")
}
