# ncbi.datasets.GeneDatasetDescriptorsApi

All URIs are relative to *https://api.ncbi.nlm.nih.gov/datasets/v1alpha*

Method | HTTP request | Description
------------- | ------------- | -------------
[**gene_descriptors_by_gene_id**](GeneDatasetDescriptorsApi.md#gene_descriptors_by_gene_id) | **GET** /gene_descriptors/gene/id/{gene_ids} | Retrieve list of descriptions of genes by gene ID


# **gene_descriptors_by_gene_id**
> V1alpha1GeneDescriptors gene_descriptors_by_gene_id(gene_ids)

Retrieve list of descriptions of genes by gene ID

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
    api_instance = ncbi.datasets.GeneDatasetDescriptorsApi(api_client)
    gene_ids = [56] # list[int] | NCBI Gene ID

    try:
        # Retrieve list of descriptions of genes by gene ID
        api_response = api_instance.gene_descriptors_by_gene_id(gene_ids)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling GeneDatasetDescriptorsApi->gene_descriptors_by_gene_id: %s\n" % e)
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

