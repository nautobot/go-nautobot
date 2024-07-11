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

// Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type UserRequest struct {
	Password string `json:"password,omitempty"`
	LastLogin time.Time `json:"last_login,omitempty"`
	// Designates that this user has all permissions without explicitly assigning them.
	IsSuperuser bool `json:"is_superuser,omitempty"`
	// Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only.
	Username string `json:"username"`
	FirstName string `json:"first_name,omitempty"`
	LastName string `json:"last_name,omitempty"`
	Email string `json:"email,omitempty"`
	// Designates whether the user can log into this admin site.
	IsStaff bool `json:"is_staff,omitempty"`
	// Designates whether this user should be treated as active. Unselect this instead of deleting accounts.
	IsActive bool `json:"is_active,omitempty"`
	DateJoined time.Time `json:"date_joined,omitempty"`
	ConfigData map[string]Object `json:"config_data,omitempty"`
	// The groups this user belongs to. A user will get all permissions granted to each of their groups.
	Groups []BulkWritableCableRequestStatus `json:"groups,omitempty"`
}
