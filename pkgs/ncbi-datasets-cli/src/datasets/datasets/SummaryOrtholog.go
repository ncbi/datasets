package datasets

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"

	openapi "main/openapi_client"
)

var (
	taxonFilter []string
)

func fetchOrthologByGeneId(geneID int32, contentType openapi.V1OrthologRequestContentType) (orthologs *openapi.V1OrthologSet, err error) {
	cli, err := createOAClient()
	if err != nil {
		return
	}
	request := cli.GeneApi.GeneOrthologsById(nil, geneID)
	request.ReturnedContent(contentType)
	if len(taxonFilter) > 0 {
		request.TaxonFilter(taxonFilter)
	}

	const retry_count = 3
	var retry_delay = time.Duration(10)
	for i := 1; i <= retry_count; i++ {
		result, resp, retry_err := request.Execute()
		err = handleHTTPResponse(resp, retry_err)
		if err == nil {
			orthologs = &result
			return
		}
		if argDebug {
			fmt.Fprintf(os.Stderr, "Error: '%s' in ortholog download. Retrying: %d of %d\n", err.Error(), i, retry_count)
		}
		time.Sleep(retry_delay * time.Millisecond)
		retry_delay *= 2
	}
	return
}

func streamOrthologs(geneInts []int32, contentType openapi.V1OrthologRequestContentType, sp streamProcessorOrtholog) error {
	firstLine := true
	orthologSets := map[int32]bool{}
	sp.insertRecordBegin()
	for _, gene_id := range geneInts {
		orthologs, err := fetchOrthologByGeneId(gene_id, contentType)
		if err != nil {
			return err
		}

		if orthologs.GetOrthologSetId() == 0 {
			// If the input gene id was invalid, the message will be stored in first gene of the returned OrthologSet
			// If the gene id was valid, but there was no related ortholog set, an empty ortholog struct is returned
			if len(orthologs.Genes.GetGenes()) > 0 {
				printMessage(&orthologs.Genes.GetGenes()[0])
			}
			continue
		}

		// Skip ortholog sets that have already been added
		if orthologSets[orthologs.GetOrthologSetId()] {
			continue
		}

		orthologSets[orthologs.GetOrthologSetId()] = true

		if !firstLine {
			sp.insertRecordDelimiter()
		}
		firstLine = false
		printResultsNoNewline(orthologs)
	}
	sp.insertRecordEnd()

	return nil
}

var summaryOrthologCmd = &cobra.Command{
	Use:   "ortholog",
	Short: "print a summary of an ortholog dataset",
	Long: `Print a summary of an ortholog dataset by NCBI Gene ID. Ortholog data is calculated by NCBI for vertebrates and insects. The summary is returned in JSON format.

Refer to NCBI's [command line quickstart](https://www.ncbi.nlm.nih.gov/datasets/docs/quickstarts/command-line-tools/) documentation for information about getting started with the command-line tools.`,

	PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
		argIDArgs, err = getArgsFromListOrFile(args, argInputFile)
		return
	},
}

func init() {
	summaryOrthologCmd.AddCommand(summaryOrthologIDCmd)
	summaryOrthologCmd.AddCommand(summaryOrthologSymbolCmd)
	summaryOrthologCmd.AddCommand(summaryOrthologAccessionCmd)

	pflags := summaryOrthologCmd.PersistentFlags()
	pflags.BoolVar(&argJsonFormat, "as-json", false, "Stream results as a JSON object instead of newline-delimited JSON-Lines")
}
