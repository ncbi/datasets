# V1reportsAnnotationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**ReleaseDate** | Pointer to **string** |  | [optional] 
**ReportUrl** | Pointer to **string** |  | [optional] 
**Stats** | Pointer to [**V1reportsFeatureCounts**](V1reportsFeatureCounts.md) |  | [optional] 
**Busco** | Pointer to [**V1reportsBuscoStat**](V1reportsBuscoStat.md) |  | [optional] 

## Methods

### NewV1reportsAnnotationInfo

`func NewV1reportsAnnotationInfo() *V1reportsAnnotationInfo`

NewV1reportsAnnotationInfo instantiates a new V1reportsAnnotationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1reportsAnnotationInfoWithDefaults

`func NewV1reportsAnnotationInfoWithDefaults() *V1reportsAnnotationInfo`

NewV1reportsAnnotationInfoWithDefaults instantiates a new V1reportsAnnotationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *V1reportsAnnotationInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1reportsAnnotationInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1reportsAnnotationInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1reportsAnnotationInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSource

`func (o *V1reportsAnnotationInfo) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *V1reportsAnnotationInfo) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *V1reportsAnnotationInfo) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *V1reportsAnnotationInfo) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetReleaseDate

`func (o *V1reportsAnnotationInfo) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *V1reportsAnnotationInfo) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *V1reportsAnnotationInfo) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *V1reportsAnnotationInfo) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### GetReportUrl

`func (o *V1reportsAnnotationInfo) GetReportUrl() string`

GetReportUrl returns the ReportUrl field if non-nil, zero value otherwise.

### GetReportUrlOk

`func (o *V1reportsAnnotationInfo) GetReportUrlOk() (*string, bool)`

GetReportUrlOk returns a tuple with the ReportUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportUrl

`func (o *V1reportsAnnotationInfo) SetReportUrl(v string)`

SetReportUrl sets ReportUrl field to given value.

### HasReportUrl

`func (o *V1reportsAnnotationInfo) HasReportUrl() bool`

HasReportUrl returns a boolean if a field has been set.

### GetStats

`func (o *V1reportsAnnotationInfo) GetStats() V1reportsFeatureCounts`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *V1reportsAnnotationInfo) GetStatsOk() (*V1reportsFeatureCounts, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *V1reportsAnnotationInfo) SetStats(v V1reportsFeatureCounts)`

SetStats sets Stats field to given value.

### HasStats

`func (o *V1reportsAnnotationInfo) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetBusco

`func (o *V1reportsAnnotationInfo) GetBusco() V1reportsBuscoStat`

GetBusco returns the Busco field if non-nil, zero value otherwise.

### GetBuscoOk

`func (o *V1reportsAnnotationInfo) GetBuscoOk() (*V1reportsBuscoStat, bool)`

GetBuscoOk returns a tuple with the Busco field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusco

`func (o *V1reportsAnnotationInfo) SetBusco(v V1reportsBuscoStat)`

SetBusco sets Busco field to given value.

### HasBusco

`func (o *V1reportsAnnotationInfo) HasBusco() bool`

HasBusco returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


