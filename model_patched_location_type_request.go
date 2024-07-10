/*
 * API Documentation
 *
 * Source of truth and network automation platform
 *
 * API version: 2.2.5 (2.2)
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package go-nautobot

// Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type PatchedLocationTypeRequest struct {
	ContentTypes []string `json:"content_types,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	// Allow Locations of this type to be parents/children of other Locations of this same type
	Nestable bool `json:"nestable,omitempty"`
	Parent *BulkWritableCircuitRequestTenant `json:"parent,omitempty"`
	CustomFields map[string]Object `json:"custom_fields,omitempty"`
	Relationships map[string]BulkWritableCableRequestRelationships `json:"relationships,omitempty"`
}
