# V1ProkaryoteGeneRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accessions** | Pointer to **[]string** |  | [optional] 
**IncludeAnnotationType** | Pointer to [**[]V1Fasta**](V1Fasta.md) |  | [optional] 
**GeneFlankConfig** | Pointer to [**V1ProkaryoteGeneRequestGeneFlankConfig**](V1ProkaryoteGeneRequestGeneFlankConfig.md) |  | [optional] 
**Taxon** | Pointer to **string** |  | [optional] 

## Methods

### NewV1ProkaryoteGeneRequest

`func NewV1ProkaryoteGeneRequest() *V1ProkaryoteGeneRequest`

NewV1ProkaryoteGeneRequest instantiates a new V1ProkaryoteGeneRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ProkaryoteGeneRequestWithDefaults

`func NewV1ProkaryoteGeneRequestWithDefaults() *V1ProkaryoteGeneRequest`

NewV1ProkaryoteGeneRequestWithDefaults instantiates a new V1ProkaryoteGeneRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessions

`func (o *V1ProkaryoteGeneRequest) GetAccessions() []string`

GetAccessions returns the Accessions field if non-nil, zero value otherwise.

### GetAccessionsOk

`func (o *V1ProkaryoteGeneRequest) GetAccessionsOk() (*[]string, bool)`

GetAccessionsOk returns a tuple with the Accessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessions

`func (o *V1ProkaryoteGeneRequest) SetAccessions(v []string)`

SetAccessions sets Accessions field to given value.

### HasAccessions

`func (o *V1ProkaryoteGeneRequest) HasAccessions() bool`

HasAccessions returns a boolean if a field has been set.

### GetIncludeAnnotationType

`func (o *V1ProkaryoteGeneRequest) GetIncludeAnnotationType() []V1Fasta`

GetIncludeAnnotationType returns the IncludeAnnotationType field if non-nil, zero value otherwise.

### GetIncludeAnnotationTypeOk

`func (o *V1ProkaryoteGeneRequest) GetIncludeAnnotationTypeOk() (*[]V1Fasta, bool)`

GetIncludeAnnotationTypeOk returns a tuple with the IncludeAnnotationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAnnotationType

`func (o *V1ProkaryoteGeneRequest) SetIncludeAnnotationType(v []V1Fasta)`

SetIncludeAnnotationType sets IncludeAnnotationType field to given value.

### HasIncludeAnnotationType

`func (o *V1ProkaryoteGeneRequest) HasIncludeAnnotationType() bool`

HasIncludeAnnotationType returns a boolean if a field has been set.

### GetGeneFlankConfig

`func (o *V1ProkaryoteGeneRequest) GetGeneFlankConfig() V1ProkaryoteGeneRequestGeneFlankConfig`

GetGeneFlankConfig returns the GeneFlankConfig field if non-nil, zero value otherwise.

### GetGeneFlankConfigOk

`func (o *V1ProkaryoteGeneRequest) GetGeneFlankConfigOk() (*V1ProkaryoteGeneRequestGeneFlankConfig, bool)`

GetGeneFlankConfigOk returns a tuple with the GeneFlankConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneFlankConfig

`func (o *V1ProkaryoteGeneRequest) SetGeneFlankConfig(v V1ProkaryoteGeneRequestGeneFlankConfig)`

SetGeneFlankConfig sets GeneFlankConfig field to given value.

### HasGeneFlankConfig

`func (o *V1ProkaryoteGeneRequest) HasGeneFlankConfig() bool`

HasGeneFlankConfig returns a boolean if a field has been set.

### GetTaxon

`func (o *V1ProkaryoteGeneRequest) GetTaxon() string`

GetTaxon returns the Taxon field if non-nil, zero value otherwise.

### GetTaxonOk

`func (o *V1ProkaryoteGeneRequest) GetTaxonOk() (*string, bool)`

GetTaxonOk returns a tuple with the Taxon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxon

`func (o *V1ProkaryoteGeneRequest) SetTaxon(v string)`

SetTaxon sets Taxon field to given value.

### HasTaxon

`func (o *V1ProkaryoteGeneRequest) HasTaxon() bool`

HasTaxon returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


