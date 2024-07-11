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
type PatchedBulkWritableClusterRequest struct {
	Id string `json:"id"`
	Name string `json:"name,omitempty"`
	Comments string `json:"comments,omitempty"`
	ClusterType *BulkWritableCableRequestStatus `json:"cluster_type,omitempty"`
	ClusterGroup *BulkWritableCircuitRequestTenant `json:"cluster_group,omitempty"`
	Tenant *BulkWritableCircuitRequestTenant `json:"tenant,omitempty"`
	Location *BulkWritableCircuitRequestTenant `json:"location,omitempty"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	CustomFields map[string]Object `json:"custom_fields,omitempty"`
	Relationships map[string]BulkWritableCableRequestRelationships `json:"relationships,omitempty"`
}
