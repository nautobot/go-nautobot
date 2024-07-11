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
type BulkWritableVlanRequest struct {
	Id string `json:"id"`
	Location *BulkWritablePrefixRequestLocation `json:"location,omitempty"`
	Vid int32 `json:"vid"`
	Name string `json:"name"`
	Description string `json:"description,omitempty"`
	VlanGroup *BulkWritableCircuitRequestTenant `json:"vlan_group,omitempty"`
	Status *BulkWritableCableRequestStatus `json:"status"`
	Role *BulkWritableCircuitRequestTenant `json:"role,omitempty"`
	Tenant *BulkWritableCircuitRequestTenant `json:"tenant,omitempty"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	CustomFields map[string]Object `json:"custom_fields,omitempty"`
	Relationships map[string]BulkWritableCableRequestRelationships `json:"relationships,omitempty"`
}
