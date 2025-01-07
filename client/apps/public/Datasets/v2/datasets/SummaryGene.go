package datasets

import (
	cmdflags "datasets_cli/v2/datasets/flags"
	openapi "datasets/openapi/v2"

	"github.com/spf13/cobra"
	"github.com/thediveo/enumflag"
)

type SummaryGeneFlag struct {
	jsonLinesLimitFlag *cmdflags.JsonLinesAndLimitFlag
	cmdFlagSet         []cmdflags.FlagInterface
}

func initSummaryGeneFlag() *SummaryGeneFlag {

	jlf := cmdflags.NewJsonLineAndLimitFlag("gene")
	sGeneFlag := &SummaryGeneFlag{
		jsonLinesLimitFlag: jlf,
		cmdFlagSet:         []cmdflags.FlagInterface{jlf},
	}
	return sGeneFlag
}

func NewGeneAccessionRequestIter(request *openapi.V2GeneDatasetReportsRequest) *DefaultPagedRequestIterator[*openapi.V2GeneDatasetReportsRequest, string] {
	accRequester := NewDefaultPagedRequestIterator[*openapi.V2GeneDatasetReportsRequest, string](request,
		func(request *openapi.V2GeneDatasetReportsRequest, accessions []string) {
			request.SetAccessions(accessions)
		},
	)

	if len(request.GetAccessions()) > PAGE_ITER_THRESHOLD {
		accRequester.ids = request.GetAccessions()
		request.SetAccessions([]string{})
	}

	return accRequester
}

func NewGeneIdRequestIter(request *openapi.V2GeneDatasetReportsRequest) *DefaultPagedRequestIterator[*openapi.V2GeneDatasetReportsRequest, int32] {
	geneIdRequester := NewDefaultPagedRequestIterator[*openapi.V2GeneDatasetReportsRequest, int32](request,
		func(request *openapi.V2GeneDatasetReportsRequest, geneIds []int32) { request.SetGeneIds(geneIds) },
	)

	if len(request.GetGeneIds()) > PAGE_ITER_THRESHOLD {
		geneIdRequester.ids = request.GetGeneIds()
		request.SetGeneIds([]int32{})
	}

	return geneIdRequester
}

func NewGeneSymbolRequestIter(request *openapi.V2GeneDatasetReportsRequest) *DefaultPagedRequestIterator[*openapi.V2GeneDatasetReportsRequest, string] {
	symbolRequester := NewDefaultPagedRequestIterator[*openapi.V2GeneDatasetReportsRequest, string](request,
		func(request *openapi.V2GeneDatasetReportsRequest, symbols []string) {
			request.SymbolsForTaxon.SetSymbols(symbols)
		},
	)

	if len(request.SymbolsForTaxon.GetSymbols()) > PAGE_ITER_THRESHOLD {
		symbolRequester.ids = request.SymbolsForTaxon.GetSymbols()
		request.SymbolsForTaxon.SetSymbols([]string{})
	}

	return symbolRequester
}

func createCheckGeneIdDuplicates() func(openapi.V2reportsGeneReportMatch, *map[string]bool) bool {
	return func(report openapi.V2reportsGeneReportMatch, duplicatesMap *map[string]bool) bool {
		geneId := ""
		if report.HasGene() {
			geneId = report.Gene.GetGeneId()
		} else if report.HasProduct() {
			geneId = report.Product.GetGeneId()
		}
		if (*duplicatesMap)[geneId] {
			return true
		}
		(*duplicatesMap)[geneId] = true
		return false
	}
}

func geneSummaryPagePrinter(sGeneFlag *SummaryGeneFlag, requestIter RequestIterator[*openapi.V2GeneDatasetReportsRequest], api GeneApi) (err error) {
	requestIter.GetRequest().SetReturnedContent(openapi.V2GeneDatasetReportsRequestContentType(GeneReportModeIds[argGeneReportMode][0]))
	if sGeneFlag.jsonLinesLimitFlag.CountOnly() {
		requestIter.GetRequest().SetReturnedContent(openapi.V2GENEDATASETREPORTSREQUESTCONTENTTYPE_IDS_ONLY)
	}

	// Look for duplicates in gene paging, since multiple gene ids can map to the same gene (used here in all cases, but not needed for queries
	// that do not do client-side paging)
	pagePrinter := NewPagePrinter[openapi.V2reportsGeneReportMatch, *openapi.V2reportsGeneDataReportPage](
		"gene",
		sGeneFlag.jsonLinesLimitFlag.JsonLines(),
		WithPrintJsonLinesResult[openapi.V2reportsGeneReportMatch, *openapi.V2reportsGeneDataReportPage](createJsonLinesPrintDescriptor(argGeneReportMode)),
		WithCheckDuplicates[openapi.V2reportsGeneReportMatch, *openapi.V2reportsGeneDataReportPage](createCheckGeneIdDuplicates()),
	)

	_, err = ProcessPages[*openapi.V2GeneDatasetReportsRequest,
		openapi.V2reportsGeneDataReportPage,
		openapi.V2reportsGeneReportMatch,
		*openapi.V2reportsGeneDataReportPage](requestIter, &api, &pagePrinter, sGeneFlag.jsonLinesLimitFlag.RetrievalCount(), sGeneFlag.jsonLinesLimitFlag.CountOnly())

	return err
}

func createSummaryGeneCmd() *cobra.Command {

	sGeneFlag := initSummaryGeneFlag()

	cmd := &cobra.Command{
		Use:   "gene",
		Short: "Print a summary of a gene dataset",
		Long: `
Print a data report containing gene metadata.  The data report is returned in JSON format.`,
		Example: `  datasets summary gene gene-id 672
  datasets summary gene symbol brca1 --taxon "mus musculus"
  datasets summary gene accession NP_000483.3`,
		Args:              cobra.NoArgs,
		PersistentPreRunE: cmdflags.ExecutePreRunEFor(sGeneFlag.cmdFlagSet),
		RunE:              ParentCommandRunE,
	}

	cmdflags.RegisterAllFlags(sGeneFlag.cmdFlagSet, cmd.PersistentFlags())

	cmd.AddCommand(createSummaryGeneIdCmd(sGeneFlag))
	cmd.AddCommand(createSummaryGeneSymbolCmd(sGeneFlag))
	cmd.AddCommand(createSummaryGeneAccessionCmd(sGeneFlag))
	cmd.AddCommand(createSummaryGeneTaxonCmd(sGeneFlag))

	pflags := cmd.PersistentFlags()

	pflags.Var(
		enumflag.New(&argGeneReportMode, "string", GeneReportModeIds, enumflag.EnumCaseInsensitive),
		"report",
		`Choose the output type:
  * gene:     Retrieve the primary gene report
  * product:  Retrieve product data report
  * ids_only: Only retrieve gene-ids
    `)

	return cmd
}

// Define your new enum flag type. It can be derived from enumflag.Flag, but
// it doesn't need to be as long as it is compatible with enumflag.Flag, so
// either an int or uint.
type GeneReportMode enumflag.Flag

// Define the enumeration values for ReportMode.
const (
	Data GeneReportMode = iota
	Product
	IdsOnly
)

// Map enumeration values to their textual representations (value
// identifiers).
var GeneReportModeIds = map[GeneReportMode][]string{
	Data:    {"COMPLETE", "GENE"},
	Product: {"PRODUCT"},
	IdsOnly: {"IDS_ONLY"},
}

var argGeneReportMode GeneReportMode
