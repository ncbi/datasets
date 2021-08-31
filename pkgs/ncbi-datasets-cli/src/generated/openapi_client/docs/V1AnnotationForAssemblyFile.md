# V1AnnotationForAssemblyFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**V1AnnotationForAssemblyType**](V1AnnotationForAssemblyType.md) |  | [optional] [default to V1ANNOTATIONFORASSEMBLYTYPE_DEFAULT]
**EstimatedSize** | Pointer to **string** |  | [optional] 

## Methods

### NewV1AnnotationForAssemblyFile

`func NewV1AnnotationForAssemblyFile() *V1AnnotationForAssemblyFile`

NewV1AnnotationForAssemblyFile instantiates a new V1AnnotationForAssemblyFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1AnnotationForAssemblyFileWithDefaults

`func NewV1AnnotationForAssemblyFileWithDefaults() *V1AnnotationForAssemblyFile`

NewV1AnnotationForAssemblyFileWithDefaults instantiates a new V1AnnotationForAssemblyFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *V1AnnotationForAssemblyFile) GetType() V1AnnotationForAssemblyType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1AnnotationForAssemblyFile) GetTypeOk() (*V1AnnotationForAssemblyType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1AnnotationForAssemblyFile) SetType(v V1AnnotationForAssemblyType)`

SetType sets Type field to given value.

### HasType

`func (o *V1AnnotationForAssemblyFile) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEstimatedSize

`func (o *V1AnnotationForAssemblyFile) GetEstimatedSize() string`

GetEstimatedSize returns the EstimatedSize field if non-nil, zero value otherwise.

### GetEstimatedSizeOk

`func (o *V1AnnotationForAssemblyFile) GetEstimatedSizeOk() (*string, bool)`

GetEstimatedSizeOk returns a tuple with the EstimatedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedSize

`func (o *V1AnnotationForAssemblyFile) SetEstimatedSize(v string)`

SetEstimatedSize sets EstimatedSize field to given value.

### HasEstimatedSize

`func (o *V1AnnotationForAssemblyFile) HasEstimatedSize() bool`

HasEstimatedSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


