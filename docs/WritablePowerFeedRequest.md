# WritablePowerFeedRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Type** | Pointer to [**PowerFeedTypeChoices**](PowerFeedTypeChoices.md) |  | [optional] 
**Supply** | Pointer to [**SupplyEnum**](SupplyEnum.md) |  | [optional] 
**Phase** | Pointer to [**PhaseEnum**](PhaseEnum.md) |  | [optional] 
**Voltage** | Pointer to **int32** |  | [optional] 
**Amperage** | Pointer to **int32** |  | [optional] 
**MaxUtilization** | Pointer to **int32** | Maximum permissible draw (percentage) | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Cable** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**PowerPanel** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Rack** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Status** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewWritablePowerFeedRequest

`func NewWritablePowerFeedRequest(name string, powerPanel BulkWritableCableRequestStatus, status BulkWritableCableRequestStatus, ) *WritablePowerFeedRequest`

NewWritablePowerFeedRequest instantiates a new WritablePowerFeedRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWritablePowerFeedRequestWithDefaults

`func NewWritablePowerFeedRequestWithDefaults() *WritablePowerFeedRequest`

NewWritablePowerFeedRequestWithDefaults instantiates a new WritablePowerFeedRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WritablePowerFeedRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WritablePowerFeedRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WritablePowerFeedRequest) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *WritablePowerFeedRequest) GetType() PowerFeedTypeChoices`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WritablePowerFeedRequest) GetTypeOk() (*PowerFeedTypeChoices, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WritablePowerFeedRequest) SetType(v PowerFeedTypeChoices)`

SetType sets Type field to given value.

### HasType

