# V1reportsVirusAssembly

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accession** | Pointer to **string** |  | [optional] 
**IsComplete** | Pointer to **bool** |  | [optional] 
**IsAnnotated** | Pointer to **bool** |  | [optional] 
**Isolate** | Pointer to [**V1reportsVirusAssemblyIsolate**](V1reportsVirusAssemblyIsolate.md) |  | [optional] 
**SourceDatabase** | Pointer to **string** |  | [optional] 
**ProteinCount** | Pointer to **int32** |  | [optional] 
**Host** | Pointer to [**V1reportsOrganism**](V1reportsOrganism.md) |  | [optional] 
**Virus** | Pointer to [**V1reportsOrganism**](V1reportsOrganism.md) |  | [optional] 
**Bioprojects** | Pointer to **[]string** |  | [optional] 
**Location** | Pointer to [**V1reportsVirusAssemblyCollectionLocation**](V1reportsVirusAssemblyCollectionLocation.md) |  | [optional] 
**UpdateDate** | Pointer to **string** |  | [optional] 
**ReleaseDate** | Pointer to **string** |  | [optional] 
**NucleotideCompleteness** | Pointer to **string** |  | [optional] 
**Completeness** | Pointer to [**V1reportsVirusAssemblyCompleteness**](V1reportsVirusAssemblyCompleteness.md) |  | [optional] [default to V1REPORTSVIRUSASSEMBLYCOMPLETENESS_UNKNOWN]
**Length** | Pointer to **int32** |  | [optional] 
**GeneCount** | Pointer to **int32** |  | [optional] 
**MaturePeptideCount** | Pointer to **int32** |  | [optional] 
**Biosample** | Pointer to **string** |  | [optional] 
**MolType** | Pointer to **string** |  | [optional] 
**Annotation** | Pointer to [**V1reportsVirusAnnotation**](V1reportsVirusAnnotation.md) |  | [optional] 
**Nucleotide** | Pointer to [**V1reportsSeqRangeSetFasta**](V1reportsSeqRangeSetFasta.md) |  | [optional] 
**PurposeOfSampling** | Pointer to [**V1reportsPurposeOfSampling**](V1reportsPurposeOfSampling.md) |  | [optional] [default to V1REPORTSPURPOSEOFSAMPLING_UNKNOWN]
**SraAccessions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewV1reportsVirusAssembly

`func NewV1reportsVirusAssembly() *V1reportsVirusAssembly`

NewV1reportsVirusAssembly instantiates a new V1reportsVirusAssembly object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1reportsVirusAssemblyWithDefaults

`func NewV1reportsVirusAssemblyWithDefaults() *V1reportsVirusAssembly`

NewV1reportsVirusAssemblyWithDefaults instantiates a new V1reportsVirusAssembly object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccession

`func (o *V1reportsVirusAssembly) GetAccession() string`

GetAccession returns the Accession field if non-nil, zero value otherwise.

### GetAccessionOk

`func (o *V1reportsVirusAssembly) GetAccessionOk() (*string, bool)`

GetAccessionOk returns a tuple with the Accession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccession

`func (o *V1reportsVirusAssembly) SetAccession(v string)`

SetAccession sets Accession field to given value.

### HasAccession

`func (o *V1reportsVirusAssembly) HasAccession() bool`

HasAccession returns a boolean if a field has been set.

### GetIsComplete

`func (o *V1reportsVirusAssembly) GetIsComplete() bool`

GetIsComplete returns the IsComplete field if non-nil, zero value otherwise.

### GetIsCompleteOk

`func (o *V1reportsVirusAssembly) GetIsCompleteOk() (*bool, bool)`

GetIsCompleteOk returns a tuple with the IsComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsComplete

`func (o *V1reportsVirusAssembly) SetIsComplete(v bool)`

SetIsComplete sets IsComplete field to given value.

### HasIsComplete

`func (o *V1reportsVirusAssembly) HasIsComplete() bool`

HasIsComplete returns a boolean if a field has been set.

### GetIsAnnotated

`func (o *V1reportsVirusAssembly) GetIsAnnotated() bool`

GetIsAnnotated returns the IsAnnotated field if non-nil, zero value otherwise.

### GetIsAnnotatedOk

`func (o *V1reportsVirusAssembly) GetIsAnnotatedOk() (*bool, bool)`

