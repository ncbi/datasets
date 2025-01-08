package datasets

import (
	"github.com/spf13/cobra"
)

func createSummaryCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "summary",
		Short: "Print a data report containing gene, genome, taxonomy or virus metadata",
		Long:  "Print a data report containing gene, genome, taxonomy or virus metadata in JSON format.",
		Example: `  datasets summary genome accession GCF_000001405.40
  datasets summary genome taxon "mus musculus"
  datasets summary gene gene-id 672
  datasets summary gene symbol brca1 --taxon "mus musculus"
  datasets summary gene accession NP_000483.3
  datasets summary taxonomy taxon "mus musculus"
  datasets summary virus genome accession NC_045512.2
  datasets summary virus genome taxon sars-cov-2 --host dog`,
		Args: cobra.NoArgs,
		RunE: ParentCommandRunE,
	}

	cmd.AddCommand(createSummaryGeneCmd())
	cmd.AddCommand(createSummaryGenomeCmd())
	cmd.AddCommand(createSummaryVirusCmd())
	cmd.AddCommand(createSummaryTaxonomyCmd())

	return cmd
}