`func (o *WritablePowerFeedRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSupply

`func (o *WritablePowerFeedRequest) GetSupply() SupplyEnum`

GetSupply returns the Supply field if non-nil, zero value otherwise.

### GetSupplyOk

`func (o *WritablePowerFeedRequest) GetSupplyOk() (*SupplyEnum, bool)`

GetSupplyOk returns a tuple with the Supply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupply

`func (o *WritablePowerFeedRequest) SetSupply(v SupplyEnum)`

SetSupply sets Supply field to given value.

### HasSupply

`func (o *WritablePowerFeedRequest) HasSupply() bool`

HasSupply returns a boolean if a field has been set.

### GetPhase

`func (o *WritablePowerFeedRequest) GetPhase() PhaseEnum`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *WritablePowerFeedRequest) GetPhaseOk() (*PhaseEnum, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *WritablePowerFeedRequest) SetPhase(v PhaseEnum)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *WritablePowerFeedRequest) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetVoltage

`func (o *WritablePowerFeedRequest) GetVoltage() int32`

GetVoltage returns the Voltage field if non-nil, zero value otherwise.

### GetVoltageOk

`func (o *WritablePowerFeedRequest) GetVoltageOk() (*int32, bool)`

GetVoltageOk returns a tuple with the Voltage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoltage

`func (o *WritablePowerFeedRequest) SetVoltage(v int32)`

SetVoltage sets Voltage field to given value.

### HasVoltage

`func (o *WritablePowerFeedRequest) HasVoltage() bool`

HasVoltage returns a boolean if a field has been set.

### GetAmperage

`func (o *WritablePowerFeedRequest) GetAmperage() int32`

GetAmperage returns the Amperage field if non-nil, zero value otherwise.

### GetAmperageOk

`func (o *WritablePowerFeedRequest) GetAmperageOk() (*int32, bool)`

GetAmperageOk returns a tuple with the Amperage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmperage

`func (o *WritablePowerFeedRequest) SetAmperage(v int32)`

SetAmperage sets Amperage field to given value.

### HasAmperage

`func (o *WritablePowerFeedRequest) HasAmperage() bool`

HasAmperage returns a boolean if a field has been set.

### GetMaxUtilization

`func (o *WritablePowerFeedRequest) GetMaxUtilization() int32`

GetMaxUtilization returns the MaxUtilization field if non-nil, zero value otherwise.

### GetMaxUtilizationOk

`func (o *WritablePowerFeedRequest) GetMaxUtilizationOk() (*int32, bool)`

GetMaxUtilizationOk returns a tuple with the MaxUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUtilization

`func (o *WritablePowerFeedRequest) SetMaxUtilization(v int32)`

SetMaxUtilization sets MaxUtilization field to given value.

### HasMaxUtilization

`func (o *WritablePowerFeedRequest) HasMaxUtilization() bool`

HasMaxUtilization returns a boolean if a field has been set.

### GetComments

`func (o *WritablePowerFeedRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *WritablePowerFeedRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *WritablePowerFeedRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *WritablePowerFeedRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetCable

`func (o *WritablePowerFeedRequest) GetCable() BulkWritableCircuitRequestTenant`

GetCable returns the Cable field if non-nil, zero value otherwise.

### GetCableOk

`func (o *WritablePowerFeedRequest) GetCableOk() (*BulkWritableCircuitRequestTenant, bool)`

GetCableOk returns a tuple with the Cable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCable

`func (o *WritablePowerFeedRequest) SetCable(v BulkWritableCircuitRequestTenant)`

SetCable sets Cable field to given value.

### HasCable

`func (o *WritablePowerFeedRequest) HasCable() bool`

HasCable returns a boolean if a field has been set.

### SetCableNil

`func (o *WritablePowerFeedRequest) SetCableNil(b bool)`

 SetCableNil sets the value for Cable to be an explicit nil

### UnsetCable
`func (o *WritablePowerFeedRequest) UnsetCable()`

UnsetCable ensures that no value is present for Cable, not even an explicit nil
### GetPowerPanel

`func (o *WritablePowerFeedRequest) GetPowerPanel() BulkWritableCableRequestStatus`

GetPowerPanel returns the PowerPanel field if non-nil, zero value otherwise.

### GetPowerPanelOk

`func (o *WritablePowerFeedRequest) GetPowerPanelOk() (*BulkWritableCableRequestStatus, bool)`

GetPowerPanelOk returns a tuple with the PowerPanel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerPanel

`func (o *WritablePowerFeedRequest) SetPowerPanel(v BulkWritableCableRequestStatus)`

SetPowerPanel sets PowerPanel field to given value.


### GetRack

`func (o *WritablePowerFeedRequest) GetRack() BulkWritableCircuitRequestTenant`

GetRack returns the Rack field if non-nil, zero value otherwise.

### GetRackOk

`func (o *WritablePowerFeedRequest) GetRackOk() (*BulkWritableCircuitRequestTenant, bool)`

GetRackOk returns a tuple with the Rack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRack

`func (o *WritablePowerFeedRequest) SetRack(v BulkWritableCircuitRequestTenant)`

SetRack sets Rack field to given value.

### HasRack

`func (o *WritablePowerFeedRequest) HasRack() bool`

HasRack returns a boolean if a field has been set.

### SetRackNil

`func (o *WritablePowerFeedRequest) SetRackNil(b bool)`

 SetRackNil sets the value for Rack to be an explicit nil

### UnsetRack
`func (o *WritablePowerFeedRequest) UnsetRack()`

UnsetRack ensures that no value is present for Rack, not even an explicit nil
### GetStatus

`func (o *WritablePowerFeedRequest) GetStatus() BulkWritableCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WritablePowerFeedRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WritablePowerFeedRequest) SetStatus(v BulkWritableCableRequestStatus)`

SetStatus sets Status field to given value.


### GetTags

`func (o *WritablePowerFeedRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WritablePowerFeedRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WritablePowerFeedRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WritablePowerFeedRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *WritablePowerFeedRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *WritablePowerFeedRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *WritablePowerFeedRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *WritablePowerFeedRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *WritablePowerFeedRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *WritablePowerFeedRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *WritablePowerFeedRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *WritablePowerFeedRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


