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
type WritablePowerOutletTemplateRequest struct {
	Name string `json:"name"`
	// Physical label
	Label string `json:"label,omitempty"`
	Description string `json:"description,omitempty"`
	Type_ *OneOfWritablePowerOutletTemplateRequestType_ `json:"type,omitempty"`
	// Phase (for three-phase feeds)
	FeedLeg *OneOfWritablePowerOutletTemplateRequestFeedLeg `json:"feed_leg,omitempty"`
	DeviceType *BulkWritableCableRequestStatus `json:"device_type"`
	PowerPortTemplate *BulkWritableCircuitRequestTenant `json:"power_port_template,omitempty"`
	CustomFields map[string]Object `json:"custom_fields,omitempty"`
	Relationships map[string]BulkWritableCableRequestRelationships `json:"relationships,omitempty"`
}
