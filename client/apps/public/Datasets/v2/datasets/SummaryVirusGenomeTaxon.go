package datasets

import (
	cmdflags "datasets_cli/v2/datasets/flags"
	openapi "datasets/openapi/v2"
	"fmt"
	"github.com/spf13/cobra"
)

func createSummaryVirusGenomeTaxCmd(vsf SummaryVirusFlag) *cobra.Command {
	inputDescription := "taxon (NCBI Taxonomy ID, scientific or common name for any virus at any tax rank)"
	iff := cmdflags.NewInputFileFlag(cmdflags.InputFileListTypeTaxon, cmdflags.AsIntegerFalse, cmdflags.WithLimit(MAX_VIRUS_TAXONS))
	flagSets := []cmdflags.FlagInterface{iff}

	cmd := &cobra.Command{
		Use:   "taxon",
		Short: fmt.Sprintf("Print a data report containing virus genome metadata by %s", inputDescription),
		Long: fmt.Sprintf(`
Print a data report containing virus genome metadata by %s. The data report is returned in JSON format

Refer to NCBI's [download and install](https://www.ncbi.nlm.nih.gov/datasets/docs/v2/download-and-install/) documentation for information about getting started with the command-line tools.`, inputDescription),
		Example: "  datasets summary virus genome taxon sars-cov-2 --host dog",
		PreRunE: cmdflags.ExecutePreRunEFor(flagSets),

		RunE: func(cmd *cobra.Command, args []string) (err error) {
			if err != nil {
				return err
			}
			var taxIdsMap, taxErr = RetrieveTaxIdsForTaxons(cmd, iff.InputIDArgs, true, openapi.V2ORGANISMQUERYREQUESTTAXONRESOURCEFILTER_ALL, "virus", 10239)
			if taxErr != nil {
				return taxErr
			}

			err = executeSummaryVirusGenomeCmd(getMapListValues(taxIdsMap), vsf, nil)

			return err

		},
	}
	cmdflags.RegisterAllFlags(flagSets, cmd.PersistentFlags())

	return cmd
}
