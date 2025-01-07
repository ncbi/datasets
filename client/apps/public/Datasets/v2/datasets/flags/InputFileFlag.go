package flags

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const (
	AsIntegerTrue  bool = true
	AsIntegerFalse bool = false
)

const (
	InputFileListTypeVirusAcc   string = "nucleotide accessions"
	InputFileListTypeGenomeAcc  string = "NCBI Assembly or BioProject accessions"
	InputFileListTypeGeneAcc    string = "NCBI Gene Accessions"
	InputFileListTypeGeneId     string = "NCBI Gene IDs"
	InputFileListTypeGeneSymbol string = "NCBI Gene Symbols"
	InputFileListTypeTaxon      string = "NCBI taxonomy identifiers (taxid or current scientific name)"
)

type InputFileFlag struct {
	InputIDArgs []string
	AsInt32List []int32
	inputType   string
	inputFile   *string
	asInt       bool
	limit       int
}

// Option defines a type for functional options.
type Option func(*InputFileFlag)

func NewInputFileFlag(inputType string, asInt bool, options ...Option) *InputFileFlag {
	iff := &InputFileFlag{
		inputType: inputType,
		inputFile: new(string),
		asInt:     asInt,
		limit:     0,
	}
	for _, opt := range options {
		opt(iff)
	}

	return iff
}

func WithLimit(limit int) Option {
	return func(iff *InputFileFlag) {
		iff.limit = limit
	}
}

func (iff *InputFileFlag) RegisterFlags(flags *pflag.FlagSet) {
	if iff.limit > 0 {
		flags.StringVar(iff.inputFile, "inputfile", *iff.inputFile, fmt.Sprintf("Read a list of %s from a file to use as input (max: %d)", iff.inputType, iff.limit))
	} else {
		flags.StringVar(iff.inputFile, "inputfile", *iff.inputFile, fmt.Sprintf("Read a list of %s from a file to use as input", iff.inputType))
	}
}

func (iff *InputFileFlag) PreRunE(cmd *cobra.Command, args []string) (err error) {
	iff.InputIDArgs, err = getArgsFromListOrFile(args, *iff.inputFile, cmd, iff.inputType)
	if err != nil {
		return err
	}
	if iff.asInt {
		iff.AsInt32List, err = strToInt32ListErr(iff.InputIDArgs)
	}
	if iff.limit > 0 && len(iff.InputIDArgs) > iff.limit {
		return fmt.Errorf("too many %s provided, limit is %d", iff.inputType, iff.limit)
	}
	return err
}

func (iff *InputFileFlag) AsStringList() []string {
	return iff.InputIDArgs
}

// Helper functions
func strToInt32ListErr(strs []string) (geneInts []int32, err error) {
	hasError := false
	lastBadInput := ""
	for _, idFullStr := range strs {
		for _, idStr := range strings.Split(idFullStr, ",") {
			geneInt64, e := strconv.ParseInt(idStr, 10, 32)
			geneInt32 := int32(geneInt64)
			if e != nil {
				hasError = true
				lastBadInput = idStr
			} else {
				geneInts = append(geneInts, geneInt32)
			}
		}
	}
	if hasError {
		err = fmt.Errorf("unable to parse input value as an integer: '%s'", lastBadInput)
	}
	return
}

func readLines(fp io.Reader) []string {
	var lines []string
	reader := bufio.NewReader(fp)
	scanner := bufio.NewScanner(reader)
	scanner.Buffer(make([]byte, 0, 1024), 64*1024) //
	for scanner.Scan() {
		var line = strings.TrimSpace(scanner.Text())
		if len(line) > 0 {
			lines = append(lines, line)
		}
	}
	return lines
}

func getArgsFromListOrFile(args []string, argInputFile string, cmd *cobra.Command, inputType string) (idArgs []string, err error) {
	var errorMsg bytes.Buffer

	if len(args) != 0 {
		if argInputFile != "" {
			err = fmt.Errorf("Accepts either argument or file, not both")
			return
		}
		if len(args) == 1 {
			idArgs = strings.Split(args[0], ",")
		} else {
			idArgs = args
		}
		return
	}

	if argInputFile == "-" {
		idArgs = readLines(os.Stdin)
	} else if argInputFile != "" {
		fp, fileErr := os.Open(argInputFile)
		if fileErr != nil {
			err = fmt.Errorf("'%s' opening input file: '%s'", fileErr.Error(), argInputFile)
			return
		}
		defer fp.Close()
		idArgs = readLines(fp)
		// Check if any geneIDs read
		if len(idArgs) == 0 {
			fmt.Fprintf(
				&errorMsg,
				"No identifiers read from file: '%s'\n       File should have 1 identifier per row and no spaces or quotes",
				argInputFile,
			)
		}
	} else {
		errorFunc := ExpectOnePositionalArgument(inputType)
		err = errorFunc(cmd, args)
		return
	}

	if errorMsg.Len() > 0 {
		err = errors.New(errorMsg.String())
	}
	return
}