GetIsAnnotatedOk returns a tuple with the IsAnnotated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnnotated

`func (o *V1reportsVirusAssembly) SetIsAnnotated(v bool)`

SetIsAnnotated sets IsAnnotated field to given value.

### HasIsAnnotated

`func (o *V1reportsVirusAssembly) HasIsAnnotated() bool`

HasIsAnnotated returns a boolean if a field has been set.

### GetIsolate

`func (o *V1reportsVirusAssembly) GetIsolate() V1reportsVirusAssemblyIsolate`

GetIsolate returns the Isolate field if non-nil, zero value otherwise.

### GetIsolateOk

`func (o *V1reportsVirusAssembly) GetIsolateOk() (*V1reportsVirusAssemblyIsolate, bool)`

GetIsolateOk returns a tuple with the Isolate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolate

`func (o *V1reportsVirusAssembly) SetIsolate(v V1reportsVirusAssemblyIsolate)`

SetIsolate sets Isolate field to given value.

### HasIsolate

`func (o *V1reportsVirusAssembly) HasIsolate() bool`

HasIsolate returns a boolean if a field has been set.

### GetSourceDatabase

`func (o *V1reportsVirusAssembly) GetSourceDatabase() string`

GetSourceDatabase returns the SourceDatabase field if non-nil, zero value otherwise.

### GetSourceDatabaseOk

`func (o *V1reportsVirusAssembly) GetSourceDatabaseOk() (*string, bool)`

GetSourceDatabaseOk returns a tuple with the SourceDatabase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDatabase

`func (o *V1reportsVirusAssembly) SetSourceDatabase(v string)`

SetSourceDatabase sets SourceDatabase field to given value.

### HasSourceDatabase

`func (o *V1reportsVirusAssembly) HasSourceDatabase() bool`

HasSourceDatabase returns a boolean if a field has been set.

### GetProteinCount

`func (o *V1reportsVirusAssembly) GetProteinCount() int32`

GetProteinCount returns the ProteinCount field if non-nil, zero value otherwise.

### GetProteinCountOk

`func (o *V1reportsVirusAssembly) GetProteinCountOk() (*int32, bool)`

GetProteinCountOk returns a tuple with the ProteinCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProteinCount

`func (o *V1reportsVirusAssembly) SetProteinCount(v int32)`

SetProteinCount sets ProteinCount field to given value.

### HasProteinCount

`func (o *V1reportsVirusAssembly) HasProteinCount() bool`

HasProteinCount returns a boolean if a field has been set.

### GetHost

`func (o *V1reportsVirusAssembly) GetHost() V1reportsOrganism`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *V1reportsVirusAssembly) GetHostOk() (*V1reportsOrganism, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *V1reportsVirusAssembly) SetHost(v V1reportsOrganism)`

SetHost sets Host field to given value.

### HasHost

`func (o *V1reportsVirusAssembly) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetVirus

`func (o *V1reportsVirusAssembly) GetVirus() V1reportsOrganism`

GetVirus returns the Virus field if non-nil, zero value otherwise.

### GetVirusOk

`func (o *V1reportsVirusAssembly) GetVirusOk() (*V1reportsOrganism, bool)`

GetVirusOk returns a tuple with the Virus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirus

`func (o *V1reportsVirusAssembly) SetVirus(v V1reportsOrganism)`

SetVirus sets Virus field to given value.

### HasVirus

`func (o *V1reportsVirusAssembly) HasVirus() bool`

HasVirus returns a boolean if a field has been set.

### GetBioprojects

`func (o *V1reportsVirusAssembly) GetBioprojects() []string`

GetBioprojects returns the Bioprojects field if non-nil, zero value otherwise.

### GetBioprojectsOk

`func (o *V1reportsVirusAssembly) GetBioprojectsOk() (*[]string, bool)`

GetBioprojectsOk returns a tuple with the Bioprojects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBioprojects

`func (o *V1reportsVirusAssembly) SetBioprojects(v []string)`

SetBioprojects sets Bioprojects field to given value.

### HasBioprojects

`func (o *V1reportsVirusAssembly) HasBioprojects() bool`

HasBioprojects returns a boolean if a field has been set.

### GetLocation

