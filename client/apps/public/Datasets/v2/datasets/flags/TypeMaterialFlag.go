package flags

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type TypeMaterialFlag struct {
	FlagInterface
	typeMaterial bool
}

func NewTypeMaterialFlag() *TypeMaterialFlag {
	return &TypeMaterialFlag{typeMaterial: false}
}

func (tm *TypeMaterialFlag) RegisterFlags(flags *pflag.FlagSet) {
	flags.BoolVar(&tm.typeMaterial, "from-type", false, "Only return records with type material")
}

func (tm *TypeMaterialFlag) PreRunE(cmd *cobra.Command, args []string) (err error) {
	return nil
}

func (tm *TypeMaterialFlag) TypeMaterial() bool {
	return tm.typeMaterial
}
