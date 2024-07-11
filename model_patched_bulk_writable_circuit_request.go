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
type PatchedBulkWritableCircuitRequest struct {
	Id string `json:"id"`
	Cid string `json:"cid,omitempty"`
	InstallDate string `json:"install_date,omitempty"`
	CommitRate int32 `json:"commit_rate,omitempty"`
	Description string `json:"description,omitempty"`
	Comments string `json:"comments,omitempty"`
	Status *BulkWritableCableRequestStatus `json:"status,omitempty"`
	Provider *BulkWritableCableRequestStatus `json:"provider,omitempty"`
	CircuitType *BulkWritableCableRequestStatus `json:"circuit_type,omitempty"`
	Tenant *BulkWritableCircuitRequestTenant `json:"tenant,omitempty"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	CustomFields map[string]Object `json:"custom_fields,omitempty"`
	Relationships map[string]BulkWritableCableRequestRelationships `json:"relationships,omitempty"`
}
