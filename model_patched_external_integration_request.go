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

// checks if the PatchedExternalIntegrationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedExternalIntegrationRequest{}

// PatchedExternalIntegrationRequest Base class to use for serializers based on OrganizationalModel or PrimaryModel.  Can also be used for models derived from BaseModel, so long as they support custom fields, notes, and relationships.
type PatchedExternalIntegrationRequest struct {
	Name *string `json:"name,omitempty"`
	RemoteUrl *string `json:"remote_url,omitempty"`
	// Verify SSL certificates when connecting to the remote system
	VerifySsl *bool `json:"verify_ssl,omitempty"`
	// Number of seconds to wait for a response
	Timeout *int32 `json:"timeout,omitempty"`
	// Optional user-defined JSON data for this integration
	ExtraConfig interface{} `json:"extra_config,omitempty"`
	HttpMethod *BulkWritableExternalIntegrationRequestHttpMethod `json:"http_method,omitempty"`
	// Headers for the HTTP request
	Headers interface{} `json:"headers,omitempty"`
	CaFilePath *string `json:"ca_file_path,omitempty"`
	SecretsGroup NullableBulkWritableExternalIntegrationRequestSecretsGroup `json:"secrets_group,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Relationships *map[string]BulkWritableCableRequestRelationshipsValue `json:"relationships,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedExternalIntegrationRequest PatchedExternalIntegrationRequest

// NewPatchedExternalIntegrationRequest instantiates a new PatchedExternalIntegrationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedExternalIntegrationRequest() *PatchedExternalIntegrationRequest {
	this := PatchedExternalIntegrationRequest{}
	return &this
}

// NewPatchedExternalIntegrationRequestWithDefaults instantiates a new PatchedExternalIntegrationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedExternalIntegrationRequestWithDefaults() *PatchedExternalIntegrationRequest {
	this := PatchedExternalIntegrationRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedExternalIntegrationRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedExternalIntegrationRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedExternalIntegrationRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedExternalIntegrationRequest) SetName(v string) {
	o.Name = &v
}

// GetRemoteUrl returns the RemoteUrl field value if set, zero value otherwise.
func (o *PatchedExternalIntegrationRequest) GetRemoteUrl() string {
	if o == nil || IsNil(o.RemoteUrl) {
		var ret string
		return ret
	}
	return *o.RemoteUrl
}

// GetRemoteUrlOk returns a tuple with the RemoteUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedExternalIntegrationRequest) GetRemoteUrlOk() (*string, bool) {
	if o == nil || IsNil(o.RemoteUrl) {
		return nil, false
	}
	return o.RemoteUrl, true
}

// HasRemoteUrl returns a boolean if a field has been set.
func (o *PatchedExternalIntegrationRequest) HasRemoteUrl() bool {
	if o != nil && !IsNil(o.RemoteUrl) {
		return true
	}

	return false
}

// SetRemoteUrl gets a reference to the given string and assigns it to the RemoteUrl field.
func (o *PatchedExternalIntegrationRequest) SetRemoteUrl(v string) {
	o.RemoteUrl = &v
}

// GetVerifySsl returns the VerifySsl field value if set, zero value otherwise.
func (o *PatchedExternalIntegrationRequest) GetVerifySsl() bool {
	if o == nil || IsNil(o.VerifySsl) {
		var ret bool
		return ret
	}
	return *o.VerifySsl
}

// GetVerifySslOk returns a tuple with the VerifySsl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedExternalIntegrationRequest) GetVerifySslOk() (*bool, bool) {
	if o == nil || IsNil(o.VerifySsl) {
		return nil, false
	}
	return o.VerifySsl, true
}

// HasVerifySsl returns a boolean if a field has been set.
func (o *PatchedExternalIntegrationRequest) HasVerifySsl() bool {
	if o != nil && !IsNil(o.VerifySsl) {
		return true
	}

	return false
}

