# BulkWritableCableToCableTerminationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
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

### NewBulkWritableCableToCableTerminationRequest

`func NewBulkWritableCableToCableTerminationRequest(id string, cableEnd CableEndEnum, cable BulkWritableCableRequestStatus, ) *BulkWritableCableToCableTerminationRequest`

NewBulkWritableCableToCableTerminationRequest instantiates a new BulkWritableCableToCableTerminationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableCableToCableTerminationRequestWithDefaults

`func NewBulkWritableCableToCableTerminationRequestWithDefaults() *BulkWritableCableToCableTerminationRequest`

NewBulkWritableCableToCableTerminationRequestWithDefaults instantiates a new BulkWritableCableToCableTerminationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableCableToCableTerminationRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableCableToCableTerminationRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableCableToCableTerminationRequest) SetId(v string)`

SetId sets Id field to given value.


### GetCableEnd

`func (o *BulkWritableCableToCableTerminationRequest) GetCableEnd() CableEndEnum`

GetCableEnd returns the CableEnd field if non-nil, zero value otherwise.

### GetCableEndOk

`func (o *BulkWritableCableToCableTerminationRequest) GetCableEndOk() (*CableEndEnum, bool)`

GetCableEndOk returns a tuple with the CableEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCableEnd

`func (o *BulkWritableCableToCableTerminationRequest) SetCableEnd(v CableEndEnum)`

SetCableEnd sets CableEnd field to given value.


### GetConnector

`func (o *BulkWritableCableToCableTerminationRequest) GetConnector() int32`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *BulkWritableCableToCableTerminationRequest) GetConnectorOk() (*int32, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *BulkWritableCableToCableTerminationRequest) SetConnector(v int32)`

SetConnector sets Connector field to given value.

### HasConnector

`func (o *BulkWritableCableToCableTerminationRequest) HasConnector() bool`

HasConnector returns a boolean if a field has been set.

### GetCable

`func (o *BulkWritableCableToCableTerminationRequest) GetCable() BulkWritableCableRequestStatus`

GetCable returns the Cable field if non-nil, zero value otherwise.

### GetCableOk

`func (o *BulkWritableCableToCableTerminationRequest) GetCableOk() (*BulkWritableCableRequestStatus, bool)`

GetCableOk returns a tuple with the Cable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCable

`func (o *BulkWritableCableToCableTerminationRequest) SetCable(v BulkWritableCableRequestStatus)`

SetCable sets Cable field to given value.


### GetCircuitTermination

`func (o *BulkWritableCableToCableTerminationRequest) GetCircuitTermination() ApprovalWorkflowUser`

GetCircuitTermination returns the CircuitTermination field if non-nil, zero value otherwise.

### GetCircuitTerminationOk

`func (o *BulkWritableCableToCableTerminationRequest) GetCircuitTerminationOk() (*ApprovalWorkflowUser, bool)`

GetCircuitTerminationOk returns a tuple with the CircuitTermination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitTermination

`func (o *BulkWritableCableToCableTerminationRequest) SetCircuitTermination(v ApprovalWorkflowUser)`

SetCircuitTermination sets CircuitTermination field to given value.

### HasCircuitTermination

`func (o *BulkWritableCableToCableTerminationRequest) HasCircuitTermination() bool`

HasCircuitTermination returns a boolean if a field has been set.

### SetCircuitTerminationNil

`func (o *BulkWritableCableToCableTerminationRequest) SetCircuitTerminationNil(b bool)`

 SetCircuitTerminationNil sets the value for CircuitTermination to be an explicit nil

### UnsetCircuitTermination
`func (o *BulkWritableCableToCableTerminationRequest) UnsetCircuitTermination()`

UnsetCircuitTermination ensures that no value is present for CircuitTermination, not even an explicit nil
### GetConsolePort

`func (o *BulkWritableCableToCableTerminationRequest) GetConsolePort() ApprovalWorkflowUser`

GetConsolePort returns the ConsolePort field if non-nil, zero value otherwise.

### GetConsolePortOk

`func (o *BulkWritableCableToCableTerminationRequest) GetConsolePortOk() (*ApprovalWorkflowUser, bool)`

GetConsolePortOk returns a tuple with the ConsolePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsolePort

`func (o *BulkWritableCableToCableTerminationRequest) SetConsolePort(v ApprovalWorkflowUser)`

SetConsolePort sets ConsolePort field to given value.

