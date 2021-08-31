# V1GenomicLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GenomicAccessionVersion** | Pointer to **string** |  | [optional] 
**SequenceName** | Pointer to **string** |  | [optional] 
**GenomicRange** | Pointer to [**V1Range**](V1Range.md) |  | [optional] 
**Exons** | Pointer to [**[]V1Range**](V1Range.md) |  | [optional] 

## Methods

### NewV1GenomicLocation

`func NewV1GenomicLocation() *V1GenomicLocation`

NewV1GenomicLocation instantiates a new V1GenomicLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1GenomicLocationWithDefaults

`func NewV1GenomicLocationWithDefaults() *V1GenomicLocation`

NewV1GenomicLocationWithDefaults instantiates a new V1GenomicLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGenomicAccessionVersion

`func (o *V1GenomicLocation) GetGenomicAccessionVersion() string`

GetGenomicAccessionVersion returns the GenomicAccessionVersion field if non-nil, zero value otherwise.

### GetGenomicAccessionVersionOk

`func (o *V1GenomicLocation) GetGenomicAccessionVersionOk() (*string, bool)`

GetGenomicAccessionVersionOk returns a tuple with the GenomicAccessionVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenomicAccessionVersion

`func (o *V1GenomicLocation) SetGenomicAccessionVersion(v string)`

SetGenomicAccessionVersion sets GenomicAccessionVersion field to given value.

### HasGenomicAccessionVersion

`func (o *V1GenomicLocation) HasGenomicAccessionVersion() bool`

HasGenomicAccessionVersion returns a boolean if a field has been set.

### GetSequenceName

`func (o *V1GenomicLocation) GetSequenceName() string`

GetSequenceName returns the SequenceName field if non-nil, zero value otherwise.

### GetSequenceNameOk

`func (o *V1GenomicLocation) GetSequenceNameOk() (*string, bool)`

GetSequenceNameOk returns a tuple with the SequenceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceName

`func (o *V1GenomicLocation) SetSequenceName(v string)`

SetSequenceName sets SequenceName field to given value.

### HasSequenceName

`func (o *V1GenomicLocation) HasSequenceName() bool`

HasSequenceName returns a boolean if a field has been set.

### GetGenomicRange

`func (o *V1GenomicLocation) GetGenomicRange() V1Range`

GetGenomicRange returns the GenomicRange field if non-nil, zero value otherwise.

### GetGenomicRangeOk

`func (o *V1GenomicLocation) GetGenomicRangeOk() (*V1Range, bool)`

GetGenomicRangeOk returns a tuple with the GenomicRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenomicRange

`func (o *V1GenomicLocation) SetGenomicRange(v V1Range)`

SetGenomicRange sets GenomicRange field to given value.

### HasGenomicRange

`func (o *V1GenomicLocation) HasGenomicRange() bool`

HasGenomicRange returns a boolean if a field has been set.

### GetExons

`func (o *V1GenomicLocation) GetExons() []V1Range`

GetExons returns the Exons field if non-nil, zero value otherwise.

### GetExonsOk

`func (o *V1GenomicLocation) GetExonsOk() (*[]V1Range, bool)`

GetExonsOk returns a tuple with the Exons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExons

`func (o *V1GenomicLocation) SetExons(v []V1Range)`

SetExons sets Exons field to given value.

### HasExons

`func (o *V1GenomicLocation) HasExons() bool`

HasExons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


