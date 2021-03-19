# ncbi.datasets.openapi.DatasetDownloadApi

All URIs are relative to *https://api.ncbi.nlm.nih.gov/datasets/v1alpha*

Method | HTTP request | Description
------------- | ------------- | -------------
[**dataset_download_get_prokaryote_gene_dataset_stream2**](DatasetDownloadApi.md#dataset_download_get_prokaryote_gene_dataset_stream2) | **POST** /protein/accession/download | Retrieve Prokaryote Gene data package.


# **dataset_download_get_prokaryote_gene_dataset_stream2**
> StreamResultOfProtobufBytesValue dataset_download_get_prokaryote_gene_dataset_stream2(body)

Retrieve Prokaryote Gene data package.

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
    api_instance = ncbi.datasets.openapi.DatasetDownloadApi(api_client)
    body = ncbi.datasets.openapi.V1alpha1ProkaryoteGeneRequest() # V1alpha1ProkaryoteGeneRequest | 

    try:
        # Retrieve Prokaryote Gene data package.
        api_response = api_instance.dataset_download_get_prokaryote_gene_dataset_stream2(body)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling DatasetDownloadApi->dataset_download_get_prokaryote_gene_dataset_stream2: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**V1alpha1ProkaryoteGeneRequest**](V1alpha1ProkaryoteGeneRequest.md)|  | 

### Return type

[**StreamResultOfProtobufBytesValue**](StreamResultOfProtobufBytesValue.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response.(streaming responses) |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

