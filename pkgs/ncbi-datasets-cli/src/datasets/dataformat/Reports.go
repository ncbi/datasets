package dataformat

import (
	"bytes"
	"errors"
	"fmt"
	"sort"
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	pb_options "ncbi/datasets/options"
	pb_reports "ncbi/datasets/v1/reports"
)

type FieldDesc struct {
	name            string
	fieldNumber     protoreflect.FieldNumber
	fieldKind       protoreflect.Kind
	repeatedMessage bool
}

type FieldDescList []FieldDesc

type ColSpec struct {
	Mnemonic  string
	Name      string
	fieldDesc FieldDescList
	isSimple  *bool
	dotname   string
}

type ListIter struct {
	currIdx int
	len     int
}

type ObjIter struct {
	obj     protoreflect.Message
	iters   []string
	listIdx map[string]*ListIter
	// cache non-repeated field values (memoization)
	simpleFields map[string]SimpleVal
}

func MakeObjIter(obj protoreflect.Message) *ObjIter {
	return &ObjIter{
		obj:          obj,
		iters:        []string{},
		listIdx:      map[string]*ListIter{},
		simpleFields: map[string]SimpleVal{},
	}
}

type SimpleVal struct {
	result     protoreflect.Value
	resultType string
}

type CommandStringsMap map[string][]string

type ReportSpec struct {
	id                 string
	cmd                string
	messageType        proto.Message
	Name               string
	mappedMessages     map[string]bool
	mapping            map[string]*ColSpec
	FileName           string
	defaultPackagePath string
	regexpMatch        string
	exampleCommands    CommandStringsMap
}

