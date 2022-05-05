package datasets

import (
	"github.com/spf13/cobra"

	openapi "datasets_cli/v1/openapi"
)

func summaryVirusGenomeRequestWithTax(taxon string) (request *openapi.V1VirusDataReportRequest, err error) {

	filter, err := setBaseFilter()
	filter.SetTaxon(taxon)
	request = openapi.NewV1VirusDataReportRequest()
	request.SetFilter(filter)
	return

}

var summaryVirusGenomeTaxCmd = &cobra.Command{
	Use:   "taxon",
	Short: "Print coronavirus genome metadata by accession",
	Long: `
Print coronavirus genome metadata (data report) by taxa in the Coronaviridae family. The data report is returned in JSON format.

Refer to NCBI's [download and install](https://www.ncbi.nlm.nih.gov/datasets/docs/download-and-install/) documentation for information about getting started with the command-line tools.`,
	Example: "  datasets summary virus genome taxon SARS-COV-2 --host dog",

	Args: cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		processAssmAccArgs(args)
		request, err := summaryVirusGenomeRequestWithTax(args[0])
		if err != nil {
			return err
		}

		return printViralMetadataWithPost(request)
	},
}

func init() {

}
