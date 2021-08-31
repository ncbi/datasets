# ncbi.datasets.openapi.ProkaryoteApi

All URIs are relative to *https://api.ncbi.nlm.nih.gov/datasets/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**download_prokaryote_gene_package**](ProkaryoteApi.md#download_prokaryote_gene_package) | **GET** /protein/accession/{accessions}/download | Get a prokaryote gene dataset by RefSeq protein accession
[**download_prokaryote_gene_package_post**](ProkaryoteApi.md#download_prokaryote_gene_package_post) | **POST** /protein/accession/download | Get a prokaryote gene dataset by RefSeq protein accession by POST


# **download_prokaryote_gene_package**
> file_type download_prokaryote_gene_package(accessions)

Get a prokaryote gene dataset by RefSeq protein accession

Get a prokaryote gene dataset including gene and protein fasta sequence, annotation and metadata by prokaryote protein accession.

### Example

* Api Key Authentication (ApiKeyAuthHeader):
```python
import time
import ncbi.datasets.openapi
from ncbi.datasets.openapi.api import prokaryote_api
from ncbi.datasets.openapi.model.v1_fasta import V1Fasta
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
    api_instance = prokaryote_api.ProkaryoteApi(api_client)
    accessions = [
        "WP_015878339.1",
    ] # [str] | WP prokaryote protein accession
    include_annotation_type = [
        V1Fasta("FASTA_GENE"),
    ] # [V1Fasta] | Select additional types of annotation to include in the data package.  If unset, no annotation is provided. (optional)
    gene_flank_config_length = 1 # int |  (optional)
    taxon = "taxon_example" # str | NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank When specified, return data from this taxon and its subtree (optional)
    filename = "ncbi_dataset.zip" # str | Output file name. (optional) if omitted the server will use the default value of "ncbi_dataset.zip"

    # example passing only required values which don't have defaults set
    try:
        # Get a prokaryote gene dataset by RefSeq protein accession
        api_response = api_instance.download_prokaryote_gene_package(accessions)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling ProkaryoteApi->download_prokaryote_gene_package: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get a prokaryote gene dataset by RefSeq protein accession
        api_response = api_instance.download_prokaryote_gene_package(accessions, include_annotation_type=include_annotation_type, gene_flank_config_length=gene_flank_config_length, taxon=taxon, filename=filename)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling ProkaryoteApi->download_prokaryote_gene_package: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessions** | **[str]**| WP prokaryote protein accession |
 **include_annotation_type** | [**[V1Fasta]**](V1Fasta.md)| Select additional types of annotation to include in the data package.  If unset, no annotation is provided. | [optional]
 **gene_flank_config_length** | **int**|  | [optional]
 **taxon** | **str**| NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank When specified, return data from this taxon and its subtree | [optional]
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

# **download_prokaryote_gene_package_post**
> file_type download_prokaryote_gene_package_post(v1_prokaryote_gene_request)

Get a prokaryote gene dataset by RefSeq protein accession by POST

Get a prokaryote gene dataset including gene and protein fasta sequence, annotation and metadata by prokaryote protein accession by POST.

### Example

* Api Key Authentication (ApiKeyAuthHeader):
```python
import time
import ncbi.datasets.openapi
from ncbi.datasets.openapi.api import prokaryote_api
from ncbi.datasets.openapi.model.v1_prokaryote_gene_request import V1ProkaryoteGeneRequest
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
    api_instance = prokaryote_api.ProkaryoteApi(api_client)
    v1_prokaryote_gene_request = V1ProkaryoteGeneRequest(
        accessions=[
            "accessions_example",
        ],
        include_annotation_type=[
            V1Fasta("FASTA_UNSPECIFIED"),
        ],
        gene_flank_config=V1ProkaryoteGeneRequestGeneFlankConfig(
            length=1,
        ),
        taxon="taxon_example",
    ) # V1ProkaryoteGeneRequest | 
    filename = "ncbi_dataset.zip" # str | Output file name. (optional) if omitted the server will use the default value of "ncbi_dataset.zip"

    # example passing only required values which don't have defaults set
    try:
        # Get a prokaryote gene dataset by RefSeq protein accession by POST
        api_response = api_instance.download_prokaryote_gene_package_post(v1_prokaryote_gene_request)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling ProkaryoteApi->download_prokaryote_gene_package_post: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        # Get a prokaryote gene dataset by RefSeq protein accession by POST
        api_response = api_instance.download_prokaryote_gene_package_post(v1_prokaryote_gene_request, filename=filename)
        pprint(api_response)
    except ncbi.datasets.openapi.ApiException as e:
        print("Exception when calling ProkaryoteApi->download_prokaryote_gene_package_post: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1_prokaryote_gene_request** | [**V1ProkaryoteGeneRequest**](V1ProkaryoteGeneRequest.md)|  |
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

