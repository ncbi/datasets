package flags

import (
	"fmt"

	openapi "datasets/openapi/v2"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/thediveo/enumflag/v2"
)

type GeneIncludeFlags enumflag.Flag

var (
	GeneDefault               []GeneIncludeFlags = []GeneIncludeFlags{Rna, Protein}
	GeneFastaDefault          []openapi.V2Fasta  = []openapi.V2Fasta{openapi.V2FASTA_RNA, openapi.V2FASTA_PROTEIN}
	GeneAccessionFastaDefault []openapi.V2Fasta  = []openapi.V2Fasta{openapi.V2FASTA_GENE, openapi.V2FASTA_PROTEIN}
)

const (
	Gene GeneIncludeFlags = iota
	Rna
	Protein
	Cds
	FivePrimeUtr
	ThreePrimeUtr
	ProductReport
	GeneIncludeNoneFlag
)

var GeneIncludeFlagIds = map[GeneIncludeFlags][]string{
	Gene:                {"gene"},
	Rna:                 {"rna"},
	Protein:             {"protein"},
	Cds:                 {"cds"},
	FivePrimeUtr:        {"5p-utr"},
	ThreePrimeUtr:       {"3p-utr"},
	ProductReport:       {"product-report"},
	GeneIncludeNoneFlag: {"None"},
}

var GeneIncludeFlagsOpenapi = map[GeneIncludeFlags]openapi.V2Fasta{
	Gene:          openapi.V2FASTA_GENE,
	Rna:           openapi.V2FASTA_RNA,
	Protein:       openapi.V2FASTA_PROTEIN,
	Cds:           openapi.V2FASTA_CDS,
	FivePrimeUtr:  openapi.V2FASTA__5_P_UTR,
	ThreePrimeUtr: openapi.V2FASTA__3_P_UTR,
}

type GeneIncludeFlag struct {
	FlagInterface
	IncludeAnnotation []GeneIncludeFlags
}

func NewGeneIncludeFlag(defaultGeneIncludeFlag []GeneIncludeFlags) *GeneIncludeFlag {
	gif := &GeneIncludeFlag{
		IncludeAnnotation: defaultGeneIncludeFlag,
	}
	return gif
}

func (gif *GeneIncludeFlag) RegisterFlags(flags *pflag.FlagSet) {
	flags.Var(
		enumflag.NewSlice(&((*gif).IncludeAnnotation), "string(,string)", GeneIncludeFlagIds, enumflag.EnumCaseInsensitive),
		"include",
		`Specify the data files to include (comma-separated).
  * gene:           gene sequence
  * rna:            transcript
  * protein:        amino acid sequences
  * cds:            nucleotide coding sequences
  * 5p-utr:         5'-UTR
  * 3p-utr:         3'-UTR
  * product-report: gene transcript and protein locations and metadata
  * none:           do not retrieve any sequence files
  `)
}

func (gif *GeneIncludeFlag) PreRunE(cmd *cobra.Command, args []string) (err error) {
	return nil
}

func (gif *GeneIncludeFlag) SetProkDownloadFlags(request *openapi.V2ProkaryoteGeneRequest) (err error) {
	// set defaults here because prok and non-prok gene-accession download have different defaults
	if len(gif.IncludeAnnotation) == 0 {
		request.SetIncludeAnnotationType(GeneAccessionFastaDefault)
		return
	}

	annotations := make([]openapi.V2Fasta, 0)
	for _, fl := range gif.IncludeAnnotation {
		if fl == GeneIncludeNoneFlag {
			continue
		}
		if fl != Gene && fl != Protein {
			return fmt.Errorf("File format %s is not supported for prokaryotic (WP_) downloads", GeneIncludeFlagIds[fl][0])
		}
		annotations = append(annotations, GeneIncludeFlagsOpenapi[fl])
	}
	request.SetIncludeAnnotationType(annotations)
	return
}

func (gif *GeneIncludeFlag) SetGeneDownloadFlags(request *openapi.V2GeneDatasetRequest) {
	// set defaults here because prok and non-prok gene-accession download have different defaults
	if len(gif.IncludeAnnotation) == 0 {
		request.SetIncludeAnnotationType(GeneFastaDefault)
		return
	}

	annotations := make([]openapi.V2Fasta, 0)
	for _, fl := range gif.IncludeAnnotation {
		if fl == GeneIncludeNoneFlag {
			continue
		}
		if fl == ProductReport {
			request.SetAuxReport(append(request.GetAuxReport(), openapi.V2GENEDATASETREQUESTGENEDATASETREPORTTYPE_PRODUCT_REPORT))
			continue
		}
		annotations = append(annotations, GeneIncludeFlagsOpenapi[fl])
	}
	request.SetIncludeAnnotationType(annotations)
}
