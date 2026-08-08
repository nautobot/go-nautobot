# PatchedCloudServiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ExtraConfig** | Pointer to **interface{}** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**CloudResourceType** | Pointer to [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 
**CloudAccount** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 
**Tags** | Pointer to [**[]ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 

## Methods

### NewPatchedCloudServiceRequest

`func NewPatchedCloudServiceRequest() *PatchedCloudServiceRequest`

NewPatchedCloudServiceRequest instantiates a new PatchedCloudServiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedCloudServiceRequestWithDefaults

`func NewPatchedCloudServiceRequestWithDefaults() *PatchedCloudServiceRequest`

NewPatchedCloudServiceRequestWithDefaults instantiates a new PatchedCloudServiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedCloudServiceRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedCloudServiceRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedCloudServiceRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedCloudServiceRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetExtraConfig

`func (o *PatchedCloudServiceRequest) GetExtraConfig() interface{}`

GetExtraConfig returns the ExtraConfig field if non-nil, zero value otherwise.

### GetExtraConfigOk

`func (o *PatchedCloudServiceRequest) GetExtraConfigOk() (*interface{}, bool)`

GetExtraConfigOk returns a tuple with the ExtraConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraConfig

`func (o *PatchedCloudServiceRequest) SetExtraConfig(v interface{})`

SetExtraConfig sets ExtraConfig field to given value.

### HasExtraConfig

`func (o *PatchedCloudServiceRequest) HasExtraConfig() bool`

HasExtraConfig returns a boolean if a field has been set.

### SetExtraConfigNil

`func (o *PatchedCloudServiceRequest) SetExtraConfigNil(b bool)`

 SetExtraConfigNil sets the value for ExtraConfig to be an explicit nil

### UnsetExtraConfig
`func (o *PatchedCloudServiceRequest) UnsetExtraConfig()`

UnsetExtraConfig ensures that no value is present for ExtraConfig, not even an explicit nil
### GetName

`func (o *PatchedCloudServiceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedCloudServiceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedCloudServiceRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedCloudServiceRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedCloudServiceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedCloudServiceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedCloudServiceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedCloudServiceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCloudResourceType

`func (o *PatchedCloudServiceRequest) GetCloudResourceType() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetCloudResourceType returns the CloudResourceType field if non-nil, zero value otherwise.

### GetCloudResourceTypeOk

`func (o *PatchedCloudServiceRequest) GetCloudResourceTypeOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetCloudResourceTypeOk returns a tuple with the CloudResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudResourceType

`func (o *PatchedCloudServiceRequest) SetCloudResourceType(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetCloudResourceType sets CloudResourceType field to given value.

### HasCloudResourceType

`func (o *PatchedCloudServiceRequest) HasCloudResourceType() bool`

HasCloudResourceType returns a boolean if a field has been set.

### GetCloudAccount

`func (o *PatchedCloudServiceRequest) GetCloudAccount() ApprovalWorkflowUser`

GetCloudAccount returns the CloudAccount field if non-nil, zero value otherwise.

### GetCloudAccountOk

`func (o *PatchedCloudServiceRequest) GetCloudAccountOk() (*ApprovalWorkflowUser, bool)`

GetCloudAccountOk returns a tuple with the CloudAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudAccount

`func (o *PatchedCloudServiceRequest) SetCloudAccount(v ApprovalWorkflowUser)`

SetCloudAccount sets CloudAccount field to given value.

### HasCloudAccount

`func (o *PatchedCloudServiceRequest) HasCloudAccount() bool`

HasCloudAccount returns a boolean if a field has been set.

### SetCloudAccountNil

`func (o *PatchedCloudServiceRequest) SetCloudAccountNil(b bool)`

 SetCloudAccountNil sets the value for CloudAccount to be an explicit nil

### UnsetCloudAccount
`func (o *PatchedCloudServiceRequest) UnsetCloudAccount()`

UnsetCloudAccount ensures that no value is present for CloudAccount, not even an explicit nil
### GetCustomFields

`func (o *PatchedCloudServiceRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedCloudServiceRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedCloudServiceRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedCloudServiceRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedCloudServiceRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedCloudServiceRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedCloudServiceRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedCloudServiceRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *PatchedCloudServiceRequest) GetTags() []ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedCloudServiceRequest) GetTagsOk() (*[]ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedCloudServiceRequest) SetTags(v []ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedCloudServiceRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


