# PatchedBulkWritableExportTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ContentType** | Pointer to **string** |  | [optional] 
**OwnerContentType** | Pointer to **NullableString** |  | [optional] 
**OwnerObjectId** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**TemplateCode** | Pointer to **string** | The list of objects being exported is passed as a context variable named &lt;code&gt;queryset&lt;/code&gt;. | [optional] 
**MimeType** | Pointer to **string** | Defaults to &lt;code&gt;text/plain&lt;/code&gt; | [optional] 
**FileExtension** | Pointer to **string** | Extension to append to the rendered filename | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewPatchedBulkWritableExportTemplateRequest

`func NewPatchedBulkWritableExportTemplateRequest(id string, ) *PatchedBulkWritableExportTemplateRequest`

NewPatchedBulkWritableExportTemplateRequest instantiates a new PatchedBulkWritableExportTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableExportTemplateRequestWithDefaults

`func NewPatchedBulkWritableExportTemplateRequestWithDefaults() *PatchedBulkWritableExportTemplateRequest`

NewPatchedBulkWritableExportTemplateRequestWithDefaults instantiates a new PatchedBulkWritableExportTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableExportTemplateRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableExportTemplateRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableExportTemplateRequest) SetId(v string)`

SetId sets Id field to given value.


### GetContentType

`func (o *PatchedBulkWritableExportTemplateRequest) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *PatchedBulkWritableExportTemplateRequest) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *PatchedBulkWritableExportTemplateRequest) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *PatchedBulkWritableExportTemplateRequest) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetOwnerContentType

`func (o *PatchedBulkWritableExportTemplateRequest) GetOwnerContentType() string`

GetOwnerContentType returns the OwnerContentType field if non-nil, zero value otherwise.

### GetOwnerContentTypeOk

`func (o *PatchedBulkWritableExportTemplateRequest) GetOwnerContentTypeOk() (*string, bool)`

GetOwnerContentTypeOk returns a tuple with the OwnerContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerContentType

`func (o *PatchedBulkWritableExportTemplateRequest) SetOwnerContentType(v string)`

SetOwnerContentType sets OwnerContentType field to given value.

### HasOwnerContentType

`func (o *PatchedBulkWritableExportTemplateRequest) HasOwnerContentType() bool`

HasOwnerContentType returns a boolean if a field has been set.

### SetOwnerContentTypeNil

`func (o *PatchedBulkWritableExportTemplateRequest) SetOwnerContentTypeNil(b bool)`

 SetOwnerContentTypeNil sets the value for OwnerContentType to be an explicit nil

### UnsetOwnerContentType
`func (o *PatchedBulkWritableExportTemplateRequest) UnsetOwnerContentType()`

UnsetOwnerContentType ensures that no value is present for OwnerContentType, not even an explicit nil
### GetOwnerObjectId

`func (o *PatchedBulkWritableExportTemplateRequest) GetOwnerObjectId() string`

GetOwnerObjectId returns the OwnerObjectId field if non-nil, zero value otherwise.

### GetOwnerObjectIdOk

`func (o *PatchedBulkWritableExportTemplateRequest) GetOwnerObjectIdOk() (*string, bool)`

GetOwnerObjectIdOk returns a tuple with the OwnerObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerObjectId

`func (o *PatchedBulkWritableExportTemplateRequest) SetOwnerObjectId(v string)`

SetOwnerObjectId sets OwnerObjectId field to given value.

### HasOwnerObjectId

`func (o *PatchedBulkWritableExportTemplateRequest) HasOwnerObjectId() bool`

HasOwnerObjectId returns a boolean if a field has been set.

### SetOwnerObjectIdNil

`func (o *PatchedBulkWritableExportTemplateRequest) SetOwnerObjectIdNil(b bool)`

 SetOwnerObjectIdNil sets the value for OwnerObjectId to be an explicit nil

### UnsetOwnerObjectId
`func (o *PatchedBulkWritableExportTemplateRequest) UnsetOwnerObjectId()`

UnsetOwnerObjectId ensures that no value is present for OwnerObjectId, not even an explicit nil
### GetName

`func (o *PatchedBulkWritableExportTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkWritableExportTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkWritableExportTemplateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkWritableExportTemplateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedBulkWritableExportTemplateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkWritableExportTemplateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkWritableExportTemplateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkWritableExportTemplateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTemplateCode

`func (o *PatchedBulkWritableExportTemplateRequest) GetTemplateCode() string`

GetTemplateCode returns the TemplateCode field if non-nil, zero value otherwise.

### GetTemplateCodeOk

`func (o *PatchedBulkWritableExportTemplateRequest) GetTemplateCodeOk() (*string, bool)`

GetTemplateCodeOk returns a tuple with the TemplateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateCode

`func (o *PatchedBulkWritableExportTemplateRequest) SetTemplateCode(v string)`

SetTemplateCode sets TemplateCode field to given value.

### HasTemplateCode

`func (o *PatchedBulkWritableExportTemplateRequest) HasTemplateCode() bool`

HasTemplateCode returns a boolean if a field has been set.

### GetMimeType

`func (o *PatchedBulkWritableExportTemplateRequest) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *PatchedBulkWritableExportTemplateRequest) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *PatchedBulkWritableExportTemplateRequest) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *PatchedBulkWritableExportTemplateRequest) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetFileExtension

`func (o *PatchedBulkWritableExportTemplateRequest) GetFileExtension() string`

GetFileExtension returns the FileExtension field if non-nil, zero value otherwise.

### GetFileExtensionOk

`func (o *PatchedBulkWritableExportTemplateRequest) GetFileExtensionOk() (*string, bool)`

GetFileExtensionOk returns a tuple with the FileExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileExtension

`func (o *PatchedBulkWritableExportTemplateRequest) SetFileExtension(v string)`

SetFileExtension sets FileExtension field to given value.

### HasFileExtension

`func (o *PatchedBulkWritableExportTemplateRequest) HasFileExtension() bool`

HasFileExtension returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedBulkWritableExportTemplateRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedBulkWritableExportTemplateRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedBulkWritableExportTemplateRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedBulkWritableExportTemplateRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