var (
	// The order here impacts the order in the doc-tree
	AllReports = []*ReportSpec{
		&ReportSpec{
			cmd:                "genome",
			messageType:        (*pb_reports.AssemblyDataReport)(nil),
			Name:               "Genome Assembly Data Report",
			FileName:           "assembly_report.html",
			defaultPackagePath: "ncbi_dataset/data/assembly_data_report.jsonl",
			regexpMatch:        "^ncbi_dataset/data/assembly_data_report.jsonl$",
			exampleCommands: CommandStringsMap{
				"tsv": []string{
					"dataformat tsv genome --inputfile human/ncbi_dataset/data/assembly_data_report.jsonl",
					"dataformat tsv genome --package human.zip",
				},
				"excel": []string{
					"dataformat excel genome --inputfile human/ncbi_dataset/data/assembly_data_report.jsonl --outputfile human-genomes.xlsx",
					"dataformat excel genome --package human.zip --outputfile human-genomes.xlsx",
				},
			},
		},
		&ReportSpec{
			cmd:         "genome-seq",
			messageType: (*pb_reports.SequenceInfo)(nil),
			Name:        "Genome Assembly Sequence Report",
			FileName:    "assembly_sequence_info_report.html",
			regexpMatch: "^ncbi_dataset/data/.+/sequence_report.jsonl$",
			exampleCommands: CommandStringsMap{
				"tsv": []string{
					"dataformat tsv genome-seq --inputfile human_package/ncbi_dataset/data/GCF_000001405.39/sequence_report.jsonl",
					"dataformat tsv genome-seq --package human.zip --inputfile GCF_000001405.39/sequence_report.jsonl",
				},
				"excel": []string{
					"dataformat excel genome-seq --inputfile human_package/ncbi_dataset/data/GCF_000001405.39/sequence_report.jsonl --outputfile human-ref-genome-seq.xlsx",
					"dataformat excel genome-seq --package human.zip --inputfile GCF_000001405.39/sequence_report.jsonl --outputfile human-ref-genome-seq.xlsx",
				},
			},
		},
		&ReportSpec{
			cmd:         "gene",
			messageType: (*pb_reports.GeneDescriptor)(nil),
			Name:        "Gene Report",
			FileName:    "gene_report.html",
			regexpMatch: "^ncbi_dataset/data/data_report.jsonl$",
			exampleCommands: CommandStringsMap{
				"tsv": []string{
					"dataformat tsv gene --inputfile gene_package/ncbi_dataset/data/data_report.jsonl",
					"dataformat tsv gene --package genes.zip",
				},
				"excel": []string{
					"dataformat excel gene --inputfile gene_package/ncbi_dataset/data/data_report.jsonl --outputfile genes.xlsx",
					"dataformat excel gene --package genes.zip --outputfile genes.xlsx",
				},
			},
		},
		&ReportSpec{
			cmd:         "virus-genome",
			messageType: (*pb_reports.VirusAssembly)(nil),
			Name:        "Virus Data Report",
			FileName:    "virus_report.html",
			regexpMatch: "^ncbi_dataset/data/data_report.jsonl$",
			exampleCommands: CommandStringsMap{
				"tsv": []string{
					"dataformat tsv virus-genome --inputfile sars2_package/ncbi_dataset/data/data_report.jsonl",
					"dataformat tsv virus-genome --package virus-sars2-refseq.zip",
				},
				"excel": []string{
					"dataformat excel virus-genome --inputfile sars2_package/ncbi_dataset/data/data_report.jsonl --outputfile sars2.xlsx",
					"dataformat excel virus-genome --package virus-sars2-refseq.zip --outputfile sars2.xlsx",
				},
			},
		},
		&ReportSpec{
			cmd:         "microbigge",
			messageType: (*pb_reports.MicroBiggeReport)(nil),
			Name:        "MicroBIGG-E Data Report",
			FileName:    "microbigge_report.html",
			regexpMatch: "^ncbi_dataset/data/data_report.jsonl$",
			exampleCommands: CommandStringsMap{
				"tsv": []string{
					"dataformat tsv microbigge --inputfile pathogen_package/ncbi_dataset/data/data_report.jsonl",
					"dataformat tsv microbigge --package virus-sars2-refseq.zip",
				},
				"excel": []string{
					"dataformat excel microbigge --inputfile pathogen_package/ncbi_dataset/data/data_report.jsonl --outputfile pathogen.xlsx",
					"dataformat excel microbigge --package van.zip --outputfile van.xlsx",
				},
			},
		},
		&ReportSpec{
			cmd:         "prok-gene",
			messageType: (*pb_reports.ProkaryoteGene)(nil),
			Name:        "Prokaryote Gene Report",
			FileName:    "prokaryote_gene_report.html",
			regexpMatch: "^ncbi_dataset/data/data_report.jsonl$",
			exampleCommands: CommandStringsMap{
				"tsv": []string{
					"dataformat tsv prok-gene --inputfile prokaryote_package/ncbi_dataset/data/data_report.jsonl",
					"dataformat tsv prok-gene --package prokaryote-package.zip",
				},
				"excel": []string{
					"dataformat excel prok-gene --inputfile prokaryote_package/ncbi_dataset/data/data_report.jsonl --outputfile prokaryote.xlsx",
					"dataformat excel prok-gene --package prokaryote-package.zip --outputfile prokaryote.xlsx",
				},
			},
		},
		&ReportSpec{
			cmd:         "prok-gene-location",
			messageType: (*pb_reports.ProkaryoteGeneLocation)(nil),
			Name:        "Prokaryote Gene Location Report",
			FileName:    "prokaryote_gene_location_report.html",
			regexpMatch: "^ncbi_dataset/data/annotation_report.jsonl$",
			exampleCommands: CommandStringsMap{
				"tsv": []string{
					"dataformat tsv prok-gene-location --inputfile prokaryote_package/ncbi_dataset/data/annotation_report.jsonl",
					"dataformat tsv prok-gene-location --package prokaryote-package.zip",
				},
				"excel": []string{
					"dataformat excel prok-gene-location --inputfile prokaryote_package/ncbi_dataset/data/annotation_report.jsonl --outputfile prokaryote.xlsx",
					"dataformat excel prok-gene-location --package prokaryote-package.zip --outputfile prokaryote.xlsx",
				},
			},
		},
	}
	allReportsByCmd = map[string]*ReportSpec{}
)

func getScalarListValue(listval protoreflect.List) (result string) {
	const scalarDelim = ","
	strs := make([]string, listval.Len())
	for i := 0; i < listval.Len(); i++ {
		strs[i] = fmt.Sprintf("%v", listval.Get(i).Interface())
	}
	result = strings.Join(strs, scalarDelim)
	return
}

func (c *ColSpec) isSimpleField() bool {
	if c.isSimple != nil {
		return *c.isSimple
	}
	retval := false
	for _, fd := range c.fieldDesc {
		if fd.repeatedMessage {
			c.isSimple = &retval
			return retval
		}
	}
	retval = true
	c.isSimple = &retval
	return retval
}

func (c *ColSpec) getDotname() string {
	if c.dotname != "" {
		return c.dotname
	}
	dotname := ""
	for _, fd := range c.fieldDesc {
		dotname += "." + string(fd.name)
	}
	c.dotname = dotname
	return c.dotname
}

