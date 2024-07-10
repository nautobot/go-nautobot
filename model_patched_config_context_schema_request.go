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
type PatchedConfigContextSchemaRequest struct {
	OwnerContentType string `json:"owner_content_type,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	// A JSON Schema document which is used to validate a config context object.
	DataSchema map[string]Object `json:"data_schema,omitempty"`
	OwnerObjectId string `json:"owner_object_id,omitempty"`
	CustomFields map[string]Object `json:"custom_fields,omitempty"`
	Relationships map[string]BulkWritableCableRequestRelationships `json:"relationships,omitempty"`
}
