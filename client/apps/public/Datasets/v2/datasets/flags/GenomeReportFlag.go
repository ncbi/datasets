package flags

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/thediveo/enumflag/v2"
)

type GenomeReportMode enumflag.Flag

// Define the enumeration values for ReportMode.
const (
	Genome GenomeReportMode = iota
	Sequence
	IdsOnly
)

var GenomeReportModeIds = map[GenomeReportMode][]string{
	Genome:   {"genome"},
	Sequence: {"sequence"},
	IdsOnly:  {"ids_only"},
}

type GenomeReportFlag struct {
	FlagInterface
	genomeReport GenomeReportMode
}

func NewGenomeReportFlag() *GenomeReportFlag {
	grf := &GenomeReportFlag{
		genomeReport: Genome,
	}
	return grf
}

func (grf *GenomeReportFlag) GenomeReport() bool {
	return grf.genomeReport == Genome
}

func (grf *GenomeReportFlag) SequenceReport() bool {
	return grf.genomeReport == Sequence
}

func (grf *GenomeReportFlag) IdsOnly() bool {
	return grf.genomeReport == IdsOnly
}

func (grf *GenomeReportFlag) RegisterFlags(flags *pflag.FlagSet) {
	flags.Var(
		enumflag.New(&grf.genomeReport, "string", GenomeReportModeIds, enumflag.EnumCaseInsensitive),
		"report",
		`Choose the output type:
  * genome:   Retrieve the primary genome report
  * sequence: Retrieve the sequence report
  * ids_only: Retrieve only the genome identifiers
  `)
}

func (grf *GenomeReportFlag) PreRunE(cmd *cobra.Command, args []string) (err error) {
	return nil
}
