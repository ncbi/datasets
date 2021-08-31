# V1GeneDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GeneId** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**TaxId** | Pointer to **string** |  | [optional] 
**Taxname** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**V1GeneDescriptorGeneType**](V1GeneDescriptorGeneType.md) |  | [optional] [default to V1GENEDESCRIPTORGENETYPE_UNKNOWN]
**RnaType** | Pointer to [**V1GeneDescriptorRnaType**](V1GeneDescriptorRnaType.md) |  | [optional] [default to V1GENEDESCRIPTORRNATYPE_RNA_UNKNOWN]
**Orientation** | Pointer to [**V1Orientation**](V1Orientation.md) |  | [optional] [default to V1ORIENTATION_NONE]
**GenomicRanges** | Pointer to [**[]V1SeqRangeSet**](V1SeqRangeSet.md) |  | [optional] 
**ReferenceStandards** | Pointer to [**[]V1GenomicRegion**](V1GenomicRegion.md) |  | [optional] 
**GenomicRegions** | Pointer to [**[]V1GenomicRegion**](V1GenomicRegion.md) |  | [optional] 
**Transcripts** | Pointer to [**[]V1Transcript**](V1Transcript.md) |  | [optional] 
**Proteins** | Pointer to [**[]V1Protein**](V1Protein.md) |  | [optional] 
**CommonName** | Pointer to **string** |  | [optional] 
**Chromosome** | Pointer to **string** |  | [optional] 
**Chromosomes** | Pointer to **[]string** |  | [optional] 
**NomenclatureAuthority** | Pointer to [**V1NomenclatureAuthority**](V1NomenclatureAuthority.md) |  | [optional] 
**SwissProtAccessions** | Pointer to **[]string** |  | [optional] 
**EnsemblGeneIds** | Pointer to **[]string** |  | [optional] 
**OmimIds** | Pointer to **[]string** |  | [optional] 
**Synonyms** | Pointer to **[]string** |  | [optional] 
**ReplacedGeneId** | Pointer to **string** |  | [optional] 
**Annotations** | Pointer to [**[]V1Annotation**](V1Annotation.md) |  | [optional] 

## Methods

### NewV1GeneDescriptor

`func NewV1GeneDescriptor() *V1GeneDescriptor`

NewV1GeneDescriptor instantiates a new V1GeneDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1GeneDescriptorWithDefaults

`func NewV1GeneDescriptorWithDefaults() *V1GeneDescriptor`

NewV1GeneDescriptorWithDefaults instantiates a new V1GeneDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeneId

`func (o *V1GeneDescriptor) GetGeneId() string`

GetGeneId returns the GeneId field if non-nil, zero value otherwise.

### GetGeneIdOk

`func (o *V1GeneDescriptor) GetGeneIdOk() (*string, bool)`

GetGeneIdOk returns a tuple with the GeneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneId

`func (o *V1GeneDescriptor) SetGeneId(v string)`

SetGeneId sets GeneId field to given value.

### HasGeneId

`func (o *V1GeneDescriptor) HasGeneId() bool`

HasGeneId returns a boolean if a field has been set.

### GetSymbol

`func (o *V1GeneDescriptor) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *V1GeneDescriptor) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *V1GeneDescriptor) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *V1GeneDescriptor) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetDescription

`func (o *V1GeneDescriptor) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V1GeneDescriptor) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V1GeneDescriptor) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V1GeneDescriptor) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTaxId

`func (o *V1GeneDescriptor) GetTaxId() string`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *V1GeneDescriptor) GetTaxIdOk() (*string, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *V1GeneDescriptor) SetTaxId(v string)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *V1GeneDescriptor) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.

### GetTaxname

`func (o *V1GeneDescriptor) GetTaxname() string`

GetTaxname returns the Taxname field if non-nil, zero value otherwise.

### GetTaxnameOk

`func (o *V1GeneDescriptor) GetTaxnameOk() (*string, bool)`

