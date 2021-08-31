# V1reportsGeneDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GeneId** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**TaxId** | Pointer to **string** |  | [optional] 
**Taxname** | Pointer to **string** |  | [optional] 
**CommonName** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**V1reportsGeneDescriptorGeneType**](V1reportsGeneDescriptorGeneType.md) |  | [optional] [default to V1REPORTSGENEDESCRIPTORGENETYPE_UNKNOWN]
**RnaType** | Pointer to [**V1reportsGeneDescriptorRnaType**](V1reportsGeneDescriptorRnaType.md) |  | [optional] [default to V1REPORTSGENEDESCRIPTORRNATYPE_RNA_UNKNOWN]
**Orientation** | Pointer to [**V1reportsOrientation**](V1reportsOrientation.md) |  | [optional] [default to V1REPORTSORIENTATION_NONE]
**GenomicRanges** | Pointer to [**[]V1reportsSeqRangeSet**](V1reportsSeqRangeSet.md) |  | [optional] 
**ReferenceStandards** | Pointer to [**[]V1reportsGenomicRegion**](V1reportsGenomicRegion.md) |  | [optional] 
**GenomicRegions** | Pointer to [**[]V1reportsGenomicRegion**](V1reportsGenomicRegion.md) |  | [optional] 
**Transcripts** | Pointer to [**[]V1reportsTranscript**](V1reportsTranscript.md) |  | [optional] 
**Proteins** | Pointer to [**[]V1reportsProtein**](V1reportsProtein.md) |  | [optional] 
**Chromosome** | Pointer to **string** |  | [optional] 
**Chromosomes** | Pointer to **[]string** |  | [optional] 
**NomenclatureAuthority** | Pointer to [**V1reportsNomenclatureAuthority**](V1reportsNomenclatureAuthority.md) |  | [optional] 
**SwissProtAccessions** | Pointer to **[]string** |  | [optional] 
**EnsemblGeneIds** | Pointer to **[]string** |  | [optional] 
**OmimIds** | Pointer to **[]string** |  | [optional] 
**Synonyms** | Pointer to **[]string** |  | [optional] 
**ReplacedGeneId** | Pointer to **string** |  | [optional] 
**Annotations** | Pointer to [**[]V1reportsAnnotation**](V1reportsAnnotation.md) |  | [optional] 

## Methods

### NewV1reportsGeneDescriptor

`func NewV1reportsGeneDescriptor() *V1reportsGeneDescriptor`

NewV1reportsGeneDescriptor instantiates a new V1reportsGeneDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1reportsGeneDescriptorWithDefaults

`func NewV1reportsGeneDescriptorWithDefaults() *V1reportsGeneDescriptor`

NewV1reportsGeneDescriptorWithDefaults instantiates a new V1reportsGeneDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeneId

`func (o *V1reportsGeneDescriptor) GetGeneId() string`

GetGeneId returns the GeneId field if non-nil, zero value otherwise.

### GetGeneIdOk

`func (o *V1reportsGeneDescriptor) GetGeneIdOk() (*string, bool)`

GetGeneIdOk returns a tuple with the GeneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneId

`func (o *V1reportsGeneDescriptor) SetGeneId(v string)`

SetGeneId sets GeneId field to given value.

### HasGeneId

`func (o *V1reportsGeneDescriptor) HasGeneId() bool`

HasGeneId returns a boolean if a field has been set.

### GetSymbol

`func (o *V1reportsGeneDescriptor) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *V1reportsGeneDescriptor) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *V1reportsGeneDescriptor) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *V1reportsGeneDescriptor) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetDescription

`func (o *V1reportsGeneDescriptor) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V1reportsGeneDescriptor) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V1reportsGeneDescriptor) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V1reportsGeneDescriptor) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTaxId

`func (o *V1reportsGeneDescriptor) GetTaxId() string`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *V1reportsGeneDescriptor) GetTaxIdOk() (*string, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *V1reportsGeneDescriptor) SetTaxId(v string)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *V1reportsGeneDescriptor) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.

### GetTaxname

`func (o *V1reportsGeneDescriptor) GetTaxname() string`

GetTaxname returns the Taxname field if non-nil, zero value otherwise.

### GetTaxnameOk

