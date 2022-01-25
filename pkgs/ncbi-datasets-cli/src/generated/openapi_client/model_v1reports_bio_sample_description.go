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

// V1reportsBioSampleDescription struct for V1reportsBioSampleDescription
type V1reportsBioSampleDescription struct {
	Title *string `json:"title,omitempty"`
	Organism *V1reportsOrganism `json:"organism,omitempty"`
	Comment *string `json:"comment,omitempty"`
}

// NewV1reportsBioSampleDescription instantiates a new V1reportsBioSampleDescription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1reportsBioSampleDescription() *V1reportsBioSampleDescription {
	this := V1reportsBioSampleDescription{}
	return &this
}

// NewV1reportsBioSampleDescriptionWithDefaults instantiates a new V1reportsBioSampleDescription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1reportsBioSampleDescriptionWithDefaults() *V1reportsBioSampleDescription {
	this := V1reportsBioSampleDescription{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *V1reportsBioSampleDescription) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsBioSampleDescription) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *V1reportsBioSampleDescription) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *V1reportsBioSampleDescription) SetTitle(v string) {
	o.Title = &v
}

// GetOrganism returns the Organism field value if set, zero value otherwise.
func (o *V1reportsBioSampleDescription) GetOrganism() V1reportsOrganism {
	if o == nil || o.Organism == nil {
		var ret V1reportsOrganism
		return ret
	}
	return *o.Organism
}

// GetOrganismOk returns a tuple with the Organism field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsBioSampleDescription) GetOrganismOk() (*V1reportsOrganism, bool) {
	if o == nil || o.Organism == nil {
		return nil, false
	}
	return o.Organism, true
}

// HasOrganism returns a boolean if a field has been set.
func (o *V1reportsBioSampleDescription) HasOrganism() bool {
	if o != nil && o.Organism != nil {
		return true
	}

	return false
}

// SetOrganism gets a reference to the given V1reportsOrganism and assigns it to the Organism field.
func (o *V1reportsBioSampleDescription) SetOrganism(v V1reportsOrganism) {
	o.Organism = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *V1reportsBioSampleDescription) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsBioSampleDescription) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *V1reportsBioSampleDescription) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *V1reportsBioSampleDescription) SetComment(v string) {
	o.Comment = &v
}

func (o V1reportsBioSampleDescription) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Title != nil  {
		toSerialize["title"] = o.Title
	}
	if o.Organism != nil  {
		toSerialize["organism"] = o.Organism
	}
	if o.Comment != nil  {
		toSerialize["comment"] = o.Comment
	}
	return json.Marshal(toSerialize)
}

type NullableV1reportsBioSampleDescription struct {
	value *V1reportsBioSampleDescription
	isSet bool
}

func (v NullableV1reportsBioSampleDescription) Get() *V1reportsBioSampleDescription {
	return v.value
}

func (v *NullableV1reportsBioSampleDescription) Set(val *V1reportsBioSampleDescription) {
	v.value = val
	v.isSet = true
}

func (v NullableV1reportsBioSampleDescription) IsSet() bool {
	return v.isSet
}

func (v *NullableV1reportsBioSampleDescription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1reportsBioSampleDescription(val *V1reportsBioSampleDescription) *NullableV1reportsBioSampleDescription {
	return &NullableV1reportsBioSampleDescription{value: val, isSet: true}
}

func (v NullableV1reportsBioSampleDescription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1reportsBioSampleDescription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


