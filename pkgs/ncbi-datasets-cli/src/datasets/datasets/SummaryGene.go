package datasets

import (
	"github.com/spf13/cobra"
)

var summaryGeneCmd = &cobra.Command{
	Use:   "gene",
	Short: "print a summary of a gene dataset",
	Long: `Print a summary of a gene dataset by NCBI Gene ID, gene symbol or RefSeq nucleotide or protein accession. The summary is returned in JSON format.

Refer to NCBI's [download and install](https://www.ncbi.nlm.nih.gov/datasets/docs/download-and-install/) documentation for information about getting started with the command-line tools.`,
	Example: `  datasets summary gene gene-id 672
  datasets summary gene symbol brca1 --taxon mouse
  datasets summary gene accession NP_000483.3`,

	PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
		argIDArgs, err = getArgsFromListOrFile(args, argInputFile)
		return
	},
}

func init() {
	summaryGeneCmd.AddCommand(summaryGeneIDCmd)
	summaryGeneCmd.AddCommand(summaryGeneSymbolCmd)
	summaryGeneCmd.AddCommand(summaryGeneAccCmd)
	summaryGeneCmd.AddCommand(summaryGeneTaxonCmd)

	pflags := summaryGeneCmd.PersistentFlags()
	pflags.BoolVar(&argJsonLinesFormat, "as-json-lines", false, "Stream results as newline delimited JSON-Lines")

}
