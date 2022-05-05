package datasets

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"

	openapi "datasets_cli/v1/openapi"
)

func getDownloadRequest(acc []string) *openapi.V1AssemblyDatasetRequest {
	req := openapi.NewV1AssemblyDatasetRequest()
	req.SetExcludeSequence(argExcludeSeq)
	if len(argChromosomes) > 0 {
		req.SetChromosomes(argChromosomes)
	}
	annotations := make([]openapi.V1AnnotationForAssemblyType, 0)
	possible_annotations := []struct {
		flag      bool
		type_enum openapi.V1AnnotationForAssemblyType
	}{
		{!argExcludeGff3, openapi.V1ANNOTATIONFORASSEMBLYTYPE_GENOME_GFF},
		{argIncludeGbff, openapi.V1ANNOTATIONFORASSEMBLYTYPE_GENOME_GBFF},
		{argIncludeGtf, openapi.V1ANNOTATIONFORASSEMBLYTYPE_GENOME_GTF},
		{!argExcludeProtein, openapi.V1ANNOTATIONFORASSEMBLYTYPE_PROT_FASTA},
		{!argExcludeRna, openapi.V1ANNOTATIONFORASSEMBLYTYPE_RNA_FASTA},
		{!argExcludeCds, openapi.V1ANNOTATIONFORASSEMBLYTYPE_CDS_FASTA},
	}
	for _, annot := range possible_annotations {
		if annot.flag {
			annotations = append(annotations, annot.type_enum)
		}
	}
	req.SetIncludeAnnotationType(annotations)
	if argDehydrated {
		req.SetHydrated(openapi.V1ASSEMBLYDATASETREQUESTRESOLUTION_DATA_REPORT_ONLY)
	}
	req.SetAssemblyAccessions(acc)

	return req
}

func downloadAssembly(req *openapi.V1AssemblyDatasetRequest, assmFilename string) (err error) {
	f, e := os.Create(assmFilename)
	if e != nil {
		err = fmt.Errorf("'%s' opening output file: %s", e, assmFilename)
		return
	}
	cli, err := createOAClient()
	if err != nil {
		return
	}
	_, resp, err := cli.GenomeApi.DownloadAssemblyPackagePost(nil, req).Execute()
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
	request := &openapi.V1AssemblyMetadataRequest{
		Bioprojects: &openapi.V1AssemblyMetadataRequestBioprojects{
			Accessions: &bioprojectAccs,
		},
	}

	err = updateAssemblyMetadataRequestOption(request)
	if err != nil {
		return
	}

	// Command line argument is used to set limit on number of records retrieved.  We want all of them at this point.
	var tmp string
	if argLimit != "" {
		tmp = argLimit
		argLimit = ""
	}
	validAccs, err = getAssemblyAccessionsWithPost(request)
	if tmp != "" {
		argLimit = tmp
	}

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

	request := &openapi.V1AssemblyDatasetRequest{
		Accessions: &accs,
	}

	result, resp, err := cli.GenomeApi.CheckAssemblyAvailabilityPost(nil, request).Execute()
	err = handleHTTPResponseWithCustomErr(resp, err, "Unexpected %s error in checking assembly availability")
	if err != nil {
		assmAccs = nil
		return
	}
	assmAccs = append(argIDArgs, result.GetValidAssemblies()...)
	warning = result.GetReason()
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
	Example: `  datasets download genome accession GCF_000001405.40 --chromosomes X,Y --exclude-gff3 --exclude-rna
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

Refer to NCBI's [download and install](https://www.ncbi.nlm.nih.gov/datasets/docs/download-and-install/) documentation for information about getting started with the command-line tools.`,
}

func init() {
	genomeCmd.AddCommand(downloadGenomeAccessionCmd)
	genomeCmd.AddCommand(downloadGenomeTaxonCmd)

	pflags := genomeCmd.PersistentFlags()
	registerHiddenStringSlicePair(pflags, &argChromosomes, "chromosomes", "c", []string{"all"}, "limit to a specified, comma-delimited list of chromosomes")
	registerHiddenBoolPair(pflags, &argExcludeSeq, "exclude-seq", "s", false, "exclude genomic.fna (genomic sequence file)")
	registerHiddenBoolPair(pflags, &argExcludeRna, "exclude-rna", "r", false, "exclude rna.fna (transcript sequence file)")
	registerHiddenBoolPair(pflags, &argExcludeProtein, "exclude-protein", "p", false, "exclude protein.faa (protein sequence file)")
	registerHiddenBoolPair(pflags, &argExcludeGff3, "exclude-gff3", "g", false, "exclude genomic.gff (gff3 annotation file)")
	registerHiddenBoolPair(pflags, &argExcludeCds, "exclude-genomic-cds", "d", false, "exclude cds_from_genomic.fna (genomic cds file)")
	registerHiddenBoolPair(pflags, &argIncludeGbff, "include-gbff", "b", false, "include genomic.gbff (GenBank flat file sequence and annotation), if available")
	registerHiddenBoolPair(pflags, &argIncludeGtf, "include-gtf", "e", false, "include genomic.gtf (gtf annotation file), if available")
	pflags.BoolVar(&argDehydrated, "dehydrated", false, "download a dehydrated zip archive including the data report and locations of data files (use the rehydrate command to retrieve data files).")

	addGenomeMetadataFilters(pflags)
}
