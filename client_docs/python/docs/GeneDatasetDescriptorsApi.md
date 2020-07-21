# ncbi.datasets.GeneDatasetDescriptorsApi

All URIs are relative to *https://api.ncbi.nlm.nih.gov/datasets/v1alpha*

Method | HTTP request | Description
------------- | ------------- | -------------
[**gene_descriptors_by_id**](GeneDatasetDescriptorsApi.md#gene_descriptors_by_id) | **GET** /gene/id/{gene_ids}/descriptors | Retrieve list of descriptions of genes by gene ID
[**gene_descriptors_by_tax_and_symbol**](GeneDatasetDescriptorsApi.md#gene_descriptors_by_tax_and_symbol) | **GET** /gene/symbol/{symbol}/taxonomy/{tax_token}/descriptors | Retrieve list of descriptions of genes by taxonomy and gene symbol
[**gene_summary_by_id**](GeneDatasetDescriptorsApi.md#gene_summary_by_id) | **GET** /gene/id/{gene_ids}/summary | Summary of gene dataset, including options to download package by gene ID
[**gene_summary_by_tax_and_symbol**](GeneDatasetDescriptorsApi.md#gene_summary_by_tax_and_symbol) | **GET** /gene/symbol/{symbol}/taxonomy/{tax_token}/summary | Summary of gene dataset, including options to download package by taxonomy and gene symbol


# **gene_descriptors_by_id**
> V1alpha1GeneDescriptors gene_descriptors_by_id(gene_ids)

Retrieve list of descriptions of genes by gene ID

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
    api_instance = ncbi.datasets.GeneDatasetDescriptorsApi(api_client)
    gene_ids = [56] # list[int] | NCBI Gene ID

    try:
        # Retrieve list of descriptions of genes by gene ID
        api_response = api_instance.gene_descriptors_by_id(gene_ids)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling GeneDatasetDescriptorsApi->gene_descriptors_by_id: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gene_ids** | [**list[int]**](int.md)| NCBI Gene ID | 

### Return type

[**V1alpha1GeneDescriptors**](V1alpha1GeneDescriptors.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **gene_descriptors_by_tax_and_symbol**
> V1alpha1GeneDescriptors gene_descriptors_by_tax_and_symbol(symbol, tax_token)

Retrieve list of descriptions of genes by taxonomy and gene symbol

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
    api_instance = ncbi.datasets.GeneDatasetDescriptorsApi(api_client)
    symbol = ['symbol_example'] # list[str] | 
tax_token = 'tax_token_example' # str | 

    try:
        # Retrieve list of descriptions of genes by taxonomy and gene symbol
        api_response = api_instance.gene_descriptors_by_tax_and_symbol(symbol, tax_token)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling GeneDatasetDescriptorsApi->gene_descriptors_by_tax_and_symbol: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | [**list[str]**](str.md)|  | 
 **tax_token** | **str**|  | 

### Return type

[**V1alpha1GeneDescriptors**](V1alpha1GeneDescriptors.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **gene_summary_by_id**
> V1alpha1DownloadSummary gene_summary_by_id(gene_ids)

Summary of gene dataset, including options to download package by gene ID

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
    api_instance = ncbi.datasets.GeneDatasetDescriptorsApi(api_client)
    gene_ids = [56] # list[int] | NCBI Gene ID

    try:
        # Summary of gene dataset, including options to download package by gene ID
        api_response = api_instance.gene_summary_by_id(gene_ids)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling GeneDatasetDescriptorsApi->gene_summary_by_id: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gene_ids** | [**list[int]**](int.md)| NCBI Gene ID | 

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **gene_summary_by_tax_and_symbol**
> V1alpha1DownloadSummary gene_summary_by_tax_and_symbol(symbol, tax_token)

Summary of gene dataset, including options to download package by taxonomy and gene symbol

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
    api_instance = ncbi.datasets.GeneDatasetDescriptorsApi(api_client)
    symbol = ['symbol_example'] # list[str] | 
tax_token = 'tax_token_example' # str | 

    try:
        # Summary of gene dataset, including options to download package by taxonomy and gene symbol
        api_response = api_instance.gene_summary_by_tax_and_symbol(symbol, tax_token)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling GeneDatasetDescriptorsApi->gene_summary_by_tax_and_symbol: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **symbol** | [**list[str]**](str.md)|  | 
 **tax_token** | **str**|  | 

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

