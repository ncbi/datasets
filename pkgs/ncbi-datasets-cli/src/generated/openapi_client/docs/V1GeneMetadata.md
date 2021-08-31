# V1GeneMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | Pointer to [**[]V1Message**](V1Message.md) |  | [optional] 
**Genes** | Pointer to [**[]V1GeneMatch**](V1GeneMatch.md) |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewV1GeneMetadata

`func NewV1GeneMetadata() *V1GeneMetadata`

NewV1GeneMetadata instantiates a new V1GeneMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1GeneMetadataWithDefaults

`func NewV1GeneMetadataWithDefaults() *V1GeneMetadata`

NewV1GeneMetadataWithDefaults instantiates a new V1GeneMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *V1GeneMetadata) GetMessages() []V1Message`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *V1GeneMetadata) GetMessagesOk() (*[]V1Message, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *V1GeneMetadata) SetMessages(v []V1Message)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *V1GeneMetadata) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetGenes

`func (o *V1GeneMetadata) GetGenes() []V1GeneMatch`

GetGenes returns the Genes field if non-nil, zero value otherwise.

### GetGenesOk

`func (o *V1GeneMetadata) GetGenesOk() (*[]V1GeneMatch, bool)`

GetGenesOk returns a tuple with the Genes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenes

`func (o *V1GeneMetadata) SetGenes(v []V1GeneMatch)`

SetGenes sets Genes field to given value.

### HasGenes

`func (o *V1GeneMetadata) HasGenes() bool`

HasGenes returns a boolean if a field has been set.

### GetTotalCount

`func (o *V1GeneMetadata) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *V1GeneMetadata) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *V1GeneMetadata) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *V1GeneMetadata) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


