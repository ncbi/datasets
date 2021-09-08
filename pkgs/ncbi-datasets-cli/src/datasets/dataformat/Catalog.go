package dataformat

import (
	"archive/zip"
	"bufio"
	"fmt"
	"io"
	"os"

	"google.golang.org/protobuf/encoding/protojson"
	pb_datasets "ncbi/datasets/v1"
)

const DATASET_CATALOG = "ncbi_dataset/data/dataset_catalog.json"

func scanEOF(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF {
		token = data
		advance = len(token)
		return
	}
	advance = 0
	for true {
		a, t, e := bufio.ScanLines(data[advance:], atEOF)
		if advance == 0 {
			break
		}
		if e != nil {
			err = e
			return
		}
		advance += a
		token = append(token[:], t[:a]...)
	}
	return
}

func emitCatalog() (err error) {
	catalog, err := getCatalog()
	pbOptsMarshal := protojson.MarshalOptions{
		Multiline:      true,
		Indent:         "  ",
		UseProtoNames:  false,
		UseEnumNumbers: false,
	}
	fmt.Print(pbOptsMarshal.Format(catalog))
	return
}

func getCatalog() (catalog *pb_datasets.Catalog, err error) {
	var file io.ReadCloser
	if len(packageFile) > 0 {
		zipreader, e := zip.OpenReader(packageFile)
		if e != nil {
			err = e
			return
		}
		defer zipreader.Close()
		fileptr, e := findSingleFileMatch(&zipreader.File, DATASET_CATALOG, DATASET_CATALOG)
		if e != nil {
			err = e
			return
		}
		if debug {
			fmt.Fprintf(os.Stderr, "matched file: %s\n", fileptr.Name)
		}
		file, err = fileptr.Open()
	} else {
		file, err = os.Open(inputFile)
	}
	if file != nil {
		defer file.Close()
	}
	if err != nil {
		return
	}

	const BUFFER_SIZE = 4e7

	scanner := bufio.NewScanner(file)
	scanner.Split(scanEOF)
	catalog = &pb_datasets.Catalog{}
	buf := make([]byte, 0, BUFFER_SIZE)
	scanner.Buffer(buf, BUFFER_SIZE)
	scanner.Scan()
	pbOptsUnmarshal := protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}
	if err = pbOptsUnmarshal.Unmarshal(scanner.Bytes(), catalog); err != nil {
		return
	}
	err = scanner.Err()
	return
}
