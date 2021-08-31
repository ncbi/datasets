# V1Range

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Begin** | Pointer to **string** |  | [optional] 
**End** | Pointer to **string** |  | [optional] 
**Orientation** | Pointer to [**V1Orientation**](V1Orientation.md) |  | [optional] [default to V1ORIENTATION_NONE]
**Order** | Pointer to **int32** |  | [optional] 
**RibosomalSlippage** | Pointer to **int32** |  | [optional] 

## Methods

### NewV1Range

`func NewV1Range() *V1Range`

NewV1Range instantiates a new V1Range object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1RangeWithDefaults

`func NewV1RangeWithDefaults() *V1Range`

NewV1RangeWithDefaults instantiates a new V1Range object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBegin

`func (o *V1Range) GetBegin() string`

GetBegin returns the Begin field if non-nil, zero value otherwise.

### GetBeginOk

`func (o *V1Range) GetBeginOk() (*string, bool)`

GetBeginOk returns a tuple with the Begin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBegin

`func (o *V1Range) SetBegin(v string)`

SetBegin sets Begin field to given value.

### HasBegin

`func (o *V1Range) HasBegin() bool`

HasBegin returns a boolean if a field has been set.

### GetEnd

`func (o *V1Range) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *V1Range) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *V1Range) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *V1Range) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetOrientation

`func (o *V1Range) GetOrientation() V1Orientation`

GetOrientation returns the Orientation field if non-nil, zero value otherwise.

### GetOrientationOk

`func (o *V1Range) GetOrientationOk() (*V1Orientation, bool)`

GetOrientationOk returns a tuple with the Orientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrientation

`func (o *V1Range) SetOrientation(v V1Orientation)`

SetOrientation sets Orientation field to given value.

### HasOrientation

`func (o *V1Range) HasOrientation() bool`

HasOrientation returns a boolean if a field has been set.

### GetOrder

`func (o *V1Range) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *V1Range) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *V1Range) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *V1Range) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetRibosomalSlippage

`func (o *V1Range) GetRibosomalSlippage() int32`

GetRibosomalSlippage returns the RibosomalSlippage field if non-nil, zero value otherwise.

### GetRibosomalSlippageOk

`func (o *V1Range) GetRibosomalSlippageOk() (*int32, bool)`

GetRibosomalSlippageOk returns a tuple with the RibosomalSlippage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRibosomalSlippage

`func (o *V1Range) SetRibosomalSlippage(v int32)`

SetRibosomalSlippage sets RibosomalSlippage field to given value.

### HasRibosomalSlippage

`func (o *V1Range) HasRibosomalSlippage() bool`

HasRibosomalSlippage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


