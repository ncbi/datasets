package datasets

import (
	cmdflags "datasets_cli/v2/datasets/flags"
	openapi "datasets/openapi/v2"

	"github.com/spf13/cobra"
)

type SummaryVirusFlag struct {
	virusFilterFlags *cmdflags.VirusFilterFlags

	virusSumRptFlag    *cmdflags.VirusSummaryReportFlag
	jsonLinesLimitFlag *cmdflags.JsonLinesAndLimitFlag
	cmdFlagSet         []cmdflags.FlagInterface
}

func initSummaryVirusGenomeFlag() SummaryVirusFlag {

	vsrf := cmdflags.NewVirusSummaryReportFlag()
	jll := cmdflags.NewJsonLineAndLimitFlag("virus")
	vff := cmdflags.NewVirusFilterFlags(cmdflags.GenomeReleasedAfterDesc, cmdflags.VirusGenomeUpdatedAfterDesc)

	svf := SummaryVirusFlag{

		virusSumRptFlag:    vsrf,
		jsonLinesLimitFlag: jll,
		virusFilterFlags:   vff,
		cmdFlagSet:         []cmdflags.FlagInterface{vsrf, jll, vff},
	}

	return svf
}

func (dvf *SummaryVirusFlag) PrepareAnnotationReportRequest(accs []string, taxons []string) *openapi.V2VirusAnnotationReportRequest {
	return dvf.virusFilterFlags.PrepareAnnotationReportRequest(accs, taxons)
}

func (dvf *SummaryVirusFlag) PrepareDatasetReportRequest(accs []string, taxons []string) *openapi.V2VirusDataReportRequest {
	return dvf.virusFilterFlags.PrepareDatasetReportRequest(accs, taxons)
}

func createSummaryVirusGenomeCmd() *cobra.Command {

	summaryVirusFlag := initSummaryVirusGenomeFlag()

	cmd := &cobra.Command{

		Use:   "genome",
		Short: "Print a data report containing virus genome metadata by accession or taxon",
		Long: `
Print a data report containing virus genome metadata by nucleotide accession or taxon. The data report is returned in JSON format.

Refer to NCBI's [download and install](https://www.ncbi.nlm.nih.gov/datasets/docs/v2/download-and-install/) documentation for information about getting started with the command-line tools.`,
		Example: `  datasets summary virus genome accession NC_045512.2
	  datasets summary virus genome taxon sars-cov-2 --host dog`,
		Args:              cobra.NoArgs,
		PersistentPreRunE: cmdflags.PersistentPreRunEFor(summaryVirusFlag.cmdFlagSet, rootCmd),
		RunE:              ParentCommandRunE,
	}

	cmdflags.RegisterAllFlags(summaryVirusFlag.cmdFlagSet, cmd.PersistentFlags())

	cmd.AddCommand(createSummaryVirusGenomeTaxCmd(summaryVirusFlag))
	cmd.AddCommand(createSummaryVirusGenomeAccCmd(summaryVirusFlag))

	return cmd

}