### HasConsolePort

`func (o *BulkWritableCableToCableTerminationRequest) HasConsolePort() bool`

HasConsolePort returns a boolean if a field has been set.

### SetConsolePortNil

`func (o *BulkWritableCableToCableTerminationRequest) SetConsolePortNil(b bool)`

 SetConsolePortNil sets the value for ConsolePort to be an explicit nil

### UnsetConsolePort
`func (o *BulkWritableCableToCableTerminationRequest) UnsetConsolePort()`

UnsetConsolePort ensures that no value is present for ConsolePort, not even an explicit nil
### GetConsoleServerPort

`func (o *BulkWritableCableToCableTerminationRequest) GetConsoleServerPort() ApprovalWorkflowUser`

GetConsoleServerPort returns the ConsoleServerPort field if non-nil, zero value otherwise.

### GetConsoleServerPortOk

`func (o *BulkWritableCableToCableTerminationRequest) GetConsoleServerPortOk() (*ApprovalWorkflowUser, bool)`

GetConsoleServerPortOk returns a tuple with the ConsoleServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleServerPort

`func (o *BulkWritableCableToCableTerminationRequest) SetConsoleServerPort(v ApprovalWorkflowUser)`

SetConsoleServerPort sets ConsoleServerPort field to given value.

### HasConsoleServerPort

`func (o *BulkWritableCableToCableTerminationRequest) HasConsoleServerPort() bool`

HasConsoleServerPort returns a boolean if a field has been set.

### SetConsoleServerPortNil

`func (o *BulkWritableCableToCableTerminationRequest) SetConsoleServerPortNil(b bool)`

 SetConsoleServerPortNil sets the value for ConsoleServerPort to be an explicit nil

### UnsetConsoleServerPort
`func (o *BulkWritableCableToCableTerminationRequest) UnsetConsoleServerPort()`

UnsetConsoleServerPort ensures that no value is present for ConsoleServerPort, not even an explicit nil
### GetFrontPort

`func (o *BulkWritableCableToCableTerminationRequest) GetFrontPort() ApprovalWorkflowUser`

GetFrontPort returns the FrontPort field if non-nil, zero value otherwise.

### GetFrontPortOk

`func (o *BulkWritableCableToCableTerminationRequest) GetFrontPortOk() (*ApprovalWorkflowUser, bool)`

GetFrontPortOk returns a tuple with the FrontPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontPort

`func (o *BulkWritableCableToCableTerminationRequest) SetFrontPort(v ApprovalWorkflowUser)`

SetFrontPort sets FrontPort field to given value.

### HasFrontPort

`func (o *BulkWritableCableToCableTerminationRequest) HasFrontPort() bool`

HasFrontPort returns a boolean if a field has been set.

### SetFrontPortNil

`func (o *BulkWritableCableToCableTerminationRequest) SetFrontPortNil(b bool)`

 SetFrontPortNil sets the value for FrontPort to be an explicit nil

### UnsetFrontPort
`func (o *BulkWritableCableToCableTerminationRequest) UnsetFrontPort()`

UnsetFrontPort ensures that no value is present for FrontPort, not even an explicit nil
### GetInterface

`func (o *BulkWritableCableToCableTerminationRequest) GetInterface() ApprovalWorkflowUser`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *BulkWritableCableToCableTerminationRequest) GetInterfaceOk() (*ApprovalWorkflowUser, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *BulkWritableCableToCableTerminationRequest) SetInterface(v ApprovalWorkflowUser)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *BulkWritableCableToCableTerminationRequest) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### SetInterfaceNil

`func (o *BulkWritableCableToCableTerminationRequest) SetInterfaceNil(b bool)`

 SetInterfaceNil sets the value for Interface to be an explicit nil

### UnsetInterface
`func (o *BulkWritableCableToCableTerminationRequest) UnsetInterface()`

UnsetInterface ensures that no value is present for Interface, not even an explicit nil
### GetPowerFeed

`func (o *BulkWritableCableToCableTerminationRequest) GetPowerFeed() ApprovalWorkflowUser`

GetPowerFeed returns the PowerFeed field if non-nil, zero value otherwise.

### GetPowerFeedOk

`func (o *BulkWritableCableToCableTerminationRequest) GetPowerFeedOk() (*ApprovalWorkflowUser, bool)`

GetPowerFeedOk returns a tuple with the PowerFeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerFeed

