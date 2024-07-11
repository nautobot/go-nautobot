/*
 * API Documentation
 *
 * Source of truth and network automation platform
 *
 * API version: 2.2.5 (2.2)
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nautobot

// Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type PatchedVrfRequest struct {
	Name string `json:"name,omitempty"`
	// Unique route distinguisher (as defined in RFC 4364)
	Rd string `json:"rd,omitempty"`
	Description string `json:"description,omitempty"`
	Namespace *BulkWritableCableRequestStatus `json:"namespace,omitempty"`
	Tenant *BulkWritableCircuitRequestTenant `json:"tenant,omitempty"`
	ImportTargets []BulkWritableCableRequestStatus `json:"import_targets,omitempty"`
	ExportTargets []BulkWritableCableRequestStatus `json:"export_targets,omitempty"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	CustomFields map[string]Object `json:"custom_fields,omitempty"`
	Relationships map[string]BulkWritableCableRequestRelationships `json:"relationships,omitempty"`
}
