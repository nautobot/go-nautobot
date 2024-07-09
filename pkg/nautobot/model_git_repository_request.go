/*
API Documentation

Source of truth and network automation platform

API version: 2.2.5 (2.2)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nautobot

import (
	"encoding/json"
	"fmt"
)

// checks if the GitRepositoryRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GitRepositoryRequest{}

// GitRepositoryRequest Git repositories defined as a data source.
type GitRepositoryRequest struct {
	ProvidedContents []BulkWritableGitRepositoryRequestProvidedContentsInner `json:"provided_contents,omitempty"`
	Name string `json:"name"`
	// Internal field name. Please use underscores rather than dashes in this key.
	Slug *string `json:"slug,omitempty" validate:"regexp=^[-a-zA-Z0-9_]+$"`
	// Only HTTP and HTTPS URLs are presently supported
	RemoteUrl string `json:"remote_url"`
	Branch *string `json:"branch,omitempty"`
	// Commit hash of the most recent fetch from the selected branch. Used for syncing between workers.
	CurrentHead *string `json:"current_head,omitempty"`
	SecretsGroup NullableBulkWritableCircuitRequestTenant `json:"secrets_group,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Relationships *map[string]BulkWritableCableRequestRelationshipsValue `json:"relationships,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GitRepositoryRequest GitRepositoryRequest

// NewGitRepositoryRequest instantiates a new GitRepositoryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGitRepositoryRequest(name string, remoteUrl string) *GitRepositoryRequest {
	this := GitRepositoryRequest{}
	this.Name = name
	this.RemoteUrl = remoteUrl
	return &this
}

// NewGitRepositoryRequestWithDefaults instantiates a new GitRepositoryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGitRepositoryRequestWithDefaults() *GitRepositoryRequest {
	this := GitRepositoryRequest{}
	return &this
}

// GetProvidedContents returns the ProvidedContents field value if set, zero value otherwise.
func (o *GitRepositoryRequest) GetProvidedContents() []BulkWritableGitRepositoryRequestProvidedContentsInner {
	if o == nil || IsNil(o.ProvidedContents) {
		var ret []BulkWritableGitRepositoryRequestProvidedContentsInner
		return ret
	}
	return o.ProvidedContents
}

// GetProvidedContentsOk returns a tuple with the ProvidedContents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitRepositoryRequest) GetProvidedContentsOk() ([]BulkWritableGitRepositoryRequestProvidedContentsInner, bool) {
	if o == nil || IsNil(o.ProvidedContents) {
		return nil, false
	}
	return o.ProvidedContents, true
}

// HasProvidedContents returns a boolean if a field has been set.
func (o *GitRepositoryRequest) HasProvidedContents() bool {
	if o != nil && !IsNil(o.ProvidedContents) {
		return true
	}

	return false
}

// SetProvidedContents gets a reference to the given []BulkWritableGitRepositoryRequestProvidedContentsInner and assigns it to the ProvidedContents field.
func (o *GitRepositoryRequest) SetProvidedContents(v []BulkWritableGitRepositoryRequestProvidedContentsInner) {
	o.ProvidedContents = v
}

// GetName returns the Name field value
func (o *GitRepositoryRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GitRepositoryRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GitRepositoryRequest) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *GitRepositoryRequest) GetSlug() string {
	if o == nil || IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitRepositoryRequest) GetSlugOk() (*string, bool) {
	if o == nil || IsNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *GitRepositoryRequest) HasSlug() bool {
	if o != nil && !IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *GitRepositoryRequest) SetSlug(v string) {
	o.Slug = &v
}

// GetRemoteUrl returns the RemoteUrl field value
func (o *GitRepositoryRequest) GetRemoteUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RemoteUrl
}

// GetRemoteUrlOk returns a tuple with the RemoteUrl field value
// and a boolean to check if the value has been set.
func (o *GitRepositoryRequest) GetRemoteUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemoteUrl, true
}

// SetRemoteUrl sets field value
func (o *GitRepositoryRequest) SetRemoteUrl(v string) {
	o.RemoteUrl = v
}

// GetBranch returns the Branch field value if set, zero value otherwise.
func (o *GitRepositoryRequest) GetBranch() string {
	if o == nil || IsNil(o.Branch) {
		var ret string
		return ret
	}
	return *o.Branch
}

// GetBranchOk returns a tuple with the Branch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitRepositoryRequest) GetBranchOk() (*string, bool) {
	if o == nil || IsNil(o.Branch) {
		return nil, false
	}
	return o.Branch, true
}

// HasBranch returns a boolean if a field has been set.
func (o *GitRepositoryRequest) HasBranch() bool {
	if o != nil && !IsNil(o.Branch) {
		return true
	}

	return false
}

// SetBranch gets a reference to the given string and assigns it to the Branch field.
func (o *GitRepositoryRequest) SetBranch(v string) {
	o.Branch = &v
}

// GetCurrentHead returns the CurrentHead field value if set, zero value otherwise.
func (o *GitRepositoryRequest) GetCurrentHead() string {
	if o == nil || IsNil(o.CurrentHead) {
		var ret string
		return ret
	}
	return *o.CurrentHead
}

// GetCurrentHeadOk returns a tuple with the CurrentHead field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitRepositoryRequest) GetCurrentHeadOk() (*string, bool) {
	if o == nil || IsNil(o.CurrentHead) {
		return nil, false
	}
	return o.CurrentHead, true
}

// HasCurrentHead returns a boolean if a field has been set.
func (o *GitRepositoryRequest) HasCurrentHead() bool {
	if o != nil && !IsNil(o.CurrentHead) {
		return true
	}

	return false
}

// SetCurrentHead gets a reference to the given string and assigns it to the CurrentHead field.
func (o *GitRepositoryRequest) SetCurrentHead(v string) {
	o.CurrentHead = &v
}

// GetSecretsGroup returns the SecretsGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GitRepositoryRequest) GetSecretsGroup() BulkWritableCircuitRequestTenant {
	if o == nil || IsNil(o.SecretsGroup.Get()) {
		var ret BulkWritableCircuitRequestTenant
		return ret
	}
	return *o.SecretsGroup.Get()
}

// GetSecretsGroupOk returns a tuple with the SecretsGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GitRepositoryRequest) GetSecretsGroupOk() (*BulkWritableCircuitRequestTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.SecretsGroup.Get(), o.SecretsGroup.IsSet()
}

// HasSecretsGroup returns a boolean if a field has been set.
func (o *GitRepositoryRequest) HasSecretsGroup() bool {
	if o != nil && o.SecretsGroup.IsSet() {
		return true
	}

	return false
}

// SetSecretsGroup gets a reference to the given NullableBulkWritableCircuitRequestTenant and assigns it to the SecretsGroup field.
func (o *GitRepositoryRequest) SetSecretsGroup(v BulkWritableCircuitRequestTenant) {
	o.SecretsGroup.Set(&v)
}
// SetSecretsGroupNil sets the value for SecretsGroup to be an explicit nil
func (o *GitRepositoryRequest) SetSecretsGroupNil() {
	o.SecretsGroup.Set(nil)
}

// UnsetSecretsGroup ensures that no value is present for SecretsGroup, not even an explicit nil
func (o *GitRepositoryRequest) UnsetSecretsGroup() {
	o.SecretsGroup.Unset()
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *GitRepositoryRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitRepositoryRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *GitRepositoryRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *GitRepositoryRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *GitRepositoryRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue {
	if o == nil || IsNil(o.Relationships) {
		var ret map[string]BulkWritableCableRequestRelationshipsValue
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitRepositoryRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *GitRepositoryRequest) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]BulkWritableCableRequestRelationshipsValue and assigns it to the Relationships field.
func (o *GitRepositoryRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue) {
	o.Relationships = &v
}

func (o GitRepositoryRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GitRepositoryRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProvidedContents) {
		toSerialize["provided_contents"] = o.ProvidedContents
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	toSerialize["remote_url"] = o.RemoteUrl
	if !IsNil(o.Branch) {
		toSerialize["branch"] = o.Branch
	}
	if !IsNil(o.CurrentHead) {
		toSerialize["current_head"] = o.CurrentHead
	}
	if o.SecretsGroup.IsSet() {
		toSerialize["secrets_group"] = o.SecretsGroup.Get()
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

func (o *GitRepositoryRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"remote_url",
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

	varGitRepositoryRequest := _GitRepositoryRequest{}

	err = json.Unmarshal(data, &varGitRepositoryRequest)

	if err != nil {
		return err
	}

	*o = GitRepositoryRequest(varGitRepositoryRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "provided_contents")
		delete(additionalProperties, "name")
		delete(additionalProperties, "slug")
		delete(additionalProperties, "remote_url")
		delete(additionalProperties, "branch")
		delete(additionalProperties, "current_head")
		delete(additionalProperties, "secrets_group")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "relationships")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGitRepositoryRequest struct {
	value *GitRepositoryRequest
	isSet bool
}

func (v NullableGitRepositoryRequest) Get() *GitRepositoryRequest {
	return v.value
}

func (v *NullableGitRepositoryRequest) Set(val *GitRepositoryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGitRepositoryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGitRepositoryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGitRepositoryRequest(val *GitRepositoryRequest) *NullableGitRepositoryRequest {
	return &NullableGitRepositoryRequest{value: val, isSet: true}
}

func (v NullableGitRepositoryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGitRepositoryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

