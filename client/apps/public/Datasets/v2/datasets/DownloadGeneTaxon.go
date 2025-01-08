package datasets

import (
	cmdflags "datasets_cli/v2/datasets/flags"
	openapi "datasets/openapi/v2"
	"fmt"

	"github.com/spf13/cobra"
)

func createDownloadGeneTaxonCmd(dgf DownloadGeneFlag) *cobra.Command {

	inputDescription := "taxon (NCBI Taxonomy ID, scientific or common name for a species)"
	svf := cmdflags.NewSkipZipValidationFlag()
	flagSets := []cmdflags.FlagInterface{svf}

	cmd := &cobra.Command{
		Use:     "taxon <taxon>",
		Short:   fmt.Sprintf("Download a gene data package by %s", inputDescription),
		Example: `  datasets download gene taxon human --include protein,cds`,
		Long: fmt.Sprintf(`
Download a gene data package by %s.  Gene data packages include gene, transcript and protein sequences and one or more data reports. Data packages are downloaded as a zip archive.

The default gene data package includes the following files:
  * rna.fna (transcript sequences)
  * protein.faa (protein sequences)
  * data_report.jsonl (data report with gene metadata)
  * dataset_catalog.json (a list of files and file types included in the data package)`, inputDescription),
		Args: cmdflags.ExpectOnePositionalArgument(inputDescription),
		RunE: func(cmd *cobra.Command, args []string) error {
			taxId, taxError := RetrieveTaxIdForTaxon(
				args[0],
				false,
				openapi.V2ORGANISMQUERYREQUESTTAXONRESOURCEFILTER_GENE,
				"gene",
			)
			if taxError != nil {
				return taxError
			}

			downloader, err := NewGeneDownloader(dgf.previewFlag.IsPreview(), dgf.geneIncludeFlag, dgf.filterFlag, WithTaxon(taxId))
			if err != nil {
				return err
			}
			return downloader.Download(svf.IsSkipValidation())
		},
	}
	cmdflags.RegisterAllFlags(flagSets, cmd.Flags())

	return cmd
}
