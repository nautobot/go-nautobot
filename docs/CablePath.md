# CablePath

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**OriginType** | **string** |  | [readonly] 
**Origin** | [**PathEndpoint**](PathEndpoint.md) |  | [readonly] 
**DestinationType** | **string** |  | [readonly] 
**Destination** | [**NullablePathEndpoint**](PathEndpoint.md) |  | [readonly] 
**Path** | [**[]CableTermination**](CableTermination.md) |  | [readonly] 
**OriginId** | **string** |  | 
**DestinationId** | Pointer to **NullableString** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**IsSplit** | Pointer to **bool** |  | [optional] 
**PeerConnector** | Pointer to **int32** |  | [optional] [default to 1]
**OriginFansOut** | Pointer to **bool** |  | [optional] 
**DestinationFansOut** | Pointer to **bool** |  | [optional] 

## Methods

### NewCablePath

`func NewCablePath(id string, originType string, origin PathEndpoint, destinationType string, destination NullablePathEndpoint, path []CableTermination, originId string, ) *CablePath`

NewCablePath instantiates a new CablePath object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCablePathWithDefaults

`func NewCablePathWithDefaults() *CablePath`

NewCablePathWithDefaults instantiates a new CablePath object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CablePath) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CablePath) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CablePath) SetId(v string)`

SetId sets Id field to given value.


### GetOriginType

`func (o *CablePath) GetOriginType() string`

GetOriginType returns the OriginType field if non-nil, zero value otherwise.

### GetOriginTypeOk

`func (o *CablePath) GetOriginTypeOk() (*string, bool)`

GetOriginTypeOk returns a tuple with the OriginType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginType

`func (o *CablePath) SetOriginType(v string)`

SetOriginType sets OriginType field to given value.


### GetOrigin

`func (o *CablePath) GetOrigin() PathEndpoint`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *CablePath) GetOriginOk() (*PathEndpoint, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *CablePath) SetOrigin(v PathEndpoint)`

SetOrigin sets Origin field to given value.


### GetDestinationType

`func (o *CablePath) GetDestinationType() string`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *CablePath) GetDestinationTypeOk() (*string, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *CablePath) SetDestinationType(v string)`

SetDestinationType sets DestinationType field to given value.


### GetDestination

`func (o *CablePath) GetDestination() PathEndpoint`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *CablePath) GetDestinationOk() (*PathEndpoint, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *CablePath) SetDestination(v PathEndpoint)`

SetDestination sets Destination field to given value.


### SetDestinationNil

`func (o *CablePath) SetDestinationNil(b bool)`

 SetDestinationNil sets the value for Destination to be an explicit nil

### UnsetDestination
`func (o *CablePath) UnsetDestination()`

UnsetDestination ensures that no value is present for Destination, not even an explicit nil
### GetPath

`func (o *CablePath) GetPath() []CableTermination`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CablePath) GetPathOk() (*[]CableTermination, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CablePath) SetPath(v []CableTermination)`

SetPath sets Path field to given value.


### GetOriginId

`func (o *CablePath) GetOriginId() string`

GetOriginId returns the OriginId field if non-nil, zero value otherwise.

### GetOriginIdOk

`func (o *CablePath) GetOriginIdOk() (*string, bool)`

GetOriginIdOk returns a tuple with the OriginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginId

`func (o *CablePath) SetOriginId(v string)`

SetOriginId sets OriginId field to given value.


### GetDestinationId

`func (o *CablePath) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *CablePath) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *CablePath) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.

### HasDestinationId

`func (o *CablePath) HasDestinationId() bool`

HasDestinationId returns a boolean if a field has been set.

### SetDestinationIdNil

`func (o *CablePath) SetDestinationIdNil(b bool)`

 SetDestinationIdNil sets the value for DestinationId to be an explicit nil

### UnsetDestinationId
`func (o *CablePath) UnsetDestinationId()`

UnsetDestinationId ensures that no value is present for DestinationId, not even an explicit nil
### GetIsActive

`func (o *CablePath) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *CablePath) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *CablePath) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *CablePath) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetIsSplit

`func (o *CablePath) GetIsSplit() bool`

GetIsSplit returns the IsSplit field if non-nil, zero value otherwise.

### GetIsSplitOk

`func (o *CablePath) GetIsSplitOk() (*bool, bool)`

GetIsSplitOk returns a tuple with the IsSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSplit

`func (o *CablePath) SetIsSplit(v bool)`

SetIsSplit sets IsSplit field to given value.

### HasIsSplit

`func (o *CablePath) HasIsSplit() bool`

HasIsSplit returns a boolean if a field has been set.

### GetPeerConnector

`func (o *CablePath) GetPeerConnector() int32`

GetPeerConnector returns the PeerConnector field if non-nil, zero value otherwise.

### GetPeerConnectorOk

`func (o *CablePath) GetPeerConnectorOk() (*int32, bool)`

GetPeerConnectorOk returns a tuple with the PeerConnector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerConnector

`func (o *CablePath) SetPeerConnector(v int32)`

SetPeerConnector sets PeerConnector field to given value.

### HasPeerConnector

`func (o *CablePath) HasPeerConnector() bool`

HasPeerConnector returns a boolean if a field has been set.

### GetOriginFansOut

`func (o *CablePath) GetOriginFansOut() bool`

GetOriginFansOut returns the OriginFansOut field if non-nil, zero value otherwise.

### GetOriginFansOutOk

`func (o *CablePath) GetOriginFansOutOk() (*bool, bool)`

GetOriginFansOutOk returns a tuple with the OriginFansOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginFansOut

`func (o *CablePath) SetOriginFansOut(v bool)`

SetOriginFansOut sets OriginFansOut field to given value.

### HasOriginFansOut

`func (o *CablePath) HasOriginFansOut() bool`

HasOriginFansOut returns a boolean if a field has been set.

### GetDestinationFansOut

`func (o *CablePath) GetDestinationFansOut() bool`

GetDestinationFansOut returns the DestinationFansOut field if non-nil, zero value otherwise.

### GetDestinationFansOutOk

`func (o *CablePath) GetDestinationFansOutOk() (*bool, bool)`

GetDestinationFansOutOk returns a tuple with the DestinationFansOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationFansOut

`func (o *CablePath) SetDestinationFansOut(v bool)`

SetDestinationFansOut sets DestinationFansOut field to given value.

### HasDestinationFansOut

`func (o *CablePath) HasDestinationFansOut() bool`

HasDestinationFansOut returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


