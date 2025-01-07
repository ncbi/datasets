package datasets

import (
	cmdflags "datasets_cli/v2/datasets/flags"

	"github.com/spf13/cobra"
)

type DownloadGeneFlag struct {
	geneIncludeFlag *cmdflags.GeneIncludeFlag
	previewFlag     *cmdflags.DownloadPreviewFlag
	filterFlag      *cmdflags.FastaFilterFlag
	cmdFlagSet      []cmdflags.FlagInterface
}

func newDownloadGeneFlag(default_include_flags []cmdflags.GeneIncludeFlags) DownloadGeneFlag {
	gif := cmdflags.NewGeneIncludeFlag(default_include_flags)
	dpf := cmdflags.NewDownloadPreviewFlag()
	fff := cmdflags.NewFastaFilterFlag()

	dgf := DownloadGeneFlag{
		geneIncludeFlag: gif,
		previewFlag:     dpf,
		filterFlag:      fff,
		cmdFlagSet:      []cmdflags.FlagInterface{gif, dpf, fff},
	}

	return dgf
}

func createGeneCmd() *cobra.Command {
	downloadGeneFlag := newDownloadGeneFlag(cmdflags.GeneDefault)

	cmd := &cobra.Command{
		Use:   "gene",
		Short: "Download a gene data package",
		Example: `  datasets download gene gene-id 672
  datasets download gene symbol brca1 --taxon "mus musculus"
  datasets download gene accession NP_000483.3
  datasets download gene gene-id 2778 --fasta-filter NC_000020.11,NM_001077490.3,NP_001070958.1`,
		Long: `
Download a gene data package.  Gene data packages include gene, transcript and protein sequences and one or more data reports. Data packages are downloaded as a zip archive.

The default gene data package for NM, NR, NP, XM, XR, XP and YP accessions:
  * rna.fna (transcript sequences)
  * protein.faa (protein sequences)
  * data_report.jsonl (data report with gene metadata)
  * dataset_catalog.json (a list of files and file types included in the data package)`,
		Args:              cobra.NoArgs,
		RunE:              ParentCommandRunE,
		PersistentPreRunE: cmdflags.PersistentPreRunEFor(downloadGeneFlag.cmdFlagSet, downloadCmd),
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			if !argNoProgress && !downloadGeneFlag.previewFlag.IsPreview() {
				progress.Stop()
			}
		},
	}

	cmdflags.RegisterAllFlags(downloadGeneFlag.cmdFlagSet, cmd.PersistentFlags())

	// Can't use the downloadGeneFlag created here in download gene accession, because default include values are different there
	cmd.AddCommand(createDownloadGeneGeneIDCmd(downloadGeneFlag))
	cmd.AddCommand(createDownloadGeneSymbolCmd(downloadGeneFlag))
	cmd.AddCommand(createDownloadGeneAccession())
	cmd.AddCommand(createDownloadGeneTaxonCmd(downloadGeneFlag))

	return cmd
}
