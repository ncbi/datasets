package flags

import (
	openapi "datasets/openapi/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/thediveo/enumflag/v2"
)

type VirusDownloadInclude enumflag.Flag

const (
	GENOME VirusDownloadInclude = iota
	CDS
	PROTEIN
	VIRUS_ANNOTATION
	VIRUS_BIOSAMPLE
	NONE
)

var VirusDownloadIncludeIds = map[VirusDownloadInclude][]string{
	GENOME:           {"genome"},
	CDS:              {"cds"},
	PROTEIN:          {"protein"},
	VIRUS_ANNOTATION: {"annotation"},
	VIRUS_BIOSAMPLE:  {"biosample"},
	NONE:             {"none"},
}

var VirusDownloadIncludeOpenapi = map[VirusDownloadInclude]openapi.V2ViralSequenceType{
	GENOME:  openapi.V2VIRALSEQUENCETYPE_GENOME,
	CDS:     openapi.V2VIRALSEQUENCETYPE_CDS,
	PROTEIN: openapi.V2VIRALSEQUENCETYPE_PROTEIN,
	NONE:    openapi.V2VIRALSEQUENCETYPE_NONE,
}

var VirusReportSelectionOpenapi = map[VirusDownloadInclude]openapi.V2VirusDatasetReportType{
	VIRUS_ANNOTATION: openapi.V2VIRUSDATASETREPORTTYPE_ANNOTATION,
	VIRUS_BIOSAMPLE:  openapi.V2VIRUSDATASETREPORTTYPE_BIOSAMPLE_REPORT,
}

type VirusDownloadIncludeFlag struct {
	FlagInterface
	argsVirusSequence []VirusDownloadInclude
	incSeqLongDesc    string
}

var GenomeDefault []VirusDownloadInclude = []VirusDownloadInclude{GENOME}

var ProteinDefault []VirusDownloadInclude = []VirusDownloadInclude{PROTEIN}

func NewVirusDownloadIncludeFlag(longDesc string) *VirusDownloadIncludeFlag {
	var defaultSeq []VirusDownloadInclude = GenomeDefault
	if longDesc == IncludeSequenceLongDescProtein {
		defaultSeq = ProteinDefault
	}
	vsf := &VirusDownloadIncludeFlag{
		argsVirusSequence: defaultSeq,
		incSeqLongDesc:    longDesc,
	}
	return vsf
}

const IncludeSequenceLongDescGenome string = `Specify virus genome sequence types to download
  * genome:     genomic sequences
  * cds:        nucleotide coding sequences
  * protein:    amino acid sequences
  * annotation: annotation report
  * biosample:  biosample report
  * none:       no sequence data, only primary data report
    `

const IncludeSequenceLongDescProtein string = `Specify virus genome sequence types to download
  * cds:        nucleotide coding sequences
  * protein:    amino acid sequences
  * annotation: annotation report
  * biosample:  biosample report
  * none:       no sequence data, only primary data report
    `

func (vsf *VirusDownloadIncludeFlag) RegisterFlags(flags *pflag.FlagSet) {
	flags.Var(
		enumflag.NewSlice(&vsf.argsVirusSequence, "string(,string)", VirusDownloadIncludeIds, enumflag.EnumCaseInsensitive),
		"include",
		vsf.incSeqLongDesc,
	)
}

func (vsf *VirusDownloadIncludeFlag) PreRunE(cmd *cobra.Command, args []string) (err error) {

	return nil

}

func (vsf *VirusDownloadIncludeFlag) PrepareDownloadRequest(request *openapi.V2VirusDatasetRequest) {
	var virusSeqOptions []openapi.V2ViralSequenceType
	var virusReportOptions []openapi.V2VirusDatasetReportType

	for _, v := range vsf.argsVirusSequence {
		if v == VIRUS_ANNOTATION || v == VIRUS_BIOSAMPLE {
			virusReportOptions = append(virusReportOptions, VirusReportSelectionOpenapi[v])
		} else {
			virusSeqOptions = append(virusSeqOptions, VirusDownloadIncludeOpenapi[v])
		}
	}
	request.SetIncludeSequence(virusSeqOptions)
	request.SetAuxReport(virusReportOptions)
}

func (vsf *VirusDownloadIncludeFlag) PrepareSarsProteinDownloadRequest(request *openapi.V2Sars2ProteinDatasetRequest) {
	var virusSeqOptions []openapi.V2ViralSequenceType
	var virusReportOptions []openapi.V2VirusDatasetReportType

	for _, v := range vsf.argsVirusSequence {
		if v == VIRUS_ANNOTATION || v == VIRUS_BIOSAMPLE {
			virusReportOptions = append(virusReportOptions, VirusReportSelectionOpenapi[v])
		} else {
			virusSeqOptions = append(virusSeqOptions, VirusDownloadIncludeOpenapi[v])
		}
	}
	request.SetIncludeSequence(virusSeqOptions)
	request.SetAuxReport(virusReportOptions)
}
