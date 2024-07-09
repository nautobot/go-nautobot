# PatchedBulkWritableImageAttachmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ContentType** | Pointer to **string** |  | [optional] 
**ObjectId** | Pointer to **string** |  | [optional] 
**Image** | Pointer to ***os.File** |  | [optional] 
**ImageHeight** | Pointer to **int32** |  | [optional] 
**ImageWidth** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchedBulkWritableImageAttachmentRequest

`func NewPatchedBulkWritableImageAttachmentRequest(id string, ) *PatchedBulkWritableImageAttachmentRequest`

NewPatchedBulkWritableImageAttachmentRequest instantiates a new PatchedBulkWritableImageAttachmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableImageAttachmentRequestWithDefaults

`func NewPatchedBulkWritableImageAttachmentRequestWithDefaults() *PatchedBulkWritableImageAttachmentRequest`

NewPatchedBulkWritableImageAttachmentRequestWithDefaults instantiates a new PatchedBulkWritableImageAttachmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableImageAttachmentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableImageAttachmentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableImageAttachmentRequest) SetId(v string)`

SetId sets Id field to given value.


### GetContentType

`func (o *PatchedBulkWritableImageAttachmentRequest) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *PatchedBulkWritableImageAttachmentRequest) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *PatchedBulkWritableImageAttachmentRequest) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *PatchedBulkWritableImageAttachmentRequest) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetObjectId

`func (o *PatchedBulkWritableImageAttachmentRequest) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *PatchedBulkWritableImageAttachmentRequest) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *PatchedBulkWritableImageAttachmentRequest) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *PatchedBulkWritableImageAttachmentRequest) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetImage

`func (o *PatchedBulkWritableImageAttachmentRequest) GetImage() *os.File`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *PatchedBulkWritableImageAttachmentRequest) GetImageOk() (**os.File, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *PatchedBulkWritableImageAttachmentRequest) SetImage(v *os.File)`

SetImage sets Image field to given value.

### HasImage

`func (o *PatchedBulkWritableImageAttachmentRequest) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetImageHeight

`func (o *PatchedBulkWritableImageAttachmentRequest) GetImageHeight() int32`

GetImageHeight returns the ImageHeight field if non-nil, zero value otherwise.

### GetImageHeightOk

`func (o *PatchedBulkWritableImageAttachmentRequest) GetImageHeightOk() (*int32, bool)`

GetImageHeightOk returns a tuple with the ImageHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageHeight

`func (o *PatchedBulkWritableImageAttachmentRequest) SetImageHeight(v int32)`

SetImageHeight sets ImageHeight field to given value.

### HasImageHeight

`func (o *PatchedBulkWritableImageAttachmentRequest) HasImageHeight() bool`

HasImageHeight returns a boolean if a field has been set.

### GetImageWidth

`func (o *PatchedBulkWritableImageAttachmentRequest) GetImageWidth() int32`

GetImageWidth returns the ImageWidth field if non-nil, zero value otherwise.

### GetImageWidthOk

`func (o *PatchedBulkWritableImageAttachmentRequest) GetImageWidthOk() (*int32, bool)`

GetImageWidthOk returns a tuple with the ImageWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageWidth

`func (o *PatchedBulkWritableImageAttachmentRequest) SetImageWidth(v int32)`

SetImageWidth sets ImageWidth field to given value.

### HasImageWidth

`func (o *PatchedBulkWritableImageAttachmentRequest) HasImageWidth() bool`

HasImageWidth returns a boolean if a field has been set.

### GetName

`func (o *PatchedBulkWritableImageAttachmentRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkWritableImageAttachmentRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkWritableImageAttachmentRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkWritableImageAttachmentRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


