# ncbi.datasets.openapi.ProkaryoteApi

All URIs are relative to *https://api.ncbi.nlm.nih.gov/datasets/v1alpha*

Method | HTTP request | Description
------------- | ------------- | -------------
[**download_prokaryote_gene_package**](ProkaryoteApi.md#download_prokaryote_gene_package) | **GET** /protein/accession/{accessions}/download | Get a prokaryote gene dataset by RefSeq protein accession
[**download_prokaryote_gene_package_post**](ProkaryoteApi.md#download_prokaryote_gene_package_post) | **POST** /protein/accession/download | Get a prokaryote gene dataset by RefSeq protein accession by POST


# **download_prokaryote_gene_package**
> file download_prokaryote_gene_package(accessions, include_annotation_type=include_annotation_type, gene_flank_config_length=gene_flank_config_length, taxon=taxon, filename=filename)

Get a prokaryote gene dataset by RefSeq protein accession

Get a prokaryote gene dataset including gene and protein fasta sequence, annotation and metadata by prokaryote protein accession.

### Example

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


# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = ncbi.datasets.openapi.ProkaryoteApi(api_client)
    accessions = ['accessions_example'] # list[str] | WP prokaryote protein accession
include_annotation_type = ['include_annotation_type_example'] # list[str] | Select additional types of annotation to include in the data package.  If unset, no annotation is provided. (optional)
gene_flank_config_length = 56 # int |  (optional)
taxon = 'taxon_example' # str | NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank When specified, return data from this taxon and its subtree. (optional)
filename = 'ncbi_dataset.zip' # str | Output file name. (optional) (default to 'ncbi_dataset.zip')

    try:
        # Get a prokaryote gene dataset by RefSeq protein accession
        api_response = api_instance.download_prokaryote_gene_package(accessions, include_annotation_type=include_annotation_type, gene_flank_config_length=gene_flank_config_length, taxon=taxon, filename=filename)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling ProkaryoteApi->download_prokaryote_gene_package: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessions** | [**list[str]**](str.md)| WP prokaryote protein accession | 
 **include_annotation_type** | [**list[str]**](str.md)| Select additional types of annotation to include in the data package.  If unset, no annotation is provided. | [optional] 
 **gene_flank_config_length** | **int**|  | [optional] 
 **taxon** | **str**| NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank When specified, return data from this taxon and its subtree. | [optional] 
 **filename** | **str**| Output file name. | [optional] [default to &#39;ncbi_dataset.zip&#39;]

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
**200** | Download selected prokaryote gene data as a zip file. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **download_prokaryote_gene_package_post**
> file download_prokaryote_gene_package_post(body, filename=filename)

Get a prokaryote gene dataset by RefSeq protein accession by POST

Get a prokaryote gene dataset including gene and protein fasta sequence, annotation and metadata by prokaryote protein accession by POST.

### Example

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


# Enter a context with an instance of the API client
with ncbi.datasets.openapi.ApiClient() as api_client:
    # Create an instance of the API class
    api_instance = ncbi.datasets.openapi.ProkaryoteApi(api_client)
    body = ncbi.datasets.openapi.V1alpha1ProkaryoteGeneRequest() # V1alpha1ProkaryoteGeneRequest | 
filename = 'ncbi_dataset.zip' # str | Output file name. (optional) (default to 'ncbi_dataset.zip')

    try:
        # Get a prokaryote gene dataset by RefSeq protein accession by POST
        api_response = api_instance.download_prokaryote_gene_package_post(body, filename=filename)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling ProkaryoteApi->download_prokaryote_gene_package_post: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**V1alpha1ProkaryoteGeneRequest**](V1alpha1ProkaryoteGeneRequest.md)|  | 
 **filename** | **str**| Output file name. | [optional] [default to &#39;ncbi_dataset.zip&#39;]

### Return type

**file**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/zip

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Download selected prokaryote gene data as a zip file. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

