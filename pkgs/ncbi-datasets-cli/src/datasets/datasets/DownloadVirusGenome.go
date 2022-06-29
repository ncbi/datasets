package datasets

import "github.com/spf13/cobra"

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
* data_report.jsonl (data report with viral metadata)
* virus_dataset.md (README containing details on sequence file data content and other information)
* dataset_catalog.json (a list of files and file types included in the dataset)

Refer to NCBI's [download and install](https://www.ncbi.nlm.nih.gov/datasets/docs/download-and-install/) documentation for information about getting started with the command-line tools.`,
	Example: `  datasets download virus genome taxon sars-cov-2 --host dog
  datasets download virus genome taxon coronaviridae --host "manis javanica"`,
}

func init() {
	registerHiddenBoolPair(downloadVirusGenomeCmd.PersistentFlags(), &argAnnotatedOnly, "annotated", "a", false, "limit to annotated coronavirus genomes")
	registerHiddenBoolPair(downloadVirusGenomeCmd.PersistentFlags(), &argCompleteOnly, "complete-only", "c", false, "limit to complete coronavirus genomes")
	downloadVirusGenomeCmd.PersistentFlags().BoolVar(&argExcludeCds, "exclude-cds", false, "exclude cds.fna (CDS sequence file)")
	registerHiddenBoolPair(downloadVirusGenomeCmd.PersistentFlags(), &argExcludeProtein, "exclude-protein", "p", false, "exclude protein.faa (protein sequence file)")
	registerHiddenBoolPair(downloadVirusGenomeCmd.PersistentFlags(), &argExcludeSeq, "exclude-seq", "s", false, "exclude genomic.fna (genomic sequence file)")
	downloadVirusGenomeCmd.PersistentFlags().StringVar(&argGeoLocation, "geo-location", "", "limit to coronavirus genomes isolated from a specified geographic location (continent, country or U.S. state)")
	downloadVirusGenomeCmd.PersistentFlags().StringVar(&argHost, "host", "", "limit to coronavirus genomes isolated from a specified host (NCBI Taxonomy ID, scientific or common name at any taxonomic rank)")
	downloadVirusGenomeCmd.PersistentFlags().StringVar(&argLineage, "lineage", "", "limit to SARS-CoV-2 genomes classified as the specified lineage (variant) by pangolin using the pangoLEARN algorithm")
	downloadVirusGenomeCmd.PersistentFlags().BoolVar(&argRefseqOnly, "refseq", false, "limit to RefSeq coronavirus genomes")
	downloadVirusGenomeCmd.PersistentFlags().StringVar(&argReleasedSince, "released-since", "", "limit to coronavirus genomes released after a specified date ("+dateFormat+")")
	downloadVirusGenomeCmd.PersistentFlags().StringVar(&argUpdatedSince, "updated-since", "", "limit to coronavirus genomes updated after a specified date ("+dateFormat+")")
}
