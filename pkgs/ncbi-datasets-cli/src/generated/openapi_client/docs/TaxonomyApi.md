# \TaxonomyApi

All URIs are relative to *https://api.ncbi.nlm.nih.gov/datasets/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TaxNameQuery**](TaxonomyApi.md#TaxNameQuery) | **Get** /taxonomy/taxon_suggest/{taxon_query} | Get a list of taxonomy names and IDs given a partial taxonomic name
[**TaxonomyFilteredSubtree**](TaxonomyApi.md#TaxonomyFilteredSubtree) | **Get** /taxonomy/taxon/{taxons}/filtered_subtree | Use taxonomic identifiers to get a filtered taxonomic subtree
[**TaxonomyFilteredSubtreePost**](TaxonomyApi.md#TaxonomyFilteredSubtreePost) | **Post** /taxonomy/filtered_subtree | Use taxonomic identifiers to get a filtered taxonomic subtree by post
[**TaxonomyMetadata**](TaxonomyApi.md#TaxonomyMetadata) | **Get** /taxonomy/taxon/{taxons} | Use taxonomic identifiers to get taxonomic metadata
[**TaxonomyMetadataPost**](TaxonomyApi.md#TaxonomyMetadataPost) | **Post** /taxonomy | Use taxonomic identifiers to get taxonomic metadata by post



## TaxNameQuery

> V1SciNameAndIds TaxNameQuery(ctx, taxonQuery).TaxRankFilter(taxRankFilter).Execute()

Get a list of taxonomy names and IDs given a partial taxonomic name



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
    resp, r, err := api_client.TaxonomyApi.TaxNameQuery(context.Background(), taxonQuery).TaxRankFilter(taxRankFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxonomyApi.TaxNameQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TaxNameQuery`: V1SciNameAndIds
    fmt.Fprintf(os.Stdout, "Response from `TaxonomyApi.TaxNameQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxonQuery** | **string** | NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank | 

### Other Parameters

Other parameters are passed through a pointer to a apiTaxNameQueryRequest struct via the builder pattern


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


## TaxonomyFilteredSubtree

> V1TaxonomyFilteredSubtreeResponse TaxonomyFilteredSubtree(ctx, taxons).SpecifiedLimit(specifiedLimit).RankLimits(rankLimits).Execute()

Use taxonomic identifiers to get a filtered taxonomic subtree



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
    taxons := []string{"Inner_example"} // []string | 
    specifiedLimit := true // bool | Limit to specified species (optional)
    rankLimits := []openapiclient.V1OrganismRankType{openapiclient.v1OrganismRankType("NO_RANK")} // []V1OrganismRankType | Limit to the provided ranks.  If empty, accept any rank. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TaxonomyApi.TaxonomyFilteredSubtree(context.Background(), taxons).SpecifiedLimit(specifiedLimit).RankLimits(rankLimits).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxonomyApi.TaxonomyFilteredSubtree``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TaxonomyFilteredSubtree`: V1TaxonomyFilteredSubtreeResponse
    fmt.Fprintf(os.Stdout, "Response from `TaxonomyApi.TaxonomyFilteredSubtree`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxons** | [**[]string**](string.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTaxonomyFilteredSubtreeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **specifiedLimit** | **bool** | Limit to specified species | 
 **rankLimits** | [**[]V1OrganismRankType**](V1OrganismRankType.md) | Limit to the provided ranks.  If empty, accept any rank. | 

### Return type

[**V1TaxonomyFilteredSubtreeResponse**](V1TaxonomyFilteredSubtreeResponse.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaxonomyFilteredSubtreePost

> V1TaxonomyFilteredSubtreeResponse TaxonomyFilteredSubtreePost(ctx).V1TaxonomyFilteredSubtreeRequest(v1TaxonomyFilteredSubtreeRequest).Execute()

Use taxonomic identifiers to get a filtered taxonomic subtree by post



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
    v1TaxonomyFilteredSubtreeRequest := *openapiclient.NewV1TaxonomyFilteredSubtreeRequest() // V1TaxonomyFilteredSubtreeRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TaxonomyApi.TaxonomyFilteredSubtreePost(context.Background()).V1TaxonomyFilteredSubtreeRequest(v1TaxonomyFilteredSubtreeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxonomyApi.TaxonomyFilteredSubtreePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TaxonomyFilteredSubtreePost`: V1TaxonomyFilteredSubtreeResponse
    fmt.Fprintf(os.Stdout, "Response from `TaxonomyApi.TaxonomyFilteredSubtreePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTaxonomyFilteredSubtreePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1TaxonomyFilteredSubtreeRequest** | [**V1TaxonomyFilteredSubtreeRequest**](V1TaxonomyFilteredSubtreeRequest.md) |  | 

### Return type

[**V1TaxonomyFilteredSubtreeResponse**](V1TaxonomyFilteredSubtreeResponse.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaxonomyMetadata

> V1TaxonomyMetadataResponse TaxonomyMetadata(ctx, taxons).ReturnedContent(returnedContent).Execute()

Use taxonomic identifiers to get taxonomic metadata



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
    taxons := []string{"Inner_example"} // []string | 
    returnedContent := openapiclient.v1TaxonomyMetadataRequestContentType("COMPLETE") // V1TaxonomyMetadataRequestContentType | Return either tax-ids alone, or entire taxononmy-metadata records (optional) (default to "COMPLETE")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TaxonomyApi.TaxonomyMetadata(context.Background(), taxons).ReturnedContent(returnedContent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxonomyApi.TaxonomyMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TaxonomyMetadata`: V1TaxonomyMetadataResponse
    fmt.Fprintf(os.Stdout, "Response from `TaxonomyApi.TaxonomyMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxons** | [**[]string**](string.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTaxonomyMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **returnedContent** | [**V1TaxonomyMetadataRequestContentType**](V1TaxonomyMetadataRequestContentType.md) | Return either tax-ids alone, or entire taxononmy-metadata records | [default to &quot;COMPLETE&quot;]

### Return type

[**V1TaxonomyMetadataResponse**](V1TaxonomyMetadataResponse.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaxonomyMetadataPost

> V1TaxonomyMetadataResponse TaxonomyMetadataPost(ctx).V1TaxonomyMetadataRequest(v1TaxonomyMetadataRequest).Execute()

Use taxonomic identifiers to get taxonomic metadata by post



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
    v1TaxonomyMetadataRequest := *openapiclient.NewV1TaxonomyMetadataRequest() // V1TaxonomyMetadataRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TaxonomyApi.TaxonomyMetadataPost(context.Background()).V1TaxonomyMetadataRequest(v1TaxonomyMetadataRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxonomyApi.TaxonomyMetadataPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TaxonomyMetadataPost`: V1TaxonomyMetadataResponse
    fmt.Fprintf(os.Stdout, "Response from `TaxonomyApi.TaxonomyMetadataPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTaxonomyMetadataPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1TaxonomyMetadataRequest** | [**V1TaxonomyMetadataRequest**](V1TaxonomyMetadataRequest.md) |  | 

### Return type

[**V1TaxonomyMetadataResponse**](V1TaxonomyMetadataResponse.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

