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

// checks if the PatchedDynamicGroupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedDynamicGroupRequest{}

// PatchedDynamicGroupRequest Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type PatchedDynamicGroupRequest struct {
	ContentType *string `json:"content_type,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	GroupType *GroupTypeEnum `json:"group_type,omitempty"`
	// A JSON-encoded dictionary of filter parameters defining membership of this group
	Filter interface{} `json:"filter,omitempty"`
	Tenant NullableBulkWritableCircuitRequestTenant `json:"tenant,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Relationships *map[string]BulkWritableCableRequestRelationshipsValue `json:"relationships,omitempty"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedDynamicGroupRequest PatchedDynamicGroupRequest

// NewPatchedDynamicGroupRequest instantiates a new PatchedDynamicGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedDynamicGroupRequest() *PatchedDynamicGroupRequest {
	this := PatchedDynamicGroupRequest{}
	return &this
}

// NewPatchedDynamicGroupRequestWithDefaults instantiates a new PatchedDynamicGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedDynamicGroupRequestWithDefaults() *PatchedDynamicGroupRequest {
	this := PatchedDynamicGroupRequest{}
	return &this
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *PatchedDynamicGroupRequest) GetContentType() string {
	if o == nil || IsNil(o.ContentType) {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedDynamicGroupRequest) GetContentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ContentType) {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *PatchedDynamicGroupRequest) HasContentType() bool {
	if o != nil && !IsNil(o.ContentType) {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *PatchedDynamicGroupRequest) SetContentType(v string) {
	o.ContentType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedDynamicGroupRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedDynamicGroupRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedDynamicGroupRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedDynamicGroupRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedDynamicGroupRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedDynamicGroupRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedDynamicGroupRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedDynamicGroupRequest) SetDescription(v string) {
	o.Description = &v
}

// GetGroupType returns the GroupType field value if set, zero value otherwise.
func (o *PatchedDynamicGroupRequest) GetGroupType() GroupTypeEnum {
	if o == nil || IsNil(o.GroupType) {
		var ret GroupTypeEnum
		return ret
	}
	return *o.GroupType
}

// GetGroupTypeOk returns a tuple with the GroupType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedDynamicGroupRequest) GetGroupTypeOk() (*GroupTypeEnum, bool) {
	if o == nil || IsNil(o.GroupType) {
		return nil, false
	}
	return o.GroupType, true
}

// HasGroupType returns a boolean if a field has been set.
func (o *PatchedDynamicGroupRequest) HasGroupType() bool {
	if o != nil && !IsNil(o.GroupType) {
		return true
	}

	return false
}

// SetGroupType gets a reference to the given GroupTypeEnum and assigns it to the GroupType field.
func (o *PatchedDynamicGroupRequest) SetGroupType(v GroupTypeEnum) {
	o.GroupType = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedDynamicGroupRequest) GetFilter() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedDynamicGroupRequest) GetFilterOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Filter) {
		return nil, false
	}
	return &o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *PatchedDynamicGroupRequest) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given interface{} and assigns it to the Filter field.
func (o *PatchedDynamicGroupRequest) SetFilter(v interface{}) {
	o.Filter = v
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedDynamicGroupRequest) GetTenant() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedDynamicGroupRequest) GetTenantOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *PatchedDynamicGroupRequest) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the Tenant field.
func (o *PatchedDynamicGroupRequest) SetTenant(v BulkWritableCircuitRequestTenant) {
	o.Tenant.Set(&v)
}
// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *PatchedDynamicGroupRequest) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *PatchedDynamicGroupRequest) UnsetTenant() {
	o.Tenant.Unset()
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedDynamicGroupRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedDynamicGroupRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedDynamicGroupRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedDynamicGroupRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PatchedDynamicGroupRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue {
	if o == nil || IsNil(o.Relationships) {
		var ret map[string]BulkWritableCableRequestRelationshipsValue
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedDynamicGroupRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PatchedDynamicGroupRequest) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]BulkWritableCableRequestRelationshipsValue and assigns it to the Relationships field.
func (o *PatchedDynamicGroupRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue) {
	o.Relationships = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedDynamicGroupRequest) GetTags() []BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Tags) {
		var ret []BulkWritableCableRequestStatus
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedDynamicGroupRequest) GetTagsOk() ([]BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedDynamicGroupRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []BulkWritableCableRequestStatus and assigns it to the Tags field.
func (o *PatchedDynamicGroupRequest) SetTags(v []BulkWritableCableRequestStatus) {
	o.Tags = v
}

func (o PatchedDynamicGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedDynamicGroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContentType) {
		toSerialize["content_type"] = o.ContentType
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.GroupType) {
		toSerialize["group_type"] = o.GroupType
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
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

func (o *PatchedDynamicGroupRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedDynamicGroupRequest := _PatchedDynamicGroupRequest{}

	err = json.Unmarshal(data, &varPatchedDynamicGroupRequest)

	if err != nil {
		return err
	}

	*o = PatchedDynamicGroupRequest(varPatchedDynamicGroupRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "content_type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "group_type")
		delete(additionalProperties, "filter")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "relationships")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedDynamicGroupRequest struct {
	value *PatchedDynamicGroupRequest
	isSet bool
}

func (v NullablePatchedDynamicGroupRequest) Get() *PatchedDynamicGroupRequest {
	return v.value
}

func (v *NullablePatchedDynamicGroupRequest) Set(val *PatchedDynamicGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedDynamicGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedDynamicGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedDynamicGroupRequest(val *PatchedDynamicGroupRequest) *NullablePatchedDynamicGroupRequest {
	return &NullablePatchedDynamicGroupRequest{value: val, isSet: true}
}

func (v NullablePatchedDynamicGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedDynamicGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


