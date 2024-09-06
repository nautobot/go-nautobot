/*
API Documentation

Source of truth and network automation platform

API version: 2.3.2 (2.3)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nautobot

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the PatchedBulkWritablePrefixRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedBulkWritablePrefixRequest{}

// PatchedBulkWritablePrefixRequest Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type PatchedBulkWritablePrefixRequest struct {
	Id string `json:"id"`
	Prefix *string `json:"prefix,omitempty"`
	Type *PrefixTypeChoices `json:"type,omitempty"`
	Location NullableBulkWritablePrefixRequestLocation `json:"location,omitempty"`
	// Date this prefix was allocated to an RIR, reserved in IPAM, etc.
	DateAllocated NullableTime `json:"date_allocated,omitempty"`
	Description *string `json:"description,omitempty"`
	Status *BulkWritableCableRequestStatus `json:"status,omitempty"`
	Role NullableBulkWritableCircuitRequestTenant `json:"role,omitempty"`
	Parent NullableBulkWritablePrefixRequestParent `json:"parent,omitempty"`
	Namespace *BulkWritableCableRequestStatus `json:"namespace,omitempty"`
	Tenant NullableBulkWritableCircuitRequestTenant `json:"tenant,omitempty"`
	Vlan NullableBulkWritableCircuitRequestTenant `json:"vlan,omitempty"`
	Rir NullableBulkWritablePrefixRequestRir `json:"rir,omitempty"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Relationships *map[string]BulkWritableCableRequestRelationshipsValue `json:"relationships,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedBulkWritablePrefixRequest PatchedBulkWritablePrefixRequest

// NewPatchedBulkWritablePrefixRequest instantiates a new PatchedBulkWritablePrefixRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedBulkWritablePrefixRequest(id string) *PatchedBulkWritablePrefixRequest {
	this := PatchedBulkWritablePrefixRequest{}
	this.Id = id
	var type_ PrefixTypeChoices = "{\"value\":\"network\",\"label\":\"Network\"}"
	this.Type = &type_
	return &this
}

// NewPatchedBulkWritablePrefixRequestWithDefaults instantiates a new PatchedBulkWritablePrefixRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedBulkWritablePrefixRequestWithDefaults() *PatchedBulkWritablePrefixRequest {
	this := PatchedBulkWritablePrefixRequest{}
	var type_ PrefixTypeChoices = "{\"value\":\"network\",\"label\":\"Network\"}"
	this.Type = &type_
	return &this
}

// GetId returns the Id field value
func (o *PatchedBulkWritablePrefixRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritablePrefixRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PatchedBulkWritablePrefixRequest) SetId(v string) {
	o.Id = v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *PatchedBulkWritablePrefixRequest) GetPrefix() string {
	if o == nil || IsNil(o.Prefix) {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritablePrefixRequest) GetPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.Prefix) {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *PatchedBulkWritablePrefixRequest) HasPrefix() bool {
	if o != nil && !IsNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *PatchedBulkWritablePrefixRequest) SetPrefix(v string) {
	o.Prefix = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PatchedBulkWritablePrefixRequest) GetType() PrefixTypeChoices {
	if o == nil || IsNil(o.Type) {
		var ret PrefixTypeChoices
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritablePrefixRequest) GetTypeOk() (*PrefixTypeChoices, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PatchedBulkWritablePrefixRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given PrefixTypeChoices and assigns it to the Type field.
func (o *PatchedBulkWritablePrefixRequest) SetType(v PrefixTypeChoices) {
	o.Type = &v
}

// GetLocation returns the Location field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBulkWritablePrefixRequest) GetLocation() BulkWritablePrefixRequestLocation {
	if o == nil || IsNil(o.Location.Get()) {
		var ret BulkWritablePrefixRequestLocation
		return ret
	}
	return *o.Location.Get()
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBulkWritablePrefixRequest) GetLocationOk() (*BulkWritablePrefixRequestLocation, bool) {
	if o == nil {
		return nil, false
	}
	return o.Location.Get(), o.Location.IsSet()
}

// HasLocation returns a boolean if a field has been set.
func (o *PatchedBulkWritablePrefixRequest) HasLocation() bool {
	if o != nil && o.Location.IsSet() {
		return true
	}

	return false
}

// SetLocation gets a reference to the given NullableBulkWritablePrefixRequestLocation and assigns it to the Location field.
func (o *PatchedBulkWritablePrefixRequest) SetLocation(v BulkWritablePrefixRequestLocation) {
	o.Location.Set(&v)
}
// SetLocationNil sets the value for Location to be an explicit nil
func (o *PatchedBulkWritablePrefixRequest) SetLocationNil() {
	o.Location.Set(nil)
}

// UnsetLocation ensures that no value is present for Location, not even an explicit nil
func (o *PatchedBulkWritablePrefixRequest) UnsetLocation() {
	o.Location.Unset()
}

// GetDateAllocated returns the DateAllocated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBulkWritablePrefixRequest) GetDateAllocated() time.Time {
	if o == nil || IsNil(o.DateAllocated.Get()) {
		var ret time.Time
		return ret
	}
	return *o.DateAllocated.Get()
}

// GetDateAllocatedOk returns a tuple with the DateAllocated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBulkWritablePrefixRequest) GetDateAllocatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DateAllocated.Get(), o.DateAllocated.IsSet()
}

// HasDateAllocated returns a boolean if a field has been set.
func (o *PatchedBulkWritablePrefixRequest) HasDateAllocated() bool {
	if o != nil && o.DateAllocated.IsSet() {
		return true
	}

	return false
}

// SetDateAllocated gets a reference to the given NullableTime and assigns it to the DateAllocated field.
func (o *PatchedBulkWritablePrefixRequest) SetDateAllocated(v time.Time) {
	o.DateAllocated.Set(&v)
}
// SetDateAllocatedNil sets the value for DateAllocated to be an explicit nil
func (o *PatchedBulkWritablePrefixRequest) SetDateAllocatedNil() {
	o.DateAllocated.Set(nil)
}

// UnsetDateAllocated ensures that no value is present for DateAllocated, not even an explicit nil
func (o *PatchedBulkWritablePrefixRequest) UnsetDateAllocated() {
	o.DateAllocated.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedBulkWritablePrefixRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritablePrefixRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedBulkWritablePrefixRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedBulkWritablePrefixRequest) SetDescription(v string) {
	o.Description = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PatchedBulkWritablePrefixRequest) GetStatus() BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Status) {
		var ret BulkWritableCableRequestStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritablePrefixRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PatchedBulkWritablePrefixRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given BulkWritableCableRequestStatus and assigns it to the Status field.
func (o *PatchedBulkWritablePrefixRequest) SetStatus(v BulkWritableCableRequestStatus) {
	o.Status = &v
}

// GetRole returns the Role field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBulkWritablePrefixRequest) GetRole() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.Role.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.Role.Get()
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBulkWritablePrefixRequest) GetRoleOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Role.Get(), o.Role.IsSet()
}

// HasRole returns a boolean if a field has been set.
func (o *PatchedBulkWritablePrefixRequest) HasRole() bool {
	if o != nil && o.Role.IsSet() {
		return true
	}

	return false
}

// SetRole gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the Role field.
func (o *PatchedBulkWritablePrefixRequest) SetRole(v BulkWritableCircuitRequestTenant) {
	o.Role.Set(&v)
}
// SetRoleNil sets the value for Role to be an explicit nil
func (o *PatchedBulkWritablePrefixRequest) SetRoleNil() {
	o.Role.Set(nil)
}

// UnsetRole ensures that no value is present for Role, not even an explicit nil
func (o *PatchedBulkWritablePrefixRequest) UnsetRole() {
	o.Role.Unset()
}

// GetParent returns the Parent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBulkWritablePrefixRequest) GetParent() BulkWritablePrefixRequestParent {
	if o == nil || IsNil(o.Parent.Get()) {
		var ret BulkWritablePrefixRequestParent
		return ret
	}
	return *o.Parent.Get()
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBulkWritablePrefixRequest) GetParentOk() (*BulkWritablePrefixRequestParent, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parent.Get(), o.Parent.IsSet()
}

// HasParent returns a boolean if a field has been set.
func (o *PatchedBulkWritablePrefixRequest) HasParent() bool {
	if o != nil && o.Parent.IsSet() {
		return true
	}

	return false
}

// SetParent gets a reference to the given NullableBulkWritablePrefixRequestParent and assigns it to the Parent field.
func (o *PatchedBulkWritablePrefixRequest) SetParent(v BulkWritablePrefixRequestParent) {
	o.Parent.Set(&v)
}
// SetParentNil sets the value for Parent to be an explicit nil
func (o *PatchedBulkWritablePrefixRequest) SetParentNil() {
	o.Parent.Set(nil)
}

// UnsetParent ensures that no value is present for Parent, not even an explicit nil
func (o *PatchedBulkWritablePrefixRequest) UnsetParent() {
	o.Parent.Unset()
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *PatchedBulkWritablePrefixRequest) GetNamespace() BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Namespace) {
		var ret BulkWritableCableRequestStatus
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritablePrefixRequest) GetNamespaceOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *PatchedBulkWritablePrefixRequest) HasNamespace() bool {
	if o != nil && !IsNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given BulkWritableCableRequestStatus and assigns it to the Namespace field.
func (o *PatchedBulkWritablePrefixRequest) SetNamespace(v BulkWritableCableRequestStatus) {
	o.Namespace = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBulkWritablePrefixRequest) GetTenant() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBulkWritablePrefixRequest) GetTenantOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *PatchedBulkWritablePrefixRequest) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the Tenant field.
func (o *PatchedBulkWritablePrefixRequest) SetTenant(v BulkWritableCircuitRequestTenant) {
	o.Tenant.Set(&v)
}
// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *PatchedBulkWritablePrefixRequest) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *PatchedBulkWritablePrefixRequest) UnsetTenant() {
	o.Tenant.Unset()
}

// GetVlan returns the Vlan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBulkWritablePrefixRequest) GetVlan() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.Vlan.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.Vlan.Get()
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBulkWritablePrefixRequest) GetVlanOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vlan.Get(), o.Vlan.IsSet()
}

// HasVlan returns a boolean if a field has been set.
func (o *PatchedBulkWritablePrefixRequest) HasVlan() bool {
	if o != nil && o.Vlan.IsSet() {
		return true
	}

	return false
}

// SetVlan gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the Vlan field.
func (o *PatchedBulkWritablePrefixRequest) SetVlan(v BulkWritableCircuitRequestTenant) {
	o.Vlan.Set(&v)
}
// SetVlanNil sets the value for Vlan to be an explicit nil
func (o *PatchedBulkWritablePrefixRequest) SetVlanNil() {
	o.Vlan.Set(nil)
}

// UnsetVlan ensures that no value is present for Vlan, not even an explicit nil
func (o *PatchedBulkWritablePrefixRequest) UnsetVlan() {
	o.Vlan.Unset()
}

// GetRir returns the Rir field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBulkWritablePrefixRequest) GetRir() BulkWritablePrefixRequestRir {
	if o == nil || IsNil(o.Rir.Get()) {
		var ret BulkWritablePrefixRequestRir
		return ret
	}
	return *o.Rir.Get()
}

// GetRirOk returns a tuple with the Rir field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBulkWritablePrefixRequest) GetRirOk() (*BulkWritablePrefixRequestRir, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rir.Get(), o.Rir.IsSet()
}

// HasRir returns a boolean if a field has been set.
func (o *PatchedBulkWritablePrefixRequest) HasRir() bool {
	if o != nil && o.Rir.IsSet() {
		return true
	}

	return false
}

// SetRir gets a reference to the given NullableBulkWritablePrefixRequestRir and assigns it to the Rir field.
func (o *PatchedBulkWritablePrefixRequest) SetRir(v BulkWritablePrefixRequestRir) {
	o.Rir.Set(&v)
}
// SetRirNil sets the value for Rir to be an explicit nil
func (o *PatchedBulkWritablePrefixRequest) SetRirNil() {
	o.Rir.Set(nil)
}

// UnsetRir ensures that no value is present for Rir, not even an explicit nil
func (o *PatchedBulkWritablePrefixRequest) UnsetRir() {
	o.Rir.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedBulkWritablePrefixRequest) GetTags() []BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Tags) {
		var ret []BulkWritableCableRequestStatus
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritablePrefixRequest) GetTagsOk() ([]BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedBulkWritablePrefixRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []BulkWritableCableRequestStatus and assigns it to the Tags field.
func (o *PatchedBulkWritablePrefixRequest) SetTags(v []BulkWritableCableRequestStatus) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedBulkWritablePrefixRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritablePrefixRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedBulkWritablePrefixRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedBulkWritablePrefixRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PatchedBulkWritablePrefixRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue {
	if o == nil || IsNil(o.Relationships) {
		var ret map[string]BulkWritableCableRequestRelationshipsValue
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritablePrefixRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PatchedBulkWritablePrefixRequest) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]BulkWritableCableRequestRelationshipsValue and assigns it to the Relationships field.
func (o *PatchedBulkWritablePrefixRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue) {
	o.Relationships = &v
}

func (o PatchedBulkWritablePrefixRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedBulkWritablePrefixRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Prefix) {
		toSerialize["prefix"] = o.Prefix
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.Location.IsSet() {
		toSerialize["location"] = o.Location.Get()
	}
	if o.DateAllocated.IsSet() {
		toSerialize["date_allocated"] = o.DateAllocated.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.Role.IsSet() {
		toSerialize["role"] = o.Role.Get()
	}
	if o.Parent.IsSet() {
		toSerialize["parent"] = o.Parent.Get()
	}
	if !IsNil(o.Namespace) {
		toSerialize["namespace"] = o.Namespace
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if o.Vlan.IsSet() {
		toSerialize["vlan"] = o.Vlan.Get()
	}
	if o.Rir.IsSet() {
		toSerialize["rir"] = o.Rir.Get()
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

func (o *PatchedBulkWritablePrefixRequest) UnmarshalJSON(data []byte) (err error) {
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

	varPatchedBulkWritablePrefixRequest := _PatchedBulkWritablePrefixRequest{}

	err = json.Unmarshal(data, &varPatchedBulkWritablePrefixRequest)

	if err != nil {
		return err
	}

	*o = PatchedBulkWritablePrefixRequest(varPatchedBulkWritablePrefixRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "prefix")
		delete(additionalProperties, "type")
		delete(additionalProperties, "location")
		delete(additionalProperties, "date_allocated")
		delete(additionalProperties, "description")
		delete(additionalProperties, "status")
		delete(additionalProperties, "role")
		delete(additionalProperties, "parent")
		delete(additionalProperties, "namespace")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "vlan")
		delete(additionalProperties, "rir")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "relationships")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedBulkWritablePrefixRequest struct {
	value *PatchedBulkWritablePrefixRequest
	isSet bool
}

func (v NullablePatchedBulkWritablePrefixRequest) Get() *PatchedBulkWritablePrefixRequest {
	return v.value
}

func (v *NullablePatchedBulkWritablePrefixRequest) Set(val *PatchedBulkWritablePrefixRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedBulkWritablePrefixRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedBulkWritablePrefixRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedBulkWritablePrefixRequest(val *PatchedBulkWritablePrefixRequest) *NullablePatchedBulkWritablePrefixRequest {
	return &NullablePatchedBulkWritablePrefixRequest{value: val, isSet: true}
}

func (v NullablePatchedBulkWritablePrefixRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedBulkWritablePrefixRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


