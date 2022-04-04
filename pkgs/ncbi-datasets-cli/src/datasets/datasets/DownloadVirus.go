package datasets

import (
	"github.com/spf13/cobra"
)

var (
	argAnnotatedOnly      bool
	argReleasedSince      string
	argHost               string
	argLineage            string
	argGeoLocation        string
	argCompleteOnly       bool
	argRetiredExcludeFlag bool
	argRetiredIncludeFlag bool
)

const dateFormat = "MM/DD/YYYY"

const messagePreamble = `
NCBI Datasets has removed the following file types from the SARS-CoV-2 Data Package:
 * PDB files (*.pdb)
 * GenBank flatfiles (genomic.gbff)
 * GenPept flatfiles (genpept.gpff)

Questions? Please contact us: info@ncbi.nlm.nih.gov

`

const virusFlagWarningMessage = messagePreamble + `
The following flags are no longer needed
 * --exclude-pdb
 * --exclude-gpff`

const virusFlagErrorMessage = messagePreamble + `
You must remove this flag to continue:
 * --include-gbff`

// virusCmd represents the virus command
var downloadVirusCmd = &cobra.Command{
	Use:   "virus [command]",
	Short: "download a coronavirus dataset",
	Long: `
Download a coronavirus genome or SARS-CoV-2 protein dataset as a zip file.

Refer to NCBI's [command line quickstart](https://www.ncbi.nlm.nih.gov/datasets/docs/quickstarts/command-line-tools/) documentation for information about getting started with the command-line tools. `,
	Example: `  datasets download virus genome taxon sars-cov-2 --host dog
  datasets download virus protein S --host dog --filename SARS2-spike-dog.zip`,
}

func init() {
	downloadVirusCmd.AddCommand(downloadVirusGenomeCmd)
	downloadVirusCmd.AddCommand(downloadVirusProteinCmd)

	downloadVirusCmd.PersistentFlags().BoolVarP(&argRetiredExcludeFlag, "exclude-gpff", "g", false, "exclude genomic.fna (genomic sequence file)")
	downloadVirusCmd.PersistentFlags().MarkHidden("exclude-gpff")

	downloadVirusCmd.PersistentFlags().BoolVarP(&argRetiredIncludeFlag, "include-gbff", "b", false, "include genomic.gbff (genome sequence and annotation in GenBank flat file format)")
	downloadVirusCmd.PersistentFlags().MarkHidden("include-gbff")

	downloadVirusCmd.PersistentFlags().BoolVar(&argRetiredExcludeFlag, "exclude-pdb", false, "exclude *.pdb (protein structure files)")
	downloadVirusCmd.PersistentFlags().MarkHidden("exclude-pdb")
}
