package datasets

import (
	cmdflags "datasets_cli/v2/datasets/flags"

	"github.com/spf13/cobra"
)

func createDownloadGeneGeneIDCmd(dgf DownloadGeneFlag) *cobra.Command {
	iff := cmdflags.NewInputFileFlag(cmdflags.InputFileListTypeGeneId, cmdflags.AsIntegerTrue)
	svf := cmdflags.NewSkipZipValidationFlag()
	otf := cmdflags.NewOrthologTaxonFilterFlag()
	flagSets := []cmdflags.FlagInterface{iff, otf, svf}

	cmd := &cobra.Command{
		Use:   "gene-id <gene-id ...>",
		Short: "Download a gene data package by NCBI Gene ID",
		Example: `  datasets download gene gene-id 672
  datasets download gene gene-id 2597 14433`,
		Long: `
Download a gene data package by NCBI Gene ID. Gene data packages include gene, transcript and protein sequences and one or more data reports. Data packages are downloaded as a zip archive.

The default gene data package includes the following files:
  * rna.fna (transcript sequences)
  * protein.faa (protein sequences)
  * data_report.jsonl (data report with gene metadata)
  * dataset_catalog.json (a list of files and file types included in the data package)`,
		PreRunE: cmdflags.ExecutePreRunEFor(flagSets),

		RunE: func(cmd *cobra.Command, args []string) error {
			downloader, err := NewGeneDownloader(dgf.previewFlag.IsPreview(), dgf.geneIncludeFlag, dgf.filterFlag, WithGeneIds(iff, otf))
			if err != nil {
				return err
			}
			return downloader.Download(svf.IsSkipValidation())
		},
	}
	cmdflags.RegisterAllFlags(flagSets, cmd.Flags())

	return cmd
}
