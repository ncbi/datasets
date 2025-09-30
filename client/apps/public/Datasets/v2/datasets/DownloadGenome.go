package datasets

import (
	"context"
	"fmt"
	"time"

	openapi "datasets/openapi/v2"
	cmdflags "datasets_cli/v2/datasets/flags"

	"github.com/spf13/cobra"
)

type DownloadSummaryInterface struct {
	ResourceUpdatedOn time.Time                               `json:"resource_updated_on"`
	RecordCount       int32                                   `json:"record_count"`
	HydratedSizeMB    int32                                   `json:"estimated_file_size_mb"`
	AvailableFiles    openapi.V2DownloadSummaryAvailableFiles `json:"included_data_files"`
	genomeApi         *openapi.GenomeAPIService
	summaryRequest    *openapi.V2AssemblyDatasetRequest
	request           *openapi.V2AssemblyDatasetReportsRequest
	Warning           string `json:"warning,omitempty"`

	DefaultPageProcessorFuncs[openapi.V2reportsAssemblyDataReport, *openapi.V2reportsAssemblyDataReportPage]
}

func (r *DownloadSummaryInterface) ReportName() string {
	return "genome"
}

func (downloadSummary *DownloadSummaryInterface) ProcessPage(ppage *openapi.V2reportsAssemblyDataReportPage) {
	var accessions []string
	for _, report := range ppage.GetReports() {
		accessions = append(accessions, report.GetAccession())
	}

	var ctx context.Context

	// call GenomeDownloadSummary with all the accessions for this page
	downloadSummary.summaryRequest.SetAccessions(accessions)
	summaryResponse, _, err := downloadSummary.genomeApi.GenomeDownloadSummaryByPost(ctx).V2AssemblyDatasetRequest(*(downloadSummary.summaryRequest)).Execute()
	if err != nil {
		fmt.Printf("GenomeDownloadSummaryByPost() failed %v\n", err)
	}

	// update record count
	downloadSummary.RecordCount = downloadSummary.RecordCount + *summaryResponse.RecordCount

	// update ResourceUpdatedOn
	downloadSummary.ResourceUpdatedOn = *summaryResponse.ResourceUpdatedOn

	hydrated, isHydratedValueSet := summaryResponse.GetHydratedOk()
	if isHydratedValueSet {
		downloadSummary.HydratedSizeMB = downloadSummary.HydratedSizeMB + hydrated.GetEstimatedFileSizeMb()
	}

	availableFiles, hasAvailableFiles := summaryResponse.GetAvailableFilesOk()

	if hasAvailableFiles {
		annotationTypeIncludes := func(fl cmdflags.GenomeIncludeAnnotation) bool {
			for _, v := range downloadSummary.summaryRequest.IncludeAnnotationType {
				if v == cmdflags.GenomeIncludeAnnotationOpenapi[fl] {
					return true
				}
			}
			return false
		}

		// all_genomic_fasta
		if availableFiles.HasAllGenomicFasta() && annotationTypeIncludes(cmdflags.GenomeSeq) {
			downloadSummaryAllGenomicFasta := downloadSummary.AvailableFiles.GetAllGenomicFasta()
			downloadSummary.AvailableFiles.SetAllGenomicFasta(updateFileSummary(downloadSummaryAllGenomicFasta, availableFiles.GetAllGenomicFasta()))
		}

		// genome_gff
		if availableFiles.HasGenomeGff() && annotationTypeIncludes(cmdflags.GenomeGff3) {
			downloadSummaryGenomeGff := downloadSummary.AvailableFiles.GetGenomeGff()
			downloadSummary.AvailableFiles.SetGenomeGff(updateFileSummary(downloadSummaryGenomeGff, availableFiles.GetGenomeGff()))
		}

		//rna_fasta
		if availableFiles.HasRnaFasta() && annotationTypeIncludes(cmdflags.GenomeRna) {
			downloadSummaryRnaFasta := downloadSummary.AvailableFiles.GetRnaFasta()
			downloadSummary.AvailableFiles.SetRnaFasta(updateFileSummary(downloadSummaryRnaFasta, availableFiles.GetRnaFasta()))
		}

		//prot_fasta
		if availableFiles.HasProtFasta() && annotationTypeIncludes(cmdflags.GenomeProtein) {
			downloadSummaryProtFasta := downloadSummary.AvailableFiles.GetProtFasta()
			downloadSummary.AvailableFiles.SetProtFasta(updateFileSummary(downloadSummaryProtFasta, availableFiles.GetProtFasta()))
		}

		//genome_gtf
		if availableFiles.HasGenomeGtf() && annotationTypeIncludes(cmdflags.GenomeGtf) {
			downloadSummaryGenomeGtf := downloadSummary.AvailableFiles.GetGenomeGtf()
			downloadSummary.AvailableFiles.SetGenomeGtf(updateFileSummary(downloadSummaryGenomeGtf, availableFiles.GetGenomeGtf()))
		}

		//cds_fasta
		if availableFiles.HasCdsFasta() && annotationTypeIncludes(cmdflags.GenomeCds) {
			downloadSummaryCdsFasta := downloadSummary.AvailableFiles.GetCdsFasta()
			downloadSummary.AvailableFiles.SetCdsFasta(updateFileSummary(downloadSummaryCdsFasta, availableFiles.GetCdsFasta()))
		}
	}
}

