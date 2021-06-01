package datasets

import (
	"os"

	"github.com/antihax/optional"
	"github.com/metakeule/fmtdate"
	"github.com/spf13/cobra"

	openapi "main/openapi_client"
)

func getVirusIncludeAnnotationTypesProtein() []openapi.V1alpha1AnnotationForVirusType {
	var argAnnotTypes []openapi.V1alpha1AnnotationForVirusType
	if !argExcludeProtein {
		argAnnotTypes = append(argAnnotTypes, openapi.V1ALPHA1ANNOTATIONFORVIRUSTYPE_CDS_FASTA)
	}
	if !argExcludeProtein {
		argAnnotTypes = append(argAnnotTypes, openapi.V1ALPHA1ANNOTATIONFORVIRUSTYPE_PROT_FASTA)
	}
	if !argExcludeGpff {
		argAnnotTypes = append(argAnnotTypes, openapi.V1ALPHA1ANNOTATIONFORVIRUSTYPE_GENOME_GPFF)
	}
	if !argExcludePdb {
		argAnnotTypes = append(argAnnotTypes, openapi.V1ALPHA1ANNOTATIONFORVIRUSTYPE_PDB_FILES)
	}
	return argAnnotTypes
}

func sars2ProteinDownloadOpts(cmd *cobra.Command) *openapi.Sars2ProteinDownloadOpts {
	opts := new(openapi.Sars2ProteinDownloadOpts)
	opts.RefseqOnly = optional.NewBool(argRefseqOnly)
	opts.AnnotatedOnly = optional.NewBool(argAnnotatedOnly)
	argAnnotTypes := getVirusIncludeAnnotationTypesProtein()
	if len(argAnnotTypes) > 0 {
		opts.IncludeAnnotationType = optional.NewInterface(argAnnotTypes)
	}
	if argReleasedSince != "" {
		date, err := fmtdate.Parse(dateFormat, argReleasedSince)
		if err != nil {
			cmd.PrintErr(cmd.UsageString())
			cmd.PrintErr("\nError: ", err, "\n")
			os.Exit(1)
		}
		opts.ReleasedSince = optional.NewTime(date)
	}
	opts.Host = optional.NewString(argHost)
	opts.GeoLocation = optional.NewString(argGeoLocation)
	opts.CompleteOnly = optional.NewBool(argCompleteOnly)
	return opts
}

func downloadVirusProtein(opts *openapi.Sars2ProteinDownloadOpts, proteinNames []string, assmFilename string) (err error) {
	f, err := os.Create(assmFilename)
	if err != nil {
		return
	}
	cli, err := createOAClient()
	if err != nil {
		return
	}
	length := int64(-1) // unknown length
	_, resp, err := cli.VirusApi.Sars2ProteinDownload(nil, proteinNames, opts)
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
Download a SARS-CoV-2 protein dataset by protein name. SARS-CoV-2 protein
datasets include CDS and protein sequence, annotation and a detailed data report. 
Datasets are downloaded as a zip file.

The default SARS-CoV-2 protein dataset includes the following files:
* cds.fna (nucleotide coding sequences)
* protein.faa (protein sequences)
* protein.gpff (protein sequence and annotation in GenPept flat file format)
* protein structures in PDB format
* data_report.jsonl (data report with viral metadata)
* virus_dataset.md (README containing details on sequence file data content and other information)
* dataset_catalog.json (a list of files and file types included in the dataset)
 
Refer to NCBI's [command line quickstart](https://www.ncbi.nlm.nih.gov/datasets/docs/quickstarts/command-line-tools/) documentation for information about getting started with the command-line tools. 

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
	Example: `  datasets download virus protein S --released-since 05/05/2021 --filename SARS2-spike-all.zip
  datasets download virus protein S E M N --refseq --filename SARS2-structural-refseq.zip`,
	Args: cobra.MinimumNArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		opts := sars2ProteinDownloadOpts(cmd)
		return downloadVirusProtein(opts, args, argAssmFilename)
	},
}

func init() {
	registerHiddenBoolPair(downloadVirusProteinCmd.PersistentFlags(), &argAnnotatedOnly, "annotated", "a", false, "limit to annotated coronavirus genomes")
	registerHiddenBoolPair(downloadVirusProteinCmd.PersistentFlags(), &argCompleteOnly, "complete-only", "c", false, "limit to complete coronavirus genomes")
	downloadVirusProteinCmd.PersistentFlags().BoolVar(&argExcludeProtein, "exclude-cds", false, "exclude cds.fna (CDS sequence file)")
	downloadVirusProteinCmd.PersistentFlags().BoolVar(&argExcludePdb, "exclude-pdb", false, "exclude *.pdb (protein structure files)")
	registerHiddenBoolPair(downloadVirusProteinCmd.PersistentFlags(), &argExcludeProtein, "exclude-protein", "p", false, "exclude protein.faa (protein sequence file)")
	registerHiddenStringPair(downloadVirusProteinCmd.PersistentFlags(), &argAssmFilename, "filename", "f", "ncbi_dataset.zip", "specify a custom file name for the downloaded dataset")
	downloadVirusProteinCmd.PersistentFlags().StringVar(&argGeoLocation, "geo-location", "", "limit to coronavirus genomes isolated from a specified geographic location (continent, country or U.S. state)")
	downloadVirusProteinCmd.PersistentFlags().StringVar(&argHost, "host", "", "limit to coronavirus genomes isolated from a specified host (NCBI Taxonomy ID, scientific or common name at any taxonomic rank)")
	registerHiddenBoolPair(downloadVirusProteinCmd.PersistentFlags(), &argExcludeGpff, "exclude-gpff", "g", false, "exclude protein.gpff (protein sequence and annotation in GenPept flat file format")
	downloadVirusProteinCmd.PersistentFlags().BoolVar(&argRefseqOnly, "refseq", false, "limit to RefSeq coronavirus genomes")
	downloadVirusProteinCmd.PersistentFlags().StringVar(&argReleasedSince, "released-since", "", "limit to coronavirus genomes released after a specified date ("+dateFormat+")")
}
