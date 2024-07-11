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
type PatchedPlatformRequest struct {
	Name string `json:"name,omitempty"`
	// The normalized network driver to use when interacting with devices, e.g. cisco_ios, arista_eos, etc. Library-specific driver names will be derived from this setting as appropriate
	NetworkDriver string `json:"network_driver,omitempty"`
	// The name of the NAPALM driver to use when Nautobot internals interact with devices
	NapalmDriver string `json:"napalm_driver,omitempty"`
	// Additional arguments to pass when initiating the NAPALM driver (JSON format)
	NapalmArgs map[string]Object `json:"napalm_args,omitempty"`
	Description string `json:"description,omitempty"`
	Manufacturer *BulkWritablePlatformRequestManufacturer `json:"manufacturer,omitempty"`
	CustomFields map[string]Object `json:"custom_fields,omitempty"`
	Relationships map[string]BulkWritableCableRequestRelationships `json:"relationships,omitempty"`
}
