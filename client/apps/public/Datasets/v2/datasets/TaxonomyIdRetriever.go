package datasets

import (
	openapi "datasets/openapi/v2"
)

type TaxonomyIdRetriever struct {
	taxIdsMap map[int32]bool
	taxRanks  []openapi.V2reportsRankType
	DefaultPageProcessorFuncs[openapi.V2reportsTaxonomyReportMatch, *openapi.V2reportsTaxonomyDataReportPage]
}

func NewTaxonomyIdRetriever() TaxonomyIdRetriever {
	return TaxonomyIdRetriever{
		taxIdsMap: make(map[int32]bool),
	}
}

func (taxidRetriever *TaxonomyIdRetriever) ReportName() string {
	return "taxonomy"
}

func (taxidRetriever *TaxonomyIdRetriever) ProcessPage(ppage *openapi.V2reportsTaxonomyDataReportPage) {
	for _, report := range ppage.GetReports() {
		if report.HasTaxonomy() {
			// Don't include the root taxid
			if report.Taxonomy.GetTaxId() != 1 {
				taxidRetriever.taxIdsMap[report.Taxonomy.GetTaxId()] = true
			}
		}
	}
}

func (taxidRetriever *TaxonomyIdRetriever) AddTaxIds(taxIds []int32) {
	for _, taxId := range taxIds {
		taxidRetriever.taxIdsMap[taxId] = true
	}
}

func (taxidRetriever *TaxonomyIdRetriever) SetRanks(ranks []openapi.V2reportsRankType) {
	taxidRetriever.taxRanks = ranks
}

func (taxidRetriever *TaxonomyIdRetriever) GetRanks() []openapi.V2reportsRankType {
	return taxidRetriever.taxRanks
}

func (taxidRetriever *TaxonomyIdRetriever) GetTaxIds() []int32 {
	taxIds := make([]int32, len(taxidRetriever.taxIdsMap))
	i := 0
	for taxId := range taxidRetriever.taxIdsMap {
		taxIds[i] = taxId
		i++
	}

	return taxIds
}
