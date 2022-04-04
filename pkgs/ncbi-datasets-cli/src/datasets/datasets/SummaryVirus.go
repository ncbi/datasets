package datasets

import (
	"github.com/spf13/cobra"
)

// summaryVirusCmd represents the GetSummary command
var summaryVirusCmd = &cobra.Command{
	Use:   "virus",
	Short: "Print coronavirus metadata",
	Long: `
Print coronavirus metadata (data report). The data report is returned in JSON Lines format.

Refer to NCBI's [command line quickstart](https://www.ncbi.nlm.nih.gov/datasets/docs/quickstarts/command-line-tools/) documentation for information about getting started with the command-line tools.`,
}

func init() {

	summaryVirusCmd.AddCommand(summaryVirusGenomeCmd)

}
