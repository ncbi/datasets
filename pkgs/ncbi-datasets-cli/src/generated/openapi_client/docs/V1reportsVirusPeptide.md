# V1reportsVirusPeptide

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accession** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OtherNames** | Pointer to **[]string** |  | [optional] 
**Nucleotide** | Pointer to [**V1reportsSeqRangeSetFasta**](V1reportsSeqRangeSetFasta.md) |  | [optional] 
**Protein** | Pointer to [**V1reportsSeqRangeSetFasta**](V1reportsSeqRangeSetFasta.md) |  | [optional] 
**PdbIds** | Pointer to **[]string** |  | [optional] 
**Cdd** | Pointer to [**[]V1reportsConservedDomain**](V1reportsConservedDomain.md) |  | [optional] 
**UniProtKb** | Pointer to [**V1reportsVirusPeptideUniProtId**](V1reportsVirusPeptideUniProtId.md) |  | [optional] 
**MaturePeptide** | Pointer to [**[]V1reportsVirusPeptide**](V1reportsVirusPeptide.md) |  | [optional] 
**ProteinCompleteness** | Pointer to [**V1reportsVirusAssemblyCompleteness**](V1reportsVirusAssemblyCompleteness.md) |  | [optional] [default to V1REPORTSVIRUSASSEMBLYCOMPLETENESS_UNKNOWN]

## Methods

### NewV1reportsVirusPeptide

`func NewV1reportsVirusPeptide() *V1reportsVirusPeptide`

NewV1reportsVirusPeptide instantiates a new V1reportsVirusPeptide object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1reportsVirusPeptideWithDefaults

`func NewV1reportsVirusPeptideWithDefaults() *V1reportsVirusPeptide`

NewV1reportsVirusPeptideWithDefaults instantiates a new V1reportsVirusPeptide object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccession

`func (o *V1reportsVirusPeptide) GetAccession() string`

GetAccession returns the Accession field if non-nil, zero value otherwise.

### GetAccessionOk

`func (o *V1reportsVirusPeptide) GetAccessionOk() (*string, bool)`

GetAccessionOk returns a tuple with the Accession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccession

`func (o *V1reportsVirusPeptide) SetAccession(v string)`

SetAccession sets Accession field to given value.

### HasAccession

`func (o *V1reportsVirusPeptide) HasAccession() bool`

HasAccession returns a boolean if a field has been set.

### GetName

