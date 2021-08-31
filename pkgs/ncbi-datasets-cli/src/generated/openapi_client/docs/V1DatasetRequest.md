# V1DatasetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GenomeV1** | Pointer to [**V1AssemblyDatasetRequest**](V1AssemblyDatasetRequest.md) |  | [optional] 
**GeneV1** | Pointer to [**V1GeneDatasetRequest**](V1GeneDatasetRequest.md) |  | [optional] 
**VirusV1** | Pointer to [**V1VirusDatasetRequest**](V1VirusDatasetRequest.md) |  | [optional] 

## Methods

### NewV1DatasetRequest

`func NewV1DatasetRequest() *V1DatasetRequest`

NewV1DatasetRequest instantiates a new V1DatasetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DatasetRequestWithDefaults

`func NewV1DatasetRequestWithDefaults() *V1DatasetRequest`

NewV1DatasetRequestWithDefaults instantiates a new V1DatasetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGenomeV1

`func (o *V1DatasetRequest) GetGenomeV1() V1AssemblyDatasetRequest`

GetGenomeV1 returns the GenomeV1 field if non-nil, zero value otherwise.

### GetGenomeV1Ok

`func (o *V1DatasetRequest) GetGenomeV1Ok() (*V1AssemblyDatasetRequest, bool)`

GetGenomeV1Ok returns a tuple with the GenomeV1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenomeV1

`func (o *V1DatasetRequest) SetGenomeV1(v V1AssemblyDatasetRequest)`

SetGenomeV1 sets GenomeV1 field to given value.

### HasGenomeV1

`func (o *V1DatasetRequest) HasGenomeV1() bool`

HasGenomeV1 returns a boolean if a field has been set.

### GetGeneV1

`func (o *V1DatasetRequest) GetGeneV1() V1GeneDatasetRequest`

GetGeneV1 returns the GeneV1 field if non-nil, zero value otherwise.

### GetGeneV1Ok

`func (o *V1DatasetRequest) GetGeneV1Ok() (*V1GeneDatasetRequest, bool)`

GetGeneV1Ok returns a tuple with the GeneV1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneV1

`func (o *V1DatasetRequest) SetGeneV1(v V1GeneDatasetRequest)`

SetGeneV1 sets GeneV1 field to given value.

### HasGeneV1

`func (o *V1DatasetRequest) HasGeneV1() bool`

HasGeneV1 returns a boolean if a field has been set.

### GetVirusV1

`func (o *V1DatasetRequest) GetVirusV1() V1VirusDatasetRequest`

GetVirusV1 returns the VirusV1 field if non-nil, zero value otherwise.

### GetVirusV1Ok

`func (o *V1DatasetRequest) GetVirusV1Ok() (*V1VirusDatasetRequest, bool)`

GetVirusV1Ok returns a tuple with the VirusV1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirusV1

`func (o *V1DatasetRequest) SetVirusV1(v V1VirusDatasetRequest)`

SetVirusV1 sets VirusV1 field to given value.

### HasVirusV1

`func (o *V1DatasetRequest) HasVirusV1() bool`

HasVirusV1 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


