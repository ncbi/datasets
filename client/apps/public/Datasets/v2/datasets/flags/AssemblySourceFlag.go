package flags

import (
	openapi "datasets/openapi/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/thediveo/enumflag/v2"
)

type GenomeAssemblySource enumflag.Flag

const (
	All GenomeAssemblySource = iota
	RefSeq
	Genbank
)

var GenomeAssemblySourceIds = map[GenomeAssemblySource][]string{
	All:     {"all"},
	RefSeq:  {"refseq"},
	Genbank: {"genbank"},
}

type AssemblySourceFlag struct {
	FlagInterface
	genomeAssemblySource GenomeAssemblySource
}

func NewAssemblySourceFlag() *AssemblySourceFlag {
	asf := &AssemblySourceFlag{
		genomeAssemblySource: All,
	}
	return asf
}

func (asf *AssemblySourceFlag) GetAssemblySourceFilter() openapi.V2AssemblyDatasetDescriptorsFilterAssemblySource {
	return openapi.V2AssemblyDatasetDescriptorsFilterAssemblySource(GenomeAssemblySourceIds[asf.genomeAssemblySource][0])
}

func (asf *AssemblySourceFlag) RegisterFlags(flags *pflag.FlagSet) {
	flags.Var(
		enumflag.New(&asf.genomeAssemblySource, "string", GenomeAssemblySourceIds, enumflag.EnumCaseInsensitive),
		"assembly-source",
		"Limit to 'RefSeq' (GCF_) or 'GenBank' (GCA_) genomes")
}

func (asf *AssemblySourceFlag) PreRunE(cmd *cobra.Command, args []string) (err error) {
	return nil
}
