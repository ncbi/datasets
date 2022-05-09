# V1MicroBiggeDatasetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OpaqueSolrQuery** | Pointer to **string** |  | [optional] 
**Files** | Pointer to [**[]V1MicroBiggeDatasetRequestFileType**](V1MicroBiggeDatasetRequestFileType.md) |  | [optional] 
**ElementFlankConfig** | Pointer to [**V1ElementFlankConfig**](V1ElementFlankConfig.md) |  | [optional] 

## Methods

### NewV1MicroBiggeDatasetRequest

`func NewV1MicroBiggeDatasetRequest() *V1MicroBiggeDatasetRequest`

NewV1MicroBiggeDatasetRequest instantiates a new V1MicroBiggeDatasetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1MicroBiggeDatasetRequestWithDefaults

`func NewV1MicroBiggeDatasetRequestWithDefaults() *V1MicroBiggeDatasetRequest`

NewV1MicroBiggeDatasetRequestWithDefaults instantiates a new V1MicroBiggeDatasetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpaqueSolrQuery

`func (o *V1MicroBiggeDatasetRequest) GetOpaqueSolrQuery() string`

GetOpaqueSolrQuery returns the OpaqueSolrQuery field if non-nil, zero value otherwise.

### GetOpaqueSolrQueryOk

`func (o *V1MicroBiggeDatasetRequest) GetOpaqueSolrQueryOk() (*string, bool)`

GetOpaqueSolrQueryOk returns a tuple with the OpaqueSolrQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpaqueSolrQuery

`func (o *V1MicroBiggeDatasetRequest) SetOpaqueSolrQuery(v string)`

SetOpaqueSolrQuery sets OpaqueSolrQuery field to given value.

### HasOpaqueSolrQuery

`func (o *V1MicroBiggeDatasetRequest) HasOpaqueSolrQuery() bool`

HasOpaqueSolrQuery returns a boolean if a field has been set.

### GetFiles

`func (o *V1MicroBiggeDatasetRequest) GetFiles() []V1MicroBiggeDatasetRequestFileType`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *V1MicroBiggeDatasetRequest) GetFilesOk() (*[]V1MicroBiggeDatasetRequestFileType, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *V1MicroBiggeDatasetRequest) SetFiles(v []V1MicroBiggeDatasetRequestFileType)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *V1MicroBiggeDatasetRequest) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetElementFlankConfig

`func (o *V1MicroBiggeDatasetRequest) GetElementFlankConfig() V1ElementFlankConfig`

GetElementFlankConfig returns the ElementFlankConfig field if non-nil, zero value otherwise.

### GetElementFlankConfigOk

`func (o *V1MicroBiggeDatasetRequest) GetElementFlankConfigOk() (*V1ElementFlankConfig, bool)`

GetElementFlankConfigOk returns a tuple with the ElementFlankConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementFlankConfig

`func (o *V1MicroBiggeDatasetRequest) SetElementFlankConfig(v V1ElementFlankConfig)`

SetElementFlankConfig sets ElementFlankConfig field to given value.

### HasElementFlankConfig

`func (o *V1MicroBiggeDatasetRequest) HasElementFlankConfig() bool`

HasElementFlankConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


