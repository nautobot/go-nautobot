# BulkWritableImageAttachmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ContentType** | **string** |  | 
**ObjectId** | **string** |  | 
**Image** | ***os.File** |  | 
**ImageHeight** | **int32** |  | 
**ImageWidth** | **int32** |  | 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewBulkWritableImageAttachmentRequest

`func NewBulkWritableImageAttachmentRequest(id string, contentType string, objectId string, image *os.File, imageHeight int32, imageWidth int32, ) *BulkWritableImageAttachmentRequest`

NewBulkWritableImageAttachmentRequest instantiates a new BulkWritableImageAttachmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableImageAttachmentRequestWithDefaults

`func NewBulkWritableImageAttachmentRequestWithDefaults() *BulkWritableImageAttachmentRequest`

NewBulkWritableImageAttachmentRequestWithDefaults instantiates a new BulkWritableImageAttachmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableImageAttachmentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableImageAttachmentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableImageAttachmentRequest) SetId(v string)`

SetId sets Id field to given value.


### GetContentType

`func (o *BulkWritableImageAttachmentRequest) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *BulkWritableImageAttachmentRequest) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *BulkWritableImageAttachmentRequest) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetObjectId

`func (o *BulkWritableImageAttachmentRequest) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *BulkWritableImageAttachmentRequest) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *BulkWritableImageAttachmentRequest) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.


### GetImage

`func (o *BulkWritableImageAttachmentRequest) GetImage() *os.File`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *BulkWritableImageAttachmentRequest) GetImageOk() (**os.File, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *BulkWritableImageAttachmentRequest) SetImage(v *os.File)`

SetImage sets Image field to given value.


### GetImageHeight

`func (o *BulkWritableImageAttachmentRequest) GetImageHeight() int32`

GetImageHeight returns the ImageHeight field if non-nil, zero value otherwise.

### GetImageHeightOk

`func (o *BulkWritableImageAttachmentRequest) GetImageHeightOk() (*int32, bool)`

GetImageHeightOk returns a tuple with the ImageHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageHeight

`func (o *BulkWritableImageAttachmentRequest) SetImageHeight(v int32)`

SetImageHeight sets ImageHeight field to given value.


### GetImageWidth

`func (o *BulkWritableImageAttachmentRequest) GetImageWidth() int32`

GetImageWidth returns the ImageWidth field if non-nil, zero value otherwise.

### GetImageWidthOk

`func (o *BulkWritableImageAttachmentRequest) GetImageWidthOk() (*int32, bool)`

GetImageWidthOk returns a tuple with the ImageWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageWidth

`func (o *BulkWritableImageAttachmentRequest) SetImageWidth(v int32)`

SetImageWidth sets ImageWidth field to given value.


### GetName

`func (o *BulkWritableImageAttachmentRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkWritableImageAttachmentRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkWritableImageAttachmentRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BulkWritableImageAttachmentRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


