# ncbi.datasets.VirusDownloadApi

All URIs are relative to *https://api.ncbi.nlm.nih.gov/datasets/v1alpha*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_virus_dataset_stream**](VirusDownloadApi.md#get_virus_dataset_stream) | **GET** /download/virus/sars2 | Retrieve a requested virus dataset and stream back reply


# **get_virus_dataset_stream**
> file get_virus_dataset_stream()

Retrieve a requested virus dataset and stream back reply

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
    api_instance = ncbi.datasets.VirusDownloadApi(api_client)
    
    try:
        # Retrieve a requested virus dataset and stream back reply
        api_response = api_instance.get_virus_dataset_stream()
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling VirusDownloadApi->get_virus_dataset_stream: %s\n" % e)
```

### Parameters
This endpoint does not need any parameter.

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

