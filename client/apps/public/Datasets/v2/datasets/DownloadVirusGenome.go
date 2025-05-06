package datasets

import (
	openapi "datasets/openapi/v2"
	cmdflags "datasets_cli/v2/datasets/flags"

	"github.com/spf13/cobra"
)

type DownloadVirusFlag struct {
	virusFilterFlags *cmdflags.VirusFilterFlags

	virusSeqFlags *cmdflags.VirusDownloadIncludeFlag
	skipValFlag   *cmdflags.SkipZipValidation
	cmdFlagSet    []cmdflags.FlagInterface
}

func (dvf *DownloadVirusFlag) addFiltersAndOptionsTo(request *openapi.V2VirusDatasetRequest) {
	dvf.virusFilterFlags.PrepareDownloadRequest(request)
	dvf.virusSeqFlags.PrepareDownloadRequest(request)
}

func (dvf *DownloadVirusFlag) prepareSars2ProteinDatasetRequest(request *openapi.V2Sars2ProteinDatasetRequest) {
	dvf.virusFilterFlags.PrepareSarsProteinDownloadRequest(request)
	dvf.virusSeqFlags.PrepareSarsProteinDownloadRequest(request)
}

func initDownloadVirusFlag(longDesc string, releasedAfterDesc string, updatedAfterDesc string) DownloadVirusFlag {

	vsf := cmdflags.NewVirusDownloadIncludeFlag(longDesc)
	svf := cmdflags.NewSkipZipValidationFlag()
	vff := cmdflags.NewVirusFilterFlags(releasedAfterDesc, updatedAfterDesc)

	dvf := DownloadVirusFlag{
		skipValFlag:      svf,
		virusFilterFlags: vff,
		virusSeqFlags:    vsf,
		cmdFlagSet:       []cmdflags.FlagInterface{svf, vff, vsf},
	}

	return dvf
}

func createDownloadVirusGenomeCmd() *cobra.Command {

	downloadVirusGenomeFlag := initDownloadVirusFlag(
		cmdflags.IncludeSequenceLongDescGenome,
		cmdflags.GenomeReleasedAfterDesc,
		cmdflags.VirusGenomeUpdatedAfterDesc,
	)

	cmd := &cobra.Command{
		Use:   "genome [command] ",
		Short: "Download a virus genome dataset by accession or taxon",
		Long: `
		Download a virus genome data package by GenBank or RefSeq nucleotide accession. Virus genome data packages include genome,
		transcript and protein sequences, annotation and one or more data reports. Data packages are downloaded as a zip archive.

		The default virus genome data package includes the following files:
		  * genomic.fna (genomic sequences)
		  * data_report.jsonl (data report with virus genome metadata)
		  * dataset_catalog.json (a list of files and file types included in the data package)`,
		Example: `  datasets download virus genome taxon sars-cov-2 --host dog --include protein
  datasets download virus genome taxon coronaviridae --host "manis javanica"`,
		RunE:              ParentCommandRunE,
		PersistentPreRunE: cmdflags.PersistentPreRunEFor(downloadVirusGenomeFlag.cmdFlagSet, downloadCmd),
	}

	flags := cmd.PersistentFlags()

	cmdflags.RegisterAllFlags(downloadVirusGenomeFlag.cmdFlagSet, flags)

	cmd.AddCommand(createDownloadVirusGenomeAccessionCmd(downloadVirusGenomeFlag))
	cmd.AddCommand(createDownloadVirusGenomeTaxonCmd(downloadVirusGenomeFlag))

	return cmd
}
