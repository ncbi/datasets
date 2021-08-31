# V1OrganismQueryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganismQuery** | Pointer to **string** |  | [optional] 
**TaxonQuery** | Pointer to **string** |  | [optional] 
**TaxRankFilter** | Pointer to [**V1OrganismQueryRequestTaxRankFilter**](V1OrganismQueryRequestTaxRankFilter.md) |  | [optional] [default to V1ORGANISMQUERYREQUESTTAXRANKFILTER_SPECIES]

## Methods

### NewV1OrganismQueryRequest

`func NewV1OrganismQueryRequest() *V1OrganismQueryRequest`

NewV1OrganismQueryRequest instantiates a new V1OrganismQueryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1OrganismQueryRequestWithDefaults

`func NewV1OrganismQueryRequestWithDefaults() *V1OrganismQueryRequest`

NewV1OrganismQueryRequestWithDefaults instantiates a new V1OrganismQueryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganismQuery

`func (o *V1OrganismQueryRequest) GetOrganismQuery() string`

GetOrganismQuery returns the OrganismQuery field if non-nil, zero value otherwise.

### GetOrganismQueryOk

`func (o *V1OrganismQueryRequest) GetOrganismQueryOk() (*string, bool)`

GetOrganismQueryOk returns a tuple with the OrganismQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganismQuery

`func (o *V1OrganismQueryRequest) SetOrganismQuery(v string)`

SetOrganismQuery sets OrganismQuery field to given value.

### HasOrganismQuery

`func (o *V1OrganismQueryRequest) HasOrganismQuery() bool`

HasOrganismQuery returns a boolean if a field has been set.

### GetTaxonQuery

`func (o *V1OrganismQueryRequest) GetTaxonQuery() string`

GetTaxonQuery returns the TaxonQuery field if non-nil, zero value otherwise.

### GetTaxonQueryOk

`func (o *V1OrganismQueryRequest) GetTaxonQueryOk() (*string, bool)`

GetTaxonQueryOk returns a tuple with the TaxonQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxonQuery

`func (o *V1OrganismQueryRequest) SetTaxonQuery(v string)`

SetTaxonQuery sets TaxonQuery field to given value.

### HasTaxonQuery

`func (o *V1OrganismQueryRequest) HasTaxonQuery() bool`

HasTaxonQuery returns a boolean if a field has been set.

### GetTaxRankFilter

`func (o *V1OrganismQueryRequest) GetTaxRankFilter() V1OrganismQueryRequestTaxRankFilter`

GetTaxRankFilter returns the TaxRankFilter field if non-nil, zero value otherwise.

### GetTaxRankFilterOk

`func (o *V1OrganismQueryRequest) GetTaxRankFilterOk() (*V1OrganismQueryRequestTaxRankFilter, bool)`

GetTaxRankFilterOk returns a tuple with the TaxRankFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRankFilter

`func (o *V1OrganismQueryRequest) SetTaxRankFilter(v V1OrganismQueryRequestTaxRankFilter)`

SetTaxRankFilter sets TaxRankFilter field to given value.

### HasTaxRankFilter

`func (o *V1OrganismQueryRequest) HasTaxRankFilter() bool`

HasTaxRankFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


