# V1alpha1AssemblyDatasetDescriptorsFilter

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**assembly_level** | [**list[AssemblyDatasetDescriptorsFilterAssemblyLevel]**](AssemblyDatasetDescriptorsFilterAssemblyLevel.md) |  | [optional] 
**assembly_source** | [**AssemblyDatasetDescriptorsFilterAssemblySource**](AssemblyDatasetDescriptorsFilterAssemblySource.md) |  | [optional] 
**first_release_date** | **datetime** |  | [optional] 
**has_annotation** | **bool** |  | [optional] 
**last_release_date** | **datetime** |  | [optional] 
**reference_only** | **bool** | If true, only return reference and representative (GCF_ and GCA_) genome assemblies. | [optional] 
**refseq_only** | **bool** | If true, only return RefSeq (GCF_) genome assemblies. Deprecated - use assembly_type instead. | [optional] 
**search_text** | **list[str]** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


