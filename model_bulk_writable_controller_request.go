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

// checks if the BulkWritableControllerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkWritableControllerRequest{}

// BulkWritableControllerRequest Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type BulkWritableControllerRequest struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	Status BulkWritableCableRequestStatus `json:"status"`
	Location BulkWritableCableRequestStatus `json:"location"`
	Platform NullableBulkWritableCircuitRequestTenant `json:"platform,omitempty"`
	Role NullableBulkWritableCircuitRequestTenant `json:"role,omitempty"`
	Tenant NullableBulkWritableCircuitRequestTenant `json:"tenant,omitempty"`
	ExternalIntegration NullableBulkWritableCircuitRequestTenant `json:"external_integration,omitempty"`
	ControllerDevice NullableBulkWritableCircuitRequestTenant `json:"controller_device,omitempty"`
	ControllerDeviceRedundancyGroup NullableBulkWritableCircuitRequestTenant `json:"controller_device_redundancy_group,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Relationships *map[string]BulkWritableCableRequestRelationshipsValue `json:"relationships,omitempty"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkWritableControllerRequest BulkWritableControllerRequest

// NewBulkWritableControllerRequest instantiates a new BulkWritableControllerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkWritableControllerRequest(id string, name string, status BulkWritableCableRequestStatus, location BulkWritableCableRequestStatus) *BulkWritableControllerRequest {
	this := BulkWritableControllerRequest{}
	this.Id = id
	this.Name = name
	this.Status = status
	this.Location = location
	return &this
}

// NewBulkWritableControllerRequestWithDefaults instantiates a new BulkWritableControllerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkWritableControllerRequestWithDefaults() *BulkWritableControllerRequest {
	this := BulkWritableControllerRequest{}
	return &this
}

// GetId returns the Id field value
func (o *BulkWritableControllerRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BulkWritableControllerRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BulkWritableControllerRequest) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *BulkWritableControllerRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BulkWritableControllerRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BulkWritableControllerRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BulkWritableControllerRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableControllerRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BulkWritableControllerRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BulkWritableControllerRequest) SetDescription(v string) {
	o.Description = &v
}

// GetStatus returns the Status field value
func (o *BulkWritableControllerRequest) GetStatus() BulkWritableCableRequestStatus {
	if o == nil {
		var ret BulkWritableCableRequestStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *BulkWritableControllerRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *BulkWritableControllerRequest) SetStatus(v BulkWritableCableRequestStatus) {
	o.Status = v
}

// GetLocation returns the Location field value
func (o *BulkWritableControllerRequest) GetLocation() BulkWritableCableRequestStatus {
	if o == nil {
		var ret BulkWritableCableRequestStatus
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *BulkWritableControllerRequest) GetLocationOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *BulkWritableControllerRequest) SetLocation(v BulkWritableCableRequestStatus) {
	o.Location = v
}

// GetPlatform returns the Platform field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkWritableControllerRequest) GetPlatform() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.Platform.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.Platform.Get()
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkWritableControllerRequest) GetPlatformOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Platform.Get(), o.Platform.IsSet()
}

// HasPlatform returns a boolean if a field has been set.
func (o *BulkWritableControllerRequest) HasPlatform() bool {
	if o != nil && o.Platform.IsSet() {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the Platform field.
func (o *BulkWritableControllerRequest) SetPlatform(v BulkWritableCircuitRequestTenant) {
	o.Platform.Set(&v)
}
// SetPlatformNil sets the value for Platform to be an explicit nil
func (o *BulkWritableControllerRequest) SetPlatformNil() {
	o.Platform.Set(nil)
}

// UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
func (o *BulkWritableControllerRequest) UnsetPlatform() {
	o.Platform.Unset()
}

// GetRole returns the Role field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkWritableControllerRequest) GetRole() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.Role.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.Role.Get()
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkWritableControllerRequest) GetRoleOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Role.Get(), o.Role.IsSet()
}

// HasRole returns a boolean if a field has been set.
func (o *BulkWritableControllerRequest) HasRole() bool {
	if o != nil && o.Role.IsSet() {
		return true
	}

	return false
}

