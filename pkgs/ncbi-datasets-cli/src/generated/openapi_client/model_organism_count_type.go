/*
 * NCBI Datasets API
 *
 * NCBI service to query and download biological sequence data across all domains of life from NCBI databases.
 *
 * API version: v1alpha
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package datasets
// OrganismCountType the model 'OrganismCountType'
type OrganismCountType string

// List of OrganismCountType
const (
	ORGANISMCOUNTTYPE_UNSPECIFIED OrganismCountType = "COUNT_TYPE_UNSPECIFIED"
	ORGANISMCOUNTTYPE_ASSEMBLY OrganismCountType = "COUNT_TYPE_ASSEMBLY"
	ORGANISMCOUNTTYPE_GENE OrganismCountType = "COUNT_TYPE_GENE"
)
