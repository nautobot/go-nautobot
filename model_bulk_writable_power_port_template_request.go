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

// checks if the BulkWritablePowerPortTemplateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkWritablePowerPortTemplateRequest{}

// BulkWritablePowerPortTemplateRequest Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type BulkWritablePowerPortTemplateRequest struct {
	Id string `json:"id"`
	Type *PowerPortTypeChoices `json:"type,omitempty"`
	Name string `json:"name"`
	// Physical label
	Label *string `json:"label,omitempty"`
	Description *string `json:"description,omitempty"`
	// Maximum power draw (watts)
	MaximumDraw NullableInt32 `json:"maximum_draw,omitempty"`
	// Allocated power draw (watts)
	AllocatedDraw NullableInt32 `json:"allocated_draw,omitempty"`
	DeviceType NullableBulkWritableCircuitRequestTenant `json:"device_type,omitempty"`
	ModuleType NullableBulkWritableCircuitRequestTenant `json:"module_type,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Relationships *map[string]BulkWritableCableRequestRelationshipsValue `json:"relationships,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkWritablePowerPortTemplateRequest BulkWritablePowerPortTemplateRequest

// NewBulkWritablePowerPortTemplateRequest instantiates a new BulkWritablePowerPortTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkWritablePowerPortTemplateRequest(id string, name string) *BulkWritablePowerPortTemplateRequest {
	this := BulkWritablePowerPortTemplateRequest{}
	this.Id = id
	this.Name = name
	return &this
}

// NewBulkWritablePowerPortTemplateRequestWithDefaults instantiates a new BulkWritablePowerPortTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkWritablePowerPortTemplateRequestWithDefaults() *BulkWritablePowerPortTemplateRequest {
	this := BulkWritablePowerPortTemplateRequest{}
	return &this
}

// GetId returns the Id field value
func (o *BulkWritablePowerPortTemplateRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BulkWritablePowerPortTemplateRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BulkWritablePowerPortTemplateRequest) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BulkWritablePowerPortTemplateRequest) GetType() PowerPortTypeChoices {
	if o == nil || IsNil(o.Type) {
		var ret PowerPortTypeChoices
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritablePowerPortTemplateRequest) GetTypeOk() (*PowerPortTypeChoices, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BulkWritablePowerPortTemplateRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given PowerPortTypeChoices and assigns it to the Type field.
func (o *BulkWritablePowerPortTemplateRequest) SetType(v PowerPortTypeChoices) {
	o.Type = &v
}

// GetName returns the Name field value
func (o *BulkWritablePowerPortTemplateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BulkWritablePowerPortTemplateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BulkWritablePowerPortTemplateRequest) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *BulkWritablePowerPortTemplateRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritablePowerPortTemplateRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *BulkWritablePowerPortTemplateRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *BulkWritablePowerPortTemplateRequest) SetLabel(v string) {
	o.Label = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BulkWritablePowerPortTemplateRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritablePowerPortTemplateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BulkWritablePowerPortTemplateRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BulkWritablePowerPortTemplateRequest) SetDescription(v string) {
	o.Description = &v
}

// GetMaximumDraw returns the MaximumDraw field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkWritablePowerPortTemplateRequest) GetMaximumDraw() int32 {
	if o == nil || IsNil(o.MaximumDraw.Get()) {
		var ret int32
		return ret
	}
	return *o.MaximumDraw.Get()
}

// GetMaximumDrawOk returns a tuple with the MaximumDraw field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkWritablePowerPortTemplateRequest) GetMaximumDrawOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaximumDraw.Get(), o.MaximumDraw.IsSet()
}

// HasMaximumDraw returns a boolean if a field has been set.
func (o *BulkWritablePowerPortTemplateRequest) HasMaximumDraw() bool {
	if o != nil && o.MaximumDraw.IsSet() {
		return true
	}

	return false
}

// SetMaximumDraw gets a reference to the given NullableInt32 and assigns it to the MaximumDraw field.
func (o *BulkWritablePowerPortTemplateRequest) SetMaximumDraw(v int32) {
	o.MaximumDraw.Set(&v)
}
// SetMaximumDrawNil sets the value for MaximumDraw to be an explicit nil
func (o *BulkWritablePowerPortTemplateRequest) SetMaximumDrawNil() {
	o.MaximumDraw.Set(nil)
}

// UnsetMaximumDraw ensures that no value is present for MaximumDraw, not even an explicit nil
func (o *BulkWritablePowerPortTemplateRequest) UnsetMaximumDraw() {
	o.MaximumDraw.Unset()
}

// GetAllocatedDraw returns the AllocatedDraw field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkWritablePowerPortTemplateRequest) GetAllocatedDraw() int32 {
	if o == nil || IsNil(o.AllocatedDraw.Get()) {
		var ret int32
		return ret
	}
	return *o.AllocatedDraw.Get()
}

// GetAllocatedDrawOk returns a tuple with the AllocatedDraw field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkWritablePowerPortTemplateRequest) GetAllocatedDrawOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllocatedDraw.Get(), o.AllocatedDraw.IsSet()
}

// HasAllocatedDraw returns a boolean if a field has been set.
func (o *BulkWritablePowerPortTemplateRequest) HasAllocatedDraw() bool {
	if o != nil && o.AllocatedDraw.IsSet() {
		return true
	}

	return false
}

