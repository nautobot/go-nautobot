# PatchedVPNProfilePhase1PolicyAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Weight** | Pointer to **int32** | Higher weights appear later in the list | [optional] 
**VpnProfile** | Pointer to [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 
**VpnPhase1Policy** | Pointer to [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewPatchedVPNProfilePhase1PolicyAssignmentRequest

`func NewPatchedVPNProfilePhase1PolicyAssignmentRequest() *PatchedVPNProfilePhase1PolicyAssignmentRequest`

NewPatchedVPNProfilePhase1PolicyAssignmentRequest instantiates a new PatchedVPNProfilePhase1PolicyAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedVPNProfilePhase1PolicyAssignmentRequestWithDefaults

`func NewPatchedVPNProfilePhase1PolicyAssignmentRequestWithDefaults() *PatchedVPNProfilePhase1PolicyAssignmentRequest`

NewPatchedVPNProfilePhase1PolicyAssignmentRequestWithDefaults instantiates a new PatchedVPNProfilePhase1PolicyAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedVPNProfilePhase1PolicyAssignmentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedVPNProfilePhase1PolicyAssignmentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedVPNProfilePhase1PolicyAssignmentRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedVPNProfilePhase1PolicyAssignmentRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetWeight

`func (o *PatchedVPNProfilePhase1PolicyAssignmentRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *PatchedVPNProfilePhase1PolicyAssignmentRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *PatchedVPNProfilePhase1PolicyAssignmentRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *PatchedVPNProfilePhase1PolicyAssignmentRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetVpnProfile

`func (o *PatchedVPNProfilePhase1PolicyAssignmentRequest) GetVpnProfile() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetVpnProfile returns the VpnProfile field if non-nil, zero value otherwise.

### GetVpnProfileOk

`func (o *PatchedVPNProfilePhase1PolicyAssignmentRequest) GetVpnProfileOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetVpnProfileOk returns a tuple with the VpnProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnProfile

`func (o *PatchedVPNProfilePhase1PolicyAssignmentRequest) SetVpnProfile(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetVpnProfile sets VpnProfile field to given value.

### HasVpnProfile

`func (o *PatchedVPNProfilePhase1PolicyAssignmentRequest) HasVpnProfile() bool`

HasVpnProfile returns a boolean if a field has been set.

### GetVpnPhase1Policy

`func (o *PatchedVPNProfilePhase1PolicyAssignmentRequest) GetVpnPhase1Policy() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetVpnPhase1Policy returns the VpnPhase1Policy field if non-nil, zero value otherwise.

### GetVpnPhase1PolicyOk

`func (o *PatchedVPNProfilePhase1PolicyAssignmentRequest) GetVpnPhase1PolicyOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetVpnPhase1PolicyOk returns a tuple with the VpnPhase1Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnPhase1Policy

`func (o *PatchedVPNProfilePhase1PolicyAssignmentRequest) SetVpnPhase1Policy(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetVpnPhase1Policy sets VpnPhase1Policy field to given value.

### HasVpnPhase1Policy

`func (o *PatchedVPNProfilePhase1PolicyAssignmentRequest) HasVpnPhase1Policy() bool`

HasVpnPhase1Policy returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedVPNProfilePhase1PolicyAssignmentRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedVPNProfilePhase1PolicyAssignmentRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedVPNProfilePhase1PolicyAssignmentRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedVPNProfilePhase1PolicyAssignmentRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedVPNProfilePhase1PolicyAssignmentRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedVPNProfilePhase1PolicyAssignmentRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedVPNProfilePhase1PolicyAssignmentRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedVPNProfilePhase1PolicyAssignmentRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


