# CircuitTermination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**CablePeerType** | **NullableString** |  | [readonly] 
**CablePeer** | [**NullableCableTermination**](CableTermination.md) |  | [readonly] 
**ConnectedEndpointType** | **NullableString** |  | [readonly] 
**ConnectedEndpoint** | [**NullablePathEndpoint**](PathEndpoint.md) |  | [readonly] 
**ConnectedEndpointReachable** | **NullableBool** |  | [readonly] 
**TermSide** | [**TermSideEnum**](TermSideEnum.md) |  | 
**PortSpeed** | Pointer to **NullableInt32** |  | [optional] 
**UpstreamSpeed** | Pointer to **NullableInt32** | Upstream speed, if different from port speed | [optional] 
**XconnectId** | Pointer to **string** |  | [optional] 
**PpInfo** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Cable** | [**NullableCircuitCircuitTerminationA**](CircuitCircuitTerminationA.md) |  | 
**Circuit** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Location** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**ProviderNetwork** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**NotesUrl** | **string** |  | [readonly] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCircuitTermination

`func NewCircuitTermination(id string, objectType string, display string, url string, naturalSlug string, cablePeerType NullableString, cablePeer NullableCableTermination, connectedEndpointType NullableString, connectedEndpoint NullablePathEndpoint, connectedEndpointReachable NullableBool, termSide TermSideEnum, cable NullableCircuitCircuitTerminationA, circuit BulkWritableCableRequestStatus, created NullableTime, lastUpdated NullableTime, notesUrl string, ) *CircuitTermination`

NewCircuitTermination instantiates a new CircuitTermination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCircuitTerminationWithDefaults

`func NewCircuitTerminationWithDefaults() *CircuitTermination`

NewCircuitTerminationWithDefaults instantiates a new CircuitTermination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CircuitTermination) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CircuitTermination) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CircuitTermination) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *CircuitTermination) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CircuitTermination) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CircuitTermination) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *CircuitTermination) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *CircuitTermination) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *CircuitTermination) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *CircuitTermination) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CircuitTermination) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CircuitTermination) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *CircuitTermination) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *CircuitTermination) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *CircuitTermination) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetCablePeerType

`func (o *CircuitTermination) GetCablePeerType() string`

GetCablePeerType returns the CablePeerType field if non-nil, zero value otherwise.

### GetCablePeerTypeOk

`func (o *CircuitTermination) GetCablePeerTypeOk() (*string, bool)`

GetCablePeerTypeOk returns a tuple with the CablePeerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCablePeerType

`func (o *CircuitTermination) SetCablePeerType(v string)`

SetCablePeerType sets CablePeerType field to given value.


### SetCablePeerTypeNil

`func (o *CircuitTermination) SetCablePeerTypeNil(b bool)`

 SetCablePeerTypeNil sets the value for CablePeerType to be an explicit nil

### UnsetCablePeerType
`func (o *CircuitTermination) UnsetCablePeerType()`

UnsetCablePeerType ensures that no value is present for CablePeerType, not even an explicit nil
### GetCablePeer

`func (o *CircuitTermination) GetCablePeer() CableTermination`

GetCablePeer returns the CablePeer field if non-nil, zero value otherwise.

### GetCablePeerOk

`func (o *CircuitTermination) GetCablePeerOk() (*CableTermination, bool)`

GetCablePeerOk returns a tuple with the CablePeer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCablePeer

`func (o *CircuitTermination) SetCablePeer(v CableTermination)`

SetCablePeer sets CablePeer field to given value.


### SetCablePeerNil

`func (o *CircuitTermination) SetCablePeerNil(b bool)`

 SetCablePeerNil sets the value for CablePeer to be an explicit nil

### UnsetCablePeer
`func (o *CircuitTermination) UnsetCablePeer()`

UnsetCablePeer ensures that no value is present for CablePeer, not even an explicit nil
### GetConnectedEndpointType

`func (o *CircuitTermination) GetConnectedEndpointType() string`

GetConnectedEndpointType returns the ConnectedEndpointType field if non-nil, zero value otherwise.

### GetConnectedEndpointTypeOk

`func (o *CircuitTermination) GetConnectedEndpointTypeOk() (*string, bool)`

GetConnectedEndpointTypeOk returns a tuple with the ConnectedEndpointType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedEndpointType

`func (o *CircuitTermination) SetConnectedEndpointType(v string)`

SetConnectedEndpointType sets ConnectedEndpointType field to given value.


### SetConnectedEndpointTypeNil

`func (o *CircuitTermination) SetConnectedEndpointTypeNil(b bool)`

 SetConnectedEndpointTypeNil sets the value for ConnectedEndpointType to be an explicit nil

### UnsetConnectedEndpointType
`func (o *CircuitTermination) UnsetConnectedEndpointType()`

UnsetConnectedEndpointType ensures that no value is present for ConnectedEndpointType, not even an explicit nil
### GetConnectedEndpoint

