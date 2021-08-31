# V1AnnotationForAssembly

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**ReleaseNumber** | Pointer to **string** |  | [optional] 
**ReleaseDate** | Pointer to **string** |  | [optional] 
**ReportUrl** | Pointer to **string** |  | [optional] 
**File** | Pointer to [**[]V1AnnotationForAssemblyFile**](V1AnnotationForAssemblyFile.md) |  | [optional] 
**Stats** | Pointer to [**V1FeatureCounts**](V1FeatureCounts.md) |  | [optional] 
**Busco** | Pointer to [**V1BuscoStat**](V1BuscoStat.md) |  | [optional] 

## Methods

### NewV1AnnotationForAssembly

`func NewV1AnnotationForAssembly() *V1AnnotationForAssembly`

NewV1AnnotationForAssembly instantiates a new V1AnnotationForAssembly object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1AnnotationForAssemblyWithDefaults

`func NewV1AnnotationForAssemblyWithDefaults() *V1AnnotationForAssembly`

NewV1AnnotationForAssemblyWithDefaults instantiates a new V1AnnotationForAssembly object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *V1AnnotationForAssembly) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1AnnotationForAssembly) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1AnnotationForAssembly) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1AnnotationForAssembly) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSource

`func (o *V1AnnotationForAssembly) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *V1AnnotationForAssembly) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *V1AnnotationForAssembly) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *V1AnnotationForAssembly) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetReleaseNumber

`func (o *V1AnnotationForAssembly) GetReleaseNumber() string`

GetReleaseNumber returns the ReleaseNumber field if non-nil, zero value otherwise.

### GetReleaseNumberOk

`func (o *V1AnnotationForAssembly) GetReleaseNumberOk() (*string, bool)`

GetReleaseNumberOk returns a tuple with the ReleaseNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseNumber

`func (o *V1AnnotationForAssembly) SetReleaseNumber(v string)`

SetReleaseNumber sets ReleaseNumber field to given value.

### HasReleaseNumber

`func (o *V1AnnotationForAssembly) HasReleaseNumber() bool`

HasReleaseNumber returns a boolean if a field has been set.

### GetReleaseDate

`func (o *V1AnnotationForAssembly) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *V1AnnotationForAssembly) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *V1AnnotationForAssembly) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *V1AnnotationForAssembly) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### GetReportUrl

`func (o *V1AnnotationForAssembly) GetReportUrl() string`

GetReportUrl returns the ReportUrl field if non-nil, zero value otherwise.

### GetReportUrlOk

`func (o *V1AnnotationForAssembly) GetReportUrlOk() (*string, bool)`

GetReportUrlOk returns a tuple with the ReportUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportUrl

`func (o *V1AnnotationForAssembly) SetReportUrl(v string)`

SetReportUrl sets ReportUrl field to given value.

### HasReportUrl

`func (o *V1AnnotationForAssembly) HasReportUrl() bool`

HasReportUrl returns a boolean if a field has been set.

### GetFile

`func (o *V1AnnotationForAssembly) GetFile() []V1AnnotationForAssemblyFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *V1AnnotationForAssembly) GetFileOk() (*[]V1AnnotationForAssemblyFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *V1AnnotationForAssembly) SetFile(v []V1AnnotationForAssemblyFile)`

SetFile sets File field to given value.

### HasFile

`func (o *V1AnnotationForAssembly) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetStats

`func (o *V1AnnotationForAssembly) GetStats() V1FeatureCounts`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *V1AnnotationForAssembly) GetStatsOk() (*V1FeatureCounts, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *V1AnnotationForAssembly) SetStats(v V1FeatureCounts)`

SetStats sets Stats field to given value.

### HasStats

`func (o *V1AnnotationForAssembly) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetBusco

`func (o *V1AnnotationForAssembly) GetBusco() V1BuscoStat`

GetBusco returns the Busco field if non-nil, zero value otherwise.

### GetBuscoOk

`func (o *V1AnnotationForAssembly) GetBuscoOk() (*V1BuscoStat, bool)`

GetBuscoOk returns a tuple with the Busco field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusco

`func (o *V1AnnotationForAssembly) SetBusco(v V1BuscoStat)`

SetBusco sets Busco field to given value.

### HasBusco

`func (o *V1AnnotationForAssembly) HasBusco() bool`

HasBusco returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


