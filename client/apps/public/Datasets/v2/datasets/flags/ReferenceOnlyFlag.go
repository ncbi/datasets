package flags

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type ReferenceOnlyFlag struct {
	FlagInterface
	referenceOnly bool
}

func NewReferenceOnlyFlag() *ReferenceOnlyFlag {
	return &ReferenceOnlyFlag{referenceOnly: false}
}

func (ro *ReferenceOnlyFlag) RegisterFlags(flags *pflag.FlagSet) {
	flags.BoolVar(&ro.referenceOnly, "reference", false, "Limit to reference genomes")
}

func (ro *ReferenceOnlyFlag) PreRunE(cmd *cobra.Command, args []string) (err error) {
	return nil
}

func (ro *ReferenceOnlyFlag) ReferenceOnly() bool {
	return ro.referenceOnly
}
