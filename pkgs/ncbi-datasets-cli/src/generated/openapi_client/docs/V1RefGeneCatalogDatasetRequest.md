# V1RefGeneCatalogDatasetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OpaqueSolrQuery** | Pointer to **string** |  | [optional] 
**Files** | Pointer to [**[]V1RefGeneCatalogDatasetRequestFileType**](V1RefGeneCatalogDatasetRequestFileType.md) |  | [optional] 
**ElementFlankConfig** | Pointer to [**V1ElementFlankConfig**](V1ElementFlankConfig.md) |  | [optional] 

## Methods

### NewV1RefGeneCatalogDatasetRequest

`func NewV1RefGeneCatalogDatasetRequest() *V1RefGeneCatalogDatasetRequest`

NewV1RefGeneCatalogDatasetRequest instantiates a new V1RefGeneCatalogDatasetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1RefGeneCatalogDatasetRequestWithDefaults

`func NewV1RefGeneCatalogDatasetRequestWithDefaults() *V1RefGeneCatalogDatasetRequest`

NewV1RefGeneCatalogDatasetRequestWithDefaults instantiates a new V1RefGeneCatalogDatasetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpaqueSolrQuery

`func (o *V1RefGeneCatalogDatasetRequest) GetOpaqueSolrQuery() string`

GetOpaqueSolrQuery returns the OpaqueSolrQuery field if non-nil, zero value otherwise.

### GetOpaqueSolrQueryOk

`func (o *V1RefGeneCatalogDatasetRequest) GetOpaqueSolrQueryOk() (*string, bool)`

GetOpaqueSolrQueryOk returns a tuple with the OpaqueSolrQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpaqueSolrQuery

`func (o *V1RefGeneCatalogDatasetRequest) SetOpaqueSolrQuery(v string)`

SetOpaqueSolrQuery sets OpaqueSolrQuery field to given value.

### HasOpaqueSolrQuery

`func (o *V1RefGeneCatalogDatasetRequest) HasOpaqueSolrQuery() bool`

HasOpaqueSolrQuery returns a boolean if a field has been set.

### GetFiles

`func (o *V1RefGeneCatalogDatasetRequest) GetFiles() []V1RefGeneCatalogDatasetRequestFileType`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *V1RefGeneCatalogDatasetRequest) GetFilesOk() (*[]V1RefGeneCatalogDatasetRequestFileType, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *V1RefGeneCatalogDatasetRequest) SetFiles(v []V1RefGeneCatalogDatasetRequestFileType)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *V1RefGeneCatalogDatasetRequest) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetElementFlankConfig

`func (o *V1RefGeneCatalogDatasetRequest) GetElementFlankConfig() V1ElementFlankConfig`

GetElementFlankConfig returns the ElementFlankConfig field if non-nil, zero value otherwise.

### GetElementFlankConfigOk

`func (o *V1RefGeneCatalogDatasetRequest) GetElementFlankConfigOk() (*V1ElementFlankConfig, bool)`

GetElementFlankConfigOk returns a tuple with the ElementFlankConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementFlankConfig

`func (o *V1RefGeneCatalogDatasetRequest) SetElementFlankConfig(v V1ElementFlankConfig)`

SetElementFlankConfig sets ElementFlankConfig field to given value.

### HasElementFlankConfig

`func (o *V1RefGeneCatalogDatasetRequest) HasElementFlankConfig() bool`

HasElementFlankConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


