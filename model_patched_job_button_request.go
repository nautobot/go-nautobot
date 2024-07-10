/*
 * API Documentation
 *
 * Source of truth and network automation platform
 *
 * API version: 2.2.5 (2.2)
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package go-nautobot

// Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type PatchedJobButtonRequest struct {
	ContentTypes []string `json:"content_types,omitempty"`
	Name string `json:"name,omitempty"`
	// Jinja2 template code for button text. Reference the object as <code>{{ obj }}</code> such as <code>{{ obj.platform.name }}</code>. Buttons which render as empty text will not be displayed.
	Text string `json:"text,omitempty"`
	Weight int32 `json:"weight,omitempty"`
	// Buttons with the same group will appear as a dropdown menu. Group dropdown buttons will inherit the button class from the button with the lowest weight in the group.
	GroupName string `json:"group_name,omitempty"`
	ButtonClass *ButtonClassEnum `json:"button_class,omitempty"`
	// Enable confirmation pop-up box. <span class='text-danger'>WARNING: unselecting this option will allow the Job to run (and commit changes) with a single click!</span>
	Confirmation bool `json:"confirmation,omitempty"`
	Job *BulkWritableJobButtonRequestJob `json:"job,omitempty"`
}
