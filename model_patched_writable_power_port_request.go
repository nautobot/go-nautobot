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
type PatchedWritablePowerPortRequest struct {
	Name string `json:"name,omitempty"`
	// Physical label
	Label string `json:"label,omitempty"`
	Description string `json:"description,omitempty"`
	// Physical port type
	Type_ *OneOfPatchedWritablePowerPortRequestType_ `json:"type,omitempty"`
	// Maximum power draw (watts)
	MaximumDraw int32 `json:"maximum_draw,omitempty"`
	// Allocated power draw (watts)
	AllocatedDraw int32 `json:"allocated_draw,omitempty"`
	Device *BulkWritableCableRequestStatus `json:"device,omitempty"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	CustomFields map[string]Object `json:"custom_fields,omitempty"`
	Relationships map[string]BulkWritableCableRequestRelationships `json:"relationships,omitempty"`
}
