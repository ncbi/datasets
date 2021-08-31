# V1TaxTreeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaxId** | Pointer to **string** |  | [optional] 
**TaxToken** | Pointer to **string** |  | [optional] 
**Taxon** | Pointer to **string** |  | [optional] 
**ChildrenOnly** | Pointer to **bool** |  | [optional] 

## Methods

### NewV1TaxTreeRequest

`func NewV1TaxTreeRequest() *V1TaxTreeRequest`

NewV1TaxTreeRequest instantiates a new V1TaxTreeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TaxTreeRequestWithDefaults

`func NewV1TaxTreeRequestWithDefaults() *V1TaxTreeRequest`

NewV1TaxTreeRequestWithDefaults instantiates a new V1TaxTreeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaxId

`func (o *V1TaxTreeRequest) GetTaxId() string`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *V1TaxTreeRequest) GetTaxIdOk() (*string, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *V1TaxTreeRequest) SetTaxId(v string)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *V1TaxTreeRequest) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.

### GetTaxToken

`func (o *V1TaxTreeRequest) GetTaxToken() string`

GetTaxToken returns the TaxToken field if non-nil, zero value otherwise.

### GetTaxTokenOk

`func (o *V1TaxTreeRequest) GetTaxTokenOk() (*string, bool)`

GetTaxTokenOk returns a tuple with the TaxToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxToken

`func (o *V1TaxTreeRequest) SetTaxToken(v string)`

SetTaxToken sets TaxToken field to given value.

### HasTaxToken

`func (o *V1TaxTreeRequest) HasTaxToken() bool`

HasTaxToken returns a boolean if a field has been set.

### GetTaxon

`func (o *V1TaxTreeRequest) GetTaxon() string`

GetTaxon returns the Taxon field if non-nil, zero value otherwise.

### GetTaxonOk

`func (o *V1TaxTreeRequest) GetTaxonOk() (*string, bool)`

GetTaxonOk returns a tuple with the Taxon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxon

`func (o *V1TaxTreeRequest) SetTaxon(v string)`

SetTaxon sets Taxon field to given value.

### HasTaxon

`func (o *V1TaxTreeRequest) HasTaxon() bool`

HasTaxon returns a boolean if a field has been set.

### GetChildrenOnly

`func (o *V1TaxTreeRequest) GetChildrenOnly() bool`

GetChildrenOnly returns the ChildrenOnly field if non-nil, zero value otherwise.

### GetChildrenOnlyOk

`func (o *V1TaxTreeRequest) GetChildrenOnlyOk() (*bool, bool)`

GetChildrenOnlyOk returns a tuple with the ChildrenOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildrenOnly

`func (o *V1TaxTreeRequest) SetChildrenOnly(v bool)`

SetChildrenOnly sets ChildrenOnly field to given value.

### HasChildrenOnly

`func (o *V1TaxTreeRequest) HasChildrenOnly() bool`

HasChildrenOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


