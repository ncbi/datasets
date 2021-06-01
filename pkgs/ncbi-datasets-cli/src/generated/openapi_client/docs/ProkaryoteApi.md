# \ProkaryoteApi

All URIs are relative to *https://api.ncbi.nlm.nih.gov/datasets/v1alpha*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadProkaryoteGenePackage**](ProkaryoteApi.md#DownloadProkaryoteGenePackage) | **Get** /protein/accession/{accessions}/download | Get a prokaryote gene dataset by RefSeq protein accession
[**DownloadProkaryoteGenePackagePost**](ProkaryoteApi.md#DownloadProkaryoteGenePackagePost) | **Post** /protein/accession/download | Get a prokaryote gene dataset by RefSeq protein accession by POST



## DownloadProkaryoteGenePackage

> *os.File DownloadProkaryoteGenePackage(ctx, accessions, optional)

Get a prokaryote gene dataset by RefSeq protein accession

Get a prokaryote gene dataset including gene and protein fasta sequence, annotation and metadata by prokaryote protein accession.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessions** | [**[]string**](string.md)| WP prokaryote protein accession | 
 **optional** | ***DownloadProkaryoteGenePackageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DownloadProkaryoteGenePackageOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeAnnotationType** | [**optional.Interface of []string**](string.md)| Select additional types of annotation to include in the data package.  If unset, no annotation is provided. | 
 **geneFlankConfigLength** | **optional.Int64**|  | 
 **taxon** | **optional.String**| NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank When specified, return data from this taxon and its subtree. | 
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


## DownloadProkaryoteGenePackagePost

> *os.File DownloadProkaryoteGenePackagePost(ctx, body, optional)

Get a prokaryote gene dataset by RefSeq protein accession by POST

Get a prokaryote gene dataset including gene and protein fasta sequence, annotation and metadata by prokaryote protein accession by POST.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**V1alpha1ProkaryoteGeneRequest**](V1alpha1ProkaryoteGeneRequest.md)|  | 
 **optional** | ***DownloadProkaryoteGenePackagePostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DownloadProkaryoteGenePackagePostOpts struct


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

