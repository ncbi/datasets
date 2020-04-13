# ncbi.datasets.AssemblyDatasetDescriptorsApi

All URIs are relative to *https://api.ncbi.nlm.nih.gov/datasets/v1alpha*

Method | HTTP request | Description
------------- | ------------- | -------------
[**assembly_descriptors_by_accession**](AssemblyDatasetDescriptorsApi.md#assembly_descriptors_by_accession) | **GET** /assembly_descriptors/accession/{assembly_accession} | Assembly descriptions by assembly accession
[**assembly_descriptors_by_organism**](AssemblyDatasetDescriptorsApi.md#assembly_descriptors_by_organism) | **GET** /assembly_descriptors/organism/{tax_name} | Assembly descriptions by taxonomic name (scientific or common name at any tax rank)
[**assembly_descriptors_by_taxid**](AssemblyDatasetDescriptorsApi.md#assembly_descriptors_by_taxid) | **GET** /assembly_descriptors/taxid/{tax_id} | Assembly descriptions by taxonomy ID


# **assembly_descriptors_by_accession**
> V1alpha1AssemblyDatasetDescriptors assembly_descriptors_by_accession(assembly_accession, cutoff=cutoff, limit=limit, filters_refseq_only=filters_refseq_only, tax_exact_match=tax_exact_match, returned_content=returned_content)

Assembly descriptions by assembly accession

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
    api_instance = ncbi.datasets.AssemblyDatasetDescriptorsApi(api_client)
    assembly_accession = 'assembly_accession_example' # str | NCBI Assembly accession
cutoff = 56 # int | Limit the number of results for the above query term. (optional)
limit = 'limit_example' # str | Limit the number of returned results (\"all\", \"none\", otherwise an integer value). (optional)
filters_refseq_only = True # bool | If true, only return RefSeq (GCF_) assemblies. (optional)
tax_exact_match = True # bool | If true, only return assemblies with the given NCBI Taxonomy ID, or name. Otherwise, assemblies from taxonomy subtree are included, too. Ignored for assembly_accession request. (optional)
returned_content = 'COMPLETE' # str |  (optional) (default to 'COMPLETE')

    try:
        # Assembly descriptions by assembly accession
        api_response = api_instance.assembly_descriptors_by_accession(assembly_accession, cutoff=cutoff, limit=limit, filters_refseq_only=filters_refseq_only, tax_exact_match=tax_exact_match, returned_content=returned_content)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling AssemblyDatasetDescriptorsApi->assembly_descriptors_by_accession: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assembly_accession** | **str**| NCBI Assembly accession | 
 **cutoff** | **int**| Limit the number of results for the above query term. | [optional] 
 **limit** | **str**| Limit the number of returned results (\&quot;all\&quot;, \&quot;none\&quot;, otherwise an integer value). | [optional] 
 **filters_refseq_only** | **bool**| If true, only return RefSeq (GCF_) assemblies. | [optional] 
 **tax_exact_match** | **bool**| If true, only return assemblies with the given NCBI Taxonomy ID, or name. Otherwise, assemblies from taxonomy subtree are included, too. Ignored for assembly_accession request. | [optional] 
 **returned_content** | **str**|  | [optional] [default to &#39;COMPLETE&#39;]

### Return type

[**V1alpha1AssemblyDatasetDescriptors**](V1alpha1AssemblyDatasetDescriptors.md)

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

# **assembly_descriptors_by_organism**
> V1alpha1AssemblyDatasetDescriptors assembly_descriptors_by_organism(tax_name, cutoff=cutoff, limit=limit, filters_refseq_only=filters_refseq_only, tax_exact_match=tax_exact_match, returned_content=returned_content)

Assembly descriptions by taxonomic name (scientific or common name at any tax rank)

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
    api_instance = ncbi.datasets.AssemblyDatasetDescriptorsApi(api_client)
    tax_name = 'tax_name_example' # str | 
cutoff = 56 # int | Limit the number of results for the above query term. (optional)
limit = 'limit_example' # str | Limit the number of returned results (\"all\", \"none\", otherwise an integer value). (optional)
filters_refseq_only = True # bool | If true, only return RefSeq (GCF_) assemblies. (optional)
tax_exact_match = True # bool | If true, only return assemblies with the given NCBI Taxonomy ID, or name. Otherwise, assemblies from taxonomy subtree are included, too. Ignored for assembly_accession request. (optional)
returned_content = 'COMPLETE' # str |  (optional) (default to 'COMPLETE')

    try:
        # Assembly descriptions by taxonomic name (scientific or common name at any tax rank)
        api_response = api_instance.assembly_descriptors_by_organism(tax_name, cutoff=cutoff, limit=limit, filters_refseq_only=filters_refseq_only, tax_exact_match=tax_exact_match, returned_content=returned_content)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling AssemblyDatasetDescriptorsApi->assembly_descriptors_by_organism: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tax_name** | **str**|  | 
 **cutoff** | **int**| Limit the number of results for the above query term. | [optional] 
 **limit** | **str**| Limit the number of returned results (\&quot;all\&quot;, \&quot;none\&quot;, otherwise an integer value). | [optional] 
 **filters_refseq_only** | **bool**| If true, only return RefSeq (GCF_) assemblies. | [optional] 
 **tax_exact_match** | **bool**| If true, only return assemblies with the given NCBI Taxonomy ID, or name. Otherwise, assemblies from taxonomy subtree are included, too. Ignored for assembly_accession request. | [optional] 
 **returned_content** | **str**|  | [optional] [default to &#39;COMPLETE&#39;]

### Return type

[**V1alpha1AssemblyDatasetDescriptors**](V1alpha1AssemblyDatasetDescriptors.md)

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

# **assembly_descriptors_by_taxid**
> V1alpha1AssemblyDatasetDescriptors assembly_descriptors_by_taxid(tax_id, cutoff=cutoff, limit=limit, filters_refseq_only=filters_refseq_only, tax_exact_match=tax_exact_match, returned_content=returned_content)

Assembly descriptions by taxonomy ID

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
    api_instance = ncbi.datasets.AssemblyDatasetDescriptorsApi(api_client)
    tax_id = 56 # int | NCBI Taxonomy ID
cutoff = 56 # int | Limit the number of results for the above query term. (optional)
limit = 'limit_example' # str | Limit the number of returned results (\"all\", \"none\", otherwise an integer value). (optional)
filters_refseq_only = True # bool | If true, only return RefSeq (GCF_) assemblies. (optional)
tax_exact_match = True # bool | If true, only return assemblies with the given NCBI Taxonomy ID, or name. Otherwise, assemblies from taxonomy subtree are included, too. Ignored for assembly_accession request. (optional)
returned_content = 'COMPLETE' # str |  (optional) (default to 'COMPLETE')

    try:
        # Assembly descriptions by taxonomy ID
        api_response = api_instance.assembly_descriptors_by_taxid(tax_id, cutoff=cutoff, limit=limit, filters_refseq_only=filters_refseq_only, tax_exact_match=tax_exact_match, returned_content=returned_content)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling AssemblyDatasetDescriptorsApi->assembly_descriptors_by_taxid: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tax_id** | **int**| NCBI Taxonomy ID | 
 **cutoff** | **int**| Limit the number of results for the above query term. | [optional] 
 **limit** | **str**| Limit the number of returned results (\&quot;all\&quot;, \&quot;none\&quot;, otherwise an integer value). | [optional] 
 **filters_refseq_only** | **bool**| If true, only return RefSeq (GCF_) assemblies. | [optional] 
 **tax_exact_match** | **bool**| If true, only return assemblies with the given NCBI Taxonomy ID, or name. Otherwise, assemblies from taxonomy subtree are included, too. Ignored for assembly_accession request. | [optional] 
 **returned_content** | **str**|  | [optional] [default to &#39;COMPLETE&#39;]

### Return type

[**V1alpha1AssemblyDatasetDescriptors**](V1alpha1AssemblyDatasetDescriptors.md)

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

