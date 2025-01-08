package flags

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type AnnotatedOnlyFlag struct {
	FlagInterface
	annotatedOnly bool
}

func NewAnnotatedOnlyFlag() *AnnotatedOnlyFlag {
	return &AnnotatedOnlyFlag{annotatedOnly: false}
}

func (ao *AnnotatedOnlyFlag) RegisterFlags(flags *pflag.FlagSet) {
	flags.BoolVar(&ao.annotatedOnly, "annotated", false, "Limit to annotated genomes")
}

func (ao *AnnotatedOnlyFlag) PreRunE(cmd *cobra.Command, args []string) (err error) {
	return nil
}

func (ao *AnnotatedOnlyFlag) AnnotatedOnly() bool {
	return ao.annotatedOnly
}
