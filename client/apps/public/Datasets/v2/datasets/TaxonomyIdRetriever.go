package datasets

import (
	openapi "datasets/openapi/v2"
)

type TaxonomyIdRetriever struct {
	taxIdsMap      map[int32]bool
	includeParents bool
	taxRanks       []openapi.V2reportsRankType
	DefaultPageProcessorFuncs[openapi.V2reportsTaxonomyReportMatch, *openapi.V2reportsTaxonomyDataReportPage]
}

func NewTaxonomyIdRetriever() TaxonomyIdRetriever {
	return TaxonomyIdRetriever{
		taxIdsMap:      make(map[int32]bool),
		includeParents: false,
	}
}

func (taxidRetriever *TaxonomyIdRetriever) SetIncludeParents(lineage bool) {
	taxidRetriever.includeParents = lineage
}

func (taxidRetriever *TaxonomyIdRetriever) ReportName() string {
	return "taxonomy"
}

func (taxidRetriever *TaxonomyIdRetriever) ProcessPage(ppage *openapi.V2reportsTaxonomyDataReportPage) {
	for _, report := range ppage.GetReports() {
		if report.HasTaxonomy() {
			taxidRetriever.taxIdsMap[report.Taxonomy.GetTaxId()] = true
			if taxidRetriever.includeParents && report.Taxonomy.HasParents() {
				for _, taxId := range report.Taxonomy.GetParents() {
					if taxId != 1 {
						// No information in the record to check rank here (if set) so that will need to be done
						// in a later step
						taxidRetriever.taxIdsMap[int32(taxId)] = true
					}
				}
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
