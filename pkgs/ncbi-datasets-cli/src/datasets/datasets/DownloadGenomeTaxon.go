package datasets

import (
	"errors"

	openapi "main/openapi_client"

	"github.com/spf13/cobra"
)

var downloadGenomeTaxonCmd = &cobra.Command{
	Use:   "taxon <taxon>",
	Short: "download a genome dataset by taxon (NCBI Taxonomy ID, scientific or common name at any tax rank)",
	Long: `
Download a genome dataset by taxon (NCBI Taxonomy ID, scientific or common name at any tax rank).  Genome datasets include genome, transcript and protein sequence, annotation and a detailed data report. Datasets are downloaded as a zip file.	

The default genome dataset includes the following files (if available):
* genomic.fna (genomic sequences)
* rna.fna (transcript sequences)
* protein.faa (protein sequences)
* genomic.gff (genome annotation in gff3 format)
* data_report.jsonl (data report with genome assembly and annotation metadata)
* dataset_catalog.json (a list of files and file types included in the dataset)

Refer to NCBI's [command line quickstart](https://www.ncbi.nlm.nih.gov/datasets/docs/quickstarts/command-line-tools/) documentation for information about getting started with the command-line tools.`,
	Example: `  datasets download genome taxon human --chromosomes 21
  datasets download genome taxon "bos taurus"
  datasets download genome taxon 10116 --exclude-seq --exclude-gff3`,
	Args: cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		request := openapi.NewV1AssemblyMetadataRequest()
		request.SetTaxon(args[0])
		err := updateAssemblyMetadataRequestOption(request)
		if err != nil {
			return err
		}

		request.SetReturnedContent(openapi.V1ASSEMBLYMETADATAREQUESTCONTENTTYPE_ASSM_ACC)

		assemblyMetadata, metadata_err := getAssemblyMetadataWithPost(request, false)
		if metadata_err != nil {
			return metadata_err
		}

		var accessions []string
		i := 0
		for _, assemblyMatch := range assemblyMetadata.GetAssemblies() {
			accessions = append(accessions, assemblyMatch.Assembly.GetAssemblyAccession())
			i++
		}
		if len(accessions) == 0 {
			return errors.New("No assembly available")
		}

		req := getDownloadRequest(accessions)
		return downloadAssembly(req, argDownloadFilename)
	},
}

func init() {
	flags := downloadGenomeTaxonCmd.Flags()
	flags.BoolVarP(&argRefseqOnly, "refseq", "", false, "limit to RefSeq (GCF_) assemblies")
	flags.BoolVarP(&argReferenceOnly, "reference", "", false, "limit to reference and representative (GCF_ and GCA_) assemblies")
	registerHiddenBoolPair(flags, &argTaxExactMatch, "tax-exact-match", "t", false, "exclude sub-species when a species-level taxon is specified")
}
