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
type SoftwareVersion struct {
	Id string `json:"id"`
	ObjectType string `json:"object_type"`
	// Human friendly display value
	Display string `json:"display"`
	Url string `json:"url"`
	NaturalSlug string `json:"natural_slug"`
	Version string `json:"version"`
	// Optional alternative label for this version
	Alias string `json:"alias,omitempty"`
	ReleaseDate string `json:"release_date,omitempty"`
	EndOfSupportDate string `json:"end_of_support_date,omitempty"`
	DocumentationUrl string `json:"documentation_url,omitempty"`
	// Is a Long Term Support version
	LongTermSupport bool `json:"long_term_support,omitempty"`
	// Is a Pre-Release version
	PreRelease bool `json:"pre_release,omitempty"`
	Platform *BulkWritableCableRequestStatus `json:"platform"`
	Status *BulkWritableCableRequestStatus `json:"status"`
	Created time.Time `json:"created"`
	LastUpdated time.Time `json:"last_updated"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	NotesUrl string `json:"notes_url"`
	CustomFields map[string]Object `json:"custom_fields,omitempty"`
}
