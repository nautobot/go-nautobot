# PatchedBulkWritableCableToCableTerminationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
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

### NewPatchedBulkWritableCableToCableTerminationRequest

`func NewPatchedBulkWritableCableToCableTerminationRequest(id string, ) *PatchedBulkWritableCableToCableTerminationRequest`

NewPatchedBulkWritableCableToCableTerminationRequest instantiates a new PatchedBulkWritableCableToCableTerminationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableCableToCableTerminationRequestWithDefaults

`func NewPatchedBulkWritableCableToCableTerminationRequestWithDefaults() *PatchedBulkWritableCableToCableTerminationRequest`

NewPatchedBulkWritableCableToCableTerminationRequestWithDefaults instantiates a new PatchedBulkWritableCableToCableTerminationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableCableToCableTerminationRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableCableToCableTerminationRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableCableToCableTerminationRequest) SetId(v string)`

SetId sets Id field to given value.


### GetCableEnd

`func (o *PatchedBulkWritableCableToCableTerminationRequest) GetCableEnd() CableEndEnum`

GetCableEnd returns the CableEnd field if non-nil, zero value otherwise.

### GetCableEndOk

`func (o *PatchedBulkWritableCableToCableTerminationRequest) GetCableEndOk() (*CableEndEnum, bool)`

GetCableEndOk returns a tuple with the CableEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCableEnd

`func (o *PatchedBulkWritableCableToCableTerminationRequest) SetCableEnd(v CableEndEnum)`

SetCableEnd sets CableEnd field to given value.

### HasCableEnd

`func (o *PatchedBulkWritableCableToCableTerminationRequest) HasCableEnd() bool`

HasCableEnd returns a boolean if a field has been set.

### GetConnector

`func (o *PatchedBulkWritableCableToCableTerminationRequest) GetConnector() int32`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *PatchedBulkWritableCableToCableTerminationRequest) GetConnectorOk() (*int32, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *PatchedBulkWritableCableToCableTerminationRequest) SetConnector(v int32)`

SetConnector sets Connector field to given value.

### HasConnector

`func (o *PatchedBulkWritableCableToCableTerminationRequest) HasConnector() bool`

HasConnector returns a boolean if a field has been set.

### GetCable

`func (o *PatchedBulkWritableCableToCableTerminationRequest) GetCable() BulkWritableCableRequestStatus`

GetCable returns the Cable field if non-nil, zero value otherwise.

### GetCableOk

`func (o *PatchedBulkWritableCableToCableTerminationRequest) GetCableOk() (*BulkWritableCableRequestStatus, bool)`

GetCableOk returns a tuple with the Cable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCable

`func (o *PatchedBulkWritableCableToCableTerminationRequest) SetCable(v BulkWritableCableRequestStatus)`

SetCable sets Cable field to given value.

### HasCable

`func (o *PatchedBulkWritableCableToCableTerminationRequest) HasCable() bool`

HasCable returns a boolean if a field has been set.

### GetCircuitTermination

`func (o *PatchedBulkWritableCableToCableTerminationRequest) GetCircuitTermination() ApprovalWorkflowUser`

GetCircuitTermination returns the CircuitTermination field if non-nil, zero value otherwise.

### GetCircuitTerminationOk

`func (o *PatchedBulkWritableCableToCableTerminationRequest) GetCircuitTerminationOk() (*ApprovalWorkflowUser, bool)`

GetCircuitTerminationOk returns a tuple with the CircuitTermination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitTermination

`func (o *PatchedBulkWritableCableToCableTerminationRequest) SetCircuitTermination(v ApprovalWorkflowUser)`

SetCircuitTermination sets CircuitTermination field to given value.

### HasCircuitTermination

`func (o *PatchedBulkWritableCableToCableTerminationRequest) HasCircuitTermination() bool`

HasCircuitTermination returns a boolean if a field has been set.

### SetCircuitTerminationNil

`func (o *PatchedBulkWritableCableToCableTerminationRequest) SetCircuitTerminationNil(b bool)`

 SetCircuitTerminationNil sets the value for CircuitTermination to be an explicit nil

### UnsetCircuitTermination
`func (o *PatchedBulkWritableCableToCableTerminationRequest) UnsetCircuitTermination()`

UnsetCircuitTermination ensures that no value is present for CircuitTermination, not even an explicit nil
### GetConsolePort

`func (o *PatchedBulkWritableCableToCableTerminationRequest) GetConsolePort() ApprovalWorkflowUser`

GetConsolePort returns the ConsolePort field if non-nil, zero value otherwise.

### GetConsolePortOk

`func (o *PatchedBulkWritableCableToCableTerminationRequest) GetConsolePortOk() (*ApprovalWorkflowUser, bool)`

