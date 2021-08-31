# V1AssemblyMetadataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Taxon** | Pointer to **string** |  | [optional] 
**Accessions** | Pointer to [**V1Accessions**](V1Accessions.md) |  | [optional] 
**Bioprojects** | Pointer to [**V1AssemblyMetadataRequestBioprojects**](V1AssemblyMetadataRequestBioprojects.md) |  | [optional] 
**Filters** | Pointer to [**V1AssemblyDatasetDescriptorsFilter**](V1AssemblyDatasetDescriptorsFilter.md) |  | [optional] 
**TaxExactMatch** | Pointer to **bool** |  | [optional] 
**ReturnedContent** | Pointer to [**V1AssemblyMetadataRequestContentType**](V1AssemblyMetadataRequestContentType.md) |  | [optional] [default to V1ASSEMBLYMETADATAREQUESTCONTENTTYPE_COMPLETE]
**Limit** | Pointer to **string** |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**PageToken** | Pointer to **string** |  | [optional] 

## Methods

### NewV1AssemblyMetadataRequest

`func NewV1AssemblyMetadataRequest() *V1AssemblyMetadataRequest`

NewV1AssemblyMetadataRequest instantiates a new V1AssemblyMetadataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1AssemblyMetadataRequestWithDefaults

`func NewV1AssemblyMetadataRequestWithDefaults() *V1AssemblyMetadataRequest`

NewV1AssemblyMetadataRequestWithDefaults instantiates a new V1AssemblyMetadataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaxon

`func (o *V1AssemblyMetadataRequest) GetTaxon() string`

GetTaxon returns the Taxon field if non-nil, zero value otherwise.

### GetTaxonOk

`func (o *V1AssemblyMetadataRequest) GetTaxonOk() (*string, bool)`

GetTaxonOk returns a tuple with the Taxon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxon

`func (o *V1AssemblyMetadataRequest) SetTaxon(v string)`

SetTaxon sets Taxon field to given value.

### HasTaxon

`func (o *V1AssemblyMetadataRequest) HasTaxon() bool`

HasTaxon returns a boolean if a field has been set.

### GetAccessions

`func (o *V1AssemblyMetadataRequest) GetAccessions() V1Accessions`

GetAccessions returns the Accessions field if non-nil, zero value otherwise.

### GetAccessionsOk

`func (o *V1AssemblyMetadataRequest) GetAccessionsOk() (*V1Accessions, bool)`

GetAccessionsOk returns a tuple with the Accessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessions

`func (o *V1AssemblyMetadataRequest) SetAccessions(v V1Accessions)`

SetAccessions sets Accessions field to given value.

### HasAccessions

`func (o *V1AssemblyMetadataRequest) HasAccessions() bool`

HasAccessions returns a boolean if a field has been set.

### GetBioprojects

`func (o *V1AssemblyMetadataRequest) GetBioprojects() V1AssemblyMetadataRequestBioprojects`

GetBioprojects returns the Bioprojects field if non-nil, zero value otherwise.

### GetBioprojectsOk

`func (o *V1AssemblyMetadataRequest) GetBioprojectsOk() (*V1AssemblyMetadataRequestBioprojects, bool)`

GetBioprojectsOk returns a tuple with the Bioprojects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBioprojects

`func (o *V1AssemblyMetadataRequest) SetBioprojects(v V1AssemblyMetadataRequestBioprojects)`

SetBioprojects sets Bioprojects field to given value.

### HasBioprojects

`func (o *V1AssemblyMetadataRequest) HasBioprojects() bool`

HasBioprojects returns a boolean if a field has been set.

### GetFilters

`func (o *V1AssemblyMetadataRequest) GetFilters() V1AssemblyDatasetDescriptorsFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *V1AssemblyMetadataRequest) GetFiltersOk() (*V1AssemblyDatasetDescriptorsFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *V1AssemblyMetadataRequest) SetFilters(v V1AssemblyDatasetDescriptorsFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *V1AssemblyMetadataRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetTaxExactMatch

`func (o *V1AssemblyMetadataRequest) GetTaxExactMatch() bool`

GetTaxExactMatch returns the TaxExactMatch field if non-nil, zero value otherwise.

### GetTaxExactMatchOk

`func (o *V1AssemblyMetadataRequest) GetTaxExactMatchOk() (*bool, bool)`

GetTaxExactMatchOk returns a tuple with the TaxExactMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxExactMatch

`func (o *V1AssemblyMetadataRequest) SetTaxExactMatch(v bool)`

SetTaxExactMatch sets TaxExactMatch field to given value.

### HasTaxExactMatch

`func (o *V1AssemblyMetadataRequest) HasTaxExactMatch() bool`

HasTaxExactMatch returns a boolean if a field has been set.

### GetReturnedContent

`func (o *V1AssemblyMetadataRequest) GetReturnedContent() V1AssemblyMetadataRequestContentType`

GetReturnedContent returns the ReturnedContent field if non-nil, zero value otherwise.

### GetReturnedContentOk

`func (o *V1AssemblyMetadataRequest) GetReturnedContentOk() (*V1AssemblyMetadataRequestContentType, bool)`

GetReturnedContentOk returns a tuple with the ReturnedContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnedContent

`func (o *V1AssemblyMetadataRequest) SetReturnedContent(v V1AssemblyMetadataRequestContentType)`

SetReturnedContent sets ReturnedContent field to given value.

### HasReturnedContent

`func (o *V1AssemblyMetadataRequest) HasReturnedContent() bool`

HasReturnedContent returns a boolean if a field has been set.

### GetLimit

`func (o *V1AssemblyMetadataRequest) GetLimit() string`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *V1AssemblyMetadataRequest) GetLimitOk() (*string, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *V1AssemblyMetadataRequest) SetLimit(v string)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *V1AssemblyMetadataRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetPageSize

`func (o *V1AssemblyMetadataRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *V1AssemblyMetadataRequest) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *V1AssemblyMetadataRequest) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *V1AssemblyMetadataRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetPageToken

`func (o *V1AssemblyMetadataRequest) GetPageToken() string`

GetPageToken returns the PageToken field if non-nil, zero value otherwise.

### GetPageTokenOk

`func (o *V1AssemblyMetadataRequest) GetPageTokenOk() (*string, bool)`

GetPageTokenOk returns a tuple with the PageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageToken

`func (o *V1AssemblyMetadataRequest) SetPageToken(v string)`

SetPageToken sets PageToken field to given value.

### HasPageToken

`func (o *V1AssemblyMetadataRequest) HasPageToken() bool`

HasPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


