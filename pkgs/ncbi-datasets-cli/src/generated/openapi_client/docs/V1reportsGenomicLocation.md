# V1reportsGenomicLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GenomicAccessionVersion** | Pointer to **string** |  | [optional] 
**SequenceName** | Pointer to **string** |  | [optional] 
**GenomicRange** | Pointer to [**V1reportsRange**](V1reportsRange.md) |  | [optional] 
**Exons** | Pointer to [**[]V1reportsRange**](V1reportsRange.md) |  | [optional] 

## Methods

### NewV1reportsGenomicLocation

`func NewV1reportsGenomicLocation() *V1reportsGenomicLocation`

NewV1reportsGenomicLocation instantiates a new V1reportsGenomicLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1reportsGenomicLocationWithDefaults

`func NewV1reportsGenomicLocationWithDefaults() *V1reportsGenomicLocation`

NewV1reportsGenomicLocationWithDefaults instantiates a new V1reportsGenomicLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGenomicAccessionVersion

`func (o *V1reportsGenomicLocation) GetGenomicAccessionVersion() string`

GetGenomicAccessionVersion returns the GenomicAccessionVersion field if non-nil, zero value otherwise.

### GetGenomicAccessionVersionOk

`func (o *V1reportsGenomicLocation) GetGenomicAccessionVersionOk() (*string, bool)`

GetGenomicAccessionVersionOk returns a tuple with the GenomicAccessionVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenomicAccessionVersion

`func (o *V1reportsGenomicLocation) SetGenomicAccessionVersion(v string)`

SetGenomicAccessionVersion sets GenomicAccessionVersion field to given value.

### HasGenomicAccessionVersion

`func (o *V1reportsGenomicLocation) HasGenomicAccessionVersion() bool`

HasGenomicAccessionVersion returns a boolean if a field has been set.

### GetSequenceName

`func (o *V1reportsGenomicLocation) GetSequenceName() string`

GetSequenceName returns the SequenceName field if non-nil, zero value otherwise.

### GetSequenceNameOk

`func (o *V1reportsGenomicLocation) GetSequenceNameOk() (*string, bool)`

GetSequenceNameOk returns a tuple with the SequenceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceName

`func (o *V1reportsGenomicLocation) SetSequenceName(v string)`

SetSequenceName sets SequenceName field to given value.

### HasSequenceName

`func (o *V1reportsGenomicLocation) HasSequenceName() bool`

HasSequenceName returns a boolean if a field has been set.

### GetGenomicRange

`func (o *V1reportsGenomicLocation) GetGenomicRange() V1reportsRange`

GetGenomicRange returns the GenomicRange field if non-nil, zero value otherwise.

### GetGenomicRangeOk

`func (o *V1reportsGenomicLocation) GetGenomicRangeOk() (*V1reportsRange, bool)`

GetGenomicRangeOk returns a tuple with the GenomicRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenomicRange

`func (o *V1reportsGenomicLocation) SetGenomicRange(v V1reportsRange)`

SetGenomicRange sets GenomicRange field to given value.

### HasGenomicRange

`func (o *V1reportsGenomicLocation) HasGenomicRange() bool`

HasGenomicRange returns a boolean if a field has been set.

### GetExons

`func (o *V1reportsGenomicLocation) GetExons() []V1reportsRange`

GetExons returns the Exons field if non-nil, zero value otherwise.

### GetExonsOk

`func (o *V1reportsGenomicLocation) GetExonsOk() (*[]V1reportsRange, bool)`

GetExonsOk returns a tuple with the Exons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExons

`func (o *V1reportsGenomicLocation) SetExons(v []V1reportsRange)`

SetExons sets Exons field to given value.

### HasExons

`func (o *V1reportsGenomicLocation) HasExons() bool`

HasExons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


