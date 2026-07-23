# VPNProfilePhase1PolicyAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Weight** | Pointer to **int32** | Higher weights appear later in the list | [optional] 
**VpnProfile** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 
**VpnPhase1Policy** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewVPNProfilePhase1PolicyAssignmentRequest

`func NewVPNProfilePhase1PolicyAssignmentRequest(vpnProfile ApprovalWorkflowStageResponseApprovalWorkflowStage, vpnPhase1Policy ApprovalWorkflowStageResponseApprovalWorkflowStage, ) *VPNProfilePhase1PolicyAssignmentRequest`

NewVPNProfilePhase1PolicyAssignmentRequest instantiates a new VPNProfilePhase1PolicyAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVPNProfilePhase1PolicyAssignmentRequestWithDefaults

`func NewVPNProfilePhase1PolicyAssignmentRequestWithDefaults() *VPNProfilePhase1PolicyAssignmentRequest`

NewVPNProfilePhase1PolicyAssignmentRequestWithDefaults instantiates a new VPNProfilePhase1PolicyAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VPNProfilePhase1PolicyAssignmentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VPNProfilePhase1PolicyAssignmentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VPNProfilePhase1PolicyAssignmentRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VPNProfilePhase1PolicyAssignmentRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetWeight

`func (o *VPNProfilePhase1PolicyAssignmentRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *VPNProfilePhase1PolicyAssignmentRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *VPNProfilePhase1PolicyAssignmentRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *VPNProfilePhase1PolicyAssignmentRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetVpnProfile

`func (o *VPNProfilePhase1PolicyAssignmentRequest) GetVpnProfile() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetVpnProfile returns the VpnProfile field if non-nil, zero value otherwise.

### GetVpnProfileOk

`func (o *VPNProfilePhase1PolicyAssignmentRequest) GetVpnProfileOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetVpnProfileOk returns a tuple with the VpnProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnProfile

`func (o *VPNProfilePhase1PolicyAssignmentRequest) SetVpnProfile(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetVpnProfile sets VpnProfile field to given value.


### GetVpnPhase1Policy

`func (o *VPNProfilePhase1PolicyAssignmentRequest) GetVpnPhase1Policy() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetVpnPhase1Policy returns the VpnPhase1Policy field if non-nil, zero value otherwise.

### GetVpnPhase1PolicyOk

`func (o *VPNProfilePhase1PolicyAssignmentRequest) GetVpnPhase1PolicyOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetVpnPhase1PolicyOk returns a tuple with the VpnPhase1Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnPhase1Policy

`func (o *VPNProfilePhase1PolicyAssignmentRequest) SetVpnPhase1Policy(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetVpnPhase1Policy sets VpnPhase1Policy field to given value.


### GetCustomFields

`func (o *VPNProfilePhase1PolicyAssignmentRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *VPNProfilePhase1PolicyAssignmentRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *VPNProfilePhase1PolicyAssignmentRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *VPNProfilePhase1PolicyAssignmentRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *VPNProfilePhase1PolicyAssignmentRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *VPNProfilePhase1PolicyAssignmentRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *VPNProfilePhase1PolicyAssignmentRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *VPNProfilePhase1PolicyAssignmentRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


