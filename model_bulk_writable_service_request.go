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
type BulkWritableServiceRequest struct {
	Id string `json:"id"`
	Protocol *ServiceProtocolChoices `json:"protocol,omitempty"`
	Ports []int32 `json:"ports"`
	Name string `json:"name"`
	Description string `json:"description,omitempty"`
	Device *BulkWritableServiceRequestDevice `json:"device,omitempty"`
	VirtualMachine *BulkWritableServiceRequestVirtualMachine `json:"virtual_machine,omitempty"`
	IpAddresses []IpAddresses `json:"ip_addresses,omitempty"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	CustomFields map[string]Object `json:"custom_fields,omitempty"`
	Relationships map[string]BulkWritableCableRequestRelationships `json:"relationships,omitempty"`
}
