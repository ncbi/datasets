package dataformat

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var (
	format string
)

type IStringWriter interface {
	Println(s string)
}

type StringWriter struct {
	IStringWriter
}

func (sw* StringWriter) Println(s string) {
	fmt.Println(s)
}

type DelimTableWriter struct {
	TableWriter
	Out IStringWriter
	Delim string
}

func newDelimWriter(delim string) *DelimTableWriter {
	sw := StringWriter{}
	writer := DelimTableWriter{
		Delim: delim,
		Out: &sw,
	}
	return &writer
}

func (t *DelimTableWriter) EmitTableHeader(rspec *ReportSpec, fields []string) {
	headerFields := make([]string, 0)
	for _, f := range fields {
		if ! rspec.hasColumn(f) {
			// How do we report back to the user there was an error?
			// should DelimTableWriter have an error collection?
			fmt.Println("Unrecognized mnemonic: ", f)
			continue
		}
		colspec := rspec.getColumn(f)
		headerFields = append(headerFields, colspec.Name)
	}
	t.Out.Println(strings.Join(headerFields, t.Delim))
}

func (t *DelimTableWriter) EmitTableRow(objIter *ObjIter, rspec *ReportSpec, fields []string) {
	var vals []string

	for _, f := range fields {
		if ! rspec.hasColumn(f) {
			// How do we report back to the user there was an error?
			// should DelimTableWriter have an error collection?
			continue
		}
		colspec := rspec.getColumn(f)
		val, _ := colspec.getValue(objIter)
		if val.IsValid() {
			vals = append(vals, val.String())
		} else {
			vals = append(vals, "")
		}
	}
	t.Out.Println(strings.Join(vals, t.Delim))
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
