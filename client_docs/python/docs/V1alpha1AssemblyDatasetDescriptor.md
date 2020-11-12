# V1alpha1AssemblyDatasetDescriptor

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**annotation_metadata** | [**V1alpha1AnnotationForAssembly**](V1alpha1AnnotationForAssembly.md) |  | [optional] 
**assembly_accession** | **str** |  | [optional] 
**assembly_category** | **str** | Category of assembly, like reference. | [optional] 
**assembly_level** | **str** |  | [optional] 
**bioproject_lineages** | [**list[V1alpha1BioProjectLineage]**](V1alpha1BioProjectLineage.md) |  | [optional] 
**chromosomes** | **list[str]** | Which chromosomes are included in this dataset. NB: Mitochondria is encoded as &#39;MT&#39;. | [optional] 
**contig_n50** | **int** |  | [optional] 
**display_name** | **str** | The name of the Assembly dataset to be displayed in a user interface. | [optional] 
**estimated_size** | **str** |  | [optional] 
**org** | [**V1alpha1Organism**](V1alpha1Organism.md) |  | [optional] 
**seq_length** | **str** |  | [optional] 
**submission_date** | **str** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