GetConsolePortOk returns a tuple with the ConsolePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsolePort

`func (o *PatchedBulkWritableCableToCableTerminationRequest) SetConsolePort(v ApprovalWorkflowUser)`

SetConsolePort sets ConsolePort field to given value.

### HasConsolePort

`func (o *PatchedBulkWritableCableToCableTerminationRequest) HasConsolePort() bool`

HasConsolePort returns a boolean if a field has been set.

### SetConsolePortNil

`func (o *PatchedBulkWritableCableToCableTerminationRequest) SetConsolePortNil(b bool)`

 SetConsolePortNil sets the value for ConsolePort to be an explicit nil

### UnsetConsolePort
`func (o *PatchedBulkWritableCableToCableTerminationRequest) UnsetConsolePort()`

UnsetConsolePort ensures that no value is present for ConsolePort, not even an explicit nil
### GetConsoleServerPort

`func (o *PatchedBulkWritableCableToCableTerminationRequest) GetConsoleServerPort() ApprovalWorkflowUser`

GetConsoleServerPort returns the ConsoleServerPort field if non-nil, zero value otherwise.

### GetConsoleServerPortOk

`func (o *PatchedBulkWritableCableToCableTerminationRequest) GetConsoleServerPortOk() (*ApprovalWorkflowUser, bool)`

GetConsoleServerPortOk returns a tuple with the ConsoleServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleServerPort

`func (o *PatchedBulkWritableCableToCableTerminationRequest) SetConsoleServerPort(v ApprovalWorkflowUser)`

SetConsoleServerPort sets ConsoleServerPort field to given value.

### HasConsoleServerPort

`func (o *PatchedBulkWritableCableToCableTerminationRequest) HasConsoleServerPort() bool`

HasConsoleServerPort returns a boolean if a field has been set.

### SetConsoleServerPortNil

`func (o *PatchedBulkWritableCableToCableTerminationRequest) SetConsoleServerPortNil(b bool)`

 SetConsoleServerPortNil sets the value for ConsoleServerPort to be an explicit nil

### UnsetConsoleServerPort
`func (o *PatchedBulkWritableCableToCableTerminationRequest) UnsetConsoleServerPort()`

UnsetConsoleServerPort ensures that no value is present for ConsoleServerPort, not even an explicit nil
### GetFrontPort

`func (o *PatchedBulkWritableCableToCableTerminationRequest) GetFrontPort() ApprovalWorkflowUser`

GetFrontPort returns the FrontPort field if non-nil, zero value otherwise.

### GetFrontPortOk

`func (o *PatchedBulkWritableCableToCableTerminationRequest) GetFrontPortOk() (*ApprovalWorkflowUser, bool)`

GetFrontPortOk returns a tuple with the FrontPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontPort

`func (o *PatchedBulkWritableCableToCableTerminationRequest) SetFrontPort(v ApprovalWorkflowUser)`

SetFrontPort sets FrontPort field to given value.

### HasFrontPort

`func (o *PatchedBulkWritableCableToCableTerminationRequest) HasFrontPort() bool`

HasFrontPort returns a boolean if a field has been set.

### SetFrontPortNil

`func (o *PatchedBulkWritableCableToCableTerminationRequest) SetFrontPortNil(b bool)`

 SetFrontPortNil sets the value for FrontPort to be an explicit nil

### UnsetFrontPort
`func (o *PatchedBulkWritableCableToCableTerminationRequest) UnsetFrontPort()`

UnsetFrontPort ensures that no value is present for FrontPort, not even an explicit nil
### GetInterface

`func (o *PatchedBulkWritableCableToCableTerminationRequest) GetInterface() ApprovalWorkflowUser`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *PatchedBulkWritableCableToCableTerminationRequest) GetInterfaceOk() (*ApprovalWorkflowUser, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *PatchedBulkWritableCableToCableTerminationRequest) SetInterface(v ApprovalWorkflowUser)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *PatchedBulkWritableCableToCableTerminationRequest) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### SetInterfaceNil

`func (o *PatchedBulkWritableCableToCableTerminationRequest) SetInterfaceNil(b bool)`

 SetInterfaceNil sets the value for Interface to be an explicit nil

### UnsetInterface
`func (o *PatchedBulkWritableCableToCableTerminationRequest) UnsetInterface()`

UnsetInterface ensures that no value is present for Interface, not even an explicit nil
### GetPowerFeed

`func (o *PatchedBulkWritableCableToCableTerminationRequest) GetPowerFeed() ApprovalWorkflowUser`

GetPowerFeed returns the PowerFeed field if non-nil, zero value otherwise.

### GetPowerFeedOk

`func (o *PatchedBulkWritableCableToCableTerminationRequest) GetPowerFeedOk() (*ApprovalWorkflowUser, bool)`

