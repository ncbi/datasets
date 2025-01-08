package flags

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type WpDownloadFlags struct {
	FlagInterface
	includeFlanksBp int
	taxonFilter     string
}

func NewWpDownloadFlags() *WpDownloadFlags {
	wdf := &WpDownloadFlags{}
	return wdf
}

func (wdf *WpDownloadFlags) RegisterFlags(flags *pflag.FlagSet) {
	flags.IntVar(&wdf.includeFlanksBp, "include-flanks-bp", 0, "Specify the length of flanking nucleotides (WP accessions only)")
	flags.StringVar(&wdf.taxonFilter, "taxon-filter", "", "Limit gene sequences and annotation report file to specified taxon (any rank, only available for WP accessions)")
}

func (wdf *WpDownloadFlags) PreRunE(cmd *cobra.Command, args []string) (err error) {
	return
}

func (wdf *WpDownloadFlags) IncludeFlankBp() int {
	return wdf.includeFlanksBp
}

func (wdf *WpDownloadFlags) TaxonFilter() string {
	return wdf.taxonFilter
}
