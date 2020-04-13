# ncbi.datasets.DownloadApi

All URIs are relative to *https://api.ncbi.nlm.nih.gov/datasets/v1alpha*

Method | HTTP request | Description
------------- | ------------- | -------------
[**download_assembly_package**](DownloadApi.md#download_assembly_package) | **GET** /download/assembly_accession/{assembly_accessions} | Retrieve a requested assembly dataset and stream back reply by assembly accession
[**download_assembly_package_post**](DownloadApi.md#download_assembly_package_post) | **POST** /download/assembly_accession | Retrieve a requested assembly dataset and stream back reply by POST
[**download_gene_package**](DownloadApi.md#download_gene_package) | **GET** /download/gene/id/{gene_ids} | Retrieve a requested gene dataset and stream back reply by gene id
[**download_gene_package_post**](DownloadApi.md#download_gene_package_post) | **POST** /download/gene | Retrieve a requested gene dataset and stream back reply by POST


# **download_assembly_package**
> file download_assembly_package(assembly_accessions, chromosomes=chromosomes, include_sequence=include_sequence, include_annotation_type=include_annotation_type)

Retrieve a requested assembly dataset and stream back reply by assembly accession

### Example

```python
from __future__ import print_function
import time
import ncbi.datasets
from ncbi.datasets.rest import ApiException
from pprint import pprint

# Enter a context with an instance of the API client
with ncbi.datasets.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = ncbi.datasets.DownloadApi(api_client)
    assembly_accessions = ['assembly_accessions_example'] # list[str] | Use 'add item' to include multiple assembly accessions.
chromosomes = ['chromosomes_example'] # list[str] | The default setting is all chromosome. Specify individual chromosome by string (1,2,MT or chr1,chr2.chrMT). Unplaced sequences are treated like their own chromosome ('Un'). The filter only applies to fasta sequence. (optional)
include_sequence = True # bool |  (optional)
include_annotation_type = ['include_annotation_type_example'] # list[str] |  (optional)

    try:
        # Retrieve a requested assembly dataset and stream back reply by assembly accession
        api_response = api_instance.download_assembly_package(assembly_accessions, chromosomes=chromosomes, include_sequence=include_sequence, include_annotation_type=include_annotation_type)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling DownloadApi->download_assembly_package: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assembly_accessions** | [**list[str]**](str.md)| Use &#39;add item&#39; to include multiple assembly accessions. | 
 **chromosomes** | [**list[str]**](str.md)| The default setting is all chromosome. Specify individual chromosome by string (1,2,MT or chr1,chr2.chrMT). Unplaced sequences are treated like their own chromosome (&#39;Un&#39;). The filter only applies to fasta sequence. | [optional] 
 **include_sequence** | **bool**|  | [optional] 
 **include_annotation_type** | [**list[str]**](str.md)|  | [optional] 

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
**200** | Successful download |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **download_assembly_package_post**
> file download_assembly_package_post(body)

Retrieve a requested assembly dataset and stream back reply by POST

### Example

```python
from __future__ import print_function
import time
import ncbi.datasets
from ncbi.datasets.rest import ApiException
from pprint import pprint

# Enter a context with an instance of the API client
with ncbi.datasets.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = ncbi.datasets.DownloadApi(api_client)
    body = ncbi.datasets.V1alpha1AssemblyDatasetRequest() # V1alpha1AssemblyDatasetRequest | 

    try:
        # Retrieve a requested assembly dataset and stream back reply by POST
        api_response = api_instance.download_assembly_package_post(body)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling DownloadApi->download_assembly_package_post: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**V1alpha1AssemblyDatasetRequest**](V1alpha1AssemblyDatasetRequest.md)|  | 

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
**200** | Successful download |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **download_gene_package**
> file download_gene_package(gene_ids)

Retrieve a requested gene dataset and stream back reply by gene id

### Example

```python
from __future__ import print_function
import time
import ncbi.datasets
from ncbi.datasets.rest import ApiException
from pprint import pprint

# Enter a context with an instance of the API client
with ncbi.datasets.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = ncbi.datasets.DownloadApi(api_client)
    gene_ids = [56] # list[int] | NCBI Gene ID

    try:
        # Retrieve a requested gene dataset and stream back reply by gene id
        api_response = api_instance.download_gene_package(gene_ids)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling DownloadApi->download_gene_package: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gene_ids** | [**list[int]**](int.md)| NCBI Gene ID | 

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
**200** | Successful download |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **download_gene_package_post**
> file download_gene_package_post(body)

Retrieve a requested gene dataset and stream back reply by POST

### Example

```python
from __future__ import print_function
import time
import ncbi.datasets
from ncbi.datasets.rest import ApiException
from pprint import pprint

# Enter a context with an instance of the API client
with ncbi.datasets.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = ncbi.datasets.DownloadApi(api_client)
    body = ncbi.datasets.V1alpha1GeneDatasetRequest() # V1alpha1GeneDatasetRequest | 

    try:
        # Retrieve a requested gene dataset and stream back reply by POST
        api_response = api_instance.download_gene_package_post(body)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling DownloadApi->download_gene_package_post: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**V1alpha1GeneDatasetRequest**](V1alpha1GeneDatasetRequest.md)|  | 

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
**200** | Successful download |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

