# PowerPort

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
**Type** | Pointer to [**PowerPortType**](PowerPortType.md) |  | [optional] 
**Name** | **string** |  | 
**Label** | Pointer to **string** | Physical label | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**MaximumDraw** | Pointer to **NullableInt32** | Maximum power draw (watts) | [optional] 
**AllocatedDraw** | Pointer to **NullableInt32** | Allocated power draw (watts) | [optional] 
**Device** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Cable** | [**NullableCircuitCircuitTerminationA**](CircuitCircuitTerminationA.md) |  | 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**NotesUrl** | **string** |  | [readonly] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewPowerPort

`func NewPowerPort(id string, objectType string, display string, url string, naturalSlug string, cablePeerType NullableString, cablePeer NullableCableTermination, connectedEndpointType NullableString, connectedEndpoint NullablePathEndpoint, connectedEndpointReachable NullableBool, name string, device BulkWritableCableRequestStatus, cable NullableCircuitCircuitTerminationA, created NullableTime, lastUpdated NullableTime, notesUrl string, ) *PowerPort`

NewPowerPort instantiates a new PowerPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPowerPortWithDefaults

`func NewPowerPortWithDefaults() *PowerPort`

NewPowerPortWithDefaults instantiates a new PowerPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PowerPort) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PowerPort) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PowerPort) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *PowerPort) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PowerPort) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PowerPort) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *PowerPort) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *PowerPort) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *PowerPort) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *PowerPort) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PowerPort) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PowerPort) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *PowerPort) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *PowerPort) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *PowerPort) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetCablePeerType

`func (o *PowerPort) GetCablePeerType() string`

GetCablePeerType returns the CablePeerType field if non-nil, zero value otherwise.

### GetCablePeerTypeOk

`func (o *PowerPort) GetCablePeerTypeOk() (*string, bool)`

GetCablePeerTypeOk returns a tuple with the CablePeerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCablePeerType

`func (o *PowerPort) SetCablePeerType(v string)`

SetCablePeerType sets CablePeerType field to given value.


### SetCablePeerTypeNil

`func (o *PowerPort) SetCablePeerTypeNil(b bool)`

 SetCablePeerTypeNil sets the value for CablePeerType to be an explicit nil

### UnsetCablePeerType
`func (o *PowerPort) UnsetCablePeerType()`

UnsetCablePeerType ensures that no value is present for CablePeerType, not even an explicit nil
### GetCablePeer

`func (o *PowerPort) GetCablePeer() CableTermination`

GetCablePeer returns the CablePeer field if non-nil, zero value otherwise.

### GetCablePeerOk

`func (o *PowerPort) GetCablePeerOk() (*CableTermination, bool)`

GetCablePeerOk returns a tuple with the CablePeer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCablePeer

`func (o *PowerPort) SetCablePeer(v CableTermination)`

SetCablePeer sets CablePeer field to given value.


### SetCablePeerNil

`func (o *PowerPort) SetCablePeerNil(b bool)`

 SetCablePeerNil sets the value for CablePeer to be an explicit nil

### UnsetCablePeer
`func (o *PowerPort) UnsetCablePeer()`

UnsetCablePeer ensures that no value is present for CablePeer, not even an explicit nil
### GetConnectedEndpointType

`func (o *PowerPort) GetConnectedEndpointType() string`

GetConnectedEndpointType returns the ConnectedEndpointType field if non-nil, zero value otherwise.

### GetConnectedEndpointTypeOk

`func (o *PowerPort) GetConnectedEndpointTypeOk() (*string, bool)`

GetConnectedEndpointTypeOk returns a tuple with the ConnectedEndpointType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedEndpointType

`func (o *PowerPort) SetConnectedEndpointType(v string)`

SetConnectedEndpointType sets ConnectedEndpointType field to given value.


### SetConnectedEndpointTypeNil

`func (o *PowerPort) SetConnectedEndpointTypeNil(b bool)`

 SetConnectedEndpointTypeNil sets the value for ConnectedEndpointType to be an explicit nil

### UnsetConnectedEndpointType
`func (o *PowerPort) UnsetConnectedEndpointType()`

UnsetConnectedEndpointType ensures that no value is present for ConnectedEndpointType, not even an explicit nil
### GetConnectedEndpoint

`func (o *PowerPort) GetConnectedEndpoint() PathEndpoint`

GetConnectedEndpoint returns the ConnectedEndpoint field if non-nil, zero value otherwise.

### GetConnectedEndpointOk

`func (o *PowerPort) GetConnectedEndpointOk() (*PathEndpoint, bool)`

GetConnectedEndpointOk returns a tuple with the ConnectedEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedEndpoint

`func (o *PowerPort) SetConnectedEndpoint(v PathEndpoint)`

SetConnectedEndpoint sets ConnectedEndpoint field to given value.


### SetConnectedEndpointNil

`func (o *PowerPort) SetConnectedEndpointNil(b bool)`

 SetConnectedEndpointNil sets the value for ConnectedEndpoint to be an explicit nil

### UnsetConnectedEndpoint
`func (o *PowerPort) UnsetConnectedEndpoint()`

UnsetConnectedEndpoint ensures that no value is present for ConnectedEndpoint, not even an explicit nil
### GetConnectedEndpointReachable

`func (o *PowerPort) GetConnectedEndpointReachable() bool`

GetConnectedEndpointReachable returns the ConnectedEndpointReachable field if non-nil, zero value otherwise.

