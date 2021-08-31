# V1reportsProkaryoteGeneLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProteinAccession** | Pointer to **string** |  | [optional] 
**RefseqGenomicLocation** | Pointer to [**V1reportsSeqRangeWithAssembly**](V1reportsSeqRangeWithAssembly.md) |  | [optional] 
**GenbankGenomicLocation** | Pointer to [**V1reportsSeqRangeWithAssembly**](V1reportsSeqRangeWithAssembly.md) |  | [optional] 
**Organism** | Pointer to [**V1reportsOrganism**](V1reportsOrganism.md) |  | [optional] 
**Completeness** | Pointer to [**V1reportsProkaryoteGeneLocationCompleteness**](V1reportsProkaryoteGeneLocationCompleteness.md) |  | [optional] [default to V1REPORTSPROKARYOTEGENELOCATIONCOMPLETENESS_COMPLETE]
**ChromosomeName** | Pointer to **string** |  | [optional] 

## Methods

### NewV1reportsProkaryoteGeneLocation

`func NewV1reportsProkaryoteGeneLocation() *V1reportsProkaryoteGeneLocation`

NewV1reportsProkaryoteGeneLocation instantiates a new V1reportsProkaryoteGeneLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1reportsProkaryoteGeneLocationWithDefaults

`func NewV1reportsProkaryoteGeneLocationWithDefaults() *V1reportsProkaryoteGeneLocation`

NewV1reportsProkaryoteGeneLocationWithDefaults instantiates a new V1reportsProkaryoteGeneLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProteinAccession

`func (o *V1reportsProkaryoteGeneLocation) GetProteinAccession() string`

GetProteinAccession returns the ProteinAccession field if non-nil, zero value otherwise.

### GetProteinAccessionOk

`func (o *V1reportsProkaryoteGeneLocation) GetProteinAccessionOk() (*string, bool)`

GetProteinAccessionOk returns a tuple with the ProteinAccession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProteinAccession

`func (o *V1reportsProkaryoteGeneLocation) SetProteinAccession(v string)`

SetProteinAccession sets ProteinAccession field to given value.

### HasProteinAccession

`func (o *V1reportsProkaryoteGeneLocation) HasProteinAccession() bool`

HasProteinAccession returns a boolean if a field has been set.

### GetRefseqGenomicLocation

`func (o *V1reportsProkaryoteGeneLocation) GetRefseqGenomicLocation() V1reportsSeqRangeWithAssembly`

GetRefseqGenomicLocation returns the RefseqGenomicLocation field if non-nil, zero value otherwise.

### GetRefseqGenomicLocationOk

`func (o *V1reportsProkaryoteGeneLocation) GetRefseqGenomicLocationOk() (*V1reportsSeqRangeWithAssembly, bool)`

GetRefseqGenomicLocationOk returns a tuple with the RefseqGenomicLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefseqGenomicLocation

`func (o *V1reportsProkaryoteGeneLocation) SetRefseqGenomicLocation(v V1reportsSeqRangeWithAssembly)`

SetRefseqGenomicLocation sets RefseqGenomicLocation field to given value.

### HasRefseqGenomicLocation

`func (o *V1reportsProkaryoteGeneLocation) HasRefseqGenomicLocation() bool`

HasRefseqGenomicLocation returns a boolean if a field has been set.

### GetGenbankGenomicLocation

`func (o *V1reportsProkaryoteGeneLocation) GetGenbankGenomicLocation() V1reportsSeqRangeWithAssembly`

GetGenbankGenomicLocation returns the GenbankGenomicLocation field if non-nil, zero value otherwise.

### GetGenbankGenomicLocationOk

`func (o *V1reportsProkaryoteGeneLocation) GetGenbankGenomicLocationOk() (*V1reportsSeqRangeWithAssembly, bool)`

GetGenbankGenomicLocationOk returns a tuple with the GenbankGenomicLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenbankGenomicLocation

`func (o *V1reportsProkaryoteGeneLocation) SetGenbankGenomicLocation(v V1reportsSeqRangeWithAssembly)`

SetGenbankGenomicLocation sets GenbankGenomicLocation field to given value.

### HasGenbankGenomicLocation

`func (o *V1reportsProkaryoteGeneLocation) HasGenbankGenomicLocation() bool`

HasGenbankGenomicLocation returns a boolean if a field has been set.

### GetOrganism

`func (o *V1reportsProkaryoteGeneLocation) GetOrganism() V1reportsOrganism`

GetOrganism returns the Organism field if non-nil, zero value otherwise.

### GetOrganismOk

`func (o *V1reportsProkaryoteGeneLocation) GetOrganismOk() (*V1reportsOrganism, bool)`

GetOrganismOk returns a tuple with the Organism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganism

`func (o *V1reportsProkaryoteGeneLocation) SetOrganism(v V1reportsOrganism)`

SetOrganism sets Organism field to given value.

### HasOrganism

`func (o *V1reportsProkaryoteGeneLocation) HasOrganism() bool`

HasOrganism returns a boolean if a field has been set.

### GetCompleteness

`func (o *V1reportsProkaryoteGeneLocation) GetCompleteness() V1reportsProkaryoteGeneLocationCompleteness`

GetCompleteness returns the Completeness field if non-nil, zero value otherwise.

### GetCompletenessOk

`func (o *V1reportsProkaryoteGeneLocation) GetCompletenessOk() (*V1reportsProkaryoteGeneLocationCompleteness, bool)`

GetCompletenessOk returns a tuple with the Completeness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleteness

`func (o *V1reportsProkaryoteGeneLocation) SetCompleteness(v V1reportsProkaryoteGeneLocationCompleteness)`

SetCompleteness sets Completeness field to given value.

### HasCompleteness

`func (o *V1reportsProkaryoteGeneLocation) HasCompleteness() bool`

HasCompleteness returns a boolean if a field has been set.

### GetChromosomeName

`func (o *V1reportsProkaryoteGeneLocation) GetChromosomeName() string`

GetChromosomeName returns the ChromosomeName field if non-nil, zero value otherwise.

### GetChromosomeNameOk

`func (o *V1reportsProkaryoteGeneLocation) GetChromosomeNameOk() (*string, bool)`

GetChromosomeNameOk returns a tuple with the ChromosomeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChromosomeName

`func (o *V1reportsProkaryoteGeneLocation) SetChromosomeName(v string)`

SetChromosomeName sets ChromosomeName field to given value.

### HasChromosomeName

`func (o *V1reportsProkaryoteGeneLocation) HasChromosomeName() bool`

HasChromosomeName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


