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
type BulkWritableIpAddressRequest struct {
	Id string `json:"id"`
	Address string `json:"address"`
	Namespace *BulkWritableIpAddressRequestNamespace `json:"namespace,omitempty"`
	Type_ *IpAddressTypeChoices `json:"type,omitempty"`
	// Hostname or FQDN (not case-sensitive)
	DnsName string `json:"dns_name,omitempty"`
	Description string `json:"description,omitempty"`
	Status *BulkWritableCableRequestStatus `json:"status"`
	Role *BulkWritableCircuitRequestTenant `json:"role,omitempty"`
	Parent *BulkWritableIpAddressRequestParent `json:"parent,omitempty"`
	Tenant *BulkWritableCircuitRequestTenant `json:"tenant,omitempty"`
	NatInside *NatInside `json:"nat_inside,omitempty"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	CustomFields map[string]Object `json:"custom_fields,omitempty"`
	Relationships map[string]BulkWritableCableRequestRelationships `json:"relationships,omitempty"`
}
