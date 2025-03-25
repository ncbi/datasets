package datasets

import (
	cmdflags "datasets_cli/v2/datasets/flags"
	openapi "datasets/openapi/v2"
	"fmt"

	"github.com/spf13/cobra"
)

var KNOWN_VIRUS_TAXONS = []string{"2697049", "197911", "2955291", "11320"}

const MAX_VIRUS_TAXONS = 100

func createDownloadVirusGenomeTaxonCmd(dvf DownloadVirusFlag) *cobra.Command {
	svf := cmdflags.NewSkipZipValidationFlag()
	iff := cmdflags.NewInputFileFlag(cmdflags.InputFileListTypeTaxon, cmdflags.AsIntegerFalse, cmdflags.WithLimit(MAX_VIRUS_TAXONS))
	flagSets := []cmdflags.FlagInterface{svf, iff}

	inputDescription := "taxon (NCBI Taxonomy ID, scientific or common name for any virus at any tax rank)"

	cmd := &cobra.Command{
		Use:   "taxon  <taxon>",
		Short: fmt.Sprintf("Download a virus genome data package by %s", inputDescription),
		Long: fmt.Sprintf(`
Download a virus genome data package by %s. Virus genome data packages include genome, transcript and protein sequences, annotation and one or more data reports. Data packages are downloaded as a zip archive.

The default virus genome data package includes the following files:
  * genomic.fna (genomic sequences)
  * data_report.jsonl (data report with virus genome metadata)
  * dataset_catalog.json (a list of files and file types included in the data package)`, inputDescription),
		Example: `  datasets download virus genome taxon sars-cov-2 --host dog --include protein
  datasets download virus genome taxon coronaviridae --host "manis javanica"`,
		PreRunE: cmdflags.ExecutePreRunEFor(flagSets),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			// To support Cloud builds, we need to allow 2697049 + FLU straight through
			// without any checks.
			var taxons []string
			if len(args) == 1 && cmdflags.Contains(KNOWN_VIRUS_TAXONS, args[0]) {
				taxons = append(taxons, args[0])
			} else {
				taxIdsMap, taxErr := RetrieveTaxIdsForTaxons(cmd, iff.InputIDArgs, true, openapi.V2ORGANISMQUERYREQUESTTAXONRESOURCEFILTER_ALL, "virus", 10239)
				if taxErr != nil {
					return taxErr
				}
				taxons = getMapListValues(taxIdsMap)
			}

			downloader, warning, err := NewVirusDownloader(VirusDownloadWithTaxon(taxons, dvf))
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
