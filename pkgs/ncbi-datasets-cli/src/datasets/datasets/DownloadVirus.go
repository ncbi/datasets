package datasets

import (
	"github.com/spf13/cobra"
)

var (
	argAnnotatedOnly bool
	argReleasedSince string
	argHost          string
	argLineage       string
	argGeoLocation   string
	argCompleteOnly  bool
	argExcludeCds    bool
	argExcludeGpff   bool
	argExcludePdb    bool
)

const dateFormat = "MM/DD/YYYY"

// virusCmd represents the virus command
var virusCmd = &cobra.Command{
	Use:   "virus [command]",
	Short: "download a coronavirus dataset",
	Long: `
Download a coronavirus genome or SARS-CoV-2 protein dataset as a zip file.

Refer to NCBI's [command line quickstart](https://www.ncbi.nlm.nih.gov/datasets/docs/quickstarts/command-line-tools/) documentation for information about getting started with the command-line tools. `,
	Example: `  datasets download virus genome taxon sars-cov-2 --released-since 05/05/2021
  datasets download virus protein S --released-since 05/05/2021 --filename SARS2-spike-all.zip`,
}

func init() {
	virusCmd.AddCommand(downloadVirusGenomeCmd)
	virusCmd.AddCommand(downloadVirusProteinCmd)
}
