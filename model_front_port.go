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
type FrontPort struct {
	Id string `json:"id"`
	ObjectType string `json:"object_type"`
	// Human friendly display value
	Display string `json:"display"`
	Url string `json:"url"`
	NaturalSlug string `json:"natural_slug"`
	CablePeerType string `json:"cable_peer_type"`
	CablePeer *AllOfFrontPortCablePeer `json:"cable_peer"`
	Type_ *FrontPortType `json:"type"`
	Name string `json:"name"`
	// Physical label
	Label string `json:"label,omitempty"`
	Description string `json:"description,omitempty"`
	RearPortPosition int32 `json:"rear_port_position,omitempty"`
	Device *BulkWritableCableRequestStatus `json:"device"`
	Cable *CircuitCircuitTerminationA `json:"cable"`
	RearPort *BulkWritableCableRequestStatus `json:"rear_port"`
	Created time.Time `json:"created"`
	LastUpdated time.Time `json:"last_updated"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	NotesUrl string `json:"notes_url"`
	CustomFields map[string]Object `json:"custom_fields,omitempty"`
}
