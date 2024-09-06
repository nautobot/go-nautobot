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

// checks if the PatchedBulkWritableContactAssociationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedBulkWritableContactAssociationRequest{}

// PatchedBulkWritableContactAssociationRequest Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type PatchedBulkWritableContactAssociationRequest struct {
	Id string `json:"id"`
	AssociatedObjectType *string `json:"associated_object_type,omitempty"`
	AssociatedObjectId *string `json:"associated_object_id,omitempty"`
	Contact NullableBulkWritableCircuitRequestTenant `json:"contact,omitempty"`
	Team NullableBulkWritableCircuitRequestTenant `json:"team,omitempty"`
	Role *BulkWritableCableRequestStatus `json:"role,omitempty"`
	Status *BulkWritableCableRequestStatus `json:"status,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Relationships *map[string]BulkWritableCableRequestRelationshipsValue `json:"relationships,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedBulkWritableContactAssociationRequest PatchedBulkWritableContactAssociationRequest

// NewPatchedBulkWritableContactAssociationRequest instantiates a new PatchedBulkWritableContactAssociationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedBulkWritableContactAssociationRequest(id string) *PatchedBulkWritableContactAssociationRequest {
	this := PatchedBulkWritableContactAssociationRequest{}
	this.Id = id
	return &this
}

// NewPatchedBulkWritableContactAssociationRequestWithDefaults instantiates a new PatchedBulkWritableContactAssociationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedBulkWritableContactAssociationRequestWithDefaults() *PatchedBulkWritableContactAssociationRequest {
	this := PatchedBulkWritableContactAssociationRequest{}
	return &this
}

// GetId returns the Id field value
func (o *PatchedBulkWritableContactAssociationRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableContactAssociationRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PatchedBulkWritableContactAssociationRequest) SetId(v string) {
	o.Id = v
}

// GetAssociatedObjectType returns the AssociatedObjectType field value if set, zero value otherwise.
func (o *PatchedBulkWritableContactAssociationRequest) GetAssociatedObjectType() string {
	if o == nil || IsNil(o.AssociatedObjectType) {
		var ret string
		return ret
	}
	return *o.AssociatedObjectType
}

// GetAssociatedObjectTypeOk returns a tuple with the AssociatedObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableContactAssociationRequest) GetAssociatedObjectTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AssociatedObjectType) {
		return nil, false
	}
	return o.AssociatedObjectType, true
}

// HasAssociatedObjectType returns a boolean if a field has been set.
func (o *PatchedBulkWritableContactAssociationRequest) HasAssociatedObjectType() bool {
	if o != nil && !IsNil(o.AssociatedObjectType) {
		return true
	}

	return false
}

// SetAssociatedObjectType gets a reference to the given string and assigns it to the AssociatedObjectType field.
func (o *PatchedBulkWritableContactAssociationRequest) SetAssociatedObjectType(v string) {
	o.AssociatedObjectType = &v
}

// GetAssociatedObjectId returns the AssociatedObjectId field value if set, zero value otherwise.
func (o *PatchedBulkWritableContactAssociationRequest) GetAssociatedObjectId() string {
	if o == nil || IsNil(o.AssociatedObjectId) {
		var ret string
		return ret
	}
	return *o.AssociatedObjectId
}

// GetAssociatedObjectIdOk returns a tuple with the AssociatedObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableContactAssociationRequest) GetAssociatedObjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.AssociatedObjectId) {
		return nil, false
	}
	return o.AssociatedObjectId, true
}

// HasAssociatedObjectId returns a boolean if a field has been set.
func (o *PatchedBulkWritableContactAssociationRequest) HasAssociatedObjectId() bool {
	if o != nil && !IsNil(o.AssociatedObjectId) {
		return true
	}

	return false
}

// SetAssociatedObjectId gets a reference to the given string and assigns it to the AssociatedObjectId field.
func (o *PatchedBulkWritableContactAssociationRequest) SetAssociatedObjectId(v string) {
	o.AssociatedObjectId = &v
}

// GetContact returns the Contact field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBulkWritableContactAssociationRequest) GetContact() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.Contact.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.Contact.Get()
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBulkWritableContactAssociationRequest) GetContactOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Contact.Get(), o.Contact.IsSet()
}

// HasContact returns a boolean if a field has been set.
func (o *PatchedBulkWritableContactAssociationRequest) HasContact() bool {
	if o != nil && o.Contact.IsSet() {
		return true
	}

	return false
}

// SetContact gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the Contact field.
func (o *PatchedBulkWritableContactAssociationRequest) SetContact(v BulkWritableCircuitRequestTenant) {
	o.Contact.Set(&v)
}
// SetContactNil sets the value for Contact to be an explicit nil
func (o *PatchedBulkWritableContactAssociationRequest) SetContactNil() {
	o.Contact.Set(nil)
}

