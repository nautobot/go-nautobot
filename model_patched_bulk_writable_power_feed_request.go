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
type PatchedBulkWritablePowerFeedRequest struct {
	Id string `json:"id"`
	Type_ *AllOfPatchedBulkWritablePowerFeedRequestType_ `json:"type,omitempty"`
	Supply *AllOfPatchedBulkWritablePowerFeedRequestSupply `json:"supply,omitempty"`
	Phase *AllOfPatchedBulkWritablePowerFeedRequestPhase `json:"phase,omitempty"`
	Name string `json:"name,omitempty"`
	Voltage int32 `json:"voltage,omitempty"`
	Amperage int32 `json:"amperage,omitempty"`
	// Maximum permissible draw (percentage)
	MaxUtilization int32 `json:"max_utilization,omitempty"`
	Comments string `json:"comments,omitempty"`
	Cable *BulkWritableCircuitRequestTenant `json:"cable,omitempty"`
	PowerPanel *BulkWritableCableRequestStatus `json:"power_panel,omitempty"`
	Rack *BulkWritableCircuitRequestTenant `json:"rack,omitempty"`
	Status *BulkWritableCableRequestStatus `json:"status,omitempty"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	CustomFields map[string]Object `json:"custom_fields,omitempty"`
	Relationships map[string]BulkWritableCableRequestRelationships `json:"relationships,omitempty"`
}
