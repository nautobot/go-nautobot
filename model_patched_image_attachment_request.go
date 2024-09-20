/*
API Documentation

Source of truth and network automation platform

API version: 2.3.4 (2.3)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nautobot

import (
	"encoding/json"
	"os"
)

// checks if the PatchedImageAttachmentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedImageAttachmentRequest{}

// PatchedImageAttachmentRequest Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type PatchedImageAttachmentRequest struct {
	ContentType *string `json:"content_type,omitempty"`
	ObjectId *string `json:"object_id,omitempty"`
	Image **os.File `json:"image,omitempty"`
	ImageHeight *int32 `json:"image_height,omitempty"`
	ImageWidth *int32 `json:"image_width,omitempty"`
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedImageAttachmentRequest PatchedImageAttachmentRequest

// NewPatchedImageAttachmentRequest instantiates a new PatchedImageAttachmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedImageAttachmentRequest() *PatchedImageAttachmentRequest {
	this := PatchedImageAttachmentRequest{}
	return &this
}

// NewPatchedImageAttachmentRequestWithDefaults instantiates a new PatchedImageAttachmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedImageAttachmentRequestWithDefaults() *PatchedImageAttachmentRequest {
	this := PatchedImageAttachmentRequest{}
	return &this
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *PatchedImageAttachmentRequest) GetContentType() string {
	if o == nil || IsNil(o.ContentType) {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedImageAttachmentRequest) GetContentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ContentType) {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *PatchedImageAttachmentRequest) HasContentType() bool {
	if o != nil && !IsNil(o.ContentType) {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *PatchedImageAttachmentRequest) SetContentType(v string) {
	o.ContentType = &v
}

// GetObjectId returns the ObjectId field value if set, zero value otherwise.
func (o *PatchedImageAttachmentRequest) GetObjectId() string {
	if o == nil || IsNil(o.ObjectId) {
		var ret string
		return ret
	}
	return *o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedImageAttachmentRequest) GetObjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectId) {
		return nil, false
	}
	return o.ObjectId, true
}

// HasObjectId returns a boolean if a field has been set.
func (o *PatchedImageAttachmentRequest) HasObjectId() bool {
	if o != nil && !IsNil(o.ObjectId) {
		return true
	}

	return false
}

// SetObjectId gets a reference to the given string and assigns it to the ObjectId field.
func (o *PatchedImageAttachmentRequest) SetObjectId(v string) {
	o.ObjectId = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *PatchedImageAttachmentRequest) GetImage() *os.File {
	if o == nil || IsNil(o.Image) {
		var ret *os.File
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedImageAttachmentRequest) GetImageOk() (**os.File, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *PatchedImageAttachmentRequest) HasImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given *os.File and assigns it to the Image field.
func (o *PatchedImageAttachmentRequest) SetImage(v *os.File) {
	o.Image = &v
}

// GetImageHeight returns the ImageHeight field value if set, zero value otherwise.
func (o *PatchedImageAttachmentRequest) GetImageHeight() int32 {
	if o == nil || IsNil(o.ImageHeight) {
		var ret int32
		return ret
	}
	return *o.ImageHeight
}

// GetImageHeightOk returns a tuple with the ImageHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedImageAttachmentRequest) GetImageHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.ImageHeight) {
		return nil, false
	}
	return o.ImageHeight, true
}

// HasImageHeight returns a boolean if a field has been set.
func (o *PatchedImageAttachmentRequest) HasImageHeight() bool {
	if o != nil && !IsNil(o.ImageHeight) {
		return true
	}

	return false
}

// SetImageHeight gets a reference to the given int32 and assigns it to the ImageHeight field.
func (o *PatchedImageAttachmentRequest) SetImageHeight(v int32) {
	o.ImageHeight = &v
}

// GetImageWidth returns the ImageWidth field value if set, zero value otherwise.
func (o *PatchedImageAttachmentRequest) GetImageWidth() int32 {
	if o == nil || IsNil(o.ImageWidth) {
		var ret int32
		return ret
	}
	return *o.ImageWidth
}

// GetImageWidthOk returns a tuple with the ImageWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedImageAttachmentRequest) GetImageWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.ImageWidth) {
		return nil, false
	}
	return o.ImageWidth, true
}

// HasImageWidth returns a boolean if a field has been set.
func (o *PatchedImageAttachmentRequest) HasImageWidth() bool {
	if o != nil && !IsNil(o.ImageWidth) {
		return true
	}

	return false
}

// SetImageWidth gets a reference to the given int32 and assigns it to the ImageWidth field.
func (o *PatchedImageAttachmentRequest) SetImageWidth(v int32) {
	o.ImageWidth = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedImageAttachmentRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedImageAttachmentRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedImageAttachmentRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedImageAttachmentRequest) SetName(v string) {
	o.Name = &v
}

func (o PatchedImageAttachmentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedImageAttachmentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContentType) {
		toSerialize["content_type"] = o.ContentType
	}
	if !IsNil(o.ObjectId) {
		toSerialize["object_id"] = o.ObjectId
	}
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	if !IsNil(o.ImageHeight) {
		toSerialize["image_height"] = o.ImageHeight
	}
	if !IsNil(o.ImageWidth) {
		toSerialize["image_width"] = o.ImageWidth
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedImageAttachmentRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedImageAttachmentRequest := _PatchedImageAttachmentRequest{}

	err = json.Unmarshal(data, &varPatchedImageAttachmentRequest)

	if err != nil {
		return err
	}

	*o = PatchedImageAttachmentRequest(varPatchedImageAttachmentRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "content_type")
		delete(additionalProperties, "object_id")
		delete(additionalProperties, "image")
		delete(additionalProperties, "image_height")
		delete(additionalProperties, "image_width")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedImageAttachmentRequest struct {
	value *PatchedImageAttachmentRequest
	isSet bool
}

func (v NullablePatchedImageAttachmentRequest) Get() *PatchedImageAttachmentRequest {
	return v.value
}

func (v *NullablePatchedImageAttachmentRequest) Set(val *PatchedImageAttachmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedImageAttachmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedImageAttachmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedImageAttachmentRequest(val *PatchedImageAttachmentRequest) *NullablePatchedImageAttachmentRequest {
	return &NullablePatchedImageAttachmentRequest{value: val, isSet: true}
}

func (v NullablePatchedImageAttachmentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedImageAttachmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


