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

// V1SortDirection the model 'V1SortDirection'
type V1SortDirection string

// List of v1SortDirection
const (
	V1SORTDIRECTION_UNSPECIFIED V1SortDirection = "SORT_DIRECTION_UNSPECIFIED"
	V1SORTDIRECTION_ASCENDING V1SortDirection = "SORT_DIRECTION_ASCENDING"
	V1SORTDIRECTION_DESCENDING V1SortDirection = "SORT_DIRECTION_DESCENDING"
)

var allowedV1SortDirectionEnumValues = []V1SortDirection{
	"SORT_DIRECTION_UNSPECIFIED",
	"SORT_DIRECTION_ASCENDING",
	"SORT_DIRECTION_DESCENDING",
}

func (v *V1SortDirection) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := V1SortDirection(value)
	for _, existing := range allowedV1SortDirectionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid V1SortDirection", value)
}

// NewV1SortDirectionFromValue returns a pointer to a valid V1SortDirection
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewV1SortDirectionFromValue(v string) (*V1SortDirection, error) {
	ev := V1SortDirection(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for V1SortDirection: valid values are %v", v, allowedV1SortDirectionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v V1SortDirection) IsValid() bool {
	for _, existing := range allowedV1SortDirectionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to v1SortDirection value
func (v V1SortDirection) Ptr() *V1SortDirection {
	return &v
}

type NullableV1SortDirection struct {
	value *V1SortDirection
	isSet bool
}

func (v NullableV1SortDirection) Get() *V1SortDirection {
	return v.value
}

func (v *NullableV1SortDirection) Set(val *V1SortDirection) {
	v.value = val
	v.isSet = true
}

func (v NullableV1SortDirection) IsSet() bool {
	return v.isSet
}

func (v *NullableV1SortDirection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1SortDirection(val *V1SortDirection) *NullableV1SortDirection {
	return &NullableV1SortDirection{value: val, isSet: true}
}

func (v NullableV1SortDirection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1SortDirection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
