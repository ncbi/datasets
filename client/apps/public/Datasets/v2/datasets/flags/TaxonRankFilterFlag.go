package flags

import (
	openapi "datasets/openapi/v2"

	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"strings"
)

type TaxonRankFilterFlag struct {
	FlagInterface
	taxonRankFilter []string
	openapiRanks    []openapi.V2reportsRankType
}

func NewTaxonRankFilterFlag() *TaxonRankFilterFlag {
	trf := &TaxonRankFilterFlag{}
	return trf
}

func (trf *TaxonRankFilterFlag) RegisterFlags(flags *pflag.FlagSet) {
	flags.StringSliceVar(&trf.taxonRankFilter, "rank", trf.taxonRankFilter, "Limit taxons to only those with the specified ranks. When enabled, the children and parents flags will be enabled as well.")
}

func (lf *TaxonRankFilterFlag) PreRunE(cmd *cobra.Command, args []string) (err error) {
	for _, rank := range lf.taxonRankFilter {
		// Provide a little flexibility - automatically convert to upper case and replace dashes with underscores
		cleanedRank := strings.Replace(strings.ToUpper(rank), "-", "_", -1)
		var rankType *openapi.V2reportsRankType
		rankType, err = openapi.NewV2reportsRankTypeFromValue(cleanedRank)
		if err == nil {
			lf.openapiRanks = append(lf.openapiRanks, *rankType)
		} else {
			rankTextReplacer := strings.NewReplacer("[", "", "]", "")
			allowedRanks := rankTextReplacer.Replace(fmt.Sprintf("%v", openapi.AllowedV2reportsRankTypeEnumValues))
			return fmt.Errorf("Invalid value '%v' for rank. Valid ranks are:\n%s", rank, allowedRanks)
		}
	}
	return
}

func (trf *TaxonRankFilterFlag) GetRanks() []openapi.V2reportsRankType {
	return trf.openapiRanks
}
