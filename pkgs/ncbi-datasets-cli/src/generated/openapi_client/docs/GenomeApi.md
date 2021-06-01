# \GenomeApi

All URIs are relative to *https://api.ncbi.nlm.nih.gov/datasets/v1alpha*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssemblyDescriptorsByAccessions**](GenomeApi.md#AssemblyDescriptorsByAccessions) | **Get** /genome/accession/{accessions} | Get genome metadata by accession
[**AssemblyDescriptorsByBioproject**](GenomeApi.md#AssemblyDescriptorsByBioproject) | **Get** /genome/bioproject/{accessions} | Get genome metadata by bioproject accession
[**AssemblyDescriptorsByTaxon**](GenomeApi.md#AssemblyDescriptorsByTaxon) | **Get** /genome/taxon/{taxon} | Get genome metadata by taxonomic identifier
[**CheckAssemblyAvailability**](GenomeApi.md#CheckAssemblyAvailability) | **Get** /genome/accession/{accessions}/check | Check the validity of genome accessions
[**CheckAssemblyAvailabilityPost**](GenomeApi.md#CheckAssemblyAvailabilityPost) | **Post** /genome/check | Check the validity of many genome accessions in a single request
[**DownloadAssemblyPackage**](GenomeApi.md#DownloadAssemblyPackage) | **Get** /genome/accession/{accessions}/download | Get a genome dataset by accession
[**DownloadAssemblyPackagePost**](GenomeApi.md#DownloadAssemblyPackagePost) | **Post** /genome/download | Get a genome dataset by post
[**GenomeDownloadSummary**](GenomeApi.md#GenomeDownloadSummary) | **Get** /genome/accession/{accessions}/download_summary | Preview genome dataset download
[**GenomeDownloadSummaryByPost**](GenomeApi.md#GenomeDownloadSummaryByPost) | **Post** /genome/download_summary | Preview genome dataset download by POST
[**GenomeMetadataByPost**](GenomeApi.md#GenomeMetadataByPost) | **Post** /genome | Get genome metadata by accession
[**GenomeTaxNameQuery**](GenomeApi.md#GenomeTaxNameQuery) | **Get** /genome/taxon_suggest/{taxon_query} | Get a list of taxonomy names and IDs found in the assembly dataset given a partial taxonomic name.
[**GenomeTaxTree**](GenomeApi.md#GenomeTaxTree) | **Get** /genome/taxon/{taxon}/tree | Get a taxonomic subtree by taxonomic identifier



## AssemblyDescriptorsByAccessions

> V1alpha1AssemblyMetadata AssemblyDescriptorsByAccessions(ctx, accessions, optional)

Get genome metadata by accession

Get detailed metadata for assembled genomes by accession in a JSON output format.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessions** | [**[]string**](string.md)|  | 
 **optional** | ***AssemblyDescriptorsByAccessionsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AssemblyDescriptorsByAccessionsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filtersReferenceOnly** | **optional.Bool**| If true, only return reference and representative (GCF_ and GCA_) genome assemblies. | 
 **filtersAssemblySource** | **optional.String**| Return only RefSeq (GCF_) or GenBank (GCA_) genome assemblies. | [default to all]
 **filtersHasAnnotation** | **optional.Bool**| Return only annotated genome assemblies. | 
 **filtersAssemblyLevel** | [**optional.Interface of []string**](string.md)| Only return genome assemblies that have one of the specified assembly levels. | 
 **filtersFirstReleaseDate** | **optional.Time**| Only return genome assemblies that were released on or after the specified date. | 
 **filtersLastReleaseDate** | **optional.Time**| Only return genome assemblies that were released on or before to the specified date. | 
 **filtersSearchText** | [**optional.Interface of []string**](string.md)| Only return results whose fields contain the specified search terms in their taxon, infraspecific, assembly name or submitter fields. | 
 **pageSize** | **optional.Int32**| The maximum number of genome assemblies to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size,  &#x60;page_token&#x60; can be used to retrieve the remaining results. | 
 **pageToken** | **optional.String**| A page token is returned from an &#x60;AssemblyMetadataRequest&#x60; call with more than &#x60;page_size&#x60; results. Use this token, along with the previous  &#x60;AssemblyMetadataRequest&#x60; parameters, to retrieve the next page of results. When &#x60;page_token&#x60; is empty, all results have been retrieved. | 

### Return type

[**V1alpha1AssemblyMetadata**](v1alpha1AssemblyMetadata.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssemblyDescriptorsByBioproject

> V1alpha1AssemblyMetadata AssemblyDescriptorsByBioproject(ctx, accessions, optional)

Get genome metadata by bioproject accession

Get detailed metadata for assembled genomes by bioproject accession in a JSON output format.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessions** | [**[]string**](string.md)|  | 
 **optional** | ***AssemblyDescriptorsByBioprojectOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AssemblyDescriptorsByBioprojectOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filtersReferenceOnly** | **optional.Bool**| If true, only return reference and representative (GCF_ and GCA_) genome assemblies. | 
 **filtersAssemblySource** | **optional.String**| Return only RefSeq (GCF_) or GenBank (GCA_) genome assemblies. | [default to all]
 **filtersHasAnnotation** | **optional.Bool**| Return only annotated genome assemblies. | 
 **filtersAssemblyLevel** | [**optional.Interface of []string**](string.md)| Only return genome assemblies that have one of the specified assembly levels. | 
 **filtersFirstReleaseDate** | **optional.Time**| Only return genome assemblies that were released on or after the specified date. | 
 **filtersLastReleaseDate** | **optional.Time**| Only return genome assemblies that were released on or before to the specified date. | 
 **filtersSearchText** | [**optional.Interface of []string**](string.md)| Only return results whose fields contain the specified search terms in their taxon, infraspecific, assembly name or submitter fields. | 
 **returnedContent** | **optional.String**| Return either assembly accessions, or entire assembly-metadata records. | [default to COMPLETE]
 **pageSize** | **optional.Int32**| The maximum number of genome assemblies to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size,  &#x60;page_token&#x60; can be used to retrieve the remaining results. | 
 **pageToken** | **optional.String**| A page token is returned from an &#x60;AssemblyMetadataRequest&#x60; call with more than &#x60;page_size&#x60; results. Use this token, along with the previous  &#x60;AssemblyMetadataRequest&#x60; parameters, to retrieve the next page of results. When &#x60;page_token&#x60; is empty, all results have been retrieved. | 

### Return type

[**V1alpha1AssemblyMetadata**](v1alpha1AssemblyMetadata.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssemblyDescriptorsByTaxon

> V1alpha1AssemblyMetadata AssemblyDescriptorsByTaxon(ctx, taxon, optional)

Get genome metadata by taxonomic identifier

Get detailed metadata on all assembled genomes for a specified NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxon** | **string**| NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank | 
 **optional** | ***AssemblyDescriptorsByTaxonOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AssemblyDescriptorsByTaxonOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filtersReferenceOnly** | **optional.Bool**| If true, only return reference and representative (GCF_ and GCA_) genome assemblies. | 
 **filtersAssemblySource** | **optional.String**| Return only RefSeq (GCF_) or GenBank (GCA_) genome assemblies. | [default to all]
 **filtersHasAnnotation** | **optional.Bool**| Return only annotated genome assemblies. | 
 **filtersAssemblyLevel** | [**optional.Interface of []string**](string.md)| Only return genome assemblies that have one of the specified assembly levels. | 
 **filtersFirstReleaseDate** | **optional.Time**| Only return genome assemblies that were released on or after the specified date. | 
 **filtersLastReleaseDate** | **optional.Time**| Only return genome assemblies that were released on or before to the specified date. | 
 **filtersSearchText** | [**optional.Interface of []string**](string.md)| Only return results whose fields contain the specified search terms in their taxon, infraspecific, assembly name or submitter fields. | 
 **taxExactMatch** | **optional.Bool**| If true, only return assemblies with the given NCBI Taxonomy ID, or name. Otherwise, assemblies from taxonomy subtree are included, too. Ignored for assembly_accession request. | 
 **returnedContent** | **optional.String**| Return either assembly accessions, or entire assembly-metadata records. | [default to COMPLETE]
 **pageSize** | **optional.Int32**| The maximum number of genome assemblies to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size,  &#x60;page_token&#x60; can be used to retrieve the remaining results. | 
 **pageToken** | **optional.String**| A page token is returned from an &#x60;AssemblyMetadataRequest&#x60; call with more than &#x60;page_size&#x60; results. Use this token, along with the previous  &#x60;AssemblyMetadataRequest&#x60; parameters, to retrieve the next page of results. When &#x60;page_token&#x60; is empty, all results have been retrieved. | 

### Return type

[**V1alpha1AssemblyMetadata**](v1alpha1AssemblyMetadata.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckAssemblyAvailability

> V1alpha1AssemblyDatasetAvailability CheckAssemblyAvailability(ctx, accessions)

Check the validity of genome accessions

The 'GET' version of check is limited by the size of the GET URL (2KB, which works out to about 140 genomic accessions).  The POST operation is provided to allow users to supply a larger number of accessions in a single request.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessions** | [**[]string**](string.md)| NCBI genome assembly accessions | 

### Return type

[**V1alpha1AssemblyDatasetAvailability**](v1alpha1AssemblyDatasetAvailability.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckAssemblyAvailabilityPost

> V1alpha1AssemblyDatasetAvailability CheckAssemblyAvailabilityPost(ctx, body)

Check the validity of many genome accessions in a single request

The 'GET' version of check is limited by the size of the GET URL (2KB, which works out to about 140 genomic accessions).  The POST operation is provided to allow users to supply a larger number of accessions in a single request.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**V1alpha1AssemblyDatasetRequest**](V1alpha1AssemblyDatasetRequest.md)|  | 

### Return type

[**V1alpha1AssemblyDatasetAvailability**](v1alpha1AssemblyDatasetAvailability.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadAssemblyPackage

> *os.File DownloadAssemblyPackage(ctx, accessions, optional)

Get a genome dataset by accession

Download a genome dataset including fasta sequence, annotation and a detailed data report by accession.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessions** | [**[]string**](string.md)| NCBI genome assembly accessions | 
 **optional** | ***DownloadAssemblyPackageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DownloadAssemblyPackageOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chromosomes** | [**optional.Interface of []string**](string.md)| The default setting is all chromosome. Specify individual chromosome by string (1,2,MT or chr1,chr2.chrMT). Unplaced sequences are treated like their own chromosome (&#39;Un&#39;). The filter only applies to fasta sequence. | 
 **excludeSequence** | **optional.Bool**| Set to true to omit the genomic sequence. | 
 **includeAnnotationType** | [**optional.Interface of []string**](string.md)| Select additional types of annotation to include in the data package.  If unset, no annotation is provided. | 
 **hydrated** | **optional.String**| Set to DATA_REPORT_ONLY, to only retrieve data-reports. | [default to FULLY_HYDRATED]
 **filename** | **optional.String**| Output file name. | [default to ncbi_dataset.zip]

### Return type

[***os.File**](*os.File.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/zip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadAssemblyPackagePost

> *os.File DownloadAssemblyPackagePost(ctx, body, optional)

Get a genome dataset by post

The 'GET' version of download is limited by the size of the GET URL (2KB, which works out to about 140 genomic accessions).  The POST operation is provided to allow users to supply a larger number of accessions in a single request.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**V1alpha1AssemblyDatasetRequest**](V1alpha1AssemblyDatasetRequest.md)|  | 
 **optional** | ***DownloadAssemblyPackagePostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DownloadAssemblyPackagePostOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filename** | **optional.String**| Output file name. | [default to ncbi_dataset.zip]

### Return type

[***os.File**](*os.File.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/zip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenomeDownloadSummary

> V1alpha1DownloadSummary GenomeDownloadSummary(ctx, accessions, optional)

Preview genome dataset download

Get a download summary by accession in a JSON output format.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessions** | [**[]string**](string.md)| NCBI genome assembly accessions | 
 **optional** | ***GenomeDownloadSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GenomeDownloadSummaryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chromosomes** | [**optional.Interface of []string**](string.md)| The default setting is all chromosome. Specify individual chromosome by string (1,2,MT or chr1,chr2.chrMT). Unplaced sequences are treated like their own chromosome (&#39;Un&#39;). The filter only applies to fasta sequence. | 
 **excludeSequence** | **optional.Bool**| Set to true to omit the genomic sequence. | 
 **includeAnnotationType** | [**optional.Interface of []string**](string.md)| Select additional types of annotation to include in the data package.  If unset, no annotation is provided. | 

### Return type

[**V1alpha1DownloadSummary**](v1alpha1DownloadSummary.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenomeDownloadSummaryByPost

> V1alpha1DownloadSummary GenomeDownloadSummaryByPost(ctx, body)

Preview genome dataset download by POST

The 'GET' version of download summary is limited by the size of the GET URL (2KB, which works out to about 140 genomic accessions).  The POST operation is provided to allow users to supply a larger number of accessions in a single request.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**V1alpha1AssemblyDatasetRequest**](V1alpha1AssemblyDatasetRequest.md)|  | 

### Return type

[**V1alpha1DownloadSummary**](v1alpha1DownloadSummary.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenomeMetadataByPost

> V1alpha1AssemblyMetadata GenomeMetadataByPost(ctx, body)

Get genome metadata by accession

Get detailed metadata for assembled genomes by accession in a JSON output format.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**V1alpha1AssemblyMetadataRequest**](V1alpha1AssemblyMetadataRequest.md)|  | 

### Return type

[**V1alpha1AssemblyMetadata**](v1alpha1AssemblyMetadata.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenomeTaxNameQuery

> V1alpha1SciNameAndIds GenomeTaxNameQuery(ctx, taxonQuery)

Get a list of taxonomy names and IDs found in the assembly dataset given a partial taxonomic name.

This endpoint retrieves a list of taxonomy names and IDs found in the assembly dataset given a partial taxonomic name of any rank.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxonQuery** | **string**| NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank | 

### Return type

[**V1alpha1SciNameAndIds**](v1alpha1SciNameAndIds.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenomeTaxTree

> V1alpha1Organism GenomeTaxTree(ctx, taxon)

Get a taxonomic subtree by taxonomic identifier

Using a NCBI Taxonomy ID or name (common or scientific) at any rank, get a subtree filtered for species with assembled genomes

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxon** | **string**| NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank | 

### Return type

[**V1alpha1Organism**](v1alpha1Organism.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

