package datasets

import (
	"github.com/spf13/cobra"

	"context"

	cmdflags "datasets_cli/v2/datasets/flags"
	openapi "datasets/openapi/v2"
)

func downloadVirusProtein(cmd *cobra.Command, proteinNames []string, assmFilename string, dvf DownloadVirusFlag) (err error) {

	cli, err := createOAClient()
	if err != nil {
		return
	}

	request := openapi.NewV2Sars2ProteinDatasetRequest()
	request.SetProteins(proteinNames)
	dvf.prepareSars2ProteinDatasetRequest(request)

	f, err := afs.Create(assmFilename)
	if err != nil {
		return
	}

	length := int64(-1) // unknown length

	_, resp, err := cli.VirusAPI.Sars2ProteinDownloadPost(context.TODO()).V2Sars2ProteinDatasetRequest(*request).Execute()
	if err = handleHTTPResponse(resp, err); err != nil {
		return
	}
	err = downloadData(&f, resp, err, assmFilename, length, dvf.skipValFlag.IsSkipValidation())
	return
}

func createDownloadVirusProteinCmd() *cobra.Command {

	downloadVirusProteinFlag := initDownloadVirusFlag(
		cmdflags.IncludeSequenceLongDescProtein,
		cmdflags.VirusProteinReleasedAfterDesc,
		cmdflags.VirusProteinUpdatedAfterDesc,
	)
	cmd := &cobra.Command{
		Use:   "protein <protein_name ...> [flags]",
		Short: "Download a SARS-CoV-2 protein dataset by protein name",
		Long: `
	Download a SARS-CoV-2 protein data package by protein name. SARS-CoV-2 protein
	data packages include CDS and protein sequence, annotation and a detailed data report.
	Datasets are downloaded as a zip file.

The default SARS-CoV-2 protein data package includes the following files:
  * cds.fna (nucleotide coding sequences)
  * protein.faa (protein sequences)
  * data_report.jsonl (data report with viral metadata)
  * dataset_catalog.json (a list of files and file types included in the data package)

	Allowed protein names are: ORF1ab, ORF1a, nsp1, nsp2, nsp3, nsp4, nsp5, nsp6, nsp7, nsp8, nsp9, nsp10, rdrp, nsp11, nsp13, nsp14, nsp15, nsp16, S, ORF3a, E, M, ORF6, ORF7a, ORF7b, ORF8, N, ORF10`,
		Example: `  datasets download virus protein S --host dog --filename SARS2-spike-dog.zip
  datasets download virus protein rdrp --refseq --filename SARS2-rdrp-refseq.zip`,
		Args:    cmdflags.ExpectAtLeastOnePositionalArgument("SARS-Cov-2 protein name or symbol"),
		PreRunE: cmdflags.ExecutePreRunEFor(downloadVirusProteinFlag.cmdFlagSet),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			return downloadVirusProtein(cmd, args, argDownloadFilename, downloadVirusProteinFlag)
		},
	}

	flags := cmd.PersistentFlags()

	cmdflags.RegisterAllFlags(downloadVirusProteinFlag.cmdFlagSet, flags)

	return cmd
}
