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

// Git repositories defined as a data source.
type GitRepository struct {
	Id string `json:"id"`
	ObjectType string `json:"object_type"`
	// Human friendly display value
	Display string `json:"display"`
	Url string `json:"url"`
	NaturalSlug string `json:"natural_slug"`
	ProvidedContents []OneOfGitRepositoryProvidedContentsItems `json:"provided_contents,omitempty"`
	Name string `json:"name"`
	// Internal field name. Please use underscores rather than dashes in this key.
	Slug string `json:"slug,omitempty"`
	// Only HTTP and HTTPS URLs are presently supported
	RemoteUrl string `json:"remote_url"`
	Branch string `json:"branch,omitempty"`
	// Commit hash of the most recent fetch from the selected branch. Used for syncing between workers.
	CurrentHead string `json:"current_head,omitempty"`
	SecretsGroup *BulkWritableCircuitRequestTenant `json:"secrets_group,omitempty"`
	Created time.Time `json:"created"`
	LastUpdated time.Time `json:"last_updated"`
	NotesUrl string `json:"notes_url"`
	CustomFields map[string]Object `json:"custom_fields,omitempty"`
}
