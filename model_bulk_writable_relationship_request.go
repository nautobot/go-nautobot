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
type BulkWritableRelationshipRequest struct {
	Id string `json:"id"`
	SourceType string `json:"source_type"`
	DestinationType string `json:"destination_type"`
	// Label of the relationship as displayed to users
	Label string `json:"label"`
	// Internal relationship key. Please use underscores rather than dashes in this key.
	Key string `json:"key,omitempty"`
	Description string `json:"description,omitempty"`
	// Cardinality of this relationship
	Type_ *AllOfBulkWritableRelationshipRequestType_ `json:"type,omitempty"`
	// Objects on the specified side MUST implement this relationship. Not permitted for symmetric relationships.
	RequiredOn *OneOfBulkWritableRelationshipRequestRequiredOn `json:"required_on,omitempty"`
	// Label for related destination objects, as displayed on the source object.
	SourceLabel string `json:"source_label,omitempty"`
	// Hide this relationship on the source object.
	SourceHidden bool `json:"source_hidden,omitempty"`
	// Filterset filter matching the applicable source objects of the selected type
	SourceFilter map[string]Object `json:"source_filter,omitempty"`
	// Label for related source objects, as displayed on the destination object.
	DestinationLabel string `json:"destination_label,omitempty"`
	// Hide this relationship on the destination object.
	DestinationHidden bool `json:"destination_hidden,omitempty"`
	// Filterset filter matching the applicable destination objects of the selected type
	DestinationFilter map[string]Object `json:"destination_filter,omitempty"`
	// Hide this field from the object's primary information tab. It will appear in the \"Advanced\" tab instead.
	AdvancedUi bool `json:"advanced_ui,omitempty"`
}
