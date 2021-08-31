# V1TaxonomyMetadataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | Pointer to [**[]V1Message**](V1Message.md) |  | [optional] 
**TaxonomyNodes** | Pointer to [**[]V1TaxonomyMatch**](V1TaxonomyMatch.md) |  | [optional] 

## Methods

### NewV1TaxonomyMetadataResponse

`func NewV1TaxonomyMetadataResponse() *V1TaxonomyMetadataResponse`

NewV1TaxonomyMetadataResponse instantiates a new V1TaxonomyMetadataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TaxonomyMetadataResponseWithDefaults

`func NewV1TaxonomyMetadataResponseWithDefaults() *V1TaxonomyMetadataResponse`

NewV1TaxonomyMetadataResponseWithDefaults instantiates a new V1TaxonomyMetadataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *V1TaxonomyMetadataResponse) GetMessages() []V1Message`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *V1TaxonomyMetadataResponse) GetMessagesOk() (*[]V1Message, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *V1TaxonomyMetadataResponse) SetMessages(v []V1Message)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *V1TaxonomyMetadataResponse) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetTaxonomyNodes

`func (o *V1TaxonomyMetadataResponse) GetTaxonomyNodes() []V1TaxonomyMatch`

GetTaxonomyNodes returns the TaxonomyNodes field if non-nil, zero value otherwise.

### GetTaxonomyNodesOk

`func (o *V1TaxonomyMetadataResponse) GetTaxonomyNodesOk() (*[]V1TaxonomyMatch, bool)`

GetTaxonomyNodesOk returns a tuple with the TaxonomyNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxonomyNodes

`func (o *V1TaxonomyMetadataResponse) SetTaxonomyNodes(v []V1TaxonomyMatch)`

SetTaxonomyNodes sets TaxonomyNodes field to given value.

### HasTaxonomyNodes

`func (o *V1TaxonomyMetadataResponse) HasTaxonomyNodes() bool`

HasTaxonomyNodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


