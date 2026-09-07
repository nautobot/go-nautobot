# CableToCableTerminationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**CableEnd** | [**CableEndEnum**](CableEndEnum.md) |  | 
**Connector** | Pointer to **int32** | The connector number on this cable end. Always 1 for standard cables. | [optional] [default to 1]
**Cable** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
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

### NewCableToCableTerminationRequest

`func NewCableToCableTerminationRequest(cableEnd CableEndEnum, cable BulkWritableCableRequestStatus, ) *CableToCableTerminationRequest`

NewCableToCableTerminationRequest instantiates a new CableToCableTerminationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCableToCableTerminationRequestWithDefaults

`func NewCableToCableTerminationRequestWithDefaults() *CableToCableTerminationRequest`

NewCableToCableTerminationRequestWithDefaults instantiates a new CableToCableTerminationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CableToCableTerminationRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CableToCableTerminationRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CableToCableTerminationRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CableToCableTerminationRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCableEnd

`func (o *CableToCableTerminationRequest) GetCableEnd() CableEndEnum`

GetCableEnd returns the CableEnd field if non-nil, zero value otherwise.

### GetCableEndOk

`func (o *CableToCableTerminationRequest) GetCableEndOk() (*CableEndEnum, bool)`

GetCableEndOk returns a tuple with the CableEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCableEnd

`func (o *CableToCableTerminationRequest) SetCableEnd(v CableEndEnum)`

SetCableEnd sets CableEnd field to given value.


### GetConnector

