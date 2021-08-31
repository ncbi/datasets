# V1reportsAssemblyStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalNumberOfChromosomes** | Pointer to **int32** |  | [optional] 
**TotalSequenceLength** | Pointer to **string** |  | [optional] 
**TotalUngappedLength** | Pointer to **string** |  | [optional] 
**NumberOfContigs** | Pointer to **int32** |  | [optional] 
**ContigN50** | Pointer to **int32** |  | [optional] 
**ContigL50** | Pointer to **int32** |  | [optional] 
**NumberOfScaffolds** | Pointer to **int32** |  | [optional] 
**ScaffoldN50** | Pointer to **int32** |  | [optional] 
**ScaffoldL50** | Pointer to **int32** |  | [optional] 
**GapsBetweenScaffoldsCount** | Pointer to **int32** |  | [optional] 
**NumberOfComponentSequences** | Pointer to **int32** |  | [optional] 
**GcCount** | Pointer to **string** |  | [optional] 

## Methods

### NewV1reportsAssemblyStats

`func NewV1reportsAssemblyStats() *V1reportsAssemblyStats`

NewV1reportsAssemblyStats instantiates a new V1reportsAssemblyStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1reportsAssemblyStatsWithDefaults

`func NewV1reportsAssemblyStatsWithDefaults() *V1reportsAssemblyStats`

NewV1reportsAssemblyStatsWithDefaults instantiates a new V1reportsAssemblyStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalNumberOfChromosomes

`func (o *V1reportsAssemblyStats) GetTotalNumberOfChromosomes() int32`

GetTotalNumberOfChromosomes returns the TotalNumberOfChromosomes field if non-nil, zero value otherwise.

### GetTotalNumberOfChromosomesOk

`func (o *V1reportsAssemblyStats) GetTotalNumberOfChromosomesOk() (*int32, bool)`

GetTotalNumberOfChromosomesOk returns a tuple with the TotalNumberOfChromosomes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNumberOfChromosomes

`func (o *V1reportsAssemblyStats) SetTotalNumberOfChromosomes(v int32)`

SetTotalNumberOfChromosomes sets TotalNumberOfChromosomes field to given value.

### HasTotalNumberOfChromosomes

`func (o *V1reportsAssemblyStats) HasTotalNumberOfChromosomes() bool`

HasTotalNumberOfChromosomes returns a boolean if a field has been set.

### GetTotalSequenceLength

`func (o *V1reportsAssemblyStats) GetTotalSequenceLength() string`

GetTotalSequenceLength returns the TotalSequenceLength field if non-nil, zero value otherwise.

### GetTotalSequenceLengthOk

`func (o *V1reportsAssemblyStats) GetTotalSequenceLengthOk() (*string, bool)`

GetTotalSequenceLengthOk returns a tuple with the TotalSequenceLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSequenceLength

`func (o *V1reportsAssemblyStats) SetTotalSequenceLength(v string)`

SetTotalSequenceLength sets TotalSequenceLength field to given value.

### HasTotalSequenceLength

`func (o *V1reportsAssemblyStats) HasTotalSequenceLength() bool`

HasTotalSequenceLength returns a boolean if a field has been set.

### GetTotalUngappedLength

`func (o *V1reportsAssemblyStats) GetTotalUngappedLength() string`

GetTotalUngappedLength returns the TotalUngappedLength field if non-nil, zero value otherwise.

### GetTotalUngappedLengthOk

`func (o *V1reportsAssemblyStats) GetTotalUngappedLengthOk() (*string, bool)`

GetTotalUngappedLengthOk returns a tuple with the TotalUngappedLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUngappedLength

`func (o *V1reportsAssemblyStats) SetTotalUngappedLength(v string)`

SetTotalUngappedLength sets TotalUngappedLength field to given value.

### HasTotalUngappedLength

`func (o *V1reportsAssemblyStats) HasTotalUngappedLength() bool`

HasTotalUngappedLength returns a boolean if a field has been set.

### GetNumberOfContigs

`func (o *V1reportsAssemblyStats) GetNumberOfContigs() int32`

GetNumberOfContigs returns the NumberOfContigs field if non-nil, zero value otherwise.

### GetNumberOfContigsOk

`func (o *V1reportsAssemblyStats) GetNumberOfContigsOk() (*int32, bool)`

GetNumberOfContigsOk returns a tuple with the NumberOfContigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfContigs

`func (o *V1reportsAssemblyStats) SetNumberOfContigs(v int32)`

SetNumberOfContigs sets NumberOfContigs field to given value.

### HasNumberOfContigs

`func (o *V1reportsAssemblyStats) HasNumberOfContigs() bool`

HasNumberOfContigs returns a boolean if a field has been set.

### GetContigN50

`func (o *V1reportsAssemblyStats) GetContigN50() int32`

GetContigN50 returns the ContigN50 field if non-nil, zero value otherwise.

### GetContigN50Ok

`func (o *V1reportsAssemblyStats) GetContigN50Ok() (*int32, bool)`

GetContigN50Ok returns a tuple with the ContigN50 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContigN50

`func (o *V1reportsAssemblyStats) SetContigN50(v int32)`

SetContigN50 sets ContigN50 field to given value.

### HasContigN50

`func (o *V1reportsAssemblyStats) HasContigN50() bool`

HasContigN50 returns a boolean if a field has been set.

### GetContigL50

`func (o *V1reportsAssemblyStats) GetContigL50() int32`

