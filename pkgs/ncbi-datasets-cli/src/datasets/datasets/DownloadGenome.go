package datasets

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	datasets_util "main/datasets/util"
	openapi "main/openapi_client"
)

func getDownloadRequest(acc []string) *openapi.V1alpha1AssemblyDatasetRequest {
	opts := new(openapi.V1alpha1AssemblyDatasetRequest)
	opts.ExcludeSequence = argExcludeSeq
	if len(argChromosomes) > 0 {
		opts.Chromosomes = argChromosomes
	}
	var argAnnotTypes []openapi.V1alpha1AnnotationForAssemblyType
	if !argExcludeGff3 {
		argAnnotTypes = append(argAnnotTypes, openapi.V1ALPHA1ANNOTATIONFORASSEMBLYTYPE_GENOME_GFF)
	}
	if argIncludeGbff {
		argAnnotTypes = append(argAnnotTypes, openapi.V1ALPHA1ANNOTATIONFORASSEMBLYTYPE_GENOME_GBFF)
	}
	if argIncludeGtf {
		argAnnotTypes = append(argAnnotTypes, openapi.V1ALPHA1ANNOTATIONFORASSEMBLYTYPE_GENOME_GTF)
	}
	if !argExcludeProtein {
		argAnnotTypes = append(argAnnotTypes, openapi.V1ALPHA1ANNOTATIONFORASSEMBLYTYPE_PROT_FASTA)
	}
	if !argExcludeRna {
		argAnnotTypes = append(argAnnotTypes, openapi.V1ALPHA1ANNOTATIONFORASSEMBLYTYPE_RNA_FASTA)
	}
	if len(argAnnotTypes) > 0 {
		opts.IncludeAnnotationType = argAnnotTypes
	}
	if argDehydrated {
		opts.Hydrated = openapi.ASSEMBLYDATASETREQUESTRESOLUTION_DATA_REPORT_ONLY
	}
	opts.AssemblyAccessions = acc

	return opts
}

func downloadAssembly(opts *openapi.V1alpha1AssemblyDatasetRequest, assmFilename string) (err error) {
	f, e := os.Create(assmFilename)
	if e != nil {
		err = fmt.Errorf("'%s' opening output file: %s", e, assmFilename)
		return
	}
	cli, err := createOAClient()
	if err != nil {
		return
	}
	_, resp, err := cli.GenomeApi.DownloadAssemblyPackagePost(nil, *opts, nil)
	if err = handleHTTPResponse(resp, err); err != nil {
		return
	}
	length := int64(-1) // unknown length
	err = downloadData(f, resp, err, assmFilename, length)
	return
}

func isValidAssemblyAcc(acc string) bool {
	isAccVerSuffix := func(accver_suffix string) bool {
		_, err := strconv.ParseFloat(accver_suffix, 32)
		return (err == nil)
	}

	validPrefixes := []string{
		"GCA_",
		"GCF_",
		"PRJNA",
		"PRJDA",
		"PRJDB",
		"PRJEA",
		"PRJEB",
	}
	for _, prefix := range validPrefixes {
		if strings.HasPrefix(acc, prefix) && isAccVerSuffix(acc[len(prefix):]) {
			return true
		}
	}
	return false
}

func lookupAssmsForBioProjects(cli *openapi.APIClient, bioprojectAccs []string) (validAccs []string, err error) {
	if len(bioprojectAccs) == 0 {
		return
	}
	request := new(openapi.V1alpha1AssemblyMetadataRequest)
	request.Bioprojects.Accessions = bioprojectAccs

	err = updateAssemblyMetadataRequestOption(request)
	if err != nil {
		return
	}

	result, meta_err := getAssemblyMetadataWithPost(request, true)
	if meta_err != nil {
		return validAccs, meta_err
	}

	validAccs = make([]string, 0, len(result.Assemblies))
	for _, assm := range result.Assemblies {
		if len(assm.Assembly.AssemblyAccession) > 0 {
			validAccs = append(validAccs, assm.Assembly.AssemblyAccession)
		}
	}
	err = datasets_util.MessagesToError(result.Messages)
	return
}

func separateAccessions(accs []string) (retAccs []string, bioprojectAccs []string, err error) {
	accNum := 0
	for _, acc := range accs {
		if isValidAssemblyAcc(acc) {
			if strings.HasPrefix(acc, "PRJ") {
				bioprojectAccs = append(bioprojectAccs, acc)
			} else {
				accs[accNum] = acc
				accNum++
			}
		} else {
			err = fmt.Errorf("invalid or unsupported assembly accession: %s", acc)
		}
	}
	retAccs = accs[:accNum]
	return
}

