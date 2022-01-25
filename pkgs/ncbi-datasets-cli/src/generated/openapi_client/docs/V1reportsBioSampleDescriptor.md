# V1reportsBioSampleDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accession** | Pointer to **string** |  | [optional] 
**LastUpdated** | Pointer to **string** |  | [optional] 
**PublicationDate** | Pointer to **string** |  | [optional] 
**SubmissionDate** | Pointer to **string** |  | [optional] 
**SampleIds** | Pointer to [**[]V1reportsBioSampleId**](V1reportsBioSampleId.md) |  | [optional] 
**Description** | Pointer to [**V1reportsBioSampleDescription**](V1reportsBioSampleDescription.md) |  | [optional] 
**Owner** | Pointer to [**V1reportsBioSampleOwner**](V1reportsBioSampleOwner.md) |  | [optional] 
**Models** | Pointer to **[]string** |  | [optional] 
**Bioprojects** | Pointer to [**[]V1reportsBioProject**](V1reportsBioProject.md) |  | [optional] 
**Package** | Pointer to **string** |  | [optional] 
**Attributes** | Pointer to [**[]V1reportsBioSampleAttribute**](V1reportsBioSampleAttribute.md) |  | [optional] 
**Status** | Pointer to [**V1reportsBioSampleStatus**](V1reportsBioSampleStatus.md) |  | [optional] 

## Methods

### NewV1reportsBioSampleDescriptor

`func NewV1reportsBioSampleDescriptor() *V1reportsBioSampleDescriptor`

NewV1reportsBioSampleDescriptor instantiates a new V1reportsBioSampleDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1reportsBioSampleDescriptorWithDefaults

`func NewV1reportsBioSampleDescriptorWithDefaults() *V1reportsBioSampleDescriptor`

NewV1reportsBioSampleDescriptorWithDefaults instantiates a new V1reportsBioSampleDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccession

`func (o *V1reportsBioSampleDescriptor) GetAccession() string`

GetAccession returns the Accession field if non-nil, zero value otherwise.

### GetAccessionOk

`func (o *V1reportsBioSampleDescriptor) GetAccessionOk() (*string, bool)`

GetAccessionOk returns a tuple with the Accession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccession

`func (o *V1reportsBioSampleDescriptor) SetAccession(v string)`

SetAccession sets Accession field to given value.

### HasAccession

`func (o *V1reportsBioSampleDescriptor) HasAccession() bool`

HasAccession returns a boolean if a field has been set.

### GetLastUpdated

