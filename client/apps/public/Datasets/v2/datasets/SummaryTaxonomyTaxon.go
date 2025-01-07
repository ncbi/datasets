package datasets

import (
	cmdflags "datasets_cli/v2/datasets/flags"
	openapi "datasets/openapi/v2"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

type SummaryTaxonomyFlag struct {
	rptFlag            *cmdflags.TaxonomyReportFlag
	childrenFlag       *cmdflags.TaxChildrenFlag
	parentsFlag        *cmdflags.TaxParentsFlag
	inputFile          *cmdflags.InputFileFlag
	rankFlag           *cmdflags.TaxonRankFilterFlag
	jsonLinesLimitFlag *cmdflags.JsonLinesAndLimitFlag
	cmdFlagSet         []cmdflags.FlagInterface
}

func initSummaryTaxonomyFlag() SummaryTaxonomyFlag {
	rf := cmdflags.NewTaxonomyReportFlag()
	cf := cmdflags.NewTaxChildrenFlag()
	lf := cmdflags.NewTaxParentsFlag()
	iff := cmdflags.NewInputFileFlag(cmdflags.InputFileListTypeTaxon, cmdflags.AsIntegerFalse)
	raf := cmdflags.NewTaxonRankFilterFlag()
	jlf := cmdflags.NewJsonLineAndLimitFlag("taxonomy")

	stf := SummaryTaxonomyFlag{
		rptFlag:            rf,
		childrenFlag:       cf,
		parentsFlag:        lf,
		inputFile:          iff,
		rankFlag:           raf,
		jsonLinesLimitFlag: jlf,
		cmdFlagSet:         []cmdflags.FlagInterface{rf, cf, lf, iff, raf, jlf},
	}

	return stf
}

func createSummaryTaxonomyTaxonCmd() *cobra.Command {
	stf := initSummaryTaxonomyFlag()
	inputDescription := "taxon (NCBI Taxonomy ID, scientific or common name at any tax rank)"

	cmd := &cobra.Command{
		Use:   "taxon",
		Short: fmt.Sprintf("Print a data report containing taxonomy metadata by %s", inputDescription),
		Long: fmt.Sprintf(`
Print a data report containing taxonomy metadata by %s. The data report is returned in JSON format.`, inputDescription),
		Example: `  datasets summary taxonomy taxon human
  datasets summary taxonomy taxon "mus musculus" "human" --report names
  datasets summary taxonomy taxon "human" --children
  datasets summary taxonomy taxon "mus musculus" --rank genus
  datasets summary taxonomy taxon "human" --parents --report names
  datasets summary taxonomy taxon 10116 --report ids_only`,

		PreRunE: cmdflags.ExecutePreRunEFor(stf.cmdFlagSet),

		RunE: func(cmd *cobra.Command, args []string) (err error) {
			var taxIdsMap map[string][]string = make(map[string][]string)
			for _, taxon := range stf.inputFile.InputIDArgs {
				taxId, taxError := RetrieveTaxIdForTaxon(
					taxon,
					true,
					openapi.V2ORGANISMQUERYREQUESTTAXONRESOURCEFILTER_ALL,
					"taxonomy",
				)
				if taxError != nil {
					fmt.Fprintln(os.Stderr, taxError)
					fmt.Fprintln(os.Stderr)
				} else {
					if _, ok := taxIdsMap[taxId]; !ok {
						taxIdsMap[taxId] = []string{taxon}
					} else {
						taxIdsMap[taxId] = append(taxIdsMap[taxId], taxon)
					}
				}
			}
			if len(taxIdsMap) == 0 {
				fmt.Fprintln(os.Stderr, "No valid taxons provided for summary - exiting")
				return nil
			}

			taxIds := make([]string, len(taxIdsMap))
			i := 0
			for t := range taxIdsMap {
				taxIds[i] = t
				i++
			}

			// report error if more than one taxid is used in conjuction with children or parents flags
			if stf.childrenFlag.GetChildren() && len(taxIds) > 1 {
				return fmt.Errorf("The 'children' flag only supports a single taxon")
			}
			if stf.parentsFlag.GetParents() && len(taxIds) > 1 {
				return fmt.Errorf("The 'parents' flag only supports a single taxon")
			}
			if len(stf.rankFlag.GetRanks()) > 0 && len(taxIds) > 1 {
				return fmt.Errorf("The 'rank' flag, which automatically searches both children and parents, only supports a single taxon")
			}

			return getTaxonomySummary(taxIds, stf, taxIdsMap)
		},
	}

	cmdflags.RegisterAllFlags(stf.cmdFlagSet, cmd.PersistentFlags())

	return cmd
}
