# V1alpha1AssemblyDatasetDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnnotationMetadata** | [**V1alpha1AnnotationForAssembly**](v1alpha1AnnotationForAssembly.md) |  | [optional] 
**AssemblyAccession** | **string** |  | [optional] 
**AssemblyCategory** | **string** | Category of assembly, like reference. | [optional] 
**AssemblyLevel** | **string** |  | [optional] 
**BioprojectLineages** | [**[]V1alpha1BioProjectLineage**](v1alpha1BioProjectLineage.md) |  | [optional] 
**BiosampleAccession** | **string** | NCBI BioSample Accession for the BioSample from which the sequences in the genome assembly were obtained. | [optional] 
**Chromosomes** | **[]string** | Which chromosomes are included in this dataset. NB: Mitochondria is encoded as &#39;MT&#39;. | [optional] 
**ContigN50** | **int64** |  | [optional] 
**DisplayName** | **string** | The name of the Assembly dataset to be displayed in a user interface. | [optional] 
**EstimatedSize** | **string** |  | [optional] 
**Org** | [**V1alpha1Organism**](v1alpha1Organism.md) |  | [optional] 
**SeqLength** | **string** |  | [optional] 
**SubmissionDate** | **string** |  | [optional] 
**Submitter** | **string** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