GetTaxnameOk returns a tuple with the Taxname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxname

`func (o *V1GeneDescriptor) SetTaxname(v string)`

SetTaxname sets Taxname field to given value.

### HasTaxname

`func (o *V1GeneDescriptor) HasTaxname() bool`

HasTaxname returns a boolean if a field has been set.

### GetType

`func (o *V1GeneDescriptor) GetType() V1GeneDescriptorGeneType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1GeneDescriptor) GetTypeOk() (*V1GeneDescriptorGeneType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1GeneDescriptor) SetType(v V1GeneDescriptorGeneType)`

SetType sets Type field to given value.

### HasType

`func (o *V1GeneDescriptor) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRnaType

`func (o *V1GeneDescriptor) GetRnaType() V1GeneDescriptorRnaType`

GetRnaType returns the RnaType field if non-nil, zero value otherwise.

### GetRnaTypeOk

`func (o *V1GeneDescriptor) GetRnaTypeOk() (*V1GeneDescriptorRnaType, bool)`

GetRnaTypeOk returns a tuple with the RnaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRnaType

`func (o *V1GeneDescriptor) SetRnaType(v V1GeneDescriptorRnaType)`

SetRnaType sets RnaType field to given value.

### HasRnaType

`func (o *V1GeneDescriptor) HasRnaType() bool`

HasRnaType returns a boolean if a field has been set.

### GetOrientation

`func (o *V1GeneDescriptor) GetOrientation() V1Orientation`

GetOrientation returns the Orientation field if non-nil, zero value otherwise.

### GetOrientationOk

`func (o *V1GeneDescriptor) GetOrientationOk() (*V1Orientation, bool)`

GetOrientationOk returns a tuple with the Orientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrientation

`func (o *V1GeneDescriptor) SetOrientation(v V1Orientation)`

SetOrientation sets Orientation field to given value.

### HasOrientation

`func (o *V1GeneDescriptor) HasOrientation() bool`

HasOrientation returns a boolean if a field has been set.

### GetGenomicRanges

`func (o *V1GeneDescriptor) GetGenomicRanges() []V1SeqRangeSet`

GetGenomicRanges returns the GenomicRanges field if non-nil, zero value otherwise.

### GetGenomicRangesOk

`func (o *V1GeneDescriptor) GetGenomicRangesOk() (*[]V1SeqRangeSet, bool)`

GetGenomicRangesOk returns a tuple with the GenomicRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenomicRanges

`func (o *V1GeneDescriptor) SetGenomicRanges(v []V1SeqRangeSet)`

SetGenomicRanges sets GenomicRanges field to given value.

### HasGenomicRanges

`func (o *V1GeneDescriptor) HasGenomicRanges() bool`

HasGenomicRanges returns a boolean if a field has been set.

### GetReferenceStandards

`func (o *V1GeneDescriptor) GetReferenceStandards() []V1GenomicRegion`

GetReferenceStandards returns the ReferenceStandards field if non-nil, zero value otherwise.

### GetReferenceStandardsOk

`func (o *V1GeneDescriptor) GetReferenceStandardsOk() (*[]V1GenomicRegion, bool)`

GetReferenceStandardsOk returns a tuple with the ReferenceStandards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceStandards

`func (o *V1GeneDescriptor) SetReferenceStandards(v []V1GenomicRegion)`

SetReferenceStandards sets ReferenceStandards field to given value.

### HasReferenceStandards

`func (o *V1GeneDescriptor) HasReferenceStandards() bool`

HasReferenceStandards returns a boolean if a field has been set.

### GetGenomicRegions

`func (o *V1GeneDescriptor) GetGenomicRegions() []V1GenomicRegion`

GetGenomicRegions returns the GenomicRegions field if non-nil, zero value otherwise.

### GetGenomicRegionsOk

`func (o *V1GeneDescriptor) GetGenomicRegionsOk() (*[]V1GenomicRegion, bool)`

GetGenomicRegionsOk returns a tuple with the GenomicRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenomicRegions

`func (o *V1GeneDescriptor) SetGenomicRegions(v []V1GenomicRegion)`

SetGenomicRegions sets GenomicRegions field to given value.

### HasGenomicRegions

`func (o *V1GeneDescriptor) HasGenomicRegions() bool`

HasGenomicRegions returns a boolean if a field has been set.

### GetTranscripts

`func (o *V1GeneDescriptor) GetTranscripts() []V1Transcript`

GetTranscripts returns the Transcripts field if non-nil, zero value otherwise.

### GetTranscriptsOk

`func (o *V1GeneDescriptor) GetTranscriptsOk() (*[]V1Transcript, bool)`

GetTranscriptsOk returns a tuple with the Transcripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscripts

`func (o *V1GeneDescriptor) SetTranscripts(v []V1Transcript)`

SetTranscripts sets Transcripts field to given value.

### HasTranscripts

`func (o *V1GeneDescriptor) HasTranscripts() bool`

HasTranscripts returns a boolean if a field has been set.

### GetProteins

`func (o *V1GeneDescriptor) GetProteins() []V1Protein`

GetProteins returns the Proteins field if non-nil, zero value otherwise.

### GetProteinsOk

`func (o *V1GeneDescriptor) GetProteinsOk() (*[]V1Protein, bool)`

GetProteinsOk returns a tuple with the Proteins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProteins

`func (o *V1GeneDescriptor) SetProteins(v []V1Protein)`

SetProteins sets Proteins field to given value.

### HasProteins

`func (o *V1GeneDescriptor) HasProteins() bool`

HasProteins returns a boolean if a field has been set.

### GetCommonName

`func (o *V1GeneDescriptor) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *V1GeneDescriptor) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *V1GeneDescriptor) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *V1GeneDescriptor) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### GetChromosome

