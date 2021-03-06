/*
 * NCBI Datasets API
 *
 * NCBI service to query and download biological sequence data across all domains of life from NCBI databases.
 *
 * API version: v1alpha
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package datasets
// V1alpha1GeneMetadata struct for V1alpha1GeneMetadata
type V1alpha1GeneMetadata struct {
	Genes []V1alpha1GeneMatch `json:"genes,omitempty"`
	Messages []V1alpha1Message `json:"messages,omitempty"`
	// The total count of available datasets (ignoring the cutoff parameter).
	TotalCount int64 `json:"total_count,omitempty"`
}
