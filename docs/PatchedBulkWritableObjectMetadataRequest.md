# PatchedBulkWritableObjectMetadataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**AssignedObjectType** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **interface{}** |  | [optional] 
**ScopedFields** | Pointer to **interface{}** | List of scoped fields, only direct fields on the model | [optional] 
**AssignedObjectId** | Pointer to **string** |  | [optional] 
**MetadataType** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**Contact** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Team** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 

## Methods

### NewPatchedBulkWritableObjectMetadataRequest

`func NewPatchedBulkWritableObjectMetadataRequest(id string, ) *PatchedBulkWritableObjectMetadataRequest`

NewPatchedBulkWritableObjectMetadataRequest instantiates a new PatchedBulkWritableObjectMetadataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableObjectMetadataRequestWithDefaults

`func NewPatchedBulkWritableObjectMetadataRequestWithDefaults() *PatchedBulkWritableObjectMetadataRequest`

NewPatchedBulkWritableObjectMetadataRequestWithDefaults instantiates a new PatchedBulkWritableObjectMetadataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableObjectMetadataRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableObjectMetadataRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableObjectMetadataRequest) SetId(v string)`

SetId sets Id field to given value.


### GetAssignedObjectType

`func (o *PatchedBulkWritableObjectMetadataRequest) GetAssignedObjectType() string`

GetAssignedObjectType returns the AssignedObjectType field if non-nil, zero value otherwise.

### GetAssignedObjectTypeOk

`func (o *PatchedBulkWritableObjectMetadataRequest) GetAssignedObjectTypeOk() (*string, bool)`

GetAssignedObjectTypeOk returns a tuple with the AssignedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObjectType

`func (o *PatchedBulkWritableObjectMetadataRequest) SetAssignedObjectType(v string)`

SetAssignedObjectType sets AssignedObjectType field to given value.

### HasAssignedObjectType

`func (o *PatchedBulkWritableObjectMetadataRequest) HasAssignedObjectType() bool`

HasAssignedObjectType returns a boolean if a field has been set.

### GetValue

`func (o *PatchedBulkWritableObjectMetadataRequest) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PatchedBulkWritableObjectMetadataRequest) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PatchedBulkWritableObjectMetadataRequest) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *PatchedBulkWritableObjectMetadataRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *PatchedBulkWritableObjectMetadataRequest) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *PatchedBulkWritableObjectMetadataRequest) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetScopedFields

`func (o *PatchedBulkWritableObjectMetadataRequest) GetScopedFields() interface{}`

GetScopedFields returns the ScopedFields field if non-nil, zero value otherwise.

### GetScopedFieldsOk

`func (o *PatchedBulkWritableObjectMetadataRequest) GetScopedFieldsOk() (*interface{}, bool)`

GetScopedFieldsOk returns a tuple with the ScopedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopedFields

`func (o *PatchedBulkWritableObjectMetadataRequest) SetScopedFields(v interface{})`

SetScopedFields sets ScopedFields field to given value.

### HasScopedFields

`func (o *PatchedBulkWritableObjectMetadataRequest) HasScopedFields() bool`

HasScopedFields returns a boolean if a field has been set.

### SetScopedFieldsNil

`func (o *PatchedBulkWritableObjectMetadataRequest) SetScopedFieldsNil(b bool)`

 SetScopedFieldsNil sets the value for ScopedFields to be an explicit nil

### UnsetScopedFields
`func (o *PatchedBulkWritableObjectMetadataRequest) UnsetScopedFields()`

UnsetScopedFields ensures that no value is present for ScopedFields, not even an explicit nil
### GetAssignedObjectId

`func (o *PatchedBulkWritableObjectMetadataRequest) GetAssignedObjectId() string`

GetAssignedObjectId returns the AssignedObjectId field if non-nil, zero value otherwise.

### GetAssignedObjectIdOk

`func (o *PatchedBulkWritableObjectMetadataRequest) GetAssignedObjectIdOk() (*string, bool)`

GetAssignedObjectIdOk returns a tuple with the AssignedObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObjectId

`func (o *PatchedBulkWritableObjectMetadataRequest) SetAssignedObjectId(v string)`

SetAssignedObjectId sets AssignedObjectId field to given value.

### HasAssignedObjectId

`func (o *PatchedBulkWritableObjectMetadataRequest) HasAssignedObjectId() bool`

HasAssignedObjectId returns a boolean if a field has been set.

### GetMetadataType

`func (o *PatchedBulkWritableObjectMetadataRequest) GetMetadataType() BulkWritableCableRequestStatus`

GetMetadataType returns the MetadataType field if non-nil, zero value otherwise.

### GetMetadataTypeOk

`func (o *PatchedBulkWritableObjectMetadataRequest) GetMetadataTypeOk() (*BulkWritableCableRequestStatus, bool)`

GetMetadataTypeOk returns a tuple with the MetadataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataType

`func (o *PatchedBulkWritableObjectMetadataRequest) SetMetadataType(v BulkWritableCableRequestStatus)`

SetMetadataType sets MetadataType field to given value.

### HasMetadataType

`func (o *PatchedBulkWritableObjectMetadataRequest) HasMetadataType() bool`

HasMetadataType returns a boolean if a field has been set.

### GetContact

`func (o *PatchedBulkWritableObjectMetadataRequest) GetContact() BulkWritableCircuitRequestTenant`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *PatchedBulkWritableObjectMetadataRequest) GetContactOk() (*BulkWritableCircuitRequestTenant, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *PatchedBulkWritableObjectMetadataRequest) SetContact(v BulkWritableCircuitRequestTenant)`

SetContact sets Contact field to given value.

### HasContact

`func (o *PatchedBulkWritableObjectMetadataRequest) HasContact() bool`

HasContact returns a boolean if a field has been set.

### SetContactNil

`func (o *PatchedBulkWritableObjectMetadataRequest) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *PatchedBulkWritableObjectMetadataRequest) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetTeam

`func (o *PatchedBulkWritableObjectMetadataRequest) GetTeam() BulkWritableCircuitRequestTenant`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *PatchedBulkWritableObjectMetadataRequest) GetTeamOk() (*BulkWritableCircuitRequestTenant, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *PatchedBulkWritableObjectMetadataRequest) SetTeam(v BulkWritableCircuitRequestTenant)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *PatchedBulkWritableObjectMetadataRequest) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### SetTeamNil

`func (o *PatchedBulkWritableObjectMetadataRequest) SetTeamNil(b bool)`

 SetTeamNil sets the value for Team to be an explicit nil

### UnsetTeam
`func (o *PatchedBulkWritableObjectMetadataRequest) UnsetTeam()`

UnsetTeam ensures that no value is present for Team, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


