# V1GeneMatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | Pointer to [**[]V1Message**](V1Message.md) |  | [optional] 
**Warnings** | Pointer to [**[]V1Warning**](V1Warning.md) |  | [optional] 
**Query** | Pointer to **[]string** |  | [optional] 
**Gene** | Pointer to [**V1GeneDescriptor**](V1GeneDescriptor.md) |  | [optional] 
**Errors** | Pointer to [**[]V1Error**](V1Error.md) |  | [optional] 

## Methods

### NewV1GeneMatch

`func NewV1GeneMatch() *V1GeneMatch`

NewV1GeneMatch instantiates a new V1GeneMatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1GeneMatchWithDefaults

`func NewV1GeneMatchWithDefaults() *V1GeneMatch`

NewV1GeneMatchWithDefaults instantiates a new V1GeneMatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *V1GeneMatch) GetMessages() []V1Message`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *V1GeneMatch) GetMessagesOk() (*[]V1Message, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *V1GeneMatch) SetMessages(v []V1Message)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *V1GeneMatch) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetWarnings

`func (o *V1GeneMatch) GetWarnings() []V1Warning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *V1GeneMatch) GetWarningsOk() (*[]V1Warning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *V1GeneMatch) SetWarnings(v []V1Warning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *V1GeneMatch) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetQuery

`func (o *V1GeneMatch) GetQuery() []string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *V1GeneMatch) GetQueryOk() (*[]string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *V1GeneMatch) SetQuery(v []string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *V1GeneMatch) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetGene

`func (o *V1GeneMatch) GetGene() V1GeneDescriptor`

GetGene returns the Gene field if non-nil, zero value otherwise.

### GetGeneOk

`func (o *V1GeneMatch) GetGeneOk() (*V1GeneDescriptor, bool)`

GetGeneOk returns a tuple with the Gene field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGene

`func (o *V1GeneMatch) SetGene(v V1GeneDescriptor)`

SetGene sets Gene field to given value.

### HasGene

`func (o *V1GeneMatch) HasGene() bool`

HasGene returns a boolean if a field has been set.

### GetErrors

`func (o *V1GeneMatch) GetErrors() []V1Error`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *V1GeneMatch) GetErrorsOk() (*[]V1Error, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *V1GeneMatch) SetErrors(v []V1Error)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *V1GeneMatch) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