`func (o *CableToCableTerminationRequest) GetConnector() int32`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *CableToCableTerminationRequest) GetConnectorOk() (*int32, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *CableToCableTerminationRequest) SetConnector(v int32)`

SetConnector sets Connector field to given value.

### HasConnector

`func (o *CableToCableTerminationRequest) HasConnector() bool`

HasConnector returns a boolean if a field has been set.

### GetCable

`func (o *CableToCableTerminationRequest) GetCable() BulkWritableCableRequestStatus`

GetCable returns the Cable field if non-nil, zero value otherwise.

### GetCableOk

`func (o *CableToCableTerminationRequest) GetCableOk() (*BulkWritableCableRequestStatus, bool)`

GetCableOk returns a tuple with the Cable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCable

`func (o *CableToCableTerminationRequest) SetCable(v BulkWritableCableRequestStatus)`

SetCable sets Cable field to given value.


### GetCircuitTermination

`func (o *CableToCableTerminationRequest) GetCircuitTermination() ApprovalWorkflowUser`

GetCircuitTermination returns the CircuitTermination field if non-nil, zero value otherwise.

### GetCircuitTerminationOk

`func (o *CableToCableTerminationRequest) GetCircuitTerminationOk() (*ApprovalWorkflowUser, bool)`

GetCircuitTerminationOk returns a tuple with the CircuitTermination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitTermination

`func (o *CableToCableTerminationRequest) SetCircuitTermination(v ApprovalWorkflowUser)`

SetCircuitTermination sets CircuitTermination field to given value.

### HasCircuitTermination

`func (o *CableToCableTerminationRequest) HasCircuitTermination() bool`

HasCircuitTermination returns a boolean if a field has been set.

### SetCircuitTerminationNil

`func (o *CableToCableTerminationRequest) SetCircuitTerminationNil(b bool)`

 SetCircuitTerminationNil sets the value for CircuitTermination to be an explicit nil

### UnsetCircuitTermination
`func (o *CableToCableTerminationRequest) UnsetCircuitTermination()`

UnsetCircuitTermination ensures that no value is present for CircuitTermination, not even an explicit nil
### GetConsolePort

`func (o *CableToCableTerminationRequest) GetConsolePort() ApprovalWorkflowUser`

GetConsolePort returns the ConsolePort field if non-nil, zero value otherwise.

### GetConsolePortOk

`func (o *CableToCableTerminationRequest) GetConsolePortOk() (*ApprovalWorkflowUser, bool)`

GetConsolePortOk returns a tuple with the ConsolePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsolePort

`func (o *CableToCableTerminationRequest) SetConsolePort(v ApprovalWorkflowUser)`

SetConsolePort sets ConsolePort field to given value.

### HasConsolePort

`func (o *CableToCableTerminationRequest) HasConsolePort() bool`

HasConsolePort returns a boolean if a field has been set.

### SetConsolePortNil

`func (o *CableToCableTerminationRequest) SetConsolePortNil(b bool)`

 SetConsolePortNil sets the value for ConsolePort to be an explicit nil

### UnsetConsolePort
`func (o *CableToCableTerminationRequest) UnsetConsolePort()`

UnsetConsolePort ensures that no value is present for ConsolePort, not even an explicit nil
### GetConsoleServerPort

`func (o *CableToCableTerminationRequest) GetConsoleServerPort() ApprovalWorkflowUser`

GetConsoleServerPort returns the ConsoleServerPort field if non-nil, zero value otherwise.

### GetConsoleServerPortOk

`func (o *CableToCableTerminationRequest) GetConsoleServerPortOk() (*ApprovalWorkflowUser, bool)`

GetConsoleServerPortOk returns a tuple with the ConsoleServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleServerPort

`func (o *CableToCableTerminationRequest) SetConsoleServerPort(v ApprovalWorkflowUser)`

SetConsoleServerPort sets ConsoleServerPort field to given value.

### HasConsoleServerPort

`func (o *CableToCableTerminationRequest) HasConsoleServerPort() bool`

HasConsoleServerPort returns a boolean if a field has been set.

### SetConsoleServerPortNil

`func (o *CableToCableTerminationRequest) SetConsoleServerPortNil(b bool)`

 SetConsoleServerPortNil sets the value for ConsoleServerPort to be an explicit nil

### UnsetConsoleServerPort
`func (o *CableToCableTerminationRequest) UnsetConsoleServerPort()`

UnsetConsoleServerPort ensures that no value is present for ConsoleServerPort, not even an explicit nil
### GetFrontPort

`func (o *CableToCableTerminationRequest) GetFrontPort() ApprovalWorkflowUser`

GetFrontPort returns the FrontPort field if non-nil, zero value otherwise.

### GetFrontPortOk

`func (o *CableToCableTerminationRequest) GetFrontPortOk() (*ApprovalWorkflowUser, bool)`

GetFrontPortOk returns a tuple with the FrontPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontPort

`func (o *CableToCableTerminationRequest) SetFrontPort(v ApprovalWorkflowUser)`

SetFrontPort sets FrontPort field to given value.

### HasFrontPort

`func (o *CableToCableTerminationRequest) HasFrontPort() bool`

HasFrontPort returns a boolean if a field has been set.

### SetFrontPortNil

`func (o *CableToCableTerminationRequest) SetFrontPortNil(b bool)`

 SetFrontPortNil sets the value for FrontPort to be an explicit nil

### UnsetFrontPort
`func (o *CableToCableTerminationRequest) UnsetFrontPort()`

UnsetFrontPort ensures that no value is present for FrontPort, not even an explicit nil
### GetInterface

`func (o *CableToCableTerminationRequest) GetInterface() ApprovalWorkflowUser`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *CableToCableTerminationRequest) GetInterfaceOk() (*ApprovalWorkflowUser, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *CableToCableTerminationRequest) SetInterface(v ApprovalWorkflowUser)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *CableToCableTerminationRequest) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### SetInterfaceNil

`func (o *CableToCableTerminationRequest) SetInterfaceNil(b bool)`

 SetInterfaceNil sets the value for Interface to be an explicit nil

### UnsetInterface
`func (o *CableToCableTerminationRequest) UnsetInterface()`

UnsetInterface ensures that no value is present for Interface, not even an explicit nil
### GetPowerFeed

`func (o *CableToCableTerminationRequest) GetPowerFeed() ApprovalWorkflowUser`

GetPowerFeed returns the PowerFeed field if non-nil, zero value otherwise.

### GetPowerFeedOk

`func (o *CableToCableTerminationRequest) GetPowerFeedOk() (*ApprovalWorkflowUser, bool)`

GetPowerFeedOk returns a tuple with the PowerFeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerFeed

`func (o *CableToCableTerminationRequest) SetPowerFeed(v ApprovalWorkflowUser)`

SetPowerFeed sets PowerFeed field to given value.

### HasPowerFeed

`func (o *CableToCableTerminationRequest) HasPowerFeed() bool`

HasPowerFeed returns a boolean if a field has been set.

### SetPowerFeedNil

`func (o *CableToCableTerminationRequest) SetPowerFeedNil(b bool)`

 SetPowerFeedNil sets the value for PowerFeed to be an explicit nil

### UnsetPowerFeed
`func (o *CableToCableTerminationRequest) UnsetPowerFeed()`

UnsetPowerFeed ensures that no value is present for PowerFeed, not even an explicit nil
### GetPowerOutlet

`func (o *CableToCableTerminationRequest) GetPowerOutlet() ApprovalWorkflowUser`

GetPowerOutlet returns the PowerOutlet field if non-nil, zero value otherwise.

### GetPowerOutletOk

`func (o *CableToCableTerminationRequest) GetPowerOutletOk() (*ApprovalWorkflowUser, bool)`

GetPowerOutletOk returns a tuple with the PowerOutlet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOutlet

`func (o *CableToCableTerminationRequest) SetPowerOutlet(v ApprovalWorkflowUser)`

SetPowerOutlet sets PowerOutlet field to given value.

### HasPowerOutlet

`func (o *CableToCableTerminationRequest) HasPowerOutlet() bool`

HasPowerOutlet returns a boolean if a field has been set.

### SetPowerOutletNil

`func (o *CableToCableTerminationRequest) SetPowerOutletNil(b bool)`

 SetPowerOutletNil sets the value for PowerOutlet to be an explicit nil

### UnsetPowerOutlet
`func (o *CableToCableTerminationRequest) UnsetPowerOutlet()`

UnsetPowerOutlet ensures that no value is present for PowerOutlet, not even an explicit nil
### GetPowerPort

`func (o *CableToCableTerminationRequest) GetPowerPort() ApprovalWorkflowUser`

GetPowerPort returns the PowerPort field if non-nil, zero value otherwise.

### GetPowerPortOk

`func (o *CableToCableTerminationRequest) GetPowerPortOk() (*ApprovalWorkflowUser, bool)`

GetPowerPortOk returns a tuple with the PowerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerPort

`func (o *CableToCableTerminationRequest) SetPowerPort(v ApprovalWorkflowUser)`

SetPowerPort sets PowerPort field to given value.

### HasPowerPort

`func (o *CableToCableTerminationRequest) HasPowerPort() bool`

HasPowerPort returns a boolean if a field has been set.

### SetPowerPortNil

`func (o *CableToCableTerminationRequest) SetPowerPortNil(b bool)`

 SetPowerPortNil sets the value for PowerPort to be an explicit nil

### UnsetPowerPort
`func (o *CableToCableTerminationRequest) UnsetPowerPort()`

UnsetPowerPort ensures that no value is present for PowerPort, not even an explicit nil
### GetRearPort

`func (o *CableToCableTerminationRequest) GetRearPort() ApprovalWorkflowUser`

GetRearPort returns the RearPort field if non-nil, zero value otherwise.

### GetRearPortOk

`func (o *CableToCableTerminationRequest) GetRearPortOk() (*ApprovalWorkflowUser, bool)`

GetRearPortOk returns a tuple with the RearPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRearPort

`func (o *CableToCableTerminationRequest) SetRearPort(v ApprovalWorkflowUser)`

SetRearPort sets RearPort field to given value.

### HasRearPort

`func (o *CableToCableTerminationRequest) HasRearPort() bool`

HasRearPort returns a boolean if a field has been set.

### SetRearPortNil

`func (o *CableToCableTerminationRequest) SetRearPortNil(b bool)`

 SetRearPortNil sets the value for RearPort to be an explicit nil

### UnsetRearPort
`func (o *CableToCableTerminationRequest) UnsetRearPort()`

UnsetRearPort ensures that no value is present for RearPort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


