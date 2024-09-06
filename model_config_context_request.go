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

// checks if the ConfigContextRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigContextRequest{}

// ConfigContextRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type ConfigContextRequest struct {
	OwnerContentType NullableString `json:"owner_content_type,omitempty"`
	Name string `json:"name"`
	OwnerObjectId NullableString `json:"owner_object_id,omitempty"`
	Weight *int32 `json:"weight,omitempty"`
	Description *string `json:"description,omitempty"`
	IsActive *bool `json:"is_active,omitempty"`
	Data interface{} `json:"data"`
	ConfigContextSchema NullableBulkWritableConfigContextRequestConfigContextSchema `json:"config_context_schema,omitempty"`
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
	AdditionalProperties map[string]interface{}
}

type _ConfigContextRequest ConfigContextRequest

// NewConfigContextRequest instantiates a new ConfigContextRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigContextRequest(name string, data interface{}) *ConfigContextRequest {
	this := ConfigContextRequest{}
	this.Name = name
	this.Data = data
	return &this
}

// NewConfigContextRequestWithDefaults instantiates a new ConfigContextRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigContextRequestWithDefaults() *ConfigContextRequest {
	this := ConfigContextRequest{}
	return &this
}

// GetOwnerContentType returns the OwnerContentType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigContextRequest) GetOwnerContentType() string {
	if o == nil || IsNil(o.OwnerContentType.Get()) {
		var ret string
		return ret
	}
	return *o.OwnerContentType.Get()
}

// GetOwnerContentTypeOk returns a tuple with the OwnerContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigContextRequest) GetOwnerContentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OwnerContentType.Get(), o.OwnerContentType.IsSet()
}

// HasOwnerContentType returns a boolean if a field has been set.
func (o *ConfigContextRequest) HasOwnerContentType() bool {
	if o != nil && o.OwnerContentType.IsSet() {
		return true
	}

	return false
}

// SetOwnerContentType gets a reference to the given NullableString and assigns it to the OwnerContentType field.
func (o *ConfigContextRequest) SetOwnerContentType(v string) {
	o.OwnerContentType.Set(&v)
}
// SetOwnerContentTypeNil sets the value for OwnerContentType to be an explicit nil
func (o *ConfigContextRequest) SetOwnerContentTypeNil() {
	o.OwnerContentType.Set(nil)
}

// UnsetOwnerContentType ensures that no value is present for OwnerContentType, not even an explicit nil
func (o *ConfigContextRequest) UnsetOwnerContentType() {
	o.OwnerContentType.Unset()
}

// GetName returns the Name field value
func (o *ConfigContextRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ConfigContextRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ConfigContextRequest) SetName(v string) {
	o.Name = v
}

// GetOwnerObjectId returns the OwnerObjectId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigContextRequest) GetOwnerObjectId() string {
	if o == nil || IsNil(o.OwnerObjectId.Get()) {
		var ret string
		return ret
	}
	return *o.OwnerObjectId.Get()
}

// GetOwnerObjectIdOk returns a tuple with the OwnerObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigContextRequest) GetOwnerObjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OwnerObjectId.Get(), o.OwnerObjectId.IsSet()
}

// HasOwnerObjectId returns a boolean if a field has been set.
func (o *ConfigContextRequest) HasOwnerObjectId() bool {
	if o != nil && o.OwnerObjectId.IsSet() {
		return true
	}

	return false
}

// SetOwnerObjectId gets a reference to the given NullableString and assigns it to the OwnerObjectId field.
func (o *ConfigContextRequest) SetOwnerObjectId(v string) {
	o.OwnerObjectId.Set(&v)
}
// SetOwnerObjectIdNil sets the value for OwnerObjectId to be an explicit nil
func (o *ConfigContextRequest) SetOwnerObjectIdNil() {
	o.OwnerObjectId.Set(nil)
}

// UnsetOwnerObjectId ensures that no value is present for OwnerObjectId, not even an explicit nil
func (o *ConfigContextRequest) UnsetOwnerObjectId() {
	o.OwnerObjectId.Unset()
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *ConfigContextRequest) GetWeight() int32 {
	if o == nil || IsNil(o.Weight) {
		var ret int32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigContextRequest) GetWeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *ConfigContextRequest) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int32 and assigns it to the Weight field.
