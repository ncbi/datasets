package flags

import (
	openapi "datasets/openapi/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/thediveo/enumflag"
)

type VirusDownloadReportSelection enumflag.Flag

var (
	ArgsVirusDownloadReportSelection []VirusDownloadReportSelection = []VirusDownloadReportSelection{DOWNLOAD_DATASET_REPORT}
)

const (
	DOWNLOAD_DATASET_REPORT VirusDownloadReportSelection = iota
	DOWNLOAD_ANNOTATION
)

var VirusDownloadReportSelectionIds = map[VirusDownloadReportSelection][]string{
	DOWNLOAD_DATASET_REPORT: {"DATASET_REPORT"},
	DOWNLOAD_ANNOTATION:     {"ANNOTATION"},
}

var VirusDownloadReportSelectionOpenapi = map[VirusDownloadReportSelection]openapi.V2VirusDatasetReportType{
	DOWNLOAD_DATASET_REPORT: openapi.V2VIRUSDATASETREPORTTYPE_DATASET_REPORT,
	DOWNLOAD_ANNOTATION:     openapi.V2VIRUSDATASETREPORTTYPE_ANNOTATION,
}

type VirusDownloadReportFlag struct {
	FlagInterface
	virusReport []VirusDownloadReportSelection
}

func NewVirusDownloadReportFlag() *VirusDownloadReportFlag {
	vrf := &VirusDownloadReportFlag{
		virusReport: ArgsVirusDownloadReportSelection,
	}
	return vrf
}

const reportFlagLongDesc string = `specify additional virus data report to download. Base data report (data_report.jsonl) will always be included in the download.
  * annotation: returns annotation report (annotation_report.jsonl)
    `

func (vrf *VirusDownloadReportFlag) RegisterFlags(flags *pflag.FlagSet) {
	flags.Var(
		enumflag.NewSlice(&ArgsVirusDownloadReportSelection, "string(,string)", VirusDownloadReportSelectionIds, enumflag.EnumCaseInsensitive),
		"report",
		reportFlagLongDesc,
	)
}

func (vrf *VirusDownloadReportFlag) PreRunE(cmd *cobra.Command, args []string) (err error) {

	return nil

}
