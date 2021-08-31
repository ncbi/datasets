# V1GeneDatasetRequestSort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | Pointer to [**V1GeneDatasetRequestSortField**](V1GeneDatasetRequestSortField.md) |  | [optional] [default to V1GENEDATASETREQUESTSORTFIELD_UNSPECIFIED]
**Direction** | Pointer to [**V1SortDirection**](V1SortDirection.md) |  | [optional] [default to V1SORTDIRECTION_UNSPECIFIED]

## Methods

### NewV1GeneDatasetRequestSort

`func NewV1GeneDatasetRequestSort() *V1GeneDatasetRequestSort`

NewV1GeneDatasetRequestSort instantiates a new V1GeneDatasetRequestSort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1GeneDatasetRequestSortWithDefaults

`func NewV1GeneDatasetRequestSortWithDefaults() *V1GeneDatasetRequestSort`

NewV1GeneDatasetRequestSortWithDefaults instantiates a new V1GeneDatasetRequestSort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *V1GeneDatasetRequestSort) GetField() V1GeneDatasetRequestSortField`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *V1GeneDatasetRequestSort) GetFieldOk() (*V1GeneDatasetRequestSortField, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *V1GeneDatasetRequestSort) SetField(v V1GeneDatasetRequestSortField)`

SetField sets Field field to given value.

### HasField

`func (o *V1GeneDatasetRequestSort) HasField() bool`

HasField returns a boolean if a field has been set.

### GetDirection

`func (o *V1GeneDatasetRequestSort) GetDirection() V1SortDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *V1GeneDatasetRequestSort) GetDirectionOk() (*V1SortDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *V1GeneDatasetRequestSort) SetDirection(v V1SortDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *V1GeneDatasetRequestSort) HasDirection() bool`

HasDirection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