GetContigL50 returns the ContigL50 field if non-nil, zero value otherwise.

### GetContigL50Ok

`func (o *V1reportsAssemblyStats) GetContigL50Ok() (*int32, bool)`

GetContigL50Ok returns a tuple with the ContigL50 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContigL50

`func (o *V1reportsAssemblyStats) SetContigL50(v int32)`

SetContigL50 sets ContigL50 field to given value.

### HasContigL50

`func (o *V1reportsAssemblyStats) HasContigL50() bool`

HasContigL50 returns a boolean if a field has been set.

### GetNumberOfScaffolds

`func (o *V1reportsAssemblyStats) GetNumberOfScaffolds() int32`

GetNumberOfScaffolds returns the NumberOfScaffolds field if non-nil, zero value otherwise.

### GetNumberOfScaffoldsOk

`func (o *V1reportsAssemblyStats) GetNumberOfScaffoldsOk() (*int32, bool)`

GetNumberOfScaffoldsOk returns a tuple with the NumberOfScaffolds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfScaffolds

`func (o *V1reportsAssemblyStats) SetNumberOfScaffolds(v int32)`

SetNumberOfScaffolds sets NumberOfScaffolds field to given value.

### HasNumberOfScaffolds

`func (o *V1reportsAssemblyStats) HasNumberOfScaffolds() bool`

HasNumberOfScaffolds returns a boolean if a field has been set.

### GetScaffoldN50

`func (o *V1reportsAssemblyStats) GetScaffoldN50() int32`

GetScaffoldN50 returns the ScaffoldN50 field if non-nil, zero value otherwise.

### GetScaffoldN50Ok

`func (o *V1reportsAssemblyStats) GetScaffoldN50Ok() (*int32, bool)`

GetScaffoldN50Ok returns a tuple with the ScaffoldN50 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaffoldN50

`func (o *V1reportsAssemblyStats) SetScaffoldN50(v int32)`

SetScaffoldN50 sets ScaffoldN50 field to given value.

### HasScaffoldN50

`func (o *V1reportsAssemblyStats) HasScaffoldN50() bool`

HasScaffoldN50 returns a boolean if a field has been set.

### GetScaffoldL50

`func (o *V1reportsAssemblyStats) GetScaffoldL50() int32`

GetScaffoldL50 returns the ScaffoldL50 field if non-nil, zero value otherwise.

### GetScaffoldL50Ok

`func (o *V1reportsAssemblyStats) GetScaffoldL50Ok() (*int32, bool)`

GetScaffoldL50Ok returns a tuple with the ScaffoldL50 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaffoldL50

`func (o *V1reportsAssemblyStats) SetScaffoldL50(v int32)`

SetScaffoldL50 sets ScaffoldL50 field to given value.

### HasScaffoldL50

`func (o *V1reportsAssemblyStats) HasScaffoldL50() bool`

HasScaffoldL50 returns a boolean if a field has been set.

### GetGapsBetweenScaffoldsCount

`func (o *V1reportsAssemblyStats) GetGapsBetweenScaffoldsCount() int32`

GetGapsBetweenScaffoldsCount returns the GapsBetweenScaffoldsCount field if non-nil, zero value otherwise.

### GetGapsBetweenScaffoldsCountOk

`func (o *V1reportsAssemblyStats) GetGapsBetweenScaffoldsCountOk() (*int32, bool)`

GetGapsBetweenScaffoldsCountOk returns a tuple with the GapsBetweenScaffoldsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGapsBetweenScaffoldsCount

`func (o *V1reportsAssemblyStats) SetGapsBetweenScaffoldsCount(v int32)`

SetGapsBetweenScaffoldsCount sets GapsBetweenScaffoldsCount field to given value.

### HasGapsBetweenScaffoldsCount

`func (o *V1reportsAssemblyStats) HasGapsBetweenScaffoldsCount() bool`

HasGapsBetweenScaffoldsCount returns a boolean if a field has been set.

### GetNumberOfComponentSequences

`func (o *V1reportsAssemblyStats) GetNumberOfComponentSequences() int32`

GetNumberOfComponentSequences returns the NumberOfComponentSequences field if non-nil, zero value otherwise.

### GetNumberOfComponentSequencesOk

`func (o *V1reportsAssemblyStats) GetNumberOfComponentSequencesOk() (*int32, bool)`

GetNumberOfComponentSequencesOk returns a tuple with the NumberOfComponentSequences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfComponentSequences

`func (o *V1reportsAssemblyStats) SetNumberOfComponentSequences(v int32)`

SetNumberOfComponentSequences sets NumberOfComponentSequences field to given value.

### HasNumberOfComponentSequences

`func (o *V1reportsAssemblyStats) HasNumberOfComponentSequences() bool`

HasNumberOfComponentSequences returns a boolean if a field has been set.

### GetGcCount

`func (o *V1reportsAssemblyStats) GetGcCount() string`

GetGcCount returns the GcCount field if non-nil, zero value otherwise.

### GetGcCountOk

`func (o *V1reportsAssemblyStats) GetGcCountOk() (*string, bool)`

GetGcCountOk returns a tuple with the GcCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcCount

`func (o *V1reportsAssemblyStats) SetGcCount(v string)`

SetGcCount sets GcCount field to given value.

### HasGcCount

`func (o *V1reportsAssemblyStats) HasGcCount() bool`

HasGcCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


