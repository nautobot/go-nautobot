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

// checks if the PatchedBulkWritableSecretsGroupAssociationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedBulkWritableSecretsGroupAssociationRequest{}

// PatchedBulkWritableSecretsGroupAssociationRequest Serializer for `SecretsGroupAssociation` objects.
type PatchedBulkWritableSecretsGroupAssociationRequest struct {
	Id string `json:"id"`
	AccessType *AccessTypeEnum `json:"access_type,omitempty"`
	SecretType *SecretTypeEnum `json:"secret_type,omitempty"`
	SecretsGroup *BulkWritableCableRequestStatus `json:"secrets_group,omitempty"`
	Secret *BulkWritableCableRequestStatus `json:"secret,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedBulkWritableSecretsGroupAssociationRequest PatchedBulkWritableSecretsGroupAssociationRequest

// NewPatchedBulkWritableSecretsGroupAssociationRequest instantiates a new PatchedBulkWritableSecretsGroupAssociationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedBulkWritableSecretsGroupAssociationRequest(id string) *PatchedBulkWritableSecretsGroupAssociationRequest {
	this := PatchedBulkWritableSecretsGroupAssociationRequest{}
	this.Id = id
	return &this
}

// NewPatchedBulkWritableSecretsGroupAssociationRequestWithDefaults instantiates a new PatchedBulkWritableSecretsGroupAssociationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedBulkWritableSecretsGroupAssociationRequestWithDefaults() *PatchedBulkWritableSecretsGroupAssociationRequest {
	this := PatchedBulkWritableSecretsGroupAssociationRequest{}
	return &this
}

// GetId returns the Id field value
func (o *PatchedBulkWritableSecretsGroupAssociationRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableSecretsGroupAssociationRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PatchedBulkWritableSecretsGroupAssociationRequest) SetId(v string) {
	o.Id = v
}

// GetAccessType returns the AccessType field value if set, zero value otherwise.
func (o *PatchedBulkWritableSecretsGroupAssociationRequest) GetAccessType() AccessTypeEnum {
	if o == nil || IsNil(o.AccessType) {
		var ret AccessTypeEnum
		return ret
	}
	return *o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableSecretsGroupAssociationRequest) GetAccessTypeOk() (*AccessTypeEnum, bool) {
	if o == nil || IsNil(o.AccessType) {
		return nil, false
	}
	return o.AccessType, true
}

// HasAccessType returns a boolean if a field has been set.
func (o *PatchedBulkWritableSecretsGroupAssociationRequest) HasAccessType() bool {
	if o != nil && !IsNil(o.AccessType) {
		return true
	}

	return false
}

// SetAccessType gets a reference to the given AccessTypeEnum and assigns it to the AccessType field.
func (o *PatchedBulkWritableSecretsGroupAssociationRequest) SetAccessType(v AccessTypeEnum) {
	o.AccessType = &v
}

// GetSecretType returns the SecretType field value if set, zero value otherwise.
func (o *PatchedBulkWritableSecretsGroupAssociationRequest) GetSecretType() SecretTypeEnum {
	if o == nil || IsNil(o.SecretType) {
		var ret SecretTypeEnum
		return ret
	}
	return *o.SecretType
}

// GetSecretTypeOk returns a tuple with the SecretType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableSecretsGroupAssociationRequest) GetSecretTypeOk() (*SecretTypeEnum, bool) {
	if o == nil || IsNil(o.SecretType) {
		return nil, false
	}
	return o.SecretType, true
}

// HasSecretType returns a boolean if a field has been set.
func (o *PatchedBulkWritableSecretsGroupAssociationRequest) HasSecretType() bool {
	if o != nil && !IsNil(o.SecretType) {
		return true
	}

	return false
}

// SetSecretType gets a reference to the given SecretTypeEnum and assigns it to the SecretType field.
func (o *PatchedBulkWritableSecretsGroupAssociationRequest) SetSecretType(v SecretTypeEnum) {
	o.SecretType = &v
}

// GetSecretsGroup returns the SecretsGroup field value if set, zero value otherwise.
func (o *PatchedBulkWritableSecretsGroupAssociationRequest) GetSecretsGroup() BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.SecretsGroup) {
		var ret BulkWritableCableRequestStatus
		return ret
	}
	return *o.SecretsGroup
}

// GetSecretsGroupOk returns a tuple with the SecretsGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableSecretsGroupAssociationRequest) GetSecretsGroupOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.SecretsGroup) {
		return nil, false
	}
	return o.SecretsGroup, true
}

// HasSecretsGroup returns a boolean if a field has been set.
func (o *PatchedBulkWritableSecretsGroupAssociationRequest) HasSecretsGroup() bool {
	if o != nil && !IsNil(o.SecretsGroup) {
		return true
	}

	return false
}

// SetSecretsGroup gets a reference to the given BulkWritableCableRequestStatus and assigns it to the SecretsGroup field.
func (o *PatchedBulkWritableSecretsGroupAssociationRequest) SetSecretsGroup(v BulkWritableCableRequestStatus) {
	o.SecretsGroup = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *PatchedBulkWritableSecretsGroupAssociationRequest) GetSecret() BulkWritableCableRequestStatus {
	if o == nil || IsNil(o.Secret) {
		var ret BulkWritableCableRequestStatus
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedBulkWritableSecretsGroupAssociationRequest) GetSecretOk() (*BulkWritableCableRequestStatus, bool) {
	if o == nil || IsNil(o.Secret) {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *PatchedBulkWritableSecretsGroupAssociationRequest) HasSecret() bool {
	if o != nil && !IsNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given BulkWritableCableRequestStatus and assigns it to the Secret field.
func (o *PatchedBulkWritableSecretsGroupAssociationRequest) SetSecret(v BulkWritableCableRequestStatus) {
	o.Secret = &v
}

func (o PatchedBulkWritableSecretsGroupAssociationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedBulkWritableSecretsGroupAssociationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
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

func (o *PatchedBulkWritableSecretsGroupAssociationRequest) UnmarshalJSON(data []byte) (err error) {
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

	varPatchedBulkWritableSecretsGroupAssociationRequest := _PatchedBulkWritableSecretsGroupAssociationRequest{}

	err = json.Unmarshal(data, &varPatchedBulkWritableSecretsGroupAssociationRequest)

	if err != nil {
		return err
	}

	*o = PatchedBulkWritableSecretsGroupAssociationRequest(varPatchedBulkWritableSecretsGroupAssociationRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "access_type")
		delete(additionalProperties, "secret_type")
		delete(additionalProperties, "secrets_group")
		delete(additionalProperties, "secret")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedBulkWritableSecretsGroupAssociationRequest struct {
	value *PatchedBulkWritableSecretsGroupAssociationRequest
	isSet bool
}

func (v NullablePatchedBulkWritableSecretsGroupAssociationRequest) Get() *PatchedBulkWritableSecretsGroupAssociationRequest {
	return v.value
}

func (v *NullablePatchedBulkWritableSecretsGroupAssociationRequest) Set(val *PatchedBulkWritableSecretsGroupAssociationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedBulkWritableSecretsGroupAssociationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedBulkWritableSecretsGroupAssociationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedBulkWritableSecretsGroupAssociationRequest(val *PatchedBulkWritableSecretsGroupAssociationRequest) *NullablePatchedBulkWritableSecretsGroupAssociationRequest {
	return &NullablePatchedBulkWritableSecretsGroupAssociationRequest{value: val, isSet: true}
}

func (v NullablePatchedBulkWritableSecretsGroupAssociationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedBulkWritableSecretsGroupAssociationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


