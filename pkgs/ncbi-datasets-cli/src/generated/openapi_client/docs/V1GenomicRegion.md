# V1GenomicRegion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GeneRange** | Pointer to [**V1SeqRangeSet**](V1SeqRangeSet.md) |  | [optional] 
**Type** | Pointer to [**V1GenomicRegionGenomicRegionType**](V1GenomicRegionGenomicRegionType.md) |  | [optional] [default to V1GENOMICREGIONGENOMICREGIONTYPE_UNKNOWN]

## Methods

### NewV1GenomicRegion

`func NewV1GenomicRegion() *V1GenomicRegion`

NewV1GenomicRegion instantiates a new V1GenomicRegion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1GenomicRegionWithDefaults

`func NewV1GenomicRegionWithDefaults() *V1GenomicRegion`

NewV1GenomicRegionWithDefaults instantiates a new V1GenomicRegion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeneRange

`func (o *V1GenomicRegion) GetGeneRange() V1SeqRangeSet`

GetGeneRange returns the GeneRange field if non-nil, zero value otherwise.

### GetGeneRangeOk

`func (o *V1GenomicRegion) GetGeneRangeOk() (*V1SeqRangeSet, bool)`

GetGeneRangeOk returns a tuple with the GeneRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneRange

`func (o *V1GenomicRegion) SetGeneRange(v V1SeqRangeSet)`

SetGeneRange sets GeneRange field to given value.

### HasGeneRange

`func (o *V1GenomicRegion) HasGeneRange() bool`

HasGeneRange returns a boolean if a field has been set.

### GetType

`func (o *V1GenomicRegion) GetType() V1GenomicRegionGenomicRegionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1GenomicRegion) GetTypeOk() (*V1GenomicRegionGenomicRegionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1GenomicRegion) SetType(v V1GenomicRegionGenomicRegionType)`

SetType sets Type field to given value.

### HasType

`func (o *V1GenomicRegion) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


