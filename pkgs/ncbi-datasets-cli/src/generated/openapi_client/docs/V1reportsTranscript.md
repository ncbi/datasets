# V1reportsTranscript

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessionVersion** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Length** | Pointer to **int32** |  | [optional] 
**GenomicRange** | Pointer to [**V1reportsSeqRangeSet**](V1reportsSeqRangeSet.md) |  | [optional] 
**Exons** | Pointer to [**V1reportsSeqRangeSet**](V1reportsSeqRangeSet.md) |  | [optional] 
**Cds** | Pointer to [**V1reportsSeqRangeSet**](V1reportsSeqRangeSet.md) |  | [optional] 
**GenomicLocations** | Pointer to [**[]V1reportsGenomicLocation**](V1reportsGenomicLocation.md) |  | [optional] 
**EnsemblTranscript** | Pointer to **string** |  | [optional] 
**Protein** | Pointer to [**V1reportsProtein**](V1reportsProtein.md) |  | [optional] 
**Type** | Pointer to [**V1reportsTranscriptTranscriptType**](V1reportsTranscriptTranscriptType.md) |  | [optional] [default to V1REPORTSTRANSCRIPTTRANSCRIPTTYPE_UNKNOWN]

## Methods

### NewV1reportsTranscript

`func NewV1reportsTranscript() *V1reportsTranscript`

NewV1reportsTranscript instantiates a new V1reportsTranscript object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1reportsTranscriptWithDefaults

`func NewV1reportsTranscriptWithDefaults() *V1reportsTranscript`

NewV1reportsTranscriptWithDefaults instantiates a new V1reportsTranscript object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessionVersion

`func (o *V1reportsTranscript) GetAccessionVersion() string`

GetAccessionVersion returns the AccessionVersion field if non-nil, zero value otherwise.

### GetAccessionVersionOk

`func (o *V1reportsTranscript) GetAccessionVersionOk() (*string, bool)`

GetAccessionVersionOk returns a tuple with the AccessionVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessionVersion

`func (o *V1reportsTranscript) SetAccessionVersion(v string)`

SetAccessionVersion sets AccessionVersion field to given value.

### HasAccessionVersion

`func (o *V1reportsTranscript) HasAccessionVersion() bool`

HasAccessionVersion returns a boolean if a field has been set.

### GetName

`func (o *V1reportsTranscript) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1reportsTranscript) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1reportsTranscript) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1reportsTranscript) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLength

`func (o *V1reportsTranscript) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *V1reportsTranscript) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *V1reportsTranscript) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *V1reportsTranscript) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetGenomicRange

`func (o *V1reportsTranscript) GetGenomicRange() V1reportsSeqRangeSet`

GetGenomicRange returns the GenomicRange field if non-nil, zero value otherwise.

### GetGenomicRangeOk

`func (o *V1reportsTranscript) GetGenomicRangeOk() (*V1reportsSeqRangeSet, bool)`

GetGenomicRangeOk returns a tuple with the GenomicRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenomicRange

`func (o *V1reportsTranscript) SetGenomicRange(v V1reportsSeqRangeSet)`

SetGenomicRange sets GenomicRange field to given value.

### HasGenomicRange

`func (o *V1reportsTranscript) HasGenomicRange() bool`

HasGenomicRange returns a boolean if a field has been set.

### GetExons

`func (o *V1reportsTranscript) GetExons() V1reportsSeqRangeSet`

GetExons returns the Exons field if non-nil, zero value otherwise.

### GetExonsOk

`func (o *V1reportsTranscript) GetExonsOk() (*V1reportsSeqRangeSet, bool)`

GetExonsOk returns a tuple with the Exons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExons

`func (o *V1reportsTranscript) SetExons(v V1reportsSeqRangeSet)`

SetExons sets Exons field to given value.

### HasExons

`func (o *V1reportsTranscript) HasExons() bool`

HasExons returns a boolean if a field has been set.

### GetCds

`func (o *V1reportsTranscript) GetCds() V1reportsSeqRangeSet`

GetCds returns the Cds field if non-nil, zero value otherwise.

### GetCdsOk

`func (o *V1reportsTranscript) GetCdsOk() (*V1reportsSeqRangeSet, bool)`

GetCdsOk returns a tuple with the Cds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCds

`func (o *V1reportsTranscript) SetCds(v V1reportsSeqRangeSet)`

SetCds sets Cds field to given value.

### HasCds

`func (o *V1reportsTranscript) HasCds() bool`

HasCds returns a boolean if a field has been set.

### GetGenomicLocations

`func (o *V1reportsTranscript) GetGenomicLocations() []V1reportsGenomicLocation`

GetGenomicLocations returns the GenomicLocations field if non-nil, zero value otherwise.

### GetGenomicLocationsOk

`func (o *V1reportsTranscript) GetGenomicLocationsOk() (*[]V1reportsGenomicLocation, bool)`

GetGenomicLocationsOk returns a tuple with the GenomicLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenomicLocations

`func (o *V1reportsTranscript) SetGenomicLocations(v []V1reportsGenomicLocation)`

SetGenomicLocations sets GenomicLocations field to given value.

### HasGenomicLocations

`func (o *V1reportsTranscript) HasGenomicLocations() bool`

HasGenomicLocations returns a boolean if a field has been set.

### GetEnsemblTranscript

`func (o *V1reportsTranscript) GetEnsemblTranscript() string`

GetEnsemblTranscript returns the EnsemblTranscript field if non-nil, zero value otherwise.

### GetEnsemblTranscriptOk

`func (o *V1reportsTranscript) GetEnsemblTranscriptOk() (*string, bool)`

GetEnsemblTranscriptOk returns a tuple with the EnsemblTranscript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnsemblTranscript

`func (o *V1reportsTranscript) SetEnsemblTranscript(v string)`

SetEnsemblTranscript sets EnsemblTranscript field to given value.

### HasEnsemblTranscript

`func (o *V1reportsTranscript) HasEnsemblTranscript() bool`

HasEnsemblTranscript returns a boolean if a field has been set.

### GetProtein

`func (o *V1reportsTranscript) GetProtein() V1reportsProtein`

GetProtein returns the Protein field if non-nil, zero value otherwise.

### GetProteinOk

`func (o *V1reportsTranscript) GetProteinOk() (*V1reportsProtein, bool)`

GetProteinOk returns a tuple with the Protein field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtein

`func (o *V1reportsTranscript) SetProtein(v V1reportsProtein)`

SetProtein sets Protein field to given value.

### HasProtein

`func (o *V1reportsTranscript) HasProtein() bool`

HasProtein returns a boolean if a field has been set.

### GetType

`func (o *V1reportsTranscript) GetType() V1reportsTranscriptTranscriptType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1reportsTranscript) GetTypeOk() (*V1reportsTranscriptTranscriptType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1reportsTranscript) SetType(v V1reportsTranscriptTranscriptType)`

SetType sets Type field to given value.

### HasType

`func (o *V1reportsTranscript) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


