/*
 * API Documentation
 *
 * Source of truth and network automation platform
 *
 * API version: 2.2.5 (2.2)
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package go-nautobot

// InterfaceRedundancyGroupAssociation Serializer.
type InterfaceRedundancyGroupAssociationRequest struct {
	Priority int32 `json:"priority"`
	InterfaceRedundancyGroup *BulkWritableCableRequestStatus `json:"interface_redundancy_group"`
	Interface_ *BulkWritableCableRequestStatus `json:"interface"`
}
