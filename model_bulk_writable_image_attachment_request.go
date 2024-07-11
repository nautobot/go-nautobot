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
	"os"
)

// Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type BulkWritableImageAttachmentRequest struct {
	Id string `json:"id"`
	ContentType string `json:"content_type"`
	ObjectId string `json:"object_id"`
	Image **os.File `json:"image"`
	ImageHeight int32 `json:"image_height"`
	ImageWidth int32 `json:"image_width"`
	Name string `json:"name,omitempty"`
}
