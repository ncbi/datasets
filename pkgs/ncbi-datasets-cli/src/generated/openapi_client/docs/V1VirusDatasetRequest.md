# V1VirusDatasetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaxId** | Pointer to **int32** |  | [optional] 
**TaxName** | Pointer to **string** |  | [optional] 
**Accession** | Pointer to **string** |  | [optional] 
**Accessions** | Pointer to **[]string** |  | [optional] 
**Taxon** | Pointer to **string** |  | [optional] 
**RefseqOnly** | Pointer to **bool** |  | [optional] 
**AnnotatedOnly** | Pointer to **bool** |  | [optional] 
**ReleasedSince** | Pointer to **time.Time** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**PangolinClassification** | Pointer to **string** |  | [optional] 
**GeoLocation** | Pointer to **string** |  | [optional] 
**CompleteOnly** | Pointer to **bool** |  | [optional] 
**TableFields** | Pointer to [**[]V1VirusTableField**](V1VirusTableField.md) |  | [optional] 
**ExcludeSequence** | Pointer to **bool** |  | [optional] 
**IncludeAnnotationType** | Pointer to [**[]V1AnnotationForVirusType**](V1AnnotationForVirusType.md) |  | [optional] 
**Format** | Pointer to [**V1TableFormat**](V1TableFormat.md) |  | [optional] [default to V1TABLEFORMAT_TSV]

## Methods

### NewV1VirusDatasetRequest

`func NewV1VirusDatasetRequest() *V1VirusDatasetRequest`

NewV1VirusDatasetRequest instantiates a new V1VirusDatasetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1VirusDatasetRequestWithDefaults

`func NewV1VirusDatasetRequestWithDefaults() *V1VirusDatasetRequest`

NewV1VirusDatasetRequestWithDefaults instantiates a new V1VirusDatasetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaxId

`func (o *V1VirusDatasetRequest) GetTaxId() int32`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *V1VirusDatasetRequest) GetTaxIdOk() (*int32, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *V1VirusDatasetRequest) SetTaxId(v int32)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *V1VirusDatasetRequest) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.

### GetTaxName

`func (o *V1VirusDatasetRequest) GetTaxName() string`

GetTaxName returns the TaxName field if non-nil, zero value otherwise.

### GetTaxNameOk

`func (o *V1VirusDatasetRequest) GetTaxNameOk() (*string, bool)`

GetTaxNameOk returns a tuple with the TaxName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxName

`func (o *V1VirusDatasetRequest) SetTaxName(v string)`

SetTaxName sets TaxName field to given value.

### HasTaxName

`func (o *V1VirusDatasetRequest) HasTaxName() bool`

HasTaxName returns a boolean if a field has been set.

### GetAccession

`func (o *V1VirusDatasetRequest) GetAccession() string`

GetAccession returns the Accession field if non-nil, zero value otherwise.

### GetAccessionOk

`func (o *V1VirusDatasetRequest) GetAccessionOk() (*string, bool)`

GetAccessionOk returns a tuple with the Accession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccession

`func (o *V1VirusDatasetRequest) SetAccession(v string)`

SetAccession sets Accession field to given value.

### HasAccession

`func (o *V1VirusDatasetRequest) HasAccession() bool`

HasAccession returns a boolean if a field has been set.

### GetAccessions

`func (o *V1VirusDatasetRequest) GetAccessions() []string`

GetAccessions returns the Accessions field if non-nil, zero value otherwise.

### GetAccessionsOk

`func (o *V1VirusDatasetRequest) GetAccessionsOk() (*[]string, bool)`

GetAccessionsOk returns a tuple with the Accessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessions

`func (o *V1VirusDatasetRequest) SetAccessions(v []string)`

SetAccessions sets Accessions field to given value.

### HasAccessions

`func (o *V1VirusDatasetRequest) HasAccessions() bool`

HasAccessions returns a boolean if a field has been set.

### GetTaxon

