# ModuleBayTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Position** | Pointer to **string** | The position of the module bay within the device or module | [optional] 
**Label** | Pointer to **string** | Physical label | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**RequiresFirstPartyModules** | Pointer to **bool** | This bay will only accept modules from the same manufacturer as the parent device or module | [optional] 
**DeviceType** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**ModuleType** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**ModuleFamily** | Pointer to [**NullableBulkWritableModuleBayTemplateRequestModuleFamily**](BulkWritableModuleBayTemplateRequestModuleFamily.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewModuleBayTemplateRequest

`func NewModuleBayTemplateRequest(name string, ) *ModuleBayTemplateRequest`

NewModuleBayTemplateRequest instantiates a new ModuleBayTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModuleBayTemplateRequestWithDefaults

`func NewModuleBayTemplateRequestWithDefaults() *ModuleBayTemplateRequest`

NewModuleBayTemplateRequestWithDefaults instantiates a new ModuleBayTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModuleBayTemplateRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModuleBayTemplateRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModuleBayTemplateRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModuleBayTemplateRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ModuleBayTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModuleBayTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModuleBayTemplateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPosition

`func (o *ModuleBayTemplateRequest) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ModuleBayTemplateRequest) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ModuleBayTemplateRequest) SetPosition(v string)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *ModuleBayTemplateRequest) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetLabel

`func (o *ModuleBayTemplateRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ModuleBayTemplateRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ModuleBayTemplateRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ModuleBayTemplateRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDescription

`func (o *ModuleBayTemplateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModuleBayTemplateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModuleBayTemplateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModuleBayTemplateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRequiresFirstPartyModules

`func (o *ModuleBayTemplateRequest) GetRequiresFirstPartyModules() bool`

GetRequiresFirstPartyModules returns the RequiresFirstPartyModules field if non-nil, zero value otherwise.

### GetRequiresFirstPartyModulesOk

`func (o *ModuleBayTemplateRequest) GetRequiresFirstPartyModulesOk() (*bool, bool)`

GetRequiresFirstPartyModulesOk returns a tuple with the RequiresFirstPartyModules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresFirstPartyModules

`func (o *ModuleBayTemplateRequest) SetRequiresFirstPartyModules(v bool)`

SetRequiresFirstPartyModules sets RequiresFirstPartyModules field to given value.

### HasRequiresFirstPartyModules

`func (o *ModuleBayTemplateRequest) HasRequiresFirstPartyModules() bool`

HasRequiresFirstPartyModules returns a boolean if a field has been set.

### GetDeviceType

`func (o *ModuleBayTemplateRequest) GetDeviceType() ApprovalWorkflowUser`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *ModuleBayTemplateRequest) GetDeviceTypeOk() (*ApprovalWorkflowUser, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *ModuleBayTemplateRequest) SetDeviceType(v ApprovalWorkflowUser)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *ModuleBayTemplateRequest) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### SetDeviceTypeNil

`func (o *ModuleBayTemplateRequest) SetDeviceTypeNil(b bool)`

 SetDeviceTypeNil sets the value for DeviceType to be an explicit nil

### UnsetDeviceType
`func (o *ModuleBayTemplateRequest) UnsetDeviceType()`

UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
### GetModuleType

`func (o *ModuleBayTemplateRequest) GetModuleType() ApprovalWorkflowUser`

GetModuleType returns the ModuleType field if non-nil, zero value otherwise.

### GetModuleTypeOk

`func (o *ModuleBayTemplateRequest) GetModuleTypeOk() (*ApprovalWorkflowUser, bool)`

GetModuleTypeOk returns a tuple with the ModuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleType

`func (o *ModuleBayTemplateRequest) SetModuleType(v ApprovalWorkflowUser)`

SetModuleType sets ModuleType field to given value.

### HasModuleType

`func (o *ModuleBayTemplateRequest) HasModuleType() bool`

HasModuleType returns a boolean if a field has been set.

### SetModuleTypeNil

`func (o *ModuleBayTemplateRequest) SetModuleTypeNil(b bool)`

 SetModuleTypeNil sets the value for ModuleType to be an explicit nil

### UnsetModuleType
`func (o *ModuleBayTemplateRequest) UnsetModuleType()`

UnsetModuleType ensures that no value is present for ModuleType, not even an explicit nil
### GetModuleFamily

`func (o *ModuleBayTemplateRequest) GetModuleFamily() BulkWritableModuleBayTemplateRequestModuleFamily`

GetModuleFamily returns the ModuleFamily field if non-nil, zero value otherwise.

### GetModuleFamilyOk

`func (o *ModuleBayTemplateRequest) GetModuleFamilyOk() (*BulkWritableModuleBayTemplateRequestModuleFamily, bool)`

GetModuleFamilyOk returns a tuple with the ModuleFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleFamily

`func (o *ModuleBayTemplateRequest) SetModuleFamily(v BulkWritableModuleBayTemplateRequestModuleFamily)`

SetModuleFamily sets ModuleFamily field to given value.

### HasModuleFamily

`func (o *ModuleBayTemplateRequest) HasModuleFamily() bool`

HasModuleFamily returns a boolean if a field has been set.

### SetModuleFamilyNil

`func (o *ModuleBayTemplateRequest) SetModuleFamilyNil(b bool)`

 SetModuleFamilyNil sets the value for ModuleFamily to be an explicit nil

### UnsetModuleFamily
`func (o *ModuleBayTemplateRequest) UnsetModuleFamily()`

UnsetModuleFamily ensures that no value is present for ModuleFamily, not even an explicit nil
### GetCustomFields

`func (o *ModuleBayTemplateRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ModuleBayTemplateRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ModuleBayTemplateRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ModuleBayTemplateRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *ModuleBayTemplateRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ModuleBayTemplateRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ModuleBayTemplateRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ModuleBayTemplateRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


