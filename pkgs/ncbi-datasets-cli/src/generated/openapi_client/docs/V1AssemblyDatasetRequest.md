# V1AssemblyDatasetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssemblyAccessions** | Pointer to **[]string** |  | [optional] 
**Accessions** | Pointer to **[]string** |  | [optional] 
**Chromosomes** | Pointer to **[]string** |  | [optional] 
**IncludeAnnotation** | Pointer to **bool** |  | [optional] 
**ExcludeSequence** | Pointer to **bool** |  | [optional] 
**IncludeAnnotationType** | Pointer to [**[]V1AnnotationForAssemblyType**](V1AnnotationForAssemblyType.md) |  | [optional] 
**Hydrated** | Pointer to [**V1AssemblyDatasetRequestResolution**](V1AssemblyDatasetRequestResolution.md) |  | [optional] [default to V1ASSEMBLYDATASETREQUESTRESOLUTION_FULLY_HYDRATED]
**IncludeTsv** | Pointer to **bool** |  | [optional] 

## Methods

### NewV1AssemblyDatasetRequest

`func NewV1AssemblyDatasetRequest() *V1AssemblyDatasetRequest`

NewV1AssemblyDatasetRequest instantiates a new V1AssemblyDatasetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1AssemblyDatasetRequestWithDefaults

`func NewV1AssemblyDatasetRequestWithDefaults() *V1AssemblyDatasetRequest`

NewV1AssemblyDatasetRequestWithDefaults instantiates a new V1AssemblyDatasetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssemblyAccessions

`func (o *V1AssemblyDatasetRequest) GetAssemblyAccessions() []string`

GetAssemblyAccessions returns the AssemblyAccessions field if non-nil, zero value otherwise.

### GetAssemblyAccessionsOk

`func (o *V1AssemblyDatasetRequest) GetAssemblyAccessionsOk() (*[]string, bool)`

GetAssemblyAccessionsOk returns a tuple with the AssemblyAccessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssemblyAccessions

`func (o *V1AssemblyDatasetRequest) SetAssemblyAccessions(v []string)`

SetAssemblyAccessions sets AssemblyAccessions field to given value.

### HasAssemblyAccessions

`func (o *V1AssemblyDatasetRequest) HasAssemblyAccessions() bool`

HasAssemblyAccessions returns a boolean if a field has been set.

### GetAccessions

`func (o *V1AssemblyDatasetRequest) GetAccessions() []string`

GetAccessions returns the Accessions field if non-nil, zero value otherwise.

### GetAccessionsOk

`func (o *V1AssemblyDatasetRequest) GetAccessionsOk() (*[]string, bool)`

GetAccessionsOk returns a tuple with the Accessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessions

`func (o *V1AssemblyDatasetRequest) SetAccessions(v []string)`

SetAccessions sets Accessions field to given value.

### HasAccessions

`func (o *V1AssemblyDatasetRequest) HasAccessions() bool`

HasAccessions returns a boolean if a field has been set.

### GetChromosomes

`func (o *V1AssemblyDatasetRequest) GetChromosomes() []string`

GetChromosomes returns the Chromosomes field if non-nil, zero value otherwise.

### GetChromosomesOk

`func (o *V1AssemblyDatasetRequest) GetChromosomesOk() (*[]string, bool)`

GetChromosomesOk returns a tuple with the Chromosomes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChromosomes

`func (o *V1AssemblyDatasetRequest) SetChromosomes(v []string)`

SetChromosomes sets Chromosomes field to given value.

### HasChromosomes

`func (o *V1AssemblyDatasetRequest) HasChromosomes() bool`

HasChromosomes returns a boolean if a field has been set.

### GetIncludeAnnotation

`func (o *V1AssemblyDatasetRequest) GetIncludeAnnotation() bool`

GetIncludeAnnotation returns the IncludeAnnotation field if non-nil, zero value otherwise.

### GetIncludeAnnotationOk

`func (o *V1AssemblyDatasetRequest) GetIncludeAnnotationOk() (*bool, bool)`

GetIncludeAnnotationOk returns a tuple with the IncludeAnnotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAnnotation

`func (o *V1AssemblyDatasetRequest) SetIncludeAnnotation(v bool)`

SetIncludeAnnotation sets IncludeAnnotation field to given value.

### HasIncludeAnnotation

`func (o *V1AssemblyDatasetRequest) HasIncludeAnnotation() bool`

HasIncludeAnnotation returns a boolean if a field has been set.

### GetExcludeSequence

`func (o *V1AssemblyDatasetRequest) GetExcludeSequence() bool`

GetExcludeSequence returns the ExcludeSequence field if non-nil, zero value otherwise.

### GetExcludeSequenceOk

`func (o *V1AssemblyDatasetRequest) GetExcludeSequenceOk() (*bool, bool)`

GetExcludeSequenceOk returns a tuple with the ExcludeSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeSequence

`func (o *V1AssemblyDatasetRequest) SetExcludeSequence(v bool)`

SetExcludeSequence sets ExcludeSequence field to given value.

### HasExcludeSequence

`func (o *V1AssemblyDatasetRequest) HasExcludeSequence() bool`

HasExcludeSequence returns a boolean if a field has been set.

### GetIncludeAnnotationType

`func (o *V1AssemblyDatasetRequest) GetIncludeAnnotationType() []V1AnnotationForAssemblyType`

GetIncludeAnnotationType returns the IncludeAnnotationType field if non-nil, zero value otherwise.

### GetIncludeAnnotationTypeOk

`func (o *V1AssemblyDatasetRequest) GetIncludeAnnotationTypeOk() (*[]V1AnnotationForAssemblyType, bool)`

GetIncludeAnnotationTypeOk returns a tuple with the IncludeAnnotationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAnnotationType

`func (o *V1AssemblyDatasetRequest) SetIncludeAnnotationType(v []V1AnnotationForAssemblyType)`

SetIncludeAnnotationType sets IncludeAnnotationType field to given value.

### HasIncludeAnnotationType

`func (o *V1AssemblyDatasetRequest) HasIncludeAnnotationType() bool`

HasIncludeAnnotationType returns a boolean if a field has been set.

### GetHydrated

`func (o *V1AssemblyDatasetRequest) GetHydrated() V1AssemblyDatasetRequestResolution`

GetHydrated returns the Hydrated field if non-nil, zero value otherwise.

### GetHydratedOk

`func (o *V1AssemblyDatasetRequest) GetHydratedOk() (*V1AssemblyDatasetRequestResolution, bool)`

GetHydratedOk returns a tuple with the Hydrated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHydrated

`func (o *V1AssemblyDatasetRequest) SetHydrated(v V1AssemblyDatasetRequestResolution)`

SetHydrated sets Hydrated field to given value.

### HasHydrated

`func (o *V1AssemblyDatasetRequest) HasHydrated() bool`

HasHydrated returns a boolean if a field has been set.

### GetIncludeTsv

`func (o *V1AssemblyDatasetRequest) GetIncludeTsv() bool`

GetIncludeTsv returns the IncludeTsv field if non-nil, zero value otherwise.

### GetIncludeTsvOk

`func (o *V1AssemblyDatasetRequest) GetIncludeTsvOk() (*bool, bool)`

GetIncludeTsvOk returns a tuple with the IncludeTsv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTsv

`func (o *V1AssemblyDatasetRequest) SetIncludeTsv(v bool)`

SetIncludeTsv sets IncludeTsv field to given value.

### HasIncludeTsv

`func (o *V1AssemblyDatasetRequest) HasIncludeTsv() bool`

HasIncludeTsv returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


