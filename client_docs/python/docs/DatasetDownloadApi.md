# ncbi.datasets.DatasetDownloadApi

All URIs are relative to *https://api.ncbi.nlm.nih.gov/datasets/v1alpha*

Method | HTTP request | Description
------------- | ------------- | -------------
[**dataset_download_version**](DatasetDownloadApi.md#dataset_download_version) | **GET** /version | Retrieve service version


# **dataset_download_version**
> V1alpha1VersionReply dataset_download_version()

Retrieve service version

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
    api_instance = ncbi.datasets.DatasetDownloadApi(api_client)
    
    try:
        # Retrieve service version
        api_response = api_instance.dataset_download_version()
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling DatasetDownloadApi->dataset_download_version: %s\n" % e)
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**V1alpha1VersionReply**](V1alpha1VersionReply.md)

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

