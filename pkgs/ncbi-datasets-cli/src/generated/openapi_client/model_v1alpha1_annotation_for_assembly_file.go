/*
 * NCBI Datasets API
 *
 * NCBI service to query and download biological sequence data across all domains of life from NCBI databases.
 *
 * API version: v1alpha
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package datasets
// V1alpha1AnnotationForAssemblyFile struct for V1alpha1AnnotationForAssemblyFile
type V1alpha1AnnotationForAssemblyFile struct {
	EstimatedSize string `json:"estimated_size,omitempty"`
	Type V1alpha1AnnotationForAssemblyType `json:"type,omitempty"`
}