`func (o *CircuitTermination) GetConnectedEndpoint() PathEndpoint`

GetConnectedEndpoint returns the ConnectedEndpoint field if non-nil, zero value otherwise.

### GetConnectedEndpointOk

`func (o *CircuitTermination) GetConnectedEndpointOk() (*PathEndpoint, bool)`

GetConnectedEndpointOk returns a tuple with the ConnectedEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedEndpoint

`func (o *CircuitTermination) SetConnectedEndpoint(v PathEndpoint)`

SetConnectedEndpoint sets ConnectedEndpoint field to given value.


### SetConnectedEndpointNil

`func (o *CircuitTermination) SetConnectedEndpointNil(b bool)`

 SetConnectedEndpointNil sets the value for ConnectedEndpoint to be an explicit nil

### UnsetConnectedEndpoint
`func (o *CircuitTermination) UnsetConnectedEndpoint()`

UnsetConnectedEndpoint ensures that no value is present for ConnectedEndpoint, not even an explicit nil
### GetConnectedEndpointReachable

`func (o *CircuitTermination) GetConnectedEndpointReachable() bool`

GetConnectedEndpointReachable returns the ConnectedEndpointReachable field if non-nil, zero value otherwise.

### GetConnectedEndpointReachableOk

`func (o *CircuitTermination) GetConnectedEndpointReachableOk() (*bool, bool)`

GetConnectedEndpointReachableOk returns a tuple with the ConnectedEndpointReachable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedEndpointReachable

`func (o *CircuitTermination) SetConnectedEndpointReachable(v bool)`

SetConnectedEndpointReachable sets ConnectedEndpointReachable field to given value.


### SetConnectedEndpointReachableNil

`func (o *CircuitTermination) SetConnectedEndpointReachableNil(b bool)`

 SetConnectedEndpointReachableNil sets the value for ConnectedEndpointReachable to be an explicit nil

### UnsetConnectedEndpointReachable
`func (o *CircuitTermination) UnsetConnectedEndpointReachable()`

UnsetConnectedEndpointReachable ensures that no value is present for ConnectedEndpointReachable, not even an explicit nil
### GetTermSide

`func (o *CircuitTermination) GetTermSide() TermSideEnum`

GetTermSide returns the TermSide field if non-nil, zero value otherwise.

### GetTermSideOk

`func (o *CircuitTermination) GetTermSideOk() (*TermSideEnum, bool)`

GetTermSideOk returns a tuple with the TermSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermSide

`func (o *CircuitTermination) SetTermSide(v TermSideEnum)`

SetTermSide sets TermSide field to given value.


### GetPortSpeed

`func (o *CircuitTermination) GetPortSpeed() int32`

GetPortSpeed returns the PortSpeed field if non-nil, zero value otherwise.

### GetPortSpeedOk

`func (o *CircuitTermination) GetPortSpeedOk() (*int32, bool)`

GetPortSpeedOk returns a tuple with the PortSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortSpeed

`func (o *CircuitTermination) SetPortSpeed(v int32)`

SetPortSpeed sets PortSpeed field to given value.

### HasPortSpeed

`func (o *CircuitTermination) HasPortSpeed() bool`

HasPortSpeed returns a boolean if a field has been set.

### SetPortSpeedNil

`func (o *CircuitTermination) SetPortSpeedNil(b bool)`

 SetPortSpeedNil sets the value for PortSpeed to be an explicit nil

### UnsetPortSpeed
`func (o *CircuitTermination) UnsetPortSpeed()`

UnsetPortSpeed ensures that no value is present for PortSpeed, not even an explicit nil
### GetUpstreamSpeed

`func (o *CircuitTermination) GetUpstreamSpeed() int32`

GetUpstreamSpeed returns the UpstreamSpeed field if non-nil, zero value otherwise.

### GetUpstreamSpeedOk

`func (o *CircuitTermination) GetUpstreamSpeedOk() (*int32, bool)`

GetUpstreamSpeedOk returns a tuple with the UpstreamSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamSpeed

`func (o *CircuitTermination) SetUpstreamSpeed(v int32)`

SetUpstreamSpeed sets UpstreamSpeed field to given value.

### HasUpstreamSpeed

`func (o *CircuitTermination) HasUpstreamSpeed() bool`

HasUpstreamSpeed returns a boolean if a field has been set.

### SetUpstreamSpeedNil

`func (o *CircuitTermination) SetUpstreamSpeedNil(b bool)`

 SetUpstreamSpeedNil sets the value for UpstreamSpeed to be an explicit nil

### UnsetUpstreamSpeed
`func (o *CircuitTermination) UnsetUpstreamSpeed()`

UnsetUpstreamSpeed ensures that no value is present for UpstreamSpeed, not even an explicit nil
### GetXconnectId

`func (o *CircuitTermination) GetXconnectId() string`

