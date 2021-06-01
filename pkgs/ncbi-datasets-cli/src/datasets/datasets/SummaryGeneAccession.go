package datasets

import (
	"errors"
	openapi "main/openapi_client"

	"github.com/spf13/cobra"
)

var summaryGeneAccCmd = &cobra.Command{
	Use:   "accession <refseq-accession ...>",
	Short: "print a summary of a gene dataset by RefSeq nucleotide or protein accession",
	Long:  `Print a summary of a gene dataset by RefSeq nucleotide or protein accession. The summary is returned in JSON format.`,
	Example: `  datasets summary gene accession NP_000483.3
  datasets summary gene accession NM_000546.6 NM_000492.4`,

	RunE: func(cmd *cobra.Command, args []string) error {
		if len(argIDArgs) == 0 {
			return errors.New("Input accessions not specified")
		}
		req := new(openapi.V1alpha1GeneDatasetRequest)
		req.Accessions = argIDArgs
		if argJsonLinesFormat {
			return streamGeneMatch(req, &JsonLinesStreamProcessor{})
		}
		return streamGeneMatch(req, &JsonStreamProcessor{wrapperName: "genes"})
	},
}

func init() {
	flags := summaryGeneAccCmd.Flags()
	registerHiddenStringPair(flags, &argInputFile, "inputfile", "i", "", "read a list of RefSeq nucleotide or protein accessions from a file to use as input")
}
