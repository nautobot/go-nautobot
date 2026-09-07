# PatchedCableToCableTerminationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**CableEnd** | Pointer to [**CableEndEnum**](CableEndEnum.md) |  | [optional] 
**Connector** | Pointer to **int32** | The connector number on this cable end. Always 1 for standard cables. | [optional] [default to 1]
**Cable** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**CircuitTermination** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**ConsolePort** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**ConsoleServerPort** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**FrontPort** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**Interface** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**PowerFeed** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**PowerOutlet** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**PowerPort** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**RearPort** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 

## Methods

### NewPatchedCableToCableTerminationRequest

`func NewPatchedCableToCableTerminationRequest() *PatchedCableToCableTerminationRequest`

NewPatchedCableToCableTerminationRequest instantiates a new PatchedCableToCableTerminationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedCableToCableTerminationRequestWithDefaults

`func NewPatchedCableToCableTerminationRequestWithDefaults() *PatchedCableToCableTerminationRequest`

NewPatchedCableToCableTerminationRequestWithDefaults instantiates a new PatchedCableToCableTerminationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedCableToCableTerminationRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedCableToCableTerminationRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedCableToCableTerminationRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedCableToCableTerminationRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCableEnd

`func (o *PatchedCableToCableTerminationRequest) GetCableEnd() CableEndEnum`

GetCableEnd returns the CableEnd field if non-nil, zero value otherwise.

### GetCableEndOk

`func (o *PatchedCableToCableTerminationRequest) GetCableEndOk() (*CableEndEnum, bool)`

GetCableEndOk returns a tuple with the CableEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCableEnd

`func (o *PatchedCableToCableTerminationRequest) SetCableEnd(v CableEndEnum)`

SetCableEnd sets CableEnd field to given value.

### HasCableEnd

`func (o *PatchedCableToCableTerminationRequest) HasCableEnd() bool`

HasCableEnd returns a boolean if a field has been set.

### GetConnector

