package datasets

import (
	"github.com/spf13/cobra"
)

// getAssemblyDescriptorsCmd represents the GetAssemblyDescriptors command
var getAssemblyDescriptorsCmd = &cobra.Command{
	Deprecated: "please use \"summary genome\"",
	Use:        "assembly-descriptors",
	Short:      "Deprecated - Retrieve descriptions of available genome assembly datasets",
	Long: `
Retrieve descriptions of available genome assemblies.  Search by assembly accession,
NCBI Taxonomy ID, scientific or common name at any tax rank across all domains of life.

Refer to NCBI's [command line quickstart](https://www.ncbi.nlm.nih.gov/datasets/docs/quickstarts/command-line-tools/) documentation for information about getting started with the command-line tools.`,
}

func init() {
}
