package datasets

import (
	openapi "datasets/openapi/v2"
	cmdflags "datasets_cli/v2/datasets/flags"
	"strings"

	"github.com/spf13/cobra"
)

func createDownloadGeneAccession() *cobra.Command {
	iff := cmdflags.NewInputFileFlag(cmdflags.InputFileListTypeGeneAcc, cmdflags.AsIntegerFalse)
	svf := cmdflags.NewSkipZipValidationFlag()
	otf := cmdflags.NewOrthologTaxonFilterFlag()
	wdf := cmdflags.NewWpDownloadFlags()
	flagSets := []cmdflags.FlagInterface{iff, otf, svf, wdf}
	// Prokaryote and non-prokaryote accessions have different defaults, so do not set them here
	dgf := newDownloadGeneFlag([]cmdflags.GeneIncludeFlags{})
	allFlags := append(dgf.cmdFlagSet, flagSets...)

	cmd := &cobra.Command{
		Use:   "accession <refseq-accession ...>",
		Short: "Download a gene data package by RefSeq nucleotide or protein accession",
		Example: `  datasets download gene accession NP_000483.3
  datasets download gene accession NM_000546.6 NM_000492.4
  datasets download gene accession WP_000769114.1`,
		Long: `
Download a gene data package by RefSeq nucleotide or protein accession. Gene data packages include gene, transcript and protein sequences and one or more data reports. Data packages are downloaded as a zip archive.

The default gene data package for NM, NR, NP, XM, XR, XP and YP accessions:
  * rna.fna (transcript sequences)
  * protein.faa (protein sequences)
  * data_report.jsonl (data report with gene metadata)
  * dataset_catalog.json (a list of files and file types included in the data package)

The default gene data package for WP accessions:
  * gene.fna (gene sequences for all genomes on which the WP is annotated)
  * protein.faa (protein sequences)
  * data_report.jsonl (data report with gene metadata)
  * dataset_catalog.json (a list of files and file types included in the data package)
  * annotation_report.jsonl (annotated locations of WP proteins on bacterial genomes)`,
		PreRunE: cmdflags.ExecutePreRunEFor(allFlags),
		RunE: func(cmd *cobra.Command, args []string) error {
			if isProkaryoteAcc(iff.AsStringList()) {

				taxIdFilter := ""
				taxError := error(nil)
				if strings.Trim(wdf.TaxonFilter(), " ") != "" {
					taxIdFilter, taxError = RetrieveTaxIdForTaxon(
						wdf.TaxonFilter(),
						true,
						openapi.V2ORGANISMQUERYREQUESTTAXONRESOURCEFILTER_ALL,
						"download gene by accession",
					)
					if taxError != nil {
						return taxError
					}
				}
				downloader, err := NewGeneProkDownloader(dgf.previewFlag.IsPreview(), taxIdFilter, dgf.geneIncludeFlag, wdf, WithProkAccessions(iff))
				if err != nil {
					return err
				}
				return downloader.Download(svf.IsSkipValidation())
			} else {
				downloader, err := NewGeneDownloader(dgf.previewFlag.IsPreview(), dgf.geneIncludeFlag, dgf.filterFlag, WithAccessions(iff, otf))
				if err != nil {
					return err
				}
				return downloader.Download(svf.IsSkipValidation())
			}
		},
	}

	cmdflags.RegisterAllFlags(flagSets, cmd.Flags())
	cmdflags.RegisterAllFlags(dgf.cmdFlagSet, cmd.Flags())

	return cmd
}

func isProkaryoteAcc(acc_list []string) bool {
	const prokPrefix string = "WP_"
	for _, accession := range acc_list {
		if !strings.HasPrefix(strings.ToUpper(accession), prokPrefix) {
			return false
		}
	}
	return true
}
