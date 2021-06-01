package dataformat

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var (
	format string
)

type DelimTableWriter struct {
	TableWriter
	delim string
}

func newDelimWriter(delim string) *DelimTableWriter {
	writer := DelimTableWriter{
		delim: delim,
	}
	return &writer
}

func (t *DelimTableWriter) emitTableHeader(rspec *ReportSpec, fields []string) {
	headerFields := make([]string, 0)
	for _, f := range fields {
		colspec := rspec.getColumn(f)
		headerFields = append(headerFields, colspec.Name)
	}
	fmt.Printf("%s\n", strings.Join(headerFields, t.delim))
}

func (t *DelimTableWriter) emitTableRow(objIter *ObjIter, rspec *ReportSpec, fields []string) {
	var vals []string
	for _, fldMnemonic := range fields {
		colspec := rspec.getColumn(fldMnemonic)
		val, _ := colspec.getValue(objIter)
		if val.IsValid() {
			vals = append(vals, val.String())
		} else {
			vals = append(vals, "")
		}
	}
	fmt.Printf("%s\n", strings.Join(vals, t.delim))
}

func getTSVCmd() *cobra.Command {
	use := "tsv"
	parentCmd := &cobra.Command{
		Use:   use,
		Short: "Convert data into TSV format",
		Long: `
Convert data to TSV format.

Refer to NCBI's [command line start](https://www.ncbi.nlm.nih.gov/datasets/docs/command-line-start) documentation for information about getting started with the command-line tools.`,
		Args: cobra.MaximumNArgs(0),
	}

	run := func(cmd *cobra.Command, args []string) (err error) {
		rpt := GetReport(cmd.Use)
		obj := rpt.newObject()
		writer := newDelimWriter("\t")
		err = emitTable(writer, obj, rpt, tableFields)
		return
	}
	for _, rpt := range AllReports {
		parentCmd.AddCommand(createFormatCmd(use, rpt, " into TSV format", run))
	}
	parentCmd.PersistentFlags().BoolVarP(&elideHeader, "elide-header", "", false, "Do not output header")
	return parentCmd
}
