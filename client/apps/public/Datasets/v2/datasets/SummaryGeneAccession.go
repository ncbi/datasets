package datasets

import (
	cmdflags "datasets_cli/v2/datasets/flags"
	openapi "datasets/openapi/v2"
	"fmt"

	"github.com/spf13/cobra"
)

func GeneDatasetReportRequestForAccessions(cli *openapi.APIClient, iff *cmdflags.InputFileFlag, otf *cmdflags.OrthologTaxonFilterFlag) (*openapi.V2GeneDatasetReportsRequest, error) {
	request := openapi.NewV2GeneDatasetReportsRequest()

	geneAccessions := iff.AsStringList()
	if !otf.IsOrthologRequested() {
		request.SetAccessions(geneAccessions)
		return request, nil
	}
	// Compute the gene-id list via ortholog
	// For each accession, -> gene_id
	// then for each gene_id -> ortholog gene_ids
	inputGeneInts, errAcc := RetrieveGeneIdsForAccessions(cli, geneAccessions)
	if errAcc != nil {
		return nil, errAcc
	}

	geneInts, err := RetrieveOrthologGeneIdsFor(cli, otf, inputGeneInts)
	if err != nil {
		return nil, err
	}
	if len(geneInts) == 0 {
		return nil, fmt.Errorf("No gene orthologs found for the selected accessions")
	}
	request.SetGeneIds(geneInts)
	return request, nil
}

func RetrieveGeneIdsForAccessions(cli *openapi.APIClient, geneAccessions []string) (geneInts []int32, err error) {
	api := GeneApi{geneApi: cli.GeneAPI}
	geneIdRetriever := NewGeneIdRetriever()

	request := openapi.NewV2GeneDatasetReportsRequest()
	request.SetAccessions(geneAccessions)
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

func createSummaryGeneAccessionCmd(sGeneFlag *SummaryGeneFlag) *cobra.Command {
	iff := cmdflags.NewInputFileFlag(cmdflags.InputFileListTypeGeneAcc, cmdflags.AsIntegerFalse)
	otf := cmdflags.NewOrthologTaxonFilterFlag()
	flagSets := []cmdflags.FlagInterface{iff, otf}

	cmd := &cobra.Command{
		Use:   "accession <refseq-accession ...>",
		Short: "Print a data report containing gene metadata by RefSeq nucleotide or protein accession",
		Long: `
Print a data report containing gene metadata by RefSeq nucleotide or protein accession. The data report is returned in JSON format.`,
		Example: `  datasets summary gene accession NP_000483.3
  datasets summary gene accession NM_000546.6 NM_000492.4`,
		PreRunE: cmdflags.ExecutePreRunEFor(flagSets),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			cli, cliErr := createOAClient()
			if cliErr != nil {
				return cliErr
			}

			request, err := GeneDatasetReportRequestForAccessions(cli, iff, otf)
			if err != nil {
				return err
			}

			api := GeneApi{geneApi: cli.GeneAPI}
			return geneSummaryPagePrinter(sGeneFlag, NewGeneAccessionRequestIter(request), api)
		},
	}

	cmdflags.RegisterAllFlags(flagSets, cmd.PersistentFlags())
	return cmd
}
