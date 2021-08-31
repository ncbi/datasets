# V1AssemblyMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | Pointer to [**[]V1Message**](V1Message.md) |  | [optional] 
**Assemblies** | Pointer to [**[]V1AssemblyMatch**](V1AssemblyMatch.md) |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 
**NextPageToken** | Pointer to **string** |  | [optional] 

## Methods

### NewV1AssemblyMetadata

`func NewV1AssemblyMetadata() *V1AssemblyMetadata`

NewV1AssemblyMetadata instantiates a new V1AssemblyMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1AssemblyMetadataWithDefaults

`func NewV1AssemblyMetadataWithDefaults() *V1AssemblyMetadata`

NewV1AssemblyMetadataWithDefaults instantiates a new V1AssemblyMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *V1AssemblyMetadata) GetMessages() []V1Message`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *V1AssemblyMetadata) GetMessagesOk() (*[]V1Message, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *V1AssemblyMetadata) SetMessages(v []V1Message)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *V1AssemblyMetadata) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetAssemblies

`func (o *V1AssemblyMetadata) GetAssemblies() []V1AssemblyMatch`

GetAssemblies returns the Assemblies field if non-nil, zero value otherwise.

### GetAssembliesOk

`func (o *V1AssemblyMetadata) GetAssembliesOk() (*[]V1AssemblyMatch, bool)`

GetAssembliesOk returns a tuple with the Assemblies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssemblies

`func (o *V1AssemblyMetadata) SetAssemblies(v []V1AssemblyMatch)`

SetAssemblies sets Assemblies field to given value.

### HasAssemblies

`func (o *V1AssemblyMetadata) HasAssemblies() bool`

HasAssemblies returns a boolean if a field has been set.

### GetTotalCount

`func (o *V1AssemblyMetadata) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *V1AssemblyMetadata) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *V1AssemblyMetadata) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *V1AssemblyMetadata) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetNextPageToken

`func (o *V1AssemblyMetadata) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *V1AssemblyMetadata) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *V1AssemblyMetadata) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *V1AssemblyMetadata) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


