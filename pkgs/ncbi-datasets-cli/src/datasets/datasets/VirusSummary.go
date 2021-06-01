package datasets

import (
	"os"

	"github.com/antihax/optional"
	"github.com/metakeule/fmtdate"
	"github.com/spf13/cobra"

	openapi "main/openapi_client"
)

func virusGenomeSummaryOpts(cmd *cobra.Command) *openapi.VirusGenomeSummaryOpts {
	opts := new(openapi.VirusGenomeSummaryOpts)
	opts.RefseqOnly = optional.NewBool(argRefseqOnly)
	opts.AnnotatedOnly = optional.NewBool(argAnnotatedOnly)
	if argReleasedSince != "" {
		date, err := fmtdate.Parse(dateFormat, argReleasedSince)
		if err != nil {
			cmd.PrintErr("\nError: ", err, "\n")
			os.Exit(1)
		}
		opts.ReleasedSince = optional.NewTime(date)
	}
	opts.Host = optional.NewString(argHost)
	return opts
}

func getVirusSummaryByName(opts *openapi.VirusGenomeSummaryOpts, taxon string, cmd *cobra.Command) (err error) {
	cli, err := createOAClient()
	if err != nil {
		return
	}
	result, resp, err := cli.VirusApi.VirusGenomeSummary(nil, taxon, opts)
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
		opts := virusGenomeSummaryOpts(cmd)
		taxon := args[0]
		return getVirusSummaryByName(opts, taxon, cmd)
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
