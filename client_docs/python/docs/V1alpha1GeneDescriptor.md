# V1alpha1GeneDescriptor

--- MongoRequiredFields:   GeneDescriptor:     required_fields:       - geneId       - symbol       - taxId       - taxname ...
## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**chromosome** | **str** |  | [optional] 
**common_name** | **str** |  | [optional] 
**description** | **str** |  | [optional] 
**gene_id** | **str** |  | [optional] 
**genomic_ranges** | [**list[V1alpha1SeqRangeSet]**](V1alpha1SeqRangeSet.md) |  | [optional] 
**orientation** | [**V1alpha1Orientation**](V1alpha1Orientation.md) |  | [optional] 
**proteins** | [**list[V1alpha1Protein]**](V1alpha1Protein.md) |  | [optional] 
**symbol** | **str** |  | [optional] 
**tax_id** | **str** |  | [optional] 
**taxname** | **str** |  | [optional] 
**transcripts** | [**list[V1alpha1Transcript]**](V1alpha1Transcript.md) |  | [optional] 
**type** | [**GeneDescriptorGeneType**](GeneDescriptorGeneType.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


