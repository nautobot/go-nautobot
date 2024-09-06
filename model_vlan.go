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

// checks if the VLAN type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VLAN{}

// VLAN Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type VLAN struct {
	Id string `json:"id"`
	ObjectType string `json:"object_type" validate:"regexp=^[a-z][a-z0-9_]+\\\\.[a-z][a-z0-9_]+$"`
	// Human friendly display value
	Display string `json:"display"`
	Url string `json:"url"`
	NaturalSlug string `json:"natural_slug"`
	PrefixCount *int32 `json:"prefix_count,omitempty"`
	Vid int32 `json:"vid"`
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	VlanGroup NullableBulkWritableCircuitRequestTenant `json:"vlan_group,omitempty"`
	Status BulkWritableCableRequestStatus `json:"status"`
	Role NullableBulkWritableCircuitRequestTenant `json:"role,omitempty"`
	Tenant NullableBulkWritableCircuitRequestTenant `json:"tenant,omitempty"`
	Locations []BulkWritableCableRequestStatus `json:"locations"`
	Created NullableTime `json:"created"`
	LastUpdated NullableTime `json:"last_updated"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	NotesUrl string `json:"notes_url"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VLAN VLAN

// NewVLAN instantiates a new VLAN object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVLAN(id string, objectType string, display string, url string, naturalSlug string, vid int32, name string, status BulkWritableCableRequestStatus, locations []BulkWritableCableRequestStatus, created NullableTime, lastUpdated NullableTime, notesUrl string) *VLAN {
	this := VLAN{}
	this.Id = id
	this.ObjectType = objectType
	this.Display = display
	this.Url = url
	this.NaturalSlug = naturalSlug
	this.Vid = vid
	this.Name = name
	this.Status = status
	this.Locations = locations
	this.Created = created
	this.LastUpdated = lastUpdated
	this.NotesUrl = notesUrl
	return &this
}

// NewVLANWithDefaults instantiates a new VLAN object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVLANWithDefaults() *VLAN {
	this := VLAN{}
	return &this
}

// GetId returns the Id field value
func (o *VLAN) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VLAN) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VLAN) SetId(v string) {
	o.Id = v
}

// GetObjectType returns the ObjectType field value
func (o *VLAN) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VLAN) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VLAN) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDisplay returns the Display field value
func (o *VLAN) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *VLAN) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *VLAN) SetDisplay(v string) {
	o.Display = v
}

// GetUrl returns the Url field value
func (o *VLAN) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *VLAN) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *VLAN) SetUrl(v string) {
	o.Url = v
}

// GetNaturalSlug returns the NaturalSlug field value
func (o *VLAN) GetNaturalSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NaturalSlug
}

// GetNaturalSlugOk returns a tuple with the NaturalSlug field value
// and a boolean to check if the value has been set.
func (o *VLAN) GetNaturalSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NaturalSlug, true
}

// SetNaturalSlug sets field value
func (o *VLAN) SetNaturalSlug(v string) {
	o.NaturalSlug = v
}

// GetPrefixCount returns the PrefixCount field value if set, zero value otherwise.
func (o *VLAN) GetPrefixCount() int32 {
	if o == nil || IsNil(o.PrefixCount) {
		var ret int32
		return ret
	}
	return *o.PrefixCount
}

// GetPrefixCountOk returns a tuple with the PrefixCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VLAN) GetPrefixCountOk() (*int32, bool) {
	if o == nil || IsNil(o.PrefixCount) {
		return nil, false
	}
	return o.PrefixCount, true
}

// HasPrefixCount returns a boolean if a field has been set.
func (o *VLAN) HasPrefixCount() bool {
	if o != nil && !IsNil(o.PrefixCount) {
		return true
	}

	return false
}

// SetPrefixCount gets a reference to the given int32 and assigns it to the PrefixCount field.
func (o *VLAN) SetPrefixCount(v int32) {
	o.PrefixCount = &v
}

// GetVid returns the Vid field value
func (o *VLAN) GetVid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Vid
}

// GetVidOk returns a tuple with the Vid field value
// and a boolean to check if the value has been set.
func (o *VLAN) GetVidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vid, true
}

// SetVid sets field value
func (o *VLAN) SetVid(v int32) {
	o.Vid = v
}

// GetName returns the Name field value
func (o *VLAN) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VLAN) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VLAN) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VLAN) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VLAN) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VLAN) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VLAN) SetDescription(v string) {
	o.Description = &v
}

// GetVlanGroup returns the VlanGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VLAN) GetVlanGroup() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.VlanGroup.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.VlanGroup.Get()
}

// GetVlanGroupOk returns a tuple with the VlanGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VLAN) GetVlanGroupOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.VlanGroup.Get(), o.VlanGroup.IsSet()
}

// HasVlanGroup returns a boolean if a field has been set.
func (o *VLAN) HasVlanGroup() bool {
	if o != nil && o.VlanGroup.IsSet() {
		return true
	}

	return false
}

// SetVlanGroup gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the VlanGroup field.
func (o *VLAN) SetVlanGroup(v BulkWritableCircuitRequestTenant) {
	o.VlanGroup.Set(&v)
}
// SetVlanGroupNil sets the value for VlanGroup to be an explicit nil
func (o *VLAN) SetVlanGroupNil() {
	o.VlanGroup.Set(nil)
}