`func (o *V1reportsGeneDescriptor) GetTaxnameOk() (*string, bool)`

GetTaxnameOk returns a tuple with the Taxname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxname

`func (o *V1reportsGeneDescriptor) SetTaxname(v string)`

SetTaxname sets Taxname field to given value.

### HasTaxname

`func (o *V1reportsGeneDescriptor) HasTaxname() bool`

HasTaxname returns a boolean if a field has been set.

### GetCommonName

`func (o *V1reportsGeneDescriptor) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *V1reportsGeneDescriptor) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *V1reportsGeneDescriptor) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *V1reportsGeneDescriptor) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### GetType

`func (o *V1reportsGeneDescriptor) GetType() V1reportsGeneDescriptorGeneType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1reportsGeneDescriptor) GetTypeOk() (*V1reportsGeneDescriptorGeneType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1reportsGeneDescriptor) SetType(v V1reportsGeneDescriptorGeneType)`

SetType sets Type field to given value.

### HasType

`func (o *V1reportsGeneDescriptor) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRnaType

`func (o *V1reportsGeneDescriptor) GetRnaType() V1reportsGeneDescriptorRnaType`

GetRnaType returns the RnaType field if non-nil, zero value otherwise.

### GetRnaTypeOk

`func (o *V1reportsGeneDescriptor) GetRnaTypeOk() (*V1reportsGeneDescriptorRnaType, bool)`

GetRnaTypeOk returns a tuple with the RnaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRnaType

`func (o *V1reportsGeneDescriptor) SetRnaType(v V1reportsGeneDescriptorRnaType)`

SetRnaType sets RnaType field to given value.

### HasRnaType

`func (o *V1reportsGeneDescriptor) HasRnaType() bool`

HasRnaType returns a boolean if a field has been set.

### GetOrientation

`func (o *V1reportsGeneDescriptor) GetOrientation() V1reportsOrientation`

GetOrientation returns the Orientation field if non-nil, zero value otherwise.

### GetOrientationOk

`func (o *V1reportsGeneDescriptor) GetOrientationOk() (*V1reportsOrientation, bool)`

GetOrientationOk returns a tuple with the Orientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrientation

`func (o *V1reportsGeneDescriptor) SetOrientation(v V1reportsOrientation)`

SetOrientation sets Orientation field to given value.

### HasOrientation

`func (o *V1reportsGeneDescriptor) HasOrientation() bool`

HasOrientation returns a boolean if a field has been set.

### GetGenomicRanges

`func (o *V1reportsGeneDescriptor) GetGenomicRanges() []V1reportsSeqRangeSet`

GetGenomicRanges returns the GenomicRanges field if non-nil, zero value otherwise.

### GetGenomicRangesOk

`func (o *V1reportsGeneDescriptor) GetGenomicRangesOk() (*[]V1reportsSeqRangeSet, bool)`

GetGenomicRangesOk returns a tuple with the GenomicRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenomicRanges

`func (o *V1reportsGeneDescriptor) SetGenomicRanges(v []V1reportsSeqRangeSet)`

SetGenomicRanges sets GenomicRanges field to given value.

### HasGenomicRanges

`func (o *V1reportsGeneDescriptor) HasGenomicRanges() bool`

HasGenomicRanges returns a boolean if a field has been set.

### GetReferenceStandards

`func (o *V1reportsGeneDescriptor) GetReferenceStandards() []V1reportsGenomicRegion`

GetReferenceStandards returns the ReferenceStandards field if non-nil, zero value otherwise.

### GetReferenceStandardsOk

`func (o *V1reportsGeneDescriptor) GetReferenceStandardsOk() (*[]V1reportsGenomicRegion, bool)`

GetReferenceStandardsOk returns a tuple with the ReferenceStandards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceStandards

`func (o *V1reportsGeneDescriptor) SetReferenceStandards(v []V1reportsGenomicRegion)`

SetReferenceStandards sets ReferenceStandards field to given value.

### HasReferenceStandards

`func (o *V1reportsGeneDescriptor) HasReferenceStandards() bool`

HasReferenceStandards returns a boolean if a field has been set.

### GetGenomicRegions

`func (o *V1reportsGeneDescriptor) GetGenomicRegions() []V1reportsGenomicRegion`

