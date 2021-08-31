# V1AssemblyDatasetDescriptorsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaxId** | Pointer to **int32** |  | [optional] 
**OrgName** | Pointer to **string** |  | [optional] 
**TaxName** | Pointer to **string** |  | [optional] 
**AssemblyAccession** | Pointer to **string** |  | [optional] 
**Cutoff** | Pointer to **int32** |  | [optional] 
**Limit** | Pointer to **string** |  | [optional] 
**Filters** | Pointer to [**V1AssemblyDatasetDescriptorsFilter**](V1AssemblyDatasetDescriptorsFilter.md) |  | [optional] 
**TaxExactMatch** | Pointer to **bool** |  | [optional] 
**ReturnedContent** | Pointer to [**V1AssemblyDatasetDescriptorsRequestContentType**](V1AssemblyDatasetDescriptorsRequestContentType.md) |  | [optional] [default to V1ASSEMBLYDATASETDESCRIPTORSREQUESTCONTENTTYPE_COMPLETE]

## Methods

### NewV1AssemblyDatasetDescriptorsRequest

`func NewV1AssemblyDatasetDescriptorsRequest() *V1AssemblyDatasetDescriptorsRequest`

NewV1AssemblyDatasetDescriptorsRequest instantiates a new V1AssemblyDatasetDescriptorsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1AssemblyDatasetDescriptorsRequestWithDefaults

`func NewV1AssemblyDatasetDescriptorsRequestWithDefaults() *V1AssemblyDatasetDescriptorsRequest`

NewV1AssemblyDatasetDescriptorsRequestWithDefaults instantiates a new V1AssemblyDatasetDescriptorsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaxId

`func (o *V1AssemblyDatasetDescriptorsRequest) GetTaxId() int32`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *V1AssemblyDatasetDescriptorsRequest) GetTaxIdOk() (*int32, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *V1AssemblyDatasetDescriptorsRequest) SetTaxId(v int32)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *V1AssemblyDatasetDescriptorsRequest) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.

### GetOrgName

`func (o *V1AssemblyDatasetDescriptorsRequest) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *V1AssemblyDatasetDescriptorsRequest) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *V1AssemblyDatasetDescriptorsRequest) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *V1AssemblyDatasetDescriptorsRequest) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetTaxName

`func (o *V1AssemblyDatasetDescriptorsRequest) GetTaxName() string`

GetTaxName returns the TaxName field if non-nil, zero value otherwise.

### GetTaxNameOk

`func (o *V1AssemblyDatasetDescriptorsRequest) GetTaxNameOk() (*string, bool)`

GetTaxNameOk returns a tuple with the TaxName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxName

`func (o *V1AssemblyDatasetDescriptorsRequest) SetTaxName(v string)`

SetTaxName sets TaxName field to given value.

### HasTaxName

`func (o *V1AssemblyDatasetDescriptorsRequest) HasTaxName() bool`

HasTaxName returns a boolean if a field has been set.

### GetAssemblyAccession

`func (o *V1AssemblyDatasetDescriptorsRequest) GetAssemblyAccession() string`

GetAssemblyAccession returns the AssemblyAccession field if non-nil, zero value otherwise.

### GetAssemblyAccessionOk

`func (o *V1AssemblyDatasetDescriptorsRequest) GetAssemblyAccessionOk() (*string, bool)`

GetAssemblyAccessionOk returns a tuple with the AssemblyAccession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssemblyAccession

`func (o *V1AssemblyDatasetDescriptorsRequest) SetAssemblyAccession(v string)`

SetAssemblyAccession sets AssemblyAccession field to given value.

### HasAssemblyAccession

`func (o *V1AssemblyDatasetDescriptorsRequest) HasAssemblyAccession() bool`

HasAssemblyAccession returns a boolean if a field has been set.

### GetCutoff

`func (o *V1AssemblyDatasetDescriptorsRequest) GetCutoff() int32`

GetCutoff returns the Cutoff field if non-nil, zero value otherwise.

### GetCutoffOk

`func (o *V1AssemblyDatasetDescriptorsRequest) GetCutoffOk() (*int32, bool)`

GetCutoffOk returns a tuple with the Cutoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCutoff

`func (o *V1AssemblyDatasetDescriptorsRequest) SetCutoff(v int32)`

SetCutoff sets Cutoff field to given value.

### HasCutoff

`func (o *V1AssemblyDatasetDescriptorsRequest) HasCutoff() bool`

HasCutoff returns a boolean if a field has been set.

### GetLimit

`func (o *V1AssemblyDatasetDescriptorsRequest) GetLimit() string`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *V1AssemblyDatasetDescriptorsRequest) GetLimitOk() (*string, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *V1AssemblyDatasetDescriptorsRequest) SetLimit(v string)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *V1AssemblyDatasetDescriptorsRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetFilters

`func (o *V1AssemblyDatasetDescriptorsRequest) GetFilters() V1AssemblyDatasetDescriptorsFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *V1AssemblyDatasetDescriptorsRequest) GetFiltersOk() (*V1AssemblyDatasetDescriptorsFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *V1AssemblyDatasetDescriptorsRequest) SetFilters(v V1AssemblyDatasetDescriptorsFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *V1AssemblyDatasetDescriptorsRequest) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetTaxExactMatch

`func (o *V1AssemblyDatasetDescriptorsRequest) GetTaxExactMatch() bool`

GetTaxExactMatch returns the TaxExactMatch field if non-nil, zero value otherwise.

### GetTaxExactMatchOk

`func (o *V1AssemblyDatasetDescriptorsRequest) GetTaxExactMatchOk() (*bool, bool)`

GetTaxExactMatchOk returns a tuple with the TaxExactMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxExactMatch

`func (o *V1AssemblyDatasetDescriptorsRequest) SetTaxExactMatch(v bool)`

SetTaxExactMatch sets TaxExactMatch field to given value.

### HasTaxExactMatch

`func (o *V1AssemblyDatasetDescriptorsRequest) HasTaxExactMatch() bool`

HasTaxExactMatch returns a boolean if a field has been set.

### GetReturnedContent

`func (o *V1AssemblyDatasetDescriptorsRequest) GetReturnedContent() V1AssemblyDatasetDescriptorsRequestContentType`

GetReturnedContent returns the ReturnedContent field if non-nil, zero value otherwise.

### GetReturnedContentOk

`func (o *V1AssemblyDatasetDescriptorsRequest) GetReturnedContentOk() (*V1AssemblyDatasetDescriptorsRequestContentType, bool)`

GetReturnedContentOk returns a tuple with the ReturnedContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnedContent

`func (o *V1AssemblyDatasetDescriptorsRequest) SetReturnedContent(v V1AssemblyDatasetDescriptorsRequestContentType)`

SetReturnedContent sets ReturnedContent field to given value.

### HasReturnedContent

`func (o *V1AssemblyDatasetDescriptorsRequest) HasReturnedContent() bool`

HasReturnedContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


