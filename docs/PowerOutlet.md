# PowerOutlet

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
**Type** | Pointer to [**PowerOutletType**](PowerOutletType.md) |  | [optional] 
**FeedLeg** | Pointer to [**PowerOutletFeedLeg**](PowerOutletFeedLeg.md) |  | [optional] 
**Name** | **string** |  | 
**Label** | Pointer to **string** | Physical label | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Device** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Module** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Cable** | [**NullableCircuitCircuitTerminationA**](CircuitCircuitTerminationA.md) |  | 
**PowerPort** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**NotesUrl** | **string** |  | [readonly] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewPowerOutlet

`func NewPowerOutlet(id string, objectType string, display string, url string, naturalSlug string, cablePeerType NullableString, cablePeer NullableCableTermination, connectedEndpointType NullableString, connectedEndpoint NullablePathEndpoint, connectedEndpointReachable NullableBool, name string, cable NullableCircuitCircuitTerminationA, created NullableTime, lastUpdated NullableTime, notesUrl string, ) *PowerOutlet`

NewPowerOutlet instantiates a new PowerOutlet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPowerOutletWithDefaults

`func NewPowerOutletWithDefaults() *PowerOutlet`

NewPowerOutletWithDefaults instantiates a new PowerOutlet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PowerOutlet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PowerOutlet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PowerOutlet) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *PowerOutlet) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PowerOutlet) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PowerOutlet) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *PowerOutlet) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *PowerOutlet) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *PowerOutlet) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *PowerOutlet) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PowerOutlet) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PowerOutlet) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *PowerOutlet) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *PowerOutlet) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *PowerOutlet) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetCablePeerType

`func (o *PowerOutlet) GetCablePeerType() string`

GetCablePeerType returns the CablePeerType field if non-nil, zero value otherwise.

### GetCablePeerTypeOk

`func (o *PowerOutlet) GetCablePeerTypeOk() (*string, bool)`

GetCablePeerTypeOk returns a tuple with the CablePeerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCablePeerType

`func (o *PowerOutlet) SetCablePeerType(v string)`

SetCablePeerType sets CablePeerType field to given value.


### SetCablePeerTypeNil

`func (o *PowerOutlet) SetCablePeerTypeNil(b bool)`

 SetCablePeerTypeNil sets the value for CablePeerType to be an explicit nil

### UnsetCablePeerType
`func (o *PowerOutlet) UnsetCablePeerType()`

UnsetCablePeerType ensures that no value is present for CablePeerType, not even an explicit nil
### GetCablePeer

`func (o *PowerOutlet) GetCablePeer() CableTermination`

GetCablePeer returns the CablePeer field if non-nil, zero value otherwise.

### GetCablePeerOk

`func (o *PowerOutlet) GetCablePeerOk() (*CableTermination, bool)`

GetCablePeerOk returns a tuple with the CablePeer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCablePeer

`func (o *PowerOutlet) SetCablePeer(v CableTermination)`

SetCablePeer sets CablePeer field to given value.


### SetCablePeerNil

`func (o *PowerOutlet) SetCablePeerNil(b bool)`

 SetCablePeerNil sets the value for CablePeer to be an explicit nil

### UnsetCablePeer
`func (o *PowerOutlet) UnsetCablePeer()`

UnsetCablePeer ensures that no value is present for CablePeer, not even an explicit nil
### GetConnectedEndpointType

`func (o *PowerOutlet) GetConnectedEndpointType() string`

GetConnectedEndpointType returns the ConnectedEndpointType field if non-nil, zero value otherwise.

### GetConnectedEndpointTypeOk

`func (o *PowerOutlet) GetConnectedEndpointTypeOk() (*string, bool)`

GetConnectedEndpointTypeOk returns a tuple with the ConnectedEndpointType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedEndpointType

`func (o *PowerOutlet) SetConnectedEndpointType(v string)`

SetConnectedEndpointType sets ConnectedEndpointType field to given value.


### SetConnectedEndpointTypeNil

