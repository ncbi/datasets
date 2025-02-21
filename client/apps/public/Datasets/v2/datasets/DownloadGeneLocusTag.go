package datasets

import (
	cmdflags "datasets_cli/v2/datasets/flags"

	"github.com/spf13/cobra"
)

func createDownloadGeneLocusTagCmd(dgf DownloadGeneFlag) *cobra.Command {
	iff := cmdflags.NewInputFileFlag(cmdflags.InputFileListTypeLocusTag, cmdflags.AsIntegerFalse)
	otf := cmdflags.NewOrthologTaxonFilterFlag()
	svf := cmdflags.NewSkipZipValidationFlag()
	gtff := cmdflags.NewGeneTaxonFilterFlag()

	flagSets := []cmdflags.FlagInterface{iff, otf, svf, gtff}

	cmd := &cobra.Command{
		Use:   "locus-tag <locus-tag ...>",
		Short: "Download a gene data package by locus tag",
		Example: `  datasets download gene locus-tag b0001
  datasets download gene locus-tag b0001 ArthCt125`,
		Long: `
Download a gene data package by locus tag. Gene data packages include gene, transcript and protein sequences and one or more data reports. Data packages are downloaded as a zip archive.

The default gene data package includes the following files:
  * rna.fna (transcript sequences)
  * protein.faa (protein sequences)
  * data_report.jsonl (data report with gene metadata)
  * dataset_catalog.json (a list of files and file types included in the data package)`,
		PreRunE: cmdflags.ExecutePreRunEFor(flagSets),
		RunE: func(cmd *cobra.Command, args []string) error {
			downloader, err := NewGeneDownloader(dgf.previewFlag.IsPreview(), dgf.geneIncludeFlag, dgf.filterFlag, WithLocusTags(iff, otf))
			if err != nil {
				return err
			}
			return downloader.Download(svf.IsSkipValidation())
		},
	}
	cmdflags.RegisterAllFlags(flagSets, cmd.Flags())

	return cmd
}
