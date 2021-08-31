/*
 * NCBI Datasets API
 *
 * ### NCBI Datasets is a resource that lets you easily gather data from NCBI. The Datasets API is still in alpha, and we're updating it often to add new functionality, iron out bugs and enhance usability. For some larger downloads, you may want to download a [dehydrated bag](https://www.ncbi.nlm.nih.gov/datasets/docs/rehydrate/), and retrieve the individual data files at a later time. 
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datasets

import (
	"encoding/json"
	"fmt"
)

// V1Fasta the model 'V1Fasta'
type V1Fasta string

// List of v1Fasta
const (
	V1FASTA_UNSPECIFIED V1Fasta = "FASTA_UNSPECIFIED"
	V1FASTA_GENE V1Fasta = "FASTA_GENE"
	V1FASTA_RNA V1Fasta = "FASTA_RNA"
	V1FASTA_PROTEIN V1Fasta = "FASTA_PROTEIN"
	V1FASTA_GENE_FLANK V1Fasta = "FASTA_GENE_FLANK"
	V1FASTA_CDS V1Fasta = "FASTA_CDS"
	V1FASTA__5_P_UTR V1Fasta = "FASTA_5P_UTR"
	V1FASTA__3_P_UTR V1Fasta = "FASTA_3P_UTR"
)

var allowedV1FastaEnumValues = []V1Fasta{
	"FASTA_UNSPECIFIED",
	"FASTA_GENE",
	"FASTA_RNA",
	"FASTA_PROTEIN",
	"FASTA_GENE_FLANK",
	"FASTA_CDS",
	"FASTA_5P_UTR",
	"FASTA_3P_UTR",
}

func (v *V1Fasta) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := V1Fasta(value)
	for _, existing := range allowedV1FastaEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid V1Fasta", value)
}

// NewV1FastaFromValue returns a pointer to a valid V1Fasta
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewV1FastaFromValue(v string) (*V1Fasta, error) {
	ev := V1Fasta(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for V1Fasta: valid values are %v", v, allowedV1FastaEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v V1Fasta) IsValid() bool {
	for _, existing := range allowedV1FastaEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to v1Fasta value
func (v V1Fasta) Ptr() *V1Fasta {
	return &v
}

type NullableV1Fasta struct {
	value *V1Fasta
	isSet bool
}

func (v NullableV1Fasta) Get() *V1Fasta {
	return v.value
}

func (v *NullableV1Fasta) Set(val *V1Fasta) {
	v.value = val
	v.isSet = true
}

func (v NullableV1Fasta) IsSet() bool {
	return v.isSet
}

func (v *NullableV1Fasta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1Fasta(val *V1Fasta) *NullableV1Fasta {
	return &NullableV1Fasta{value: val, isSet: true}
}

func (v NullableV1Fasta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1Fasta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
