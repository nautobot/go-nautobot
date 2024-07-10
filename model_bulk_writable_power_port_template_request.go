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
type BulkWritablePowerPortTemplateRequest struct {
	Id string `json:"id"`
	Type_ *PowerPortTypeChoices `json:"type,omitempty"`
	Name string `json:"name"`
	// Physical label
	Label string `json:"label,omitempty"`
	Description string `json:"description,omitempty"`
	// Maximum power draw (watts)
	MaximumDraw int32 `json:"maximum_draw,omitempty"`
	// Allocated power draw (watts)
	AllocatedDraw int32 `json:"allocated_draw,omitempty"`
	DeviceType *BulkWritableCableRequestStatus `json:"device_type"`
	CustomFields map[string]Object `json:"custom_fields,omitempty"`
	Relationships map[string]BulkWritableCableRequestRelationships `json:"relationships,omitempty"`
}
