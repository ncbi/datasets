package flags

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var excludePackageContentsTemplate = "The flag '--%s' has been disabled in datasets v14+. Data package file contents can be specified using the '--%s' flag."
var baseMessageTemplate = "The flag '--%s' has been disabled in datasets v14+ and replaced by the '--%s' flag."

var genomePackageContents = `
The default genome data package includes:
  * genomic.fna             (genome)
  * data_report.jsonl
  * dataset_catalog.json`

var genePackageContents = `
The default gene data package includes:
  * rna.fna                 (rna)
  * protein.faa             (protein)
  * data_report.jsonl
  * dataset_catalog.json`

var virusGenomePackageContents = `
The default virus data package includes:
  * genomic.fna             (genome)
  * data_report.jsonl
  * dataset_catalog.json`

var virusProteinPackageContents = `
The default protein virus data package includes:
  * data_report.jsonl
  * dataset_catalog.json`

var deprecatedFlags = map[string]map[string]string{
	"datasets.download.genome": {
		"message-template":    excludePackageContentsTemplate,
		"message-usage":       genomePackageContents,
		"exclude-genomic-cds": "include",
		"exclude-gff3":        "include",
		"exclude-protein":     "include",
		"exclude-rna":         "include",
		"exclude-seq":         "include",
	},
	"datasets.download.gene": {
		"message-template": excludePackageContentsTemplate,
		"message-usage":    genePackageContents,
		"exclude-gene":     "include",
		"exclude-rna":      "include",
		"exclude-protein":  "include",
	},
	"datasets.download.virus.protein": {
		"message-template": excludePackageContentsTemplate,
		"message-usage":    virusProteinPackageContents,
		"exclude-cds":      "include",
		"exclude-protein":  "include",
	},
	"datasets.download.virus.genome": {
		"message-template": excludePackageContentsTemplate,
		"message-usage":    virusGenomePackageContents,
		"exclude-seq":      "include",
		"exclude-cds":      "include",
		"exclude-protein":  "include",
	},
	"datasets": {
		"message-template": baseMessageTemplate,
		"message-usage":    "",
		"released-since":   "released-after",
		"updated-since":    "updated-after",
	},
}

var examples = map[string]map[string]string{
	"datasets": {
		"released-after":    "10/25/2022",
		"released-before":   "10/25/2022",
		"updated-after":     "10/25/2022",
		"api-key":           "ABCDE12345",
		"assembly-level":    "scaffold",
		"assembly-source":   "genbank",
		"assembly-version":  "latest",
		"chromosomes":       "1,2,22",
		"directory":         "my_download",
		"fasta-filter":      "NM_001077490.3",
		"fasta-filter-file": "accession.txt",
		"filename":          "my_download.zip",
		"geo-location":      "Canada",
		"host":              "dog",
		"include":           "rna,protein",
		"include-flanks-bp": "200",
		"inputfile":         "list.txt",
		"limit":             "10",
		"lineage":           "BA.1",
		"list":              "files_rehydrate.txt",
		"match":             "genomic",
		"max-workers":       "20",
		"ortholog":          "canidae",
		"search":            "ST540",
		"taxon":             "ferret",
		"taxon-filter":      "Acinetobacter lwoffii",
	},
	"datasets.summary.genome": {
		"report": "sequence",
	},
	"datasets.summary.gene": {
		"report": "product",
	},
	"datasets.summary.virus": {
		"report": "annotation",
	},
}

var missingPositionalArgumentTemplate = `Missing %s as the next positional argument.

Sample Commands
%s`

var tooManyPositionalArgumentsMsg = `Only one argument for %s is allowed

Sample Commands
%s`

func ExpectOnePositionalArgument(missingDesc string) cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) == 1 {
			return nil
		}
		pos_msg := tooManyPositionalArgumentsMsg
		if len(args) == 0 {
			pos_msg = missingPositionalArgumentTemplate
		}
		if cmd != nil {
			return fmt.Errorf(pos_msg, missingDesc, cmd.Example)
		} else {
			return fmt.Errorf(pos_msg, missingDesc, "<Examples>")
		}
	}
}

func ExpectAtLeastOnePositionalArgument(missingDesc string) cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf(missingPositionalArgumentTemplate, missingDesc, cmd.Example)
		}
		return nil
	}
}

