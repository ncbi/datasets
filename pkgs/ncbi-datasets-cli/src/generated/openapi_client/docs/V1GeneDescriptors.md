# V1GeneDescriptors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Genes** | Pointer to [**[]V1GeneDescriptor**](V1GeneDescriptor.md) |  | [optional] 
**Errors** | Pointer to [**[]V1Error**](V1Error.md) |  | [optional] 

## Methods

### NewV1GeneDescriptors

`func NewV1GeneDescriptors() *V1GeneDescriptors`

NewV1GeneDescriptors instantiates a new V1GeneDescriptors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1GeneDescriptorsWithDefaults

`func NewV1GeneDescriptorsWithDefaults() *V1GeneDescriptors`

NewV1GeneDescriptorsWithDefaults instantiates a new V1GeneDescriptors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGenes

`func (o *V1GeneDescriptors) GetGenes() []V1GeneDescriptor`

GetGenes returns the Genes field if non-nil, zero value otherwise.

### GetGenesOk

`func (o *V1GeneDescriptors) GetGenesOk() (*[]V1GeneDescriptor, bool)`

GetGenesOk returns a tuple with the Genes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenes

`func (o *V1GeneDescriptors) SetGenes(v []V1GeneDescriptor)`

SetGenes sets Genes field to given value.

### HasGenes

`func (o *V1GeneDescriptors) HasGenes() bool`

HasGenes returns a boolean if a field has been set.

### GetErrors

`func (o *V1GeneDescriptors) GetErrors() []V1Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V1GeneDescriptors) GetErrorsOk() (*[]V1Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V1GeneDescriptors) SetErrors(v []V1Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V1GeneDescriptors) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


