package flags

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type SkipZipValidation struct {
	FlagInterface
	skipValidation bool
}

func NewSkipZipValidationFlag() *SkipZipValidation {
	return &SkipZipValidation{}
}

func (sv *SkipZipValidation) RegisterFlags(flags *pflag.FlagSet) {
	flags.BoolVar(&sv.skipValidation, "fast-zip-validation", false, "Skip zip checksum validation after download")
}

func (lf *SkipZipValidation) PreRunE(cmd *cobra.Command, args []string) (err error) {
	return nil
}

func (sv *SkipZipValidation) IsSkipValidation() bool {
	return sv.skipValidation
}
