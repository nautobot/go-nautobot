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

// Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type TokenRequest struct {
	Key string `json:"key,omitempty"`
	Expires time.Time `json:"expires,omitempty"`
	// Permit create/update/delete operations using this key
	WriteEnabled bool `json:"write_enabled,omitempty"`
	Description string `json:"description,omitempty"`
}
