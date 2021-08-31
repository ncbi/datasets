# V1SeqRangeSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessionVersion** | Pointer to **string** |  | [optional] 
**Range** | Pointer to [**[]V1Range**](V1Range.md) |  | [optional] 

## Methods

### NewV1SeqRangeSet

`func NewV1SeqRangeSet() *V1SeqRangeSet`

NewV1SeqRangeSet instantiates a new V1SeqRangeSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SeqRangeSetWithDefaults

`func NewV1SeqRangeSetWithDefaults() *V1SeqRangeSet`

NewV1SeqRangeSetWithDefaults instantiates a new V1SeqRangeSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessionVersion

`func (o *V1SeqRangeSet) GetAccessionVersion() string`

GetAccessionVersion returns the AccessionVersion field if non-nil, zero value otherwise.

### GetAccessionVersionOk

`func (o *V1SeqRangeSet) GetAccessionVersionOk() (*string, bool)`

GetAccessionVersionOk returns a tuple with the AccessionVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessionVersion

`func (o *V1SeqRangeSet) SetAccessionVersion(v string)`

SetAccessionVersion sets AccessionVersion field to given value.

### HasAccessionVersion

`func (o *V1SeqRangeSet) HasAccessionVersion() bool`

HasAccessionVersion returns a boolean if a field has been set.

### GetRange

`func (o *V1SeqRangeSet) GetRange() []V1Range`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *V1SeqRangeSet) GetRangeOk() (*[]V1Range, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *V1SeqRangeSet) SetRange(v []V1Range)`

SetRange sets Range field to given value.

### HasRange

`func (o *V1SeqRangeSet) HasRange() bool`

HasRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


