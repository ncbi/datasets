# ncbi.datasets.GeneApi

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
[**gene_metadata_by_post**](GeneApi.md#gene_metadata_by_post) | **POST** /gene | Get gene metadata
[**gene_metadata_by_tax_and_symbol**](GeneApi.md#gene_metadata_by_tax_and_symbol) | **GET** /gene/symbol/{symbols}/taxon/{taxon} | Get gene metadata by gene symbol.
[**gene_tax_name_query**](GeneApi.md#gene_tax_name_query) | **GET** /gene/taxon_suggest/{taxon_query} | Get a list of taxonomy names and IDs found in the gene dataset given a partial taxonomic name.
[**gene_tax_tree**](GeneApi.md#gene_tax_tree) | **GET** /gene/taxon/{taxon}/tree | Retrieve tax tree


# **download_gene_package**
> file download_gene_package(gene_ids, include_annotation_type=include_annotation_type, filename=filename)

Get a gene dataset by gene ID

Get a gene dataset including gene, transcript and protein fasta sequence, annotation and metadata by gene ID.

### Example

```python
from __future__ import print_function
import time
import ncbi.datasets
from ncbi.datasets.rest import ApiException
from pprint import pprint
# Defining the host is optional and defaults to https://api.ncbi.nlm.nih.gov/datasets/v1alpha
# See configuration.py for a list of all supported configuration parameters.
configuration = ncbi.datasets.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1alpha"
)


# Enter a context with an instance of the API client
with ncbi.datasets.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = ncbi.datasets.GeneApi(api_client)
    gene_ids = [56] # list[int] | 
include_annotation_type = ['include_annotation_type_example'] # list[str] | Select additional types of annotation to include in the data package.  If unset, no annotation is provided. (optional)
filename = 'ncbi_dataset.zip' # str | Output file name. (optional) (default to 'ncbi_dataset.zip')

    try:
        # Get a gene dataset by gene ID
        api_response = api_instance.download_gene_package(gene_ids, include_annotation_type=include_annotation_type, filename=filename)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling GeneApi->download_gene_package: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gene_ids** | [**list[int]**](int.md)|  | 
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
**200** | Download selected gene data as a zip file. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **download_gene_package_post**
> file download_gene_package_post(body, filename=filename)

Get a gene dataset by POST

Get a gene dataset including gene, transcript and protein fasta sequence, annotation and metadata by POST.

### Example

```python
from __future__ import print_function
import time
import ncbi.datasets
from ncbi.datasets.rest import ApiException
from pprint import pprint
# Defining the host is optional and defaults to https://api.ncbi.nlm.nih.gov/datasets/v1alpha
# See configuration.py for a list of all supported configuration parameters.
configuration = ncbi.datasets.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1alpha"
)


# Enter a context with an instance of the API client
with ncbi.datasets.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = ncbi.datasets.GeneApi(api_client)
    body = ncbi.datasets.V1alpha1GeneDatasetRequest() # V1alpha1GeneDatasetRequest | 
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

No authorization required

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
> V1alpha1DownloadSummary gene_download_summary_by_accession(accessions)

Get gene download summary by RefSeq Accession

Get gene download summary by RefSeq Accession in a JSON output format.

### Example

```python
from __future__ import print_function
import time
import ncbi.datasets
from ncbi.datasets.rest import ApiException
from pprint import pprint
# Defining the host is optional and defaults to https://api.ncbi.nlm.nih.gov/datasets/v1alpha
# See configuration.py for a list of all supported configuration parameters.
configuration = ncbi.datasets.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1alpha"
)


# Enter a context with an instance of the API client
with ncbi.datasets.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = ncbi.datasets.GeneApi(api_client)
    accessions = ['accessions_example'] # list[str] | RNA or Protein accessions.

    try:
        # Get gene download summary by RefSeq Accession
        api_response = api_instance.gene_download_summary_by_accession(accessions)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling GeneApi->gene_download_summary_by_accession: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessions** | [**list[str]**](str.md)| RNA or Protein accessions. | 

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

# **gene_download_summary_by_id**
> V1alpha1DownloadSummary gene_download_summary_by_id(gene_ids)

Get gene download summary by GeneID

Get a download summary by GeneID in a JSON output format.

### Example

```python
from __future__ import print_function
import time
import ncbi.datasets
from ncbi.datasets.rest import ApiException
from pprint import pprint
# Defining the host is optional and defaults to https://api.ncbi.nlm.nih.gov/datasets/v1alpha
# See configuration.py for a list of all supported configuration parameters.
configuration = ncbi.datasets.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1alpha"
)


# Enter a context with an instance of the API client
with ncbi.datasets.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = ncbi.datasets.GeneApi(api_client)
    gene_ids = [56] # list[int] | 

    try:
        # Get gene download summary by GeneID
        api_response = api_instance.gene_download_summary_by_id(gene_ids)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling GeneApi->gene_download_summary_by_id: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gene_ids** | [**list[int]**](int.md)|  | 

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

# **gene_download_summary_by_post**
> V1alpha1DownloadSummary gene_download_summary_by_post(body)

Get gene download summary

Get gene download summary in a JSON output format.

### Example

```python
from __future__ import print_function
import time
import ncbi.datasets
from ncbi.datasets.rest import ApiException
from pprint import pprint
# Defining the host is optional and defaults to https://api.ncbi.nlm.nih.gov/datasets/v1alpha
# See configuration.py for a list of all supported configuration parameters.
configuration = ncbi.datasets.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1alpha"
)


# Enter a context with an instance of the API client
with ncbi.datasets.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = ncbi.datasets.GeneApi(api_client)
    body = ncbi.datasets.V1alpha1GeneDatasetRequest() # V1alpha1GeneDatasetRequest | 

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

No authorization required

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
> V1alpha1DownloadSummary gene_download_summary_by_tax_and_symbol(symbols, taxon)

Get gene download summary by gene symbol.

Get gene download summary by gene symbol in a JSON output format.

### Example

```python
from __future__ import print_function
import time
import ncbi.datasets
from ncbi.datasets.rest import ApiException
from pprint import pprint
# Defining the host is optional and defaults to https://api.ncbi.nlm.nih.gov/datasets/v1alpha
# See configuration.py for a list of all supported configuration parameters.
configuration = ncbi.datasets.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1alpha"
)


# Enter a context with an instance of the API client
with ncbi.datasets.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = ncbi.datasets.GeneApi(api_client)
    symbols = ['symbols_example'] # list[str] | 
taxon = 'taxon_example' # str | 

    try:
        # Get gene download summary by gene symbol.
        api_response = api_instance.gene_download_summary_by_tax_and_symbol(symbols, taxon)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling GeneApi->gene_download_summary_by_tax_and_symbol: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbols** | [**list[str]**](str.md)|  | 
 **taxon** | **str**|  | 

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

# **gene_metadata_by_accession**
> V1alpha1GeneMetadata gene_metadata_by_accession(accessions, returned_content=returned_content, sort_schema_field=sort_schema_field, sort_schema_direction=sort_schema_direction)

Get gene metadata by RefSeq Accession

Get detailed gene metadata by RefSeq Accession in a JSON output format.

### Example

```python
from __future__ import print_function
import time
import ncbi.datasets
from ncbi.datasets.rest import ApiException
from pprint import pprint
# Defining the host is optional and defaults to https://api.ncbi.nlm.nih.gov/datasets/v1alpha
# See configuration.py for a list of all supported configuration parameters.
configuration = ncbi.datasets.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1alpha"
)


# Enter a context with an instance of the API client
with ncbi.datasets.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = ncbi.datasets.GeneApi(api_client)
    accessions = ['accessions_example'] # list[str] | RNA or Protein accessions.
returned_content = 'COMPLETE' # str | Return either gene-ids, or entire gene metadata. (optional) (default to 'COMPLETE')
sort_schema_field = 'SORT_FIELD_UNSPECIFIED' # str | Select a field to sort on. (optional) (default to 'SORT_FIELD_UNSPECIFIED')
sort_schema_direction = 'SORT_DIRECTION_UNSPECIFIED' # str | Select a direction for the sort. (optional) (default to 'SORT_DIRECTION_UNSPECIFIED')

    try:
        # Get gene metadata by RefSeq Accession
        api_response = api_instance.gene_metadata_by_accession(accessions, returned_content=returned_content, sort_schema_field=sort_schema_field, sort_schema_direction=sort_schema_direction)
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

### Return type

[**V1alpha1GeneMetadata**](V1alpha1GeneMetadata.md)

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

# **gene_metadata_by_id**
> V1alpha1GeneMetadata gene_metadata_by_id(gene_ids, returned_content=returned_content, sort_schema_field=sort_schema_field, sort_schema_direction=sort_schema_direction)

Get gene metadata by GeneID

Get detailed gene metadata by GeneID in a JSON output format.

### Example

```python
from __future__ import print_function
import time
import ncbi.datasets
from ncbi.datasets.rest import ApiException
from pprint import pprint
# Defining the host is optional and defaults to https://api.ncbi.nlm.nih.gov/datasets/v1alpha
# See configuration.py for a list of all supported configuration parameters.
configuration = ncbi.datasets.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1alpha"
)


