/*
API Documentation

Source of truth and network automation platform

API version: 2.3.2 (2.3)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nautobot

import (
	"encoding/json"
	"fmt"
)

// checks if the PatchedBulkWritableVirtualMachineRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedBulkWritableVirtualMachineRequest{}

// PatchedBulkWritableVirtualMachineRequest Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type PatchedBulkWritableVirtualMachineRequest struct {
	Id string `json:"id"`
	LocalConfigContextData interface{} `json:"local_config_context_data,omitempty"`
	LocalConfigContextDataOwnerObjectId NullableString `json:"local_config_context_data_owner_object_id,omitempty"`
	Name *string `json:"name,omitempty"`
	Vcpus NullableInt32 `json:"vcpus,omitempty"`
	Memory NullableInt32 `json:"memory,omitempty"`
	Disk NullableInt32 `json:"disk,omitempty"`
	Comments *string `json:"comments,omitempty"`
	LocalConfigContextSchema NullableBulkWritableConfigContextRequestConfigContextSchema `json:"local_config_context_schema,omitempty"`
	LocalConfigContextDataOwnerContentType NullableBulkWritableCircuitRequestTenant `json:"local_config_context_data_owner_content_type,omitempty"`
	Cluster *BulkWritableCableRequestStatus `json:"cluster,omitempty"`
	Tenant NullableBulkWritableCircuitRequestTenant `json:"tenant,omitempty"`
	Platform NullableBulkWritableCircuitRequestTenant `json:"platform,omitempty"`
	Status *BulkWritableCableRequestStatus `json:"status,omitempty"`
	Role NullableBulkWritableCircuitRequestTenant `json:"role,omitempty"`
	PrimaryIp4 NullablePrimaryIPv4 `json:"primary_ip4,omitempty"`
	PrimaryIp6 NullablePrimaryIPv6 `json:"primary_ip6,omitempty"`
	SoftwareVersion NullableBulkWritableVirtualMachineRequestSoftwareVersion `json:"software_version,omitempty"`
	// Override the software image files associated with the software version for this virtual machine
	SoftwareImageFiles []SoftwareImageFiles `json:"software_image_files,omitempty"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Relationships *map[string]BulkWritableCableRequestRelationshipsValue `json:"relationships,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedBulkWritableVirtualMachineRequest PatchedBulkWritableVirtualMachineRequest

// NewPatchedBulkWritableVirtualMachineRequest instantiates a new PatchedBulkWritableVirtualMachineRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedBulkWritableVirtualMachineRequest(id string) *PatchedBulkWritableVirtualMachineRequest {
	this := PatchedBulkWritableVirtualMachineRequest{}
	this.Id = id
	return &this
}

// NewPatchedBulkWritableVirtualMachineRequestWithDefaults instantiates a new PatchedBulkWritableVirtualMachineRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedBulkWritableVirtualMachineRequestWithDefaults() *PatchedBulkWritableVirtualMachineRequest {
	this := PatchedBulkWritableVirtualMachineRequest{}
	return &this
}

// GetId returns the Id field value
func (o *PatchedBulkWritableVirtualMachineRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableVirtualMachineRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PatchedBulkWritableVirtualMachineRequest) SetId(v string) {
	o.Id = v
}

// GetLocalConfigContextData returns the LocalConfigContextData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBulkWritableVirtualMachineRequest) GetLocalConfigContextData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.LocalConfigContextData
}

// GetLocalConfigContextDataOk returns a tuple with the LocalConfigContextData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBulkWritableVirtualMachineRequest) GetLocalConfigContextDataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.LocalConfigContextData) {
		return nil, false
	}
	return &o.LocalConfigContextData, true
}

// HasLocalConfigContextData returns a boolean if a field has been set.
func (o *PatchedBulkWritableVirtualMachineRequest) HasLocalConfigContextData() bool {
	if o != nil && !IsNil(o.LocalConfigContextData) {
		return true
	}

	return false
}

// SetLocalConfigContextData gets a reference to the given interface{} and assigns it to the LocalConfigContextData field.
func (o *PatchedBulkWritableVirtualMachineRequest) SetLocalConfigContextData(v interface{}) {
	o.LocalConfigContextData = v
}

// GetLocalConfigContextDataOwnerObjectId returns the LocalConfigContextDataOwnerObjectId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBulkWritableVirtualMachineRequest) GetLocalConfigContextDataOwnerObjectId() string {
	if o == nil || IsNil(o.LocalConfigContextDataOwnerObjectId.Get()) {
		var ret string
		return ret
	}
	return *o.LocalConfigContextDataOwnerObjectId.Get()
}

// GetLocalConfigContextDataOwnerObjectIdOk returns a tuple with the LocalConfigContextDataOwnerObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBulkWritableVirtualMachineRequest) GetLocalConfigContextDataOwnerObjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LocalConfigContextDataOwnerObjectId.Get(), o.LocalConfigContextDataOwnerObjectId.IsSet()
}

// HasLocalConfigContextDataOwnerObjectId returns a boolean if a field has been set.
func (o *PatchedBulkWritableVirtualMachineRequest) HasLocalConfigContextDataOwnerObjectId() bool {
	if o != nil && o.LocalConfigContextDataOwnerObjectId.IsSet() {
		return true
	}

	return false
}

// SetLocalConfigContextDataOwnerObjectId gets a reference to the given NullableString and assigns it to the LocalConfigContextDataOwnerObjectId field.
func (o *PatchedBulkWritableVirtualMachineRequest) SetLocalConfigContextDataOwnerObjectId(v string) {
	o.LocalConfigContextDataOwnerObjectId.Set(&v)
}
// SetLocalConfigContextDataOwnerObjectIdNil sets the value for LocalConfigContextDataOwnerObjectId to be an explicit nil
func (o *PatchedBulkWritableVirtualMachineRequest) SetLocalConfigContextDataOwnerObjectIdNil() {
	o.LocalConfigContextDataOwnerObjectId.Set(nil)
}

// UnsetLocalConfigContextDataOwnerObjectId ensures that no value is present for LocalConfigContextDataOwnerObjectId, not even an explicit nil
func (o *PatchedBulkWritableVirtualMachineRequest) UnsetLocalConfigContextDataOwnerObjectId() {
	o.LocalConfigContextDataOwnerObjectId.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedBulkWritableVirtualMachineRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableVirtualMachineRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedBulkWritableVirtualMachineRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedBulkWritableVirtualMachineRequest) SetName(v string) {
	o.Name = &v
}

// GetVcpus returns the Vcpus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBulkWritableVirtualMachineRequest) GetVcpus() int32 {
	if o == nil || IsNil(o.Vcpus.Get()) {
		var ret int32
		return ret
	}
	return *o.Vcpus.Get()
}

// GetVcpusOk returns a tuple with the Vcpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBulkWritableVirtualMachineRequest) GetVcpusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vcpus.Get(), o.Vcpus.IsSet()
}

// HasVcpus returns a boolean if a field has been set.
func (o *PatchedBulkWritableVirtualMachineRequest) HasVcpus() bool {
	if o != nil && o.Vcpus.IsSet() {
		return true
	}

	return false
}

// SetVcpus gets a reference to the given NullableInt32 and assigns it to the Vcpus field.
func (o *PatchedBulkWritableVirtualMachineRequest) SetVcpus(v int32) {
	o.Vcpus.Set(&v)
}
// SetVcpusNil sets the value for Vcpus to be an explicit nil
func (o *PatchedBulkWritableVirtualMachineRequest) SetVcpusNil() {
	o.Vcpus.Set(nil)
}

// UnsetVcpus ensures that no value is present for Vcpus, not even an explicit nil
func (o *PatchedBulkWritableVirtualMachineRequest) UnsetVcpus() {
	o.Vcpus.Unset()
}

// GetMemory returns the Memory field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBulkWritableVirtualMachineRequest) GetMemory() int32 {
	if o == nil || IsNil(o.Memory.Get()) {
		var ret int32
		return ret
	}
	return *o.Memory.Get()
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBulkWritableVirtualMachineRequest) GetMemoryOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Memory.Get(), o.Memory.IsSet()
}

// HasMemory returns a boolean if a field has been set.
func (o *PatchedBulkWritableVirtualMachineRequest) HasMemory() bool {
	if o != nil && o.Memory.IsSet() {
		return true
	}

	return false
}

// SetMemory gets a reference to the given NullableInt32 and assigns it to the Memory field.
func (o *PatchedBulkWritableVirtualMachineRequest) SetMemory(v int32) {
	o.Memory.Set(&v)
}
// SetMemoryNil sets the value for Memory to be an explicit nil
func (o *PatchedBulkWritableVirtualMachineRequest) SetMemoryNil() {
	o.Memory.Set(nil)
}

// UnsetMemory ensures that no value is present for Memory, not even an explicit nil
func (o *PatchedBulkWritableVirtualMachineRequest) UnsetMemory() {
	o.Memory.Unset()
}

// GetDisk returns the Disk field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBulkWritableVirtualMachineRequest) GetDisk() int32 {
	if o == nil || IsNil(o.Disk.Get()) {
		var ret int32
		return ret
	}
	return *o.Disk.Get()
}

// GetDiskOk returns a tuple with the Disk field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBulkWritableVirtualMachineRequest) GetDiskOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Disk.Get(), o.Disk.IsSet()
}

// HasDisk returns a boolean if a field has been set.
func (o *PatchedBulkWritableVirtualMachineRequest) HasDisk() bool {
	if o != nil && o.Disk.IsSet() {
		return true
	}

	return false
}

// SetDisk gets a reference to the given NullableInt32 and assigns it to the Disk field.
func (o *PatchedBulkWritableVirtualMachineRequest) SetDisk(v int32) {
	o.Disk.Set(&v)
}
// SetDiskNil sets the value for Disk to be an explicit nil
func (o *PatchedBulkWritableVirtualMachineRequest) SetDiskNil() {
	o.Disk.Set(nil)
}

// UnsetDisk ensures that no value is present for Disk, not even an explicit nil
func (o *PatchedBulkWritableVirtualMachineRequest) UnsetDisk() {
	o.Disk.Unset()
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *PatchedBulkWritableVirtualMachineRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableVirtualMachineRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *PatchedBulkWritableVirtualMachineRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *PatchedBulkWritableVirtualMachineRequest) SetComments(v string) {
	o.Comments = &v
}

// GetLocalConfigContextSchema returns the LocalConfigContextSchema field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBulkWritableVirtualMachineRequest) GetLocalConfigContextSchema() BulkWritableConfigContextRequestConfigContextSchema {
	if o == nil || IsNil(o.LocalConfigContextSchema.Get()) {
		var ret BulkWritableConfigContextRequestConfigContextSchema
		return ret
	}
	return *o.LocalConfigContextSchema.Get()
}

// GetLocalConfigContextSchemaOk returns a tuple with the LocalConfigContextSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBulkWritableVirtualMachineRequest) GetLocalConfigContextSchemaOk() (*BulkWritableConfigContextRequestConfigContextSchema, bool) {
	if o == nil {
		return nil, false
	}
	return o.LocalConfigContextSchema.Get(), o.LocalConfigContextSchema.IsSet()
}

// HasLocalConfigContextSchema returns a boolean if a field has been set.
func (o *PatchedBulkWritableVirtualMachineRequest) HasLocalConfigContextSchema() bool {
	if o != nil && o.LocalConfigContextSchema.IsSet() {
		return true
	}

	return false
}

// SetLocalConfigContextSchema gets a reference to the given NullableBulkWritableConfigContextRequestConfigContextSchema and assigns it to the LocalConfigContextSchema field.
func (o *PatchedBulkWritableVirtualMachineRequest) SetLocalConfigContextSchema(v BulkWritableConfigContextRequestConfigContextSchema) {
	o.LocalConfigContextSchema.Set(&v)
}
// SetLocalConfigContextSchemaNil sets the value for LocalConfigContextSchema to be an explicit nil
func (o *PatchedBulkWritableVirtualMachineRequest) SetLocalConfigContextSchemaNil() {
	o.LocalConfigContextSchema.Set(nil)
}

// UnsetLocalConfigContextSchema ensures that no value is present for LocalConfigContextSchema, not even an explicit nil
func (o *PatchedBulkWritableVirtualMachineRequest) UnsetLocalConfigContextSchema() {
	o.LocalConfigContextSchema.Unset()
}

// GetLocalConfigContextDataOwnerContentType returns the LocalConfigContextDataOwnerContentType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBulkWritableVirtualMachineRequest) GetLocalConfigContextDataOwnerContentType() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.LocalConfigContextDataOwnerContentType.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.LocalConfigContextDataOwnerContentType.Get()
}

// GetLocalConfigContextDataOwnerContentTypeOk returns a tuple with the LocalConfigContextDataOwnerContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBulkWritableVirtualMachineRequest) GetLocalConfigContextDataOwnerContentTypeOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.LocalConfigContextDataOwnerContentType.Get(), o.LocalConfigContextDataOwnerContentType.IsSet()
}

// HasLocalConfigContextDataOwnerContentType returns a boolean if a field has been set.
func (o *PatchedBulkWritableVirtualMachineRequest) HasLocalConfigContextDataOwnerContentType() bool {
	if o != nil && o.LocalConfigContextDataOwnerContentType.IsSet() {
		return true
	}

	return false
}

// SetLocalConfigContextDataOwnerContentType gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the LocalConfigContextDataOwnerContentType field.
func (o *PatchedBulkWritableVirtualMachineRequest) SetLocalConfigContextDataOwnerContentType(v BulkWritableCircuitRequestTenant) {
	o.LocalConfigContextDataOwnerContentType.Set(&v)
}
// SetLocalConfigContextDataOwnerContentTypeNil sets the value for LocalConfigContextDataOwnerContentType to be an explicit nil
func (o *PatchedBulkWritableVirtualMachineRequest) SetLocalConfigContextDataOwnerContentTypeNil() {
	o.LocalConfigContextDataOwnerContentType.Set(nil)
}

// UnsetLocalConfigContextDataOwnerContentType ensures that no value is present for LocalConfigContextDataOwnerContentType, not even an explicit nil
func (o *PatchedBulkWritableVirtualMachineRequest) UnsetLocalConfigContextDataOwnerContentType() {
	o.LocalConfigContextDataOwnerContentType.Unset()
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *PatchedBulkWritableVirtualMachineRequest) GetCluster() BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Cluster) {
		var ret BulkWritableCableRequestStatus
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableVirtualMachineRequest) GetClusterOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *PatchedBulkWritableVirtualMachineRequest) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given BulkWritableCableRequestStatus and assigns it to the Cluster field.
func (o *PatchedBulkWritableVirtualMachineRequest) SetCluster(v BulkWritableCableRequestStatus) {
	o.Cluster = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBulkWritableVirtualMachineRequest) GetTenant() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBulkWritableVirtualMachineRequest) GetTenantOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *PatchedBulkWritableVirtualMachineRequest) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the Tenant field.
func (o *PatchedBulkWritableVirtualMachineRequest) SetTenant(v BulkWritableCircuitRequestTenant) {
	o.Tenant.Set(&v)
}
// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *PatchedBulkWritableVirtualMachineRequest) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *PatchedBulkWritableVirtualMachineRequest) UnsetTenant() {
	o.Tenant.Unset()
}

// GetPlatform returns the Platform field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBulkWritableVirtualMachineRequest) GetPlatform() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.Platform.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.Platform.Get()
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBulkWritableVirtualMachineRequest) GetPlatformOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Platform.Get(), o.Platform.IsSet()
}

// HasPlatform returns a boolean if a field has been set.
func (o *PatchedBulkWritableVirtualMachineRequest) HasPlatform() bool {
	if o != nil && o.Platform.IsSet() {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the Platform field.
func (o *PatchedBulkWritableVirtualMachineRequest) SetPlatform(v BulkWritableCircuitRequestTenant) {
	o.Platform.Set(&v)
}
// SetPlatformNil sets the value for Platform to be an explicit nil
func (o *PatchedBulkWritableVirtualMachineRequest) SetPlatformNil() {
	o.Platform.Set(nil)
}

// UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
func (o *PatchedBulkWritableVirtualMachineRequest) UnsetPlatform() {
	o.Platform.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PatchedBulkWritableVirtualMachineRequest) GetStatus() BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Status) {
		var ret BulkWritableCableRequestStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableVirtualMachineRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PatchedBulkWritableVirtualMachineRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given BulkWritableCableRequestStatus and assigns it to the Status field.
func (o *PatchedBulkWritableVirtualMachineRequest) SetStatus(v BulkWritableCableRequestStatus) {
	o.Status = &v
}

// GetRole returns the Role field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBulkWritableVirtualMachineRequest) GetRole() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.Role.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.Role.Get()
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBulkWritableVirtualMachineRequest) GetRoleOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Role.Get(), o.Role.IsSet()
}

// HasRole returns a boolean if a field has been set.
func (o *PatchedBulkWritableVirtualMachineRequest) HasRole() bool {
	if o != nil && o.Role.IsSet() {
		return true
	}

	return false
}

// SetRole gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the Role field.
func (o *PatchedBulkWritableVirtualMachineRequest) SetRole(v BulkWritableCircuitRequestTenant) {
	o.Role.Set(&v)
}
// SetRoleNil sets the value for Role to be an explicit nil
func (o *PatchedBulkWritableVirtualMachineRequest) SetRoleNil() {
	o.Role.Set(nil)
}

// UnsetRole ensures that no value is present for Role, not even an explicit nil
func (o *PatchedBulkWritableVirtualMachineRequest) UnsetRole() {
	o.Role.Unset()
}

// GetPrimaryIp4 returns the PrimaryIp4 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBulkWritableVirtualMachineRequest) GetPrimaryIp4() PrimaryIPv4 {
	if o == nil || IsNil(o.PrimaryIp4.Get()) {
		var ret PrimaryIPv4
		return ret
	}
	return *o.PrimaryIp4.Get()
}

// GetPrimaryIp4Ok returns a tuple with the PrimaryIp4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBulkWritableVirtualMachineRequest) GetPrimaryIp4Ok() (*PrimaryIPv4, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrimaryIp4.Get(), o.PrimaryIp4.IsSet()
}

// HasPrimaryIp4 returns a boolean if a field has been set.
func (o *PatchedBulkWritableVirtualMachineRequest) HasPrimaryIp4() bool {
	if o != nil && o.PrimaryIp4.IsSet() {
		return true
	}

	return false
}

// SetPrimaryIp4 gets a reference to the given NullablePrimaryIPv4 and assigns it to the PrimaryIp4 field.
func (o *PatchedBulkWritableVirtualMachineRequest) SetPrimaryIp4(v PrimaryIPv4) {
	o.PrimaryIp4.Set(&v)
}
// SetPrimaryIp4Nil sets the value for PrimaryIp4 to be an explicit nil
func (o *PatchedBulkWritableVirtualMachineRequest) SetPrimaryIp4Nil() {
	o.PrimaryIp4.Set(nil)
}

// UnsetPrimaryIp4 ensures that no value is present for PrimaryIp4, not even an explicit nil
func (o *PatchedBulkWritableVirtualMachineRequest) UnsetPrimaryIp4() {
	o.PrimaryIp4.Unset()
}

// GetPrimaryIp6 returns the PrimaryIp6 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBulkWritableVirtualMachineRequest) GetPrimaryIp6() PrimaryIPv6 {
	if o == nil || IsNil(o.PrimaryIp6.Get()) {
		var ret PrimaryIPv6
		return ret
	}
	return *o.PrimaryIp6.Get()
}

// GetPrimaryIp6Ok returns a tuple with the PrimaryIp6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBulkWritableVirtualMachineRequest) GetPrimaryIp6Ok() (*PrimaryIPv6, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrimaryIp6.Get(), o.PrimaryIp6.IsSet()
}

// HasPrimaryIp6 returns a boolean if a field has been set.
func (o *PatchedBulkWritableVirtualMachineRequest) HasPrimaryIp6() bool {
	if o != nil && o.PrimaryIp6.IsSet() {
		return true
	}

	return false
}

// SetPrimaryIp6 gets a reference to the given NullablePrimaryIPv6 and assigns it to the PrimaryIp6 field.
func (o *PatchedBulkWritableVirtualMachineRequest) SetPrimaryIp6(v PrimaryIPv6) {
	o.PrimaryIp6.Set(&v)
}
// SetPrimaryIp6Nil sets the value for PrimaryIp6 to be an explicit nil
func (o *PatchedBulkWritableVirtualMachineRequest) SetPrimaryIp6Nil() {
	o.PrimaryIp6.Set(nil)
}

// UnsetPrimaryIp6 ensures that no value is present for PrimaryIp6, not even an explicit nil
func (o *PatchedBulkWritableVirtualMachineRequest) UnsetPrimaryIp6() {
	o.PrimaryIp6.Unset()
}

// GetSoftwareVersion returns the SoftwareVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBulkWritableVirtualMachineRequest) GetSoftwareVersion() BulkWritableVirtualMachineRequestSoftwareVersion {
	if o == nil || IsNil(o.SoftwareVersion.Get()) {
		var ret BulkWritableVirtualMachineRequestSoftwareVersion
		return ret
	}
	return *o.SoftwareVersion.Get()
}

// GetSoftwareVersionOk returns a tuple with the SoftwareVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBulkWritableVirtualMachineRequest) GetSoftwareVersionOk() (*BulkWritableVirtualMachineRequestSoftwareVersion, bool) {
	if o == nil {
		return nil, false
	}
	return o.SoftwareVersion.Get(), o.SoftwareVersion.IsSet()
}

// HasSoftwareVersion returns a boolean if a field has been set.
func (o *PatchedBulkWritableVirtualMachineRequest) HasSoftwareVersion() bool {
	if o != nil && o.SoftwareVersion.IsSet() {
		return true
	}

	return false
}

// SetSoftwareVersion gets a reference to the given NullableBulkWritableVirtualMachineRequestSoftwareVersion and assigns it to the SoftwareVersion field.
func (o *PatchedBulkWritableVirtualMachineRequest) SetSoftwareVersion(v BulkWritableVirtualMachineRequestSoftwareVersion) {
	o.SoftwareVersion.Set(&v)
}
// SetSoftwareVersionNil sets the value for SoftwareVersion to be an explicit nil
func (o *PatchedBulkWritableVirtualMachineRequest) SetSoftwareVersionNil() {
	o.SoftwareVersion.Set(nil)
}

// UnsetSoftwareVersion ensures that no value is present for SoftwareVersion, not even an explicit nil
func (o *PatchedBulkWritableVirtualMachineRequest) UnsetSoftwareVersion() {
	o.SoftwareVersion.Unset()
}

// GetSoftwareImageFiles returns the SoftwareImageFiles field value if set, zero value otherwise.
func (o *PatchedBulkWritableVirtualMachineRequest) GetSoftwareImageFiles() []SoftwareImageFiles {
	if o == nil || IsNil(o.SoftwareImageFiles) {
		var ret []SoftwareImageFiles
		return ret
	}
	return o.SoftwareImageFiles
}

// GetSoftwareImageFilesOk returns a tuple with the SoftwareImageFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableVirtualMachineRequest) GetSoftwareImageFilesOk() ([]SoftwareImageFiles, bool) {
	if o == nil || IsNil(o.SoftwareImageFiles) {
		return nil, false
	}
	return o.SoftwareImageFiles, true
}

// HasSoftwareImageFiles returns a boolean if a field has been set.
func (o *PatchedBulkWritableVirtualMachineRequest) HasSoftwareImageFiles() bool {
	if o != nil && !IsNil(o.SoftwareImageFiles) {
		return true
	}

	return false
}

// SetSoftwareImageFiles gets a reference to the given []SoftwareImageFiles and assigns it to the SoftwareImageFiles field.
func (o *PatchedBulkWritableVirtualMachineRequest) SetSoftwareImageFiles(v []SoftwareImageFiles) {
	o.SoftwareImageFiles = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedBulkWritableVirtualMachineRequest) GetTags() []BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Tags) {
		var ret []BulkWritableCableRequestStatus
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableVirtualMachineRequest) GetTagsOk() ([]BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedBulkWritableVirtualMachineRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []BulkWritableCableRequestStatus and assigns it to the Tags field.
func (o *PatchedBulkWritableVirtualMachineRequest) SetTags(v []BulkWritableCableRequestStatus) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedBulkWritableVirtualMachineRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableVirtualMachineRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedBulkWritableVirtualMachineRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedBulkWritableVirtualMachineRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PatchedBulkWritableVirtualMachineRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue {
	if o == nil || IsNil(o.Relationships) {
		var ret map[string]BulkWritableCableRequestRelationshipsValue
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableVirtualMachineRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PatchedBulkWritableVirtualMachineRequest) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]BulkWritableCableRequestRelationshipsValue and assigns it to the Relationships field.
func (o *PatchedBulkWritableVirtualMachineRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue) {
	o.Relationships = &v
}

func (o PatchedBulkWritableVirtualMachineRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedBulkWritableVirtualMachineRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if o.LocalConfigContextData != nil {
		toSerialize["local_config_context_data"] = o.LocalConfigContextData
	}
	if o.LocalConfigContextDataOwnerObjectId.IsSet() {
		toSerialize["local_config_context_data_owner_object_id"] = o.LocalConfigContextDataOwnerObjectId.Get()
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.Vcpus.IsSet() {
		toSerialize["vcpus"] = o.Vcpus.Get()
	}
	if o.Memory.IsSet() {
		toSerialize["memory"] = o.Memory.Get()
	}
	if o.Disk.IsSet() {
		toSerialize["disk"] = o.Disk.Get()
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	if o.LocalConfigContextSchema.IsSet() {
		toSerialize["local_config_context_schema"] = o.LocalConfigContextSchema.Get()
	}
	if o.LocalConfigContextDataOwnerContentType.IsSet() {
		toSerialize["local_config_context_data_owner_content_type"] = o.LocalConfigContextDataOwnerContentType.Get()
	}
	if !IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if o.Platform.IsSet() {
		toSerialize["platform"] = o.Platform.Get()
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.Role.IsSet() {
		toSerialize["role"] = o.Role.Get()
	}
	if o.PrimaryIp4.IsSet() {
		toSerialize["primary_ip4"] = o.PrimaryIp4.Get()
	}
	if o.PrimaryIp6.IsSet() {
		toSerialize["primary_ip6"] = o.PrimaryIp6.Get()
	}
	if o.SoftwareVersion.IsSet() {
		toSerialize["software_version"] = o.SoftwareVersion.Get()
	}
	if !IsNil(o.SoftwareImageFiles) {
		toSerialize["software_image_files"] = o.SoftwareImageFiles
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedBulkWritableVirtualMachineRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPatchedBulkWritableVirtualMachineRequest := _PatchedBulkWritableVirtualMachineRequest{}

	err = json.Unmarshal(data, &varPatchedBulkWritableVirtualMachineRequest)

	if err != nil {
		return err
	}

	*o = PatchedBulkWritableVirtualMachineRequest(varPatchedBulkWritableVirtualMachineRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "local_config_context_data")
		delete(additionalProperties, "local_config_context_data_owner_object_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "vcpus")
		delete(additionalProperties, "memory")
		delete(additionalProperties, "disk")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "local_config_context_schema")
		delete(additionalProperties, "local_config_context_data_owner_content_type")
		delete(additionalProperties, "cluster")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "platform")
		delete(additionalProperties, "status")
		delete(additionalProperties, "role")
		delete(additionalProperties, "primary_ip4")
		delete(additionalProperties, "primary_ip6")
		delete(additionalProperties, "software_version")
		delete(additionalProperties, "software_image_files")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "relationships")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedBulkWritableVirtualMachineRequest struct {
	value *PatchedBulkWritableVirtualMachineRequest
	isSet bool
}

func (v NullablePatchedBulkWritableVirtualMachineRequest) Get() *PatchedBulkWritableVirtualMachineRequest {
	return v.value
}

func (v *NullablePatchedBulkWritableVirtualMachineRequest) Set(val *PatchedBulkWritableVirtualMachineRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedBulkWritableVirtualMachineRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedBulkWritableVirtualMachineRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedBulkWritableVirtualMachineRequest(val *PatchedBulkWritableVirtualMachineRequest) *NullablePatchedBulkWritableVirtualMachineRequest {
	return &NullablePatchedBulkWritableVirtualMachineRequest{value: val, isSet: true}
}

func (v NullablePatchedBulkWritableVirtualMachineRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedBulkWritableVirtualMachineRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


