/*
 * API Documentation
 *
 * Source of truth and network automation platform
 *
 * API version: 2.2.5 (2.2)
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nautobot

// Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type VrfPrefixAssignment struct {
	Id string `json:"id"`
	ObjectType string `json:"object_type"`
	// Human friendly display value
	Display string `json:"display"`
	Url string `json:"url"`
	NaturalSlug string `json:"natural_slug"`
	Vrf *BulkWritableCableRequestStatus `json:"vrf"`
	Prefix *BulkWritableCableRequestStatus `json:"prefix"`
}
