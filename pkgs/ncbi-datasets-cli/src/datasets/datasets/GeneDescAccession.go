package datasets

import (
	"errors"

	"github.com/spf13/cobra"

	openapi "main/openapi_client"
)

var geneDescrAccCmd = &cobra.Command{
	Deprecated: "please use \"summary gene accession\"",
	Use:        "acc <gene-accession ...>",
	Short:      "Gene descriptions by gene accession",

	RunE: func(cmd *cobra.Command, args []string) error {
		if len(argIDArgs) == 0 {
			return errors.New("Input accessions not specified")
		}
		req := openapi.NewV1GeneDatasetRequest()
		req.SetAccessions(argIDArgs)
		return streamGeneMatch(req, &JsonStreamProcessor{wrapperName: "genes"})
	},
}
