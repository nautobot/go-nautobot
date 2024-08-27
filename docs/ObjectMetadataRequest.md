# ObjectMetadataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignedObjectType** | **string** |  | 
**Value** | Pointer to **interface{}** |  | [optional] 
**ScopedFields** | Pointer to **interface{}** | List of scoped fields, only direct fields on the model | [optional] 
**AssignedObjectId** | **string** |  | 
**MetadataType** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Contact** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Team** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 

## Methods

### NewObjectMetadataRequest

`func NewObjectMetadataRequest(assignedObjectType string, assignedObjectId string, metadataType BulkWritableCableRequestStatus, ) *ObjectMetadataRequest`

NewObjectMetadataRequest instantiates a new ObjectMetadataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectMetadataRequestWithDefaults

`func NewObjectMetadataRequestWithDefaults() *ObjectMetadataRequest`

NewObjectMetadataRequestWithDefaults instantiates a new ObjectMetadataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignedObjectType

`func (o *ObjectMetadataRequest) GetAssignedObjectType() string`

GetAssignedObjectType returns the AssignedObjectType field if non-nil, zero value otherwise.

### GetAssignedObjectTypeOk

`func (o *ObjectMetadataRequest) GetAssignedObjectTypeOk() (*string, bool)`

GetAssignedObjectTypeOk returns a tuple with the AssignedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObjectType

`func (o *ObjectMetadataRequest) SetAssignedObjectType(v string)`

SetAssignedObjectType sets AssignedObjectType field to given value.


### GetValue

`func (o *ObjectMetadataRequest) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ObjectMetadataRequest) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ObjectMetadataRequest) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *ObjectMetadataRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ObjectMetadataRequest) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ObjectMetadataRequest) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetScopedFields

`func (o *ObjectMetadataRequest) GetScopedFields() interface{}`

GetScopedFields returns the ScopedFields field if non-nil, zero value otherwise.

### GetScopedFieldsOk

`func (o *ObjectMetadataRequest) GetScopedFieldsOk() (*interface{}, bool)`

GetScopedFieldsOk returns a tuple with the ScopedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopedFields

`func (o *ObjectMetadataRequest) SetScopedFields(v interface{})`

SetScopedFields sets ScopedFields field to given value.

### HasScopedFields

`func (o *ObjectMetadataRequest) HasScopedFields() bool`

HasScopedFields returns a boolean if a field has been set.

### SetScopedFieldsNil

`func (o *ObjectMetadataRequest) SetScopedFieldsNil(b bool)`

 SetScopedFieldsNil sets the value for ScopedFields to be an explicit nil

### UnsetScopedFields
`func (o *ObjectMetadataRequest) UnsetScopedFields()`

UnsetScopedFields ensures that no value is present for ScopedFields, not even an explicit nil
### GetAssignedObjectId

`func (o *ObjectMetadataRequest) GetAssignedObjectId() string`

GetAssignedObjectId returns the AssignedObjectId field if non-nil, zero value otherwise.

### GetAssignedObjectIdOk

`func (o *ObjectMetadataRequest) GetAssignedObjectIdOk() (*string, bool)`

GetAssignedObjectIdOk returns a tuple with the AssignedObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObjectId

`func (o *ObjectMetadataRequest) SetAssignedObjectId(v string)`

SetAssignedObjectId sets AssignedObjectId field to given value.


### GetMetadataType

`func (o *ObjectMetadataRequest) GetMetadataType() BulkWritableCableRequestStatus`

GetMetadataType returns the MetadataType field if non-nil, zero value otherwise.

### GetMetadataTypeOk

`func (o *ObjectMetadataRequest) GetMetadataTypeOk() (*BulkWritableCableRequestStatus, bool)`

GetMetadataTypeOk returns a tuple with the MetadataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataType

`func (o *ObjectMetadataRequest) SetMetadataType(v BulkWritableCableRequestStatus)`

SetMetadataType sets MetadataType field to given value.


### GetContact

`func (o *ObjectMetadataRequest) GetContact() BulkWritableCircuitRequestTenant`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *ObjectMetadataRequest) GetContactOk() (*BulkWritableCircuitRequestTenant, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *ObjectMetadataRequest) SetContact(v BulkWritableCircuitRequestTenant)`

SetContact sets Contact field to given value.

### HasContact

`func (o *ObjectMetadataRequest) HasContact() bool`

HasContact returns a boolean if a field has been set.

### SetContactNil

`func (o *ObjectMetadataRequest) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *ObjectMetadataRequest) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetTeam

`func (o *ObjectMetadataRequest) GetTeam() BulkWritableCircuitRequestTenant`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *ObjectMetadataRequest) GetTeamOk() (*BulkWritableCircuitRequestTenant, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *ObjectMetadataRequest) SetTeam(v BulkWritableCircuitRequestTenant)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *ObjectMetadataRequest) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### SetTeamNil

`func (o *ObjectMetadataRequest) SetTeamNil(b bool)`

 SetTeamNil sets the value for Team to be an explicit nil

### UnsetTeam
`func (o *ObjectMetadataRequest) UnsetTeam()`

UnsetTeam ensures that no value is present for Team, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


