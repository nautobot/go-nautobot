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
type BulkWritableConfigContextRequest struct {
	Id string `json:"id"`
	OwnerContentType string `json:"owner_content_type,omitempty"`
	Name string `json:"name"`
	OwnerObjectId string `json:"owner_object_id,omitempty"`
	Weight int32 `json:"weight,omitempty"`
	Description string `json:"description,omitempty"`
	IsActive bool `json:"is_active,omitempty"`
	Data map[string]Object `json:"data"`
	ConfigContextSchema *BulkWritableConfigContextRequestConfigContextSchema `json:"config_context_schema,omitempty"`
	Locations []BulkWritableCableRequestStatus `json:"locations,omitempty"`
	Roles []BulkWritableCableRequestStatus `json:"roles,omitempty"`
	DeviceTypes []BulkWritableCableRequestStatus `json:"device_types,omitempty"`
	DeviceRedundancyGroups []BulkWritableCableRequestStatus `json:"device_redundancy_groups,omitempty"`
	Platforms []BulkWritableCableRequestStatus `json:"platforms,omitempty"`
	ClusterGroups []BulkWritableCableRequestStatus `json:"cluster_groups,omitempty"`
	Clusters []BulkWritableCableRequestStatus `json:"clusters,omitempty"`
	TenantGroups []BulkWritableCableRequestStatus `json:"tenant_groups,omitempty"`
	Tenants []BulkWritableCableRequestStatus `json:"tenants,omitempty"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
}
