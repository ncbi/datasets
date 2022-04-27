# ncbi.datasets.openapi.GenomeApi

All URIs are relative to *https://api.ncbi.nlm.nih.gov/datasets/v1*

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
[**genome_metadata_by_post**](GenomeApi.md#genome_metadata_by_post) | **POST** /genome | Get genome metadata by variety of identifiers
[**genome_tax_name_query**](GenomeApi.md#genome_tax_name_query) | **GET** /genome/taxon_suggest/{taxon_query} | Get a list of taxonomy names and IDs found in the assembly dataset given a partial taxonomic name
[**genome_tax_tree**](GenomeApi.md#genome_tax_tree) | **GET** /genome/taxon/{taxon}/tree | Get a taxonomic subtree by taxonomic identifier


# **assembly_descriptors_by_accessions**
> V1AssemblyMetadata assembly_descriptors_by_accessions(accessions)

Get genome metadata by accession

Get detailed metadata for assembled genomes by accession in a JSON output format.

### Example

* Api Key Authentication (ApiKeyAuthHeader):

```python
import time
import ncbi.datasets.openapi
from ncbi.datasets.openapi.api import genome_api
from ncbi.datasets.openapi.model.v1_assembly_metadata import V1AssemblyMetadata
from ncbi.datasets.openapi.model.v1_assembly_dataset_descriptors_filter_assembly_level import V1AssemblyDatasetDescriptorsFilterAssemblyLevel
from ncbi.datasets.openapi.model.rpc_status import RpcStatus
from ncbi.datasets.openapi.model.v1_assembly_dataset_descriptors_filter_assembly_source import V1AssemblyDatasetDescriptorsFilterAssemblySource
from pprint import pprint
# Defining the host is optional and defaults to https://api.ncbi.nlm.nih.gov/datasets/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = ncbi.datasets.openapi.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: ApiKeyAuthHeader
configuration.api_key['ApiKeyAuthHeader'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['ApiKeyAuthHeader'] = 'Bearer'

# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = genome_api.GenomeApi(api_client)
    accessions = ["GCF_000001405.40","GCF_000001635.27"] # [str] | 
    filters_reference_only = True # bool | If true, only return reference and representative (GCF_ and GCA_) genome assemblies. (optional) if omitted the server will use the default value of False
    filters_assembly_source = V1AssemblyDatasetDescriptorsFilterAssemblySource("refseq") # V1AssemblyDatasetDescriptorsFilterAssemblySource | Return only RefSeq (GCF_) or GenBank (GCA_) genome assemblies (optional)
    filters_has_annotation = True # bool | Return only annotated genome assemblies (optional) if omitted the server will use the default value of False
    filters_assembly_level = [
        V1AssemblyDatasetDescriptorsFilterAssemblyLevel("["chromosome","complete_genome"]"),
    ] # [V1AssemblyDatasetDescriptorsFilterAssemblyLevel] | Only return genome assemblies that have one of the specified assembly levels. By default, do not filter. (optional)
    filters_first_release_date = dateutil_parser('01/10/2015') # datetime | Only return genome assemblies that were released on or after the specified date By default, do not filter. (optional)
    filters_last_release_date = dateutil_parser('01/10/2021') # datetime | Only return genome assemblies that were released on or before to the specified date By default, do not filter. (optional)
    filters_search_text = [
        "Genome Reference Consortium",
    ] # [str] | Only return results whose fields contain the specified search terms in their taxon, infraspecific, assembly name or submitter fields By default, do not filter (optional)
    page_size = 20 # int | The maximum number of genome assemblies to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, `page_token` can be used to retrieve the remaining results. (optional) if omitted the server will use the default value of 20
    page_token = "page_token_example" # str | A page token is returned from an `AssemblyMetadataRequest` call with more than `page_size` results. Use this token, along with the previous `AssemblyMetadataRequest` parameters, to retrieve the next page of results. When `page_token` is empty, all results have been retrieved. (optional)

    # example passing only required values which don't have defaults set
    try:
        # Get genome metadata by accession
        api_response = api_instance.assembly_descriptors_by_accessions(accessions)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling GenomeApi->assembly_descriptors_by_accessions: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get genome metadata by accession
        api_response = api_instance.assembly_descriptors_by_accessions(accessions, filters_reference_only=filters_reference_only, filters_assembly_source=filters_assembly_source, filters_has_annotation=filters_has_annotation, filters_assembly_level=filters_assembly_level, filters_first_release_date=filters_first_release_date, filters_last_release_date=filters_last_release_date, filters_search_text=filters_search_text, page_size=page_size, page_token=page_token)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling GenomeApi->assembly_descriptors_by_accessions: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessions** | **[str]**|  |
 **filters_reference_only** | **bool**| If true, only return reference and representative (GCF_ and GCA_) genome assemblies. | [optional] if omitted the server will use the default value of False
 **filters_assembly_source** | **V1AssemblyDatasetDescriptorsFilterAssemblySource**| Return only RefSeq (GCF_) or GenBank (GCA_) genome assemblies | [optional]
 **filters_has_annotation** | **bool**| Return only annotated genome assemblies | [optional] if omitted the server will use the default value of False
 **filters_assembly_level** | [**[V1AssemblyDatasetDescriptorsFilterAssemblyLevel]**](V1AssemblyDatasetDescriptorsFilterAssemblyLevel.md)| Only return genome assemblies that have one of the specified assembly levels. By default, do not filter. | [optional]
 **filters_first_release_date** | **datetime**| Only return genome assemblies that were released on or after the specified date By default, do not filter. | [optional]
 **filters_last_release_date** | **datetime**| Only return genome assemblies that were released on or before to the specified date By default, do not filter. | [optional]
 **filters_search_text** | **[str]**| Only return results whose fields contain the specified search terms in their taxon, infraspecific, assembly name or submitter fields By default, do not filter | [optional]
 **page_size** | **int**| The maximum number of genome assemblies to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, &#x60;page_token&#x60; can be used to retrieve the remaining results. | [optional] if omitted the server will use the default value of 20
 **page_token** | **str**| A page token is returned from an &#x60;AssemblyMetadataRequest&#x60; call with more than &#x60;page_size&#x60; results. Use this token, along with the previous &#x60;AssemblyMetadataRequest&#x60; parameters, to retrieve the next page of results. When &#x60;page_token&#x60; is empty, all results have been retrieved. | [optional]

### Return type

[**V1AssemblyMetadata**](V1AssemblyMetadata.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **assembly_descriptors_by_bioproject**
> V1AssemblyMetadata assembly_descriptors_by_bioproject(accessions)

Get genome metadata by bioproject accession

Get detailed metadata for assembled genomes by bioproject accession in a JSON output format.

### Example

* Api Key Authentication (ApiKeyAuthHeader):

```python
import time
import ncbi.datasets.openapi
from ncbi.datasets.openapi.api import genome_api
from ncbi.datasets.openapi.model.v1_assembly_metadata_request_content_type import V1AssemblyMetadataRequestContentType
from ncbi.datasets.openapi.model.v1_assembly_metadata import V1AssemblyMetadata
from ncbi.datasets.openapi.model.v1_assembly_dataset_descriptors_filter_assembly_level import V1AssemblyDatasetDescriptorsFilterAssemblyLevel
from ncbi.datasets.openapi.model.rpc_status import RpcStatus
from ncbi.datasets.openapi.model.v1_assembly_dataset_descriptors_filter_assembly_source import V1AssemblyDatasetDescriptorsFilterAssemblySource
from pprint import pprint
# Defining the host is optional and defaults to https://api.ncbi.nlm.nih.gov/datasets/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = ncbi.datasets.openapi.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: ApiKeyAuthHeader
configuration.api_key['ApiKeyAuthHeader'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['ApiKeyAuthHeader'] = 'Bearer'

# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = genome_api.GenomeApi(api_client)
    accessions = [
        "PRJNA489243",
    ] # [str] | 
    filters_reference_only = True # bool | If true, only return reference and representative (GCF_ and GCA_) genome assemblies. (optional) if omitted the server will use the default value of False
    filters_assembly_source = V1AssemblyDatasetDescriptorsFilterAssemblySource("refseq") # V1AssemblyDatasetDescriptorsFilterAssemblySource | Return only RefSeq (GCF_) or GenBank (GCA_) genome assemblies (optional)
    filters_has_annotation = True # bool | Return only annotated genome assemblies (optional) if omitted the server will use the default value of False
    filters_assembly_level = [
        V1AssemblyDatasetDescriptorsFilterAssemblyLevel("["chromosome","complete_genome"]"),
    ] # [V1AssemblyDatasetDescriptorsFilterAssemblyLevel] | Only return genome assemblies that have one of the specified assembly levels. By default, do not filter. (optional)
    filters_first_release_date = dateutil_parser('01/10/2015') # datetime | Only return genome assemblies that were released on or after the specified date By default, do not filter. (optional)
    filters_last_release_date = dateutil_parser('01/10/2021') # datetime | Only return genome assemblies that were released on or before to the specified date By default, do not filter. (optional)
    filters_search_text = [
        "Genome Reference Consortium",
    ] # [str] | Only return results whose fields contain the specified search terms in their taxon, infraspecific, assembly name or submitter fields By default, do not filter (optional)
    returned_content = V1AssemblyMetadataRequestContentType("COMPLETE") # V1AssemblyMetadataRequestContentType | Return either assembly accessions, or entire assembly-metadata records (optional)
    page_size = 20 # int | The maximum number of genome assemblies to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, `page_token` can be used to retrieve the remaining results. (optional) if omitted the server will use the default value of 20
    page_token = "page_token_example" # str | A page token is returned from an `AssemblyMetadataRequest` call with more than `page_size` results. Use this token, along with the previous `AssemblyMetadataRequest` parameters, to retrieve the next page of results. When `page_token` is empty, all results have been retrieved. (optional)

    # example passing only required values which don't have defaults set
    try:
        # Get genome metadata by bioproject accession
        api_response = api_instance.assembly_descriptors_by_bioproject(accessions)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling GenomeApi->assembly_descriptors_by_bioproject: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get genome metadata by bioproject accession
        api_response = api_instance.assembly_descriptors_by_bioproject(accessions, filters_reference_only=filters_reference_only, filters_assembly_source=filters_assembly_source, filters_has_annotation=filters_has_annotation, filters_assembly_level=filters_assembly_level, filters_first_release_date=filters_first_release_date, filters_last_release_date=filters_last_release_date, filters_search_text=filters_search_text, returned_content=returned_content, page_size=page_size, page_token=page_token)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling GenomeApi->assembly_descriptors_by_bioproject: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessions** | **[str]**|  |
 **filters_reference_only** | **bool**| If true, only return reference and representative (GCF_ and GCA_) genome assemblies. | [optional] if omitted the server will use the default value of False
 **filters_assembly_source** | **V1AssemblyDatasetDescriptorsFilterAssemblySource**| Return only RefSeq (GCF_) or GenBank (GCA_) genome assemblies | [optional]
 **filters_has_annotation** | **bool**| Return only annotated genome assemblies | [optional] if omitted the server will use the default value of False
 **filters_assembly_level** | [**[V1AssemblyDatasetDescriptorsFilterAssemblyLevel]**](V1AssemblyDatasetDescriptorsFilterAssemblyLevel.md)| Only return genome assemblies that have one of the specified assembly levels. By default, do not filter. | [optional]
 **filters_first_release_date** | **datetime**| Only return genome assemblies that were released on or after the specified date By default, do not filter. | [optional]
 **filters_last_release_date** | **datetime**| Only return genome assemblies that were released on or before to the specified date By default, do not filter. | [optional]
 **filters_search_text** | **[str]**| Only return results whose fields contain the specified search terms in their taxon, infraspecific, assembly name or submitter fields By default, do not filter | [optional]
 **returned_content** | **V1AssemblyMetadataRequestContentType**| Return either assembly accessions, or entire assembly-metadata records | [optional]
 **page_size** | **int**| The maximum number of genome assemblies to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, &#x60;page_token&#x60; can be used to retrieve the remaining results. | [optional] if omitted the server will use the default value of 20
 **page_token** | **str**| A page token is returned from an &#x60;AssemblyMetadataRequest&#x60; call with more than &#x60;page_size&#x60; results. Use this token, along with the previous &#x60;AssemblyMetadataRequest&#x60; parameters, to retrieve the next page of results. When &#x60;page_token&#x60; is empty, all results have been retrieved. | [optional]

### Return type

[**V1AssemblyMetadata**](V1AssemblyMetadata.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **assembly_descriptors_by_taxon**
> V1AssemblyMetadata assembly_descriptors_by_taxon(taxon)

Get genome metadata by taxonomic identifier

Get detailed metadata on all assembled genomes for a specified NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank.

### Example

* Api Key Authentication (ApiKeyAuthHeader):

```python
import time
import ncbi.datasets.openapi
from ncbi.datasets.openapi.api import genome_api
from ncbi.datasets.openapi.model.v1_assembly_metadata_request_content_type import V1AssemblyMetadataRequestContentType
from ncbi.datasets.openapi.model.v1_assembly_metadata import V1AssemblyMetadata
from ncbi.datasets.openapi.model.v1_assembly_dataset_descriptors_filter_assembly_level import V1AssemblyDatasetDescriptorsFilterAssemblyLevel
from ncbi.datasets.openapi.model.rpc_status import RpcStatus
from ncbi.datasets.openapi.model.v1_assembly_dataset_descriptors_filter_assembly_source import V1AssemblyDatasetDescriptorsFilterAssemblySource
from pprint import pprint
# Defining the host is optional and defaults to https://api.ncbi.nlm.nih.gov/datasets/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = ncbi.datasets.openapi.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: ApiKeyAuthHeader
configuration.api_key['ApiKeyAuthHeader'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['ApiKeyAuthHeader'] = 'Bearer'

# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = genome_api.GenomeApi(api_client)
    taxon = "9606" # str | NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank
    filters_reference_only = True # bool | If true, only return reference and representative (GCF_ and GCA_) genome assemblies. (optional) if omitted the server will use the default value of False
    filters_assembly_source = V1AssemblyDatasetDescriptorsFilterAssemblySource("refseq") # V1AssemblyDatasetDescriptorsFilterAssemblySource | Return only RefSeq (GCF_) or GenBank (GCA_) genome assemblies (optional)
    filters_has_annotation = True # bool | Return only annotated genome assemblies (optional) if omitted the server will use the default value of False
    filters_assembly_level = [
        V1AssemblyDatasetDescriptorsFilterAssemblyLevel("["chromosome","complete_genome"]"),
    ] # [V1AssemblyDatasetDescriptorsFilterAssemblyLevel] | Only return genome assemblies that have one of the specified assembly levels. By default, do not filter. (optional)
    filters_first_release_date = dateutil_parser('01/10/2015') # datetime | Only return genome assemblies that were released on or after the specified date By default, do not filter. (optional)
    filters_last_release_date = dateutil_parser('01/10/2021') # datetime | Only return genome assemblies that were released on or before to the specified date By default, do not filter. (optional)
    filters_search_text = [
        "Genome Reference Consortium",
    ] # [str] | Only return results whose fields contain the specified search terms in their taxon, infraspecific, assembly name or submitter fields By default, do not filter (optional)
    tax_exact_match = False # bool | If true, only return assemblies with the given NCBI Taxonomy ID, or name. Otherwise, assemblies from taxonomy subtree are included, too. Ignored for assembly_accession request. (optional) if omitted the server will use the default value of False
    returned_content = V1AssemblyMetadataRequestContentType("COMPLETE") # V1AssemblyMetadataRequestContentType | Return either assembly accessions, or entire assembly-metadata records (optional)
    page_size = 20 # int | The maximum number of genome assemblies to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, `page_token` can be used to retrieve the remaining results. (optional) if omitted the server will use the default value of 20
    page_token = "page_token_example" # str | A page token is returned from an `AssemblyMetadataRequest` call with more than `page_size` results. Use this token, along with the previous `AssemblyMetadataRequest` parameters, to retrieve the next page of results. When `page_token` is empty, all results have been retrieved. (optional)

    # example passing only required values which don't have defaults set
    try:
        # Get genome metadata by taxonomic identifier
        api_response = api_instance.assembly_descriptors_by_taxon(taxon)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling GenomeApi->assembly_descriptors_by_taxon: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get genome metadata by taxonomic identifier
        api_response = api_instance.assembly_descriptors_by_taxon(taxon, filters_reference_only=filters_reference_only, filters_assembly_source=filters_assembly_source, filters_has_annotation=filters_has_annotation, filters_assembly_level=filters_assembly_level, filters_first_release_date=filters_first_release_date, filters_last_release_date=filters_last_release_date, filters_search_text=filters_search_text, tax_exact_match=tax_exact_match, returned_content=returned_content, page_size=page_size, page_token=page_token)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling GenomeApi->assembly_descriptors_by_taxon: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taxon** | **str**| NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank |
 **filters_reference_only** | **bool**| If true, only return reference and representative (GCF_ and GCA_) genome assemblies. | [optional] if omitted the server will use the default value of False
 **filters_assembly_source** | **V1AssemblyDatasetDescriptorsFilterAssemblySource**| Return only RefSeq (GCF_) or GenBank (GCA_) genome assemblies | [optional]
 **filters_has_annotation** | **bool**| Return only annotated genome assemblies | [optional] if omitted the server will use the default value of False
 **filters_assembly_level** | [**[V1AssemblyDatasetDescriptorsFilterAssemblyLevel]**](V1AssemblyDatasetDescriptorsFilterAssemblyLevel.md)| Only return genome assemblies that have one of the specified assembly levels. By default, do not filter. | [optional]
 **filters_first_release_date** | **datetime**| Only return genome assemblies that were released on or after the specified date By default, do not filter. | [optional]
 **filters_last_release_date** | **datetime**| Only return genome assemblies that were released on or before to the specified date By default, do not filter. | [optional]
 **filters_search_text** | **[str]**| Only return results whose fields contain the specified search terms in their taxon, infraspecific, assembly name or submitter fields By default, do not filter | [optional]
 **tax_exact_match** | **bool**| If true, only return assemblies with the given NCBI Taxonomy ID, or name. Otherwise, assemblies from taxonomy subtree are included, too. Ignored for assembly_accession request. | [optional] if omitted the server will use the default value of False
 **returned_content** | **V1AssemblyMetadataRequestContentType**| Return either assembly accessions, or entire assembly-metadata records | [optional]
 **page_size** | **int**| The maximum number of genome assemblies to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, &#x60;page_token&#x60; can be used to retrieve the remaining results. | [optional] if omitted the server will use the default value of 20
 **page_token** | **str**| A page token is returned from an &#x60;AssemblyMetadataRequest&#x60; call with more than &#x60;page_size&#x60; results. Use this token, along with the previous &#x60;AssemblyMetadataRequest&#x60; parameters, to retrieve the next page of results. When &#x60;page_token&#x60; is empty, all results have been retrieved. | [optional]

### Return type

[**V1AssemblyMetadata**](V1AssemblyMetadata.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **check_assembly_availability**
> V1AssemblyDatasetAvailability check_assembly_availability(accessions)

Check the validity of genome accessions

The 'GET' version of check is limited by the size of the GET URL (2KB, which works out to about 140 genomic accessions).  The POST operation is provided to allow users to supply a larger number of accessions in a single request.

### Example

* Api Key Authentication (ApiKeyAuthHeader):

```python
import time
import ncbi.datasets.openapi
from ncbi.datasets.openapi.api import genome_api
from ncbi.datasets.openapi.model.v1_assembly_dataset_availability import V1AssemblyDatasetAvailability
from ncbi.datasets.openapi.model.rpc_status import RpcStatus
from pprint import pprint
# Defining the host is optional and defaults to https://api.ncbi.nlm.nih.gov/datasets/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = ncbi.datasets.openapi.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: ApiKeyAuthHeader
configuration.api_key['ApiKeyAuthHeader'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['ApiKeyAuthHeader'] = 'Bearer'

# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = genome_api.GenomeApi(api_client)
    accessions = ["GCF_000001405.40","GCF_000001635.27"] # [str] | NCBI genome assembly accessions

    # example passing only required values which don't have defaults set
    try:
        # Check the validity of genome accessions
        api_response = api_instance.check_assembly_availability(accessions)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling GenomeApi->check_assembly_availability: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessions** | **[str]**| NCBI genome assembly accessions |

### Return type

[**V1AssemblyDatasetAvailability**](V1AssemblyDatasetAvailability.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **check_assembly_availability_post**
> V1AssemblyDatasetAvailability check_assembly_availability_post(v1_assembly_dataset_request)

Check the validity of many genome accessions in a single request

The 'GET' version of check is limited by the size of the GET URL (2KB, which works out to about 140 genomic accessions).  The POST operation is provided to allow users to supply a larger number of accessions in a single request.

### Example

* Api Key Authentication (ApiKeyAuthHeader):

```python
import time
import ncbi.datasets.openapi
from ncbi.datasets.openapi.api import genome_api
from ncbi.datasets.openapi.model.v1_assembly_dataset_availability import V1AssemblyDatasetAvailability
from ncbi.datasets.openapi.model.v1_assembly_dataset_request import V1AssemblyDatasetRequest
from ncbi.datasets.openapi.model.rpc_status import RpcStatus
from pprint import pprint
# Defining the host is optional and defaults to https://api.ncbi.nlm.nih.gov/datasets/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = ncbi.datasets.openapi.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: ApiKeyAuthHeader
configuration.api_key['ApiKeyAuthHeader'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['ApiKeyAuthHeader'] = 'Bearer'

# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = genome_api.GenomeApi(api_client)
    v1_assembly_dataset_request = V1AssemblyDatasetRequest(
        assembly_accessions=[
            "assembly_accessions_example",
        ],
        accessions=[
            "accessions_example",
        ],
        chromosomes=[
            "chromosomes_example",
        ],
        include_annotation=True,
        exclude_sequence=True,
        include_annotation_type=[
            V1AnnotationForAssemblyType("DEFAULT"),
        ],
        hydrated=V1AssemblyDatasetRequestResolution("FULLY_HYDRATED"),
        include_tsv=True,
    ) # V1AssemblyDatasetRequest | 

    # example passing only required values which don't have defaults set
    try:
        # Check the validity of many genome accessions in a single request
        api_response = api_instance.check_assembly_availability_post(v1_assembly_dataset_request)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling GenomeApi->check_assembly_availability_post: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1_assembly_dataset_request** | [**V1AssemblyDatasetRequest**](V1AssemblyDatasetRequest.md)|  |

### Return type

[**V1AssemblyDatasetAvailability**](V1AssemblyDatasetAvailability.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **download_assembly_package**
> file_type download_assembly_package(accessions)

Get a genome dataset by accession

Download a genome dataset including fasta sequence, annotation and a detailed data report by accession.

### Example

* Api Key Authentication (ApiKeyAuthHeader):

```python
import time
import ncbi.datasets.openapi
from ncbi.datasets.openapi.api import genome_api
from ncbi.datasets.openapi.model.v1_annotation_for_assembly_type import V1AnnotationForAssemblyType
from ncbi.datasets.openapi.model.v1_assembly_dataset_request_resolution import V1AssemblyDatasetRequestResolution
from ncbi.datasets.openapi.model.rpc_status import RpcStatus
from pprint import pprint
# Defining the host is optional and defaults to https://api.ncbi.nlm.nih.gov/datasets/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = ncbi.datasets.openapi.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: ApiKeyAuthHeader
configuration.api_key['ApiKeyAuthHeader'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['ApiKeyAuthHeader'] = 'Bearer'

# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = genome_api.GenomeApi(api_client)
    accessions = ["GCF_000001405.40","GCF_000001635.27"] # [str] | NCBI genome assembly accessions
    chromosomes = [1,2,3,"X","Y","MT"] # [str] | The default setting is all chromosome. Specify individual chromosome by string (1,2,MT or chr1,chr2.chrMT). Unplaced sequences are treated like their own chromosome ('Un'). The filter only applies to fasta sequence. (optional)
    exclude_sequence = True # bool | Set to true to omit the genomic sequence. (optional) if omitted the server will use the default value of False
    include_annotation_type = [
        V1AnnotationForAssemblyType("["PROT_FASTA","RNA_FASTA"]"),
    ] # [V1AnnotationForAssemblyType] | Select additional types of annotation to include in the data package.  If unset, no annotation is provided. (optional)
    hydrated = V1AssemblyDatasetRequestResolution("FULLY_HYDRATED") # V1AssemblyDatasetRequestResolution | Set to DATA_REPORT_ONLY, to only retrieve data-reports. (optional)
    filename = "ncbi_dataset.zip" # str | Output file name. (optional) if omitted the server will use the default value of "ncbi_dataset.zip"

    # example passing only required values which don't have defaults set
    try:
        # Get a genome dataset by accession
        api_response = api_instance.download_assembly_package(accessions)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling GenomeApi->download_assembly_package: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get a genome dataset by accession
        api_response = api_instance.download_assembly_package(accessions, chromosomes=chromosomes, exclude_sequence=exclude_sequence, include_annotation_type=include_annotation_type, hydrated=hydrated, filename=filename)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling GenomeApi->download_assembly_package: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessions** | **[str]**| NCBI genome assembly accessions |
 **chromosomes** | **[str]**| The default setting is all chromosome. Specify individual chromosome by string (1,2,MT or chr1,chr2.chrMT). Unplaced sequences are treated like their own chromosome (&#39;Un&#39;). The filter only applies to fasta sequence. | [optional]
 **exclude_sequence** | **bool**| Set to true to omit the genomic sequence. | [optional] if omitted the server will use the default value of False
 **include_annotation_type** | [**[V1AnnotationForAssemblyType]**](V1AnnotationForAssemblyType.md)| Select additional types of annotation to include in the data package.  If unset, no annotation is provided. | [optional]
 **hydrated** | **V1AssemblyDatasetRequestResolution**| Set to DATA_REPORT_ONLY, to only retrieve data-reports. | [optional]
 **filename** | **str**| Output file name. | [optional] if omitted the server will use the default value of "ncbi_dataset.zip"

### Return type

**file_type**

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/zip, application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Download selected genome assemblies and associated annotation data as a zip file |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **download_assembly_package_post**
> file_type download_assembly_package_post(v1_assembly_dataset_request)

Get a genome dataset by post

The 'GET' version of download is limited by the size of the GET URL (2KB, which works out to about 140 genomic accessions).  The POST operation is provided to allow users to supply a larger number of accessions in a single request.

### Example

* Api Key Authentication (ApiKeyAuthHeader):

```python
import time
import ncbi.datasets.openapi
from ncbi.datasets.openapi.api import genome_api
from ncbi.datasets.openapi.model.v1_assembly_dataset_request import V1AssemblyDatasetRequest
from ncbi.datasets.openapi.model.rpc_status import RpcStatus
from pprint import pprint
# Defining the host is optional and defaults to https://api.ncbi.nlm.nih.gov/datasets/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = ncbi.datasets.openapi.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: ApiKeyAuthHeader
configuration.api_key['ApiKeyAuthHeader'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['ApiKeyAuthHeader'] = 'Bearer'

# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = genome_api.GenomeApi(api_client)
    v1_assembly_dataset_request = V1AssemblyDatasetRequest(
        assembly_accessions=[
            "assembly_accessions_example",
        ],
        accessions=[
            "accessions_example",
        ],
        chromosomes=[
            "chromosomes_example",
        ],
        include_annotation=True,
        exclude_sequence=True,
        include_annotation_type=[
            V1AnnotationForAssemblyType("DEFAULT"),
        ],
        hydrated=V1AssemblyDatasetRequestResolution("FULLY_HYDRATED"),
        include_tsv=True,
    ) # V1AssemblyDatasetRequest | 
    filename = "ncbi_dataset.zip" # str | Output file name. (optional) if omitted the server will use the default value of "ncbi_dataset.zip"

    # example passing only required values which don't have defaults set
    try:
        # Get a genome dataset by post
        api_response = api_instance.download_assembly_package_post(v1_assembly_dataset_request)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling GenomeApi->download_assembly_package_post: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get a genome dataset by post
        api_response = api_instance.download_assembly_package_post(v1_assembly_dataset_request, filename=filename)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling GenomeApi->download_assembly_package_post: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1_assembly_dataset_request** | [**V1AssemblyDatasetRequest**](V1AssemblyDatasetRequest.md)|  |
 **filename** | **str**| Output file name. | [optional] if omitted the server will use the default value of "ncbi_dataset.zip"

### Return type

**file_type**

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/zip, application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Download selected genome assemblies and associated annotation data as a zip file |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **genome_download_summary**
> V1DownloadSummary genome_download_summary(accessions)

Preview genome dataset download

Get a download summary by accession in a JSON output format.

### Example

* Api Key Authentication (ApiKeyAuthHeader):

```python
import time
import ncbi.datasets.openapi
from ncbi.datasets.openapi.api import genome_api
from ncbi.datasets.openapi.model.v1_download_summary import V1DownloadSummary
from ncbi.datasets.openapi.model.v1_annotation_for_assembly_type import V1AnnotationForAssemblyType
from ncbi.datasets.openapi.model.rpc_status import RpcStatus
from pprint import pprint
# Defining the host is optional and defaults to https://api.ncbi.nlm.nih.gov/datasets/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = ncbi.datasets.openapi.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: ApiKeyAuthHeader
configuration.api_key['ApiKeyAuthHeader'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['ApiKeyAuthHeader'] = 'Bearer'

# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = genome_api.GenomeApi(api_client)
    accessions = ["GCF_000001405.40","GCF_000001635.27"] # [str] | NCBI genome assembly accessions
    chromosomes = [1,2,3,"X","Y","MT"] # [str] | The default setting is all chromosome. Specify individual chromosome by string (1,2,MT or chr1,chr2.chrMT). Unplaced sequences are treated like their own chromosome ('Un'). The filter only applies to fasta sequence. (optional)
    exclude_sequence = True # bool | Set to true to omit the genomic sequence. (optional) if omitted the server will use the default value of False
    include_annotation_type = [
        V1AnnotationForAssemblyType("["PROT_FASTA","RNA_FASTA"]"),
    ] # [V1AnnotationForAssemblyType] | Select additional types of annotation to include in the data package.  If unset, no annotation is provided. (optional)

    # example passing only required values which don't have defaults set
    try:
        # Preview genome dataset download
        api_response = api_instance.genome_download_summary(accessions)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling GenomeApi->genome_download_summary: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Preview genome dataset download
        api_response = api_instance.genome_download_summary(accessions, chromosomes=chromosomes, exclude_sequence=exclude_sequence, include_annotation_type=include_annotation_type)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling GenomeApi->genome_download_summary: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessions** | **[str]**| NCBI genome assembly accessions |
 **chromosomes** | **[str]**| The default setting is all chromosome. Specify individual chromosome by string (1,2,MT or chr1,chr2.chrMT). Unplaced sequences are treated like their own chromosome (&#39;Un&#39;). The filter only applies to fasta sequence. | [optional]
 **exclude_sequence** | **bool**| Set to true to omit the genomic sequence. | [optional] if omitted the server will use the default value of False
 **include_annotation_type** | [**[V1AnnotationForAssemblyType]**](V1AnnotationForAssemblyType.md)| Select additional types of annotation to include in the data package.  If unset, no annotation is provided. | [optional]

### Return type

[**V1DownloadSummary**](V1DownloadSummary.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **genome_download_summary_by_post**
> V1DownloadSummary genome_download_summary_by_post(v1_assembly_dataset_request)

Preview genome dataset download by POST

The 'GET' version of download summary is limited by the size of the GET URL (2KB, which works out to about 140 genomic accessions).  The POST operation is provided to allow users to supply a larger number of accessions in a single request.

### Example

* Api Key Authentication (ApiKeyAuthHeader):

```python
import time
import ncbi.datasets.openapi
from ncbi.datasets.openapi.api import genome_api
from ncbi.datasets.openapi.model.v1_download_summary import V1DownloadSummary
from ncbi.datasets.openapi.model.v1_assembly_dataset_request import V1AssemblyDatasetRequest
from ncbi.datasets.openapi.model.rpc_status import RpcStatus
from pprint import pprint
# Defining the host is optional and defaults to https://api.ncbi.nlm.nih.gov/datasets/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = ncbi.datasets.openapi.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: ApiKeyAuthHeader
configuration.api_key['ApiKeyAuthHeader'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['ApiKeyAuthHeader'] = 'Bearer'

# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = genome_api.GenomeApi(api_client)
    v1_assembly_dataset_request = V1AssemblyDatasetRequest(
        assembly_accessions=[
            "assembly_accessions_example",
        ],
        accessions=[
            "accessions_example",
        ],
        chromosomes=[
            "chromosomes_example",
        ],
        include_annotation=True,
        exclude_sequence=True,
        include_annotation_type=[
            V1AnnotationForAssemblyType("DEFAULT"),
        ],
        hydrated=V1AssemblyDatasetRequestResolution("FULLY_HYDRATED"),
        include_tsv=True,
    ) # V1AssemblyDatasetRequest | 

    # example passing only required values which don't have defaults set
    try:
        # Preview genome dataset download by POST
        api_response = api_instance.genome_download_summary_by_post(v1_assembly_dataset_request)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling GenomeApi->genome_download_summary_by_post: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1_assembly_dataset_request** | [**V1AssemblyDatasetRequest**](V1AssemblyDatasetRequest.md)|  |

### Return type

[**V1DownloadSummary**](V1DownloadSummary.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **genome_metadata_by_post**
> V1AssemblyMetadata genome_metadata_by_post(v1_assembly_metadata_request)

Get genome metadata by variety of identifiers

Get detailed metadata for assembled genomes.

### Example

* Api Key Authentication (ApiKeyAuthHeader):

```python
import time
import ncbi.datasets.openapi
from ncbi.datasets.openapi.api import genome_api
from ncbi.datasets.openapi.model.v1_assembly_metadata_request import V1AssemblyMetadataRequest
from ncbi.datasets.openapi.model.v1_assembly_metadata import V1AssemblyMetadata
from ncbi.datasets.openapi.model.rpc_status import RpcStatus
from pprint import pprint
# Defining the host is optional and defaults to https://api.ncbi.nlm.nih.gov/datasets/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = ncbi.datasets.openapi.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: ApiKeyAuthHeader
configuration.api_key['ApiKeyAuthHeader'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['ApiKeyAuthHeader'] = 'Bearer'

# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = genome_api.GenomeApi(api_client)
    v1_assembly_metadata_request = V1AssemblyMetadataRequest(
        taxon="taxon_example",
        accessions=V1Accessions(
            accessions=[
                "accessions_example",
            ],
        ),
        bioprojects=V1AssemblyMetadataRequestBioprojects(
            accessions=[
                "accessions_example",
            ],
        ),
        filters=V1AssemblyDatasetDescriptorsFilter(
            reference_only=True,
            refseq_only=True,
            assembly_source=V1AssemblyDatasetDescriptorsFilterAssemblySource("all"),
            has_annotation=True,
            assembly_level=[
                V1AssemblyDatasetDescriptorsFilterAssemblyLevel("chromosome"),
            ],
            first_release_date=dateutil_parser('1970-01-01T00:00:00.00Z'),
            last_release_date=dateutil_parser('1970-01-01T00:00:00.00Z'),
            search_text=[
                "search_text_example",
            ],
        ),
        tax_exact_match=True,
        returned_content=V1AssemblyMetadataRequestContentType("COMPLETE"),
        limit="limit_example",
        page_size=1,
        page_token="page_token_example",
    ) # V1AssemblyMetadataRequest | 

    # example passing only required values which don't have defaults set
    try:
        # Get genome metadata by variety of identifiers
        api_response = api_instance.genome_metadata_by_post(v1_assembly_metadata_request)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling GenomeApi->genome_metadata_by_post: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1_assembly_metadata_request** | [**V1AssemblyMetadataRequest**](V1AssemblyMetadataRequest.md)|  |

### Return type

[**V1AssemblyMetadata**](V1AssemblyMetadata.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **genome_tax_name_query**
> V1SciNameAndIds genome_tax_name_query(taxon_query)

Get a list of taxonomy names and IDs found in the assembly dataset given a partial taxonomic name

This endpoint retrieves a list of taxonomy names and IDs found in the assembly dataset given a partial taxonomic name of any rank.

### Example

* Api Key Authentication (ApiKeyAuthHeader):

```python
import time
import ncbi.datasets.openapi
from ncbi.datasets.openapi.api import genome_api
from ncbi.datasets.openapi.model.v1_sci_name_and_ids import V1SciNameAndIds
from ncbi.datasets.openapi.model.rpc_status import RpcStatus
from ncbi.datasets.openapi.model.v1_organism_query_request_tax_rank_filter import V1OrganismQueryRequestTaxRankFilter
from pprint import pprint
# Defining the host is optional and defaults to https://api.ncbi.nlm.nih.gov/datasets/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = ncbi.datasets.openapi.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: ApiKeyAuthHeader
configuration.api_key['ApiKeyAuthHeader'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['ApiKeyAuthHeader'] = 'Bearer'

# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = genome_api.GenomeApi(api_client)
    taxon_query = "hum" # str | NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank
    tax_rank_filter = V1OrganismQueryRequestTaxRankFilter("species") # V1OrganismQueryRequestTaxRankFilter | Set the scope of searched tax ranks (optional)

    # example passing only required values which don't have defaults set
    try:
        # Get a list of taxonomy names and IDs found in the assembly dataset given a partial taxonomic name
        api_response = api_instance.genome_tax_name_query(taxon_query)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling GenomeApi->genome_tax_name_query: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get a list of taxonomy names and IDs found in the assembly dataset given a partial taxonomic name
        api_response = api_instance.genome_tax_name_query(taxon_query, tax_rank_filter=tax_rank_filter)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling GenomeApi->genome_tax_name_query: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taxon_query** | **str**| NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank |
 **tax_rank_filter** | **V1OrganismQueryRequestTaxRankFilter**| Set the scope of searched tax ranks | [optional]

### Return type

[**V1SciNameAndIds**](V1SciNameAndIds.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **genome_tax_tree**
> V1Organism genome_tax_tree(taxon)

Get a taxonomic subtree by taxonomic identifier

Using a NCBI Taxonomy ID or name (common or scientific) at any rank, get a subtree filtered for species with assembled genomes.

### Example

* Api Key Authentication (ApiKeyAuthHeader):

```python
import time
import ncbi.datasets.openapi
from ncbi.datasets.openapi.api import genome_api
from ncbi.datasets.openapi.model.v1_organism import V1Organism
from ncbi.datasets.openapi.model.rpc_status import RpcStatus
from pprint import pprint
# Defining the host is optional and defaults to https://api.ncbi.nlm.nih.gov/datasets/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = ncbi.datasets.openapi.Configuration(
    host = "https://api.ncbi.nlm.nih.gov/datasets/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: ApiKeyAuthHeader
configuration.api_key['ApiKeyAuthHeader'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['ApiKeyAuthHeader'] = 'Bearer'

# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = genome_api.GenomeApi(api_client)
    taxon = "9606" # str | NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank
    children_only = False # bool | Only report the children of the requested taxon and not their descendants (optional) if omitted the server will use the default value of False

    # example passing only required values which don't have defaults set
    try:
        # Get a taxonomic subtree by taxonomic identifier
        api_response = api_instance.genome_tax_tree(taxon)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling GenomeApi->genome_tax_tree: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get a taxonomic subtree by taxonomic identifier
        api_response = api_instance.genome_tax_tree(taxon, children_only=children_only)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling GenomeApi->genome_tax_tree: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taxon** | **str**| NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank |
 **children_only** | **bool**| Only report the children of the requested taxon and not their descendants | [optional] if omitted the server will use the default value of False

### Return type

[**V1Organism**](V1Organism.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

