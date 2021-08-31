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

// V1AssemblyMetadataRequest struct for V1AssemblyMetadataRequest
type V1AssemblyMetadataRequest struct {
	Taxon *string `json:"taxon,omitempty"`
	Accessions *V1Accessions `json:"accessions,omitempty"`
	Bioprojects *V1AssemblyMetadataRequestBioprojects `json:"bioprojects,omitempty"`
	Filters *V1AssemblyDatasetDescriptorsFilter `json:"filters,omitempty"`
	TaxExactMatch *bool `json:"tax_exact_match,omitempty"`
	ReturnedContent *V1AssemblyMetadataRequestContentType `json:"returned_content,omitempty"`
	Limit *string `json:"limit,omitempty"`
	PageSize *int32 `json:"page_size,omitempty"`
	PageToken *string `json:"page_token,omitempty"`
}

// NewV1AssemblyMetadataRequest instantiates a new V1AssemblyMetadataRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1AssemblyMetadataRequest() *V1AssemblyMetadataRequest {
	this := V1AssemblyMetadataRequest{}
	var returnedContent V1AssemblyMetadataRequestContentType = V1ASSEMBLYMETADATAREQUESTCONTENTTYPE_COMPLETE
	this.ReturnedContent = &returnedContent
	return &this
}

// NewV1AssemblyMetadataRequestWithDefaults instantiates a new V1AssemblyMetadataRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1AssemblyMetadataRequestWithDefaults() *V1AssemblyMetadataRequest {
	this := V1AssemblyMetadataRequest{}
	var returnedContent V1AssemblyMetadataRequestContentType = V1ASSEMBLYMETADATAREQUESTCONTENTTYPE_COMPLETE
	this.ReturnedContent = &returnedContent
	return &this
}

// GetTaxon returns the Taxon field value if set, zero value otherwise.
func (o *V1AssemblyMetadataRequest) GetTaxon() string {
	if o == nil || o.Taxon == nil {
		var ret string
		return ret
	}
	return *o.Taxon
}

// GetTaxonOk returns a tuple with the Taxon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AssemblyMetadataRequest) GetTaxonOk() (*string, bool) {
	if o == nil || o.Taxon == nil {
		return nil, false
	}
	return o.Taxon, true
}

// HasTaxon returns a boolean if a field has been set.
func (o *V1AssemblyMetadataRequest) HasTaxon() bool {
	if o != nil && o.Taxon != nil {
		return true
	}

	return false
}

// SetTaxon gets a reference to the given string and assigns it to the Taxon field.
func (o *V1AssemblyMetadataRequest) SetTaxon(v string) {
	o.Taxon = &v
}

// GetAccessions returns the Accessions field value if set, zero value otherwise.
func (o *V1AssemblyMetadataRequest) GetAccessions() V1Accessions {
	if o == nil || o.Accessions == nil {
		var ret V1Accessions
		return ret
	}
	return *o.Accessions
}

// GetAccessionsOk returns a tuple with the Accessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AssemblyMetadataRequest) GetAccessionsOk() (*V1Accessions, bool) {
	if o == nil || o.Accessions == nil {
		return nil, false
	}
	return o.Accessions, true
}

// HasAccessions returns a boolean if a field has been set.
func (o *V1AssemblyMetadataRequest) HasAccessions() bool {
	if o != nil && o.Accessions != nil {
		return true
	}

	return false
}

// SetAccessions gets a reference to the given V1Accessions and assigns it to the Accessions field.
func (o *V1AssemblyMetadataRequest) SetAccessions(v V1Accessions) {
	o.Accessions = &v
}

// GetBioprojects returns the Bioprojects field value if set, zero value otherwise.
func (o *V1AssemblyMetadataRequest) GetBioprojects() V1AssemblyMetadataRequestBioprojects {
	if o == nil || o.Bioprojects == nil {
		var ret V1AssemblyMetadataRequestBioprojects
		return ret
	}
	return *o.Bioprojects
}

// GetBioprojectsOk returns a tuple with the Bioprojects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AssemblyMetadataRequest) GetBioprojectsOk() (*V1AssemblyMetadataRequestBioprojects, bool) {
	if o == nil || o.Bioprojects == nil {
		return nil, false
	}
	return o.Bioprojects, true
}

// HasBioprojects returns a boolean if a field has been set.
func (o *V1AssemblyMetadataRequest) HasBioprojects() bool {
	if o != nil && o.Bioprojects != nil {
		return true
	}

	return false
}