### GetConnectedEndpointReachableOk

`func (o *PowerPort) GetConnectedEndpointReachableOk() (*bool, bool)`

GetConnectedEndpointReachableOk returns a tuple with the ConnectedEndpointReachable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedEndpointReachable

`func (o *PowerPort) SetConnectedEndpointReachable(v bool)`

SetConnectedEndpointReachable sets ConnectedEndpointReachable field to given value.


### SetConnectedEndpointReachableNil

`func (o *PowerPort) SetConnectedEndpointReachableNil(b bool)`

 SetConnectedEndpointReachableNil sets the value for ConnectedEndpointReachable to be an explicit nil

### UnsetConnectedEndpointReachable
`func (o *PowerPort) UnsetConnectedEndpointReachable()`

UnsetConnectedEndpointReachable ensures that no value is present for ConnectedEndpointReachable, not even an explicit nil
### GetType

`func (o *PowerPort) GetType() PowerPortType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PowerPort) GetTypeOk() (*PowerPortType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PowerPort) SetType(v PowerPortType)`

SetType sets Type field to given value.

### HasType

`func (o *PowerPort) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *PowerPort) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PowerPort) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PowerPort) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *PowerPort) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PowerPort) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PowerPort) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PowerPort) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDescription

`func (o *PowerPort) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PowerPort) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PowerPort) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PowerPort) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMaximumDraw

`func (o *PowerPort) GetMaximumDraw() int32`

GetMaximumDraw returns the MaximumDraw field if non-nil, zero value otherwise.

### GetMaximumDrawOk

`func (o *PowerPort) GetMaximumDrawOk() (*int32, bool)`

GetMaximumDrawOk returns a tuple with the MaximumDraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumDraw

`func (o *PowerPort) SetMaximumDraw(v int32)`

SetMaximumDraw sets MaximumDraw field to given value.

### HasMaximumDraw

`func (o *PowerPort) HasMaximumDraw() bool`

HasMaximumDraw returns a boolean if a field has been set.

### SetMaximumDrawNil

`func (o *PowerPort) SetMaximumDrawNil(b bool)`

 SetMaximumDrawNil sets the value for MaximumDraw to be an explicit nil

### UnsetMaximumDraw
`func (o *PowerPort) UnsetMaximumDraw()`

UnsetMaximumDraw ensures that no value is present for MaximumDraw, not even an explicit nil
### GetAllocatedDraw

`func (o *PowerPort) GetAllocatedDraw() int32`

GetAllocatedDraw returns the AllocatedDraw field if non-nil, zero value otherwise.

### GetAllocatedDrawOk

`func (o *PowerPort) GetAllocatedDrawOk() (*int32, bool)`

GetAllocatedDrawOk returns a tuple with the AllocatedDraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedDraw

`func (o *PowerPort) SetAllocatedDraw(v int32)`

SetAllocatedDraw sets AllocatedDraw field to given value.

### HasAllocatedDraw

`func (o *PowerPort) HasAllocatedDraw() bool`

HasAllocatedDraw returns a boolean if a field has been set.

### SetAllocatedDrawNil

`func (o *PowerPort) SetAllocatedDrawNil(b bool)`

 SetAllocatedDrawNil sets the value for AllocatedDraw to be an explicit nil

### UnsetAllocatedDraw
`func (o *PowerPort) UnsetAllocatedDraw()`

UnsetAllocatedDraw ensures that no value is present for AllocatedDraw, not even an explicit nil
### GetDevice

`func (o *PowerPort) GetDevice() BulkWritableCableRequestStatus`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *PowerPort) GetDeviceOk() (*BulkWritableCableRequestStatus, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *PowerPort) SetDevice(v BulkWritableCableRequestStatus)`

SetDevice sets Device field to given value.


### GetCable

`func (o *PowerPort) GetCable() CircuitCircuitTerminationA`

GetCable returns the Cable field if non-nil, zero value otherwise.

### GetCableOk

`func (o *PowerPort) GetCableOk() (*CircuitCircuitTerminationA, bool)`

GetCableOk returns a tuple with the Cable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCable

`func (o *PowerPort) SetCable(v CircuitCircuitTerminationA)`

SetCable sets Cable field to given value.


### SetCableNil

`func (o *PowerPort) SetCableNil(b bool)`

 SetCableNil sets the value for Cable to be an explicit nil

### UnsetCable
`func (o *PowerPort) UnsetCable()`

UnsetCable ensures that no value is present for Cable, not even an explicit nil
### GetCreated

`func (o *PowerPort) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PowerPort) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PowerPort) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *PowerPort) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *PowerPort) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *PowerPort) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *PowerPort) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *PowerPort) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *PowerPort) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *PowerPort) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetTags

`func (o *PowerPort) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PowerPort) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PowerPort) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PowerPort) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetNotesUrl

`func (o *PowerPort) GetNotesUrl() string`

GetNotesUrl returns the NotesUrl field if non-nil, zero value otherwise.

### GetNotesUrlOk

`func (o *PowerPort) GetNotesUrlOk() (*string, bool)`

GetNotesUrlOk returns a tuple with the NotesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesUrl

`func (o *PowerPort) SetNotesUrl(v string)`

SetNotesUrl sets NotesUrl field to given value.


### GetCustomFields

`func (o *PowerPort) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PowerPort) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PowerPort) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PowerPort) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


