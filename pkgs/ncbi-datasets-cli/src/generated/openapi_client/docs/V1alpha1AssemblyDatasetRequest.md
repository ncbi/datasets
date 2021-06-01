# V1alpha1AssemblyDatasetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accessions** | **[]string** |  | [optional] 
**AssemblyAccessions** | **[]string** |  | [optional] 
**Chromosomes** | **[]string** | The default setting is all chromosome. Specify individual chromosome by string (1,2,MT or chr1,chr2.chrMT). Unplaced sequences are treated like their own chromosome (&#39;Un&#39;). The filter only applies to fasta sequence. | [optional] 
**ExcludeSequence** | **bool** | Set to true to omit the genomic sequence. | [optional] 
**Hydrated** | [**AssemblyDatasetRequestResolution**](AssemblyDatasetRequestResolution.md) |  | [optional] 
**IncludeAnnotation** | **bool** |  | [optional] 
**IncludeAnnotationType** | [**[]V1alpha1AnnotationForAssemblyType**](v1alpha1AnnotationForAssemblyType.md) | Select additional types of annotation to include in the data package.  If unset, no annotation is provided. | [optional] 
**IncludeTsv** | **bool** | Set to true to include a TSV represention of the data-report. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


