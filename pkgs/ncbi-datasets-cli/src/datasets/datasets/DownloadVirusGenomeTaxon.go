package datasets

import (
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

func downloadVirusGenomeOrg(cmd *cobra.Command, taxon string, assmFilename string) error {
	cli, err := createOAClient()
	if err != nil {
		return err
	}

	if !checkTaxonWithinScope(cli, taxon) {
		fmt.Print("The download virus genome taxon command only supports coronavirus taxa.\nFor data on other viruses, please use the download genome taxon command.\n")
		return fmt.Errorf("taxa %s is out of scope", taxon)
	}

	request := cli.VirusApi.VirusGenomeDownload(nil, taxon)
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
		{argIncludeGbff, openapi.V1ANNOTATIONFORVIRUSTYPE_GENOME_GBFF},
		{!argExcludeGpff, openapi.V1ANNOTATIONFORVIRUSTYPE_GENOME_GPFF},
		{!argExcludePdb, openapi.V1ANNOTATIONFORVIRUSTYPE_PDB_FILES},
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

var downloadVirusGenomeCmd = &cobra.Command{
	Use:   "genome [command] ",
	Short: "download a coronavirus genome dataset by taxon",
	Long: `
Download a coronavirus genome dataset including genome, CDS and protein sequence, annotation
and a detailed data report. Coronavirus genome datasets are limited to the Coronaviridae family
including SARS-CoV-2. Coronavirus genome datasets can be specified by taxon. Datasets are
downloaded as a zip file.

The default coronavirus genome dataset includes the following files (if available):
* genomic.fna (genomic sequences)
* cds.fna (nucleotide coding sequences)
* protein.faa (protein sequences)
* protein.gpff (protein sequence and annotation in GenPept flat file format)
* protein structures in PDB format
* data_report.jsonl (data report with viral metadata)
* virus_dataset.md (README containing details on sequence file data content and other information)
* dataset_catalog.json (a list of files and file types included in the dataset)

Refer to NCBI's [command line quickstart](https://www.ncbi.nlm.nih.gov/datasets/docs/quickstarts/command-line-tools/) documentation for information about getting started with the command-line tools.`,
	Example: `  datasets download virus genome taxon sars-cov-2 --host dog
  datasets download virus genome taxon coronaviridae --host "manis javanica"`,
}

// virusCmd represents the virus command
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
* protein.gpff (protein sequence and annotation in GenPept flat file format)
* protein structures in PDB format
* data_report.jsonl (data report with viral metadata)
* virus_dataset.md (README containing details on sequence file data content and other information)
* dataset_catalog.json (a list of files and file types included in the dataset)

Refer to NCBI's [command line quickstart](https://www.ncbi.nlm.nih.gov/datasets/docs/quickstarts/command-line-tools/) documentation for information about getting started with the command-line tools.`,
	Example: `  datasets download virus genome taxon sars-cov-2 --host dog
  datasets download virus genome taxon coronaviridae --host "manis javanica"`,
	Args: cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		taxon := args[0]
		return downloadVirusGenomeOrg(cmd, taxon, argDownloadFilename)
	},
}

func init() {
	downloadVirusGenomeCmd.AddCommand(downloadVirusGenomeOrgCmd)

	registerHiddenBoolPair(downloadVirusGenomeOrgCmd.PersistentFlags(), &argAnnotatedOnly, "annotated", "a", false, "limit to annotated coronavirus genomes")
	registerHiddenBoolPair(downloadVirusGenomeOrgCmd.PersistentFlags(), &argCompleteOnly, "complete-only", "c", false, "limit to complete coronavirus genomes")
	downloadVirusGenomeOrgCmd.PersistentFlags().BoolVar(&argExcludeCds, "exclude-cds", false, "exclude cds.fna (CDS sequence file)")
	downloadVirusGenomeOrgCmd.PersistentFlags().BoolVar(&argExcludePdb, "exclude-pdb", false, "exclude *.pdb (protein structure files)")
	registerHiddenBoolPair(downloadVirusGenomeOrgCmd.PersistentFlags(), &argExcludeProtein, "exclude-protein", "p", false, "exclude protein.faa (protein sequence file)")
	registerHiddenBoolPair(downloadVirusGenomeOrgCmd.PersistentFlags(), &argExcludeSeq, "exclude-seq", "s", false, "exclude genomic.fna (genomic sequence file)")
	downloadVirusGenomeOrgCmd.PersistentFlags().StringVar(&argGeoLocation, "geo-location", "", "limit to coronavirus genomes isolated from a specified geographic location (continent, country or U.S. state)")
	downloadVirusGenomeOrgCmd.PersistentFlags().StringVar(&argHost, "host", "", "limit to coronavirus genomes isolated from a specified host (NCBI Taxonomy ID, scientific or common name at any taxonomic rank)")
	downloadVirusGenomeOrgCmd.PersistentFlags().StringVar(&argLineage, "lineage", "", "limit to SARS-CoV-2 genomes classified as the specified lineage (variant) by pangolin using the pangoLEARN algorithm")
	registerHiddenBoolPair(downloadVirusGenomeOrgCmd.PersistentFlags(), &argIncludeGbff, "include-gbff", "b", false, "include genomic.gbff (genome sequence and annotation in GenBank flat file format)")
	registerHiddenBoolPair(downloadVirusGenomeOrgCmd.PersistentFlags(), &argExcludeGpff, "exclude-gpff", "g", false, "exclude protein.gpff (protein sequence and annotation in GenPept flat file format")
	downloadVirusGenomeOrgCmd.PersistentFlags().BoolVar(&argRefseqOnly, "refseq", false, "limit to RefSeq coronavirus genomes")
	downloadVirusGenomeOrgCmd.PersistentFlags().StringVar(&argReleasedSince, "released-since", "", "limit to coronavirus genomes released after a specified date ("+dateFormat+")")
}
