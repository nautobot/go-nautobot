# PatchedWritablePowerFeedRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**PowerFeedTypeChoices**](PowerFeedTypeChoices.md) |  | [optional] 
**PowerPath** | Pointer to [**PatchedWritablePowerFeedRequestPowerPath**](PatchedWritablePowerFeedRequestPowerPath.md) |  | [optional] 
**Supply** | Pointer to [**SupplyEnum**](SupplyEnum.md) |  | [optional] 
**Phase** | Pointer to [**PhaseEnum**](PhaseEnum.md) |  | [optional] 
**Voltage** | Pointer to **int32** |  | [optional] 
**Amperage** | Pointer to **int32** |  | [optional] 
**MaxUtilization** | Pointer to **int32** | Maximum permissible draw (percentage) | [optional] 
**BreakerPosition** | Pointer to **NullableInt32** | Starting circuit breaker position in panel | [optional] 
**BreakerPoleCount** | Pointer to [**NullableBreakerPoleCountEnum**](BreakerPoleCountEnum.md) | Number of breaker poles | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Cable** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**PowerPanel** | Pointer to [**BulkWritablePowerFeedRequestPowerPanel**](BulkWritablePowerFeedRequestPowerPanel.md) |  | [optional] 
**DestinationPanel** | Pointer to [**NullableBulkWritablePowerFeedRequestDestinationPanel**](BulkWritablePowerFeedRequestDestinationPanel.md) |  | [optional] 
**Rack** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**Status** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewPatchedWritablePowerFeedRequest

`func NewPatchedWritablePowerFeedRequest() *PatchedWritablePowerFeedRequest`

NewPatchedWritablePowerFeedRequest instantiates a new PatchedWritablePowerFeedRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritablePowerFeedRequestWithDefaults

`func NewPatchedWritablePowerFeedRequestWithDefaults() *PatchedWritablePowerFeedRequest`

