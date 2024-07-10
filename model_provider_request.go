/*
 * API Documentation
 *
 * Source of truth and network automation platform
 *
 * API version: 2.2.5 (2.2)
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package go-nautobot

// Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type ProviderRequest struct {
	Name string `json:"name"`
	// 32-bit autonomous system number
	Asn int64 `json:"asn,omitempty"`
	Account string `json:"account,omitempty"`
	PortalUrl string `json:"portal_url,omitempty"`
	NocContact string `json:"noc_contact,omitempty"`
	AdminContact string `json:"admin_contact,omitempty"`
	Comments string `json:"comments,omitempty"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	CustomFields map[string]Object `json:"custom_fields,omitempty"`
	Relationships map[string]BulkWritableCableRequestRelationships `json:"relationships,omitempty"`
}
