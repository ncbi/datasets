# ncbi.datasets.openapi.GeneApi

All URIs are relative to *https://api.ncbi.nlm.nih.gov/datasets/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**download_gene_package**](GeneApi.md#download_gene_package) | **GET** /gene/id/{gene_ids}/download | Get a gene dataset by gene ID
[**download_gene_package_post**](GeneApi.md#download_gene_package_post) | **POST** /gene/download | Get a gene dataset by POST
[**gene_download_summary_by_accession**](GeneApi.md#gene_download_summary_by_accession) | **GET** /gene/accession/{accessions}/download_summary | Get gene download summary by RefSeq Accession
[**gene_download_summary_by_id**](GeneApi.md#gene_download_summary_by_id) | **GET** /gene/id/{gene_ids}/download_summary | Get gene download summary by GeneID
[**gene_download_summary_by_post**](GeneApi.md#gene_download_summary_by_post) | **POST** /gene/download_summary | Get gene download summary
[**gene_download_summary_by_tax_and_symbol**](GeneApi.md#gene_download_summary_by_tax_and_symbol) | **GET** /gene/symbol/{symbols}/taxon/{taxon}/download_summary | Get gene download summary by gene symbol
[**gene_metadata_by_accession**](GeneApi.md#gene_metadata_by_accession) | **GET** /gene/accession/{accessions} | Get gene metadata by RefSeq Accession
[**gene_metadata_by_id**](GeneApi.md#gene_metadata_by_id) | **GET** /gene/id/{gene_ids} | Get gene metadata by GeneID
[**gene_metadata_by_post**](GeneApi.md#gene_metadata_by_post) | **POST** /gene | Get gene metadata as JSON
[**gene_metadata_by_tax_and_symbol**](GeneApi.md#gene_metadata_by_tax_and_symbol) | **GET** /gene/symbol/{symbols}/taxon/{taxon} | Get gene metadata by gene symbol
[**gene_metadata_stream_by_post**](GeneApi.md#gene_metadata_stream_by_post) | **POST** /gene/stream | Get gene metadata
[**gene_orthologs_by_id**](GeneApi.md#gene_orthologs_by_id) | **GET** /gene/id/{gene_id}/orthologs | Get gene orthologs by gene ID
[**gene_tax_name_query**](GeneApi.md#gene_tax_name_query) | **GET** /gene/taxon_suggest/{taxon_query} | Get a list of taxonomy names and IDs found in the gene dataset given a partial taxonomic name
[**gene_tax_tree**](GeneApi.md#gene_tax_tree) | **GET** /gene/taxon/{taxon}/tree | Get a taxonomic subtree by taxonomic identifier


# **download_gene_package**
> file_type download_gene_package(gene_ids)

Get a gene dataset by gene ID

Get a gene dataset including gene, transcript and protein fasta sequence, annotation and metadata by gene ID.

### Example

* Api Key Authentication (ApiKeyAuthHeader):

```python
import time
import ncbi.datasets.openapi
from ncbi.datasets.openapi.api import gene_api
from ncbi.datasets.openapi.model.v1_fasta import V1Fasta
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
    api_instance = gene_api.GeneApi(api_client)
    gene_ids = [
        59067,
    ] # [int] | NCBI gene ids
    include_annotation_type = [
        V1Fasta("FASTA_UNSPECIFIED"),
    ] # [V1Fasta] | Select additional types of annotation to include in the data package.  If unset, no annotation is provided. (optional)
    fasta_filter = [
        "fasta_filter_example",
    ] # [str] | Limit the FASTA sequences in the datasets package to these transcript and protein accessions (optional)
    filename = "ncbi_dataset.zip" # str | Output file name. (optional) if omitted the server will use the default value of "ncbi_dataset.zip"

    # example passing only required values which don't have defaults set
    try:
        # Get a gene dataset by gene ID
        api_response = api_instance.download_gene_package(gene_ids)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling GeneApi->download_gene_package: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get a gene dataset by gene ID
        api_response = api_instance.download_gene_package(gene_ids, include_annotation_type=include_annotation_type, fasta_filter=fasta_filter, filename=filename)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling GeneApi->download_gene_package: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gene_ids** | **[int]**| NCBI gene ids |
 **include_annotation_type** | [**[V1Fasta]**](V1Fasta.md)| Select additional types of annotation to include in the data package.  If unset, no annotation is provided. | [optional]
 **fasta_filter** | **[str]**| Limit the FASTA sequences in the datasets package to these transcript and protein accessions | [optional]
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

# **download_gene_package_post**
> file_type download_gene_package_post(v1_gene_dataset_request)

Get a gene dataset by POST

Get a gene dataset including gene, transcript and protein fasta sequence, annotation and metadata by POST.

### Example

* Api Key Authentication (ApiKeyAuthHeader):

```python
import time
import ncbi.datasets.openapi
from ncbi.datasets.openapi.api import gene_api
from ncbi.datasets.openapi.model.v1_gene_dataset_request import V1GeneDatasetRequest
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
    api_instance = gene_api.GeneApi(api_client)
    v1_gene_dataset_request = V1GeneDatasetRequest(
        gene_ids=[
            1,
        ],
        accessions=[
            "accessions_example",
        ],
        symbols_for_taxon=V1GeneDatasetRequestSymbolsForTaxon(
            symbols=[
                "symbols_example",
            ],
            taxon="taxon_example",
        ),
        taxon="taxon_example",
        include_annotation_type=[
            V1Fasta("FASTA_UNSPECIFIED"),
        ],
        returned_content=V1GeneDatasetRequestContentType("COMPLETE"),
        sort_schema=V1GeneDatasetRequestSort(
            field=V1GeneDatasetRequestSortField("SORT_FIELD_UNSPECIFIED"),
            direction=V1SortDirection("SORT_DIRECTION_UNSPECIFIED"),
        ),
        limit="limit_example",
        fasta_filter=[
            "fasta_filter_example",
        ],
    ) # V1GeneDatasetRequest | 
    filename = "ncbi_dataset.zip" # str | Output file name. (optional) if omitted the server will use the default value of "ncbi_dataset.zip"

    # example passing only required values which don't have defaults set
    try:
        # Get a gene dataset by POST
        api_response = api_instance.download_gene_package_post(v1_gene_dataset_request)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling GeneApi->download_gene_package_post: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get a gene dataset by POST
        api_response = api_instance.download_gene_package_post(v1_gene_dataset_request, filename=filename)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling GeneApi->download_gene_package_post: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1_gene_dataset_request** | [**V1GeneDatasetRequest**](V1GeneDatasetRequest.md)|  |
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

# **gene_download_summary_by_accession**
> V1DownloadSummary gene_download_summary_by_accession(accessions)

Get gene download summary by RefSeq Accession

Get gene download summary by RefSeq Accession in a JSON output format.

### Example

* Api Key Authentication (ApiKeyAuthHeader):

```python
import time
import ncbi.datasets.openapi
from ncbi.datasets.openapi.api import gene_api
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
    api_instance = gene_api.GeneApi(api_client)
    accessions = ["NM_021803.4","NM_181078.3"] # [str] | RNA or Protein accessions.
    fasta_filter = [
        "fasta_filter_example",
    ] # [str] | Limit the FASTA sequences in the datasets package to these transcript and protein accessions (optional)

    # example passing only required values which don't have defaults set
    try:
        # Get gene download summary by RefSeq Accession
        api_response = api_instance.gene_download_summary_by_accession(accessions)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling GeneApi->gene_download_summary_by_accession: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get gene download summary by RefSeq Accession
        api_response = api_instance.gene_download_summary_by_accession(accessions, fasta_filter=fasta_filter)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling GeneApi->gene_download_summary_by_accession: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessions** | **[str]**| RNA or Protein accessions. |
 **fasta_filter** | **[str]**| Limit the FASTA sequences in the datasets package to these transcript and protein accessions | [optional]

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

# **gene_download_summary_by_id**
> V1DownloadSummary gene_download_summary_by_id(gene_ids)

Get gene download summary by GeneID

Get a download summary by GeneID in a JSON output format.

### Example

* Api Key Authentication (ApiKeyAuthHeader):

```python
import time
import ncbi.datasets.openapi
from ncbi.datasets.openapi.api import gene_api
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
    api_instance = gene_api.GeneApi(api_client)
    gene_ids = [
        59067,
    ] # [int] | NCBI gene ids
    fasta_filter = [
        "fasta_filter_example",
    ] # [str] | Limit the FASTA sequences in the datasets package to these transcript and protein accessions (optional)

    # example passing only required values which don't have defaults set
    try:
        # Get gene download summary by GeneID
        api_response = api_instance.gene_download_summary_by_id(gene_ids)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling GeneApi->gene_download_summary_by_id: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get gene download summary by GeneID
        api_response = api_instance.gene_download_summary_by_id(gene_ids, fasta_filter=fasta_filter)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling GeneApi->gene_download_summary_by_id: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gene_ids** | **[int]**| NCBI gene ids |
 **fasta_filter** | **[str]**| Limit the FASTA sequences in the datasets package to these transcript and protein accessions | [optional]

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

# **gene_download_summary_by_post**
> V1DownloadSummary gene_download_summary_by_post(v1_gene_dataset_request)

Get gene download summary

Get gene download summary in a JSON output format.

### Example

* Api Key Authentication (ApiKeyAuthHeader):

```python
import time
import ncbi.datasets.openapi
from ncbi.datasets.openapi.api import gene_api
from ncbi.datasets.openapi.model.v1_download_summary import V1DownloadSummary
from ncbi.datasets.openapi.model.v1_gene_dataset_request import V1GeneDatasetRequest
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
    api_instance = gene_api.GeneApi(api_client)
    v1_gene_dataset_request = V1GeneDatasetRequest(
        gene_ids=[
            1,
        ],
        accessions=[
            "accessions_example",
        ],
        symbols_for_taxon=V1GeneDatasetRequestSymbolsForTaxon(
            symbols=[
                "symbols_example",
            ],
            taxon="taxon_example",
        ),
        taxon="taxon_example",
        include_annotation_type=[
            V1Fasta("FASTA_UNSPECIFIED"),
        ],
        returned_content=V1GeneDatasetRequestContentType("COMPLETE"),
        sort_schema=V1GeneDatasetRequestSort(
            field=V1GeneDatasetRequestSortField("SORT_FIELD_UNSPECIFIED"),
            direction=V1SortDirection("SORT_DIRECTION_UNSPECIFIED"),
        ),
        limit="limit_example",
        fasta_filter=[
            "fasta_filter_example",
        ],
    ) # V1GeneDatasetRequest | 

    # example passing only required values which don't have defaults set
    try:
        # Get gene download summary
        api_response = api_instance.gene_download_summary_by_post(v1_gene_dataset_request)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling GeneApi->gene_download_summary_by_post: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1_gene_dataset_request** | [**V1GeneDatasetRequest**](V1GeneDatasetRequest.md)|  |

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

# **gene_download_summary_by_tax_and_symbol**
> V1DownloadSummary gene_download_summary_by_tax_and_symbol(symbols, taxon)

Get gene download summary by gene symbol

Get gene download summary by gene symbol in a JSON output format.

### Example

* Api Key Authentication (ApiKeyAuthHeader):

```python
import time
import ncbi.datasets.openapi
from ncbi.datasets.openapi.api import gene_api
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
    api_instance = gene_api.GeneApi(api_client)
    symbols = [
        "GNAS",
    ] # [str] | Gene symbol
    taxon = "9606" # str | Taxon for provided gene symbol
    fasta_filter = [
        "fasta_filter_example",
    ] # [str] | Limit the FASTA sequences in the datasets package to these transcript and protein accessions (optional)

    # example passing only required values which don't have defaults set
    try:
        # Get gene download summary by gene symbol
        api_response = api_instance.gene_download_summary_by_tax_and_symbol(symbols, taxon)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling GeneApi->gene_download_summary_by_tax_and_symbol: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get gene download summary by gene symbol
        api_response = api_instance.gene_download_summary_by_tax_and_symbol(symbols, taxon, fasta_filter=fasta_filter)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling GeneApi->gene_download_summary_by_tax_and_symbol: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbols** | **[str]**| Gene symbol |
 **taxon** | **str**| Taxon for provided gene symbol |
 **fasta_filter** | **[str]**| Limit the FASTA sequences in the datasets package to these transcript and protein accessions | [optional]

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

# **gene_metadata_by_accession**
> V1GeneMetadata gene_metadata_by_accession(accessions)

Get gene metadata by RefSeq Accession

Get detailed gene metadata by RefSeq Accession in a JSON output format.

### Example

* Api Key Authentication (ApiKeyAuthHeader):

```python
import time
import ncbi.datasets.openapi
from ncbi.datasets.openapi.api import gene_api
from ncbi.datasets.openapi.model.v1_gene_dataset_request_content_type import V1GeneDatasetRequestContentType
from ncbi.datasets.openapi.model.v1_sort_direction import V1SortDirection
from ncbi.datasets.openapi.model.v1_gene_dataset_request_sort_field import V1GeneDatasetRequestSortField
from ncbi.datasets.openapi.model.v1_gene_metadata import V1GeneMetadata
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
    api_instance = gene_api.GeneApi(api_client)
    accessions = ["NM_021803.4","NM_181078.3"] # [str] | RNA or Protein accessions.
    returned_content = V1GeneDatasetRequestContentType("COMPLETE") # V1GeneDatasetRequestContentType | Return either gene-ids, or entire gene metadata (optional)
    sort_schema_field = V1GeneDatasetRequestSortField("SORT_FIELD_UNSPECIFIED") # V1GeneDatasetRequestSortField | Select a field to sort on. (optional)
    sort_schema_direction = V1SortDirection("SORT_DIRECTION_UNSPECIFIED") # V1SortDirection | Select a direction for the sort. (optional)
    limit = "limit_example" # str | Limit the number of returned results (\"all\", \"none\", otherwise an integer value) (optional)

    # example passing only required values which don't have defaults set
    try:
        # Get gene metadata by RefSeq Accession
        api_response = api_instance.gene_metadata_by_accession(accessions)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling GeneApi->gene_metadata_by_accession: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get gene metadata by RefSeq Accession
        api_response = api_instance.gene_metadata_by_accession(accessions, returned_content=returned_content, sort_schema_field=sort_schema_field, sort_schema_direction=sort_schema_direction, limit=limit)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling GeneApi->gene_metadata_by_accession: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessions** | **[str]**| RNA or Protein accessions. |
 **returned_content** | **V1GeneDatasetRequestContentType**| Return either gene-ids, or entire gene metadata | [optional]
 **sort_schema_field** | **V1GeneDatasetRequestSortField**| Select a field to sort on. | [optional]
 **sort_schema_direction** | **V1SortDirection**| Select a direction for the sort. | [optional]
 **limit** | **str**| Limit the number of returned results (\&quot;all\&quot;, \&quot;none\&quot;, otherwise an integer value) | [optional]

### Return type

[**V1GeneMetadata**](V1GeneMetadata.md)

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

# **gene_metadata_by_id**
> V1GeneMetadata gene_metadata_by_id(gene_ids)

Get gene metadata by GeneID

Get detailed gene metadata by GeneID in a JSON output format.

### Example

* Api Key Authentication (ApiKeyAuthHeader):

```python
import time
import ncbi.datasets.openapi
from ncbi.datasets.openapi.api import gene_api
from ncbi.datasets.openapi.model.v1_gene_dataset_request_content_type import V1GeneDatasetRequestContentType
from ncbi.datasets.openapi.model.v1_sort_direction import V1SortDirection
from ncbi.datasets.openapi.model.v1_gene_dataset_request_sort_field import V1GeneDatasetRequestSortField
from ncbi.datasets.openapi.model.v1_gene_metadata import V1GeneMetadata
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
    api_instance = gene_api.GeneApi(api_client)
    gene_ids = [
        59067,
    ] # [int] | NCBI gene ids
    returned_content = V1GeneDatasetRequestContentType("COMPLETE") # V1GeneDatasetRequestContentType | Return either gene-ids, or entire gene metadata (optional)
    sort_schema_field = V1GeneDatasetRequestSortField("SORT_FIELD_UNSPECIFIED") # V1GeneDatasetRequestSortField | Select a field to sort on. (optional)
    sort_schema_direction = V1SortDirection("SORT_DIRECTION_UNSPECIFIED") # V1SortDirection | Select a direction for the sort. (optional)
    limit = "limit_example" # str | Limit the number of returned results (\"all\", \"none\", otherwise an integer value) (optional)

    # example passing only required values which don't have defaults set
    try:
        # Get gene metadata by GeneID
        api_response = api_instance.gene_metadata_by_id(gene_ids)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling GeneApi->gene_metadata_by_id: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get gene metadata by GeneID
        api_response = api_instance.gene_metadata_by_id(gene_ids, returned_content=returned_content, sort_schema_field=sort_schema_field, sort_schema_direction=sort_schema_direction, limit=limit)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling GeneApi->gene_metadata_by_id: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gene_ids** | **[int]**| NCBI gene ids |
 **returned_content** | **V1GeneDatasetRequestContentType**| Return either gene-ids, or entire gene metadata | [optional]
 **sort_schema_field** | **V1GeneDatasetRequestSortField**| Select a field to sort on. | [optional]
 **sort_schema_direction** | **V1SortDirection**| Select a direction for the sort. | [optional]
 **limit** | **str**| Limit the number of returned results (\&quot;all\&quot;, \&quot;none\&quot;, otherwise an integer value) | [optional]

### Return type

[**V1GeneMetadata**](V1GeneMetadata.md)

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

# **gene_metadata_by_post**
> V1GeneMetadata gene_metadata_by_post(v1_gene_dataset_request)

Get gene metadata as JSON

Get detailed gene metadata in a JSON output format.

### Example

* Api Key Authentication (ApiKeyAuthHeader):

```python
import time
import ncbi.datasets.openapi
from ncbi.datasets.openapi.api import gene_api
from ncbi.datasets.openapi.model.v1_gene_metadata import V1GeneMetadata
from ncbi.datasets.openapi.model.v1_gene_dataset_request import V1GeneDatasetRequest
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
    api_instance = gene_api.GeneApi(api_client)
    v1_gene_dataset_request = V1GeneDatasetRequest(
        gene_ids=[
            1,
        ],
        accessions=[
            "accessions_example",
        ],
        symbols_for_taxon=V1GeneDatasetRequestSymbolsForTaxon(
            symbols=[
                "symbols_example",
            ],
            taxon="taxon_example",
        ),
        taxon="taxon_example",
        include_annotation_type=[
            V1Fasta("FASTA_UNSPECIFIED"),
        ],
        returned_content=V1GeneDatasetRequestContentType("COMPLETE"),
        sort_schema=V1GeneDatasetRequestSort(
            field=V1GeneDatasetRequestSortField("SORT_FIELD_UNSPECIFIED"),
            direction=V1SortDirection("SORT_DIRECTION_UNSPECIFIED"),
        ),
        limit="limit_example",
        fasta_filter=[
            "fasta_filter_example",
        ],
    ) # V1GeneDatasetRequest | 

    # example passing only required values which don't have defaults set
    try:
        # Get gene metadata as JSON
        api_response = api_instance.gene_metadata_by_post(v1_gene_dataset_request)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling GeneApi->gene_metadata_by_post: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1_gene_dataset_request** | [**V1GeneDatasetRequest**](V1GeneDatasetRequest.md)|  |

### Return type

[**V1GeneMetadata**](V1GeneMetadata.md)

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

# **gene_metadata_by_tax_and_symbol**
> V1GeneMetadata gene_metadata_by_tax_and_symbol(symbols, taxon)

Get gene metadata by gene symbol

Get detailed gene metadata by gene symbol in a JSON output format.

### Example

* Api Key Authentication (ApiKeyAuthHeader):

```python
import time
import ncbi.datasets.openapi
from ncbi.datasets.openapi.api import gene_api
from ncbi.datasets.openapi.model.v1_gene_dataset_request_content_type import V1GeneDatasetRequestContentType
from ncbi.datasets.openapi.model.v1_sort_direction import V1SortDirection
from ncbi.datasets.openapi.model.v1_gene_dataset_request_sort_field import V1GeneDatasetRequestSortField
from ncbi.datasets.openapi.model.v1_gene_metadata import V1GeneMetadata
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
    api_instance = gene_api.GeneApi(api_client)
    symbols = [
        "GNAS",
    ] # [str] | Gene symbol
    taxon = "9606" # str | Taxon for provided gene symbol
    accessions = ["NM_021803.4","NM_181078.3"] # [str] | RNA or Protein accessions. (optional)
    returned_content = V1GeneDatasetRequestContentType("COMPLETE") # V1GeneDatasetRequestContentType | Return either gene-ids, or entire gene metadata (optional)
    sort_schema_field = V1GeneDatasetRequestSortField("SORT_FIELD_UNSPECIFIED") # V1GeneDatasetRequestSortField | Select a field to sort on. (optional)
    sort_schema_direction = V1SortDirection("SORT_DIRECTION_UNSPECIFIED") # V1SortDirection | Select a direction for the sort. (optional)
    limit = "limit_example" # str | Limit the number of returned results (\"all\", \"none\", otherwise an integer value) (optional)

    # example passing only required values which don't have defaults set
    try:
        # Get gene metadata by gene symbol
        api_response = api_instance.gene_metadata_by_tax_and_symbol(symbols, taxon)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling GeneApi->gene_metadata_by_tax_and_symbol: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get gene metadata by gene symbol
        api_response = api_instance.gene_metadata_by_tax_and_symbol(symbols, taxon, accessions=accessions, returned_content=returned_content, sort_schema_field=sort_schema_field, sort_schema_direction=sort_schema_direction, limit=limit)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling GeneApi->gene_metadata_by_tax_and_symbol: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbols** | **[str]**| Gene symbol |
 **taxon** | **str**| Taxon for provided gene symbol |
 **accessions** | **[str]**| RNA or Protein accessions. | [optional]
 **returned_content** | **V1GeneDatasetRequestContentType**| Return either gene-ids, or entire gene metadata | [optional]
 **sort_schema_field** | **V1GeneDatasetRequestSortField**| Select a field to sort on. | [optional]
 **sort_schema_direction** | **V1SortDirection**| Select a direction for the sort. | [optional]
 **limit** | **str**| Limit the number of returned results (\&quot;all\&quot;, \&quot;none\&quot;, otherwise an integer value) | [optional]

### Return type

[**V1GeneMetadata**](V1GeneMetadata.md)

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

# **gene_metadata_stream_by_post**
> V1GeneMatch gene_metadata_stream_by_post(v1_gene_dataset_request)

Get gene metadata

Get detailed gene metadata in a streaming, JSON-lines output format.

### Example

* Api Key Authentication (ApiKeyAuthHeader):

```python
import time
import ncbi.datasets.openapi
from ncbi.datasets.openapi.api import gene_api
from ncbi.datasets.openapi.model.v1_gene_match import V1GeneMatch
from ncbi.datasets.openapi.model.v1_gene_dataset_request import V1GeneDatasetRequest
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
    api_instance = gene_api.GeneApi(api_client)
    v1_gene_dataset_request = V1GeneDatasetRequest(
        gene_ids=[
            1,
        ],
        accessions=[
            "accessions_example",
        ],
        symbols_for_taxon=V1GeneDatasetRequestSymbolsForTaxon(
            symbols=[
                "symbols_example",
            ],
            taxon="taxon_example",
        ),
        taxon="taxon_example",
        include_annotation_type=[
            V1Fasta("FASTA_UNSPECIFIED"),
        ],
        returned_content=V1GeneDatasetRequestContentType("COMPLETE"),
        sort_schema=V1GeneDatasetRequestSort(
            field=V1GeneDatasetRequestSortField("SORT_FIELD_UNSPECIFIED"),
            direction=V1SortDirection("SORT_DIRECTION_UNSPECIFIED"),
        ),
        limit="limit_example",
        fasta_filter=[
            "fasta_filter_example",
        ],
    ) # V1GeneDatasetRequest | 

    # example passing only required values which don't have defaults set
    try:
        # Get gene metadata
        api_response = api_instance.gene_metadata_stream_by_post(v1_gene_dataset_request)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling GeneApi->gene_metadata_stream_by_post: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1_gene_dataset_request** | [**V1GeneDatasetRequest**](V1GeneDatasetRequest.md)|  |

### Return type

[**V1GeneMatch**](V1GeneMatch.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/x-ndjson, application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **gene_orthologs_by_id**
> V1OrthologSet gene_orthologs_by_id(gene_id)

Get gene orthologs by gene ID

Get detailed gene metadata for an ortholog set by gene ID in a JSON output format.

### Example

* Api Key Authentication (ApiKeyAuthHeader):

```python
import time
import ncbi.datasets.openapi
from ncbi.datasets.openapi.api import gene_api
from ncbi.datasets.openapi.model.v1_ortholog_set import V1OrthologSet
from ncbi.datasets.openapi.model.rpc_status import RpcStatus
from ncbi.datasets.openapi.model.v1_ortholog_request_content_type import V1OrthologRequestContentType
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
    api_instance = gene_api.GeneApi(api_client)
    gene_id = 2778 # int | 
    returned_content = V1OrthologRequestContentType("COMPLETE") # V1OrthologRequestContentType | Return either gene-ids, or entire gene metadata (optional)
    taxon_filter = [9606,10090] # [str] | Filter genes by taxa (optional)

    # example passing only required values which don't have defaults set
    try:
        # Get gene orthologs by gene ID
        api_response = api_instance.gene_orthologs_by_id(gene_id)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling GeneApi->gene_orthologs_by_id: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get gene orthologs by gene ID
        api_response = api_instance.gene_orthologs_by_id(gene_id, returned_content=returned_content, taxon_filter=taxon_filter)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling GeneApi->gene_orthologs_by_id: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gene_id** | **int**|  |
 **returned_content** | **V1OrthologRequestContentType**| Return either gene-ids, or entire gene metadata | [optional]
 **taxon_filter** | **[str]**| Filter genes by taxa | [optional]

### Return type

[**V1OrthologSet**](V1OrthologSet.md)

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

# **gene_tax_name_query**
> V1SciNameAndIds gene_tax_name_query(taxon_query)

Get a list of taxonomy names and IDs found in the gene dataset given a partial taxonomic name

This endpoint retrieves a list of taxonomy names and IDs found in the gene dataset given a partial taxonomic name of any rank.

### Example

* Api Key Authentication (ApiKeyAuthHeader):

```python
import time
import ncbi.datasets.openapi
from ncbi.datasets.openapi.api import gene_api
from ncbi.datasets.openapi.model.v1_sci_name_and_ids import V1SciNameAndIds
from ncbi.datasets.openapi.model.rpc_status import RpcStatus
from ncbi.datasets.openapi.model.v1_organism_query_request_tax_rank_filter import V1OrganismQueryRequestTaxRankFilter
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
    api_instance = gene_api.GeneApi(api_client)
    taxon_query = "hum" # str | NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank
    tax_rank_filter = V1OrganismQueryRequestTaxRankFilter("species") # V1OrganismQueryRequestTaxRankFilter | Set the scope of searched tax ranks (optional)

    # example passing only required values which don't have defaults set
    try:
        # Get a list of taxonomy names and IDs found in the gene dataset given a partial taxonomic name
        api_response = api_instance.gene_tax_name_query(taxon_query)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling GeneApi->gene_tax_name_query: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get a list of taxonomy names and IDs found in the gene dataset given a partial taxonomic name
        api_response = api_instance.gene_tax_name_query(taxon_query, tax_rank_filter=tax_rank_filter)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling GeneApi->gene_tax_name_query: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taxon_query** | **str**| NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank |
 **tax_rank_filter** | **V1OrganismQueryRequestTaxRankFilter**| Set the scope of searched tax ranks | [optional]

### Return type

[**V1SciNameAndIds**](V1SciNameAndIds.md)

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

# **gene_tax_tree**
> V1Organism gene_tax_tree(taxon)

Get a taxonomic subtree by taxonomic identifier

Using an NCBI Taxonomy ID or name (common or scientific) at any rank, get a subtree filtered for species with assembled genomes.

### Example

* Api Key Authentication (ApiKeyAuthHeader):

```python
import time
import ncbi.datasets.openapi
from ncbi.datasets.openapi.api import gene_api
from ncbi.datasets.openapi.model.v1_organism import V1Organism
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
    api_instance = gene_api.GeneApi(api_client)
    taxon = "9606" # str | NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank
    children_only = False # bool | Only report the children of the requested taxon and not their descendants (optional) if omitted the server will use the default value of False

    # example passing only required values which don't have defaults set
    try:
        # Get a taxonomic subtree by taxonomic identifier
        api_response = api_instance.gene_tax_tree(taxon)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling GeneApi->gene_tax_tree: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get a taxonomic subtree by taxonomic identifier
        api_response = api_instance.gene_tax_tree(taxon, children_only=children_only)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling GeneApi->gene_tax_tree: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taxon** | **str**| NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank |
 **children_only** | **bool**| Only report the children of the requested taxon and not their descendants | [optional] if omitted the server will use the default value of False

### Return type

[**V1Organism**](V1Organism.md)

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

