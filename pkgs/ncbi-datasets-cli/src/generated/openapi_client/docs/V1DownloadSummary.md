# V1DownloadSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordCount** | Pointer to **int32** |  | [optional] 
**AssemblyCount** | Pointer to **int32** |  | [optional] 
**ResourceUpdatedOn** | Pointer to **time.Time** |  | [optional] 
**Hydrated** | Pointer to [**V1DownloadSummaryHydrated**](V1DownloadSummaryHydrated.md) |  | [optional] 
**Dehydrated** | Pointer to [**V1DownloadSummaryDehydrated**](V1DownloadSummaryDehydrated.md) |  | [optional] 
**Errors** | Pointer to [**[]V1Error**](V1Error.md) |  | [optional] 
**Messages** | Pointer to [**[]V1Message**](V1Message.md) |  | [optional] 
**AvailableFiles** | Pointer to [**V1DownloadSummaryAvailableFiles**](V1DownloadSummaryAvailableFiles.md) |  | [optional] 

## Methods

### NewV1DownloadSummary

`func NewV1DownloadSummary() *V1DownloadSummary`

NewV1DownloadSummary instantiates a new V1DownloadSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DownloadSummaryWithDefaults

`func NewV1DownloadSummaryWithDefaults() *V1DownloadSummary`

NewV1DownloadSummaryWithDefaults instantiates a new V1DownloadSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordCount

`func (o *V1DownloadSummary) GetRecordCount() int32`

GetRecordCount returns the RecordCount field if non-nil, zero value otherwise.

### GetRecordCountOk

`func (o *V1DownloadSummary) GetRecordCountOk() (*int32, bool)`

GetRecordCountOk returns a tuple with the RecordCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordCount

`func (o *V1DownloadSummary) SetRecordCount(v int32)`

SetRecordCount sets RecordCount field to given value.

### HasRecordCount

`func (o *V1DownloadSummary) HasRecordCount() bool`

HasRecordCount returns a boolean if a field has been set.

### GetAssemblyCount

`func (o *V1DownloadSummary) GetAssemblyCount() int32`

GetAssemblyCount returns the AssemblyCount field if non-nil, zero value otherwise.

### GetAssemblyCountOk

`func (o *V1DownloadSummary) GetAssemblyCountOk() (*int32, bool)`

GetAssemblyCountOk returns a tuple with the AssemblyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssemblyCount

`func (o *V1DownloadSummary) SetAssemblyCount(v int32)`

SetAssemblyCount sets AssemblyCount field to given value.

### HasAssemblyCount

`func (o *V1DownloadSummary) HasAssemblyCount() bool`

HasAssemblyCount returns a boolean if a field has been set.

### GetResourceUpdatedOn

`func (o *V1DownloadSummary) GetResourceUpdatedOn() time.Time`

GetResourceUpdatedOn returns the ResourceUpdatedOn field if non-nil, zero value otherwise.

### GetResourceUpdatedOnOk

`func (o *V1DownloadSummary) GetResourceUpdatedOnOk() (*time.Time, bool)`

GetResourceUpdatedOnOk returns a tuple with the ResourceUpdatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceUpdatedOn

`func (o *V1DownloadSummary) SetResourceUpdatedOn(v time.Time)`

SetResourceUpdatedOn sets ResourceUpdatedOn field to given value.

### HasResourceUpdatedOn

`func (o *V1DownloadSummary) HasResourceUpdatedOn() bool`

HasResourceUpdatedOn returns a boolean if a field has been set.

### GetHydrated

`func (o *V1DownloadSummary) GetHydrated() V1DownloadSummaryHydrated`

GetHydrated returns the Hydrated field if non-nil, zero value otherwise.

### GetHydratedOk

`func (o *V1DownloadSummary) GetHydratedOk() (*V1DownloadSummaryHydrated, bool)`

GetHydratedOk returns a tuple with the Hydrated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHydrated

`func (o *V1DownloadSummary) SetHydrated(v V1DownloadSummaryHydrated)`

SetHydrated sets Hydrated field to given value.

### HasHydrated

`func (o *V1DownloadSummary) HasHydrated() bool`

HasHydrated returns a boolean if a field has been set.

### GetDehydrated

`func (o *V1DownloadSummary) GetDehydrated() V1DownloadSummaryDehydrated`

GetDehydrated returns the Dehydrated field if non-nil, zero value otherwise.

### GetDehydratedOk

`func (o *V1DownloadSummary) GetDehydratedOk() (*V1DownloadSummaryDehydrated, bool)`

GetDehydratedOk returns a tuple with the Dehydrated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDehydrated

`func (o *V1DownloadSummary) SetDehydrated(v V1DownloadSummaryDehydrated)`

SetDehydrated sets Dehydrated field to given value.

### HasDehydrated

`func (o *V1DownloadSummary) HasDehydrated() bool`

HasDehydrated returns a boolean if a field has been set.

### GetErrors

`func (o *V1DownloadSummary) GetErrors() []V1Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V1DownloadSummary) GetErrorsOk() (*[]V1Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V1DownloadSummary) SetErrors(v []V1Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V1DownloadSummary) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetMessages

`func (o *V1DownloadSummary) GetMessages() []V1Message`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *V1DownloadSummary) GetMessagesOk() (*[]V1Message, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *V1DownloadSummary) SetMessages(v []V1Message)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *V1DownloadSummary) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetAvailableFiles

`func (o *V1DownloadSummary) GetAvailableFiles() V1DownloadSummaryAvailableFiles`

GetAvailableFiles returns the AvailableFiles field if non-nil, zero value otherwise.

### GetAvailableFilesOk

`func (o *V1DownloadSummary) GetAvailableFilesOk() (*V1DownloadSummaryAvailableFiles, bool)`

GetAvailableFilesOk returns a tuple with the AvailableFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableFiles

`func (o *V1DownloadSummary) SetAvailableFiles(v V1DownloadSummaryAvailableFiles)`

SetAvailableFiles sets AvailableFiles field to given value.

### HasAvailableFiles

`func (o *V1DownloadSummary) HasAvailableFiles() bool`

HasAvailableFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


