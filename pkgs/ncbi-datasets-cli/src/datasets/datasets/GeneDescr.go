package datasets

import (
	"github.com/spf13/cobra"
)

// getGeneDescriptorsCmd represents the GetGeneDescriptors command
var getGeneDescriptorsCmd = &cobra.Command{
	Deprecated: "please use \"summary gene\"",
	Use:        "gene-descriptors",
	Short:      "Retrieve metadata of available gene datasets",
	Long: `
Retrieve descriptions of available NCBI Gene records for all domains of life.

Refer to NCBI's [command line quickstart](https://www.ncbi.nlm.nih.gov/datasets/docs/quickstarts/command-line-tools/) documentation for information about getting started with the command-line tools.
`,

	PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
		argIDArgs, err = getArgsFromListOrFile(args, argInputFile)
		return
	},
}

func init() {
	getGeneDescriptorsCmd.AddCommand(geneDescrIDCmd)
	getGeneDescriptorsCmd.AddCommand(geneDescrSymbolCmd)
	getGeneDescriptorsCmd.AddCommand(geneDescrAccCmd)

	getGeneDescriptorsCmd.PersistentFlags().StringVarP(&argInputFile, "inputfile", "i", "", "file to read list of gene identifiers")
}
