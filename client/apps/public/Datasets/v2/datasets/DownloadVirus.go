package datasets

import (
	"github.com/spf13/cobra"
)

func createVirusCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "virus",
		Short: "Download a virus data package",
		Long:  "Download a virus genome or SARS-CoV-2 protein data package as a zip file.",
		Example: `  datasets download virus genome taxon sars-cov-2 --host dog
  datasets download virus protein S --host dog --filename SARS2-spike-dog.zip`,
		Args: cobra.NoArgs,
		RunE: ParentCommandRunE,
	}

	cmd.AddCommand(createDownloadVirusGenomeCmd())
	cmd.AddCommand(createDownloadVirusProteinCmd())

	return cmd
}
