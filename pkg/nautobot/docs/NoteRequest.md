# NoteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignedObjectType** | **string** |  | 
**AssignedObjectId** | **string** |  | 
**Note** | **string** |  | 
**User** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 

## Methods

### NewNoteRequest

`func NewNoteRequest(assignedObjectType string, assignedObjectId string, note string, ) *NoteRequest`

NewNoteRequest instantiates a new NoteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoteRequestWithDefaults

`func NewNoteRequestWithDefaults() *NoteRequest`

NewNoteRequestWithDefaults instantiates a new NoteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignedObjectType

`func (o *NoteRequest) GetAssignedObjectType() string`

GetAssignedObjectType returns the AssignedObjectType field if non-nil, zero value otherwise.

### GetAssignedObjectTypeOk

`func (o *NoteRequest) GetAssignedObjectTypeOk() (*string, bool)`

GetAssignedObjectTypeOk returns a tuple with the AssignedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObjectType

`func (o *NoteRequest) SetAssignedObjectType(v string)`

SetAssignedObjectType sets AssignedObjectType field to given value.


### GetAssignedObjectId

`func (o *NoteRequest) GetAssignedObjectId() string`

GetAssignedObjectId returns the AssignedObjectId field if non-nil, zero value otherwise.

### GetAssignedObjectIdOk

`func (o *NoteRequest) GetAssignedObjectIdOk() (*string, bool)`

GetAssignedObjectIdOk returns a tuple with the AssignedObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObjectId

`func (o *NoteRequest) SetAssignedObjectId(v string)`

SetAssignedObjectId sets AssignedObjectId field to given value.


### GetNote

`func (o *NoteRequest) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *NoteRequest) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *NoteRequest) SetNote(v string)`

SetNote sets Note field to given value.


### GetUser

`func (o *NoteRequest) GetUser() BulkWritableCircuitRequestTenant`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *NoteRequest) GetUserOk() (*BulkWritableCircuitRequestTenant, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *NoteRequest) SetUser(v BulkWritableCircuitRequestTenant)`

SetUser sets User field to given value.

### HasUser

`func (o *NoteRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *NoteRequest) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *NoteRequest) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


