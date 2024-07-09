# PatchedBulkWritableNoteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**AssignedObjectType** | Pointer to **string** |  | [optional] 
**AssignedObjectId** | Pointer to **string** |  | [optional] 
**Note** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 

## Methods

### NewPatchedBulkWritableNoteRequest

`func NewPatchedBulkWritableNoteRequest(id string, ) *PatchedBulkWritableNoteRequest`

NewPatchedBulkWritableNoteRequest instantiates a new PatchedBulkWritableNoteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableNoteRequestWithDefaults

`func NewPatchedBulkWritableNoteRequestWithDefaults() *PatchedBulkWritableNoteRequest`

NewPatchedBulkWritableNoteRequestWithDefaults instantiates a new PatchedBulkWritableNoteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableNoteRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableNoteRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableNoteRequest) SetId(v string)`

SetId sets Id field to given value.


### GetAssignedObjectType

`func (o *PatchedBulkWritableNoteRequest) GetAssignedObjectType() string`

GetAssignedObjectType returns the AssignedObjectType field if non-nil, zero value otherwise.

### GetAssignedObjectTypeOk

`func (o *PatchedBulkWritableNoteRequest) GetAssignedObjectTypeOk() (*string, bool)`

GetAssignedObjectTypeOk returns a tuple with the AssignedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObjectType

`func (o *PatchedBulkWritableNoteRequest) SetAssignedObjectType(v string)`

SetAssignedObjectType sets AssignedObjectType field to given value.

### HasAssignedObjectType

`func (o *PatchedBulkWritableNoteRequest) HasAssignedObjectType() bool`

HasAssignedObjectType returns a boolean if a field has been set.

### GetAssignedObjectId

`func (o *PatchedBulkWritableNoteRequest) GetAssignedObjectId() string`

GetAssignedObjectId returns the AssignedObjectId field if non-nil, zero value otherwise.

### GetAssignedObjectIdOk

`func (o *PatchedBulkWritableNoteRequest) GetAssignedObjectIdOk() (*string, bool)`

GetAssignedObjectIdOk returns a tuple with the AssignedObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObjectId

`func (o *PatchedBulkWritableNoteRequest) SetAssignedObjectId(v string)`

SetAssignedObjectId sets AssignedObjectId field to given value.

### HasAssignedObjectId

`func (o *PatchedBulkWritableNoteRequest) HasAssignedObjectId() bool`

HasAssignedObjectId returns a boolean if a field has been set.

### GetNote

`func (o *PatchedBulkWritableNoteRequest) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *PatchedBulkWritableNoteRequest) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *PatchedBulkWritableNoteRequest) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *PatchedBulkWritableNoteRequest) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetUser

`func (o *PatchedBulkWritableNoteRequest) GetUser() BulkWritableCircuitRequestTenant`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PatchedBulkWritableNoteRequest) GetUserOk() (*BulkWritableCircuitRequestTenant, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PatchedBulkWritableNoteRequest) SetUser(v BulkWritableCircuitRequestTenant)`

SetUser sets User field to given value.

### HasUser

`func (o *PatchedBulkWritableNoteRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *PatchedBulkWritableNoteRequest) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *PatchedBulkWritableNoteRequest) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


