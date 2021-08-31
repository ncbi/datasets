# V1reportsRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Begin** | Pointer to **string** |  | [optional] 
**End** | Pointer to **string** |  | [optional] 
**Orientation** | Pointer to [**V1reportsOrientation**](V1reportsOrientation.md) |  | [optional] [default to V1REPORTSORIENTATION_NONE]
**Order** | Pointer to **int32** |  | [optional] 

## Methods

### NewV1reportsRange

`func NewV1reportsRange() *V1reportsRange`

NewV1reportsRange instantiates a new V1reportsRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1reportsRangeWithDefaults

`func NewV1reportsRangeWithDefaults() *V1reportsRange`

NewV1reportsRangeWithDefaults instantiates a new V1reportsRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBegin

`func (o *V1reportsRange) GetBegin() string`

GetBegin returns the Begin field if non-nil, zero value otherwise.

### GetBeginOk

`func (o *V1reportsRange) GetBeginOk() (*string, bool)`

GetBeginOk returns a tuple with the Begin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBegin

`func (o *V1reportsRange) SetBegin(v string)`

SetBegin sets Begin field to given value.

### HasBegin

`func (o *V1reportsRange) HasBegin() bool`

HasBegin returns a boolean if a field has been set.

### GetEnd

`func (o *V1reportsRange) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *V1reportsRange) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *V1reportsRange) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *V1reportsRange) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetOrientation

`func (o *V1reportsRange) GetOrientation() V1reportsOrientation`

GetOrientation returns the Orientation field if non-nil, zero value otherwise.

### GetOrientationOk

`func (o *V1reportsRange) GetOrientationOk() (*V1reportsOrientation, bool)`

GetOrientationOk returns a tuple with the Orientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrientation

`func (o *V1reportsRange) SetOrientation(v V1reportsOrientation)`

SetOrientation sets Orientation field to given value.

### HasOrientation

`func (o *V1reportsRange) HasOrientation() bool`

HasOrientation returns a boolean if a field has been set.

### GetOrder

`func (o *V1reportsRange) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *V1reportsRange) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *V1reportsRange) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *V1reportsRange) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


