# V1TaxonomyFilteredSubtreeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RootNodes** | Pointer to **[]int32** |  | [optional] 
**Edges** | Pointer to [**V1TaxonomyFilteredSubtreeResponseEdgesEntry**](V1TaxonomyFilteredSubtreeResponseEdgesEntry.md) |  | [optional] 
**Warnings** | Pointer to [**[]V1Warning**](V1Warning.md) |  | [optional] 
**Errors** | Pointer to [**[]V1Error**](V1Error.md) |  | [optional] 

## Methods

### NewV1TaxonomyFilteredSubtreeResponse

`func NewV1TaxonomyFilteredSubtreeResponse() *V1TaxonomyFilteredSubtreeResponse`

NewV1TaxonomyFilteredSubtreeResponse instantiates a new V1TaxonomyFilteredSubtreeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TaxonomyFilteredSubtreeResponseWithDefaults

`func NewV1TaxonomyFilteredSubtreeResponseWithDefaults() *V1TaxonomyFilteredSubtreeResponse`

NewV1TaxonomyFilteredSubtreeResponseWithDefaults instantiates a new V1TaxonomyFilteredSubtreeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRootNodes

`func (o *V1TaxonomyFilteredSubtreeResponse) GetRootNodes() []int32`

GetRootNodes returns the RootNodes field if non-nil, zero value otherwise.

### GetRootNodesOk

`func (o *V1TaxonomyFilteredSubtreeResponse) GetRootNodesOk() (*[]int32, bool)`

GetRootNodesOk returns a tuple with the RootNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootNodes

`func (o *V1TaxonomyFilteredSubtreeResponse) SetRootNodes(v []int32)`

SetRootNodes sets RootNodes field to given value.

### HasRootNodes

`func (o *V1TaxonomyFilteredSubtreeResponse) HasRootNodes() bool`

HasRootNodes returns a boolean if a field has been set.

### GetEdges

`func (o *V1TaxonomyFilteredSubtreeResponse) GetEdges() V1TaxonomyFilteredSubtreeResponseEdgesEntry`

GetEdges returns the Edges field if non-nil, zero value otherwise.

### GetEdgesOk

`func (o *V1TaxonomyFilteredSubtreeResponse) GetEdgesOk() (*V1TaxonomyFilteredSubtreeResponseEdgesEntry, bool)`

GetEdgesOk returns a tuple with the Edges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdges

`func (o *V1TaxonomyFilteredSubtreeResponse) SetEdges(v V1TaxonomyFilteredSubtreeResponseEdgesEntry)`

SetEdges sets Edges field to given value.

### HasEdges

`func (o *V1TaxonomyFilteredSubtreeResponse) HasEdges() bool`

HasEdges returns a boolean if a field has been set.

### GetWarnings

`func (o *V1TaxonomyFilteredSubtreeResponse) GetWarnings() []V1Warning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V1TaxonomyFilteredSubtreeResponse) GetWarningsOk() (*[]V1Warning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V1TaxonomyFilteredSubtreeResponse) SetWarnings(v []V1Warning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V1TaxonomyFilteredSubtreeResponse) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetErrors

`func (o *V1TaxonomyFilteredSubtreeResponse) GetErrors() []V1Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V1TaxonomyFilteredSubtreeResponse) GetErrorsOk() (*[]V1Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V1TaxonomyFilteredSubtreeResponse) SetErrors(v []V1Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V1TaxonomyFilteredSubtreeResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


