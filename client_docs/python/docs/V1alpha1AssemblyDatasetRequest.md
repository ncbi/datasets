# V1alpha1AssemblyDatasetRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**assembly_accessions** | **list[str]** | Use &#39;add item&#39; to include multiple assembly accessions. | [optional] 
**chromosomes** | **list[str]** | The default setting is all chromosome. Specify individual chromosome by string (1,2,MT or chr1,chr2.chrMT). Unplaced sequences are treated like their own chromosome (&#39;Un&#39;). The filter only applies to fasta sequence. | [optional] 
**hydrated** | [**AssemblyDatasetRequestResolution**](AssemblyDatasetRequestResolution.md) |  | [optional] 
**include_annotation** | **bool** |  | [optional] 
**include_annotation_type** | [**list[AnnotationForAssemblyType]**](AnnotationForAssemblyType.md) |  | [optional] 
**include_sequence** | **bool** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


