/*
 * NCBI Datasets API
 *
 * NCBI service to query and download biological sequence data across all domains of life from NCBI databases.
 *
 * API version: v1alpha
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package datasets
// Datasetsv1alpha1Annotation struct for Datasetsv1alpha1Annotation
type Datasetsv1alpha1Annotation struct {
	AssembliesInScope []V1alpha1AnnotatedAssemblies `json:"assemblies_in_scope,omitempty"`
	ReleaseDate string `json:"release_date,omitempty"`
	ReleaseName string `json:"release_name,omitempty"`
}
