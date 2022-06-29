# V1Sars2ProteinDatasetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Proteins** | Pointer to **[]string** |  | [optional] 
**RefseqOnly** | Pointer to **bool** |  | [optional] 
**AnnotatedOnly** | Pointer to **bool** |  | [optional] 
**ReleasedSince** | Pointer to **time.Time** |  | [optional] 
**UpdatedSince** | Pointer to **time.Time** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**GeoLocation** | Pointer to **string** |  | [optional] 
**CompleteOnly** | Pointer to **bool** |  | [optional] 
**TableFields** | Pointer to [**[]V1VirusTableField**](V1VirusTableField.md) |  | [optional] 
**ExcludeSequence** | Pointer to **bool** |  | [optional] 
**IncludeAnnotationType** | Pointer to [**[]V1AnnotationForVirusType**](V1AnnotationForVirusType.md) |  | [optional] 
**Format** | Pointer to [**V1TableFormat**](V1TableFormat.md) |  | [optional] [default to V1TABLEFORMAT_TSV]

## Methods

### NewV1Sars2ProteinDatasetRequest

`func NewV1Sars2ProteinDatasetRequest() *V1Sars2ProteinDatasetRequest`

NewV1Sars2ProteinDatasetRequest instantiates a new V1Sars2ProteinDatasetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1Sars2ProteinDatasetRequestWithDefaults

`func NewV1Sars2ProteinDatasetRequestWithDefaults() *V1Sars2ProteinDatasetRequest`

NewV1Sars2ProteinDatasetRequestWithDefaults instantiates a new V1Sars2ProteinDatasetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProteins

`func (o *V1Sars2ProteinDatasetRequest) GetProteins() []string`

GetProteins returns the Proteins field if non-nil, zero value otherwise.

### GetProteinsOk

`func (o *V1Sars2ProteinDatasetRequest) GetProteinsOk() (*[]string, bool)`

GetProteinsOk returns a tuple with the Proteins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProteins

`func (o *V1Sars2ProteinDatasetRequest) SetProteins(v []string)`

SetProteins sets Proteins field to given value.

### HasProteins

`func (o *V1Sars2ProteinDatasetRequest) HasProteins() bool`

HasProteins returns a boolean if a field has been set.

### GetRefseqOnly

`func (o *V1Sars2ProteinDatasetRequest) GetRefseqOnly() bool`

GetRefseqOnly returns the RefseqOnly field if non-nil, zero value otherwise.

### GetRefseqOnlyOk

`func (o *V1Sars2ProteinDatasetRequest) GetRefseqOnlyOk() (*bool, bool)`

GetRefseqOnlyOk returns a tuple with the RefseqOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefseqOnly

`func (o *V1Sars2ProteinDatasetRequest) SetRefseqOnly(v bool)`

SetRefseqOnly sets RefseqOnly field to given value.

### HasRefseqOnly

`func (o *V1Sars2ProteinDatasetRequest) HasRefseqOnly() bool`

HasRefseqOnly returns a boolean if a field has been set.

### GetAnnotatedOnly

`func (o *V1Sars2ProteinDatasetRequest) GetAnnotatedOnly() bool`

GetAnnotatedOnly returns the AnnotatedOnly field if non-nil, zero value otherwise.

### GetAnnotatedOnlyOk

`func (o *V1Sars2ProteinDatasetRequest) GetAnnotatedOnlyOk() (*bool, bool)`

GetAnnotatedOnlyOk returns a tuple with the AnnotatedOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotatedOnly

`func (o *V1Sars2ProteinDatasetRequest) SetAnnotatedOnly(v bool)`

SetAnnotatedOnly sets AnnotatedOnly field to given value.

### HasAnnotatedOnly

`func (o *V1Sars2ProteinDatasetRequest) HasAnnotatedOnly() bool`

HasAnnotatedOnly returns a boolean if a field has been set.

### GetReleasedSince

`func (o *V1Sars2ProteinDatasetRequest) GetReleasedSince() time.Time`

GetReleasedSince returns the ReleasedSince field if non-nil, zero value otherwise.

### GetReleasedSinceOk

`func (o *V1Sars2ProteinDatasetRequest) GetReleasedSinceOk() (*time.Time, bool)`

GetReleasedSinceOk returns a tuple with the ReleasedSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasedSince

`func (o *V1Sars2ProteinDatasetRequest) SetReleasedSince(v time.Time)`

SetReleasedSince sets ReleasedSince field to given value.

### HasReleasedSince

`func (o *V1Sars2ProteinDatasetRequest) HasReleasedSince() bool`

HasReleasedSince returns a boolean if a field has been set.

### GetUpdatedSince

`func (o *V1Sars2ProteinDatasetRequest) GetUpdatedSince() time.Time`

GetUpdatedSince returns the UpdatedSince field if non-nil, zero value otherwise.

### GetUpdatedSinceOk

`func (o *V1Sars2ProteinDatasetRequest) GetUpdatedSinceOk() (*time.Time, bool)`

