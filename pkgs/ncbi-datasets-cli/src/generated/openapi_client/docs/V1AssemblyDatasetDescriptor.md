# V1AssemblyDatasetDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssemblyAccession** | Pointer to **string** |  | [optional] 
**PairedAssemblyAccession** | Pointer to **string** |  | [optional] 
**BioprojectLineages** | Pointer to [**[]V1BioProjectLineage**](V1BioProjectLineage.md) |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Org** | Pointer to [**V1Organism**](V1Organism.md) |  | [optional] 
**Chromosomes** | Pointer to [**[]V1AssemblyDatasetDescriptorChromosome**](V1AssemblyDatasetDescriptorChromosome.md) |  | [optional] 
**AssemblyCategory** | Pointer to **string** |  | [optional] 
**AnnotationMetadata** | Pointer to [**V1AnnotationForAssembly**](V1AnnotationForAssembly.md) |  | [optional] 
**AssemblyLevel** | Pointer to **string** |  | [optional] 
**Submitter** | Pointer to **string** |  | [optional] 
**SubmissionDate** | Pointer to **string** |  | [optional] 
**ContigN50** | Pointer to **int32** |  | [optional] 
**EstimatedSize** | Pointer to **string** |  | [optional] 
**SeqLength** | Pointer to **string** |  | [optional] 
**BiosampleAccession** | Pointer to **string** |  | [optional] 
**Biosample** | Pointer to [**V1reportsBioSampleDescriptor**](V1reportsBioSampleDescriptor.md) |  | [optional] 
**BlastUrl** | Pointer to **string** |  | [optional] 
**GcCount** | Pointer to **string** |  | [optional] 

## Methods

### NewV1AssemblyDatasetDescriptor

`func NewV1AssemblyDatasetDescriptor() *V1AssemblyDatasetDescriptor`

NewV1AssemblyDatasetDescriptor instantiates a new V1AssemblyDatasetDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1AssemblyDatasetDescriptorWithDefaults

`func NewV1AssemblyDatasetDescriptorWithDefaults() *V1AssemblyDatasetDescriptor`

NewV1AssemblyDatasetDescriptorWithDefaults instantiates a new V1AssemblyDatasetDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssemblyAccession

`func (o *V1AssemblyDatasetDescriptor) GetAssemblyAccession() string`

GetAssemblyAccession returns the AssemblyAccession field if non-nil, zero value otherwise.

### GetAssemblyAccessionOk

`func (o *V1AssemblyDatasetDescriptor) GetAssemblyAccessionOk() (*string, bool)`

GetAssemblyAccessionOk returns a tuple with the AssemblyAccession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssemblyAccession

`func (o *V1AssemblyDatasetDescriptor) SetAssemblyAccession(v string)`

SetAssemblyAccession sets AssemblyAccession field to given value.

### HasAssemblyAccession

`func (o *V1AssemblyDatasetDescriptor) HasAssemblyAccession() bool`

HasAssemblyAccession returns a boolean if a field has been set.

### GetPairedAssemblyAccession

`func (o *V1AssemblyDatasetDescriptor) GetPairedAssemblyAccession() string`

GetPairedAssemblyAccession returns the PairedAssemblyAccession field if non-nil, zero value otherwise.

### GetPairedAssemblyAccessionOk

`func (o *V1AssemblyDatasetDescriptor) GetPairedAssemblyAccessionOk() (*string, bool)`

GetPairedAssemblyAccessionOk returns a tuple with the PairedAssemblyAccession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairedAssemblyAccession

`func (o *V1AssemblyDatasetDescriptor) SetPairedAssemblyAccession(v string)`

SetPairedAssemblyAccession sets PairedAssemblyAccession field to given value.

### HasPairedAssemblyAccession

`func (o *V1AssemblyDatasetDescriptor) HasPairedAssemblyAccession() bool`

HasPairedAssemblyAccession returns a boolean if a field has been set.

### GetBioprojectLineages

`func (o *V1AssemblyDatasetDescriptor) GetBioprojectLineages() []V1BioProjectLineage`

GetBioprojectLineages returns the BioprojectLineages field if non-nil, zero value otherwise.

### GetBioprojectLineagesOk

`func (o *V1AssemblyDatasetDescriptor) GetBioprojectLineagesOk() (*[]V1BioProjectLineage, bool)`

GetBioprojectLineagesOk returns a tuple with the BioprojectLineages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBioprojectLineages

`func (o *V1AssemblyDatasetDescriptor) SetBioprojectLineages(v []V1BioProjectLineage)`

SetBioprojectLineages sets BioprojectLineages field to given value.

### HasBioprojectLineages

`func (o *V1AssemblyDatasetDescriptor) HasBioprojectLineages() bool`

HasBioprojectLineages returns a boolean if a field has been set.

### GetDisplayName

