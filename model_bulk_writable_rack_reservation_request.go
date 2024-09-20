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

// checks if the BulkWritableRackReservationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkWritableRackReservationRequest{}

// BulkWritableRackReservationRequest Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type BulkWritableRackReservationRequest struct {
	Id string `json:"id"`
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

type _BulkWritableRackReservationRequest BulkWritableRackReservationRequest

// NewBulkWritableRackReservationRequest instantiates a new BulkWritableRackReservationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkWritableRackReservationRequest(id string, units interface{}, description string, rack BulkWritableCableRequestStatus) *BulkWritableRackReservationRequest {
	this := BulkWritableRackReservationRequest{}
	this.Id = id
	this.Units = units
	this.Description = description
	this.Rack = rack
	return &this
}

// NewBulkWritableRackReservationRequestWithDefaults instantiates a new BulkWritableRackReservationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkWritableRackReservationRequestWithDefaults() *BulkWritableRackReservationRequest {
	this := BulkWritableRackReservationRequest{}
	return &this
}

// GetId returns the Id field value
func (o *BulkWritableRackReservationRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BulkWritableRackReservationRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BulkWritableRackReservationRequest) SetId(v string) {
	o.Id = v
}

// GetUnits returns the Units field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *BulkWritableRackReservationRequest) GetUnits() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Units
}

// GetUnitsOk returns a tuple with the Units field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkWritableRackReservationRequest) GetUnitsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Units) {
		return nil, false
	}
	return &o.Units, true
}

// SetUnits sets field value
func (o *BulkWritableRackReservationRequest) SetUnits(v interface{}) {
	o.Units = v
}

// GetDescription returns the Description field value
func (o *BulkWritableRackReservationRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *BulkWritableRackReservationRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *BulkWritableRackReservationRequest) SetDescription(v string) {
	o.Description = v
}

// GetRack returns the Rack field value
func (o *BulkWritableRackReservationRequest) GetRack() BulkWritableCableRequestStatus {
	if o == nil {
		var ret BulkWritableCableRequestStatus
		return ret
	}

	return o.Rack
}

// GetRackOk returns a tuple with the Rack field value
// and a boolean to check if the value has been set.
func (o *BulkWritableRackReservationRequest) GetRackOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rack, true
}

// SetRack sets field value
func (o *BulkWritableRackReservationRequest) SetRack(v BulkWritableCableRequestStatus) {
	o.Rack = v
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkWritableRackReservationRequest) GetTenant() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkWritableRackReservationRequest) GetTenantOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *BulkWritableRackReservationRequest) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the Tenant field.
func (o *BulkWritableRackReservationRequest) SetTenant(v BulkWritableCircuitRequestTenant) {
	o.Tenant.Set(&v)
}
// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *BulkWritableRackReservationRequest) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *BulkWritableRackReservationRequest) UnsetTenant() {
	o.Tenant.Unset()
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *BulkWritableRackReservationRequest) GetUser() BulkWritableRackReservationRequestUser {
	if o == nil || IsNil(o.User) {
		var ret BulkWritableRackReservationRequestUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableRackReservationRequest) GetUserOk() (*BulkWritableRackReservationRequestUser, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *BulkWritableRackReservationRequest) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given BulkWritableRackReservationRequestUser and assigns it to the User field.
func (o *BulkWritableRackReservationRequest) SetUser(v BulkWritableRackReservationRequestUser) {
	o.User = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *BulkWritableRackReservationRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableRackReservationRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *BulkWritableRackReservationRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *BulkWritableRackReservationRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *BulkWritableRackReservationRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue {
	if o == nil || IsNil(o.Relationships) {
		var ret map[string]BulkWritableCableRequestRelationshipsValue
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableRackReservationRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *BulkWritableRackReservationRequest) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]BulkWritableCableRequestRelationshipsValue and assigns it to the Relationships field.
func (o *BulkWritableRackReservationRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue) {
	o.Relationships = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *BulkWritableRackReservationRequest) GetTags() []BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Tags) {
		var ret []BulkWritableCableRequestStatus
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableRackReservationRequest) GetTagsOk() ([]BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *BulkWritableRackReservationRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []BulkWritableCableRequestStatus and assigns it to the Tags field.
func (o *BulkWritableRackReservationRequest) SetTags(v []BulkWritableCableRequestStatus) {
	o.Tags = v
}

func (o BulkWritableRackReservationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkWritableRackReservationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
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

func (o *BulkWritableRackReservationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varBulkWritableRackReservationRequest := _BulkWritableRackReservationRequest{}

	err = json.Unmarshal(data, &varBulkWritableRackReservationRequest)

	if err != nil {
		return err
	}

	*o = BulkWritableRackReservationRequest(varBulkWritableRackReservationRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
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

type NullableBulkWritableRackReservationRequest struct {
	value *BulkWritableRackReservationRequest
	isSet bool
}

func (v NullableBulkWritableRackReservationRequest) Get() *BulkWritableRackReservationRequest {
	return v.value
}

func (v *NullableBulkWritableRackReservationRequest) Set(val *BulkWritableRackReservationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkWritableRackReservationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkWritableRackReservationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkWritableRackReservationRequest(val *BulkWritableRackReservationRequest) *NullableBulkWritableRackReservationRequest {
	return &NullableBulkWritableRackReservationRequest{value: val, isSet: true}
}

func (v NullableBulkWritableRackReservationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkWritableRackReservationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


