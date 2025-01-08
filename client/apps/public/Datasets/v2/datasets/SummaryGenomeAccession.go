package datasets

import (
	cmdflags "datasets_cli/v2/datasets/flags"

	"github.com/spf13/cobra"
)

func createSummaryGenomeAccessionCmd(sgf SummaryGenomeFlag, assemblyRequestFlag AssemblyRequestFlag) *cobra.Command {
	iff := cmdflags.NewInputFileFlag(cmdflags.InputFileListTypeGenomeAcc, cmdflags.AsIntegerFalse)
	flagSets := []cmdflags.FlagInterface{iff}

	cmd := &cobra.Command{
		Use:   "accession <accession ...>",
		Short: "Print a data report containing assembled genome metadata by Assembly or BioProject accession",
		Long: `
Print a data report containing assembled genome metadata by Assembly or BioProject accession. The data report is returned in JSON format.`,
		Example: `  datasets summary genome accession GCF_000001405.40
  datasets summary genome accession GCA_003774525.2 GCA_000001635
  datasets summary genome accession GCF_000001405.40 --report sequence --as-json-lines
  datasets summary genome accession PRJNA31257`,

		PreRunE: cmdflags.ExecutePreRunEFor(flagSets),

		RunE: func(cmd *cobra.Command, args []string) error {
			// This could potentially change avf.AssemblyVersion
			accessions := assemblyRequestFlag.assemblyVersionFlag.UpdateForInputAccessions(iff.InputIDArgs)
			request, err := GetGenomeReportsAccessionRequest(accessions)
			if err != nil {
				return err
			}
			return getGenomeSummary(NewGenomeAccessionRequestIter(request), sgf, assemblyRequestFlag)
		},
	}

	cmdflags.RegisterAllFlags(flagSets, cmd.PersistentFlags())

	return cmd
}
