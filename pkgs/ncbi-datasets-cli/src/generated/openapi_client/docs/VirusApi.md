# \VirusApi

All URIs are relative to *https://api.ncbi.nlm.nih.gov/datasets/v1alpha*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Sars2ProteinDownload**](VirusApi.md#Sars2ProteinDownload) | **Get** /virus/taxon/sars2/protein/{proteins}/download | Download SARS-CoV-2 protein and CDS datasets by protein name
[**Sars2ProteinSummary**](VirusApi.md#Sars2ProteinSummary) | **Get** /virus/taxon/sars2/protein/{proteins} | Summary of SARS-CoV-2 protein and CDS datasets by protein name
[**Sars2ProteinTable**](VirusApi.md#Sars2ProteinTable) | **Get** /virus/taxon/sars2/protein/{proteins}/table | Get SARS-CoV-2 protein metadata in a tabular format.
[**VirusGenomeDownload**](VirusApi.md#VirusGenomeDownload) | **Get** /virus/taxon/{taxon}/genome/download | Download Coronavirus genome datasets by taxon
[**VirusGenomeSummary**](VirusApi.md#VirusGenomeSummary) | **Get** /virus/taxon/{taxon}/genome | Get summary data for Coronaviridae genomes by taxon
[**VirusGenomeTable**](VirusApi.md#VirusGenomeTable) | **Get** /virus/taxon/{taxon}/genome/table | Get viral genomic metadata in a tabular format.



## Sars2ProteinDownload

> *os.File Sars2ProteinDownload(ctx, proteins, optional)

Download SARS-CoV-2 protein and CDS datasets by protein name

Download a SARS-CoV-2 protein datasets

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**proteins** | [**[]string**](string.md)| Which proteins to retrieve in the data package | 
 **optional** | ***Sars2ProteinDownloadOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a Sars2ProteinDownloadOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **refseqOnly** | **optional.Bool**| If true, limit results to RefSeq genomes. | 
 **annotatedOnly** | **optional.Bool**| If true, limit results to annotated genomes. | 
 **releasedSince** | **optional.Time**| If set, limit results to viral genomes that have been released after a specified date and time. April 1, 2020 midnight UTC should be formatted as follows: 2020-04-01T00:00:00.000Z. | 
 **host** | **optional.String**| If set, limit results to genomes extracted from this host (Taxonomy ID or name). | 
 **geoLocation** | **optional.String**| Assemblies from this location (country and state, or continent). | 
 **completeOnly** | **optional.Bool**| only include complete genomes. | 
 **includeAnnotationType** | [**optional.Interface of []string**](string.md)| Select additional types of annotation to include in the data package.  If unset, no annotation is provided. | 
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


## Sars2ProteinSummary

> V1alpha1DownloadSummary Sars2ProteinSummary(ctx, proteins, optional)

Summary of SARS-CoV-2 protein and CDS datasets by protein name

Download a summary of available SARS-CoV-2 protein datasets

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**proteins** | [**[]string**](string.md)| Which proteins to retrieve in the data package | 
 **optional** | ***Sars2ProteinSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a Sars2ProteinSummaryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **refseqOnly** | **optional.Bool**| If true, limit results to RefSeq genomes. | 
 **annotatedOnly** | **optional.Bool**| If true, limit results to annotated genomes. | 
 **releasedSince** | **optional.Time**| If set, limit results to viral genomes that have been released after a specified date and time. April 1, 2020 midnight UTC should be formatted as follows: 2020-04-01T00:00:00.000Z. | 
 **host** | **optional.String**| If set, limit results to genomes extracted from this host (Taxonomy ID or name). | 
 **geoLocation** | **optional.String**| Assemblies from this location (country and state, or continent). | 
 **completeOnly** | **optional.Bool**| only include complete genomes. | 
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


## Sars2ProteinTable

> StreamResultOfV1alpha1TabularOutput Sars2ProteinTable(ctx, proteins, optional)

Get SARS-CoV-2 protein metadata in a tabular format.

Get protein metadata in tabular format for SARS-CoV-2 genomes.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**proteins** | [**[]string**](string.md)| Which proteins to retrieve in the data package | 
 **optional** | ***Sars2ProteinTableOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a Sars2ProteinTableOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **refseqOnly** | **optional.Bool**| If true, limit results to RefSeq genomes. | 
 **annotatedOnly** | **optional.Bool**| If true, limit results to annotated genomes. | 
 **releasedSince** | **optional.Time**| If set, limit results to viral genomes that have been released after a specified date and time. April 1, 2020 midnight UTC should be formatted as follows: 2020-04-01T00:00:00.000Z. | 
 **host** | **optional.String**| If set, limit results to genomes extracted from this host (Taxonomy ID or name). | 
 **geoLocation** | **optional.String**| Assemblies from this location (country and state, or continent). | 
 **completeOnly** | **optional.Bool**| only include complete genomes. | 
 **tableFields** | [**optional.Interface of []string**](string.md)| Specify which fields to include in the tabular report. | 
 **format** | **optional.String**| Choose download format. | [default to tsv]

### Return type

[**StreamResultOfV1alpha1TabularOutput**](Stream_result_of_v1alpha1TabularOutput.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/tsv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VirusGenomeDownload

> *os.File VirusGenomeDownload(ctx, taxon, optional)

Download Coronavirus genome datasets by taxon

Download a Coronavirus genome datasets by taxon

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxon** | **string**| NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank | 
 **optional** | ***VirusGenomeDownloadOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a VirusGenomeDownloadOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **refseqOnly** | **optional.Bool**| If true, limit results to RefSeq genomes. | 
 **annotatedOnly** | **optional.Bool**| If true, limit results to annotated genomes. | 
 **releasedSince** | **optional.Time**| If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as follows: 2020-04-01T00:00:00.000Z. | 
 **host** | **optional.String**| If set, limit results to genomes extracted from this host (Taxonomy ID or name). | 
 **pangolinClassification** | **optional.String**| If set, limit results to genomes classified to this lineage by the PangoLearn tool. | 
 **geoLocation** | **optional.String**| Assemblies from this location (country and state, or continent). | 
 **completeOnly** | **optional.Bool**| only include complete genomes. | 
 **excludeSequence** | **optional.Bool**| Set to true to omit the genomic sequence. | 
 **includeAnnotationType** | [**optional.Interface of []string**](string.md)| Select additional types of annotation to include in the data package.  If unset, no annotation is provided. | 
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


## VirusGenomeSummary

> V1alpha1DownloadSummary VirusGenomeSummary(ctx, taxon, optional)

Get summary data for Coronaviridae genomes by taxon

Get summary data and download by command line instructions for Coronaviridae genomes by taxon.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxon** | **string**| NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank | 
 **optional** | ***VirusGenomeSummaryOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a VirusGenomeSummaryOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **refseqOnly** | **optional.Bool**| If true, limit results to RefSeq genomes. | 
 **annotatedOnly** | **optional.Bool**| If true, limit results to annotated genomes. | 
 **releasedSince** | **optional.Time**| If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as follows: 2020-04-01T00:00:00.000Z. | 
 **host** | **optional.String**| If set, limit results to genomes extracted from this host (Taxonomy ID or name). | 
 **pangolinClassification** | **optional.String**| If set, limit results to genomes classified to this lineage by the PangoLearn tool. | 
 **geoLocation** | **optional.String**| Assemblies from this location (country and state, or continent). | 
 **completeOnly** | **optional.Bool**| only include complete genomes. | 
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


## VirusGenomeTable

> StreamResultOfV1alpha1TabularOutput VirusGenomeTable(ctx, taxon, optional)

Get viral genomic metadata in a tabular format.

Get viral genomic metadata in tabular format for Coronaviridae genomes by taxon.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxon** | **string**| NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank | 
 **optional** | ***VirusGenomeTableOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a VirusGenomeTableOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **refseqOnly** | **optional.Bool**| If true, limit results to RefSeq genomes. | 
 **annotatedOnly** | **optional.Bool**| If true, limit results to annotated genomes. | 
 **releasedSince** | **optional.Time**| If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as follows: 2020-04-01T00:00:00.000Z. | 
 **host** | **optional.String**| If set, limit results to genomes extracted from this host (Taxonomy ID or name). | 
 **pangolinClassification** | **optional.String**| If set, limit results to genomes classified to this lineage by the PangoLearn tool. | 
 **geoLocation** | **optional.String**| Assemblies from this location (country and state, or continent). | 
 **completeOnly** | **optional.Bool**| only include complete genomes. | 
 **tableFields** | [**optional.Interface of []string**](string.md)| Specify which fields to include in the tabular report. | 
 **format** | **optional.String**| Choose download format (tsv, csv or jsonl). | [default to tsv]

### Return type

[**StreamResultOfV1alpha1TabularOutput**](Stream_result_of_v1alpha1TabularOutput.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/tsv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

