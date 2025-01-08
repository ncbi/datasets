package datasets

import (
	openapi "datasets/openapi/v2"
)

type GenomeAccRetriever struct {
	assemblyAccessions []string
	DefaultPageProcessorFuncs[openapi.V2reportsAssemblyDataReport, *openapi.V2reportsAssemblyDataReportPage]
}

func (r *GenomeAccRetriever) ReportName() string {
	return "genome"
}

func (assmRetriever *GenomeAccRetriever) ProcessPage(ppage *openapi.V2reportsAssemblyDataReportPage) {
	for _, report := range ppage.GetReports() {
		assmRetriever.assemblyAccessions = append(assmRetriever.assemblyAccessions, report.GetAccession())
	}
}
