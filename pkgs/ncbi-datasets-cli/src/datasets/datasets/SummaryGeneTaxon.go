package datasets

import (
	openapi "main/openapi_client"

	"github.com/spf13/cobra"
)

func cmdRunSummaryGeneTaxon(cmd *cobra.Command, args []string) (err error) {

	req := openapi.NewV1GeneDatasetRequest()
	req.SetTaxon(argIDArgs[0])
	if argJsonLinesFormat {
		err = streamGeneMatch(req, &JsonLinesStreamProcessor{})
	} else {
		err = streamGeneMatch(req, &JsonStreamProcessor{wrapperName: "genes"})
	}
	return
}

var summaryGeneTaxonCmd = &cobra.Command{
	Use:     "taxon <taxon> [flags]",
	Short:   "print a summary of a gene dataset by taxon",
	Long:    `Print a summary of a gene dataset by taxon. The summary is returned in JSON format.`,
	Example: `  datasets summary gene taxon human`,
	Args:    cobra.MaximumNArgs(1),
	RunE:    cmdRunSummaryGeneTaxon,
}
