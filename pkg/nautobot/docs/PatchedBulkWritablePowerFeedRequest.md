# PatchedBulkWritablePowerFeedRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | Pointer to [**PowerFeedTypeChoices**](PowerFeedTypeChoices.md) |  | [optional] [default to "{\"value\":\"primary\",\"label\":\"Primary\"}"]
**Supply** | Pointer to [**SupplyEnum**](SupplyEnum.md) |  | [optional] [default to "{\"value\":\"ac\",\"label\":\"AC\"}"]
**Phase** | Pointer to [**PhaseEnum**](PhaseEnum.md) |  | [optional] [default to "{\"value\":\"single-phase\",\"label\":\"Single phase\"}"]
**Name** | Pointer to **string** |  | [optional] 
**Voltage** | Pointer to **int32** |  | [optional] 
**Amperage** | Pointer to **int32** |  | [optional] 
**MaxUtilization** | Pointer to **int32** | Maximum permissible draw (percentage) | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Cable** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**PowerPanel** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**Rack** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Status** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewPatchedBulkWritablePowerFeedRequest

`func NewPatchedBulkWritablePowerFeedRequest(id string, ) *PatchedBulkWritablePowerFeedRequest`

NewPatchedBulkWritablePowerFeedRequest instantiates a new PatchedBulkWritablePowerFeedRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritablePowerFeedRequestWithDefaults

`func NewPatchedBulkWritablePowerFeedRequestWithDefaults() *PatchedBulkWritablePowerFeedRequest`

