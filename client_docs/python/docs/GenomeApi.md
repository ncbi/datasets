# ncbi.datasets.openapi.GenomeApi

All URIs are relative to *https://api.ncbi.nlm.nih.gov/datasets/v1alpha*

Method | HTTP request | Description
------------- | ------------- | -------------
[**assembly_descriptors_by_accessions**](GenomeApi.md#assembly_descriptors_by_accessions) | **GET** /genome/accession/{accessions} | Get genome metadata by accession
[**assembly_descriptors_by_bioproject**](GenomeApi.md#assembly_descriptors_by_bioproject) | **GET** /genome/bioproject/{accessions} | Get genome metadata by bioproject accession
[**assembly_descriptors_by_taxon**](GenomeApi.md#assembly_descriptors_by_taxon) | **GET** /genome/taxon/{taxon} | Get genome metadata by taxonomic identifier
[**check_assembly_availability**](GenomeApi.md#check_assembly_availability) | **GET** /genome/accession/{accessions}/check | Check the validity of genome accessions
[**check_assembly_availability_post**](GenomeApi.md#check_assembly_availability_post) | **POST** /genome/check | Check the validity of many genome accessions in a single request
[**download_assembly_package**](GenomeApi.md#download_assembly_package) | **GET** /genome/accession/{accessions}/download | Get a genome dataset by accession
[**download_assembly_package_post**](GenomeApi.md#download_assembly_package_post) | **POST** /genome/download | Get a genome dataset by post
[**genome_download_summary**](GenomeApi.md#genome_download_summary) | **GET** /genome/accession/{accessions}/download_summary | Preview genome dataset download
[**genome_download_summary_by_post**](GenomeApi.md#genome_download_summary_by_post) | **POST** /genome/download_summary | Preview genome dataset download by POST
[**genome_metadata_by_post**](GenomeApi.md#genome_metadata_by_post) | **POST** /genome | Get genome metadata by accession
[**genome_tax_name_query**](GenomeApi.md#genome_tax_name_query) | **GET** /genome/taxon_suggest/{taxon_query} | Get a list of taxonomy names and IDs found in the assembly dataset given a partial taxonomic name.
[**genome_tax_tree**](GenomeApi.md#genome_tax_tree) | **GET** /genome/taxon/{taxon}/tree | Get a taxonomic subtree by taxonomic identifier


# **assembly_descriptors_by_accessions**
> V1alpha1AssemblyMetadata assembly_descriptors_by_accessions(accessions, filters_reference_only=filters_reference_only, filters_assembly_source=filters_assembly_source, filters_has_annotation=filters_has_annotation, filters_assembly_level=filters_assembly_level, filters_first_release_date=filters_first_release_date, filters_last_release_date=filters_last_release_date, filters_search_text=filters_search_text, page_size=page_size, page_token=page_token)

Get genome metadata by accession

Get detailed metadata for assembled genomes by accession in a JSON output format.

### Example

* Api Key Authentication (ApiKeyAuthHeader):
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

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: ApiKeyAuthHeader
configuration = ncbi.datasets.openapi.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1alpha",
    api_key = {
        'api-key': 'YOUR_API_KEY'
    }
)
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['api-key'] = 'Bearer'

# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = ncbi.datasets.openapi.GenomeApi(api_client)
    accessions = ['accessions_example'] # list[str] | 
filters_reference_only = True # bool | If true, only return reference and representative (GCF_ and GCA_) genome assemblies. (optional)
filters_assembly_source = 'all' # str | Return only RefSeq (GCF_) or GenBank (GCA_) genome assemblies. (optional) (default to 'all')
filters_has_annotation = True # bool | Return only annotated genome assemblies. (optional)
filters_assembly_level = ['filters_assembly_level_example'] # list[str] | Only return genome assemblies that have one of the specified assembly levels. (optional)
filters_first_release_date = '2013-10-20T19:20:30+01:00' # datetime | Only return genome assemblies that were released on or after the specified date. (optional)
filters_last_release_date = '2013-10-20T19:20:30+01:00' # datetime | Only return genome assemblies that were released on or before to the specified date. (optional)
filters_search_text = ['filters_search_text_example'] # list[str] | Only return results whose fields contain the specified search terms in their taxon, infraspecific, assembly name or submitter fields. (optional)
page_size = 56 # int | The maximum number of genome assemblies to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size,  `page_token` can be used to retrieve the remaining results. (optional)
page_token = 'page_token_example' # str | A page token is returned from an `AssemblyMetadataRequest` call with more than `page_size` results. Use this token, along with the previous  `AssemblyMetadataRequest` parameters, to retrieve the next page of results. When `page_token` is empty, all results have been retrieved. (optional)

    try:
        # Get genome metadata by accession
        api_response = api_instance.assembly_descriptors_by_accessions(accessions, filters_reference_only=filters_reference_only, filters_assembly_source=filters_assembly_source, filters_has_annotation=filters_has_annotation, filters_assembly_level=filters_assembly_level, filters_first_release_date=filters_first_release_date, filters_last_release_date=filters_last_release_date, filters_search_text=filters_search_text, page_size=page_size, page_token=page_token)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling GenomeApi->assembly_descriptors_by_accessions: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessions** | [**list[str]**](str.md)|  | 
 **filters_reference_only** | **bool**| If true, only return reference and representative (GCF_ and GCA_) genome assemblies. | [optional] 
 **filters_assembly_source** | **str**| Return only RefSeq (GCF_) or GenBank (GCA_) genome assemblies. | [optional] [default to &#39;all&#39;]
 **filters_has_annotation** | **bool**| Return only annotated genome assemblies. | [optional] 
 **filters_assembly_level** | [**list[str]**](str.md)| Only return genome assemblies that have one of the specified assembly levels. | [optional] 
 **filters_first_release_date** | **datetime**| Only return genome assemblies that were released on or after the specified date. | [optional] 
 **filters_last_release_date** | **datetime**| Only return genome assemblies that were released on or before to the specified date. | [optional] 
 **filters_search_text** | [**list[str]**](str.md)| Only return results whose fields contain the specified search terms in their taxon, infraspecific, assembly name or submitter fields. | [optional] 
 **page_size** | **int**| The maximum number of genome assemblies to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size,  &#x60;page_token&#x60; can be used to retrieve the remaining results. | [optional] 
 **page_token** | **str**| A page token is returned from an &#x60;AssemblyMetadataRequest&#x60; call with more than &#x60;page_size&#x60; results. Use this token, along with the previous  &#x60;AssemblyMetadataRequest&#x60; parameters, to retrieve the next page of results. When &#x60;page_token&#x60; is empty, all results have been retrieved. | [optional] 

### Return type

[**V1alpha1AssemblyMetadata**](V1alpha1AssemblyMetadata.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **assembly_descriptors_by_bioproject**
> V1alpha1AssemblyMetadata assembly_descriptors_by_bioproject(accessions, filters_reference_only=filters_reference_only, filters_assembly_source=filters_assembly_source, filters_has_annotation=filters_has_annotation, filters_assembly_level=filters_assembly_level, filters_first_release_date=filters_first_release_date, filters_last_release_date=filters_last_release_date, filters_search_text=filters_search_text, returned_content=returned_content, page_size=page_size, page_token=page_token)

Get genome metadata by bioproject accession

Get detailed metadata for assembled genomes by bioproject accession in a JSON output format.

### Example

* Api Key Authentication (ApiKeyAuthHeader):
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

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: ApiKeyAuthHeader
configuration = ncbi.datasets.openapi.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1alpha",
    api_key = {
        'api-key': 'YOUR_API_KEY'
    }
)
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['api-key'] = 'Bearer'

# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = ncbi.datasets.openapi.GenomeApi(api_client)
    accessions = ['accessions_example'] # list[str] | 
filters_reference_only = True # bool | If true, only return reference and representative (GCF_ and GCA_) genome assemblies. (optional)
filters_assembly_source = 'all' # str | Return only RefSeq (GCF_) or GenBank (GCA_) genome assemblies. (optional) (default to 'all')
filters_has_annotation = True # bool | Return only annotated genome assemblies. (optional)
filters_assembly_level = ['filters_assembly_level_example'] # list[str] | Only return genome assemblies that have one of the specified assembly levels. (optional)
filters_first_release_date = '2013-10-20T19:20:30+01:00' # datetime | Only return genome assemblies that were released on or after the specified date. (optional)
filters_last_release_date = '2013-10-20T19:20:30+01:00' # datetime | Only return genome assemblies that were released on or before to the specified date. (optional)
filters_search_text = ['filters_search_text_example'] # list[str] | Only return results whose fields contain the specified search terms in their taxon, infraspecific, assembly name or submitter fields. (optional)
returned_content = 'COMPLETE' # str | Return either assembly accessions, or entire assembly-metadata records. (optional) (default to 'COMPLETE')
page_size = 56 # int | The maximum number of genome assemblies to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size,  `page_token` can be used to retrieve the remaining results. (optional)
page_token = 'page_token_example' # str | A page token is returned from an `AssemblyMetadataRequest` call with more than `page_size` results. Use this token, along with the previous  `AssemblyMetadataRequest` parameters, to retrieve the next page of results. When `page_token` is empty, all results have been retrieved. (optional)

    try:
        # Get genome metadata by bioproject accession
        api_response = api_instance.assembly_descriptors_by_bioproject(accessions, filters_reference_only=filters_reference_only, filters_assembly_source=filters_assembly_source, filters_has_annotation=filters_has_annotation, filters_assembly_level=filters_assembly_level, filters_first_release_date=filters_first_release_date, filters_last_release_date=filters_last_release_date, filters_search_text=filters_search_text, returned_content=returned_content, page_size=page_size, page_token=page_token)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling GenomeApi->assembly_descriptors_by_bioproject: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessions** | [**list[str]**](str.md)|  | 
 **filters_reference_only** | **bool**| If true, only return reference and representative (GCF_ and GCA_) genome assemblies. | [optional] 
 **filters_assembly_source** | **str**| Return only RefSeq (GCF_) or GenBank (GCA_) genome assemblies. | [optional] [default to &#39;all&#39;]
 **filters_has_annotation** | **bool**| Return only annotated genome assemblies. | [optional] 
 **filters_assembly_level** | [**list[str]**](str.md)| Only return genome assemblies that have one of the specified assembly levels. | [optional] 
 **filters_first_release_date** | **datetime**| Only return genome assemblies that were released on or after the specified date. | [optional] 
 **filters_last_release_date** | **datetime**| Only return genome assemblies that were released on or before to the specified date. | [optional] 
 **filters_search_text** | [**list[str]**](str.md)| Only return results whose fields contain the specified search terms in their taxon, infraspecific, assembly name or submitter fields. | [optional] 
 **returned_content** | **str**| Return either assembly accessions, or entire assembly-metadata records. | [optional] [default to &#39;COMPLETE&#39;]
 **page_size** | **int**| The maximum number of genome assemblies to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size,  &#x60;page_token&#x60; can be used to retrieve the remaining results. | [optional] 
 **page_token** | **str**| A page token is returned from an &#x60;AssemblyMetadataRequest&#x60; call with more than &#x60;page_size&#x60; results. Use this token, along with the previous  &#x60;AssemblyMetadataRequest&#x60; parameters, to retrieve the next page of results. When &#x60;page_token&#x60; is empty, all results have been retrieved. | [optional] 

### Return type

[**V1alpha1AssemblyMetadata**](V1alpha1AssemblyMetadata.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **assembly_descriptors_by_taxon**
> V1alpha1AssemblyMetadata assembly_descriptors_by_taxon(taxon, filters_reference_only=filters_reference_only, filters_assembly_source=filters_assembly_source, filters_has_annotation=filters_has_annotation, filters_assembly_level=filters_assembly_level, filters_first_release_date=filters_first_release_date, filters_last_release_date=filters_last_release_date, filters_search_text=filters_search_text, tax_exact_match=tax_exact_match, returned_content=returned_content, page_size=page_size, page_token=page_token)

Get genome metadata by taxonomic identifier

Get detailed metadata on all assembled genomes for a specified NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank

### Example

* Api Key Authentication (ApiKeyAuthHeader):
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

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: ApiKeyAuthHeader
configuration = ncbi.datasets.openapi.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1alpha",
    api_key = {
        'api-key': 'YOUR_API_KEY'
    }
)
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['api-key'] = 'Bearer'

# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = ncbi.datasets.openapi.GenomeApi(api_client)
    taxon = 'taxon_example' # str | NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank
filters_reference_only = True # bool | If true, only return reference and representative (GCF_ and GCA_) genome assemblies. (optional)
filters_assembly_source = 'all' # str | Return only RefSeq (GCF_) or GenBank (GCA_) genome assemblies. (optional) (default to 'all')
filters_has_annotation = True # bool | Return only annotated genome assemblies. (optional)
filters_assembly_level = ['filters_assembly_level_example'] # list[str] | Only return genome assemblies that have one of the specified assembly levels. (optional)
filters_first_release_date = '2013-10-20T19:20:30+01:00' # datetime | Only return genome assemblies that were released on or after the specified date. (optional)
filters_last_release_date = '2013-10-20T19:20:30+01:00' # datetime | Only return genome assemblies that were released on or before to the specified date. (optional)
filters_search_text = ['filters_search_text_example'] # list[str] | Only return results whose fields contain the specified search terms in their taxon, infraspecific, assembly name or submitter fields. (optional)
tax_exact_match = True # bool | If true, only return assemblies with the given NCBI Taxonomy ID, or name. Otherwise, assemblies from taxonomy subtree are included, too. Ignored for assembly_accession request. (optional)
returned_content = 'COMPLETE' # str | Return either assembly accessions, or entire assembly-metadata records. (optional) (default to 'COMPLETE')
page_size = 56 # int | The maximum number of genome assemblies to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size,  `page_token` can be used to retrieve the remaining results. (optional)
page_token = 'page_token_example' # str | A page token is returned from an `AssemblyMetadataRequest` call with more than `page_size` results. Use this token, along with the previous  `AssemblyMetadataRequest` parameters, to retrieve the next page of results. When `page_token` is empty, all results have been retrieved. (optional)

    try:
        # Get genome metadata by taxonomic identifier
        api_response = api_instance.assembly_descriptors_by_taxon(taxon, filters_reference_only=filters_reference_only, filters_assembly_source=filters_assembly_source, filters_has_annotation=filters_has_annotation, filters_assembly_level=filters_assembly_level, filters_first_release_date=filters_first_release_date, filters_last_release_date=filters_last_release_date, filters_search_text=filters_search_text, tax_exact_match=tax_exact_match, returned_content=returned_content, page_size=page_size, page_token=page_token)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling GenomeApi->assembly_descriptors_by_taxon: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taxon** | **str**| NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank | 
 **filters_reference_only** | **bool**| If true, only return reference and representative (GCF_ and GCA_) genome assemblies. | [optional] 
 **filters_assembly_source** | **str**| Return only RefSeq (GCF_) or GenBank (GCA_) genome assemblies. | [optional] [default to &#39;all&#39;]
 **filters_has_annotation** | **bool**| Return only annotated genome assemblies. | [optional] 
 **filters_assembly_level** | [**list[str]**](str.md)| Only return genome assemblies that have one of the specified assembly levels. | [optional] 
 **filters_first_release_date** | **datetime**| Only return genome assemblies that were released on or after the specified date. | [optional] 
 **filters_last_release_date** | **datetime**| Only return genome assemblies that were released on or before to the specified date. | [optional] 
 **filters_search_text** | [**list[str]**](str.md)| Only return results whose fields contain the specified search terms in their taxon, infraspecific, assembly name or submitter fields. | [optional] 
 **tax_exact_match** | **bool**| If true, only return assemblies with the given NCBI Taxonomy ID, or name. Otherwise, assemblies from taxonomy subtree are included, too. Ignored for assembly_accession request. | [optional] 
 **returned_content** | **str**| Return either assembly accessions, or entire assembly-metadata records. | [optional] [default to &#39;COMPLETE&#39;]
 **page_size** | **int**| The maximum number of genome assemblies to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size,  &#x60;page_token&#x60; can be used to retrieve the remaining results. | [optional] 
 **page_token** | **str**| A page token is returned from an &#x60;AssemblyMetadataRequest&#x60; call with more than &#x60;page_size&#x60; results. Use this token, along with the previous  &#x60;AssemblyMetadataRequest&#x60; parameters, to retrieve the next page of results. When &#x60;page_token&#x60; is empty, all results have been retrieved. | [optional] 

### Return type

[**V1alpha1AssemblyMetadata**](V1alpha1AssemblyMetadata.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **check_assembly_availability**
> V1alpha1AssemblyDatasetAvailability check_assembly_availability(accessions)

Check the validity of genome accessions

The 'GET' version of check is limited by the size of the GET URL (2KB, which works out to about 140 genomic accessions).  The POST operation is provided to allow users to supply a larger number of accessions in a single request.

### Example

* Api Key Authentication (ApiKeyAuthHeader):
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

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: ApiKeyAuthHeader
configuration = ncbi.datasets.openapi.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1alpha",
    api_key = {
        'api-key': 'YOUR_API_KEY'
    }
)
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['api-key'] = 'Bearer'

# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = ncbi.datasets.openapi.GenomeApi(api_client)
    accessions = ['accessions_example'] # list[str] | NCBI genome assembly accessions

    try:
        # Check the validity of genome accessions
        api_response = api_instance.check_assembly_availability(accessions)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling GenomeApi->check_assembly_availability: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessions** | [**list[str]**](str.md)| NCBI genome assembly accessions | 

### Return type

[**V1alpha1AssemblyDatasetAvailability**](V1alpha1AssemblyDatasetAvailability.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **check_assembly_availability_post**
> V1alpha1AssemblyDatasetAvailability check_assembly_availability_post(body)

Check the validity of many genome accessions in a single request

The 'GET' version of check is limited by the size of the GET URL (2KB, which works out to about 140 genomic accessions).  The POST operation is provided to allow users to supply a larger number of accessions in a single request.

### Example

* Api Key Authentication (ApiKeyAuthHeader):
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

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: ApiKeyAuthHeader
configuration = ncbi.datasets.openapi.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1alpha",
    api_key = {
        'api-key': 'YOUR_API_KEY'
    }
)
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['api-key'] = 'Bearer'

# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = ncbi.datasets.openapi.GenomeApi(api_client)
    body = ncbi.datasets.openapi.V1alpha1AssemblyDatasetRequest() # V1alpha1AssemblyDatasetRequest | 

    try:
        # Check the validity of many genome accessions in a single request
        api_response = api_instance.check_assembly_availability_post(body)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling GenomeApi->check_assembly_availability_post: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**V1alpha1AssemblyDatasetRequest**](V1alpha1AssemblyDatasetRequest.md)|  | 

### Return type

[**V1alpha1AssemblyDatasetAvailability**](V1alpha1AssemblyDatasetAvailability.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **download_assembly_package**
> file download_assembly_package(accessions, chromosomes=chromosomes, exclude_sequence=exclude_sequence, include_annotation_type=include_annotation_type, hydrated=hydrated, filename=filename)

Get a genome dataset by accession

Download a genome dataset including fasta sequence, annotation and a detailed data report by accession.

### Example

* Api Key Authentication (ApiKeyAuthHeader):
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

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: ApiKeyAuthHeader
configuration = ncbi.datasets.openapi.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1alpha",
    api_key = {
        'api-key': 'YOUR_API_KEY'
    }
)
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['api-key'] = 'Bearer'

# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = ncbi.datasets.openapi.GenomeApi(api_client)
    accessions = ['accessions_example'] # list[str] | NCBI genome assembly accessions
chromosomes = ['chromosomes_example'] # list[str] | The default setting is all chromosome. Specify individual chromosome by string (1,2,MT or chr1,chr2.chrMT). Unplaced sequences are treated like their own chromosome ('Un'). The filter only applies to fasta sequence. (optional)
exclude_sequence = True # bool | Set to true to omit the genomic sequence. (optional)
include_annotation_type = ['include_annotation_type_example'] # list[str] | Select additional types of annotation to include in the data package.  If unset, no annotation is provided. (optional)
hydrated = 'FULLY_HYDRATED' # str | Set to DATA_REPORT_ONLY, to only retrieve data-reports. (optional) (default to 'FULLY_HYDRATED')
filename = 'ncbi_dataset.zip' # str | Output file name. (optional) (default to 'ncbi_dataset.zip')

    try:
        # Get a genome dataset by accession
        api_response = api_instance.download_assembly_package(accessions, chromosomes=chromosomes, exclude_sequence=exclude_sequence, include_annotation_type=include_annotation_type, hydrated=hydrated, filename=filename)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling GenomeApi->download_assembly_package: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessions** | [**list[str]**](str.md)| NCBI genome assembly accessions | 
 **chromosomes** | [**list[str]**](str.md)| The default setting is all chromosome. Specify individual chromosome by string (1,2,MT or chr1,chr2.chrMT). Unplaced sequences are treated like their own chromosome (&#39;Un&#39;). The filter only applies to fasta sequence. | [optional] 
 **exclude_sequence** | **bool**| Set to true to omit the genomic sequence. | [optional] 
 **include_annotation_type** | [**list[str]**](str.md)| Select additional types of annotation to include in the data package.  If unset, no annotation is provided. | [optional] 
 **hydrated** | **str**| Set to DATA_REPORT_ONLY, to only retrieve data-reports. | [optional] [default to &#39;FULLY_HYDRATED&#39;]
 **filename** | **str**| Output file name. | [optional] [default to &#39;ncbi_dataset.zip&#39;]

### Return type

**file**

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/zip

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Download selected genome assemblies and associated annotation data as a zip file. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **download_assembly_package_post**
> file download_assembly_package_post(body, filename=filename)

Get a genome dataset by post

The 'GET' version of download is limited by the size of the GET URL (2KB, which works out to about 140 genomic accessions).  The POST operation is provided to allow users to supply a larger number of accessions in a single request.

### Example

* Api Key Authentication (ApiKeyAuthHeader):
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

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: ApiKeyAuthHeader
configuration = ncbi.datasets.openapi.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1alpha",
    api_key = {
        'api-key': 'YOUR_API_KEY'
    }
)
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['api-key'] = 'Bearer'

# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = ncbi.datasets.openapi.GenomeApi(api_client)
    body = ncbi.datasets.openapi.V1alpha1AssemblyDatasetRequest() # V1alpha1AssemblyDatasetRequest | 
filename = 'ncbi_dataset.zip' # str | Output file name. (optional) (default to 'ncbi_dataset.zip')

    try:
        # Get a genome dataset by post
        api_response = api_instance.download_assembly_package_post(body, filename=filename)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling GenomeApi->download_assembly_package_post: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**V1alpha1AssemblyDatasetRequest**](V1alpha1AssemblyDatasetRequest.md)|  | 
 **filename** | **str**| Output file name. | [optional] [default to &#39;ncbi_dataset.zip&#39;]

### Return type

**file**

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/zip

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Download selected genome assemblies and associated annotation data as a zip file. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **genome_download_summary**
> V1alpha1DownloadSummary genome_download_summary(accessions, chromosomes=chromosomes, exclude_sequence=exclude_sequence, include_annotation_type=include_annotation_type)

Preview genome dataset download

Get a download summary by accession in a JSON output format.

### Example

* Api Key Authentication (ApiKeyAuthHeader):
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

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: ApiKeyAuthHeader
configuration = ncbi.datasets.openapi.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1alpha",
    api_key = {
        'api-key': 'YOUR_API_KEY'
    }
)
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['api-key'] = 'Bearer'

# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = ncbi.datasets.openapi.GenomeApi(api_client)
    accessions = ['accessions_example'] # list[str] | NCBI genome assembly accessions
chromosomes = ['chromosomes_example'] # list[str] | The default setting is all chromosome. Specify individual chromosome by string (1,2,MT or chr1,chr2.chrMT). Unplaced sequences are treated like their own chromosome ('Un'). The filter only applies to fasta sequence. (optional)
exclude_sequence = True # bool | Set to true to omit the genomic sequence. (optional)
include_annotation_type = ['include_annotation_type_example'] # list[str] | Select additional types of annotation to include in the data package.  If unset, no annotation is provided. (optional)

    try:
        # Preview genome dataset download
        api_response = api_instance.genome_download_summary(accessions, chromosomes=chromosomes, exclude_sequence=exclude_sequence, include_annotation_type=include_annotation_type)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling GenomeApi->genome_download_summary: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessions** | [**list[str]**](str.md)| NCBI genome assembly accessions | 
 **chromosomes** | [**list[str]**](str.md)| The default setting is all chromosome. Specify individual chromosome by string (1,2,MT or chr1,chr2.chrMT). Unplaced sequences are treated like their own chromosome (&#39;Un&#39;). The filter only applies to fasta sequence. | [optional] 
 **exclude_sequence** | **bool**| Set to true to omit the genomic sequence. | [optional] 
 **include_annotation_type** | [**list[str]**](str.md)| Select additional types of annotation to include in the data package.  If unset, no annotation is provided. | [optional] 

### Return type

[**V1alpha1DownloadSummary**](V1alpha1DownloadSummary.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **genome_download_summary_by_post**
> V1alpha1DownloadSummary genome_download_summary_by_post(body)

Preview genome dataset download by POST

The 'GET' version of download summary is limited by the size of the GET URL (2KB, which works out to about 140 genomic accessions).  The POST operation is provided to allow users to supply a larger number of accessions in a single request.

### Example

* Api Key Authentication (ApiKeyAuthHeader):
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

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: ApiKeyAuthHeader
configuration = ncbi.datasets.openapi.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1alpha",
    api_key = {
        'api-key': 'YOUR_API_KEY'
    }
)
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['api-key'] = 'Bearer'

# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = ncbi.datasets.openapi.GenomeApi(api_client)
    body = ncbi.datasets.openapi.V1alpha1AssemblyDatasetRequest() # V1alpha1AssemblyDatasetRequest | 

    try:
        # Preview genome dataset download by POST
        api_response = api_instance.genome_download_summary_by_post(body)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling GenomeApi->genome_download_summary_by_post: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**V1alpha1AssemblyDatasetRequest**](V1alpha1AssemblyDatasetRequest.md)|  | 

### Return type

[**V1alpha1DownloadSummary**](V1alpha1DownloadSummary.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **genome_metadata_by_post**
> V1alpha1AssemblyMetadata genome_metadata_by_post(body)

Get genome metadata by accession

Get detailed metadata for assembled genomes by accession in a JSON output format.

### Example

* Api Key Authentication (ApiKeyAuthHeader):
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

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: ApiKeyAuthHeader
configuration = ncbi.datasets.openapi.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1alpha",
    api_key = {
        'api-key': 'YOUR_API_KEY'
    }
)
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['api-key'] = 'Bearer'

# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = ncbi.datasets.openapi.GenomeApi(api_client)
    body = ncbi.datasets.openapi.V1alpha1AssemblyMetadataRequest() # V1alpha1AssemblyMetadataRequest | 

    try:
        # Get genome metadata by accession
        api_response = api_instance.genome_metadata_by_post(body)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling GenomeApi->genome_metadata_by_post: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**V1alpha1AssemblyMetadataRequest**](V1alpha1AssemblyMetadataRequest.md)|  | 

### Return type

[**V1alpha1AssemblyMetadata**](V1alpha1AssemblyMetadata.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **genome_tax_name_query**
> V1alpha1SciNameAndIds genome_tax_name_query(taxon_query)

Get a list of taxonomy names and IDs found in the assembly dataset given a partial taxonomic name.

This endpoint retrieves a list of taxonomy names and IDs found in the assembly dataset given a partial taxonomic name of any rank.

### Example

* Api Key Authentication (ApiKeyAuthHeader):
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

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: ApiKeyAuthHeader
configuration = ncbi.datasets.openapi.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1alpha",
    api_key = {
        'api-key': 'YOUR_API_KEY'
    }
)
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['api-key'] = 'Bearer'

# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = ncbi.datasets.openapi.GenomeApi(api_client)
    taxon_query = 'taxon_query_example' # str | NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank

    try:
        # Get a list of taxonomy names and IDs found in the assembly dataset given a partial taxonomic name.
        api_response = api_instance.genome_tax_name_query(taxon_query)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling GenomeApi->genome_tax_name_query: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taxon_query** | **str**| NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank | 

### Return type

[**V1alpha1SciNameAndIds**](V1alpha1SciNameAndIds.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **genome_tax_tree**
> V1alpha1Organism genome_tax_tree(taxon)

Get a taxonomic subtree by taxonomic identifier

Using a NCBI Taxonomy ID or name (common or scientific) at any rank, get a subtree filtered for species with assembled genomes

### Example

* Api Key Authentication (ApiKeyAuthHeader):
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

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: ApiKeyAuthHeader
configuration = ncbi.datasets.openapi.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1alpha",
    api_key = {
        'api-key': 'YOUR_API_KEY'
    }
)
# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['api-key'] = 'Bearer'

# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = ncbi.datasets.openapi.GenomeApi(api_client)
    taxon = 'taxon_example' # str | NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank

    try:
        # Get a taxonomic subtree by taxonomic identifier
        api_response = api_instance.genome_tax_tree(taxon)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling GenomeApi->genome_tax_tree: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taxon** | **str**| NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank | 

### Return type

[**V1alpha1Organism**](V1alpha1Organism.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

