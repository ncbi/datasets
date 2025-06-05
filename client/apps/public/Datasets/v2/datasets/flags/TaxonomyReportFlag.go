package flags

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/thediveo/enumflag/v2"
)

type TaxonomyReportMode enumflag.Flag

// Define the enumeration values for ReportMode.
const (
	Taxonomy TaxonomyReportMode = iota
	Names
	TaxIdsOnly
)

var TaxonomyReportModeIds = map[TaxonomyReportMode][]string{
	Taxonomy:   {"taxonomy"},
	Names:      {"names"},
	TaxIdsOnly: {"ids_only"},
}

type TaxonomyReportFlag struct {
	FlagInterface
	taxonomyReport TaxonomyReportMode
}

func NewTaxonomyReportFlag() *TaxonomyReportFlag {
	trf := &TaxonomyReportFlag{
		taxonomyReport: Taxonomy,
	}
	return trf
}

func (trf *TaxonomyReportFlag) TaxonomyReport() bool {
	return trf.taxonomyReport == Taxonomy
}

func (trf *TaxonomyReportFlag) NamesReport() bool {
	return trf.taxonomyReport == Names
}

func (trf *TaxonomyReportFlag) IdsOnly() bool {
	return trf.taxonomyReport == TaxIdsOnly
}

func (trf *TaxonomyReportFlag) RegisterFlags(flags *pflag.FlagSet) {
	flags.Var(
		enumflag.New(&trf.taxonomyReport, "string", TaxonomyReportModeIds, enumflag.EnumCaseInsensitive),
		"report",
		`Choose the output type:
  * taxonomy:   Retrieve the primary taxonomy report
  * names: Retrieve the taxonomy names report
  * ids_only: Retrieve only the taxonomy identifiers
  `)
}

func (trf *TaxonomyReportFlag) PreRunE(cmd *cobra.Command, args []string) (err error) {
	return nil
}
