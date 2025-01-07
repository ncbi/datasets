package datasets

import (
	openapi "datasets/openapi/v2"
	"strconv"
)

// This implements the PageProcessor interface.
// Client-side paged geneids can return dupicates, so we use a map.
type GeneIdRetriever struct {
	geneIdsMap map[int32]bool
	DefaultPageProcessorFuncs[openapi.V2reportsGeneReportMatch, *openapi.V2reportsGeneDataReportPage]
}

func (r *GeneIdRetriever) ReportName() string {
	return "gene"
}

func (geneIdRetriever *GeneIdRetriever) ProcessPage(ppage *openapi.V2reportsGeneDataReportPage) {
	for _, report := range ppage.GetReports() {
		gene := report.GetGene()
		// While 32 specifies what the integer must fit, it still returns an in64, thus the additional cast
		geneId64, e := strconv.ParseInt(gene.GetGeneId(), 10, 32)
		if e == nil {
			// Ignore errors
			geneIdRetriever.geneIdsMap[int32(geneId64)] = true
		}
	}
}

func (geneIdRetriever *GeneIdRetriever) GetGeneIds() []int32 {
	keys := make([]int32, len(geneIdRetriever.geneIdsMap))
	i := 0
	for k := range geneIdRetriever.geneIdsMap {
		keys[i] = k
		i++
	}
	return keys
}

func (geneIdRetriever *GeneIdRetriever) RetrievalCount(totalCount int) int {
	return len(geneIdRetriever.geneIdsMap)
}

func NewGeneIdRetriever() GeneIdRetriever {
	return GeneIdRetriever{
		geneIdsMap: make(map[int32]bool),
	}
}

type GeneCounts struct {
	Genes       int `json:"genes"`
	Transcripts int `json:"transcripts"`
	Proteins    int `json:"proteins"`
}

type GeneCountRetriever struct {
	geneCounts GeneCounts
	geneIdsMap map[string]bool
	DefaultPageProcessorFuncs[openapi.V2reportsGeneReportMatch, *openapi.V2reportsGeneDataReportPage]
}

func (r *GeneCountRetriever) ReportName() string {
	return "gene"
}

func (geneCountRetriever *GeneCountRetriever) ProcessPage(ppage *openapi.V2reportsGeneDataReportPage) {
	for _, report := range ppage.GetReports() {
		gene := report.GetGene()
		if !geneCountRetriever.geneIdsMap[gene.GetGeneId()] {
			geneCountRetriever.geneCounts.Genes += 1
			geneCountRetriever.geneCounts.Transcripts += int(gene.GetTranscriptCount())
			geneCountRetriever.geneCounts.Proteins += int(gene.GetProteinCount())
			geneCountRetriever.geneIdsMap[gene.GetGeneId()] = true
		}
	}
}

func NewGeneCountRetriever() GeneCountRetriever {
	return GeneCountRetriever{
		geneIdsMap: make(map[string]bool),
	}
}
