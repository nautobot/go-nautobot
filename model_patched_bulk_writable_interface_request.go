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
type PatchedBulkWritableInterfaceRequest struct {
	Id string `json:"id"`
	Type_ *InterfaceTypeChoices `json:"type,omitempty"`
	Mode *ModeEnum `json:"mode,omitempty"`
	MacAddress string `json:"mac_address,omitempty"`
	Name string `json:"name,omitempty"`
	// Physical label
	Label string `json:"label,omitempty"`
	Description string `json:"description,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
	Mtu int32 `json:"mtu,omitempty"`
	// This interface is used only for out-of-band management
	MgmtOnly bool `json:"mgmt_only,omitempty"`
	Device *BulkWritableCableRequestStatus `json:"device,omitempty"`
	Status *BulkWritableCableRequestStatus `json:"status,omitempty"`
	ParentInterface *BulkWritableInterfaceRequestParentInterface `json:"parent_interface,omitempty"`
	Bridge *BridgeInterface `json:"bridge,omitempty"`
	Lag *ParentLag `json:"lag,omitempty"`
	UntaggedVlan *BulkWritableCircuitRequestTenant `json:"untagged_vlan,omitempty"`
	Vrf *BulkWritableCircuitRequestTenant `json:"vrf,omitempty"`
	TaggedVlans []TaggedVlans `json:"tagged_vlans,omitempty"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	CustomFields map[string]Object `json:"custom_fields,omitempty"`
	Relationships map[string]BulkWritableCableRequestRelationships `json:"relationships,omitempty"`
}
