package datasets

import (
	"errors"
	"fmt"
	"os"

	"github.com/metakeule/fmtdate"
	"github.com/spf13/cobra"

	openapi "datasets_cli/v1/openapi"
)

func checkTaxonWithinScope(cli *openapi.APIClient, taxon string) (is_cov bool) {
	// conservatively set default to true
	is_cov = true
	query_min_ord, query_max_ord, err := findTaxonOrd(cli, taxon)
	if err == nil {
		cov_min_ord, cov_max_ord, err := findTaxonOrd(cli, "Coronaviridae")
		if err == nil && (query_min_ord < cov_min_ord || query_max_ord > cov_max_ord) {
			is_cov = false
		}
	}
	return
}

func findTaxonOrd(cli *openapi.APIClient, taxon string) (min_ord int32, max_ord int32, err error) {
	min_ord = 0
	max_ord = 0
	request := cli.GeneApi.GeneTaxTree(nil, taxon)
	request.ChildrenOnly(false)
	org, resp, err := request.Execute()
	if err = handleHTTPResponse(resp, err); err == nil {
		min_ord = org.GetMinOrd()
		max_ord = org.GetMaxOrd()
	}
	return
}

func downloadVirusGenomeRequestWithTaxon(taxon string) (request openapi.ApiVirusGenomeDownloadRequest, err error) {
	cli, err := createOAClient()
	if err != nil {
		return
	}
	if !checkTaxonWithinScope(cli, taxon) {
		fmt.Print("The download virus genome taxon command only supports coronavirus taxa.\nFor data on other viruses, please use the download genome taxon command.\n")
		err = fmt.Errorf("taxa %s is out of scope", taxon)
		return
	}

	request = cli.VirusApi.VirusGenomeDownload(nil, taxon)
	return
}

func downloadVirusGenome(request openapi.ApiVirusGenomeDownloadRequest, assmFilename string) (err error) {
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

	if argUpdatedSince != "" {
		date, e := fmtdate.Parse(dateFormat, argUpdatedSince)
		if e != nil {
			return e
		}
		request.UpdatedSince(date)
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

func downloadVirusGenomeOrg(cmd *cobra.Command, taxon string, assmFilename string) error {
	request, err := downloadVirusGenomeRequestWithTaxon(taxon)
	if err != nil {
		return err
	}
	return downloadVirusGenome(request, assmFilename)
}

var downloadVirusGenomeOrgCmd = &cobra.Command{
	Use:   "taxon  <taxon>",
	Short: "Request genome data by taxonomic id or name. Allowed taxon are limited to all taxa under Coronaviridae, e.g. sars2 or betacoronavirus",
	Long: `
Download a coronavirus genome dataset by taxon (NCBI Taxonomy ID, scientific or common name
for any taxonomic group in the coronavirus family). Coronavirus genome data packages include genome,
CDS and protein sequence, annotation and a detailed data report. Datasets are downloaded as a zip file.

The default coronavirus genome dataset includes the following files (if available):
* genomic.fna (genomic sequences)
* cds.fna (nucleotide coding sequences)
* protein.faa (protein sequences)
* data_report.jsonl (data report with viral metadata)
* virus_dataset.md (README containing details on sequence file data content and other information)
* dataset_catalog.json (a list of files and file types included in the dataset)

Refer to NCBI's [download and install](https://www.ncbi.nlm.nih.gov/datasets/docs/download-and-install/) documentation for information about getting started with the command-line tools.`,
	Example: `  datasets download virus genome taxon sars-cov-2 --host dog
  datasets download virus genome taxon coronaviridae --host "manis javanica"`,
	Args: cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) (err error) {
		taxon := args[0]
		if argRetiredIncludeFlag {
			err = errors.New(virusFlagErrorMessage)
		}

		if argRetiredExcludeFlag {
			cmd.PrintErrln(virusFlagWarningMessage)
		}
		return downloadVirusGenomeOrg(cmd, taxon, argDownloadFilename)
	},
}

func init() {
	downloadVirusGenomeCmd.AddCommand(downloadVirusGenomeOrgCmd)
}
