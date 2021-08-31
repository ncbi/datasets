# V1Annotation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReleaseName** | Pointer to **string** |  | [optional] 
**ReleaseDate** | Pointer to **string** |  | [optional] 
**AssembliesInScope** | Pointer to [**[]V1AnnotatedAssemblies**](V1AnnotatedAssemblies.md) |  | [optional] 

## Methods

### NewV1Annotation

`func NewV1Annotation() *V1Annotation`

NewV1Annotation instantiates a new V1Annotation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1AnnotationWithDefaults

`func NewV1AnnotationWithDefaults() *V1Annotation`

NewV1AnnotationWithDefaults instantiates a new V1Annotation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReleaseName

`func (o *V1Annotation) GetReleaseName() string`

GetReleaseName returns the ReleaseName field if non-nil, zero value otherwise.

### GetReleaseNameOk

`func (o *V1Annotation) GetReleaseNameOk() (*string, bool)`

GetReleaseNameOk returns a tuple with the ReleaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseName

`func (o *V1Annotation) SetReleaseName(v string)`

SetReleaseName sets ReleaseName field to given value.

### HasReleaseName

`func (o *V1Annotation) HasReleaseName() bool`

HasReleaseName returns a boolean if a field has been set.

### GetReleaseDate

`func (o *V1Annotation) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *V1Annotation) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *V1Annotation) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *V1Annotation) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### GetAssembliesInScope

`func (o *V1Annotation) GetAssembliesInScope() []V1AnnotatedAssemblies`

GetAssembliesInScope returns the AssembliesInScope field if non-nil, zero value otherwise.

### GetAssembliesInScopeOk

`func (o *V1Annotation) GetAssembliesInScopeOk() (*[]V1AnnotatedAssemblies, bool)`

GetAssembliesInScopeOk returns a tuple with the AssembliesInScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssembliesInScope

`func (o *V1Annotation) SetAssembliesInScope(v []V1AnnotatedAssemblies)`

SetAssembliesInScope sets AssembliesInScope field to given value.

### HasAssembliesInScope

`func (o *V1Annotation) HasAssembliesInScope() bool`

HasAssembliesInScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


