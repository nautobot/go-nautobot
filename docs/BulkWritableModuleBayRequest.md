# BulkWritableModuleBayRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Position** | Pointer to **string** | The position of the module bay within the parent device/module | [optional] 
**Label** | Pointer to **string** | Physical label | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ParentDevice** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**ParentModule** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewBulkWritableModuleBayRequest

`func NewBulkWritableModuleBayRequest(id string, name string, ) *BulkWritableModuleBayRequest`

NewBulkWritableModuleBayRequest instantiates a new BulkWritableModuleBayRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableModuleBayRequestWithDefaults

`func NewBulkWritableModuleBayRequestWithDefaults() *BulkWritableModuleBayRequest`

NewBulkWritableModuleBayRequestWithDefaults instantiates a new BulkWritableModuleBayRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableModuleBayRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableModuleBayRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableModuleBayRequest) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BulkWritableModuleBayRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkWritableModuleBayRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkWritableModuleBayRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPosition

`func (o *BulkWritableModuleBayRequest) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *BulkWritableModuleBayRequest) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *BulkWritableModuleBayRequest) SetPosition(v string)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *BulkWritableModuleBayRequest) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetLabel

`func (o *BulkWritableModuleBayRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *BulkWritableModuleBayRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *BulkWritableModuleBayRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *BulkWritableModuleBayRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDescription

`func (o *BulkWritableModuleBayRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkWritableModuleBayRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkWritableModuleBayRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkWritableModuleBayRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParentDevice

`func (o *BulkWritableModuleBayRequest) GetParentDevice() BulkWritableCircuitRequestTenant`

GetParentDevice returns the ParentDevice field if non-nil, zero value otherwise.

### GetParentDeviceOk

`func (o *BulkWritableModuleBayRequest) GetParentDeviceOk() (*BulkWritableCircuitRequestTenant, bool)`

GetParentDeviceOk returns a tuple with the ParentDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentDevice

`func (o *BulkWritableModuleBayRequest) SetParentDevice(v BulkWritableCircuitRequestTenant)`

SetParentDevice sets ParentDevice field to given value.

### HasParentDevice

`func (o *BulkWritableModuleBayRequest) HasParentDevice() bool`

HasParentDevice returns a boolean if a field has been set.

### SetParentDeviceNil

`func (o *BulkWritableModuleBayRequest) SetParentDeviceNil(b bool)`

 SetParentDeviceNil sets the value for ParentDevice to be an explicit nil

### UnsetParentDevice
`func (o *BulkWritableModuleBayRequest) UnsetParentDevice()`

UnsetParentDevice ensures that no value is present for ParentDevice, not even an explicit nil
### GetParentModule

`func (o *BulkWritableModuleBayRequest) GetParentModule() BulkWritableCircuitRequestTenant`

GetParentModule returns the ParentModule field if non-nil, zero value otherwise.

### GetParentModuleOk

`func (o *BulkWritableModuleBayRequest) GetParentModuleOk() (*BulkWritableCircuitRequestTenant, bool)`

GetParentModuleOk returns a tuple with the ParentModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentModule

`func (o *BulkWritableModuleBayRequest) SetParentModule(v BulkWritableCircuitRequestTenant)`

SetParentModule sets ParentModule field to given value.

### HasParentModule

`func (o *BulkWritableModuleBayRequest) HasParentModule() bool`

HasParentModule returns a boolean if a field has been set.

### SetParentModuleNil

`func (o *BulkWritableModuleBayRequest) SetParentModuleNil(b bool)`

 SetParentModuleNil sets the value for ParentModule to be an explicit nil

### UnsetParentModule
`func (o *BulkWritableModuleBayRequest) UnsetParentModule()`

UnsetParentModule ensures that no value is present for ParentModule, not even an explicit nil
### GetCustomFields

`func (o *BulkWritableModuleBayRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkWritableModuleBayRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkWritableModuleBayRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkWritableModuleBayRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *BulkWritableModuleBayRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BulkWritableModuleBayRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BulkWritableModuleBayRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BulkWritableModuleBayRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *BulkWritableModuleBayRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkWritableModuleBayRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkWritableModuleBayRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkWritableModuleBayRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


