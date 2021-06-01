package datasets

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"

	openapi "main/openapi_client"
)

var (
	argFlank       int
	argTaxonFilter string
)

var downloadGeneAccessionCmd = &cobra.Command{
	Use:   "accession <refseq-accession ...>",
	Short: "download a gene dataset by RefSeq nucleotide or protein accession",
	Example: `  datasets download gene accession NP_000483.3
  datasets download gene accession NM_000546.6 NM_000492.4
  datasets download gene accession WP_004675351.1`,
	Long: `
Download a gene dataset by RefSeq nucleotide or protein accession. Gene datasets include gene, transcript and protein sequence, a data table and a data report. Datasets are downloaded as a zip file.

The default gene dataset includes the following files:
 * gene.fna (gene sequences)
 * rna.fna (transcript sequences)
 * protein.faa (protein sequences)
 * data_report.jsonl (data report with gene metadata)
 * data_table.tsv (data table with gene metadata, one transcript per row)
 * dataset_catalog.json (a list of files and file types included in the dataset)
 * annotation_report.jsonl (included with prokaryotic gene packages)

Refer to [NCBI Datasets Gene Package](https://www.ncbi.nlm.nih.gov/datasets/docs/data-packages/gene/) documentation for more information about the gene package. `,

	RunE: cmdDownloadGeneAccession,
}

var dlGeneAccessionCmd = &cobra.Command{
	Deprecated: "please use \"datasets download gene accession\"",
	Use:        "acc <gene_accession ...>",
	Short:      "Download data by gene accession",
	Example:    "  datasets download gene acc NM_001347423.2",
	Long: `
Download data by gene accession. Data is returned as a zip archive.
The default package includes a gene data report and genomic, RNA and protein fasta sequence.

Refer to NCBI's [command line quickstart](https://www.ncbi.nlm.nih.gov/datasets/docs/quickstarts/command-line-tools/) documentation for information about getting started with the command-line tools.
`,
	RunE: cmdDownloadGeneAccession,
}

func isProkaryoteAcc(acc_list []string) bool {
	const prokPrefix string = "WP_"
	for _, accession := range acc_list {
		if strings.HasPrefix(accession, prokPrefix) == false {
			return false
		}
	}
	return true
}

func downloadProkaryoteGeneForRequest(req *openapi.V1alpha1ProkaryoteGeneRequest) (err error) {

	f, e := os.Create(argGeneFilename)
	if e != nil {
		err = fmt.Errorf("'%s' opening output file: %s", e, argGeneFilename)
		return
	}
	defer f.Close()

	cli, err := createOAClient()
	if err != nil {
		return
	}

	if !argExcludeGene {
		req.IncludeAnnotationType = append(req.IncludeAnnotationType, openapi.V1ALPHA1FASTA_GENE)
	}
	if !argExcludeProtein {
		req.IncludeAnnotationType = append(req.IncludeAnnotationType, openapi.V1ALPHA1FASTA_PROTEIN)
	}

	flank_length := int64(argFlank)
	if flank_length > 0 {
		req.GeneFlankConfig.Length = flank_length
		req.IncludeAnnotationType = append(req.IncludeAnnotationType, openapi.V1ALPHA1FASTA_GENE_FLANK)
	}

	if argTaxonFilter != "" {
		req.Taxon = argTaxonFilter
	}

	_, resp, err := cli.ProkaryoteApi.DownloadProkaryoteGenePackagePost(nil, *req, nil)
	if err = handleHTTPResponse(resp, err); err != nil {
		return
	}
	length := int64(-1) // unknown length
	err = downloadData(f, resp, err, argGeneFilename, length)
	return
}

func cmdDownloadGeneAccession(cmd *cobra.Command, args []string) error {
	if isProkaryoteAcc(argIDArgs) {
		req := new(openapi.V1alpha1ProkaryoteGeneRequest)
		req.Accessions = argIDArgs
		return downloadProkaryoteGeneForRequest(req)
	} else {
		req := new(openapi.V1alpha1GeneDatasetRequest)
		req.Accessions = argIDArgs
		return downloadGeneForRequest(req)
	}
}

func init() {
	registerHiddenStringPair(
		downloadGeneAccessionCmd.Flags(),
		&argInputFile,
		"inputfile",
		"i",
		"",
		"read a list of RefSeq nucleotide or protein accessions from a file to use as input",
	)

	registerHiddenStringPair(
		dlGeneAccessionCmd.Flags(),
		&argInputFile,
		"inputfile",
		"i",
		"",
		"read a list of RefSeq nucleotide or protein accessions from a file to use as input",
	)

	flags := downloadGeneAccessionCmd.Flags()
	registerHiddenIntPair(flags, &argFlank, "include-flanks-bp", "k", 0, "include gene flanking sequence, limited to prokaryotic genes (default: gene_flanks.fna is omitted from package)")
	registerHiddenStringPair(flags, &argTaxonFilter, "taxon-filter", "", "", "limit genes to a specified taxon (any rank)")
}
