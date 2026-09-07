# CableToCableTermination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
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

### NewCableToCableTermination

`func NewCableToCableTermination(objectType string, display string, url string, naturalSlug string, cableEnd CableEndEnum, cable BulkWritableCableRequestStatus, ) *CableToCableTermination`

NewCableToCableTermination instantiates a new CableToCableTermination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCableToCableTerminationWithDefaults

`func NewCableToCableTerminationWithDefaults() *CableToCableTermination`

NewCableToCableTerminationWithDefaults instantiates a new CableToCableTermination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CableToCableTermination) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CableToCableTermination) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CableToCableTermination) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CableToCableTermination) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectType

`func (o *CableToCableTermination) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CableToCableTermination) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CableToCableTermination) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *CableToCableTermination) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *CableToCableTermination) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *CableToCableTermination) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *CableToCableTermination) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CableToCableTermination) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CableToCableTermination) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *CableToCableTermination) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *CableToCableTermination) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *CableToCableTermination) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetCableEnd

`func (o *CableToCableTermination) GetCableEnd() CableEndEnum`

GetCableEnd returns the CableEnd field if non-nil, zero value otherwise.

### GetCableEndOk

`func (o *CableToCableTermination) GetCableEndOk() (*CableEndEnum, bool)`

GetCableEndOk returns a tuple with the CableEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCableEnd

`func (o *CableToCableTermination) SetCableEnd(v CableEndEnum)`

SetCableEnd sets CableEnd field to given value.


### GetConnector

`func (o *CableToCableTermination) GetConnector() int32`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *CableToCableTermination) GetConnectorOk() (*int32, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *CableToCableTermination) SetConnector(v int32)`

SetConnector sets Connector field to given value.

### HasConnector

`func (o *CableToCableTermination) HasConnector() bool`

HasConnector returns a boolean if a field has been set.

### GetCable

`func (o *CableToCableTermination) GetCable() BulkWritableCableRequestStatus`

GetCable returns the Cable field if non-nil, zero value otherwise.

### GetCableOk

`func (o *CableToCableTermination) GetCableOk() (*BulkWritableCableRequestStatus, bool)`

GetCableOk returns a tuple with the Cable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCable

`func (o *CableToCableTermination) SetCable(v BulkWritableCableRequestStatus)`

SetCable sets Cable field to given value.


### GetCircuitTermination

`func (o *CableToCableTermination) GetCircuitTermination() ApprovalWorkflowUser`

GetCircuitTermination returns the CircuitTermination field if non-nil, zero value otherwise.

### GetCircuitTerminationOk

`func (o *CableToCableTermination) GetCircuitTerminationOk() (*ApprovalWorkflowUser, bool)`

GetCircuitTerminationOk returns a tuple with the CircuitTermination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitTermination

`func (o *CableToCableTermination) SetCircuitTermination(v ApprovalWorkflowUser)`

SetCircuitTermination sets CircuitTermination field to given value.

### HasCircuitTermination

`func (o *CableToCableTermination) HasCircuitTermination() bool`

HasCircuitTermination returns a boolean if a field has been set.

### SetCircuitTerminationNil

`func (o *CableToCableTermination) SetCircuitTerminationNil(b bool)`

 SetCircuitTerminationNil sets the value for CircuitTermination to be an explicit nil

### UnsetCircuitTermination
`func (o *CableToCableTermination) UnsetCircuitTermination()`

UnsetCircuitTermination ensures that no value is present for CircuitTermination, not even an explicit nil
### GetConsolePort

`func (o *CableToCableTermination) GetConsolePort() ApprovalWorkflowUser`

GetConsolePort returns the ConsolePort field if non-nil, zero value otherwise.

### GetConsolePortOk

`func (o *CableToCableTermination) GetConsolePortOk() (*ApprovalWorkflowUser, bool)`

GetConsolePortOk returns a tuple with the ConsolePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsolePort

