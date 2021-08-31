# V1TaxonomyMatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Warnings** | Pointer to [**[]V1Warning**](V1Warning.md) |  | [optional] 
**Errors** | Pointer to [**[]V1Error**](V1Error.md) |  | [optional] 
**Query** | Pointer to **[]string** |  | [optional] 
**Taxonomy** | Pointer to [**V1TaxonomyNode**](V1TaxonomyNode.md) |  | [optional] 

## Methods

### NewV1TaxonomyMatch

`func NewV1TaxonomyMatch() *V1TaxonomyMatch`

NewV1TaxonomyMatch instantiates a new V1TaxonomyMatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TaxonomyMatchWithDefaults

`func NewV1TaxonomyMatchWithDefaults() *V1TaxonomyMatch`

NewV1TaxonomyMatchWithDefaults instantiates a new V1TaxonomyMatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWarnings

`func (o *V1TaxonomyMatch) GetWarnings() []V1Warning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V1TaxonomyMatch) GetWarningsOk() (*[]V1Warning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V1TaxonomyMatch) SetWarnings(v []V1Warning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V1TaxonomyMatch) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetErrors

`func (o *V1TaxonomyMatch) GetErrors() []V1Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V1TaxonomyMatch) GetErrorsOk() (*[]V1Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V1TaxonomyMatch) SetErrors(v []V1Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V1TaxonomyMatch) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetQuery

`func (o *V1TaxonomyMatch) GetQuery() []string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *V1TaxonomyMatch) GetQueryOk() (*[]string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *V1TaxonomyMatch) SetQuery(v []string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *V1TaxonomyMatch) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetTaxonomy

`func (o *V1TaxonomyMatch) GetTaxonomy() V1TaxonomyNode`

GetTaxonomy returns the Taxonomy field if non-nil, zero value otherwise.

### GetTaxonomyOk

`func (o *V1TaxonomyMatch) GetTaxonomyOk() (*V1TaxonomyNode, bool)`

GetTaxonomyOk returns a tuple with the Taxonomy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxonomy

`func (o *V1TaxonomyMatch) SetTaxonomy(v V1TaxonomyNode)`

SetTaxonomy sets Taxonomy field to given value.

### HasTaxonomy

`func (o *V1TaxonomyMatch) HasTaxonomy() bool`

HasTaxonomy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