`func (o *PatchedCableToCableTerminationRequest) GetConnector() int32`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *PatchedCableToCableTerminationRequest) GetConnectorOk() (*int32, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *PatchedCableToCableTerminationRequest) SetConnector(v int32)`

SetConnector sets Connector field to given value.

### HasConnector

`func (o *PatchedCableToCableTerminationRequest) HasConnector() bool`

HasConnector returns a boolean if a field has been set.

### GetCable

`func (o *PatchedCableToCableTerminationRequest) GetCable() BulkWritableCableRequestStatus`

GetCable returns the Cable field if non-nil, zero value otherwise.

### GetCableOk

`func (o *PatchedCableToCableTerminationRequest) GetCableOk() (*BulkWritableCableRequestStatus, bool)`

GetCableOk returns a tuple with the Cable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCable

`func (o *PatchedCableToCableTerminationRequest) SetCable(v BulkWritableCableRequestStatus)`

SetCable sets Cable field to given value.

### HasCable

`func (o *PatchedCableToCableTerminationRequest) HasCable() bool`

HasCable returns a boolean if a field has been set.

### GetCircuitTermination

`func (o *PatchedCableToCableTerminationRequest) GetCircuitTermination() ApprovalWorkflowUser`

GetCircuitTermination returns the CircuitTermination field if non-nil, zero value otherwise.

### GetCircuitTerminationOk

`func (o *PatchedCableToCableTerminationRequest) GetCircuitTerminationOk() (*ApprovalWorkflowUser, bool)`

GetCircuitTerminationOk returns a tuple with the CircuitTermination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitTermination

`func (o *PatchedCableToCableTerminationRequest) SetCircuitTermination(v ApprovalWorkflowUser)`

SetCircuitTermination sets CircuitTermination field to given value.

### HasCircuitTermination

`func (o *PatchedCableToCableTerminationRequest) HasCircuitTermination() bool`

HasCircuitTermination returns a boolean if a field has been set.

### SetCircuitTerminationNil

`func (o *PatchedCableToCableTerminationRequest) SetCircuitTerminationNil(b bool)`

 SetCircuitTerminationNil sets the value for CircuitTermination to be an explicit nil

### UnsetCircuitTermination
`func (o *PatchedCableToCableTerminationRequest) UnsetCircuitTermination()`

UnsetCircuitTermination ensures that no value is present for CircuitTermination, not even an explicit nil
### GetConsolePort

`func (o *PatchedCableToCableTerminationRequest) GetConsolePort() ApprovalWorkflowUser`

GetConsolePort returns the ConsolePort field if non-nil, zero value otherwise.

### GetConsolePortOk

`func (o *PatchedCableToCableTerminationRequest) GetConsolePortOk() (*ApprovalWorkflowUser, bool)`

GetConsolePortOk returns a tuple with the ConsolePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsolePort

`func (o *PatchedCableToCableTerminationRequest) SetConsolePort(v ApprovalWorkflowUser)`

SetConsolePort sets ConsolePort field to given value.

### HasConsolePort

`func (o *PatchedCableToCableTerminationRequest) HasConsolePort() bool`

HasConsolePort returns a boolean if a field has been set.

### SetConsolePortNil

`func (o *PatchedCableToCableTerminationRequest) SetConsolePortNil(b bool)`

 SetConsolePortNil sets the value for ConsolePort to be an explicit nil

### UnsetConsolePort
`func (o *PatchedCableToCableTerminationRequest) UnsetConsolePort()`

UnsetConsolePort ensures that no value is present for ConsolePort, not even an explicit nil
### GetConsoleServerPort

`func (o *PatchedCableToCableTerminationRequest) GetConsoleServerPort() ApprovalWorkflowUser`

GetConsoleServerPort returns the ConsoleServerPort field if non-nil, zero value otherwise.

### GetConsoleServerPortOk

`func (o *PatchedCableToCableTerminationRequest) GetConsoleServerPortOk() (*ApprovalWorkflowUser, bool)`

GetConsoleServerPortOk returns a tuple with the ConsoleServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleServerPort

`func (o *PatchedCableToCableTerminationRequest) SetConsoleServerPort(v ApprovalWorkflowUser)`

SetConsoleServerPort sets ConsoleServerPort field to given value.

### HasConsoleServerPort

`func (o *PatchedCableToCableTerminationRequest) HasConsoleServerPort() bool`

HasConsoleServerPort returns a boolean if a field has been set.

### SetConsoleServerPortNil

`func (o *PatchedCableToCableTerminationRequest) SetConsoleServerPortNil(b bool)`

 SetConsoleServerPortNil sets the value for ConsoleServerPort to be an explicit nil

### UnsetConsoleServerPort
`func (o *PatchedCableToCableTerminationRequest) UnsetConsoleServerPort()`

UnsetConsoleServerPort ensures that no value is present for ConsoleServerPort, not even an explicit nil
### GetFrontPort

`func (o *PatchedCableToCableTerminationRequest) GetFrontPort() ApprovalWorkflowUser`

GetFrontPort returns the FrontPort field if non-nil, zero value otherwise.

### GetFrontPortOk

`func (o *PatchedCableToCableTerminationRequest) GetFrontPortOk() (*ApprovalWorkflowUser, bool)`

GetFrontPortOk returns a tuple with the FrontPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontPort

`func (o *PatchedCableToCableTerminationRequest) SetFrontPort(v ApprovalWorkflowUser)`

SetFrontPort sets FrontPort field to given value.

### HasFrontPort

`func (o *PatchedCableToCableTerminationRequest) HasFrontPort() bool`

HasFrontPort returns a boolean if a field has been set.

### SetFrontPortNil

`func (o *PatchedCableToCableTerminationRequest) SetFrontPortNil(b bool)`

 SetFrontPortNil sets the value for FrontPort to be an explicit nil

### UnsetFrontPort
`func (o *PatchedCableToCableTerminationRequest) UnsetFrontPort()`

UnsetFrontPort ensures that no value is present for FrontPort, not even an explicit nil
### GetInterface

`func (o *PatchedCableToCableTerminationRequest) GetInterface() ApprovalWorkflowUser`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *PatchedCableToCableTerminationRequest) GetInterfaceOk() (*ApprovalWorkflowUser, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *PatchedCableToCableTerminationRequest) SetInterface(v ApprovalWorkflowUser)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *PatchedCableToCableTerminationRequest) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### SetInterfaceNil

`func (o *PatchedCableToCableTerminationRequest) SetInterfaceNil(b bool)`

 SetInterfaceNil sets the value for Interface to be an explicit nil

### UnsetInterface
`func (o *PatchedCableToCableTerminationRequest) UnsetInterface()`

UnsetInterface ensures that no value is present for Interface, not even an explicit nil
### GetPowerFeed

`func (o *PatchedCableToCableTerminationRequest) GetPowerFeed() ApprovalWorkflowUser`

GetPowerFeed returns the PowerFeed field if non-nil, zero value otherwise.

### GetPowerFeedOk

`func (o *PatchedCableToCableTerminationRequest) GetPowerFeedOk() (*ApprovalWorkflowUser, bool)`

GetPowerFeedOk returns a tuple with the PowerFeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerFeed

`func (o *PatchedCableToCableTerminationRequest) SetPowerFeed(v ApprovalWorkflowUser)`

SetPowerFeed sets PowerFeed field to given value.

### HasPowerFeed

`func (o *PatchedCableToCableTerminationRequest) HasPowerFeed() bool`

HasPowerFeed returns a boolean if a field has been set.

### SetPowerFeedNil

`func (o *PatchedCableToCableTerminationRequest) SetPowerFeedNil(b bool)`

 SetPowerFeedNil sets the value for PowerFeed to be an explicit nil

### UnsetPowerFeed
`func (o *PatchedCableToCableTerminationRequest) UnsetPowerFeed()`

UnsetPowerFeed ensures that no value is present for PowerFeed, not even an explicit nil
### GetPowerOutlet

`func (o *PatchedCableToCableTerminationRequest) GetPowerOutlet() ApprovalWorkflowUser`

GetPowerOutlet returns the PowerOutlet field if non-nil, zero value otherwise.

### GetPowerOutletOk

`func (o *PatchedCableToCableTerminationRequest) GetPowerOutletOk() (*ApprovalWorkflowUser, bool)`

GetPowerOutletOk returns a tuple with the PowerOutlet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOutlet

`func (o *PatchedCableToCableTerminationRequest) SetPowerOutlet(v ApprovalWorkflowUser)`

SetPowerOutlet sets PowerOutlet field to given value.

### HasPowerOutlet

`func (o *PatchedCableToCableTerminationRequest) HasPowerOutlet() bool`

HasPowerOutlet returns a boolean if a field has been set.

### SetPowerOutletNil

`func (o *PatchedCableToCableTerminationRequest) SetPowerOutletNil(b bool)`

 SetPowerOutletNil sets the value for PowerOutlet to be an explicit nil

### UnsetPowerOutlet
`func (o *PatchedCableToCableTerminationRequest) UnsetPowerOutlet()`

UnsetPowerOutlet ensures that no value is present for PowerOutlet, not even an explicit nil
### GetPowerPort

`func (o *PatchedCableToCableTerminationRequest) GetPowerPort() ApprovalWorkflowUser`

GetPowerPort returns the PowerPort field if non-nil, zero value otherwise.

### GetPowerPortOk

`func (o *PatchedCableToCableTerminationRequest) GetPowerPortOk() (*ApprovalWorkflowUser, bool)`

GetPowerPortOk returns a tuple with the PowerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerPort

`func (o *PatchedCableToCableTerminationRequest) SetPowerPort(v ApprovalWorkflowUser)`

SetPowerPort sets PowerPort field to given value.

### HasPowerPort

`func (o *PatchedCableToCableTerminationRequest) HasPowerPort() bool`

HasPowerPort returns a boolean if a field has been set.

### SetPowerPortNil

`func (o *PatchedCableToCableTerminationRequest) SetPowerPortNil(b bool)`

 SetPowerPortNil sets the value for PowerPort to be an explicit nil

### UnsetPowerPort
`func (o *PatchedCableToCableTerminationRequest) UnsetPowerPort()`

UnsetPowerPort ensures that no value is present for PowerPort, not even an explicit nil
### GetRearPort

`func (o *PatchedCableToCableTerminationRequest) GetRearPort() ApprovalWorkflowUser`

GetRearPort returns the RearPort field if non-nil, zero value otherwise.

### GetRearPortOk

`func (o *PatchedCableToCableTerminationRequest) GetRearPortOk() (*ApprovalWorkflowUser, bool)`

GetRearPortOk returns a tuple with the RearPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRearPort

`func (o *PatchedCableToCableTerminationRequest) SetRearPort(v ApprovalWorkflowUser)`

SetRearPort sets RearPort field to given value.

### HasRearPort

`func (o *PatchedCableToCableTerminationRequest) HasRearPort() bool`

HasRearPort returns a boolean if a field has been set.

### SetRearPortNil

`func (o *PatchedCableToCableTerminationRequest) SetRearPortNil(b bool)`

 SetRearPortNil sets the value for RearPort to be an explicit nil

### UnsetRearPort
`func (o *PatchedCableToCableTerminationRequest) UnsetRearPort()`

UnsetRearPort ensures that no value is present for RearPort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


