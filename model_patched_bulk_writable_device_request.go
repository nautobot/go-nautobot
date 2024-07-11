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
type PatchedBulkWritableDeviceRequest struct {
	Id string `json:"id"`
	Face *FaceEnum `json:"face,omitempty"`
	LocalConfigContextData map[string]Object `json:"local_config_context_data,omitempty"`
	LocalConfigContextDataOwnerObjectId string `json:"local_config_context_data_owner_object_id,omitempty"`
	Name string `json:"name,omitempty"`
	Serial string `json:"serial,omitempty"`
	// A unique tag used to identify this device
	AssetTag string `json:"asset_tag,omitempty"`
	// The lowest-numbered unit occupied by the device
	Position int32 `json:"position,omitempty"`
	// The priority the device has in the device redundancy group.
	DeviceRedundancyGroupPriority int32 `json:"device_redundancy_group_priority,omitempty"`
	VcPosition int32 `json:"vc_position,omitempty"`
	VcPriority int32 `json:"vc_priority,omitempty"`
	Comments string `json:"comments,omitempty"`
	LocalConfigContextSchema *BulkWritableConfigContextRequestConfigContextSchema `json:"local_config_context_schema,omitempty"`
	LocalConfigContextDataOwnerContentType *BulkWritableCircuitRequestTenant `json:"local_config_context_data_owner_content_type,omitempty"`
	DeviceType *BulkWritableCableRequestStatus `json:"device_type,omitempty"`
	Status *BulkWritableCableRequestStatus `json:"status,omitempty"`
	Role *BulkWritableCableRequestStatus `json:"role,omitempty"`
	Tenant *BulkWritableCircuitRequestTenant `json:"tenant,omitempty"`
	Platform *BulkWritableCircuitRequestTenant `json:"platform,omitempty"`
	Location *BulkWritableCableRequestStatus `json:"location,omitempty"`
	Rack *BulkWritableCircuitRequestTenant `json:"rack,omitempty"`
	PrimaryIp4 *PrimaryIpv4 `json:"primary_ip4,omitempty"`
	PrimaryIp6 *PrimaryIpv6 `json:"primary_ip6,omitempty"`
	Cluster *BulkWritableCircuitRequestTenant `json:"cluster,omitempty"`
	VirtualChassis *BulkWritableCircuitRequestTenant `json:"virtual_chassis,omitempty"`
	DeviceRedundancyGroup *BulkWritableCircuitRequestTenant `json:"device_redundancy_group,omitempty"`
	SoftwareVersion *BulkWritableDeviceRequestSoftwareVersion `json:"software_version,omitempty"`
	SecretsGroup *BulkWritableCircuitRequestTenant `json:"secrets_group,omitempty"`
	ControllerManagedDeviceGroup *BulkWritableCircuitRequestTenant `json:"controller_managed_device_group,omitempty"`
	// Override the software image files associated with the software version for this device
	SoftwareImageFiles []SoftwareImageFiles `json:"software_image_files,omitempty"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	CustomFields map[string]Object `json:"custom_fields,omitempty"`
	Relationships map[string]BulkWritableCableRequestRelationships `json:"relationships,omitempty"`
	ParentBay *BulkWritableCircuitRequestTenant `json:"parent_bay,omitempty"`
}
