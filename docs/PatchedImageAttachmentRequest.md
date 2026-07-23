# PatchedImageAttachmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ContentType** | Pointer to **string** |  | [optional] 
**ObjectId** | Pointer to **string** |  | [optional] 
**Image** | Pointer to ***os.File** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchedImageAttachmentRequest

`func NewPatchedImageAttachmentRequest() *PatchedImageAttachmentRequest`

NewPatchedImageAttachmentRequest instantiates a new PatchedImageAttachmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedImageAttachmentRequestWithDefaults

`func NewPatchedImageAttachmentRequestWithDefaults() *PatchedImageAttachmentRequest`

NewPatchedImageAttachmentRequestWithDefaults instantiates a new PatchedImageAttachmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedImageAttachmentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedImageAttachmentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedImageAttachmentRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedImageAttachmentRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetContentType

`func (o *PatchedImageAttachmentRequest) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *PatchedImageAttachmentRequest) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *PatchedImageAttachmentRequest) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *PatchedImageAttachmentRequest) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetObjectId

`func (o *PatchedImageAttachmentRequest) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *PatchedImageAttachmentRequest) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *PatchedImageAttachmentRequest) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *PatchedImageAttachmentRequest) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetImage

`func (o *PatchedImageAttachmentRequest) GetImage() *os.File`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *PatchedImageAttachmentRequest) GetImageOk() (**os.File, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *PatchedImageAttachmentRequest) SetImage(v *os.File)`

SetImage sets Image field to given value.

### HasImage

`func (o *PatchedImageAttachmentRequest) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetName

`func (o *PatchedImageAttachmentRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedImageAttachmentRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedImageAttachmentRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedImageAttachmentRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


