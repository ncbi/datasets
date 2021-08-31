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

// V1Orientation the model 'V1Orientation'
type V1Orientation string

// List of v1Orientation
const (
	V1ORIENTATION_NONE V1Orientation = "none"
	V1ORIENTATION_PLUS V1Orientation = "plus"
	V1ORIENTATION_MINUS V1Orientation = "minus"
)

var allowedV1OrientationEnumValues = []V1Orientation{
	"none",
	"plus",
	"minus",
}

func (v *V1Orientation) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := V1Orientation(value)
	for _, existing := range allowedV1OrientationEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid V1Orientation", value)
}

// NewV1OrientationFromValue returns a pointer to a valid V1Orientation
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewV1OrientationFromValue(v string) (*V1Orientation, error) {
	ev := V1Orientation(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for V1Orientation: valid values are %v", v, allowedV1OrientationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v V1Orientation) IsValid() bool {
	for _, existing := range allowedV1OrientationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to v1Orientation value
func (v V1Orientation) Ptr() *V1Orientation {
	return &v
}

type NullableV1Orientation struct {
	value *V1Orientation
	isSet bool
}

func (v NullableV1Orientation) Get() *V1Orientation {
	return v.value
}

func (v *NullableV1Orientation) Set(val *V1Orientation) {
	v.value = val
	v.isSet = true
}

func (v NullableV1Orientation) IsSet() bool {
	return v.isSet
}

func (v *NullableV1Orientation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1Orientation(val *V1Orientation) *NullableV1Orientation {
	return &NullableV1Orientation{value: val, isSet: true}
}

func (v NullableV1Orientation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1Orientation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
