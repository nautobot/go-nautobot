# PatchedObjectMetadataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignedObjectType** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **interface{}** |  | [optional] 
**ScopedFields** | Pointer to **interface{}** | List of scoped fields, only direct fields on the model | [optional] 
**AssignedObjectId** | Pointer to **string** |  | [optional] 
**MetadataType** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**Contact** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Team** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 

## Methods

### NewPatchedObjectMetadataRequest

`func NewPatchedObjectMetadataRequest() *PatchedObjectMetadataRequest`

NewPatchedObjectMetadataRequest instantiates a new PatchedObjectMetadataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedObjectMetadataRequestWithDefaults

`func NewPatchedObjectMetadataRequestWithDefaults() *PatchedObjectMetadataRequest`

NewPatchedObjectMetadataRequestWithDefaults instantiates a new PatchedObjectMetadataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignedObjectType

`func (o *PatchedObjectMetadataRequest) GetAssignedObjectType() string`

GetAssignedObjectType returns the AssignedObjectType field if non-nil, zero value otherwise.

### GetAssignedObjectTypeOk

`func (o *PatchedObjectMetadataRequest) GetAssignedObjectTypeOk() (*string, bool)`

GetAssignedObjectTypeOk returns a tuple with the AssignedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObjectType

`func (o *PatchedObjectMetadataRequest) SetAssignedObjectType(v string)`

SetAssignedObjectType sets AssignedObjectType field to given value.

### HasAssignedObjectType

`func (o *PatchedObjectMetadataRequest) HasAssignedObjectType() bool`

HasAssignedObjectType returns a boolean if a field has been set.

### GetValue

`func (o *PatchedObjectMetadataRequest) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PatchedObjectMetadataRequest) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PatchedObjectMetadataRequest) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *PatchedObjectMetadataRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *PatchedObjectMetadataRequest) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *PatchedObjectMetadataRequest) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetScopedFields

`func (o *PatchedObjectMetadataRequest) GetScopedFields() interface{}`

GetScopedFields returns the ScopedFields field if non-nil, zero value otherwise.

### GetScopedFieldsOk

`func (o *PatchedObjectMetadataRequest) GetScopedFieldsOk() (*interface{}, bool)`

GetScopedFieldsOk returns a tuple with the ScopedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopedFields

`func (o *PatchedObjectMetadataRequest) SetScopedFields(v interface{})`

SetScopedFields sets ScopedFields field to given value.

### HasScopedFields

`func (o *PatchedObjectMetadataRequest) HasScopedFields() bool`

HasScopedFields returns a boolean if a field has been set.

### SetScopedFieldsNil

`func (o *PatchedObjectMetadataRequest) SetScopedFieldsNil(b bool)`

 SetScopedFieldsNil sets the value for ScopedFields to be an explicit nil

### UnsetScopedFields
`func (o *PatchedObjectMetadataRequest) UnsetScopedFields()`

UnsetScopedFields ensures that no value is present for ScopedFields, not even an explicit nil
### GetAssignedObjectId

`func (o *PatchedObjectMetadataRequest) GetAssignedObjectId() string`

GetAssignedObjectId returns the AssignedObjectId field if non-nil, zero value otherwise.

### GetAssignedObjectIdOk

`func (o *PatchedObjectMetadataRequest) GetAssignedObjectIdOk() (*string, bool)`

GetAssignedObjectIdOk returns a tuple with the AssignedObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObjectId

`func (o *PatchedObjectMetadataRequest) SetAssignedObjectId(v string)`

SetAssignedObjectId sets AssignedObjectId field to given value.

### HasAssignedObjectId

`func (o *PatchedObjectMetadataRequest) HasAssignedObjectId() bool`

HasAssignedObjectId returns a boolean if a field has been set.

### GetMetadataType

`func (o *PatchedObjectMetadataRequest) GetMetadataType() BulkWritableCableRequestStatus`

GetMetadataType returns the MetadataType field if non-nil, zero value otherwise.

### GetMetadataTypeOk

`func (o *PatchedObjectMetadataRequest) GetMetadataTypeOk() (*BulkWritableCableRequestStatus, bool)`

GetMetadataTypeOk returns a tuple with the MetadataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataType

`func (o *PatchedObjectMetadataRequest) SetMetadataType(v BulkWritableCableRequestStatus)`

SetMetadataType sets MetadataType field to given value.

### HasMetadataType

`func (o *PatchedObjectMetadataRequest) HasMetadataType() bool`

HasMetadataType returns a boolean if a field has been set.

### GetContact

`func (o *PatchedObjectMetadataRequest) GetContact() BulkWritableCircuitRequestTenant`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *PatchedObjectMetadataRequest) GetContactOk() (*BulkWritableCircuitRequestTenant, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *PatchedObjectMetadataRequest) SetContact(v BulkWritableCircuitRequestTenant)`

SetContact sets Contact field to given value.

### HasContact

`func (o *PatchedObjectMetadataRequest) HasContact() bool`

HasContact returns a boolean if a field has been set.

### SetContactNil

`func (o *PatchedObjectMetadataRequest) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *PatchedObjectMetadataRequest) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetTeam

`func (o *PatchedObjectMetadataRequest) GetTeam() BulkWritableCircuitRequestTenant`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *PatchedObjectMetadataRequest) GetTeamOk() (*BulkWritableCircuitRequestTenant, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *PatchedObjectMetadataRequest) SetTeam(v BulkWritableCircuitRequestTenant)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *PatchedObjectMetadataRequest) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### SetTeamNil

`func (o *PatchedObjectMetadataRequest) SetTeamNil(b bool)`

 SetTeamNil sets the value for Team to be an explicit nil

### UnsetTeam
`func (o *PatchedObjectMetadataRequest) UnsetTeam()`

UnsetTeam ensures that no value is present for Team, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


