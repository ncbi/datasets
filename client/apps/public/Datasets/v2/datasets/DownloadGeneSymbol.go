package datasets

import (
	openapi "datasets/openapi/v2"
	cmdflags "datasets_cli/v2/datasets/flags"

	"github.com/spf13/cobra"
)

func createDownloadGeneSymbolCmd(dgf DownloadGeneFlag) *cobra.Command {
	iff := cmdflags.NewInputFileFlag(cmdflags.InputFileListTypeGeneSymbol, cmdflags.AsIntegerFalse)
	otf := cmdflags.NewOrthologTaxonFilterFlag()
	svf := cmdflags.NewSkipZipValidationFlag()
	gtff := cmdflags.NewGeneTaxonFilterFlag()

	flagSets := []cmdflags.FlagInterface{iff, otf, svf, gtff}

	cmd := &cobra.Command{
		Use:   "symbol <gene_symbol ...>",
		Short: "Download a gene data package by gene symbol",
		Example: `  datasets download gene symbol tp53
  datasets download gene symbol brca1 --taxon "mus musculus"`,
		Long: `
Download a gene data package by gene symbol and taxon (NCBI Taxonomy ID, scientific or common name for a species). If no taxon is specified, data will be returned for human (--taxon human).  Gene data packages include gene, transcript and protein sequences and one or more data reports. Data packages are downloaded as a zip archive.

The default gene data package includes the following files:
  * rna.fna (transcript sequences)
  * protein.faa (protein sequences)
  * data_report.jsonl (data report with gene metadata)
  * dataset_catalog.json (a list of files and file types included in the data package)`,
		PreRunE: cmdflags.ExecutePreRunEFor(flagSets),
		RunE: func(cmd *cobra.Command, args []string) error {
			taxId, taxError := RetrieveTaxIdForTaxon(
				gtff.Taxon,
				false,
				openapi.V2ORGANISMQUERYREQUESTTAXONRESOURCEFILTER_GENE,
				"gene",
			)
			if taxError != nil {
				return taxError
			}

			downloader, err := NewGeneDownloader(dgf.previewFlag.IsPreview(), dgf.geneIncludeFlag, dgf.filterFlag, WithSymbolAndTaxon(iff, otf, taxId))
			if err != nil {
				return err
			}
			return downloader.Download(svf.IsSkipValidation())
		},
	}
	cmdflags.RegisterAllFlags(flagSets, cmd.Flags())

	return cmd
}
