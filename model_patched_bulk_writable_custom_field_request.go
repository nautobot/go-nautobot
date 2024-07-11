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
type PatchedBulkWritableCustomFieldRequest struct {
	Id string `json:"id"`
	ContentTypes []string `json:"content_types,omitempty"`
	Type_ *CustomFieldTypeChoices `json:"type,omitempty"`
	FilterLogic *FilterLogicEnum `json:"filter_logic,omitempty"`
	Label string `json:"label,omitempty"`
	// Human-readable grouping that this custom field belongs to.
	Grouping string `json:"grouping,omitempty"`
	// Internal field name. Please use underscores rather than dashes in this key.
	Key string `json:"key,omitempty"`
	// A helpful description for this field.
	Description string `json:"description,omitempty"`
	// If true, this field is required when creating new objects or editing an existing object.
	Required bool `json:"required,omitempty"`
	// Default value for the field (must be a JSON value). Encapsulate strings with double quotes (e.g. \"Foo\").
	Default_ map[string]Object `json:"default,omitempty"`
	// Fields with higher weights appear lower in a form.
	Weight int32 `json:"weight,omitempty"`
	// Minimum allowed value (for numeric fields) or length (for text fields).
	ValidationMinimum int64 `json:"validation_minimum,omitempty"`
	// Maximum allowed value (for numeric fields) or length (for text fields).
	ValidationMaximum int64 `json:"validation_maximum,omitempty"`
	// Regular expression to enforce on text field values. Use ^ and $ to force matching of entire string. For example, <code>^[A-Z]{3}$</code> will limit values to exactly three uppercase letters. Regular expression on select and multi-select will be applied at <code>Custom Field Choices</code> definition.
	ValidationRegex string `json:"validation_regex,omitempty"`
	// Hide this field from the object's primary information tab. It will appear in the \"Advanced\" tab instead.
	AdvancedUi bool `json:"advanced_ui,omitempty"`
}
