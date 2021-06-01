# V1alpha1AssemblyMetadataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accessions** | [**Datasetsv1alpha1Accessions**](datasetsv1alpha1Accessions.md) |  | [optional] 
**Bioprojects** | [**AssemblyMetadataRequestBioprojects**](AssemblyMetadataRequestBioprojects.md) |  | [optional] 
**Filters** | [**V1alpha1AssemblyDatasetDescriptorsFilter**](v1alpha1AssemblyDatasetDescriptorsFilter.md) |  | [optional] 
**Limit** | **string** |  | [optional] 
**PageSize** | **int32** | The maximum number of genome assemblies to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size,  &#x60;page_token&#x60; can be used to retrieve the remaining results. | [optional] 
**PageToken** | **string** | A page token is returned from an &#x60;AssemblyMetadataRequest&#x60; call with more than &#x60;page_size&#x60; results. Use this token, along with the previous  &#x60;AssemblyMetadataRequest&#x60; parameters, to retrieve the next page of results. When &#x60;page_token&#x60; is empty, all results have been retrieved. | [optional] 
**ReturnedContent** | [**V1alpha1AssemblyMetadataRequestContentType**](v1alpha1AssemblyMetadataRequestContentType.md) |  | [optional] 
**TaxExactMatch** | **bool** | If true, only return assemblies with the given NCBI Taxonomy ID, or name. Otherwise, assemblies from taxonomy subtree are included, too. Ignored for assembly_accession request. | [optional] 
**Taxon** | **string** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


