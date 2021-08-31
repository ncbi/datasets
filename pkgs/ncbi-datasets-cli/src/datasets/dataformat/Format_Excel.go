package dataformat

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	xlsx "github.com/tealeg/xlsx/v3"
)

var (
	outputFile string
)

type ExcelWriter struct {
	TableWriter
	worksheet *xlsx.Sheet
}

func newExcelWriter(workbook *xlsx.File, sheetname string) *ExcelWriter {
	sheet, err := workbook.AddSheet(sheetname)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil
	}
	return &ExcelWriter{
		worksheet: sheet,
	}
}

func (xlw *ExcelWriter) close() {
	if xlw.worksheet != nil {
		xlw.worksheet.Close()
	}
}

func (xlw *ExcelWriter) EmitTableHeader(rspec *ReportSpec, fields []string) {
	row := xlw.worksheet.AddRow()
	for _, f := range fields {
		colSpec := rspec.getColumn(f)
		cell := row.AddCell()
		cell.SetString(colSpec.Name)
	}
}

func (xlw *ExcelWriter) EmitTableRow(objIter *ObjIter, rspec *ReportSpec, fields []string) {
	row := xlw.worksheet.AddRow()
	for _, fldMnemonic := range fields {
		colSpec := rspec.getColumn(fldMnemonic)
		val, valType := colSpec.getValue(objIter)
		cell := row.AddCell()
		if !val.IsValid() {
			continue
		}
		switch valType {
		case "int":
			cell.SetInt64(val.Int())
			cell.SetFormat("0")
			//cell.SetFormat("###,###")
		case "uint":
			cell.SetInt64(int64(val.Uint()))
			cell.SetFormat("0")
		case "float":
			cell.SetFloat(val.Float())
		default:
			cell.SetString(val.String())
		}
	}
}

func createExcelWriter(sheetname string) (*xlsx.File, *ExcelWriter) {
	workbook := xlsx.NewFile()
	writer := newExcelWriter(workbook, sheetname)
	return workbook, writer
}

func validateFilename(filename string, ext string) bool {
	return filepath.Ext(filename) == ext
}

func getExcelCmd() *cobra.Command {
	use := "excel"
	parentCmd := &cobra.Command{
		Use:   use,
		Short: "Convert data into an Excel workbook",
		Long: `
Convert data into an Excel workbook.

Refer to NCBI's [command line start](https://www.ncbi.nlm.nih.gov/datasets/docs/command-line-start) documentation for information about getting started with the command-line tools.`,
		Args: cobra.MaximumNArgs(0),
	}

	run := func(cmd *cobra.Command, args []string) (err error) {
		rpt := GetReport(cmd.Use)
		obj := rpt.newObject()
		workbook, writer := createExcelWriter(rpt.Name)
		if workbook == nil {
			return
		}
		defer writer.close()
		err = emitTable(writer, obj, rpt, tableFields)
		if err != nil {
			return
		}
		if !validateFilename(outputFile, ".xlsx") {
			fmt.Fprintln(os.Stderr, "Workbook must have .xlsx extension")
			outputFile += ".xlsx"
		}
		workbook.Save(outputFile)
		fmt.Println("saved Excel workbook", outputFile)
		return
	}
	for _, rpt := range AllReports {
		parentCmd.AddCommand(createFormatCmd(use, rpt, " into an Excel workbook", run))
	}
	parentCmd.PersistentFlags().StringVarP(&outputFile, "outputfile", "o", "", "Excel workbook file")
	parentCmd.MarkPersistentFlagRequired("outputfile")
	return parentCmd
}
