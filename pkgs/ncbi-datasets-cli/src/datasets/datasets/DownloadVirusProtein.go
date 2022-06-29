package datasets

import (
	"errors"
	"os"
	"time"

	"github.com/metakeule/fmtdate"
	"github.com/spf13/cobra"

	openapi "datasets_cli/v1/openapi"
)

func downloadVirusProtein(cmd *cobra.Command, proteinNames []string, assmFilename string) (err error) {

	cli, err := createOAClient()
	if err != nil {
		return
	}

	request := cli.VirusApi.Sars2ProteinDownload(nil, proteinNames)
	request.RefseqOnly(argRefseqOnly)
	request.AnnotatedOnly(argAnnotatedOnly)
	request.Host(argHost)
	request.GeoLocation(argGeoLocation)
	request.CompleteOnly(argCompleteOnly)

	if argReleasedSince != "" {
		date := time.Time{}
		date, err = fmtdate.Parse(dateFormat, argReleasedSince)
		if err != nil {
			return
		}
		request.ReleasedSince(date)
	}

	if argUpdatedSince != "" {
		date, e := fmtdate.Parse(dateFormat, argUpdatedSince)
		if e != nil {
			return e
		}
		request.UpdatedSince(date)
	}

	annotations := make([]openapi.V1AnnotationForVirusType, 0)
	possible_annotations := []struct {
		flag      bool
		type_enum openapi.V1AnnotationForVirusType
	}{
		{!argExcludeCds, openapi.V1ANNOTATIONFORVIRUSTYPE_CDS_FASTA},
		{!argExcludeProtein, openapi.V1ANNOTATIONFORVIRUSTYPE_PROT_FASTA},
	}
	for _, annot := range possible_annotations {
		if annot.flag {
			annotations = append(annotations, annot.type_enum)
		}
	}
	request.IncludeAnnotationType(annotations)

	f, err := os.Create(assmFilename)
	if err != nil {
		return
	}

	length := int64(-1) // unknown length
	_, resp, err := request.Execute()
	if err = handleHTTPResponse(resp, err); err != nil {
		return
	}
	err = downloadData(f, resp, err, assmFilename, length)
	return
}

// virusCmd represents the virus command
var downloadVirusProteinCmd = &cobra.Command{
	Use:   "protein <protein_name ...> [flags]",
	Short: "download a SARS-CoV-2 protein dataset by protein name",
	Long: `
Download a SARS-CoV-2 protein data package by protein name. SARS-CoV-2 protein
data packages include CDS and protein sequence, annotation and a detailed data report.
Data packages are downloaded as a zip file.

The default SARS-CoV-2 protein data package includes the following files:
* cds.fna (nucleotide coding sequences)
* protein.faa (protein sequences)
* data_report.jsonl (data report with viral metadata)
* virus_dataset.md (README containing details on sequence file data content and other information)
* dataset_catalog.json (a list of files and file types included in the data package)

Refer to NCBI's [download and install](https://www.ncbi.nlm.nih.gov/datasets/docs/download-and-install/) documentation for information about getting started with the command-line tools.

Allowed protein names:
* ORF1ab
* ORF1a
* nsp1
* nsp2
* nsp3
* nsp4
* nsp5
* nsp6
* nsp7
* nsp8
* nsp9
* nsp10
* rdrp
* nsp11
* nsp13
* nsp14
* nsp15
* nsp16
* S
* ORF3a
* E
* M
* ORF6
* ORF7a
* ORF7b
* ORF8
* N
* ORF10
`,
	Example: `  datasets download virus protein S --host dog --filename SARS2-spike-dog.zip
  datasets download virus protein S E M N --refseq --filename SARS2-structural-refseq.zip`,
	Args: cobra.MinimumNArgs(1),

	RunE: func(cmd *cobra.Command, args []string) (err error) {
		if argRetiredIncludeFlag {
			err = errors.New(virusFlagErrorMessage)
		}

		if argRetiredExcludeFlag {
			cmd.PrintErrln(virusFlagWarningMessage)
		}
		return downloadVirusProtein(cmd, args, argDownloadFilename)
	},
}

func init() {
	registerHiddenBoolPair(downloadVirusProteinCmd.PersistentFlags(), &argAnnotatedOnly, "annotated", "a", false, "limit to annotated coronavirus genomes")
	registerHiddenBoolPair(downloadVirusProteinCmd.PersistentFlags(), &argCompleteOnly, "complete-only", "c", false, "limit to complete coronavirus genomes")
	downloadVirusProteinCmd.PersistentFlags().BoolVar(&argExcludeCds, "exclude-cds", false, "exclude cds.fna (CDS sequence file)")
	registerHiddenBoolPair(downloadVirusProteinCmd.PersistentFlags(), &argExcludeProtein, "exclude-protein", "p", false, "exclude protein.faa (protein sequence file)")
	downloadVirusProteinCmd.PersistentFlags().StringVar(&argGeoLocation, "geo-location", "", "limit to coronavirus genomes isolated from a specified geographic location (continent, country or U.S. state)")
	downloadVirusProteinCmd.PersistentFlags().StringVar(&argHost, "host", "", "limit to coronavirus genomes isolated from a specified host (NCBI Taxonomy ID, scientific or common name at any taxonomic rank)")
	downloadVirusProteinCmd.PersistentFlags().BoolVar(&argRefseqOnly, "refseq", false, "limit to RefSeq coronavirus genomes")
	downloadVirusProteinCmd.PersistentFlags().StringVar(&argReleasedSince, "released-since", "", "limit to coronavirus genomes released after a specified date ("+dateFormat+")")
	downloadVirusProteinCmd.PersistentFlags().StringVar(&argUpdatedSince, "updated-since", "", "limit to coronavirus genomes updated after a specified date ("+dateFormat+")")
}