`func (o *V1VirusDatasetRequest) GetTaxon() string`

GetTaxon returns the Taxon field if non-nil, zero value otherwise.

### GetTaxonOk

`func (o *V1VirusDatasetRequest) GetTaxonOk() (*string, bool)`

GetTaxonOk returns a tuple with the Taxon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxon

`func (o *V1VirusDatasetRequest) SetTaxon(v string)`

SetTaxon sets Taxon field to given value.

### HasTaxon

`func (o *V1VirusDatasetRequest) HasTaxon() bool`

HasTaxon returns a boolean if a field has been set.

### GetRefseqOnly

`func (o *V1VirusDatasetRequest) GetRefseqOnly() bool`

GetRefseqOnly returns the RefseqOnly field if non-nil, zero value otherwise.

### GetRefseqOnlyOk

`func (o *V1VirusDatasetRequest) GetRefseqOnlyOk() (*bool, bool)`

GetRefseqOnlyOk returns a tuple with the RefseqOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefseqOnly

`func (o *V1VirusDatasetRequest) SetRefseqOnly(v bool)`

SetRefseqOnly sets RefseqOnly field to given value.

### HasRefseqOnly

`func (o *V1VirusDatasetRequest) HasRefseqOnly() bool`

HasRefseqOnly returns a boolean if a field has been set.

### GetAnnotatedOnly

`func (o *V1VirusDatasetRequest) GetAnnotatedOnly() bool`

GetAnnotatedOnly returns the AnnotatedOnly field if non-nil, zero value otherwise.

### GetAnnotatedOnlyOk

`func (o *V1VirusDatasetRequest) GetAnnotatedOnlyOk() (*bool, bool)`

GetAnnotatedOnlyOk returns a tuple with the AnnotatedOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotatedOnly

`func (o *V1VirusDatasetRequest) SetAnnotatedOnly(v bool)`

SetAnnotatedOnly sets AnnotatedOnly field to given value.

### HasAnnotatedOnly

`func (o *V1VirusDatasetRequest) HasAnnotatedOnly() bool`

HasAnnotatedOnly returns a boolean if a field has been set.

### GetReleasedSince

`func (o *V1VirusDatasetRequest) GetReleasedSince() time.Time`

GetReleasedSince returns the ReleasedSince field if non-nil, zero value otherwise.

### GetReleasedSinceOk

`func (o *V1VirusDatasetRequest) GetReleasedSinceOk() (*time.Time, bool)`

GetReleasedSinceOk returns a tuple with the ReleasedSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasedSince

`func (o *V1VirusDatasetRequest) SetReleasedSince(v time.Time)`

SetReleasedSince sets ReleasedSince field to given value.

### HasReleasedSince

`func (o *V1VirusDatasetRequest) HasReleasedSince() bool`

HasReleasedSince returns a boolean if a field has been set.

### GetHost

