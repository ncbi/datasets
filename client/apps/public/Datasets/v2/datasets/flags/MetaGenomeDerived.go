package flags

import (
	openapi "datasets/openapi/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/thediveo/enumflag/v2"
)

type MetagenomeDerived enumflag.Flag

// Define the enumeration values for MetagenomeDerived
const (
	MetagenomeDerived_All MetagenomeDerived = iota
	Only
	Exclude
)

var MetagenomeDerivedIds = map[MetagenomeDerived][]string{
	MetagenomeDerived_All: {"all", "METAGENOME_DERIVED_UNSET"},
	Only:                  {"only", "metagenome_derived_only"},
	Exclude:               {"exclude", "metagenome_derived_exclude"},
}

type MetaGenomeDerivedFlag struct {
	FlagInterface
	metaGenomeDerived MetagenomeDerived
}

func NewMetaGenomeDerivedFlag() *MetaGenomeDerivedFlag {
	mgd := &MetaGenomeDerivedFlag{
		metaGenomeDerived: MetagenomeDerived_All,
	}
	return mgd
}

func (mgd *MetaGenomeDerivedFlag) GetMetaGenomeFilter() openapi.V2AssemblyDatasetDescriptorsFilterMetagenomeDerivedFilter {
	return openapi.V2AssemblyDatasetDescriptorsFilterMetagenomeDerivedFilter(MetagenomeDerivedIds[mgd.metaGenomeDerived][1])
}

func (mgd *MetaGenomeDerivedFlag) RegisterFlags(flags *pflag.FlagSet) {
	flags.Var(
		enumflag.New(&mgd.metaGenomeDerived, "string", MetagenomeDerivedIds, enumflag.EnumCaseInsensitive),
		"mag",
		"Limit to metagenome assembled genomes (only) or remove them from the results (exclude)")
}

func (mgd *MetaGenomeDerivedFlag) PreRunE(cmd *cobra.Command, args []string) (err error) {
	return nil
}
