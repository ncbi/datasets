package flags

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type ChromosomesFlag struct {
	FlagInterface
	chromosomes []string
}

func NewChromosomesFlag() *ChromosomesFlag {
	cf := &ChromosomesFlag{}
	return cf
}

func (cf *ChromosomesFlag) RegisterFlags(flags *pflag.FlagSet) {
	flags.StringSliceVar(&cf.chromosomes, "chromosomes", []string{}, "Limit to a specified, comma-delimited list of chromosomes, or 'all' for all chromosomes")
}

func (cf *ChromosomesFlag) PreRunE(cmd *cobra.Command, args []string) (err error) {
	return nil
}

func (cf *ChromosomesFlag) GetChromosomes() []string {
	return cf.chromosomes
}
