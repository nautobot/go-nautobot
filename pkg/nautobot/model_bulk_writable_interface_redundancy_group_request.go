/*
API Documentation

Source of truth and network automation platform

API version: 2.2.5 (2.2)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nautobot

import (
	"encoding/json"
	"fmt"
)

// checks if the BulkWritableInterfaceRedundancyGroupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkWritableInterfaceRedundancyGroupRequest{}

// BulkWritableInterfaceRedundancyGroupRequest InterfaceRedundancyGroup Serializer.
type BulkWritableInterfaceRedundancyGroupRequest struct {
	Id string `json:"id"`
	Protocol InterfaceRedundancyGroupProtocolChoices `json:"protocol"`
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	ProtocolGroupId *string `json:"protocol_group_id,omitempty"`
	Status BulkWritableCableRequestStatus `json:"status"`
	SecretsGroup NullableBulkWritableCircuitRequestTenant `json:"secrets_group,omitempty"`
	VirtualIp NullableBulkWritableCircuitRequestTenant `json:"virtual_ip,omitempty"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Relationships *map[string]BulkWritableCableRequestRelationshipsValue `json:"relationships,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkWritableInterfaceRedundancyGroupRequest BulkWritableInterfaceRedundancyGroupRequest

// NewBulkWritableInterfaceRedundancyGroupRequest instantiates a new BulkWritableInterfaceRedundancyGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkWritableInterfaceRedundancyGroupRequest(id string, protocol InterfaceRedundancyGroupProtocolChoices, name string, status BulkWritableCableRequestStatus) *BulkWritableInterfaceRedundancyGroupRequest {
	this := BulkWritableInterfaceRedundancyGroupRequest{}
	this.Id = id
	this.Protocol = protocol
	this.Name = name
	this.Status = status
	return &this
}

// NewBulkWritableInterfaceRedundancyGroupRequestWithDefaults instantiates a new BulkWritableInterfaceRedundancyGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkWritableInterfaceRedundancyGroupRequestWithDefaults() *BulkWritableInterfaceRedundancyGroupRequest {
	this := BulkWritableInterfaceRedundancyGroupRequest{}
	return &this
}

// GetId returns the Id field value
func (o *BulkWritableInterfaceRedundancyGroupRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BulkWritableInterfaceRedundancyGroupRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BulkWritableInterfaceRedundancyGroupRequest) SetId(v string) {
	o.Id = v
}

// GetProtocol returns the Protocol field value
func (o *BulkWritableInterfaceRedundancyGroupRequest) GetProtocol() InterfaceRedundancyGroupProtocolChoices {
	if o == nil {
		var ret InterfaceRedundancyGroupProtocolChoices
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *BulkWritableInterfaceRedundancyGroupRequest) GetProtocolOk() (*InterfaceRedundancyGroupProtocolChoices, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *BulkWritableInterfaceRedundancyGroupRequest) SetProtocol(v InterfaceRedundancyGroupProtocolChoices) {
	o.Protocol = v
}

// GetName returns the Name field value
func (o *BulkWritableInterfaceRedundancyGroupRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BulkWritableInterfaceRedundancyGroupRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BulkWritableInterfaceRedundancyGroupRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BulkWritableInterfaceRedundancyGroupRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableInterfaceRedundancyGroupRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BulkWritableInterfaceRedundancyGroupRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BulkWritableInterfaceRedundancyGroupRequest) SetDescription(v string) {
	o.Description = &v
}

// GetProtocolGroupId returns the ProtocolGroupId field value if set, zero value otherwise.
func (o *BulkWritableInterfaceRedundancyGroupRequest) GetProtocolGroupId() string {
	if o == nil || IsNil(o.ProtocolGroupId) {
		var ret string
		return ret
	}
	return *o.ProtocolGroupId
}

// GetProtocolGroupIdOk returns a tuple with the ProtocolGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableInterfaceRedundancyGroupRequest) GetProtocolGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProtocolGroupId) {
		return nil, false
	}
	return o.ProtocolGroupId, true
}

// HasProtocolGroupId returns a boolean if a field has been set.
func (o *BulkWritableInterfaceRedundancyGroupRequest) HasProtocolGroupId() bool {
	if o != nil && !IsNil(o.ProtocolGroupId) {
		return true
	}

	return false
}

// SetProtocolGroupId gets a reference to the given string and assigns it to the ProtocolGroupId field.
func (o *BulkWritableInterfaceRedundancyGroupRequest) SetProtocolGroupId(v string) {
	o.ProtocolGroupId = &v
}

// GetStatus returns the Status field value
func (o *BulkWritableInterfaceRedundancyGroupRequest) GetStatus() BulkWritableCableRequestStatus {
	if o == nil {
		var ret BulkWritableCableRequestStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *BulkWritableInterfaceRedundancyGroupRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *BulkWritableInterfaceRedundancyGroupRequest) SetStatus(v BulkWritableCableRequestStatus) {
	o.Status = v
}

// GetSecretsGroup returns the SecretsGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkWritableInterfaceRedundancyGroupRequest) GetSecretsGroup() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.SecretsGroup.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.SecretsGroup.Get()
}

// GetSecretsGroupOk returns a tuple with the SecretsGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkWritableInterfaceRedundancyGroupRequest) GetSecretsGroupOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.SecretsGroup.Get(), o.SecretsGroup.IsSet()
}

// HasSecretsGroup returns a boolean if a field has been set.
func (o *BulkWritableInterfaceRedundancyGroupRequest) HasSecretsGroup() bool {
	if o != nil && o.SecretsGroup.IsSet() {
		return true
	}

	return false
}

// SetSecretsGroup gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the SecretsGroup field.
func (o *BulkWritableInterfaceRedundancyGroupRequest) SetSecretsGroup(v BulkWritableCircuitRequestTenant) {
	o.SecretsGroup.Set(&v)
}
// SetSecretsGroupNil sets the value for SecretsGroup to be an explicit nil
func (o *BulkWritableInterfaceRedundancyGroupRequest) SetSecretsGroupNil() {
	o.SecretsGroup.Set(nil)
}

// UnsetSecretsGroup ensures that no value is present for SecretsGroup, not even an explicit nil
func (o *BulkWritableInterfaceRedundancyGroupRequest) UnsetSecretsGroup() {
	o.SecretsGroup.Unset()
}

// GetVirtualIp returns the VirtualIp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkWritableInterfaceRedundancyGroupRequest) GetVirtualIp() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.VirtualIp.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.VirtualIp.Get()
}

// GetVirtualIpOk returns a tuple with the VirtualIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkWritableInterfaceRedundancyGroupRequest) GetVirtualIpOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.VirtualIp.Get(), o.VirtualIp.IsSet()
}

// HasVirtualIp returns a boolean if a field has been set.
func (o *BulkWritableInterfaceRedundancyGroupRequest) HasVirtualIp() bool {
	if o != nil && o.VirtualIp.IsSet() {
		return true
	}

	return false
}

// SetVirtualIp gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the VirtualIp field.
func (o *BulkWritableInterfaceRedundancyGroupRequest) SetVirtualIp(v BulkWritableCircuitRequestTenant) {
	o.VirtualIp.Set(&v)
}
// SetVirtualIpNil sets the value for VirtualIp to be an explicit nil
func (o *BulkWritableInterfaceRedundancyGroupRequest) SetVirtualIpNil() {
	o.VirtualIp.Set(nil)
}

// UnsetVirtualIp ensures that no value is present for VirtualIp, not even an explicit nil
func (o *BulkWritableInterfaceRedundancyGroupRequest) UnsetVirtualIp() {
	o.VirtualIp.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *BulkWritableInterfaceRedundancyGroupRequest) GetTags() []BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Tags) {
		var ret []BulkWritableCableRequestStatus
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableInterfaceRedundancyGroupRequest) GetTagsOk() ([]BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *BulkWritableInterfaceRedundancyGroupRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []BulkWritableCableRequestStatus and assigns it to the Tags field.
func (o *BulkWritableInterfaceRedundancyGroupRequest) SetTags(v []BulkWritableCableRequestStatus) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *BulkWritableInterfaceRedundancyGroupRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableInterfaceRedundancyGroupRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *BulkWritableInterfaceRedundancyGroupRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *BulkWritableInterfaceRedundancyGroupRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *BulkWritableInterfaceRedundancyGroupRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue {
	if o == nil || IsNil(o.Relationships) {
		var ret map[string]BulkWritableCableRequestRelationshipsValue
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableInterfaceRedundancyGroupRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *BulkWritableInterfaceRedundancyGroupRequest) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]BulkWritableCableRequestRelationshipsValue and assigns it to the Relationships field.
func (o *BulkWritableInterfaceRedundancyGroupRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue) {
	o.Relationships = &v
}

func (o BulkWritableInterfaceRedundancyGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkWritableInterfaceRedundancyGroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["protocol"] = o.Protocol
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ProtocolGroupId) {
		toSerialize["protocol_group_id"] = o.ProtocolGroupId
	}
	toSerialize["status"] = o.Status
	if o.SecretsGroup.IsSet() {
		toSerialize["secrets_group"] = o.SecretsGroup.Get()
	}
	if o.VirtualIp.IsSet() {
		toSerialize["virtual_ip"] = o.VirtualIp.Get()
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

func (o *BulkWritableInterfaceRedundancyGroupRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"protocol",
		"name",
		"status",
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

	varBulkWritableInterfaceRedundancyGroupRequest := _BulkWritableInterfaceRedundancyGroupRequest{}

	err = json.Unmarshal(data, &varBulkWritableInterfaceRedundancyGroupRequest)

	if err != nil {
		return err
	}

	*o = BulkWritableInterfaceRedundancyGroupRequest(varBulkWritableInterfaceRedundancyGroupRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "protocol")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "protocol_group_id")
		delete(additionalProperties, "status")
		delete(additionalProperties, "secrets_group")
		delete(additionalProperties, "virtual_ip")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "relationships")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBulkWritableInterfaceRedundancyGroupRequest struct {
	value *BulkWritableInterfaceRedundancyGroupRequest
	isSet bool
}

func (v NullableBulkWritableInterfaceRedundancyGroupRequest) Get() *BulkWritableInterfaceRedundancyGroupRequest {
	return v.value
}

func (v *NullableBulkWritableInterfaceRedundancyGroupRequest) Set(val *BulkWritableInterfaceRedundancyGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkWritableInterfaceRedundancyGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkWritableInterfaceRedundancyGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkWritableInterfaceRedundancyGroupRequest(val *BulkWritableInterfaceRedundancyGroupRequest) *NullableBulkWritableInterfaceRedundancyGroupRequest {
	return &NullableBulkWritableInterfaceRedundancyGroupRequest{value: val, isSet: true}
}

func (v NullableBulkWritableInterfaceRedundancyGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkWritableInterfaceRedundancyGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

