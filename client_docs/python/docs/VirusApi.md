# ncbi.datasets.openapi.VirusApi

All URIs are relative to *https://api.ncbi.nlm.nih.gov/datasets/v1alpha*

Method | HTTP request | Description
------------- | ------------- | -------------
[**sars2_protein_download**](VirusApi.md#sars2_protein_download) | **GET** /virus/taxon/sars2/protein/{proteins}/download | Download SARS-CoV-2 protein and CDS datasets by protein name
[**sars2_protein_summary**](VirusApi.md#sars2_protein_summary) | **GET** /virus/taxon/sars2/protein/{proteins} | Summary of SARS-CoV-2 protein and CDS datasets by protein name
[**sars2_protein_table**](VirusApi.md#sars2_protein_table) | **GET** /virus/taxon/sars2/protein/{proteins}/table | Get SARS-CoV-2 protein metadata in a tabular format.
[**virus_genome_download**](VirusApi.md#virus_genome_download) | **GET** /virus/taxon/{taxon}/genome/download | Download Coronavirus genome datasets by taxon
[**virus_genome_summary**](VirusApi.md#virus_genome_summary) | **GET** /virus/taxon/{taxon}/genome | Get summary data for Coronaviridae genomes by taxon
[**virus_genome_table**](VirusApi.md#virus_genome_table) | **GET** /virus/taxon/{taxon}/genome/table | Get viral genomic metadata in a tabular format.


# **sars2_protein_download**
> file sars2_protein_download(proteins, refseq_only=refseq_only, annotated_only=annotated_only, released_since=released_since, host=host, geo_location=geo_location, complete_only=complete_only, include_annotation_type=include_annotation_type, filename=filename)

Download SARS-CoV-2 protein and CDS datasets by protein name

Download a SARS-CoV-2 protein datasets

### Example

```python
from __future__ import print_function
import time
import ncbi.datasets.openapi
from ncbi.datasets.openapi.rest import ApiException
from pprint import pprint
# Defining the host is optional and defaults to https://api.ncbi.nlm.nih.gov/datasets/v1alpha
# See configuration.py for a list of all supported configuration parameters.
configuration = ncbi.datasets.openapi.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1alpha"
)


# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = ncbi.datasets.openapi.VirusApi(api_client)
    proteins = ['proteins_example'] # list[str] | Which proteins to retrieve in the data package
refseq_only = True # bool | If true, limit results to RefSeq genomes. (optional)
annotated_only = True # bool | If true, limit results to annotated genomes. (optional)
released_since = '2013-10-20T19:20:30+01:00' # datetime | If set, limit results to viral genomes that have been released after a specified date and time. April 1, 2020 midnight UTC should be formatted as follows: 2020-04-01T00:00:00.000Z. (optional)
host = 'host_example' # str | If set, limit results to genomes extracted from this host (Taxonomy ID or name) or its taxonomic descendants. (optional)
geo_location = 'geo_location_example' # str | Assemblies from this location (country and state, or continent). (optional)
complete_only = True # bool | only include complete genomes. (optional)
include_annotation_type = ['include_annotation_type_example'] # list[str] | Select additional types of annotation to include in the data package.  If unset, no annotation is provided. (optional)
filename = 'ncbi_dataset.zip' # str | Output file name. (optional) (default to 'ncbi_dataset.zip')

    try:
        # Download SARS-CoV-2 protein and CDS datasets by protein name
        api_response = api_instance.sars2_protein_download(proteins, refseq_only=refseq_only, annotated_only=annotated_only, released_since=released_since, host=host, geo_location=geo_location, complete_only=complete_only, include_annotation_type=include_annotation_type, filename=filename)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling VirusApi->sars2_protein_download: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **proteins** | [**list[str]**](str.md)| Which proteins to retrieve in the data package | 
 **refseq_only** | **bool**| If true, limit results to RefSeq genomes. | [optional] 
 **annotated_only** | **bool**| If true, limit results to annotated genomes. | [optional] 
 **released_since** | **datetime**| If set, limit results to viral genomes that have been released after a specified date and time. April 1, 2020 midnight UTC should be formatted as follows: 2020-04-01T00:00:00.000Z. | [optional] 
 **host** | **str**| If set, limit results to genomes extracted from this host (Taxonomy ID or name) or its taxonomic descendants. | [optional] 
 **geo_location** | **str**| Assemblies from this location (country and state, or continent). | [optional] 
 **complete_only** | **bool**| only include complete genomes. | [optional] 
 **include_annotation_type** | [**list[str]**](str.md)| Select additional types of annotation to include in the data package.  If unset, no annotation is provided. | [optional] 
 **filename** | **str**| Output file name. | [optional] [default to &#39;ncbi_dataset.zip&#39;]

### Return type

**file**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/zip

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Download selected SARS-CoV-2 proteins and associated annotation data as a zip file. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **sars2_protein_summary**
> V1alpha1DownloadSummary sars2_protein_summary(proteins, refseq_only=refseq_only, annotated_only=annotated_only, released_since=released_since, host=host, geo_location=geo_location, complete_only=complete_only, include_annotation_type=include_annotation_type)

Summary of SARS-CoV-2 protein and CDS datasets by protein name

Download a summary of available SARS-CoV-2 protein datasets

### Example

```python
from __future__ import print_function
import time
import ncbi.datasets.openapi
from ncbi.datasets.openapi.rest import ApiException
from pprint import pprint
# Defining the host is optional and defaults to https://api.ncbi.nlm.nih.gov/datasets/v1alpha
# See configuration.py for a list of all supported configuration parameters.
configuration = ncbi.datasets.openapi.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1alpha"
)


# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = ncbi.datasets.openapi.VirusApi(api_client)
    proteins = ['proteins_example'] # list[str] | Which proteins to retrieve in the data package
refseq_only = True # bool | If true, limit results to RefSeq genomes. (optional)
annotated_only = True # bool | If true, limit results to annotated genomes. (optional)
released_since = '2013-10-20T19:20:30+01:00' # datetime | If set, limit results to viral genomes that have been released after a specified date and time. April 1, 2020 midnight UTC should be formatted as follows: 2020-04-01T00:00:00.000Z. (optional)
host = 'host_example' # str | If set, limit results to genomes extracted from this host (Taxonomy ID or name) or its taxonomic descendants. (optional)
geo_location = 'geo_location_example' # str | Assemblies from this location (country and state, or continent). (optional)
complete_only = True # bool | only include complete genomes. (optional)
include_annotation_type = ['include_annotation_type_example'] # list[str] | Select additional types of annotation to include in the data package.  If unset, no annotation is provided. (optional)

    try:
        # Summary of SARS-CoV-2 protein and CDS datasets by protein name
        api_response = api_instance.sars2_protein_summary(proteins, refseq_only=refseq_only, annotated_only=annotated_only, released_since=released_since, host=host, geo_location=geo_location, complete_only=complete_only, include_annotation_type=include_annotation_type)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling VirusApi->sars2_protein_summary: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **proteins** | [**list[str]**](str.md)| Which proteins to retrieve in the data package | 
 **refseq_only** | **bool**| If true, limit results to RefSeq genomes. | [optional] 
 **annotated_only** | **bool**| If true, limit results to annotated genomes. | [optional] 
 **released_since** | **datetime**| If set, limit results to viral genomes that have been released after a specified date and time. April 1, 2020 midnight UTC should be formatted as follows: 2020-04-01T00:00:00.000Z. | [optional] 
 **host** | **str**| If set, limit results to genomes extracted from this host (Taxonomy ID or name) or its taxonomic descendants. | [optional] 
 **geo_location** | **str**| Assemblies from this location (country and state, or continent). | [optional] 
 **complete_only** | **bool**| only include complete genomes. | [optional] 
 **include_annotation_type** | [**list[str]**](str.md)| Select additional types of annotation to include in the data package.  If unset, no annotation is provided. | [optional] 

### Return type

[**V1alpha1DownloadSummary**](V1alpha1DownloadSummary.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **sars2_protein_table**
> StreamResultOfV1alpha1TabularOutput sars2_protein_table(proteins, refseq_only=refseq_only, annotated_only=annotated_only, released_since=released_since, host=host, geo_location=geo_location, complete_only=complete_only, table_fields=table_fields, format=format)

Get SARS-CoV-2 protein metadata in a tabular format.

Get protein metadata in tabular format for SARS-CoV-2 genomes.

### Example

```python
from __future__ import print_function
import time
import ncbi.datasets.openapi
from ncbi.datasets.openapi.rest import ApiException
from pprint import pprint
# Defining the host is optional and defaults to https://api.ncbi.nlm.nih.gov/datasets/v1alpha
# See configuration.py for a list of all supported configuration parameters.
configuration = ncbi.datasets.openapi.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1alpha"
)


# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = ncbi.datasets.openapi.VirusApi(api_client)
    proteins = ['proteins_example'] # list[str] | Which proteins to retrieve in the data package
refseq_only = True # bool | If true, limit results to RefSeq genomes. (optional)
annotated_only = True # bool | If true, limit results to annotated genomes. (optional)
released_since = '2013-10-20T19:20:30+01:00' # datetime | If set, limit results to viral genomes that have been released after a specified date and time. April 1, 2020 midnight UTC should be formatted as follows: 2020-04-01T00:00:00.000Z. (optional)
host = 'host_example' # str | If set, limit results to genomes extracted from this host (Taxonomy ID or name) or its taxonomic descendants. (optional)
geo_location = 'geo_location_example' # str | Assemblies from this location (country and state, or continent). (optional)
complete_only = True # bool | only include complete genomes. (optional)
table_fields = ['table_fields_example'] # list[str] | Specify which fields to include in the tabular report. (optional)
format = 'tsv' # str | Choose download format. (optional) (default to 'tsv')

    try:
        # Get SARS-CoV-2 protein metadata in a tabular format.
        api_response = api_instance.sars2_protein_table(proteins, refseq_only=refseq_only, annotated_only=annotated_only, released_since=released_since, host=host, geo_location=geo_location, complete_only=complete_only, table_fields=table_fields, format=format)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling VirusApi->sars2_protein_table: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **proteins** | [**list[str]**](str.md)| Which proteins to retrieve in the data package | 
 **refseq_only** | **bool**| If true, limit results to RefSeq genomes. | [optional] 
 **annotated_only** | **bool**| If true, limit results to annotated genomes. | [optional] 
 **released_since** | **datetime**| If set, limit results to viral genomes that have been released after a specified date and time. April 1, 2020 midnight UTC should be formatted as follows: 2020-04-01T00:00:00.000Z. | [optional] 
 **host** | **str**| If set, limit results to genomes extracted from this host (Taxonomy ID or name) or its taxonomic descendants. | [optional] 
 **geo_location** | **str**| Assemblies from this location (country and state, or continent). | [optional] 
 **complete_only** | **bool**| only include complete genomes. | [optional] 
 **table_fields** | [**list[str]**](str.md)| Specify which fields to include in the tabular report. | [optional] 
 **format** | **str**| Choose download format. | [optional] [default to &#39;tsv&#39;]

### Return type

[**StreamResultOfV1alpha1TabularOutput**](StreamResultOfV1alpha1TabularOutput.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/tsv

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response.(streaming responses) |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **virus_genome_download**
> file virus_genome_download(taxon, refseq_only=refseq_only, annotated_only=annotated_only, released_since=released_since, host=host, geo_location=geo_location, complete_only=complete_only, exclude_sequence=exclude_sequence, include_annotation_type=include_annotation_type, filename=filename)

Download Coronavirus genome datasets by taxon

Download a Coronavirus genome datasets by taxon

### Example

```python
from __future__ import print_function
import time
import ncbi.datasets.openapi
from ncbi.datasets.openapi.rest import ApiException
from pprint import pprint
# Defining the host is optional and defaults to https://api.ncbi.nlm.nih.gov/datasets/v1alpha
# See configuration.py for a list of all supported configuration parameters.
configuration = ncbi.datasets.openapi.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1alpha"
)


# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = ncbi.datasets.openapi.VirusApi(api_client)
    taxon = 'taxon_example' # str | NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank
refseq_only = True # bool | If true, limit results to RefSeq genomes. (optional)
annotated_only = True # bool | If true, limit results to annotated genomes. (optional)
released_since = '2013-10-20T19:20:30+01:00' # datetime | If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as follows: 2020-04-01T00:00:00.000Z. (optional)
host = 'host_example' # str | If set, limit results to genomes extracted from this host (Taxonomy ID or name) or its taxonomic descendants. (optional)
geo_location = 'geo_location_example' # str | Assemblies from this location (country and state, or continent). (optional)
complete_only = True # bool | only include complete genomes. (optional)
exclude_sequence = True # bool | Set to true to omit the genomic sequence. (optional)
include_annotation_type = ['include_annotation_type_example'] # list[str] | Select additional types of annotation to include in the data package.  If unset, no annotation is provided. (optional)
filename = 'ncbi_dataset.zip' # str | Output file name. (optional) (default to 'ncbi_dataset.zip')

    try:
        # Download Coronavirus genome datasets by taxon
        api_response = api_instance.virus_genome_download(taxon, refseq_only=refseq_only, annotated_only=annotated_only, released_since=released_since, host=host, geo_location=geo_location, complete_only=complete_only, exclude_sequence=exclude_sequence, include_annotation_type=include_annotation_type, filename=filename)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling VirusApi->virus_genome_download: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taxon** | **str**| NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank | 
 **refseq_only** | **bool**| If true, limit results to RefSeq genomes. | [optional] 
 **annotated_only** | **bool**| If true, limit results to annotated genomes. | [optional] 
 **released_since** | **datetime**| If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as follows: 2020-04-01T00:00:00.000Z. | [optional] 
 **host** | **str**| If set, limit results to genomes extracted from this host (Taxonomy ID or name) or its taxonomic descendants. | [optional] 
 **geo_location** | **str**| Assemblies from this location (country and state, or continent). | [optional] 
 **complete_only** | **bool**| only include complete genomes. | [optional] 
 **exclude_sequence** | **bool**| Set to true to omit the genomic sequence. | [optional] 
 **include_annotation_type** | [**list[str]**](str.md)| Select additional types of annotation to include in the data package.  If unset, no annotation is provided. | [optional] 
 **filename** | **str**| Output file name. | [optional] [default to &#39;ncbi_dataset.zip&#39;]

### Return type

**file**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/zip

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Download selected viral genome and associated annotation data as a zip file. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **virus_genome_summary**
> V1alpha1DownloadSummary virus_genome_summary(taxon, refseq_only=refseq_only, annotated_only=annotated_only, released_since=released_since, host=host, geo_location=geo_location, complete_only=complete_only, exclude_sequence=exclude_sequence, include_annotation_type=include_annotation_type)

Get summary data for Coronaviridae genomes by taxon

Get summary data and download by command line instructions for Coronaviridae genomes by taxon.

### Example

```python
from __future__ import print_function
import time
import ncbi.datasets.openapi
from ncbi.datasets.openapi.rest import ApiException
from pprint import pprint
# Defining the host is optional and defaults to https://api.ncbi.nlm.nih.gov/datasets/v1alpha
# See configuration.py for a list of all supported configuration parameters.
configuration = ncbi.datasets.openapi.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1alpha"
)


# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = ncbi.datasets.openapi.VirusApi(api_client)
    taxon = 'taxon_example' # str | NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank
refseq_only = True # bool | If true, limit results to RefSeq genomes. (optional)
annotated_only = True # bool | If true, limit results to annotated genomes. (optional)
released_since = '2013-10-20T19:20:30+01:00' # datetime | If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as follows: 2020-04-01T00:00:00.000Z. (optional)
host = 'host_example' # str | If set, limit results to genomes extracted from this host (Taxonomy ID or name) or its taxonomic descendants. (optional)
geo_location = 'geo_location_example' # str | Assemblies from this location (country and state, or continent). (optional)
complete_only = True # bool | only include complete genomes. (optional)
exclude_sequence = True # bool | Set to true to omit the genomic sequence. (optional)
include_annotation_type = ['include_annotation_type_example'] # list[str] | Select additional types of annotation to include in the data package.  If unset, no annotation is provided. (optional)

    try:
        # Get summary data for Coronaviridae genomes by taxon
        api_response = api_instance.virus_genome_summary(taxon, refseq_only=refseq_only, annotated_only=annotated_only, released_since=released_since, host=host, geo_location=geo_location, complete_only=complete_only, exclude_sequence=exclude_sequence, include_annotation_type=include_annotation_type)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling VirusApi->virus_genome_summary: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taxon** | **str**| NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank | 
 **refseq_only** | **bool**| If true, limit results to RefSeq genomes. | [optional] 
 **annotated_only** | **bool**| If true, limit results to annotated genomes. | [optional] 
 **released_since** | **datetime**| If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as follows: 2020-04-01T00:00:00.000Z. | [optional] 
 **host** | **str**| If set, limit results to genomes extracted from this host (Taxonomy ID or name) or its taxonomic descendants. | [optional] 
 **geo_location** | **str**| Assemblies from this location (country and state, or continent). | [optional] 
 **complete_only** | **bool**| only include complete genomes. | [optional] 
 **exclude_sequence** | **bool**| Set to true to omit the genomic sequence. | [optional] 
 **include_annotation_type** | [**list[str]**](str.md)| Select additional types of annotation to include in the data package.  If unset, no annotation is provided. | [optional] 

### Return type

[**V1alpha1DownloadSummary**](V1alpha1DownloadSummary.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **virus_genome_table**
> StreamResultOfV1alpha1TabularOutput virus_genome_table(taxon, refseq_only=refseq_only, annotated_only=annotated_only, released_since=released_since, host=host, geo_location=geo_location, complete_only=complete_only, table_fields=table_fields, format=format)

Get viral genomic metadata in a tabular format.

Get viral genomic metadata in tabular format for Coronaviridae genomes by taxon.

### Example

```python
from __future__ import print_function
import time
import ncbi.datasets.openapi
from ncbi.datasets.openapi.rest import ApiException
from pprint import pprint
# Defining the host is optional and defaults to https://api.ncbi.nlm.nih.gov/datasets/v1alpha
# See configuration.py for a list of all supported configuration parameters.
configuration = ncbi.datasets.openapi.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1alpha"
)


# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = ncbi.datasets.openapi.VirusApi(api_client)
    taxon = 'taxon_example' # str | NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank
refseq_only = True # bool | If true, limit results to RefSeq genomes. (optional)
annotated_only = True # bool | If true, limit results to annotated genomes. (optional)
released_since = '2013-10-20T19:20:30+01:00' # datetime | If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as follows: 2020-04-01T00:00:00.000Z. (optional)
host = 'host_example' # str | If set, limit results to genomes extracted from this host (Taxonomy ID or name) or its taxonomic descendants. (optional)
geo_location = 'geo_location_example' # str | Assemblies from this location (country and state, or continent). (optional)
complete_only = True # bool | only include complete genomes. (optional)
table_fields = ['table_fields_example'] # list[str] | Specify which fields to include in the tabular report. (optional)
format = 'tsv' # str | Choose download format (tsv, csv or jsonl). (optional) (default to 'tsv')

    try:
        # Get viral genomic metadata in a tabular format.
        api_response = api_instance.virus_genome_table(taxon, refseq_only=refseq_only, annotated_only=annotated_only, released_since=released_since, host=host, geo_location=geo_location, complete_only=complete_only, table_fields=table_fields, format=format)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling VirusApi->virus_genome_table: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taxon** | **str**| NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank | 
 **refseq_only** | **bool**| If true, limit results to RefSeq genomes. | [optional] 
 **annotated_only** | **bool**| If true, limit results to annotated genomes. | [optional] 
 **released_since** | **datetime**| If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as follows: 2020-04-01T00:00:00.000Z. | [optional] 
 **host** | **str**| If set, limit results to genomes extracted from this host (Taxonomy ID or name) or its taxonomic descendants. | [optional] 
 **geo_location** | **str**| Assemblies from this location (country and state, or continent). | [optional] 
 **complete_only** | **bool**| only include complete genomes. | [optional] 
 **table_fields** | [**list[str]**](str.md)| Specify which fields to include in the tabular report. | [optional] 
 **format** | **str**| Choose download format (tsv, csv or jsonl). | [optional] [default to &#39;tsv&#39;]

### Return type

[**StreamResultOfV1alpha1TabularOutput**](StreamResultOfV1alpha1TabularOutput.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/tsv

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response.(streaming responses) |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

