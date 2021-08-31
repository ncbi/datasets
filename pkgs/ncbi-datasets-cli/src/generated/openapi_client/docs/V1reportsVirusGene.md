# V1reportsVirusGene

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**GeneId** | Pointer to **int32** |  | [optional] 
**Nucleotide** | Pointer to [**V1reportsSeqRangeSetFasta**](V1reportsSeqRangeSetFasta.md) |  | [optional] 
**Cds** | Pointer to [**[]V1reportsVirusPeptide**](V1reportsVirusPeptide.md) |  | [optional] 

## Methods

### NewV1reportsVirusGene

`func NewV1reportsVirusGene() *V1reportsVirusGene`

NewV1reportsVirusGene instantiates a new V1reportsVirusGene object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1reportsVirusGeneWithDefaults

`func NewV1reportsVirusGeneWithDefaults() *V1reportsVirusGene`

NewV1reportsVirusGeneWithDefaults instantiates a new V1reportsVirusGene object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *V1reportsVirusGene) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1reportsVirusGene) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1reportsVirusGene) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1reportsVirusGene) HasName() bool`

HasName returns a boolean if a field has been set.

### GetGeneId

`func (o *V1reportsVirusGene) GetGeneId() int32`

GetGeneId returns the GeneId field if non-nil, zero value otherwise.

### GetGeneIdOk

`func (o *V1reportsVirusGene) GetGeneIdOk() (*int32, bool)`

GetGeneIdOk returns a tuple with the GeneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneId

`func (o *V1reportsVirusGene) SetGeneId(v int32)`

SetGeneId sets GeneId field to given value.

### HasGeneId

`func (o *V1reportsVirusGene) HasGeneId() bool`

HasGeneId returns a boolean if a field has been set.

### GetNucleotide

`func (o *V1reportsVirusGene) GetNucleotide() V1reportsSeqRangeSetFasta`

GetNucleotide returns the Nucleotide field if non-nil, zero value otherwise.

### GetNucleotideOk

`func (o *V1reportsVirusGene) GetNucleotideOk() (*V1reportsSeqRangeSetFasta, bool)`

GetNucleotideOk returns a tuple with the Nucleotide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNucleotide

`func (o *V1reportsVirusGene) SetNucleotide(v V1reportsSeqRangeSetFasta)`

SetNucleotide sets Nucleotide field to given value.

### HasNucleotide

`func (o *V1reportsVirusGene) HasNucleotide() bool`

HasNucleotide returns a boolean if a field has been set.

### GetCds

`func (o *V1reportsVirusGene) GetCds() []V1reportsVirusPeptide`

GetCds returns the Cds field if non-nil, zero value otherwise.

### GetCdsOk

`func (o *V1reportsVirusGene) GetCdsOk() (*[]V1reportsVirusPeptide, bool)`

GetCdsOk returns a tuple with the Cds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCds

`func (o *V1reportsVirusGene) SetCds(v []V1reportsVirusPeptide)`

SetCds sets Cds field to given value.

### HasCds

`func (o *V1reportsVirusGene) HasCds() bool`

HasCds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