`func (o *PowerOutlet) SetConnectedEndpointTypeNil(b bool)`

 SetConnectedEndpointTypeNil sets the value for ConnectedEndpointType to be an explicit nil

### UnsetConnectedEndpointType
`func (o *PowerOutlet) UnsetConnectedEndpointType()`

UnsetConnectedEndpointType ensures that no value is present for ConnectedEndpointType, not even an explicit nil
### GetConnectedEndpoint

`func (o *PowerOutlet) GetConnectedEndpoint() PathEndpoint`

GetConnectedEndpoint returns the ConnectedEndpoint field if non-nil, zero value otherwise.

### GetConnectedEndpointOk

`func (o *PowerOutlet) GetConnectedEndpointOk() (*PathEndpoint, bool)`

GetConnectedEndpointOk returns a tuple with the ConnectedEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedEndpoint

`func (o *PowerOutlet) SetConnectedEndpoint(v PathEndpoint)`

SetConnectedEndpoint sets ConnectedEndpoint field to given value.


### SetConnectedEndpointNil

`func (o *PowerOutlet) SetConnectedEndpointNil(b bool)`

 SetConnectedEndpointNil sets the value for ConnectedEndpoint to be an explicit nil

### UnsetConnectedEndpoint
`func (o *PowerOutlet) UnsetConnectedEndpoint()`

UnsetConnectedEndpoint ensures that no value is present for ConnectedEndpoint, not even an explicit nil
### GetConnectedEndpointReachable

`func (o *PowerOutlet) GetConnectedEndpointReachable() bool`

GetConnectedEndpointReachable returns the ConnectedEndpointReachable field if non-nil, zero value otherwise.

### GetConnectedEndpointReachableOk

`func (o *PowerOutlet) GetConnectedEndpointReachableOk() (*bool, bool)`

GetConnectedEndpointReachableOk returns a tuple with the ConnectedEndpointReachable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedEndpointReachable

`func (o *PowerOutlet) SetConnectedEndpointReachable(v bool)`

SetConnectedEndpointReachable sets ConnectedEndpointReachable field to given value.


### SetConnectedEndpointReachableNil

`func (o *PowerOutlet) SetConnectedEndpointReachableNil(b bool)`

 SetConnectedEndpointReachableNil sets the value for ConnectedEndpointReachable to be an explicit nil

### UnsetConnectedEndpointReachable
`func (o *PowerOutlet) UnsetConnectedEndpointReachable()`

UnsetConnectedEndpointReachable ensures that no value is present for ConnectedEndpointReachable, not even an explicit nil
### GetType

`func (o *PowerOutlet) GetType() PowerOutletType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PowerOutlet) GetTypeOk() (*PowerOutletType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PowerOutlet) SetType(v PowerOutletType)`

SetType sets Type field to given value.

### HasType

`func (o *PowerOutlet) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFeedLeg

`func (o *PowerOutlet) GetFeedLeg() PowerOutletFeedLeg`

GetFeedLeg returns the FeedLeg field if non-nil, zero value otherwise.

### GetFeedLegOk

`func (o *PowerOutlet) GetFeedLegOk() (*PowerOutletFeedLeg, bool)`

GetFeedLegOk returns a tuple with the FeedLeg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedLeg

`func (o *PowerOutlet) SetFeedLeg(v PowerOutletFeedLeg)`

SetFeedLeg sets FeedLeg field to given value.

### HasFeedLeg

`func (o *PowerOutlet) HasFeedLeg() bool`

HasFeedLeg returns a boolean if a field has been set.

### GetName

`func (o *PowerOutlet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PowerOutlet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PowerOutlet) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *PowerOutlet) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PowerOutlet) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PowerOutlet) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PowerOutlet) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDescription

