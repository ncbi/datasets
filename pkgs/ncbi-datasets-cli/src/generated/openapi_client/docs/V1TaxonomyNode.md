# V1TaxonomyNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaxId** | Pointer to **int32** |  | [optional] 
**OrganismName** | Pointer to **string** |  | [optional] 
**CommonName** | Pointer to **string** |  | [optional] 
**Lineage** | Pointer to **[]int32** |  | [optional] 
**Children** | Pointer to **[]int32** |  | [optional] 
**DescendentWithDescribedSpeciesNamesCount** | Pointer to **int32** |  | [optional] 
**Rank** | Pointer to [**V1OrganismRankType**](V1OrganismRankType.md) |  | [optional] [default to V1ORGANISMRANKTYPE_NO_RANK]
**HasDescribedSpeciesName** | Pointer to **bool** |  | [optional] 
**Counts** | Pointer to [**[]V1TaxonomyNodeCountByType**](V1TaxonomyNodeCountByType.md) |  | [optional] 
**MinOrd** | Pointer to **int32** |  | [optional] 
**MaxOrd** | Pointer to **int32** |  | [optional] 
**IndexedNames** | Pointer to **[]string** |  | [optional] 

## Methods

### NewV1TaxonomyNode

`func NewV1TaxonomyNode() *V1TaxonomyNode`

NewV1TaxonomyNode instantiates a new V1TaxonomyNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TaxonomyNodeWithDefaults

`func NewV1TaxonomyNodeWithDefaults() *V1TaxonomyNode`

NewV1TaxonomyNodeWithDefaults instantiates a new V1TaxonomyNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaxId

`func (o *V1TaxonomyNode) GetTaxId() int32`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *V1TaxonomyNode) GetTaxIdOk() (*int32, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *V1TaxonomyNode) SetTaxId(v int32)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *V1TaxonomyNode) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.

### GetOrganismName

`func (o *V1TaxonomyNode) GetOrganismName() string`

GetOrganismName returns the OrganismName field if non-nil, zero value otherwise.

### GetOrganismNameOk

`func (o *V1TaxonomyNode) GetOrganismNameOk() (*string, bool)`

GetOrganismNameOk returns a tuple with the OrganismName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganismName

`func (o *V1TaxonomyNode) SetOrganismName(v string)`

SetOrganismName sets OrganismName field to given value.

### HasOrganismName

`func (o *V1TaxonomyNode) HasOrganismName() bool`

HasOrganismName returns a boolean if a field has been set.

### GetCommonName

`func (o *V1TaxonomyNode) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *V1TaxonomyNode) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *V1TaxonomyNode) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *V1TaxonomyNode) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### GetLineage

`func (o *V1TaxonomyNode) GetLineage() []int32`

GetLineage returns the Lineage field if non-nil, zero value otherwise.

### GetLineageOk

`func (o *V1TaxonomyNode) GetLineageOk() (*[]int32, bool)`

GetLineageOk returns a tuple with the Lineage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineage

`func (o *V1TaxonomyNode) SetLineage(v []int32)`

SetLineage sets Lineage field to given value.

### HasLineage

`func (o *V1TaxonomyNode) HasLineage() bool`

HasLineage returns a boolean if a field has been set.

### GetChildren

`func (o *V1TaxonomyNode) GetChildren() []int32`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *V1TaxonomyNode) GetChildrenOk() (*[]int32, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *V1TaxonomyNode) SetChildren(v []int32)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *V1TaxonomyNode) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetDescendentWithDescribedSpeciesNamesCount

`func (o *V1TaxonomyNode) GetDescendentWithDescribedSpeciesNamesCount() int32`

GetDescendentWithDescribedSpeciesNamesCount returns the DescendentWithDescribedSpeciesNamesCount field if non-nil, zero value otherwise.

### GetDescendentWithDescribedSpeciesNamesCountOk

`func (o *V1TaxonomyNode) GetDescendentWithDescribedSpeciesNamesCountOk() (*int32, bool)`

GetDescendentWithDescribedSpeciesNamesCountOk returns a tuple with the DescendentWithDescribedSpeciesNamesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescendentWithDescribedSpeciesNamesCount

`func (o *V1TaxonomyNode) SetDescendentWithDescribedSpeciesNamesCount(v int32)`

SetDescendentWithDescribedSpeciesNamesCount sets DescendentWithDescribedSpeciesNamesCount field to given value.

### HasDescendentWithDescribedSpeciesNamesCount

`func (o *V1TaxonomyNode) HasDescendentWithDescribedSpeciesNamesCount() bool`

HasDescendentWithDescribedSpeciesNamesCount returns a boolean if a field has been set.

### GetRank

`func (o *V1TaxonomyNode) GetRank() V1OrganismRankType`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *V1TaxonomyNode) GetRankOk() (*V1OrganismRankType, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *V1TaxonomyNode) SetRank(v V1OrganismRankType)`

