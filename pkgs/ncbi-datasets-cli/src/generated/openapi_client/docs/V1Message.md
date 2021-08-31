# V1Message

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**V1Error**](V1Error.md) |  | [optional] 
**Warning** | Pointer to [**V1Warning**](V1Warning.md) |  | [optional] 

## Methods

### NewV1Message

`func NewV1Message() *V1Message`

NewV1Message instantiates a new V1Message object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1MessageWithDefaults

`func NewV1MessageWithDefaults() *V1Message`

NewV1MessageWithDefaults instantiates a new V1Message object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *V1Message) GetError() V1Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *V1Message) GetErrorOk() (*V1Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *V1Message) SetError(v V1Error)`

SetError sets Error field to given value.

### HasError

`func (o *V1Message) HasError() bool`

HasError returns a boolean if a field has been set.

### GetWarning

`func (o *V1Message) GetWarning() V1Warning`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *V1Message) GetWarningOk() (*V1Warning, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *V1Message) SetWarning(v V1Warning)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *V1Message) HasWarning() bool`

HasWarning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


