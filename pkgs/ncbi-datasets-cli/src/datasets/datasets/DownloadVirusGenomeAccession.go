package datasets

import (
	"errors"
	"fmt"
	"os"

	"github.com/metakeule/fmtdate"
	"github.com/spf13/cobra"

	openapi "datasets_cli/v1/openapi"
)

func downloadVirusGenomeAccession(accessions []string, assmFilename string) (err error) {
	request := openapi.NewV1VirusDatasetRequest()
	request.SetAccessions(accessions)

	request.SetRefseqOnly(argRefseqOnly)
	request.SetAnnotatedOnly(argAnnotatedOnly)
	request.SetExcludeSequence(argExcludeSeq)
	request.SetHost(argHost)
	request.SetPangolinClassification(argLineage)
	request.SetGeoLocation(argGeoLocation)
	request.SetCompleteOnly(argCompleteOnly)

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
	request.SetIncludeAnnotationType(annotations)

	if argReleasedSince != "" {
		date, e := fmtdate.Parse(dateFormat, argReleasedSince)
		if e != nil {
			return e
		}
		request.SetReleasedSince(date)
	}

	if argUpdatedSince != "" {
		date, e := fmtdate.Parse(dateFormat, argUpdatedSince)
		if e != nil {
			return e
		}
		request.SetUpdatedSince(date)
	}

	f, e := os.Create(assmFilename)
	if e != nil {
		return fmt.Errorf("%s opening output file: %s", e, assmFilename)
	}

	cli, err := createOAClient()
	if err != nil {
		return
	}

	_, resp, err := cli.VirusApi.VirusGenomeDownloadPost(nil, request).Execute()
	if err != nil {
		return err
	}
	if err = handleHTTPResponse(resp, err); err != nil {
		return err
	}
	length := int64(-1) // unknown length
	return downloadData(f, resp, err, assmFilename, length)
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
		return downloadVirusGenomeAccession(argIDArgs, argDownloadFilename)
	},
}

func init() {
	downloadVirusGenomeCmd.AddCommand(downloadVirusGenomeAccCmd)

	registerHiddenStringPair(downloadVirusGenomeAccCmd.Flags(), &argInputFile, "input-file", "i", "", "read a list of nucleotide accessions from a text file - file should have 1 identifier per row and no spaces or quotes")
}
