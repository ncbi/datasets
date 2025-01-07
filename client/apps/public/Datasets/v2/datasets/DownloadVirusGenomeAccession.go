package datasets

import (
	cmdflags "datasets_cli/v2/datasets/flags"
	"errors"
	"github.com/spf13/cobra"
)

func initInputFlagVirusAccession() (iff *cmdflags.InputFileFlag, svf *cmdflags.SkipZipValidation, flagSets []cmdflags.FlagInterface) {
	iff = cmdflags.NewInputFileFlag(cmdflags.InputFileListTypeVirusAcc, cmdflags.AsIntegerFalse)
	svf = cmdflags.NewSkipZipValidationFlag()
	flagSets = []cmdflags.FlagInterface{iff, svf}

	return iff, svf, flagSets
}

func createDownloadVirusGenomeAccessionCmd(dvf DownloadVirusFlag) *cobra.Command {

	iff, svf, flagSets := initInputFlagVirusAccession()

	cmd := &cobra.Command{
		Use:   "accession <accession ...>",
		Short: "Download a virus genome data package by accession.",
		Long: `
Download a virus genome data package by GenBank or RefSeq nucleotide accession. Virus genome data packages include genome, transcript and protein sequences, annotation and one or more data reports. Data packages are downloaded as a zip archive.

The default virus genome data package includes the following files:
  * genomic.fna (genomic sequences)
  * data_report.jsonl (data report with virus genome metadata)
  * dataset_catalog.json (a list of files and file types included in the data package)`,
		Example: `  datasets download virus genome accession NC_045512.2 --include genome,cds,protein`,
		PreRunE: cmdflags.ExecutePreRunEFor(flagSets),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			idArgs := iff.InputIDArgs
			if len(idArgs) == 0 {
				return errors.New("Input accessions not specified")
			}
			downloader, warning, err := NewVirusDownloader(VirusDownloadWithAccession(idArgs, dvf))
			if err != nil {
				if warning != "" {
					cmd.Println(warning)
				}
				return err
			}
			if warning != "" {
				cmd.Println(warning)
			}
			return downloader.Download(svf.IsSkipValidation())

		},
	}
	cmdflags.RegisterAllFlags(flagSets, cmd.Flags())
	return cmd
}