GetGenomicRegions returns the GenomicRegions field if non-nil, zero value otherwise.

### GetGenomicRegionsOk

`func (o *V1reportsGeneDescriptor) GetGenomicRegionsOk() (*[]V1reportsGenomicRegion, bool)`

GetGenomicRegionsOk returns a tuple with the GenomicRegions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenomicRegions

`func (o *V1reportsGeneDescriptor) SetGenomicRegions(v []V1reportsGenomicRegion)`

SetGenomicRegions sets GenomicRegions field to given value.

### HasGenomicRegions

`func (o *V1reportsGeneDescriptor) HasGenomicRegions() bool`

HasGenomicRegions returns a boolean if a field has been set.

### GetTranscripts

`func (o *V1reportsGeneDescriptor) GetTranscripts() []V1reportsTranscript`

GetTranscripts returns the Transcripts field if non-nil, zero value otherwise.

### GetTranscriptsOk

`func (o *V1reportsGeneDescriptor) GetTranscriptsOk() (*[]V1reportsTranscript, bool)`

GetTranscriptsOk returns a tuple with the Transcripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscripts

`func (o *V1reportsGeneDescriptor) SetTranscripts(v []V1reportsTranscript)`

SetTranscripts sets Transcripts field to given value.

### HasTranscripts

`func (o *V1reportsGeneDescriptor) HasTranscripts() bool`

HasTranscripts returns a boolean if a field has been set.

### GetProteins

`func (o *V1reportsGeneDescriptor) GetProteins() []V1reportsProtein`

GetProteins returns the Proteins field if non-nil, zero value otherwise.

### GetProteinsOk

`func (o *V1reportsGeneDescriptor) GetProteinsOk() (*[]V1reportsProtein, bool)`

GetProteinsOk returns a tuple with the Proteins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProteins

`func (o *V1reportsGeneDescriptor) SetProteins(v []V1reportsProtein)`

SetProteins sets Proteins field to given value.

### HasProteins

`func (o *V1reportsGeneDescriptor) HasProteins() bool`

HasProteins returns a boolean if a field has been set.

### GetChromosome

`func (o *V1reportsGeneDescriptor) GetChromosome() string`

GetChromosome returns the Chromosome field if non-nil, zero value otherwise.

### GetChromosomeOk

`func (o *V1reportsGeneDescriptor) GetChromosomeOk() (*string, bool)`

GetChromosomeOk returns a tuple with the Chromosome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChromosome

`func (o *V1reportsGeneDescriptor) SetChromosome(v string)`

SetChromosome sets Chromosome field to given value.

### HasChromosome

`func (o *V1reportsGeneDescriptor) HasChromosome() bool`

HasChromosome returns a boolean if a field has been set.

### GetChromosomes

`func (o *V1reportsGeneDescriptor) GetChromosomes() []string`

GetChromosomes returns the Chromosomes field if non-nil, zero value otherwise.

### GetChromosomesOk

`func (o *V1reportsGeneDescriptor) GetChromosomesOk() (*[]string, bool)`

GetChromosomesOk returns a tuple with the Chromosomes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChromosomes

`func (o *V1reportsGeneDescriptor) SetChromosomes(v []string)`

SetChromosomes sets Chromosomes field to given value.

### HasChromosomes

`func (o *V1reportsGeneDescriptor) HasChromosomes() bool`

HasChromosomes returns a boolean if a field has been set.

### GetNomenclatureAuthority

`func (o *V1reportsGeneDescriptor) GetNomenclatureAuthority() V1reportsNomenclatureAuthority`

GetNomenclatureAuthority returns the NomenclatureAuthority field if non-nil, zero value otherwise.

### GetNomenclatureAuthorityOk

`func (o *V1reportsGeneDescriptor) GetNomenclatureAuthorityOk() (*V1reportsNomenclatureAuthority, bool)`

GetNomenclatureAuthorityOk returns a tuple with the NomenclatureAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNomenclatureAuthority

`func (o *V1reportsGeneDescriptor) SetNomenclatureAuthority(v V1reportsNomenclatureAuthority)`

SetNomenclatureAuthority sets NomenclatureAuthority field to given value.

### HasNomenclatureAuthority

`func (o *V1reportsGeneDescriptor) HasNomenclatureAuthority() bool`

