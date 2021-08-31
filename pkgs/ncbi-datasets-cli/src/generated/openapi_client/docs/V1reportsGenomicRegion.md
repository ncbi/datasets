# V1reportsGenomicRegion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GeneRange** | Pointer to [**V1reportsSeqRangeSet**](V1reportsSeqRangeSet.md) |  | [optional] 
**Type** | Pointer to [**V1reportsGenomicRegionGenomicRegionType**](V1reportsGenomicRegionGenomicRegionType.md) |  | [optional] [default to V1REPORTSGENOMICREGIONGENOMICREGIONTYPE_UNKNOWN]

## Methods

### NewV1reportsGenomicRegion

`func NewV1reportsGenomicRegion() *V1reportsGenomicRegion`

NewV1reportsGenomicRegion instantiates a new V1reportsGenomicRegion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1reportsGenomicRegionWithDefaults

`func NewV1reportsGenomicRegionWithDefaults() *V1reportsGenomicRegion`

NewV1reportsGenomicRegionWithDefaults instantiates a new V1reportsGenomicRegion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeneRange

`func (o *V1reportsGenomicRegion) GetGeneRange() V1reportsSeqRangeSet`

GetGeneRange returns the GeneRange field if non-nil, zero value otherwise.

### GetGeneRangeOk

`func (o *V1reportsGenomicRegion) GetGeneRangeOk() (*V1reportsSeqRangeSet, bool)`

GetGeneRangeOk returns a tuple with the GeneRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneRange

`func (o *V1reportsGenomicRegion) SetGeneRange(v V1reportsSeqRangeSet)`

SetGeneRange sets GeneRange field to given value.

### HasGeneRange

`func (o *V1reportsGenomicRegion) HasGeneRange() bool`

HasGeneRange returns a boolean if a field has been set.

### GetType

`func (o *V1reportsGenomicRegion) GetType() V1reportsGenomicRegionGenomicRegionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1reportsGenomicRegion) GetTypeOk() (*V1reportsGenomicRegionGenomicRegionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1reportsGenomicRegion) SetType(v V1reportsGenomicRegionGenomicRegionType)`

SetType sets Type field to given value.

### HasType

`func (o *V1reportsGenomicRegion) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


