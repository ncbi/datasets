# ncbi.datasets.AssemblyDatasetDescriptorsApi

All URIs are relative to *https://api.ncbi.nlm.nih.gov/datasets/v1alpha*

Method | HTTP request | Description
------------- | ------------- | -------------
[**assembly_descriptors_by_accession**](AssemblyDatasetDescriptorsApi.md#assembly_descriptors_by_accession) | **GET** /assembly_descriptors/accession/{assembly_accession} | Assembly descriptions by assembly accession
[**assembly_descriptors_by_organism**](AssemblyDatasetDescriptorsApi.md#assembly_descriptors_by_organism) | **GET** /assembly_descriptors/organism/{tax_name} | Assembly descriptions by taxonomic name (scientific or common name at any tax rank)
[**assembly_descriptors_by_taxid**](AssemblyDatasetDescriptorsApi.md#assembly_descriptors_by_taxid) | **GET** /assembly_descriptors/taxid/{tax_id} | Assembly descriptions by taxonomy ID
[**genome_summary_by_accession**](AssemblyDatasetDescriptorsApi.md#genome_summary_by_accession) | **GET** /genome/summary/accession/{assembly_accessions} | Summary of assembly dataset, including options to download package
[**genome_tax_name_query**](AssemblyDatasetDescriptorsApi.md#genome_tax_name_query) | **GET** /genome/tax_name_query/{organism_query}/names | Retrieve list of taxonomy names and is for OrganismQuery
[**genome_tax_tree**](AssemblyDatasetDescriptorsApi.md#genome_tax_tree) | **GET** /genome/taxonomy/{tax_token}/tree | Retrieve tax tree


# **assembly_descriptors_by_accession**
> V1alpha1AssemblyDatasetDescriptors assembly_descriptors_by_accession(assembly_accession, limit=limit, filters_refseq_only=filters_refseq_only, tax_exact_match=tax_exact_match, returned_content=returned_content)

Assembly descriptions by assembly accession

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
    api_instance = ncbi.datasets.AssemblyDatasetDescriptorsApi(api_client)
    assembly_accession = 'assembly_accession_example' # str | NCBI Assembly accession
limit = 'limit_example' # str | Limit the number of returned results (\"all\", \"none\", otherwise an integer value). (optional)
filters_refseq_only = True # bool | If true, only return RefSeq (GCF_) assemblies. (optional)
tax_exact_match = True # bool | If true, only return assemblies with the given NCBI Taxonomy ID, or name. Otherwise, assemblies from taxonomy subtree are included, too. Ignored for assembly_accession request. (optional)
returned_content = 'COMPLETE' # str |  (optional) (default to 'COMPLETE')

    try:
        # Assembly descriptions by assembly accession
        api_response = api_instance.assembly_descriptors_by_accession(assembly_accession, limit=limit, filters_refseq_only=filters_refseq_only, tax_exact_match=tax_exact_match, returned_content=returned_content)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling AssemblyDatasetDescriptorsApi->assembly_descriptors_by_accession: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assembly_accession** | **str**| NCBI Assembly accession | 
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
> V1alpha1AssemblyDatasetDescriptors assembly_descriptors_by_organism(tax_name, limit=limit, filters_refseq_only=filters_refseq_only, tax_exact_match=tax_exact_match, returned_content=returned_content)

Assembly descriptions by taxonomic name (scientific or common name at any tax rank)

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
    api_instance = ncbi.datasets.AssemblyDatasetDescriptorsApi(api_client)
    tax_name = 'tax_name_example' # str | 
limit = 'limit_example' # str | Limit the number of returned results (\"all\", \"none\", otherwise an integer value). (optional)
filters_refseq_only = True # bool | If true, only return RefSeq (GCF_) assemblies. (optional)
tax_exact_match = True # bool | If true, only return assemblies with the given NCBI Taxonomy ID, or name. Otherwise, assemblies from taxonomy subtree are included, too. Ignored for assembly_accession request. (optional)
returned_content = 'COMPLETE' # str |  (optional) (default to 'COMPLETE')

    try:
        # Assembly descriptions by taxonomic name (scientific or common name at any tax rank)
        api_response = api_instance.assembly_descriptors_by_organism(tax_name, limit=limit, filters_refseq_only=filters_refseq_only, tax_exact_match=tax_exact_match, returned_content=returned_content)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling AssemblyDatasetDescriptorsApi->assembly_descriptors_by_organism: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tax_name** | **str**|  | 
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
> V1alpha1AssemblyDatasetDescriptors assembly_descriptors_by_taxid(tax_id, limit=limit, filters_refseq_only=filters_refseq_only, tax_exact_match=tax_exact_match, returned_content=returned_content)

Assembly descriptions by taxonomy ID

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
    api_instance = ncbi.datasets.AssemblyDatasetDescriptorsApi(api_client)
    tax_id = 56 # int | NCBI Taxonomy ID
limit = 'limit_example' # str | Limit the number of returned results (\"all\", \"none\", otherwise an integer value). (optional)
filters_refseq_only = True # bool | If true, only return RefSeq (GCF_) assemblies. (optional)
tax_exact_match = True # bool | If true, only return assemblies with the given NCBI Taxonomy ID, or name. Otherwise, assemblies from taxonomy subtree are included, too. Ignored for assembly_accession request. (optional)
returned_content = 'COMPLETE' # str |  (optional) (default to 'COMPLETE')

    try:
        # Assembly descriptions by taxonomy ID
        api_response = api_instance.assembly_descriptors_by_taxid(tax_id, limit=limit, filters_refseq_only=filters_refseq_only, tax_exact_match=tax_exact_match, returned_content=returned_content)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling AssemblyDatasetDescriptorsApi->assembly_descriptors_by_taxid: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tax_id** | **int**| NCBI Taxonomy ID | 
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

# **genome_summary_by_accession**
> V1alpha1DownloadSummary genome_summary_by_accession(assembly_accessions, chromosomes=chromosomes, include_sequence=include_sequence, include_annotation_type=include_annotation_type, hydrated=hydrated)

Summary of assembly dataset, including options to download package

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
    api_instance = ncbi.datasets.AssemblyDatasetDescriptorsApi(api_client)
    assembly_accessions = ['assembly_accessions_example'] # list[str] | Use 'add item' to include multiple assembly accessions.
chromosomes = ['chromosomes_example'] # list[str] | The default setting is all chromosome. Specify individual chromosome by string (1,2,MT or chr1,chr2.chrMT). Unplaced sequences are treated like their own chromosome ('Un'). The filter only applies to fasta sequence. (optional)
include_sequence = True # bool |  (optional)
include_annotation_type = ['include_annotation_type_example'] # list[str] |  (optional)
hydrated = 'FULLY_HYDRATED' # str |  - FULLY_HYDRATED: By default, supply a fully hydrated bag. (optional) (default to 'FULLY_HYDRATED')

    try:
        # Summary of assembly dataset, including options to download package
        api_response = api_instance.genome_summary_by_accession(assembly_accessions, chromosomes=chromosomes, include_sequence=include_sequence, include_annotation_type=include_annotation_type, hydrated=hydrated)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling AssemblyDatasetDescriptorsApi->genome_summary_by_accession: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assembly_accessions** | [**list[str]**](str.md)| Use &#39;add item&#39; to include multiple assembly accessions. | 
 **chromosomes** | [**list[str]**](str.md)| The default setting is all chromosome. Specify individual chromosome by string (1,2,MT or chr1,chr2.chrMT). Unplaced sequences are treated like their own chromosome (&#39;Un&#39;). The filter only applies to fasta sequence. | [optional] 
 **include_sequence** | **bool**|  | [optional] 
 **include_annotation_type** | [**list[str]**](str.md)|  | [optional] 
 **hydrated** | **str**|  - FULLY_HYDRATED: By default, supply a fully hydrated bag. | [optional] [default to &#39;FULLY_HYDRATED&#39;]

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

# **genome_tax_name_query**
> V1alpha1SciNameAndIds genome_tax_name_query(organism_query)

Retrieve list of taxonomy names and is for OrganismQuery

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
    api_instance = ncbi.datasets.AssemblyDatasetDescriptorsApi(api_client)
    organism_query = 'organism_query_example' # str | Prefix of organism name (common or scientific) to search

    try:
        # Retrieve list of taxonomy names and is for OrganismQuery
        api_response = api_instance.genome_tax_name_query(organism_query)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling AssemblyDatasetDescriptorsApi->genome_tax_name_query: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organism_query** | **str**| Prefix of organism name (common or scientific) to search | 

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **genome_tax_tree**
> V1alpha1Organism genome_tax_tree(tax_token)

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
    api_instance = ncbi.datasets.AssemblyDatasetDescriptorsApi(api_client)
    tax_token = 'tax_token_example' # str | 

    try:
        # Retrieve tax tree
        api_response = api_instance.genome_tax_tree(tax_token)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling AssemblyDatasetDescriptorsApi->genome_tax_tree: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tax_token** | **str**|  | 

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

