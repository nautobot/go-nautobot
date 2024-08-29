/*
API Documentation

Source of truth and network automation platform

API version: 2.3.1 (2.3)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nautobot

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the ObjectMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObjectMetadata{}

// ObjectMetadata Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type ObjectMetadata struct {
	Id string `json:"id"`
	ObjectType string `json:"object_type" validate:"regexp=^[a-z][a-z0-9_]+\\\\.[a-z][a-z0-9_]+$"`
	// Human friendly display value
	Display string `json:"display"`
	Url string `json:"url"`
	NaturalSlug string `json:"natural_slug"`
	AssignedObjectType string `json:"assigned_object_type"`
	AssignedObject NullableObjectMetadataAssignedObject `json:"assigned_object"`
	Value interface{} `json:"value,omitempty"`
	// List of scoped fields, only direct fields on the model
	ScopedFields interface{} `json:"scoped_fields,omitempty"`
	AssignedObjectId string `json:"assigned_object_id"`
	MetadataType BulkWritableCableRequestStatus `json:"metadata_type"`
	Contact NullableBulkWritableCircuitRequestTenant `json:"contact,omitempty"`
	Team NullableBulkWritableCircuitRequestTenant `json:"team,omitempty"`
	Created NullableTime `json:"created"`
	LastUpdated NullableTime `json:"last_updated"`
	AdditionalProperties map[string]interface{}
}

type _ObjectMetadata ObjectMetadata

// NewObjectMetadata instantiates a new ObjectMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectMetadata(id string, objectType string, display string, url string, naturalSlug string, assignedObjectType string, assignedObject NullableObjectMetadataAssignedObject, assignedObjectId string, metadataType BulkWritableCableRequestStatus, created NullableTime, lastUpdated NullableTime) *ObjectMetadata {
	this := ObjectMetadata{}
	this.Id = id
	this.ObjectType = objectType
	this.Display = display
	this.Url = url
	this.NaturalSlug = naturalSlug
	this.AssignedObjectType = assignedObjectType
	this.AssignedObject = assignedObject
	this.AssignedObjectId = assignedObjectId
	this.MetadataType = metadataType
	this.Created = created
	this.LastUpdated = lastUpdated
	return &this
}

// NewObjectMetadataWithDefaults instantiates a new ObjectMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectMetadataWithDefaults() *ObjectMetadata {
	this := ObjectMetadata{}
	return &this
}

// GetId returns the Id field value
func (o *ObjectMetadata) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObjectMetadata) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ObjectMetadata) SetId(v string) {
	o.Id = v
}

// GetObjectType returns the ObjectType field value
func (o *ObjectMetadata) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ObjectMetadata) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ObjectMetadata) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDisplay returns the Display field value
func (o *ObjectMetadata) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *ObjectMetadata) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *ObjectMetadata) SetDisplay(v string) {
	o.Display = v
}

// GetUrl returns the Url field value
func (o *ObjectMetadata) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ObjectMetadata) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ObjectMetadata) SetUrl(v string) {
	o.Url = v
}

// GetNaturalSlug returns the NaturalSlug field value
func (o *ObjectMetadata) GetNaturalSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NaturalSlug
}

// GetNaturalSlugOk returns a tuple with the NaturalSlug field value
// and a boolean to check if the value has been set.
func (o *ObjectMetadata) GetNaturalSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NaturalSlug, true
}

// SetNaturalSlug sets field value
func (o *ObjectMetadata) SetNaturalSlug(v string) {
	o.NaturalSlug = v
}

// GetAssignedObjectType returns the AssignedObjectType field value
func (o *ObjectMetadata) GetAssignedObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssignedObjectType
}

// GetAssignedObjectTypeOk returns a tuple with the AssignedObjectType field value
// and a boolean to check if the value has been set.
func (o *ObjectMetadata) GetAssignedObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssignedObjectType, true
}

// SetAssignedObjectType sets field value
func (o *ObjectMetadata) SetAssignedObjectType(v string) {
	o.AssignedObjectType = v
}

// GetAssignedObject returns the AssignedObject field value
// If the value is explicit nil, the zero value for ObjectMetadataAssignedObject will be returned
func (o *ObjectMetadata) GetAssignedObject() ObjectMetadataAssignedObject {
	if o == nil || o.AssignedObject.Get() == nil {
		var ret ObjectMetadataAssignedObject
		return ret
	}

	return *o.AssignedObject.Get()
}

// GetAssignedObjectOk returns a tuple with the AssignedObject field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ObjectMetadata) GetAssignedObjectOk() (*ObjectMetadataAssignedObject, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssignedObject.Get(), o.AssignedObject.IsSet()
}

// SetAssignedObject sets field value
func (o *ObjectMetadata) SetAssignedObject(v ObjectMetadataAssignedObject) {
	o.AssignedObject.Set(&v)
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ObjectMetadata) GetValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ObjectMetadata) GetValueOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return &o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ObjectMetadata) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given interface{} and assigns it to the Value field.
func (o *ObjectMetadata) SetValue(v interface{}) {
	o.Value = v
}

// GetScopedFields returns the ScopedFields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ObjectMetadata) GetScopedFields() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ScopedFields
}

// GetScopedFieldsOk returns a tuple with the ScopedFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ObjectMetadata) GetScopedFieldsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ScopedFields) {
		return nil, false
	}
	return &o.ScopedFields, true
}

// HasScopedFields returns a boolean if a field has been set.
func (o *ObjectMetadata) HasScopedFields() bool {
	if o != nil && !IsNil(o.ScopedFields) {
		return true
	}

	return false
}

// SetScopedFields gets a reference to the given interface{} and assigns it to the ScopedFields field.
func (o *ObjectMetadata) SetScopedFields(v interface{}) {
	o.ScopedFields = v
}

// GetAssignedObjectId returns the AssignedObjectId field value
func (o *ObjectMetadata) GetAssignedObjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssignedObjectId
}

// GetAssignedObjectIdOk returns a tuple with the AssignedObjectId field value
// and a boolean to check if the value has been set.
func (o *ObjectMetadata) GetAssignedObjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssignedObjectId, true
}

// SetAssignedObjectId sets field value
func (o *ObjectMetadata) SetAssignedObjectId(v string) {
	o.AssignedObjectId = v
}

// GetMetadataType returns the MetadataType field value
func (o *ObjectMetadata) GetMetadataType() BulkWritableCableRequestStatus {
	if o == nil {
		var ret BulkWritableCableRequestStatus
		return ret
	}

	return o.MetadataType
}

// GetMetadataTypeOk returns a tuple with the MetadataType field value
// and a boolean to check if the value has been set.
func (o *ObjectMetadata) GetMetadataTypeOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetadataType, true
}

// SetMetadataType sets field value
func (o *ObjectMetadata) SetMetadataType(v BulkWritableCableRequestStatus) {
	o.MetadataType = v
}

// GetContact returns the Contact field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ObjectMetadata) GetContact() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.Contact.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.Contact.Get()
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ObjectMetadata) GetContactOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Contact.Get(), o.Contact.IsSet()
}

// HasContact returns a boolean if a field has been set.
func (o *ObjectMetadata) HasContact() bool {
	if o != nil && o.Contact.IsSet() {
		return true
	}

	return false
}

// SetContact gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the Contact field.
func (o *ObjectMetadata) SetContact(v BulkWritableCircuitRequestTenant) {
	o.Contact.Set(&v)
}
// SetContactNil sets the value for Contact to be an explicit nil
func (o *ObjectMetadata) SetContactNil() {
	o.Contact.Set(nil)
}

// UnsetContact ensures that no value is present for Contact, not even an explicit nil
func (o *ObjectMetadata) UnsetContact() {
	o.Contact.Unset()
}

// GetTeam returns the Team field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ObjectMetadata) GetTeam() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.Team.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.Team.Get()
}

// GetTeamOk returns a tuple with the Team field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ObjectMetadata) GetTeamOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Team.Get(), o.Team.IsSet()
}

// HasTeam returns a boolean if a field has been set.
func (o *ObjectMetadata) HasTeam() bool {
	if o != nil && o.Team.IsSet() {
		return true
	}

	return false
}

// SetTeam gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the Team field.
func (o *ObjectMetadata) SetTeam(v BulkWritableCircuitRequestTenant) {
	o.Team.Set(&v)
}
// SetTeamNil sets the value for Team to be an explicit nil
func (o *ObjectMetadata) SetTeamNil() {
	o.Team.Set(nil)
}

// UnsetTeam ensures that no value is present for Team, not even an explicit nil
func (o *ObjectMetadata) UnsetTeam() {
	o.Team.Unset()
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *ObjectMetadata) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ObjectMetadata) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *ObjectMetadata) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *ObjectMetadata) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ObjectMetadata) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *ObjectMetadata) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

func (o ObjectMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObjectMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object_type"] = o.ObjectType
	toSerialize["display"] = o.Display
	toSerialize["url"] = o.Url
	toSerialize["natural_slug"] = o.NaturalSlug
	toSerialize["assigned_object_type"] = o.AssignedObjectType
	toSerialize["assigned_object"] = o.AssignedObject.Get()
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.ScopedFields != nil {
		toSerialize["scoped_fields"] = o.ScopedFields
	}
	toSerialize["assigned_object_id"] = o.AssignedObjectId
	toSerialize["metadata_type"] = o.MetadataType
	if o.Contact.IsSet() {
		toSerialize["contact"] = o.Contact.Get()
	}
	if o.Team.IsSet() {
		toSerialize["team"] = o.Team.Get()
	}
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ObjectMetadata) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"object_type",
		"display",
		"url",
		"natural_slug",
		"assigned_object_type",
		"assigned_object",
		"assigned_object_id",
		"metadata_type",
		"created",
		"last_updated",
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

	varObjectMetadata := _ObjectMetadata{}

	err = json.Unmarshal(data, &varObjectMetadata)

	if err != nil {
		return err
	}

	*o = ObjectMetadata(varObjectMetadata)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "object_type")
		delete(additionalProperties, "display")
		delete(additionalProperties, "url")
		delete(additionalProperties, "natural_slug")
		delete(additionalProperties, "assigned_object_type")
		delete(additionalProperties, "assigned_object")
		delete(additionalProperties, "value")
		delete(additionalProperties, "scoped_fields")
		delete(additionalProperties, "assigned_object_id")
		delete(additionalProperties, "metadata_type")
		delete(additionalProperties, "contact")
		delete(additionalProperties, "team")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableObjectMetadata struct {
	value *ObjectMetadata
	isSet bool
}

func (v NullableObjectMetadata) Get() *ObjectMetadata {
	return v.value
}

func (v *NullableObjectMetadata) Set(val *ObjectMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectMetadata(val *ObjectMetadata) *NullableObjectMetadata {
	return &NullableObjectMetadata{value: val, isSet: true}
}

func (v NullableObjectMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

