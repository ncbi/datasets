package flags

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type TaxParentsFlag struct {
	FlagInterface
	taxParents bool
}

func NewTaxParentsFlag() *TaxParentsFlag {
	return &TaxParentsFlag{taxParents: false}
}

func (tpf *TaxParentsFlag) RegisterFlags(flags *pflag.FlagSet) {
	flags.BoolVar(&tpf.taxParents, "parents", false, "Include all parents of the requested taxon")
}

func (tpf *TaxParentsFlag) PreRunE(cmd *cobra.Command, args []string) (err error) {
	return nil
}

func (tpf *TaxParentsFlag) GetParents() bool {
	return tpf.taxParents
}

func (tpf *TaxParentsFlag) SetParents(val bool) {
	tpf.taxParents = val
}