// SetRole gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the Role field.
func (o *BulkWritableControllerRequest) SetRole(v BulkWritableCircuitRequestTenant) {
	o.Role.Set(&v)
}
// SetRoleNil sets the value for Role to be an explicit nil
func (o *BulkWritableControllerRequest) SetRoleNil() {
	o.Role.Set(nil)
}

// UnsetRole ensures that no value is present for Role, not even an explicit nil
func (o *BulkWritableControllerRequest) UnsetRole() {
	o.Role.Unset()
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkWritableControllerRequest) GetTenant() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkWritableControllerRequest) GetTenantOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *BulkWritableControllerRequest) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the Tenant field.
func (o *BulkWritableControllerRequest) SetTenant(v BulkWritableCircuitRequestTenant) {
	o.Tenant.Set(&v)
}
// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *BulkWritableControllerRequest) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *BulkWritableControllerRequest) UnsetTenant() {
	o.Tenant.Unset()
}

// GetExternalIntegration returns the ExternalIntegration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkWritableControllerRequest) GetExternalIntegration() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.ExternalIntegration.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.ExternalIntegration.Get()
}

// GetExternalIntegrationOk returns a tuple with the ExternalIntegration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkWritableControllerRequest) GetExternalIntegrationOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalIntegration.Get(), o.ExternalIntegration.IsSet()
}

// HasExternalIntegration returns a boolean if a field has been set.
func (o *BulkWritableControllerRequest) HasExternalIntegration() bool {
	if o != nil && o.ExternalIntegration.IsSet() {
		return true
	}

	return false
}

// SetExternalIntegration gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the ExternalIntegration field.
func (o *BulkWritableControllerRequest) SetExternalIntegration(v BulkWritableCircuitRequestTenant) {
	o.ExternalIntegration.Set(&v)
}
// SetExternalIntegrationNil sets the value for ExternalIntegration to be an explicit nil
func (o *BulkWritableControllerRequest) SetExternalIntegrationNil() {
	o.ExternalIntegration.Set(nil)
}

// UnsetExternalIntegration ensures that no value is present for ExternalIntegration, not even an explicit nil
func (o *BulkWritableControllerRequest) UnsetExternalIntegration() {
	o.ExternalIntegration.Unset()
}

// GetControllerDevice returns the ControllerDevice field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkWritableControllerRequest) GetControllerDevice() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.ControllerDevice.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.ControllerDevice.Get()
}

// GetControllerDeviceOk returns a tuple with the ControllerDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkWritableControllerRequest) GetControllerDeviceOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.ControllerDevice.Get(), o.ControllerDevice.IsSet()
}

// HasControllerDevice returns a boolean if a field has been set.
func (o *BulkWritableControllerRequest) HasControllerDevice() bool {
	if o != nil && o.ControllerDevice.IsSet() {
		return true
	}

	return false
}

// SetControllerDevice gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the ControllerDevice field.
func (o *BulkWritableControllerRequest) SetControllerDevice(v BulkWritableCircuitRequestTenant) {
	o.ControllerDevice.Set(&v)
}
// SetControllerDeviceNil sets the value for ControllerDevice to be an explicit nil
func (o *BulkWritableControllerRequest) SetControllerDeviceNil() {
	o.ControllerDevice.Set(nil)
}

// UnsetControllerDevice ensures that no value is present for ControllerDevice, not even an explicit nil
func (o *BulkWritableControllerRequest) UnsetControllerDevice() {
	o.ControllerDevice.Unset()
}

// GetControllerDeviceRedundancyGroup returns the ControllerDeviceRedundancyGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkWritableControllerRequest) GetControllerDeviceRedundancyGroup() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.ControllerDeviceRedundancyGroup.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.ControllerDeviceRedundancyGroup.Get()
}

// GetControllerDeviceRedundancyGroupOk returns a tuple with the ControllerDeviceRedundancyGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkWritableControllerRequest) GetControllerDeviceRedundancyGroupOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.ControllerDeviceRedundancyGroup.Get(), o.ControllerDeviceRedundancyGroup.IsSet()
}

