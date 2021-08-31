# ncbi.datasets.openapi.TaxonomyApi

All URIs are relative to *https://api.ncbi.nlm.nih.gov/datasets/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**tax_name_query**](TaxonomyApi.md#tax_name_query) | **GET** /taxonomy/taxon_suggest/{taxon_query} | Get a list of taxonomy names and IDs given a partial taxonomic name
[**taxonomy_filtered_subtree**](TaxonomyApi.md#taxonomy_filtered_subtree) | **GET** /taxonomy/taxon/{taxons}/filtered_subtree | Use taxonomic identifiers to get a filtered taxonomic subtree
[**taxonomy_filtered_subtree_post**](TaxonomyApi.md#taxonomy_filtered_subtree_post) | **POST** /taxonomy/filtered_subtree | Use taxonomic identifiers to get a filtered taxonomic subtree by post
[**taxonomy_metadata**](TaxonomyApi.md#taxonomy_metadata) | **GET** /taxonomy/taxon/{taxons} | Use taxonomic identifiers to get taxonomic metadata
[**taxonomy_metadata_post**](TaxonomyApi.md#taxonomy_metadata_post) | **POST** /taxonomy | Use taxonomic identifiers to get taxonomic metadata by post


# **tax_name_query**
> V1SciNameAndIds tax_name_query(taxon_query)

Get a list of taxonomy names and IDs given a partial taxonomic name

This endpoint retrieves a list of taxonomy names and IDs given a partial taxonomic name of any rank.

### Example

* Api Key Authentication (ApiKeyAuthHeader):
```python
import time
import ncbi.datasets.openapi
from ncbi.datasets.openapi.api import taxonomy_api
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
    api_instance = taxonomy_api.TaxonomyApi(api_client)
    taxon_query = "hum" # str | NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank
    tax_rank_filter = V1OrganismQueryRequestTaxRankFilter("species") # V1OrganismQueryRequestTaxRankFilter | Set the scope of searched tax ranks (optional)

    # example passing only required values which don't have defaults set
    try:
        # Get a list of taxonomy names and IDs given a partial taxonomic name
        api_response = api_instance.tax_name_query(taxon_query)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling TaxonomyApi->tax_name_query: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get a list of taxonomy names and IDs given a partial taxonomic name
        api_response = api_instance.tax_name_query(taxon_query, tax_rank_filter=tax_rank_filter)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling TaxonomyApi->tax_name_query: %s\n" % e)
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

# **taxonomy_filtered_subtree**
> V1TaxonomyFilteredSubtreeResponse taxonomy_filtered_subtree(taxons)

Use taxonomic identifiers to get a filtered taxonomic subtree

Using NCBI Taxonomy IDs or names (common or scientific) at any rank, get a filtered taxonomic subtree that includes the full parent lineage and all immediate children from the selected taxonomic ranks in JSON format.

### Example

* Api Key Authentication (ApiKeyAuthHeader):
```python
import time
import ncbi.datasets.openapi
from ncbi.datasets.openapi.api import taxonomy_api
from ncbi.datasets.openapi.model.v1_organism_rank_type import V1OrganismRankType
from ncbi.datasets.openapi.model.v1_taxonomy_filtered_subtree_response import V1TaxonomyFilteredSubtreeResponse
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
    api_instance = taxonomy_api.TaxonomyApi(api_client)
    taxons = [9606,10090] # [str] | 
    specified_limit = True # bool | Limit to specified species (optional)
    rank_limits = [
        V1OrganismRankType("NO_RANK"),
    ] # [V1OrganismRankType] | Limit to the provided ranks.  If empty, accept any rank. (optional)

    # example passing only required values which don't have defaults set
    try:
        # Use taxonomic identifiers to get a filtered taxonomic subtree
        api_response = api_instance.taxonomy_filtered_subtree(taxons)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling TaxonomyApi->taxonomy_filtered_subtree: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Use taxonomic identifiers to get a filtered taxonomic subtree
        api_response = api_instance.taxonomy_filtered_subtree(taxons, specified_limit=specified_limit, rank_limits=rank_limits)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling TaxonomyApi->taxonomy_filtered_subtree: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taxons** | **[str]**|  |
 **specified_limit** | **bool**| Limit to specified species | [optional]
 **rank_limits** | [**[V1OrganismRankType]**](V1OrganismRankType.md)| Limit to the provided ranks.  If empty, accept any rank. | [optional]

### Return type

[**V1TaxonomyFilteredSubtreeResponse**](V1TaxonomyFilteredSubtreeResponse.md)

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

# **taxonomy_filtered_subtree_post**
> V1TaxonomyFilteredSubtreeResponse taxonomy_filtered_subtree_post(v1_taxonomy_filtered_subtree_request)

Use taxonomic identifiers to get a filtered taxonomic subtree by post

Using NCBI Taxonomy IDs or names (common or scientific) at any rank, get a filtered taxonomic subtree that includes the full parent lineage and all immediate children from the selected taxonomic ranks in JSON format.

### Example

* Api Key Authentication (ApiKeyAuthHeader):
```python
import time
import ncbi.datasets.openapi
from ncbi.datasets.openapi.api import taxonomy_api
from ncbi.datasets.openapi.model.v1_taxonomy_filtered_subtree_request import V1TaxonomyFilteredSubtreeRequest
from ncbi.datasets.openapi.model.v1_taxonomy_filtered_subtree_response import V1TaxonomyFilteredSubtreeResponse
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
    api_instance = taxonomy_api.TaxonomyApi(api_client)
    v1_taxonomy_filtered_subtree_request = V1TaxonomyFilteredSubtreeRequest(
        taxons=[
            "taxons_example",
        ],
        specified_limit=True,
        rank_limits=[
            V1OrganismRankType("NO_RANK"),
        ],
    ) # V1TaxonomyFilteredSubtreeRequest | 

    # example passing only required values which don't have defaults set
    try:
        # Use taxonomic identifiers to get a filtered taxonomic subtree by post
        api_response = api_instance.taxonomy_filtered_subtree_post(v1_taxonomy_filtered_subtree_request)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling TaxonomyApi->taxonomy_filtered_subtree_post: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1_taxonomy_filtered_subtree_request** | [**V1TaxonomyFilteredSubtreeRequest**](V1TaxonomyFilteredSubtreeRequest.md)|  |

### Return type

[**V1TaxonomyFilteredSubtreeResponse**](V1TaxonomyFilteredSubtreeResponse.md)

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

# **taxonomy_metadata**
> V1TaxonomyMetadataResponse taxonomy_metadata(taxons)

Use taxonomic identifiers to get taxonomic metadata

Using NCBI Taxonomy IDs or names (common or scientific) at any rank, get metadata about a taxonomic node including taxonomic identifiers, lineage information, child nodes, and gene and genome counts in JSON format.

### Example

* Api Key Authentication (ApiKeyAuthHeader):
```python
import time
import ncbi.datasets.openapi
from ncbi.datasets.openapi.api import taxonomy_api
from ncbi.datasets.openapi.model.v1_taxonomy_metadata_request_content_type import V1TaxonomyMetadataRequestContentType
from ncbi.datasets.openapi.model.rpc_status import RpcStatus
from ncbi.datasets.openapi.model.v1_taxonomy_metadata_response import V1TaxonomyMetadataResponse
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
    api_instance = taxonomy_api.TaxonomyApi(api_client)
    taxons = [9606,10090] # [str] | 
    returned_content = V1TaxonomyMetadataRequestContentType("COMPLETE") # V1TaxonomyMetadataRequestContentType | Return either tax-ids alone, or entire taxononmy-metadata records (optional)

    # example passing only required values which don't have defaults set
    try:
        # Use taxonomic identifiers to get taxonomic metadata
        api_response = api_instance.taxonomy_metadata(taxons)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling TaxonomyApi->taxonomy_metadata: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Use taxonomic identifiers to get taxonomic metadata
        api_response = api_instance.taxonomy_metadata(taxons, returned_content=returned_content)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling TaxonomyApi->taxonomy_metadata: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taxons** | **[str]**|  |
 **returned_content** | **V1TaxonomyMetadataRequestContentType**| Return either tax-ids alone, or entire taxononmy-metadata records | [optional]

### Return type

[**V1TaxonomyMetadataResponse**](V1TaxonomyMetadataResponse.md)

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

# **taxonomy_metadata_post**
> V1TaxonomyMetadataResponse taxonomy_metadata_post(v1_taxonomy_metadata_request)

Use taxonomic identifiers to get taxonomic metadata by post

Using NCBI Taxonomy IDs or names (common or scientific) at any rank, get metadata about a taxonomic node including taxonomic identifiers, lineage information, child nodes, and gene and genome counts in JSON format.

### Example

* Api Key Authentication (ApiKeyAuthHeader):
```python
import time
import ncbi.datasets.openapi
from ncbi.datasets.openapi.api import taxonomy_api
from ncbi.datasets.openapi.model.v1_taxonomy_metadata_request import V1TaxonomyMetadataRequest
from ncbi.datasets.openapi.model.rpc_status import RpcStatus
from ncbi.datasets.openapi.model.v1_taxonomy_metadata_response import V1TaxonomyMetadataResponse
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
    api_instance = taxonomy_api.TaxonomyApi(api_client)
    v1_taxonomy_metadata_request = V1TaxonomyMetadataRequest(
        taxons=[
            "taxons_example",
        ],
        returned_content=V1TaxonomyMetadataRequestContentType("COMPLETE"),
    ) # V1TaxonomyMetadataRequest | 

    # example passing only required values which don't have defaults set
    try:
        # Use taxonomic identifiers to get taxonomic metadata by post
        api_response = api_instance.taxonomy_metadata_post(v1_taxonomy_metadata_request)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling TaxonomyApi->taxonomy_metadata_post: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1_taxonomy_metadata_request** | [**V1TaxonomyMetadataRequest**](V1TaxonomyMetadataRequest.md)|  |

### Return type

[**V1TaxonomyMetadataResponse**](V1TaxonomyMetadataResponse.md)

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

