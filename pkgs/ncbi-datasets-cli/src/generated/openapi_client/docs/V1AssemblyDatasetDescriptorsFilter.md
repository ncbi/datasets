# V1AssemblyDatasetDescriptorsFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReferenceOnly** | Pointer to **bool** |  | [optional] 
**RefseqOnly** | Pointer to **bool** |  | [optional] 
**AssemblySource** | Pointer to [**V1AssemblyDatasetDescriptorsFilterAssemblySource**](V1AssemblyDatasetDescriptorsFilterAssemblySource.md) |  | [optional] [default to V1ASSEMBLYDATASETDESCRIPTORSFILTERASSEMBLYSOURCE_ALL]
**HasAnnotation** | Pointer to **bool** |  | [optional] 
**AssemblyLevel** | Pointer to [**[]V1AssemblyDatasetDescriptorsFilterAssemblyLevel**](V1AssemblyDatasetDescriptorsFilterAssemblyLevel.md) |  | [optional] 
**FirstReleaseDate** | Pointer to **time.Time** |  | [optional] 
**LastReleaseDate** | Pointer to **time.Time** |  | [optional] 
**SearchText** | Pointer to **[]string** |  | [optional] 

## Methods

### NewV1AssemblyDatasetDescriptorsFilter

`func NewV1AssemblyDatasetDescriptorsFilter() *V1AssemblyDatasetDescriptorsFilter`

NewV1AssemblyDatasetDescriptorsFilter instantiates a new V1AssemblyDatasetDescriptorsFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1AssemblyDatasetDescriptorsFilterWithDefaults

`func NewV1AssemblyDatasetDescriptorsFilterWithDefaults() *V1AssemblyDatasetDescriptorsFilter`

NewV1AssemblyDatasetDescriptorsFilterWithDefaults instantiates a new V1AssemblyDatasetDescriptorsFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferenceOnly

`func (o *V1AssemblyDatasetDescriptorsFilter) GetReferenceOnly() bool`

GetReferenceOnly returns the ReferenceOnly field if non-nil, zero value otherwise.

### GetReferenceOnlyOk

`func (o *V1AssemblyDatasetDescriptorsFilter) GetReferenceOnlyOk() (*bool, bool)`

GetReferenceOnlyOk returns a tuple with the ReferenceOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceOnly

`func (o *V1AssemblyDatasetDescriptorsFilter) SetReferenceOnly(v bool)`

SetReferenceOnly sets ReferenceOnly field to given value.

### HasReferenceOnly

`func (o *V1AssemblyDatasetDescriptorsFilter) HasReferenceOnly() bool`

HasReferenceOnly returns a boolean if a field has been set.

### GetRefseqOnly

`func (o *V1AssemblyDatasetDescriptorsFilter) GetRefseqOnly() bool`

GetRefseqOnly returns the RefseqOnly field if non-nil, zero value otherwise.

### GetRefseqOnlyOk

`func (o *V1AssemblyDatasetDescriptorsFilter) GetRefseqOnlyOk() (*bool, bool)`

GetRefseqOnlyOk returns a tuple with the RefseqOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefseqOnly

`func (o *V1AssemblyDatasetDescriptorsFilter) SetRefseqOnly(v bool)`

SetRefseqOnly sets RefseqOnly field to given value.

### HasRefseqOnly

`func (o *V1AssemblyDatasetDescriptorsFilter) HasRefseqOnly() bool`

HasRefseqOnly returns a boolean if a field has been set.

### GetAssemblySource

`func (o *V1AssemblyDatasetDescriptorsFilter) GetAssemblySource() V1AssemblyDatasetDescriptorsFilterAssemblySource`

GetAssemblySource returns the AssemblySource field if non-nil, zero value otherwise.

### GetAssemblySourceOk

`func (o *V1AssemblyDatasetDescriptorsFilter) GetAssemblySourceOk() (*V1AssemblyDatasetDescriptorsFilterAssemblySource, bool)`

GetAssemblySourceOk returns a tuple with the AssemblySource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssemblySource

`func (o *V1AssemblyDatasetDescriptorsFilter) SetAssemblySource(v V1AssemblyDatasetDescriptorsFilterAssemblySource)`

SetAssemblySource sets AssemblySource field to given value.

### HasAssemblySource

`func (o *V1AssemblyDatasetDescriptorsFilter) HasAssemblySource() bool`

