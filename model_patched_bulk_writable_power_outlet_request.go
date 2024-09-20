/*
API Documentation

Source of truth and network automation platform

API version: 2.3.4 (2.3)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nautobot

import (
	"encoding/json"
	"fmt"
)

// checks if the PatchedBulkWritablePowerOutletRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedBulkWritablePowerOutletRequest{}

// PatchedBulkWritablePowerOutletRequest Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type PatchedBulkWritablePowerOutletRequest struct {
	Id string `json:"id"`
	Type *PowerOutletTypeChoices `json:"type,omitempty"`
	FeedLeg *FeedLegEnum `json:"feed_leg,omitempty"`
	Name *string `json:"name,omitempty"`
	// Physical label
	Label *string `json:"label,omitempty"`
	Description *string `json:"description,omitempty"`
	Device NullableBulkWritableCircuitRequestTenant `json:"device,omitempty"`
	Module NullableBulkWritableCircuitRequestTenant `json:"module,omitempty"`
	PowerPort NullableBulkWritableCircuitRequestTenant `json:"power_port,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Relationships *map[string]BulkWritableCableRequestRelationshipsValue `json:"relationships,omitempty"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedBulkWritablePowerOutletRequest PatchedBulkWritablePowerOutletRequest

// NewPatchedBulkWritablePowerOutletRequest instantiates a new PatchedBulkWritablePowerOutletRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedBulkWritablePowerOutletRequest(id string) *PatchedBulkWritablePowerOutletRequest {
	this := PatchedBulkWritablePowerOutletRequest{}
	this.Id = id
	return &this
}

// NewPatchedBulkWritablePowerOutletRequestWithDefaults instantiates a new PatchedBulkWritablePowerOutletRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedBulkWritablePowerOutletRequestWithDefaults() *PatchedBulkWritablePowerOutletRequest {
	this := PatchedBulkWritablePowerOutletRequest{}
	return &this
}

// GetId returns the Id field value
func (o *PatchedBulkWritablePowerOutletRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritablePowerOutletRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PatchedBulkWritablePowerOutletRequest) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PatchedBulkWritablePowerOutletRequest) GetType() PowerOutletTypeChoices {
	if o == nil || IsNil(o.Type) {
		var ret PowerOutletTypeChoices
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritablePowerOutletRequest) GetTypeOk() (*PowerOutletTypeChoices, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PatchedBulkWritablePowerOutletRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given PowerOutletTypeChoices and assigns it to the Type field.
func (o *PatchedBulkWritablePowerOutletRequest) SetType(v PowerOutletTypeChoices) {
	o.Type = &v
}

// GetFeedLeg returns the FeedLeg field value if set, zero value otherwise.
func (o *PatchedBulkWritablePowerOutletRequest) GetFeedLeg() FeedLegEnum {
	if o == nil || IsNil(o.FeedLeg) {
		var ret FeedLegEnum
		return ret
	}
	return *o.FeedLeg
}

// GetFeedLegOk returns a tuple with the FeedLeg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritablePowerOutletRequest) GetFeedLegOk() (*FeedLegEnum, bool) {
	if o == nil || IsNil(o.FeedLeg) {
		return nil, false
	}
	return o.FeedLeg, true
}

// HasFeedLeg returns a boolean if a field has been set.
func (o *PatchedBulkWritablePowerOutletRequest) HasFeedLeg() bool {
	if o != nil && !IsNil(o.FeedLeg) {
		return true
	}

	return false
}

// SetFeedLeg gets a reference to the given FeedLegEnum and assigns it to the FeedLeg field.
func (o *PatchedBulkWritablePowerOutletRequest) SetFeedLeg(v FeedLegEnum) {
	o.FeedLeg = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedBulkWritablePowerOutletRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritablePowerOutletRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedBulkWritablePowerOutletRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedBulkWritablePowerOutletRequest) SetName(v string) {
	o.Name = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PatchedBulkWritablePowerOutletRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritablePowerOutletRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PatchedBulkWritablePowerOutletRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PatchedBulkWritablePowerOutletRequest) SetLabel(v string) {
	o.Label = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedBulkWritablePowerOutletRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritablePowerOutletRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedBulkWritablePowerOutletRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedBulkWritablePowerOutletRequest) SetDescription(v string) {
	o.Description = &v
}

// GetDevice returns the Device field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBulkWritablePowerOutletRequest) GetDevice() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.Device.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.Device.Get()
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBulkWritablePowerOutletRequest) GetDeviceOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Device.Get(), o.Device.IsSet()
}

// HasDevice returns a boolean if a field has been set.
func (o *PatchedBulkWritablePowerOutletRequest) HasDevice() bool {
	if o != nil && o.Device.IsSet() {
		return true
	}

	return false
}

// SetDevice gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the Device field.
func (o *PatchedBulkWritablePowerOutletRequest) SetDevice(v BulkWritableCircuitRequestTenant) {
	o.Device.Set(&v)
}
// SetDeviceNil sets the value for Device to be an explicit nil
func (o *PatchedBulkWritablePowerOutletRequest) SetDeviceNil() {
	o.Device.Set(nil)
}

// UnsetDevice ensures that no value is present for Device, not even an explicit nil
func (o *PatchedBulkWritablePowerOutletRequest) UnsetDevice() {
	o.Device.Unset()
}

// GetModule returns the Module field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBulkWritablePowerOutletRequest) GetModule() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.Module.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.Module.Get()
}

// GetModuleOk returns a tuple with the Module field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBulkWritablePowerOutletRequest) GetModuleOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Module.Get(), o.Module.IsSet()
}

// HasModule returns a boolean if a field has been set.
func (o *PatchedBulkWritablePowerOutletRequest) HasModule() bool {
	if o != nil && o.Module.IsSet() {
		return true
	}

	return false
}

// SetModule gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the Module field.
func (o *PatchedBulkWritablePowerOutletRequest) SetModule(v BulkWritableCircuitRequestTenant) {
	o.Module.Set(&v)
}
// SetModuleNil sets the value for Module to be an explicit nil
func (o *PatchedBulkWritablePowerOutletRequest) SetModuleNil() {
	o.Module.Set(nil)
}

// UnsetModule ensures that no value is present for Module, not even an explicit nil
func (o *PatchedBulkWritablePowerOutletRequest) UnsetModule() {
	o.Module.Unset()
}

// GetPowerPort returns the PowerPort field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBulkWritablePowerOutletRequest) GetPowerPort() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.PowerPort.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.PowerPort.Get()
}

// GetPowerPortOk returns a tuple with the PowerPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBulkWritablePowerOutletRequest) GetPowerPortOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.PowerPort.Get(), o.PowerPort.IsSet()
}

// HasPowerPort returns a boolean if a field has been set.
func (o *PatchedBulkWritablePowerOutletRequest) HasPowerPort() bool {
	if o != nil && o.PowerPort.IsSet() {
		return true
	}

	return false
}

// SetPowerPort gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the PowerPort field.
func (o *PatchedBulkWritablePowerOutletRequest) SetPowerPort(v BulkWritableCircuitRequestTenant) {
	o.PowerPort.Set(&v)
}
// SetPowerPortNil sets the value for PowerPort to be an explicit nil
func (o *PatchedBulkWritablePowerOutletRequest) SetPowerPortNil() {
	o.PowerPort.Set(nil)
}

// UnsetPowerPort ensures that no value is present for PowerPort, not even an explicit nil
func (o *PatchedBulkWritablePowerOutletRequest) UnsetPowerPort() {
	o.PowerPort.Unset()
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedBulkWritablePowerOutletRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritablePowerOutletRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedBulkWritablePowerOutletRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedBulkWritablePowerOutletRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PatchedBulkWritablePowerOutletRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue {
	if o == nil || IsNil(o.Relationships) {
		var ret map[string]BulkWritableCableRequestRelationshipsValue
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritablePowerOutletRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PatchedBulkWritablePowerOutletRequest) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]BulkWritableCableRequestRelationshipsValue and assigns it to the Relationships field.
func (o *PatchedBulkWritablePowerOutletRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue) {
	o.Relationships = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedBulkWritablePowerOutletRequest) GetTags() []BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Tags) {
		var ret []BulkWritableCableRequestStatus
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritablePowerOutletRequest) GetTagsOk() ([]BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedBulkWritablePowerOutletRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []BulkWritableCableRequestStatus and assigns it to the Tags field.
func (o *PatchedBulkWritablePowerOutletRequest) SetTags(v []BulkWritableCableRequestStatus) {
	o.Tags = v
}

func (o PatchedBulkWritablePowerOutletRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedBulkWritablePowerOutletRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.FeedLeg) {
		toSerialize["feed_leg"] = o.FeedLeg
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.Device.IsSet() {
		toSerialize["device"] = o.Device.Get()
	}
	if o.Module.IsSet() {
		toSerialize["module"] = o.Module.Get()
	}
	if o.PowerPort.IsSet() {
		toSerialize["power_port"] = o.PowerPort.Get()
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

func (o *PatchedBulkWritablePowerOutletRequest) UnmarshalJSON(data []byte) (err error) {
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

	varPatchedBulkWritablePowerOutletRequest := _PatchedBulkWritablePowerOutletRequest{}

	err = json.Unmarshal(data, &varPatchedBulkWritablePowerOutletRequest)

	if err != nil {
		return err
	}

	*o = PatchedBulkWritablePowerOutletRequest(varPatchedBulkWritablePowerOutletRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "feed_leg")
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "description")
		delete(additionalProperties, "device")
		delete(additionalProperties, "module")
		delete(additionalProperties, "power_port")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "relationships")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedBulkWritablePowerOutletRequest struct {
	value *PatchedBulkWritablePowerOutletRequest
	isSet bool
}

func (v NullablePatchedBulkWritablePowerOutletRequest) Get() *PatchedBulkWritablePowerOutletRequest {
	return v.value
}

func (v *NullablePatchedBulkWritablePowerOutletRequest) Set(val *PatchedBulkWritablePowerOutletRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedBulkWritablePowerOutletRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedBulkWritablePowerOutletRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedBulkWritablePowerOutletRequest(val *PatchedBulkWritablePowerOutletRequest) *NullablePatchedBulkWritablePowerOutletRequest {
	return &NullablePatchedBulkWritablePowerOutletRequest{value: val, isSet: true}
}

func (v NullablePatchedBulkWritablePowerOutletRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedBulkWritablePowerOutletRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


