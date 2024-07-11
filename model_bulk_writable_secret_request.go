/*
 * API Documentation
 *
 * Source of truth and network automation platform
 *
 * API version: 2.2.5 (2.2)
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nautobot

// Serializer for `Secret` objects.
type BulkWritableSecretRequest struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description,omitempty"`
	Provider string `json:"provider"`
	Parameters map[string]Object `json:"parameters,omitempty"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	CustomFields map[string]Object `json:"custom_fields,omitempty"`
	Relationships map[string]BulkWritableCableRequestRelationships `json:"relationships,omitempty"`
}
