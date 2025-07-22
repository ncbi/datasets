package flags

import (
	openapi "datasets/openapi/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type VirusFilterFlags struct {
	FlagInterface
	annotatedOnly     bool
	completeOnly      bool
	geoLocation       string
	geoUsaState       string
	lineage           string
	refseqOnly        bool
	hostTaxonFlag     *HostTaxonFilterFlag
	releasedAfterFlag *ReleasedAfterFlag
	updatedAfterFlag  *UpdatedAfterFlag
}

func NewVirusFilterFlags(releasedAfterDesc string, updatedAfterDesc string) *VirusFilterFlags {
	vff := &VirusFilterFlags{
		hostTaxonFlag:     NewHostTaxonFilterFlag(),
		releasedAfterFlag: NewReleasedAfterFlag(releasedAfterDesc),
		updatedAfterFlag:  NewUpdatedAfterFlag(updatedAfterDesc),
	}
	return vff
}

func (vff *VirusFilterFlags) RegisterFlags(flags *pflag.FlagSet) {
	flags.BoolVar(&vff.annotatedOnly, "annotated", false, "Limit to annotated genomes")

	flags.BoolVar(&vff.completeOnly, "complete-only", false, "Limit to complete sequences, as defined by submitter")
	flags.StringVar(&vff.geoLocation, "geo-location", "", "Limit to genomes isolated from a specified geographic location (continent or country)")
	flags.StringVar(&vff.geoUsaState, "usa-state", "", "Limit to genomes isolated from a specified U.S. state (two-letter abbreviation)")
	flags.StringVar(&vff.lineage, "lineage", "", "Limit results by Pango lineage (only SARS-CoV-2)")
	flags.BoolVar(&vff.refseqOnly, "refseq", false, "Limit to RefSeq genomes")

	vff.releasedAfterFlag.RegisterFlags(flags)
	vff.updatedAfterFlag.RegisterFlags(flags)
	vff.hostTaxonFlag.RegisterFlags(flags)
}

func (vff *VirusFilterFlags) PreRunE(cmd *cobra.Command, args []string) (err error) {
	for _, f := range []FlagInterface{vff.releasedAfterFlag, vff.updatedAfterFlag, vff.hostTaxonFlag} {
		err := f.PreRunE(cmd, args)
		if err != nil {
			return err
		}
	}
	return nil
}

// Use for summary of data_reports
func (vff *VirusFilterFlags) PrepareDownloadRequest(request *openapi.V2VirusDatasetRequest) {
	request.SetRefseqOnly(vff.refseqOnly)
	request.SetAnnotatedOnly(vff.annotatedOnly)
	request.SetPangolinClassification(vff.lineage)
	request.SetGeoLocation(vff.geoLocation)
	request.SetUsaState(vff.geoUsaState)
	request.SetCompleteOnly(vff.completeOnly)

	request.SetReleasedSince(vff.releasedAfterFlag.ReleasedAfterFlagDateAsTime())
	request.SetUpdatedSince(vff.updatedAfterFlag.UpdatedAfterFlagDateAsTime())
	request.SetHost(vff.hostTaxonFlag.HostTaxIdValue())
}

func (vff *VirusFilterFlags) PrepareSarsProteinDownloadRequest(request *openapi.V2Sars2ProteinDatasetRequest) {
	request.SetRefseqOnly(vff.refseqOnly)
	request.SetAnnotatedOnly(vff.annotatedOnly)
	request.SetPangolinClassification(vff.lineage)
	request.SetGeoLocation(vff.geoLocation)
	request.SetUsaState(vff.geoUsaState)
	request.SetCompleteOnly(vff.completeOnly)

	request.SetReleasedSince(vff.releasedAfterFlag.ReleasedAfterFlagDateAsTime())
	request.SetUpdatedSince(vff.updatedAfterFlag.UpdatedAfterFlagDateAsTime())
	request.SetHost(vff.hostTaxonFlag.HostTaxIdValue())
}

func (vff *VirusFilterFlags) PrepareAnnotationReportRequest(accs []string, taxons []string) *openapi.V2VirusAnnotationReportRequest {
	filter := *openapi.NewV2VirusAnnotationFilter()
	filter.SetRefseqOnly(vff.refseqOnly)
	filter.SetAnnotatedOnly(vff.annotatedOnly)
	filter.SetPangolinClassification(vff.lineage)
	filter.SetGeoLocation(vff.geoLocation)
	filter.SetUsaState(vff.geoUsaState)
	filter.SetCompleteOnly(vff.completeOnly)

	filter.SetReleasedSince(vff.releasedAfterFlag.ReleasedAfterFlagDateAsTime())
	filter.SetUpdatedSince(vff.updatedAfterFlag.UpdatedAfterFlagDateAsTime())
	filter.SetHost(vff.hostTaxonFlag.HostTaxIdValue())

	filter.SetAccessions(accs)
	filter.SetTaxons(taxons)

	request := openapi.NewV2VirusAnnotationReportRequest()
	request.SetFilter(filter)
	request.SetPageSize(1000)
	return request
}

func (vff *VirusFilterFlags) PrepareDatasetReportRequest(accs []string, taxons []string) *openapi.V2VirusDataReportRequest {
	filter := *openapi.NewV2VirusDatasetFilter()
	filter.SetRefseqOnly(vff.refseqOnly)
	filter.SetAnnotatedOnly(vff.annotatedOnly)
	filter.SetPangolinClassification(vff.lineage)
	filter.SetGeoLocation(vff.geoLocation)
	filter.SetUsaState(vff.geoUsaState)
	filter.SetCompleteOnly(vff.completeOnly)

	filter.SetReleasedSince(vff.releasedAfterFlag.ReleasedAfterFlagDateAsTime())
	filter.SetUpdatedSince(vff.updatedAfterFlag.UpdatedAfterFlagDateAsTime())
	filter.SetHost(vff.hostTaxonFlag.HostTaxIdValue())

	filter.SetAccessions(accs)
	filter.SetTaxons(taxons)

	request := openapi.NewV2VirusDataReportRequest()
	request.SetFilter(filter)
	request.SetPageSize(1000)
	return request
}