func updateFileSummary(downloadSummaryFileSummary openapi.V2DownloadSummaryFileSummary, fileSummary openapi.V2DownloadSummaryFileSummary) openapi.V2DownloadSummaryFileSummary {
	fileCount := fileSummary.GetFileCount()
	sizeMb := fileSummary.GetSizeMb()
	downloadSummaryFileSummary.SetFileCount(downloadSummaryFileSummary.GetFileCount() + fileCount)
	downloadSummaryFileSummary.SetSizeMb(downloadSummaryFileSummary.GetSizeMb() + sizeMb)
	return downloadSummaryFileSummary
}

func getDownloadSummary(dgf DownloadGenomeFlag, requestIter RequestIterator[*openapi.V2AssemblyDatasetReportsRequest]) (err error) {
	api, apiErr := getGenomeApi()
	if apiErr != nil {
		return apiErr
	}
	var summaryRequest = openapi.NewV2AssemblyDatasetRequest()
	arg_warn, arg_err := SetRequestArgs(summaryRequest, dgf)
	if arg_err != nil {
		return arg_err
	}

	downloadSummaryRetriever := DownloadSummaryInterface{request: requestIter.GetRequest(), Warning: arg_warn}
	downloadSummaryRetriever.summaryRequest = summaryRequest
	downloadSummaryRetriever.genomeApi = api.genomeApi

	_, metaErr := ProcessAllPages[*openapi.V2AssemblyDatasetReportsRequest,
		openapi.V2reportsAssemblyDataReportPage,
		openapi.V2reportsAssemblyDataReport,
		*openapi.V2reportsAssemblyDataReportPage](
		requestIter,
		api,
		&downloadSummaryRetriever,
	)

	if !argNoProgress {
		progress.Stop()
	}
	if metaErr == nil {
		printResults(downloadSummaryRetriever)
	}
	return metaErr
}

type DownloadGenomeFlag struct {
	genomeIncludeAnnotationFlag *cmdflags.GenomeIncludeAnnotationFlag
	skipValidationFlag          *cmdflags.SkipZipValidation
	downloadPreviewFlag         *cmdflags.DownloadPreviewFlag
	dehydratedFlag              *cmdflags.DehydratedFlag
	chromosomesFlag             *cmdflags.ChromosomesFlag

	cmdFlagSet []cmdflags.FlagInterface
}

func initDownloadGenomeFlag() DownloadGenomeFlag {
	svf := cmdflags.NewSkipZipValidationFlag()
	giaf := cmdflags.NewGenomeIncludeAnnotationFlag([]cmdflags.GenomeIncludeAnnotation{cmdflags.GenomeSeq})
	dpf := cmdflags.NewDownloadPreviewFlag()
	df := cmdflags.NewDehydratedFlag()
	cf := cmdflags.NewChromosomesFlag()

	dgf := DownloadGenomeFlag{
		skipValidationFlag:          svf,
		genomeIncludeAnnotationFlag: giaf,
		downloadPreviewFlag:         dpf,
		dehydratedFlag:              df,
		chromosomesFlag:             cf,

		cmdFlagSet: []cmdflags.FlagInterface{svf, giaf, dpf, df, cf},
	}

	return dgf
}

func createGenomeCmd() *cobra.Command {
	downloadGenomeFlag := initDownloadGenomeFlag()
	assemblyRequestFlag := initAssemblyRequestFlag()
	allFlags := append(downloadGenomeFlag.cmdFlagSet, assemblyRequestFlag.cmdFlagSet...)

	cmd := &cobra.Command{
		Use:   "genome",
		Short: "Download a genome data package",
		Example: `  datasets download genome accession GCF_000001405.40 --chromosomes X,Y --include genome,gff3,rna
  datasets download genome taxon "bos taurus" --dehydrated
  datasets download genome taxon human --assembly-level chromosome,complete --dehydrated
  datasets download genome taxon "house mouse" --search C57BL/6J --search "Broad Institute" --dehydrated`,
		Long: `
Download a genome data package. Genome data packages may include genome, transcript and protein sequences, annotation and one or more data reports. Data packages are downloaded as a zip archive.

The default genome data package includes the following files:
  * <accession>_<assembly_name>_genomic.fna (genomic sequences)
  * assembly_data_report.jsonl (data report with genome assembly and annotation metadata)
  * dataset_catalog.json (a list of files and file types included in the data package)`,
		Args:              cobra.NoArgs,
		PersistentPreRunE: cmdflags.PersistentPreRunEFor(allFlags, downloadCmd),
		RunE:              ParentCommandRunE,
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			if !argNoProgress && !downloadGenomeFlag.downloadPreviewFlag.IsPreview() {
				progress.Stop()
			}
		},
	}

	cmd.AddCommand(createDownloadGenomeAccessionCmd(downloadGenomeFlag, assemblyRequestFlag))
	cmd.AddCommand(createDownloadGenomeTaxonCmd(downloadGenomeFlag, assemblyRequestFlag))

	cmdflags.RegisterAllFlags(allFlags, cmd.PersistentFlags())

	return cmd
}