// HasControllerDeviceRedundancyGroup returns a boolean if a field has been set.
func (o *BulkWritableControllerRequest) HasControllerDeviceRedundancyGroup() bool {
	if o != nil && o.ControllerDeviceRedundancyGroup.IsSet() {
		return true
	}

	return false
}

// SetControllerDeviceRedundancyGroup gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the ControllerDeviceRedundancyGroup field.
func (o *BulkWritableControllerRequest) SetControllerDeviceRedundancyGroup(v BulkWritableCircuitRequestTenant) {
	o.ControllerDeviceRedundancyGroup.Set(&v)
}
// SetControllerDeviceRedundancyGroupNil sets the value for ControllerDeviceRedundancyGroup to be an explicit nil
func (o *BulkWritableControllerRequest) SetControllerDeviceRedundancyGroupNil() {
	o.ControllerDeviceRedundancyGroup.Set(nil)
}

// UnsetControllerDeviceRedundancyGroup ensures that no value is present for ControllerDeviceRedundancyGroup, not even an explicit nil
func (o *BulkWritableControllerRequest) UnsetControllerDeviceRedundancyGroup() {
	o.ControllerDeviceRedundancyGroup.Unset()
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *BulkWritableControllerRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableControllerRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *BulkWritableControllerRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *BulkWritableControllerRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *BulkWritableControllerRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue {
	if o == nil || IsNil(o.Relationships) {
		var ret map[string]BulkWritableCableRequestRelationshipsValue
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableControllerRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *BulkWritableControllerRequest) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]BulkWritableCableRequestRelationshipsValue and assigns it to the Relationships field.
func (o *BulkWritableControllerRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue) {
	o.Relationships = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *BulkWritableControllerRequest) GetTags() []BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Tags) {
		var ret []BulkWritableCableRequestStatus
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableControllerRequest) GetTagsOk() ([]BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *BulkWritableControllerRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []BulkWritableCableRequestStatus and assigns it to the Tags field.
func (o *BulkWritableControllerRequest) SetTags(v []BulkWritableCableRequestStatus) {
	o.Tags = v
}

func (o BulkWritableControllerRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkWritableControllerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["status"] = o.Status
	toSerialize["location"] = o.Location
	if o.Platform.IsSet() {
		toSerialize["platform"] = o.Platform.Get()
	}
	if o.Role.IsSet() {
		toSerialize["role"] = o.Role.Get()
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if o.ExternalIntegration.IsSet() {
		toSerialize["external_integration"] = o.ExternalIntegration.Get()
	}
	if o.ControllerDevice.IsSet() {
		toSerialize["controller_device"] = o.ControllerDevice.Get()
	}
	if o.ControllerDeviceRedundancyGroup.IsSet() {
		toSerialize["controller_device_redundancy_group"] = o.ControllerDeviceRedundancyGroup.Get()
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BulkWritableControllerRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"status",
		"location",
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

	varBulkWritableControllerRequest := _BulkWritableControllerRequest{}

	err = json.Unmarshal(data, &varBulkWritableControllerRequest)

	if err != nil {
		return err
	}

	*o = BulkWritableControllerRequest(varBulkWritableControllerRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "status")
		delete(additionalProperties, "location")
		delete(additionalProperties, "platform")
		delete(additionalProperties, "role")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "external_integration")
		delete(additionalProperties, "controller_device")
		delete(additionalProperties, "controller_device_redundancy_group")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "relationships")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBulkWritableControllerRequest struct {
	value *BulkWritableControllerRequest
	isSet bool
}

func (v NullableBulkWritableControllerRequest) Get() *BulkWritableControllerRequest {
	return v.value
}

func (v *NullableBulkWritableControllerRequest) Set(val *BulkWritableControllerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkWritableControllerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkWritableControllerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkWritableControllerRequest(val *BulkWritableControllerRequest) *NullableBulkWritableControllerRequest {
	return &NullableBulkWritableControllerRequest{value: val, isSet: true}
}

func (v NullableBulkWritableControllerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkWritableControllerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


