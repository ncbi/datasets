package datasets

import (
	"fmt"
	"os"

	"github.com/antihax/optional"
	"github.com/metakeule/fmtdate"
	"github.com/spf13/cobra"

	openapi "main/openapi_client"
)

func getVirusIncludeAnnotationTypesGenome() []openapi.V1alpha1AnnotationForVirusType {
	var argAnnotTypes []openapi.V1alpha1AnnotationForVirusType
	if !argExcludeProtein {
		argAnnotTypes = append(argAnnotTypes, openapi.V1ALPHA1ANNOTATIONFORVIRUSTYPE_CDS_FASTA)
	}
	if !argExcludeProtein {
		argAnnotTypes = append(argAnnotTypes, openapi.V1ALPHA1ANNOTATIONFORVIRUSTYPE_PROT_FASTA)
	}
	if argIncludeGbff {
		argAnnotTypes = append(argAnnotTypes, openapi.V1ALPHA1ANNOTATIONFORVIRUSTYPE_GENOME_GBFF)
	}
	if !argExcludeGpff {
		argAnnotTypes = append(argAnnotTypes, openapi.V1ALPHA1ANNOTATIONFORVIRUSTYPE_GENOME_GPFF)
	}
	if !argExcludePdb {
		argAnnotTypes = append(argAnnotTypes, openapi.V1ALPHA1ANNOTATIONFORVIRUSTYPE_PDB_FILES)
	}
	return argAnnotTypes
}

func virusGenomeDownloadOpts(cmd *cobra.Command) (opts *openapi.VirusGenomeDownloadOpts, err error) {
	opts = new(openapi.VirusGenomeDownloadOpts)
	opts.RefseqOnly = optional.NewBool(argRefseqOnly)
	opts.AnnotatedOnly = optional.NewBool(argAnnotatedOnly)
	opts.ExcludeSequence = optional.NewBool(argExcludeSeq)
	argAnnotTypes := getVirusIncludeAnnotationTypesGenome()
	if len(argAnnotTypes) > 0 {
		opts.IncludeAnnotationType = optional.NewInterface(argAnnotTypes)
	}
	if argReleasedSince != "" {
		date, e := fmtdate.Parse(dateFormat, argReleasedSince)
		if e != nil {
			err = e
			return
		}
		opts.ReleasedSince = optional.NewTime(date)
	}
	opts.Host = optional.NewString(argHost)
	opts.PangolinClassification = optional.NewString(argLineage)
	opts.GeoLocation = optional.NewString(argGeoLocation)
	opts.CompleteOnly = optional.NewBool(argCompleteOnly)
	return
}

func checkCovTaxon(taxon string) (is_cov bool) {
	cli, err := createOAClient()
	// conservatively set default to true
	is_cov = true
	if err == nil {
		query_min_ord, query_max_ord, err := findTaxonOrd(cli, taxon)
		if err == nil {
			cov_min_ord, cov_max_ord, err := findTaxonOrd(cli, "Coronaviridae")
			if err == nil && (query_min_ord < cov_min_ord || query_max_ord > cov_max_ord) {
				is_cov = false
			}
		}
	}
	return
}

func findTaxonOrd(cli *openapi.APIClient, taxon string) (min_ord int64, max_ord int64, err error) {
	min_ord = 0
	max_ord = 0
	org, resp, err := cli.GeneApi.GeneTaxTree(nil, taxon)
	if err = handleHTTPResponse(resp, err); err == nil {
		min_ord = org.MinOrd
		max_ord = org.MaxOrd
	}
	return
}

func downloadVirusGenomeOrg(opts *openapi.VirusGenomeDownloadOpts, taxon string, assmFilename string) error {
	taxon_within_cov := checkCovTaxon(taxon)
	if !taxon_within_cov {
		fmt.Print("The download virus genome taxon command only supports coronavirus taxa.\nFor data on other viruses, please use the download genome taxon command.\n")
		return fmt.Errorf("taxa %s is out of scope", taxon)
	}
	f, e := os.Create(assmFilename)
	if e != nil {
		return fmt.Errorf("%s opening output file: %s", e, assmFilename)
	}
	cli, err := createOAClient()
	if err != nil {
		return err
	}
	_, resp, err := cli.VirusApi.VirusGenomeDownload(nil, taxon, opts)
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
	Example: `  datasets download virus genome taxon sars-cov-2 --released-since 05/05/2021
  datasets download virus genome taxon coronaviridae --host "manis javanica"`,
}

