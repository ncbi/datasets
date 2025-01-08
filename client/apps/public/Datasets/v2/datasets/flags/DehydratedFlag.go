package flags

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type DehydratedFlag struct {
	FlagInterface
	dehydrated bool
}

func NewDehydratedFlag() *DehydratedFlag {
	return &DehydratedFlag{dehydrated: false}
}

func (df *DehydratedFlag) RegisterFlags(flags *pflag.FlagSet) {
	flags.BoolVar(&df.dehydrated, "dehydrated", false, "Download a dehydrated zip archive including the data report and locations of data files (use the rehydrate command to retrieve data files).")
}

func (df *DehydratedFlag) PreRunE(cmd *cobra.Command, args []string) (err error) {
	return nil
}

func (df *DehydratedFlag) Dehydrated() bool {
	return df.dehydrated
}
