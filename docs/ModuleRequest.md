# ModuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Serial** | Pointer to **NullableString** |  | [optional] 
**AssetTag** | Pointer to **NullableString** | A unique tag used to identify this module | [optional] 
**ModuleType** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 
**ParentModuleBay** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**Status** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 
**Role** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**Tenant** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**Location** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 
**Tags** | Pointer to [**[]ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 

## Methods

### NewModuleRequest

`func NewModuleRequest(moduleType ApprovalWorkflowStageResponseApprovalWorkflowStage, status ApprovalWorkflowStageResponseApprovalWorkflowStage, ) *ModuleRequest`

NewModuleRequest instantiates a new ModuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModuleRequestWithDefaults

`func NewModuleRequestWithDefaults() *ModuleRequest`

NewModuleRequestWithDefaults instantiates a new ModuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModuleRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModuleRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModuleRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModuleRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSerial

`func (o *ModuleRequest) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *ModuleRequest) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *ModuleRequest) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *ModuleRequest) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### SetSerialNil

`func (o *ModuleRequest) SetSerialNil(b bool)`

 SetSerialNil sets the value for Serial to be an explicit nil

### UnsetSerial
`func (o *ModuleRequest) UnsetSerial()`

UnsetSerial ensures that no value is present for Serial, not even an explicit nil
### GetAssetTag

`func (o *ModuleRequest) GetAssetTag() string`

GetAssetTag returns the AssetTag field if non-nil, zero value otherwise.

### GetAssetTagOk

`func (o *ModuleRequest) GetAssetTagOk() (*string, bool)`

GetAssetTagOk returns a tuple with the AssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTag

`func (o *ModuleRequest) SetAssetTag(v string)`

SetAssetTag sets AssetTag field to given value.

### HasAssetTag

`func (o *ModuleRequest) HasAssetTag() bool`

HasAssetTag returns a boolean if a field has been set.

### SetAssetTagNil

`func (o *ModuleRequest) SetAssetTagNil(b bool)`

 SetAssetTagNil sets the value for AssetTag to be an explicit nil

### UnsetAssetTag
`func (o *ModuleRequest) UnsetAssetTag()`

UnsetAssetTag ensures that no value is present for AssetTag, not even an explicit nil
### GetModuleType

`func (o *ModuleRequest) GetModuleType() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetModuleType returns the ModuleType field if non-nil, zero value otherwise.

### GetModuleTypeOk

`func (o *ModuleRequest) GetModuleTypeOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetModuleTypeOk returns a tuple with the ModuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleType

`func (o *ModuleRequest) SetModuleType(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetModuleType sets ModuleType field to given value.


### GetParentModuleBay

`func (o *ModuleRequest) GetParentModuleBay() ApprovalWorkflowUser`

GetParentModuleBay returns the ParentModuleBay field if non-nil, zero value otherwise.

### GetParentModuleBayOk

`func (o *ModuleRequest) GetParentModuleBayOk() (*ApprovalWorkflowUser, bool)`

GetParentModuleBayOk returns a tuple with the ParentModuleBay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentModuleBay

`func (o *ModuleRequest) SetParentModuleBay(v ApprovalWorkflowUser)`

SetParentModuleBay sets ParentModuleBay field to given value.

### HasParentModuleBay

`func (o *ModuleRequest) HasParentModuleBay() bool`

HasParentModuleBay returns a boolean if a field has been set.

### SetParentModuleBayNil

`func (o *ModuleRequest) SetParentModuleBayNil(b bool)`

 SetParentModuleBayNil sets the value for ParentModuleBay to be an explicit nil

### UnsetParentModuleBay
`func (o *ModuleRequest) UnsetParentModuleBay()`

UnsetParentModuleBay ensures that no value is present for ParentModuleBay, not even an explicit nil
### GetStatus

`func (o *ModuleRequest) GetStatus() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModuleRequest) GetStatusOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModuleRequest) SetStatus(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetStatus sets Status field to given value.


### GetRole

`func (o *ModuleRequest) GetRole() ApprovalWorkflowUser`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ModuleRequest) GetRoleOk() (*ApprovalWorkflowUser, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ModuleRequest) SetRole(v ApprovalWorkflowUser)`

SetRole sets Role field to given value.

### HasRole

`func (o *ModuleRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *ModuleRequest) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *ModuleRequest) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetTenant

`func (o *ModuleRequest) GetTenant() ApprovalWorkflowUser`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *ModuleRequest) GetTenantOk() (*ApprovalWorkflowUser, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *ModuleRequest) SetTenant(v ApprovalWorkflowUser)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *ModuleRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *ModuleRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *ModuleRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetLocation

`func (o *ModuleRequest) GetLocation() ApprovalWorkflowUser`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ModuleRequest) GetLocationOk() (*ApprovalWorkflowUser, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ModuleRequest) SetLocation(v ApprovalWorkflowUser)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ModuleRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *ModuleRequest) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *ModuleRequest) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetCustomFields

`func (o *ModuleRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ModuleRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ModuleRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ModuleRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *ModuleRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ModuleRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ModuleRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ModuleRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *ModuleRequest) GetTags() []ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ModuleRequest) GetTagsOk() (*[]ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ModuleRequest) SetTags(v []ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ModuleRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


