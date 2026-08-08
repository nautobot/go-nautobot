# PatchedVPNProfilePhase2PolicyAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Weight** | Pointer to **int32** | Higher weights appear later in the list | [optional] 
**VpnProfile** | Pointer to [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 
**VpnPhase2Policy** | Pointer to [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewPatchedVPNProfilePhase2PolicyAssignmentRequest

`func NewPatchedVPNProfilePhase2PolicyAssignmentRequest() *PatchedVPNProfilePhase2PolicyAssignmentRequest`

NewPatchedVPNProfilePhase2PolicyAssignmentRequest instantiates a new PatchedVPNProfilePhase2PolicyAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedVPNProfilePhase2PolicyAssignmentRequestWithDefaults

`func NewPatchedVPNProfilePhase2PolicyAssignmentRequestWithDefaults() *PatchedVPNProfilePhase2PolicyAssignmentRequest`

NewPatchedVPNProfilePhase2PolicyAssignmentRequestWithDefaults instantiates a new PatchedVPNProfilePhase2PolicyAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedVPNProfilePhase2PolicyAssignmentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedVPNProfilePhase2PolicyAssignmentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedVPNProfilePhase2PolicyAssignmentRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedVPNProfilePhase2PolicyAssignmentRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetWeight

`func (o *PatchedVPNProfilePhase2PolicyAssignmentRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *PatchedVPNProfilePhase2PolicyAssignmentRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *PatchedVPNProfilePhase2PolicyAssignmentRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *PatchedVPNProfilePhase2PolicyAssignmentRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetVpnProfile

`func (o *PatchedVPNProfilePhase2PolicyAssignmentRequest) GetVpnProfile() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetVpnProfile returns the VpnProfile field if non-nil, zero value otherwise.

### GetVpnProfileOk

`func (o *PatchedVPNProfilePhase2PolicyAssignmentRequest) GetVpnProfileOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetVpnProfileOk returns a tuple with the VpnProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnProfile

`func (o *PatchedVPNProfilePhase2PolicyAssignmentRequest) SetVpnProfile(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetVpnProfile sets VpnProfile field to given value.

### HasVpnProfile

`func (o *PatchedVPNProfilePhase2PolicyAssignmentRequest) HasVpnProfile() bool`

HasVpnProfile returns a boolean if a field has been set.

### GetVpnPhase2Policy

`func (o *PatchedVPNProfilePhase2PolicyAssignmentRequest) GetVpnPhase2Policy() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetVpnPhase2Policy returns the VpnPhase2Policy field if non-nil, zero value otherwise.

### GetVpnPhase2PolicyOk

`func (o *PatchedVPNProfilePhase2PolicyAssignmentRequest) GetVpnPhase2PolicyOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetVpnPhase2PolicyOk returns a tuple with the VpnPhase2Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnPhase2Policy

`func (o *PatchedVPNProfilePhase2PolicyAssignmentRequest) SetVpnPhase2Policy(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetVpnPhase2Policy sets VpnPhase2Policy field to given value.

### HasVpnPhase2Policy

`func (o *PatchedVPNProfilePhase2PolicyAssignmentRequest) HasVpnPhase2Policy() bool`

HasVpnPhase2Policy returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedVPNProfilePhase2PolicyAssignmentRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedVPNProfilePhase2PolicyAssignmentRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedVPNProfilePhase2PolicyAssignmentRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedVPNProfilePhase2PolicyAssignmentRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedVPNProfilePhase2PolicyAssignmentRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedVPNProfilePhase2PolicyAssignmentRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedVPNProfilePhase2PolicyAssignmentRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedVPNProfilePhase2PolicyAssignmentRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


