package datasets

import (
	cmdflags "datasets_cli/v2/datasets/flags"

	"github.com/spf13/cobra"
)

func createSummaryVirusCmd() *cobra.Command {
	flagSets := []cmdflags.FlagInterface{}

	cmd := &cobra.Command{
		Use:   "virus",
		Short: "Print a data report containing virus genome metadata",
		Long: `
Print a data report containing virus genome metadata by accession or taxon. The data report is returned in JSON format.`,
		Args: cobra.NoArgs,
		RunE: ParentCommandRunE,
	}

	cmdflags.RegisterAllFlags(flagSets, cmd.Flags())

	cmd.AddCommand(createSummaryVirusGenomeCmd())

	return cmd
}
