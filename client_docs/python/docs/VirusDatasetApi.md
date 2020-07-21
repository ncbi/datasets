# ncbi.datasets.VirusDatasetApi

All URIs are relative to *https://api.ncbi.nlm.nih.gov/datasets/v1alpha*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_sars2_summary**](VirusDatasetApi.md#get_sars2_summary) | **GET** /virus/summary/sars2/protein/{proteins} | Summary of SARS-CoV-2 protein and CDS datasets by protein name
[**get_summary**](VirusDatasetApi.md#get_summary) | **GET** /virus/summary/taxid/{tax_id} | Summary of Coronavirus genome datasets, including options to download package by taxonomy ID
[**get_summary_by_name**](VirusDatasetApi.md#get_summary_by_name) | **GET** /virus/summary/organism/{tax_name} | Summary of Coronavirus genome datasets, including options to download package by taxonomy name


# **get_sars2_summary**
> V1alpha1DownloadSummary get_sars2_summary(proteins, refseq_only=refseq_only, annotated_only=annotated_only, released_since=released_since, host=host)

Summary of SARS-CoV-2 protein and CDS datasets by protein name

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
    api_instance = ncbi.datasets.VirusDatasetApi(api_client)
    proteins = ['proteins_example'] # list[str] | Which proteins to retrieve in the data package
refseq_only = True # bool | The source database is RefSeq. (optional)
annotated_only = True # bool | Only those assemblies with annotation as asserted by Virus. (optional)
released_since = '2013-10-20T19:20:30+01:00' # datetime | Only include viral genomes that have been released after a specified date and time. April 1, 2020 midnight UTC should be formatted as follows: 2020-04-01T00:00:00.000Z. (optional)
host = 'host_example' # str | Assemblies that were taken from this host (tax-id or name) or its taxonomic descendents. (optional)

    try:
        # Summary of SARS-CoV-2 protein and CDS datasets by protein name
        api_response = api_instance.get_sars2_summary(proteins, refseq_only=refseq_only, annotated_only=annotated_only, released_since=released_since, host=host)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling VirusDatasetApi->get_sars2_summary: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **proteins** | [**list[str]**](str.md)| Which proteins to retrieve in the data package | 
 **refseq_only** | **bool**| The source database is RefSeq. | [optional] 
 **annotated_only** | **bool**| Only those assemblies with annotation as asserted by Virus. | [optional] 
 **released_since** | **datetime**| Only include viral genomes that have been released after a specified date and time. April 1, 2020 midnight UTC should be formatted as follows: 2020-04-01T00:00:00.000Z. | [optional] 
 **host** | **str**| Assemblies that were taken from this host (tax-id or name) or its taxonomic descendents. | [optional] 

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

# **get_summary**
> V1alpha1DownloadSummary get_summary(tax_id, refseq_only=refseq_only, annotated_only=annotated_only, released_since=released_since, host=host)

Summary of Coronavirus genome datasets, including options to download package by taxonomy ID

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
    api_instance = ncbi.datasets.VirusDatasetApi(api_client)
    tax_id = 56 # int | NCBI Taxonomy ID
refseq_only = True # bool | The source database is RefSeq. (optional)
annotated_only = True # bool | Only those assemblies with annotation as asserted by Virus. (optional)
released_since = '2013-10-20T19:20:30+01:00' # datetime | Only include viral genomes that have been released after a specified date and time. April 1, 2020 midnight UTC should be formatted as follows: 2020-04-01T00:00:00.000Z. (optional)
host = 'host_example' # str | Assemblies that were taken from this host (tax-id or name) or its taxonomic descendents. (optional)

    try:
        # Summary of Coronavirus genome datasets, including options to download package by taxonomy ID
        api_response = api_instance.get_summary(tax_id, refseq_only=refseq_only, annotated_only=annotated_only, released_since=released_since, host=host)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling VirusDatasetApi->get_summary: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tax_id** | **int**| NCBI Taxonomy ID | 
 **refseq_only** | **bool**| The source database is RefSeq. | [optional] 
 **annotated_only** | **bool**| Only those assemblies with annotation as asserted by Virus. | [optional] 
 **released_since** | **datetime**| Only include viral genomes that have been released after a specified date and time. April 1, 2020 midnight UTC should be formatted as follows: 2020-04-01T00:00:00.000Z. | [optional] 
 **host** | **str**| Assemblies that were taken from this host (tax-id or name) or its taxonomic descendents. | [optional] 

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

# **get_summary_by_name**
> V1alpha1DownloadSummary get_summary_by_name(tax_name, refseq_only=refseq_only, annotated_only=annotated_only, released_since=released_since, host=host)

Summary of Coronavirus genome datasets, including options to download package by taxonomy name

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
    api_instance = ncbi.datasets.VirusDatasetApi(api_client)
    tax_name = 'tax_name_example' # str | Common or scientific name at any tax rank for all domains of life
refseq_only = True # bool | The source database is RefSeq. (optional)
annotated_only = True # bool | Only those assemblies with annotation as asserted by Virus. (optional)
released_since = '2013-10-20T19:20:30+01:00' # datetime | Only include viral genomes that have been released after a specified date and time. April 1, 2020 midnight UTC should be formatted as follows: 2020-04-01T00:00:00.000Z. (optional)
host = 'host_example' # str | Assemblies that were taken from this host (tax-id or name) or its taxonomic descendents. (optional)

    try:
        # Summary of Coronavirus genome datasets, including options to download package by taxonomy name
        api_response = api_instance.get_summary_by_name(tax_name, refseq_only=refseq_only, annotated_only=annotated_only, released_since=released_since, host=host)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling VirusDatasetApi->get_summary_by_name: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tax_name** | **str**| Common or scientific name at any tax rank for all domains of life | 
 **refseq_only** | **bool**| The source database is RefSeq. | [optional] 
 **annotated_only** | **bool**| Only those assemblies with annotation as asserted by Virus. | [optional] 
 **released_since** | **datetime**| Only include viral genomes that have been released after a specified date and time. April 1, 2020 midnight UTC should be formatted as follows: 2020-04-01T00:00:00.000Z. | [optional] 
 **host** | **str**| Assemblies that were taken from this host (tax-id or name) or its taxonomic descendents. | [optional] 

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