// UnsetContact ensures that no value is present for Contact, not even an explicit nil
func (o *PatchedBulkWritableContactAssociationRequest) UnsetContact() {
	o.Contact.Unset()
}

// GetTeam returns the Team field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBulkWritableContactAssociationRequest) GetTeam() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.Team.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.Team.Get()
}

// GetTeamOk returns a tuple with the Team field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBulkWritableContactAssociationRequest) GetTeamOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Team.Get(), o.Team.IsSet()
}

// HasTeam returns a boolean if a field has been set.
func (o *PatchedBulkWritableContactAssociationRequest) HasTeam() bool {
	if o != nil && o.Team.IsSet() {
		return true
	}

	return false
}

// SetTeam gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the Team field.
func (o *PatchedBulkWritableContactAssociationRequest) SetTeam(v BulkWritableCircuitRequestTenant) {
	o.Team.Set(&v)
}
// SetTeamNil sets the value for Team to be an explicit nil
func (o *PatchedBulkWritableContactAssociationRequest) SetTeamNil() {
	o.Team.Set(nil)
}

// UnsetTeam ensures that no value is present for Team, not even an explicit nil
func (o *PatchedBulkWritableContactAssociationRequest) UnsetTeam() {
	o.Team.Unset()
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *PatchedBulkWritableContactAssociationRequest) GetRole() BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Role) {
		var ret BulkWritableCableRequestStatus
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableContactAssociationRequest) GetRoleOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *PatchedBulkWritableContactAssociationRequest) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given BulkWritableCableRequestStatus and assigns it to the Role field.
func (o *PatchedBulkWritableContactAssociationRequest) SetRole(v BulkWritableCableRequestStatus) {
	o.Role = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PatchedBulkWritableContactAssociationRequest) GetStatus() BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Status) {
		var ret BulkWritableCableRequestStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableContactAssociationRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PatchedBulkWritableContactAssociationRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given BulkWritableCableRequestStatus and assigns it to the Status field.
func (o *PatchedBulkWritableContactAssociationRequest) SetStatus(v BulkWritableCableRequestStatus) {
	o.Status = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedBulkWritableContactAssociationRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableContactAssociationRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedBulkWritableContactAssociationRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedBulkWritableContactAssociationRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PatchedBulkWritableContactAssociationRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue {
	if o == nil || IsNil(o.Relationships) {
		var ret map[string]BulkWritableCableRequestRelationshipsValue
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableContactAssociationRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PatchedBulkWritableContactAssociationRequest) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]BulkWritableCableRequestRelationshipsValue and assigns it to the Relationships field.
func (o *PatchedBulkWritableContactAssociationRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue) {
	o.Relationships = &v
}

func (o PatchedBulkWritableContactAssociationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedBulkWritableContactAssociationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.AssociatedObjectType) {
		toSerialize["associated_object_type"] = o.AssociatedObjectType
	}
	if !IsNil(o.AssociatedObjectId) {
		toSerialize["associated_object_id"] = o.AssociatedObjectId
	}
	if o.Contact.IsSet() {
		toSerialize["contact"] = o.Contact.Get()
	}
	if o.Team.IsSet() {
		toSerialize["team"] = o.Team.Get()
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedBulkWritableContactAssociationRequest) UnmarshalJSON(data []byte) (err error) {
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

	varPatchedBulkWritableContactAssociationRequest := _PatchedBulkWritableContactAssociationRequest{}

	err = json.Unmarshal(data, &varPatchedBulkWritableContactAssociationRequest)

	if err != nil {
		return err
	}

	*o = PatchedBulkWritableContactAssociationRequest(varPatchedBulkWritableContactAssociationRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "associated_object_type")
		delete(additionalProperties, "associated_object_id")
		delete(additionalProperties, "contact")
		delete(additionalProperties, "team")
		delete(additionalProperties, "role")
		delete(additionalProperties, "status")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "relationships")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedBulkWritableContactAssociationRequest struct {
	value *PatchedBulkWritableContactAssociationRequest
	isSet bool
}

func (v NullablePatchedBulkWritableContactAssociationRequest) Get() *PatchedBulkWritableContactAssociationRequest {
	return v.value
}

func (v *NullablePatchedBulkWritableContactAssociationRequest) Set(val *PatchedBulkWritableContactAssociationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedBulkWritableContactAssociationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedBulkWritableContactAssociationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedBulkWritableContactAssociationRequest(val *PatchedBulkWritableContactAssociationRequest) *NullablePatchedBulkWritableContactAssociationRequest {
	return &NullablePatchedBulkWritableContactAssociationRequest{value: val, isSet: true}
}

func (v NullablePatchedBulkWritableContactAssociationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedBulkWritableContactAssociationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


