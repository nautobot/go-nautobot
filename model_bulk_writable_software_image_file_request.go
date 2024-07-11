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
type BulkWritableSoftwareImageFileRequest struct {
	Id string `json:"id"`
	ImageFileName string `json:"image_file_name"`
	ImageFileChecksum string `json:"image_file_checksum,omitempty"`
	// Hashing algorithm for image file checksum
	HashingAlgorithm *OneOfBulkWritableSoftwareImageFileRequestHashingAlgorithm `json:"hashing_algorithm,omitempty"`
	// Image file size in bytes
	ImageFileSize int64 `json:"image_file_size,omitempty"`
	DownloadUrl string `json:"download_url,omitempty"`
	// Is the default image for this software version
	DefaultImage bool `json:"default_image,omitempty"`
	SoftwareVersion *BulkWritableCableRequestStatus `json:"software_version"`
	Status *BulkWritableCableRequestStatus `json:"status"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	CustomFields map[string]Object `json:"custom_fields,omitempty"`
	Relationships map[string]BulkWritableCableRequestRelationships `json:"relationships,omitempty"`
}