func (o *ConfigContextRequest) SetWeight(v int32) {
	o.Weight = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ConfigContextRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigContextRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ConfigContextRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ConfigContextRequest) SetDescription(v string) {
	o.Description = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *ConfigContextRequest) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigContextRequest) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *ConfigContextRequest) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *ConfigContextRequest) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetData returns the Data field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ConfigContextRequest) GetData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigContextRequest) GetDataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ConfigContextRequest) SetData(v interface{}) {
	o.Data = v
}

// GetConfigContextSchema returns the ConfigContextSchema field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigContextRequest) GetConfigContextSchema() BulkWritableConfigContextRequestConfigContextSchema {
	if o == nil || IsNil(o.ConfigContextSchema.Get()) {
		var ret BulkWritableConfigContextRequestConfigContextSchema
		return ret
	}
	return *o.ConfigContextSchema.Get()
}

// GetConfigContextSchemaOk returns a tuple with the ConfigContextSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigContextRequest) GetConfigContextSchemaOk() (*BulkWritableConfigContextRequestConfigContextSchema, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigContextSchema.Get(), o.ConfigContextSchema.IsSet()
}

// HasConfigContextSchema returns a boolean if a field has been set.
func (o *ConfigContextRequest) HasConfigContextSchema() bool {
	if o != nil && o.ConfigContextSchema.IsSet() {
		return true
	}

	return false
}

// SetConfigContextSchema gets a reference to the given NullableBulkWritableConfigContextRequestConfigContextSchema and assigns it to the ConfigContextSchema field.
func (o *ConfigContextRequest) SetConfigContextSchema(v BulkWritableConfigContextRequestConfigContextSchema) {
	o.ConfigContextSchema.Set(&v)
}
// SetConfigContextSchemaNil sets the value for ConfigContextSchema to be an explicit nil
func (o *ConfigContextRequest) SetConfigContextSchemaNil() {
	o.ConfigContextSchema.Set(nil)
}

// UnsetConfigContextSchema ensures that no value is present for ConfigContextSchema, not even an explicit nil
func (o *ConfigContextRequest) UnsetConfigContextSchema() {
	o.ConfigContextSchema.Unset()
}

// GetLocations returns the Locations field value if set, zero value otherwise.
func (o *ConfigContextRequest) GetLocations() []BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Locations) {
		var ret []BulkWritableCableRequestStatus
		return ret
	}
	return o.Locations
}

// GetLocationsOk returns a tuple with the Locations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigContextRequest) GetLocationsOk() ([]BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Locations) {
		return nil, false
	}
	return o.Locations, true
}

// HasLocations returns a boolean if a field has been set.
func (o *ConfigContextRequest) HasLocations() bool {
	if o != nil && !IsNil(o.Locations) {
		return true
	}

	return false
}

// SetLocations gets a reference to the given []BulkWritableCableRequestStatus and assigns it to the Locations field.
func (o *ConfigContextRequest) SetLocations(v []BulkWritableCableRequestStatus) {
	o.Locations = v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *ConfigContextRequest) GetRoles() []BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Roles) {
		var ret []BulkWritableCableRequestStatus
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigContextRequest) GetRolesOk() ([]BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *ConfigContextRequest) HasRoles() bool {
	if o != nil && !IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []BulkWritableCableRequestStatus and assigns it to the Roles field.
func (o *ConfigContextRequest) SetRoles(v []BulkWritableCableRequestStatus) {
	o.Roles = v
}

// GetDeviceTypes returns the DeviceTypes field value if set, zero value otherwise.
func (o *ConfigContextRequest) GetDeviceTypes() []BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.DeviceTypes) {
		var ret []BulkWritableCableRequestStatus
		return ret
	}
	return o.DeviceTypes
}

// GetDeviceTypesOk returns a tuple with the DeviceTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigContextRequest) GetDeviceTypesOk() ([]BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.DeviceTypes) {
		return nil, false
	}
	return o.DeviceTypes, true
}

// HasDeviceTypes returns a boolean if a field has been set.
func (o *ConfigContextRequest) HasDeviceTypes() bool {
	if o != nil && !IsNil(o.DeviceTypes) {
		return true
	}

	return false
}

// SetDeviceTypes gets a reference to the given []BulkWritableCableRequestStatus and assigns it to the DeviceTypes field.
func (o *ConfigContextRequest) SetDeviceTypes(v []BulkWritableCableRequestStatus) {
	o.DeviceTypes = v
}

