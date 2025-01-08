package flags

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type ExcludeAtypicalFlag struct {
	FlagInterface
	excludeAtypical bool
}

func NewExcludeAtypicalFlag() *ExcludeAtypicalFlag {
	return &ExcludeAtypicalFlag{excludeAtypical: false}
}

func (tm *ExcludeAtypicalFlag) RegisterFlags(flags *pflag.FlagSet) {
	flags.BoolVar(&tm.excludeAtypical, "exclude-atypical", false, "Exclude atypical assemblies")
}

func (lf *ExcludeAtypicalFlag) PreRunE(cmd *cobra.Command, args []string) (err error) {
	return nil
}

func (tm *ExcludeAtypicalFlag) ExcludeAtypical() bool {
	return tm.excludeAtypical
}
