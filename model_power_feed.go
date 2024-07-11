/*
 * API Documentation
 *
 * Source of truth and network automation platform
 *
 * API version: 2.2.5 (2.2)
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nautobot
import (
	"time"
)

// Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type PowerFeed struct {
	Id string `json:"id"`
	ObjectType string `json:"object_type"`
	// Human friendly display value
	Display string `json:"display"`
	Url string `json:"url"`
	NaturalSlug string `json:"natural_slug"`
	CablePeerType string `json:"cable_peer_type"`
	CablePeer *AllOfPowerFeedCablePeer `json:"cable_peer"`
	ConnectedEndpointType string `json:"connected_endpoint_type"`
	ConnectedEndpoint *AllOfPowerFeedConnectedEndpoint `json:"connected_endpoint"`
	ConnectedEndpointReachable bool `json:"connected_endpoint_reachable"`
	Type_ *PowerFeedType `json:"type,omitempty"`
	Supply *PowerFeedSupply `json:"supply,omitempty"`
	Phase *PowerFeedPhase `json:"phase,omitempty"`
	Name string `json:"name"`
	Voltage int32 `json:"voltage,omitempty"`
	Amperage int32 `json:"amperage,omitempty"`
	// Maximum permissible draw (percentage)
	MaxUtilization int32 `json:"max_utilization,omitempty"`
	AvailablePower int32 `json:"available_power"`
	Comments string `json:"comments,omitempty"`
	Cable *BulkWritableCircuitRequestTenant `json:"cable,omitempty"`
	PowerPanel *BulkWritableCableRequestStatus `json:"power_panel"`
	Rack *BulkWritableCircuitRequestTenant `json:"rack,omitempty"`
	Status *BulkWritableCableRequestStatus `json:"status"`
	Created time.Time `json:"created"`
	LastUpdated time.Time `json:"last_updated"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	NotesUrl string `json:"notes_url"`
	CustomFields map[string]Object `json:"custom_fields,omitempty"`
}
