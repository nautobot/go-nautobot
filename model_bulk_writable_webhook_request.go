/*
API Documentation

Source of truth and network automation platform

API version: 2.3.1 (2.3)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nautobot

import (
	"encoding/json"
	"fmt"
)

// checks if the BulkWritableWebhookRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkWritableWebhookRequest{}

// BulkWritableWebhookRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type BulkWritableWebhookRequest struct {
	Id string `json:"id"`
	ContentTypes []string `json:"content_types"`
	Name string `json:"name"`
	// Call this webhook when a matching object is created.
	TypeCreate *bool `json:"type_create,omitempty"`
	// Call this webhook when a matching object is updated.
	TypeUpdate *bool `json:"type_update,omitempty"`
	// Call this webhook when a matching object is deleted.
	TypeDelete *bool `json:"type_delete,omitempty"`
	// A POST will be sent to this URL when the webhook is called.
	PayloadUrl string `json:"payload_url"`
	Enabled *bool `json:"enabled,omitempty"`
	HttpMethod *HttpMethodEnum `json:"http_method,omitempty"`
	// The complete list of official content types is available <a href=\"https://www.iana.org/assignments/media-types/media-types.xhtml\">here</a>.
	HttpContentType *string `json:"http_content_type,omitempty"`
	// User-supplied HTTP headers to be sent with the request in addition to the HTTP content type. Headers should be defined in the format <code>Name: Value</code>. Jinja2 template processing is supported with the same context as the request body (below).
	AdditionalHeaders *string `json:"additional_headers,omitempty"`
	// Jinja2 template for a custom request body. If blank, a JSON object representing the change will be included. Available context data includes: <code>event</code>, <code>model</code>, <code>timestamp</code>, <code>username</code>, <code>request_id</code>, and <code>data</code>.
	BodyTemplate *string `json:"body_template,omitempty"`
	// When provided, the request will include a 'X-Hook-Signature' header containing a HMAC hex digest of the payload body using the secret as the key. The secret is not transmitted in the request.
	Secret *string `json:"secret,omitempty"`
	// Enable SSL certificate verification. Disable with caution!
	SslVerification *bool `json:"ssl_verification,omitempty"`
	// The specific CA certificate file to use for SSL verification. Leave blank to use the system defaults.
	CaFilePath *string `json:"ca_file_path,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkWritableWebhookRequest BulkWritableWebhookRequest

// NewBulkWritableWebhookRequest instantiates a new BulkWritableWebhookRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkWritableWebhookRequest(id string, contentTypes []string, name string, payloadUrl string) *BulkWritableWebhookRequest {
	this := BulkWritableWebhookRequest{}
	this.Id = id
	this.ContentTypes = contentTypes
	this.Name = name
	this.PayloadUrl = payloadUrl
	return &this
}

// NewBulkWritableWebhookRequestWithDefaults instantiates a new BulkWritableWebhookRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkWritableWebhookRequestWithDefaults() *BulkWritableWebhookRequest {
	this := BulkWritableWebhookRequest{}
	return &this
}

// GetId returns the Id field value
func (o *BulkWritableWebhookRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BulkWritableWebhookRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BulkWritableWebhookRequest) SetId(v string) {
	o.Id = v
}

// GetContentTypes returns the ContentTypes field value
func (o *BulkWritableWebhookRequest) GetContentTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ContentTypes
}

// GetContentTypesOk returns a tuple with the ContentTypes field value
// and a boolean to check if the value has been set.
func (o *BulkWritableWebhookRequest) GetContentTypesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContentTypes, true
}

// SetContentTypes sets field value
func (o *BulkWritableWebhookRequest) SetContentTypes(v []string) {
	o.ContentTypes = v
}

// GetName returns the Name field value
func (o *BulkWritableWebhookRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BulkWritableWebhookRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BulkWritableWebhookRequest) SetName(v string) {
	o.Name = v
}

// GetTypeCreate returns the TypeCreate field value if set, zero value otherwise.
func (o *BulkWritableWebhookRequest) GetTypeCreate() bool {
	if o == nil || IsNil(o.TypeCreate) {
		var ret bool
		return ret
	}
	return *o.TypeCreate
}

// GetTypeCreateOk returns a tuple with the TypeCreate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableWebhookRequest) GetTypeCreateOk() (*bool, bool) {
	if o == nil || IsNil(o.TypeCreate) {
		return nil, false
	}
	return o.TypeCreate, true
}

// HasTypeCreate returns a boolean if a field has been set.
func (o *BulkWritableWebhookRequest) HasTypeCreate() bool {
	if o != nil && !IsNil(o.TypeCreate) {
		return true
	}

	return false
}

// SetTypeCreate gets a reference to the given bool and assigns it to the TypeCreate field.
func (o *BulkWritableWebhookRequest) SetTypeCreate(v bool) {
	o.TypeCreate = &v
}

// GetTypeUpdate returns the TypeUpdate field value if set, zero value otherwise.
func (o *BulkWritableWebhookRequest) GetTypeUpdate() bool {
	if o == nil || IsNil(o.TypeUpdate) {
		var ret bool
		return ret
	}
	return *o.TypeUpdate
}

// GetTypeUpdateOk returns a tuple with the TypeUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableWebhookRequest) GetTypeUpdateOk() (*bool, bool) {
	if o == nil || IsNil(o.TypeUpdate) {
		return nil, false
	}
	return o.TypeUpdate, true
}

// HasTypeUpdate returns a boolean if a field has been set.
func (o *BulkWritableWebhookRequest) HasTypeUpdate() bool {
	if o != nil && !IsNil(o.TypeUpdate) {
		return true
	}

	return false
}

// SetTypeUpdate gets a reference to the given bool and assigns it to the TypeUpdate field.
func (o *BulkWritableWebhookRequest) SetTypeUpdate(v bool) {
	o.TypeUpdate = &v
}

// GetTypeDelete returns the TypeDelete field value if set, zero value otherwise.
func (o *BulkWritableWebhookRequest) GetTypeDelete() bool {
	if o == nil || IsNil(o.TypeDelete) {
		var ret bool
		return ret
	}
	return *o.TypeDelete
}

// GetTypeDeleteOk returns a tuple with the TypeDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableWebhookRequest) GetTypeDeleteOk() (*bool, bool) {
	if o == nil || IsNil(o.TypeDelete) {
		return nil, false
	}
	return o.TypeDelete, true
}

// HasTypeDelete returns a boolean if a field has been set.
func (o *BulkWritableWebhookRequest) HasTypeDelete() bool {
	if o != nil && !IsNil(o.TypeDelete) {
		return true
	}

	return false
}

// SetTypeDelete gets a reference to the given bool and assigns it to the TypeDelete field.
func (o *BulkWritableWebhookRequest) SetTypeDelete(v bool) {
	o.TypeDelete = &v
}

// GetPayloadUrl returns the PayloadUrl field value
func (o *BulkWritableWebhookRequest) GetPayloadUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PayloadUrl
}

// GetPayloadUrlOk returns a tuple with the PayloadUrl field value
// and a boolean to check if the value has been set.
func (o *BulkWritableWebhookRequest) GetPayloadUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PayloadUrl, true
}

// SetPayloadUrl sets field value
func (o *BulkWritableWebhookRequest) SetPayloadUrl(v string) {
	o.PayloadUrl = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *BulkWritableWebhookRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableWebhookRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *BulkWritableWebhookRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *BulkWritableWebhookRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetHttpMethod returns the HttpMethod field value if set, zero value otherwise.
func (o *BulkWritableWebhookRequest) GetHttpMethod() HttpMethodEnum {
	if o == nil || IsNil(o.HttpMethod) {
		var ret HttpMethodEnum
		return ret
	}
	return *o.HttpMethod
}

// GetHttpMethodOk returns a tuple with the HttpMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableWebhookRequest) GetHttpMethodOk() (*HttpMethodEnum, bool) {
	if o == nil || IsNil(o.HttpMethod) {
		return nil, false
	}
	return o.HttpMethod, true
}

// HasHttpMethod returns a boolean if a field has been set.
func (o *BulkWritableWebhookRequest) HasHttpMethod() bool {
	if o != nil && !IsNil(o.HttpMethod) {
		return true
	}

	return false
}

// SetHttpMethod gets a reference to the given HttpMethodEnum and assigns it to the HttpMethod field.
func (o *BulkWritableWebhookRequest) SetHttpMethod(v HttpMethodEnum) {
	o.HttpMethod = &v
}

// GetHttpContentType returns the HttpContentType field value if set, zero value otherwise.
func (o *BulkWritableWebhookRequest) GetHttpContentType() string {
	if o == nil || IsNil(o.HttpContentType) {
		var ret string
		return ret
	}
	return *o.HttpContentType
}

// GetHttpContentTypeOk returns a tuple with the HttpContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableWebhookRequest) GetHttpContentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.HttpContentType) {
		return nil, false
	}
	return o.HttpContentType, true
}

// HasHttpContentType returns a boolean if a field has been set.
func (o *BulkWritableWebhookRequest) HasHttpContentType() bool {
	if o != nil && !IsNil(o.HttpContentType) {
		return true
	}

	return false
}

// SetHttpContentType gets a reference to the given string and assigns it to the HttpContentType field.
func (o *BulkWritableWebhookRequest) SetHttpContentType(v string) {
	o.HttpContentType = &v
}

// GetAdditionalHeaders returns the AdditionalHeaders field value if set, zero value otherwise.
func (o *BulkWritableWebhookRequest) GetAdditionalHeaders() string {
	if o == nil || IsNil(o.AdditionalHeaders) {
		var ret string
		return ret
	}
	return *o.AdditionalHeaders
}

// GetAdditionalHeadersOk returns a tuple with the AdditionalHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableWebhookRequest) GetAdditionalHeadersOk() (*string, bool) {
	if o == nil || IsNil(o.AdditionalHeaders) {
		return nil, false
	}
	return o.AdditionalHeaders, true
}

// HasAdditionalHeaders returns a boolean if a field has been set.
func (o *BulkWritableWebhookRequest) HasAdditionalHeaders() bool {
	if o != nil && !IsNil(o.AdditionalHeaders) {
		return true
	}

	return false
}

// SetAdditionalHeaders gets a reference to the given string and assigns it to the AdditionalHeaders field.
func (o *BulkWritableWebhookRequest) SetAdditionalHeaders(v string) {
	o.AdditionalHeaders = &v
}

// GetBodyTemplate returns the BodyTemplate field value if set, zero value otherwise.
func (o *BulkWritableWebhookRequest) GetBodyTemplate() string {
	if o == nil || IsNil(o.BodyTemplate) {
		var ret string
		return ret
	}
	return *o.BodyTemplate
}

// GetBodyTemplateOk returns a tuple with the BodyTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableWebhookRequest) GetBodyTemplateOk() (*string, bool) {
	if o == nil || IsNil(o.BodyTemplate) {
		return nil, false
	}
	return o.BodyTemplate, true
}

// HasBodyTemplate returns a boolean if a field has been set.
func (o *BulkWritableWebhookRequest) HasBodyTemplate() bool {
	if o != nil && !IsNil(o.BodyTemplate) {
		return true
	}

	return false
}

// SetBodyTemplate gets a reference to the given string and assigns it to the BodyTemplate field.
func (o *BulkWritableWebhookRequest) SetBodyTemplate(v string) {
	o.BodyTemplate = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *BulkWritableWebhookRequest) GetSecret() string {
	if o == nil || IsNil(o.Secret) {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableWebhookRequest) GetSecretOk() (*string, bool) {
	if o == nil || IsNil(o.Secret) {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *BulkWritableWebhookRequest) HasSecret() bool {
	if o != nil && !IsNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *BulkWritableWebhookRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetSslVerification returns the SslVerification field value if set, zero value otherwise.
func (o *BulkWritableWebhookRequest) GetSslVerification() bool {
	if o == nil || IsNil(o.SslVerification) {
		var ret bool
		return ret
	}
	return *o.SslVerification
}

// GetSslVerificationOk returns a tuple with the SslVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableWebhookRequest) GetSslVerificationOk() (*bool, bool) {
	if o == nil || IsNil(o.SslVerification) {
		return nil, false
	}
	return o.SslVerification, true
}

// HasSslVerification returns a boolean if a field has been set.
func (o *BulkWritableWebhookRequest) HasSslVerification() bool {
	if o != nil && !IsNil(o.SslVerification) {
		return true
	}

	return false
}

// SetSslVerification gets a reference to the given bool and assigns it to the SslVerification field.
func (o *BulkWritableWebhookRequest) SetSslVerification(v bool) {
	o.SslVerification = &v
}

// GetCaFilePath returns the CaFilePath field value if set, zero value otherwise.
func (o *BulkWritableWebhookRequest) GetCaFilePath() string {
	if o == nil || IsNil(o.CaFilePath) {
		var ret string
		return ret
	}
	return *o.CaFilePath
}

// GetCaFilePathOk returns a tuple with the CaFilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkWritableWebhookRequest) GetCaFilePathOk() (*string, bool) {
	if o == nil || IsNil(o.CaFilePath) {
		return nil, false
	}
	return o.CaFilePath, true
}

// HasCaFilePath returns a boolean if a field has been set.
func (o *BulkWritableWebhookRequest) HasCaFilePath() bool {
	if o != nil && !IsNil(o.CaFilePath) {
		return true
	}

	return false
}

// SetCaFilePath gets a reference to the given string and assigns it to the CaFilePath field.
func (o *BulkWritableWebhookRequest) SetCaFilePath(v string) {
	o.CaFilePath = &v
}

func (o BulkWritableWebhookRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkWritableWebhookRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["content_types"] = o.ContentTypes
	toSerialize["name"] = o.Name
	if !IsNil(o.TypeCreate) {
		toSerialize["type_create"] = o.TypeCreate
	}
	if !IsNil(o.TypeUpdate) {
		toSerialize["type_update"] = o.TypeUpdate
	}
	if !IsNil(o.TypeDelete) {
		toSerialize["type_delete"] = o.TypeDelete
	}
	toSerialize["payload_url"] = o.PayloadUrl
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.HttpMethod) {
		toSerialize["http_method"] = o.HttpMethod
	}
	if !IsNil(o.HttpContentType) {
		toSerialize["http_content_type"] = o.HttpContentType
	}
	if !IsNil(o.AdditionalHeaders) {
		toSerialize["additional_headers"] = o.AdditionalHeaders
	}
	if !IsNil(o.BodyTemplate) {
		toSerialize["body_template"] = o.BodyTemplate
	}
	if !IsNil(o.Secret) {
		toSerialize["secret"] = o.Secret
	}
	if !IsNil(o.SslVerification) {
		toSerialize["ssl_verification"] = o.SslVerification
	}
	if !IsNil(o.CaFilePath) {
		toSerialize["ca_file_path"] = o.CaFilePath
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BulkWritableWebhookRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"content_types",
		"name",
		"payload_url",
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

	varBulkWritableWebhookRequest := _BulkWritableWebhookRequest{}

	err = json.Unmarshal(data, &varBulkWritableWebhookRequest)

	if err != nil {
		return err
	}

	*o = BulkWritableWebhookRequest(varBulkWritableWebhookRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "content_types")
		delete(additionalProperties, "name")
		delete(additionalProperties, "type_create")
		delete(additionalProperties, "type_update")
		delete(additionalProperties, "type_delete")
		delete(additionalProperties, "payload_url")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "http_method")
		delete(additionalProperties, "http_content_type")
		delete(additionalProperties, "additional_headers")
		delete(additionalProperties, "body_template")
		delete(additionalProperties, "secret")
		delete(additionalProperties, "ssl_verification")
		delete(additionalProperties, "ca_file_path")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBulkWritableWebhookRequest struct {
	value *BulkWritableWebhookRequest
	isSet bool
}

func (v NullableBulkWritableWebhookRequest) Get() *BulkWritableWebhookRequest {
	return v.value
}

func (v *NullableBulkWritableWebhookRequest) Set(val *BulkWritableWebhookRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkWritableWebhookRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkWritableWebhookRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkWritableWebhookRequest(val *BulkWritableWebhookRequest) *NullableBulkWritableWebhookRequest {
	return &NullableBulkWritableWebhookRequest{value: val, isSet: true}
}

func (v NullableBulkWritableWebhookRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkWritableWebhookRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

