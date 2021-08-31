# V1SleepRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SleepMsec** | Pointer to **int32** |  | [optional] 
**ErrorRate** | Pointer to **float32** |  | [optional] 

## Methods

### NewV1SleepRequest

`func NewV1SleepRequest() *V1SleepRequest`

NewV1SleepRequest instantiates a new V1SleepRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SleepRequestWithDefaults

`func NewV1SleepRequestWithDefaults() *V1SleepRequest`

NewV1SleepRequestWithDefaults instantiates a new V1SleepRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSleepMsec

`func (o *V1SleepRequest) GetSleepMsec() int32`

GetSleepMsec returns the SleepMsec field if non-nil, zero value otherwise.

### GetSleepMsecOk

`func (o *V1SleepRequest) GetSleepMsecOk() (*int32, bool)`

GetSleepMsecOk returns a tuple with the SleepMsec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepMsec

`func (o *V1SleepRequest) SetSleepMsec(v int32)`

SetSleepMsec sets SleepMsec field to given value.

### HasSleepMsec

`func (o *V1SleepRequest) HasSleepMsec() bool`

HasSleepMsec returns a boolean if a field has been set.

### GetErrorRate

`func (o *V1SleepRequest) GetErrorRate() float32`

GetErrorRate returns the ErrorRate field if non-nil, zero value otherwise.

### GetErrorRateOk

`func (o *V1SleepRequest) GetErrorRateOk() (*float32, bool)`

GetErrorRateOk returns a tuple with the ErrorRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorRate

`func (o *V1SleepRequest) SetErrorRate(v float32)`

SetErrorRate sets ErrorRate field to given value.

### HasErrorRate

`func (o *V1SleepRequest) HasErrorRate() bool`

HasErrorRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