func gencollAccessionsForMixedAccessions(cli *openapi.APIClient, accs []string) (assmAccs []string, err error) {
	accs, bioprojectAccs, err := separateAccessions(accs)
	if err != nil {
		return
	}
	assmAccs, err = lookupAssmsForBioProjects(cli, bioprojectAccs)
	if err != nil {
		return
	}
	assmAccs = append(accs, assmAccs...)
	return
}

// TODO : BABY-1392 - Upgrade CLI to use AssemblyMetadata over checkAssemblyAvailability
func checkAssemblyAvailability(accs []string) (assmAccs []string, warning string, err error) {
	warning = ""
	accs, bioprojectAccs, err := separateAccessions(accs)
	if err != nil {
		return
	}
	cli, err := createOAClient()
	if err != nil {
		return
	}
	argIDArgs, err := lookupAssmsForBioProjects(cli, bioprojectAccs)
	if err != nil {
		assmAccs = nil
		return
	} else if len(accs) == 0 && len(argIDArgs) == 0 {
		assmAccs = nil
		err = errors.New("no assemblies found for requested accessions")
		return
	}

	opts := new(openapi.V1alpha1AssemblyDatasetRequest)
	opts.AssemblyAccessions = accs

	result, resp, err := cli.GenomeApi.CheckAssemblyAvailabilityPost(nil, *opts)
	err = handleHTTPResponseWithCustomErr(resp, err, "Unexpected %s error in checking assembly availability")
	if err != nil {
		assmAccs = nil
		return
	}
	assmAccs = append(argIDArgs, result.ValidAssemblies...)
	warning = result.Reason
	return
}

func processAssmAccArgs(args []string) (err error) {
	argIDArgs, err = getArgsFromListOrFile(args, argInputFile)
	if err == nil && len(argIDArgs) == 0 {
		err = errors.New("Input accessions not specified")
	}
	return
}

var genomeCmd = &cobra.Command{
	Use:   "genome",
	Short: "download a genome dataset",
	Example: `  datasets download genome accession GCF_000001405.39 --chromosomes X,Y --exclude-gff3 --exclude-rna
  datasets download genome taxon "bos taurus" --dehydrated
  datasets download genome taxon human --assembly-level chromosome,complete_genome --dehydrated
  datasets download genome taxon mouse --search C57BL/6J --search "Broad Institute" --dehydrated`,
	Long: `
Download a genome dataset including genome, transcript and protein sequence, annotation and a detailed data report. 
Genome datasets can be specified by NCBI Assembly or BioProject accession or taxon. Datasets are downloaded as a zip file.

The default genome dataset includes the following files (if available):
* genomic.fna (genomic sequences)
* rna.fna (transcript sequences)
* protein.faa (protein sequences)
* genomic.gff (genome annotation in gff3 format)
* data_report.jsonl (data report with genome assembly and annotation metadata)
* dataset_catalog.json (a list of files and file types included in the dataset)

Refer to NCBI's [command line quickstart](https://www.ncbi.nlm.nih.gov/datasets/docs/quickstarts/command-line-tools/) documentation for information about getting started with the command-line tools.`,

	PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
		err = processAssmAccArgs(args)
		return
	},
}

func init() {
	genomeCmd.AddCommand(downloadGenomeAccessionCmd)
	genomeCmd.AddCommand(downloadGenomeTaxonCmd)

	pflags := genomeCmd.PersistentFlags()
	registerHiddenStringPair(pflags, &argAssmFilename, "filename", "f", "ncbi_dataset.zip", "specify a custom file name for the downloaded dataset")
	registerHiddenStringSlicePair(pflags, &argChromosomes, "chromosomes", "c", []string{"all"}, "limit to a specified, comma-delimited list of chromosomes")
	registerHiddenBoolPair(pflags, &argExcludeSeq, "exclude-seq", "s", false, "exclude genomic.fna (genomic sequence file)")
	registerHiddenBoolPair(pflags, &argExcludeRna, "exclude-rna", "r", false, "exclude rna.fna (transcript sequence file)")
	registerHiddenBoolPair(pflags, &argExcludeProtein, "exclude-protein", "p", false, "exclude protein.faa (protein sequence file)")
	registerHiddenBoolPair(pflags, &argExcludeGff3, "exclude-gff3", "g", false, "exclude genomic.gff (gff3 annotation file)")
	registerHiddenBoolPair(pflags, &argIncludeGbff, "include-gbff", "b", false, "include genomic.gbff (GenBank flat file sequence and annotation), if available")
	registerHiddenBoolPair(pflags, &argIncludeGtf, "include-gtf", "e", false, "include genomic.gtf (gtf annotation file), if available")
	pflags.BoolVar(&argDehydrated, "dehydrated", false, "download a dehydrated zip archive including the data report and locations of data files (use the rehydrate command to retrieve data files).")

	addGenomeMetadataFilters(pflags)
}
