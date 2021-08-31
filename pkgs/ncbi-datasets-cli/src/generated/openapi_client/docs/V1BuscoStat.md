# V1BuscoStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuscoLineage** | Pointer to **string** |  | [optional] 
**BuscoVer** | Pointer to **string** |  | [optional] 
**Complete** | Pointer to **float32** |  | [optional] 
**SingleCopy** | Pointer to **float32** |  | [optional] 
**Duplicated** | Pointer to **float32** |  | [optional] 
**Fragmented** | Pointer to **float32** |  | [optional] 
**Missing** | Pointer to **float32** |  | [optional] 
**TotalCount** | Pointer to **string** |  | [optional] 

## Methods

### NewV1BuscoStat

`func NewV1BuscoStat() *V1BuscoStat`

NewV1BuscoStat instantiates a new V1BuscoStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1BuscoStatWithDefaults

`func NewV1BuscoStatWithDefaults() *V1BuscoStat`

NewV1BuscoStatWithDefaults instantiates a new V1BuscoStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuscoLineage

`func (o *V1BuscoStat) GetBuscoLineage() string`

GetBuscoLineage returns the BuscoLineage field if non-nil, zero value otherwise.

### GetBuscoLineageOk

`func (o *V1BuscoStat) GetBuscoLineageOk() (*string, bool)`

GetBuscoLineageOk returns a tuple with the BuscoLineage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuscoLineage

`func (o *V1BuscoStat) SetBuscoLineage(v string)`

SetBuscoLineage sets BuscoLineage field to given value.

### HasBuscoLineage

`func (o *V1BuscoStat) HasBuscoLineage() bool`

HasBuscoLineage returns a boolean if a field has been set.

### GetBuscoVer

`func (o *V1BuscoStat) GetBuscoVer() string`

GetBuscoVer returns the BuscoVer field if non-nil, zero value otherwise.

### GetBuscoVerOk

`func (o *V1BuscoStat) GetBuscoVerOk() (*string, bool)`

GetBuscoVerOk returns a tuple with the BuscoVer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuscoVer

`func (o *V1BuscoStat) SetBuscoVer(v string)`

SetBuscoVer sets BuscoVer field to given value.

### HasBuscoVer

`func (o *V1BuscoStat) HasBuscoVer() bool`

HasBuscoVer returns a boolean if a field has been set.

### GetComplete

`func (o *V1BuscoStat) GetComplete() float32`

GetComplete returns the Complete field if non-nil, zero value otherwise.

### GetCompleteOk

`func (o *V1BuscoStat) GetCompleteOk() (*float32, bool)`

GetCompleteOk returns a tuple with the Complete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplete

`func (o *V1BuscoStat) SetComplete(v float32)`

SetComplete sets Complete field to given value.

### HasComplete

`func (o *V1BuscoStat) HasComplete() bool`

HasComplete returns a boolean if a field has been set.

### GetSingleCopy

`func (o *V1BuscoStat) GetSingleCopy() float32`

GetSingleCopy returns the SingleCopy field if non-nil, zero value otherwise.

### GetSingleCopyOk

`func (o *V1BuscoStat) GetSingleCopyOk() (*float32, bool)`

GetSingleCopyOk returns a tuple with the SingleCopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleCopy

`func (o *V1BuscoStat) SetSingleCopy(v float32)`

SetSingleCopy sets SingleCopy field to given value.

### HasSingleCopy

`func (o *V1BuscoStat) HasSingleCopy() bool`

HasSingleCopy returns a boolean if a field has been set.

### GetDuplicated

`func (o *V1BuscoStat) GetDuplicated() float32`

GetDuplicated returns the Duplicated field if non-nil, zero value otherwise.

### GetDuplicatedOk

`func (o *V1BuscoStat) GetDuplicatedOk() (*float32, bool)`

GetDuplicatedOk returns a tuple with the Duplicated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicated

`func (o *V1BuscoStat) SetDuplicated(v float32)`

SetDuplicated sets Duplicated field to given value.

### HasDuplicated

`func (o *V1BuscoStat) HasDuplicated() bool`

HasDuplicated returns a boolean if a field has been set.

### GetFragmented

`func (o *V1BuscoStat) GetFragmented() float32`

GetFragmented returns the Fragmented field if non-nil, zero value otherwise.

### GetFragmentedOk

`func (o *V1BuscoStat) GetFragmentedOk() (*float32, bool)`

GetFragmentedOk returns a tuple with the Fragmented field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFragmented

`func (o *V1BuscoStat) SetFragmented(v float32)`

SetFragmented sets Fragmented field to given value.

### HasFragmented

`func (o *V1BuscoStat) HasFragmented() bool`

HasFragmented returns a boolean if a field has been set.

### GetMissing

`func (o *V1BuscoStat) GetMissing() float32`

GetMissing returns the Missing field if non-nil, zero value otherwise.

### GetMissingOk

`func (o *V1BuscoStat) GetMissingOk() (*float32, bool)`

GetMissingOk returns a tuple with the Missing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissing

`func (o *V1BuscoStat) SetMissing(v float32)`

SetMissing sets Missing field to given value.

### HasMissing

`func (o *V1BuscoStat) HasMissing() bool`

HasMissing returns a boolean if a field has been set.

### GetTotalCount

`func (o *V1BuscoStat) GetTotalCount() string`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *V1BuscoStat) GetTotalCountOk() (*string, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *V1BuscoStat) SetTotalCount(v string)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *V1BuscoStat) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


