package datasets

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	openapi "datasets_cli/v1/openapi"
)

var (
	argFastaFilterFilename string
	argFastaFilter         []string
)

func singleGeneIdForRequest(req *openapi.V1GeneDatasetRequest) (int32, error) {
	req.SetReturnedContent(openapi.V1GENEDATASETREQUESTCONTENTTYPE_IDS_ONLY)
	gsp := &GeneIdStreamProcessor{}
	err := streamGeneMatch(req, gsp)
	if err != nil {
		err = fmt.Errorf("Failure to read response from server: %w", err)
		return 0, err
	}

	if len(gsp.GeneIds) == 0 {
		err = errors.New("No genes found for search term")
		return 0, err
	}

	if len(gsp.GeneIds) > 1 {
		err = errors.New("multiple genes found, symbol and taxon ambiguous")
		return 0, err
	}
	return gsp.GeneIds[0], nil
}

func allGeneIdForRequest(req *openapi.V1GeneDatasetRequest) ([]int32, error) {
	req.SetReturnedContent(openapi.V1GENEDATASETREQUESTCONTENTTYPE_IDS_ONLY)
	gsp := &GeneIdStreamProcessor{}
	err := streamGeneMatch(req, gsp)
	if err != nil {
		err = fmt.Errorf("Failure to read response from server: %w", err)
		return []int32{}, err
	}
	return gsp.GeneIds, nil
}

// downloadGeneCmd represents the gene command
var downloadGeneCmd = &cobra.Command{
	Use:   "gene",
	Short: "download a gene dataset",
	Example: `  datasets download gene gene-id 672
  datasets download gene symbol brca1 --taxon mouse
  datasets download gene accession NP_000483.3
  datasets download gene gene-id 2778 --fasta-filter NC_000020.11,NM_001077490.3,NP_001070958.1`,
	Long: `
Download a gene data package including gene, transcript and protein sequence, a data table and a data report. Gene data packages can be specified by NCBI Gene ID, symbol or RefSeq accession. Data packages are downloaded as a zip file.

The default gene dataset includes the following files:
 * gene.fna (gene sequences)
 * rna.fna (transcript sequences)
 * protein.faa (protein sequences)
 * data_report.jsonl (data report with gene metadata)
 * data_table.tsv (data table with gene metadata, one transcript per row)
 * dataset_catalog.json (a list of files and file types included in the dataset)

Refer to NCBI's [command line quickstart](https://www.ncbi.nlm.nih.gov/datasets/docs/quickstarts/command-line-tools/) documentation for information about getting started with the command-line tools.
`,
	RunE: cmdDownloadGeneGeneID,
}

func cmdDownloadGeneGeneID(cmd *cobra.Command, args []string) error {
	req := openapi.NewV1GeneDatasetRequest()
	req.SetGeneIds(strToInt32List(argIDArgs))
	return downloadGeneForRequest(req)
}

func downloadGeneForRequest(req *openapi.V1GeneDatasetRequest) (err error) {
	var geneIds []int32
	geneIds, err = allGeneIdForRequest(req)
	if err != nil {
		return
	}
	if len(geneIds) == 0 {
		err = errors.New("No valid NCBI gene identifiers - Exiting")
		return
	}
	err = downloadGeneByIDs(geneIds, argExcludeGene, argExcludeRna, argExcludeProtein, argIncludeCds, argInclude5putr, argInclude3putr, argDownloadFilename)
	return
}

func downloadGeneByIDs(geneIDs []int32, excludeGene bool, excludeRNA bool, excludeProtein bool, includeCds bool, include5utr bool, include3utr bool, geneFilename string) (err error) {
	f, e := os.Create(geneFilename)
	if e != nil {
		err = fmt.Errorf("'%s' opening output file: %s", e, geneFilename)
		return
	}
	defer f.Close()

	cli, err := createOAClient()
	if err != nil {
		return
	}
	req := openapi.NewV1GeneDatasetRequest()
	req.SetGeneIds(geneIDs)

	fastaFilterArgs := []string{}
	if argFastaFilterFilename != "" {
		fp, fileErr := os.Open(argFastaFilterFilename)
		if fileErr != nil {
			err = fmt.Errorf("'%s' opening input file: '%s'", fileErr.Error(), argFastaFilterFilename)
			return
		}
		defer fp.Close()
		fastaFilterArgs = readLines(fp)
		// Check if any accessions were read
		if len(fastaFilterArgs) == 0 {
			err = fmt.Errorf(
				"No identifiers read from file: '%s'\n       File should have 1 identifier per row and no spaces or quotes",
				argFastaFilterFilename,
			)
			return
		}
	}

	if len(argFastaFilter) > 0 {
		fastaFilterArgs = append(fastaFilterArgs, argFastaFilter...)
	}
	if len(fastaFilterArgs) > 0 {
		req.SetFastaFilter(fastaFilterArgs)
	}

	annotations := make([]openapi.V1Fasta, 0)
	possible_annotations := []struct {
		flag      bool
		type_enum openapi.V1Fasta
	}{
		{!excludeGene, openapi.V1FASTA_GENE},
		{!excludeRNA, openapi.V1FASTA_RNA},
		{!excludeProtein, openapi.V1FASTA_PROTEIN},
		{includeCds, openapi.V1FASTA_CDS},
		{include5utr, openapi.V1FASTA__5_P_UTR},
		{include3utr, openapi.V1FASTA__3_P_UTR},
	}
	for _, annot := range possible_annotations {
		if annot.flag {
			annotations = append(annotations, annot.type_enum)
		}
	}
	req.SetIncludeAnnotationType(annotations)

	_, resp, err := cli.GeneApi.DownloadGenePackagePost(nil, req).Execute()
	if err = handleHTTPResponse(resp, err); err != nil {
		return
	}
	length := int64(-1) // unknown length
	err = downloadData(f, resp, err, geneFilename, length)
	return
}

func init() {
	downloadGeneCmd.AddCommand(downloadGeneGeneIDCmd)
	downloadGeneCmd.AddCommand(downloadGeneSymbolCmd)
	downloadGeneCmd.AddCommand(downloadGeneAccessionCmd)
	downloadGeneCmd.AddCommand(dlGeneAccessionCmd)
	downloadGeneCmd.AddCommand(downloadGeneTaxonCmd)

	flags := downloadGeneCmd.PersistentFlags()
	registerHiddenStringSlicePair(flags, &argFastaFilter, "fasta-filter", "a", []string{}, "limit gene fasta download to a specific list of accessions")
	registerHiddenStringPair(flags, &argFastaFilterFilename, "fasta-filter-file", "", "", "file of accessions to limit gene fasta download")

	registerHiddenBoolPair(flags, &argExcludeGene, "exclude-gene", "g", false, "exclude gene.fna (gene sequence file)")
	registerHiddenBoolPair(flags, &argExcludeRna, "exclude-rna", "r", false, "exclude rna.fna (transcript sequence file)")
	registerHiddenBoolPair(flags, &argExcludeProtein, "exclude-protein", "p", false, "exclude protein.faa (protein sequence file)")

	registerHiddenBoolPair(flags, &argIncludeCds, "include-cds", "c", false, "include cds.fna (CDS sequence file)")
	registerHiddenBoolPair(flags, &argInclude5putr, "include-5p-utr", "5", false, "include 5p_utr.fna (5'-UTR sequence file)")
	registerHiddenBoolPair(flags, &argInclude3putr, "include-3p-utr", "3", false, "include 3p_utr.fna (3'-UTR sequence file)")
}
