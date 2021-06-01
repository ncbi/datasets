# \GeneApi

All URIs are relative to *https://api.ncbi.nlm.nih.gov/datasets/v1alpha*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadGenePackage**](GeneApi.md#DownloadGenePackage) | **Get** /gene/id/{gene_ids}/download | Get a gene dataset by gene ID
[**DownloadGenePackagePost**](GeneApi.md#DownloadGenePackagePost) | **Post** /gene/download | Get a gene dataset by POST
[**GeneDownloadSummaryByAccession**](GeneApi.md#GeneDownloadSummaryByAccession) | **Get** /gene/accession/{accessions}/download_summary | Get gene download summary by RefSeq Accession
[**GeneDownloadSummaryById**](GeneApi.md#GeneDownloadSummaryById) | **Get** /gene/id/{gene_ids}/download_summary | Get gene download summary by GeneID
[**GeneDownloadSummaryByPost**](GeneApi.md#GeneDownloadSummaryByPost) | **Post** /gene/download_summary | Get gene download summary
[**GeneDownloadSummaryByTaxAndSymbol**](GeneApi.md#GeneDownloadSummaryByTaxAndSymbol) | **Get** /gene/symbol/{symbols}/taxon/{taxon}/download_summary | Get gene download summary by gene symbol.
[**GeneMetadataByAccession**](GeneApi.md#GeneMetadataByAccession) | **Get** /gene/accession/{accessions} | Get gene metadata by RefSeq Accession
[**GeneMetadataById**](GeneApi.md#GeneMetadataById) | **Get** /gene/id/{gene_ids} | Get gene metadata by GeneID
[**GeneMetadataByPost**](GeneApi.md#GeneMetadataByPost) | **Post** /gene | Get gene metadata as JSON
[**GeneMetadataByTaxAndSymbol**](GeneApi.md#GeneMetadataByTaxAndSymbol) | **Get** /gene/symbol/{symbols}/taxon/{taxon} | Get gene metadata by gene symbol.
[**GeneMetadataStreamByPost**](GeneApi.md#GeneMetadataStreamByPost) | **Post** /gene/stream | Get gene metadata
[**GeneOrthologsById**](GeneApi.md#GeneOrthologsById) | **Get** /gene/id/{gene_id}/orthologs | Get gene orthologs by gene ID
[**GeneTaxNameQuery**](GeneApi.md#GeneTaxNameQuery) | **Get** /gene/taxon_suggest/{taxon_query} | Get a list of taxonomy names and IDs found in the gene dataset given a partial taxonomic name.
[**GeneTaxTree**](GeneApi.md#GeneTaxTree) | **Get** /gene/taxon/{taxon}/tree | Retrieve tax tree



## DownloadGenePackage

> *os.File DownloadGenePackage(ctx, geneIds, optional)

Get a gene dataset by gene ID

Get a gene dataset including gene, transcript and protein fasta sequence, annotation and metadata by gene ID.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**geneIds** | [**[]int64**](int64.md)| NCBI gene ids | 
 **optional** | ***DownloadGenePackageOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DownloadGenePackageOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeAnnotationType** | [**optional.Interface of []string**](string.md)| Select additional types of annotation to include in the data package.  If unset, no annotation is provided. | 
 **fastaFilter** | [**optional.Interface of []string**](string.md)| Limit the FASTA sequences in the datasets package to these transcript and protein accessions. | 
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


## DownloadGenePackagePost

> *os.File DownloadGenePackagePost(ctx, body, optional)

Get a gene dataset by POST

Get a gene dataset including gene, transcript and protein fasta sequence, annotation and metadata by POST.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**V1alpha1GeneDatasetRequest**](V1alpha1GeneDatasetRequest.md)|  | 
 **optional** | ***DownloadGenePackagePostOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DownloadGenePackagePostOpts struct


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


## GeneDownloadSummaryByAccession

> V1alpha1DownloadSummary GeneDownloadSummaryByAccession(ctx, accessions, optional)

Get gene download summary by RefSeq Accession

Get gene download summary by RefSeq Accession in a JSON output format.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessions** | [**[]string**](string.md)| RNA or Protein accessions. | 
 **optional** | ***GeneDownloadSummaryByAccessionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GeneDownloadSummaryByAccessionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.String**| Limit the number of returned results (\&quot;all\&quot;, \&quot;none\&quot;, otherwise an integer value). | 
 **fastaFilter** | [**optional.Interface of []string**](string.md)| Limit the FASTA sequences in the datasets package to these transcript and protein accessions. | 

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


## GeneDownloadSummaryById

> V1alpha1DownloadSummary GeneDownloadSummaryById(ctx, geneIds, optional)

Get gene download summary by GeneID

Get a download summary by GeneID in a JSON output format.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**geneIds** | [**[]int64**](int64.md)| NCBI gene ids | 
 **optional** | ***GeneDownloadSummaryByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GeneDownloadSummaryByIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.String**| Limit the number of returned results (\&quot;all\&quot;, \&quot;none\&quot;, otherwise an integer value). | 
 **fastaFilter** | [**optional.Interface of []string**](string.md)| Limit the FASTA sequences in the datasets package to these transcript and protein accessions. | 

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


## GeneDownloadSummaryByPost

> V1alpha1DownloadSummary GeneDownloadSummaryByPost(ctx, body)

Get gene download summary

Get gene download summary in a JSON output format.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**V1alpha1GeneDatasetRequest**](V1alpha1GeneDatasetRequest.md)|  | 

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


## GeneDownloadSummaryByTaxAndSymbol

> V1alpha1DownloadSummary GeneDownloadSummaryByTaxAndSymbol(ctx, symbols, taxon, optional)

Get gene download summary by gene symbol.

Get gene download summary by gene symbol in a JSON output format.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**symbols** | [**[]string**](string.md)|  | 
**taxon** | **string**|  | 
 **optional** | ***GeneDownloadSummaryByTaxAndSymbolOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GeneDownloadSummaryByTaxAndSymbolOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fastaFilter** | [**optional.Interface of []string**](string.md)| Limit the FASTA sequences in the datasets package to these transcript and protein accessions. | 

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


## GeneMetadataByAccession

> V1alpha1GeneMetadata GeneMetadataByAccession(ctx, accessions, optional)

Get gene metadata by RefSeq Accession

Get detailed gene metadata by RefSeq Accession in a JSON output format.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessions** | [**[]string**](string.md)| RNA or Protein accessions. | 
 **optional** | ***GeneMetadataByAccessionOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GeneMetadataByAccessionOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **returnedContent** | **optional.String**| Return either gene-ids, or entire gene metadata. | [default to COMPLETE]
 **sortSchemaField** | **optional.String**| Select a field to sort on. | [default to SORT_FIELD_UNSPECIFIED]
 **sortSchemaDirection** | **optional.String**| Select a direction for the sort. | [default to SORT_DIRECTION_UNSPECIFIED]
 **limit** | **optional.String**| Limit the number of returned results (\&quot;all\&quot;, \&quot;none\&quot;, otherwise an integer value). | 

### Return type

[**V1alpha1GeneMetadata**](v1alpha1GeneMetadata.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GeneMetadataById

> V1alpha1GeneMetadata GeneMetadataById(ctx, geneIds, optional)

Get gene metadata by GeneID

Get detailed gene metadata by GeneID in a JSON output format.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**geneIds** | [**[]int64**](int64.md)| NCBI gene ids | 
 **optional** | ***GeneMetadataByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GeneMetadataByIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **returnedContent** | **optional.String**| Return either gene-ids, or entire gene metadata. | [default to COMPLETE]
 **sortSchemaField** | **optional.String**| Select a field to sort on. | [default to SORT_FIELD_UNSPECIFIED]
 **sortSchemaDirection** | **optional.String**| Select a direction for the sort. | [default to SORT_DIRECTION_UNSPECIFIED]
 **limit** | **optional.String**| Limit the number of returned results (\&quot;all\&quot;, \&quot;none\&quot;, otherwise an integer value). | 

### Return type

[**V1alpha1GeneMetadata**](v1alpha1GeneMetadata.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GeneMetadataByPost

> V1alpha1GeneMetadata GeneMetadataByPost(ctx, body)

Get gene metadata as JSON

Get detailed gene metadata in a JSON output format.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**V1alpha1GeneDatasetRequest**](V1alpha1GeneDatasetRequest.md)|  | 

### Return type

[**V1alpha1GeneMetadata**](v1alpha1GeneMetadata.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GeneMetadataByTaxAndSymbol

> V1alpha1GeneMetadata GeneMetadataByTaxAndSymbol(ctx, symbols, taxon, optional)

Get gene metadata by gene symbol.

Get detailed gene metadata by gene symbol in a JSON output format.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**symbols** | [**[]string**](string.md)|  | 
**taxon** | **string**|  | 
 **optional** | ***GeneMetadataByTaxAndSymbolOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GeneMetadataByTaxAndSymbolOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **returnedContent** | **optional.String**| Return either gene-ids, or entire gene metadata. | [default to COMPLETE]
 **sortSchemaField** | **optional.String**| Select a field to sort on. | [default to SORT_FIELD_UNSPECIFIED]
 **sortSchemaDirection** | **optional.String**| Select a direction for the sort. | [default to SORT_DIRECTION_UNSPECIFIED]
 **limit** | **optional.String**| Limit the number of returned results (\&quot;all\&quot;, \&quot;none\&quot;, otherwise an integer value). | 

### Return type

[**V1alpha1GeneMetadata**](v1alpha1GeneMetadata.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GeneMetadataStreamByPost

> V1alpha1GeneMatch GeneMetadataStreamByPost(ctx, body)

Get gene metadata

Get detailed gene metadata in a streaming, JSON-lines output format.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**V1alpha1GeneDatasetRequest**](V1alpha1GeneDatasetRequest.md)|  | 

### Return type

[**V1alpha1GeneMatch**](v1alpha1GeneMatch.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/x-ndjson

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GeneOrthologsById

> V1alpha1OrthologSet GeneOrthologsById(ctx, geneId, optional)

Get gene orthologs by gene ID

Get detailed gene metadata for an ortholog set by gene ID in a JSON output format.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**geneId** | **int64**|  | 
 **optional** | ***GeneOrthologsByIdOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GeneOrthologsByIdOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **returnedContent** | **optional.String**| Return either gene-ids, or entire gene metadata. | [default to COMPLETE]
 **taxonFilter** | [**optional.Interface of []string**](string.md)| Filter genes by taxa. | 

### Return type

[**V1alpha1OrthologSet**](v1alpha1OrthologSet.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GeneTaxNameQuery

> V1alpha1SciNameAndIds GeneTaxNameQuery(ctx, taxonQuery)

Get a list of taxonomy names and IDs found in the gene dataset given a partial taxonomic name.

This endpoint retrieves a list of taxonomy names and IDs found in the gene dataset given a partial taxonomic name of any rank.

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


## GeneTaxTree

> V1alpha1Organism GeneTaxTree(ctx, taxon)

Retrieve tax tree

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

