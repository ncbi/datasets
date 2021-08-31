# V1reportsBioSampleDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Organism** | Pointer to [**V1reportsOrganism**](V1reportsOrganism.md) |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 

## Methods

### NewV1reportsBioSampleDescription

`func NewV1reportsBioSampleDescription() *V1reportsBioSampleDescription`

NewV1reportsBioSampleDescription instantiates a new V1reportsBioSampleDescription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1reportsBioSampleDescriptionWithDefaults

`func NewV1reportsBioSampleDescriptionWithDefaults() *V1reportsBioSampleDescription`

NewV1reportsBioSampleDescriptionWithDefaults instantiates a new V1reportsBioSampleDescription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *V1reportsBioSampleDescription) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *V1reportsBioSampleDescription) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *V1reportsBioSampleDescription) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *V1reportsBioSampleDescription) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetOrganism

`func (o *V1reportsBioSampleDescription) GetOrganism() V1reportsOrganism`

GetOrganism returns the Organism field if non-nil, zero value otherwise.

### GetOrganismOk

`func (o *V1reportsBioSampleDescription) GetOrganismOk() (*V1reportsOrganism, bool)`

GetOrganismOk returns a tuple with the Organism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganism

`func (o *V1reportsBioSampleDescription) SetOrganism(v V1reportsOrganism)`

SetOrganism sets Organism field to given value.

### HasOrganism

`func (o *V1reportsBioSampleDescription) HasOrganism() bool`

HasOrganism returns a boolean if a field has been set.

### GetComment

`func (o *V1reportsBioSampleDescription) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *V1reportsBioSampleDescription) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *V1reportsBioSampleDescription) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *V1reportsBioSampleDescription) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


