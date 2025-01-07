package flags

import (
	openapi "datasets/openapi/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/thediveo/enumflag"
)

type GenomeAssemblyLevel enumflag.Flag

const (
	Chromosome GenomeAssemblyLevel = iota
	Scaffold
	Contig
	Complete
)

var GenomeAssemblyLevelIds = map[GenomeAssemblyLevel][]string{
	Chromosome: {"chromosome"},
	Scaffold:   {"scaffold"},
	Contig:     {"contig"},
	Complete:   {"complete"},
}

var GenomeAssemblyLevelOpenapi = map[GenomeAssemblyLevel]openapi.V2reportsAssemblyLevel{
	Chromosome: openapi.V2REPORTSASSEMBLYLEVEL_CHROMOSOME,
	Scaffold:   openapi.V2REPORTSASSEMBLYLEVEL_SCAFFOLD,
	Contig:     openapi.V2REPORTSASSEMBLYLEVEL_CONTIG,
	Complete:   openapi.V2REPORTSASSEMBLYLEVEL_COMPLETE_GENOME,
}

type GenomeAssemblyLevelFlag struct {
	FlagInterface
	genomeAssemblyLevels []GenomeAssemblyLevel
}

func NewGenomeAssemblyLevelFlag() *GenomeAssemblyLevelFlag {
	gal := &GenomeAssemblyLevelFlag{
		genomeAssemblyLevels: []GenomeAssemblyLevel{},
	}
	return gal
}

func (gal *GenomeAssemblyLevelFlag) GetAssemblyLevels() (assemblyLevels []openapi.V2reportsAssemblyLevel) {
	for _, level := range gal.genomeAssemblyLevels {
		assemblyLevels = append(assemblyLevels, GenomeAssemblyLevelOpenapi[level])
	}
	return
}

func (gal *GenomeAssemblyLevelFlag) RegisterFlags(flags *pflag.FlagSet) {
	flags.Var(
		enumflag.NewSlice(&gal.genomeAssemblyLevels, "string", GenomeAssemblyLevelIds, enumflag.EnumCaseInsensitive),
		"assembly-level",
		`Limit to genomes at one or more assembly levels (comma-separated):
  * chromosome
  * complete
  * contig
  * scaffold
  `)
}

func (gal *GenomeAssemblyLevelFlag) PreRunE(cmd *cobra.Command, args []string) (err error) {
	return nil
}
