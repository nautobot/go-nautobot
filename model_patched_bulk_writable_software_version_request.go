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

// checks if the PatchedBulkWritableSoftwareVersionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedBulkWritableSoftwareVersionRequest{}

// PatchedBulkWritableSoftwareVersionRequest Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type PatchedBulkWritableSoftwareVersionRequest struct {
	Id string `json:"id"`
	Version *string `json:"version,omitempty"`
	// Optional alternative label for this version
	Alias *string `json:"alias,omitempty"`
	ReleaseDate NullableString `json:"release_date,omitempty"`
	EndOfSupportDate NullableString `json:"end_of_support_date,omitempty"`
	DocumentationUrl *string `json:"documentation_url,omitempty"`
	// Is a Long Term Support version
	LongTermSupport *bool `json:"long_term_support,omitempty"`
	// Is a Pre-Release version
	PreRelease *bool `json:"pre_release,omitempty"`
	Platform *BulkWritableCableRequestStatus `json:"platform,omitempty"`
	Status *BulkWritableCableRequestStatus `json:"status,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Relationships *map[string]BulkWritableCableRequestRelationshipsValue `json:"relationships,omitempty"`
	Tags []BulkWritableCableRequestStatus `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedBulkWritableSoftwareVersionRequest PatchedBulkWritableSoftwareVersionRequest

// NewPatchedBulkWritableSoftwareVersionRequest instantiates a new PatchedBulkWritableSoftwareVersionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedBulkWritableSoftwareVersionRequest(id string) *PatchedBulkWritableSoftwareVersionRequest {
	this := PatchedBulkWritableSoftwareVersionRequest{}
	this.Id = id
	return &this
}

// NewPatchedBulkWritableSoftwareVersionRequestWithDefaults instantiates a new PatchedBulkWritableSoftwareVersionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedBulkWritableSoftwareVersionRequestWithDefaults() *PatchedBulkWritableSoftwareVersionRequest {
	this := PatchedBulkWritableSoftwareVersionRequest{}
	return &this
}

// GetId returns the Id field value
func (o *PatchedBulkWritableSoftwareVersionRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableSoftwareVersionRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PatchedBulkWritableSoftwareVersionRequest) SetId(v string) {
	o.Id = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *PatchedBulkWritableSoftwareVersionRequest) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableSoftwareVersionRequest) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *PatchedBulkWritableSoftwareVersionRequest) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *PatchedBulkWritableSoftwareVersionRequest) SetVersion(v string) {
	o.Version = &v
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *PatchedBulkWritableSoftwareVersionRequest) GetAlias() string {
	if o == nil || IsNil(o.Alias) {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableSoftwareVersionRequest) GetAliasOk() (*string, bool) {
	if o == nil || IsNil(o.Alias) {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *PatchedBulkWritableSoftwareVersionRequest) HasAlias() bool {
	if o != nil && !IsNil(o.Alias) {
		return true
	}

	return false
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *PatchedBulkWritableSoftwareVersionRequest) SetAlias(v string) {
	o.Alias = &v
}

// GetReleaseDate returns the ReleaseDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBulkWritableSoftwareVersionRequest) GetReleaseDate() string {
	if o == nil || IsNil(o.ReleaseDate.Get()) {
		var ret string
		return ret
	}
	return *o.ReleaseDate.Get()
}

// GetReleaseDateOk returns a tuple with the ReleaseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBulkWritableSoftwareVersionRequest) GetReleaseDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReleaseDate.Get(), o.ReleaseDate.IsSet()
}

// HasReleaseDate returns a boolean if a field has been set.
func (o *PatchedBulkWritableSoftwareVersionRequest) HasReleaseDate() bool {
	if o != nil && o.ReleaseDate.IsSet() {
		return true
	}

	return false
}

// SetReleaseDate gets a reference to the given NullableString and assigns it to the ReleaseDate field.
func (o *PatchedBulkWritableSoftwareVersionRequest) SetReleaseDate(v string) {
	o.ReleaseDate.Set(&v)
}
// SetReleaseDateNil sets the value for ReleaseDate to be an explicit nil
func (o *PatchedBulkWritableSoftwareVersionRequest) SetReleaseDateNil() {
	o.ReleaseDate.Set(nil)
}

// UnsetReleaseDate ensures that no value is present for ReleaseDate, not even an explicit nil
func (o *PatchedBulkWritableSoftwareVersionRequest) UnsetReleaseDate() {
	o.ReleaseDate.Unset()
}

// GetEndOfSupportDate returns the EndOfSupportDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedBulkWritableSoftwareVersionRequest) GetEndOfSupportDate() string {
	if o == nil || IsNil(o.EndOfSupportDate.Get()) {
		var ret string
		return ret
	}
	return *o.EndOfSupportDate.Get()
}

