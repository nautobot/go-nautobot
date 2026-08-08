# VPNProfilePhase2PolicyAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Weight** | Pointer to **int32** | Higher weights appear later in the list | [optional] 
**VpnProfile** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 
**VpnPhase2Policy** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewVPNProfilePhase2PolicyAssignmentRequest

`func NewVPNProfilePhase2PolicyAssignmentRequest(vpnProfile ApprovalWorkflowStageResponseApprovalWorkflowStage, vpnPhase2Policy ApprovalWorkflowStageResponseApprovalWorkflowStage, ) *VPNProfilePhase2PolicyAssignmentRequest`

NewVPNProfilePhase2PolicyAssignmentRequest instantiates a new VPNProfilePhase2PolicyAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVPNProfilePhase2PolicyAssignmentRequestWithDefaults

`func NewVPNProfilePhase2PolicyAssignmentRequestWithDefaults() *VPNProfilePhase2PolicyAssignmentRequest`

NewVPNProfilePhase2PolicyAssignmentRequestWithDefaults instantiates a new VPNProfilePhase2PolicyAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VPNProfilePhase2PolicyAssignmentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VPNProfilePhase2PolicyAssignmentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VPNProfilePhase2PolicyAssignmentRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VPNProfilePhase2PolicyAssignmentRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetWeight

`func (o *VPNProfilePhase2PolicyAssignmentRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *VPNProfilePhase2PolicyAssignmentRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *VPNProfilePhase2PolicyAssignmentRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *VPNProfilePhase2PolicyAssignmentRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetVpnProfile

`func (o *VPNProfilePhase2PolicyAssignmentRequest) GetVpnProfile() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetVpnProfile returns the VpnProfile field if non-nil, zero value otherwise.

### GetVpnProfileOk

`func (o *VPNProfilePhase2PolicyAssignmentRequest) GetVpnProfileOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetVpnProfileOk returns a tuple with the VpnProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnProfile

`func (o *VPNProfilePhase2PolicyAssignmentRequest) SetVpnProfile(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetVpnProfile sets VpnProfile field to given value.


### GetVpnPhase2Policy

`func (o *VPNProfilePhase2PolicyAssignmentRequest) GetVpnPhase2Policy() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetVpnPhase2Policy returns the VpnPhase2Policy field if non-nil, zero value otherwise.

### GetVpnPhase2PolicyOk

`func (o *VPNProfilePhase2PolicyAssignmentRequest) GetVpnPhase2PolicyOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetVpnPhase2PolicyOk returns a tuple with the VpnPhase2Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnPhase2Policy

`func (o *VPNProfilePhase2PolicyAssignmentRequest) SetVpnPhase2Policy(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetVpnPhase2Policy sets VpnPhase2Policy field to given value.


### GetCustomFields

`func (o *VPNProfilePhase2PolicyAssignmentRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *VPNProfilePhase2PolicyAssignmentRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *VPNProfilePhase2PolicyAssignmentRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *VPNProfilePhase2PolicyAssignmentRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *VPNProfilePhase2PolicyAssignmentRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *VPNProfilePhase2PolicyAssignmentRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *VPNProfilePhase2PolicyAssignmentRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *VPNProfilePhase2PolicyAssignmentRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


