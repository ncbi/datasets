# ncbi.datasets.openapi.GeneApi

All URIs are relative to *https://api.ncbi.nlm.nih.gov/datasets/v1alpha*

Method | HTTP request | Description
------------- | ------------- | -------------
[**download_gene_package**](GeneApi.md#download_gene_package) | **GET** /gene/id/{gene_ids}/download | Get a gene dataset by gene ID
[**download_gene_package_post**](GeneApi.md#download_gene_package_post) | **POST** /gene/download | Get a gene dataset by POST
[**gene_download_summary_by_accession**](GeneApi.md#gene_download_summary_by_accession) | **GET** /gene/accession/{accessions}/download_summary | Get gene download summary by RefSeq Accession
[**gene_download_summary_by_id**](GeneApi.md#gene_download_summary_by_id) | **GET** /gene/id/{gene_ids}/download_summary | Get gene download summary by GeneID
[**gene_download_summary_by_post**](GeneApi.md#gene_download_summary_by_post) | **POST** /gene/download_summary | Get gene download summary
[**gene_download_summary_by_tax_and_symbol**](GeneApi.md#gene_download_summary_by_tax_and_symbol) | **GET** /gene/symbol/{symbols}/taxon/{taxon}/download_summary | Get gene download summary by gene symbol.
[**gene_metadata_by_accession**](GeneApi.md#gene_metadata_by_accession) | **GET** /gene/accession/{accessions} | Get gene metadata by RefSeq Accession
[**gene_metadata_by_id**](GeneApi.md#gene_metadata_by_id) | **GET** /gene/id/{gene_ids} | Get gene metadata by GeneID
[**gene_metadata_by_post**](GeneApi.md#gene_metadata_by_post) | **POST** /gene | Get gene metadata as JSON
[**gene_metadata_by_tax_and_symbol**](GeneApi.md#gene_metadata_by_tax_and_symbol) | **GET** /gene/symbol/{symbols}/taxon/{taxon} | Get gene metadata by gene symbol.
[**gene_metadata_stream_by_post**](GeneApi.md#gene_metadata_stream_by_post) | **POST** /gene/stream | Get gene metadata
[**gene_orthologs_by_id**](GeneApi.md#gene_orthologs_by_id) | **GET** /gene/id/{gene_id}/orthologs | Get gene orthologs by gene ID
[**gene_tax_name_query**](GeneApi.md#gene_tax_name_query) | **GET** /gene/taxon_suggest/{taxon_query} | Get a list of taxonomy names and IDs found in the gene dataset given a partial taxonomic name.
[**gene_tax_tree**](GeneApi.md#gene_tax_tree) | **GET** /gene/taxon/{taxon}/tree | Retrieve tax tree


# **download_gene_package**
> file download_gene_package(gene_ids, include_annotation_type=include_annotation_type, fasta_filter=fasta_filter, filename=filename)

Get a gene dataset by gene ID

Get a gene dataset including gene, transcript and protein fasta sequence, annotation and metadata by gene ID.

### Example

* Api Key Authentication (ApiKeyAuthHeader):
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

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: ApiKeyAuthHeader
configuration = ncbi.datasets.openapi.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1alpha",
    api_key = {
        'api-key': 'YOUR_API_KEY'
    }
)
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['api-key'] = 'Bearer'

# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = ncbi.datasets.openapi.GeneApi(api_client)
    gene_ids = [56] # list[int] | NCBI gene ids
include_annotation_type = ['include_annotation_type_example'] # list[str] | Select additional types of annotation to include in the data package.  If unset, no annotation is provided. (optional)
fasta_filter = ['fasta_filter_example'] # list[str] | Limit the FASTA sequences in the datasets package to these transcript and protein accessions. (optional)
filename = 'ncbi_dataset.zip' # str | Output file name. (optional) (default to 'ncbi_dataset.zip')

    try:
        # Get a gene dataset by gene ID
        api_response = api_instance.download_gene_package(gene_ids, include_annotation_type=include_annotation_type, fasta_filter=fasta_filter, filename=filename)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling GeneApi->download_gene_package: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gene_ids** | [**list[int]**](int.md)| NCBI gene ids | 
 **include_annotation_type** | [**list[str]**](str.md)| Select additional types of annotation to include in the data package.  If unset, no annotation is provided. | [optional] 
 **fasta_filter** | [**list[str]**](str.md)| Limit the FASTA sequences in the datasets package to these transcript and protein accessions. | [optional] 
 **filename** | **str**| Output file name. | [optional] [default to &#39;ncbi_dataset.zip&#39;]

### Return type

**file**

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/zip

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Download selected gene data as a zip file. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **download_gene_package_post**
> file download_gene_package_post(body, filename=filename)

Get a gene dataset by POST

Get a gene dataset including gene, transcript and protein fasta sequence, annotation and metadata by POST.

### Example

* Api Key Authentication (ApiKeyAuthHeader):
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

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: ApiKeyAuthHeader
configuration = ncbi.datasets.openapi.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1alpha",
    api_key = {
        'api-key': 'YOUR_API_KEY'
    }
)
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['api-key'] = 'Bearer'

# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = ncbi.datasets.openapi.GeneApi(api_client)
    body = ncbi.datasets.openapi.V1alpha1GeneDatasetRequest() # V1alpha1GeneDatasetRequest | 
filename = 'ncbi_dataset.zip' # str | Output file name. (optional) (default to 'ncbi_dataset.zip')

    try:
        # Get a gene dataset by POST
        api_response = api_instance.download_gene_package_post(body, filename=filename)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling GeneApi->download_gene_package_post: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**V1alpha1GeneDatasetRequest**](V1alpha1GeneDatasetRequest.md)|  | 
 **filename** | **str**| Output file name. | [optional] [default to &#39;ncbi_dataset.zip&#39;]

### Return type

**file**

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/zip

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Download selected gene data as a zip file. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **gene_download_summary_by_accession**
> V1alpha1DownloadSummary gene_download_summary_by_accession(accessions, limit=limit, fasta_filter=fasta_filter)

Get gene download summary by RefSeq Accession

Get gene download summary by RefSeq Accession in a JSON output format.

### Example

* Api Key Authentication (ApiKeyAuthHeader):
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

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: ApiKeyAuthHeader
configuration = ncbi.datasets.openapi.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1alpha",
    api_key = {
        'api-key': 'YOUR_API_KEY'
    }
)
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['api-key'] = 'Bearer'

# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = ncbi.datasets.openapi.GeneApi(api_client)
    accessions = ['accessions_example'] # list[str] | RNA or Protein accessions.
limit = 'limit_example' # str | Limit the number of returned results (\"all\", \"none\", otherwise an integer value). (optional)
fasta_filter = ['fasta_filter_example'] # list[str] | Limit the FASTA sequences in the datasets package to these transcript and protein accessions. (optional)

    try:
        # Get gene download summary by RefSeq Accession
        api_response = api_instance.gene_download_summary_by_accession(accessions, limit=limit, fasta_filter=fasta_filter)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling GeneApi->gene_download_summary_by_accession: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessions** | [**list[str]**](str.md)| RNA or Protein accessions. | 
 **limit** | **str**| Limit the number of returned results (\&quot;all\&quot;, \&quot;none\&quot;, otherwise an integer value). | [optional] 
 **fasta_filter** | [**list[str]**](str.md)| Limit the FASTA sequences in the datasets package to these transcript and protein accessions. | [optional] 

### Return type

[**V1alpha1DownloadSummary**](V1alpha1DownloadSummary.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **gene_download_summary_by_id**
> V1alpha1DownloadSummary gene_download_summary_by_id(gene_ids, limit=limit, fasta_filter=fasta_filter)

Get gene download summary by GeneID

Get a download summary by GeneID in a JSON output format.

### Example

* Api Key Authentication (ApiKeyAuthHeader):
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

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: ApiKeyAuthHeader
configuration = ncbi.datasets.openapi.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1alpha",
    api_key = {
        'api-key': 'YOUR_API_KEY'
    }
)
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['api-key'] = 'Bearer'

# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = ncbi.datasets.openapi.GeneApi(api_client)
    gene_ids = [56] # list[int] | NCBI gene ids
limit = 'limit_example' # str | Limit the number of returned results (\"all\", \"none\", otherwise an integer value). (optional)
fasta_filter = ['fasta_filter_example'] # list[str] | Limit the FASTA sequences in the datasets package to these transcript and protein accessions. (optional)

    try:
        # Get gene download summary by GeneID
        api_response = api_instance.gene_download_summary_by_id(gene_ids, limit=limit, fasta_filter=fasta_filter)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling GeneApi->gene_download_summary_by_id: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gene_ids** | [**list[int]**](int.md)| NCBI gene ids | 
 **limit** | **str**| Limit the number of returned results (\&quot;all\&quot;, \&quot;none\&quot;, otherwise an integer value). | [optional] 
 **fasta_filter** | [**list[str]**](str.md)| Limit the FASTA sequences in the datasets package to these transcript and protein accessions. | [optional] 

### Return type

[**V1alpha1DownloadSummary**](V1alpha1DownloadSummary.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **gene_download_summary_by_post**
> V1alpha1DownloadSummary gene_download_summary_by_post(body)

Get gene download summary

Get gene download summary in a JSON output format.

### Example

* Api Key Authentication (ApiKeyAuthHeader):
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

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: ApiKeyAuthHeader
configuration = ncbi.datasets.openapi.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1alpha",
    api_key = {
        'api-key': 'YOUR_API_KEY'
    }
)
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['api-key'] = 'Bearer'

# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = ncbi.datasets.openapi.GeneApi(api_client)
    body = ncbi.datasets.openapi.V1alpha1GeneDatasetRequest() # V1alpha1GeneDatasetRequest | 

    try:
        # Get gene download summary
        api_response = api_instance.gene_download_summary_by_post(body)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling GeneApi->gene_download_summary_by_post: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**V1alpha1GeneDatasetRequest**](V1alpha1GeneDatasetRequest.md)|  | 

### Return type

[**V1alpha1DownloadSummary**](V1alpha1DownloadSummary.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **gene_download_summary_by_tax_and_symbol**
> V1alpha1DownloadSummary gene_download_summary_by_tax_and_symbol(symbols, taxon, fasta_filter=fasta_filter)

Get gene download summary by gene symbol.

Get gene download summary by gene symbol in a JSON output format.

### Example

* Api Key Authentication (ApiKeyAuthHeader):
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

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: ApiKeyAuthHeader
configuration = ncbi.datasets.openapi.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1alpha",
    api_key = {
        'api-key': 'YOUR_API_KEY'
    }
)
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['api-key'] = 'Bearer'

# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = ncbi.datasets.openapi.GeneApi(api_client)
    symbols = ['symbols_example'] # list[str] | 
taxon = 'taxon_example' # str | 
fasta_filter = ['fasta_filter_example'] # list[str] | Limit the FASTA sequences in the datasets package to these transcript and protein accessions. (optional)

    try:
        # Get gene download summary by gene symbol.
        api_response = api_instance.gene_download_summary_by_tax_and_symbol(symbols, taxon, fasta_filter=fasta_filter)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling GeneApi->gene_download_summary_by_tax_and_symbol: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbols** | [**list[str]**](str.md)|  | 
 **taxon** | **str**|  | 
 **fasta_filter** | [**list[str]**](str.md)| Limit the FASTA sequences in the datasets package to these transcript and protein accessions. | [optional] 

### Return type

[**V1alpha1DownloadSummary**](V1alpha1DownloadSummary.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **gene_metadata_by_accession**
> V1alpha1GeneMetadata gene_metadata_by_accession(accessions, returned_content=returned_content, sort_schema_field=sort_schema_field, sort_schema_direction=sort_schema_direction, limit=limit)

Get gene metadata by RefSeq Accession

Get detailed gene metadata by RefSeq Accession in a JSON output format.

### Example

* Api Key Authentication (ApiKeyAuthHeader):
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

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: ApiKeyAuthHeader
configuration = ncbi.datasets.openapi.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1alpha",
    api_key = {
        'api-key': 'YOUR_API_KEY'
    }
)
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['api-key'] = 'Bearer'

# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = ncbi.datasets.openapi.GeneApi(api_client)
    accessions = ['accessions_example'] # list[str] | RNA or Protein accessions.
returned_content = 'COMPLETE' # str | Return either gene-ids, or entire gene metadata. (optional) (default to 'COMPLETE')
sort_schema_field = 'SORT_FIELD_UNSPECIFIED' # str | Select a field to sort on. (optional) (default to 'SORT_FIELD_UNSPECIFIED')
sort_schema_direction = 'SORT_DIRECTION_UNSPECIFIED' # str | Select a direction for the sort. (optional) (default to 'SORT_DIRECTION_UNSPECIFIED')
limit = 'limit_example' # str | Limit the number of returned results (\"all\", \"none\", otherwise an integer value). (optional)

    try:
        # Get gene metadata by RefSeq Accession
        api_response = api_instance.gene_metadata_by_accession(accessions, returned_content=returned_content, sort_schema_field=sort_schema_field, sort_schema_direction=sort_schema_direction, limit=limit)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling GeneApi->gene_metadata_by_accession: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessions** | [**list[str]**](str.md)| RNA or Protein accessions. | 
 **returned_content** | **str**| Return either gene-ids, or entire gene metadata. | [optional] [default to &#39;COMPLETE&#39;]
 **sort_schema_field** | **str**| Select a field to sort on. | [optional] [default to &#39;SORT_FIELD_UNSPECIFIED&#39;]
 **sort_schema_direction** | **str**| Select a direction for the sort. | [optional] [default to &#39;SORT_DIRECTION_UNSPECIFIED&#39;]
 **limit** | **str**| Limit the number of returned results (\&quot;all\&quot;, \&quot;none\&quot;, otherwise an integer value). | [optional] 

### Return type

[**V1alpha1GeneMetadata**](V1alpha1GeneMetadata.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **gene_metadata_by_id**
> V1alpha1GeneMetadata gene_metadata_by_id(gene_ids, returned_content=returned_content, sort_schema_field=sort_schema_field, sort_schema_direction=sort_schema_direction, limit=limit)

Get gene metadata by GeneID

Get detailed gene metadata by GeneID in a JSON output format.

### Example

* Api Key Authentication (ApiKeyAuthHeader):
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

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: ApiKeyAuthHeader
configuration = ncbi.datasets.openapi.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1alpha",
    api_key = {
        'api-key': 'YOUR_API_KEY'
    }
)
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['api-key'] = 'Bearer'

# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = ncbi.datasets.openapi.GeneApi(api_client)
    gene_ids = [56] # list[int] | NCBI gene ids
returned_content = 'COMPLETE' # str | Return either gene-ids, or entire gene metadata. (optional) (default to 'COMPLETE')
sort_schema_field = 'SORT_FIELD_UNSPECIFIED' # str | Select a field to sort on. (optional) (default to 'SORT_FIELD_UNSPECIFIED')
sort_schema_direction = 'SORT_DIRECTION_UNSPECIFIED' # str | Select a direction for the sort. (optional) (default to 'SORT_DIRECTION_UNSPECIFIED')
limit = 'limit_example' # str | Limit the number of returned results (\"all\", \"none\", otherwise an integer value). (optional)

    try:
        # Get gene metadata by GeneID
        api_response = api_instance.gene_metadata_by_id(gene_ids, returned_content=returned_content, sort_schema_field=sort_schema_field, sort_schema_direction=sort_schema_direction, limit=limit)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling GeneApi->gene_metadata_by_id: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gene_ids** | [**list[int]**](int.md)| NCBI gene ids | 
 **returned_content** | **str**| Return either gene-ids, or entire gene metadata. | [optional] [default to &#39;COMPLETE&#39;]
 **sort_schema_field** | **str**| Select a field to sort on. | [optional] [default to &#39;SORT_FIELD_UNSPECIFIED&#39;]
 **sort_schema_direction** | **str**| Select a direction for the sort. | [optional] [default to &#39;SORT_DIRECTION_UNSPECIFIED&#39;]
 **limit** | **str**| Limit the number of returned results (\&quot;all\&quot;, \&quot;none\&quot;, otherwise an integer value). | [optional] 

### Return type

[**V1alpha1GeneMetadata**](V1alpha1GeneMetadata.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **gene_metadata_by_post**
> V1alpha1GeneMetadata gene_metadata_by_post(body)

Get gene metadata as JSON

Get detailed gene metadata in a JSON output format.

### Example

* Api Key Authentication (ApiKeyAuthHeader):
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

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: ApiKeyAuthHeader
configuration = ncbi.datasets.openapi.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1alpha",
    api_key = {
        'api-key': 'YOUR_API_KEY'
    }
)
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['api-key'] = 'Bearer'

# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = ncbi.datasets.openapi.GeneApi(api_client)
    body = ncbi.datasets.openapi.V1alpha1GeneDatasetRequest() # V1alpha1GeneDatasetRequest | 

    try:
        # Get gene metadata as JSON
        api_response = api_instance.gene_metadata_by_post(body)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling GeneApi->gene_metadata_by_post: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**V1alpha1GeneDatasetRequest**](V1alpha1GeneDatasetRequest.md)|  | 

### Return type

[**V1alpha1GeneMetadata**](V1alpha1GeneMetadata.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **gene_metadata_by_tax_and_symbol**
> V1alpha1GeneMetadata gene_metadata_by_tax_and_symbol(symbols, taxon, returned_content=returned_content, sort_schema_field=sort_schema_field, sort_schema_direction=sort_schema_direction, limit=limit)

Get gene metadata by gene symbol.

Get detailed gene metadata by gene symbol in a JSON output format.

### Example

* Api Key Authentication (ApiKeyAuthHeader):
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

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: ApiKeyAuthHeader
configuration = ncbi.datasets.openapi.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1alpha",
    api_key = {
        'api-key': 'YOUR_API_KEY'
    }
)
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['api-key'] = 'Bearer'

# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = ncbi.datasets.openapi.GeneApi(api_client)
    symbols = ['symbols_example'] # list[str] | 
taxon = 'taxon_example' # str | 
returned_content = 'COMPLETE' # str | Return either gene-ids, or entire gene metadata. (optional) (default to 'COMPLETE')
sort_schema_field = 'SORT_FIELD_UNSPECIFIED' # str | Select a field to sort on. (optional) (default to 'SORT_FIELD_UNSPECIFIED')
sort_schema_direction = 'SORT_DIRECTION_UNSPECIFIED' # str | Select a direction for the sort. (optional) (default to 'SORT_DIRECTION_UNSPECIFIED')
limit = 'limit_example' # str | Limit the number of returned results (\"all\", \"none\", otherwise an integer value). (optional)

    try:
        # Get gene metadata by gene symbol.
        api_response = api_instance.gene_metadata_by_tax_and_symbol(symbols, taxon, returned_content=returned_content, sort_schema_field=sort_schema_field, sort_schema_direction=sort_schema_direction, limit=limit)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling GeneApi->gene_metadata_by_tax_and_symbol: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbols** | [**list[str]**](str.md)|  | 
 **taxon** | **str**|  | 
 **returned_content** | **str**| Return either gene-ids, or entire gene metadata. | [optional] [default to &#39;COMPLETE&#39;]
 **sort_schema_field** | **str**| Select a field to sort on. | [optional] [default to &#39;SORT_FIELD_UNSPECIFIED&#39;]
 **sort_schema_direction** | **str**| Select a direction for the sort. | [optional] [default to &#39;SORT_DIRECTION_UNSPECIFIED&#39;]
 **limit** | **str**| Limit the number of returned results (\&quot;all\&quot;, \&quot;none\&quot;, otherwise an integer value). | [optional] 

### Return type

[**V1alpha1GeneMetadata**](V1alpha1GeneMetadata.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **gene_metadata_stream_by_post**
> V1alpha1GeneMatch gene_metadata_stream_by_post(body)

Get gene metadata

Get detailed gene metadata in a streaming, JSON-lines output format.

### Example

* Api Key Authentication (ApiKeyAuthHeader):
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

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: ApiKeyAuthHeader
configuration = ncbi.datasets.openapi.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1alpha",
    api_key = {
        'api-key': 'YOUR_API_KEY'
    }
)
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['api-key'] = 'Bearer'

# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = ncbi.datasets.openapi.GeneApi(api_client)
    body = ncbi.datasets.openapi.V1alpha1GeneDatasetRequest() # V1alpha1GeneDatasetRequest | 

    try:
        # Get gene metadata
        api_response = api_instance.gene_metadata_stream_by_post(body)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling GeneApi->gene_metadata_stream_by_post: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**V1alpha1GeneDatasetRequest**](V1alpha1GeneDatasetRequest.md)|  | 

### Return type

[**V1alpha1GeneMatch**](V1alpha1GeneMatch.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/x-ndjson

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successful request, returning a stream of individual JSON Text payloads. |  * Transfer-Encoding - chunked <br>  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **gene_orthologs_by_id**
> V1alpha1OrthologSet gene_orthologs_by_id(gene_id, returned_content=returned_content, taxon_filter=taxon_filter)

Get gene orthologs by gene ID

Get detailed gene metadata for an ortholog set by gene ID in a JSON output format.

### Example

* Api Key Authentication (ApiKeyAuthHeader):
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

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: ApiKeyAuthHeader
configuration = ncbi.datasets.openapi.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1alpha",
    api_key = {
        'api-key': 'YOUR_API_KEY'
    }
)
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['api-key'] = 'Bearer'

# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = ncbi.datasets.openapi.GeneApi(api_client)
    gene_id = 56 # int | 
returned_content = 'COMPLETE' # str | Return either gene-ids, or entire gene metadata. (optional) (default to 'COMPLETE')
taxon_filter = ['taxon_filter_example'] # list[str] | Filter genes by taxa. (optional)

    try:
        # Get gene orthologs by gene ID
        api_response = api_instance.gene_orthologs_by_id(gene_id, returned_content=returned_content, taxon_filter=taxon_filter)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling GeneApi->gene_orthologs_by_id: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gene_id** | **int**|  | 
 **returned_content** | **str**| Return either gene-ids, or entire gene metadata. | [optional] [default to &#39;COMPLETE&#39;]
 **taxon_filter** | [**list[str]**](str.md)| Filter genes by taxa. | [optional] 

### Return type

[**V1alpha1OrthologSet**](V1alpha1OrthologSet.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **gene_tax_name_query**
> V1alpha1SciNameAndIds gene_tax_name_query(taxon_query)

Get a list of taxonomy names and IDs found in the gene dataset given a partial taxonomic name.

This endpoint retrieves a list of taxonomy names and IDs found in the gene dataset given a partial taxonomic name of any rank.

### Example

* Api Key Authentication (ApiKeyAuthHeader):
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

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: ApiKeyAuthHeader
configuration = ncbi.datasets.openapi.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1alpha",
    api_key = {
        'api-key': 'YOUR_API_KEY'
    }
)
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['api-key'] = 'Bearer'

# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = ncbi.datasets.openapi.GeneApi(api_client)
    taxon_query = 'taxon_query_example' # str | NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank

    try:
        # Get a list of taxonomy names and IDs found in the gene dataset given a partial taxonomic name.
        api_response = api_instance.gene_tax_name_query(taxon_query)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling GeneApi->gene_tax_name_query: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taxon_query** | **str**| NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank | 

### Return type

[**V1alpha1SciNameAndIds**](V1alpha1SciNameAndIds.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **gene_tax_tree**
> V1alpha1Organism gene_tax_tree(taxon)

Retrieve tax tree

### Example

* Api Key Authentication (ApiKeyAuthHeader):
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

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: ApiKeyAuthHeader
configuration = ncbi.datasets.openapi.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1alpha",
    api_key = {
        'api-key': 'YOUR_API_KEY'
    }
)
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['api-key'] = 'Bearer'

# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = ncbi.datasets.openapi.GeneApi(api_client)
    taxon = 'taxon_example' # str | NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank

    try:
        # Retrieve tax tree
        api_response = api_instance.gene_tax_tree(taxon)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling GeneApi->gene_tax_tree: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taxon** | **str**| NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank | 

### Return type

[**V1alpha1Organism**](V1alpha1Organism.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

