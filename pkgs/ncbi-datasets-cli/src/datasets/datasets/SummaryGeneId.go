package datasets

import (
	"github.com/spf13/cobra"
	openapi "main/openapi_client"
)

func cmdRunSummaryGeneID(cmd *cobra.Command, args []string) (err error) {
	geneInts := strToInt64List(argIDArgs)

	req := new(openapi.V1alpha1GeneDatasetRequest)
	req.GeneIds = geneInts
	if argJsonLinesFormat {
		err = streamGeneMatch(req, &JsonLinesStreamProcessor{})
	} else {
		err = streamGeneMatch(req, &JsonStreamProcessor{wrapperName: "genes"})
	}
	return
}

var summaryGeneIDCmd = &cobra.Command{
	Use:   "gene-id <gene-id ...> [flags]",
	Short: "print a summary of a gene dataset by NCBI Gene ID",
	Long:  `Print a summary of a gene dataset by NCBI Gene ID. The summary is returned in JSON format.`,
	Example: `  datasets summary gene gene-id 672
  datasets summary gene gene-id 2597 14433`,
	RunE: cmdRunSummaryGeneID,
}

func init() {
	flags := summaryGeneIDCmd.Flags()
	registerHiddenStringPair(flags, &argInputFile, "inputfile", "i", "", "read a list of NCBI Gene IDs from a file to use as input")
}
