# V1Transcript

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessionVersion** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Length** | Pointer to **int32** |  | [optional] 
**GenomicRange** | Pointer to [**V1SeqRangeSet**](V1SeqRangeSet.md) |  | [optional] 
**Exons** | Pointer to [**V1SeqRangeSet**](V1SeqRangeSet.md) |  | [optional] 
**Cds** | Pointer to [**V1SeqRangeSet**](V1SeqRangeSet.md) |  | [optional] 
**GenomicLocations** | Pointer to [**[]V1GenomicLocation**](V1GenomicLocation.md) |  | [optional] 
**EnsemblTranscript** | Pointer to **string** |  | [optional] 
**Protein** | Pointer to [**V1Protein**](V1Protein.md) |  | [optional] 
**Type** | Pointer to [**V1TranscriptTranscriptType**](V1TranscriptTranscriptType.md) |  | [optional] [default to V1TRANSCRIPTTRANSCRIPTTYPE_UNKNOWN]

## Methods

### NewV1Transcript

`func NewV1Transcript() *V1Transcript`

NewV1Transcript instantiates a new V1Transcript object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TranscriptWithDefaults

`func NewV1TranscriptWithDefaults() *V1Transcript`

NewV1TranscriptWithDefaults instantiates a new V1Transcript object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessionVersion

`func (o *V1Transcript) GetAccessionVersion() string`

GetAccessionVersion returns the AccessionVersion field if non-nil, zero value otherwise.

### GetAccessionVersionOk

`func (o *V1Transcript) GetAccessionVersionOk() (*string, bool)`

GetAccessionVersionOk returns a tuple with the AccessionVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessionVersion

`func (o *V1Transcript) SetAccessionVersion(v string)`

SetAccessionVersion sets AccessionVersion field to given value.

### HasAccessionVersion

`func (o *V1Transcript) HasAccessionVersion() bool`

HasAccessionVersion returns a boolean if a field has been set.

### GetName

`func (o *V1Transcript) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1Transcript) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1Transcript) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1Transcript) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLength

`func (o *V1Transcript) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *V1Transcript) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *V1Transcript) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *V1Transcript) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetGenomicRange

`func (o *V1Transcript) GetGenomicRange() V1SeqRangeSet`

GetGenomicRange returns the GenomicRange field if non-nil, zero value otherwise.

### GetGenomicRangeOk

`func (o *V1Transcript) GetGenomicRangeOk() (*V1SeqRangeSet, bool)`

GetGenomicRangeOk returns a tuple with the GenomicRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenomicRange

`func (o *V1Transcript) SetGenomicRange(v V1SeqRangeSet)`

SetGenomicRange sets GenomicRange field to given value.

### HasGenomicRange

`func (o *V1Transcript) HasGenomicRange() bool`

HasGenomicRange returns a boolean if a field has been set.

### GetExons

`func (o *V1Transcript) GetExons() V1SeqRangeSet`

GetExons returns the Exons field if non-nil, zero value otherwise.

### GetExonsOk

`func (o *V1Transcript) GetExonsOk() (*V1SeqRangeSet, bool)`

GetExonsOk returns a tuple with the Exons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExons

`func (o *V1Transcript) SetExons(v V1SeqRangeSet)`

SetExons sets Exons field to given value.

### HasExons

`func (o *V1Transcript) HasExons() bool`

HasExons returns a boolean if a field has been set.

### GetCds

`func (o *V1Transcript) GetCds() V1SeqRangeSet`

GetCds returns the Cds field if non-nil, zero value otherwise.

### GetCdsOk

`func (o *V1Transcript) GetCdsOk() (*V1SeqRangeSet, bool)`

GetCdsOk returns a tuple with the Cds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCds

`func (o *V1Transcript) SetCds(v V1SeqRangeSet)`

SetCds sets Cds field to given value.

### HasCds

`func (o *V1Transcript) HasCds() bool`

HasCds returns a boolean if a field has been set.

### GetGenomicLocations

`func (o *V1Transcript) GetGenomicLocations() []V1GenomicLocation`

GetGenomicLocations returns the GenomicLocations field if non-nil, zero value otherwise.

### GetGenomicLocationsOk

`func (o *V1Transcript) GetGenomicLocationsOk() (*[]V1GenomicLocation, bool)`

GetGenomicLocationsOk returns a tuple with the GenomicLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenomicLocations

`func (o *V1Transcript) SetGenomicLocations(v []V1GenomicLocation)`

SetGenomicLocations sets GenomicLocations field to given value.

### HasGenomicLocations

`func (o *V1Transcript) HasGenomicLocations() bool`

HasGenomicLocations returns a boolean if a field has been set.

### GetEnsemblTranscript

`func (o *V1Transcript) GetEnsemblTranscript() string`

GetEnsemblTranscript returns the EnsemblTranscript field if non-nil, zero value otherwise.

### GetEnsemblTranscriptOk

`func (o *V1Transcript) GetEnsemblTranscriptOk() (*string, bool)`

GetEnsemblTranscriptOk returns a tuple with the EnsemblTranscript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnsemblTranscript

`func (o *V1Transcript) SetEnsemblTranscript(v string)`

SetEnsemblTranscript sets EnsemblTranscript field to given value.

### HasEnsemblTranscript

`func (o *V1Transcript) HasEnsemblTranscript() bool`

HasEnsemblTranscript returns a boolean if a field has been set.

### GetProtein

`func (o *V1Transcript) GetProtein() V1Protein`

GetProtein returns the Protein field if non-nil, zero value otherwise.

### GetProteinOk

`func (o *V1Transcript) GetProteinOk() (*V1Protein, bool)`

GetProteinOk returns a tuple with the Protein field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtein

`func (o *V1Transcript) SetProtein(v V1Protein)`

SetProtein sets Protein field to given value.

### HasProtein

`func (o *V1Transcript) HasProtein() bool`

HasProtein returns a boolean if a field has been set.

### GetType

`func (o *V1Transcript) GetType() V1TranscriptTranscriptType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1Transcript) GetTypeOk() (*V1TranscriptTranscriptType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1Transcript) SetType(v V1TranscriptTranscriptType)`

SetType sets Type field to given value.

### HasType

`func (o *V1Transcript) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


