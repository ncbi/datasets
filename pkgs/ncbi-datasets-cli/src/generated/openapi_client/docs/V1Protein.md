# V1Protein

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessionVersion** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Length** | Pointer to **int32** |  | [optional] 
**IsoformName** | Pointer to **string** |  | [optional] 
**EnsemblProtein** | Pointer to **string** |  | [optional] 
**MaturePeptides** | Pointer to [**[]V1MaturePeptide**](V1MaturePeptide.md) |  | [optional] 

## Methods

### NewV1Protein

`func NewV1Protein() *V1Protein`

NewV1Protein instantiates a new V1Protein object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ProteinWithDefaults

`func NewV1ProteinWithDefaults() *V1Protein`

NewV1ProteinWithDefaults instantiates a new V1Protein object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessionVersion

`func (o *V1Protein) GetAccessionVersion() string`

GetAccessionVersion returns the AccessionVersion field if non-nil, zero value otherwise.

### GetAccessionVersionOk

`func (o *V1Protein) GetAccessionVersionOk() (*string, bool)`

GetAccessionVersionOk returns a tuple with the AccessionVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessionVersion

`func (o *V1Protein) SetAccessionVersion(v string)`

SetAccessionVersion sets AccessionVersion field to given value.

### HasAccessionVersion

`func (o *V1Protein) HasAccessionVersion() bool`

HasAccessionVersion returns a boolean if a field has been set.

### GetName

`func (o *V1Protein) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1Protein) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1Protein) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1Protein) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLength

`func (o *V1Protein) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *V1Protein) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *V1Protein) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *V1Protein) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetIsoformName

`func (o *V1Protein) GetIsoformName() string`

GetIsoformName returns the IsoformName field if non-nil, zero value otherwise.

### GetIsoformNameOk

`func (o *V1Protein) GetIsoformNameOk() (*string, bool)`

GetIsoformNameOk returns a tuple with the IsoformName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoformName

`func (o *V1Protein) SetIsoformName(v string)`

SetIsoformName sets IsoformName field to given value.

### HasIsoformName

`func (o *V1Protein) HasIsoformName() bool`

HasIsoformName returns a boolean if a field has been set.

### GetEnsemblProtein

`func (o *V1Protein) GetEnsemblProtein() string`

GetEnsemblProtein returns the EnsemblProtein field if non-nil, zero value otherwise.

### GetEnsemblProteinOk

`func (o *V1Protein) GetEnsemblProteinOk() (*string, bool)`

GetEnsemblProteinOk returns a tuple with the EnsemblProtein field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnsemblProtein

`func (o *V1Protein) SetEnsemblProtein(v string)`

SetEnsemblProtein sets EnsemblProtein field to given value.

### HasEnsemblProtein

`func (o *V1Protein) HasEnsemblProtein() bool`

HasEnsemblProtein returns a boolean if a field has been set.

### GetMaturePeptides

`func (o *V1Protein) GetMaturePeptides() []V1MaturePeptide`

GetMaturePeptides returns the MaturePeptides field if non-nil, zero value otherwise.

### GetMaturePeptidesOk

`func (o *V1Protein) GetMaturePeptidesOk() (*[]V1MaturePeptide, bool)`

GetMaturePeptidesOk returns a tuple with the MaturePeptides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaturePeptides

`func (o *V1Protein) SetMaturePeptides(v []V1MaturePeptide)`

SetMaturePeptides sets MaturePeptides field to given value.

### HasMaturePeptides

`func (o *V1Protein) HasMaturePeptides() bool`

HasMaturePeptides returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


