package dataformat

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	datasets_util "main/datasets/util"
	datasets_command "main/datasets/util/command"
)

var (
	debug       bool
	AppVersion  = "undefined"
	userMessage string
	tableFields []string
	inputFile   string
	packageFile string
)

var rootCmd = &cobra.Command{
	Use:   "dataformat",
	Short: "NCBI Datasets Dataformat",
	Long: `dataformat is a command-line tool to convert JSON-lines formatted NCBI Datasets reports into other formats (Excel, TSV).

To learn more about what fields are available in the reports and how to convert them into 
tabular data, visit the [data report schema](https://www.ncbi.nlm.nih.gov/datasets/docs/data-report-schemas-overview/) documentation.

Refer to NCBI's [command line start](https://www.ncbi.nlm.nih.gov/datasets/docs/command-line-start) documentation for information about getting started with the command-line tools.`,
	Args: cobra.MaximumNArgs(0),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// set the default value for the --fields flag, specific to the command specified
		rpt := GetReport(cmd.Use)
		if rpt == nil {
			return nil
		}
		fields := rpt.GetAllFieldMnemonics()
		if f := cmd.Flag("fields"); f != nil {
			f.DefValue = strings.Join(fields, ",")
		}
		if len(tableFields) == 0 {
			tableFields = fields
		}
		if len(rpt.defaultPackagePath) > 0 {
			if f := cmd.Flag("inputfile"); f != nil {
				f.DefValue = rpt.defaultPackagePath
				if len(inputFile) == 0 {
					inputFile = rpt.defaultPackagePath
				}
			}
		}
		return nil
	},
}

func inputFilePreRunE(cmd *cobra.Command, args []string) (err error) {
	reqFlags := []string{"inputfile", "package"}
	ok := false
	flags := cmd.Flags()
	for _, c := range reqFlags {
		if flags.Changed(c) {
			ok = true
			break
		}
	}
	if !ok {
		err = fmt.Errorf("at least one of the following flags must be specified: %s", reqFlags)
	}
	return
}

func addFileInputFlags(cmd *cobra.Command) {
	pflags := cmd.PersistentFlags()
	pflags.StringVar(&inputFile, "inputfile", "", "input file")
	//cmd.MarkPersistentFlagRequired("inputfile")
	pflags.StringVar(&packageFile, "package", "", "datasets package (zip archive), inputfile parameter is relative to the root path inside the archive")
}

var catalogCmd = &cobra.Command{
	Use:     "catalog [flags]",
	Short:   "print the package catalog",
	Long:    "Print the datasets package catalog.",
	Args:    cobra.MaximumNArgs(0),
	PreRunE: inputFilePreRunE,
	RunE: func(cmd *cobra.Command, args []string) error {
		return emitCatalog()
	},
}

var versionCmd = &cobra.Command{
	Use:   "version [flags]",
	Short: "print the version of this client and exit",
	Long:  "Print the version of this client and exit.",
	Args:  cobra.MaximumNArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(AppVersion)
		return nil
	},
}

// return top-level cobra command, used for documenting the CLI
func ReturnTopLevelCobraCommand() *cobra.Command {
	return rootCmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	exitval := 0
	err := rootCmd.Execute()
	if err != nil {
		exitval = 1
	}
	if userMessage != "" {
		fmt.Fprintln(os.Stderr, userMessage)
		exitval = 1
	}
	os.Exit(exitval)
}

func outputFieldNames(reportName, flagContent string) (retval []string) {
	if rpt := GetReport(reportName); rpt != nil {
		retval = rpt.GetAllFieldMnemonics()
	}
	return
}

func init() {
	rptCmds := make([]string, len(AllReports))
	for i, rpt := range AllReports {
		rptCmds[i] = rpt.cmd
	}
	reportsSection := &datasets_util.UsageSection{
		SectionText: "Report Commands",
		Commands:    rptCmds,
	}
	datasets_util.AddUsageSections(
		"dataformat",
		&datasets_util.UsageSections{
			&datasets_util.UsageSection{
				SectionText: "Formatting Commands",
				Commands:    []string{"tsv", "excel"},
			},
			&datasets_util.UsageSection{
				SectionText: "Data extraction Commands",
				Commands:    []string{"catalog"},
			},
			&datasets_util.UsageSection{
				SectionText: "Miscellaneous Commands",
				Commands:    []string{"completion", "version", "help"},
			},
		},
	)
	datasets_util.AddUsageSections(
		"tsv",
		&datasets_util.UsageSections{reportsSection},
	)
	datasets_util.AddUsageSections(
		"excel",
		&datasets_util.UsageSections{reportsSection},
	)
	datasets_util.InitRootCommand(rootCmd)

	pflags := rootCmd.PersistentFlags()
	pflags.BoolVar(&debug, "debug", false, "Emit debugging info")
	pflags.MarkHidden("debug")

	cobra.EnableCommandSorting = false

	rootCmd.AddCommand(getTSVCmd())
	rootCmd.AddCommand(getExcelCmd())

	addFileInputFlags(catalogCmd)
	rootCmd.AddCommand(catalogCmd)

	rootCmd.AddCommand(datasets_command.NewAutocompleteCmd(rootCmd))
	rootCmd.AddCommand(versionCmd)

	datasets_util.AddFlagAdjustment(&datasets_util.FlagAdjustment{
		RegexpMatch: `--fields strings[ ]+`,
		Content:     outputFieldNames,
	})
}
