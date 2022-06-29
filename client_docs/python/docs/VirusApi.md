# ncbi.datasets.openapi.VirusApi

All URIs are relative to *https://api.ncbi.nlm.nih.gov/datasets/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**sars2_protein_download**](VirusApi.md#sars2_protein_download) | **GET** /virus/taxon/sars2/protein/{proteins}/download | Download SARS-CoV-2 protein and CDS datasets by protein name
[**sars2_protein_summary**](VirusApi.md#sars2_protein_summary) | **GET** /virus/taxon/sars2/protein/{proteins} | Summary of SARS-CoV-2 protein and CDS datasets by protein name
[**sars2_protein_table**](VirusApi.md#sars2_protein_table) | **GET** /virus/taxon/sars2/protein/{proteins}/table | Get SARS-CoV-2 protein metadata in a tabular format.
[**virus_genome_download**](VirusApi.md#virus_genome_download) | **GET** /virus/taxon/{taxon}/genome/download | Download a coronavirus genome dataset by taxon
[**virus_genome_download_accession**](VirusApi.md#virus_genome_download_accession) | **GET** /virus/accession/{accessions}/genome/download | Download a coronavirus genome dataset by accession
[**virus_genome_download_post**](VirusApi.md#virus_genome_download_post) | **POST** /virus/genome/download | Get a coronavirus genome dataset by post
[**virus_genome_summary**](VirusApi.md#virus_genome_summary) | **GET** /virus/taxon/{taxon}/genome | Get summary data for Coronaviridae genomes by taxon
[**virus_genome_summary_accession**](VirusApi.md#virus_genome_summary_accession) | **GET** /virus/accession/{accessions}/genome | Get summary data for Coronaviridae genomes by accession
[**virus_genome_summary_post**](VirusApi.md#virus_genome_summary_post) | **POST** /virus/genome | Get summary data for Coronaviridae genomes by post
[**virus_genome_table**](VirusApi.md#virus_genome_table) | **GET** /virus/taxon/{taxon}/genome/table | Get viral genomic metadata in a tabular format.
[**virus_reports_by_acessions**](VirusApi.md#virus_reports_by_acessions) | **GET** /virus/accession/{accessions}/dataset_report | Get virus metadata by accession
[**virus_reports_by_post**](VirusApi.md#virus_reports_by_post) | **POST** /virus | Get virus metadata by POST
[**virus_reports_by_taxon**](VirusApi.md#virus_reports_by_taxon) | **GET** /virus/taxon/{taxon}/dataset_report | Get virus metadata by taxon


# **sars2_protein_download**
> file_type sars2_protein_download(proteins)

Download SARS-CoV-2 protein and CDS datasets by protein name

Download a SARS-CoV-2 protein datasets

### Example

* Api Key Authentication (ApiKeyAuthHeader):

```python
import time
import ncbi.datasets.openapi
from ncbi.datasets.openapi.api import virus_api
from ncbi.datasets.openapi.model.v1_annotation_for_virus_type import V1AnnotationForVirusType
from ncbi.datasets.openapi.model.rpc_status import RpcStatus
from pprint import pprint
# Defining the host is optional and defaults to https://api.ncbi.nlm.nih.gov/datasets/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = ncbi.datasets.openapi.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: ApiKeyAuthHeader
configuration.api_key['ApiKeyAuthHeader'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['ApiKeyAuthHeader'] = 'Bearer'

# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = virus_api.VirusApi(api_client)
    proteins = [
        "spike protein",
    ] # [str] | Which proteins to retrieve in the data package
    refseq_only = True # bool | If true, limit results to RefSeq genomes. (optional) if omitted the server will use the default value of False
    annotated_only = True # bool | If true, limit results to annotated genomes. (optional) if omitted the server will use the default value of False
    released_since = dateutil_parser('2020-08-01') # datetime | If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as '2020-04-01T00:00:00.000Z' (optional)
    updated_since = dateutil_parser('2021-07-18') # datetime |  (optional)
    host = "human" # str | If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default (optional)
    geo_location = "USA" # str | Assemblies from this location (country and state, or continent) (optional)
    complete_only = True # bool | only include complete genomes. (optional) if omitted the server will use the default value of False
    include_annotation_type = [
        V1AnnotationForVirusType("CDS_FASTA"),
    ] # [V1AnnotationForVirusType] | Select additional types of annotation to include in the data package.  If unset, no annotation is provided. (optional)
    filename = "ncbi_dataset.zip" # str | Output file name. (optional) if omitted the server will use the default value of "ncbi_dataset.zip"

    # example passing only required values which don't have defaults set
    try:
        # Download SARS-CoV-2 protein and CDS datasets by protein name
        api_response = api_instance.sars2_protein_download(proteins)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling VirusApi->sars2_protein_download: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Download SARS-CoV-2 protein and CDS datasets by protein name
        api_response = api_instance.sars2_protein_download(proteins, refseq_only=refseq_only, annotated_only=annotated_only, released_since=released_since, updated_since=updated_since, host=host, geo_location=geo_location, complete_only=complete_only, include_annotation_type=include_annotation_type, filename=filename)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling VirusApi->sars2_protein_download: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **proteins** | **[str]**| Which proteins to retrieve in the data package |
 **refseq_only** | **bool**| If true, limit results to RefSeq genomes. | [optional] if omitted the server will use the default value of False
 **annotated_only** | **bool**| If true, limit results to annotated genomes. | [optional] if omitted the server will use the default value of False
 **released_since** | **datetime**| If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as &#39;2020-04-01T00:00:00.000Z&#39; | [optional]
 **updated_since** | **datetime**|  | [optional]
 **host** | **str**| If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default | [optional]
 **geo_location** | **str**| Assemblies from this location (country and state, or continent) | [optional]
 **complete_only** | **bool**| only include complete genomes. | [optional] if omitted the server will use the default value of False
 **include_annotation_type** | [**[V1AnnotationForVirusType]**](V1AnnotationForVirusType.md)| Select additional types of annotation to include in the data package.  If unset, no annotation is provided. | [optional]
 **filename** | **str**| Output file name. | [optional] if omitted the server will use the default value of "ncbi_dataset.zip"

### Return type

**file_type**

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/zip, application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Download selected genome assemblies and associated annotation data as a zip file |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **sars2_protein_summary**
> V1DownloadSummary sars2_protein_summary(proteins)

Summary of SARS-CoV-2 protein and CDS datasets by protein name

Download a summary of available SARS-CoV-2 protein datasets

### Example

* Api Key Authentication (ApiKeyAuthHeader):

```python
import time
import ncbi.datasets.openapi
from ncbi.datasets.openapi.api import virus_api
from ncbi.datasets.openapi.model.v1_annotation_for_virus_type import V1AnnotationForVirusType
from ncbi.datasets.openapi.model.v1_download_summary import V1DownloadSummary
from ncbi.datasets.openapi.model.rpc_status import RpcStatus
from pprint import pprint
# Defining the host is optional and defaults to https://api.ncbi.nlm.nih.gov/datasets/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = ncbi.datasets.openapi.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: ApiKeyAuthHeader
configuration.api_key['ApiKeyAuthHeader'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['ApiKeyAuthHeader'] = 'Bearer'

# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = virus_api.VirusApi(api_client)
    proteins = [
        "spike protein",
    ] # [str] | Which proteins to retrieve in the data package
    refseq_only = True # bool | If true, limit results to RefSeq genomes. (optional) if omitted the server will use the default value of False
    annotated_only = True # bool | If true, limit results to annotated genomes. (optional) if omitted the server will use the default value of False
    released_since = dateutil_parser('2020-08-01') # datetime | If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as '2020-04-01T00:00:00.000Z' (optional)
    updated_since = dateutil_parser('2021-07-18') # datetime |  (optional)
    host = "human" # str | If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default (optional)
    geo_location = "USA" # str | Assemblies from this location (country and state, or continent) (optional)
    complete_only = True # bool | only include complete genomes. (optional) if omitted the server will use the default value of False
    include_annotation_type = [
        V1AnnotationForVirusType("CDS_FASTA"),
    ] # [V1AnnotationForVirusType] | Select additional types of annotation to include in the data package.  If unset, no annotation is provided. (optional)

    # example passing only required values which don't have defaults set
    try:
        # Summary of SARS-CoV-2 protein and CDS datasets by protein name
        api_response = api_instance.sars2_protein_summary(proteins)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling VirusApi->sars2_protein_summary: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Summary of SARS-CoV-2 protein and CDS datasets by protein name
        api_response = api_instance.sars2_protein_summary(proteins, refseq_only=refseq_only, annotated_only=annotated_only, released_since=released_since, updated_since=updated_since, host=host, geo_location=geo_location, complete_only=complete_only, include_annotation_type=include_annotation_type)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling VirusApi->sars2_protein_summary: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **proteins** | **[str]**| Which proteins to retrieve in the data package |
 **refseq_only** | **bool**| If true, limit results to RefSeq genomes. | [optional] if omitted the server will use the default value of False
 **annotated_only** | **bool**| If true, limit results to annotated genomes. | [optional] if omitted the server will use the default value of False
 **released_since** | **datetime**| If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as &#39;2020-04-01T00:00:00.000Z&#39; | [optional]
 **updated_since** | **datetime**|  | [optional]
 **host** | **str**| If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default | [optional]
 **geo_location** | **str**| Assemblies from this location (country and state, or continent) | [optional]
 **complete_only** | **bool**| only include complete genomes. | [optional] if omitted the server will use the default value of False
 **include_annotation_type** | [**[V1AnnotationForVirusType]**](V1AnnotationForVirusType.md)| Select additional types of annotation to include in the data package.  If unset, no annotation is provided. | [optional]

### Return type

[**V1DownloadSummary**](V1DownloadSummary.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **sars2_protein_table**
> V1TabularOutput sars2_protein_table(proteins)

Get SARS-CoV-2 protein metadata in a tabular format.

Get protein metadata in tabular format for SARS-CoV-2 genomes.

### Example

* Api Key Authentication (ApiKeyAuthHeader):

```python
import time
import ncbi.datasets.openapi
from ncbi.datasets.openapi.api import virus_api
from ncbi.datasets.openapi.model.v1_virus_table_field import V1VirusTableField
from ncbi.datasets.openapi.model.v1_tabular_output import V1TabularOutput
from ncbi.datasets.openapi.model.v1_table_format import V1TableFormat
from ncbi.datasets.openapi.model.rpc_status import RpcStatus
from pprint import pprint
# Defining the host is optional and defaults to https://api.ncbi.nlm.nih.gov/datasets/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = ncbi.datasets.openapi.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: ApiKeyAuthHeader
configuration.api_key['ApiKeyAuthHeader'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['ApiKeyAuthHeader'] = 'Bearer'

# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = virus_api.VirusApi(api_client)
    proteins = [
        "spike protein",
    ] # [str] | Which proteins to retrieve in the data package
    refseq_only = True # bool | If true, limit results to RefSeq genomes. (optional) if omitted the server will use the default value of False
    annotated_only = True # bool | If true, limit results to annotated genomes. (optional) if omitted the server will use the default value of False
    released_since = dateutil_parser('2020-08-01') # datetime | If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as '2020-04-01T00:00:00.000Z' (optional)
    updated_since = dateutil_parser('2021-07-18') # datetime |  (optional)
    host = "human" # str | If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default (optional)
    geo_location = "USA" # str | Assemblies from this location (country and state, or continent) (optional)
    complete_only = True # bool | only include complete genomes. (optional) if omitted the server will use the default value of False
    table_fields = [
        V1VirusTableField("["nucleotide_accession","nucleotide_length","nuc_completeness"]"),
    ] # [V1VirusTableField] | Specify which fields to include in the tabular report (optional)
    format = V1TableFormat("tsv") # V1TableFormat | Choose download format (tsv, csv or jsonl) (optional)

    # example passing only required values which don't have defaults set
    try:
        # Get SARS-CoV-2 protein metadata in a tabular format.
        api_response = api_instance.sars2_protein_table(proteins)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling VirusApi->sars2_protein_table: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get SARS-CoV-2 protein metadata in a tabular format.
        api_response = api_instance.sars2_protein_table(proteins, refseq_only=refseq_only, annotated_only=annotated_only, released_since=released_since, updated_since=updated_since, host=host, geo_location=geo_location, complete_only=complete_only, table_fields=table_fields, format=format)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling VirusApi->sars2_protein_table: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **proteins** | **[str]**| Which proteins to retrieve in the data package |
 **refseq_only** | **bool**| If true, limit results to RefSeq genomes. | [optional] if omitted the server will use the default value of False
 **annotated_only** | **bool**| If true, limit results to annotated genomes. | [optional] if omitted the server will use the default value of False
 **released_since** | **datetime**| If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as &#39;2020-04-01T00:00:00.000Z&#39; | [optional]
 **updated_since** | **datetime**|  | [optional]
 **host** | **str**| If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default | [optional]
 **geo_location** | **str**| Assemblies from this location (country and state, or continent) | [optional]
 **complete_only** | **bool**| only include complete genomes. | [optional] if omitted the server will use the default value of False
 **table_fields** | [**[V1VirusTableField]**](V1VirusTableField.md)| Specify which fields to include in the tabular report | [optional]
 **format** | **V1TableFormat**| Choose download format (tsv, csv or jsonl) | [optional]

### Return type

[**V1TabularOutput**](V1TabularOutput.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **virus_genome_download**
> file_type virus_genome_download(taxon)

Download a coronavirus genome dataset by taxon

Download a coronavirus genome dataset by taxon

### Example

* Api Key Authentication (ApiKeyAuthHeader):

```python
import time
import ncbi.datasets.openapi
from ncbi.datasets.openapi.api import virus_api
from ncbi.datasets.openapi.model.v1_annotation_for_virus_type import V1AnnotationForVirusType
from ncbi.datasets.openapi.model.rpc_status import RpcStatus
from pprint import pprint
# Defining the host is optional and defaults to https://api.ncbi.nlm.nih.gov/datasets/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = ncbi.datasets.openapi.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: ApiKeyAuthHeader
configuration.api_key['ApiKeyAuthHeader'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['ApiKeyAuthHeader'] = 'Bearer'

# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = virus_api.VirusApi(api_client)
    taxon = "2697049" # str | NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank
    refseq_only = True # bool | If true, limit results to RefSeq genomes. (optional) if omitted the server will use the default value of False
    annotated_only = True # bool | If true, limit results to annotated genomes. (optional) if omitted the server will use the default value of False
    released_since = dateutil_parser('2020-08-01') # datetime | If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as '2020-04-01T00:00:00.000Z' (optional)
    updated_since = dateutil_parser('2021-07-18') # datetime |  (optional)
    host = "human" # str | If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default (optional)
    pangolin_classification = "pangolin_classification_example" # str | If set, limit results to genomes classified to this lineage by the PangoLearn tool. (optional)
    geo_location = "USA" # str | Assemblies from this location (country and state, or continent) (optional)
    complete_only = True # bool | only include complete genomes. (optional) if omitted the server will use the default value of False
    exclude_sequence = True # bool | Set to true to omit the genomic sequence. (optional) if omitted the server will use the default value of False
    include_annotation_type = [
        V1AnnotationForVirusType("CDS_FASTA"),
    ] # [V1AnnotationForVirusType] | Select additional types of annotation to include in the data package.  If unset, no annotation is provided. (optional)
    filename = "ncbi_dataset.zip" # str | Output file name. (optional) if omitted the server will use the default value of "ncbi_dataset.zip"

    # example passing only required values which don't have defaults set
    try:
        # Download a coronavirus genome dataset by taxon
        api_response = api_instance.virus_genome_download(taxon)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling VirusApi->virus_genome_download: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Download a coronavirus genome dataset by taxon
        api_response = api_instance.virus_genome_download(taxon, refseq_only=refseq_only, annotated_only=annotated_only, released_since=released_since, updated_since=updated_since, host=host, pangolin_classification=pangolin_classification, geo_location=geo_location, complete_only=complete_only, exclude_sequence=exclude_sequence, include_annotation_type=include_annotation_type, filename=filename)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling VirusApi->virus_genome_download: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taxon** | **str**| NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank |
 **refseq_only** | **bool**| If true, limit results to RefSeq genomes. | [optional] if omitted the server will use the default value of False
 **annotated_only** | **bool**| If true, limit results to annotated genomes. | [optional] if omitted the server will use the default value of False
 **released_since** | **datetime**| If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as &#39;2020-04-01T00:00:00.000Z&#39; | [optional]
 **updated_since** | **datetime**|  | [optional]
 **host** | **str**| If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default | [optional]
 **pangolin_classification** | **str**| If set, limit results to genomes classified to this lineage by the PangoLearn tool. | [optional]
 **geo_location** | **str**| Assemblies from this location (country and state, or continent) | [optional]
 **complete_only** | **bool**| only include complete genomes. | [optional] if omitted the server will use the default value of False
 **exclude_sequence** | **bool**| Set to true to omit the genomic sequence. | [optional] if omitted the server will use the default value of False
 **include_annotation_type** | [**[V1AnnotationForVirusType]**](V1AnnotationForVirusType.md)| Select additional types of annotation to include in the data package.  If unset, no annotation is provided. | [optional]
 **filename** | **str**| Output file name. | [optional] if omitted the server will use the default value of "ncbi_dataset.zip"

### Return type

**file_type**

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/zip, application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Download selected genome assemblies and associated annotation data as a zip file |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **virus_genome_download_accession**
> file_type virus_genome_download_accession(accessions)

Download a coronavirus genome dataset by accession

Download a coronavirus genome dataset by accession

### Example

* Api Key Authentication (ApiKeyAuthHeader):

```python
import time
import ncbi.datasets.openapi
from ncbi.datasets.openapi.api import virus_api
from ncbi.datasets.openapi.model.v1_annotation_for_virus_type import V1AnnotationForVirusType
from ncbi.datasets.openapi.model.rpc_status import RpcStatus
from pprint import pprint
# Defining the host is optional and defaults to https://api.ncbi.nlm.nih.gov/datasets/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = ncbi.datasets.openapi.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: ApiKeyAuthHeader
configuration.api_key['ApiKeyAuthHeader'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['ApiKeyAuthHeader'] = 'Bearer'

# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = virus_api.VirusApi(api_client)
    accessions = [
        "NC_045512.2",
    ] # [str] | Accessions accessions = 16;
    refseq_only = True # bool | If true, limit results to RefSeq genomes. (optional) if omitted the server will use the default value of False
    annotated_only = True # bool | If true, limit results to annotated genomes. (optional) if omitted the server will use the default value of False
    released_since = dateutil_parser('2020-08-01') # datetime | If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as '2020-04-01T00:00:00.000Z' (optional)
    updated_since = dateutil_parser('2021-07-18') # datetime |  (optional)
    host = "human" # str | If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default (optional)
    pangolin_classification = "pangolin_classification_example" # str | If set, limit results to genomes classified to this lineage by the PangoLearn tool. (optional)
    geo_location = "USA" # str | Assemblies from this location (country and state, or continent) (optional)
    complete_only = True # bool | only include complete genomes. (optional) if omitted the server will use the default value of False
    exclude_sequence = True # bool | Set to true to omit the genomic sequence. (optional) if omitted the server will use the default value of False
    include_annotation_type = [
        V1AnnotationForVirusType("CDS_FASTA"),
    ] # [V1AnnotationForVirusType] | Select additional types of annotation to include in the data package.  If unset, no annotation is provided. (optional)
    filename = "ncbi_dataset.zip" # str | Output file name. (optional) if omitted the server will use the default value of "ncbi_dataset.zip"

    # example passing only required values which don't have defaults set
    try:
        # Download a coronavirus genome dataset by accession
        api_response = api_instance.virus_genome_download_accession(accessions)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling VirusApi->virus_genome_download_accession: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Download a coronavirus genome dataset by accession
        api_response = api_instance.virus_genome_download_accession(accessions, refseq_only=refseq_only, annotated_only=annotated_only, released_since=released_since, updated_since=updated_since, host=host, pangolin_classification=pangolin_classification, geo_location=geo_location, complete_only=complete_only, exclude_sequence=exclude_sequence, include_annotation_type=include_annotation_type, filename=filename)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling VirusApi->virus_genome_download_accession: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessions** | **[str]**| Accessions accessions &#x3D; 16; |
 **refseq_only** | **bool**| If true, limit results to RefSeq genomes. | [optional] if omitted the server will use the default value of False
 **annotated_only** | **bool**| If true, limit results to annotated genomes. | [optional] if omitted the server will use the default value of False
 **released_since** | **datetime**| If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as &#39;2020-04-01T00:00:00.000Z&#39; | [optional]
 **updated_since** | **datetime**|  | [optional]
 **host** | **str**| If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default | [optional]
 **pangolin_classification** | **str**| If set, limit results to genomes classified to this lineage by the PangoLearn tool. | [optional]
 **geo_location** | **str**| Assemblies from this location (country and state, or continent) | [optional]
 **complete_only** | **bool**| only include complete genomes. | [optional] if omitted the server will use the default value of False
 **exclude_sequence** | **bool**| Set to true to omit the genomic sequence. | [optional] if omitted the server will use the default value of False
 **include_annotation_type** | [**[V1AnnotationForVirusType]**](V1AnnotationForVirusType.md)| Select additional types of annotation to include in the data package.  If unset, no annotation is provided. | [optional]
 **filename** | **str**| Output file name. | [optional] if omitted the server will use the default value of "ncbi_dataset.zip"

### Return type

**file_type**

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/zip, application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Download selected genome assemblies and associated annotation data as a zip file |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **virus_genome_download_post**
> file_type virus_genome_download_post(v1_virus_dataset_request)

Get a coronavirus genome dataset by post

The 'GET' version of download is limited by the size of the GET URL (2KB, which works out to about 140 genomic accessions).  The POST operation is provided to allow users to supply a larger number of accessions in a single request.

### Example

* Api Key Authentication (ApiKeyAuthHeader):

```python
import time
import ncbi.datasets.openapi
from ncbi.datasets.openapi.api import virus_api
from ncbi.datasets.openapi.model.v1_virus_dataset_request import V1VirusDatasetRequest
from ncbi.datasets.openapi.model.rpc_status import RpcStatus
from pprint import pprint
# Defining the host is optional and defaults to https://api.ncbi.nlm.nih.gov/datasets/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = ncbi.datasets.openapi.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: ApiKeyAuthHeader
configuration.api_key['ApiKeyAuthHeader'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['ApiKeyAuthHeader'] = 'Bearer'

# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = virus_api.VirusApi(api_client)
    v1_virus_dataset_request = V1VirusDatasetRequest(
        tax_id=1,
        tax_name="tax_name_example",
        accession="accession_example",
        accessions=[
            "accessions_example",
        ],
        taxon="taxon_example",
        refseq_only=True,
        annotated_only=True,
        released_since=dateutil_parser('1970-01-01T00:00:00.00Z'),
        updated_since=dateutil_parser('1970-01-01T00:00:00.00Z'),
        host="host_example",
        pangolin_classification="pangolin_classification_example",
        geo_location="geo_location_example",
        complete_only=True,
        table_fields=[
            V1VirusTableField("unspecified"),
        ],
        exclude_sequence=True,
        include_annotation_type=[
            V1AnnotationForVirusType("DEFAULT"),
        ],
        format=V1TableFormat("tsv"),
    ) # V1VirusDatasetRequest | 
    filename = "ncbi_dataset.zip" # str | Output file name. (optional) if omitted the server will use the default value of "ncbi_dataset.zip"

    # example passing only required values which don't have defaults set
    try:
        # Get a coronavirus genome dataset by post
        api_response = api_instance.virus_genome_download_post(v1_virus_dataset_request)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling VirusApi->virus_genome_download_post: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get a coronavirus genome dataset by post
        api_response = api_instance.virus_genome_download_post(v1_virus_dataset_request, filename=filename)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling VirusApi->virus_genome_download_post: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1_virus_dataset_request** | [**V1VirusDatasetRequest**](V1VirusDatasetRequest.md)|  |
 **filename** | **str**| Output file name. | [optional] if omitted the server will use the default value of "ncbi_dataset.zip"

### Return type

**file_type**

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/zip, application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Download selected genome assemblies and associated annotation data as a zip file |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **virus_genome_summary**
> V1DownloadSummary virus_genome_summary(taxon)

Get summary data for Coronaviridae genomes by taxon

Get summary data and download by command line instructions for Coronaviridae genomes by taxon.

### Example

* Api Key Authentication (ApiKeyAuthHeader):

```python
import time
import ncbi.datasets.openapi
from ncbi.datasets.openapi.api import virus_api
from ncbi.datasets.openapi.model.v1_annotation_for_virus_type import V1AnnotationForVirusType
from ncbi.datasets.openapi.model.v1_download_summary import V1DownloadSummary
from ncbi.datasets.openapi.model.rpc_status import RpcStatus
from pprint import pprint
# Defining the host is optional and defaults to https://api.ncbi.nlm.nih.gov/datasets/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = ncbi.datasets.openapi.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: ApiKeyAuthHeader
configuration.api_key['ApiKeyAuthHeader'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['ApiKeyAuthHeader'] = 'Bearer'

# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = virus_api.VirusApi(api_client)
    taxon = "2697049" # str | NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank
    refseq_only = True # bool | If true, limit results to RefSeq genomes. (optional) if omitted the server will use the default value of False
    annotated_only = True # bool | If true, limit results to annotated genomes. (optional) if omitted the server will use the default value of False
    released_since = dateutil_parser('2020-08-01') # datetime | If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as '2020-04-01T00:00:00.000Z' (optional)
    updated_since = dateutil_parser('2021-07-18') # datetime |  (optional)
    host = "human" # str | If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default (optional)
    pangolin_classification = "pangolin_classification_example" # str | If set, limit results to genomes classified to this lineage by the PangoLearn tool. (optional)
    geo_location = "USA" # str | Assemblies from this location (country and state, or continent) (optional)
    complete_only = True # bool | only include complete genomes. (optional) if omitted the server will use the default value of False
    exclude_sequence = True # bool | Set to true to omit the genomic sequence. (optional) if omitted the server will use the default value of False
    include_annotation_type = [
        V1AnnotationForVirusType("CDS_FASTA"),
    ] # [V1AnnotationForVirusType] | Select additional types of annotation to include in the data package.  If unset, no annotation is provided. (optional)

    # example passing only required values which don't have defaults set
    try:
        # Get summary data for Coronaviridae genomes by taxon
        api_response = api_instance.virus_genome_summary(taxon)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling VirusApi->virus_genome_summary: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get summary data for Coronaviridae genomes by taxon
        api_response = api_instance.virus_genome_summary(taxon, refseq_only=refseq_only, annotated_only=annotated_only, released_since=released_since, updated_since=updated_since, host=host, pangolin_classification=pangolin_classification, geo_location=geo_location, complete_only=complete_only, exclude_sequence=exclude_sequence, include_annotation_type=include_annotation_type)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling VirusApi->virus_genome_summary: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taxon** | **str**| NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank |
 **refseq_only** | **bool**| If true, limit results to RefSeq genomes. | [optional] if omitted the server will use the default value of False
 **annotated_only** | **bool**| If true, limit results to annotated genomes. | [optional] if omitted the server will use the default value of False
 **released_since** | **datetime**| If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as &#39;2020-04-01T00:00:00.000Z&#39; | [optional]
 **updated_since** | **datetime**|  | [optional]
 **host** | **str**| If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default | [optional]
 **pangolin_classification** | **str**| If set, limit results to genomes classified to this lineage by the PangoLearn tool. | [optional]
 **geo_location** | **str**| Assemblies from this location (country and state, or continent) | [optional]
 **complete_only** | **bool**| only include complete genomes. | [optional] if omitted the server will use the default value of False
 **exclude_sequence** | **bool**| Set to true to omit the genomic sequence. | [optional] if omitted the server will use the default value of False
 **include_annotation_type** | [**[V1AnnotationForVirusType]**](V1AnnotationForVirusType.md)| Select additional types of annotation to include in the data package.  If unset, no annotation is provided. | [optional]

### Return type

[**V1DownloadSummary**](V1DownloadSummary.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **virus_genome_summary_accession**
> V1DownloadSummary virus_genome_summary_accession(accessions)

Get summary data for Coronaviridae genomes by accession

Get summary data for Coronaviridae genomes by accession

### Example

* Api Key Authentication (ApiKeyAuthHeader):

```python
import time
import ncbi.datasets.openapi
from ncbi.datasets.openapi.api import virus_api
from ncbi.datasets.openapi.model.v1_annotation_for_virus_type import V1AnnotationForVirusType
from ncbi.datasets.openapi.model.v1_download_summary import V1DownloadSummary
from ncbi.datasets.openapi.model.rpc_status import RpcStatus
from pprint import pprint
# Defining the host is optional and defaults to https://api.ncbi.nlm.nih.gov/datasets/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = ncbi.datasets.openapi.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: ApiKeyAuthHeader
configuration.api_key['ApiKeyAuthHeader'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['ApiKeyAuthHeader'] = 'Bearer'

# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = virus_api.VirusApi(api_client)
    accessions = [
        "NC_045512.2",
    ] # [str] | Accessions accessions = 16;
    refseq_only = True # bool | If true, limit results to RefSeq genomes. (optional) if omitted the server will use the default value of False
    annotated_only = True # bool | If true, limit results to annotated genomes. (optional) if omitted the server will use the default value of False
    released_since = dateutil_parser('2020-08-01') # datetime | If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as '2020-04-01T00:00:00.000Z' (optional)
    updated_since = dateutil_parser('2021-07-18') # datetime |  (optional)
    host = "human" # str | If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default (optional)
    pangolin_classification = "pangolin_classification_example" # str | If set, limit results to genomes classified to this lineage by the PangoLearn tool. (optional)
    geo_location = "USA" # str | Assemblies from this location (country and state, or continent) (optional)
    complete_only = True # bool | only include complete genomes. (optional) if omitted the server will use the default value of False
    exclude_sequence = True # bool | Set to true to omit the genomic sequence. (optional) if omitted the server will use the default value of False
    include_annotation_type = [
        V1AnnotationForVirusType("CDS_FASTA"),
    ] # [V1AnnotationForVirusType] | Select additional types of annotation to include in the data package.  If unset, no annotation is provided. (optional)

    # example passing only required values which don't have defaults set
    try:
        # Get summary data for Coronaviridae genomes by accession
        api_response = api_instance.virus_genome_summary_accession(accessions)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling VirusApi->virus_genome_summary_accession: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get summary data for Coronaviridae genomes by accession
        api_response = api_instance.virus_genome_summary_accession(accessions, refseq_only=refseq_only, annotated_only=annotated_only, released_since=released_since, updated_since=updated_since, host=host, pangolin_classification=pangolin_classification, geo_location=geo_location, complete_only=complete_only, exclude_sequence=exclude_sequence, include_annotation_type=include_annotation_type)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling VirusApi->virus_genome_summary_accession: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessions** | **[str]**| Accessions accessions &#x3D; 16; |
 **refseq_only** | **bool**| If true, limit results to RefSeq genomes. | [optional] if omitted the server will use the default value of False
 **annotated_only** | **bool**| If true, limit results to annotated genomes. | [optional] if omitted the server will use the default value of False
 **released_since** | **datetime**| If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as &#39;2020-04-01T00:00:00.000Z&#39; | [optional]
 **updated_since** | **datetime**|  | [optional]
 **host** | **str**| If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default | [optional]
 **pangolin_classification** | **str**| If set, limit results to genomes classified to this lineage by the PangoLearn tool. | [optional]
 **geo_location** | **str**| Assemblies from this location (country and state, or continent) | [optional]
 **complete_only** | **bool**| only include complete genomes. | [optional] if omitted the server will use the default value of False
 **exclude_sequence** | **bool**| Set to true to omit the genomic sequence. | [optional] if omitted the server will use the default value of False
 **include_annotation_type** | [**[V1AnnotationForVirusType]**](V1AnnotationForVirusType.md)| Select additional types of annotation to include in the data package.  If unset, no annotation is provided. | [optional]

### Return type

[**V1DownloadSummary**](V1DownloadSummary.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **virus_genome_summary_post**
> V1DownloadSummary virus_genome_summary_post(v1_virus_dataset_request)

Get summary data for Coronaviridae genomes by post

The 'GET' version is limited by the size of the GET URL (2KB, which works out to about 140 genomic accessions).  The POST operation is provided to allow users to supply a larger number of accessions in a single request.

### Example

* Api Key Authentication (ApiKeyAuthHeader):

```python
import time
import ncbi.datasets.openapi
from ncbi.datasets.openapi.api import virus_api
from ncbi.datasets.openapi.model.v1_download_summary import V1DownloadSummary
from ncbi.datasets.openapi.model.v1_virus_dataset_request import V1VirusDatasetRequest
from ncbi.datasets.openapi.model.rpc_status import RpcStatus
from pprint import pprint
# Defining the host is optional and defaults to https://api.ncbi.nlm.nih.gov/datasets/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = ncbi.datasets.openapi.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: ApiKeyAuthHeader
configuration.api_key['ApiKeyAuthHeader'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['ApiKeyAuthHeader'] = 'Bearer'

# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = virus_api.VirusApi(api_client)
    v1_virus_dataset_request = V1VirusDatasetRequest(
        tax_id=1,
        tax_name="tax_name_example",
        accession="accession_example",
        accessions=[
            "accessions_example",
        ],
        taxon="taxon_example",
        refseq_only=True,
        annotated_only=True,
        released_since=dateutil_parser('1970-01-01T00:00:00.00Z'),
        updated_since=dateutil_parser('1970-01-01T00:00:00.00Z'),
        host="host_example",
        pangolin_classification="pangolin_classification_example",
        geo_location="geo_location_example",
        complete_only=True,
        table_fields=[
            V1VirusTableField("unspecified"),
        ],
        exclude_sequence=True,
        include_annotation_type=[
            V1AnnotationForVirusType("DEFAULT"),
        ],
        format=V1TableFormat("tsv"),
    ) # V1VirusDatasetRequest | 

    # example passing only required values which don't have defaults set
    try:
        # Get summary data for Coronaviridae genomes by post
        api_response = api_instance.virus_genome_summary_post(v1_virus_dataset_request)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling VirusApi->virus_genome_summary_post: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1_virus_dataset_request** | [**V1VirusDatasetRequest**](V1VirusDatasetRequest.md)|  |

### Return type

[**V1DownloadSummary**](V1DownloadSummary.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **virus_genome_table**
> V1TabularOutput virus_genome_table(taxon)

Get viral genomic metadata in a tabular format.

Get viral genomic metadata in tabular format for Coronaviridae genomes by taxon.

### Example

* Api Key Authentication (ApiKeyAuthHeader):

```python
import time
import ncbi.datasets.openapi
from ncbi.datasets.openapi.api import virus_api
from ncbi.datasets.openapi.model.v1_virus_table_field import V1VirusTableField
from ncbi.datasets.openapi.model.v1_tabular_output import V1TabularOutput
from ncbi.datasets.openapi.model.v1_table_format import V1TableFormat
from ncbi.datasets.openapi.model.rpc_status import RpcStatus
from pprint import pprint
# Defining the host is optional and defaults to https://api.ncbi.nlm.nih.gov/datasets/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = ncbi.datasets.openapi.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: ApiKeyAuthHeader
configuration.api_key['ApiKeyAuthHeader'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['ApiKeyAuthHeader'] = 'Bearer'

# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = virus_api.VirusApi(api_client)
    taxon = "2697049" # str | NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank
    refseq_only = True # bool | If true, limit results to RefSeq genomes. (optional) if omitted the server will use the default value of False
    annotated_only = True # bool | If true, limit results to annotated genomes. (optional) if omitted the server will use the default value of False
    released_since = dateutil_parser('2020-08-01') # datetime | If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as '2020-04-01T00:00:00.000Z' (optional)
    updated_since = dateutil_parser('2021-07-18') # datetime |  (optional)
    host = "human" # str | If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default (optional)
    pangolin_classification = "pangolin_classification_example" # str | If set, limit results to genomes classified to this lineage by the PangoLearn tool. (optional)
    geo_location = "USA" # str | Assemblies from this location (country and state, or continent) (optional)
    complete_only = True # bool | only include complete genomes. (optional) if omitted the server will use the default value of False
    table_fields = [
        V1VirusTableField("["nucleotide_accession","nucleotide_length","nuc_completeness"]"),
    ] # [V1VirusTableField] | Specify which fields to include in the tabular report (optional)
    format = V1TableFormat("tsv") # V1TableFormat | Choose download format (tsv, csv or jsonl) (optional)

    # example passing only required values which don't have defaults set
    try:
        # Get viral genomic metadata in a tabular format.
        api_response = api_instance.virus_genome_table(taxon)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling VirusApi->virus_genome_table: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get viral genomic metadata in a tabular format.
        api_response = api_instance.virus_genome_table(taxon, refseq_only=refseq_only, annotated_only=annotated_only, released_since=released_since, updated_since=updated_since, host=host, pangolin_classification=pangolin_classification, geo_location=geo_location, complete_only=complete_only, table_fields=table_fields, format=format)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling VirusApi->virus_genome_table: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taxon** | **str**| NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank |
 **refseq_only** | **bool**| If true, limit results to RefSeq genomes. | [optional] if omitted the server will use the default value of False
 **annotated_only** | **bool**| If true, limit results to annotated genomes. | [optional] if omitted the server will use the default value of False
 **released_since** | **datetime**| If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as &#39;2020-04-01T00:00:00.000Z&#39; | [optional]
 **updated_since** | **datetime**|  | [optional]
 **host** | **str**| If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default | [optional]
 **pangolin_classification** | **str**| If set, limit results to genomes classified to this lineage by the PangoLearn tool. | [optional]
 **geo_location** | **str**| Assemblies from this location (country and state, or continent) | [optional]
 **complete_only** | **bool**| only include complete genomes. | [optional] if omitted the server will use the default value of False
 **table_fields** | [**[V1VirusTableField]**](V1VirusTableField.md)| Specify which fields to include in the tabular report | [optional]
 **format** | **V1TableFormat**| Choose download format (tsv, csv or jsonl) | [optional]

### Return type

[**V1TabularOutput**](V1TabularOutput.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **virus_reports_by_acessions**
> V1reportsVirusDataReportPage virus_reports_by_acessions(accessions)

Get virus metadata by accession

Get virus metadata by accesion. By default, in paged JSON format, but also available as tabular (accept: x-tabular) or json-lines (accept: x-jsonlines)

### Example

* Api Key Authentication (ApiKeyAuthHeader):

```python
import time
import ncbi.datasets.openapi
from ncbi.datasets.openapi.api import virus_api
from ncbi.datasets.openapi.model.v1_virus_data_report_request_content_type import V1VirusDataReportRequestContentType
from ncbi.datasets.openapi.model.rpc_status import RpcStatus
from ncbi.datasets.openapi.model.v1reports_virus_data_report_page import V1reportsVirusDataReportPage
from pprint import pprint
# Defining the host is optional and defaults to https://api.ncbi.nlm.nih.gov/datasets/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = ncbi.datasets.openapi.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: ApiKeyAuthHeader
configuration.api_key['ApiKeyAuthHeader'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['ApiKeyAuthHeader'] = 'Bearer'

# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = virus_api.VirusApi(api_client)
    accessions = [
        "NC_045512.2,NC_038294.1",
    ] # [str] | genome sequence accessions
    filter_refseq_only = True # bool | If true, limit results to RefSeq genomes. (optional) if omitted the server will use the default value of False
    filter_annotated_only = True # bool | If true, limit results to annotated genomes. (optional) if omitted the server will use the default value of False
    filter_released_since = dateutil_parser('2020-08-01') # datetime | If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as '2020-04-01T00:00:00.000Z' (optional)
    filter_updated_since = dateutil_parser('2021-07-18') # datetime |  (optional)
    filter_host = "human" # str | If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default (optional)
    filter_pangolin_classification = "filter.pangolin_classification_example" # str | If set, limit results to genomes classified to this lineage by the PangoLearn tool. (optional)
    filter_geo_location = "USA" # str | Assemblies from this location (country and state, or continent) (optional)
    filter_complete_only = True # bool | only include complete genomes. (optional) if omitted the server will use the default value of False
    returned_content = V1VirusDataReportRequestContentType("COMPLETE") # V1VirusDataReportRequestContentType | Return either virus genome accessions, or complete virus metadata (optional)
    table_fields = ["accession","is_complete","is_annotated"] # [str] | Specify which fields to include in the tabular report (optional)
    page_size = 20 # int | The maximum number of virus data reports to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, `page_token` can be used to retrieve the remaining results. (optional) if omitted the server will use the default value of 20
    page_token = "page_token_example" # str | A page token is returned from a `GetVirusDataReports` call with more than `page_size` results. Use this token, along with the previous `VirusDataReportRequest` parameters, to retrieve the next page of results. When `page_token` is empty, all results have been retrieved. (optional)

    # example passing only required values which don't have defaults set
    try:
        # Get virus metadata by accession
        api_response = api_instance.virus_reports_by_acessions(accessions)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling VirusApi->virus_reports_by_acessions: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get virus metadata by accession
        api_response = api_instance.virus_reports_by_acessions(accessions, filter_refseq_only=filter_refseq_only, filter_annotated_only=filter_annotated_only, filter_released_since=filter_released_since, filter_updated_since=filter_updated_since, filter_host=filter_host, filter_pangolin_classification=filter_pangolin_classification, filter_geo_location=filter_geo_location, filter_complete_only=filter_complete_only, returned_content=returned_content, table_fields=table_fields, page_size=page_size, page_token=page_token)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling VirusApi->virus_reports_by_acessions: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessions** | **[str]**| genome sequence accessions |
 **filter_refseq_only** | **bool**| If true, limit results to RefSeq genomes. | [optional] if omitted the server will use the default value of False
 **filter_annotated_only** | **bool**| If true, limit results to annotated genomes. | [optional] if omitted the server will use the default value of False
 **filter_released_since** | **datetime**| If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as &#39;2020-04-01T00:00:00.000Z&#39; | [optional]
 **filter_updated_since** | **datetime**|  | [optional]
 **filter_host** | **str**| If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default | [optional]
 **filter_pangolin_classification** | **str**| If set, limit results to genomes classified to this lineage by the PangoLearn tool. | [optional]
 **filter_geo_location** | **str**| Assemblies from this location (country and state, or continent) | [optional]
 **filter_complete_only** | **bool**| only include complete genomes. | [optional] if omitted the server will use the default value of False
 **returned_content** | **V1VirusDataReportRequestContentType**| Return either virus genome accessions, or complete virus metadata | [optional]
 **table_fields** | **[str]**| Specify which fields to include in the tabular report | [optional]
 **page_size** | **int**| The maximum number of virus data reports to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, &#x60;page_token&#x60; can be used to retrieve the remaining results. | [optional] if omitted the server will use the default value of 20
 **page_token** | **str**| A page token is returned from a &#x60;GetVirusDataReports&#x60; call with more than &#x60;page_size&#x60; results. Use this token, along with the previous &#x60;VirusDataReportRequest&#x60; parameters, to retrieve the next page of results. When &#x60;page_token&#x60; is empty, all results have been retrieved. | [optional]

### Return type

[**V1reportsVirusDataReportPage**](V1reportsVirusDataReportPage.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/tab-separated-values, application/x-tabular


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **virus_reports_by_post**
> V1reportsVirusDataReportPage virus_reports_by_post(v1_virus_data_report_request)

Get virus metadata by POST

Get virus metadata. By default, in paged JSON format, but also available as tabular (accept: x-tabular) or json-lines (accept: x-jsonlines)

### Example

* Api Key Authentication (ApiKeyAuthHeader):

```python
import time
import ncbi.datasets.openapi
from ncbi.datasets.openapi.api import virus_api
from ncbi.datasets.openapi.model.v1_virus_data_report_request import V1VirusDataReportRequest
from ncbi.datasets.openapi.model.rpc_status import RpcStatus
from ncbi.datasets.openapi.model.v1reports_virus_data_report_page import V1reportsVirusDataReportPage
from pprint import pprint
# Defining the host is optional and defaults to https://api.ncbi.nlm.nih.gov/datasets/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = ncbi.datasets.openapi.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: ApiKeyAuthHeader
configuration.api_key['ApiKeyAuthHeader'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['ApiKeyAuthHeader'] = 'Bearer'

# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = virus_api.VirusApi(api_client)
    v1_virus_data_report_request = V1VirusDataReportRequest(
        filter=V1VirusDatasetFilter(
            accessions=[
                "accessions_example",
            ],
            taxon="taxon_example",
            refseq_only=True,
            annotated_only=True,
            released_since=dateutil_parser('1970-01-01T00:00:00.00Z'),
            updated_since=dateutil_parser('1970-01-01T00:00:00.00Z'),
            host="host_example",
            pangolin_classification="pangolin_classification_example",
            geo_location="geo_location_example",
            complete_only=True,
        ),
        returned_content=V1VirusDataReportRequestContentType("COMPLETE"),
        table_fields=[
            "table_fields_example",
        ],
        page_size=1,
        page_token="page_token_example",
    ) # V1VirusDataReportRequest | 

    # example passing only required values which don't have defaults set
    try:
        # Get virus metadata by POST
        api_response = api_instance.virus_reports_by_post(v1_virus_data_report_request)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling VirusApi->virus_reports_by_post: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1_virus_data_report_request** | [**V1VirusDataReportRequest**](V1VirusDataReportRequest.md)|  |

### Return type

[**V1reportsVirusDataReportPage**](V1reportsVirusDataReportPage.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/tab-separated-values, application/x-tabular


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **virus_reports_by_taxon**
> V1reportsVirusDataReportPage virus_reports_by_taxon(taxon)

Get virus metadata by taxon

Get virus metadata by taxon. By default, in paged JSON format, but also available as tabular (accept: x-tabular) or json-lines (accept: x-jsonlines)

### Example

* Api Key Authentication (ApiKeyAuthHeader):

```python
import time
import ncbi.datasets.openapi
from ncbi.datasets.openapi.api import virus_api
from ncbi.datasets.openapi.model.v1_virus_data_report_request_content_type import V1VirusDataReportRequestContentType
from ncbi.datasets.openapi.model.rpc_status import RpcStatus
from ncbi.datasets.openapi.model.v1reports_virus_data_report_page import V1reportsVirusDataReportPage
from pprint import pprint
# Defining the host is optional and defaults to https://api.ncbi.nlm.nih.gov/datasets/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = ncbi.datasets.openapi.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: ApiKeyAuthHeader
configuration.api_key['ApiKeyAuthHeader'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['ApiKeyAuthHeader'] = 'Bearer'

# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = virus_api.VirusApi(api_client)
    taxon = "2697049" # str | NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank
    filter_refseq_only = True # bool | If true, limit results to RefSeq genomes. (optional) if omitted the server will use the default value of False
    filter_annotated_only = True # bool | If true, limit results to annotated genomes. (optional) if omitted the server will use the default value of False
    filter_released_since = dateutil_parser('2020-08-01') # datetime | If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as '2020-04-01T00:00:00.000Z' (optional)
    filter_updated_since = dateutil_parser('2021-07-18') # datetime |  (optional)
    filter_host = "human" # str | If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default (optional)
    filter_pangolin_classification = "filter.pangolin_classification_example" # str | If set, limit results to genomes classified to this lineage by the PangoLearn tool. (optional)
    filter_geo_location = "USA" # str | Assemblies from this location (country and state, or continent) (optional)
    filter_complete_only = True # bool | only include complete genomes. (optional) if omitted the server will use the default value of False
    returned_content = V1VirusDataReportRequestContentType("COMPLETE") # V1VirusDataReportRequestContentType | Return either virus genome accessions, or complete virus metadata (optional)
    table_fields = ["accession","is_complete","is_annotated"] # [str] | Specify which fields to include in the tabular report (optional)
    page_size = 20 # int | The maximum number of virus data reports to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, `page_token` can be used to retrieve the remaining results. (optional) if omitted the server will use the default value of 20
    page_token = "page_token_example" # str | A page token is returned from a `GetVirusDataReports` call with more than `page_size` results. Use this token, along with the previous `VirusDataReportRequest` parameters, to retrieve the next page of results. When `page_token` is empty, all results have been retrieved. (optional)

    # example passing only required values which don't have defaults set
    try:
        # Get virus metadata by taxon
        api_response = api_instance.virus_reports_by_taxon(taxon)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling VirusApi->virus_reports_by_taxon: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get virus metadata by taxon
        api_response = api_instance.virus_reports_by_taxon(taxon, filter_refseq_only=filter_refseq_only, filter_annotated_only=filter_annotated_only, filter_released_since=filter_released_since, filter_updated_since=filter_updated_since, filter_host=filter_host, filter_pangolin_classification=filter_pangolin_classification, filter_geo_location=filter_geo_location, filter_complete_only=filter_complete_only, returned_content=returned_content, table_fields=table_fields, page_size=page_size, page_token=page_token)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling VirusApi->virus_reports_by_taxon: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taxon** | **str**| NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank |
 **filter_refseq_only** | **bool**| If true, limit results to RefSeq genomes. | [optional] if omitted the server will use the default value of False
 **filter_annotated_only** | **bool**| If true, limit results to annotated genomes. | [optional] if omitted the server will use the default value of False
 **filter_released_since** | **datetime**| If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as &#39;2020-04-01T00:00:00.000Z&#39; | [optional]
 **filter_updated_since** | **datetime**|  | [optional]
 **filter_host** | **str**| If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default | [optional]
 **filter_pangolin_classification** | **str**| If set, limit results to genomes classified to this lineage by the PangoLearn tool. | [optional]
 **filter_geo_location** | **str**| Assemblies from this location (country and state, or continent) | [optional]
 **filter_complete_only** | **bool**| only include complete genomes. | [optional] if omitted the server will use the default value of False
 **returned_content** | **V1VirusDataReportRequestContentType**| Return either virus genome accessions, or complete virus metadata | [optional]
 **table_fields** | **[str]**| Specify which fields to include in the tabular report | [optional]
 **page_size** | **int**| The maximum number of virus data reports to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, &#x60;page_token&#x60; can be used to retrieve the remaining results. | [optional] if omitted the server will use the default value of 20
 **page_token** | **str**| A page token is returned from a &#x60;GetVirusDataReports&#x60; call with more than &#x60;page_size&#x60; results. Use this token, along with the previous &#x60;VirusDataReportRequest&#x60; parameters, to retrieve the next page of results. When &#x60;page_token&#x60; is empty, all results have been retrieved. | [optional]

### Return type

[**V1reportsVirusDataReportPage**](V1reportsVirusDataReportPage.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/tab-separated-values, application/x-tabular


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

