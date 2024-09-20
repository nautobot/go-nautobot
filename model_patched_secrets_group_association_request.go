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

// checks if the PatchedSecretsGroupAssociationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedSecretsGroupAssociationRequest{}

// PatchedSecretsGroupAssociationRequest Serializer for `SecretsGroupAssociation` objects.
type PatchedSecretsGroupAssociationRequest struct {
	AccessType *AccessTypeEnum `json:"access_type,omitempty"`
	SecretType *SecretTypeEnum `json:"secret_type,omitempty"`
	SecretsGroup *BulkWritableCableRequestStatus `json:"secrets_group,omitempty"`
	Secret *BulkWritableCableRequestStatus `json:"secret,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedSecretsGroupAssociationRequest PatchedSecretsGroupAssociationRequest

// NewPatchedSecretsGroupAssociationRequest instantiates a new PatchedSecretsGroupAssociationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedSecretsGroupAssociationRequest() *PatchedSecretsGroupAssociationRequest {
	this := PatchedSecretsGroupAssociationRequest{}
	return &this
}

// NewPatchedSecretsGroupAssociationRequestWithDefaults instantiates a new PatchedSecretsGroupAssociationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedSecretsGroupAssociationRequestWithDefaults() *PatchedSecretsGroupAssociationRequest {
	this := PatchedSecretsGroupAssociationRequest{}
	return &this
}

// GetAccessType returns the AccessType field value if set, zero value otherwise.
func (o *PatchedSecretsGroupAssociationRequest) GetAccessType() AccessTypeEnum {
	if o == nil || IsNil(o.AccessType) {
		var ret AccessTypeEnum
		return ret
	}
	return *o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSecretsGroupAssociationRequest) GetAccessTypeOk() (*AccessTypeEnum, bool) {
	if o == nil || IsNil(o.AccessType) {
		return nil, false
	}
	return o.AccessType, true
}

// HasAccessType returns a boolean if a field has been set.
func (o *PatchedSecretsGroupAssociationRequest) HasAccessType() bool {
	if o != nil && !IsNil(o.AccessType) {
		return true
	}

	return false
}

// SetAccessType gets a reference to the given AccessTypeEnum and assigns it to the AccessType field.
func (o *PatchedSecretsGroupAssociationRequest) SetAccessType(v AccessTypeEnum) {
	o.AccessType = &v
}

// GetSecretType returns the SecretType field value if set, zero value otherwise.
func (o *PatchedSecretsGroupAssociationRequest) GetSecretType() SecretTypeEnum {
	if o == nil || IsNil(o.SecretType) {
		var ret SecretTypeEnum
		return ret
	}
	return *o.SecretType
}

// GetSecretTypeOk returns a tuple with the SecretType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSecretsGroupAssociationRequest) GetSecretTypeOk() (*SecretTypeEnum, bool) {
	if o == nil || IsNil(o.SecretType) {
		return nil, false
	}
	return o.SecretType, true
}

// HasSecretType returns a boolean if a field has been set.
func (o *PatchedSecretsGroupAssociationRequest) HasSecretType() bool {
	if o != nil && !IsNil(o.SecretType) {
		return true
	}

	return false
}

// SetSecretType gets a reference to the given SecretTypeEnum and assigns it to the SecretType field.
func (o *PatchedSecretsGroupAssociationRequest) SetSecretType(v SecretTypeEnum) {
	o.SecretType = &v
}

// GetSecretsGroup returns the SecretsGroup field value if set, zero value otherwise.
func (o *PatchedSecretsGroupAssociationRequest) GetSecretsGroup() BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.SecretsGroup) {
		var ret BulkWritableCableRequestStatus
		return ret
	}
	return *o.SecretsGroup
}

// GetSecretsGroupOk returns a tuple with the SecretsGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSecretsGroupAssociationRequest) GetSecretsGroupOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.SecretsGroup) {
		return nil, false
	}
	return o.SecretsGroup, true
}

// HasSecretsGroup returns a boolean if a field has been set.
func (o *PatchedSecretsGroupAssociationRequest) HasSecretsGroup() bool {
	if o != nil && !IsNil(o.SecretsGroup) {
		return true
	}

	return false
}

// SetSecretsGroup gets a reference to the given BulkWritableCableRequestStatus and assigns it to the SecretsGroup field.
func (o *PatchedSecretsGroupAssociationRequest) SetSecretsGroup(v BulkWritableCableRequestStatus) {
	o.SecretsGroup = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *PatchedSecretsGroupAssociationRequest) GetSecret() BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Secret) {
		var ret BulkWritableCableRequestStatus
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedSecretsGroupAssociationRequest) GetSecretOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Secret) {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *PatchedSecretsGroupAssociationRequest) HasSecret() bool {
	if o != nil && !IsNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given BulkWritableCableRequestStatus and assigns it to the Secret field.
func (o *PatchedSecretsGroupAssociationRequest) SetSecret(v BulkWritableCableRequestStatus) {
	o.Secret = &v
}

func (o PatchedSecretsGroupAssociationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedSecretsGroupAssociationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessType) {
		toSerialize["access_type"] = o.AccessType
	}
	if !IsNil(o.SecretType) {
		toSerialize["secret_type"] = o.SecretType
	}
	if !IsNil(o.SecretsGroup) {
		toSerialize["secrets_group"] = o.SecretsGroup
	}
	if !IsNil(o.Secret) {
		toSerialize["secret"] = o.Secret
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedSecretsGroupAssociationRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedSecretsGroupAssociationRequest := _PatchedSecretsGroupAssociationRequest{}

	err = json.Unmarshal(data, &varPatchedSecretsGroupAssociationRequest)

	if err != nil {
		return err
	}

	*o = PatchedSecretsGroupAssociationRequest(varPatchedSecretsGroupAssociationRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "access_type")
		delete(additionalProperties, "secret_type")
		delete(additionalProperties, "secrets_group")
		delete(additionalProperties, "secret")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedSecretsGroupAssociationRequest struct {
	value *PatchedSecretsGroupAssociationRequest
	isSet bool
}

func (v NullablePatchedSecretsGroupAssociationRequest) Get() *PatchedSecretsGroupAssociationRequest {
	return v.value
}

func (v *NullablePatchedSecretsGroupAssociationRequest) Set(val *PatchedSecretsGroupAssociationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedSecretsGroupAssociationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedSecretsGroupAssociationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedSecretsGroupAssociationRequest(val *PatchedSecretsGroupAssociationRequest) *NullablePatchedSecretsGroupAssociationRequest {
	return &NullablePatchedSecretsGroupAssociationRequest{value: val, isSet: true}
}

func (v NullablePatchedSecretsGroupAssociationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedSecretsGroupAssociationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


