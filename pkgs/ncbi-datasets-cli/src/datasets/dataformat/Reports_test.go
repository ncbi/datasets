package dataformat

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/reflect/protoreflect"

	pb_reports "ncbi/datasets/v1/reports"
)

func TestReport(t *testing.T) {
	rspec1 := &ReportSpec{}
	rspec1.populateMapping()
	rspec1.addColumn(&ColSpec{Mnemonic: "a"})
	require.True(t, rspec1.hasColumn("a"))
	require.NotNil(t, rspec1.getColumn("a"))
	require.Equal(t, rspec1.getColumn("a").Mnemonic, "a")
	require.ElementsMatch(t, rspec1.GetAllFieldMnemonics(), []string{"a"})

	rspec1.addColumn(&ColSpec{Mnemonic: "c"})
	rspec1.addColumn(&ColSpec{Mnemonic: "b"})
	require.ElementsMatch(t, rspec1.GetAllFieldMnemonics(), []string{"c", "b", "a"})
	require.True(t, rspec1.hasColumns([]string{}))
	require.True(t, rspec1.hasColumns([]string{"b", "a"}))
	require.False(t, rspec1.hasColumns([]string{"a", "d"}))

	val, err := rspec1.fieldsAreValid([]string{"a"})
	require.True(t, val)
	require.NoError(t, err)

	val, err = rspec1.fieldsAreValid([]string{"a", "d", "e"})
	require.False(t, val)
	require.Error(t, err)
}

func TestAssemblyReportMapping(t *testing.T) {
	rspec := &ReportSpec{
		messageType: (*pb_reports.AssemblyDataReport)(nil),
	}
	initReportSpec(rspec, false)

	require.NotNil(t, rspec.newObject())
	require.False(t, rspec.hasColumn("a"))
	require.True(t, rspec.hasColumn("organism-name"))
	require.True(t, rspec.hasColumn("assminfo-genbank-assm-accession"))

	reportTester(
		t,
		rspec,
		&pb_reports.AssemblyDataReport{
			OrganismName: "foo",
			CommonName:   "bar",
			AnnotationInfo: &pb_reports.AnnotationInfo{
				Name:        "From INSDC submitter",
				ReleaseDate: "Sep 20, 2018",
				Source:      "CNAG",
				Stats: &pb_reports.FeatureCounts{
					GeneCounts: &pb_reports.GeneCounts{
						NonCoding:     123,
						Other:         456,
						ProteinCoding: 789,
						Total:         111,
					},
				},
			},
		},
		[]*TestFields{
			&TestFields{Mnemonic: "common-name", AllowEmpty: true},
			&TestFields{Mnemonic: "organism-name", AllowEmpty: true},
			&TestFields{Mnemonic: "annotinfo-source", AllowEmpty: true},
			&TestFields{Mnemonic: "annotinfo-featcount-gene-other", AllowEmpty: false},
			&TestFields{Mnemonic: "annotinfo-featcount-gene-protein-coding", AllowEmpty: false},
			&TestFields{
				Mnemonic:   "annotinfo-featcount-gene-non-coding",
				AllowEmpty: false,
				NonEmptyValidators: []ValidatorFunc{
					valIfNotZero,
					func(t *testing.T, val *string) {
						if val != nil {
							require.Equal(t, *val, "123")
						}
					},
				},
			},
			&TestFields{
				Mnemonic:   "annotinfo-featcount-gene-total",
				AllowEmpty: false,
				NonEmptyValidators: []ValidatorFunc{
					valIfNotZero,
					func(t *testing.T, val *string) {
						if val != nil {
							require.Equal(t, *val, "111")
						}
					},
				},
			},
		},
	)
}

type ValidatorFunc func(*testing.T, *string)
type TestFields struct {
	Mnemonic           string
	AllowEmpty         bool
	NonEmptyValidators []ValidatorFunc
}

func reportTester(t *testing.T, rspec *ReportSpec, rptObj protoreflect.ProtoMessage, testCols []*TestFields) {
	require.NotNil(t, rptObj)
	require.False(t, rspec.hasColumn("a"))

	objIter := MakeObjIter(rptObj.ProtoReflect())
	for ok := true; ok; ok = objIter.Next() {
		for _, testField := range testCols {
			if !rspec.hasColumn(testField.Mnemonic) {
				require.Equal(t, testField.Mnemonic, "unrecognized field mnemonic")
			}
			c := rspec.getColumn(testField.Mnemonic)
			v, _ := c.getValue(objIter)
			require.NotNil(t, v)
			if !testField.AllowEmpty {
				require.NotEmpty(t, v)
			}
			for _, fnVal := range testField.NonEmptyValidators {
				strVal := v.String()
				fnVal(t, &strVal)
			}
		}
	}
}

var (
	valNotEmpty = func(t *testing.T, val *string) {
		require.NotEmpty(t, *val)
	}

	valNotZero = func(t *testing.T, val *string) {
		require.NotZero(t, *val)
	}

	valIfNotZero = func(t *testing.T, val *string) {
		if val != nil && len(*val) > 0 {
			valNotZero(t, val)
		}
	}

	valIfNotEmpty = func(t *testing.T, val *string) {
		if val != nil && len(*val) > 0 {
			valNotEmpty(t, val)
		}
	}
)

