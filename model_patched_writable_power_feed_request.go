/*
API Documentation

Source of truth and network automation platform

API version: 2.3.1 (2.3)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nautobot

import (
	"encoding/json"
)

// checks if the PatchedWritablePowerFeedRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedWritablePowerFeedRequest{}

// PatchedWritablePowerFeedRequest Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type PatchedWritablePowerFeedRequest struct {
	Name *string `json:"name,omitempty"`
	Type *PowerFeedTypeChoices `json:"type,omitempty"`
	Supply *SupplyEnum `json:"supply,omitempty"`
	Phase *PhaseEnum `json:"phase,omitempty"`
	Voltage *int32 `json:"voltage,omitempty"`
	Amperage *int32 `json:"amperage,omitempty"`
	// Maximum permissible draw (percentage)
	MaxUtilization *int32 `json:"max_utilization,omitempty"`
	Comments *string `json:"comments,omitempty"`
	Cable NullableBulkWritableCircuitRequestTenant `json:"cable,omitempty"`
	PowerPanel *BulkWritableCableRequestStatus `json:"power_panel,omitempty"`
	Rack NullableBulkWritableCircuitRequestTenant `json:"rack,omitempty"`
	Status *BulkWritableCableRequestStatus `json:"status,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Relationships *map[string]BulkWritableCableRequestRelationshipsValue `json:"relationships,omitempty"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedWritablePowerFeedRequest PatchedWritablePowerFeedRequest

// NewPatchedWritablePowerFeedRequest instantiates a new PatchedWritablePowerFeedRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritablePowerFeedRequest() *PatchedWritablePowerFeedRequest {
	this := PatchedWritablePowerFeedRequest{}
	return &this
}

// NewPatchedWritablePowerFeedRequestWithDefaults instantiates a new PatchedWritablePowerFeedRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritablePowerFeedRequestWithDefaults() *PatchedWritablePowerFeedRequest {
	this := PatchedWritablePowerFeedRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedWritablePowerFeedRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerFeedRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedWritablePowerFeedRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedWritablePowerFeedRequest) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PatchedWritablePowerFeedRequest) GetType() PowerFeedTypeChoices {
	if o == nil || IsNil(o.Type) {
		var ret PowerFeedTypeChoices
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerFeedRequest) GetTypeOk() (*PowerFeedTypeChoices, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PatchedWritablePowerFeedRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given PowerFeedTypeChoices and assigns it to the Type field.
func (o *PatchedWritablePowerFeedRequest) SetType(v PowerFeedTypeChoices) {
	o.Type = &v
}

// GetSupply returns the Supply field value if set, zero value otherwise.
func (o *PatchedWritablePowerFeedRequest) GetSupply() SupplyEnum {
	if o == nil || IsNil(o.Supply) {
		var ret SupplyEnum
		return ret
	}
	return *o.Supply
}

// GetSupplyOk returns a tuple with the Supply field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerFeedRequest) GetSupplyOk() (*SupplyEnum, bool) {
	if o == nil || IsNil(o.Supply) {
		return nil, false
	}
	return o.Supply, true
}

// HasSupply returns a boolean if a field has been set.
func (o *PatchedWritablePowerFeedRequest) HasSupply() bool {
	if o != nil && !IsNil(o.Supply) {
		return true
	}

	return false
}

// SetSupply gets a reference to the given SupplyEnum and assigns it to the Supply field.
func (o *PatchedWritablePowerFeedRequest) SetSupply(v SupplyEnum) {
	o.Supply = &v
}

// GetPhase returns the Phase field value if set, zero value otherwise.
func (o *PatchedWritablePowerFeedRequest) GetPhase() PhaseEnum {
	if o == nil || IsNil(o.Phase) {
		var ret PhaseEnum
		return ret
	}
	return *o.Phase
}

// GetPhaseOk returns a tuple with the Phase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerFeedRequest) GetPhaseOk() (*PhaseEnum, bool) {
	if o == nil || IsNil(o.Phase) {
		return nil, false
	}
	return o.Phase, true
}

// HasPhase returns a boolean if a field has been set.
func (o *PatchedWritablePowerFeedRequest) HasPhase() bool {
	if o != nil && !IsNil(o.Phase) {
		return true
	}

	return false
}

// SetPhase gets a reference to the given PhaseEnum and assigns it to the Phase field.
func (o *PatchedWritablePowerFeedRequest) SetPhase(v PhaseEnum) {
	o.Phase = &v
}

// GetVoltage returns the Voltage field value if set, zero value otherwise.
func (o *PatchedWritablePowerFeedRequest) GetVoltage() int32 {
	if o == nil || IsNil(o.Voltage) {
		var ret int32
		return ret
	}
	return *o.Voltage
}

// GetVoltageOk returns a tuple with the Voltage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerFeedRequest) GetVoltageOk() (*int32, bool) {
	if o == nil || IsNil(o.Voltage) {
		return nil, false
	}
	return o.Voltage, true
}

// HasVoltage returns a boolean if a field has been set.
func (o *PatchedWritablePowerFeedRequest) HasVoltage() bool {
	if o != nil && !IsNil(o.Voltage) {
		return true
	}

	return false
}

// SetVoltage gets a reference to the given int32 and assigns it to the Voltage field.
func (o *PatchedWritablePowerFeedRequest) SetVoltage(v int32) {
	o.Voltage = &v
}

// GetAmperage returns the Amperage field value if set, zero value otherwise.
func (o *PatchedWritablePowerFeedRequest) GetAmperage() int32 {
	if o == nil || IsNil(o.Amperage) {
		var ret int32
		return ret
	}
	return *o.Amperage
}

// GetAmperageOk returns a tuple with the Amperage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerFeedRequest) GetAmperageOk() (*int32, bool) {
	if o == nil || IsNil(o.Amperage) {
		return nil, false
	}
	return o.Amperage, true
}

// HasAmperage returns a boolean if a field has been set.
func (o *PatchedWritablePowerFeedRequest) HasAmperage() bool {
	if o != nil && !IsNil(o.Amperage) {
		return true
	}

	return false
}

// SetAmperage gets a reference to the given int32 and assigns it to the Amperage field.
func (o *PatchedWritablePowerFeedRequest) SetAmperage(v int32) {
	o.Amperage = &v
}

// GetMaxUtilization returns the MaxUtilization field value if set, zero value otherwise.
func (o *PatchedWritablePowerFeedRequest) GetMaxUtilization() int32 {
	if o == nil || IsNil(o.MaxUtilization) {
		var ret int32
		return ret
	}
	return *o.MaxUtilization
}

// GetMaxUtilizationOk returns a tuple with the MaxUtilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerFeedRequest) GetMaxUtilizationOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxUtilization) {
		return nil, false
	}
	return o.MaxUtilization, true
}

// HasMaxUtilization returns a boolean if a field has been set.
func (o *PatchedWritablePowerFeedRequest) HasMaxUtilization() bool {
	if o != nil && !IsNil(o.MaxUtilization) {
		return true
	}

	return false
}

// SetMaxUtilization gets a reference to the given int32 and assigns it to the MaxUtilization field.
func (o *PatchedWritablePowerFeedRequest) SetMaxUtilization(v int32) {
	o.MaxUtilization = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *PatchedWritablePowerFeedRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerFeedRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *PatchedWritablePowerFeedRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *PatchedWritablePowerFeedRequest) SetComments(v string) {
	o.Comments = &v
}

// GetCable returns the Cable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritablePowerFeedRequest) GetCable() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.Cable.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.Cable.Get()
}

// GetCableOk returns a tuple with the Cable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritablePowerFeedRequest) GetCableOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cable.Get(), o.Cable.IsSet()
}

// HasCable returns a boolean if a field has been set.
func (o *PatchedWritablePowerFeedRequest) HasCable() bool {
	if o != nil && o.Cable.IsSet() {
		return true
	}

	return false
}

// SetCable gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the Cable field.
func (o *PatchedWritablePowerFeedRequest) SetCable(v BulkWritableCircuitRequestTenant) {
	o.Cable.Set(&v)
}
// SetCableNil sets the value for Cable to be an explicit nil
func (o *PatchedWritablePowerFeedRequest) SetCableNil() {
	o.Cable.Set(nil)
}

// UnsetCable ensures that no value is present for Cable, not even an explicit nil
func (o *PatchedWritablePowerFeedRequest) UnsetCable() {
	o.Cable.Unset()
}

// GetPowerPanel returns the PowerPanel field value if set, zero value otherwise.
func (o *PatchedWritablePowerFeedRequest) GetPowerPanel() BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.PowerPanel) {
		var ret BulkWritableCableRequestStatus
		return ret
	}
	return *o.PowerPanel
}

// GetPowerPanelOk returns a tuple with the PowerPanel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerFeedRequest) GetPowerPanelOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.PowerPanel) {
		return nil, false
	}
	return o.PowerPanel, true
}

// HasPowerPanel returns a boolean if a field has been set.
func (o *PatchedWritablePowerFeedRequest) HasPowerPanel() bool {
	if o != nil && !IsNil(o.PowerPanel) {
		return true
	}

	return false
}

// SetPowerPanel gets a reference to the given BulkWritableCableRequestStatus and assigns it to the PowerPanel field.
func (o *PatchedWritablePowerFeedRequest) SetPowerPanel(v BulkWritableCableRequestStatus) {
	o.PowerPanel = &v
}

// GetRack returns the Rack field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedWritablePowerFeedRequest) GetRack() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.Rack.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.Rack.Get()
}

// GetRackOk returns a tuple with the Rack field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedWritablePowerFeedRequest) GetRackOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rack.Get(), o.Rack.IsSet()
}

// HasRack returns a boolean if a field has been set.
func (o *PatchedWritablePowerFeedRequest) HasRack() bool {
	if o != nil && o.Rack.IsSet() {
		return true
	}

	return false
}

// SetRack gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the Rack field.
func (o *PatchedWritablePowerFeedRequest) SetRack(v BulkWritableCircuitRequestTenant) {
	o.Rack.Set(&v)
}
// SetRackNil sets the value for Rack to be an explicit nil
func (o *PatchedWritablePowerFeedRequest) SetRackNil() {
	o.Rack.Set(nil)
}

// UnsetRack ensures that no value is present for Rack, not even an explicit nil
func (o *PatchedWritablePowerFeedRequest) UnsetRack() {
	o.Rack.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PatchedWritablePowerFeedRequest) GetStatus() BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Status) {
		var ret BulkWritableCableRequestStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerFeedRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PatchedWritablePowerFeedRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given BulkWritableCableRequestStatus and assigns it to the Status field.
func (o *PatchedWritablePowerFeedRequest) SetStatus(v BulkWritableCableRequestStatus) {
	o.Status = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedWritablePowerFeedRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerFeedRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedWritablePowerFeedRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedWritablePowerFeedRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PatchedWritablePowerFeedRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue {
	if o == nil || IsNil(o.Relationships) {
		var ret map[string]BulkWritableCableRequestRelationshipsValue
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerFeedRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PatchedWritablePowerFeedRequest) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]BulkWritableCableRequestRelationshipsValue and assigns it to the Relationships field.
func (o *PatchedWritablePowerFeedRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue) {
	o.Relationships = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedWritablePowerFeedRequest) GetTags() []BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Tags) {
		var ret []BulkWritableCableRequestStatus
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritablePowerFeedRequest) GetTagsOk() ([]BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedWritablePowerFeedRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []BulkWritableCableRequestStatus and assigns it to the Tags field.
func (o *PatchedWritablePowerFeedRequest) SetTags(v []BulkWritableCableRequestStatus) {
	o.Tags = v
}

func (o PatchedWritablePowerFeedRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedWritablePowerFeedRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Supply) {
		toSerialize["supply"] = o.Supply
	}
	if !IsNil(o.Phase) {
		toSerialize["phase"] = o.Phase
	}
	if !IsNil(o.Voltage) {
		toSerialize["voltage"] = o.Voltage
	}
	if !IsNil(o.Amperage) {
		toSerialize["amperage"] = o.Amperage
	}
	if !IsNil(o.MaxUtilization) {
		toSerialize["max_utilization"] = o.MaxUtilization
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	if o.Cable.IsSet() {
		toSerialize["cable"] = o.Cable.Get()
	}
	if !IsNil(o.PowerPanel) {
		toSerialize["power_panel"] = o.PowerPanel
	}
	if o.Rack.IsSet() {
		toSerialize["rack"] = o.Rack.Get()
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
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

func (o *PatchedWritablePowerFeedRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedWritablePowerFeedRequest := _PatchedWritablePowerFeedRequest{}

	err = json.Unmarshal(data, &varPatchedWritablePowerFeedRequest)

	if err != nil {
		return err
	}

	*o = PatchedWritablePowerFeedRequest(varPatchedWritablePowerFeedRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "supply")
		delete(additionalProperties, "phase")
		delete(additionalProperties, "voltage")
		delete(additionalProperties, "amperage")
		delete(additionalProperties, "max_utilization")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "cable")
		delete(additionalProperties, "power_panel")
		delete(additionalProperties, "rack")
		delete(additionalProperties, "status")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "relationships")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedWritablePowerFeedRequest struct {
	value *PatchedWritablePowerFeedRequest
	isSet bool
}

func (v NullablePatchedWritablePowerFeedRequest) Get() *PatchedWritablePowerFeedRequest {
	return v.value
}

func (v *NullablePatchedWritablePowerFeedRequest) Set(val *PatchedWritablePowerFeedRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritablePowerFeedRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritablePowerFeedRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritablePowerFeedRequest(val *PatchedWritablePowerFeedRequest) *NullablePatchedWritablePowerFeedRequest {
	return &NullablePatchedWritablePowerFeedRequest{value: val, isSet: true}
}

func (v NullablePatchedWritablePowerFeedRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritablePowerFeedRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

