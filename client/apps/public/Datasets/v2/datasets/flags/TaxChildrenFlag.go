package flags

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type TaxChildrenFlag struct {
	FlagInterface
	taxChildren bool
}

func NewTaxChildrenFlag() *TaxChildrenFlag {
	return &TaxChildrenFlag{taxChildren: false}
}

func (tcf *TaxChildrenFlag) RegisterFlags(flags *pflag.FlagSet) {
	flags.BoolVar(&tcf.taxChildren, "children", false, "Return all taxonomic children of the requested taxon")
}

func (tcf *TaxChildrenFlag) PreRunE(cmd *cobra.Command, args []string) (err error) {
	return nil
}

func (tcf *TaxChildrenFlag) GetChildren() bool {
	return tcf.taxChildren
}

func (tcf *TaxChildrenFlag) SetChildren(val bool) {
	tcf.taxChildren = val
}
