package datasets

import (
	openapi "datasets/openapi/v2"
	cmdflags "datasets_cli/v2/datasets/flags"

	"github.com/spf13/cobra"
)

func createDownloadGenomeAccessionCmd(dgf DownloadGenomeFlag, assemblyRequestFlag AssemblyRequestFlag) *cobra.Command {
	iff := cmdflags.NewInputFileFlag(cmdflags.InputFileListTypeGenomeAcc, cmdflags.AsIntegerFalse)
	flagSets := []cmdflags.FlagInterface{iff}

	cmd := &cobra.Command{
		Use:   "accession <accession ...>",
		Short: "Download a genome data package by Assembly or BioProject accession",
		Long: `
Download a genome data package by Assembly or BioProject accession. Genome data packages may include assembled genome, transcript and protein sequences, annotation and one or more data reports. Data packages are downloaded as a zip archive.

The default genome data package includes the following files:
  * <accession>_<assembly_name>_genomic.fna (genomic sequences)
  * assembly_data_report.jsonl (data report with genome assembly and annotation metadata)
  * dataset_catalog.json (a list of files and file types included in the data package)`,
		Example: `  datasets download genome accession GCF_000001405.40 --chromosomes X,Y --include protein,cds
  datasets download genome accession GCA_003774525.2 GCA_000001635 --chromosomes X,Y,Un.9
  datasets download genome accession GCA_003774525.2 --preview
  datasets download genome accession PRJNA289059 --include none`,
		PreRunE: cmdflags.ExecutePreRunEFor(flagSets),
		RunE: func(cmd *cobra.Command, args []string) error {
			// This could potentially change AssemblyVersion
			accessions := assemblyRequestFlag.assemblyVersionFlag.UpdateForInputAccessions(iff.InputIDArgs)

			if dgf.downloadPreviewFlag.IsPreview() {
				request, err := GetGenomeReportsAccessionRequest(accessions)
				if err != nil {
					return err
				}
				request.SetReturnedContent(openapi.V2ASSEMBLYDATASETREPORTSREQUESTCONTENTTYPE_ASSM_ACC)
				err = updateAssemblyReportRequestOption(request, assemblyRequestFlag)
				if err != nil {
					return err
				}
				return getDownloadSummary(dgf, NewGenomeAccessionRequestIter(request))
			}
			downloader, warning, err := NewGenomeDownloader(GenomeWithAccessions(accessions), dgf, assemblyRequestFlag)
			if warning != "" {
				cmd.PrintErrln(warning)
			}
			if err != nil {
				return err
			}
			err = downloader.Download(dgf.skipValidationFlag.IsSkipValidation())
			return err
		},
	}

	cmdflags.RegisterAllFlags(flagSets, cmd.Flags())

	return cmd
}