`func (o *V1GeneDescriptor) GetChromosome() string`

GetChromosome returns the Chromosome field if non-nil, zero value otherwise.

### GetChromosomeOk

`func (o *V1GeneDescriptor) GetChromosomeOk() (*string, bool)`

GetChromosomeOk returns a tuple with the Chromosome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChromosome

`func (o *V1GeneDescriptor) SetChromosome(v string)`

SetChromosome sets Chromosome field to given value.

### HasChromosome

`func (o *V1GeneDescriptor) HasChromosome() bool`

HasChromosome returns a boolean if a field has been set.

### GetChromosomes

`func (o *V1GeneDescriptor) GetChromosomes() []string`

GetChromosomes returns the Chromosomes field if non-nil, zero value otherwise.

### GetChromosomesOk

`func (o *V1GeneDescriptor) GetChromosomesOk() (*[]string, bool)`

GetChromosomesOk returns a tuple with the Chromosomes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChromosomes

`func (o *V1GeneDescriptor) SetChromosomes(v []string)`

SetChromosomes sets Chromosomes field to given value.

### HasChromosomes

`func (o *V1GeneDescriptor) HasChromosomes() bool`

HasChromosomes returns a boolean if a field has been set.

### GetNomenclatureAuthority

`func (o *V1GeneDescriptor) GetNomenclatureAuthority() V1NomenclatureAuthority`

GetNomenclatureAuthority returns the NomenclatureAuthority field if non-nil, zero value otherwise.

### GetNomenclatureAuthorityOk

`func (o *V1GeneDescriptor) GetNomenclatureAuthorityOk() (*V1NomenclatureAuthority, bool)`

GetNomenclatureAuthorityOk returns a tuple with the NomenclatureAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNomenclatureAuthority

`func (o *V1GeneDescriptor) SetNomenclatureAuthority(v V1NomenclatureAuthority)`

SetNomenclatureAuthority sets NomenclatureAuthority field to given value.

### HasNomenclatureAuthority

`func (o *V1GeneDescriptor) HasNomenclatureAuthority() bool`

HasNomenclatureAuthority returns a boolean if a field has been set.