GetUpdatedSinceOk returns a tuple with the UpdatedSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedSince

`func (o *V1Sars2ProteinDatasetRequest) SetUpdatedSince(v time.Time)`

SetUpdatedSince sets UpdatedSince field to given value.

### HasUpdatedSince

`func (o *V1Sars2ProteinDatasetRequest) HasUpdatedSince() bool`

HasUpdatedSince returns a boolean if a field has been set.

### GetHost

`func (o *V1Sars2ProteinDatasetRequest) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *V1Sars2ProteinDatasetRequest) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *V1Sars2ProteinDatasetRequest) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *V1Sars2ProteinDatasetRequest) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetGeoLocation

`func (o *V1Sars2ProteinDatasetRequest) GetGeoLocation() string`

GetGeoLocation returns the GeoLocation field if non-nil, zero value otherwise.

### GetGeoLocationOk

`func (o *V1Sars2ProteinDatasetRequest) GetGeoLocationOk() (*string, bool)`

GetGeoLocationOk returns a tuple with the GeoLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoLocation

`func (o *V1Sars2ProteinDatasetRequest) SetGeoLocation(v string)`

SetGeoLocation sets GeoLocation field to given value.

### HasGeoLocation

`func (o *V1Sars2ProteinDatasetRequest) HasGeoLocation() bool`

HasGeoLocation returns a boolean if a field has been set.

### GetCompleteOnly

`func (o *V1Sars2ProteinDatasetRequest) GetCompleteOnly() bool`

GetCompleteOnly returns the CompleteOnly field if non-nil, zero value otherwise.

### GetCompleteOnlyOk

`func (o *V1Sars2ProteinDatasetRequest) GetCompleteOnlyOk() (*bool, bool)`

GetCompleteOnlyOk returns a tuple with the CompleteOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleteOnly

`func (o *V1Sars2ProteinDatasetRequest) SetCompleteOnly(v bool)`

SetCompleteOnly sets CompleteOnly field to given value.

### HasCompleteOnly

`func (o *V1Sars2ProteinDatasetRequest) HasCompleteOnly() bool`

HasCompleteOnly returns a boolean if a field has been set.

### GetTableFields

`func (o *V1Sars2ProteinDatasetRequest) GetTableFields() []V1VirusTableField`

GetTableFields returns the TableFields field if non-nil, zero value otherwise.

### GetTableFieldsOk

`func (o *V1Sars2ProteinDatasetRequest) GetTableFieldsOk() (*[]V1VirusTableField, bool)`

GetTableFieldsOk returns a tuple with the TableFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableFields

`func (o *V1Sars2ProteinDatasetRequest) SetTableFields(v []V1VirusTableField)`

SetTableFields sets TableFields field to given value.

### HasTableFields

`func (o *V1Sars2ProteinDatasetRequest) HasTableFields() bool`

HasTableFields returns a boolean if a field has been set.

### GetExcludeSequence

`func (o *V1Sars2ProteinDatasetRequest) GetExcludeSequence() bool`

GetExcludeSequence returns the ExcludeSequence field if non-nil, zero value otherwise.

### GetExcludeSequenceOk

`func (o *V1Sars2ProteinDatasetRequest) GetExcludeSequenceOk() (*bool, bool)`

GetExcludeSequenceOk returns a tuple with the ExcludeSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeSequence

`func (o *V1Sars2ProteinDatasetRequest) SetExcludeSequence(v bool)`

SetExcludeSequence sets ExcludeSequence field to given value.

### HasExcludeSequence

`func (o *V1Sars2ProteinDatasetRequest) HasExcludeSequence() bool`

HasExcludeSequence returns a boolean if a field has been set.

### GetIncludeAnnotationType

`func (o *V1Sars2ProteinDatasetRequest) GetIncludeAnnotationType() []V1AnnotationForVirusType`

GetIncludeAnnotationType returns the IncludeAnnotationType field if non-nil, zero value otherwise.

### GetIncludeAnnotationTypeOk

`func (o *V1Sars2ProteinDatasetRequest) GetIncludeAnnotationTypeOk() (*[]V1AnnotationForVirusType, bool)`

GetIncludeAnnotationTypeOk returns a tuple with the IncludeAnnotationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAnnotationType

`func (o *V1Sars2ProteinDatasetRequest) SetIncludeAnnotationType(v []V1AnnotationForVirusType)`

SetIncludeAnnotationType sets IncludeAnnotationType field to given value.

### HasIncludeAnnotationType

`func (o *V1Sars2ProteinDatasetRequest) HasIncludeAnnotationType() bool`

HasIncludeAnnotationType returns a boolean if a field has been set.

### GetFormat

`func (o *V1Sars2ProteinDatasetRequest) GetFormat() V1TableFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *V1Sars2ProteinDatasetRequest) GetFormatOk() (*V1TableFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *V1Sars2ProteinDatasetRequest) SetFormat(v V1TableFormat)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *V1Sars2ProteinDatasetRequest) HasFormat() bool`

HasFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


