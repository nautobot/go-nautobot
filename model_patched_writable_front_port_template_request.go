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

// checks if the PatchedWritableFrontPortTemplateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedWritableFrontPortTemplateRequest{}

// PatchedWritableFrontPortTemplateRequest Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type PatchedWritableFrontPortTemplateRequest struct {
	Name *string `json:"name,omitempty"`
	// Physical label
	Label *string `json:"label,omitempty"`
	Description *string `json:"description,omitempty"`
	Type *PortTypeChoices `json:"type,omitempty"`
	RearPortPosition *int32 `json:"rear_port_position,omitempty"`
	DeviceType NullableBulkWritableCircuitRequestTenant `json:"device_type,omitempty"`
	ModuleType NullableBulkWritableCircuitRequestTenant `json:"module_type,omitempty"`
	RearPortTemplate *BulkWritableCableRequestStatus `json:"rear_port_template,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Relationships *map[string]BulkWritableCableRequestRelationshipsValue `json:"relationships,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedWritableFrontPortTemplateRequest PatchedWritableFrontPortTemplateRequest

// NewPatchedWritableFrontPortTemplateRequest instantiates a new PatchedWritableFrontPortTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritableFrontPortTemplateRequest() *PatchedWritableFrontPortTemplateRequest {
	this := PatchedWritableFrontPortTemplateRequest{}
	return &this
}

// NewPatchedWritableFrontPortTemplateRequestWithDefaults instantiates a new PatchedWritableFrontPortTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritableFrontPortTemplateRequestWithDefaults() *PatchedWritableFrontPortTemplateRequest {
	this := PatchedWritableFrontPortTemplateRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedWritableFrontPortTemplateRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableFrontPortTemplateRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedWritableFrontPortTemplateRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedWritableFrontPortTemplateRequest) SetName(v string) {
	o.Name = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PatchedWritableFrontPortTemplateRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableFrontPortTemplateRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PatchedWritableFrontPortTemplateRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PatchedWritableFrontPortTemplateRequest) SetLabel(v string) {
	o.Label = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedWritableFrontPortTemplateRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableFrontPortTemplateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedWritableFrontPortTemplateRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedWritableFrontPortTemplateRequest) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PatchedWritableFrontPortTemplateRequest) GetType() PortTypeChoices {
	if o == nil || IsNil(o.Type) {
		var ret PortTypeChoices
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableFrontPortTemplateRequest) GetTypeOk() (*PortTypeChoices, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PatchedWritableFrontPortTemplateRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given PortTypeChoices and assigns it to the Type field.
func (o *PatchedWritableFrontPortTemplateRequest) SetType(v PortTypeChoices) {
	o.Type = &v
}

// GetRearPortPosition returns the RearPortPosition field value if set, zero value otherwise.
func (o *PatchedWritableFrontPortTemplateRequest) GetRearPortPosition() int32 {
	if o == nil || IsNil(o.RearPortPosition) {
		var ret int32
		return ret
	}
	return *o.RearPortPosition
}

// GetRearPortPositionOk returns a tuple with the RearPortPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableFrontPortTemplateRequest) GetRearPortPositionOk() (*int32, bool) {
	if o == nil || IsNil(o.RearPortPosition) {
		return nil, false
	}
	return o.RearPortPosition, true
}

// HasRearPortPosition returns a boolean if a field has been set.
func (o *PatchedWritableFrontPortTemplateRequest) HasRearPortPosition() bool {
	if o != nil && !IsNil(o.RearPortPosition) {
		return true
	}

	return false
}

// SetRearPortPosition gets a reference to the given int32 and assigns it to the RearPortPosition field.
func (o *PatchedWritableFrontPortTemplateRequest) SetRearPortPosition(v int32) {
	o.RearPortPosition = &v
}

// GetDeviceType returns the DeviceType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableFrontPortTemplateRequest) GetDeviceType() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.DeviceType.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.DeviceType.Get()
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableFrontPortTemplateRequest) GetDeviceTypeOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceType.Get(), o.DeviceType.IsSet()
}

// HasDeviceType returns a boolean if a field has been set.
func (o *PatchedWritableFrontPortTemplateRequest) HasDeviceType() bool {
	if o != nil && o.DeviceType.IsSet() {
		return true
	}

	return false
}

// SetDeviceType gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the DeviceType field.
func (o *PatchedWritableFrontPortTemplateRequest) SetDeviceType(v BulkWritableCircuitRequestTenant) {
	o.DeviceType.Set(&v)
}
// SetDeviceTypeNil sets the value for DeviceType to be an explicit nil
func (o *PatchedWritableFrontPortTemplateRequest) SetDeviceTypeNil() {
	o.DeviceType.Set(nil)
}

// UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
func (o *PatchedWritableFrontPortTemplateRequest) UnsetDeviceType() {
	o.DeviceType.Unset()
}

