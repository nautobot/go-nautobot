# BulkWritableNoteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**AssignedObjectType** | **string** |  | 
**AssignedObjectId** | **string** |  | 
**Note** | **string** |  | 
**User** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 

## Methods

### NewBulkWritableNoteRequest

`func NewBulkWritableNoteRequest(id string, assignedObjectType string, assignedObjectId string, note string, ) *BulkWritableNoteRequest`

NewBulkWritableNoteRequest instantiates a new BulkWritableNoteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableNoteRequestWithDefaults

`func NewBulkWritableNoteRequestWithDefaults() *BulkWritableNoteRequest`

NewBulkWritableNoteRequestWithDefaults instantiates a new BulkWritableNoteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableNoteRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableNoteRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableNoteRequest) SetId(v string)`

SetId sets Id field to given value.


### GetAssignedObjectType

`func (o *BulkWritableNoteRequest) GetAssignedObjectType() string`

GetAssignedObjectType returns the AssignedObjectType field if non-nil, zero value otherwise.

### GetAssignedObjectTypeOk

`func (o *BulkWritableNoteRequest) GetAssignedObjectTypeOk() (*string, bool)`

GetAssignedObjectTypeOk returns a tuple with the AssignedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObjectType

`func (o *BulkWritableNoteRequest) SetAssignedObjectType(v string)`

SetAssignedObjectType sets AssignedObjectType field to given value.


### GetAssignedObjectId

`func (o *BulkWritableNoteRequest) GetAssignedObjectId() string`

GetAssignedObjectId returns the AssignedObjectId field if non-nil, zero value otherwise.

### GetAssignedObjectIdOk

`func (o *BulkWritableNoteRequest) GetAssignedObjectIdOk() (*string, bool)`

GetAssignedObjectIdOk returns a tuple with the AssignedObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObjectId

`func (o *BulkWritableNoteRequest) SetAssignedObjectId(v string)`

SetAssignedObjectId sets AssignedObjectId field to given value.


### GetNote

`func (o *BulkWritableNoteRequest) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *BulkWritableNoteRequest) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *BulkWritableNoteRequest) SetNote(v string)`

SetNote sets Note field to given value.


### GetUser

`func (o *BulkWritableNoteRequest) GetUser() BulkWritableCircuitRequestTenant`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *BulkWritableNoteRequest) GetUserOk() (*BulkWritableCircuitRequestTenant, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *BulkWritableNoteRequest) SetUser(v BulkWritableCircuitRequestTenant)`

SetUser sets User field to given value.

### HasUser

`func (o *BulkWritableNoteRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *BulkWritableNoteRequest) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *BulkWritableNoteRequest) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


