# V1alpha1GeneDatasetRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**accessions** | [**V1alpha1GeneDatasetRequestAccessions**](V1alpha1GeneDatasetRequestAccessions.md) |  | [optional] 
**gene_ids** | [**GeneDatasetRequestGeneIds**](GeneDatasetRequestGeneIds.md) |  | [optional] 
**include_annotation_type** | [**list[GeneDatasetRequestFasta]**](GeneDatasetRequestFasta.md) | Select additional types of annotation to include in the data package.  If unset, no annotation is provided. | [optional] 
**limit** | **str** |  | [optional] 
**returned_content** | [**V1alpha1GeneDatasetRequestContentType**](V1alpha1GeneDatasetRequestContentType.md) |  | [optional] 
**sort_schema** | [**GeneDatasetRequestSort**](GeneDatasetRequestSort.md) |  | [optional] 
**symbols_for_taxon** | [**GeneDatasetRequestSymbolsForTaxon**](GeneDatasetRequestSymbolsForTaxon.md) |  | [optional] 
**taxon** | **str** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


