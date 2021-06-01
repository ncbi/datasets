# V1alpha1GeneDatasetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accessions** | **[]string** | RNA or Protein accessions. | [optional] 
**FastaFilter** | **[]string** |  | [optional] 
**GeneIds** | **[]int64** |  | [optional] 
**IncludeAnnotationType** | [**[]V1alpha1Fasta**](v1alpha1Fasta.md) | Select additional types of annotation to include in the data package.  If unset, no annotation is provided. | [optional] 
**Limit** | **string** |  | [optional] 
**ReturnedContent** | [**V1alpha1GeneDatasetRequestContentType**](v1alpha1GeneDatasetRequestContentType.md) |  | [optional] 
**SortSchema** | [**GeneDatasetRequestSort**](GeneDatasetRequestSort.md) |  | [optional] 
**SymbolsForTaxon** | [**GeneDatasetRequestSymbolsForTaxon**](GeneDatasetRequestSymbolsForTaxon.md) |  | [optional] 
**Taxon** | **string** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


