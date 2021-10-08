package datasets

import (
	openapi "datasets_cli/v1/openapi"
	"fmt"
	"sort"
	"strings"

	pb_datasets "ncbi/datasets/v1"

	"github.com/antihax/optional"
	"github.com/metakeule/fmtdate"
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
)

var (
	argGenomeReleasedSince string
	argGenomeReleasedUntil string
	argSearchText          []string
	argAssemblyLevel       string
	argAssemblySource      string

	argAnnotatedGenomesOnly bool
)

const (
	MAX_PAGE_SIZE int = 1000
)

func getAssemblyLevelStrings() (result []string) {
	for name := range pb_datasets.AssemblyDatasetDescriptorsFilter_AssemblyLevel_value {
		result = append(result, name)
	}
	sort.Strings(result)

	return
}

func checkAssemblyLevel(requestLevel string) (*string, error) {
	requestLower := strings.ToLower(requestLevel)
	for level := range pb_datasets.AssemblyDatasetDescriptorsFilter_AssemblyLevel_value {
		if requestLower == string(level) {
			return &level, nil
		}
	}

	err := fmt.Errorf("Request level %s is not valid assembly level.", requestLevel)
	return nil, err
}

func getAssemblyLevels() (assemblyLevels []openapi.V1AssemblyDatasetDescriptorsFilterAssemblyLevel, err error) {
	if argAssemblyLevel != "" {
		levels := strings.Split(argAssemblyLevel, ",")
		for _, level := range levels {
			checkedLevel, level_err := checkAssemblyLevel(level)
			if level_err != nil {
				return assemblyLevels, level_err
			}
			assemblyLevels = append(assemblyLevels, openapi.V1AssemblyDatasetDescriptorsFilterAssemblyLevel(*checkedLevel))
		}
		if len(assemblyLevels) > 0 {
			return assemblyLevels, nil
		}
	}
	return assemblyLevels, nil
}

func getSearchStrings() optional.Interface {
	if len(argSearchText) != 0 {
		return optional.NewInterface(argSearchText)
	}
	return optional.EmptyInterface()
}

func getAssemblySource() (string, error) {
	if argAssemblySource != "" {
		if argRefseqOnly {
			err := fmt.Errorf("The deprecated '--refseq' option cannot be used in conjunction with '--assembly-source'")
			return "", err
		}

		requestLower := strings.ToLower(argAssemblySource)
		for asmSource := range pb_datasets.AssemblyDatasetDescriptorsFilter_AssemblySource_value {
			if requestLower == string(asmSource) {
				return requestLower, nil
			}
		}
		err := fmt.Errorf("Source %s is not valid assembly source.", argAssemblySource)
		return "", err
	} else if argRefseqOnly {
		return string(openapi.V1ASSEMBLYDATASETDESCRIPTORSFILTERASSEMBLYSOURCE_REFSEQ), nil
	}

	return string(openapi.V1ASSEMBLYDATASETDESCRIPTORSFILTERASSEMBLYSOURCE_ALL), nil
}

func getDate(argDate string) (optional.Time, error) {
	if argDate != "" {
		formattedDate, err := fmtdate.Parse(dateFormat, argDate)
		if err != nil {
			return optional.EmptyTime(), err
		}
		return optional.NewTime(formattedDate), nil
	}
	return optional.EmptyTime(), nil
}

func addGenomeMetadataFilters(pflags *flag.FlagSet) {
	pflags.BoolVar(&argRefseqOnly, "refseq", false, "limit to RefSeq (GCF_) assemblies")
	pflags.MarkHidden("refseq")
	pflags.BoolVar(&argReferenceOnly, "reference", false, "limit to reference and representative (GCF_ and GCA_) assemblies")
	pflags.BoolVarP(&argAnnotatedGenomesOnly, "annotated", "a", false, "only include genomes with annotation")
	pflags.StringVarP(&argGenomeReleasedSince, "released-since", "", "", "only include genomes that have been released after a specified date ("+dateFormat+")")
	pflags.StringVarP(&argGenomeReleasedUntil, "released-before", "", "", "only include genomes that have been released before a specified date ("+dateFormat+")")
	pflags.StringVarP(&argAssemblySource, "assembly-source", "", "", "restrict assemblies to refseq or genbank only")
	assemblyLevels := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(getAssemblyLevelStrings())), ", "), "[]")
	pflags.StringVarP(&argAssemblyLevel, "assembly-level", "", "", "restrict assemblies to a comma-separated list of one or more of: "+assemblyLevels)
	pflags.StringSliceVar(&argSearchText, "search", []string{}, `only include genomes that have the specified text in the
searchable fields: species and infraspecies, assembly name and submitter
To provide multiple strings '--search' can be included multiple times`)
}

var summaryGenomeCmd = &cobra.Command{
	Use:   "genome",
	Short: "print a summary of a genome dataset",
	Long: `
Print a summary of a genome dataset by assembly accession, bioproject accession or taxon. The summary is returned in JSON format.

Refer to NCBI's [command line quickstart](https://www.ncbi.nlm.nih.gov/datasets/docs/quickstarts/command-line-tools/) documentation for information about getting started with the command-line tools.`,
	Example: `  datasets summary genome accession GCF_000001405.39
  datasets summary genome taxon mouse
  datasets summary genome taxon human --assembly-level chromosome,complete_genome
  datasets summary genome taxon mouse --search C57BL/6J --search "Broad Institute"`,
}

func init() {
	summaryGenomeCmd.AddCommand(summaryGenomeAccessionCmd)
	summaryGenomeCmd.AddCommand(summaryGenomeTaxonCmd)

	pflags := summaryGenomeCmd.PersistentFlags()
	registerHiddenStringPair(pflags, &argLimit, "limit", "l", "all", `limit the number of genome summaries returned
* "all": returns all matching genome summaries
* "none": returns only the number of matching genome summaries
* a number: returns the specified number of matching genome summaries
`)
	pflags.BoolVar(&argAssmAccsOnly, "assmaccs", false, "return only assembly accessions")
	pflags.BoolVar(&argJsonLinesFormat, "as-json-lines", false, "Stream results as newline delimited JSON-Lines")

	addGenomeMetadataFilters(pflags)
}
