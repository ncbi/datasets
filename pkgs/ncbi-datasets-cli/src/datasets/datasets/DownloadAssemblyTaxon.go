package datasets

import (
	"errors"

	openapi "main/openapi_client"

	"github.com/spf13/cobra"
)

var downloadAssmTaxonCmd = &cobra.Command{
	Deprecated: "please use \"datasets download genome taxon\".",
	Use:        "taxon <taxon>",
	Short:      "Download assembly data by taxonomy id, scientific or common name at any tax rank",
	Example:    "  datasets download assembly taxon 6689",
	Long: `
Download assembly data by taxid. Data is returned as a zip archive.
The default download package for a given assembly (or set of assemblies) includes all chromosomes and
unlocalized sequences and excludes any available annotation data.

Refer to NCBI's [command line quickstart](https://www.ncbi.nlm.nih.gov/datasets/docs/quickstarts/command-line-tools/) documentation for information about getting started with the command-line tools.`,
	Args: cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		request := openapi.NewV1AssemblyMetadataRequest()
		request.SetTaxon(args[0])
		err := updateAssemblyMetadataRequestOption(request)
		if err != nil {
			return err
		}

		request.SetReturnedContent(openapi.V1ASSEMBLYMETADATAREQUESTCONTENTTYPE_ASSM_ACC)

		assemblyMetadata, post_err := getAssemblyMetadataWithPost(request, false)
		if post_err != nil {
			return post_err
		}

		var accessions []string
		i := 0
		for _, assemblyMatch := range assemblyMetadata.GetAssemblies() {
			accessions = append(accessions, assemblyMatch.Assembly.GetAssemblyAccession())
			i++
		}
		if len(accessions) == 0 {
			err = errors.New("No assembly available")
			return err
		}

		req := getDownloadRequest(accessions)
		return downloadAssembly(req, argDownloadFilename)
	},
}

func init() {
	downloadAssmTaxonCmd.PersistentFlags().BoolVarP(&argRefseqOnly, "refseq", "", false, "if true, only download RefSeq (GCF_) assemblies")
	downloadAssmTaxonCmd.PersistentFlags().BoolVarP(&argTaxExactMatch, "tax-exact-match", "t", false, "use this flag with species-level taxids. when this flag is specified, sub-species are excluded")
}
