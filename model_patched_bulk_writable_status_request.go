/*
 * API Documentation
 *
 * Source of truth and network automation platform
 *
 * API version: 2.2.5 (2.2)
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nautobot

// Serializer for `Status` objects.
type PatchedBulkWritableStatusRequest struct {
	Id string `json:"id"`
	ContentTypes []string `json:"content_types,omitempty"`
	Name string `json:"name,omitempty"`
	// RGB color in hexadecimal (e.g. 00ff00)
	Color string `json:"color,omitempty"`
	Description string `json:"description,omitempty"`
	CustomFields map[string]Object `json:"custom_fields,omitempty"`
	Relationships map[string]BulkWritableCableRequestRelationships `json:"relationships,omitempty"`
}
