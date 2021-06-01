package datasets

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	openapi "main/openapi_client"
)

var (
	argGeneFilename        string
	argFastaFilterFilename string
	argFastaFilter         []string
)

func singleGeneIdForRequest(req *openapi.V1alpha1GeneDatasetRequest) (int64, error) {
	req.ReturnedContent = openapi.V1ALPHA1GENEDATASETREQUESTCONTENTTYPE_IDS_ONLY
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

func allGeneIdForRequest(req *openapi.V1alpha1GeneDatasetRequest) ([]int64, error) {
	req.ReturnedContent = openapi.V1ALPHA1GENEDATASETREQUESTCONTENTTYPE_IDS_ONLY
	gsp := &GeneIdStreamProcessor{}
	err := streamGeneMatch(req, gsp)
	if err != nil {
		err = fmt.Errorf("Failure to read response from server: %w", err)
		return []int64{}, err
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
Download a gene dataset including gene, transcript and protein sequence, a data table and a data report. Gene datasets can be specified by NCBI Gene ID, symbol or RefSeq accession. Datasets are downloaded as a zip file.

The default gene dataset includes the following files:
 * gene.fna (gene sequences)
 * rna.fna (transcript sequences)
 * protein.faa (protein sequences)
 * data_report.jsonl (data report with gene metadata)
 * data_table.tsv (data table with gene metadata, one transcript per row)
 * dataset_catalog.json (a list of files and file types included in the dataset)

Refer to NCBI's [command line quickstart](https://www.ncbi.nlm.nih.gov/datasets/docs/quickstarts/command-line-tools/) documentation for information about getting started with the command-line tools.
`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
		argIDArgs, err = getArgsFromListOrFile(args, argInputFile)
		if err == nil && len(argIDArgs) == 0 {
			err = errors.New("input genes not specified")
		}
		return
	},

	RunE: cmdDownloadGeneGeneID,
}

func cmdDownloadGeneGeneID(cmd *cobra.Command, args []string) error {
	req := new(openapi.V1alpha1GeneDatasetRequest)
	req.GeneIds = strToInt64List(argIDArgs)
	return downloadGeneForRequest(req)
}

func downloadGeneForRequest(req *openapi.V1alpha1GeneDatasetRequest) (err error) {
	var geneIds []int64
	geneIds, err = allGeneIdForRequest(req)
	if err != nil {
		return
	}
	if len(geneIds) == 0 {
		err = errors.New("No valid NCBI gene identifiers - Exiting")
		return
	}
	err = downloadGeneByIDs(geneIds, argExcludeGene, argExcludeRna, argExcludeProtein, argGeneFilename)
	return
}

func downloadGeneByIDs(geneIDs []int64, excludeGene bool, excludeRNA bool, excludeProtein bool, geneFilename string) (err error) {
	f, e := os.Create(geneFilename)
	if e != nil {
		err = fmt.Errorf("'%s' opening output file: %s", e, geneFilename)
		return
	}
	defer f.Close()
	req := new(openapi.V1alpha1GeneDatasetRequest)
	req.GeneIds = geneIDs

	accessionArgs := []string{}
	if argFastaFilterFilename != "" {
		fp, fileErr := os.Open(argFastaFilterFilename)
		if fileErr != nil {
			err = fmt.Errorf("'%s' opening input file: '%s'", fileErr.Error(), argFastaFilterFilename)
			return
		}
		defer fp.Close()
		accessionArgs = readLines(fp)
		// Check if any accessions were read
		if len(accessionArgs) == 0 {
			err = fmt.Errorf(
				"No identifiers read from file: '%s'\n       File should have 1 identifier per row and no spaces or quotes",
				argFastaFilterFilename,
			)
			return
		}
	}

	if len(argFastaFilter) > 0 {
		accessionArgs = append(accessionArgs, argFastaFilter...)
	}
	if len(accessionArgs) > 0 {
		req.FastaFilter = accessionArgs
	}

	cli, err := createOAClient()
	if err != nil {
		return
	}

	if !excludeGene {
		req.IncludeAnnotationType = append(req.IncludeAnnotationType, openapi.V1ALPHA1FASTA_GENE)
	}
	if !excludeRNA {
		req.IncludeAnnotationType = append(req.IncludeAnnotationType, openapi.V1ALPHA1FASTA_RNA)
	}
	if !excludeProtein {
		req.IncludeAnnotationType = append(req.IncludeAnnotationType, openapi.V1ALPHA1FASTA_PROTEIN)
	}
	_, resp, err := cli.GeneApi.DownloadGenePackagePost(nil, *req, nil)
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
	registerHiddenStringPair(flags, &argGeneFilename, "filename", "f", "ncbi_dataset.zip", "specify a custom file name for the downloaded dataset")

	registerHiddenBoolPair(flags, &argExcludeGene, "exclude-gene", "g", false, "exclude gene.fna (gene sequence file)")
	registerHiddenBoolPair(flags, &argExcludeRna, "exclude-rna", "r", false, "exclude rna.fna (transcript sequence file)")
	registerHiddenBoolPair(flags, &argExcludeProtein, "exclude-protein", "p", false, "exclude protein.faa (protein sequence file)")
}