NewPatchedBulkWritablePowerFeedRequestWithDefaults instantiates a new PatchedBulkWritablePowerFeedRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritablePowerFeedRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritablePowerFeedRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritablePowerFeedRequest) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *PatchedBulkWritablePowerFeedRequest) GetType() PowerFeedTypeChoices`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedBulkWritablePowerFeedRequest) GetTypeOk() (*PowerFeedTypeChoices, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedBulkWritablePowerFeedRequest) SetType(v PowerFeedTypeChoices)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedBulkWritablePowerFeedRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSupply

`func (o *PatchedBulkWritablePowerFeedRequest) GetSupply() SupplyEnum`

GetSupply returns the Supply field if non-nil, zero value otherwise.

### GetSupplyOk

`func (o *PatchedBulkWritablePowerFeedRequest) GetSupplyOk() (*SupplyEnum, bool)`

GetSupplyOk returns a tuple with the Supply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupply

`func (o *PatchedBulkWritablePowerFeedRequest) SetSupply(v SupplyEnum)`

SetSupply sets Supply field to given value.

### HasSupply

`func (o *PatchedBulkWritablePowerFeedRequest) HasSupply() bool`

HasSupply returns a boolean if a field has been set.

### GetPhase

`func (o *PatchedBulkWritablePowerFeedRequest) GetPhase() PhaseEnum`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *PatchedBulkWritablePowerFeedRequest) GetPhaseOk() (*PhaseEnum, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *PatchedBulkWritablePowerFeedRequest) SetPhase(v PhaseEnum)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *PatchedBulkWritablePowerFeedRequest) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetName

`func (o *PatchedBulkWritablePowerFeedRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkWritablePowerFeedRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkWritablePowerFeedRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkWritablePowerFeedRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVoltage

`func (o *PatchedBulkWritablePowerFeedRequest) GetVoltage() int32`

GetVoltage returns the Voltage field if non-nil, zero value otherwise.

### GetVoltageOk

`func (o *PatchedBulkWritablePowerFeedRequest) GetVoltageOk() (*int32, bool)`

GetVoltageOk returns a tuple with the Voltage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoltage

`func (o *PatchedBulkWritablePowerFeedRequest) SetVoltage(v int32)`

SetVoltage sets Voltage field to given value.

### HasVoltage

`func (o *PatchedBulkWritablePowerFeedRequest) HasVoltage() bool`

HasVoltage returns a boolean if a field has been set.

### GetAmperage

`func (o *PatchedBulkWritablePowerFeedRequest) GetAmperage() int32`

GetAmperage returns the Amperage field if non-nil, zero value otherwise.

### GetAmperageOk

`func (o *PatchedBulkWritablePowerFeedRequest) GetAmperageOk() (*int32, bool)`

GetAmperageOk returns a tuple with the Amperage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmperage

`func (o *PatchedBulkWritablePowerFeedRequest) SetAmperage(v int32)`

SetAmperage sets Amperage field to given value.

### HasAmperage

`func (o *PatchedBulkWritablePowerFeedRequest) HasAmperage() bool`

HasAmperage returns a boolean if a field has been set.

### GetMaxUtilization

`func (o *PatchedBulkWritablePowerFeedRequest) GetMaxUtilization() int32`

GetMaxUtilization returns the MaxUtilization field if non-nil, zero value otherwise.

### GetMaxUtilizationOk

`func (o *PatchedBulkWritablePowerFeedRequest) GetMaxUtilizationOk() (*int32, bool)`

GetMaxUtilizationOk returns a tuple with the MaxUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUtilization

`func (o *PatchedBulkWritablePowerFeedRequest) SetMaxUtilization(v int32)`

SetMaxUtilization sets MaxUtilization field to given value.

### HasMaxUtilization

`func (o *PatchedBulkWritablePowerFeedRequest) HasMaxUtilization() bool`

HasMaxUtilization returns a boolean if a field has been set.

### GetComments

`func (o *PatchedBulkWritablePowerFeedRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedBulkWritablePowerFeedRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedBulkWritablePowerFeedRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedBulkWritablePowerFeedRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetCable

`func (o *PatchedBulkWritablePowerFeedRequest) GetCable() BulkWritableCircuitRequestTenant`

GetCable returns the Cable field if non-nil, zero value otherwise.

### GetCableOk

`func (o *PatchedBulkWritablePowerFeedRequest) GetCableOk() (*BulkWritableCircuitRequestTenant, bool)`

GetCableOk returns a tuple with the Cable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCable

`func (o *PatchedBulkWritablePowerFeedRequest) SetCable(v BulkWritableCircuitRequestTenant)`

SetCable sets Cable field to given value.

### HasCable

`func (o *PatchedBulkWritablePowerFeedRequest) HasCable() bool`

HasCable returns a boolean if a field has been set.

### SetCableNil

`func (o *PatchedBulkWritablePowerFeedRequest) SetCableNil(b bool)`

 SetCableNil sets the value for Cable to be an explicit nil

### UnsetCable
`func (o *PatchedBulkWritablePowerFeedRequest) UnsetCable()`

UnsetCable ensures that no value is present for Cable, not even an explicit nil
### GetPowerPanel

`func (o *PatchedBulkWritablePowerFeedRequest) GetPowerPanel() BulkWritableCableRequestStatus`

GetPowerPanel returns the PowerPanel field if non-nil, zero value otherwise.

### GetPowerPanelOk

`func (o *PatchedBulkWritablePowerFeedRequest) GetPowerPanelOk() (*BulkWritableCableRequestStatus, bool)`

GetPowerPanelOk returns a tuple with the PowerPanel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerPanel

`func (o *PatchedBulkWritablePowerFeedRequest) SetPowerPanel(v BulkWritableCableRequestStatus)`

SetPowerPanel sets PowerPanel field to given value.

### HasPowerPanel

`func (o *PatchedBulkWritablePowerFeedRequest) HasPowerPanel() bool`

HasPowerPanel returns a boolean if a field has been set.

### GetRack

`func (o *PatchedBulkWritablePowerFeedRequest) GetRack() BulkWritableCircuitRequestTenant`

GetRack returns the Rack field if non-nil, zero value otherwise.

### GetRackOk

`func (o *PatchedBulkWritablePowerFeedRequest) GetRackOk() (*BulkWritableCircuitRequestTenant, bool)`

GetRackOk returns a tuple with the Rack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRack

`func (o *PatchedBulkWritablePowerFeedRequest) SetRack(v BulkWritableCircuitRequestTenant)`

SetRack sets Rack field to given value.

### HasRack

`func (o *PatchedBulkWritablePowerFeedRequest) HasRack() bool`

HasRack returns a boolean if a field has been set.

### SetRackNil

`func (o *PatchedBulkWritablePowerFeedRequest) SetRackNil(b bool)`

 SetRackNil sets the value for Rack to be an explicit nil

### UnsetRack
`func (o *PatchedBulkWritablePowerFeedRequest) UnsetRack()`

UnsetRack ensures that no value is present for Rack, not even an explicit nil
### GetStatus

`func (o *PatchedBulkWritablePowerFeedRequest) GetStatus() BulkWritableCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedBulkWritablePowerFeedRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedBulkWritablePowerFeedRequest) SetStatus(v BulkWritableCableRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedBulkWritablePowerFeedRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *PatchedBulkWritablePowerFeedRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkWritablePowerFeedRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkWritablePowerFeedRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkWritablePowerFeedRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedBulkWritablePowerFeedRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkWritablePowerFeedRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkWritablePowerFeedRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkWritablePowerFeedRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedBulkWritablePowerFeedRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedBulkWritablePowerFeedRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedBulkWritablePowerFeedRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedBulkWritablePowerFeedRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


