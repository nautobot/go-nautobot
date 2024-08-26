# ConsoleServerPort

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
**Type** | Pointer to [**ConsolePortType**](ConsolePortType.md) |  | [optional] 
**Name** | **string** |  | 
**Label** | Pointer to **string** | Physical label | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Device** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Cable** | [**NullableCircuitCircuitTerminationA**](CircuitCircuitTerminationA.md) |  | 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**NotesUrl** | **string** |  | [readonly] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewConsoleServerPort

`func NewConsoleServerPort(id string, objectType string, display string, url string, naturalSlug string, cablePeerType NullableString, cablePeer NullableCableTermination, connectedEndpointType NullableString, connectedEndpoint NullablePathEndpoint, connectedEndpointReachable NullableBool, name string, device BulkWritableCableRequestStatus, cable NullableCircuitCircuitTerminationA, created NullableTime, lastUpdated NullableTime, notesUrl string, ) *ConsoleServerPort`

NewConsoleServerPort instantiates a new ConsoleServerPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsoleServerPortWithDefaults

`func NewConsoleServerPortWithDefaults() *ConsoleServerPort`

NewConsoleServerPortWithDefaults instantiates a new ConsoleServerPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConsoleServerPort) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConsoleServerPort) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConsoleServerPort) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *ConsoleServerPort) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ConsoleServerPort) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ConsoleServerPort) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *ConsoleServerPort) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *ConsoleServerPort) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *ConsoleServerPort) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *ConsoleServerPort) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ConsoleServerPort) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ConsoleServerPort) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *ConsoleServerPort) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *ConsoleServerPort) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *ConsoleServerPort) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetCablePeerType

`func (o *ConsoleServerPort) GetCablePeerType() string`

GetCablePeerType returns the CablePeerType field if non-nil, zero value otherwise.

### GetCablePeerTypeOk

`func (o *ConsoleServerPort) GetCablePeerTypeOk() (*string, bool)`

GetCablePeerTypeOk returns a tuple with the CablePeerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCablePeerType

`func (o *ConsoleServerPort) SetCablePeerType(v string)`

SetCablePeerType sets CablePeerType field to given value.


### SetCablePeerTypeNil

`func (o *ConsoleServerPort) SetCablePeerTypeNil(b bool)`

 SetCablePeerTypeNil sets the value for CablePeerType to be an explicit nil

### UnsetCablePeerType
`func (o *ConsoleServerPort) UnsetCablePeerType()`

UnsetCablePeerType ensures that no value is present for CablePeerType, not even an explicit nil
### GetCablePeer

`func (o *ConsoleServerPort) GetCablePeer() CableTermination`

GetCablePeer returns the CablePeer field if non-nil, zero value otherwise.

### GetCablePeerOk

`func (o *ConsoleServerPort) GetCablePeerOk() (*CableTermination, bool)`

GetCablePeerOk returns a tuple with the CablePeer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCablePeer

`func (o *ConsoleServerPort) SetCablePeer(v CableTermination)`

SetCablePeer sets CablePeer field to given value.


### SetCablePeerNil

`func (o *ConsoleServerPort) SetCablePeerNil(b bool)`

 SetCablePeerNil sets the value for CablePeer to be an explicit nil

### UnsetCablePeer
`func (o *ConsoleServerPort) UnsetCablePeer()`

UnsetCablePeer ensures that no value is present for CablePeer, not even an explicit nil
### GetConnectedEndpointType

`func (o *ConsoleServerPort) GetConnectedEndpointType() string`

GetConnectedEndpointType returns the ConnectedEndpointType field if non-nil, zero value otherwise.

### GetConnectedEndpointTypeOk

`func (o *ConsoleServerPort) GetConnectedEndpointTypeOk() (*string, bool)`

GetConnectedEndpointTypeOk returns a tuple with the ConnectedEndpointType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedEndpointType

`func (o *ConsoleServerPort) SetConnectedEndpointType(v string)`

SetConnectedEndpointType sets ConnectedEndpointType field to given value.


### SetConnectedEndpointTypeNil

`func (o *ConsoleServerPort) SetConnectedEndpointTypeNil(b bool)`

 SetConnectedEndpointTypeNil sets the value for ConnectedEndpointType to be an explicit nil

### UnsetConnectedEndpointType
`func (o *ConsoleServerPort) UnsetConnectedEndpointType()`

UnsetConnectedEndpointType ensures that no value is present for ConnectedEndpointType, not even an explicit nil
### GetConnectedEndpoint

`func (o *ConsoleServerPort) GetConnectedEndpoint() PathEndpoint`

GetConnectedEndpoint returns the ConnectedEndpoint field if non-nil, zero value otherwise.

### GetConnectedEndpointOk

