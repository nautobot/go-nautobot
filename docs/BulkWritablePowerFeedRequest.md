# BulkWritablePowerFeedRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | Pointer to [**PowerFeedTypeChoices**](PowerFeedTypeChoices.md) |  | [optional] [default to POWERFEEDTYPECHOICES_PRIMARY]
**PowerPath** | Pointer to [**PowerPathEnum**](PowerPathEnum.md) |  | [optional] 
**Supply** | Pointer to [**SupplyEnum**](SupplyEnum.md) |  | [optional] [default to SUPPLYENUM_AC]
**Phase** | Pointer to [**PhaseEnum**](PhaseEnum.md) |  | [optional] [default to PHASEENUM_SINGLE_PHASE]
**BreakerPoleCount** | Pointer to [**NullableBreakerPoleCountEnum**](BreakerPoleCountEnum.md) |  | [optional] 
**Name** | **string** |  | 
**Voltage** | Pointer to **int32** |  | [optional] 
**Amperage** | Pointer to **int32** |  | [optional] 
**MaxUtilization** | Pointer to **int32** | Maximum permissible draw (percentage) | [optional] 
**BreakerPosition** | Pointer to **NullableInt32** | Starting circuit breaker position in panel | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Cable** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**PowerPanel** | [**BulkWritablePowerFeedRequestPowerPanel**](BulkWritablePowerFeedRequestPowerPanel.md) |  | 
**DestinationPanel** | Pointer to [**NullableBulkWritablePowerFeedRequestDestinationPanel**](BulkWritablePowerFeedRequestDestinationPanel.md) |  | [optional] 
**Rack** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Status** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewBulkWritablePowerFeedRequest

`func NewBulkWritablePowerFeedRequest(id string, name string, powerPanel BulkWritablePowerFeedRequestPowerPanel, status BulkWritableCableRequestStatus, ) *BulkWritablePowerFeedRequest`

NewBulkWritablePowerFeedRequest instantiates a new BulkWritablePowerFeedRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritablePowerFeedRequestWithDefaults

`func NewBulkWritablePowerFeedRequestWithDefaults() *BulkWritablePowerFeedRequest`

NewBulkWritablePowerFeedRequestWithDefaults instantiates a new BulkWritablePowerFeedRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritablePowerFeedRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritablePowerFeedRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritablePowerFeedRequest) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *BulkWritablePowerFeedRequest) GetType() PowerFeedTypeChoices`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BulkWritablePowerFeedRequest) GetTypeOk() (*PowerFeedTypeChoices, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BulkWritablePowerFeedRequest) SetType(v PowerFeedTypeChoices)`

SetType sets Type field to given value.

### HasType

`func (o *BulkWritablePowerFeedRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPowerPath

`func (o *BulkWritablePowerFeedRequest) GetPowerPath() PowerPathEnum`

GetPowerPath returns the PowerPath field if non-nil, zero value otherwise.

### GetPowerPathOk

`func (o *BulkWritablePowerFeedRequest) GetPowerPathOk() (*PowerPathEnum, bool)`

GetPowerPathOk returns a tuple with the PowerPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerPath

`func (o *BulkWritablePowerFeedRequest) SetPowerPath(v PowerPathEnum)`

SetPowerPath sets PowerPath field to given value.

### HasPowerPath

`func (o *BulkWritablePowerFeedRequest) HasPowerPath() bool`

HasPowerPath returns a boolean if a field has been set.

### GetSupply

`func (o *BulkWritablePowerFeedRequest) GetSupply() SupplyEnum`

GetSupply returns the Supply field if non-nil, zero value otherwise.

### GetSupplyOk

`func (o *BulkWritablePowerFeedRequest) GetSupplyOk() (*SupplyEnum, bool)`

GetSupplyOk returns a tuple with the Supply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupply

`func (o *BulkWritablePowerFeedRequest) SetSupply(v SupplyEnum)`

SetSupply sets Supply field to given value.

### HasSupply

`func (o *BulkWritablePowerFeedRequest) HasSupply() bool`

HasSupply returns a boolean if a field has been set.

### GetPhase

`func (o *BulkWritablePowerFeedRequest) GetPhase() PhaseEnum`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *BulkWritablePowerFeedRequest) GetPhaseOk() (*PhaseEnum, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *BulkWritablePowerFeedRequest) SetPhase(v PhaseEnum)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *BulkWritablePowerFeedRequest) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetBreakerPoleCount

`func (o *BulkWritablePowerFeedRequest) GetBreakerPoleCount() BreakerPoleCountEnum`

GetBreakerPoleCount returns the BreakerPoleCount field if non-nil, zero value otherwise.

### GetBreakerPoleCountOk

`func (o *BulkWritablePowerFeedRequest) GetBreakerPoleCountOk() (*BreakerPoleCountEnum, bool)`

GetBreakerPoleCountOk returns a tuple with the BreakerPoleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakerPoleCount

`func (o *BulkWritablePowerFeedRequest) SetBreakerPoleCount(v BreakerPoleCountEnum)`

SetBreakerPoleCount sets BreakerPoleCount field to given value.

### HasBreakerPoleCount

`func (o *BulkWritablePowerFeedRequest) HasBreakerPoleCount() bool`

HasBreakerPoleCount returns a boolean if a field has been set.

### SetBreakerPoleCountNil

`func (o *BulkWritablePowerFeedRequest) SetBreakerPoleCountNil(b bool)`

 SetBreakerPoleCountNil sets the value for BreakerPoleCount to be an explicit nil

### UnsetBreakerPoleCount
`func (o *BulkWritablePowerFeedRequest) UnsetBreakerPoleCount()`

UnsetBreakerPoleCount ensures that no value is present for BreakerPoleCount, not even an explicit nil
### GetName

`func (o *BulkWritablePowerFeedRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkWritablePowerFeedRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkWritablePowerFeedRequest) SetName(v string)`

SetName sets Name field to given value.


### GetVoltage

`func (o *BulkWritablePowerFeedRequest) GetVoltage() int32`

GetVoltage returns the Voltage field if non-nil, zero value otherwise.

### GetVoltageOk

`func (o *BulkWritablePowerFeedRequest) GetVoltageOk() (*int32, bool)`

GetVoltageOk returns a tuple with the Voltage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoltage

`func (o *BulkWritablePowerFeedRequest) SetVoltage(v int32)`

SetVoltage sets Voltage field to given value.

### HasVoltage

`func (o *BulkWritablePowerFeedRequest) HasVoltage() bool`

HasVoltage returns a boolean if a field has been set.

### GetAmperage

`func (o *BulkWritablePowerFeedRequest) GetAmperage() int32`

GetAmperage returns the Amperage field if non-nil, zero value otherwise.

### GetAmperageOk

`func (o *BulkWritablePowerFeedRequest) GetAmperageOk() (*int32, bool)`

GetAmperageOk returns a tuple with the Amperage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmperage

`func (o *BulkWritablePowerFeedRequest) SetAmperage(v int32)`

SetAmperage sets Amperage field to given value.

### HasAmperage

`func (o *BulkWritablePowerFeedRequest) HasAmperage() bool`

HasAmperage returns a boolean if a field has been set.

### GetMaxUtilization

`func (o *BulkWritablePowerFeedRequest) GetMaxUtilization() int32`

GetMaxUtilization returns the MaxUtilization field if non-nil, zero value otherwise.

### GetMaxUtilizationOk

`func (o *BulkWritablePowerFeedRequest) GetMaxUtilizationOk() (*int32, bool)`

GetMaxUtilizationOk returns a tuple with the MaxUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUtilization

`func (o *BulkWritablePowerFeedRequest) SetMaxUtilization(v int32)`

SetMaxUtilization sets MaxUtilization field to given value.

### HasMaxUtilization

`func (o *BulkWritablePowerFeedRequest) HasMaxUtilization() bool`

HasMaxUtilization returns a boolean if a field has been set.

### GetBreakerPosition

`func (o *BulkWritablePowerFeedRequest) GetBreakerPosition() int32`

GetBreakerPosition returns the BreakerPosition field if non-nil, zero value otherwise.

### GetBreakerPositionOk

`func (o *BulkWritablePowerFeedRequest) GetBreakerPositionOk() (*int32, bool)`

GetBreakerPositionOk returns a tuple with the BreakerPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakerPosition

`func (o *BulkWritablePowerFeedRequest) SetBreakerPosition(v int32)`

SetBreakerPosition sets BreakerPosition field to given value.

### HasBreakerPosition

`func (o *BulkWritablePowerFeedRequest) HasBreakerPosition() bool`

HasBreakerPosition returns a boolean if a field has been set.

### SetBreakerPositionNil

`func (o *BulkWritablePowerFeedRequest) SetBreakerPositionNil(b bool)`

 SetBreakerPositionNil sets the value for BreakerPosition to be an explicit nil

### UnsetBreakerPosition
`func (o *BulkWritablePowerFeedRequest) UnsetBreakerPosition()`

UnsetBreakerPosition ensures that no value is present for BreakerPosition, not even an explicit nil
### GetComments

`func (o *BulkWritablePowerFeedRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *BulkWritablePowerFeedRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *BulkWritablePowerFeedRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *BulkWritablePowerFeedRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetCable

`func (o *BulkWritablePowerFeedRequest) GetCable() BulkWritableCircuitRequestTenant`

GetCable returns the Cable field if non-nil, zero value otherwise.

### GetCableOk

`func (o *BulkWritablePowerFeedRequest) GetCableOk() (*BulkWritableCircuitRequestTenant, bool)`

GetCableOk returns a tuple with the Cable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCable

`func (o *BulkWritablePowerFeedRequest) SetCable(v BulkWritableCircuitRequestTenant)`

SetCable sets Cable field to given value.

### HasCable

`func (o *BulkWritablePowerFeedRequest) HasCable() bool`

HasCable returns a boolean if a field has been set.

### SetCableNil

`func (o *BulkWritablePowerFeedRequest) SetCableNil(b bool)`

 SetCableNil sets the value for Cable to be an explicit nil

### UnsetCable
`func (o *BulkWritablePowerFeedRequest) UnsetCable()`

UnsetCable ensures that no value is present for Cable, not even an explicit nil
### GetPowerPanel

`func (o *BulkWritablePowerFeedRequest) GetPowerPanel() BulkWritablePowerFeedRequestPowerPanel`

GetPowerPanel returns the PowerPanel field if non-nil, zero value otherwise.

### GetPowerPanelOk

`func (o *BulkWritablePowerFeedRequest) GetPowerPanelOk() (*BulkWritablePowerFeedRequestPowerPanel, bool)`

GetPowerPanelOk returns a tuple with the PowerPanel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerPanel

`func (o *BulkWritablePowerFeedRequest) SetPowerPanel(v BulkWritablePowerFeedRequestPowerPanel)`

SetPowerPanel sets PowerPanel field to given value.


### GetDestinationPanel

`func (o *BulkWritablePowerFeedRequest) GetDestinationPanel() BulkWritablePowerFeedRequestDestinationPanel`

GetDestinationPanel returns the DestinationPanel field if non-nil, zero value otherwise.

### GetDestinationPanelOk

`func (o *BulkWritablePowerFeedRequest) GetDestinationPanelOk() (*BulkWritablePowerFeedRequestDestinationPanel, bool)`

GetDestinationPanelOk returns a tuple with the DestinationPanel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPanel

`func (o *BulkWritablePowerFeedRequest) SetDestinationPanel(v BulkWritablePowerFeedRequestDestinationPanel)`

SetDestinationPanel sets DestinationPanel field to given value.

### HasDestinationPanel

`func (o *BulkWritablePowerFeedRequest) HasDestinationPanel() bool`

HasDestinationPanel returns a boolean if a field has been set.

### SetDestinationPanelNil

`func (o *BulkWritablePowerFeedRequest) SetDestinationPanelNil(b bool)`

 SetDestinationPanelNil sets the value for DestinationPanel to be an explicit nil

### UnsetDestinationPanel
`func (o *BulkWritablePowerFeedRequest) UnsetDestinationPanel()`

UnsetDestinationPanel ensures that no value is present for DestinationPanel, not even an explicit nil
### GetRack

`func (o *BulkWritablePowerFeedRequest) GetRack() BulkWritableCircuitRequestTenant`

GetRack returns the Rack field if non-nil, zero value otherwise.

### GetRackOk

`func (o *BulkWritablePowerFeedRequest) GetRackOk() (*BulkWritableCircuitRequestTenant, bool)`

GetRackOk returns a tuple with the Rack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRack

`func (o *BulkWritablePowerFeedRequest) SetRack(v BulkWritableCircuitRequestTenant)`

SetRack sets Rack field to given value.

### HasRack

`func (o *BulkWritablePowerFeedRequest) HasRack() bool`

HasRack returns a boolean if a field has been set.

### SetRackNil

`func (o *BulkWritablePowerFeedRequest) SetRackNil(b bool)`

 SetRackNil sets the value for Rack to be an explicit nil

### UnsetRack
`func (o *BulkWritablePowerFeedRequest) UnsetRack()`

UnsetRack ensures that no value is present for Rack, not even an explicit nil
### GetStatus

`func (o *BulkWritablePowerFeedRequest) GetStatus() BulkWritableCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkWritablePowerFeedRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkWritablePowerFeedRequest) SetStatus(v BulkWritableCableRequestStatus)`

SetStatus sets Status field to given value.


### GetCustomFields

`func (o *BulkWritablePowerFeedRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkWritablePowerFeedRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkWritablePowerFeedRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkWritablePowerFeedRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *BulkWritablePowerFeedRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BulkWritablePowerFeedRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BulkWritablePowerFeedRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BulkWritablePowerFeedRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *BulkWritablePowerFeedRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkWritablePowerFeedRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkWritablePowerFeedRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkWritablePowerFeedRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


