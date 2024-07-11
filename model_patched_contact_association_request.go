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
type PatchedContactAssociationRequest struct {
	AssociatedObjectType string `json:"associated_object_type,omitempty"`
	AssociatedObjectId string `json:"associated_object_id,omitempty"`
	Contact *BulkWritableCircuitRequestTenant `json:"contact,omitempty"`
	Team *BulkWritableCircuitRequestTenant `json:"team,omitempty"`
	Role *BulkWritableCableRequestStatus `json:"role,omitempty"`
	Status *BulkWritableCableRequestStatus `json:"status,omitempty"`
	CustomFields map[string]Object `json:"custom_fields,omitempty"`
	Relationships map[string]BulkWritableCableRequestRelationships `json:"relationships,omitempty"`
}
