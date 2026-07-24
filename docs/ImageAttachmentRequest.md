# ImageAttachmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ContentType** | **string** |  | 
**ObjectId** | **string** |  | 
**Image** | ***os.File** |  | 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewImageAttachmentRequest

`func NewImageAttachmentRequest(contentType string, objectId string, image *os.File, ) *ImageAttachmentRequest`

NewImageAttachmentRequest instantiates a new ImageAttachmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageAttachmentRequestWithDefaults

`func NewImageAttachmentRequestWithDefaults() *ImageAttachmentRequest`

NewImageAttachmentRequestWithDefaults instantiates a new ImageAttachmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ImageAttachmentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImageAttachmentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImageAttachmentRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ImageAttachmentRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetContentType

`func (o *ImageAttachmentRequest) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ImageAttachmentRequest) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *ImageAttachmentRequest) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetObjectId

`func (o *ImageAttachmentRequest) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *ImageAttachmentRequest) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *ImageAttachmentRequest) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.


### GetImage

`func (o *ImageAttachmentRequest) GetImage() *os.File`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *ImageAttachmentRequest) GetImageOk() (**os.File, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *ImageAttachmentRequest) SetImage(v *os.File)`

SetImage sets Image field to given value.


### GetName

`func (o *ImageAttachmentRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImageAttachmentRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImageAttachmentRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ImageAttachmentRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