`func (o *CableToCableTermination) SetConsolePort(v ApprovalWorkflowUser)`

SetConsolePort sets ConsolePort field to given value.

### HasConsolePort

`func (o *CableToCableTermination) HasConsolePort() bool`

HasConsolePort returns a boolean if a field has been set.

### SetConsolePortNil

`func (o *CableToCableTermination) SetConsolePortNil(b bool)`

 SetConsolePortNil sets the value for ConsolePort to be an explicit nil

### UnsetConsolePort
`func (o *CableToCableTermination) UnsetConsolePort()`

UnsetConsolePort ensures that no value is present for ConsolePort, not even an explicit nil
### GetConsoleServerPort

`func (o *CableToCableTermination) GetConsoleServerPort() ApprovalWorkflowUser`

GetConsoleServerPort returns the ConsoleServerPort field if non-nil, zero value otherwise.

### GetConsoleServerPortOk

`func (o *CableToCableTermination) GetConsoleServerPortOk() (*ApprovalWorkflowUser, bool)`

GetConsoleServerPortOk returns a tuple with the ConsoleServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleServerPort

`func (o *CableToCableTermination) SetConsoleServerPort(v ApprovalWorkflowUser)`

SetConsoleServerPort sets ConsoleServerPort field to given value.

### HasConsoleServerPort

`func (o *CableToCableTermination) HasConsoleServerPort() bool`

HasConsoleServerPort returns a boolean if a field has been set.

### SetConsoleServerPortNil

`func (o *CableToCableTermination) SetConsoleServerPortNil(b bool)`

 SetConsoleServerPortNil sets the value for ConsoleServerPort to be an explicit nil

### UnsetConsoleServerPort
`func (o *CableToCableTermination) UnsetConsoleServerPort()`

UnsetConsoleServerPort ensures that no value is present for ConsoleServerPort, not even an explicit nil
### GetFrontPort

`func (o *CableToCableTermination) GetFrontPort() ApprovalWorkflowUser`

GetFrontPort returns the FrontPort field if non-nil, zero value otherwise.

### GetFrontPortOk

`func (o *CableToCableTermination) GetFrontPortOk() (*ApprovalWorkflowUser, bool)`

GetFrontPortOk returns a tuple with the FrontPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrontPort

`func (o *CableToCableTermination) SetFrontPort(v ApprovalWorkflowUser)`

SetFrontPort sets FrontPort field to given value.

### HasFrontPort

`func (o *CableToCableTermination) HasFrontPort() bool`

HasFrontPort returns a boolean if a field has been set.

### SetFrontPortNil

`func (o *CableToCableTermination) SetFrontPortNil(b bool)`

 SetFrontPortNil sets the value for FrontPort to be an explicit nil

### UnsetFrontPort
`func (o *CableToCableTermination) UnsetFrontPort()`

UnsetFrontPort ensures that no value is present for FrontPort, not even an explicit nil
### GetInterface

`func (o *CableToCableTermination) GetInterface() ApprovalWorkflowUser`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *CableToCableTermination) GetInterfaceOk() (*ApprovalWorkflowUser, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *CableToCableTermination) SetInterface(v ApprovalWorkflowUser)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *CableToCableTermination) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### SetInterfaceNil

`func (o *CableToCableTermination) SetInterfaceNil(b bool)`

 SetInterfaceNil sets the value for Interface to be an explicit nil

### UnsetInterface
`func (o *CableToCableTermination) UnsetInterface()`

UnsetInterface ensures that no value is present for Interface, not even an explicit nil
### GetPowerFeed

`func (o *CableToCableTermination) GetPowerFeed() ApprovalWorkflowUser`

GetPowerFeed returns the PowerFeed field if non-nil, zero value otherwise.

### GetPowerFeedOk

`func (o *CableToCableTermination) GetPowerFeedOk() (*ApprovalWorkflowUser, bool)`

