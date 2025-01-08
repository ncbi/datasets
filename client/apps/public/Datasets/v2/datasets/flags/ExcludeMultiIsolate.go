package flags

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type ExcludeMultiIsolateFlag struct {
	FlagInterface
	excludeMultiIsolate bool
}

func NewExcludeMultiIsolateFlag() *ExcludeMultiIsolateFlag {
	return &ExcludeMultiIsolateFlag{excludeMultiIsolate: false}
}

func (mi *ExcludeMultiIsolateFlag) RegisterFlags(flags *pflag.FlagSet) {
	flags.BoolVar(&mi.excludeMultiIsolate, "exclude-multi-isolate", false, "Exclude assemblies from multi-isolate projects")
}

func (lf *ExcludeMultiIsolateFlag) PreRunE(cmd *cobra.Command, args []string) (err error) {
	return nil
}

func (mi *ExcludeMultiIsolateFlag) ExcludeMultiIsolate() bool {
	return mi.excludeMultiIsolate
}
