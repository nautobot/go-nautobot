/*
 * API Documentation
 *
 * Source of truth and network automation platform
 *
 * API version: 2.2.5 (2.2)
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package go-nautobot
import (
	"time"
)

// Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type IpAddress struct {
	Id string `json:"id"`
	ObjectType string `json:"object_type"`
	// Human friendly display value
	Display string `json:"display"`
	Url string `json:"url"`
	NaturalSlug string `json:"natural_slug"`
	Address string `json:"address"`
	// IPv4 or IPv6 host address
	Host string `json:"host"`
	// Length of the network mask, in bits.
	MaskLength int32 `json:"mask_length"`
	Type_ *IpAddressTypeChoices `json:"type,omitempty"`
	IpVersion int32 `json:"ip_version"`
	// Hostname or FQDN (not case-sensitive)
	DnsName string `json:"dns_name,omitempty"`
	Description string `json:"description,omitempty"`
	Status *BulkWritableCableRequestStatus `json:"status"`
	Role *BulkWritableCircuitRequestTenant `json:"role,omitempty"`
	Parent *BulkWritableIpAddressRequestParent `json:"parent,omitempty"`
	Tenant *BulkWritableCircuitRequestTenant `json:"tenant,omitempty"`
	NatInside *NatInside `json:"nat_inside,omitempty"`
	Created time.Time `json:"created"`
	LastUpdated time.Time `json:"last_updated"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	NotesUrl string `json:"notes_url"`
	CustomFields map[string]Object `json:"custom_fields,omitempty"`
	NatOutsideList []BulkWritableCableRequestStatus `json:"nat_outside_list"`
	Interfaces []BulkWritableCableRequestStatus `json:"interfaces"`
	VmInterfaces []BulkWritableCableRequestStatus `json:"vm_interfaces"`
}
