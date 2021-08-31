# V1Organism

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaxId** | Pointer to **string** |  | [optional] 
**SciName** | Pointer to **string** |  | [optional] 
**CommonName** | Pointer to **string** |  | [optional] 
**BlastNode** | Pointer to **bool** |  | [optional] 
**Breed** | Pointer to **string** |  | [optional] 
**Cultivar** | Pointer to **string** |  | [optional] 
**Ecotype** | Pointer to **string** |  | [optional] 
**Isolate** | Pointer to **string** |  | [optional] 
**Sex** | Pointer to **string** |  | [optional] 
**Strain** | Pointer to **string** |  | [optional] 
**SearchText** | Pointer to **[]string** |  | [optional] 
**Rank** | Pointer to [**V1OrganismRankType**](V1OrganismRankType.md) |  | [optional] [default to V1ORGANISMRANKTYPE_NO_RANK]
**ParentTaxId** | Pointer to **string** |  | [optional] 
**AssemblyCount** | Pointer to **string** |  | [optional] 
**AssemblyCounts** | Pointer to [**V1OrganismCounts**](V1OrganismCounts.md) |  | [optional] 
**Counts** | Pointer to [**[]V1OrganismCountByType**](V1OrganismCountByType.md) |  | [optional] 
**Children** | Pointer to [**[]V1Organism**](V1Organism.md) |  | [optional] 
**Merged** | Pointer to [**[]V1Organism**](V1Organism.md) |  | [optional] 
**MergedTaxIds** | Pointer to **[]string** |  | [optional] 
**MinOrd** | Pointer to **int32** |  | [optional] 
**MaxOrd** | Pointer to **int32** |  | [optional] 
**Weight** | Pointer to **int32** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Icon** | Pointer to **bool** |  | [optional] 

## Methods

### NewV1Organism

`func NewV1Organism() *V1Organism`

NewV1Organism instantiates a new V1Organism object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1OrganismWithDefaults

`func NewV1OrganismWithDefaults() *V1Organism`

NewV1OrganismWithDefaults instantiates a new V1Organism object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaxId

`func (o *V1Organism) GetTaxId() string`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *V1Organism) GetTaxIdOk() (*string, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *V1Organism) SetTaxId(v string)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *V1Organism) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.

### GetSciName

`func (o *V1Organism) GetSciName() string`

GetSciName returns the SciName field if non-nil, zero value otherwise.

### GetSciNameOk

`func (o *V1Organism) GetSciNameOk() (*string, bool)`

GetSciNameOk returns a tuple with the SciName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSciName

`func (o *V1Organism) SetSciName(v string)`

SetSciName sets SciName field to given value.

### HasSciName

`func (o *V1Organism) HasSciName() bool`

HasSciName returns a boolean if a field has been set.

### GetCommonName