`func (o *V1reportsBioSampleDescriptor) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *V1reportsBioSampleDescriptor) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *V1reportsBioSampleDescriptor) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *V1reportsBioSampleDescriptor) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetPublicationDate

`func (o *V1reportsBioSampleDescriptor) GetPublicationDate() string`

GetPublicationDate returns the PublicationDate field if non-nil, zero value otherwise.

### GetPublicationDateOk

`func (o *V1reportsBioSampleDescriptor) GetPublicationDateOk() (*string, bool)`

GetPublicationDateOk returns a tuple with the PublicationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicationDate

`func (o *V1reportsBioSampleDescriptor) SetPublicationDate(v string)`

SetPublicationDate sets PublicationDate field to given value.

### HasPublicationDate

`func (o *V1reportsBioSampleDescriptor) HasPublicationDate() bool`

HasPublicationDate returns a boolean if a field has been set.

### GetSubmissionDate

`func (o *V1reportsBioSampleDescriptor) GetSubmissionDate() string`

GetSubmissionDate returns the SubmissionDate field if non-nil, zero value otherwise.

### GetSubmissionDateOk

`func (o *V1reportsBioSampleDescriptor) GetSubmissionDateOk() (*string, bool)`

GetSubmissionDateOk returns a tuple with the SubmissionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionDate

`func (o *V1reportsBioSampleDescriptor) SetSubmissionDate(v string)`

SetSubmissionDate sets SubmissionDate field to given value.

### HasSubmissionDate

`func (o *V1reportsBioSampleDescriptor) HasSubmissionDate() bool`

HasSubmissionDate returns a boolean if a field has been set.

### GetSampleIds

`func (o *V1reportsBioSampleDescriptor) GetSampleIds() []V1reportsBioSampleId`

GetSampleIds returns the SampleIds field if non-nil, zero value otherwise.

### GetSampleIdsOk

`func (o *V1reportsBioSampleDescriptor) GetSampleIdsOk() (*[]V1reportsBioSampleId, bool)`

GetSampleIdsOk returns a tuple with the SampleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleIds

`func (o *V1reportsBioSampleDescriptor) SetSampleIds(v []V1reportsBioSampleId)`

SetSampleIds sets SampleIds field to given value.

### HasSampleIds

`func (o *V1reportsBioSampleDescriptor) HasSampleIds() bool`

HasSampleIds returns a boolean if a field has been set.

### GetDescription

`func (o *V1reportsBioSampleDescriptor) GetDescription() V1reportsBioSampleDescription`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V1reportsBioSampleDescriptor) GetDescriptionOk() (*V1reportsBioSampleDescription, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V1reportsBioSampleDescriptor) SetDescription(v V1reportsBioSampleDescription)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V1reportsBioSampleDescriptor) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *V1reportsBioSampleDescriptor) GetOwner() V1reportsBioSampleOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *V1reportsBioSampleDescriptor) GetOwnerOk() (*V1reportsBioSampleOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *V1reportsBioSampleDescriptor) SetOwner(v V1reportsBioSampleOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *V1reportsBioSampleDescriptor) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetModels

`func (o *V1reportsBioSampleDescriptor) GetModels() []string`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *V1reportsBioSampleDescriptor) GetModelsOk() (*[]string, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *V1reportsBioSampleDescriptor) SetModels(v []string)`

SetModels sets Models field to given value.

### HasModels

`func (o *V1reportsBioSampleDescriptor) HasModels() bool`

HasModels returns a boolean if a field has been set.

### GetBioprojects

`func (o *V1reportsBioSampleDescriptor) GetBioprojects() []V1reportsBioProject`

GetBioprojects returns the Bioprojects field if non-nil, zero value otherwise.

### GetBioprojectsOk

`func (o *V1reportsBioSampleDescriptor) GetBioprojectsOk() (*[]V1reportsBioProject, bool)`

GetBioprojectsOk returns a tuple with the Bioprojects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBioprojects

`func (o *V1reportsBioSampleDescriptor) SetBioprojects(v []V1reportsBioProject)`

SetBioprojects sets Bioprojects field to given value.

### HasBioprojects

`func (o *V1reportsBioSampleDescriptor) HasBioprojects() bool`

HasBioprojects returns a boolean if a field has been set.

### GetPackage

`func (o *V1reportsBioSampleDescriptor) GetPackage() string`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *V1reportsBioSampleDescriptor) GetPackageOk() (*string, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *V1reportsBioSampleDescriptor) SetPackage(v string)`

SetPackage sets Package field to given value.

### HasPackage

`func (o *V1reportsBioSampleDescriptor) HasPackage() bool`

HasPackage returns a boolean if a field has been set.

### GetAttributes

`func (o *V1reportsBioSampleDescriptor) GetAttributes() []V1reportsBioSampleAttribute`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *V1reportsBioSampleDescriptor) GetAttributesOk() (*[]V1reportsBioSampleAttribute, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *V1reportsBioSampleDescriptor) SetAttributes(v []V1reportsBioSampleAttribute)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *V1reportsBioSampleDescriptor) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetStatus

`func (o *V1reportsBioSampleDescriptor) GetStatus() V1reportsBioSampleStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *V1reportsBioSampleDescriptor) GetStatusOk() (*V1reportsBioSampleStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *V1reportsBioSampleDescriptor) SetStatus(v V1reportsBioSampleStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *V1reportsBioSampleDescriptor) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