GetPowerFeedOk returns a tuple with the PowerFeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerFeed

`func (o *CableToCableTermination) SetPowerFeed(v ApprovalWorkflowUser)`

SetPowerFeed sets PowerFeed field to given value.

### HasPowerFeed

`func (o *CableToCableTermination) HasPowerFeed() bool`

HasPowerFeed returns a boolean if a field has been set.

### SetPowerFeedNil

`func (o *CableToCableTermination) SetPowerFeedNil(b bool)`

 SetPowerFeedNil sets the value for PowerFeed to be an explicit nil

### UnsetPowerFeed
`func (o *CableToCableTermination) UnsetPowerFeed()`

UnsetPowerFeed ensures that no value is present for PowerFeed, not even an explicit nil
### GetPowerOutlet

`func (o *CableToCableTermination) GetPowerOutlet() ApprovalWorkflowUser`

GetPowerOutlet returns the PowerOutlet field if non-nil, zero value otherwise.

### GetPowerOutletOk

`func (o *CableToCableTermination) GetPowerOutletOk() (*ApprovalWorkflowUser, bool)`

GetPowerOutletOk returns a tuple with the PowerOutlet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOutlet

`func (o *CableToCableTermination) SetPowerOutlet(v ApprovalWorkflowUser)`

SetPowerOutlet sets PowerOutlet field to given value.

### HasPowerOutlet

`func (o *CableToCableTermination) HasPowerOutlet() bool`

HasPowerOutlet returns a boolean if a field has been set.

### SetPowerOutletNil

`func (o *CableToCableTermination) SetPowerOutletNil(b bool)`

 SetPowerOutletNil sets the value for PowerOutlet to be an explicit nil

### UnsetPowerOutlet
`func (o *CableToCableTermination) UnsetPowerOutlet()`

UnsetPowerOutlet ensures that no value is present for PowerOutlet, not even an explicit nil
### GetPowerPort

`func (o *CableToCableTermination) GetPowerPort() ApprovalWorkflowUser`

GetPowerPort returns the PowerPort field if non-nil, zero value otherwise.

### GetPowerPortOk

`func (o *CableToCableTermination) GetPowerPortOk() (*ApprovalWorkflowUser, bool)`

GetPowerPortOk returns a tuple with the PowerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerPort

`func (o *CableToCableTermination) SetPowerPort(v ApprovalWorkflowUser)`

SetPowerPort sets PowerPort field to given value.

### HasPowerPort

`func (o *CableToCableTermination) HasPowerPort() bool`

HasPowerPort returns a boolean if a field has been set.

### SetPowerPortNil

`func (o *CableToCableTermination) SetPowerPortNil(b bool)`

 SetPowerPortNil sets the value for PowerPort to be an explicit nil

### UnsetPowerPort
`func (o *CableToCableTermination) UnsetPowerPort()`

UnsetPowerPort ensures that no value is present for PowerPort, not even an explicit nil
### GetRearPort

`func (o *CableToCableTermination) GetRearPort() ApprovalWorkflowUser`

GetRearPort returns the RearPort field if non-nil, zero value otherwise.

### GetRearPortOk

`func (o *CableToCableTermination) GetRearPortOk() (*ApprovalWorkflowUser, bool)`

GetRearPortOk returns a tuple with the RearPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRearPort

`func (o *CableToCableTermination) SetRearPort(v ApprovalWorkflowUser)`

SetRearPort sets RearPort field to given value.

### HasRearPort

`func (o *CableToCableTermination) HasRearPort() bool`

HasRearPort returns a boolean if a field has been set.

### SetRearPortNil

`func (o *CableToCableTermination) SetRearPortNil(b bool)`

 SetRearPortNil sets the value for RearPort to be an explicit nil

### UnsetRearPort
`func (o *CableToCableTermination) UnsetRearPort()`

UnsetRearPort ensures that no value is present for RearPort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


