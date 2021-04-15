# V1alpha1AssemblyMetadataRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**accessions** | [**Datasetsv1alpha1Accessions**](Datasetsv1alpha1Accessions.md) |  | [optional] 
**bioprojects** | [**AssemblyMetadataRequestBioprojects**](AssemblyMetadataRequestBioprojects.md) |  | [optional] 
**filters** | [**V1alpha1AssemblyDatasetDescriptorsFilter**](V1alpha1AssemblyDatasetDescriptorsFilter.md) |  | [optional] 
**limit** | **str** |  | [optional] 
**returned_content** | [**V1alpha1AssemblyMetadataRequestContentType**](V1alpha1AssemblyMetadataRequestContentType.md) |  | [optional] 
**tax_exact_match** | **bool** | If true, only return assemblies with the given NCBI Taxonomy ID, or name. Otherwise, assemblies from taxonomy subtree are included, too. Ignored for assembly_accession request. | [optional] 
**taxon** | **str** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