// SetBioprojects gets a reference to the given V1AssemblyMetadataRequestBioprojects and assigns it to the Bioprojects field.
func (o *V1AssemblyMetadataRequest) SetBioprojects(v V1AssemblyMetadataRequestBioprojects) {
	o.Bioprojects = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *V1AssemblyMetadataRequest) GetFilters() V1AssemblyDatasetDescriptorsFilter {
	if o == nil || o.Filters == nil {
		var ret V1AssemblyDatasetDescriptorsFilter
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AssemblyMetadataRequest) GetFiltersOk() (*V1AssemblyDatasetDescriptorsFilter, bool) {
	if o == nil || o.Filters == nil {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *V1AssemblyMetadataRequest) HasFilters() bool {
	if o != nil && o.Filters != nil {
		return true
	}

	return false
}

// SetFilters gets a reference to the given V1AssemblyDatasetDescriptorsFilter and assigns it to the Filters field.
func (o *V1AssemblyMetadataRequest) SetFilters(v V1AssemblyDatasetDescriptorsFilter) {
	o.Filters = &v
}

// GetTaxExactMatch returns the TaxExactMatch field value if set, zero value otherwise.
func (o *V1AssemblyMetadataRequest) GetTaxExactMatch() bool {
	if o == nil || o.TaxExactMatch == nil {
		var ret bool
		return ret
	}
	return *o.TaxExactMatch
}

// GetTaxExactMatchOk returns a tuple with the TaxExactMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AssemblyMetadataRequest) GetTaxExactMatchOk() (*bool, bool) {
	if o == nil || o.TaxExactMatch == nil {
		return nil, false
	}
	return o.TaxExactMatch, true
}

// HasTaxExactMatch returns a boolean if a field has been set.
func (o *V1AssemblyMetadataRequest) HasTaxExactMatch() bool {
	if o != nil && o.TaxExactMatch != nil {
		return true
	}

	return false
}

// SetTaxExactMatch gets a reference to the given bool and assigns it to the TaxExactMatch field.
func (o *V1AssemblyMetadataRequest) SetTaxExactMatch(v bool) {
	o.TaxExactMatch = &v
}

// GetReturnedContent returns the ReturnedContent field value if set, zero value otherwise.
func (o *V1AssemblyMetadataRequest) GetReturnedContent() V1AssemblyMetadataRequestContentType {
	if o == nil || o.ReturnedContent == nil {
		var ret V1AssemblyMetadataRequestContentType
		return ret
	}
	return *o.ReturnedContent
}

// GetReturnedContentOk returns a tuple with the ReturnedContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AssemblyMetadataRequest) GetReturnedContentOk() (*V1AssemblyMetadataRequestContentType, bool) {
	if o == nil || o.ReturnedContent == nil {
		return nil, false
	}
	return o.ReturnedContent, true
}

// HasReturnedContent returns a boolean if a field has been set.
func (o *V1AssemblyMetadataRequest) HasReturnedContent() bool {
	if o != nil && o.ReturnedContent != nil {
		return true
	}

	return false
}

// SetReturnedContent gets a reference to the given V1AssemblyMetadataRequestContentType and assigns it to the ReturnedContent field.
func (o *V1AssemblyMetadataRequest) SetReturnedContent(v V1AssemblyMetadataRequestContentType) {
	o.ReturnedContent = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *V1AssemblyMetadataRequest) GetLimit() string {
	if o == nil || o.Limit == nil {
		var ret string
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AssemblyMetadataRequest) GetLimitOk() (*string, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *V1AssemblyMetadataRequest) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given string and assigns it to the Limit field.
func (o *V1AssemblyMetadataRequest) SetLimit(v string) {
	o.Limit = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *V1AssemblyMetadataRequest) GetPageSize() int32 {
	if o == nil || o.PageSize == nil {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AssemblyMetadataRequest) GetPageSizeOk() (*int32, bool) {
	if o == nil || o.PageSize == nil {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *V1AssemblyMetadataRequest) HasPageSize() bool {
	if o != nil && o.PageSize != nil {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *V1AssemblyMetadataRequest) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetPageToken returns the PageToken field value if set, zero value otherwise.
func (o *V1AssemblyMetadataRequest) GetPageToken() string {
	if o == nil || o.PageToken == nil {
		var ret string
		return ret
	}
	return *o.PageToken
}

// GetPageTokenOk returns a tuple with the PageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1AssemblyMetadataRequest) GetPageTokenOk() (*string, bool) {
	if o == nil || o.PageToken == nil {
		return nil, false
	}
	return o.PageToken, true
}

// HasPageToken returns a boolean if a field has been set.
func (o *V1AssemblyMetadataRequest) HasPageToken() bool {
	if o != nil && o.PageToken != nil {
		return true
	}

	return false
}

// SetPageToken gets a reference to the given string and assigns it to the PageToken field.
func (o *V1AssemblyMetadataRequest) SetPageToken(v string) {
	o.PageToken = &v
}

func (o V1AssemblyMetadataRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Taxon != nil  {
		toSerialize["taxon"] = o.Taxon
	}
	if o.Accessions != nil  {
		toSerialize["accessions"] = o.Accessions
	}
	if o.Bioprojects != nil  {
		toSerialize["bioprojects"] = o.Bioprojects
	}
	if o.Filters != nil  {
		toSerialize["filters"] = o.Filters
	}
	if o.TaxExactMatch != nil  {
		toSerialize["tax_exact_match"] = o.TaxExactMatch
	}
	if o.ReturnedContent != nil  {
		toSerialize["returned_content"] = o.ReturnedContent
	}
	if o.Limit != nil  {
		toSerialize["limit"] = o.Limit
	}
	if o.PageSize != nil  {
		toSerialize["page_size"] = o.PageSize
	}
	if o.PageToken != nil  {
		toSerialize["page_token"] = o.PageToken
	}
	return json.Marshal(toSerialize)
}

type NullableV1AssemblyMetadataRequest struct {
	value *V1AssemblyMetadataRequest
	isSet bool
}

func (v NullableV1AssemblyMetadataRequest) Get() *V1AssemblyMetadataRequest {
	return v.value
}

func (v *NullableV1AssemblyMetadataRequest) Set(val *V1AssemblyMetadataRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV1AssemblyMetadataRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV1AssemblyMetadataRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1AssemblyMetadataRequest(val *V1AssemblyMetadataRequest) *NullableV1AssemblyMetadataRequest {
	return &NullableV1AssemblyMetadataRequest{value: val, isSet: true}
}

func (v NullableV1AssemblyMetadataRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1AssemblyMetadataRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