`func (o *PowerOutlet) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PowerOutlet) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PowerOutlet) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PowerOutlet) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDevice

`func (o *PowerOutlet) GetDevice() BulkWritableCircuitRequestTenant`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *PowerOutlet) GetDeviceOk() (*BulkWritableCircuitRequestTenant, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *PowerOutlet) SetDevice(v BulkWritableCircuitRequestTenant)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *PowerOutlet) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### SetDeviceNil

`func (o *PowerOutlet) SetDeviceNil(b bool)`

 SetDeviceNil sets the value for Device to be an explicit nil

### UnsetDevice
`func (o *PowerOutlet) UnsetDevice()`

UnsetDevice ensures that no value is present for Device, not even an explicit nil
### GetModule

`func (o *PowerOutlet) GetModule() BulkWritableCircuitRequestTenant`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *PowerOutlet) GetModuleOk() (*BulkWritableCircuitRequestTenant, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *PowerOutlet) SetModule(v BulkWritableCircuitRequestTenant)`

SetModule sets Module field to given value.

### HasModule

`func (o *PowerOutlet) HasModule() bool`

HasModule returns a boolean if a field has been set.

### SetModuleNil

`func (o *PowerOutlet) SetModuleNil(b bool)`

 SetModuleNil sets the value for Module to be an explicit nil

### UnsetModule
`func (o *PowerOutlet) UnsetModule()`

UnsetModule ensures that no value is present for Module, not even an explicit nil
### GetCable

`func (o *PowerOutlet) GetCable() CircuitCircuitTerminationA`

GetCable returns the Cable field if non-nil, zero value otherwise.

### GetCableOk

`func (o *PowerOutlet) GetCableOk() (*CircuitCircuitTerminationA, bool)`

GetCableOk returns a tuple with the Cable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCable

`func (o *PowerOutlet) SetCable(v CircuitCircuitTerminationA)`

SetCable sets Cable field to given value.


### SetCableNil

`func (o *PowerOutlet) SetCableNil(b bool)`

 SetCableNil sets the value for Cable to be an explicit nil

### UnsetCable
`func (o *PowerOutlet) UnsetCable()`

UnsetCable ensures that no value is present for Cable, not even an explicit nil
### GetPowerPort

`func (o *PowerOutlet) GetPowerPort() BulkWritableCircuitRequestTenant`

GetPowerPort returns the PowerPort field if non-nil, zero value otherwise.

### GetPowerPortOk

`func (o *PowerOutlet) GetPowerPortOk() (*BulkWritableCircuitRequestTenant, bool)`

GetPowerPortOk returns a tuple with the PowerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerPort

`func (o *PowerOutlet) SetPowerPort(v BulkWritableCircuitRequestTenant)`

SetPowerPort sets PowerPort field to given value.

### HasPowerPort

`func (o *PowerOutlet) HasPowerPort() bool`

HasPowerPort returns a boolean if a field has been set.

### SetPowerPortNil

`func (o *PowerOutlet) SetPowerPortNil(b bool)`

 SetPowerPortNil sets the value for PowerPort to be an explicit nil

### UnsetPowerPort
`func (o *PowerOutlet) UnsetPowerPort()`

UnsetPowerPort ensures that no value is present for PowerPort, not even an explicit nil
### GetCreated

`func (o *PowerOutlet) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PowerOutlet) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PowerOutlet) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *PowerOutlet) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *PowerOutlet) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *PowerOutlet) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *PowerOutlet) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *PowerOutlet) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *PowerOutlet) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *PowerOutlet) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetNotesUrl

`func (o *PowerOutlet) GetNotesUrl() string`

GetNotesUrl returns the NotesUrl field if non-nil, zero value otherwise.

### GetNotesUrlOk

`func (o *PowerOutlet) GetNotesUrlOk() (*string, bool)`

GetNotesUrlOk returns a tuple with the NotesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesUrl

`func (o *PowerOutlet) SetNotesUrl(v string)`

SetNotesUrl sets NotesUrl field to given value.


### GetCustomFields

`func (o *PowerOutlet) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PowerOutlet) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PowerOutlet) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PowerOutlet) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetTags

`func (o *PowerOutlet) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PowerOutlet) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PowerOutlet) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PowerOutlet) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


