# V1alpha1GeneDatasetRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**accessions** | **list[str]** | RNA or Protein accessions. | [optional] 
**fasta_filter** | **list[str]** |  | [optional] 
**gene_ids** | **list[int]** |  | [optional] 
**include_annotation_type** | [**list[V1alpha1Fasta]**](V1alpha1Fasta.md) | Select additional types of annotation to include in the data package.  If unset, no annotation is provided. | [optional] 
**limit** | **str** |  | [optional] 
**returned_content** | [**V1alpha1GeneDatasetRequestContentType**](V1alpha1GeneDatasetRequestContentType.md) |  | [optional] 
**sort_schema** | [**GeneDatasetRequestSort**](GeneDatasetRequestSort.md) |  | [optional] 
**symbols_for_taxon** | [**GeneDatasetRequestSymbolsForTaxon**](GeneDatasetRequestSymbolsForTaxon.md) |  | [optional] 
**taxon** | **str** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


