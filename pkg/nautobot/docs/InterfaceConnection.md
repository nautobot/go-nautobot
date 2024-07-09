# InterfaceConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**InterfaceA** | [**Interface**](Interface.md) |  | [readonly] 
**InterfaceB** | [**Interface**](Interface.md) |  | 
**ConnectedEndpointReachable** | **NullableBool** |  | [readonly] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 

## Methods

### NewInterfaceConnection

`func NewInterfaceConnection(id string, objectType string, display string, interfaceA Interface, interfaceB Interface, connectedEndpointReachable NullableBool, created NullableTime, lastUpdated NullableTime, ) *InterfaceConnection`

NewInterfaceConnection instantiates a new InterfaceConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterfaceConnectionWithDefaults

`func NewInterfaceConnectionWithDefaults() *InterfaceConnection`

NewInterfaceConnectionWithDefaults instantiates a new InterfaceConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InterfaceConnection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InterfaceConnection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InterfaceConnection) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *InterfaceConnection) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *InterfaceConnection) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *InterfaceConnection) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *InterfaceConnection) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *InterfaceConnection) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *InterfaceConnection) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetInterfaceA

`func (o *InterfaceConnection) GetInterfaceA() Interface`

GetInterfaceA returns the InterfaceA field if non-nil, zero value otherwise.

### GetInterfaceAOk

`func (o *InterfaceConnection) GetInterfaceAOk() (*Interface, bool)`

GetInterfaceAOk returns a tuple with the InterfaceA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceA

`func (o *InterfaceConnection) SetInterfaceA(v Interface)`

SetInterfaceA sets InterfaceA field to given value.


### GetInterfaceB

`func (o *InterfaceConnection) GetInterfaceB() Interface`

GetInterfaceB returns the InterfaceB field if non-nil, zero value otherwise.

### GetInterfaceBOk

`func (o *InterfaceConnection) GetInterfaceBOk() (*Interface, bool)`

GetInterfaceBOk returns a tuple with the InterfaceB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceB

`func (o *InterfaceConnection) SetInterfaceB(v Interface)`

SetInterfaceB sets InterfaceB field to given value.


### GetConnectedEndpointReachable

`func (o *InterfaceConnection) GetConnectedEndpointReachable() bool`

GetConnectedEndpointReachable returns the ConnectedEndpointReachable field if non-nil, zero value otherwise.

### GetConnectedEndpointReachableOk

`func (o *InterfaceConnection) GetConnectedEndpointReachableOk() (*bool, bool)`

GetConnectedEndpointReachableOk returns a tuple with the ConnectedEndpointReachable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedEndpointReachable

`func (o *InterfaceConnection) SetConnectedEndpointReachable(v bool)`

SetConnectedEndpointReachable sets ConnectedEndpointReachable field to given value.


### SetConnectedEndpointReachableNil

`func (o *InterfaceConnection) SetConnectedEndpointReachableNil(b bool)`

 SetConnectedEndpointReachableNil sets the value for ConnectedEndpointReachable to be an explicit nil

### UnsetConnectedEndpointReachable
`func (o *InterfaceConnection) UnsetConnectedEndpointReachable()`

UnsetConnectedEndpointReachable ensures that no value is present for ConnectedEndpointReachable, not even an explicit nil
### GetCreated

`func (o *InterfaceConnection) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *InterfaceConnection) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *InterfaceConnection) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *InterfaceConnection) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *InterfaceConnection) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *InterfaceConnection) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *InterfaceConnection) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *InterfaceConnection) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *InterfaceConnection) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *InterfaceConnection) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