func (it *ObjIter) Next() bool {
	for true {
		if it.iters == nil || len(it.iters) == 0 {
			return false
		}
		idx := len(it.iters) - 1
		iterName := it.iters[idx]
		iterPtr := it.listIdx[iterName]
		nextIdx := iterPtr.currIdx + 1
		if nextIdx < (*iterPtr).len {
			(*iterPtr).currIdx = nextIdx
			return true
		}
		// remove exhausted iterator from the top of the stack
		(*it).iters = it.iters[:idx]
		delete(it.listIdx, iterName)
	}
	return false
}

func (it *ObjIter) getIndex(dotname string, len int) (idx int, ok bool) {
	if len < 1 {
		ok = false
		return
	}
	if idxVal, listOk := it.listIdx[dotname]; listOk {
		if idxVal.currIdx >= len {
			panic("iterator extended beyond end of list")
		}
		idx = idxVal.currIdx
		ok = true
		return
	}
	idx = 0
	it.listIdx[dotname] = &ListIter{
		currIdx: idx,
		len:     len,
	}
	it.iters = append(it.iters, dotname)
	ok = true
	return
}

func (objIt *ObjIter) getSimpleFieldValue(dotname string) *SimpleVal {
	if simpleVal, ok := objIt.simpleFields[dotname]; ok {
		return &simpleVal
	}
	return nil
}

func kindToType(kind protoreflect.Kind) (resultType string) {
	switch kind {
	case protoreflect.Int32Kind, protoreflect.Sint32Kind,
		protoreflect.Int64Kind, protoreflect.Sint64Kind:
		resultType = "int"
	case protoreflect.Uint32Kind, protoreflect.Uint64Kind:
		resultType = "uint"
	case protoreflect.Sfixed32Kind, protoreflect.Fixed32Kind, protoreflect.FloatKind,
		protoreflect.Sfixed64Kind, protoreflect.Fixed64Kind, protoreflect.DoubleKind:
		resultType = "float"
	default:
		resultType = "string"
	}
	return
}

func (c *ColSpec) getValue(objIt *ObjIter) (result protoreflect.Value, resultType string) {
	msg := objIt.obj
	fullDotname := c.getDotname()
	if simpleVal := objIt.getSimpleFieldValue(fullDotname); simpleVal != nil {
		result = simpleVal.result
		resultType = simpleVal.resultType
		return
	}
	fieldDepth := len(c.fieldDesc)
	dotname := ""
	for i, fDesc := range c.fieldDesc {
		fldDesc := msg.Descriptor().Fields().ByNumber(fDesc.fieldNumber)
		if !msg.Has(fldDesc) {
			return
		}
		dotname += "." + fDesc.name
		v := msg.Get(fldDesc)
		if fDesc.fieldKind == protoreflect.EnumKind {
			enumVal := v.Enum()
			enumName := fldDesc.Enum().Values().ByNumber(enumVal).Name()
			enumNameVal := protoreflect.ValueOfString(string(enumName))
			result = enumNameVal
			resultType = kindToType(fDesc.fieldKind)
			return
		}
		lastField := (i == fieldDepth-1)
		if lastField {
			if !fldDesc.IsList() {
				result = v
				resultType = kindToType(fDesc.fieldKind)
			} else if fDesc.fieldKind != protoreflect.MessageKind {
				result = protoreflect.ValueOfString(getScalarListValue(v.List()))
				resultType = kindToType(fDesc.fieldKind)
			}
			if c.isSimpleField() {
				objIt.simpleFields[fullDotname] = SimpleVal{
					result:     result,
					resultType: resultType,
				}
			}
			return
		} else {
			if fDesc.repeatedMessage {
				listval := v.List()
				if idx, ok := objIt.getIndex(dotname, listval.Len()); ok {
					msg = listval.Get(idx).Message()
				}
			} else {
				msg = v.Message()
			}
		}
	}
	return
}

func (rpt *ReportSpec) addColumn(col *ColSpec) {
	rpt.mapping[col.Mnemonic] = col
}

func (rpt *ReportSpec) hasColumn(field string) bool {
	_, ok := rpt.mapping[field]
	return ok
}

func (rpt *ReportSpec) hasColumns(fields []string) (val bool) {
	val, _ = rpt.fieldsAreValid(fields)
	return
}

