package flags

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type FastaFilterFlag struct {
	FlagInterface
	fastaFilter     []string
	fastaFilterFile string
	filters         []string
}

func NewFastaFilterFlag() *FastaFilterFlag {
	return &FastaFilterFlag{}
}

func (fff *FastaFilterFlag) RegisterFlags(flags *pflag.FlagSet) {
	flags.StringSliceVar(&fff.fastaFilter, "fasta-filter", []string{}, "Limit protein and RNA sequence files to the specified RefSeq nucleotide and protein accessions")
	flags.StringVar(&fff.fastaFilterFile, "fasta-filter-file", "", "Limit protein and RNA sequence files to the specified RefSeq nucleotide and protein accessions included in the specified file")
}

func (fff *FastaFilterFlag) PreRunE(cmd *cobra.Command, args []string) (err error) {
	if fff.fastaFilterFile != "" {
		fp, fileErr := os.Open(fff.fastaFilterFile)
		if fileErr != nil {
			return fmt.Errorf("'%s' opening input file: '%s'", fileErr.Error(), fff.fastaFilterFile)
		}
		defer fp.Close()
		fff.filters = readLines(fp)
		// Check if any accessions were read
		if len(fff.filters) == 0 {
			return fmt.Errorf(
				"No identifiers read from file: '%s'\n       File should have 1 identifier per row and no spaces or quotes",
				fff.fastaFilterFile,
			)
		}
	}

	// Set other fields
	if len(fff.fastaFilter) > 0 {
		fff.filters = append(fff.filters, fff.fastaFilter...)
	}

	return nil
}

func (fff *FastaFilterFlag) Filters() []string {
	return fff.filters
}
