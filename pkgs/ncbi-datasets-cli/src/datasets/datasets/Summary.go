package datasets

import (
	"github.com/spf13/cobra"
)

var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "print a summary of a gene or genome dataset",
	Long: `
Print a summary of a gene or genome dataset in JSON format.

Refer to NCBI's [command line quickstart](https://www.ncbi.nlm.nih.gov/datasets/docs/quickstarts/command-line-tools/) documentation for information about getting started with the command-line tools.`,
	Example: `  datasets summary genome accession GCF_000001405.39
  datasets summary genome taxon "mus musculus"
  datasets summary gene gene-id 672
  datasets summary gene symbol brca1 --taxon mouse
  datasets summary gene accession NP_000483.3
  datasets summary virus genome accession NC_045512.2
  datasets summary virus genome taxon SARS-COV-2 --host dog`,
}

func init() {
	summaryCmd.AddCommand(summaryVirusCmd)
	summaryCmd.AddCommand(summaryGeneCmd)
	summaryCmd.AddCommand(summaryGenomeCmd)
	summaryCmd.AddCommand(summaryOrthologCmd)
}
