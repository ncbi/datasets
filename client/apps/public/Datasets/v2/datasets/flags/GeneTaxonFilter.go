package flags

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type GeneTaxonFilterFlag struct {
	FlagInterface
	Taxon string
}

func NewGeneTaxonFilterFlag() *GeneTaxonFilterFlag {
	return &GeneTaxonFilterFlag{}
}

func (gtff *GeneTaxonFilterFlag) RegisterFlags(flags *pflag.FlagSet) {
	flags.StringVar(&gtff.Taxon, "taxon", "human", "Define species (NCBI taxid, common or scientific name) for gene symbol")
}

func (gtff *GeneTaxonFilterFlag) PreRunE(cmd *cobra.Command, args []string) (err error) {
	return nil
}
