# V1OrganismCountByType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**V1CountType**](V1CountType.md) |  | [optional] [default to V1COUNTTYPE_UNSPECIFIED]
**Counts** | Pointer to [**V1OrganismCounts**](V1OrganismCounts.md) |  | [optional] 

## Methods

### NewV1OrganismCountByType

`func NewV1OrganismCountByType() *V1OrganismCountByType`

NewV1OrganismCountByType instantiates a new V1OrganismCountByType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1OrganismCountByTypeWithDefaults

`func NewV1OrganismCountByTypeWithDefaults() *V1OrganismCountByType`

NewV1OrganismCountByTypeWithDefaults instantiates a new V1OrganismCountByType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *V1OrganismCountByType) GetType() V1CountType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1OrganismCountByType) GetTypeOk() (*V1CountType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1OrganismCountByType) SetType(v V1CountType)`

SetType sets Type field to given value.

### HasType

`func (o *V1OrganismCountByType) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCounts

`func (o *V1OrganismCountByType) GetCounts() V1OrganismCounts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *V1OrganismCountByType) GetCountsOk() (*V1OrganismCounts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *V1OrganismCountByType) SetCounts(v V1OrganismCounts)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *V1OrganismCountByType) HasCounts() bool`

HasCounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


