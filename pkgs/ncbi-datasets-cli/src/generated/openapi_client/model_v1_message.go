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
)

// V1Message struct for V1Message
type V1Message struct {
	Error *V1Error `json:"error,omitempty"`
	Warning *V1Warning `json:"warning,omitempty"`
}

// NewV1Message instantiates a new V1Message object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1Message() *V1Message {
	this := V1Message{}
	return &this
}

// NewV1MessageWithDefaults instantiates a new V1Message object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1MessageWithDefaults() *V1Message {
	this := V1Message{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *V1Message) GetError() V1Error {
	if o == nil || o.Error == nil {
		var ret V1Error
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Message) GetErrorOk() (*V1Error, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *V1Message) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given V1Error and assigns it to the Error field.
func (o *V1Message) SetError(v V1Error) {
	o.Error = &v
}

// GetWarning returns the Warning field value if set, zero value otherwise.
func (o *V1Message) GetWarning() V1Warning {
	if o == nil || o.Warning == nil {
		var ret V1Warning
		return ret
	}
	return *o.Warning
}

// GetWarningOk returns a tuple with the Warning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Message) GetWarningOk() (*V1Warning, bool) {
	if o == nil || o.Warning == nil {
		return nil, false
	}
	return o.Warning, true
}

// HasWarning returns a boolean if a field has been set.
func (o *V1Message) HasWarning() bool {
	if o != nil && o.Warning != nil {
		return true
	}

	return false
}

// SetWarning gets a reference to the given V1Warning and assigns it to the Warning field.
func (o *V1Message) SetWarning(v V1Warning) {
	o.Warning = &v
}

func (o V1Message) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Error != nil  {
		toSerialize["error"] = o.Error
	}
	if o.Warning != nil  {
		toSerialize["warning"] = o.Warning
	}
	return json.Marshal(toSerialize)
}

type NullableV1Message struct {
	value *V1Message
	isSet bool
}

func (v NullableV1Message) Get() *V1Message {
	return v.value
}

func (v *NullableV1Message) Set(val *V1Message) {
	v.value = val
	v.isSet = true
}

func (v NullableV1Message) IsSet() bool {
	return v.isSet
}

func (v *NullableV1Message) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1Message(val *V1Message) *NullableV1Message {
	return &NullableV1Message{value: val, isSet: true}
}

func (v NullableV1Message) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1Message) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

