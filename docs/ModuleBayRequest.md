# ModuleBayRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewModuleBayRequest

`func NewModuleBayRequest(name string, ) *ModuleBayRequest`

NewModuleBayRequest instantiates a new ModuleBayRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModuleBayRequestWithDefaults

`func NewModuleBayRequestWithDefaults() *ModuleBayRequest`

NewModuleBayRequestWithDefaults instantiates a new ModuleBayRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ModuleBayRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModuleBayRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModuleBayRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPosition

`func (o *ModuleBayRequest) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ModuleBayRequest) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ModuleBayRequest) SetPosition(v string)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *ModuleBayRequest) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetLabel

`func (o *ModuleBayRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ModuleBayRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ModuleBayRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ModuleBayRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDescription

`func (o *ModuleBayRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModuleBayRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModuleBayRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModuleBayRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParentDevice

`func (o *ModuleBayRequest) GetParentDevice() BulkWritableCircuitRequestTenant`

GetParentDevice returns the ParentDevice field if non-nil, zero value otherwise.

### GetParentDeviceOk

`func (o *ModuleBayRequest) GetParentDeviceOk() (*BulkWritableCircuitRequestTenant, bool)`

GetParentDeviceOk returns a tuple with the ParentDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentDevice

`func (o *ModuleBayRequest) SetParentDevice(v BulkWritableCircuitRequestTenant)`

SetParentDevice sets ParentDevice field to given value.

### HasParentDevice

`func (o *ModuleBayRequest) HasParentDevice() bool`

HasParentDevice returns a boolean if a field has been set.

### SetParentDeviceNil

`func (o *ModuleBayRequest) SetParentDeviceNil(b bool)`

 SetParentDeviceNil sets the value for ParentDevice to be an explicit nil

### UnsetParentDevice
`func (o *ModuleBayRequest) UnsetParentDevice()`

UnsetParentDevice ensures that no value is present for ParentDevice, not even an explicit nil
### GetParentModule

`func (o *ModuleBayRequest) GetParentModule() BulkWritableCircuitRequestTenant`

GetParentModule returns the ParentModule field if non-nil, zero value otherwise.

### GetParentModuleOk

`func (o *ModuleBayRequest) GetParentModuleOk() (*BulkWritableCircuitRequestTenant, bool)`

GetParentModuleOk returns a tuple with the ParentModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentModule

`func (o *ModuleBayRequest) SetParentModule(v BulkWritableCircuitRequestTenant)`

SetParentModule sets ParentModule field to given value.

### HasParentModule

`func (o *ModuleBayRequest) HasParentModule() bool`

HasParentModule returns a boolean if a field has been set.

### SetParentModuleNil

`func (o *ModuleBayRequest) SetParentModuleNil(b bool)`

 SetParentModuleNil sets the value for ParentModule to be an explicit nil

### UnsetParentModule
`func (o *ModuleBayRequest) UnsetParentModule()`

UnsetParentModule ensures that no value is present for ParentModule, not even an explicit nil
### GetCustomFields

`func (o *ModuleBayRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ModuleBayRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ModuleBayRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ModuleBayRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *ModuleBayRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ModuleBayRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ModuleBayRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ModuleBayRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *ModuleBayRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ModuleBayRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ModuleBayRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ModuleBayRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


