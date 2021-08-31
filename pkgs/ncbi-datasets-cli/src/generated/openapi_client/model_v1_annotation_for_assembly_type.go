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

// V1AnnotationForAssemblyType the model 'V1AnnotationForAssemblyType'
type V1AnnotationForAssemblyType string

// List of v1AnnotationForAssemblyType
const (
	V1ANNOTATIONFORASSEMBLYTYPE_DEFAULT V1AnnotationForAssemblyType = "DEFAULT"
	V1ANNOTATIONFORASSEMBLYTYPE_GENOME_GFF V1AnnotationForAssemblyType = "GENOME_GFF"
	V1ANNOTATIONFORASSEMBLYTYPE_GENOME_GBFF V1AnnotationForAssemblyType = "GENOME_GBFF"
	V1ANNOTATIONFORASSEMBLYTYPE_GENOME_GB V1AnnotationForAssemblyType = "GENOME_GB"
	V1ANNOTATIONFORASSEMBLYTYPE_RNA_FASTA V1AnnotationForAssemblyType = "RNA_FASTA"
	V1ANNOTATIONFORASSEMBLYTYPE_PROT_FASTA V1AnnotationForAssemblyType = "PROT_FASTA"
	V1ANNOTATIONFORASSEMBLYTYPE_GENOME_GTF V1AnnotationForAssemblyType = "GENOME_GTF"
	V1ANNOTATIONFORASSEMBLYTYPE_CDS_FASTA V1AnnotationForAssemblyType = "CDS_FASTA"
)

var allowedV1AnnotationForAssemblyTypeEnumValues = []V1AnnotationForAssemblyType{
	"DEFAULT",
	"GENOME_GFF",
	"GENOME_GBFF",
	"GENOME_GB",
	"RNA_FASTA",
	"PROT_FASTA",
	"GENOME_GTF",
	"CDS_FASTA",
}

func (v *V1AnnotationForAssemblyType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := V1AnnotationForAssemblyType(value)
	for _, existing := range allowedV1AnnotationForAssemblyTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid V1AnnotationForAssemblyType", value)
}

// NewV1AnnotationForAssemblyTypeFromValue returns a pointer to a valid V1AnnotationForAssemblyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewV1AnnotationForAssemblyTypeFromValue(v string) (*V1AnnotationForAssemblyType, error) {
	ev := V1AnnotationForAssemblyType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for V1AnnotationForAssemblyType: valid values are %v", v, allowedV1AnnotationForAssemblyTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v V1AnnotationForAssemblyType) IsValid() bool {
	for _, existing := range allowedV1AnnotationForAssemblyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to v1AnnotationForAssemblyType value
func (v V1AnnotationForAssemblyType) Ptr() *V1AnnotationForAssemblyType {
	return &v
}

type NullableV1AnnotationForAssemblyType struct {
	value *V1AnnotationForAssemblyType
	isSet bool
}

func (v NullableV1AnnotationForAssemblyType) Get() *V1AnnotationForAssemblyType {
	return v.value
}

func (v *NullableV1AnnotationForAssemblyType) Set(val *V1AnnotationForAssemblyType) {
	v.value = val
	v.isSet = true
}

func (v NullableV1AnnotationForAssemblyType) IsSet() bool {
	return v.isSet
}

func (v *NullableV1AnnotationForAssemblyType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1AnnotationForAssemblyType(val *V1AnnotationForAssemblyType) *NullableV1AnnotationForAssemblyType {
	return &NullableV1AnnotationForAssemblyType{value: val, isSet: true}
}

func (v NullableV1AnnotationForAssemblyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1AnnotationForAssemblyType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
