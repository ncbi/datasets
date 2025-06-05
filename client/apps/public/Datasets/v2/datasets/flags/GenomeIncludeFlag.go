package flags

import (
	openapi "datasets/openapi/v2"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/thediveo/enumflag/v2"
)

type GenomeIncludeAnnotation enumflag.Flag

const (
	GenomeGff3 GenomeIncludeAnnotation = iota
	GenomeGbff
	GenomeGtf
	GenomeProtein
	GenomeRna
	GenomeCds
	GenomeSeq
	GenomeSequenceReport
	GenomeIncludeNoneFlag
)

// Map enumeration values to their textual representations (value
// identifiers).
var GenomeIncludeAnnotationIds = map[GenomeIncludeAnnotation][]string{
	GenomeGff3:            {"gff3"},
	GenomeGbff:            {"gbff"},
	GenomeGtf:             {"gtf"},
	GenomeProtein:         {"protein"},
	GenomeRna:             {"rna"},
	GenomeCds:             {"cds"},
	GenomeSeq:             {"genome"},
	GenomeSequenceReport:  {"seq-report"},
	GenomeIncludeNoneFlag: {"none"},
}

var GenomeIncludeAnnotationOpenapi = map[GenomeIncludeAnnotation]openapi.V2AnnotationForAssemblyType{
	GenomeGff3:           openapi.V2ANNOTATIONFORASSEMBLYTYPE_GENOME_GFF,
	GenomeGbff:           openapi.V2ANNOTATIONFORASSEMBLYTYPE_GENOME_GBFF,
	GenomeGtf:            openapi.V2ANNOTATIONFORASSEMBLYTYPE_GENOME_GTF,
	GenomeProtein:        openapi.V2ANNOTATIONFORASSEMBLYTYPE_PROT_FASTA,
	GenomeRna:            openapi.V2ANNOTATIONFORASSEMBLYTYPE_RNA_FASTA,
	GenomeCds:            openapi.V2ANNOTATIONFORASSEMBLYTYPE_CDS_FASTA,
	GenomeSeq:            openapi.V2ANNOTATIONFORASSEMBLYTYPE_GENOME_FASTA,
	GenomeSequenceReport: openapi.V2ANNOTATIONFORASSEMBLYTYPE_SEQUENCE_REPORT,
}

type GenomeIncludeAnnotationFlag struct {
	FlagInterface
	IncludeAnnotation []GenomeIncludeAnnotation
}

func NewGenomeIncludeAnnotationFlag(defaultGenomeIncludeAnnotationFlag []GenomeIncludeAnnotation) *GenomeIncludeAnnotationFlag {
	giaf := &GenomeIncludeAnnotationFlag{
		IncludeAnnotation: defaultGenomeIncludeAnnotationFlag,
	}
	return giaf
}

func (giaf *GenomeIncludeAnnotationFlag) RegisterFlags(flags *pflag.FlagSet) {
	flags.Var(
		enumflag.NewSlice(&((*giaf).IncludeAnnotation), "string(,string)", GenomeIncludeAnnotationIds, enumflag.EnumCaseInsensitive),
		"include",
		`Specify the data files to include (comma-separated).
  * genome:     genomic sequence
  * rna:        transcript
  * protein:    amnio acid sequences
  * cds:        nucleotide coding sequences
  * gff3:       general feature file
  * gtf:        gene transfer format
  * gbff:       GenBank flat file
  * seq-report: sequence report file
  * none:       do not retrieve any sequence files
  `)
}

func (giaf *GenomeIncludeAnnotationFlag) PreRunE(cmd *cobra.Command, args []string) (err error) {
	return nil
}
