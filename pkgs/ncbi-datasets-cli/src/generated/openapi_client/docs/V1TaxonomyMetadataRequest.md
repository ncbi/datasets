# V1TaxonomyMetadataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Taxons** | Pointer to **[]string** |  | [optional] 
**ReturnedContent** | Pointer to [**V1TaxonomyMetadataRequestContentType**](V1TaxonomyMetadataRequestContentType.md) |  | [optional] [default to V1TAXONOMYMETADATAREQUESTCONTENTTYPE_COMPLETE]

## Methods

### NewV1TaxonomyMetadataRequest

`func NewV1TaxonomyMetadataRequest() *V1TaxonomyMetadataRequest`

NewV1TaxonomyMetadataRequest instantiates a new V1TaxonomyMetadataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TaxonomyMetadataRequestWithDefaults

`func NewV1TaxonomyMetadataRequestWithDefaults() *V1TaxonomyMetadataRequest`

NewV1TaxonomyMetadataRequestWithDefaults instantiates a new V1TaxonomyMetadataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaxons

`func (o *V1TaxonomyMetadataRequest) GetTaxons() []string`

GetTaxons returns the Taxons field if non-nil, zero value otherwise.

### GetTaxonsOk

`func (o *V1TaxonomyMetadataRequest) GetTaxonsOk() (*[]string, bool)`

GetTaxonsOk returns a tuple with the Taxons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxons

`func (o *V1TaxonomyMetadataRequest) SetTaxons(v []string)`

SetTaxons sets Taxons field to given value.

### HasTaxons

`func (o *V1TaxonomyMetadataRequest) HasTaxons() bool`

HasTaxons returns a boolean if a field has been set.

### GetReturnedContent

`func (o *V1TaxonomyMetadataRequest) GetReturnedContent() V1TaxonomyMetadataRequestContentType`

GetReturnedContent returns the ReturnedContent field if non-nil, zero value otherwise.

### GetReturnedContentOk

`func (o *V1TaxonomyMetadataRequest) GetReturnedContentOk() (*V1TaxonomyMetadataRequestContentType, bool)`

GetReturnedContentOk returns a tuple with the ReturnedContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnedContent

`func (o *V1TaxonomyMetadataRequest) SetReturnedContent(v V1TaxonomyMetadataRequestContentType)`

SetReturnedContent sets ReturnedContent field to given value.

### HasReturnedContent

`func (o *V1TaxonomyMetadataRequest) HasReturnedContent() bool`

HasReturnedContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


