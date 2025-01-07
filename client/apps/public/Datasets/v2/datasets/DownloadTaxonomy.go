package datasets

import (
	"fmt"

	"github.com/spf13/cobra"
)

func createTaxonomyCmd() *cobra.Command {
	inputDescription := "taxonomy taxon (NCBI Taxonomy ID, scientific or common name at any tax rank)"
	cmd := &cobra.Command{
		Use:   "taxonomy",
		Short: "Download a taxonomy data package",
		Long: fmt.Sprintf(`
Download a taxonomy data package by %s.

The default taxonomy data package includes the following files:
  * taxonomy_report.jsonl
  * taxonomy_summary.tsv
  * dataset_catalog.json (a list of files and file types included in the data package)
A taxonomy names data report can also be added to the package
  * names_report.jsonl`, inputDescription),
		Example: `  datasets download taxonomy taxon "bos taurus"
  datasets download taxonomy taxon human,"drosophila melanogaster" --include names
  datasets download taxonomy taxon 10116 --parents --children`,

		Args: cobra.NoArgs,
		RunE: ParentCommandRunE,
	}

	cmd.AddCommand(createDownloadTaxonomyTaxonCmd())
	return cmd
}
