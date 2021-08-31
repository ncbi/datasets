package datasets

import (
	"github.com/metakeule/fmtdate"
	"github.com/spf13/cobra"
	"time"
)

func getVirusSummaryByName(cmd *cobra.Command, taxon string) (err error) {
	cli, err := createOAClient()
	if err != nil {
		return
	}
	request := cli.VirusApi.VirusGenomeSummary(nil, taxon)
	request.RefseqOnly(argRefseqOnly)
	request.AnnotatedOnly(argAnnotatedOnly)
	if argReleasedSince != "" {
		date := time.Time{}
		date, err = fmtdate.Parse(dateFormat, argReleasedSince)
		if err != nil {
			return
		}
		request.ReleasedSince(date)
	}
	request.Host(argHost)

	result, resp, err := request.Execute()
	if err = handleHTTPResponse(resp, err); err != nil {
		return
	}
	printResults(&result)
	return
}

// virusSummaryCmd represents the GetSummary command
var virusSummaryOrgCmd = &cobra.Command{
	Use:   "taxon <taxonomic-name>",
	Short: "Retrieve summary of virus genomes by taxonomic name (taxonomy id, scientific or common name at any tax rank)",
	Example: `  datasets virus-summary taxon sars2
  datasets virus-summary taxon sars2 --refseq
  datasets virus-summary taxon Coronaviridae --released-since 05/05/2021`,
	Long: `
Retrieve a summary of virus genomes by taxonomic name (tax-id, scientific or common name at any tax rank).

Refer to NCBI's [command line quickstart](https://www.ncbi.nlm.nih.gov/datasets/docs/quickstarts/command-line-tools/) documentation for information about getting started with the command-line tools.`,
	Args: cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		taxon := args[0]
		return getVirusSummaryByName(cmd, taxon)
	},
}

// virusSummaryCmd represents the GetSummary command
var virusSummaryCmd = &cobra.Command{
	Deprecated: "will remove",
	Use:        "virus-summary",
	Short:      "Retrieve summary of available virus genome datasets",
	Long: `
Retrieve summary of available virus genomes.

Refer to NCBI's [command line quickstart](https://www.ncbi.nlm.nih.gov/datasets/docs/quickstarts/command-line-tools/) documentation for information about getting started with the command-line tools.`,
}

func init() {
	virusSummaryCmd.AddCommand(virusSummaryOrgCmd)

	virusSummaryCmd.PersistentFlags().BoolVarP(&argRefseqOnly, "refseq", "", false, "Only include RefSeq genomes")
	virusSummaryCmd.PersistentFlags().BoolVarP(&argAnnotatedOnly, "annotated", "a", false, "Only include viral genomes with annotation")
	virusSummaryCmd.PersistentFlags().StringVarP(&argReleasedSince, "released-since", "", "", "("+dateFormat+") Only include viral genomes that have been released after a specified date")
	virusSummaryCmd.PersistentFlags().StringVarP(&argHost, "host", "", "", "Only include genomes isolated from the specified host. Specify tax-name or tax-id at any taxonomic rank.")
}