// SetAllocatedDraw gets a reference to the given NullableInt32 and assigns it to the AllocatedDraw field.
func (o *BulkWritablePowerPortTemplateRequest) SetAllocatedDraw(v int32) {
	o.AllocatedDraw.Set(&v)
}
// SetAllocatedDrawNil sets the value for AllocatedDraw to be an explicit nil
func (o *BulkWritablePowerPortTemplateRequest) SetAllocatedDrawNil() {
	o.AllocatedDraw.Set(nil)
}

// UnsetAllocatedDraw ensures that no value is present for AllocatedDraw, not even an explicit nil
func (o *BulkWritablePowerPortTemplateRequest) UnsetAllocatedDraw() {
	o.AllocatedDraw.Unset()
}

// GetDeviceType returns the DeviceType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkWritablePowerPortTemplateRequest) GetDeviceType() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.DeviceType.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.DeviceType.Get()
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkWritablePowerPortTemplateRequest) GetDeviceTypeOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceType.Get(), o.DeviceType.IsSet()
}

// HasDeviceType returns a boolean if a field has been set.
func (o *BulkWritablePowerPortTemplateRequest) HasDeviceType() bool {
	if o != nil && o.DeviceType.IsSet() {
		return true
	}

	return false
}

// SetDeviceType gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the DeviceType field.
func (o *BulkWritablePowerPortTemplateRequest) SetDeviceType(v BulkWritableCircuitRequestTenant) {
	o.DeviceType.Set(&v)
}
// SetDeviceTypeNil sets the value for DeviceType to be an explicit nil
func (o *BulkWritablePowerPortTemplateRequest) SetDeviceTypeNil() {
	o.DeviceType.Set(nil)
}

// UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
func (o *BulkWritablePowerPortTemplateRequest) UnsetDeviceType() {
	o.DeviceType.Unset()
}

// GetModuleType returns the ModuleType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkWritablePowerPortTemplateRequest) GetModuleType() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.ModuleType.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.ModuleType.Get()
}

// GetModuleTypeOk returns a tuple with the ModuleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkWritablePowerPortTemplateRequest) GetModuleTypeOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModuleType.Get(), o.ModuleType.IsSet()
}

// HasModuleType returns a boolean if a field has been set.
func (o *BulkWritablePowerPortTemplateRequest) HasModuleType() bool {
	if o != nil && o.ModuleType.IsSet() {
		return true
	}

	return false
}

// SetModuleType gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the ModuleType field.
func (o *BulkWritablePowerPortTemplateRequest) SetModuleType(v BulkWritableCircuitRequestTenant) {
	o.ModuleType.Set(&v)
}
// SetModuleTypeNil sets the value for ModuleType to be an explicit nil
func (o *BulkWritablePowerPortTemplateRequest) SetModuleTypeNil() {
	o.ModuleType.Set(nil)
}

// UnsetModuleType ensures that no value is present for ModuleType, not even an explicit nil
func (o *BulkWritablePowerPortTemplateRequest) UnsetModuleType() {
	o.ModuleType.Unset()
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *BulkWritablePowerPortTemplateRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritablePowerPortTemplateRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *BulkWritablePowerPortTemplateRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *BulkWritablePowerPortTemplateRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *BulkWritablePowerPortTemplateRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue {
	if o == nil || IsNil(o.Relationships) {
		var ret map[string]BulkWritableCableRequestRelationshipsValue
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritablePowerPortTemplateRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *BulkWritablePowerPortTemplateRequest) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]BulkWritableCableRequestRelationshipsValue and assigns it to the Relationships field.
func (o *BulkWritablePowerPortTemplateRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue) {
	o.Relationships = &v
}

func (o BulkWritablePowerPortTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkWritablePowerPortTemplateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.MaximumDraw.IsSet() {
		toSerialize["maximum_draw"] = o.MaximumDraw.Get()
	}
	if o.AllocatedDraw.IsSet() {
		toSerialize["allocated_draw"] = o.AllocatedDraw.Get()
	}
	if o.DeviceType.IsSet() {
		toSerialize["device_type"] = o.DeviceType.Get()
	}
	if o.ModuleType.IsSet() {
		toSerialize["module_type"] = o.ModuleType.Get()
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

func (o *BulkWritablePowerPortTemplateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
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

	varBulkWritablePowerPortTemplateRequest := _BulkWritablePowerPortTemplateRequest{}

	err = json.Unmarshal(data, &varBulkWritablePowerPortTemplateRequest)

	if err != nil {
		return err
	}

	*o = BulkWritablePowerPortTemplateRequest(varBulkWritablePowerPortTemplateRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "description")
		delete(additionalProperties, "maximum_draw")
		delete(additionalProperties, "allocated_draw")
		delete(additionalProperties, "device_type")
		delete(additionalProperties, "module_type")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "relationships")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBulkWritablePowerPortTemplateRequest struct {
	value *BulkWritablePowerPortTemplateRequest
	isSet bool
}

func (v NullableBulkWritablePowerPortTemplateRequest) Get() *BulkWritablePowerPortTemplateRequest {
	return v.value
}

func (v *NullableBulkWritablePowerPortTemplateRequest) Set(val *BulkWritablePowerPortTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkWritablePowerPortTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkWritablePowerPortTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkWritablePowerPortTemplateRequest(val *BulkWritablePowerPortTemplateRequest) *NullableBulkWritablePowerPortTemplateRequest {
	return &NullableBulkWritablePowerPortTemplateRequest{value: val, isSet: true}
}

func (v NullableBulkWritablePowerPortTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkWritablePowerPortTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


