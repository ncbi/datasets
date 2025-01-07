package flags

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type SearchStringFlag struct {
	FlagInterface
	searchStrings []string
}

func NewSearchStringFlag() *SearchStringFlag {
	ssf := &SearchStringFlag{}
	return ssf
}

func (ssf *SearchStringFlag) RegisterFlags(flags *pflag.FlagSet) {
	flags.StringSliceVar(&ssf.searchStrings, "search", []string{}, `Limit results to genomes with specified text in the searchable fields:
species and infraspecies, assembly name and submitter.
To search multiple strings, use the flag multiple times.`)
}

func (ssf *SearchStringFlag) PreRunE(cmd *cobra.Command, args []string) (err error) {
	return nil
}

func (ssf *SearchStringFlag) GetSearchText() []string {
	return ssf.searchStrings
}
