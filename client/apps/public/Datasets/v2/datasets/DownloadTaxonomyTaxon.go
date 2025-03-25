package datasets

import (
	"fmt"

	cmdflags "datasets_cli/v2/datasets/flags"
	openapi "datasets/openapi/v2"

	"github.com/spf13/cobra"
)

type DownloadTaxonomyFlag struct {
	rptFlag        *cmdflags.TaxonomyDownloadIncludeFlag
	childrenFlag   *cmdflags.TaxChildrenFlag
	parentsFlag    *cmdflags.TaxParentsFlag
	inputFile      *cmdflags.InputFileFlag
	skipValidation *cmdflags.SkipZipValidation
	rankFlag       *cmdflags.TaxonRankFilterFlag

	cmdFlagSet []cmdflags.FlagInterface
}

func initDownloadTaxonomyFlag() DownloadTaxonomyFlag {
	rf := cmdflags.NewTaxonomyDownloadIncludeFlag()
	cf := cmdflags.NewTaxChildrenFlag()
	lf := cmdflags.NewTaxParentsFlag()
	iff := cmdflags.NewInputFileFlag(cmdflags.InputFileListTypeTaxon, cmdflags.AsIntegerFalse)
	sv := cmdflags.NewSkipZipValidationFlag()
	raf := cmdflags.NewTaxonRankFilterFlag()

	stf := DownloadTaxonomyFlag{
		rptFlag:        rf,
		childrenFlag:   cf,
		parentsFlag:    lf,
		inputFile:      iff,
		skipValidation: sv,
		rankFlag:       raf,
		cmdFlagSet:     []cmdflags.FlagInterface{rf, cf, lf, iff, sv, raf},
	}

	return stf
}

func createDownloadTaxonomyTaxonCmd() *cobra.Command {
	downloadTaxonomyFlags := initDownloadTaxonomyFlag()

	inputDescription := "taxonomy taxon (NCBI Taxonomy ID, scientific or common name at any tax rank)"
	cmd := &cobra.Command{
		Use:   "taxon <taxon...>",
		Short: fmt.Sprintf("Download a taxonomy package by %s", inputDescription),
		Long: fmt.Sprintf(`
Download a taxonomy data package by %s.

The default taxonomy data package includes the following files:
  * taxonomy_report.jsonl
  * taxonomy_summary.tsv
  * dataset_catalog.json (a list of files and file types included in the data package)
A taxonomy names data report can also be added to the package
  * names_report.jsonl`, inputDescription),
		Example: `  datasets download taxonomy taxon "bos taurus"
  datasets download taxonomy taxon human,"drosophila melanogaster" --include names
  datasets download taxonomy taxon 10116 --parents --children`,

		PreRunE: cmdflags.ExecutePreRunEFor(downloadTaxonomyFlags.cmdFlagSet),

		RunE: func(cmd *cobra.Command, args []string) error {

			// Convert all taxons to valid taxids (and save as a set) before calling download
			var taxIdsMap, taxErr = RetrieveTaxIdsForTaxons(
				cmd,
				downloadTaxonomyFlags.inputFile.InputIDArgs,
				true,
				openapi.V2ORGANISMQUERYREQUESTTAXONRESOURCEFILTER_ALL,
				"taxonomy",
			)
			if taxErr != nil {
				return taxErr
			}

			taxIds, err := strToInt32ListErr(getMapKeys(taxIdsMap))
			if err != nil {
				return err
			}

			// report error if more than one taxid is used in conjuction with children flag.
			if downloadTaxonomyFlags.childrenFlag.GetChildren() && len(taxIds) > 1 {
				return fmt.Errorf("The 'children' flag only supports a single taxon")
			}
			if downloadTaxonomyFlags.parentsFlag.GetParents() && len(taxIds) > 1 {
				return fmt.Errorf("The 'parents' flag only supports a single taxon")
			}
			if len(downloadTaxonomyFlags.rankFlag.GetRanks()) > 0 && len(taxIds) > 1 {
				return fmt.Errorf("The 'rank' flag, which automatically searches both children and parents, only supports a single taxon")
			}

			downloader, warning, err := NewTaxonomyDownloader(taxIds, downloadTaxonomyFlags)
			if warning != "" {
				cmd.PrintErrln(warning)
			}
			if err != nil {
				return err
			}
			return downloader.Download(downloadTaxonomyFlags.skipValidation.IsSkipValidation())
		},
	}

	cmdflags.RegisterAllFlags(downloadTaxonomyFlags.cmdFlagSet, cmd.PersistentFlags())

	return cmd
}
