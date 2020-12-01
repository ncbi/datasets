# ncbi.datasets.GenomeApi

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
[**genome_tax_name_query**](GenomeApi.md#genome_tax_name_query) | **GET** /genome/taxon_suggest/{taxon_query} | Get a list of taxonomy names and IDs found in the assembly dataset given a partial taxonomic name.
[**genome_tax_tree**](GenomeApi.md#genome_tax_tree) | **GET** /genome/taxon/{taxon}/tree | Get a taxonomic subtree by taxonomic identifier


# **assembly_descriptors_by_accessions**
> V1alpha1AssemblyMetadata assembly_descriptors_by_accessions(accessions, limit=limit, filters_reference_only=filters_reference_only, filters_refseq_only=filters_refseq_only, returned_content=returned_content)

Get genome metadata by accession

Get detailed metadata for assembled genomes by accession in a JSON output format.

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
    api_instance = ncbi.datasets.GenomeApi(api_client)
    accessions = ['accessions_example'] # list[str] | 
limit = 'limit_example' # str | Limit the number of returned results (\"all\", \"none\", otherwise an integer value). (optional)
filters_reference_only = True # bool | If true, only return reference and representative (GCF_ and GCA_) assemblies. (optional)
filters_refseq_only = True # bool | If true, only return RefSeq (GCF_) assemblies. (optional)
returned_content = 'COMPLETE' # str | Return either assembly accessions, or entire assembly-metadata records. (optional) (default to 'COMPLETE')

    try:
        # Get genome metadata by accession
        api_response = api_instance.assembly_descriptors_by_accessions(accessions, limit=limit, filters_reference_only=filters_reference_only, filters_refseq_only=filters_refseq_only, returned_content=returned_content)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling GenomeApi->assembly_descriptors_by_accessions: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessions** | [**list[str]**](str.md)|  | 
 **limit** | **str**| Limit the number of returned results (\&quot;all\&quot;, \&quot;none\&quot;, otherwise an integer value). | [optional] 
 **filters_reference_only** | **bool**| If true, only return reference and representative (GCF_ and GCA_) assemblies. | [optional] 
 **filters_refseq_only** | **bool**| If true, only return RefSeq (GCF_) assemblies. | [optional] 
 **returned_content** | **str**| Return either assembly accessions, or entire assembly-metadata records. | [optional] [default to &#39;COMPLETE&#39;]

### Return type

[**V1alpha1AssemblyMetadata**](V1alpha1AssemblyMetadata.md)

### Authorization

No authorization required

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
> V1alpha1AssemblyMetadata assembly_descriptors_by_bioproject(accessions, limit=limit, filters_reference_only=filters_reference_only, filters_refseq_only=filters_refseq_only, returned_content=returned_content)

Get genome metadata by bioproject accession

Get detailed metadata for assembled genomes by bioproject accession in a JSON output format.

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
    api_instance = ncbi.datasets.GenomeApi(api_client)
    accessions = ['accessions_example'] # list[str] | 
limit = 'limit_example' # str | Limit the number of returned results (\"all\", \"none\", otherwise an integer value). (optional)
filters_reference_only = True # bool | If true, only return reference and representative (GCF_ and GCA_) assemblies. (optional)
filters_refseq_only = True # bool | If true, only return RefSeq (GCF_) assemblies. (optional)
returned_content = 'COMPLETE' # str | Return either assembly accessions, or entire assembly-metadata records. (optional) (default to 'COMPLETE')

    try:
        # Get genome metadata by bioproject accession
        api_response = api_instance.assembly_descriptors_by_bioproject(accessions, limit=limit, filters_reference_only=filters_reference_only, filters_refseq_only=filters_refseq_only, returned_content=returned_content)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling GenomeApi->assembly_descriptors_by_bioproject: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessions** | [**list[str]**](str.md)|  | 
 **limit** | **str**| Limit the number of returned results (\&quot;all\&quot;, \&quot;none\&quot;, otherwise an integer value). | [optional] 
 **filters_reference_only** | **bool**| If true, only return reference and representative (GCF_ and GCA_) assemblies. | [optional] 
 **filters_refseq_only** | **bool**| If true, only return RefSeq (GCF_) assemblies. | [optional] 
 **returned_content** | **str**| Return either assembly accessions, or entire assembly-metadata records. | [optional] [default to &#39;COMPLETE&#39;]

### Return type

[**V1alpha1AssemblyMetadata**](V1alpha1AssemblyMetadata.md)

### Authorization

No authorization required

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
> V1alpha1AssemblyMetadata assembly_descriptors_by_taxon(taxon, limit=limit, filters_reference_only=filters_reference_only, filters_refseq_only=filters_refseq_only, tax_exact_match=tax_exact_match, returned_content=returned_content)

Get genome metadata by taxonomic identifier

Get detailed metadata on all assembled genomes for a specified NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank

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
    api_instance = ncbi.datasets.GenomeApi(api_client)
    taxon = 'taxon_example' # str | NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank
limit = 'limit_example' # str | Limit the number of returned results (\"all\", \"none\", otherwise an integer value). (optional)
filters_reference_only = True # bool | If true, only return reference and representative (GCF_ and GCA_) assemblies. (optional)
filters_refseq_only = True # bool | If true, only return RefSeq (GCF_) assemblies. (optional)
tax_exact_match = True # bool | If true, only return assemblies with the given NCBI Taxonomy ID, or name. Otherwise, assemblies from taxonomy subtree are included, too. Ignored for assembly_accession request. (optional)
returned_content = 'COMPLETE' # str | Return either assembly accessions, or entire assembly-metadata records. (optional) (default to 'COMPLETE')

    try:
        # Get genome metadata by taxonomic identifier
        api_response = api_instance.assembly_descriptors_by_taxon(taxon, limit=limit, filters_reference_only=filters_reference_only, filters_refseq_only=filters_refseq_only, tax_exact_match=tax_exact_match, returned_content=returned_content)
        pprint(api_response)
    except ApiException as e:
        print("Exception when calling GenomeApi->assembly_descriptors_by_taxon: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taxon** | **str**| NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank | 
 **limit** | **str**| Limit the number of returned results (\&quot;all\&quot;, \&quot;none\&quot;, otherwise an integer value). | [optional] 
 **filters_reference_only** | **bool**| If true, only return reference and representative (GCF_ and GCA_) assemblies. | [optional] 
 **filters_refseq_only** | **bool**| If true, only return RefSeq (GCF_) assemblies. | [optional] 
 **tax_exact_match** | **bool**| If true, only return assemblies with the given NCBI Taxonomy ID, or name. Otherwise, assemblies from taxonomy subtree are included, too. Ignored for assembly_accession request. | [optional] 
 **returned_content** | **str**| Return either assembly accessions, or entire assembly-metadata records. | [optional] [default to &#39;COMPLETE&#39;]

### Return type

[**V1alpha1AssemblyMetadata**](V1alpha1AssemblyMetadata.md)

### Authorization

No authorization required

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
    api_instance = ncbi.datasets.GenomeApi(api_client)
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

No authorization required

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
    api_instance = ncbi.datasets.GenomeApi(api_client)
    body = ncbi.datasets.V1alpha1AssemblyDatasetRequest() # V1alpha1AssemblyDatasetRequest | 

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

No authorization required

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
    api_instance = ncbi.datasets.GenomeApi(api_client)
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

No authorization required

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
    api_instance = ncbi.datasets.GenomeApi(api_client)
    body = ncbi.datasets.V1alpha1AssemblyDatasetRequest() # V1alpha1AssemblyDatasetRequest | 
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

No authorization required

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
    api_instance = ncbi.datasets.GenomeApi(api_client)
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

No authorization required

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
    api_instance = ncbi.datasets.GenomeApi(api_client)
    body = ncbi.datasets.V1alpha1AssemblyDatasetRequest() # V1alpha1AssemblyDatasetRequest | 

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

No authorization required

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
    api_instance = ncbi.datasets.GenomeApi(api_client)
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

No authorization required

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
    api_instance = ncbi.datasets.GenomeApi(api_client)
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

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