// GetModuleType returns the ModuleType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritableFrontPortTemplateRequest) GetModuleType() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.ModuleType.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.ModuleType.Get()
}

// GetModuleTypeOk returns a tuple with the ModuleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritableFrontPortTemplateRequest) GetModuleTypeOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModuleType.Get(), o.ModuleType.IsSet()
}

// HasModuleType returns a boolean if a field has been set.
func (o *PatchedWritableFrontPortTemplateRequest) HasModuleType() bool {
	if o != nil && o.ModuleType.IsSet() {
		return true
	}

	return false
}

// SetModuleType gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the ModuleType field.
func (o *PatchedWritableFrontPortTemplateRequest) SetModuleType(v BulkWritableCircuitRequestTenant) {
	o.ModuleType.Set(&v)
}
// SetModuleTypeNil sets the value for ModuleType to be an explicit nil
func (o *PatchedWritableFrontPortTemplateRequest) SetModuleTypeNil() {
	o.ModuleType.Set(nil)
}

// UnsetModuleType ensures that no value is present for ModuleType, not even an explicit nil
func (o *PatchedWritableFrontPortTemplateRequest) UnsetModuleType() {
	o.ModuleType.Unset()
}

// GetRearPortTemplate returns the RearPortTemplate field value if set, zero value otherwise.
func (o *PatchedWritableFrontPortTemplateRequest) GetRearPortTemplate() BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.RearPortTemplate) {
		var ret BulkWritableCableRequestStatus
		return ret
	}
	return *o.RearPortTemplate
}

// GetRearPortTemplateOk returns a tuple with the RearPortTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableFrontPortTemplateRequest) GetRearPortTemplateOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.RearPortTemplate) {
		return nil, false
	}
	return o.RearPortTemplate, true
}

// HasRearPortTemplate returns a boolean if a field has been set.
func (o *PatchedWritableFrontPortTemplateRequest) HasRearPortTemplate() bool {
	if o != nil && !IsNil(o.RearPortTemplate) {
		return true
	}

	return false
}

// SetRearPortTemplate gets a reference to the given BulkWritableCableRequestStatus and assigns it to the RearPortTemplate field.
func (o *PatchedWritableFrontPortTemplateRequest) SetRearPortTemplate(v BulkWritableCableRequestStatus) {
	o.RearPortTemplate = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedWritableFrontPortTemplateRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableFrontPortTemplateRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedWritableFrontPortTemplateRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedWritableFrontPortTemplateRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PatchedWritableFrontPortTemplateRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue {
	if o == nil || IsNil(o.Relationships) {
		var ret map[string]BulkWritableCableRequestRelationshipsValue
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableFrontPortTemplateRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PatchedWritableFrontPortTemplateRequest) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]BulkWritableCableRequestRelationshipsValue and assigns it to the Relationships field.
func (o *PatchedWritableFrontPortTemplateRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue) {
	o.Relationships = &v
}

func (o PatchedWritableFrontPortTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedWritableFrontPortTemplateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.RearPortPosition) {
		toSerialize["rear_port_position"] = o.RearPortPosition
	}
	if o.DeviceType.IsSet() {
		toSerialize["device_type"] = o.DeviceType.Get()
	}
	if o.ModuleType.IsSet() {
		toSerialize["module_type"] = o.ModuleType.Get()
	}
	if !IsNil(o.RearPortTemplate) {
		toSerialize["rear_port_template"] = o.RearPortTemplate
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

func (o *PatchedWritableFrontPortTemplateRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedWritableFrontPortTemplateRequest := _PatchedWritableFrontPortTemplateRequest{}

	err = json.Unmarshal(data, &varPatchedWritableFrontPortTemplateRequest)

	if err != nil {
		return err
	}

	*o = PatchedWritableFrontPortTemplateRequest(varPatchedWritableFrontPortTemplateRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "description")
		delete(additionalProperties, "type")
		delete(additionalProperties, "rear_port_position")
		delete(additionalProperties, "device_type")
		delete(additionalProperties, "module_type")
		delete(additionalProperties, "rear_port_template")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "relationships")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedWritableFrontPortTemplateRequest struct {
	value *PatchedWritableFrontPortTemplateRequest
	isSet bool
}

func (v NullablePatchedWritableFrontPortTemplateRequest) Get() *PatchedWritableFrontPortTemplateRequest {
	return v.value
}

func (v *NullablePatchedWritableFrontPortTemplateRequest) Set(val *PatchedWritableFrontPortTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableFrontPortTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableFrontPortTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableFrontPortTemplateRequest(val *PatchedWritableFrontPortTemplateRequest) *NullablePatchedWritableFrontPortTemplateRequest {
	return &NullablePatchedWritableFrontPortTemplateRequest{value: val, isSet: true}
}

func (v NullablePatchedWritableFrontPortTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableFrontPortTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


