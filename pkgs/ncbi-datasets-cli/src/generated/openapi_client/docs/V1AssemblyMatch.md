# V1AssemblyMatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | Pointer to [**[]V1Message**](V1Message.md) |  | [optional] 
**Assembly** | Pointer to [**V1AssemblyDatasetDescriptor**](V1AssemblyDatasetDescriptor.md) |  | [optional] 

## Methods

### NewV1AssemblyMatch

`func NewV1AssemblyMatch() *V1AssemblyMatch`

NewV1AssemblyMatch instantiates a new V1AssemblyMatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1AssemblyMatchWithDefaults

`func NewV1AssemblyMatchWithDefaults() *V1AssemblyMatch`

NewV1AssemblyMatchWithDefaults instantiates a new V1AssemblyMatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *V1AssemblyMatch) GetMessages() []V1Message`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *V1AssemblyMatch) GetMessagesOk() (*[]V1Message, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *V1AssemblyMatch) SetMessages(v []V1Message)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *V1AssemblyMatch) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetAssembly

`func (o *V1AssemblyMatch) GetAssembly() V1AssemblyDatasetDescriptor`

GetAssembly returns the Assembly field if non-nil, zero value otherwise.

### GetAssemblyOk

`func (o *V1AssemblyMatch) GetAssemblyOk() (*V1AssemblyDatasetDescriptor, bool)`

GetAssemblyOk returns a tuple with the Assembly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssembly

`func (o *V1AssemblyMatch) SetAssembly(v V1AssemblyDatasetDescriptor)`

SetAssembly sets Assembly field to given value.

### HasAssembly

`func (o *V1AssemblyMatch) HasAssembly() bool`

HasAssembly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


