# V1OrthologSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrthologSetId** | Pointer to **int32** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Genes** | Pointer to [**V1GeneMetadata**](V1GeneMetadata.md) |  | [optional] 

## Methods

### NewV1OrthologSet

`func NewV1OrthologSet() *V1OrthologSet`

NewV1OrthologSet instantiates a new V1OrthologSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1OrthologSetWithDefaults

`func NewV1OrthologSetWithDefaults() *V1OrthologSet`

NewV1OrthologSetWithDefaults instantiates a new V1OrthologSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrthologSetId

`func (o *V1OrthologSet) GetOrthologSetId() int32`

GetOrthologSetId returns the OrthologSetId field if non-nil, zero value otherwise.

### GetOrthologSetIdOk

`func (o *V1OrthologSet) GetOrthologSetIdOk() (*int32, bool)`

GetOrthologSetIdOk returns a tuple with the OrthologSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrthologSetId

`func (o *V1OrthologSet) SetOrthologSetId(v int32)`

SetOrthologSetId sets OrthologSetId field to given value.

### HasOrthologSetId

`func (o *V1OrthologSet) HasOrthologSetId() bool`

HasOrthologSetId returns a boolean if a field has been set.

### GetMethod

`func (o *V1OrthologSet) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *V1OrthologSet) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *V1OrthologSet) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *V1OrthologSet) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetGenes

`func (o *V1OrthologSet) GetGenes() V1GeneMetadata`

GetGenes returns the Genes field if non-nil, zero value otherwise.

### GetGenesOk

`func (o *V1OrthologSet) GetGenesOk() (*V1GeneMetadata, bool)`

GetGenesOk returns a tuple with the Genes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenes

`func (o *V1OrthologSet) SetGenes(v V1GeneMetadata)`

SetGenes sets Genes field to given value.

### HasGenes

`func (o *V1OrthologSet) HasGenes() bool`

HasGenes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


