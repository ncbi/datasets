package flags

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type TaxExactMatch struct {
	FlagInterface
	taxExactMatch bool
}

func NewTaxExactMatchFlag() *TaxExactMatch {
	return &TaxExactMatch{taxExactMatch: false}
}

func (sv *TaxExactMatch) RegisterFlags(flags *pflag.FlagSet) {
	flags.BoolVar(&sv.taxExactMatch, "tax-exact-match", false, "Exclude sub-species when a species-level taxon is specified")
}

func (lf *TaxExactMatch) PreRunE(cmd *cobra.Command, args []string) (err error) {
	return nil
}

func (sv *TaxExactMatch) IsTaxExactMatch() bool {
	return sv.taxExactMatch
}