var dateMissingArgumentTemplate = "--%s must have a date argument formatted as MM/DD/YYYY"
var stringMissingArgumentTemplate = "--%s must have a string argument"
var missingArgumentMessages = map[string]map[string]string{
	"datasets.": {
		"released-after":    dateMissingArgumentTemplate,
		"released-before":   dateMissingArgumentTemplate,
		"updated-after":     dateMissingArgumentTemplate,
		"api-key":           stringMissingArgumentTemplate,
		"assembly-level":    stringMissingArgumentTemplate,
		"assembly-source":   stringMissingArgumentTemplate,
		"assembly-version":  stringMissingArgumentTemplate,
		"chromosomes":       stringMissingArgumentTemplate,
		"directory":         stringMissingArgumentTemplate,
		"fasta-filter":      stringMissingArgumentTemplate,
		"fasta-filter-file": stringMissingArgumentTemplate,
		"filename":          stringMissingArgumentTemplate,
		"geo-location":      stringMissingArgumentTemplate,
		"host":              stringMissingArgumentTemplate,
		"include":           stringMissingArgumentTemplate,
		"include-flanks-bp": stringMissingArgumentTemplate,
		"inputfile":         stringMissingArgumentTemplate,
		"limit":             stringMissingArgumentTemplate,
		"lineage":           stringMissingArgumentTemplate,
		"list":              stringMissingArgumentTemplate,
		"match":             stringMissingArgumentTemplate,
		"max-workers":       stringMissingArgumentTemplate,
		"ortholog":          stringMissingArgumentTemplate,
		"refseq":            stringMissingArgumentTemplate,
		"report":            stringMissingArgumentTemplate,
		"search":            stringMissingArgumentTemplate,
		"taxon":             stringMissingArgumentTemplate,
		"taxon-filter":      stringMissingArgumentTemplate,
	},
}

func signatureFor(cmd *cobra.Command) string {
	signature := []string{cmd.Use}
	cmd.VisitParents(
		func(cmd *cobra.Command) {
			signature = append([]string{cmd.Use}, signature...)
		},
	)
	return strings.Join(signature, ".")
}

func exampleFor(cmdSig string, argument string) (string, bool) {
	for k, v := range examples {
		if !strings.HasPrefix(cmdSig, k) {
			continue
		}
		if example, ok := v[argument]; ok {
			return example, true
		}
	}
	return "", false
}

func HandleUnknownFlag(cmd *cobra.Command, inErr error, prefixLen int) error {
	cmdSig := signatureFor(cmd)
	unknownFlag := inErr.Error()[prefixLen+3:]

	for k, v := range deprecatedFlags {
		if !strings.HasPrefix(cmdSig, k) {
			continue
		}
		if newFlag, ok := v[unknownFlag]; ok {
			if example, ok2 := exampleFor(cmdSig, newFlag); ok2 {
				return fmt.Errorf(v["message-template"]+"  Example: '--%s %s'\n"+v["message-usage"], unknownFlag, newFlag, newFlag, example)
			}
			return fmt.Errorf(v["message-template"]+"\n"+v["message-usage"], unknownFlag, newFlag)
		}
	}
	return nil
}

func HandleMissingArgumentValue(cmd *cobra.Command, inErr error, prefixLen int) error {
	cmdSig := signatureFor(cmd)
	missingArgument := inErr.Error()[prefixLen+3:]

	for k, v := range missingArgumentMessages {
		if !strings.HasPrefix(cmdSig, k) {
			continue
		}
		if newMessageF, ok := v[missingArgument]; ok {
			if example, ok2 := exampleFor(cmdSig, missingArgument); ok2 {
				return fmt.Errorf(newMessageF+".  Example: '--%s %s'", missingArgument, missingArgument, example)
			}
			return fmt.Errorf(newMessageF, missingArgument)
		}
	}
	return nil
}

type handlerInterface func(cmd *cobra.Command, inErr error, prefixLen int) error

var handlers = map[string]handlerInterface{
	"unknown flag:":           HandleUnknownFlag,
	"flag needs an argument:": HandleMissingArgumentValue,
}

func ErrorFlagHandler(cmd *cobra.Command, inErr error) error {
	for prefix, handler := range handlers {
		if strings.HasPrefix(inErr.Error(), prefix) {
			if err := handler(cmd, inErr, len(prefix)); err != nil {
				return err
			}
		}
	}
	return inErr
}
