# V1alpha1AssemblyDatasetDescriptorsFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssemblyLevel** | [**[]AssemblyDatasetDescriptorsFilterAssemblyLevel**](AssemblyDatasetDescriptorsFilterAssemblyLevel.md) |  | [optional] 
**AssemblySource** | [**AssemblyDatasetDescriptorsFilterAssemblySource**](AssemblyDatasetDescriptorsFilterAssemblySource.md) |  | [optional] 
**FirstReleaseDate** | [**time.Time**](time.Time.md) |  | [optional] 
**HasAnnotation** | **bool** |  | [optional] 
**LastReleaseDate** | [**time.Time**](time.Time.md) |  | [optional] 
**ReferenceOnly** | **bool** | If true, only return reference and representative (GCF_ and GCA_) genome assemblies. | [optional] 
**RefseqOnly** | **bool** | If true, only return RefSeq (GCF_) genome assemblies. Deprecated - use assembly_type instead. | [optional] 
**SearchText** | **[]string** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


