# V1reportsOrganism

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaxId** | Pointer to **int32** |  | [optional] 
**SciName** | Pointer to **string** |  | [optional] 
**OrganismName** | Pointer to **string** |  | [optional] 
**CommonName** | Pointer to **string** |  | [optional] 
**Lineage** | Pointer to [**[]V1reportsLineageOrganism**](V1reportsLineageOrganism.md) |  | [optional] 
**Strain** | Pointer to **string** |  | [optional] 
**PangolinClassification** | Pointer to **string** |  | [optional] 

## Methods

### NewV1reportsOrganism

`func NewV1reportsOrganism() *V1reportsOrganism`

NewV1reportsOrganism instantiates a new V1reportsOrganism object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1reportsOrganismWithDefaults

`func NewV1reportsOrganismWithDefaults() *V1reportsOrganism`

NewV1reportsOrganismWithDefaults instantiates a new V1reportsOrganism object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaxId

`func (o *V1reportsOrganism) GetTaxId() int32`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *V1reportsOrganism) GetTaxIdOk() (*int32, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *V1reportsOrganism) SetTaxId(v int32)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *V1reportsOrganism) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.

### GetSciName

`func (o *V1reportsOrganism) GetSciName() string`

GetSciName returns the SciName field if non-nil, zero value otherwise.

### GetSciNameOk

`func (o *V1reportsOrganism) GetSciNameOk() (*string, bool)`

GetSciNameOk returns a tuple with the SciName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSciName

`func (o *V1reportsOrganism) SetSciName(v string)`

SetSciName sets SciName field to given value.

### HasSciName

`func (o *V1reportsOrganism) HasSciName() bool`

HasSciName returns a boolean if a field has been set.

### GetOrganismName

`func (o *V1reportsOrganism) GetOrganismName() string`

GetOrganismName returns the OrganismName field if non-nil, zero value otherwise.

### GetOrganismNameOk

`func (o *V1reportsOrganism) GetOrganismNameOk() (*string, bool)`

GetOrganismNameOk returns a tuple with the OrganismName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganismName

`func (o *V1reportsOrganism) SetOrganismName(v string)`

SetOrganismName sets OrganismName field to given value.

### HasOrganismName

`func (o *V1reportsOrganism) HasOrganismName() bool`

HasOrganismName returns a boolean if a field has been set.

### GetCommonName

`func (o *V1reportsOrganism) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *V1reportsOrganism) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *V1reportsOrganism) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *V1reportsOrganism) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### GetLineage

`func (o *V1reportsOrganism) GetLineage() []V1reportsLineageOrganism`

GetLineage returns the Lineage field if non-nil, zero value otherwise.

### GetLineageOk

`func (o *V1reportsOrganism) GetLineageOk() (*[]V1reportsLineageOrganism, bool)`

GetLineageOk returns a tuple with the Lineage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineage

`func (o *V1reportsOrganism) SetLineage(v []V1reportsLineageOrganism)`

SetLineage sets Lineage field to given value.

### HasLineage

`func (o *V1reportsOrganism) HasLineage() bool`

HasLineage returns a boolean if a field has been set.

### GetStrain

`func (o *V1reportsOrganism) GetStrain() string`

GetStrain returns the Strain field if non-nil, zero value otherwise.

### GetStrainOk

`func (o *V1reportsOrganism) GetStrainOk() (*string, bool)`

GetStrainOk returns a tuple with the Strain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrain

`func (o *V1reportsOrganism) SetStrain(v string)`

SetStrain sets Strain field to given value.

### HasStrain

`func (o *V1reportsOrganism) HasStrain() bool`

HasStrain returns a boolean if a field has been set.

### GetPangolinClassification

`func (o *V1reportsOrganism) GetPangolinClassification() string`

GetPangolinClassification returns the PangolinClassification field if non-nil, zero value otherwise.

### GetPangolinClassificationOk

`func (o *V1reportsOrganism) GetPangolinClassificationOk() (*string, bool)`

GetPangolinClassificationOk returns a tuple with the PangolinClassification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPangolinClassification

`func (o *V1reportsOrganism) SetPangolinClassification(v string)`

SetPangolinClassification sets PangolinClassification field to given value.

### HasPangolinClassification

`func (o *V1reportsOrganism) HasPangolinClassification() bool`

HasPangolinClassification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


