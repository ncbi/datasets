# V1TaxonomyFilteredSubtreeResponseEdge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VisibleChildren** | Pointer to **[]int32** |  | [optional] 
**ChildrenStatus** | Pointer to [**V1TaxonomyFilteredSubtreeResponseEdgeChildStatus**](V1TaxonomyFilteredSubtreeResponseEdgeChildStatus.md) |  | [optional] [default to V1TAXONOMYFILTEREDSUBTREERESPONSEEDGECHILDSTATUS_UNSPECIFIED]

## Methods

### NewV1TaxonomyFilteredSubtreeResponseEdge

`func NewV1TaxonomyFilteredSubtreeResponseEdge() *V1TaxonomyFilteredSubtreeResponseEdge`

NewV1TaxonomyFilteredSubtreeResponseEdge instantiates a new V1TaxonomyFilteredSubtreeResponseEdge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TaxonomyFilteredSubtreeResponseEdgeWithDefaults

`func NewV1TaxonomyFilteredSubtreeResponseEdgeWithDefaults() *V1TaxonomyFilteredSubtreeResponseEdge`

NewV1TaxonomyFilteredSubtreeResponseEdgeWithDefaults instantiates a new V1TaxonomyFilteredSubtreeResponseEdge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVisibleChildren

`func (o *V1TaxonomyFilteredSubtreeResponseEdge) GetVisibleChildren() []int32`

GetVisibleChildren returns the VisibleChildren field if non-nil, zero value otherwise.

### GetVisibleChildrenOk

`func (o *V1TaxonomyFilteredSubtreeResponseEdge) GetVisibleChildrenOk() (*[]int32, bool)`

GetVisibleChildrenOk returns a tuple with the VisibleChildren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleChildren

`func (o *V1TaxonomyFilteredSubtreeResponseEdge) SetVisibleChildren(v []int32)`

SetVisibleChildren sets VisibleChildren field to given value.

### HasVisibleChildren

`func (o *V1TaxonomyFilteredSubtreeResponseEdge) HasVisibleChildren() bool`

HasVisibleChildren returns a boolean if a field has been set.

### GetChildrenStatus

`func (o *V1TaxonomyFilteredSubtreeResponseEdge) GetChildrenStatus() V1TaxonomyFilteredSubtreeResponseEdgeChildStatus`

GetChildrenStatus returns the ChildrenStatus field if non-nil, zero value otherwise.

### GetChildrenStatusOk

`func (o *V1TaxonomyFilteredSubtreeResponseEdge) GetChildrenStatusOk() (*V1TaxonomyFilteredSubtreeResponseEdgeChildStatus, bool)`

GetChildrenStatusOk returns a tuple with the ChildrenStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildrenStatus

`func (o *V1TaxonomyFilteredSubtreeResponseEdge) SetChildrenStatus(v V1TaxonomyFilteredSubtreeResponseEdgeChildStatus)`

SetChildrenStatus sets ChildrenStatus field to given value.

### HasChildrenStatus

`func (o *V1TaxonomyFilteredSubtreeResponseEdge) HasChildrenStatus() bool`

HasChildrenStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


