# RearPort

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
**Type** | [**FrontPortType**](FrontPortType.md) |  | 
**Name** | **string** |  | 
**Label** | Pointer to **string** | Physical label | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Positions** | Pointer to **int32** |  | [optional] 
**Device** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Cable** | [**NullableCircuitCircuitTerminationA**](CircuitCircuitTerminationA.md) |  | 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**NotesUrl** | **string** |  | [readonly] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewRearPort

`func NewRearPort(id string, objectType string, display string, url string, naturalSlug string, cablePeerType NullableString, cablePeer NullableCableTermination, type_ FrontPortType, name string, device BulkWritableCableRequestStatus, cable NullableCircuitCircuitTerminationA, created NullableTime, lastUpdated NullableTime, notesUrl string, ) *RearPort`

NewRearPort instantiates a new RearPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRearPortWithDefaults

`func NewRearPortWithDefaults() *RearPort`

NewRearPortWithDefaults instantiates a new RearPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RearPort) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RearPort) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RearPort) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *RearPort) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *RearPort) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *RearPort) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *RearPort) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *RearPort) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *RearPort) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *RearPort) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RearPort) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RearPort) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *RearPort) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *RearPort) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *RearPort) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetCablePeerType

`func (o *RearPort) GetCablePeerType() string`

GetCablePeerType returns the CablePeerType field if non-nil, zero value otherwise.

### GetCablePeerTypeOk

`func (o *RearPort) GetCablePeerTypeOk() (*string, bool)`

GetCablePeerTypeOk returns a tuple with the CablePeerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCablePeerType

`func (o *RearPort) SetCablePeerType(v string)`

SetCablePeerType sets CablePeerType field to given value.


### SetCablePeerTypeNil

`func (o *RearPort) SetCablePeerTypeNil(b bool)`

 SetCablePeerTypeNil sets the value for CablePeerType to be an explicit nil

### UnsetCablePeerType
`func (o *RearPort) UnsetCablePeerType()`

UnsetCablePeerType ensures that no value is present for CablePeerType, not even an explicit nil
### GetCablePeer

`func (o *RearPort) GetCablePeer() CableTermination`

GetCablePeer returns the CablePeer field if non-nil, zero value otherwise.

### GetCablePeerOk

`func (o *RearPort) GetCablePeerOk() (*CableTermination, bool)`

GetCablePeerOk returns a tuple with the CablePeer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCablePeer

`func (o *RearPort) SetCablePeer(v CableTermination)`

SetCablePeer sets CablePeer field to given value.


### SetCablePeerNil

`func (o *RearPort) SetCablePeerNil(b bool)`

 SetCablePeerNil sets the value for CablePeer to be an explicit nil

### UnsetCablePeer
`func (o *RearPort) UnsetCablePeer()`

UnsetCablePeer ensures that no value is present for CablePeer, not even an explicit nil
### GetType

`func (o *RearPort) GetType() FrontPortType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RearPort) GetTypeOk() (*FrontPortType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RearPort) SetType(v FrontPortType)`

SetType sets Type field to given value.


### GetName

`func (o *RearPort) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RearPort) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RearPort) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *RearPort) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *RearPort) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *RearPort) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *RearPort) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDescription

`func (o *RearPort) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RearPort) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RearPort) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RearPort) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPositions

`func (o *RearPort) GetPositions() int32`

GetPositions returns the Positions field if non-nil, zero value otherwise.

### GetPositionsOk

`func (o *RearPort) GetPositionsOk() (*int32, bool)`

GetPositionsOk returns a tuple with the Positions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositions

`func (o *RearPort) SetPositions(v int32)`

SetPositions sets Positions field to given value.

### HasPositions

`func (o *RearPort) HasPositions() bool`

HasPositions returns a boolean if a field has been set.

### GetDevice

`func (o *RearPort) GetDevice() BulkWritableCableRequestStatus`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *RearPort) GetDeviceOk() (*BulkWritableCableRequestStatus, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *RearPort) SetDevice(v BulkWritableCableRequestStatus)`

SetDevice sets Device field to given value.


### GetCable

`func (o *RearPort) GetCable() CircuitCircuitTerminationA`

GetCable returns the Cable field if non-nil, zero value otherwise.

### GetCableOk

`func (o *RearPort) GetCableOk() (*CircuitCircuitTerminationA, bool)`

GetCableOk returns a tuple with the Cable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCable

`func (o *RearPort) SetCable(v CircuitCircuitTerminationA)`

SetCable sets Cable field to given value.


### SetCableNil

`func (o *RearPort) SetCableNil(b bool)`

 SetCableNil sets the value for Cable to be an explicit nil

### UnsetCable
`func (o *RearPort) UnsetCable()`

UnsetCable ensures that no value is present for Cable, not even an explicit nil
### GetCreated

`func (o *RearPort) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RearPort) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RearPort) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *RearPort) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *RearPort) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *RearPort) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *RearPort) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *RearPort) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *RearPort) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *RearPort) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetTags

`func (o *RearPort) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RearPort) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RearPort) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RearPort) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetNotesUrl

`func (o *RearPort) GetNotesUrl() string`

GetNotesUrl returns the NotesUrl field if non-nil, zero value otherwise.

### GetNotesUrlOk

`func (o *RearPort) GetNotesUrlOk() (*string, bool)`

GetNotesUrlOk returns a tuple with the NotesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesUrl

`func (o *RearPort) SetNotesUrl(v string)`

SetNotesUrl sets NotesUrl field to given value.


### GetCustomFields

`func (o *RearPort) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *RearPort) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *RearPort) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *RearPort) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


