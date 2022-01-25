/*
NCBI Datasets API

### NCBI Datasets is a resource that lets you easily gather data from NCBI. The Datasets API is still in alpha, and we're updating it often to add new functionality, iron out bugs and enhance usability. For some larger downloads, you may want to download a [dehydrated bag](https://www.ncbi.nlm.nih.gov/datasets/docs/rehydrate/), and retrieve the individual data files at a later time. 

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datasets

import (
	"encoding/json"
)

// V1reportsBioSampleAttribute struct for V1reportsBioSampleAttribute
type V1reportsBioSampleAttribute struct {
	Name *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewV1reportsBioSampleAttribute instantiates a new V1reportsBioSampleAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1reportsBioSampleAttribute() *V1reportsBioSampleAttribute {
	this := V1reportsBioSampleAttribute{}
	return &this
}

// NewV1reportsBioSampleAttributeWithDefaults instantiates a new V1reportsBioSampleAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1reportsBioSampleAttributeWithDefaults() *V1reportsBioSampleAttribute {
	this := V1reportsBioSampleAttribute{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V1reportsBioSampleAttribute) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsBioSampleAttribute) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V1reportsBioSampleAttribute) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V1reportsBioSampleAttribute) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *V1reportsBioSampleAttribute) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsBioSampleAttribute) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *V1reportsBioSampleAttribute) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *V1reportsBioSampleAttribute) SetValue(v string) {
	o.Value = &v
}

func (o V1reportsBioSampleAttribute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil  {
		toSerialize["name"] = o.Name
	}
	if o.Value != nil  {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableV1reportsBioSampleAttribute struct {
	value *V1reportsBioSampleAttribute
	isSet bool
}

func (v NullableV1reportsBioSampleAttribute) Get() *V1reportsBioSampleAttribute {
	return v.value
}

func (v *NullableV1reportsBioSampleAttribute) Set(val *V1reportsBioSampleAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullableV1reportsBioSampleAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableV1reportsBioSampleAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1reportsBioSampleAttribute(val *V1reportsBioSampleAttribute) *NullableV1reportsBioSampleAttribute {
	return &NullableV1reportsBioSampleAttribute{value: val, isSet: true}
}

func (v NullableV1reportsBioSampleAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1reportsBioSampleAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


