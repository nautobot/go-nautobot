/*
 * API Documentation
 *
 * Source of truth and network automation platform
 *
 * API version: 2.2.5 (2.2)
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package go-nautobot
import (
	"time"
)

// Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type DeviceFamily struct {
	Id string `json:"id"`
	ObjectType string `json:"object_type"`
	// Human friendly display value
	Display string `json:"display"`
	Url string `json:"url"`
	NaturalSlug string `json:"natural_slug"`
	DeviceTypeCount int32 `json:"device_type_count"`
	Name string `json:"name"`
	Description string `json:"description,omitempty"`
	Created time.Time `json:"created"`
	LastUpdated time.Time `json:"last_updated"`
	NotesUrl string `json:"notes_url"`
	CustomFields map[string]Object `json:"custom_fields,omitempty"`
}
