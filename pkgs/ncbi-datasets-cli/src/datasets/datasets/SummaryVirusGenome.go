package datasets

import (
	openapi "datasets_cli/v1/openapi"
	"github.com/metakeule/fmtdate"
	"github.com/spf13/cobra"
)

var summaryVirusGenomeCmd = &cobra.Command{
	Use:   "genome",
	Short: "Print coronavirus genome metadata by accession or taxon",
	Long: `
Print coronavirus genome metadata (data report) by nucleotide accession or taxon. The data report is returned in JSON Lines format.

Refer to NCBI's [download and install](https://www.ncbi.nlm.nih.gov/datasets/docs/download-and-install/) documentation for information about getting started with the command-line tools.`,
	Example: `  datasets summary virus genome accession NC_045512.2
  datasets summary virus genome taxon SARS-COV-2 --host dog`,
}

func setBaseFilter() (openapi.V1VirusDatasetFilter, error) {
	filter := *openapi.NewV1VirusDatasetFilter()
	filter.SetRefseqOnly(argRefseqOnly)
	filter.SetAnnotatedOnly(argAnnotatedOnly)
	if argReleasedSince != "" {
		date, e := fmtdate.Parse(dateFormat, argReleasedSince)
		if e != nil {
			return filter, e
		}
		filter.SetReleasedSince(date)
	}
	filter.SetHost(argHost)
	filter.SetPangolinClassification(argLineage)
	filter.SetGeoLocation(argGeoLocation)
	filter.SetCompleteOnly(argCompleteOnly)
	return filter, nil
}

func init() {

	summaryVirusGenomeCmd.AddCommand(summaryVirusGenomeTaxCmd)
	summaryVirusGenomeCmd.AddCommand(summaryVirusGenomeAccCmd)

	registerHiddenStringPair(summaryVirusGenomeCmd.PersistentFlags(), &argLimit, "limit", "l", "all", `limit the number of genome summaries returned
	* "all": returns all matching virus data reports
	* "none": returns only the number of matching virus data reports
	* a number: returns the specified number of matching data reports
	`)
	summaryVirusGenomeCmd.PersistentFlags().BoolVarP(&argAnnotatedOnly, "annotated", "a", false, "Only include viral genomes with annotation")
	summaryVirusGenomeCmd.PersistentFlags().BoolVar(&argCompleteOnly, "complete-only", false, "only include complete coronavirus data reports")
	summaryVirusGenomeCmd.PersistentFlags().StringVar(&argGeoLocation, "geo-location", "", "limit to coronavirus genomes isolated from a specified geographic location (continent, country or U.S. state)")
	summaryVirusGenomeCmd.PersistentFlags().StringVar(&argHost, "host", "", "limit to coronavirus genomes isolated from a specified host (NCBI Taxonomy ID, scientific or common name at any taxonomic rank)")
	summaryVirusGenomeCmd.PersistentFlags().StringVar(&argLineage, "lineage", "", "limit to SARS-CoV-2 genomes classified as the specified lineage (variant) by pangolin using the pangoLEARN algorithm")
	summaryVirusGenomeCmd.PersistentFlags().BoolVar(&argRefseqOnly, "refseq", false, "limit to RefSeq coronavirus genomes")
	summaryVirusGenomeCmd.PersistentFlags().StringVar(&argReleasedSince, "released-since", "", "limit to coronavirus genomes released after a specified date ("+dateFormat+")")
	summaryVirusGenomeCmd.PersistentFlags().BoolVar(&argJsonLinesFormat, "as-json-lines", false, "Stream results as newline delimited JSON-Lines")
}