GetPowerFeedOk returns a tuple with the PowerFeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerFeed

`func (o *PatchedBulkWritableCableToCableTerminationRequest) SetPowerFeed(v ApprovalWorkflowUser)`

SetPowerFeed sets PowerFeed field to given value.

### HasPowerFeed

`func (o *PatchedBulkWritableCableToCableTerminationRequest) HasPowerFeed() bool`

HasPowerFeed returns a boolean if a field has been set.

### SetPowerFeedNil

`func (o *PatchedBulkWritableCableToCableTerminationRequest) SetPowerFeedNil(b bool)`

 SetPowerFeedNil sets the value for PowerFeed to be an explicit nil

### UnsetPowerFeed
`func (o *PatchedBulkWritableCableToCableTerminationRequest) UnsetPowerFeed()`

UnsetPowerFeed ensures that no value is present for PowerFeed, not even an explicit nil
### GetPowerOutlet

`func (o *PatchedBulkWritableCableToCableTerminationRequest) GetPowerOutlet() ApprovalWorkflowUser`

GetPowerOutlet returns the PowerOutlet field if non-nil, zero value otherwise.

### GetPowerOutletOk

`func (o *PatchedBulkWritableCableToCableTerminationRequest) GetPowerOutletOk() (*ApprovalWorkflowUser, bool)`

GetPowerOutletOk returns a tuple with the PowerOutlet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOutlet

`func (o *PatchedBulkWritableCableToCableTerminationRequest) SetPowerOutlet(v ApprovalWorkflowUser)`

SetPowerOutlet sets PowerOutlet field to given value.

### HasPowerOutlet

`func (o *PatchedBulkWritableCableToCableTerminationRequest) HasPowerOutlet() bool`

HasPowerOutlet returns a boolean if a field has been set.

### SetPowerOutletNil

`func (o *PatchedBulkWritableCableToCableTerminationRequest) SetPowerOutletNil(b bool)`

 SetPowerOutletNil sets the value for PowerOutlet to be an explicit nil

### UnsetPowerOutlet
`func (o *PatchedBulkWritableCableToCableTerminationRequest) UnsetPowerOutlet()`

UnsetPowerOutlet ensures that no value is present for PowerOutlet, not even an explicit nil
### GetPowerPort

`func (o *PatchedBulkWritableCableToCableTerminationRequest) GetPowerPort() ApprovalWorkflowUser`

GetPowerPort returns the PowerPort field if non-nil, zero value otherwise.

### GetPowerPortOk

`func (o *PatchedBulkWritableCableToCableTerminationRequest) GetPowerPortOk() (*ApprovalWorkflowUser, bool)`

GetPowerPortOk returns a tuple with the PowerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerPort

`func (o *PatchedBulkWritableCableToCableTerminationRequest) SetPowerPort(v ApprovalWorkflowUser)`

SetPowerPort sets PowerPort field to given value.

### HasPowerPort

`func (o *PatchedBulkWritableCableToCableTerminationRequest) HasPowerPort() bool`

HasPowerPort returns a boolean if a field has been set.

### SetPowerPortNil

`func (o *PatchedBulkWritableCableToCableTerminationRequest) SetPowerPortNil(b bool)`

 SetPowerPortNil sets the value for PowerPort to be an explicit nil

### UnsetPowerPort
`func (o *PatchedBulkWritableCableToCableTerminationRequest) UnsetPowerPort()`

UnsetPowerPort ensures that no value is present for PowerPort, not even an explicit nil
### GetRearPort

`func (o *PatchedBulkWritableCableToCableTerminationRequest) GetRearPort() ApprovalWorkflowUser`

GetRearPort returns the RearPort field if non-nil, zero value otherwise.

### GetRearPortOk

`func (o *PatchedBulkWritableCableToCableTerminationRequest) GetRearPortOk() (*ApprovalWorkflowUser, bool)`

GetRearPortOk returns a tuple with the RearPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRearPort

`func (o *PatchedBulkWritableCableToCableTerminationRequest) SetRearPort(v ApprovalWorkflowUser)`

SetRearPort sets RearPort field to given value.

### HasRearPort

`func (o *PatchedBulkWritableCableToCableTerminationRequest) HasRearPort() bool`

HasRearPort returns a boolean if a field has been set.

### SetRearPortNil

`func (o *PatchedBulkWritableCableToCableTerminationRequest) SetRearPortNil(b bool)`

 SetRearPortNil sets the value for RearPort to be an explicit nil

### UnsetRearPort
`func (o *PatchedBulkWritableCableToCableTerminationRequest) UnsetRearPort()`

UnsetRearPort ensures that no value is present for RearPort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


