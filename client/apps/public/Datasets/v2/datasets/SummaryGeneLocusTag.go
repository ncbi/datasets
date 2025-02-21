package datasets

import (
	cmdflags "datasets_cli/v2/datasets/flags"
	openapi "datasets/openapi/v2"
	"fmt"

	"github.com/spf13/cobra"
)

func GeneDatasetReportRequestForLocusTags(cli *openapi.APIClient, iff *cmdflags.InputFileFlag, otf *cmdflags.OrthologTaxonFilterFlag) (*openapi.V2GeneDatasetReportsRequest, error) {
	request := openapi.NewV2GeneDatasetReportsRequest()

	geneLocusTags := iff.AsStringList()
	if !otf.IsOrthologRequested() {
		request.SetLocusTags(geneLocusTags)
		return request, nil
	}
	// Compute the gene-id list via ortholog
	// For each locus-tag, -> gene_id
	// then for each gene_id -> ortholog gene_ids
	inputGeneInts, errLocusTag := RetrieveGeneIdsForLocusTags(cli, geneLocusTags)
	if errLocusTag != nil {
		return nil, errLocusTag
	}

	geneInts, err := RetrieveOrthologGeneIdsFor(cli, otf, inputGeneInts)
	if err != nil {
		return nil, err
	}
	if len(geneInts) == 0 {
		return nil, fmt.Errorf("No gene orthologs found for the selected locus-tags")
	}
	request.SetGeneIds(geneInts)
	return request, nil
}

func RetrieveGeneIdsForLocusTags(cli *openapi.APIClient, geneLocusTags []string) (geneInts []int32, err error) {
	api := GeneApi{geneApi: cli.GeneAPI}
	geneIdRetriever := NewGeneIdRetriever()

	request := openapi.NewV2GeneDatasetReportsRequest()
	request.SetLocusTags(geneLocusTags)
	request.SetReturnedContent(openapi.V2GENEDATASETREPORTSREQUESTCONTENTTYPE_IDS_ONLY)
	_, err = ProcessAllPages[
		*openapi.V2GeneDatasetReportsRequest,
		openapi.V2reportsGeneDataReportPage,
		openapi.V2reportsGeneReportMatch,
		*openapi.V2reportsGeneDataReportPage](NewGeneAccessionRequestIter(request), &api, &geneIdRetriever)
	if err != nil {
		return nil, err
	}

	return geneIdRetriever.GetGeneIds(), nil
}

func createSummaryGeneLocusTagCmd(sGeneFlag *SummaryGeneFlag) *cobra.Command {
	iff := cmdflags.NewInputFileFlag(cmdflags.InputFileListTypeLocusTag, cmdflags.AsIntegerFalse)
	otf := cmdflags.NewOrthologTaxonFilterFlag()
	flagSets := []cmdflags.FlagInterface{iff, otf}

	cmd := &cobra.Command{
		Use:   "locus-tag <locus-tag ...>",
		Short: "Print a data report containing gene metadata by locus tag",
		Long: `
Print a data report containing gene metadata by locus tag. The data report is returned in JSON format.`,
		Example: `  datasets summary gene locus-tag b0001
  datasets summary gene locus-tag b0001 ArthCt125`,
		PreRunE: cmdflags.ExecutePreRunEFor(flagSets),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			cli, cliErr := createOAClient()
			if cliErr != nil {
				return cliErr
			}

			request, err := GeneDatasetReportRequestForLocusTags(cli, iff, otf)
			if err != nil {
				return err
			}

			api := GeneApi{geneApi: cli.GeneAPI}
			return geneSummaryPagePrinter(sGeneFlag, NewGeneLocusTagRequestIter(request), api)
		},
	}

	cmdflags.RegisterAllFlags(flagSets, cmd.PersistentFlags())
	return cmd
}