`func (o *ConsoleServerPort) GetConnectedEndpointOk() (*PathEndpoint, bool)`

GetConnectedEndpointOk returns a tuple with the ConnectedEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedEndpoint

`func (o *ConsoleServerPort) SetConnectedEndpoint(v PathEndpoint)`

SetConnectedEndpoint sets ConnectedEndpoint field to given value.


### SetConnectedEndpointNil

`func (o *ConsoleServerPort) SetConnectedEndpointNil(b bool)`

 SetConnectedEndpointNil sets the value for ConnectedEndpoint to be an explicit nil

### UnsetConnectedEndpoint
`func (o *ConsoleServerPort) UnsetConnectedEndpoint()`

UnsetConnectedEndpoint ensures that no value is present for ConnectedEndpoint, not even an explicit nil
### GetConnectedEndpointReachable

`func (o *ConsoleServerPort) GetConnectedEndpointReachable() bool`

GetConnectedEndpointReachable returns the ConnectedEndpointReachable field if non-nil, zero value otherwise.

### GetConnectedEndpointReachableOk

`func (o *ConsoleServerPort) GetConnectedEndpointReachableOk() (*bool, bool)`

GetConnectedEndpointReachableOk returns a tuple with the ConnectedEndpointReachable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedEndpointReachable

`func (o *ConsoleServerPort) SetConnectedEndpointReachable(v bool)`

SetConnectedEndpointReachable sets ConnectedEndpointReachable field to given value.


### SetConnectedEndpointReachableNil

`func (o *ConsoleServerPort) SetConnectedEndpointReachableNil(b bool)`

 SetConnectedEndpointReachableNil sets the value for ConnectedEndpointReachable to be an explicit nil

### UnsetConnectedEndpointReachable
`func (o *ConsoleServerPort) UnsetConnectedEndpointReachable()`

UnsetConnectedEndpointReachable ensures that no value is present for ConnectedEndpointReachable, not even an explicit nil
### GetType

`func (o *ConsoleServerPort) GetType() ConsolePortType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConsoleServerPort) GetTypeOk() (*ConsolePortType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConsoleServerPort) SetType(v ConsolePortType)`

SetType sets Type field to given value.

### HasType

`func (o *ConsoleServerPort) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *ConsoleServerPort) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConsoleServerPort) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConsoleServerPort) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *ConsoleServerPort) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ConsoleServerPort) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ConsoleServerPort) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ConsoleServerPort) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDescription

`func (o *ConsoleServerPort) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConsoleServerPort) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConsoleServerPort) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConsoleServerPort) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDevice

`func (o *ConsoleServerPort) GetDevice() BulkWritableCableRequestStatus`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *ConsoleServerPort) GetDeviceOk() (*BulkWritableCableRequestStatus, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *ConsoleServerPort) SetDevice(v BulkWritableCableRequestStatus)`

SetDevice sets Device field to given value.


### GetCable

`func (o *ConsoleServerPort) GetCable() CircuitCircuitTerminationA`

GetCable returns the Cable field if non-nil, zero value otherwise.

### GetCableOk

`func (o *ConsoleServerPort) GetCableOk() (*CircuitCircuitTerminationA, bool)`

GetCableOk returns a tuple with the Cable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCable

`func (o *ConsoleServerPort) SetCable(v CircuitCircuitTerminationA)`

SetCable sets Cable field to given value.


### SetCableNil

`func (o *ConsoleServerPort) SetCableNil(b bool)`

 SetCableNil sets the value for Cable to be an explicit nil

### UnsetCable
`func (o *ConsoleServerPort) UnsetCable()`

UnsetCable ensures that no value is present for Cable, not even an explicit nil
### GetCreated

`func (o *ConsoleServerPort) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ConsoleServerPort) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ConsoleServerPort) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *ConsoleServerPort) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *ConsoleServerPort) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *ConsoleServerPort) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ConsoleServerPort) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ConsoleServerPort) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *ConsoleServerPort) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *ConsoleServerPort) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetTags

`func (o *ConsoleServerPort) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ConsoleServerPort) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ConsoleServerPort) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ConsoleServerPort) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetNotesUrl

`func (o *ConsoleServerPort) GetNotesUrl() string`

GetNotesUrl returns the NotesUrl field if non-nil, zero value otherwise.

### GetNotesUrlOk

`func (o *ConsoleServerPort) GetNotesUrlOk() (*string, bool)`

GetNotesUrlOk returns a tuple with the NotesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesUrl

`func (o *ConsoleServerPort) SetNotesUrl(v string)`

SetNotesUrl sets NotesUrl field to given value.


### GetCustomFields

`func (o *ConsoleServerPort) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ConsoleServerPort) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ConsoleServerPort) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ConsoleServerPort) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


