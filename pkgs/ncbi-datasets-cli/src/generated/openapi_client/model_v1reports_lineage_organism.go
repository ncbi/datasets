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

// V1reportsLineageOrganism struct for V1reportsLineageOrganism
type V1reportsLineageOrganism struct {
	TaxId *int32 `json:"tax_id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewV1reportsLineageOrganism instantiates a new V1reportsLineageOrganism object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1reportsLineageOrganism() *V1reportsLineageOrganism {
	this := V1reportsLineageOrganism{}
	return &this
}

// NewV1reportsLineageOrganismWithDefaults instantiates a new V1reportsLineageOrganism object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1reportsLineageOrganismWithDefaults() *V1reportsLineageOrganism {
	this := V1reportsLineageOrganism{}
	return &this
}

// GetTaxId returns the TaxId field value if set, zero value otherwise.
func (o *V1reportsLineageOrganism) GetTaxId() int32 {
	if o == nil || o.TaxId == nil {
		var ret int32
		return ret
	}
	return *o.TaxId
}

// GetTaxIdOk returns a tuple with the TaxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsLineageOrganism) GetTaxIdOk() (*int32, bool) {
	if o == nil || o.TaxId == nil {
		return nil, false
	}
	return o.TaxId, true
}

// HasTaxId returns a boolean if a field has been set.
func (o *V1reportsLineageOrganism) HasTaxId() bool {
	if o != nil && o.TaxId != nil {
		return true
	}

	return false
}

// SetTaxId gets a reference to the given int32 and assigns it to the TaxId field.
func (o *V1reportsLineageOrganism) SetTaxId(v int32) {
	o.TaxId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V1reportsLineageOrganism) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsLineageOrganism) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V1reportsLineageOrganism) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V1reportsLineageOrganism) SetName(v string) {
	o.Name = &v
}

func (o V1reportsLineageOrganism) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TaxId != nil  {
		toSerialize["tax_id"] = o.TaxId
	}
	if o.Name != nil  {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableV1reportsLineageOrganism struct {
	value *V1reportsLineageOrganism
	isSet bool
}

func (v NullableV1reportsLineageOrganism) Get() *V1reportsLineageOrganism {
	return v.value
}

func (v *NullableV1reportsLineageOrganism) Set(val *V1reportsLineageOrganism) {
	v.value = val
	v.isSet = true
}

func (v NullableV1reportsLineageOrganism) IsSet() bool {
	return v.isSet
}

func (v *NullableV1reportsLineageOrganism) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1reportsLineageOrganism(val *V1reportsLineageOrganism) *NullableV1reportsLineageOrganism {
	return &NullableV1reportsLineageOrganism{value: val, isSet: true}
}

func (v NullableV1reportsLineageOrganism) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1reportsLineageOrganism) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


