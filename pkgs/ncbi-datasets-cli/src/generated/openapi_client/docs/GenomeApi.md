# \GenomeApi

All URIs are relative to *https://api.ncbi.nlm.nih.gov/datasets/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssemblyDescriptorsByAccessions**](GenomeApi.md#AssemblyDescriptorsByAccessions) | **Get** /genome/accession/{accessions} | Get genome metadata by accession
[**AssemblyDescriptorsByBioproject**](GenomeApi.md#AssemblyDescriptorsByBioproject) | **Get** /genome/bioproject/{accessions} | Get genome metadata by bioproject accession
[**AssemblyDescriptorsByTaxon**](GenomeApi.md#AssemblyDescriptorsByTaxon) | **Get** /genome/taxon/{taxon} | Get genome metadata by taxonomic identifier
[**CheckAssemblyAvailability**](GenomeApi.md#CheckAssemblyAvailability) | **Get** /genome/accession/{accessions}/check | Check the validity of genome accessions
[**CheckAssemblyAvailabilityPost**](GenomeApi.md#CheckAssemblyAvailabilityPost) | **Post** /genome/check | Check the validity of many genome accessions in a single request
[**DownloadAssemblyPackage**](GenomeApi.md#DownloadAssemblyPackage) | **Get** /genome/accession/{accessions}/download | Get a genome dataset by accession
[**DownloadAssemblyPackagePost**](GenomeApi.md#DownloadAssemblyPackagePost) | **Post** /genome/download | Get a genome dataset by post
[**GenomeDownloadSummary**](GenomeApi.md#GenomeDownloadSummary) | **Get** /genome/accession/{accessions}/download_summary | Preview genome dataset download
[**GenomeDownloadSummaryByPost**](GenomeApi.md#GenomeDownloadSummaryByPost) | **Post** /genome/download_summary | Preview genome dataset download by POST
[**GenomeMetadataByPost**](GenomeApi.md#GenomeMetadataByPost) | **Post** /genome | Get genome metadata by variety of identifiers
[**GenomeTaxNameQuery**](GenomeApi.md#GenomeTaxNameQuery) | **Get** /genome/taxon_suggest/{taxon_query} | Get a list of taxonomy names and IDs found in the assembly dataset given a partial taxonomic name
[**GenomeTaxTree**](GenomeApi.md#GenomeTaxTree) | **Get** /genome/taxon/{taxon}/tree | Get a taxonomic subtree by taxonomic identifier



## AssemblyDescriptorsByAccessions

> V1AssemblyMetadata AssemblyDescriptorsByAccessions(ctx, accessions).FiltersReferenceOnly(filtersReferenceOnly).FiltersAssemblySource(filtersAssemblySource).FiltersHasAnnotation(filtersHasAnnotation).FiltersExcludeAtypical(filtersExcludeAtypical).FiltersAssemblyLevel(filtersAssemblyLevel).FiltersFirstReleaseDate(filtersFirstReleaseDate).FiltersLastReleaseDate(filtersLastReleaseDate).FiltersSearchText(filtersSearchText).PageSize(pageSize).PageToken(pageToken).Execute()

Get genome metadata by accession



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
    accessions := []string{"Inner_example"} // []string | 
    filtersReferenceOnly := true // bool | If true, only return reference and representative (GCF_ and GCA_) genome assemblies. (optional) (default to false)
    filtersAssemblySource := openapiclient.v1AssemblyDatasetDescriptorsFilterAssemblySource("all") // V1AssemblyDatasetDescriptorsFilterAssemblySource | Return only RefSeq (GCF_) or GenBank (GCA_) genome assemblies (optional) (default to "all")
    filtersHasAnnotation := true // bool | Return only annotated genome assemblies (optional) (default to false)
    filtersExcludeAtypical := false // bool | If true, exclude atypical genomes, i.e. genomes that have assembly issues or are otherwise atypical (optional) (default to false)
    filtersAssemblyLevel := []openapiclient.V1AssemblyDatasetDescriptorsFilterAssemblyLevel{openapiclient.v1AssemblyDatasetDescriptorsFilterAssemblyLevel("chromosome")} // []V1AssemblyDatasetDescriptorsFilterAssemblyLevel | Only return genome assemblies that have one of the specified assembly levels. By default, do not filter. (optional)
    filtersFirstReleaseDate := time.Now() // time.Time | Only return genome assemblies that were released on or after the specified date By default, do not filter. (optional)
    filtersLastReleaseDate := time.Now() // time.Time | Only return genome assemblies that were released on or before to the specified date By default, do not filter. (optional)
    filtersSearchText := []string{"Inner_example"} // []string | Only return results whose fields contain the specified search terms in their taxon, infraspecific, assembly name or submitter fields By default, do not filter (optional)
    pageSize := int32(56) // int32 | The maximum number of genome assemblies to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, `page_token` can be used to retrieve the remaining results. (optional) (default to 20)
    pageToken := "pageToken_example" // string | A page token is returned from an `AssemblyMetadataRequest` call with more than `page_size` results. Use this token, along with the previous `AssemblyMetadataRequest` parameters, to retrieve the next page of results. When `page_token` is empty, all results have been retrieved. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GenomeApi.AssemblyDescriptorsByAccessions(context.Background(), accessions).FiltersReferenceOnly(filtersReferenceOnly).FiltersAssemblySource(filtersAssemblySource).FiltersHasAnnotation(filtersHasAnnotation).FiltersExcludeAtypical(filtersExcludeAtypical).FiltersAssemblyLevel(filtersAssemblyLevel).FiltersFirstReleaseDate(filtersFirstReleaseDate).FiltersLastReleaseDate(filtersLastReleaseDate).FiltersSearchText(filtersSearchText).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenomeApi.AssemblyDescriptorsByAccessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssemblyDescriptorsByAccessions`: V1AssemblyMetadata
    fmt.Fprintf(os.Stdout, "Response from `GenomeApi.AssemblyDescriptorsByAccessions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessions** | [**[]string**](string.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssemblyDescriptorsByAccessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filtersReferenceOnly** | **bool** | If true, only return reference and representative (GCF_ and GCA_) genome assemblies. | [default to false]
 **filtersAssemblySource** | [**V1AssemblyDatasetDescriptorsFilterAssemblySource**](V1AssemblyDatasetDescriptorsFilterAssemblySource.md) | Return only RefSeq (GCF_) or GenBank (GCA_) genome assemblies | [default to &quot;all&quot;]
 **filtersHasAnnotation** | **bool** | Return only annotated genome assemblies | [default to false]
 **filtersExcludeAtypical** | **bool** | If true, exclude atypical genomes, i.e. genomes that have assembly issues or are otherwise atypical | [default to false]
 **filtersAssemblyLevel** | [**[]V1AssemblyDatasetDescriptorsFilterAssemblyLevel**](V1AssemblyDatasetDescriptorsFilterAssemblyLevel.md) | Only return genome assemblies that have one of the specified assembly levels. By default, do not filter. | 
 **filtersFirstReleaseDate** | **time.Time** | Only return genome assemblies that were released on or after the specified date By default, do not filter. | 
 **filtersLastReleaseDate** | **time.Time** | Only return genome assemblies that were released on or before to the specified date By default, do not filter. | 
 **filtersSearchText** | **[]string** | Only return results whose fields contain the specified search terms in their taxon, infraspecific, assembly name or submitter fields By default, do not filter | 
 **pageSize** | **int32** | The maximum number of genome assemblies to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, &#x60;page_token&#x60; can be used to retrieve the remaining results. | [default to 20]
 **pageToken** | **string** | A page token is returned from an &#x60;AssemblyMetadataRequest&#x60; call with more than &#x60;page_size&#x60; results. Use this token, along with the previous &#x60;AssemblyMetadataRequest&#x60; parameters, to retrieve the next page of results. When &#x60;page_token&#x60; is empty, all results have been retrieved. | 

### Return type

[**V1AssemblyMetadata**](V1AssemblyMetadata.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssemblyDescriptorsByBioproject

> V1AssemblyMetadata AssemblyDescriptorsByBioproject(ctx, accessions).FiltersReferenceOnly(filtersReferenceOnly).FiltersAssemblySource(filtersAssemblySource).FiltersHasAnnotation(filtersHasAnnotation).FiltersExcludePairedReports(filtersExcludePairedReports).FiltersExcludeAtypical(filtersExcludeAtypical).FiltersAssemblyVersion(filtersAssemblyVersion).FiltersAssemblyLevel(filtersAssemblyLevel).FiltersFirstReleaseDate(filtersFirstReleaseDate).FiltersLastReleaseDate(filtersLastReleaseDate).FiltersSearchText(filtersSearchText).ReturnedContent(returnedContent).PageSize(pageSize).PageToken(pageToken).Execute()

Get genome metadata by bioproject accession



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
    accessions := []string{"Inner_example"} // []string | 
    filtersReferenceOnly := true // bool | If true, only return reference and representative (GCF_ and GCA_) genome assemblies. (optional) (default to false)
    filtersAssemblySource := openapiclient.v1AssemblyDatasetDescriptorsFilterAssemblySource("all") // V1AssemblyDatasetDescriptorsFilterAssemblySource | Return only RefSeq (GCF_) or GenBank (GCA_) genome assemblies (optional) (default to "all")
    filtersHasAnnotation := true // bool | Return only annotated genome assemblies (optional) (default to false)
    filtersExcludePairedReports := false // bool | For paired (GCA/GCF) records, only return the primary record (optional) (default to false)
    filtersExcludeAtypical := false // bool | If true, exclude atypical genomes, i.e. genomes that have assembly issues or are otherwise atypical (optional) (default to false)
    filtersAssemblyVersion := openapiclient.v1AssemblyDatasetDescriptorsFilterAssemblyVersion("current") // V1AssemblyDatasetDescriptorsFilterAssemblyVersion | Return all assemblies, including replaced and suppressed, or only current assemblies (optional) (default to "current")
    filtersAssemblyLevel := []openapiclient.V1AssemblyDatasetDescriptorsFilterAssemblyLevel{openapiclient.v1AssemblyDatasetDescriptorsFilterAssemblyLevel("chromosome")} // []V1AssemblyDatasetDescriptorsFilterAssemblyLevel | Only return genome assemblies that have one of the specified assembly levels. By default, do not filter. (optional)
    filtersFirstReleaseDate := time.Now() // time.Time | Only return genome assemblies that were released on or after the specified date By default, do not filter. (optional)
    filtersLastReleaseDate := time.Now() // time.Time | Only return genome assemblies that were released on or before to the specified date By default, do not filter. (optional)
    filtersSearchText := []string{"Inner_example"} // []string | Only return results whose fields contain the specified search terms in their taxon, infraspecific, assembly name or submitter fields By default, do not filter (optional)
    returnedContent := openapiclient.v1AssemblyMetadataRequestContentType("COMPLETE") // V1AssemblyMetadataRequestContentType | Return either assembly accessions, or entire assembly-metadata records (optional) (default to "COMPLETE")
    pageSize := int32(56) // int32 | The maximum number of genome assemblies to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, `page_token` can be used to retrieve the remaining results. (optional) (default to 20)
    pageToken := "pageToken_example" // string | A page token is returned from an `AssemblyMetadataRequest` call with more than `page_size` results. Use this token, along with the previous `AssemblyMetadataRequest` parameters, to retrieve the next page of results. When `page_token` is empty, all results have been retrieved. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GenomeApi.AssemblyDescriptorsByBioproject(context.Background(), accessions).FiltersReferenceOnly(filtersReferenceOnly).FiltersAssemblySource(filtersAssemblySource).FiltersHasAnnotation(filtersHasAnnotation).FiltersExcludePairedReports(filtersExcludePairedReports).FiltersExcludeAtypical(filtersExcludeAtypical).FiltersAssemblyVersion(filtersAssemblyVersion).FiltersAssemblyLevel(filtersAssemblyLevel).FiltersFirstReleaseDate(filtersFirstReleaseDate).FiltersLastReleaseDate(filtersLastReleaseDate).FiltersSearchText(filtersSearchText).ReturnedContent(returnedContent).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenomeApi.AssemblyDescriptorsByBioproject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssemblyDescriptorsByBioproject`: V1AssemblyMetadata
    fmt.Fprintf(os.Stdout, "Response from `GenomeApi.AssemblyDescriptorsByBioproject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessions** | [**[]string**](string.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssemblyDescriptorsByBioprojectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filtersReferenceOnly** | **bool** | If true, only return reference and representative (GCF_ and GCA_) genome assemblies. | [default to false]
 **filtersAssemblySource** | [**V1AssemblyDatasetDescriptorsFilterAssemblySource**](V1AssemblyDatasetDescriptorsFilterAssemblySource.md) | Return only RefSeq (GCF_) or GenBank (GCA_) genome assemblies | [default to &quot;all&quot;]
 **filtersHasAnnotation** | **bool** | Return only annotated genome assemblies | [default to false]
 **filtersExcludePairedReports** | **bool** | For paired (GCA/GCF) records, only return the primary record | [default to false]
 **filtersExcludeAtypical** | **bool** | If true, exclude atypical genomes, i.e. genomes that have assembly issues or are otherwise atypical | [default to false]
 **filtersAssemblyVersion** | [**V1AssemblyDatasetDescriptorsFilterAssemblyVersion**](V1AssemblyDatasetDescriptorsFilterAssemblyVersion.md) | Return all assemblies, including replaced and suppressed, or only current assemblies | [default to &quot;current&quot;]
 **filtersAssemblyLevel** | [**[]V1AssemblyDatasetDescriptorsFilterAssemblyLevel**](V1AssemblyDatasetDescriptorsFilterAssemblyLevel.md) | Only return genome assemblies that have one of the specified assembly levels. By default, do not filter. | 
 **filtersFirstReleaseDate** | **time.Time** | Only return genome assemblies that were released on or after the specified date By default, do not filter. | 
 **filtersLastReleaseDate** | **time.Time** | Only return genome assemblies that were released on or before to the specified date By default, do not filter. | 
 **filtersSearchText** | **[]string** | Only return results whose fields contain the specified search terms in their taxon, infraspecific, assembly name or submitter fields By default, do not filter | 
 **returnedContent** | [**V1AssemblyMetadataRequestContentType**](V1AssemblyMetadataRequestContentType.md) | Return either assembly accessions, or entire assembly-metadata records | [default to &quot;COMPLETE&quot;]
 **pageSize** | **int32** | The maximum number of genome assemblies to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, &#x60;page_token&#x60; can be used to retrieve the remaining results. | [default to 20]
 **pageToken** | **string** | A page token is returned from an &#x60;AssemblyMetadataRequest&#x60; call with more than &#x60;page_size&#x60; results. Use this token, along with the previous &#x60;AssemblyMetadataRequest&#x60; parameters, to retrieve the next page of results. When &#x60;page_token&#x60; is empty, all results have been retrieved. | 

### Return type

[**V1AssemblyMetadata**](V1AssemblyMetadata.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssemblyDescriptorsByTaxon

> V1AssemblyMetadata AssemblyDescriptorsByTaxon(ctx, taxon).FiltersReferenceOnly(filtersReferenceOnly).FiltersAssemblySource(filtersAssemblySource).FiltersHasAnnotation(filtersHasAnnotation).FiltersExcludePairedReports(filtersExcludePairedReports).FiltersExcludeAtypical(filtersExcludeAtypical).FiltersAssemblyVersion(filtersAssemblyVersion).FiltersAssemblyLevel(filtersAssemblyLevel).FiltersFirstReleaseDate(filtersFirstReleaseDate).FiltersLastReleaseDate(filtersLastReleaseDate).FiltersSearchText(filtersSearchText).TaxExactMatch(taxExactMatch).ReturnedContent(returnedContent).PageSize(pageSize).PageToken(pageToken).Execute()

Get genome metadata by taxonomic identifier



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
    taxon := "9606" // string | NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank
    filtersReferenceOnly := true // bool | If true, only return reference and representative (GCF_ and GCA_) genome assemblies. (optional) (default to false)
    filtersAssemblySource := openapiclient.v1AssemblyDatasetDescriptorsFilterAssemblySource("all") // V1AssemblyDatasetDescriptorsFilterAssemblySource | Return only RefSeq (GCF_) or GenBank (GCA_) genome assemblies (optional) (default to "all")
    filtersHasAnnotation := true // bool | Return only annotated genome assemblies (optional) (default to false)
    filtersExcludePairedReports := false // bool | For paired (GCA/GCF) records, only return the primary record (optional) (default to false)
    filtersExcludeAtypical := false // bool | If true, exclude atypical genomes, i.e. genomes that have assembly issues or are otherwise atypical (optional) (default to false)
    filtersAssemblyVersion := openapiclient.v1AssemblyDatasetDescriptorsFilterAssemblyVersion("current") // V1AssemblyDatasetDescriptorsFilterAssemblyVersion | Return all assemblies, including replaced and suppressed, or only current assemblies (optional) (default to "current")
    filtersAssemblyLevel := []openapiclient.V1AssemblyDatasetDescriptorsFilterAssemblyLevel{openapiclient.v1AssemblyDatasetDescriptorsFilterAssemblyLevel("chromosome")} // []V1AssemblyDatasetDescriptorsFilterAssemblyLevel | Only return genome assemblies that have one of the specified assembly levels. By default, do not filter. (optional)
    filtersFirstReleaseDate := time.Now() // time.Time | Only return genome assemblies that were released on or after the specified date By default, do not filter. (optional)
    filtersLastReleaseDate := time.Now() // time.Time | Only return genome assemblies that were released on or before to the specified date By default, do not filter. (optional)
    filtersSearchText := []string{"Inner_example"} // []string | Only return results whose fields contain the specified search terms in their taxon, infraspecific, assembly name or submitter fields By default, do not filter (optional)
    taxExactMatch := true // bool | If true, only return assemblies with the given NCBI Taxonomy ID, or name. Otherwise, assemblies from taxonomy subtree are included, too. Ignored for assembly_accession request. (optional) (default to false)
    returnedContent := openapiclient.v1AssemblyMetadataRequestContentType("COMPLETE") // V1AssemblyMetadataRequestContentType | Return either assembly accessions, or entire assembly-metadata records (optional) (default to "COMPLETE")
    pageSize := int32(56) // int32 | The maximum number of genome assemblies to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, `page_token` can be used to retrieve the remaining results. (optional) (default to 20)
    pageToken := "pageToken_example" // string | A page token is returned from an `AssemblyMetadataRequest` call with more than `page_size` results. Use this token, along with the previous `AssemblyMetadataRequest` parameters, to retrieve the next page of results. When `page_token` is empty, all results have been retrieved. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GenomeApi.AssemblyDescriptorsByTaxon(context.Background(), taxon).FiltersReferenceOnly(filtersReferenceOnly).FiltersAssemblySource(filtersAssemblySource).FiltersHasAnnotation(filtersHasAnnotation).FiltersExcludePairedReports(filtersExcludePairedReports).FiltersExcludeAtypical(filtersExcludeAtypical).FiltersAssemblyVersion(filtersAssemblyVersion).FiltersAssemblyLevel(filtersAssemblyLevel).FiltersFirstReleaseDate(filtersFirstReleaseDate).FiltersLastReleaseDate(filtersLastReleaseDate).FiltersSearchText(filtersSearchText).TaxExactMatch(taxExactMatch).ReturnedContent(returnedContent).PageSize(pageSize).PageToken(pageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenomeApi.AssemblyDescriptorsByTaxon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssemblyDescriptorsByTaxon`: V1AssemblyMetadata
    fmt.Fprintf(os.Stdout, "Response from `GenomeApi.AssemblyDescriptorsByTaxon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxon** | **string** | NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssemblyDescriptorsByTaxonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filtersReferenceOnly** | **bool** | If true, only return reference and representative (GCF_ and GCA_) genome assemblies. | [default to false]
 **filtersAssemblySource** | [**V1AssemblyDatasetDescriptorsFilterAssemblySource**](V1AssemblyDatasetDescriptorsFilterAssemblySource.md) | Return only RefSeq (GCF_) or GenBank (GCA_) genome assemblies | [default to &quot;all&quot;]
 **filtersHasAnnotation** | **bool** | Return only annotated genome assemblies | [default to false]
 **filtersExcludePairedReports** | **bool** | For paired (GCA/GCF) records, only return the primary record | [default to false]
 **filtersExcludeAtypical** | **bool** | If true, exclude atypical genomes, i.e. genomes that have assembly issues or are otherwise atypical | [default to false]
 **filtersAssemblyVersion** | [**V1AssemblyDatasetDescriptorsFilterAssemblyVersion**](V1AssemblyDatasetDescriptorsFilterAssemblyVersion.md) | Return all assemblies, including replaced and suppressed, or only current assemblies | [default to &quot;current&quot;]
 **filtersAssemblyLevel** | [**[]V1AssemblyDatasetDescriptorsFilterAssemblyLevel**](V1AssemblyDatasetDescriptorsFilterAssemblyLevel.md) | Only return genome assemblies that have one of the specified assembly levels. By default, do not filter. | 
 **filtersFirstReleaseDate** | **time.Time** | Only return genome assemblies that were released on or after the specified date By default, do not filter. | 
 **filtersLastReleaseDate** | **time.Time** | Only return genome assemblies that were released on or before to the specified date By default, do not filter. | 
 **filtersSearchText** | **[]string** | Only return results whose fields contain the specified search terms in their taxon, infraspecific, assembly name or submitter fields By default, do not filter | 
 **taxExactMatch** | **bool** | If true, only return assemblies with the given NCBI Taxonomy ID, or name. Otherwise, assemblies from taxonomy subtree are included, too. Ignored for assembly_accession request. | [default to false]
 **returnedContent** | [**V1AssemblyMetadataRequestContentType**](V1AssemblyMetadataRequestContentType.md) | Return either assembly accessions, or entire assembly-metadata records | [default to &quot;COMPLETE&quot;]
 **pageSize** | **int32** | The maximum number of genome assemblies to return. Default is 20 and maximum is 1000. If the number of results exceeds the page size, &#x60;page_token&#x60; can be used to retrieve the remaining results. | [default to 20]
 **pageToken** | **string** | A page token is returned from an &#x60;AssemblyMetadataRequest&#x60; call with more than &#x60;page_size&#x60; results. Use this token, along with the previous &#x60;AssemblyMetadataRequest&#x60; parameters, to retrieve the next page of results. When &#x60;page_token&#x60; is empty, all results have been retrieved. | 

### Return type

[**V1AssemblyMetadata**](V1AssemblyMetadata.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckAssemblyAvailability

> V1AssemblyDatasetAvailability CheckAssemblyAvailability(ctx, accessions).Execute()

Check the validity of genome accessions



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
    accessions := []string{"Inner_example"} // []string | NCBI genome assembly accessions

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GenomeApi.CheckAssemblyAvailability(context.Background(), accessions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenomeApi.CheckAssemblyAvailability``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckAssemblyAvailability`: V1AssemblyDatasetAvailability
    fmt.Fprintf(os.Stdout, "Response from `GenomeApi.CheckAssemblyAvailability`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessions** | [**[]string**](string.md) | NCBI genome assembly accessions | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckAssemblyAvailabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V1AssemblyDatasetAvailability**](V1AssemblyDatasetAvailability.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckAssemblyAvailabilityPost

> V1AssemblyDatasetAvailability CheckAssemblyAvailabilityPost(ctx).V1AssemblyDatasetRequest(v1AssemblyDatasetRequest).Execute()

Check the validity of many genome accessions in a single request



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
    v1AssemblyDatasetRequest := *openapiclient.NewV1AssemblyDatasetRequest() // V1AssemblyDatasetRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GenomeApi.CheckAssemblyAvailabilityPost(context.Background()).V1AssemblyDatasetRequest(v1AssemblyDatasetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenomeApi.CheckAssemblyAvailabilityPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckAssemblyAvailabilityPost`: V1AssemblyDatasetAvailability
    fmt.Fprintf(os.Stdout, "Response from `GenomeApi.CheckAssemblyAvailabilityPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckAssemblyAvailabilityPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1AssemblyDatasetRequest** | [**V1AssemblyDatasetRequest**](V1AssemblyDatasetRequest.md) |  | 

### Return type

[**V1AssemblyDatasetAvailability**](V1AssemblyDatasetAvailability.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadAssemblyPackage

> *os.File DownloadAssemblyPackage(ctx, accessions).Chromosomes(chromosomes).ExcludeSequence(excludeSequence).IncludeAnnotationType(includeAnnotationType).Hydrated(hydrated).Filename(filename).Execute()

Get a genome dataset by accession



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
    accessions := []string{"Inner_example"} // []string | NCBI genome assembly accessions
    chromosomes := []string{"Inner_example"} // []string | The default setting is all chromosome. Specify individual chromosome by string (1,2,MT or chr1,chr2.chrMT). Unplaced sequences are treated like their own chromosome ('Un'). The filter only applies to fasta sequence. (optional)
    excludeSequence := true // bool | Set to true to omit the genomic sequence. (optional) (default to false)
    includeAnnotationType := []openapiclient.V1AnnotationForAssemblyType{openapiclient.v1AnnotationForAssemblyType("DEFAULT")} // []V1AnnotationForAssemblyType | Select additional types of annotation to include in the data package.  If unset, no annotation is provided. (optional)
    hydrated := openapiclient.v1AssemblyDatasetRequestResolution("FULLY_HYDRATED") // V1AssemblyDatasetRequestResolution | Set to DATA_REPORT_ONLY, to only retrieve data-reports. (optional) (default to "FULLY_HYDRATED")
    filename := "filename_example" // string | Output file name. (optional) (default to "ncbi_dataset.zip")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GenomeApi.DownloadAssemblyPackage(context.Background(), accessions).Chromosomes(chromosomes).ExcludeSequence(excludeSequence).IncludeAnnotationType(includeAnnotationType).Hydrated(hydrated).Filename(filename).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenomeApi.DownloadAssemblyPackage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadAssemblyPackage`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `GenomeApi.DownloadAssemblyPackage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessions** | [**[]string**](string.md) | NCBI genome assembly accessions | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadAssemblyPackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chromosomes** | **[]string** | The default setting is all chromosome. Specify individual chromosome by string (1,2,MT or chr1,chr2.chrMT). Unplaced sequences are treated like their own chromosome (&#39;Un&#39;). The filter only applies to fasta sequence. | 
 **excludeSequence** | **bool** | Set to true to omit the genomic sequence. | [default to false]
 **includeAnnotationType** | [**[]V1AnnotationForAssemblyType**](V1AnnotationForAssemblyType.md) | Select additional types of annotation to include in the data package.  If unset, no annotation is provided. | 
 **hydrated** | [**V1AssemblyDatasetRequestResolution**](V1AssemblyDatasetRequestResolution.md) | Set to DATA_REPORT_ONLY, to only retrieve data-reports. | [default to &quot;FULLY_HYDRATED&quot;]
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


## DownloadAssemblyPackagePost

> *os.File DownloadAssemblyPackagePost(ctx).V1AssemblyDatasetRequest(v1AssemblyDatasetRequest).Filename(filename).Execute()

Get a genome dataset by post



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
    v1AssemblyDatasetRequest := *openapiclient.NewV1AssemblyDatasetRequest() // V1AssemblyDatasetRequest | 
    filename := "filename_example" // string | Output file name. (optional) (default to "ncbi_dataset.zip")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GenomeApi.DownloadAssemblyPackagePost(context.Background()).V1AssemblyDatasetRequest(v1AssemblyDatasetRequest).Filename(filename).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenomeApi.DownloadAssemblyPackagePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadAssemblyPackagePost`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `GenomeApi.DownloadAssemblyPackagePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDownloadAssemblyPackagePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1AssemblyDatasetRequest** | [**V1AssemblyDatasetRequest**](V1AssemblyDatasetRequest.md) |  | 
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


## GenomeDownloadSummary

> V1DownloadSummary GenomeDownloadSummary(ctx, accessions).Chromosomes(chromosomes).ExcludeSequence(excludeSequence).IncludeAnnotationType(includeAnnotationType).Execute()

Preview genome dataset download



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
    accessions := []string{"Inner_example"} // []string | NCBI genome assembly accessions
    chromosomes := []string{"Inner_example"} // []string | The default setting is all chromosome. Specify individual chromosome by string (1,2,MT or chr1,chr2.chrMT). Unplaced sequences are treated like their own chromosome ('Un'). The filter only applies to fasta sequence. (optional)
    excludeSequence := true // bool | Set to true to omit the genomic sequence. (optional) (default to false)
    includeAnnotationType := []openapiclient.V1AnnotationForAssemblyType{openapiclient.v1AnnotationForAssemblyType("DEFAULT")} // []V1AnnotationForAssemblyType | Select additional types of annotation to include in the data package.  If unset, no annotation is provided. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GenomeApi.GenomeDownloadSummary(context.Background(), accessions).Chromosomes(chromosomes).ExcludeSequence(excludeSequence).IncludeAnnotationType(includeAnnotationType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenomeApi.GenomeDownloadSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenomeDownloadSummary`: V1DownloadSummary
    fmt.Fprintf(os.Stdout, "Response from `GenomeApi.GenomeDownloadSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessions** | [**[]string**](string.md) | NCBI genome assembly accessions | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenomeDownloadSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **chromosomes** | **[]string** | The default setting is all chromosome. Specify individual chromosome by string (1,2,MT or chr1,chr2.chrMT). Unplaced sequences are treated like their own chromosome (&#39;Un&#39;). The filter only applies to fasta sequence. | 
 **excludeSequence** | **bool** | Set to true to omit the genomic sequence. | [default to false]
 **includeAnnotationType** | [**[]V1AnnotationForAssemblyType**](V1AnnotationForAssemblyType.md) | Select additional types of annotation to include in the data package.  If unset, no annotation is provided. | 

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


## GenomeDownloadSummaryByPost

> V1DownloadSummary GenomeDownloadSummaryByPost(ctx).V1AssemblyDatasetRequest(v1AssemblyDatasetRequest).Execute()

Preview genome dataset download by POST



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
    v1AssemblyDatasetRequest := *openapiclient.NewV1AssemblyDatasetRequest() // V1AssemblyDatasetRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GenomeApi.GenomeDownloadSummaryByPost(context.Background()).V1AssemblyDatasetRequest(v1AssemblyDatasetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenomeApi.GenomeDownloadSummaryByPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenomeDownloadSummaryByPost`: V1DownloadSummary
    fmt.Fprintf(os.Stdout, "Response from `GenomeApi.GenomeDownloadSummaryByPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenomeDownloadSummaryByPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1AssemblyDatasetRequest** | [**V1AssemblyDatasetRequest**](V1AssemblyDatasetRequest.md) |  | 

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


## GenomeMetadataByPost

> V1AssemblyMetadata GenomeMetadataByPost(ctx).V1AssemblyMetadataRequest(v1AssemblyMetadataRequest).Execute()

Get genome metadata by variety of identifiers



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
    v1AssemblyMetadataRequest := *openapiclient.NewV1AssemblyMetadataRequest() // V1AssemblyMetadataRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.GenomeApi.GenomeMetadataByPost(context.Background()).V1AssemblyMetadataRequest(v1AssemblyMetadataRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenomeApi.GenomeMetadataByPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenomeMetadataByPost`: V1AssemblyMetadata
    fmt.Fprintf(os.Stdout, "Response from `GenomeApi.GenomeMetadataByPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenomeMetadataByPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v1AssemblyMetadataRequest** | [**V1AssemblyMetadataRequest**](V1AssemblyMetadataRequest.md) |  | 

### Return type

[**V1AssemblyMetadata**](V1AssemblyMetadata.md)

### Authorization

[ApiKeyAuthHeader](../README.md#ApiKeyAuthHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenomeTaxNameQuery

> V1SciNameAndIds GenomeTaxNameQuery(ctx, taxonQuery).TaxRankFilter(taxRankFilter).Execute()

Get a list of taxonomy names and IDs found in the assembly dataset given a partial taxonomic name



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
    resp, r, err := api_client.GenomeApi.GenomeTaxNameQuery(context.Background(), taxonQuery).TaxRankFilter(taxRankFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenomeApi.GenomeTaxNameQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenomeTaxNameQuery`: V1SciNameAndIds
    fmt.Fprintf(os.Stdout, "Response from `GenomeApi.GenomeTaxNameQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxonQuery** | **string** | NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenomeTaxNameQueryRequest struct via the builder pattern


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


## GenomeTaxTree

> V1Organism GenomeTaxTree(ctx, taxon).ChildrenOnly(childrenOnly).Execute()

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
    resp, r, err := api_client.GenomeApi.GenomeTaxTree(context.Background(), taxon).ChildrenOnly(childrenOnly).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GenomeApi.GenomeTaxTree``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenomeTaxTree`: V1Organism
    fmt.Fprintf(os.Stdout, "Response from `GenomeApi.GenomeTaxTree`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taxon** | **string** | NCBI Taxonomy ID or name (common or scientific) at any taxonomic rank | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenomeTaxTreeRequest struct via the builder pattern


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