// GetDeviceRedundancyGroups returns the DeviceRedundancyGroups field value if set, zero value otherwise.
func (o *ConfigContextRequest) GetDeviceRedundancyGroups() []BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.DeviceRedundancyGroups) {
		var ret []BulkWritableCableRequestStatus
		return ret
	}
	return o.DeviceRedundancyGroups
}

// GetDeviceRedundancyGroupsOk returns a tuple with the DeviceRedundancyGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigContextRequest) GetDeviceRedundancyGroupsOk() ([]BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.DeviceRedundancyGroups) {
		return nil, false
	}
	return o.DeviceRedundancyGroups, true
}

// HasDeviceRedundancyGroups returns a boolean if a field has been set.
func (o *ConfigContextRequest) HasDeviceRedundancyGroups() bool {
	if o != nil && !IsNil(o.DeviceRedundancyGroups) {
		return true
	}

	return false
}

// SetDeviceRedundancyGroups gets a reference to the given []BulkWritableCableRequestStatus and assigns it to the DeviceRedundancyGroups field.
func (o *ConfigContextRequest) SetDeviceRedundancyGroups(v []BulkWritableCableRequestStatus) {
	o.DeviceRedundancyGroups = v
}

// GetPlatforms returns the Platforms field value if set, zero value otherwise.
func (o *ConfigContextRequest) GetPlatforms() []BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Platforms) {
		var ret []BulkWritableCableRequestStatus
		return ret
	}
	return o.Platforms
}

// GetPlatformsOk returns a tuple with the Platforms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigContextRequest) GetPlatformsOk() ([]BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Platforms) {
		return nil, false
	}
	return o.Platforms, true
}

// HasPlatforms returns a boolean if a field has been set.
func (o *ConfigContextRequest) HasPlatforms() bool {
	if o != nil && !IsNil(o.Platforms) {
		return true
	}

	return false
}

// SetPlatforms gets a reference to the given []BulkWritableCableRequestStatus and assigns it to the Platforms field.
func (o *ConfigContextRequest) SetPlatforms(v []BulkWritableCableRequestStatus) {
	o.Platforms = v
}

// GetClusterGroups returns the ClusterGroups field value if set, zero value otherwise.
func (o *ConfigContextRequest) GetClusterGroups() []BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.ClusterGroups) {
		var ret []BulkWritableCableRequestStatus
		return ret
	}
	return o.ClusterGroups
}

// GetClusterGroupsOk returns a tuple with the ClusterGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigContextRequest) GetClusterGroupsOk() ([]BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.ClusterGroups) {
		return nil, false
	}
	return o.ClusterGroups, true
}

// HasClusterGroups returns a boolean if a field has been set.
func (o *ConfigContextRequest) HasClusterGroups() bool {
	if o != nil && !IsNil(o.ClusterGroups) {
		return true
	}

	return false
}

// SetClusterGroups gets a reference to the given []BulkWritableCableRequestStatus and assigns it to the ClusterGroups field.
func (o *ConfigContextRequest) SetClusterGroups(v []BulkWritableCableRequestStatus) {
	o.ClusterGroups = v
}

// GetClusters returns the Clusters field value if set, zero value otherwise.
func (o *ConfigContextRequest) GetClusters() []BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Clusters) {
		var ret []BulkWritableCableRequestStatus
		return ret
	}
	return o.Clusters
}

// GetClustersOk returns a tuple with the Clusters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigContextRequest) GetClustersOk() ([]BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Clusters) {
		return nil, false
	}
	return o.Clusters, true
}

// HasClusters returns a boolean if a field has been set.
func (o *ConfigContextRequest) HasClusters() bool {
	if o != nil && !IsNil(o.Clusters) {
		return true
	}

	return false
}

// SetClusters gets a reference to the given []BulkWritableCableRequestStatus and assigns it to the Clusters field.
func (o *ConfigContextRequest) SetClusters(v []BulkWritableCableRequestStatus) {
	o.Clusters = v
}

// GetTenantGroups returns the TenantGroups field value if set, zero value otherwise.
func (o *ConfigContextRequest) GetTenantGroups() []BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.TenantGroups) {
		var ret []BulkWritableCableRequestStatus
		return ret
	}
	return o.TenantGroups
}

// GetTenantGroupsOk returns a tuple with the TenantGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigContextRequest) GetTenantGroupsOk() ([]BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.TenantGroups) {
		return nil, false
	}
	return o.TenantGroups, true
}

// HasTenantGroups returns a boolean if a field has been set.
func (o *ConfigContextRequest) HasTenantGroups() bool {
	if o != nil && !IsNil(o.TenantGroups) {
		return true
	}

	return false
}

