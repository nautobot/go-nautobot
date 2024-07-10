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
type CircuitTerminationRequest struct {
	TermSide *AllOfCircuitTerminationRequestTermSide `json:"term_side"`
	PortSpeed int32 `json:"port_speed,omitempty"`
	// Upstream speed, if different from port speed
	UpstreamSpeed int32 `json:"upstream_speed,omitempty"`
	XconnectId string `json:"xconnect_id,omitempty"`
	PpInfo string `json:"pp_info,omitempty"`
	Description string `json:"description,omitempty"`
	Circuit *BulkWritableCableRequestStatus `json:"circuit"`
	Location *BulkWritableCircuitRequestTenant `json:"location,omitempty"`
	ProviderNetwork *BulkWritableCircuitRequestTenant `json:"provider_network,omitempty"`
	CustomFields map[string]Object `json:"custom_fields,omitempty"`
	Relationships map[string]BulkWritableCableRequestRelationships `json:"relationships,omitempty"`
}