`func (o *V1reportsVirusAssembly) GetLocation() V1reportsVirusAssemblyCollectionLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *V1reportsVirusAssembly) GetLocationOk() (*V1reportsVirusAssemblyCollectionLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *V1reportsVirusAssembly) SetLocation(v V1reportsVirusAssemblyCollectionLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *V1reportsVirusAssembly) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetUpdateDate

`func (o *V1reportsVirusAssembly) GetUpdateDate() string`

GetUpdateDate returns the UpdateDate field if non-nil, zero value otherwise.

### GetUpdateDateOk

`func (o *V1reportsVirusAssembly) GetUpdateDateOk() (*string, bool)`

GetUpdateDateOk returns a tuple with the UpdateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateDate

`func (o *V1reportsVirusAssembly) SetUpdateDate(v string)`

SetUpdateDate sets UpdateDate field to given value.

### HasUpdateDate

`func (o *V1reportsVirusAssembly) HasUpdateDate() bool`

HasUpdateDate returns a boolean if a field has been set.

### GetReleaseDate

`func (o *V1reportsVirusAssembly) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *V1reportsVirusAssembly) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *V1reportsVirusAssembly) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *V1reportsVirusAssembly) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### GetNucleotideCompleteness

`func (o *V1reportsVirusAssembly) GetNucleotideCompleteness() string`

GetNucleotideCompleteness returns the NucleotideCompleteness field if non-nil, zero value otherwise.

### GetNucleotideCompletenessOk

`func (o *V1reportsVirusAssembly) GetNucleotideCompletenessOk() (*string, bool)`

GetNucleotideCompletenessOk returns a tuple with the NucleotideCompleteness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNucleotideCompleteness

`func (o *V1reportsVirusAssembly) SetNucleotideCompleteness(v string)`

SetNucleotideCompleteness sets NucleotideCompleteness field to given value.

### HasNucleotideCompleteness

`func (o *V1reportsVirusAssembly) HasNucleotideCompleteness() bool`

HasNucleotideCompleteness returns a boolean if a field has been set.

### GetCompleteness

`func (o *V1reportsVirusAssembly) GetCompleteness() V1reportsVirusAssemblyCompleteness`

GetCompleteness returns the Completeness field if non-nil, zero value otherwise.

### GetCompletenessOk

`func (o *V1reportsVirusAssembly) GetCompletenessOk() (*V1reportsVirusAssemblyCompleteness, bool)`

GetCompletenessOk returns a tuple with the Completeness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleteness

`func (o *V1reportsVirusAssembly) SetCompleteness(v V1reportsVirusAssemblyCompleteness)`

SetCompleteness sets Completeness field to given value.

### HasCompleteness

`func (o *V1reportsVirusAssembly) HasCompleteness() bool`

HasCompleteness returns a boolean if a field has been set.

### GetLength

