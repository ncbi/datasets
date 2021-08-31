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

// V1reportsBiosample struct for V1reportsBiosample
type V1reportsBiosample struct {
	GeographicOrigin *string `json:"geographic_origin,omitempty"`
	Source *string `json:"source,omitempty"`
	Type *string `json:"type,omitempty"`
	Accession *string `json:"accession,omitempty"`
	Assembly *string `json:"assembly,omitempty"`
	CollectionDate *string `json:"collection_date,omitempty"`
}

// NewV1reportsBiosample instantiates a new V1reportsBiosample object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1reportsBiosample() *V1reportsBiosample {
	this := V1reportsBiosample{}
	return &this
}

// NewV1reportsBiosampleWithDefaults instantiates a new V1reportsBiosample object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1reportsBiosampleWithDefaults() *V1reportsBiosample {
	this := V1reportsBiosample{}
	return &this
}

// GetGeographicOrigin returns the GeographicOrigin field value if set, zero value otherwise.
func (o *V1reportsBiosample) GetGeographicOrigin() string {
	if o == nil || o.GeographicOrigin == nil {
		var ret string
		return ret
	}
	return *o.GeographicOrigin
}

// GetGeographicOriginOk returns a tuple with the GeographicOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsBiosample) GetGeographicOriginOk() (*string, bool) {
	if o == nil || o.GeographicOrigin == nil {
		return nil, false
	}
	return o.GeographicOrigin, true
}

// HasGeographicOrigin returns a boolean if a field has been set.
func (o *V1reportsBiosample) HasGeographicOrigin() bool {
	if o != nil && o.GeographicOrigin != nil {
		return true
	}

	return false
}

// SetGeographicOrigin gets a reference to the given string and assigns it to the GeographicOrigin field.
func (o *V1reportsBiosample) SetGeographicOrigin(v string) {
	o.GeographicOrigin = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *V1reportsBiosample) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsBiosample) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *V1reportsBiosample) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *V1reportsBiosample) SetSource(v string) {
	o.Source = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *V1reportsBiosample) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsBiosample) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *V1reportsBiosample) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *V1reportsBiosample) SetType(v string) {
	o.Type = &v
}

// GetAccession returns the Accession field value if set, zero value otherwise.
func (o *V1reportsBiosample) GetAccession() string {
	if o == nil || o.Accession == nil {
		var ret string
		return ret
	}
	return *o.Accession
}

// GetAccessionOk returns a tuple with the Accession field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsBiosample) GetAccessionOk() (*string, bool) {
	if o == nil || o.Accession == nil {
		return nil, false
	}
	return o.Accession, true
}

// HasAccession returns a boolean if a field has been set.
func (o *V1reportsBiosample) HasAccession() bool {
	if o != nil && o.Accession != nil {
		return true
	}

	return false
}

// SetAccession gets a reference to the given string and assigns it to the Accession field.
func (o *V1reportsBiosample) SetAccession(v string) {
	o.Accession = &v
}

// GetAssembly returns the Assembly field value if set, zero value otherwise.
func (o *V1reportsBiosample) GetAssembly() string {
	if o == nil || o.Assembly == nil {
		var ret string
		return ret
	}
	return *o.Assembly
}

// GetAssemblyOk returns a tuple with the Assembly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsBiosample) GetAssemblyOk() (*string, bool) {
	if o == nil || o.Assembly == nil {
		return nil, false
	}
	return o.Assembly, true
}

// HasAssembly returns a boolean if a field has been set.
func (o *V1reportsBiosample) HasAssembly() bool {
	if o != nil && o.Assembly != nil {
		return true
	}

	return false
}

// SetAssembly gets a reference to the given string and assigns it to the Assembly field.
func (o *V1reportsBiosample) SetAssembly(v string) {
	o.Assembly = &v
}

// GetCollectionDate returns the CollectionDate field value if set, zero value otherwise.
func (o *V1reportsBiosample) GetCollectionDate() string {
	if o == nil || o.CollectionDate == nil {
		var ret string
		return ret
	}
	return *o.CollectionDate
}

// GetCollectionDateOk returns a tuple with the CollectionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1reportsBiosample) GetCollectionDateOk() (*string, bool) {
	if o == nil || o.CollectionDate == nil {
		return nil, false
	}
	return o.CollectionDate, true
}

// HasCollectionDate returns a boolean if a field has been set.
func (o *V1reportsBiosample) HasCollectionDate() bool {
	if o != nil && o.CollectionDate != nil {
		return true
	}

	return false
}

// SetCollectionDate gets a reference to the given string and assigns it to the CollectionDate field.
func (o *V1reportsBiosample) SetCollectionDate(v string) {
	o.CollectionDate = &v
}

func (o V1reportsBiosample) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GeographicOrigin != nil  {
		toSerialize["geographic_origin"] = o.GeographicOrigin
	}
	if o.Source != nil  {
		toSerialize["source"] = o.Source
	}
	if o.Type != nil  {
		toSerialize["type"] = o.Type
	}
	if o.Accession != nil  {
		toSerialize["accession"] = o.Accession
	}
	if o.Assembly != nil  {
		toSerialize["assembly"] = o.Assembly
	}
	if o.CollectionDate != nil  {
		toSerialize["collection_date"] = o.CollectionDate
	}
	return json.Marshal(toSerialize)
}

type NullableV1reportsBiosample struct {
	value *V1reportsBiosample
	isSet bool
}

func (v NullableV1reportsBiosample) Get() *V1reportsBiosample {
	return v.value
}

func (v *NullableV1reportsBiosample) Set(val *V1reportsBiosample) {
	v.value = val
	v.isSet = true
}

func (v NullableV1reportsBiosample) IsSet() bool {
	return v.isSet
}

func (v *NullableV1reportsBiosample) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1reportsBiosample(val *V1reportsBiosample) *NullableV1reportsBiosample {
	return &NullableV1reportsBiosample{value: val, isSet: true}
}

func (v NullableV1reportsBiosample) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1reportsBiosample) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

