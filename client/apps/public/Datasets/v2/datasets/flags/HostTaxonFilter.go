package flags

import (
	openapi "datasets/openapi/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"strings"
)

type HostTaxonFilterFlag struct {
	FlagInterface
	hostTaxonFilter string
	hostTaxId       string
}

func NewHostTaxonFilterFlag() *HostTaxonFilterFlag {
	htf := &HostTaxonFilterFlag{}
	return htf
}

func (htf *HostTaxonFilterFlag) RegisterFlags(flags *pflag.FlagSet) {
	flags.StringVar(&htf.hostTaxonFilter, "host", htf.hostTaxonFilter, "Limit to virus genomes isolated from a specified host species")
}

func (htf *HostTaxonFilterFlag) PreRunE(cmd *cobra.Command, args []string) (err error) {
	// Shortcut to support Cache package build without taxon service lookup
	if htf.hostTaxonFilter == "9606" {
		htf.hostTaxId = htf.hostTaxonFilter
		return nil
	}
	if strings.TrimSpace(htf.hostTaxonFilter) != "" {
		taxId, taxError := RetrieveTaxIdForTaxon(
			htf.hostTaxonFilter,
			true,
			openapi.V2ORGANISMQUERYREQUESTTAXONRESOURCEFILTER_ALL,
			"viral host",
		)
		if taxError != nil {
			return taxError
		}
		htf.hostTaxId = taxId
	}
	return nil
}

func (htf *HostTaxonFilterFlag) IsHostProvided() bool {
	return len(htf.hostTaxonFilter) > 0
}

func (htf *HostTaxonFilterFlag) HostTaxIdValue() string {
	return htf.hostTaxId
}
