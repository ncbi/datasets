# V1VirusDatasetFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accessions** | Pointer to **[]string** |  | [optional] 
**Taxon** | Pointer to **string** |  | [optional] 
**RefseqOnly** | Pointer to **bool** |  | [optional] 
**AnnotatedOnly** | Pointer to **bool** |  | [optional] 
**ReleasedSince** | Pointer to **time.Time** |  | [optional] 
**UpdatedSince** | Pointer to **time.Time** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**PangolinClassification** | Pointer to **string** |  | [optional] 
**GeoLocation** | Pointer to **string** |  | [optional] 
**CompleteOnly** | Pointer to **bool** |  | [optional] 

## Methods

### NewV1VirusDatasetFilter

`func NewV1VirusDatasetFilter() *V1VirusDatasetFilter`

NewV1VirusDatasetFilter instantiates a new V1VirusDatasetFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1VirusDatasetFilterWithDefaults

`func NewV1VirusDatasetFilterWithDefaults() *V1VirusDatasetFilter`

NewV1VirusDatasetFilterWithDefaults instantiates a new V1VirusDatasetFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessions

`func (o *V1VirusDatasetFilter) GetAccessions() []string`

GetAccessions returns the Accessions field if non-nil, zero value otherwise.

### GetAccessionsOk

`func (o *V1VirusDatasetFilter) GetAccessionsOk() (*[]string, bool)`

GetAccessionsOk returns a tuple with the Accessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessions

`func (o *V1VirusDatasetFilter) SetAccessions(v []string)`

SetAccessions sets Accessions field to given value.

### HasAccessions

`func (o *V1VirusDatasetFilter) HasAccessions() bool`

HasAccessions returns a boolean if a field has been set.

### GetTaxon

`func (o *V1VirusDatasetFilter) GetTaxon() string`

GetTaxon returns the Taxon field if non-nil, zero value otherwise.

### GetTaxonOk

`func (o *V1VirusDatasetFilter) GetTaxonOk() (*string, bool)`

GetTaxonOk returns a tuple with the Taxon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxon

`func (o *V1VirusDatasetFilter) SetTaxon(v string)`

SetTaxon sets Taxon field to given value.

### HasTaxon

`func (o *V1VirusDatasetFilter) HasTaxon() bool`

HasTaxon returns a boolean if a field has been set.

### GetRefseqOnly

`func (o *V1VirusDatasetFilter) GetRefseqOnly() bool`

GetRefseqOnly returns the RefseqOnly field if non-nil, zero value otherwise.

### GetRefseqOnlyOk

`func (o *V1VirusDatasetFilter) GetRefseqOnlyOk() (*bool, bool)`

GetRefseqOnlyOk returns a tuple with the RefseqOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefseqOnly

`func (o *V1VirusDatasetFilter) SetRefseqOnly(v bool)`

SetRefseqOnly sets RefseqOnly field to given value.

### HasRefseqOnly

`func (o *V1VirusDatasetFilter) HasRefseqOnly() bool`

HasRefseqOnly returns a boolean if a field has been set.

### GetAnnotatedOnly

`func (o *V1VirusDatasetFilter) GetAnnotatedOnly() bool`

GetAnnotatedOnly returns the AnnotatedOnly field if non-nil, zero value otherwise.

### GetAnnotatedOnlyOk

`func (o *V1VirusDatasetFilter) GetAnnotatedOnlyOk() (*bool, bool)`

GetAnnotatedOnlyOk returns a tuple with the AnnotatedOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotatedOnly

`func (o *V1VirusDatasetFilter) SetAnnotatedOnly(v bool)`

SetAnnotatedOnly sets AnnotatedOnly field to given value.

### HasAnnotatedOnly

`func (o *V1VirusDatasetFilter) HasAnnotatedOnly() bool`

HasAnnotatedOnly returns a boolean if a field has been set.

### GetReleasedSince

`func (o *V1VirusDatasetFilter) GetReleasedSince() time.Time`

GetReleasedSince returns the ReleasedSince field if non-nil, zero value otherwise.

### GetReleasedSinceOk

`func (o *V1VirusDatasetFilter) GetReleasedSinceOk() (*time.Time, bool)`

GetReleasedSinceOk returns a tuple with the ReleasedSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleasedSince

`func (o *V1VirusDatasetFilter) SetReleasedSince(v time.Time)`

SetReleasedSince sets ReleasedSince field to given value.

### HasReleasedSince

`func (o *V1VirusDatasetFilter) HasReleasedSince() bool`

HasReleasedSince returns a boolean if a field has been set.

### GetUpdatedSince

`func (o *V1VirusDatasetFilter) GetUpdatedSince() time.Time`

GetUpdatedSince returns the UpdatedSince field if non-nil, zero value otherwise.

### GetUpdatedSinceOk

`func (o *V1VirusDatasetFilter) GetUpdatedSinceOk() (*time.Time, bool)`

GetUpdatedSinceOk returns a tuple with the UpdatedSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedSince

`func (o *V1VirusDatasetFilter) SetUpdatedSince(v time.Time)`

SetUpdatedSince sets UpdatedSince field to given value.

### HasUpdatedSince

`func (o *V1VirusDatasetFilter) HasUpdatedSince() bool`

HasUpdatedSince returns a boolean if a field has been set.

### GetHost

`func (o *V1VirusDatasetFilter) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *V1VirusDatasetFilter) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *V1VirusDatasetFilter) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *V1VirusDatasetFilter) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPangolinClassification

`func (o *V1VirusDatasetFilter) GetPangolinClassification() string`

GetPangolinClassification returns the PangolinClassification field if non-nil, zero value otherwise.

### GetPangolinClassificationOk

`func (o *V1VirusDatasetFilter) GetPangolinClassificationOk() (*string, bool)`

GetPangolinClassificationOk returns a tuple with the PangolinClassification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPangolinClassification

`func (o *V1VirusDatasetFilter) SetPangolinClassification(v string)`

SetPangolinClassification sets PangolinClassification field to given value.

### HasPangolinClassification

`func (o *V1VirusDatasetFilter) HasPangolinClassification() bool`

HasPangolinClassification returns a boolean if a field has been set.

### GetGeoLocation

`func (o *V1VirusDatasetFilter) GetGeoLocation() string`

GetGeoLocation returns the GeoLocation field if non-nil, zero value otherwise.

### GetGeoLocationOk

`func (o *V1VirusDatasetFilter) GetGeoLocationOk() (*string, bool)`

GetGeoLocationOk returns a tuple with the GeoLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoLocation

`func (o *V1VirusDatasetFilter) SetGeoLocation(v string)`

SetGeoLocation sets GeoLocation field to given value.

### HasGeoLocation

`func (o *V1VirusDatasetFilter) HasGeoLocation() bool`

HasGeoLocation returns a boolean if a field has been set.

### GetCompleteOnly

`func (o *V1VirusDatasetFilter) GetCompleteOnly() bool`

GetCompleteOnly returns the CompleteOnly field if non-nil, zero value otherwise.

### GetCompleteOnlyOk

`func (o *V1VirusDatasetFilter) GetCompleteOnlyOk() (*bool, bool)`

GetCompleteOnlyOk returns a tuple with the CompleteOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleteOnly

`func (o *V1VirusDatasetFilter) SetCompleteOnly(v bool)`

SetCompleteOnly sets CompleteOnly field to given value.

### HasCompleteOnly

`func (o *V1VirusDatasetFilter) HasCompleteOnly() bool`

HasCompleteOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