`func (o *V1AssemblyDatasetDescriptor) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *V1AssemblyDatasetDescriptor) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *V1AssemblyDatasetDescriptor) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *V1AssemblyDatasetDescriptor) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetOrg

`func (o *V1AssemblyDatasetDescriptor) GetOrg() V1Organism`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *V1AssemblyDatasetDescriptor) GetOrgOk() (*V1Organism, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *V1AssemblyDatasetDescriptor) SetOrg(v V1Organism)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *V1AssemblyDatasetDescriptor) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetChromosomes

`func (o *V1AssemblyDatasetDescriptor) GetChromosomes() []V1AssemblyDatasetDescriptorChromosome`

GetChromosomes returns the Chromosomes field if non-nil, zero value otherwise.

### GetChromosomesOk

`func (o *V1AssemblyDatasetDescriptor) GetChromosomesOk() (*[]V1AssemblyDatasetDescriptorChromosome, bool)`

GetChromosomesOk returns a tuple with the Chromosomes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChromosomes

`func (o *V1AssemblyDatasetDescriptor) SetChromosomes(v []V1AssemblyDatasetDescriptorChromosome)`

SetChromosomes sets Chromosomes field to given value.

### HasChromosomes

`func (o *V1AssemblyDatasetDescriptor) HasChromosomes() bool`

HasChromosomes returns a boolean if a field has been set.

### GetAssemblyCategory

`func (o *V1AssemblyDatasetDescriptor) GetAssemblyCategory() string`

GetAssemblyCategory returns the AssemblyCategory field if non-nil, zero value otherwise.

### GetAssemblyCategoryOk

`func (o *V1AssemblyDatasetDescriptor) GetAssemblyCategoryOk() (*string, bool)`

GetAssemblyCategoryOk returns a tuple with the AssemblyCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssemblyCategory

`func (o *V1AssemblyDatasetDescriptor) SetAssemblyCategory(v string)`

SetAssemblyCategory sets AssemblyCategory field to given value.

### HasAssemblyCategory

`func (o *V1AssemblyDatasetDescriptor) HasAssemblyCategory() bool`

HasAssemblyCategory returns a boolean if a field has been set.

### GetAnnotationMetadata

`func (o *V1AssemblyDatasetDescriptor) GetAnnotationMetadata() V1AnnotationForAssembly`

GetAnnotationMetadata returns the AnnotationMetadata field if non-nil, zero value otherwise.

### GetAnnotationMetadataOk

`func (o *V1AssemblyDatasetDescriptor) GetAnnotationMetadataOk() (*V1AnnotationForAssembly, bool)`

GetAnnotationMetadataOk returns a tuple with the AnnotationMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotationMetadata

`func (o *V1AssemblyDatasetDescriptor) SetAnnotationMetadata(v V1AnnotationForAssembly)`

SetAnnotationMetadata sets AnnotationMetadata field to given value.

### HasAnnotationMetadata

`func (o *V1AssemblyDatasetDescriptor) HasAnnotationMetadata() bool`

HasAnnotationMetadata returns a boolean if a field has been set.

### GetAssemblyLevel

`func (o *V1AssemblyDatasetDescriptor) GetAssemblyLevel() string`

GetAssemblyLevel returns the AssemblyLevel field if non-nil, zero value otherwise.

### GetAssemblyLevelOk

`func (o *V1AssemblyDatasetDescriptor) GetAssemblyLevelOk() (*string, bool)`

GetAssemblyLevelOk returns a tuple with the AssemblyLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssemblyLevel

`func (o *V1AssemblyDatasetDescriptor) SetAssemblyLevel(v string)`

SetAssemblyLevel sets AssemblyLevel field to given value.

### HasAssemblyLevel

`func (o *V1AssemblyDatasetDescriptor) HasAssemblyLevel() bool`

HasAssemblyLevel returns a boolean if a field has been set.

### GetSubmitter

`func (o *V1AssemblyDatasetDescriptor) GetSubmitter() string`

GetSubmitter returns the Submitter field if non-nil, zero value otherwise.

### GetSubmitterOk

`func (o *V1AssemblyDatasetDescriptor) GetSubmitterOk() (*string, bool)`

GetSubmitterOk returns a tuple with the Submitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitter

`func (o *V1AssemblyDatasetDescriptor) SetSubmitter(v string)`

SetSubmitter sets Submitter field to given value.

### HasSubmitter

`func (o *V1AssemblyDatasetDescriptor) HasSubmitter() bool`

HasSubmitter returns a boolean if a field has been set.

### GetSubmissionDate

`func (o *V1AssemblyDatasetDescriptor) GetSubmissionDate() string`

GetSubmissionDate returns the SubmissionDate field if non-nil, zero value otherwise.

### GetSubmissionDateOk

`func (o *V1AssemblyDatasetDescriptor) GetSubmissionDateOk() (*string, bool)`

GetSubmissionDateOk returns a tuple with the SubmissionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionDate

`func (o *V1AssemblyDatasetDescriptor) SetSubmissionDate(v string)`

SetSubmissionDate sets SubmissionDate field to given value.

### HasSubmissionDate

`func (o *V1AssemblyDatasetDescriptor) HasSubmissionDate() bool`

HasSubmissionDate returns a boolean if a field has been set.

### GetContigN50

`func (o *V1AssemblyDatasetDescriptor) GetContigN50() int32`

GetContigN50 returns the ContigN50 field if non-nil, zero value otherwise.

### GetContigN50Ok

`func (o *V1AssemblyDatasetDescriptor) GetContigN50Ok() (*int32, bool)`

GetContigN50Ok returns a tuple with the ContigN50 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContigN50

`func (o *V1AssemblyDatasetDescriptor) SetContigN50(v int32)`

SetContigN50 sets ContigN50 field to given value.

### HasContigN50

`func (o *V1AssemblyDatasetDescriptor) HasContigN50() bool`

HasContigN50 returns a boolean if a field has been set.

### GetEstimatedSize

`func (o *V1AssemblyDatasetDescriptor) GetEstimatedSize() string`

GetEstimatedSize returns the EstimatedSize field if non-nil, zero value otherwise.

### GetEstimatedSizeOk

`func (o *V1AssemblyDatasetDescriptor) GetEstimatedSizeOk() (*string, bool)`

GetEstimatedSizeOk returns a tuple with the EstimatedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedSize

`func (o *V1AssemblyDatasetDescriptor) SetEstimatedSize(v string)`

SetEstimatedSize sets EstimatedSize field to given value.

### HasEstimatedSize

`func (o *V1AssemblyDatasetDescriptor) HasEstimatedSize() bool`

HasEstimatedSize returns a boolean if a field has been set.

### GetSeqLength

`func (o *V1AssemblyDatasetDescriptor) GetSeqLength() string`

GetSeqLength returns the SeqLength field if non-nil, zero value otherwise.

### GetSeqLengthOk

`func (o *V1AssemblyDatasetDescriptor) GetSeqLengthOk() (*string, bool)`

GetSeqLengthOk returns a tuple with the SeqLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeqLength

`func (o *V1AssemblyDatasetDescriptor) SetSeqLength(v string)`

SetSeqLength sets SeqLength field to given value.

### HasSeqLength

`func (o *V1AssemblyDatasetDescriptor) HasSeqLength() bool`

HasSeqLength returns a boolean if a field has been set.

### GetBiosampleAccession

`func (o *V1AssemblyDatasetDescriptor) GetBiosampleAccession() string`

GetBiosampleAccession returns the BiosampleAccession field if non-nil, zero value otherwise.

### GetBiosampleAccessionOk

`func (o *V1AssemblyDatasetDescriptor) GetBiosampleAccessionOk() (*string, bool)`

GetBiosampleAccessionOk returns a tuple with the BiosampleAccession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosampleAccession

`func (o *V1AssemblyDatasetDescriptor) SetBiosampleAccession(v string)`

SetBiosampleAccession sets BiosampleAccession field to given value.

### HasBiosampleAccession

`func (o *V1AssemblyDatasetDescriptor) HasBiosampleAccession() bool`

HasBiosampleAccession returns a boolean if a field has been set.

### GetBiosample

`func (o *V1AssemblyDatasetDescriptor) GetBiosample() V1reportsBioSampleDescriptor`

GetBiosample returns the Biosample field if non-nil, zero value otherwise.

### GetBiosampleOk

`func (o *V1AssemblyDatasetDescriptor) GetBiosampleOk() (*V1reportsBioSampleDescriptor, bool)`

GetBiosampleOk returns a tuple with the Biosample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosample

`func (o *V1AssemblyDatasetDescriptor) SetBiosample(v V1reportsBioSampleDescriptor)`

SetBiosample sets Biosample field to given value.

### HasBiosample

`func (o *V1AssemblyDatasetDescriptor) HasBiosample() bool`

HasBiosample returns a boolean if a field has been set.

### GetBlastUrl

`func (o *V1AssemblyDatasetDescriptor) GetBlastUrl() string`

GetBlastUrl returns the BlastUrl field if non-nil, zero value otherwise.

### GetBlastUrlOk

`func (o *V1AssemblyDatasetDescriptor) GetBlastUrlOk() (*string, bool)`

GetBlastUrlOk returns a tuple with the BlastUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlastUrl

`func (o *V1AssemblyDatasetDescriptor) SetBlastUrl(v string)`

SetBlastUrl sets BlastUrl field to given value.

### HasBlastUrl

`func (o *V1AssemblyDatasetDescriptor) HasBlastUrl() bool`

HasBlastUrl returns a boolean if a field has been set.

### GetGcCount

`func (o *V1AssemblyDatasetDescriptor) GetGcCount() string`

GetGcCount returns the GcCount field if non-nil, zero value otherwise.

### GetGcCountOk

`func (o *V1AssemblyDatasetDescriptor) GetGcCountOk() (*string, bool)`

GetGcCountOk returns a tuple with the GcCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcCount

`func (o *V1AssemblyDatasetDescriptor) SetGcCount(v string)`

SetGcCount sets GcCount field to given value.

### HasGcCount

`func (o *V1AssemblyDatasetDescriptor) HasGcCount() bool`

HasGcCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


