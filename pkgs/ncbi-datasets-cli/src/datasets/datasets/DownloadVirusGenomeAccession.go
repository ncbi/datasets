package datasets

import (
	"errors"
	"fmt"
	"os"

	"github.com/metakeule/fmtdate"
	"github.com/spf13/cobra"

	openapi "datasets_cli/v1/openapi"
)

func downloadVirusGenomeRequestWithAcc(args []string) (request openapi.ApiVirusGenomeDownloadAccessionRequest, err error) {
	cli, err := createOAClient()
	if err != nil {
		return
	}
	request = cli.VirusApi.VirusGenomeDownloadAccession(nil, args)
	return
}

func downloadVirusGenomeAccession(request openapi.ApiVirusGenomeDownloadAccessionRequest, assmFilename string) (err error) {
	request.RefseqOnly(argRefseqOnly)
	request.AnnotatedOnly(argAnnotatedOnly)
	request.ExcludeSequence(argExcludeSeq)
	request.Host(argHost)
	request.PangolinClassification(argLineage)
	request.GeoLocation(argGeoLocation)
	request.CompleteOnly(argCompleteOnly)

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

	if argReleasedSince != "" {
		date, e := fmtdate.Parse(dateFormat, argReleasedSince)
		if e != nil {
			return e
		}
		request.ReleasedSince(date)
	}

	f, e := os.Create(assmFilename)
	if e != nil {
		return fmt.Errorf("%s opening output file: %s", e, assmFilename)
	}

	_, resp, err := request.Execute()
	if err != nil {
		return err
	}
	if err = handleHTTPResponse(resp, err); err != nil {
		return err
	}
	length := int64(-1) // unknown length
	return downloadData(f, resp, err, assmFilename, length)
}

func downloadVirusGenomeAcc(cmd *cobra.Command, args []string, assmFilename string) error {
	request, err := downloadVirusGenomeRequestWithAcc(args)
	if err != nil {
		return err
	}
	return downloadVirusGenomeAccession(request, assmFilename)
}

var downloadVirusGenomeAccCmd = &cobra.Command{
	Use:   "accession <accession ...>",
	Short: "Request genome data by accessions.",
	Long: `
Download a coronavirus genome dataset by nucleotide accessions. Coronavirus genome data packages include genome,
CDS and protein sequence, annotation and a detailed data report. Datasets are downloaded as a zip file.

The default coronavirus genome dataset includes the following files (if available):
* genomic.fna (genomic sequences)
* cds.fna (nucleotide coding sequences)
* protein.faa (protein sequences)
* data_report.jsonl (data report with viral metadata)
* virus_dataset.md (README containing details on sequence file data content and other information)
* dataset_catalog.json (a list of files and file types included in the dataset)

Refer to NCBI's [download and install](https://www.ncbi.nlm.nih.gov/datasets/docs/download-and-install/) documentation for information about getting started with the command-line tools.`,
	Example: `  datasets download virus genome accession NC_045512.2`,

	RunE: func(cmd *cobra.Command, args []string) (err error) {
		if len(argIDArgs) == 0 {
			return errors.New("Input accessions not specified")
		}
		if argRetiredIncludeFlag {
			err = errors.New(virusFlagErrorMessage)
		}
		if argRetiredExcludeFlag {
			cmd.PrintErrln(virusFlagWarningMessage)
		}
		return downloadVirusGenomeAcc(cmd, argIDArgs, argDownloadFilename)
	},
}

func init() {
	downloadVirusGenomeCmd.AddCommand(downloadVirusGenomeAccCmd)

	registerHiddenStringPair(downloadVirusGenomeAccCmd.Flags(), &argInputFile, "inputfile", "i", "", "read a list of accessions from a file to use as input")
}
