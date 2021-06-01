package datasets

import (
	"errors"

	"github.com/spf13/cobra"

	openapi "main/openapi_client"
)

func cmdRunSummaryOrthologAccession(cmd *cobra.Command, args []string) (err error) {
	req := new(openapi.V1alpha1GeneDatasetRequest)
	req.Accessions = argIDArgs
	var geneInts []int64
	geneInts, err = allGeneIdForRequest(req)
	if err != nil {
		return
	}
	if len(geneInts) == 0 {
		err = errors.New("No genes found for search term")
		return
	}
	if !argJsonFormat {
		err = streamOrthologs(geneInts, openapi.V1ALPHA1ORTHOLOGREQUESTCONTENTTYPE_COMPLETE, &JsonLinesStreamProcessor{})
	} else {
		err = streamOrthologs(geneInts, openapi.V1ALPHA1ORTHOLOGREQUESTCONTENTTYPE_COMPLETE, &JsonStreamProcessor{wrapperName: "orthologs"})
	}

	return
}

var summaryOrthologAccessionCmd = &cobra.Command{
	Use:   "accession <accession>",
	Short: "print a summary of an ortholog dataset by RefSeq nucleotide or protein accession",
	Long:  `Print a summary of an ortholog dataset by RefSeq nucleotide or protein accession. Ortholog data is calculated by NCBI for vertebrates and insects. The summary is returned in JSON format.`,
	Example: `  datasets summary ortholog accession NP_000483.3
  datasets summary ortholog accession NM_000546.6`,
	RunE: cmdRunSummaryOrthologAccession,
}

func init() {
	flags := summaryOrthologAccessionCmd.Flags()
	registerHiddenStringSlicePair(flags, &taxonFilter, "taxon-filter", "", []string{}, "limit results to ortholog data for a specified taxonomic group")
	registerHiddenStringPair(flags, &argInputFile, "inputfile", "i", "", "read a list of RefSeq nucleotide or protein accessions from a file to use as input")
}