// GetEndOfSupportDateOk returns a tuple with the EndOfSupportDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedBulkWritableSoftwareVersionRequest) GetEndOfSupportDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndOfSupportDate.Get(), o.EndOfSupportDate.IsSet()
}

// HasEndOfSupportDate returns a boolean if a field has been set.
func (o *PatchedBulkWritableSoftwareVersionRequest) HasEndOfSupportDate() bool {
	if o != nil && o.EndOfSupportDate.IsSet() {
		return true
	}

	return false
}

// SetEndOfSupportDate gets a reference to the given NullableString and assigns it to the EndOfSupportDate field.
func (o *PatchedBulkWritableSoftwareVersionRequest) SetEndOfSupportDate(v string) {
	o.EndOfSupportDate.Set(&v)
}
// SetEndOfSupportDateNil sets the value for EndOfSupportDate to be an explicit nil
func (o *PatchedBulkWritableSoftwareVersionRequest) SetEndOfSupportDateNil() {
	o.EndOfSupportDate.Set(nil)
}

// UnsetEndOfSupportDate ensures that no value is present for EndOfSupportDate, not even an explicit nil
func (o *PatchedBulkWritableSoftwareVersionRequest) UnsetEndOfSupportDate() {
	o.EndOfSupportDate.Unset()
}

// GetDocumentationUrl returns the DocumentationUrl field value if set, zero value otherwise.
func (o *PatchedBulkWritableSoftwareVersionRequest) GetDocumentationUrl() string {
	if o == nil || IsNil(o.DocumentationUrl) {
		var ret string
		return ret
	}
	return *o.DocumentationUrl
}

// GetDocumentationUrlOk returns a tuple with the DocumentationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableSoftwareVersionRequest) GetDocumentationUrlOk() (*string, bool) {
	if o == nil || IsNil(o.DocumentationUrl) {
		return nil, false
	}
	return o.DocumentationUrl, true
}

// HasDocumentationUrl returns a boolean if a field has been set.
func (o *PatchedBulkWritableSoftwareVersionRequest) HasDocumentationUrl() bool {
	if o != nil && !IsNil(o.DocumentationUrl) {
		return true
	}

	return false
}

// SetDocumentationUrl gets a reference to the given string and assigns it to the DocumentationUrl field.
func (o *PatchedBulkWritableSoftwareVersionRequest) SetDocumentationUrl(v string) {
	o.DocumentationUrl = &v
}

// GetLongTermSupport returns the LongTermSupport field value if set, zero value otherwise.
func (o *PatchedBulkWritableSoftwareVersionRequest) GetLongTermSupport() bool {
	if o == nil || IsNil(o.LongTermSupport) {
		var ret bool
		return ret
	}
	return *o.LongTermSupport
}

// GetLongTermSupportOk returns a tuple with the LongTermSupport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableSoftwareVersionRequest) GetLongTermSupportOk() (*bool, bool) {
	if o == nil || IsNil(o.LongTermSupport) {
		return nil, false
	}
	return o.LongTermSupport, true
}

// HasLongTermSupport returns a boolean if a field has been set.
func (o *PatchedBulkWritableSoftwareVersionRequest) HasLongTermSupport() bool {
	if o != nil && !IsNil(o.LongTermSupport) {
		return true
	}

	return false
}

// SetLongTermSupport gets a reference to the given bool and assigns it to the LongTermSupport field.
func (o *PatchedBulkWritableSoftwareVersionRequest) SetLongTermSupport(v bool) {
	o.LongTermSupport = &v
}

// GetPreRelease returns the PreRelease field value if set, zero value otherwise.
func (o *PatchedBulkWritableSoftwareVersionRequest) GetPreRelease() bool {
	if o == nil || IsNil(o.PreRelease) {
		var ret bool
		return ret
	}
	return *o.PreRelease
}

// GetPreReleaseOk returns a tuple with the PreRelease field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableSoftwareVersionRequest) GetPreReleaseOk() (*bool, bool) {
	if o == nil || IsNil(o.PreRelease) {
		return nil, false
	}
	return o.PreRelease, true
}

// HasPreRelease returns a boolean if a field has been set.
func (o *PatchedBulkWritableSoftwareVersionRequest) HasPreRelease() bool {
	if o != nil && !IsNil(o.PreRelease) {
		return true
	}

	return false
}