`func (o *V1VirusDatasetRequest) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *V1VirusDatasetRequest) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *V1VirusDatasetRequest) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *V1VirusDatasetRequest) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPangolinClassification

`func (o *V1VirusDatasetRequest) GetPangolinClassification() string`

GetPangolinClassification returns the PangolinClassification field if non-nil, zero value otherwise.

### GetPangolinClassificationOk

`func (o *V1VirusDatasetRequest) GetPangolinClassificationOk() (*string, bool)`

GetPangolinClassificationOk returns a tuple with the PangolinClassification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPangolinClassification

`func (o *V1VirusDatasetRequest) SetPangolinClassification(v string)`

SetPangolinClassification sets PangolinClassification field to given value.

### HasPangolinClassification

`func (o *V1VirusDatasetRequest) HasPangolinClassification() bool`

HasPangolinClassification returns a boolean if a field has been set.

### GetGeoLocation

`func (o *V1VirusDatasetRequest) GetGeoLocation() string`

GetGeoLocation returns the GeoLocation field if non-nil, zero value otherwise.

### GetGeoLocationOk

`func (o *V1VirusDatasetRequest) GetGeoLocationOk() (*string, bool)`

GetGeoLocationOk returns a tuple with the GeoLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoLocation

`func (o *V1VirusDatasetRequest) SetGeoLocation(v string)`

SetGeoLocation sets GeoLocation field to given value.

### HasGeoLocation

`func (o *V1VirusDatasetRequest) HasGeoLocation() bool`

HasGeoLocation returns a boolean if a field has been set.

### GetCompleteOnly

`func (o *V1VirusDatasetRequest) GetCompleteOnly() bool`

GetCompleteOnly returns the CompleteOnly field if non-nil, zero value otherwise.

### GetCompleteOnlyOk

`func (o *V1VirusDatasetRequest) GetCompleteOnlyOk() (*bool, bool)`

GetCompleteOnlyOk returns a tuple with the CompleteOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleteOnly

`func (o *V1VirusDatasetRequest) SetCompleteOnly(v bool)`

SetCompleteOnly sets CompleteOnly field to given value.

### HasCompleteOnly

`func (o *V1VirusDatasetRequest) HasCompleteOnly() bool`

HasCompleteOnly returns a boolean if a field has been set.

### GetTableFields

`func (o *V1VirusDatasetRequest) GetTableFields() []V1VirusTableField`

GetTableFields returns the TableFields field if non-nil, zero value otherwise.

### GetTableFieldsOk

`func (o *V1VirusDatasetRequest) GetTableFieldsOk() (*[]V1VirusTableField, bool)`

GetTableFieldsOk returns a tuple with the TableFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableFields

`func (o *V1VirusDatasetRequest) SetTableFields(v []V1VirusTableField)`

SetTableFields sets TableFields field to given value.

### HasTableFields

`func (o *V1VirusDatasetRequest) HasTableFields() bool`

HasTableFields returns a boolean if a field has been set.

### GetExcludeSequence

`func (o *V1VirusDatasetRequest) GetExcludeSequence() bool`

GetExcludeSequence returns the ExcludeSequence field if non-nil, zero value otherwise.

### GetExcludeSequenceOk

`func (o *V1VirusDatasetRequest) GetExcludeSequenceOk() (*bool, bool)`

GetExcludeSequenceOk returns a tuple with the ExcludeSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeSequence

`func (o *V1VirusDatasetRequest) SetExcludeSequence(v bool)`

SetExcludeSequence sets ExcludeSequence field to given value.

### HasExcludeSequence

`func (o *V1VirusDatasetRequest) HasExcludeSequence() bool`

HasExcludeSequence returns a boolean if a field has been set.

### GetIncludeAnnotationType

`func (o *V1VirusDatasetRequest) GetIncludeAnnotationType() []V1AnnotationForVirusType`

GetIncludeAnnotationType returns the IncludeAnnotationType field if non-nil, zero value otherwise.

### GetIncludeAnnotationTypeOk

`func (o *V1VirusDatasetRequest) GetIncludeAnnotationTypeOk() (*[]V1AnnotationForVirusType, bool)`

GetIncludeAnnotationTypeOk returns a tuple with the IncludeAnnotationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAnnotationType

`func (o *V1VirusDatasetRequest) SetIncludeAnnotationType(v []V1AnnotationForVirusType)`

SetIncludeAnnotationType sets IncludeAnnotationType field to given value.

### HasIncludeAnnotationType

`func (o *V1VirusDatasetRequest) HasIncludeAnnotationType() bool`

HasIncludeAnnotationType returns a boolean if a field has been set.

### GetFormat

`func (o *V1VirusDatasetRequest) GetFormat() V1TableFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *V1VirusDatasetRequest) GetFormatOk() (*V1TableFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *V1VirusDatasetRequest) SetFormat(v V1TableFormat)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *V1VirusDatasetRequest) HasFormat() bool`

HasFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