`func (o *V1Organism) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *V1Organism) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *V1Organism) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *V1Organism) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### GetBlastNode

`func (o *V1Organism) GetBlastNode() bool`

GetBlastNode returns the BlastNode field if non-nil, zero value otherwise.

### GetBlastNodeOk

`func (o *V1Organism) GetBlastNodeOk() (*bool, bool)`

GetBlastNodeOk returns a tuple with the BlastNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlastNode

`func (o *V1Organism) SetBlastNode(v bool)`

SetBlastNode sets BlastNode field to given value.

### HasBlastNode

`func (o *V1Organism) HasBlastNode() bool`

HasBlastNode returns a boolean if a field has been set.

### GetBreed

`func (o *V1Organism) GetBreed() string`

GetBreed returns the Breed field if non-nil, zero value otherwise.

### GetBreedOk

`func (o *V1Organism) GetBreedOk() (*string, bool)`

GetBreedOk returns a tuple with the Breed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreed

`func (o *V1Organism) SetBreed(v string)`

SetBreed sets Breed field to given value.

### HasBreed

`func (o *V1Organism) HasBreed() bool`

HasBreed returns a boolean if a field has been set.

### GetCultivar

`func (o *V1Organism) GetCultivar() string`

GetCultivar returns the Cultivar field if non-nil, zero value otherwise.

### GetCultivarOk

`func (o *V1Organism) GetCultivarOk() (*string, bool)`

GetCultivarOk returns a tuple with the Cultivar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCultivar

`func (o *V1Organism) SetCultivar(v string)`

SetCultivar sets Cultivar field to given value.

### HasCultivar

`func (o *V1Organism) HasCultivar() bool`

HasCultivar returns a boolean if a field has been set.

### GetEcotype

`func (o *V1Organism) GetEcotype() string`

GetEcotype returns the Ecotype field if non-nil, zero value otherwise.

### GetEcotypeOk

`func (o *V1Organism) GetEcotypeOk() (*string, bool)`

GetEcotypeOk returns a tuple with the Ecotype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcotype

`func (o *V1Organism) SetEcotype(v string)`

SetEcotype sets Ecotype field to given value.

### HasEcotype

`func (o *V1Organism) HasEcotype() bool`

HasEcotype returns a boolean if a field has been set.

### GetIsolate

`func (o *V1Organism) GetIsolate() string`

GetIsolate returns the Isolate field if non-nil, zero value otherwise.

### GetIsolateOk

`func (o *V1Organism) GetIsolateOk() (*string, bool)`

GetIsolateOk returns a tuple with the Isolate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolate

`func (o *V1Organism) SetIsolate(v string)`

SetIsolate sets Isolate field to given value.

### HasIsolate

`func (o *V1Organism) HasIsolate() bool`

HasIsolate returns a boolean if a field has been set.

### GetSex

`func (o *V1Organism) GetSex() string`

GetSex returns the Sex field if non-nil, zero value otherwise.

### GetSexOk

`func (o *V1Organism) GetSexOk() (*string, bool)`

GetSexOk returns a tuple with the Sex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSex

`func (o *V1Organism) SetSex(v string)`

SetSex sets Sex field to given value.

### HasSex

`func (o *V1Organism) HasSex() bool`

HasSex returns a boolean if a field has been set.

### GetStrain

`func (o *V1Organism) GetStrain() string`

GetStrain returns the Strain field if non-nil, zero value otherwise.

### GetStrainOk

`func (o *V1Organism) GetStrainOk() (*string, bool)`

GetStrainOk returns a tuple with the Strain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrain

`func (o *V1Organism) SetStrain(v string)`

SetStrain sets Strain field to given value.

### HasStrain

`func (o *V1Organism) HasStrain() bool`

HasStrain returns a boolean if a field has been set.

### GetSearchText

`func (o *V1Organism) GetSearchText() []string`

GetSearchText returns the SearchText field if non-nil, zero value otherwise.

### GetSearchTextOk

`func (o *V1Organism) GetSearchTextOk() (*[]string, bool)`

GetSearchTextOk returns a tuple with the SearchText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchText

`func (o *V1Organism) SetSearchText(v []string)`

SetSearchText sets SearchText field to given value.

### HasSearchText

`func (o *V1Organism) HasSearchText() bool`

HasSearchText returns a boolean if a field has been set.

### GetRank

`func (o *V1Organism) GetRank() V1OrganismRankType`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *V1Organism) GetRankOk() (*V1OrganismRankType, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *V1Organism) SetRank(v V1OrganismRankType)`

SetRank sets Rank field to given value.

### HasRank

`func (o *V1Organism) HasRank() bool`

HasRank returns a boolean if a field has been set.

### GetParentTaxId

`func (o *V1Organism) GetParentTaxId() string`

GetParentTaxId returns the ParentTaxId field if non-nil, zero value otherwise.

### GetParentTaxIdOk

`func (o *V1Organism) GetParentTaxIdOk() (*string, bool)`

GetParentTaxIdOk returns a tuple with the ParentTaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTaxId

`func (o *V1Organism) SetParentTaxId(v string)`

SetParentTaxId sets ParentTaxId field to given value.

### HasParentTaxId

`func (o *V1Organism) HasParentTaxId() bool`

HasParentTaxId returns a boolean if a field has been set.

### GetAssemblyCount

`func (o *V1Organism) GetAssemblyCount() string`

GetAssemblyCount returns the AssemblyCount field if non-nil, zero value otherwise.

### GetAssemblyCountOk

`func (o *V1Organism) GetAssemblyCountOk() (*string, bool)`

GetAssemblyCountOk returns a tuple with the AssemblyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssemblyCount

`func (o *V1Organism) SetAssemblyCount(v string)`

SetAssemblyCount sets AssemblyCount field to given value.

### HasAssemblyCount

`func (o *V1Organism) HasAssemblyCount() bool`

