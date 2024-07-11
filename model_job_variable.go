/*
 * API Documentation
 *
 * Source of truth and network automation platform
 *
 * API version: 2.2.5 (2.2)
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nautobot

// Serializer used for responses from the JobModelViewSet.variables() detail endpoint.
type JobVariable struct {
	Name string `json:"name"`
	Type_ string `json:"type"`
	Label string `json:"label"`
	HelpText string `json:"help_text"`
	Default_ map[string]Object `json:"default"`
	Required bool `json:"required"`
	MinLength int32 `json:"min_length"`
	MaxLength int32 `json:"max_length"`
	MinValue int32 `json:"min_value"`
	MaxValue int32 `json:"max_value"`
	Choices map[string]Object `json:"choices"`
	Model string `json:"model"`
}