`func (o *BulkWritableCableToCableTerminationRequest) SetPowerFeed(v ApprovalWorkflowUser)`

SetPowerFeed sets PowerFeed field to given value.

### HasPowerFeed

`func (o *BulkWritableCableToCableTerminationRequest) HasPowerFeed() bool`

HasPowerFeed returns a boolean if a field has been set.

### SetPowerFeedNil

`func (o *BulkWritableCableToCableTerminationRequest) SetPowerFeedNil(b bool)`

 SetPowerFeedNil sets the value for PowerFeed to be an explicit nil

### UnsetPowerFeed
`func (o *BulkWritableCableToCableTerminationRequest) UnsetPowerFeed()`

UnsetPowerFeed ensures that no value is present for PowerFeed, not even an explicit nil
### GetPowerOutlet

`func (o *BulkWritableCableToCableTerminationRequest) GetPowerOutlet() ApprovalWorkflowUser`

GetPowerOutlet returns the PowerOutlet field if non-nil, zero value otherwise.

### GetPowerOutletOk

`func (o *BulkWritableCableToCableTerminationRequest) GetPowerOutletOk() (*ApprovalWorkflowUser, bool)`

GetPowerOutletOk returns a tuple with the PowerOutlet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOutlet

`func (o *BulkWritableCableToCableTerminationRequest) SetPowerOutlet(v ApprovalWorkflowUser)`

SetPowerOutlet sets PowerOutlet field to given value.

### HasPowerOutlet

`func (o *BulkWritableCableToCableTerminationRequest) HasPowerOutlet() bool`

HasPowerOutlet returns a boolean if a field has been set.

### SetPowerOutletNil

`func (o *BulkWritableCableToCableTerminationRequest) SetPowerOutletNil(b bool)`

 SetPowerOutletNil sets the value for PowerOutlet to be an explicit nil

### UnsetPowerOutlet
`func (o *BulkWritableCableToCableTerminationRequest) UnsetPowerOutlet()`

UnsetPowerOutlet ensures that no value is present for PowerOutlet, not even an explicit nil
### GetPowerPort

`func (o *BulkWritableCableToCableTerminationRequest) GetPowerPort() ApprovalWorkflowUser`

GetPowerPort returns the PowerPort field if non-nil, zero value otherwise.

### GetPowerPortOk

`func (o *BulkWritableCableToCableTerminationRequest) GetPowerPortOk() (*ApprovalWorkflowUser, bool)`

GetPowerPortOk returns a tuple with the PowerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerPort

`func (o *BulkWritableCableToCableTerminationRequest) SetPowerPort(v ApprovalWorkflowUser)`

SetPowerPort sets PowerPort field to given value.

### HasPowerPort

`func (o *BulkWritableCableToCableTerminationRequest) HasPowerPort() bool`

HasPowerPort returns a boolean if a field has been set.

### SetPowerPortNil

`func (o *BulkWritableCableToCableTerminationRequest) SetPowerPortNil(b bool)`

 SetPowerPortNil sets the value for PowerPort to be an explicit nil

### UnsetPowerPort
`func (o *BulkWritableCableToCableTerminationRequest) UnsetPowerPort()`

UnsetPowerPort ensures that no value is present for PowerPort, not even an explicit nil
### GetRearPort

`func (o *BulkWritableCableToCableTerminationRequest) GetRearPort() ApprovalWorkflowUser`

GetRearPort returns the RearPort field if non-nil, zero value otherwise.

### GetRearPortOk

`func (o *BulkWritableCableToCableTerminationRequest) GetRearPortOk() (*ApprovalWorkflowUser, bool)`

GetRearPortOk returns a tuple with the RearPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRearPort

`func (o *BulkWritableCableToCableTerminationRequest) SetRearPort(v ApprovalWorkflowUser)`

SetRearPort sets RearPort field to given value.

### HasRearPort

`func (o *BulkWritableCableToCableTerminationRequest) HasRearPort() bool`

HasRearPort returns a boolean if a field has been set.

### SetRearPortNil

`func (o *BulkWritableCableToCableTerminationRequest) SetRearPortNil(b bool)`

 SetRearPortNil sets the value for RearPort to be an explicit nil

### UnsetRearPort
`func (o *BulkWritableCableToCableTerminationRequest) UnsetRearPort()`

UnsetRearPort ensures that no value is present for RearPort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