HasAssemblySource returns a boolean if a field has been set.

### GetHasAnnotation

`func (o *V1AssemblyDatasetDescriptorsFilter) GetHasAnnotation() bool`

GetHasAnnotation returns the HasAnnotation field if non-nil, zero value otherwise.

### GetHasAnnotationOk

`func (o *V1AssemblyDatasetDescriptorsFilter) GetHasAnnotationOk() (*bool, bool)`

GetHasAnnotationOk returns a tuple with the HasAnnotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAnnotation

`func (o *V1AssemblyDatasetDescriptorsFilter) SetHasAnnotation(v bool)`

SetHasAnnotation sets HasAnnotation field to given value.

### HasHasAnnotation

`func (o *V1AssemblyDatasetDescriptorsFilter) HasHasAnnotation() bool`

HasHasAnnotation returns a boolean if a field has been set.

### GetAssemblyLevel

`func (o *V1AssemblyDatasetDescriptorsFilter) GetAssemblyLevel() []V1AssemblyDatasetDescriptorsFilterAssemblyLevel`

GetAssemblyLevel returns the AssemblyLevel field if non-nil, zero value otherwise.

### GetAssemblyLevelOk

`func (o *V1AssemblyDatasetDescriptorsFilter) GetAssemblyLevelOk() (*[]V1AssemblyDatasetDescriptorsFilterAssemblyLevel, bool)`

GetAssemblyLevelOk returns a tuple with the AssemblyLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssemblyLevel

`func (o *V1AssemblyDatasetDescriptorsFilter) SetAssemblyLevel(v []V1AssemblyDatasetDescriptorsFilterAssemblyLevel)`

SetAssemblyLevel sets AssemblyLevel field to given value.

### HasAssemblyLevel

`func (o *V1AssemblyDatasetDescriptorsFilter) HasAssemblyLevel() bool`

HasAssemblyLevel returns a boolean if a field has been set.

### GetFirstReleaseDate

`func (o *V1AssemblyDatasetDescriptorsFilter) GetFirstReleaseDate() time.Time`

GetFirstReleaseDate returns the FirstReleaseDate field if non-nil, zero value otherwise.

### GetFirstReleaseDateOk

`func (o *V1AssemblyDatasetDescriptorsFilter) GetFirstReleaseDateOk() (*time.Time, bool)`

GetFirstReleaseDateOk returns a tuple with the FirstReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstReleaseDate

`func (o *V1AssemblyDatasetDescriptorsFilter) SetFirstReleaseDate(v time.Time)`

SetFirstReleaseDate sets FirstReleaseDate field to given value.

### HasFirstReleaseDate

`func (o *V1AssemblyDatasetDescriptorsFilter) HasFirstReleaseDate() bool`

HasFirstReleaseDate returns a boolean if a field has been set.

### GetLastReleaseDate

`func (o *V1AssemblyDatasetDescriptorsFilter) GetLastReleaseDate() time.Time`

GetLastReleaseDate returns the LastReleaseDate field if non-nil, zero value otherwise.

### GetLastReleaseDateOk

`func (o *V1AssemblyDatasetDescriptorsFilter) GetLastReleaseDateOk() (*time.Time, bool)`

GetLastReleaseDateOk returns a tuple with the LastReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReleaseDate

`func (o *V1AssemblyDatasetDescriptorsFilter) SetLastReleaseDate(v time.Time)`

SetLastReleaseDate sets LastReleaseDate field to given value.

### HasLastReleaseDate

`func (o *V1AssemblyDatasetDescriptorsFilter) HasLastReleaseDate() bool`

HasLastReleaseDate returns a boolean if a field has been set.

### GetSearchText

`func (o *V1AssemblyDatasetDescriptorsFilter) GetSearchText() []string`

GetSearchText returns the SearchText field if non-nil, zero value otherwise.

### GetSearchTextOk

`func (o *V1AssemblyDatasetDescriptorsFilter) GetSearchTextOk() (*[]string, bool)`

GetSearchTextOk returns a tuple with the SearchText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchText

`func (o *V1AssemblyDatasetDescriptorsFilter) SetSearchText(v []string)`

SetSearchText sets SearchText field to given value.

### HasSearchText

`func (o *V1AssemblyDatasetDescriptorsFilter) HasSearchText() bool`

HasSearchText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


