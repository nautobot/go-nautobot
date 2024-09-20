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

// checks if the RackReservationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RackReservationRequest{}

// RackReservationRequest Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type RackReservationRequest struct {
	// List of rack unit numbers to reserve
	Units interface{} `json:"units"`
	Description string `json:"description"`
	Rack BulkWritableCableRequestStatus `json:"rack"`
	Tenant NullableBulkWritableCircuitRequestTenant `json:"tenant,omitempty"`
	User *BulkWritableRackReservationRequestUser `json:"user,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Relationships *map[string]BulkWritableCableRequestRelationshipsValue `json:"relationships,omitempty"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RackReservationRequest RackReservationRequest

// NewRackReservationRequest instantiates a new RackReservationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRackReservationRequest(units interface{}, description string, rack BulkWritableCableRequestStatus) *RackReservationRequest {
	this := RackReservationRequest{}
	this.Units = units
	this.Description = description
	this.Rack = rack
	return &this
}

// NewRackReservationRequestWithDefaults instantiates a new RackReservationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRackReservationRequestWithDefaults() *RackReservationRequest {
	this := RackReservationRequest{}
	return &this
}

// GetUnits returns the Units field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *RackReservationRequest) GetUnits() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Units
}

// GetUnitsOk returns a tuple with the Units field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RackReservationRequest) GetUnitsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Units) {
		return nil, false
	}
	return &o.Units, true
}

// SetUnits sets field value
func (o *RackReservationRequest) SetUnits(v interface{}) {
	o.Units = v
}

// GetDescription returns the Description field value
func (o *RackReservationRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *RackReservationRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *RackReservationRequest) SetDescription(v string) {
	o.Description = v
}

// GetRack returns the Rack field value
func (o *RackReservationRequest) GetRack() BulkWritableCableRequestStatus {
	if o == nil {
		var ret BulkWritableCableRequestStatus
		return ret
	}

	return o.Rack
}

// GetRackOk returns a tuple with the Rack field value
// and a boolean to check if the value has been set.
func (o *RackReservationRequest) GetRackOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rack, true
}

// SetRack sets field value
func (o *RackReservationRequest) SetRack(v BulkWritableCableRequestStatus) {
	o.Rack = v
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RackReservationRequest) GetTenant() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RackReservationRequest) GetTenantOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *RackReservationRequest) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the Tenant field.
func (o *RackReservationRequest) SetTenant(v BulkWritableCircuitRequestTenant) {
	o.Tenant.Set(&v)
}
// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *RackReservationRequest) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *RackReservationRequest) UnsetTenant() {
	o.Tenant.Unset()
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *RackReservationRequest) GetUser() BulkWritableRackReservationRequestUser {
	if o == nil || IsNil(o.User) {
		var ret BulkWritableRackReservationRequestUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackReservationRequest) GetUserOk() (*BulkWritableRackReservationRequestUser, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *RackReservationRequest) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given BulkWritableRackReservationRequestUser and assigns it to the User field.
func (o *RackReservationRequest) SetUser(v BulkWritableRackReservationRequestUser) {
	o.User = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *RackReservationRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackReservationRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *RackReservationRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *RackReservationRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *RackReservationRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue {
	if o == nil || IsNil(o.Relationships) {
		var ret map[string]BulkWritableCableRequestRelationshipsValue
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackReservationRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *RackReservationRequest) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]BulkWritableCableRequestRelationshipsValue and assigns it to the Relationships field.
func (o *RackReservationRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue) {
	o.Relationships = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *RackReservationRequest) GetTags() []BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Tags) {
		var ret []BulkWritableCableRequestStatus
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackReservationRequest) GetTagsOk() ([]BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *RackReservationRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []BulkWritableCableRequestStatus and assigns it to the Tags field.
func (o *RackReservationRequest) SetTags(v []BulkWritableCableRequestStatus) {
	o.Tags = v
}

func (o RackReservationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RackReservationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Units != nil {
		toSerialize["units"] = o.Units
	}
	toSerialize["description"] = o.Description
	toSerialize["rack"] = o.Rack
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
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

func (o *RackReservationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"units",
		"description",
		"rack",
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

	varRackReservationRequest := _RackReservationRequest{}

	err = json.Unmarshal(data, &varRackReservationRequest)

	if err != nil {
		return err
	}

	*o = RackReservationRequest(varRackReservationRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "units")
		delete(additionalProperties, "description")
		delete(additionalProperties, "rack")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "user")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "relationships")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRackReservationRequest struct {
	value *RackReservationRequest
	isSet bool
}

func (v NullableRackReservationRequest) Get() *RackReservationRequest {
	return v.value
}

func (v *NullableRackReservationRequest) Set(val *RackReservationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRackReservationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRackReservationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRackReservationRequest(val *RackReservationRequest) *NullableRackReservationRequest {
	return &NullableRackReservationRequest{value: val, isSet: true}
}

func (v NullableRackReservationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRackReservationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


