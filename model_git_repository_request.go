/*
 * API Documentation
 *
 * Source of truth and network automation platform
 *
 * API version: 2.2.5 (2.2)
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package go-nautobot

// Git repositories defined as a data source.
type GitRepositoryRequest struct {
	ProvidedContents []OneOfGitRepositoryRequestProvidedContentsItems `json:"provided_contents,omitempty"`
	Name string `json:"name"`
	// Internal field name. Please use underscores rather than dashes in this key.
	Slug string `json:"slug,omitempty"`
	// Only HTTP and HTTPS URLs are presently supported
	RemoteUrl string `json:"remote_url"`
	Branch string `json:"branch,omitempty"`
	// Commit hash of the most recent fetch from the selected branch. Used for syncing between workers.
	CurrentHead string `json:"current_head,omitempty"`
	SecretsGroup *BulkWritableCircuitRequestTenant `json:"secrets_group,omitempty"`
	CustomFields map[string]Object `json:"custom_fields,omitempty"`
	Relationships map[string]BulkWritableCableRequestRelationships `json:"relationships,omitempty"`
}
