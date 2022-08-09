# \VirusApi

All URIs are relative to *https://api.ncbi.nlm.nih.gov/datasets/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Sars2ProteinDownload**](VirusApi.md#Sars2ProteinDownload) | **Get** /virus/taxon/sars2/protein/{proteins}/download | Download SARS-CoV-2 protein and CDS datasets by protein name
[**Sars2ProteinSummary**](VirusApi.md#Sars2ProteinSummary) | **Get** /virus/taxon/sars2/protein/{proteins} | Summary of SARS-CoV-2 protein and CDS datasets by protein name
[**Sars2ProteinTable**](VirusApi.md#Sars2ProteinTable) | **Get** /virus/taxon/sars2/protein/{proteins}/table | Get SARS-CoV-2 protein metadata in a tabular format.
[**VirusGenomeDownload**](VirusApi.md#VirusGenomeDownload) | **Get** /virus/taxon/{taxon}/genome/download | Download a coronavirus genome dataset by taxon
[**VirusGenomeDownloadAccession**](VirusApi.md#VirusGenomeDownloadAccession) | **Get** /virus/accession/{accessions}/genome/download | Download a coronavirus genome dataset by accession
[**VirusGenomeDownloadPost**](VirusApi.md#VirusGenomeDownloadPost) | **Post** /virus/genome/download | Get a coronavirus genome dataset by post
[**VirusGenomeSummary**](VirusApi.md#VirusGenomeSummary) | **Get** /virus/taxon/{taxon}/genome | Get summary data for Coronaviridae genomes by taxon
[**VirusGenomeSummaryAccession**](VirusApi.md#VirusGenomeSummaryAccession) | **Get** /virus/accession/{accessions}/genome | Get summary data for Coronaviridae genomes by accession
[**VirusGenomeSummaryPost**](VirusApi.md#VirusGenomeSummaryPost) | **Post** /virus/genome | Get summary data for Coronaviridae genomes by post
[**VirusGenomeTable**](VirusApi.md#VirusGenomeTable) | **Get** /virus/taxon/{taxon}/genome/table | Get viral genomic metadata in a tabular format.
[**VirusReportsByAcessions**](VirusApi.md#VirusReportsByAcessions) | **Get** /virus/accession/{accessions}/dataset_report | Get virus metadata by accession
[**VirusReportsByPost**](VirusApi.md#VirusReportsByPost) | **Post** /virus | Get virus metadata by POST
[**VirusReportsByTaxon**](VirusApi.md#VirusReportsByTaxon) | **Get** /virus/taxon/{taxon}/dataset_report | Get virus metadata by taxon



## Sars2ProteinDownload

> *os.File Sars2ProteinDownload(ctx, proteins).RefseqOnly(refseqOnly).AnnotatedOnly(annotatedOnly).ReleasedSince(releasedSince).UpdatedSince(updatedSince).Host(host).GeoLocation(geoLocation).CompleteOnly(completeOnly).IncludeAnnotationType(includeAnnotationType).Filename(filename).Execute()

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
    updatedSince := time.Now() // time.Time |  (optional)
    host := "human" // string | If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default (optional)
    geoLocation := "USA" // string | Assemblies from this location (country and state, or continent) (optional)
    completeOnly := true // bool | only include complete genomes. (optional) (default to false)
    includeAnnotationType := []openapiclient.V1AnnotationForVirusType{openapiclient.v1AnnotationForVirusType("DEFAULT")} // []V1AnnotationForVirusType | Select additional types of annotation to include in the data package.  If unset, no annotation is provided. (optional)
    filename := "filename_example" // string | Output file name. (optional) (default to "ncbi_dataset.zip")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirusApi.Sars2ProteinDownload(context.Background(), proteins).RefseqOnly(refseqOnly).AnnotatedOnly(annotatedOnly).ReleasedSince(releasedSince).UpdatedSince(updatedSince).Host(host).GeoLocation(geoLocation).CompleteOnly(completeOnly).IncludeAnnotationType(includeAnnotationType).Filename(filename).Execute()
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
 **updatedSince** | **time.Time** |  | 
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

> V1DownloadSummary Sars2ProteinSummary(ctx, proteins).RefseqOnly(refseqOnly).AnnotatedOnly(annotatedOnly).ReleasedSince(releasedSince).UpdatedSince(updatedSince).Host(host).GeoLocation(geoLocation).CompleteOnly(completeOnly).IncludeAnnotationType(includeAnnotationType).Execute()

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
    updatedSince := time.Now() // time.Time |  (optional)
    host := "human" // string | If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default (optional)
    geoLocation := "USA" // string | Assemblies from this location (country and state, or continent) (optional)
    completeOnly := true // bool | only include complete genomes. (optional) (default to false)
    includeAnnotationType := []openapiclient.V1AnnotationForVirusType{openapiclient.v1AnnotationForVirusType("DEFAULT")} // []V1AnnotationForVirusType | Select additional types of annotation to include in the data package.  If unset, no annotation is provided. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirusApi.Sars2ProteinSummary(context.Background(), proteins).RefseqOnly(refseqOnly).AnnotatedOnly(annotatedOnly).ReleasedSince(releasedSince).UpdatedSince(updatedSince).Host(host).GeoLocation(geoLocation).CompleteOnly(completeOnly).IncludeAnnotationType(includeAnnotationType).Execute()
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
 **updatedSince** | **time.Time** |  | 
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

> V1TabularOutput Sars2ProteinTable(ctx, proteins).RefseqOnly(refseqOnly).AnnotatedOnly(annotatedOnly).ReleasedSince(releasedSince).UpdatedSince(updatedSince).Host(host).GeoLocation(geoLocation).CompleteOnly(completeOnly).TableFields(tableFields).Format(format).Execute()

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
    updatedSince := time.Now() // time.Time |  (optional)
    host := "human" // string | If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default (optional)
    geoLocation := "USA" // string | Assemblies from this location (country and state, or continent) (optional)
    completeOnly := true // bool | only include complete genomes. (optional) (default to false)
    tableFields := []openapiclient.V1VirusTableField{openapiclient.v1VirusTableField("unspecified")} // []V1VirusTableField | Specify which fields to include in the tabular report (optional)
    format := openapiclient.v1TableFormat("tsv") // V1TableFormat | Choose download format (tsv, csv or jsonl) (optional) (default to "tsv")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirusApi.Sars2ProteinTable(context.Background(), proteins).RefseqOnly(refseqOnly).AnnotatedOnly(annotatedOnly).ReleasedSince(releasedSince).UpdatedSince(updatedSince).Host(host).GeoLocation(geoLocation).CompleteOnly(completeOnly).TableFields(tableFields).Format(format).Execute()
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
 **updatedSince** | **time.Time** |  | 
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

> *os.File VirusGenomeDownload(ctx, taxon).RefseqOnly(refseqOnly).AnnotatedOnly(annotatedOnly).ReleasedSince(releasedSince).UpdatedSince(updatedSince).Host(host).PangolinClassification(pangolinClassification).GeoLocation(geoLocation).CompleteOnly(completeOnly).ExcludeSequence(excludeSequence).IncludeAnnotationType(includeAnnotationType).Filename(filename).Execute()

Download a coronavirus genome dataset by taxon



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
    updatedSince := time.Now() // time.Time |  (optional)
    host := "human" // string | If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default (optional)
    pangolinClassification := "pangolinClassification_example" // string | If set, limit results to genomes classified to this lineage by the PangoLearn tool. (optional)
    geoLocation := "USA" // string | Assemblies from this location (country and state, or continent) (optional)
    completeOnly := true // bool | only include complete genomes. (optional) (default to false)
    excludeSequence := true // bool | Set to true to omit the genomic sequence. (optional) (default to false)
    includeAnnotationType := []openapiclient.V1AnnotationForVirusType{openapiclient.v1AnnotationForVirusType("DEFAULT")} // []V1AnnotationForVirusType | Select additional types of annotation to include in the data package.  If unset, no annotation is provided. (optional)
    filename := "filename_example" // string | Output file name. (optional) (default to "ncbi_dataset.zip")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirusApi.VirusGenomeDownload(context.Background(), taxon).RefseqOnly(refseqOnly).AnnotatedOnly(annotatedOnly).ReleasedSince(releasedSince).UpdatedSince(updatedSince).Host(host).PangolinClassification(pangolinClassification).GeoLocation(geoLocation).CompleteOnly(completeOnly).ExcludeSequence(excludeSequence).IncludeAnnotationType(includeAnnotationType).Filename(filename).Execute()
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
 **updatedSince** | **time.Time** |  | 
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


## VirusGenomeDownloadAccession

> *os.File VirusGenomeDownloadAccession(ctx, accessions).RefseqOnly(refseqOnly).AnnotatedOnly(annotatedOnly).ReleasedSince(releasedSince).UpdatedSince(updatedSince).Host(host).PangolinClassification(pangolinClassification).GeoLocation(geoLocation).CompleteOnly(completeOnly).ExcludeSequence(excludeSequence).IncludeAnnotationType(includeAnnotationType).Filename(filename).Execute()

Download a coronavirus genome dataset by accession



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
    accessions := []string{"Inner_example"} // []string | Accessions accessions = 16;
    refseqOnly := true // bool | If true, limit results to RefSeq genomes. (optional) (default to false)
    annotatedOnly := true // bool | If true, limit results to annotated genomes. (optional) (default to false)
    releasedSince := time.Now() // time.Time | If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as '2020-04-01T00:00:00.000Z' (optional)
    updatedSince := time.Now() // time.Time |  (optional)
    host := "human" // string | If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default (optional)
    pangolinClassification := "pangolinClassification_example" // string | If set, limit results to genomes classified to this lineage by the PangoLearn tool. (optional)
    geoLocation := "USA" // string | Assemblies from this location (country and state, or continent) (optional)
    completeOnly := true // bool | only include complete genomes. (optional) (default to false)
    excludeSequence := true // bool | Set to true to omit the genomic sequence. (optional) (default to false)
    includeAnnotationType := []openapiclient.V1AnnotationForVirusType{openapiclient.v1AnnotationForVirusType("DEFAULT")} // []V1AnnotationForVirusType | Select additional types of annotation to include in the data package.  If unset, no annotation is provided. (optional)
    filename := "filename_example" // string | Output file name. (optional) (default to "ncbi_dataset.zip")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirusApi.VirusGenomeDownloadAccession(context.Background(), accessions).RefseqOnly(refseqOnly).AnnotatedOnly(annotatedOnly).ReleasedSince(releasedSince).UpdatedSince(updatedSince).Host(host).PangolinClassification(pangolinClassification).GeoLocation(geoLocation).CompleteOnly(completeOnly).ExcludeSequence(excludeSequence).IncludeAnnotationType(includeAnnotationType).Filename(filename).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirusApi.VirusGenomeDownloadAccession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VirusGenomeDownloadAccession`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `VirusApi.VirusGenomeDownloadAccession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessions** | [**[]string**](string.md) | Accessions accessions &#x3D; 16; | 

### Other Parameters

Other parameters are passed through a pointer to a apiVirusGenomeDownloadAccessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **refseqOnly** | **bool** | If true, limit results to RefSeq genomes. | [default to false]
 **annotatedOnly** | **bool** | If true, limit results to annotated genomes. | [default to false]
 **releasedSince** | **time.Time** | If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as &#39;2020-04-01T00:00:00.000Z&#39; | 
 **updatedSince** | **time.Time** |  | 
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


## VirusGenomeDownloadPost

> *os.File VirusGenomeDownloadPost(ctx).V1VirusDatasetRequest(v1VirusDatasetRequest).Filename(filename).Execute()

Get a coronavirus genome dataset by post



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    v1VirusDatasetRequest := *openapiclient.NewV1VirusDatasetRequest() // V1VirusDatasetRequest | 
    filename := "filename_example" // string | Output file name. (optional) (default to "ncbi_dataset.zip")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirusApi.VirusGenomeDownloadPost(context.Background()).V1VirusDatasetRequest(v1VirusDatasetRequest).Filename(filename).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirusApi.VirusGenomeDownloadPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VirusGenomeDownloadPost`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `VirusApi.VirusGenomeDownloadPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVirusGenomeDownloadPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1VirusDatasetRequest** | [**V1VirusDatasetRequest**](V1VirusDatasetRequest.md) |  | 
 **filename** | **string** | Output file name. | [default to &quot;ncbi_dataset.zip&quot;]

### Return type

[***os.File**](*os.File.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/zip, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VirusGenomeSummary

> V1DownloadSummary VirusGenomeSummary(ctx, taxon).RefseqOnly(refseqOnly).AnnotatedOnly(annotatedOnly).ReleasedSince(releasedSince).UpdatedSince(updatedSince).Host(host).PangolinClassification(pangolinClassification).GeoLocation(geoLocation).CompleteOnly(completeOnly).ExcludeSequence(excludeSequence).IncludeAnnotationType(includeAnnotationType).Execute()

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
    updatedSince := time.Now() // time.Time |  (optional)
    host := "human" // string | If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default (optional)
    pangolinClassification := "pangolinClassification_example" // string | If set, limit results to genomes classified to this lineage by the PangoLearn tool. (optional)
    geoLocation := "USA" // string | Assemblies from this location (country and state, or continent) (optional)
    completeOnly := true // bool | only include complete genomes. (optional) (default to false)
    excludeSequence := true // bool | Set to true to omit the genomic sequence. (optional) (default to false)
    includeAnnotationType := []openapiclient.V1AnnotationForVirusType{openapiclient.v1AnnotationForVirusType("DEFAULT")} // []V1AnnotationForVirusType | Select additional types of annotation to include in the data package.  If unset, no annotation is provided. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirusApi.VirusGenomeSummary(context.Background(), taxon).RefseqOnly(refseqOnly).AnnotatedOnly(annotatedOnly).ReleasedSince(releasedSince).UpdatedSince(updatedSince).Host(host).PangolinClassification(pangolinClassification).GeoLocation(geoLocation).CompleteOnly(completeOnly).ExcludeSequence(excludeSequence).IncludeAnnotationType(includeAnnotationType).Execute()
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
 **updatedSince** | **time.Time** |  | 
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


## VirusGenomeSummaryAccession

> V1DownloadSummary VirusGenomeSummaryAccession(ctx, accessions).RefseqOnly(refseqOnly).AnnotatedOnly(annotatedOnly).ReleasedSince(releasedSince).UpdatedSince(updatedSince).Host(host).PangolinClassification(pangolinClassification).GeoLocation(geoLocation).CompleteOnly(completeOnly).ExcludeSequence(excludeSequence).IncludeAnnotationType(includeAnnotationType).Execute()

Get summary data for Coronaviridae genomes by accession



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
    accessions := []string{"Inner_example"} // []string | Accessions accessions = 16;
    refseqOnly := true // bool | If true, limit results to RefSeq genomes. (optional) (default to false)
    annotatedOnly := true // bool | If true, limit results to annotated genomes. (optional) (default to false)
    releasedSince := time.Now() // time.Time | If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as '2020-04-01T00:00:00.000Z' (optional)
    updatedSince := time.Now() // time.Time |  (optional)
    host := "human" // string | If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default (optional)
    pangolinClassification := "pangolinClassification_example" // string | If set, limit results to genomes classified to this lineage by the PangoLearn tool. (optional)
    geoLocation := "USA" // string | Assemblies from this location (country and state, or continent) (optional)
    completeOnly := true // bool | only include complete genomes. (optional) (default to false)
    excludeSequence := true // bool | Set to true to omit the genomic sequence. (optional) (default to false)
    includeAnnotationType := []openapiclient.V1AnnotationForVirusType{openapiclient.v1AnnotationForVirusType("DEFAULT")} // []V1AnnotationForVirusType | Select additional types of annotation to include in the data package.  If unset, no annotation is provided. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirusApi.VirusGenomeSummaryAccession(context.Background(), accessions).RefseqOnly(refseqOnly).AnnotatedOnly(annotatedOnly).ReleasedSince(releasedSince).UpdatedSince(updatedSince).Host(host).PangolinClassification(pangolinClassification).GeoLocation(geoLocation).CompleteOnly(completeOnly).ExcludeSequence(excludeSequence).IncludeAnnotationType(includeAnnotationType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirusApi.VirusGenomeSummaryAccession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VirusGenomeSummaryAccession`: V1DownloadSummary
    fmt.Fprintf(os.Stdout, "Response from `VirusApi.VirusGenomeSummaryAccession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessions** | [**[]string**](string.md) | Accessions accessions &#x3D; 16; | 

### Other Parameters

Other parameters are passed through a pointer to a apiVirusGenomeSummaryAccessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **refseqOnly** | **bool** | If true, limit results to RefSeq genomes. | [default to false]
 **annotatedOnly** | **bool** | If true, limit results to annotated genomes. | [default to false]
 **releasedSince** | **time.Time** | If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as &#39;2020-04-01T00:00:00.000Z&#39; | 
 **updatedSince** | **time.Time** |  | 
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


## VirusGenomeSummaryPost

> V1DownloadSummary VirusGenomeSummaryPost(ctx).V1VirusDatasetRequest(v1VirusDatasetRequest).Execute()

Get summary data for Coronaviridae genomes by post



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    v1VirusDatasetRequest := *openapiclient.NewV1VirusDatasetRequest() // V1VirusDatasetRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirusApi.VirusGenomeSummaryPost(context.Background()).V1VirusDatasetRequest(v1VirusDatasetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirusApi.VirusGenomeSummaryPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VirusGenomeSummaryPost`: V1DownloadSummary
    fmt.Fprintf(os.Stdout, "Response from `VirusApi.VirusGenomeSummaryPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVirusGenomeSummaryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1VirusDatasetRequest** | [**V1VirusDatasetRequest**](V1VirusDatasetRequest.md) |  | 

### Return type

[**V1DownloadSummary**](V1DownloadSummary.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VirusGenomeTable

> V1TabularOutput VirusGenomeTable(ctx, taxon).RefseqOnly(refseqOnly).AnnotatedOnly(annotatedOnly).ReleasedSince(releasedSince).UpdatedSince(updatedSince).Host(host).PangolinClassification(pangolinClassification).GeoLocation(geoLocation).CompleteOnly(completeOnly).TableFields(tableFields).Format(format).Execute()

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
    updatedSince := time.Now() // time.Time |  (optional)
    host := "human" // string | If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default (optional)
    pangolinClassification := "pangolinClassification_example" // string | If set, limit results to genomes classified to this lineage by the PangoLearn tool. (optional)
    geoLocation := "USA" // string | Assemblies from this location (country and state, or continent) (optional)
    completeOnly := true // bool | only include complete genomes. (optional) (default to false)
    tableFields := []openapiclient.V1VirusTableField{openapiclient.v1VirusTableField("unspecified")} // []V1VirusTableField | Specify which fields to include in the tabular report (optional)
    format := openapiclient.v1TableFormat("tsv") // V1TableFormat | Choose download format (tsv, csv or jsonl) (optional) (default to "tsv")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirusApi.VirusGenomeTable(context.Background(), taxon).RefseqOnly(refseqOnly).AnnotatedOnly(annotatedOnly).ReleasedSince(releasedSince).UpdatedSince(updatedSince).Host(host).PangolinClassification(pangolinClassification).GeoLocation(geoLocation).CompleteOnly(completeOnly).TableFields(tableFields).Format(format).Execute()
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
 **updatedSince** | **time.Time** |  | 
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


## VirusReportsByAcessions

> V1reportsVirusDataReportPage VirusReportsByAcessions(ctx, accessions).FilterRefseqOnly(filterRefseqOnly).FilterAnnotatedOnly(filterAnnotatedOnly).FilterReleasedSince(filterReleasedSince).FilterUpdatedSince(filterUpdatedSince).FilterHost(filterHost).FilterPangolinClassification(filterPangolinClassification).FilterGeoLocation(filterGeoLocation).FilterCompleteOnly(filterCompleteOnly).ReturnedContent(returnedContent).TableFields(tableFields).PageSize(pageSize).PageToken(pageToken).Execute()

Get virus metadata by accession



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
    accessions := []string{"Inner_example"} // []string | genome sequence accessions
    filterRefseqOnly := true // bool | If true, limit results to RefSeq genomes. (optional) (default to false)
    filterAnnotatedOnly := true // bool | If true, limit results to annotated genomes. (optional) (default to false)
    filterReleasedSince := time.Now() // time.Time | If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as '2020-04-01T00:00:00.000Z' (optional)
    filterUpdatedSince := time.Now() // time.Time |  (optional)
    filterHost := "human" // string | If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default (optional)
    filterPangolinClassification := "filterPangolinClassification_example" // string | If set, limit results to genomes classified to this lineage by the PangoLearn tool. (optional)
    filterGeoLocation := "USA" // string | Assemblies from this location (country and state, or continent) (optional)
    filterCompleteOnly := true // bool | only include complete genomes. (optional) (default to false)
    returnedContent := openapiclient.v1VirusDataReportRequestContentType("COMPLETE") // V1VirusDataReportRequestContentType | Return either virus genome accessions, or complete virus metadata (optional) (default to "COMPLETE")
    tableFields := []string{"TableFields_example"} // []string | Specify which fields to include in the tabular report (optional)
    pageSize := int32(56) // int32 | The maximum number of virus data reports to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, `page_token` can be used to retrieve the remaining results. (optional) (default to 20)
    pageToken := "pageToken_example" // string | A page token is returned from a `GetVirusDataReports` call with more than `page_size` results. Use this token, along with the previous `VirusDataReportRequest` parameters, to retrieve the next page of results. When `page_token` is empty, all results have been retrieved. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirusApi.VirusReportsByAcessions(context.Background(), accessions).FilterRefseqOnly(filterRefseqOnly).FilterAnnotatedOnly(filterAnnotatedOnly).FilterReleasedSince(filterReleasedSince).FilterUpdatedSince(filterUpdatedSince).FilterHost(filterHost).FilterPangolinClassification(filterPangolinClassification).FilterGeoLocation(filterGeoLocation).FilterCompleteOnly(filterCompleteOnly).ReturnedContent(returnedContent).TableFields(tableFields).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirusApi.VirusReportsByAcessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VirusReportsByAcessions`: V1reportsVirusDataReportPage
    fmt.Fprintf(os.Stdout, "Response from `VirusApi.VirusReportsByAcessions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessions** | [**[]string**](string.md) | genome sequence accessions | 

### Other Parameters

Other parameters are passed through a pointer to a apiVirusReportsByAcessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterRefseqOnly** | **bool** | If true, limit results to RefSeq genomes. | [default to false]
 **filterAnnotatedOnly** | **bool** | If true, limit results to annotated genomes. | [default to false]
 **filterReleasedSince** | **time.Time** | If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as &#39;2020-04-01T00:00:00.000Z&#39; | 
 **filterUpdatedSince** | **time.Time** |  | 
 **filterHost** | **string** | If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default | 
 **filterPangolinClassification** | **string** | If set, limit results to genomes classified to this lineage by the PangoLearn tool. | 
 **filterGeoLocation** | **string** | Assemblies from this location (country and state, or continent) | 
 **filterCompleteOnly** | **bool** | only include complete genomes. | [default to false]
 **returnedContent** | [**V1VirusDataReportRequestContentType**](V1VirusDataReportRequestContentType.md) | Return either virus genome accessions, or complete virus metadata | [default to &quot;COMPLETE&quot;]
 **tableFields** | **[]string** | Specify which fields to include in the tabular report | 
 **pageSize** | **int32** | The maximum number of virus data reports to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, &#x60;page_token&#x60; can be used to retrieve the remaining results. | [default to 20]
 **pageToken** | **string** | A page token is returned from a &#x60;GetVirusDataReports&#x60; call with more than &#x60;page_size&#x60; results. Use this token, along with the previous &#x60;VirusDataReportRequest&#x60; parameters, to retrieve the next page of results. When &#x60;page_token&#x60; is empty, all results have been retrieved. | 

### Return type

[**V1reportsVirusDataReportPage**](V1reportsVirusDataReportPage.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/tab-separated-values

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VirusReportsByPost

> V1reportsVirusDataReportPage VirusReportsByPost(ctx).V1VirusDataReportRequest(v1VirusDataReportRequest).Execute()

Get virus metadata by POST



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    v1VirusDataReportRequest := *openapiclient.NewV1VirusDataReportRequest() // V1VirusDataReportRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirusApi.VirusReportsByPost(context.Background()).V1VirusDataReportRequest(v1VirusDataReportRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirusApi.VirusReportsByPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VirusReportsByPost`: V1reportsVirusDataReportPage
    fmt.Fprintf(os.Stdout, "Response from `VirusApi.VirusReportsByPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVirusReportsByPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1VirusDataReportRequest** | [**V1VirusDataReportRequest**](V1VirusDataReportRequest.md) |  | 

### Return type

[**V1reportsVirusDataReportPage**](V1reportsVirusDataReportPage.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/tab-separated-values

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VirusReportsByTaxon

> V1reportsVirusDataReportPage VirusReportsByTaxon(ctx, taxon).FilterRefseqOnly(filterRefseqOnly).FilterAnnotatedOnly(filterAnnotatedOnly).FilterReleasedSince(filterReleasedSince).FilterUpdatedSince(filterUpdatedSince).FilterHost(filterHost).FilterPangolinClassification(filterPangolinClassification).FilterGeoLocation(filterGeoLocation).FilterCompleteOnly(filterCompleteOnly).ReturnedContent(returnedContent).TableFields(tableFields).PageSize(pageSize).PageToken(pageToken).Execute()

Get virus metadata by taxon



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
    filterRefseqOnly := true // bool | If true, limit results to RefSeq genomes. (optional) (default to false)
    filterAnnotatedOnly := true // bool | If true, limit results to annotated genomes. (optional) (default to false)
    filterReleasedSince := time.Now() // time.Time | If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as '2020-04-01T00:00:00.000Z' (optional)
    filterUpdatedSince := time.Now() // time.Time |  (optional)
    filterHost := "human" // string | If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default (optional)
    filterPangolinClassification := "filterPangolinClassification_example" // string | If set, limit results to genomes classified to this lineage by the PangoLearn tool. (optional)
    filterGeoLocation := "USA" // string | Assemblies from this location (country and state, or continent) (optional)
    filterCompleteOnly := true // bool | only include complete genomes. (optional) (default to false)
    returnedContent := openapiclient.v1VirusDataReportRequestContentType("COMPLETE") // V1VirusDataReportRequestContentType | Return either virus genome accessions, or complete virus metadata (optional) (default to "COMPLETE")
    tableFields := []string{"TableFields_example"} // []string | Specify which fields to include in the tabular report (optional)
    pageSize := int32(56) // int32 | The maximum number of virus data reports to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, `page_token` can be used to retrieve the remaining results. (optional) (default to 20)
    pageToken := "pageToken_example" // string | A page token is returned from a `GetVirusDataReports` call with more than `page_size` results. Use this token, along with the previous `VirusDataReportRequest` parameters, to retrieve the next page of results. When `page_token` is empty, all results have been retrieved. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VirusApi.VirusReportsByTaxon(context.Background(), taxon).FilterRefseqOnly(filterRefseqOnly).FilterAnnotatedOnly(filterAnnotatedOnly).FilterReleasedSince(filterReleasedSince).FilterUpdatedSince(filterUpdatedSince).FilterHost(filterHost).FilterPangolinClassification(filterPangolinClassification).FilterGeoLocation(filterGeoLocation).FilterCompleteOnly(filterCompleteOnly).ReturnedContent(returnedContent).TableFields(tableFields).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirusApi.VirusReportsByTaxon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VirusReportsByTaxon`: V1reportsVirusDataReportPage
    fmt.Fprintf(os.Stdout, "Response from `VirusApi.VirusReportsByTaxon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxon** | **string** | NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank | 

### Other Parameters

Other parameters are passed through a pointer to a apiVirusReportsByTaxonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filterRefseqOnly** | **bool** | If true, limit results to RefSeq genomes. | [default to false]
 **filterAnnotatedOnly** | **bool** | If true, limit results to annotated genomes. | [default to false]
 **filterReleasedSince** | **time.Time** | If set, limit results to viral genomes that have been released after a specified date (and optionally, time). April 1, 2020 midnight UTC should be formatted as &#39;2020-04-01T00:00:00.000Z&#39; | 
 **filterUpdatedSince** | **time.Time** |  | 
 **filterHost** | **string** | If set, limit results to genomes extracted from this host (Taxonomy ID or name) All hosts by default | 
 **filterPangolinClassification** | **string** | If set, limit results to genomes classified to this lineage by the PangoLearn tool. | 
 **filterGeoLocation** | **string** | Assemblies from this location (country and state, or continent) | 
 **filterCompleteOnly** | **bool** | only include complete genomes. | [default to false]
 **returnedContent** | [**V1VirusDataReportRequestContentType**](V1VirusDataReportRequestContentType.md) | Return either virus genome accessions, or complete virus metadata | [default to &quot;COMPLETE&quot;]
 **tableFields** | **[]string** | Specify which fields to include in the tabular report | 
 **pageSize** | **int32** | The maximum number of virus data reports to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, &#x60;page_token&#x60; can be used to retrieve the remaining results. | [default to 20]
 **pageToken** | **string** | A page token is returned from a &#x60;GetVirusDataReports&#x60; call with more than &#x60;page_size&#x60; results. Use this token, along with the previous &#x60;VirusDataReportRequest&#x60; parameters, to retrieve the next page of results. When &#x60;page_token&#x60; is empty, all results have been retrieved. | 

### Return type

[**V1reportsVirusDataReportPage**](V1reportsVirusDataReportPage.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/tab-separated-values

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

