package flags

import (
	openapi "datasets/openapi/v2"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/thediveo/enumflag"
)

type TaxonomyDownloadIncludeFlags enumflag.Flag

var (
	ArgTaxonomyDownloadIncludeFlags []TaxonomyDownloadIncludeFlags
)

const (
	AuxNames TaxonomyDownloadIncludeFlags = iota
)

var TaxonomyDownloadIncludeFlagIds = map[TaxonomyDownloadIncludeFlags][]string{
	AuxNames: {"names"},
}

var TaxonomyDownloadIncludeFlagOpenapi = map[TaxonomyDownloadIncludeFlags]openapi.V2TaxonomyDatasetRequestTaxonomyReportType{
	AuxNames: openapi.V2TAXONOMYDATASETREQUESTTAXONOMYREPORTTYPE_NAMES_REPORT,
}

type TaxonomyDownloadIncludeFlag struct {
	FlagInterface
	IncludeReports []TaxonomyDownloadIncludeFlags
}

func NewTaxonomyDownloadIncludeFlag() *TaxonomyDownloadIncludeFlag {
	tif := &TaxonomyDownloadIncludeFlag{
		IncludeReports: []TaxonomyDownloadIncludeFlags{},
	}
	return tif
}

func (tif *TaxonomyDownloadIncludeFlag) RegisterFlags(flags *pflag.FlagSet) {
	flags.Var(
		enumflag.NewSlice(&((*tif).IncludeReports), "string", TaxonomyDownloadIncludeFlagIds, enumflag.EnumCaseInsensitive),
		"include",
		`Add report to download:
  * names:          taxonomy names report
  `)
}

func (tif *TaxonomyDownloadIncludeFlag) PreRunE(cmd *cobra.Command, args []string) (err error) {
	return nil
}
