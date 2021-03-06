package datasets

import (
	"errors"
	"fmt"
	openapi "main/openapi_client"

	"github.com/spf13/cobra"
)

var downloadOrthologCmd = &cobra.Command{
	Use:     "ortholog",
	Short:   "download an ortholog dataset",
	Example: `  datasets download ortholog gene-id 672`,
	Long: `
Download an ortholog dataset including gene, transcript and protein sequence, a data table and a data report. Ortholog data is calculated by NCBI for vertebrates and insects. Ortholog datasets can be specified by NCBI Gene ID, symbol or RefSeq accession. Datasets are downloaded as a zip file.

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
		return
	},
}

func downloadOrthologsByGeneId(geneInts []int64) error {
	geneIdSet := map[string]bool{}
	orthoGenes := []string{}
	for _, gene_id := range geneInts {
		orthologs, err := fetchOrthologByGeneId(gene_id, openapi.V1ALPHA1ORTHOLOGREQUESTCONTENTTYPE_IDS_ONLY)
		if err != nil {
			return err
		}
		for _, g := range orthologs.Genes.Genes {
			if err := printMessage(&g); err != nil {
				return err
			}
			if len(g.Gene.GeneId) > 0 && !geneIdSet[g.Gene.GeneId] {
				geneIdSet[g.Gene.GeneId] = true
				orthoGenes = append(orthoGenes, g.Gene.GeneId)
			}
		}
	}

	if len(orthoGenes) == 0 {
		return errors.New("no valid NCBI gene identifiers, exiting")
	}
	orthoGeneInts := strToInt64List(orthoGenes)
	fmt.Printf("Found %d genes in set\n", len(orthoGenes))
	err := downloadGeneByIDs(orthoGeneInts, argExcludeGene, argExcludeRna, argExcludeProtein, argGeneFilename)
	return err
}

func downloadOrthologByGeneReq(req *openapi.V1alpha1GeneDatasetRequest) (err error) {
	var geneInts []int64
	geneInts, err = allGeneIdForRequest(req)
	if err != nil {
		return
	}
	err = downloadOrthologsByGeneId(geneInts)
	return
}

func init() {
	downloadOrthologCmd.AddCommand(downloadOrthologGeneIDCmd)
	downloadOrthologCmd.AddCommand(downloadOrthologSymbolCmd)
	downloadOrthologCmd.AddCommand(downloadOrthologAccessionCmd)

	flags := downloadOrthologCmd.PersistentFlags()
	registerHiddenStringPair(flags, &argGeneFilename, "filename", "f", "ncbi_dataset.zip", "specify a custom file name for the downloaded dataset")

	registerHiddenBoolPair(flags, &argExcludeGene, "exclude-gene", "g", false, "exclude gene.fna (gene sequence file)")
	registerHiddenBoolPair(flags, &argExcludeRna, "exclude-rna", "r", false, "exclude rna.fna (transcript sequence file)")
	registerHiddenBoolPair(flags, &argExcludeProtein, "exclude-protein", "p", false, "exclude protein.faa (protein sequence file)")
	flags.StringSliceVar(&taxonFilter, "taxon-filter", []string{}, "limit results to ortholog data for a specified taxonomic group")
}