HasNomenclatureAuthority returns a boolean if a field has been set.

### GetSwissProtAccessions

`func (o *V1reportsGeneDescriptor) GetSwissProtAccessions() []string`

GetSwissProtAccessions returns the SwissProtAccessions field if non-nil, zero value otherwise.

### GetSwissProtAccessionsOk

`func (o *V1reportsGeneDescriptor) GetSwissProtAccessionsOk() (*[]string, bool)`

GetSwissProtAccessionsOk returns a tuple with the SwissProtAccessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwissProtAccessions

`func (o *V1reportsGeneDescriptor) SetSwissProtAccessions(v []string)`

SetSwissProtAccessions sets SwissProtAccessions field to given value.

### HasSwissProtAccessions

`func (o *V1reportsGeneDescriptor) HasSwissProtAccessions() bool`

HasSwissProtAccessions returns a boolean if a field has been set.

### GetEnsemblGeneIds

`func (o *V1reportsGeneDescriptor) GetEnsemblGeneIds() []string`

GetEnsemblGeneIds returns the EnsemblGeneIds field if non-nil, zero value otherwise.

### GetEnsemblGeneIdsOk

`func (o *V1reportsGeneDescriptor) GetEnsemblGeneIdsOk() (*[]string, bool)`

GetEnsemblGeneIdsOk returns a tuple with the EnsemblGeneIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnsemblGeneIds

`func (o *V1reportsGeneDescriptor) SetEnsemblGeneIds(v []string)`

SetEnsemblGeneIds sets EnsemblGeneIds field to given value.

### HasEnsemblGeneIds

`func (o *V1reportsGeneDescriptor) HasEnsemblGeneIds() bool`

HasEnsemblGeneIds returns a boolean if a field has been set.

### GetOmimIds

`func (o *V1reportsGeneDescriptor) GetOmimIds() []string`

GetOmimIds returns the OmimIds field if non-nil, zero value otherwise.

### GetOmimIdsOk

`func (o *V1reportsGeneDescriptor) GetOmimIdsOk() (*[]string, bool)`

GetOmimIdsOk returns a tuple with the OmimIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmimIds

`func (o *V1reportsGeneDescriptor) SetOmimIds(v []string)`

SetOmimIds sets OmimIds field to given value.

### HasOmimIds

`func (o *V1reportsGeneDescriptor) HasOmimIds() bool`

HasOmimIds returns a boolean if a field has been set.

### GetSynonyms

`func (o *V1reportsGeneDescriptor) GetSynonyms() []string`

GetSynonyms returns the Synonyms field if non-nil, zero value otherwise.

### GetSynonymsOk

`func (o *V1reportsGeneDescriptor) GetSynonymsOk() (*[]string, bool)`

GetSynonymsOk returns a tuple with the Synonyms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynonyms

`func (o *V1reportsGeneDescriptor) SetSynonyms(v []string)`

SetSynonyms sets Synonyms field to given value.

### HasSynonyms

`func (o *V1reportsGeneDescriptor) HasSynonyms() bool`

HasSynonyms returns a boolean if a field has been set.

### GetReplacedGeneId

`func (o *V1reportsGeneDescriptor) GetReplacedGeneId() string`

GetReplacedGeneId returns the ReplacedGeneId field if non-nil, zero value otherwise.

### GetReplacedGeneIdOk

`func (o *V1reportsGeneDescriptor) GetReplacedGeneIdOk() (*string, bool)`

GetReplacedGeneIdOk returns a tuple with the ReplacedGeneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacedGeneId

`func (o *V1reportsGeneDescriptor) SetReplacedGeneId(v string)`

SetReplacedGeneId sets ReplacedGeneId field to given value.

### HasReplacedGeneId

`func (o *V1reportsGeneDescriptor) HasReplacedGeneId() bool`

HasReplacedGeneId returns a boolean if a field has been set.

### GetAnnotations

`func (o *V1reportsGeneDescriptor) GetAnnotations() []V1reportsAnnotation`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *V1reportsGeneDescriptor) GetAnnotationsOk() (*[]V1reportsAnnotation, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *V1reportsGeneDescriptor) SetAnnotations(v []V1reportsAnnotation)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *V1reportsGeneDescriptor) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


