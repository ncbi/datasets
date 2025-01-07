package datasets

import (
	cmdflags "datasets_cli/v2/datasets/flags"
	openapi "datasets/openapi/v2"
	"fmt"

	"github.com/spf13/cobra"
)

func createSummaryGeneTaxonCmd(sGeneFlag *SummaryGeneFlag) *cobra.Command {
	inputDescription := "taxon (NCBI Taxonomy ID, scientific or common name for a species)"

	cmd := &cobra.Command{
		Use:     "taxon <taxon> [flags]",
		Short:   fmt.Sprintf("Print a data report containing gene metadata by %s", inputDescription),
		Long:    fmt.Sprintf(`Print a data report containing gene metadata by %s. The data report is returned in JSON format.`, inputDescription),
		Example: `  datasets summary gene taxon human`,
		Args:    cmdflags.ExpectOnePositionalArgument(inputDescription),
		PreRunE: cmdflags.ExecutePreRunEFor(sGeneFlag.cmdFlagSet),
		RunE: func(cmd *cobra.Command, args []string) error {

			taxId, taxError := RetrieveTaxIdForTaxon(
				args[0],
				false,
				openapi.V2ORGANISMQUERYREQUESTTAXONRESOURCEFILTER_GENE,
				"gene",
			)
			if taxError != nil {
				return taxError
			}

			request := openapi.NewV2GeneDatasetReportsRequest()
			request.SetTaxon(taxId)

			cli, cliErr := createOAClient()
			if cliErr != nil {
				return cliErr
			}
			api := GeneApi{geneApi: cli.GeneAPI}
			return geneSummaryPagePrinter(sGeneFlag, NewDefaultRequestIterator(request), api)
		},
	}

	return cmd
}
