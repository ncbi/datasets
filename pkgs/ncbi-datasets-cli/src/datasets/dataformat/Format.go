package dataformat

import (
	"archive/zip"
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	pb_datasets "ncbi/datasets/v1"

	"github.com/spf13/cobra"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var (
	elideHeader = false
)

type ITableWriter interface {
	EmitTableHeader(rspec *ReportSpec, fields []string)
	EmitTableRow(objIter *ObjIter, rspec *ReportSpec, fields []string)
}

type TableWriter struct {
}

func constructRegexpPattern(pattern string) string {
	return ".*" + regexp.QuoteMeta(pattern) + ".*"
}

func buildFilenameRegexp(pattern string) (*regexp.Regexp, error) {
	return regexp.Compile(constructRegexpPattern(pattern))
}

func findSingleFileMatch(files *[]*zip.File, fileRegexp string, ifile string) (fileptr *zip.File, err error) {
	matchedFiles := make([]*zip.File, 0)
	reportRegexp := regexp.MustCompile(fileRegexp)
	userRegexp, err := buildFilenameRegexp(ifile)
	if err != nil {
		return
	}
	for _, f := range *files {
		if !(reportRegexp.MatchString(f.Name) && userRegexp.MatchString(f.Name)) {
			continue
		}
		matchedFiles = append(matchedFiles, f)
	}
	if len(matchedFiles) == 0 {
		err = fmt.Errorf("no matching files found for [%s]\n", inputFile)
		return
	} else if len(matchedFiles) > 1 {
		var str bytes.Buffer
		fmt.Fprintf(&str, "multiple matching files found for [%s]\n", inputFile)
		for _, f := range matchedFiles {
			fmt.Fprintf(&str, "\t%s\n", f.Name)
		}
		err = errors.New(str.String())
		return
	}
	fileptr = matchedFiles[0]
	return
}

func contains(itemList []pb_datasets.Catalog_ApiVersion, search_val pb_datasets.Catalog_ApiVersion) bool {
	for _, val := range itemList {
		if val == search_val {
			return true
		}
	}
	return false
}

func emitTable(writer ITableWriter, obj protoreflect.Message, rspec *ReportSpec, fields []string) (err error) {
	if ok, e := rspec.fieldsAreValid(fields); !ok {
		err = e
		return
	}

	var file io.ReadCloser
	if len(packageFile) > 0 {
		zipreader, e := zip.OpenReader(packageFile)
		if e != nil {
			err = e
			return
		}
		defer zipreader.Close()
		fileptr, e := findSingleFileMatch(&zipreader.File, rspec.regexpMatch, inputFile)
		if e != nil {
			err = e
			return
		}
		if debug {
			fmt.Fprintf(os.Stderr, "matched file: %s\n", fileptr.Name)
		}

		catalog, e := getCatalog()
		if e != nil {
			err = e
			return
		}
		if !contains(supportedApiVersions, catalog.ApiVersion) {
			err = fmt.Errorf("The archive %s is version: %s and is not supported by this version (%s) of the dataformat tool", packageFile, catalog.ApiVersion.String(), AppVersion)
			if err != nil {
				return
			}
		}
		file, err = fileptr.Open()
	} else if len(inputFile) > 0 {
		file, err = os.Open(inputFile)
	} else {
		err = fmt.Errorf("--inputfile and/or --packagefile must be specified")
	}
	if file != nil {
		defer file.Close()
	}
	if err != nil {
		return
	}

	if !elideHeader {
		writer.EmitTableHeader(rspec, fields)
	}
	scanner := bufio.NewScanner(file)
	const BUFFER_SIZE = 4e6
	buf := make([]byte, 0, BUFFER_SIZE)
	scanner.Buffer(buf, BUFFER_SIZE)
	pb_opts := protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}
	for scanner.Scan() {
		line := scanner.Bytes()
		if err = pb_opts.Unmarshal(line, obj.Interface()); err != nil {
			return
		}
		objIter := MakeObjIter(obj)
		for ok := true; ok; ok = objIter.Next() {
			writer.EmitTableRow(objIter, rspec, fields)
		}
	}
	err = scanner.Err()
	return
}

func createFormatCmd(
	parentUse string,
	rpt *ReportSpec,
	shortDescSuffix string,
	runE func(cmd *cobra.Command, args []string) error,
) *cobra.Command {
	shortdesc := fmt.Sprint("Convert ", rpt.Name, shortDescSuffix)
	cmd := &cobra.Command{
		Use:   rpt.cmd,
		Short: shortdesc,
		Long: fmt.Sprintf(`
%s.

Refer to NCBI's [command line start](https://www.ncbi.nlm.nih.gov/datasets/docs/command-line-start) documentation for information about getting started with the command-line tools.`, shortdesc),
		Args:    cobra.MaximumNArgs(0),
		RunE:    runE,
		PreRunE: inputFilePreRunE,
	}
	if len(rpt.exampleCommands) > 0 {
		if examples, ok := rpt.exampleCommands[parentUse]; ok {
			cmd.Example = "  " + strings.Join(examples, "\n  ")
		}
	}

	addFileInputFlags(cmd)
	pflags := cmd.PersistentFlags()
	pflags.StringSliceVar(&tableFields, "fields", []string{}, "comma-separated list of fields")
	return cmd
}
