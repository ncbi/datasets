# V1Warning

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GeneWarningCode** | Pointer to [**V1WarningGeneWarningCode**](V1WarningGeneWarningCode.md) |  | [optional] [default to V1WARNINGGENEWARNINGCODE_UNKNOWN_GENE_WARNING_CODE]
**Reason** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**ReplacedId** | Pointer to [**V1WarningReplacedId**](V1WarningReplacedId.md) |  | [optional] 
**UnrecognizedIdentifier** | Pointer to **string** |  | [optional] 

## Methods

### NewV1Warning

`func NewV1Warning() *V1Warning`

NewV1Warning instantiates a new V1Warning object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1WarningWithDefaults

`func NewV1WarningWithDefaults() *V1Warning`

NewV1WarningWithDefaults instantiates a new V1Warning object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeneWarningCode

`func (o *V1Warning) GetGeneWarningCode() V1WarningGeneWarningCode`

GetGeneWarningCode returns the GeneWarningCode field if non-nil, zero value otherwise.

### GetGeneWarningCodeOk

`func (o *V1Warning) GetGeneWarningCodeOk() (*V1WarningGeneWarningCode, bool)`

GetGeneWarningCodeOk returns a tuple with the GeneWarningCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneWarningCode

`func (o *V1Warning) SetGeneWarningCode(v V1WarningGeneWarningCode)`

SetGeneWarningCode sets GeneWarningCode field to given value.

### HasGeneWarningCode

`func (o *V1Warning) HasGeneWarningCode() bool`

HasGeneWarningCode returns a boolean if a field has been set.

### GetReason

`func (o *V1Warning) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *V1Warning) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *V1Warning) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *V1Warning) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetMessage

`func (o *V1Warning) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V1Warning) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V1Warning) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *V1Warning) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetReplacedId

`func (o *V1Warning) GetReplacedId() V1WarningReplacedId`

GetReplacedId returns the ReplacedId field if non-nil, zero value otherwise.

### GetReplacedIdOk

`func (o *V1Warning) GetReplacedIdOk() (*V1WarningReplacedId, bool)`

GetReplacedIdOk returns a tuple with the ReplacedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacedId

`func (o *V1Warning) SetReplacedId(v V1WarningReplacedId)`

SetReplacedId sets ReplacedId field to given value.

### HasReplacedId

`func (o *V1Warning) HasReplacedId() bool`

HasReplacedId returns a boolean if a field has been set.

### GetUnrecognizedIdentifier

`func (o *V1Warning) GetUnrecognizedIdentifier() string`

GetUnrecognizedIdentifier returns the UnrecognizedIdentifier field if non-nil, zero value otherwise.

### GetUnrecognizedIdentifierOk

`func (o *V1Warning) GetUnrecognizedIdentifierOk() (*string, bool)`

GetUnrecognizedIdentifierOk returns a tuple with the UnrecognizedIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrecognizedIdentifier

`func (o *V1Warning) SetUnrecognizedIdentifier(v string)`

SetUnrecognizedIdentifier sets UnrecognizedIdentifier field to given value.

### HasUnrecognizedIdentifier

`func (o *V1Warning) HasUnrecognizedIdentifier() bool`

HasUnrecognizedIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


