/*
 * NCBI Datasets API
 *
 * NCBI service to query and download biological sequence data across all domains of life from NCBI databases.
 *
 * API version: v1alpha
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package datasets
// GenomicRegionGenomicRegionType the model 'GenomicRegionGenomicRegionType'
type GenomicRegionGenomicRegionType string

// List of GenomicRegionGenomicRegionType
const (
	GENOMICREGIONGENOMICREGIONTYPE_UNKNOWN GenomicRegionGenomicRegionType = "UNKNOWN"
	GENOMICREGIONGENOMICREGIONTYPE_REFSEQ_GENE GenomicRegionGenomicRegionType = "REFSEQ_GENE"
	GENOMICREGIONGENOMICREGIONTYPE_PSEUDOGENE GenomicRegionGenomicRegionType = "PSEUDOGENE"
	GENOMICREGIONGENOMICREGIONTYPE_BIOLOGICAL_REGION GenomicRegionGenomicRegionType = "BIOLOGICAL_REGION"
	GENOMICREGIONGENOMICREGIONTYPE_OTHER GenomicRegionGenomicRegionType = "OTHER"
)
