# V1OrthologRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GeneId** | Pointer to **int32** |  | [optional] 
**ReturnedContent** | Pointer to [**V1OrthologRequestContentType**](V1OrthologRequestContentType.md) |  | [optional] [default to V1ORTHOLOGREQUESTCONTENTTYPE_COMPLETE]
**TaxonFilter** | Pointer to **[]string** |  | [optional] 

## Methods

### NewV1OrthologRequest

`func NewV1OrthologRequest() *V1OrthologRequest`

NewV1OrthologRequest instantiates a new V1OrthologRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1OrthologRequestWithDefaults

`func NewV1OrthologRequestWithDefaults() *V1OrthologRequest`

NewV1OrthologRequestWithDefaults instantiates a new V1OrthologRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeneId

`func (o *V1OrthologRequest) GetGeneId() int32`

GetGeneId returns the GeneId field if non-nil, zero value otherwise.

### GetGeneIdOk

`func (o *V1OrthologRequest) GetGeneIdOk() (*int32, bool)`

GetGeneIdOk returns a tuple with the GeneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneId

`func (o *V1OrthologRequest) SetGeneId(v int32)`

SetGeneId sets GeneId field to given value.

### HasGeneId

`func (o *V1OrthologRequest) HasGeneId() bool`

HasGeneId returns a boolean if a field has been set.

### GetReturnedContent

`func (o *V1OrthologRequest) GetReturnedContent() V1OrthologRequestContentType`

GetReturnedContent returns the ReturnedContent field if non-nil, zero value otherwise.

### GetReturnedContentOk

`func (o *V1OrthologRequest) GetReturnedContentOk() (*V1OrthologRequestContentType, bool)`

GetReturnedContentOk returns a tuple with the ReturnedContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnedContent

`func (o *V1OrthologRequest) SetReturnedContent(v V1OrthologRequestContentType)`

SetReturnedContent sets ReturnedContent field to given value.

### HasReturnedContent

`func (o *V1OrthologRequest) HasReturnedContent() bool`

HasReturnedContent returns a boolean if a field has been set.

### GetTaxonFilter

`func (o *V1OrthologRequest) GetTaxonFilter() []string`

GetTaxonFilter returns the TaxonFilter field if non-nil, zero value otherwise.

### GetTaxonFilterOk

`func (o *V1OrthologRequest) GetTaxonFilterOk() (*[]string, bool)`

GetTaxonFilterOk returns a tuple with the TaxonFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxonFilter

`func (o *V1OrthologRequest) SetTaxonFilter(v []string)`

SetTaxonFilter sets TaxonFilter field to given value.

### HasTaxonFilter

`func (o *V1OrthologRequest) HasTaxonFilter() bool`

HasTaxonFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


