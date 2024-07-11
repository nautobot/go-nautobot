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

// Serializer for `Role` objects.
type Role struct {
	Id string `json:"id"`
	ObjectType string `json:"object_type"`
	// Human friendly display value
	Display string `json:"display"`
	Url string `json:"url"`
	NaturalSlug string `json:"natural_slug"`
	ContentTypes []string `json:"content_types"`
	Name string `json:"name"`
	// RGB color in hexadecimal (e.g. 00ff00)
	Color string `json:"color,omitempty"`
	Description string `json:"description,omitempty"`
	Weight int32 `json:"weight,omitempty"`
	Created time.Time `json:"created"`
	LastUpdated time.Time `json:"last_updated"`
	NotesUrl string `json:"notes_url"`
	CustomFields map[string]Object `json:"custom_fields,omitempty"`
}
