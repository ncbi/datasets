# V1reportsProkaryoteGene

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accession** | Pointer to **string** |  | [optional] 
**GeneSymbol** | Pointer to **string** |  | [optional] 
**ProteinName** | Pointer to **string** |  | [optional] 
**ProteinLength** | Pointer to **int32** |  | [optional] 
**TaxonomyScope** | Pointer to [**V1reportsOrganism**](V1reportsOrganism.md) |  | [optional] 
**NumberOfGenomeMappings** | Pointer to **int32** |  | [optional] 
**ProteinNameEvidence** | Pointer to [**V1reportsProkaryoteGeneProteinNameEvidence**](V1reportsProkaryoteGeneProteinNameEvidence.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**EcNumber** | Pointer to **[]string** |  | [optional] 

## Methods

### NewV1reportsProkaryoteGene

`func NewV1reportsProkaryoteGene() *V1reportsProkaryoteGene`

NewV1reportsProkaryoteGene instantiates a new V1reportsProkaryoteGene object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1reportsProkaryoteGeneWithDefaults

`func NewV1reportsProkaryoteGeneWithDefaults() *V1reportsProkaryoteGene`

NewV1reportsProkaryoteGeneWithDefaults instantiates a new V1reportsProkaryoteGene object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccession

`func (o *V1reportsProkaryoteGene) GetAccession() string`

GetAccession returns the Accession field if non-nil, zero value otherwise.

### GetAccessionOk

`func (o *V1reportsProkaryoteGene) GetAccessionOk() (*string, bool)`

GetAccessionOk returns a tuple with the Accession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccession

`func (o *V1reportsProkaryoteGene) SetAccession(v string)`

SetAccession sets Accession field to given value.

### HasAccession

`func (o *V1reportsProkaryoteGene) HasAccession() bool`

HasAccession returns a boolean if a field has been set.

### GetGeneSymbol

`func (o *V1reportsProkaryoteGene) GetGeneSymbol() string`

GetGeneSymbol returns the GeneSymbol field if non-nil, zero value otherwise.

### GetGeneSymbolOk

`func (o *V1reportsProkaryoteGene) GetGeneSymbolOk() (*string, bool)`

GetGeneSymbolOk returns a tuple with the GeneSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneSymbol

`func (o *V1reportsProkaryoteGene) SetGeneSymbol(v string)`

SetGeneSymbol sets GeneSymbol field to given value.

### HasGeneSymbol

`func (o *V1reportsProkaryoteGene) HasGeneSymbol() bool`

HasGeneSymbol returns a boolean if a field has been set.

### GetProteinName

`func (o *V1reportsProkaryoteGene) GetProteinName() string`

GetProteinName returns the ProteinName field if non-nil, zero value otherwise.

### GetProteinNameOk

`func (o *V1reportsProkaryoteGene) GetProteinNameOk() (*string, bool)`

GetProteinNameOk returns a tuple with the ProteinName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProteinName

`func (o *V1reportsProkaryoteGene) SetProteinName(v string)`

SetProteinName sets ProteinName field to given value.

### HasProteinName

`func (o *V1reportsProkaryoteGene) HasProteinName() bool`

HasProteinName returns a boolean if a field has been set.

### GetProteinLength

`func (o *V1reportsProkaryoteGene) GetProteinLength() int32`

GetProteinLength returns the ProteinLength field if non-nil, zero value otherwise.

### GetProteinLengthOk

`func (o *V1reportsProkaryoteGene) GetProteinLengthOk() (*int32, bool)`

GetProteinLengthOk returns a tuple with the ProteinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProteinLength

`func (o *V1reportsProkaryoteGene) SetProteinLength(v int32)`

SetProteinLength sets ProteinLength field to given value.

### HasProteinLength

`func (o *V1reportsProkaryoteGene) HasProteinLength() bool`

HasProteinLength returns a boolean if a field has been set.

### GetTaxonomyScope

`func (o *V1reportsProkaryoteGene) GetTaxonomyScope() V1reportsOrganism`

GetTaxonomyScope returns the TaxonomyScope field if non-nil, zero value otherwise.

### GetTaxonomyScopeOk

`func (o *V1reportsProkaryoteGene) GetTaxonomyScopeOk() (*V1reportsOrganism, bool)`

GetTaxonomyScopeOk returns a tuple with the TaxonomyScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxonomyScope

`func (o *V1reportsProkaryoteGene) SetTaxonomyScope(v V1reportsOrganism)`

SetTaxonomyScope sets TaxonomyScope field to given value.

### HasTaxonomyScope

`func (o *V1reportsProkaryoteGene) HasTaxonomyScope() bool`

HasTaxonomyScope returns a boolean if a field has been set.

### GetNumberOfGenomeMappings

`func (o *V1reportsProkaryoteGene) GetNumberOfGenomeMappings() int32`

GetNumberOfGenomeMappings returns the NumberOfGenomeMappings field if non-nil, zero value otherwise.

### GetNumberOfGenomeMappingsOk

`func (o *V1reportsProkaryoteGene) GetNumberOfGenomeMappingsOk() (*int32, bool)`

GetNumberOfGenomeMappingsOk returns a tuple with the NumberOfGenomeMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfGenomeMappings

`func (o *V1reportsProkaryoteGene) SetNumberOfGenomeMappings(v int32)`

SetNumberOfGenomeMappings sets NumberOfGenomeMappings field to given value.

### HasNumberOfGenomeMappings

`func (o *V1reportsProkaryoteGene) HasNumberOfGenomeMappings() bool`

HasNumberOfGenomeMappings returns a boolean if a field has been set.

### GetProteinNameEvidence

`func (o *V1reportsProkaryoteGene) GetProteinNameEvidence() V1reportsProkaryoteGeneProteinNameEvidence`

GetProteinNameEvidence returns the ProteinNameEvidence field if non-nil, zero value otherwise.

### GetProteinNameEvidenceOk

`func (o *V1reportsProkaryoteGene) GetProteinNameEvidenceOk() (*V1reportsProkaryoteGeneProteinNameEvidence, bool)`

GetProteinNameEvidenceOk returns a tuple with the ProteinNameEvidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProteinNameEvidence

`func (o *V1reportsProkaryoteGene) SetProteinNameEvidence(v V1reportsProkaryoteGeneProteinNameEvidence)`

SetProteinNameEvidence sets ProteinNameEvidence field to given value.

### HasProteinNameEvidence

`func (o *V1reportsProkaryoteGene) HasProteinNameEvidence() bool`

HasProteinNameEvidence returns a boolean if a field has been set.

### GetDescription

`func (o *V1reportsProkaryoteGene) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V1reportsProkaryoteGene) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V1reportsProkaryoteGene) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V1reportsProkaryoteGene) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEcNumber

`func (o *V1reportsProkaryoteGene) GetEcNumber() []string`

GetEcNumber returns the EcNumber field if non-nil, zero value otherwise.

### GetEcNumberOk

`func (o *V1reportsProkaryoteGene) GetEcNumberOk() (*[]string, bool)`

GetEcNumberOk returns a tuple with the EcNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcNumber

`func (o *V1reportsProkaryoteGene) SetEcNumber(v []string)`

SetEcNumber sets EcNumber field to given value.

### HasEcNumber

`func (o *V1reportsProkaryoteGene) HasEcNumber() bool`

HasEcNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


