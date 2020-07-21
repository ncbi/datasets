# ncbi.datasets.TaxTreeApi

All URIs are relative to *https://api.ncbi.nlm.nih.gov/datasets/v1alpha*

Method | HTTP request | Description
------------- | ------------- | -------------
[**tax_tree_by_tax_id**](TaxTreeApi.md#tax_tree_by_tax_id) | **GET** /tax_tree/taxid/{tax_id} | Retrieve tax tree by taxonomy ID


# **tax_tree_by_tax_id**
> V1alpha1Organism tax_tree_by_tax_id(tax_id)

Retrieve tax tree by taxonomy ID

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
    api_instance = ncbi.datasets.TaxTreeApi(api_client)
    tax_id = 'tax_id_example' # str | 

    try:
        # Retrieve tax tree by taxonomy ID
        api_response = api_instance.tax_tree_by_tax_id(tax_id)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling TaxTreeApi->tax_tree_by_tax_id: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tax_id** | **str**|  | 

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