### GetSwissProtAccessions

`func (o *V1GeneDescriptor) GetSwissProtAccessions() []string`

GetSwissProtAccessions returns the SwissProtAccessions field if non-nil, zero value otherwise.

### GetSwissProtAccessionsOk

`func (o *V1GeneDescriptor) GetSwissProtAccessionsOk() (*[]string, bool)`

GetSwissProtAccessionsOk returns a tuple with the SwissProtAccessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwissProtAccessions

`func (o *V1GeneDescriptor) SetSwissProtAccessions(v []string)`

SetSwissProtAccessions sets SwissProtAccessions field to given value.

### HasSwissProtAccessions

`func (o *V1GeneDescriptor) HasSwissProtAccessions() bool`

HasSwissProtAccessions returns a boolean if a field has been set.

### GetEnsemblGeneIds

`func (o *V1GeneDescriptor) GetEnsemblGeneIds() []string`

GetEnsemblGeneIds returns the EnsemblGeneIds field if non-nil, zero value otherwise.

### GetEnsemblGeneIdsOk

`func (o *V1GeneDescriptor) GetEnsemblGeneIdsOk() (*[]string, bool)`

GetEnsemblGeneIdsOk returns a tuple with the EnsemblGeneIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnsemblGeneIds

`func (o *V1GeneDescriptor) SetEnsemblGeneIds(v []string)`

SetEnsemblGeneIds sets EnsemblGeneIds field to given value.

### HasEnsemblGeneIds

`func (o *V1GeneDescriptor) HasEnsemblGeneIds() bool`

HasEnsemblGeneIds returns a boolean if a field has been set.

### GetOmimIds

`func (o *V1GeneDescriptor) GetOmimIds() []string`

GetOmimIds returns the OmimIds field if non-nil, zero value otherwise.

### GetOmimIdsOk

`func (o *V1GeneDescriptor) GetOmimIdsOk() (*[]string, bool)`

GetOmimIdsOk returns a tuple with the OmimIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmimIds

`func (o *V1GeneDescriptor) SetOmimIds(v []string)`

SetOmimIds sets OmimIds field to given value.

### HasOmimIds

`func (o *V1GeneDescriptor) HasOmimIds() bool`

HasOmimIds returns a boolean if a field has been set.

### GetSynonyms

`func (o *V1GeneDescriptor) GetSynonyms() []string`

GetSynonyms returns the Synonyms field if non-nil, zero value otherwise.

### GetSynonymsOk

`func (o *V1GeneDescriptor) GetSynonymsOk() (*[]string, bool)`

GetSynonymsOk returns a tuple with the Synonyms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynonyms

`func (o *V1GeneDescriptor) SetSynonyms(v []string)`

SetSynonyms sets Synonyms field to given value.

### HasSynonyms

`func (o *V1GeneDescriptor) HasSynonyms() bool`

HasSynonyms returns a boolean if a field has been set.

### GetReplacedGeneId

`func (o *V1GeneDescriptor) GetReplacedGeneId() string`

GetReplacedGeneId returns the ReplacedGeneId field if non-nil, zero value otherwise.

### GetReplacedGeneIdOk

`func (o *V1GeneDescriptor) GetReplacedGeneIdOk() (*string, bool)`

GetReplacedGeneIdOk returns a tuple with the ReplacedGeneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacedGeneId

`func (o *V1GeneDescriptor) SetReplacedGeneId(v string)`

SetReplacedGeneId sets ReplacedGeneId field to given value.

### HasReplacedGeneId

`func (o *V1GeneDescriptor) HasReplacedGeneId() bool`

HasReplacedGeneId returns a boolean if a field has been set.

### GetAnnotations

`func (o *V1GeneDescriptor) GetAnnotations() []V1Annotation`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *V1GeneDescriptor) GetAnnotationsOk() (*[]V1Annotation, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *V1GeneDescriptor) SetAnnotations(v []V1Annotation)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *V1GeneDescriptor) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