// SetPreRelease gets a reference to the given bool and assigns it to the PreRelease field.
func (o *PatchedBulkWritableSoftwareVersionRequest) SetPreRelease(v bool) {
	o.PreRelease = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *PatchedBulkWritableSoftwareVersionRequest) GetPlatform() BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Platform) {
		var ret BulkWritableCableRequestStatus
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableSoftwareVersionRequest) GetPlatformOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Platform) {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *PatchedBulkWritableSoftwareVersionRequest) HasPlatform() bool {
	if o != nil && !IsNil(o.Platform) {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given BulkWritableCableRequestStatus and assigns it to the Platform field.
func (o *PatchedBulkWritableSoftwareVersionRequest) SetPlatform(v BulkWritableCableRequestStatus) {
	o.Platform = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PatchedBulkWritableSoftwareVersionRequest) GetStatus() BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Status) {
		var ret BulkWritableCableRequestStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableSoftwareVersionRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PatchedBulkWritableSoftwareVersionRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given BulkWritableCableRequestStatus and assigns it to the Status field.
func (o *PatchedBulkWritableSoftwareVersionRequest) SetStatus(v BulkWritableCableRequestStatus) {
	o.Status = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedBulkWritableSoftwareVersionRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableSoftwareVersionRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedBulkWritableSoftwareVersionRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedBulkWritableSoftwareVersionRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PatchedBulkWritableSoftwareVersionRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue {
	if o == nil || IsNil(o.Relationships) {
		var ret map[string]BulkWritableCableRequestRelationshipsValue
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableSoftwareVersionRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PatchedBulkWritableSoftwareVersionRequest) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]BulkWritableCableRequestRelationshipsValue and assigns it to the Relationships field.
func (o *PatchedBulkWritableSoftwareVersionRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue) {
	o.Relationships = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedBulkWritableSoftwareVersionRequest) GetTags() []BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Tags) {
		var ret []BulkWritableCableRequestStatus
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableSoftwareVersionRequest) GetTagsOk() ([]BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedBulkWritableSoftwareVersionRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []BulkWritableCableRequestStatus and assigns it to the Tags field.
func (o *PatchedBulkWritableSoftwareVersionRequest) SetTags(v []BulkWritableCableRequestStatus) {
	o.Tags = v
}

func (o PatchedBulkWritableSoftwareVersionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedBulkWritableSoftwareVersionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.Alias) {
		toSerialize["alias"] = o.Alias
	}
	if o.ReleaseDate.IsSet() {
		toSerialize["release_date"] = o.ReleaseDate.Get()
	}
	if o.EndOfSupportDate.IsSet() {
		toSerialize["end_of_support_date"] = o.EndOfSupportDate.Get()
	}
	if !IsNil(o.DocumentationUrl) {
		toSerialize["documentation_url"] = o.DocumentationUrl
	}
	if !IsNil(o.LongTermSupport) {
		toSerialize["long_term_support"] = o.LongTermSupport
	}
	if !IsNil(o.PreRelease) {
		toSerialize["pre_release"] = o.PreRelease
	}
	if !IsNil(o.Platform) {
		toSerialize["platform"] = o.Platform
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

func (o *PatchedBulkWritableSoftwareVersionRequest) UnmarshalJSON(data []byte) (err error) {
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

	varPatchedBulkWritableSoftwareVersionRequest := _PatchedBulkWritableSoftwareVersionRequest{}

	err = json.Unmarshal(data, &varPatchedBulkWritableSoftwareVersionRequest)

	if err != nil {
		return err
	}

	*o = PatchedBulkWritableSoftwareVersionRequest(varPatchedBulkWritableSoftwareVersionRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "version")
		delete(additionalProperties, "alias")
		delete(additionalProperties, "release_date")
		delete(additionalProperties, "end_of_support_date")
		delete(additionalProperties, "documentation_url")
		delete(additionalProperties, "long_term_support")
		delete(additionalProperties, "pre_release")
		delete(additionalProperties, "platform")
		delete(additionalProperties, "status")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "relationships")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedBulkWritableSoftwareVersionRequest struct {
	value *PatchedBulkWritableSoftwareVersionRequest
	isSet bool
}

func (v NullablePatchedBulkWritableSoftwareVersionRequest) Get() *PatchedBulkWritableSoftwareVersionRequest {
	return v.value
}

func (v *NullablePatchedBulkWritableSoftwareVersionRequest) Set(val *PatchedBulkWritableSoftwareVersionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedBulkWritableSoftwareVersionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedBulkWritableSoftwareVersionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedBulkWritableSoftwareVersionRequest(val *PatchedBulkWritableSoftwareVersionRequest) *NullablePatchedBulkWritableSoftwareVersionRequest {
	return &NullablePatchedBulkWritableSoftwareVersionRequest{value: val, isSet: true}
}

func (v NullablePatchedBulkWritableSoftwareVersionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedBulkWritableSoftwareVersionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


