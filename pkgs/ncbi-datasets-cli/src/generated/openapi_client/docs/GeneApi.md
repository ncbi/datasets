# \GeneApi

All URIs are relative to *https://api.ncbi.nlm.nih.gov/datasets/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadGenePackage**](GeneApi.md#DownloadGenePackage) | **Get** /gene/id/{gene_ids}/download | Get a gene dataset by gene ID
[**DownloadGenePackagePost**](GeneApi.md#DownloadGenePackagePost) | **Post** /gene/download | Get a gene dataset by POST
[**GeneDownloadSummaryByAccession**](GeneApi.md#GeneDownloadSummaryByAccession) | **Get** /gene/accession/{accessions}/download_summary | Get gene download summary by RefSeq Accession
[**GeneDownloadSummaryById**](GeneApi.md#GeneDownloadSummaryById) | **Get** /gene/id/{gene_ids}/download_summary | Get gene download summary by GeneID
[**GeneDownloadSummaryByPost**](GeneApi.md#GeneDownloadSummaryByPost) | **Post** /gene/download_summary | Get gene download summary
[**GeneDownloadSummaryByTaxAndSymbol**](GeneApi.md#GeneDownloadSummaryByTaxAndSymbol) | **Get** /gene/symbol/{symbols}/taxon/{taxon}/download_summary | Get gene download summary by gene symbol
[**GeneMetadataByAccession**](GeneApi.md#GeneMetadataByAccession) | **Get** /gene/accession/{accessions} | Get gene metadata by RefSeq Accession
[**GeneMetadataById**](GeneApi.md#GeneMetadataById) | **Get** /gene/id/{gene_ids} | Get gene metadata by GeneID
[**GeneMetadataByPost**](GeneApi.md#GeneMetadataByPost) | **Post** /gene | Get gene metadata as JSON
[**GeneMetadataByTaxAndSymbol**](GeneApi.md#GeneMetadataByTaxAndSymbol) | **Get** /gene/symbol/{symbols}/taxon/{taxon} | Get gene metadata by gene symbol
[**GeneMetadataStreamByPost**](GeneApi.md#GeneMetadataStreamByPost) | **Post** /gene/stream | Get gene metadata
[**GeneOrthologsById**](GeneApi.md#GeneOrthologsById) | **Get** /gene/id/{gene_id}/orthologs | Get gene orthologs by gene ID
[**GeneTaxNameQuery**](GeneApi.md#GeneTaxNameQuery) | **Get** /gene/taxon_suggest/{taxon_query} | Get a list of taxonomy names and IDs found in the gene dataset given a partial taxonomic name
[**GeneTaxTree**](GeneApi.md#GeneTaxTree) | **Get** /gene/taxon/{taxon}/tree | Get a taxonomic subtree by taxonomic identifier



## DownloadGenePackage

> *os.File DownloadGenePackage(ctx, geneIds).IncludeAnnotationType(includeAnnotationType).FastaFilter(fastaFilter).Filename(filename).Execute()

Get a gene dataset by gene ID



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
    geneIds := []int32{int32(123)} // []int32 | NCBI gene ids
    includeAnnotationType := []openapiclient.V1Fasta{openapiclient.v1Fasta("FASTA_UNSPECIFIED")} // []V1Fasta | Select additional types of annotation to include in the data package.  If unset, no annotation is provided. (optional)
    fastaFilter := []string{"Inner_example"} // []string | Limit the FASTA sequences in the datasets package to these transcript and protein accessions (optional)
    filename := "filename_example" // string | Output file name. (optional) (default to "ncbi_dataset.zip")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GeneApi.DownloadGenePackage(context.Background(), geneIds).IncludeAnnotationType(includeAnnotationType).FastaFilter(fastaFilter).Filename(filename).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GeneApi.DownloadGenePackage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadGenePackage`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `GeneApi.DownloadGenePackage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**geneIds** | [**[]int32**](int32.md) | NCBI gene ids | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadGenePackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeAnnotationType** | [**[]V1Fasta**](V1Fasta.md) | Select additional types of annotation to include in the data package.  If unset, no annotation is provided. | 
 **fastaFilter** | **[]string** | Limit the FASTA sequences in the datasets package to these transcript and protein accessions | 
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


## DownloadGenePackagePost

> *os.File DownloadGenePackagePost(ctx).V1GeneDatasetRequest(v1GeneDatasetRequest).Filename(filename).Execute()

Get a gene dataset by POST



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
    v1GeneDatasetRequest := *openapiclient.NewV1GeneDatasetRequest() // V1GeneDatasetRequest | 
    filename := "filename_example" // string | Output file name. (optional) (default to "ncbi_dataset.zip")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GeneApi.DownloadGenePackagePost(context.Background()).V1GeneDatasetRequest(v1GeneDatasetRequest).Filename(filename).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GeneApi.DownloadGenePackagePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadGenePackagePost`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `GeneApi.DownloadGenePackagePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDownloadGenePackagePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1GeneDatasetRequest** | [**V1GeneDatasetRequest**](V1GeneDatasetRequest.md) |  | 
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


## GeneDownloadSummaryByAccession

> V1DownloadSummary GeneDownloadSummaryByAccession(ctx, accessions).FastaFilter(fastaFilter).Execute()

Get gene download summary by RefSeq Accession



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
    accessions := []string{"Inner_example"} // []string | RNA or Protein accessions.
    fastaFilter := []string{"Inner_example"} // []string | Limit the FASTA sequences in the datasets package to these transcript and protein accessions (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GeneApi.GeneDownloadSummaryByAccession(context.Background(), accessions).FastaFilter(fastaFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GeneApi.GeneDownloadSummaryByAccession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GeneDownloadSummaryByAccession`: V1DownloadSummary
    fmt.Fprintf(os.Stdout, "Response from `GeneApi.GeneDownloadSummaryByAccession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessions** | [**[]string**](string.md) | RNA or Protein accessions. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGeneDownloadSummaryByAccessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fastaFilter** | **[]string** | Limit the FASTA sequences in the datasets package to these transcript and protein accessions | 

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


## GeneDownloadSummaryById

> V1DownloadSummary GeneDownloadSummaryById(ctx, geneIds).FastaFilter(fastaFilter).Execute()

Get gene download summary by GeneID



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
    geneIds := []int32{int32(123)} // []int32 | NCBI gene ids
    fastaFilter := []string{"Inner_example"} // []string | Limit the FASTA sequences in the datasets package to these transcript and protein accessions (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GeneApi.GeneDownloadSummaryById(context.Background(), geneIds).FastaFilter(fastaFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GeneApi.GeneDownloadSummaryById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GeneDownloadSummaryById`: V1DownloadSummary
    fmt.Fprintf(os.Stdout, "Response from `GeneApi.GeneDownloadSummaryById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**geneIds** | [**[]int32**](int32.md) | NCBI gene ids | 

### Other Parameters

Other parameters are passed through a pointer to a apiGeneDownloadSummaryByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fastaFilter** | **[]string** | Limit the FASTA sequences in the datasets package to these transcript and protein accessions | 

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


## GeneDownloadSummaryByPost

> V1DownloadSummary GeneDownloadSummaryByPost(ctx).V1GeneDatasetRequest(v1GeneDatasetRequest).Execute()

Get gene download summary



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
    v1GeneDatasetRequest := *openapiclient.NewV1GeneDatasetRequest() // V1GeneDatasetRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GeneApi.GeneDownloadSummaryByPost(context.Background()).V1GeneDatasetRequest(v1GeneDatasetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GeneApi.GeneDownloadSummaryByPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GeneDownloadSummaryByPost`: V1DownloadSummary
    fmt.Fprintf(os.Stdout, "Response from `GeneApi.GeneDownloadSummaryByPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGeneDownloadSummaryByPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1GeneDatasetRequest** | [**V1GeneDatasetRequest**](V1GeneDatasetRequest.md) |  | 

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


## GeneDownloadSummaryByTaxAndSymbol

> V1DownloadSummary GeneDownloadSummaryByTaxAndSymbol(ctx, symbols, taxon).FastaFilter(fastaFilter).Execute()

Get gene download summary by gene symbol



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
    symbols := []string{"Inner_example"} // []string | Gene symbol
    taxon := "9606" // string | Taxon for provided gene symbol
    fastaFilter := []string{"Inner_example"} // []string | Limit the FASTA sequences in the datasets package to these transcript and protein accessions (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GeneApi.GeneDownloadSummaryByTaxAndSymbol(context.Background(), symbols, taxon).FastaFilter(fastaFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GeneApi.GeneDownloadSummaryByTaxAndSymbol``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GeneDownloadSummaryByTaxAndSymbol`: V1DownloadSummary
    fmt.Fprintf(os.Stdout, "Response from `GeneApi.GeneDownloadSummaryByTaxAndSymbol`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**symbols** | [**[]string**](string.md) | Gene symbol | 
**taxon** | **string** | Taxon for provided gene symbol | 

### Other Parameters

Other parameters are passed through a pointer to a apiGeneDownloadSummaryByTaxAndSymbolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fastaFilter** | **[]string** | Limit the FASTA sequences in the datasets package to these transcript and protein accessions | 

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


## GeneMetadataByAccession

> V1GeneMetadata GeneMetadataByAccession(ctx, accessions).ReturnedContent(returnedContent).SortSchemaField(sortSchemaField).SortSchemaDirection(sortSchemaDirection).Limit(limit).Execute()

Get gene metadata by RefSeq Accession



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
    accessions := []string{"Inner_example"} // []string | RNA or Protein accessions.
    returnedContent := openapiclient.v1GeneDatasetRequestContentType("COMPLETE") // V1GeneDatasetRequestContentType | Return either gene-ids, or entire gene metadata (optional) (default to "COMPLETE")
    sortSchemaField := openapiclient.v1GeneDatasetRequestSortField("SORT_FIELD_UNSPECIFIED") // V1GeneDatasetRequestSortField | Select a field to sort on. (optional) (default to "SORT_FIELD_UNSPECIFIED")
    sortSchemaDirection := openapiclient.v1SortDirection("SORT_DIRECTION_UNSPECIFIED") // V1SortDirection | Select a direction for the sort. (optional) (default to "SORT_DIRECTION_UNSPECIFIED")
    limit := "limit_example" // string | Limit the number of returned results (\"all\", \"none\", otherwise an integer value) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GeneApi.GeneMetadataByAccession(context.Background(), accessions).ReturnedContent(returnedContent).SortSchemaField(sortSchemaField).SortSchemaDirection(sortSchemaDirection).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GeneApi.GeneMetadataByAccession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GeneMetadataByAccession`: V1GeneMetadata
    fmt.Fprintf(os.Stdout, "Response from `GeneApi.GeneMetadataByAccession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessions** | [**[]string**](string.md) | RNA or Protein accessions. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGeneMetadataByAccessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **returnedContent** | [**V1GeneDatasetRequestContentType**](V1GeneDatasetRequestContentType.md) | Return either gene-ids, or entire gene metadata | [default to &quot;COMPLETE&quot;]
 **sortSchemaField** | [**V1GeneDatasetRequestSortField**](V1GeneDatasetRequestSortField.md) | Select a field to sort on. | [default to &quot;SORT_FIELD_UNSPECIFIED&quot;]
 **sortSchemaDirection** | [**V1SortDirection**](V1SortDirection.md) | Select a direction for the sort. | [default to &quot;SORT_DIRECTION_UNSPECIFIED&quot;]
 **limit** | **string** | Limit the number of returned results (\&quot;all\&quot;, \&quot;none\&quot;, otherwise an integer value) | 

### Return type

[**V1GeneMetadata**](V1GeneMetadata.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GeneMetadataById

> V1GeneMetadata GeneMetadataById(ctx, geneIds).ReturnedContent(returnedContent).SortSchemaField(sortSchemaField).SortSchemaDirection(sortSchemaDirection).Limit(limit).Execute()

Get gene metadata by GeneID



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
    geneIds := []int32{int32(123)} // []int32 | NCBI gene ids
    returnedContent := openapiclient.v1GeneDatasetRequestContentType("COMPLETE") // V1GeneDatasetRequestContentType | Return either gene-ids, or entire gene metadata (optional) (default to "COMPLETE")
    sortSchemaField := openapiclient.v1GeneDatasetRequestSortField("SORT_FIELD_UNSPECIFIED") // V1GeneDatasetRequestSortField | Select a field to sort on. (optional) (default to "SORT_FIELD_UNSPECIFIED")
    sortSchemaDirection := openapiclient.v1SortDirection("SORT_DIRECTION_UNSPECIFIED") // V1SortDirection | Select a direction for the sort. (optional) (default to "SORT_DIRECTION_UNSPECIFIED")
    limit := "limit_example" // string | Limit the number of returned results (\"all\", \"none\", otherwise an integer value) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GeneApi.GeneMetadataById(context.Background(), geneIds).ReturnedContent(returnedContent).SortSchemaField(sortSchemaField).SortSchemaDirection(sortSchemaDirection).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GeneApi.GeneMetadataById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GeneMetadataById`: V1GeneMetadata
    fmt.Fprintf(os.Stdout, "Response from `GeneApi.GeneMetadataById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**geneIds** | [**[]int32**](int32.md) | NCBI gene ids | 

### Other Parameters

Other parameters are passed through a pointer to a apiGeneMetadataByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **returnedContent** | [**V1GeneDatasetRequestContentType**](V1GeneDatasetRequestContentType.md) | Return either gene-ids, or entire gene metadata | [default to &quot;COMPLETE&quot;]
 **sortSchemaField** | [**V1GeneDatasetRequestSortField**](V1GeneDatasetRequestSortField.md) | Select a field to sort on. | [default to &quot;SORT_FIELD_UNSPECIFIED&quot;]
 **sortSchemaDirection** | [**V1SortDirection**](V1SortDirection.md) | Select a direction for the sort. | [default to &quot;SORT_DIRECTION_UNSPECIFIED&quot;]
 **limit** | **string** | Limit the number of returned results (\&quot;all\&quot;, \&quot;none\&quot;, otherwise an integer value) | 

### Return type

[**V1GeneMetadata**](V1GeneMetadata.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GeneMetadataByPost

> V1GeneMetadata GeneMetadataByPost(ctx).V1GeneDatasetRequest(v1GeneDatasetRequest).Execute()

Get gene metadata as JSON



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
    v1GeneDatasetRequest := *openapiclient.NewV1GeneDatasetRequest() // V1GeneDatasetRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GeneApi.GeneMetadataByPost(context.Background()).V1GeneDatasetRequest(v1GeneDatasetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GeneApi.GeneMetadataByPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GeneMetadataByPost`: V1GeneMetadata
    fmt.Fprintf(os.Stdout, "Response from `GeneApi.GeneMetadataByPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGeneMetadataByPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1GeneDatasetRequest** | [**V1GeneDatasetRequest**](V1GeneDatasetRequest.md) |  | 

### Return type

[**V1GeneMetadata**](V1GeneMetadata.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GeneMetadataByTaxAndSymbol

> V1GeneMetadata GeneMetadataByTaxAndSymbol(ctx, symbols, taxon).Accessions(accessions).ReturnedContent(returnedContent).SortSchemaField(sortSchemaField).SortSchemaDirection(sortSchemaDirection).Limit(limit).Execute()

Get gene metadata by gene symbol



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
    symbols := []string{"Inner_example"} // []string | Gene symbol
    taxon := "9606" // string | Taxon for provided gene symbol
    accessions := []string{"Inner_example"} // []string | RNA or Protein accessions. (optional)
    returnedContent := openapiclient.v1GeneDatasetRequestContentType("COMPLETE") // V1GeneDatasetRequestContentType | Return either gene-ids, or entire gene metadata (optional) (default to "COMPLETE")
    sortSchemaField := openapiclient.v1GeneDatasetRequestSortField("SORT_FIELD_UNSPECIFIED") // V1GeneDatasetRequestSortField | Select a field to sort on. (optional) (default to "SORT_FIELD_UNSPECIFIED")
    sortSchemaDirection := openapiclient.v1SortDirection("SORT_DIRECTION_UNSPECIFIED") // V1SortDirection | Select a direction for the sort. (optional) (default to "SORT_DIRECTION_UNSPECIFIED")
    limit := "limit_example" // string | Limit the number of returned results (\"all\", \"none\", otherwise an integer value) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GeneApi.GeneMetadataByTaxAndSymbol(context.Background(), symbols, taxon).Accessions(accessions).ReturnedContent(returnedContent).SortSchemaField(sortSchemaField).SortSchemaDirection(sortSchemaDirection).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GeneApi.GeneMetadataByTaxAndSymbol``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GeneMetadataByTaxAndSymbol`: V1GeneMetadata
    fmt.Fprintf(os.Stdout, "Response from `GeneApi.GeneMetadataByTaxAndSymbol`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**symbols** | [**[]string**](string.md) | Gene symbol | 
**taxon** | **string** | Taxon for provided gene symbol | 

### Other Parameters

Other parameters are passed through a pointer to a apiGeneMetadataByTaxAndSymbolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accessions** | **[]string** | RNA or Protein accessions. | 
 **returnedContent** | [**V1GeneDatasetRequestContentType**](V1GeneDatasetRequestContentType.md) | Return either gene-ids, or entire gene metadata | [default to &quot;COMPLETE&quot;]
 **sortSchemaField** | [**V1GeneDatasetRequestSortField**](V1GeneDatasetRequestSortField.md) | Select a field to sort on. | [default to &quot;SORT_FIELD_UNSPECIFIED&quot;]
 **sortSchemaDirection** | [**V1SortDirection**](V1SortDirection.md) | Select a direction for the sort. | [default to &quot;SORT_DIRECTION_UNSPECIFIED&quot;]
 **limit** | **string** | Limit the number of returned results (\&quot;all\&quot;, \&quot;none\&quot;, otherwise an integer value) | 

### Return type

[**V1GeneMetadata**](V1GeneMetadata.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GeneMetadataStreamByPost

> V1GeneMatch GeneMetadataStreamByPost(ctx).V1GeneDatasetRequest(v1GeneDatasetRequest).Execute()

Get gene metadata



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
    v1GeneDatasetRequest := *openapiclient.NewV1GeneDatasetRequest() // V1GeneDatasetRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GeneApi.GeneMetadataStreamByPost(context.Background()).V1GeneDatasetRequest(v1GeneDatasetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GeneApi.GeneMetadataStreamByPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GeneMetadataStreamByPost`: V1GeneMatch
    fmt.Fprintf(os.Stdout, "Response from `GeneApi.GeneMetadataStreamByPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGeneMetadataStreamByPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1GeneDatasetRequest** | [**V1GeneDatasetRequest**](V1GeneDatasetRequest.md) |  | 

### Return type

[**V1GeneMatch**](V1GeneMatch.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/x-ndjson, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GeneOrthologsById

> V1OrthologSet GeneOrthologsById(ctx, geneId).ReturnedContent(returnedContent).TaxonFilter(taxonFilter).Execute()

Get gene orthologs by gene ID



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
    geneId := int32(2778) // int32 | 
    returnedContent := openapiclient.v1OrthologRequestContentType("COMPLETE") // V1OrthologRequestContentType | Return either gene-ids, or entire gene metadata (optional) (default to "COMPLETE")
    taxonFilter := []string{"Inner_example"} // []string | Filter genes by taxa (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GeneApi.GeneOrthologsById(context.Background(), geneId).ReturnedContent(returnedContent).TaxonFilter(taxonFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GeneApi.GeneOrthologsById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GeneOrthologsById`: V1OrthologSet
    fmt.Fprintf(os.Stdout, "Response from `GeneApi.GeneOrthologsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**geneId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGeneOrthologsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **returnedContent** | [**V1OrthologRequestContentType**](V1OrthologRequestContentType.md) | Return either gene-ids, or entire gene metadata | [default to &quot;COMPLETE&quot;]
 **taxonFilter** | **[]string** | Filter genes by taxa | 

### Return type

[**V1OrthologSet**](V1OrthologSet.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GeneTaxNameQuery

> V1SciNameAndIds GeneTaxNameQuery(ctx, taxonQuery).TaxRankFilter(taxRankFilter).Execute()

Get a list of taxonomy names and IDs found in the gene dataset given a partial taxonomic name



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
    taxonQuery := "hum" // string | NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank
    taxRankFilter := openapiclient.v1OrganismQueryRequestTaxRankFilter("species") // V1OrganismQueryRequestTaxRankFilter | Set the scope of searched tax ranks (optional) (default to "species")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GeneApi.GeneTaxNameQuery(context.Background(), taxonQuery).TaxRankFilter(taxRankFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GeneApi.GeneTaxNameQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GeneTaxNameQuery`: V1SciNameAndIds
    fmt.Fprintf(os.Stdout, "Response from `GeneApi.GeneTaxNameQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxonQuery** | **string** | NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank | 

### Other Parameters

Other parameters are passed through a pointer to a apiGeneTaxNameQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **taxRankFilter** | [**V1OrganismQueryRequestTaxRankFilter**](V1OrganismQueryRequestTaxRankFilter.md) | Set the scope of searched tax ranks | [default to &quot;species&quot;]

### Return type

[**V1SciNameAndIds**](V1SciNameAndIds.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GeneTaxTree

> V1Organism GeneTaxTree(ctx, taxon).ChildrenOnly(childrenOnly).Execute()

Get a taxonomic subtree by taxonomic identifier



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
    taxon := "9606" // string | NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank
    childrenOnly := true // bool | Only report the children of the requested taxon and not their descendants (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GeneApi.GeneTaxTree(context.Background(), taxon).ChildrenOnly(childrenOnly).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GeneApi.GeneTaxTree``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GeneTaxTree`: V1Organism
    fmt.Fprintf(os.Stdout, "Response from `GeneApi.GeneTaxTree`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxon** | **string** | NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank | 

### Other Parameters

Other parameters are passed through a pointer to a apiGeneTaxTreeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **childrenOnly** | **bool** | Only report the children of the requested taxon and not their descendants | [default to false]

### Return type

[**V1Organism**](V1Organism.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