GetXconnectId returns the XconnectId field if non-nil, zero value otherwise.

### GetXconnectIdOk

`func (o *CircuitTermination) GetXconnectIdOk() (*string, bool)`

GetXconnectIdOk returns a tuple with the XconnectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXconnectId

`func (o *CircuitTermination) SetXconnectId(v string)`

SetXconnectId sets XconnectId field to given value.

### HasXconnectId

`func (o *CircuitTermination) HasXconnectId() bool`

HasXconnectId returns a boolean if a field has been set.

### GetPpInfo

`func (o *CircuitTermination) GetPpInfo() string`

GetPpInfo returns the PpInfo field if non-nil, zero value otherwise.

### GetPpInfoOk

`func (o *CircuitTermination) GetPpInfoOk() (*string, bool)`

GetPpInfoOk returns a tuple with the PpInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpInfo

`func (o *CircuitTermination) SetPpInfo(v string)`

SetPpInfo sets PpInfo field to given value.

### HasPpInfo

`func (o *CircuitTermination) HasPpInfo() bool`

HasPpInfo returns a boolean if a field has been set.

### GetDescription

`func (o *CircuitTermination) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CircuitTermination) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CircuitTermination) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CircuitTermination) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCable

`func (o *CircuitTermination) GetCable() CircuitCircuitTerminationA`

GetCable returns the Cable field if non-nil, zero value otherwise.

### GetCableOk

`func (o *CircuitTermination) GetCableOk() (*CircuitCircuitTerminationA, bool)`

GetCableOk returns a tuple with the Cable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCable

`func (o *CircuitTermination) SetCable(v CircuitCircuitTerminationA)`

SetCable sets Cable field to given value.


### SetCableNil

`func (o *CircuitTermination) SetCableNil(b bool)`

 SetCableNil sets the value for Cable to be an explicit nil

### UnsetCable
`func (o *CircuitTermination) UnsetCable()`

UnsetCable ensures that no value is present for Cable, not even an explicit nil
### GetCircuit

`func (o *CircuitTermination) GetCircuit() BulkWritableCableRequestStatus`

GetCircuit returns the Circuit field if non-nil, zero value otherwise.

### GetCircuitOk

`func (o *CircuitTermination) GetCircuitOk() (*BulkWritableCableRequestStatus, bool)`

GetCircuitOk returns a tuple with the Circuit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuit

`func (o *CircuitTermination) SetCircuit(v BulkWritableCableRequestStatus)`

SetCircuit sets Circuit field to given value.


### GetLocation

`func (o *CircuitTermination) GetLocation() BulkWritableCircuitRequestTenant`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CircuitTermination) GetLocationOk() (*BulkWritableCircuitRequestTenant, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *CircuitTermination) SetLocation(v BulkWritableCircuitRequestTenant)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *CircuitTermination) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *CircuitTermination) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *CircuitTermination) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetProviderNetwork

`func (o *CircuitTermination) GetProviderNetwork() BulkWritableCircuitRequestTenant`

GetProviderNetwork returns the ProviderNetwork field if non-nil, zero value otherwise.

### GetProviderNetworkOk

`func (o *CircuitTermination) GetProviderNetworkOk() (*BulkWritableCircuitRequestTenant, bool)`

GetProviderNetworkOk returns a tuple with the ProviderNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderNetwork

`func (o *CircuitTermination) SetProviderNetwork(v BulkWritableCircuitRequestTenant)`

SetProviderNetwork sets ProviderNetwork field to given value.

### HasProviderNetwork

`func (o *CircuitTermination) HasProviderNetwork() bool`

HasProviderNetwork returns a boolean if a field has been set.

### SetProviderNetworkNil

`func (o *CircuitTermination) SetProviderNetworkNil(b bool)`

 SetProviderNetworkNil sets the value for ProviderNetwork to be an explicit nil

### UnsetProviderNetwork
`func (o *CircuitTermination) UnsetProviderNetwork()`

UnsetProviderNetwork ensures that no value is present for ProviderNetwork, not even an explicit nil
### GetCreated

`func (o *CircuitTermination) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CircuitTermination) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CircuitTermination) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *CircuitTermination) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *CircuitTermination) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *CircuitTermination) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *CircuitTermination) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *CircuitTermination) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *CircuitTermination) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *CircuitTermination) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetNotesUrl

`func (o *CircuitTermination) GetNotesUrl() string`

GetNotesUrl returns the NotesUrl field if non-nil, zero value otherwise.

### GetNotesUrlOk

`func (o *CircuitTermination) GetNotesUrlOk() (*string, bool)`

GetNotesUrlOk returns a tuple with the NotesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesUrl

`func (o *CircuitTermination) SetNotesUrl(v string)`

SetNotesUrl sets NotesUrl field to given value.


### GetCustomFields

`func (o *CircuitTermination) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CircuitTermination) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CircuitTermination) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CircuitTermination) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