// SetVerifySsl gets a reference to the given bool and assigns it to the VerifySsl field.
func (o *PatchedExternalIntegrationRequest) SetVerifySsl(v bool) {
	o.VerifySsl = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *PatchedExternalIntegrationRequest) GetTimeout() int32 {
	if o == nil || IsNil(o.Timeout) {
		var ret int32
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedExternalIntegrationRequest) GetTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.Timeout) {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *PatchedExternalIntegrationRequest) HasTimeout() bool {
	if o != nil && !IsNil(o.Timeout) {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int32 and assigns it to the Timeout field.
func (o *PatchedExternalIntegrationRequest) SetTimeout(v int32) {
	o.Timeout = &v
}

// GetExtraConfig returns the ExtraConfig field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedExternalIntegrationRequest) GetExtraConfig() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ExtraConfig
}

// GetExtraConfigOk returns a tuple with the ExtraConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedExternalIntegrationRequest) GetExtraConfigOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ExtraConfig) {
		return nil, false
	}
	return &o.ExtraConfig, true
}

// HasExtraConfig returns a boolean if a field has been set.
func (o *PatchedExternalIntegrationRequest) HasExtraConfig() bool {
	if o != nil && !IsNil(o.ExtraConfig) {
		return true
	}

	return false
}

// SetExtraConfig gets a reference to the given interface{} and assigns it to the ExtraConfig field.
func (o *PatchedExternalIntegrationRequest) SetExtraConfig(v interface{}) {
	o.ExtraConfig = v
}

// GetHttpMethod returns the HttpMethod field value if set, zero value otherwise.
func (o *PatchedExternalIntegrationRequest) GetHttpMethod() BulkWritableExternalIntegrationRequestHttpMethod {
	if o == nil || IsNil(o.HttpMethod) {
		var ret BulkWritableExternalIntegrationRequestHttpMethod
		return ret
	}
	return *o.HttpMethod
}

// GetHttpMethodOk returns a tuple with the HttpMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedExternalIntegrationRequest) GetHttpMethodOk() (*BulkWritableExternalIntegrationRequestHttpMethod, bool) {
	if o == nil || IsNil(o.HttpMethod) {
		return nil, false
	}
	return o.HttpMethod, true
}

// HasHttpMethod returns a boolean if a field has been set.
func (o *PatchedExternalIntegrationRequest) HasHttpMethod() bool {
	if o != nil && !IsNil(o.HttpMethod) {
		return true
	}

	return false
}

// SetHttpMethod gets a reference to the given BulkWritableExternalIntegrationRequestHttpMethod and assigns it to the HttpMethod field.
func (o *PatchedExternalIntegrationRequest) SetHttpMethod(v BulkWritableExternalIntegrationRequestHttpMethod) {
	o.HttpMethod = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedExternalIntegrationRequest) GetHeaders() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedExternalIntegrationRequest) GetHeadersOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Headers) {
		return nil, false
	}
	return &o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *PatchedExternalIntegrationRequest) HasHeaders() bool {
	if o != nil && !IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given interface{} and assigns it to the Headers field.
func (o *PatchedExternalIntegrationRequest) SetHeaders(v interface{}) {
	o.Headers = v
}

// GetCaFilePath returns the CaFilePath field value if set, zero value otherwise.
func (o *PatchedExternalIntegrationRequest) GetCaFilePath() string {
	if o == nil || IsNil(o.CaFilePath) {
		var ret string
		return ret
	}
	return *o.CaFilePath
}

// GetCaFilePathOk returns a tuple with the CaFilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedExternalIntegrationRequest) GetCaFilePathOk() (*string, bool) {
	if o == nil || IsNil(o.CaFilePath) {
		return nil, false
	}
	return o.CaFilePath, true
}

// HasCaFilePath returns a boolean if a field has been set.
func (o *PatchedExternalIntegrationRequest) HasCaFilePath() bool {
	if o != nil && !IsNil(o.CaFilePath) {
		return true
	}

	return false
}

// SetCaFilePath gets a reference to the given string and assigns it to the CaFilePath field.
func (o *PatchedExternalIntegrationRequest) SetCaFilePath(v string) {
	o.CaFilePath = &v
}

// GetSecretsGroup returns the SecretsGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedExternalIntegrationRequest) GetSecretsGroup() BulkWritableExternalIntegrationRequestSecretsGroup {
	if o == nil || IsNil(o.SecretsGroup.Get()) {
		var ret BulkWritableExternalIntegrationRequestSecretsGroup
		return ret
	}
	return *o.SecretsGroup.Get()
}

