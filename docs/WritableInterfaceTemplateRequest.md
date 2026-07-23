# WritableInterfaceTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Label** | Pointer to **string** | Physical label | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Type** | [**InterfaceTypeChoices**](InterfaceTypeChoices.md) |  | 
**PortType** | Pointer to [**PatchedWritableInterfaceRequestPortType**](PatchedWritableInterfaceRequestPortType.md) |  | [optional] 
**MgmtOnly** | Pointer to **bool** |  | [optional] 
**Speed** | Pointer to **NullableInt32** |  | [optional] 
**Duplex** | Pointer to [**BulkWritableInterfaceTemplateRequestDuplex**](BulkWritableInterfaceTemplateRequestDuplex.md) |  | [optional] 
**DeviceType** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**ModuleType** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewWritableInterfaceTemplateRequest

`func NewWritableInterfaceTemplateRequest(name string, type_ InterfaceTypeChoices, ) *WritableInterfaceTemplateRequest`

NewWritableInterfaceTemplateRequest instantiates a new WritableInterfaceTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritableInterfaceTemplateRequestWithDefaults

`func NewWritableInterfaceTemplateRequestWithDefaults() *WritableInterfaceTemplateRequest`

NewWritableInterfaceTemplateRequestWithDefaults instantiates a new WritableInterfaceTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WritableInterfaceTemplateRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WritableInterfaceTemplateRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WritableInterfaceTemplateRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WritableInterfaceTemplateRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *WritableInterfaceTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WritableInterfaceTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WritableInterfaceTemplateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *WritableInterfaceTemplateRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *WritableInterfaceTemplateRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *WritableInterfaceTemplateRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *WritableInterfaceTemplateRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDescription

`func (o *WritableInterfaceTemplateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WritableInterfaceTemplateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WritableInterfaceTemplateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WritableInterfaceTemplateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *WritableInterfaceTemplateRequest) GetType() InterfaceTypeChoices`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WritableInterfaceTemplateRequest) GetTypeOk() (*InterfaceTypeChoices, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WritableInterfaceTemplateRequest) SetType(v InterfaceTypeChoices)`

SetType sets Type field to given value.


### GetPortType

`func (o *WritableInterfaceTemplateRequest) GetPortType() PatchedWritableInterfaceRequestPortType`

GetPortType returns the PortType field if non-nil, zero value otherwise.

### GetPortTypeOk

`func (o *WritableInterfaceTemplateRequest) GetPortTypeOk() (*PatchedWritableInterfaceRequestPortType, bool)`

GetPortTypeOk returns a tuple with the PortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortType

`func (o *WritableInterfaceTemplateRequest) SetPortType(v PatchedWritableInterfaceRequestPortType)`

SetPortType sets PortType field to given value.

### HasPortType

`func (o *WritableInterfaceTemplateRequest) HasPortType() bool`

HasPortType returns a boolean if a field has been set.

### GetMgmtOnly

`func (o *WritableInterfaceTemplateRequest) GetMgmtOnly() bool`

GetMgmtOnly returns the MgmtOnly field if non-nil, zero value otherwise.

### GetMgmtOnlyOk

`func (o *WritableInterfaceTemplateRequest) GetMgmtOnlyOk() (*bool, bool)`

GetMgmtOnlyOk returns a tuple with the MgmtOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtOnly

`func (o *WritableInterfaceTemplateRequest) SetMgmtOnly(v bool)`

SetMgmtOnly sets MgmtOnly field to given value.

### HasMgmtOnly

`func (o *WritableInterfaceTemplateRequest) HasMgmtOnly() bool`

HasMgmtOnly returns a boolean if a field has been set.

### GetSpeed

`func (o *WritableInterfaceTemplateRequest) GetSpeed() int32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *WritableInterfaceTemplateRequest) GetSpeedOk() (*int32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *WritableInterfaceTemplateRequest) SetSpeed(v int32)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *WritableInterfaceTemplateRequest) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### SetSpeedNil

`func (o *WritableInterfaceTemplateRequest) SetSpeedNil(b bool)`

 SetSpeedNil sets the value for Speed to be an explicit nil

### UnsetSpeed
`func (o *WritableInterfaceTemplateRequest) UnsetSpeed()`

UnsetSpeed ensures that no value is present for Speed, not even an explicit nil
### GetDuplex

`func (o *WritableInterfaceTemplateRequest) GetDuplex() BulkWritableInterfaceTemplateRequestDuplex`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *WritableInterfaceTemplateRequest) GetDuplexOk() (*BulkWritableInterfaceTemplateRequestDuplex, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *WritableInterfaceTemplateRequest) SetDuplex(v BulkWritableInterfaceTemplateRequestDuplex)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *WritableInterfaceTemplateRequest) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetDeviceType

`func (o *WritableInterfaceTemplateRequest) GetDeviceType() ApprovalWorkflowUser`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *WritableInterfaceTemplateRequest) GetDeviceTypeOk() (*ApprovalWorkflowUser, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *WritableInterfaceTemplateRequest) SetDeviceType(v ApprovalWorkflowUser)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *WritableInterfaceTemplateRequest) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### SetDeviceTypeNil

`func (o *WritableInterfaceTemplateRequest) SetDeviceTypeNil(b bool)`

 SetDeviceTypeNil sets the value for DeviceType to be an explicit nil

### UnsetDeviceType
`func (o *WritableInterfaceTemplateRequest) UnsetDeviceType()`

UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
### GetModuleType

`func (o *WritableInterfaceTemplateRequest) GetModuleType() ApprovalWorkflowUser`

GetModuleType returns the ModuleType field if non-nil, zero value otherwise.

### GetModuleTypeOk

`func (o *WritableInterfaceTemplateRequest) GetModuleTypeOk() (*ApprovalWorkflowUser, bool)`

GetModuleTypeOk returns a tuple with the ModuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleType

`func (o *WritableInterfaceTemplateRequest) SetModuleType(v ApprovalWorkflowUser)`

SetModuleType sets ModuleType field to given value.

### HasModuleType

`func (o *WritableInterfaceTemplateRequest) HasModuleType() bool`

HasModuleType returns a boolean if a field has been set.

### SetModuleTypeNil

`func (o *WritableInterfaceTemplateRequest) SetModuleTypeNil(b bool)`

 SetModuleTypeNil sets the value for ModuleType to be an explicit nil

### UnsetModuleType
`func (o *WritableInterfaceTemplateRequest) UnsetModuleType()`

UnsetModuleType ensures that no value is present for ModuleType, not even an explicit nil
### GetCustomFields

`func (o *WritableInterfaceTemplateRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritableInterfaceTemplateRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritableInterfaceTemplateRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritableInterfaceTemplateRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *WritableInterfaceTemplateRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *WritableInterfaceTemplateRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *WritableInterfaceTemplateRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *WritableInterfaceTemplateRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