func (rpt *ReportSpec) fieldsAreValid(fields []string) (valid bool, err error) {
	valid = true
	var errStr bytes.Buffer
	for _, f := range fields {
		if !rpt.hasColumn(f) {
			valid = false
			fmt.Fprintf(&errStr, "field [%s] not recognized", f)
		}
	}
	if errStr.Len() > 0 {
		err = errors.New(errStr.String())
	}
	return
}

func (rpt *ReportSpec) getColumn(field string) *ColSpec {
	return rpt.mapping[field]
}

func (rpt *ReportSpec) newObject() protoreflect.Message {
	return rpt.messageType.ProtoReflect().New()
}

func (rpt *ReportSpec) GetAllFields() (cols []*ColSpec) {
	fields := rpt.GetAllFieldMnemonics()
	cols = make([]*ColSpec, len(fields))
	for i, field := range fields {
		cols[i] = rpt.mapping[field]
	}
	return
}

func (rpt *ReportSpec) GetAllFieldMnemonics() []string {
	fields := make([]string, len(rpt.mapping))
	i := 0
	for _, col := range rpt.mapping {
		fields[i] = col.Mnemonic
		i++
	}
	sort.StringSlice(fields).Sort()
	return fields
}

func (rpt *ReportSpec) populateMappingForObj(
	obj protoreflect.Message,
	fieldDescs FieldDescList,
	prefixMnemonic string,
	prefixColName string,
) {
	objDesc := obj.Descriptor()
	msgName := prefixMnemonic
	if mapped, found := rpt.mappedMessages[msgName]; found && mapped {
		return
	}
	rpt.mappedMessages[msgName] = true
	flds := objDesc.Fields()
	for i := 0; i < flds.Len(); i++ {
		fld := flds.Get(i)
		isList := fld.IsList()
		isMsg := (fld.Kind() == protoreflect.MessageKind)
		isMsgList := (isList && isMsg)
		opts := fld.Options().ProtoReflect()
		newFieldDesc := FieldDesc{
			name:            string(fld.Name()),
			fieldNumber:     fld.Number(),
			fieldKind:       fld.Kind(),
			repeatedMessage: isMsgList,
		}
		newFieldDescs := append(fieldDescs, newFieldDesc)
		if opts == nil || !opts.IsValid() {
			if isMsg && !isList {
				// a message field containing a "partial" tabular description
				rpt.populateMappingForObj(
					obj.Get(fld).Message(),
					newFieldDescs,
					prefixMnemonic,
					prefixColName,
				)
			} else if debug {
				fmt.Println("field options not found for", objDesc.Name(), fld.Name(), "skipping")
			}
			continue
		}
		extTab := proto.GetExtension(opts.Interface(), pb_options.E_Tabular)
		if extTab == nil {
			if debug {
				fmt.Println("tabular field options not found for", objDesc.Name(), fld.Name(), "skipping")
			}
			continue
		}
		tbl_opts := extTab.(*pb_options.TabularOptions)
		if tbl_opts == nil {
			if debug {
				fmt.Println("TabularOptions not found for", objDesc.Name(), fld.Name(), "skipping")
			}
			continue
		}
		mnemonic := prefixMnemonic + tbl_opts.Mnemonic
		colName := prefixColName + tbl_opts.ColumnName
		if isMsg {
			// a message field containing a "partial" tabular description
			var newObj protoreflect.Message
			if isList {
				newObj = obj.Get(fld).List().NewElement().Message()
			} else {
				newObj = obj.Get(fld).Message()
			}
			fds := make(FieldDescList, len(newFieldDescs))
			copy(fds, newFieldDescs)
			rpt.populateMappingForObj(
				newObj,
				fds,
				mnemonic,
				colName,
			)
		} else {
			rpt.addColumn(&ColSpec{
				Mnemonic:  mnemonic,
				Name:      colName,
				fieldDesc: newFieldDescs,
			})
		}
	}
}

func (rpt *ReportSpec) populateMapping() {
	if rpt.mappedMessages == nil {
		rpt.mappedMessages = map[string]bool{}
	}
	if rpt.mapping == nil {
		rpt.mapping = map[string]*ColSpec{}
	}
	if rpt.messageType != nil {
		obj := rpt.newObject()
		rpt.populateMappingForObj(obj, FieldDescList{}, "", "")
	}
}

func GetReport(name string) *ReportSpec {
	return allReportsByCmd[name]
}

func init() {
	for _, rpt := range AllReports {
		allReportsByCmd[rpt.cmd] = rpt
		rpt.populateMapping()
	}
}
