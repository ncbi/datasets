package flags

import (
	openapi "datasets/openapi/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var (
	RetrieveTaxIdForTaxon func(string, bool, openapi.V2OrganismQueryRequestTaxonResourceFilter, string, ...int32) (string, error)
)

type RegisterFlagsType func(flags *pflag.FlagSet)
type PreRunEType func(cmd *cobra.Command, args []string) error

type FlagInterface interface {
	RegisterFlags(flags *pflag.FlagSet)
	PreRunE(cmd *cobra.Command, args []string) error
}

func ExecutePreRunEFor(flagSets []FlagInterface) PreRunEType {
	return func(cmd *cobra.Command, args []string) error {
		for _, f := range flagSets {
			err := f.PreRunE(cmd, args)
			if err != nil {
				return err
			}
		}
		return nil
	}
}

func PersistentPreRunEFor(flagSets []FlagInterface, parentCmd *cobra.Command) PreRunEType {
	return func(cmd *cobra.Command, args []string) error {
		if err := ExecutePreRunEFor(flagSets)(cmd, args); err != nil {
			return err
		}
		if parentCmd == nil {
			return nil
		}
		if parentCmd.PersistentPreRunE != nil {
			return parentCmd.PersistentPreRunE(cmd, args)
		}
		return nil
	}
}

func RegisterAllFlags(flagSets []FlagInterface, flags *pflag.FlagSet) {
	for _, f := range flagSets {
		f.RegisterFlags(flags)
	}
}
