# V1alpha1AssemblyDatasetRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**accessions** | **list[str]** |  | [optional] 
**assembly_accessions** | **list[str]** |  | [optional] 
**chromosomes** | **list[str]** | The default setting is all chromosome. Specify individual chromosome by string (1,2,MT or chr1,chr2.chrMT). Unplaced sequences are treated like their own chromosome (&#39;Un&#39;). The filter only applies to fasta sequence. | [optional] 
**exclude_sequence** | **bool** | Set to true to omit the genomic sequence. | [optional] 
**hydrated** | [**AssemblyDatasetRequestResolution**](AssemblyDatasetRequestResolution.md) |  | [optional] 
**include_annotation** | **bool** |  | [optional] 
**include_annotation_type** | [**list[V1alpha1AnnotationForAssemblyType]**](V1alpha1AnnotationForAssemblyType.md) | Select additional types of annotation to include in the data package.  If unset, no annotation is provided. | [optional] 
**include_tsv** | **bool** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


