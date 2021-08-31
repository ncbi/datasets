# V1GeneDatasetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GeneIds** | Pointer to **[]int32** |  | [optional] 
**Accessions** | Pointer to **[]string** |  | [optional] 
**SymbolsForTaxon** | Pointer to [**V1GeneDatasetRequestSymbolsForTaxon**](V1GeneDatasetRequestSymbolsForTaxon.md) |  | [optional] 
**Taxon** | Pointer to **string** |  | [optional] 
**IncludeAnnotationType** | Pointer to [**[]V1Fasta**](V1Fasta.md) |  | [optional] 
**ReturnedContent** | Pointer to [**V1GeneDatasetRequestContentType**](V1GeneDatasetRequestContentType.md) |  | [optional] [default to V1GENEDATASETREQUESTCONTENTTYPE_COMPLETE]
**SortSchema** | Pointer to [**V1GeneDatasetRequestSort**](V1GeneDatasetRequestSort.md) |  | [optional] 
**Limit** | Pointer to **string** |  | [optional] 
**FastaFilter** | Pointer to **[]string** |  | [optional] 

## Methods

### NewV1GeneDatasetRequest

`func NewV1GeneDatasetRequest() *V1GeneDatasetRequest`

NewV1GeneDatasetRequest instantiates a new V1GeneDatasetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1GeneDatasetRequestWithDefaults

`func NewV1GeneDatasetRequestWithDefaults() *V1GeneDatasetRequest`

NewV1GeneDatasetRequestWithDefaults instantiates a new V1GeneDatasetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeneIds

`func (o *V1GeneDatasetRequest) GetGeneIds() []int32`

GetGeneIds returns the GeneIds field if non-nil, zero value otherwise.

### GetGeneIdsOk

`func (o *V1GeneDatasetRequest) GetGeneIdsOk() (*[]int32, bool)`

GetGeneIdsOk returns a tuple with the GeneIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneIds

`func (o *V1GeneDatasetRequest) SetGeneIds(v []int32)`

SetGeneIds sets GeneIds field to given value.

### HasGeneIds

`func (o *V1GeneDatasetRequest) HasGeneIds() bool`

HasGeneIds returns a boolean if a field has been set.

### GetAccessions

`func (o *V1GeneDatasetRequest) GetAccessions() []string`

GetAccessions returns the Accessions field if non-nil, zero value otherwise.

### GetAccessionsOk

`func (o *V1GeneDatasetRequest) GetAccessionsOk() (*[]string, bool)`

GetAccessionsOk returns a tuple with the Accessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessions

`func (o *V1GeneDatasetRequest) SetAccessions(v []string)`

SetAccessions sets Accessions field to given value.

### HasAccessions

`func (o *V1GeneDatasetRequest) HasAccessions() bool`

HasAccessions returns a boolean if a field has been set.

### GetSymbolsForTaxon

`func (o *V1GeneDatasetRequest) GetSymbolsForTaxon() V1GeneDatasetRequestSymbolsForTaxon`

GetSymbolsForTaxon returns the SymbolsForTaxon field if non-nil, zero value otherwise.

### GetSymbolsForTaxonOk

`func (o *V1GeneDatasetRequest) GetSymbolsForTaxonOk() (*V1GeneDatasetRequestSymbolsForTaxon, bool)`

GetSymbolsForTaxonOk returns a tuple with the SymbolsForTaxon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbolsForTaxon

`func (o *V1GeneDatasetRequest) SetSymbolsForTaxon(v V1GeneDatasetRequestSymbolsForTaxon)`

SetSymbolsForTaxon sets SymbolsForTaxon field to given value.

### HasSymbolsForTaxon

`func (o *V1GeneDatasetRequest) HasSymbolsForTaxon() bool`

HasSymbolsForTaxon returns a boolean if a field has been set.

### GetTaxon

`func (o *V1GeneDatasetRequest) GetTaxon() string`

GetTaxon returns the Taxon field if non-nil, zero value otherwise.

### GetTaxonOk

`func (o *V1GeneDatasetRequest) GetTaxonOk() (*string, bool)`

GetTaxonOk returns a tuple with the Taxon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxon

`func (o *V1GeneDatasetRequest) SetTaxon(v string)`

SetTaxon sets Taxon field to given value.

### HasTaxon

`func (o *V1GeneDatasetRequest) HasTaxon() bool`

HasTaxon returns a boolean if a field has been set.

### GetIncludeAnnotationType

`func (o *V1GeneDatasetRequest) GetIncludeAnnotationType() []V1Fasta`

GetIncludeAnnotationType returns the IncludeAnnotationType field if non-nil, zero value otherwise.

### GetIncludeAnnotationTypeOk

`func (o *V1GeneDatasetRequest) GetIncludeAnnotationTypeOk() (*[]V1Fasta, bool)`

GetIncludeAnnotationTypeOk returns a tuple with the IncludeAnnotationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAnnotationType

`func (o *V1GeneDatasetRequest) SetIncludeAnnotationType(v []V1Fasta)`

SetIncludeAnnotationType sets IncludeAnnotationType field to given value.

### HasIncludeAnnotationType

`func (o *V1GeneDatasetRequest) HasIncludeAnnotationType() bool`

HasIncludeAnnotationType returns a boolean if a field has been set.

### GetReturnedContent

`func (o *V1GeneDatasetRequest) GetReturnedContent() V1GeneDatasetRequestContentType`

GetReturnedContent returns the ReturnedContent field if non-nil, zero value otherwise.

### GetReturnedContentOk

`func (o *V1GeneDatasetRequest) GetReturnedContentOk() (*V1GeneDatasetRequestContentType, bool)`

GetReturnedContentOk returns a tuple with the ReturnedContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnedContent

`func (o *V1GeneDatasetRequest) SetReturnedContent(v V1GeneDatasetRequestContentType)`

SetReturnedContent sets ReturnedContent field to given value.

### HasReturnedContent

`func (o *V1GeneDatasetRequest) HasReturnedContent() bool`

HasReturnedContent returns a boolean if a field has been set.

### GetSortSchema

`func (o *V1GeneDatasetRequest) GetSortSchema() V1GeneDatasetRequestSort`

GetSortSchema returns the SortSchema field if non-nil, zero value otherwise.

### GetSortSchemaOk

`func (o *V1GeneDatasetRequest) GetSortSchemaOk() (*V1GeneDatasetRequestSort, bool)`

GetSortSchemaOk returns a tuple with the SortSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortSchema

`func (o *V1GeneDatasetRequest) SetSortSchema(v V1GeneDatasetRequestSort)`

SetSortSchema sets SortSchema field to given value.

### HasSortSchema

`func (o *V1GeneDatasetRequest) HasSortSchema() bool`

HasSortSchema returns a boolean if a field has been set.

### GetLimit

`func (o *V1GeneDatasetRequest) GetLimit() string`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *V1GeneDatasetRequest) GetLimitOk() (*string, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *V1GeneDatasetRequest) SetLimit(v string)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *V1GeneDatasetRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetFastaFilter

`func (o *V1GeneDatasetRequest) GetFastaFilter() []string`

GetFastaFilter returns the FastaFilter field if non-nil, zero value otherwise.

### GetFastaFilterOk

`func (o *V1GeneDatasetRequest) GetFastaFilterOk() (*[]string, bool)`

GetFastaFilterOk returns a tuple with the FastaFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFastaFilter

`func (o *V1GeneDatasetRequest) SetFastaFilter(v []string)`

SetFastaFilter sets FastaFilter field to given value.

### HasFastaFilter

`func (o *V1GeneDatasetRequest) HasFastaFilter() bool`

HasFastaFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