// UnsetVlanGroup ensures that no value is present for VlanGroup, not even an explicit nil
func (o *VLAN) UnsetVlanGroup() {
	o.VlanGroup.Unset()
}

// GetStatus returns the Status field value
func (o *VLAN) GetStatus() BulkWritableCableRequestStatus {
	if o == nil {
		var ret BulkWritableCableRequestStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *VLAN) GetStatusOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *VLAN) SetStatus(v BulkWritableCableRequestStatus) {
	o.Status = v
}

// GetRole returns the Role field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VLAN) GetRole() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.Role.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.Role.Get()
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VLAN) GetRoleOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Role.Get(), o.Role.IsSet()
}

// HasRole returns a boolean if a field has been set.
func (o *VLAN) HasRole() bool {
	if o != nil && o.Role.IsSet() {
		return true
	}

	return false
}

// SetRole gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the Role field.
func (o *VLAN) SetRole(v BulkWritableCircuitRequestTenant) {
	o.Role.Set(&v)
}
// SetRoleNil sets the value for Role to be an explicit nil
func (o *VLAN) SetRoleNil() {
	o.Role.Set(nil)
}

// UnsetRole ensures that no value is present for Role, not even an explicit nil
func (o *VLAN) UnsetRole() {
	o.Role.Unset()
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VLAN) GetTenant() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VLAN) GetTenantOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *VLAN) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the Tenant field.
func (o *VLAN) SetTenant(v BulkWritableCircuitRequestTenant) {
	o.Tenant.Set(&v)
}
// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *VLAN) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *VLAN) UnsetTenant() {
	o.Tenant.Unset()
}

// GetLocations returns the Locations field value
func (o *VLAN) GetLocations() []BulkWritableCableRequestStatus {
	if o == nil {
		var ret []BulkWritableCableRequestStatus
		return ret
	}

	return o.Locations
}

// GetLocationsOk returns a tuple with the Locations field value
// and a boolean to check if the value has been set.
func (o *VLAN) GetLocationsOk() ([]BulkWritableCableRequestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.Locations, true
}

// SetLocations sets field value
func (o *VLAN) SetLocations(v []BulkWritableCableRequestStatus) {
	o.Locations = v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *VLAN) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VLAN) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *VLAN) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *VLAN) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VLAN) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *VLAN) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *VLAN) GetTags() []BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Tags) {
		var ret []BulkWritableCableRequestStatus
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VLAN) GetTagsOk() ([]BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *VLAN) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []BulkWritableCableRequestStatus and assigns it to the Tags field.
func (o *VLAN) SetTags(v []BulkWritableCableRequestStatus) {
	o.Tags = v
}

// GetNotesUrl returns the NotesUrl field value
func (o *VLAN) GetNotesUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotesUrl
}

// GetNotesUrlOk returns a tuple with the NotesUrl field value
// and a boolean to check if the value has been set.
func (o *VLAN) GetNotesUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotesUrl, true
}

// SetNotesUrl sets field value
func (o *VLAN) SetNotesUrl(v string) {
	o.NotesUrl = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *VLAN) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VLAN) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *VLAN) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *VLAN) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o VLAN) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VLAN) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object_type"] = o.ObjectType
	toSerialize["display"] = o.Display
	toSerialize["url"] = o.Url
	toSerialize["natural_slug"] = o.NaturalSlug
	if !IsNil(o.PrefixCount) {
		toSerialize["prefix_count"] = o.PrefixCount
	}
	toSerialize["vid"] = o.Vid
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.VlanGroup.IsSet() {
		toSerialize["vlan_group"] = o.VlanGroup.Get()
	}
	toSerialize["status"] = o.Status
	if o.Role.IsSet() {
		toSerialize["role"] = o.Role.Get()
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	toSerialize["locations"] = o.Locations
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["notes_url"] = o.NotesUrl
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VLAN) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"object_type",
		"display",
		"url",
		"natural_slug",
		"vid",
		"name",
		"status",
		"locations",
		"created",
		"last_updated",
		"notes_url",
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

	varVLAN := _VLAN{}

	err = json.Unmarshal(data, &varVLAN)

	if err != nil {
		return err
	}

	*o = VLAN(varVLAN)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "object_type")
		delete(additionalProperties, "display")
		delete(additionalProperties, "url")
		delete(additionalProperties, "natural_slug")
		delete(additionalProperties, "prefix_count")
		delete(additionalProperties, "vid")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "vlan_group")
		delete(additionalProperties, "status")
		delete(additionalProperties, "role")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "locations")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "notes_url")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVLAN struct {
	value *VLAN
	isSet bool
}

func (v NullableVLAN) Get() *VLAN {
	return v.value
}

func (v *NullableVLAN) Set(val *VLAN) {
	v.value = val
	v.isSet = true
}

func (v NullableVLAN) IsSet() bool {
	return v.isSet
}

func (v *NullableVLAN) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVLAN(val *VLAN) *NullableVLAN {
	return &NullableVLAN{value: val, isSet: true}
}

func (v NullableVLAN) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVLAN) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