# Enter a context with an instance of the API client
with ncbi.datasets.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = ncbi.datasets.GeneApi(api_client)
    gene_ids = [56] # list[int] | 
returned_content = 'COMPLETE' # str | Return either gene-ids, or entire gene metadata. (optional) (default to 'COMPLETE')
sort_schema_field = 'SORT_FIELD_UNSPECIFIED' # str | Select a field to sort on. (optional) (default to 'SORT_FIELD_UNSPECIFIED')
sort_schema_direction = 'SORT_DIRECTION_UNSPECIFIED' # str | Select a direction for the sort. (optional) (default to 'SORT_DIRECTION_UNSPECIFIED')

    try:
        # Get gene metadata by GeneID
        api_response = api_instance.gene_metadata_by_id(gene_ids, returned_content=returned_content, sort_schema_field=sort_schema_field, sort_schema_direction=sort_schema_direction)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling GeneApi->gene_metadata_by_id: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gene_ids** | [**list[int]**](int.md)|  | 
 **returned_content** | **str**| Return either gene-ids, or entire gene metadata. | [optional] [default to &#39;COMPLETE&#39;]
 **sort_schema_field** | **str**| Select a field to sort on. | [optional] [default to &#39;SORT_FIELD_UNSPECIFIED&#39;]
 **sort_schema_direction** | **str**| Select a direction for the sort. | [optional] [default to &#39;SORT_DIRECTION_UNSPECIFIED&#39;]

### Return type

[**V1alpha1GeneMetadata**](V1alpha1GeneMetadata.md)

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

# **gene_metadata_by_post**
> V1alpha1GeneMetadata gene_metadata_by_post(body)

Get gene metadata

Get detailed gene metadata in a JSON output format.

### Example

```python
from __future__ import print_function
import time
import ncbi.datasets
from ncbi.datasets.rest import ApiException
from pprint import pprint
# Defining the host is optional and defaults to https://api.ncbi.nlm.nih.gov/datasets/v1alpha
# See configuration.py for a list of all supported configuration parameters.
configuration = ncbi.datasets.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1alpha"
)


# Enter a context with an instance of the API client
with ncbi.datasets.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = ncbi.datasets.GeneApi(api_client)
    body = ncbi.datasets.V1alpha1GeneDatasetRequest() # V1alpha1GeneDatasetRequest | 

    try:
        # Get gene metadata
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

No authorization required

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
> V1alpha1GeneMetadata gene_metadata_by_tax_and_symbol(symbols, taxon, returned_content=returned_content, sort_schema_field=sort_schema_field, sort_schema_direction=sort_schema_direction)

Get gene metadata by gene symbol.

Get detailed gene metadata by gene symbol in a JSON output format.

### Example

```python
from __future__ import print_function
import time
import ncbi.datasets
from ncbi.datasets.rest import ApiException
from pprint import pprint
# Defining the host is optional and defaults to https://api.ncbi.nlm.nih.gov/datasets/v1alpha
# See configuration.py for a list of all supported configuration parameters.
configuration = ncbi.datasets.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1alpha"
)


# Enter a context with an instance of the API client
with ncbi.datasets.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = ncbi.datasets.GeneApi(api_client)
    symbols = ['symbols_example'] # list[str] | 
taxon = 'taxon_example' # str | 
returned_content = 'COMPLETE' # str | Return either gene-ids, or entire gene metadata. (optional) (default to 'COMPLETE')
sort_schema_field = 'SORT_FIELD_UNSPECIFIED' # str | Select a field to sort on. (optional) (default to 'SORT_FIELD_UNSPECIFIED')
sort_schema_direction = 'SORT_DIRECTION_UNSPECIFIED' # str | Select a direction for the sort. (optional) (default to 'SORT_DIRECTION_UNSPECIFIED')

    try:
        # Get gene metadata by gene symbol.
        api_response = api_instance.gene_metadata_by_tax_and_symbol(symbols, taxon, returned_content=returned_content, sort_schema_field=sort_schema_field, sort_schema_direction=sort_schema_direction)
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

### Return type

[**V1alpha1GeneMetadata**](V1alpha1GeneMetadata.md)

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

# **gene_tax_name_query**
> V1alpha1SciNameAndIds gene_tax_name_query(taxon_query)

Get a list of taxonomy names and IDs found in the gene dataset given a partial taxonomic name.

This endpoint retrieves a list of taxonomy names and IDs found in the gene dataset given a partial taxonomic name of any rank.

### Example

```python
from __future__ import print_function
import time
import ncbi.datasets
from ncbi.datasets.rest import ApiException
from pprint import pprint
# Defining the host is optional and defaults to https://api.ncbi.nlm.nih.gov/datasets/v1alpha
# See configuration.py for a list of all supported configuration parameters.
configuration = ncbi.datasets.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1alpha"
)


# Enter a context with an instance of the API client
with ncbi.datasets.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = ncbi.datasets.GeneApi(api_client)
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

# **gene_tax_tree**
> V1alpha1Organism gene_tax_tree(taxon)

Retrieve tax tree

### Example

```python
from __future__ import print_function
import time
import ncbi.datasets
from ncbi.datasets.rest import ApiException
from pprint import pprint
# Defining the host is optional and defaults to https://api.ncbi.nlm.nih.gov/datasets/v1alpha
# See configuration.py for a list of all supported configuration parameters.
configuration = ncbi.datasets.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1alpha"
)


# Enter a context with an instance of the API client
with ncbi.datasets.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = ncbi.datasets.GeneApi(api_client)
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