`func (o *V1reportsVirusAssembly) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *V1reportsVirusAssembly) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *V1reportsVirusAssembly) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *V1reportsVirusAssembly) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetGeneCount

`func (o *V1reportsVirusAssembly) GetGeneCount() int32`

GetGeneCount returns the GeneCount field if non-nil, zero value otherwise.

### GetGeneCountOk

`func (o *V1reportsVirusAssembly) GetGeneCountOk() (*int32, bool)`

GetGeneCountOk returns a tuple with the GeneCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneCount

`func (o *V1reportsVirusAssembly) SetGeneCount(v int32)`

SetGeneCount sets GeneCount field to given value.

### HasGeneCount

`func (o *V1reportsVirusAssembly) HasGeneCount() bool`

HasGeneCount returns a boolean if a field has been set.

### GetMaturePeptideCount

`func (o *V1reportsVirusAssembly) GetMaturePeptideCount() int32`

GetMaturePeptideCount returns the MaturePeptideCount field if non-nil, zero value otherwise.

### GetMaturePeptideCountOk

`func (o *V1reportsVirusAssembly) GetMaturePeptideCountOk() (*int32, bool)`

GetMaturePeptideCountOk returns a tuple with the MaturePeptideCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaturePeptideCount

`func (o *V1reportsVirusAssembly) SetMaturePeptideCount(v int32)`

SetMaturePeptideCount sets MaturePeptideCount field to given value.

### HasMaturePeptideCount

`func (o *V1reportsVirusAssembly) HasMaturePeptideCount() bool`

HasMaturePeptideCount returns a boolean if a field has been set.

### GetBiosample

`func (o *V1reportsVirusAssembly) GetBiosample() string`

GetBiosample returns the Biosample field if non-nil, zero value otherwise.

### GetBiosampleOk

`func (o *V1reportsVirusAssembly) GetBiosampleOk() (*string, bool)`

GetBiosampleOk returns a tuple with the Biosample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosample

`func (o *V1reportsVirusAssembly) SetBiosample(v string)`

SetBiosample sets Biosample field to given value.

### HasBiosample

`func (o *V1reportsVirusAssembly) HasBiosample() bool`

HasBiosample returns a boolean if a field has been set.

### GetMolType

`func (o *V1reportsVirusAssembly) GetMolType() string`

GetMolType returns the MolType field if non-nil, zero value otherwise.

### GetMolTypeOk

`func (o *V1reportsVirusAssembly) GetMolTypeOk() (*string, bool)`

GetMolTypeOk returns a tuple with the MolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMolType

`func (o *V1reportsVirusAssembly) SetMolType(v string)`

SetMolType sets MolType field to given value.

### HasMolType

`func (o *V1reportsVirusAssembly) HasMolType() bool`

HasMolType returns a boolean if a field has been set.

### GetAnnotation

`func (o *V1reportsVirusAssembly) GetAnnotation() V1reportsVirusAnnotation`

GetAnnotation returns the Annotation field if non-nil, zero value otherwise.

### GetAnnotationOk

`func (o *V1reportsVirusAssembly) GetAnnotationOk() (*V1reportsVirusAnnotation, bool)`

GetAnnotationOk returns a tuple with the Annotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotation

`func (o *V1reportsVirusAssembly) SetAnnotation(v V1reportsVirusAnnotation)`

SetAnnotation sets Annotation field to given value.

### HasAnnotation

`func (o *V1reportsVirusAssembly) HasAnnotation() bool`

HasAnnotation returns a boolean if a field has been set.

### GetNucleotide

`func (o *V1reportsVirusAssembly) GetNucleotide() V1reportsSeqRangeSetFasta`

GetNucleotide returns the Nucleotide field if non-nil, zero value otherwise.

### GetNucleotideOk

`func (o *V1reportsVirusAssembly) GetNucleotideOk() (*V1reportsSeqRangeSetFasta, bool)`

GetNucleotideOk returns a tuple with the Nucleotide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNucleotide

`func (o *V1reportsVirusAssembly) SetNucleotide(v V1reportsSeqRangeSetFasta)`

SetNucleotide sets Nucleotide field to given value.

### HasNucleotide

`func (o *V1reportsVirusAssembly) HasNucleotide() bool`

HasNucleotide returns a boolean if a field has been set.

### GetPurposeOfSampling

`func (o *V1reportsVirusAssembly) GetPurposeOfSampling() V1reportsPurposeOfSampling`

GetPurposeOfSampling returns the PurposeOfSampling field if non-nil, zero value otherwise.

### GetPurposeOfSamplingOk

`func (o *V1reportsVirusAssembly) GetPurposeOfSamplingOk() (*V1reportsPurposeOfSampling, bool)`

GetPurposeOfSamplingOk returns a tuple with the PurposeOfSampling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurposeOfSampling

`func (o *V1reportsVirusAssembly) SetPurposeOfSampling(v V1reportsPurposeOfSampling)`

SetPurposeOfSampling sets PurposeOfSampling field to given value.

### HasPurposeOfSampling

`func (o *V1reportsVirusAssembly) HasPurposeOfSampling() bool`

HasPurposeOfSampling returns a boolean if a field has been set.

### GetSraAccessions

`func (o *V1reportsVirusAssembly) GetSraAccessions() []string`

GetSraAccessions returns the SraAccessions field if non-nil, zero value otherwise.

### GetSraAccessionsOk

`func (o *V1reportsVirusAssembly) GetSraAccessionsOk() (*[]string, bool)`

GetSraAccessionsOk returns a tuple with the SraAccessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSraAccessions

`func (o *V1reportsVirusAssembly) SetSraAccessions(v []string)`

SetSraAccessions sets SraAccessions field to given value.

### HasSraAccessions

`func (o *V1reportsVirusAssembly) HasSraAccessions() bool`

HasSraAccessions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