HasAssemblyCount returns a boolean if a field has been set.

### GetAssemblyCounts

`func (o *V1Organism) GetAssemblyCounts() V1OrganismCounts`

GetAssemblyCounts returns the AssemblyCounts field if non-nil, zero value otherwise.

### GetAssemblyCountsOk

`func (o *V1Organism) GetAssemblyCountsOk() (*V1OrganismCounts, bool)`

GetAssemblyCountsOk returns a tuple with the AssemblyCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssemblyCounts

`func (o *V1Organism) SetAssemblyCounts(v V1OrganismCounts)`

SetAssemblyCounts sets AssemblyCounts field to given value.

### HasAssemblyCounts

`func (o *V1Organism) HasAssemblyCounts() bool`

HasAssemblyCounts returns a boolean if a field has been set.

### GetCounts

`func (o *V1Organism) GetCounts() []V1OrganismCountByType`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *V1Organism) GetCountsOk() (*[]V1OrganismCountByType, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *V1Organism) SetCounts(v []V1OrganismCountByType)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *V1Organism) HasCounts() bool`

HasCounts returns a boolean if a field has been set.

### GetChildren

`func (o *V1Organism) GetChildren() []V1Organism`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *V1Organism) GetChildrenOk() (*[]V1Organism, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *V1Organism) SetChildren(v []V1Organism)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *V1Organism) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetMerged

`func (o *V1Organism) GetMerged() []V1Organism`

GetMerged returns the Merged field if non-nil, zero value otherwise.

### GetMergedOk

`func (o *V1Organism) GetMergedOk() (*[]V1Organism, bool)`

GetMergedOk returns a tuple with the Merged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerged

`func (o *V1Organism) SetMerged(v []V1Organism)`

SetMerged sets Merged field to given value.

### HasMerged

`func (o *V1Organism) HasMerged() bool`

HasMerged returns a boolean if a field has been set.

### GetMergedTaxIds

`func (o *V1Organism) GetMergedTaxIds() []string`

GetMergedTaxIds returns the MergedTaxIds field if non-nil, zero value otherwise.

### GetMergedTaxIdsOk

`func (o *V1Organism) GetMergedTaxIdsOk() (*[]string, bool)`

GetMergedTaxIdsOk returns a tuple with the MergedTaxIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergedTaxIds

`func (o *V1Organism) SetMergedTaxIds(v []string)`

SetMergedTaxIds sets MergedTaxIds field to given value.

### HasMergedTaxIds

`func (o *V1Organism) HasMergedTaxIds() bool`

HasMergedTaxIds returns a boolean if a field has been set.

### GetMinOrd

`func (o *V1Organism) GetMinOrd() int32`

GetMinOrd returns the MinOrd field if non-nil, zero value otherwise.

### GetMinOrdOk

`func (o *V1Organism) GetMinOrdOk() (*int32, bool)`

GetMinOrdOk returns a tuple with the MinOrd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinOrd

`func (o *V1Organism) SetMinOrd(v int32)`

SetMinOrd sets MinOrd field to given value.

### HasMinOrd

`func (o *V1Organism) HasMinOrd() bool`

HasMinOrd returns a boolean if a field has been set.

### GetMaxOrd

`func (o *V1Organism) GetMaxOrd() int32`

GetMaxOrd returns the MaxOrd field if non-nil, zero value otherwise.

### GetMaxOrdOk

`func (o *V1Organism) GetMaxOrdOk() (*int32, bool)`

GetMaxOrdOk returns a tuple with the MaxOrd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxOrd

`func (o *V1Organism) SetMaxOrd(v int32)`

SetMaxOrd sets MaxOrd field to given value.

### HasMaxOrd

`func (o *V1Organism) HasMaxOrd() bool`

HasMaxOrd returns a boolean if a field has been set.

### GetWeight

`func (o *V1Organism) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *V1Organism) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *V1Organism) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *V1Organism) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetKey

`func (o *V1Organism) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *V1Organism) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *V1Organism) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *V1Organism) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetTitle

`func (o *V1Organism) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *V1Organism) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *V1Organism) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *V1Organism) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetIcon

`func (o *V1Organism) GetIcon() bool`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *V1Organism) GetIconOk() (*bool, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *V1Organism) SetIcon(v bool)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *V1Organism) HasIcon() bool`

HasIcon returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


