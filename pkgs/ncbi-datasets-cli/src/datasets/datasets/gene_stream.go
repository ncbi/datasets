package datasets

import (
	"bufio"
	"errors"
	"fmt"
	openapi "main/openapi_client"
	"os"
	"strconv"
	"strings"
)

type streamProcessor interface {
	insertRecordBegin()
	insertRecordEnd()
	insertRecordDelimiter()
	processor(b []byte) error
}

func bytesToGeneConvert(b []byte) (*openapi.V1GeneMatch, error) {
	geneMatchObj := new(openapi.NullableV1GeneMatch)
	if err := geneMatchObj.UnmarshalJSON(b); err != nil {
		return nil, fmt.Errorf("Failure to parse JSON: %s. %w", string(b), err)
	}
	if argDebug {
		fmt.Println("Parsed object: ", geneMatchObj.Get())
	}
	return geneMatchObj.Get(), nil
}

type JsonLinesStreamProcessor struct{}

func (j *JsonLinesStreamProcessor) insertRecordBegin()     {}
func (j *JsonLinesStreamProcessor) insertRecordEnd()       { fmt.Print("\n") }
func (j *JsonLinesStreamProcessor) insertRecordDelimiter() { fmt.Print("\n") }
func (j *JsonLinesStreamProcessor) processor(b []byte) error {
	if geneMatchObj, err := bytesToGeneConvert(b); err != nil {
		return err
	} else {
		printResultsNoNewline(geneMatchObj)
	}
	return nil
}

type JsonStreamProcessor struct {
	JsonLinesStreamProcessor
	wrapperName string
}

func (j *JsonStreamProcessor) insertRecordBegin()     { fmt.Printf("{\"%s\":[", j.wrapperName) }
func (j *JsonStreamProcessor) insertRecordEnd()       { fmt.Print("]}\n") }
func (j *JsonStreamProcessor) insertRecordDelimiter() { fmt.Print(",") }

func streamGeneMatch(req *openapi.V1GeneDatasetRequest, sp streamProcessor) error {
	cli, err := createOAClient()
	if err != nil {
		return err
	}
	_, resp, err := cli.GeneApi.GeneMetadataStreamByPost(nil, req).Execute()
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	scanner := bufio.NewScanner(resp.Body)
	const BUFFER_SIZE = 4e6
	buf := make([]byte, 0, BUFFER_SIZE)
	scanner.Buffer(buf, BUFFER_SIZE)
	firstLine := true

	sp.insertRecordBegin()
	for scanner.Scan() {
		if !firstLine {
			sp.insertRecordDelimiter()
		}
		firstLine = false
		if err := sp.processor(scanner.Bytes()); err != nil {
			return err
		}
	}
	sp.insertRecordEnd()
	return scanner.Err()
}

type GeneIdStreamProcessor struct {
	GeneIds []int32
}

func printInfo(reason string, identifiers []string, message string) {
	fmt.Fprintf(os.Stderr, message)
	if argDebug {
		err := fmt.Errorf("Detailed Message: For input: %s : %s", strings.Join(identifiers, ", "), reason)
		fmt.Fprintf(os.Stderr, err.Error())
	}
}

func printMessage(geneMatchObj *openapi.V1GeneMatch) error {
	geneIdStr := geneMatchObj.Gene.GetGeneId()
	if len(geneMatchObj.GetErrors()) > 0 {
		printInfo(geneMatchObj.GetErrors()[0].GetReason(), geneMatchObj.GetErrors()[0].GetInvalidIdentifiers(), geneMatchObj.GetErrors()[0].GetMessage())
	} else if len(geneMatchObj.GetWarnings()) > 0 {
		printInfo(geneMatchObj.GetWarnings()[0].GetReason(), geneMatchObj.GetQuery(), geneMatchObj.GetWarnings()[0].GetMessage())
	} else if len(geneIdStr) == 0 {
		return errors.New("Records is missing a gene-id")
	}
	return nil
}

func (g *GeneIdStreamProcessor) insertRecordBegin()     {}
func (g *GeneIdStreamProcessor) insertRecordEnd()       {}
func (g *GeneIdStreamProcessor) insertRecordDelimiter() {}
func (g *GeneIdStreamProcessor) processor(b []byte) error {
	if geneMatchObj, err := bytesToGeneConvert(b); err != nil {
		return err
	} else {
		if err := printMessage(geneMatchObj); err != nil {
			return err
		}
		if len(geneMatchObj.Gene.GetGeneId()) == 0 {
			return nil
		}
		if geneId, e := strconv.ParseInt(geneMatchObj.Gene.GetGeneId(), 10, 32); e == nil {
			g.GeneIds = append(g.GeneIds, int32(geneId))
		} else {
			fmt.Println("Failure to parse integer: ", e)
			return e
		}
	}
	return nil
}
