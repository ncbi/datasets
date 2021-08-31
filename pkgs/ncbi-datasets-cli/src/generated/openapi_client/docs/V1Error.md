# V1Error

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssemblyErrorCode** | Pointer to [**V1ErrorAssemblyErrorCode**](V1ErrorAssemblyErrorCode.md) |  | [optional] [default to V1ERRORASSEMBLYERRORCODE_UNKNOWN_ASSEMBLY_ERROR_CODE]
**GeneErrorCode** | Pointer to [**V1ErrorGeneErrorCode**](V1ErrorGeneErrorCode.md) |  | [optional] [default to V1ERRORGENEERRORCODE_UNKNOWN_GENE_ERROR_CODE]
**VirusErrorCode** | Pointer to [**V1ErrorVirusErrorCode**](V1ErrorVirusErrorCode.md) |  | [optional] [default to V1ERRORVIRUSERRORCODE_UNKNOWN_VIRUS_ERROR_CODE]
**Reason** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**ValidIdentifiers** | Pointer to **[]string** |  | [optional] 
**InvalidIdentifiers** | Pointer to **[]string** |  | [optional] 

## Methods

### NewV1Error

`func NewV1Error() *V1Error`

NewV1Error instantiates a new V1Error object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ErrorWithDefaults

`func NewV1ErrorWithDefaults() *V1Error`

NewV1ErrorWithDefaults instantiates a new V1Error object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssemblyErrorCode

`func (o *V1Error) GetAssemblyErrorCode() V1ErrorAssemblyErrorCode`

GetAssemblyErrorCode returns the AssemblyErrorCode field if non-nil, zero value otherwise.

### GetAssemblyErrorCodeOk

`func (o *V1Error) GetAssemblyErrorCodeOk() (*V1ErrorAssemblyErrorCode, bool)`

GetAssemblyErrorCodeOk returns a tuple with the AssemblyErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssemblyErrorCode

`func (o *V1Error) SetAssemblyErrorCode(v V1ErrorAssemblyErrorCode)`

SetAssemblyErrorCode sets AssemblyErrorCode field to given value.

### HasAssemblyErrorCode

`func (o *V1Error) HasAssemblyErrorCode() bool`

HasAssemblyErrorCode returns a boolean if a field has been set.

### GetGeneErrorCode

`func (o *V1Error) GetGeneErrorCode() V1ErrorGeneErrorCode`

GetGeneErrorCode returns the GeneErrorCode field if non-nil, zero value otherwise.

### GetGeneErrorCodeOk

`func (o *V1Error) GetGeneErrorCodeOk() (*V1ErrorGeneErrorCode, bool)`

GetGeneErrorCodeOk returns a tuple with the GeneErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneErrorCode

`func (o *V1Error) SetGeneErrorCode(v V1ErrorGeneErrorCode)`

SetGeneErrorCode sets GeneErrorCode field to given value.

### HasGeneErrorCode

`func (o *V1Error) HasGeneErrorCode() bool`

HasGeneErrorCode returns a boolean if a field has been set.

### GetVirusErrorCode

`func (o *V1Error) GetVirusErrorCode() V1ErrorVirusErrorCode`

GetVirusErrorCode returns the VirusErrorCode field if non-nil, zero value otherwise.

### GetVirusErrorCodeOk

`func (o *V1Error) GetVirusErrorCodeOk() (*V1ErrorVirusErrorCode, bool)`

GetVirusErrorCodeOk returns a tuple with the VirusErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirusErrorCode

`func (o *V1Error) SetVirusErrorCode(v V1ErrorVirusErrorCode)`

SetVirusErrorCode sets VirusErrorCode field to given value.

### HasVirusErrorCode

`func (o *V1Error) HasVirusErrorCode() bool`

HasVirusErrorCode returns a boolean if a field has been set.

### GetReason

`func (o *V1Error) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *V1Error) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *V1Error) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *V1Error) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetMessage

`func (o *V1Error) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V1Error) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V1Error) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *V1Error) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetValidIdentifiers

`func (o *V1Error) GetValidIdentifiers() []string`

GetValidIdentifiers returns the ValidIdentifiers field if non-nil, zero value otherwise.

### GetValidIdentifiersOk

`func (o *V1Error) GetValidIdentifiersOk() (*[]string, bool)`

GetValidIdentifiersOk returns a tuple with the ValidIdentifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidIdentifiers

`func (o *V1Error) SetValidIdentifiers(v []string)`

SetValidIdentifiers sets ValidIdentifiers field to given value.

### HasValidIdentifiers

`func (o *V1Error) HasValidIdentifiers() bool`

HasValidIdentifiers returns a boolean if a field has been set.

### GetInvalidIdentifiers

`func (o *V1Error) GetInvalidIdentifiers() []string`

GetInvalidIdentifiers returns the InvalidIdentifiers field if non-nil, zero value otherwise.

### GetInvalidIdentifiersOk

`func (o *V1Error) GetInvalidIdentifiersOk() (*[]string, bool)`

GetInvalidIdentifiersOk returns a tuple with the InvalidIdentifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidIdentifiers

`func (o *V1Error) SetInvalidIdentifiers(v []string)`

SetInvalidIdentifiers sets InvalidIdentifiers field to given value.

### HasInvalidIdentifiers

`func (o *V1Error) HasInvalidIdentifiers() bool`

HasInvalidIdentifiers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


