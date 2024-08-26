# PatchedNoteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignedObjectType** | Pointer to **string** |  | [optional] 
**AssignedObjectId** | Pointer to **string** |  | [optional] 
**Note** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 

## Methods

### NewPatchedNoteRequest

`func NewPatchedNoteRequest() *PatchedNoteRequest`

NewPatchedNoteRequest instantiates a new PatchedNoteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedNoteRequestWithDefaults

`func NewPatchedNoteRequestWithDefaults() *PatchedNoteRequest`

NewPatchedNoteRequestWithDefaults instantiates a new PatchedNoteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignedObjectType

`func (o *PatchedNoteRequest) GetAssignedObjectType() string`

GetAssignedObjectType returns the AssignedObjectType field if non-nil, zero value otherwise.

### GetAssignedObjectTypeOk

`func (o *PatchedNoteRequest) GetAssignedObjectTypeOk() (*string, bool)`

GetAssignedObjectTypeOk returns a tuple with the AssignedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObjectType

`func (o *PatchedNoteRequest) SetAssignedObjectType(v string)`

SetAssignedObjectType sets AssignedObjectType field to given value.

### HasAssignedObjectType

`func (o *PatchedNoteRequest) HasAssignedObjectType() bool`

HasAssignedObjectType returns a boolean if a field has been set.

### GetAssignedObjectId

`func (o *PatchedNoteRequest) GetAssignedObjectId() string`

GetAssignedObjectId returns the AssignedObjectId field if non-nil, zero value otherwise.

### GetAssignedObjectIdOk

`func (o *PatchedNoteRequest) GetAssignedObjectIdOk() (*string, bool)`

GetAssignedObjectIdOk returns a tuple with the AssignedObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObjectId

`func (o *PatchedNoteRequest) SetAssignedObjectId(v string)`

SetAssignedObjectId sets AssignedObjectId field to given value.

### HasAssignedObjectId

`func (o *PatchedNoteRequest) HasAssignedObjectId() bool`

HasAssignedObjectId returns a boolean if a field has been set.

### GetNote

`func (o *PatchedNoteRequest) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *PatchedNoteRequest) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *PatchedNoteRequest) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *PatchedNoteRequest) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetUser

`func (o *PatchedNoteRequest) GetUser() BulkWritableCircuitRequestTenant`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PatchedNoteRequest) GetUserOk() (*BulkWritableCircuitRequestTenant, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PatchedNoteRequest) SetUser(v BulkWritableCircuitRequestTenant)`

SetUser sets User field to given value.

### HasUser

`func (o *PatchedNoteRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *PatchedNoteRequest) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *PatchedNoteRequest) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


