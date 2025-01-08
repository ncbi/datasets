package flags

import (
	openapi "datasets/openapi/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"strings"
)

type OrthologTaxonFilterFlag struct {
	FlagInterface
	orthologTaxonFilter []string
	orthologTaxIdFilter []string
}

// Might need to parameterize for download vs summary to fix the help text.
func NewOrthologTaxonFilterFlag() *OrthologTaxonFilterFlag {
	otf := &OrthologTaxonFilterFlag{}
	return otf
}

func (otf *OrthologTaxonFilterFlag) RegisterFlags(flags *pflag.FlagSet) {
	flags.StringSliceVar(&otf.orthologTaxonFilter, "ortholog", otf.orthologTaxonFilter, "Retrieves data for an ortholog set. Provide one or more taxa (any rank, limited to vertebrates and insects) to filter results or 'all' for the complete set.")
}

func (otf *OrthologTaxonFilterFlag) PreRunE(cmd *cobra.Command, args []string) (err error) {
	if Contains(otf.orthologTaxonFilter, "all") {
		return nil
	}
	for _, requestedTaxon := range otf.orthologTaxonFilter {
		if strings.TrimSpace(requestedTaxon) != "" {
			taxId, taxError := RetrieveTaxIdForTaxon(
				requestedTaxon,
				true,
				openapi.V2ORGANISMQUERYREQUESTTAXONRESOURCEFILTER_GENE,
				"gene orthologs",
			)
			if taxError != nil {
				return taxError
			}
			otf.orthologTaxIdFilter = append(otf.orthologTaxIdFilter, taxId)
		}
	}

	return nil
}

func (otf *OrthologTaxonFilterFlag) IsOrthologRequested() bool {
	return len(otf.orthologTaxonFilter) > 0
}

func (otf *OrthologTaxonFilterFlag) RequestAllOrthologs() bool {
	return Contains(otf.orthologTaxonFilter, "all")
}

func (otf *OrthologTaxonFilterFlag) OrthologTaxonValue() []string {
	return otf.orthologTaxIdFilter
}

// Helper
// Contains is a case insensitive match, finding needle in a haystack
func Contains(haystack []string, needle string) bool {
	for _, a := range haystack {
		if strings.ToLower(a) == strings.ToLower(needle) {
			return true
		}
	}
	return false
}