SetRank sets Rank field to given value.

### HasRank

`func (o *V1TaxonomyNode) HasRank() bool`

HasRank returns a boolean if a field has been set.

### GetHasDescribedSpeciesName

`func (o *V1TaxonomyNode) GetHasDescribedSpeciesName() bool`

GetHasDescribedSpeciesName returns the HasDescribedSpeciesName field if non-nil, zero value otherwise.

### GetHasDescribedSpeciesNameOk

`func (o *V1TaxonomyNode) GetHasDescribedSpeciesNameOk() (*bool, bool)`

GetHasDescribedSpeciesNameOk returns a tuple with the HasDescribedSpeciesName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDescribedSpeciesName

`func (o *V1TaxonomyNode) SetHasDescribedSpeciesName(v bool)`

SetHasDescribedSpeciesName sets HasDescribedSpeciesName field to given value.

### HasHasDescribedSpeciesName

`func (o *V1TaxonomyNode) HasHasDescribedSpeciesName() bool`

HasHasDescribedSpeciesName returns a boolean if a field has been set.

### GetCounts

`func (o *V1TaxonomyNode) GetCounts() []V1TaxonomyNodeCountByType`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *V1TaxonomyNode) GetCountsOk() (*[]V1TaxonomyNodeCountByType, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *V1TaxonomyNode) SetCounts(v []V1TaxonomyNodeCountByType)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *V1TaxonomyNode) HasCounts() bool`

HasCounts returns a boolean if a field has been set.

### GetMinOrd

`func (o *V1TaxonomyNode) GetMinOrd() int32`

GetMinOrd returns the MinOrd field if non-nil, zero value otherwise.

### GetMinOrdOk

`func (o *V1TaxonomyNode) GetMinOrdOk() (*int32, bool)`

GetMinOrdOk returns a tuple with the MinOrd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinOrd

`func (o *V1TaxonomyNode) SetMinOrd(v int32)`

SetMinOrd sets MinOrd field to given value.

### HasMinOrd

`func (o *V1TaxonomyNode) HasMinOrd() bool`

HasMinOrd returns a boolean if a field has been set.

### GetMaxOrd

`func (o *V1TaxonomyNode) GetMaxOrd() int32`

GetMaxOrd returns the MaxOrd field if non-nil, zero value otherwise.

### GetMaxOrdOk

`func (o *V1TaxonomyNode) GetMaxOrdOk() (*int32, bool)`

GetMaxOrdOk returns a tuple with the MaxOrd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxOrd

`func (o *V1TaxonomyNode) SetMaxOrd(v int32)`

SetMaxOrd sets MaxOrd field to given value.

### HasMaxOrd

`func (o *V1TaxonomyNode) HasMaxOrd() bool`

HasMaxOrd returns a boolean if a field has been set.

### GetIndexedNames

`func (o *V1TaxonomyNode) GetIndexedNames() []string`

GetIndexedNames returns the IndexedNames field if non-nil, zero value otherwise.

### GetIndexedNamesOk

`func (o *V1TaxonomyNode) GetIndexedNamesOk() (*[]string, bool)`

GetIndexedNamesOk returns a tuple with the IndexedNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexedNames

`func (o *V1TaxonomyNode) SetIndexedNames(v []string)`

SetIndexedNames sets IndexedNames field to given value.

### HasIndexedNames

`func (o *V1TaxonomyNode) HasIndexedNames() bool`

HasIndexedNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


