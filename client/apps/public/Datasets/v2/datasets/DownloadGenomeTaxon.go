package datasets

import (
	"fmt"

	cmdflags "datasets_cli/v2/datasets/flags"
	openapi "datasets/openapi/v2"

	"github.com/spf13/cobra"
)

func createDownloadGenomeTaxonCmd(dgf DownloadGenomeFlag, assemblyRequestFlag AssemblyRequestFlag) *cobra.Command {
	tem := cmdflags.NewTaxExactMatchFlag()
	iff := cmdflags.NewInputFileFlag(cmdflags.InputFileListTypeTaxon, cmdflags.AsIntegerFalse)
	flagSets := []cmdflags.FlagInterface{tem, iff}

	inputDescription := "taxon (NCBI Taxonomy ID, scientific or common name at any tax rank)"
	cmd := &cobra.Command{
		Use:   "taxon <taxon>",
		Short: fmt.Sprintf("Download a genome data package by %s", inputDescription),
		Long: fmt.Sprintf(`
Download a genome data package by %s. Genome data packages may include genome, transcript and protein sequences, annotation and one or more data reports. Data packages are downloaded as a zip archive.

The default genome data package includes the following files:
  * <accession>_<assembly_name>_genomic.fna (genomic sequences)
  * assembly_data_report.jsonl (data report with genome assembly and annotation metadata)
  * dataset_catalog.json (a list of files and file types included in the data package)`, inputDescription),
		Example: `  datasets download genome taxon human --chromosomes 21 --include none
  datasets download genome taxon "bos taurus"
  datasets download genome taxon human --preview
  datasets download genome taxon 10116 --include rna,protein`,

		PreRunE: cmdflags.ExecutePreRunEFor(flagSets),

		RunE: func(cmd *cobra.Command, args []string) error {
			var taxons []string
			for _, arg := range iff.InputIDArgs {
				_, taxError := RetrieveTaxIdForTaxon(
					arg,
					true,
					openapi.V2ORGANISMQUERYREQUESTTAXONRESOURCEFILTER_GENOME,
					"genome",
				)
				if taxError != nil {
					return taxError
				}
				taxons = append(taxons, arg)
			}

			if dgf.downloadPreviewFlag.IsPreview() {
				request := GetGenomeReportsTaxonRequest(taxons, tem.IsTaxExactMatch())
				request.SetReturnedContent(openapi.V2ASSEMBLYDATASETREPORTSREQUESTCONTENTTYPE_ASSM_ACC)
				err := updateAssemblyReportRequestOption(request, assemblyRequestFlag)
				if err != nil {
					return err
				}
				return getDownloadSummary(dgf, NewDefaultRequestIterator(request))
			}
			downloader, warning, err := NewGenomeDownloader(GenomeWithTaxon(taxons, tem.IsTaxExactMatch()), dgf, assemblyRequestFlag)
			if warning != "" {
				cmd.PrintErrln(warning)
			}
			if err != nil {
				return err
			}
			return downloader.Download(dgf.skipValidationFlag.IsSkipValidation())
		},
	}

	cmdflags.RegisterAllFlags(flagSets, cmd.Flags())

	return cmd
}