// SetTenantGroups gets a reference to the given []BulkWritableCableRequestStatus and assigns it to the TenantGroups field.
func (o *ConfigContextRequest) SetTenantGroups(v []BulkWritableCableRequestStatus) {
	o.TenantGroups = v
}

// GetTenants returns the Tenants field value if set, zero value otherwise.
func (o *ConfigContextRequest) GetTenants() []BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Tenants) {
		var ret []BulkWritableCableRequestStatus
		return ret
	}
	return o.Tenants
}

// GetTenantsOk returns a tuple with the Tenants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigContextRequest) GetTenantsOk() ([]BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Tenants) {
		return nil, false
	}
	return o.Tenants, true
}

// HasTenants returns a boolean if a field has been set.
func (o *ConfigContextRequest) HasTenants() bool {
	if o != nil && !IsNil(o.Tenants) {
		return true
	}

	return false
}

// SetTenants gets a reference to the given []BulkWritableCableRequestStatus and assigns it to the Tenants field.
func (o *ConfigContextRequest) SetTenants(v []BulkWritableCableRequestStatus) {
	o.Tenants = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ConfigContextRequest) GetTags() []BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Tags) {
		var ret []BulkWritableCableRequestStatus
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigContextRequest) GetTagsOk() ([]BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ConfigContextRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []BulkWritableCableRequestStatus and assigns it to the Tags field.
func (o *ConfigContextRequest) SetTags(v []BulkWritableCableRequestStatus) {
	o.Tags = v
}

func (o ConfigContextRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigContextRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.OwnerContentType.IsSet() {
		toSerialize["owner_content_type"] = o.OwnerContentType.Get()
	}
	toSerialize["name"] = o.Name
	if o.OwnerObjectId.IsSet() {
		toSerialize["owner_object_id"] = o.OwnerObjectId.Get()
	}
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.IsActive) {
		toSerialize["is_active"] = o.IsActive
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.ConfigContextSchema.IsSet() {
		toSerialize["config_context_schema"] = o.ConfigContextSchema.Get()
	}
	if !IsNil(o.Locations) {
		toSerialize["locations"] = o.Locations
	}
	if !IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	if !IsNil(o.DeviceTypes) {
		toSerialize["device_types"] = o.DeviceTypes
	}
	if !IsNil(o.DeviceRedundancyGroups) {
		toSerialize["device_redundancy_groups"] = o.DeviceRedundancyGroups
	}
	if !IsNil(o.Platforms) {
		toSerialize["platforms"] = o.Platforms
	}
	if !IsNil(o.ClusterGroups) {
		toSerialize["cluster_groups"] = o.ClusterGroups
	}
	if !IsNil(o.Clusters) {
		toSerialize["clusters"] = o.Clusters
	}
	if !IsNil(o.TenantGroups) {
		toSerialize["tenant_groups"] = o.TenantGroups
	}
	if !IsNil(o.Tenants) {
		toSerialize["tenants"] = o.Tenants
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConfigContextRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"data",
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

	varConfigContextRequest := _ConfigContextRequest{}

	err = json.Unmarshal(data, &varConfigContextRequest)

	if err != nil {
		return err
	}

	*o = ConfigContextRequest(varConfigContextRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "owner_content_type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "owner_object_id")
		delete(additionalProperties, "weight")
		delete(additionalProperties, "description")
		delete(additionalProperties, "is_active")
		delete(additionalProperties, "data")
		delete(additionalProperties, "config_context_schema")
		delete(additionalProperties, "locations")
		delete(additionalProperties, "roles")
		delete(additionalProperties, "device_types")
		delete(additionalProperties, "device_redundancy_groups")
		delete(additionalProperties, "platforms")
		delete(additionalProperties, "cluster_groups")
		delete(additionalProperties, "clusters")
		delete(additionalProperties, "tenant_groups")
		delete(additionalProperties, "tenants")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConfigContextRequest struct {
	value *ConfigContextRequest
	isSet bool
}

func (v NullableConfigContextRequest) Get() *ConfigContextRequest {
	return v.value
}

func (v *NullableConfigContextRequest) Set(val *ConfigContextRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigContextRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigContextRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigContextRequest(val *ConfigContextRequest) *NullableConfigContextRequest {
	return &NullableConfigContextRequest{value: val, isSet: true}
}

func (v NullableConfigContextRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigContextRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


