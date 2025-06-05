package flags

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/thediveo/enumflag/v2"
)

type VirusSummaryReportMode enumflag.Flag

// Define the enumeration values for FooMode.
const (
	DEFAULT VirusSummaryReportMode = iota
	ANNOTATION
)

// Map enumeration values to their textual representations (value
// identifiers).
var VirusSummaryReportModeIds = map[VirusSummaryReportMode][]string{
	DEFAULT:    {"virus"},
	ANNOTATION: {"annotation"},
}

var ArgsVirusSummaryReportMode VirusSummaryReportMode

type VirusSummaryReportFlag struct {
	FlagInterface
	virusReport VirusSummaryReportMode
}

func NewVirusSummaryReportFlag() *VirusSummaryReportFlag {
	vsrf := &VirusSummaryReportFlag{
		virusReport: ArgsVirusSummaryReportMode,
	}
	return vsrf
}

// TODO:   *  "biosample": returns a biosample report
const summaryReportFlagLongDesc string = `Specify report virus genome report summary type
  * virus:      returns a primary virus data report
  * annotation: returns a virus annotation report
    `

func (vsrf *VirusSummaryReportFlag) RegisterFlags(flags *pflag.FlagSet) {
	flags.Var(
		enumflag.New(&ArgsVirusSummaryReportMode, "string", VirusSummaryReportModeIds, enumflag.EnumCaseInsensitive),
		"report",
		summaryReportFlagLongDesc)
}

func (vsrf *VirusSummaryReportFlag) PreRunE(cmd *cobra.Command, args []string) (err error) {

	return nil

}