NewPatchedWritablePowerFeedRequestWithDefaults instantiates a new PatchedWritablePowerFeedRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedWritablePowerFeedRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedWritablePowerFeedRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedWritablePowerFeedRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedWritablePowerFeedRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PatchedWritablePowerFeedRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedWritablePowerFeedRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedWritablePowerFeedRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedWritablePowerFeedRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *PatchedWritablePowerFeedRequest) GetType() PowerFeedTypeChoices`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedWritablePowerFeedRequest) GetTypeOk() (*PowerFeedTypeChoices, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedWritablePowerFeedRequest) SetType(v PowerFeedTypeChoices)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedWritablePowerFeedRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPowerPath

`func (o *PatchedWritablePowerFeedRequest) GetPowerPath() PatchedWritablePowerFeedRequestPowerPath`

GetPowerPath returns the PowerPath field if non-nil, zero value otherwise.

### GetPowerPathOk

`func (o *PatchedWritablePowerFeedRequest) GetPowerPathOk() (*PatchedWritablePowerFeedRequestPowerPath, bool)`

GetPowerPathOk returns a tuple with the PowerPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerPath

`func (o *PatchedWritablePowerFeedRequest) SetPowerPath(v PatchedWritablePowerFeedRequestPowerPath)`

SetPowerPath sets PowerPath field to given value.

### HasPowerPath

`func (o *PatchedWritablePowerFeedRequest) HasPowerPath() bool`

HasPowerPath returns a boolean if a field has been set.

### GetSupply

`func (o *PatchedWritablePowerFeedRequest) GetSupply() SupplyEnum`

GetSupply returns the Supply field if non-nil, zero value otherwise.

### GetSupplyOk

`func (o *PatchedWritablePowerFeedRequest) GetSupplyOk() (*SupplyEnum, bool)`

GetSupplyOk returns a tuple with the Supply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupply

`func (o *PatchedWritablePowerFeedRequest) SetSupply(v SupplyEnum)`

SetSupply sets Supply field to given value.

### HasSupply

`func (o *PatchedWritablePowerFeedRequest) HasSupply() bool`

HasSupply returns a boolean if a field has been set.

### GetPhase

`func (o *PatchedWritablePowerFeedRequest) GetPhase() PhaseEnum`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *PatchedWritablePowerFeedRequest) GetPhaseOk() (*PhaseEnum, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *PatchedWritablePowerFeedRequest) SetPhase(v PhaseEnum)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *PatchedWritablePowerFeedRequest) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetVoltage

`func (o *PatchedWritablePowerFeedRequest) GetVoltage() int32`

GetVoltage returns the Voltage field if non-nil, zero value otherwise.

### GetVoltageOk

`func (o *PatchedWritablePowerFeedRequest) GetVoltageOk() (*int32, bool)`

GetVoltageOk returns a tuple with the Voltage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoltage

`func (o *PatchedWritablePowerFeedRequest) SetVoltage(v int32)`

SetVoltage sets Voltage field to given value.

### HasVoltage

`func (o *PatchedWritablePowerFeedRequest) HasVoltage() bool`

HasVoltage returns a boolean if a field has been set.

### GetAmperage

`func (o *PatchedWritablePowerFeedRequest) GetAmperage() int32`

GetAmperage returns the Amperage field if non-nil, zero value otherwise.

### GetAmperageOk

`func (o *PatchedWritablePowerFeedRequest) GetAmperageOk() (*int32, bool)`

GetAmperageOk returns a tuple with the Amperage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmperage

`func (o *PatchedWritablePowerFeedRequest) SetAmperage(v int32)`

SetAmperage sets Amperage field to given value.

### HasAmperage

`func (o *PatchedWritablePowerFeedRequest) HasAmperage() bool`

HasAmperage returns a boolean if a field has been set.

### GetMaxUtilization

`func (o *PatchedWritablePowerFeedRequest) GetMaxUtilization() int32`

GetMaxUtilization returns the MaxUtilization field if non-nil, zero value otherwise.

### GetMaxUtilizationOk

`func (o *PatchedWritablePowerFeedRequest) GetMaxUtilizationOk() (*int32, bool)`

GetMaxUtilizationOk returns a tuple with the MaxUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUtilization

`func (o *PatchedWritablePowerFeedRequest) SetMaxUtilization(v int32)`

SetMaxUtilization sets MaxUtilization field to given value.

### HasMaxUtilization

`func (o *PatchedWritablePowerFeedRequest) HasMaxUtilization() bool`

HasMaxUtilization returns a boolean if a field has been set.

### GetBreakerPosition

`func (o *PatchedWritablePowerFeedRequest) GetBreakerPosition() int32`

GetBreakerPosition returns the BreakerPosition field if non-nil, zero value otherwise.

### GetBreakerPositionOk

`func (o *PatchedWritablePowerFeedRequest) GetBreakerPositionOk() (*int32, bool)`

GetBreakerPositionOk returns a tuple with the BreakerPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakerPosition

`func (o *PatchedWritablePowerFeedRequest) SetBreakerPosition(v int32)`

SetBreakerPosition sets BreakerPosition field to given value.

### HasBreakerPosition

`func (o *PatchedWritablePowerFeedRequest) HasBreakerPosition() bool`

HasBreakerPosition returns a boolean if a field has been set.

### SetBreakerPositionNil

`func (o *PatchedWritablePowerFeedRequest) SetBreakerPositionNil(b bool)`

 SetBreakerPositionNil sets the value for BreakerPosition to be an explicit nil

### UnsetBreakerPosition
`func (o *PatchedWritablePowerFeedRequest) UnsetBreakerPosition()`

UnsetBreakerPosition ensures that no value is present for BreakerPosition, not even an explicit nil
### GetBreakerPoleCount

`func (o *PatchedWritablePowerFeedRequest) GetBreakerPoleCount() BreakerPoleCountEnum`

GetBreakerPoleCount returns the BreakerPoleCount field if non-nil, zero value otherwise.

### GetBreakerPoleCountOk

`func (o *PatchedWritablePowerFeedRequest) GetBreakerPoleCountOk() (*BreakerPoleCountEnum, bool)`

GetBreakerPoleCountOk returns a tuple with the BreakerPoleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakerPoleCount

`func (o *PatchedWritablePowerFeedRequest) SetBreakerPoleCount(v BreakerPoleCountEnum)`

SetBreakerPoleCount sets BreakerPoleCount field to given value.

### HasBreakerPoleCount

`func (o *PatchedWritablePowerFeedRequest) HasBreakerPoleCount() bool`

HasBreakerPoleCount returns a boolean if a field has been set.

### SetBreakerPoleCountNil

`func (o *PatchedWritablePowerFeedRequest) SetBreakerPoleCountNil(b bool)`

 SetBreakerPoleCountNil sets the value for BreakerPoleCount to be an explicit nil

### UnsetBreakerPoleCount
`func (o *PatchedWritablePowerFeedRequest) UnsetBreakerPoleCount()`

UnsetBreakerPoleCount ensures that no value is present for BreakerPoleCount, not even an explicit nil
### GetComments

`func (o *PatchedWritablePowerFeedRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedWritablePowerFeedRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedWritablePowerFeedRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedWritablePowerFeedRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetCable

`func (o *PatchedWritablePowerFeedRequest) GetCable() ApprovalWorkflowUser`

GetCable returns the Cable field if non-nil, zero value otherwise.

### GetCableOk

`func (o *PatchedWritablePowerFeedRequest) GetCableOk() (*ApprovalWorkflowUser, bool)`

GetCableOk returns a tuple with the Cable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCable

`func (o *PatchedWritablePowerFeedRequest) SetCable(v ApprovalWorkflowUser)`

SetCable sets Cable field to given value.

### HasCable

`func (o *PatchedWritablePowerFeedRequest) HasCable() bool`

HasCable returns a boolean if a field has been set.

### SetCableNil

`func (o *PatchedWritablePowerFeedRequest) SetCableNil(b bool)`

 SetCableNil sets the value for Cable to be an explicit nil

### UnsetCable
`func (o *PatchedWritablePowerFeedRequest) UnsetCable()`

UnsetCable ensures that no value is present for Cable, not even an explicit nil
### GetPowerPanel

`func (o *PatchedWritablePowerFeedRequest) GetPowerPanel() BulkWritablePowerFeedRequestPowerPanel`

GetPowerPanel returns the PowerPanel field if non-nil, zero value otherwise.

### GetPowerPanelOk

`func (o *PatchedWritablePowerFeedRequest) GetPowerPanelOk() (*BulkWritablePowerFeedRequestPowerPanel, bool)`

GetPowerPanelOk returns a tuple with the PowerPanel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerPanel

`func (o *PatchedWritablePowerFeedRequest) SetPowerPanel(v BulkWritablePowerFeedRequestPowerPanel)`

SetPowerPanel sets PowerPanel field to given value.

### HasPowerPanel

`func (o *PatchedWritablePowerFeedRequest) HasPowerPanel() bool`

HasPowerPanel returns a boolean if a field has been set.

### GetDestinationPanel

`func (o *PatchedWritablePowerFeedRequest) GetDestinationPanel() BulkWritablePowerFeedRequestDestinationPanel`

GetDestinationPanel returns the DestinationPanel field if non-nil, zero value otherwise.

### GetDestinationPanelOk

`func (o *PatchedWritablePowerFeedRequest) GetDestinationPanelOk() (*BulkWritablePowerFeedRequestDestinationPanel, bool)`

GetDestinationPanelOk returns a tuple with the DestinationPanel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPanel

`func (o *PatchedWritablePowerFeedRequest) SetDestinationPanel(v BulkWritablePowerFeedRequestDestinationPanel)`

SetDestinationPanel sets DestinationPanel field to given value.

### HasDestinationPanel

`func (o *PatchedWritablePowerFeedRequest) HasDestinationPanel() bool`

HasDestinationPanel returns a boolean if a field has been set.

### SetDestinationPanelNil

`func (o *PatchedWritablePowerFeedRequest) SetDestinationPanelNil(b bool)`

 SetDestinationPanelNil sets the value for DestinationPanel to be an explicit nil

### UnsetDestinationPanel
`func (o *PatchedWritablePowerFeedRequest) UnsetDestinationPanel()`

UnsetDestinationPanel ensures that no value is present for DestinationPanel, not even an explicit nil
### GetRack

`func (o *PatchedWritablePowerFeedRequest) GetRack() ApprovalWorkflowUser`

GetRack returns the Rack field if non-nil, zero value otherwise.

### GetRackOk

`func (o *PatchedWritablePowerFeedRequest) GetRackOk() (*ApprovalWorkflowUser, bool)`

GetRackOk returns a tuple with the Rack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRack

`func (o *PatchedWritablePowerFeedRequest) SetRack(v ApprovalWorkflowUser)`

SetRack sets Rack field to given value.

### HasRack

`func (o *PatchedWritablePowerFeedRequest) HasRack() bool`

HasRack returns a boolean if a field has been set.

### SetRackNil

`func (o *PatchedWritablePowerFeedRequest) SetRackNil(b bool)`

 SetRackNil sets the value for Rack to be an explicit nil

### UnsetRack
`func (o *PatchedWritablePowerFeedRequest) UnsetRack()`

UnsetRack ensures that no value is present for Rack, not even an explicit nil
### GetStatus

`func (o *PatchedWritablePowerFeedRequest) GetStatus() BulkWritableCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedWritablePowerFeedRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedWritablePowerFeedRequest) SetStatus(v BulkWritableCableRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedWritablePowerFeedRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedWritablePowerFeedRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritablePowerFeedRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritablePowerFeedRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritablePowerFeedRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedWritablePowerFeedRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedWritablePowerFeedRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedWritablePowerFeedRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedWritablePowerFeedRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *PatchedWritablePowerFeedRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedWritablePowerFeedRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedWritablePowerFeedRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedWritablePowerFeedRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


