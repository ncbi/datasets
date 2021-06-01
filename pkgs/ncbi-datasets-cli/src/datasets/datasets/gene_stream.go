package datasets

import (
	"bufio"
	"encoding/json"
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

func bytesToGeneConvert(b []byte) (*openapi.V1alpha1GeneMatch, error) {
	geneMatchObj := new(openapi.V1alpha1GeneMatch)
	if err := json.Unmarshal(b, geneMatchObj); err != nil {
		return nil, fmt.Errorf("Failure to parse JSON: %s. %w", string(b), err)
	}
	if argDebug {
		//fmt.Println("Parsed object: ", geneMatchObj)
	}
	return geneMatchObj, nil
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

func streamGeneMatch(req *openapi.V1alpha1GeneDatasetRequest, sp streamProcessor) error {
	cli, err := createOAClient()
	if err != nil {
		return err
	}
	_, resp, err := cli.GeneApi.GeneMetadataStreamByPost(nil, *req)
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
	GeneIds []int64
}

func printInfo(reason string, identifiers []string, message string) {
	fmt.Fprintf(os.Stderr, message)
	if argDebug {
		err := fmt.Errorf("Detailed Message: For input: %s : %s", strings.Join(identifiers, ", "), reason)
		fmt.Fprintf(os.Stderr, err.Error())
	}
}

func printMessage(geneMatchObj *openapi.V1alpha1GeneMatch) error {
	geneIdStr := geneMatchObj.Gene.GeneId
	if len(geneMatchObj.Errors) > 0 {
		printInfo(geneMatchObj.Errors[0].Reason, geneMatchObj.Errors[0].InvalidIdentifiers, geneMatchObj.Errors[0].Message)
	} else if len(geneMatchObj.Warnings) > 0 {
		printInfo(geneMatchObj.Warnings[0].Reason, geneMatchObj.Query, geneMatchObj.Warnings[0].Message)
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
		if len(geneMatchObj.Gene.GeneId) == 0 {
			return nil
		}
		if geneId, e := strconv.ParseInt(geneMatchObj.Gene.GeneId, 10, 64); e == nil {
			g.GeneIds = append(g.GeneIds, geneId)
		} else {
			fmt.Println("Failure to parse integer: ", e)
			return e
		}
	}
	return nil
}
