# V1VirusDataReportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**V1VirusDatasetFilter**](V1VirusDatasetFilter.md) |  | [optional] 
**ReturnedContent** | Pointer to [**V1VirusDataReportRequestContentType**](V1VirusDataReportRequestContentType.md) |  | [optional] [default to V1VIRUSDATAREPORTREQUESTCONTENTTYPE_COMPLETE]
**TableFields** | Pointer to **[]string** |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**PageToken** | Pointer to **string** |  | [optional] 

## Methods

### NewV1VirusDataReportRequest

`func NewV1VirusDataReportRequest() *V1VirusDataReportRequest`

NewV1VirusDataReportRequest instantiates a new V1VirusDataReportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1VirusDataReportRequestWithDefaults

`func NewV1VirusDataReportRequestWithDefaults() *V1VirusDataReportRequest`

NewV1VirusDataReportRequestWithDefaults instantiates a new V1VirusDataReportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *V1VirusDataReportRequest) GetFilter() V1VirusDatasetFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *V1VirusDataReportRequest) GetFilterOk() (*V1VirusDatasetFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *V1VirusDataReportRequest) SetFilter(v V1VirusDatasetFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *V1VirusDataReportRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetReturnedContent

`func (o *V1VirusDataReportRequest) GetReturnedContent() V1VirusDataReportRequestContentType`

GetReturnedContent returns the ReturnedContent field if non-nil, zero value otherwise.

### GetReturnedContentOk

`func (o *V1VirusDataReportRequest) GetReturnedContentOk() (*V1VirusDataReportRequestContentType, bool)`

GetReturnedContentOk returns a tuple with the ReturnedContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnedContent

`func (o *V1VirusDataReportRequest) SetReturnedContent(v V1VirusDataReportRequestContentType)`

SetReturnedContent sets ReturnedContent field to given value.

### HasReturnedContent

`func (o *V1VirusDataReportRequest) HasReturnedContent() bool`

HasReturnedContent returns a boolean if a field has been set.

### GetTableFields

`func (o *V1VirusDataReportRequest) GetTableFields() []string`

GetTableFields returns the TableFields field if non-nil, zero value otherwise.

### GetTableFieldsOk

`func (o *V1VirusDataReportRequest) GetTableFieldsOk() (*[]string, bool)`

GetTableFieldsOk returns a tuple with the TableFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableFields

`func (o *V1VirusDataReportRequest) SetTableFields(v []string)`

SetTableFields sets TableFields field to given value.

### HasTableFields

`func (o *V1VirusDataReportRequest) HasTableFields() bool`

HasTableFields returns a boolean if a field has been set.

### GetPageSize

`func (o *V1VirusDataReportRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *V1VirusDataReportRequest) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *V1VirusDataReportRequest) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *V1VirusDataReportRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetPageToken

`func (o *V1VirusDataReportRequest) GetPageToken() string`

GetPageToken returns the PageToken field if non-nil, zero value otherwise.

### GetPageTokenOk

`func (o *V1VirusDataReportRequest) GetPageTokenOk() (*string, bool)`

GetPageTokenOk returns a tuple with the PageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageToken

`func (o *V1VirusDataReportRequest) SetPageToken(v string)`

SetPageToken sets PageToken field to given value.

### HasPageToken

`func (o *V1VirusDataReportRequest) HasPageToken() bool`

HasPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