`func (o *V1reportsVirusPeptide) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1reportsVirusPeptide) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1reportsVirusPeptide) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1reportsVirusPeptide) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOtherNames

`func (o *V1reportsVirusPeptide) GetOtherNames() []string`

GetOtherNames returns the OtherNames field if non-nil, zero value otherwise.

### GetOtherNamesOk

`func (o *V1reportsVirusPeptide) GetOtherNamesOk() (*[]string, bool)`

GetOtherNamesOk returns a tuple with the OtherNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherNames

`func (o *V1reportsVirusPeptide) SetOtherNames(v []string)`

SetOtherNames sets OtherNames field to given value.

### HasOtherNames

`func (o *V1reportsVirusPeptide) HasOtherNames() bool`

HasOtherNames returns a boolean if a field has been set.

### GetNucleotide

`func (o *V1reportsVirusPeptide) GetNucleotide() V1reportsSeqRangeSetFasta`

GetNucleotide returns the Nucleotide field if non-nil, zero value otherwise.

### GetNucleotideOk

`func (o *V1reportsVirusPeptide) GetNucleotideOk() (*V1reportsSeqRangeSetFasta, bool)`

GetNucleotideOk returns a tuple with the Nucleotide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNucleotide

`func (o *V1reportsVirusPeptide) SetNucleotide(v V1reportsSeqRangeSetFasta)`

SetNucleotide sets Nucleotide field to given value.

### HasNucleotide

`func (o *V1reportsVirusPeptide) HasNucleotide() bool`

HasNucleotide returns a boolean if a field has been set.

### GetProtein

`func (o *V1reportsVirusPeptide) GetProtein() V1reportsSeqRangeSetFasta`

GetProtein returns the Protein field if non-nil, zero value otherwise.

### GetProteinOk

`func (o *V1reportsVirusPeptide) GetProteinOk() (*V1reportsSeqRangeSetFasta, bool)`

GetProteinOk returns a tuple with the Protein field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtein

`func (o *V1reportsVirusPeptide) SetProtein(v V1reportsSeqRangeSetFasta)`

SetProtein sets Protein field to given value.

### HasProtein

`func (o *V1reportsVirusPeptide) HasProtein() bool`

HasProtein returns a boolean if a field has been set.

### GetPdbIds

`func (o *V1reportsVirusPeptide) GetPdbIds() []string`

GetPdbIds returns the PdbIds field if non-nil, zero value otherwise.

### GetPdbIdsOk

`func (o *V1reportsVirusPeptide) GetPdbIdsOk() (*[]string, bool)`

GetPdbIdsOk returns a tuple with the PdbIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdbIds

`func (o *V1reportsVirusPeptide) SetPdbIds(v []string)`

SetPdbIds sets PdbIds field to given value.

### HasPdbIds

`func (o *V1reportsVirusPeptide) HasPdbIds() bool`

HasPdbIds returns a boolean if a field has been set.

### GetCdd

`func (o *V1reportsVirusPeptide) GetCdd() []V1reportsConservedDomain`

GetCdd returns the Cdd field if non-nil, zero value otherwise.

### GetCddOk

`func (o *V1reportsVirusPeptide) GetCddOk() (*[]V1reportsConservedDomain, bool)`

GetCddOk returns a tuple with the Cdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdd

`func (o *V1reportsVirusPeptide) SetCdd(v []V1reportsConservedDomain)`

SetCdd sets Cdd field to given value.

### HasCdd

`func (o *V1reportsVirusPeptide) HasCdd() bool`

HasCdd returns a boolean if a field has been set.

### GetUniProtKb

`func (o *V1reportsVirusPeptide) GetUniProtKb() V1reportsVirusPeptideUniProtId`

GetUniProtKb returns the UniProtKb field if non-nil, zero value otherwise.

### GetUniProtKbOk

`func (o *V1reportsVirusPeptide) GetUniProtKbOk() (*V1reportsVirusPeptideUniProtId, bool)`

GetUniProtKbOk returns a tuple with the UniProtKb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniProtKb

`func (o *V1reportsVirusPeptide) SetUniProtKb(v V1reportsVirusPeptideUniProtId)`

SetUniProtKb sets UniProtKb field to given value.

### HasUniProtKb

`func (o *V1reportsVirusPeptide) HasUniProtKb() bool`

HasUniProtKb returns a boolean if a field has been set.

### GetMaturePeptide

`func (o *V1reportsVirusPeptide) GetMaturePeptide() []V1reportsVirusPeptide`

GetMaturePeptide returns the MaturePeptide field if non-nil, zero value otherwise.

### GetMaturePeptideOk

`func (o *V1reportsVirusPeptide) GetMaturePeptideOk() (*[]V1reportsVirusPeptide, bool)`

GetMaturePeptideOk returns a tuple with the MaturePeptide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaturePeptide

`func (o *V1reportsVirusPeptide) SetMaturePeptide(v []V1reportsVirusPeptide)`

SetMaturePeptide sets MaturePeptide field to given value.

### HasMaturePeptide

`func (o *V1reportsVirusPeptide) HasMaturePeptide() bool`

HasMaturePeptide returns a boolean if a field has been set.

### GetProteinCompleteness

`func (o *V1reportsVirusPeptide) GetProteinCompleteness() V1reportsVirusAssemblyCompleteness`

GetProteinCompleteness returns the ProteinCompleteness field if non-nil, zero value otherwise.

### GetProteinCompletenessOk

`func (o *V1reportsVirusPeptide) GetProteinCompletenessOk() (*V1reportsVirusAssemblyCompleteness, bool)`

GetProteinCompletenessOk returns a tuple with the ProteinCompleteness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProteinCompleteness

`func (o *V1reportsVirusPeptide) SetProteinCompleteness(v V1reportsVirusAssemblyCompleteness)`

SetProteinCompleteness sets ProteinCompleteness field to given value.

### HasProteinCompleteness

`func (o *V1reportsVirusPeptide) HasProteinCompleteness() bool`

HasProteinCompleteness returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


