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
type PatchedBulkWritableJobHookRequest struct {
	Id string `json:"id"`
	ContentTypes []string `json:"content_types,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
	Name string `json:"name,omitempty"`
	// Call this job hook when a matching object is created.
	TypeCreate bool `json:"type_create,omitempty"`
	// Call this job hook when a matching object is deleted.
	TypeDelete bool `json:"type_delete,omitempty"`
	// Call this job hook when a matching object is updated.
	TypeUpdate bool `json:"type_update,omitempty"`
	Job *BulkWritableJobHookRequestJob `json:"job,omitempty"`
	CustomFields map[string]Object `json:"custom_fields,omitempty"`
	Relationships map[string]BulkWritableCableRequestRelationships `json:"relationships,omitempty"`
}