// GetSecretsGroupOk returns a tuple with the SecretsGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedExternalIntegrationRequest) GetSecretsGroupOk() (*BulkWritableExternalIntegrationRequestSecretsGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.SecretsGroup.Get(), o.SecretsGroup.IsSet()
}

// HasSecretsGroup returns a boolean if a field has been set.
func (o *PatchedExternalIntegrationRequest) HasSecretsGroup() bool {
	if o != nil && o.SecretsGroup.IsSet() {
		return true
	}

	return false
}

// SetSecretsGroup gets a reference to the given NullableBulkWritableExternalIntegrationRequestSecretsGroup and assigns it to the SecretsGroup field.
func (o *PatchedExternalIntegrationRequest) SetSecretsGroup(v BulkWritableExternalIntegrationRequestSecretsGroup) {
	o.SecretsGroup.Set(&v)
}
// SetSecretsGroupNil sets the value for SecretsGroup to be an explicit nil
func (o *PatchedExternalIntegrationRequest) SetSecretsGroupNil() {
	o.SecretsGroup.Set(nil)
}

// UnsetSecretsGroup ensures that no value is present for SecretsGroup, not even an explicit nil
func (o *PatchedExternalIntegrationRequest) UnsetSecretsGroup() {
	o.SecretsGroup.Unset()
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedExternalIntegrationRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedExternalIntegrationRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedExternalIntegrationRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedExternalIntegrationRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *PatchedExternalIntegrationRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue {
	if o == nil || IsNil(o.Relationships) {
		var ret map[string]BulkWritableCableRequestRelationshipsValue
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedExternalIntegrationRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *PatchedExternalIntegrationRequest) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given map[string]BulkWritableCableRequestRelationshipsValue and assigns it to the Relationships field.
func (o *PatchedExternalIntegrationRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue) {
	o.Relationships = &v
}

func (o PatchedExternalIntegrationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedExternalIntegrationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.RemoteUrl) {
		toSerialize["remote_url"] = o.RemoteUrl
	}
	if !IsNil(o.VerifySsl) {
		toSerialize["verify_ssl"] = o.VerifySsl
	}
	if !IsNil(o.Timeout) {
		toSerialize["timeout"] = o.Timeout
	}
	if o.ExtraConfig != nil {
		toSerialize["extra_config"] = o.ExtraConfig
	}
	if !IsNil(o.HttpMethod) {
		toSerialize["http_method"] = o.HttpMethod
	}
	if o.Headers != nil {
		toSerialize["headers"] = o.Headers
	}
	if !IsNil(o.CaFilePath) {
		toSerialize["ca_file_path"] = o.CaFilePath
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

func (o *PatchedExternalIntegrationRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedExternalIntegrationRequest := _PatchedExternalIntegrationRequest{}

	err = json.Unmarshal(data, &varPatchedExternalIntegrationRequest)

	if err != nil {
		return err
	}

	*o = PatchedExternalIntegrationRequest(varPatchedExternalIntegrationRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "remote_url")
		delete(additionalProperties, "verify_ssl")
		delete(additionalProperties, "timeout")
		delete(additionalProperties, "extra_config")
		delete(additionalProperties, "http_method")
		delete(additionalProperties, "headers")
		delete(additionalProperties, "ca_file_path")
		delete(additionalProperties, "secrets_group")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "relationships")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedExternalIntegrationRequest struct {
	value *PatchedExternalIntegrationRequest
	isSet bool
}

func (v NullablePatchedExternalIntegrationRequest) Get() *PatchedExternalIntegrationRequest {
	return v.value
}

func (v *NullablePatchedExternalIntegrationRequest) Set(val *PatchedExternalIntegrationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedExternalIntegrationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedExternalIntegrationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedExternalIntegrationRequest(val *PatchedExternalIntegrationRequest) *NullablePatchedExternalIntegrationRequest {
	return &NullablePatchedExternalIntegrationRequest{value: val, isSet: true}
}

func (v NullablePatchedExternalIntegrationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedExternalIntegrationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


