package datasets

import (
	"github.com/spf13/cobra"
)

var geneDescrIDCmd = &cobra.Command{
	Deprecated: "please use \"summary gene gene-id\"",
	Use:        "gene-id <gene-id ...>",
	Short:      "Gene descriptions by NCBI GeneID",
	RunE:       cmdRunSummaryGeneID,
}

func init() {
	registerHiddenStringPair(geneDescrIDCmd.Flags(), &argInputFile, "inputfile", "i", "", "read a list of NCBI Gene IDs from a file to use as input")
}