// virusCmd represents the virus command
var downloadVirusGenomeOrgCmd = &cobra.Command{
	Use:   "taxon  <taxon>",
	Short: "Request genome data by taxonomic id or name. Allowed taxon are limited to all taxa under Coronaviridae, e.g. sars2 or betacoronavirus",
	Long: `
Download a coronavirus genome dataset by taxon (NCBI Taxonomy ID, scientific or common name
for any taxonomic group in the coronavirus family). Coronavirus genome datasets include genome,
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
	Example: `  datasets download virus genome taxon sars-cov-2 --released-since 05/05/2021
  datasets download virus genome taxon coronaviridae --host "manis javanica"`,
	Args: cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		if opts, err := virusGenomeDownloadOpts(cmd); err == nil {
			taxon := args[0]
			return downloadVirusGenomeOrg(opts, taxon, argAssmFilename)
		} else {
			return err
		}
	},
}

func init() {
	downloadVirusGenomeCmd.AddCommand(downloadVirusGenomeOrgCmd)

	registerHiddenBoolPair(downloadVirusGenomeOrgCmd.PersistentFlags(), &argAnnotatedOnly, "annotated", "a", false, "limit to annotated coronavirus genomes")
	registerHiddenBoolPair(downloadVirusGenomeOrgCmd.PersistentFlags(), &argCompleteOnly, "complete-only", "c", false, "limit to complete coronavirus genomes")
	downloadVirusGenomeOrgCmd.PersistentFlags().BoolVar(&argExcludeProtein, "exclude-cds", false, "exclude cds.fna (CDS sequence file)")
	downloadVirusGenomeOrgCmd.PersistentFlags().BoolVar(&argExcludePdb, "exclude-pdb", false, "exclude *.pdb (protein structure files)")
	registerHiddenBoolPair(downloadVirusGenomeOrgCmd.PersistentFlags(), &argExcludeProtein, "exclude-protein", "p", false, "exclude protein.faa (protein sequence file)")
	registerHiddenBoolPair(downloadVirusGenomeOrgCmd.PersistentFlags(), &argExcludeSeq, "exclude-seq", "s", false, "exclude genomic.fna (genomic sequence file)")
	registerHiddenStringPair(downloadVirusGenomeOrgCmd.PersistentFlags(), &argAssmFilename, "filename", "f", "ncbi_dataset.zip", "specify a custom file name for the downloaded dataset")
	downloadVirusGenomeOrgCmd.PersistentFlags().StringVar(&argGeoLocation, "geo-location", "", "limit to coronavirus genomes isolated from a specified geographic location (continent, country or U.S. state)")
	downloadVirusGenomeOrgCmd.PersistentFlags().StringVar(&argHost, "host", "", "limit to coronavirus genomes isolated from a specified host (NCBI Taxonomy ID, scientific or common name at any taxonomic rank)")
	downloadVirusGenomeOrgCmd.PersistentFlags().StringVar(&argLineage, "lineage", "", "limit to SARS-CoV-2 genomes classified as the specified lineage (variant) by pangolin using the pangoLEARN algorithm")
	registerHiddenBoolPair(downloadVirusGenomeOrgCmd.PersistentFlags(), &argIncludeGbff, "include-gbff", "b", false, "include genomic.gbff (genome sequence and annotation in GenBank flat file format)")
	registerHiddenBoolPair(downloadVirusGenomeOrgCmd.PersistentFlags(), &argExcludeGpff, "exclude-gpff", "g", false, "exclude protein.gpff (protein sequence and annotation in GenPept flat file format")
	downloadVirusGenomeOrgCmd.PersistentFlags().BoolVar(&argRefseqOnly, "refseq", false, "limit to RefSeq coronavirus genomes")
	downloadVirusGenomeOrgCmd.PersistentFlags().StringVar(&argReleasedSince, "released-since", "", "limit to coronavirus genomes released after a specified date ("+dateFormat+")")
}