func initReportSpec(rspec *ReportSpec, debug bool) {
	rspec.populateMapping()
	if !debug {
		return
	}
	for mnemonic, colspec := range rspec.mapping {
		fmt.Printf("%s [len: %d] [simple field: %t]:\n", mnemonic, len(colspec.fieldDesc), colspec.isSimpleField())
		fmt.Println(colspec.fieldDesc)
		for _, fldspec := range colspec.fieldDesc {
			fmt.Printf("  %s [repeated message: %t]\n", fldspec.name, fldspec.repeatedMessage)
		}
	}
}

func TestGeneReportMapping(t *testing.T) {
	rspec := &ReportSpec{
		messageType: (*pb_reports.GeneDescriptor)(nil),
	}
	initReportSpec(rspec, false)

	reportTester(
		t,
		rspec,
		&pb_reports.GeneDescriptor{
			GeneId:      2,
			TaxId:       1,
			Chromosomes: []string{"a", "b", "c"},
		},
		[]*TestFields{
			&TestFields{
				Mnemonic:   "gene-id",
				AllowEmpty: false,
				NonEmptyValidators: []ValidatorFunc{
					valNotEmpty,
					valNotZero,
				},
			},
			&TestFields{
				Mnemonic:   "chromosomes",
				AllowEmpty: false,
				NonEmptyValidators: []ValidatorFunc{
					func(t *testing.T, val *string) {
						if val != nil {
							require.Equal(t, *val, "a,b,c")
						}
					},
				},
			},
			&TestFields{Mnemonic: "rna-type", AllowEmpty: true},
			&TestFields{Mnemonic: "genomic-range-range-start", AllowEmpty: true},
		},
	)

	reportTester(
		t,
		rspec,
		&pb_reports.GeneDescriptor{},
		[]*TestFields{
			&TestFields{Mnemonic: "gene-id", AllowEmpty: true, NonEmptyValidators: []ValidatorFunc{valIfNotZero}},
			&TestFields{Mnemonic: "rna-type", AllowEmpty: true},
			&TestFields{Mnemonic: "genomic-range-range-start", AllowEmpty: true, NonEmptyValidators: []ValidatorFunc{valIfNotZero}},
		},
	)

	reportTester(
		t,
		rspec,
		&pb_reports.GeneDescriptor{
			GeneId: 12,
			TaxId:  1,
			GenomicRanges: []*pb_reports.SeqRangeSet{
				&pb_reports.SeqRangeSet{
					Range: []*pb_reports.Range{
						&pb_reports.Range{
							Begin: 3,
						},
						&pb_reports.Range{
							Begin: 13,
							End:   15,
						},
					},
				},
			},
			GenomicRegions: []*pb_reports.GenomicRegion{
				&pb_reports.GenomicRegion{
					GeneRange: &pb_reports.SeqRangeSet{
						Range: []*pb_reports.Range{
							&pb_reports.Range{
								Begin: 1,
							},
						},
					},
				},
				&pb_reports.GenomicRegion{
					GeneRange: &pb_reports.SeqRangeSet{
						Range: []*pb_reports.Range{
							&pb_reports.Range{
								Begin: 2,
							},
						},
					},
				},
			},
			Transcripts: []*pb_reports.Transcript{
				&pb_reports.Transcript{
					GenomicLocations: []*pb_reports.GenomicLocation{
						&pb_reports.GenomicLocation{
							Exons: []*pb_reports.Range{
								&pb_reports.Range{
									Begin: 1,
								},
							},
						},
					},
				},
				&pb_reports.Transcript{
					GenomicLocations: []*pb_reports.GenomicLocation{
						&pb_reports.GenomicLocation{
							Exons: []*pb_reports.Range{
								&pb_reports.Range{
									Begin: 2,
								},
							},
						},
					},
				},
			},
		},
		[]*TestFields{
			&TestFields{Mnemonic: "gene-id", AllowEmpty: true, NonEmptyValidators: []ValidatorFunc{valIfNotZero}},
			&TestFields{Mnemonic: "rna-type", AllowEmpty: true},
			&TestFields{
				Mnemonic:   "genomic-range-range-start",
				AllowEmpty: true,
				NonEmptyValidators: []ValidatorFunc{
					valNotZero,
					func(t *testing.T, val *string) {
						if v, err := strconv.Atoi(*val); err == nil {
							require.Greater(t, v, 2)
						}
					},
				},
			},
			&TestFields{
				Mnemonic:   "genomic-range-range-stop",
				AllowEmpty: true,
			},
			&TestFields{
				Mnemonic:   "genomic-region-gene-range-range-start",
				AllowEmpty: true,
			},
			&TestFields{
				Mnemonic:   "genomic-region-gene-range-range-stop",
				AllowEmpty: true,
			},
			&TestFields{
				Mnemonic:   "transcript-genomic-location-exon-start",
				AllowEmpty: false,
			},
		},
	)
}

func TestReportNames(t *testing.T) {
	for _, rptName := range []string{"gene", "genome", "prok-gene"} {
		require.NotNil(t, GetReport(rptName))
	}
}
