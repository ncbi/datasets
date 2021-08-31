# V1TaxonomyFilteredSubtreeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Taxons** | Pointer to **[]string** |  | [optional] 
**SpecifiedLimit** | Pointer to **bool** |  | [optional] 
**RankLimits** | Pointer to [**[]V1OrganismRankType**](V1OrganismRankType.md) |  | [optional] 

## Methods

### NewV1TaxonomyFilteredSubtreeRequest

`func NewV1TaxonomyFilteredSubtreeRequest() *V1TaxonomyFilteredSubtreeRequest`

NewV1TaxonomyFilteredSubtreeRequest instantiates a new V1TaxonomyFilteredSubtreeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TaxonomyFilteredSubtreeRequestWithDefaults

`func NewV1TaxonomyFilteredSubtreeRequestWithDefaults() *V1TaxonomyFilteredSubtreeRequest`

NewV1TaxonomyFilteredSubtreeRequestWithDefaults instantiates a new V1TaxonomyFilteredSubtreeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaxons

`func (o *V1TaxonomyFilteredSubtreeRequest) GetTaxons() []string`

GetTaxons returns the Taxons field if non-nil, zero value otherwise.

### GetTaxonsOk

`func (o *V1TaxonomyFilteredSubtreeRequest) GetTaxonsOk() (*[]string, bool)`

GetTaxonsOk returns a tuple with the Taxons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxons

`func (o *V1TaxonomyFilteredSubtreeRequest) SetTaxons(v []string)`

SetTaxons sets Taxons field to given value.

### HasTaxons

`func (o *V1TaxonomyFilteredSubtreeRequest) HasTaxons() bool`

HasTaxons returns a boolean if a field has been set.

### GetSpecifiedLimit

`func (o *V1TaxonomyFilteredSubtreeRequest) GetSpecifiedLimit() bool`

GetSpecifiedLimit returns the SpecifiedLimit field if non-nil, zero value otherwise.

### GetSpecifiedLimitOk

`func (o *V1TaxonomyFilteredSubtreeRequest) GetSpecifiedLimitOk() (*bool, bool)`

GetSpecifiedLimitOk returns a tuple with the SpecifiedLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecifiedLimit

`func (o *V1TaxonomyFilteredSubtreeRequest) SetSpecifiedLimit(v bool)`

SetSpecifiedLimit sets SpecifiedLimit field to given value.

### HasSpecifiedLimit

`func (o *V1TaxonomyFilteredSubtreeRequest) HasSpecifiedLimit() bool`

HasSpecifiedLimit returns a boolean if a field has been set.

### GetRankLimits

`func (o *V1TaxonomyFilteredSubtreeRequest) GetRankLimits() []V1OrganismRankType`

GetRankLimits returns the RankLimits field if non-nil, zero value otherwise.

### GetRankLimitsOk

`func (o *V1TaxonomyFilteredSubtreeRequest) GetRankLimitsOk() (*[]V1OrganismRankType, bool)`

GetRankLimitsOk returns a tuple with the RankLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRankLimits

`func (o *V1TaxonomyFilteredSubtreeRequest) SetRankLimits(v []V1OrganismRankType)`

SetRankLimits sets RankLimits field to given value.

### HasRankLimits

`func (o *V1TaxonomyFilteredSubtreeRequest) HasRankLimits() bool`

HasRankLimits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


