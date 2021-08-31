package datasets

import (
	"github.com/spf13/cobra"

	openapi "main/openapi_client"
)

var summaryGenomeTaxonCmd = &cobra.Command{
	Use:   "taxon",
	Short: "print a summary of a genome dataset by taxon (NCBI Taxonomy ID, scientific or common name at any tax rank)",
	Long: `
Print a summary of a genome dataset by taxon (NCBI Taxonomy ID, scientific or common name at any tax rank). The summary is returned in JSON format.

Refer to NCBI's [command line quickstart](https://www.ncbi.nlm.nih.gov/datasets/docs/quickstarts/command-line-tools/) documentation for information about getting started with the command-line tools.`,
	Example: `  datasets summary genome taxon human
  datasets summary genome taxon "mus musculus"
  datasets summary genome taxon 10116`,

	Args: cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) (err error) {
		request := openapi.NewV1AssemblyMetadataRequest()
		request.SetTaxon(args[0])

		err = updateAssemblyMetadataRequestOption(request)
		if err != nil {
			return err
		}

		result, metadata_err := getAssemblyMetadataWithPost(request, false)
		if metadata_err != nil {
			return metadata_err
		}
		printResults(&result)
		return nil
	},
}

func init() {
	summaryGenomeTaxonCmd.PersistentFlags().BoolVar(&argTaxExactMatch, "tax-exact-match", false, "exclude sub-species when a species-level taxon is specified")
}
