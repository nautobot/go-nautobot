/*
API Documentation

Source of truth and network automation platform

API version: 2.3.4 (2.3)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nautobot

import (
	"encoding/json"
)

// checks if the PatchedWritableInterfaceRedundancyGroupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedWritableInterfaceRedundancyGroupRequest{}

// PatchedWritableInterfaceRedundancyGroupRequest InterfaceRedundancyGroup Serializer.
type PatchedWritableInterfaceRedundancyGroupRequest struct {
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Protocol *RedundancyProtocol `json:"protocol,omitempty"`
	ProtocolGroupId *string `json:"protocol_group_id,omitempty"`
	Status *BulkWritableCableRequestStatus `json:"status,omitempty"`
	SecretsGroup NullableBulkWritableCircuitRequestTenant `json:"secrets_group,omitempty"`
	VirtualIp NullableBulkWritableCircuitRequestTenant `json:"virtual_ip,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Relationships *map[string]BulkWritableCableRequestRelationshipsValue `json:"relationships,omitempty"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedWritableInterfaceRedundancyGroupRequest PatchedWritableInterfaceRedundancyGroupRequest

// NewPatchedWritableInterfaceRedundancyGroupRequest instantiates a new PatchedWritableInterfaceRedundancyGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritableInterfaceRedundancyGroupRequest() *PatchedWritableInterfaceRedundancyGroupRequest {
	this := PatchedWritableInterfaceRedundancyGroupRequest{}
	return &this
}

// NewPatchedWritableInterfaceRedundancyGroupRequestWithDefaults instantiates a new PatchedWritableInterfaceRedundancyGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritableInterfaceRedundancyGroupRequestWithDefaults() *PatchedWritableInterfaceRedundancyGroupRequest {
	this := PatchedWritableInterfaceRedundancyGroupRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedWritableInterfaceRedundancyGroupRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableInterfaceRedundancyGroupRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedWritableInterfaceRedundancyGroupRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedWritableInterfaceRedundancyGroupRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedWritableInterfaceRedundancyGroupRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableInterfaceRedundancyGroupRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedWritableInterfaceRedundancyGroupRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedWritableInterfaceRedundancyGroupRequest) SetDescription(v string) {
	o.Description = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *PatchedWritableInterfaceRedundancyGroupRequest) GetProtocol() RedundancyProtocol {
	if o == nil || IsNil(o.Protocol) {
		var ret RedundancyProtocol
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableInterfaceRedundancyGroupRequest) GetProtocolOk() (*RedundancyProtocol, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *PatchedWritableInterfaceRedundancyGroupRequest) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given RedundancyProtocol and assigns it to the Protocol field.
func (o *PatchedWritableInterfaceRedundancyGroupRequest) SetProtocol(v RedundancyProtocol) {
	o.Protocol = &v
}

// GetProtocolGroupId returns the ProtocolGroupId field value if set, zero value otherwise.
func (o *PatchedWritableInterfaceRedundancyGroupRequest) GetProtocolGroupId() string {
	if o == nil || IsNil(o.ProtocolGroupId) {
		var ret string
		return ret
	}
	return *o.ProtocolGroupId
}

// GetProtocolGroupIdOk returns a tuple with the ProtocolGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableInterfaceRedundancyGroupRequest) GetProtocolGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProtocolGroupId) {
		return nil, false
	}
	return o.ProtocolGroupId, true
}

// HasProtocolGroupId returns a boolean if a field has been set.
func (o *PatchedWritableInterfaceRedundancyGroupRequest) HasProtocolGroupId() bool {
	if o != nil && !IsNil(o.ProtocolGroupId) {
		return true
	}

	return false
}

// SetProtocolGroupId gets a reference to the given string and assigns it to the ProtocolGroupId field.
func (o *PatchedWritableInterfaceRedundancyGroupRequest) SetProtocolGroupId(v string) {
	o.ProtocolGroupId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PatchedWritableInterfaceRedundancyGroupRequest) GetStatus() BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Status) {
		var ret BulkWritableCableRequestStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableInterfaceRedundancyGroupRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PatchedWritableInterfaceRedundancyGroupRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given BulkWritableCableRequestStatus and assigns it to the Status field.
func (o *PatchedWritableInterfaceRedundancyGroupRequest) SetStatus(v BulkWritableCableRequestStatus) {
	o.Status = &v
}

// GetSecretsGroup returns the SecretsGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableInterfaceRedundancyGroupRequest) GetSecretsGroup() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.SecretsGroup.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.SecretsGroup.Get()
}

// GetSecretsGroupOk returns a tuple with the SecretsGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableInterfaceRedundancyGroupRequest) GetSecretsGroupOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.SecretsGroup.Get(), o.SecretsGroup.IsSet()
}

// HasSecretsGroup returns a boolean if a field has been set.
func (o *PatchedWritableInterfaceRedundancyGroupRequest) HasSecretsGroup() bool {
	if o != nil && o.SecretsGroup.IsSet() {
		return true
	}

	return false
}

// SetSecretsGroup gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the SecretsGroup field.
func (o *PatchedWritableInterfaceRedundancyGroupRequest) SetSecretsGroup(v BulkWritableCircuitRequestTenant) {
	o.SecretsGroup.Set(&v)
}
// SetSecretsGroupNil sets the value for SecretsGroup to be an explicit nil
func (o *PatchedWritableInterfaceRedundancyGroupRequest) SetSecretsGroupNil() {
	o.SecretsGroup.Set(nil)
}

// UnsetSecretsGroup ensures that no value is present for SecretsGroup, not even an explicit nil
func (o *PatchedWritableInterfaceRedundancyGroupRequest) UnsetSecretsGroup() {
	o.SecretsGroup.Unset()
}

// GetVirtualIp returns the VirtualIp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableInterfaceRedundancyGroupRequest) GetVirtualIp() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.VirtualIp.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.VirtualIp.Get()
}

// GetVirtualIpOk returns a tuple with the VirtualIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableInterfaceRedundancyGroupRequest) GetVirtualIpOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.VirtualIp.Get(), o.VirtualIp.IsSet()
}

// HasVirtualIp returns a boolean if a field has been set.
func (o *PatchedWritableInterfaceRedundancyGroupRequest) HasVirtualIp() bool {
	if o != nil && o.VirtualIp.IsSet() {
		return true
	}

	return false
}

// SetVirtualIp gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the VirtualIp field.
func (o *PatchedWritableInterfaceRedundancyGroupRequest) SetVirtualIp(v BulkWritableCircuitRequestTenant) {
	o.VirtualIp.Set(&v)
}
// SetVirtualIpNil sets the value for VirtualIp to be an explicit nil
func (o *PatchedWritableInterfaceRedundancyGroupRequest) SetVirtualIpNil() {
	o.VirtualIp.Set(nil)
}

// UnsetVirtualIp ensures that no value is present for VirtualIp, not even an explicit nil
func (o *PatchedWritableInterfaceRedundancyGroupRequest) UnsetVirtualIp() {
	o.VirtualIp.Unset()
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedWritableInterfaceRedundancyGroupRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableInterfaceRedundancyGroupRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedWritableInterfaceRedundancyGroupRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedWritableInterfaceRedundancyGroupRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PatchedWritableInterfaceRedundancyGroupRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue {
	if o == nil || IsNil(o.Relationships) {
		var ret map[string]BulkWritableCableRequestRelationshipsValue
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableInterfaceRedundancyGroupRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PatchedWritableInterfaceRedundancyGroupRequest) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]BulkWritableCableRequestRelationshipsValue and assigns it to the Relationships field.
func (o *PatchedWritableInterfaceRedundancyGroupRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue) {
	o.Relationships = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedWritableInterfaceRedundancyGroupRequest) GetTags() []BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Tags) {
		var ret []BulkWritableCableRequestStatus
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableInterfaceRedundancyGroupRequest) GetTagsOk() ([]BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedWritableInterfaceRedundancyGroupRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []BulkWritableCableRequestStatus and assigns it to the Tags field.
func (o *PatchedWritableInterfaceRedundancyGroupRequest) SetTags(v []BulkWritableCableRequestStatus) {
	o.Tags = v
}

func (o PatchedWritableInterfaceRedundancyGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedWritableInterfaceRedundancyGroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !IsNil(o.ProtocolGroupId) {
		toSerialize["protocol_group_id"] = o.ProtocolGroupId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.SecretsGroup.IsSet() {
		toSerialize["secrets_group"] = o.SecretsGroup.Get()
	}
	if o.VirtualIp.IsSet() {
		toSerialize["virtual_ip"] = o.VirtualIp.Get()
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

func (o *PatchedWritableInterfaceRedundancyGroupRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedWritableInterfaceRedundancyGroupRequest := _PatchedWritableInterfaceRedundancyGroupRequest{}

	err = json.Unmarshal(data, &varPatchedWritableInterfaceRedundancyGroupRequest)

	if err != nil {
		return err
	}

	*o = PatchedWritableInterfaceRedundancyGroupRequest(varPatchedWritableInterfaceRedundancyGroupRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "protocol")
		delete(additionalProperties, "protocol_group_id")
		delete(additionalProperties, "status")
		delete(additionalProperties, "secrets_group")
		delete(additionalProperties, "virtual_ip")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "relationships")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedWritableInterfaceRedundancyGroupRequest struct {
	value *PatchedWritableInterfaceRedundancyGroupRequest
	isSet bool
}

func (v NullablePatchedWritableInterfaceRedundancyGroupRequest) Get() *PatchedWritableInterfaceRedundancyGroupRequest {
	return v.value
}

func (v *NullablePatchedWritableInterfaceRedundancyGroupRequest) Set(val *PatchedWritableInterfaceRedundancyGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableInterfaceRedundancyGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableInterfaceRedundancyGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableInterfaceRedundancyGroupRequest(val *PatchedWritableInterfaceRedundancyGroupRequest) *NullablePatchedWritableInterfaceRedundancyGroupRequest {
	return &NullablePatchedWritableInterfaceRedundancyGroupRequest{value: val, isSet: true}
}

func (v NullablePatchedWritableInterfaceRedundancyGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableInterfaceRedundancyGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


