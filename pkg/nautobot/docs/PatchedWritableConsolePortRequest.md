# PatchedWritableConsolePortRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** | Physical label | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**PatchedWritableConsolePortRequestType**](PatchedWritableConsolePortRequestType.md) |  | [optional] 
**Device** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewPatchedWritableConsolePortRequest

`func NewPatchedWritableConsolePortRequest() *PatchedWritableConsolePortRequest`

NewPatchedWritableConsolePortRequest instantiates a new PatchedWritableConsolePortRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableConsolePortRequestWithDefaults

`func NewPatchedWritableConsolePortRequestWithDefaults() *PatchedWritableConsolePortRequest`

NewPatchedWritableConsolePortRequestWithDefaults instantiates a new PatchedWritableConsolePortRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedWritableConsolePortRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedWritableConsolePortRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedWritableConsolePortRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedWritableConsolePortRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabel

`func (o *PatchedWritableConsolePortRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PatchedWritableConsolePortRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PatchedWritableConsolePortRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PatchedWritableConsolePortRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedWritableConsolePortRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedWritableConsolePortRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedWritableConsolePortRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedWritableConsolePortRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *PatchedWritableConsolePortRequest) GetType() PatchedWritableConsolePortRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedWritableConsolePortRequest) GetTypeOk() (*PatchedWritableConsolePortRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedWritableConsolePortRequest) SetType(v PatchedWritableConsolePortRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedWritableConsolePortRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDevice

`func (o *PatchedWritableConsolePortRequest) GetDevice() BulkWritableCableRequestStatus`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *PatchedWritableConsolePortRequest) GetDeviceOk() (*BulkWritableCableRequestStatus, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *PatchedWritableConsolePortRequest) SetDevice(v BulkWritableCableRequestStatus)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *PatchedWritableConsolePortRequest) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetTags

`func (o *PatchedWritableConsolePortRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedWritableConsolePortRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedWritableConsolePortRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedWritableConsolePortRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedWritableConsolePortRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritableConsolePortRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritableConsolePortRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritableConsolePortRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedWritableConsolePortRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedWritableConsolePortRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedWritableConsolePortRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedWritableConsolePortRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


