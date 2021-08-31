# \VirusApi

All URIs are relative to *https://api.ncbi.nlm.nih.gov/datasets/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Sars2ProteinDownload**](VirusApi.md#Sars2ProteinDownload) | **Get** /virus/taxon/sars2/protein/{proteins}/download | Download SARS-CoV-2 protein and CDS datasets by protein name
[**Sars2ProteinSummary**](VirusApi.md#Sars2ProteinSummary) | **Get** /virus/taxon/sars2/protein/{proteins} | Summary of SARS-CoV-2 protein and CDS datasets by protein name
[**Sars2ProteinTable**](VirusApi.md#Sars2ProteinTable) | **Get** /virus/taxon/sars2/protein/{proteins}/table | Get SARS-CoV-2 protein metadata in a tabular format.
[**VirusGenomeDownload**](VirusApi.md#VirusGenomeDownload) | **Get** /virus/taxon/{taxon}/genome/download | Download Coronavirus genome datasets by taxon
[**VirusGenomeSummary**](VirusApi.md#VirusGenomeSummary) | **Get** /virus/taxon/{taxon}/genome | Get summary data for Coronaviridae genomes by taxon
[**VirusGenomeTable**](VirusApi.md#VirusGenomeTable) | **Get** /virus/taxon/{taxon}/genome/table | Get viral genomic metadata in a tabular format.



## Sars2ProteinDownload

> *os.File Sars2ProteinDownload(ctx, proteins).RefseqOnly(refseqOnly).AnnotatedOnly(annotatedOnly).ReleasedSince(releasedSince).Host(host).GeoLocation(geoLocation).CompleteOnly(completeOnly).IncludeAnnotationType(includeAnnotationType).Filename(filename).Execute()

Download SARS-CoV-2 protein and CDS datasets by protein name



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    proteins := []string{"Inner_example"} // []string | Which proteins to retrieve in the data package
    refseqOnly := true // bool | If true, limit results to RefSeq genomes. (optional) (default to false)
    annotatedOnly := true // bool | If true, limit results to annotated genomes. (optional) (default to false)
    releasedSince := time.Now() // time.Time | If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as '2020-04-01T00:00:00.000Z' (optional)
    host := "human" // string | If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default (optional)
    geoLocation := "USA" // string | Assemblies from this location (country and state, or continent) (optional)
    completeOnly := true // bool | only include complete genomes. (optional) (default to false)
    includeAnnotationType := []openapiclient.V1AnnotationForVirusType{openapiclient.v1AnnotationForVirusType("DEFAULT")} // []V1AnnotationForVirusType | Select additional types of annotation to include in the data package.  If unset, no annotation is provided. (optional)
    filename := "filename_example" // string | Output file name. (optional) (default to "ncbi_dataset.zip")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirusApi.Sars2ProteinDownload(context.Background(), proteins).RefseqOnly(refseqOnly).AnnotatedOnly(annotatedOnly).ReleasedSince(releasedSince).Host(host).GeoLocation(geoLocation).CompleteOnly(completeOnly).IncludeAnnotationType(includeAnnotationType).Filename(filename).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirusApi.Sars2ProteinDownload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Sars2ProteinDownload`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `VirusApi.Sars2ProteinDownload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**proteins** | [**[]string**](string.md) | Which proteins to retrieve in the data package | 

### Other Parameters

Other parameters are passed through a pointer to a apiSars2ProteinDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **refseqOnly** | **bool** | If true, limit results to RefSeq genomes. | [default to false]
 **annotatedOnly** | **bool** | If true, limit results to annotated genomes. | [default to false]
 **releasedSince** | **time.Time** | If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as &#39;2020-04-01T00:00:00.000Z&#39; | 
 **host** | **string** | If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default | 
 **geoLocation** | **string** | Assemblies from this location (country and state, or continent) | 
 **completeOnly** | **bool** | only include complete genomes. | [default to false]
 **includeAnnotationType** | [**[]V1AnnotationForVirusType**](V1AnnotationForVirusType.md) | Select additional types of annotation to include in the data package.  If unset, no annotation is provided. | 
 **filename** | **string** | Output file name. | [default to &quot;ncbi_dataset.zip&quot;]

### Return type

[***os.File**](*os.File.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/zip, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Sars2ProteinSummary

> V1DownloadSummary Sars2ProteinSummary(ctx, proteins).RefseqOnly(refseqOnly).AnnotatedOnly(annotatedOnly).ReleasedSince(releasedSince).Host(host).GeoLocation(geoLocation).CompleteOnly(completeOnly).IncludeAnnotationType(includeAnnotationType).Execute()

Summary of SARS-CoV-2 protein and CDS datasets by protein name



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    proteins := []string{"Inner_example"} // []string | Which proteins to retrieve in the data package
    refseqOnly := true // bool | If true, limit results to RefSeq genomes. (optional) (default to false)
    annotatedOnly := true // bool | If true, limit results to annotated genomes. (optional) (default to false)
    releasedSince := time.Now() // time.Time | If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as '2020-04-01T00:00:00.000Z' (optional)
    host := "human" // string | If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default (optional)
    geoLocation := "USA" // string | Assemblies from this location (country and state, or continent) (optional)
    completeOnly := true // bool | only include complete genomes. (optional) (default to false)
    includeAnnotationType := []openapiclient.V1AnnotationForVirusType{openapiclient.v1AnnotationForVirusType("DEFAULT")} // []V1AnnotationForVirusType | Select additional types of annotation to include in the data package.  If unset, no annotation is provided. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirusApi.Sars2ProteinSummary(context.Background(), proteins).RefseqOnly(refseqOnly).AnnotatedOnly(annotatedOnly).ReleasedSince(releasedSince).Host(host).GeoLocation(geoLocation).CompleteOnly(completeOnly).IncludeAnnotationType(includeAnnotationType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirusApi.Sars2ProteinSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Sars2ProteinSummary`: V1DownloadSummary
    fmt.Fprintf(os.Stdout, "Response from `VirusApi.Sars2ProteinSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**proteins** | [**[]string**](string.md) | Which proteins to retrieve in the data package | 

### Other Parameters

Other parameters are passed through a pointer to a apiSars2ProteinSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **refseqOnly** | **bool** | If true, limit results to RefSeq genomes. | [default to false]
 **annotatedOnly** | **bool** | If true, limit results to annotated genomes. | [default to false]
 **releasedSince** | **time.Time** | If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as &#39;2020-04-01T00:00:00.000Z&#39; | 
 **host** | **string** | If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default | 
 **geoLocation** | **string** | Assemblies from this location (country and state, or continent) | 
 **completeOnly** | **bool** | only include complete genomes. | [default to false]
 **includeAnnotationType** | [**[]V1AnnotationForVirusType**](V1AnnotationForVirusType.md) | Select additional types of annotation to include in the data package.  If unset, no annotation is provided. | 

### Return type

[**V1DownloadSummary**](V1DownloadSummary.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Sars2ProteinTable

> V1TabularOutput Sars2ProteinTable(ctx, proteins).RefseqOnly(refseqOnly).AnnotatedOnly(annotatedOnly).ReleasedSince(releasedSince).Host(host).GeoLocation(geoLocation).CompleteOnly(completeOnly).TableFields(tableFields).Format(format).Execute()

Get SARS-CoV-2 protein metadata in a tabular format.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    proteins := []string{"Inner_example"} // []string | Which proteins to retrieve in the data package
    refseqOnly := true // bool | If true, limit results to RefSeq genomes. (optional) (default to false)
    annotatedOnly := true // bool | If true, limit results to annotated genomes. (optional) (default to false)
    releasedSince := time.Now() // time.Time | If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as '2020-04-01T00:00:00.000Z' (optional)
    host := "human" // string | If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default (optional)
    geoLocation := "USA" // string | Assemblies from this location (country and state, or continent) (optional)
    completeOnly := true // bool | only include complete genomes. (optional) (default to false)
    tableFields := []openapiclient.V1VirusTableField{openapiclient.v1VirusTableField("unspecified")} // []V1VirusTableField | Specify which fields to include in the tabular report (optional)
    format := openapiclient.v1TableFormat("tsv") // V1TableFormat | Choose download format (tsv, csv or jsonl) (optional) (default to "tsv")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirusApi.Sars2ProteinTable(context.Background(), proteins).RefseqOnly(refseqOnly).AnnotatedOnly(annotatedOnly).ReleasedSince(releasedSince).Host(host).GeoLocation(geoLocation).CompleteOnly(completeOnly).TableFields(tableFields).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirusApi.Sars2ProteinTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Sars2ProteinTable`: V1TabularOutput
    fmt.Fprintf(os.Stdout, "Response from `VirusApi.Sars2ProteinTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**proteins** | [**[]string**](string.md) | Which proteins to retrieve in the data package | 

### Other Parameters

Other parameters are passed through a pointer to a apiSars2ProteinTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **refseqOnly** | **bool** | If true, limit results to RefSeq genomes. | [default to false]
 **annotatedOnly** | **bool** | If true, limit results to annotated genomes. | [default to false]
 **releasedSince** | **time.Time** | If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as &#39;2020-04-01T00:00:00.000Z&#39; | 
 **host** | **string** | If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default | 
 **geoLocation** | **string** | Assemblies from this location (country and state, or continent) | 
 **completeOnly** | **bool** | only include complete genomes. | [default to false]
 **tableFields** | [**[]V1VirusTableField**](V1VirusTableField.md) | Specify which fields to include in the tabular report | 
 **format** | [**V1TableFormat**](V1TableFormat.md) | Choose download format (tsv, csv or jsonl) | [default to &quot;tsv&quot;]

### Return type

[**V1TabularOutput**](V1TabularOutput.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VirusGenomeDownload

> *os.File VirusGenomeDownload(ctx, taxon).RefseqOnly(refseqOnly).AnnotatedOnly(annotatedOnly).ReleasedSince(releasedSince).Host(host).PangolinClassification(pangolinClassification).GeoLocation(geoLocation).CompleteOnly(completeOnly).ExcludeSequence(excludeSequence).IncludeAnnotationType(includeAnnotationType).Filename(filename).Execute()

Download Coronavirus genome datasets by taxon



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    taxon := "2697049" // string | NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank
    refseqOnly := true // bool | If true, limit results to RefSeq genomes. (optional) (default to false)
    annotatedOnly := true // bool | If true, limit results to annotated genomes. (optional) (default to false)
    releasedSince := time.Now() // time.Time | If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as '2020-04-01T00:00:00.000Z' (optional)
    host := "human" // string | If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default (optional)
    pangolinClassification := "pangolinClassification_example" // string | If set, limit results to genomes classified to this lineage by the PangoLearn tool. (optional)
    geoLocation := "USA" // string | Assemblies from this location (country and state, or continent) (optional)
    completeOnly := true // bool | only include complete genomes. (optional) (default to false)
    excludeSequence := true // bool | Set to true to omit the genomic sequence. (optional) (default to false)
    includeAnnotationType := []openapiclient.V1AnnotationForVirusType{openapiclient.v1AnnotationForVirusType("DEFAULT")} // []V1AnnotationForVirusType | Select additional types of annotation to include in the data package.  If unset, no annotation is provided. (optional)
    filename := "filename_example" // string | Output file name. (optional) (default to "ncbi_dataset.zip")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirusApi.VirusGenomeDownload(context.Background(), taxon).RefseqOnly(refseqOnly).AnnotatedOnly(annotatedOnly).ReleasedSince(releasedSince).Host(host).PangolinClassification(pangolinClassification).GeoLocation(geoLocation).CompleteOnly(completeOnly).ExcludeSequence(excludeSequence).IncludeAnnotationType(includeAnnotationType).Filename(filename).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirusApi.VirusGenomeDownload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VirusGenomeDownload`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `VirusApi.VirusGenomeDownload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxon** | **string** | NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank | 

### Other Parameters

Other parameters are passed through a pointer to a apiVirusGenomeDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **refseqOnly** | **bool** | If true, limit results to RefSeq genomes. | [default to false]
 **annotatedOnly** | **bool** | If true, limit results to annotated genomes. | [default to false]
 **releasedSince** | **time.Time** | If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as &#39;2020-04-01T00:00:00.000Z&#39; | 
 **host** | **string** | If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default | 
 **pangolinClassification** | **string** | If set, limit results to genomes classified to this lineage by the PangoLearn tool. | 
 **geoLocation** | **string** | Assemblies from this location (country and state, or continent) | 
 **completeOnly** | **bool** | only include complete genomes. | [default to false]
 **excludeSequence** | **bool** | Set to true to omit the genomic sequence. | [default to false]
 **includeAnnotationType** | [**[]V1AnnotationForVirusType**](V1AnnotationForVirusType.md) | Select additional types of annotation to include in the data package.  If unset, no annotation is provided. | 
 **filename** | **string** | Output file name. | [default to &quot;ncbi_dataset.zip&quot;]

### Return type

[***os.File**](*os.File.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/zip, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VirusGenomeSummary

> V1DownloadSummary VirusGenomeSummary(ctx, taxon).RefseqOnly(refseqOnly).AnnotatedOnly(annotatedOnly).ReleasedSince(releasedSince).Host(host).PangolinClassification(pangolinClassification).GeoLocation(geoLocation).CompleteOnly(completeOnly).ExcludeSequence(excludeSequence).IncludeAnnotationType(includeAnnotationType).Execute()

Get summary data for Coronaviridae genomes by taxon



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    taxon := "2697049" // string | NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank
    refseqOnly := true // bool | If true, limit results to RefSeq genomes. (optional) (default to false)
    annotatedOnly := true // bool | If true, limit results to annotated genomes. (optional) (default to false)
    releasedSince := time.Now() // time.Time | If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as '2020-04-01T00:00:00.000Z' (optional)
    host := "human" // string | If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default (optional)
    pangolinClassification := "pangolinClassification_example" // string | If set, limit results to genomes classified to this lineage by the PangoLearn tool. (optional)
    geoLocation := "USA" // string | Assemblies from this location (country and state, or continent) (optional)
    completeOnly := true // bool | only include complete genomes. (optional) (default to false)
    excludeSequence := true // bool | Set to true to omit the genomic sequence. (optional) (default to false)
    includeAnnotationType := []openapiclient.V1AnnotationForVirusType{openapiclient.v1AnnotationForVirusType("DEFAULT")} // []V1AnnotationForVirusType | Select additional types of annotation to include in the data package.  If unset, no annotation is provided. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirusApi.VirusGenomeSummary(context.Background(), taxon).RefseqOnly(refseqOnly).AnnotatedOnly(annotatedOnly).ReleasedSince(releasedSince).Host(host).PangolinClassification(pangolinClassification).GeoLocation(geoLocation).CompleteOnly(completeOnly).ExcludeSequence(excludeSequence).IncludeAnnotationType(includeAnnotationType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirusApi.VirusGenomeSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VirusGenomeSummary`: V1DownloadSummary
    fmt.Fprintf(os.Stdout, "Response from `VirusApi.VirusGenomeSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxon** | **string** | NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank | 

### Other Parameters

Other parameters are passed through a pointer to a apiVirusGenomeSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **refseqOnly** | **bool** | If true, limit results to RefSeq genomes. | [default to false]
 **annotatedOnly** | **bool** | If true, limit results to annotated genomes. | [default to false]
 **releasedSince** | **time.Time** | If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as &#39;2020-04-01T00:00:00.000Z&#39; | 
 **host** | **string** | If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default | 
 **pangolinClassification** | **string** | If set, limit results to genomes classified to this lineage by the PangoLearn tool. | 
 **geoLocation** | **string** | Assemblies from this location (country and state, or continent) | 
 **completeOnly** | **bool** | only include complete genomes. | [default to false]
 **excludeSequence** | **bool** | Set to true to omit the genomic sequence. | [default to false]
 **includeAnnotationType** | [**[]V1AnnotationForVirusType**](V1AnnotationForVirusType.md) | Select additional types of annotation to include in the data package.  If unset, no annotation is provided. | 

### Return type

[**V1DownloadSummary**](V1DownloadSummary.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VirusGenomeTable

> V1TabularOutput VirusGenomeTable(ctx, taxon).RefseqOnly(refseqOnly).AnnotatedOnly(annotatedOnly).ReleasedSince(releasedSince).Host(host).PangolinClassification(pangolinClassification).GeoLocation(geoLocation).CompleteOnly(completeOnly).TableFields(tableFields).Format(format).Execute()

Get viral genomic metadata in a tabular format.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    taxon := "2697049" // string | NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank
    refseqOnly := true // bool | If true, limit results to RefSeq genomes. (optional) (default to false)
    annotatedOnly := true // bool | If true, limit results to annotated genomes. (optional) (default to false)
    releasedSince := time.Now() // time.Time | If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as '2020-04-01T00:00:00.000Z' (optional)
    host := "human" // string | If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default (optional)
    pangolinClassification := "pangolinClassification_example" // string | If set, limit results to genomes classified to this lineage by the PangoLearn tool. (optional)
    geoLocation := "USA" // string | Assemblies from this location (country and state, or continent) (optional)
    completeOnly := true // bool | only include complete genomes. (optional) (default to false)
    tableFields := []openapiclient.V1VirusTableField{openapiclient.v1VirusTableField("unspecified")} // []V1VirusTableField | Specify which fields to include in the tabular report (optional)
    format := openapiclient.v1TableFormat("tsv") // V1TableFormat | Choose download format (tsv, csv or jsonl) (optional) (default to "tsv")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirusApi.VirusGenomeTable(context.Background(), taxon).RefseqOnly(refseqOnly).AnnotatedOnly(annotatedOnly).ReleasedSince(releasedSince).Host(host).PangolinClassification(pangolinClassification).GeoLocation(geoLocation).CompleteOnly(completeOnly).TableFields(tableFields).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirusApi.VirusGenomeTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VirusGenomeTable`: V1TabularOutput
    fmt.Fprintf(os.Stdout, "Response from `VirusApi.VirusGenomeTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxon** | **string** | NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank | 

### Other Parameters

Other parameters are passed through a pointer to a apiVirusGenomeTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **refseqOnly** | **bool** | If true, limit results to RefSeq genomes. | [default to false]
 **annotatedOnly** | **bool** | If true, limit results to annotated genomes. | [default to false]
 **releasedSince** | **time.Time** | If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as &#39;2020-04-01T00:00:00.000Z&#39; | 
 **host** | **string** | If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default | 
 **pangolinClassification** | **string** | If set, limit results to genomes classified to this lineage by the PangoLearn tool. | 
 **geoLocation** | **string** | Assemblies from this location (country and state, or continent) | 
 **completeOnly** | **bool** | only include complete genomes. | [default to false]
 **tableFields** | [**[]V1VirusTableField**](V1VirusTableField.md) | Specify which fields to include in the tabular report | 
 **format** | [**V1TableFormat**](V1TableFormat.md) | Choose download format (tsv, csv or jsonl) | [default to &quot;tsv&quot;]

### Return type

[**V1TabularOutput**](V1TabularOutput.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

