# BulkWritableObjectMetadataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**AssignedObjectType** | **string** |  | 
**Value** | Pointer to **interface{}** |  | [optional] 
**ScopedFields** | Pointer to **interface{}** | List of scoped fields, only direct fields on the model | [optional] 
**AssignedObjectId** | **string** |  | 
**MetadataType** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Contact** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Team** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 

## Methods

### NewBulkWritableObjectMetadataRequest

`func NewBulkWritableObjectMetadataRequest(id string, assignedObjectType string, assignedObjectId string, metadataType BulkWritableCableRequestStatus, ) *BulkWritableObjectMetadataRequest`

NewBulkWritableObjectMetadataRequest instantiates a new BulkWritableObjectMetadataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableObjectMetadataRequestWithDefaults

`func NewBulkWritableObjectMetadataRequestWithDefaults() *BulkWritableObjectMetadataRequest`

NewBulkWritableObjectMetadataRequestWithDefaults instantiates a new BulkWritableObjectMetadataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableObjectMetadataRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableObjectMetadataRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableObjectMetadataRequest) SetId(v string)`

SetId sets Id field to given value.


### GetAssignedObjectType

`func (o *BulkWritableObjectMetadataRequest) GetAssignedObjectType() string`

GetAssignedObjectType returns the AssignedObjectType field if non-nil, zero value otherwise.

### GetAssignedObjectTypeOk

`func (o *BulkWritableObjectMetadataRequest) GetAssignedObjectTypeOk() (*string, bool)`

GetAssignedObjectTypeOk returns a tuple with the AssignedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObjectType

`func (o *BulkWritableObjectMetadataRequest) SetAssignedObjectType(v string)`

SetAssignedObjectType sets AssignedObjectType field to given value.


### GetValue

`func (o *BulkWritableObjectMetadataRequest) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BulkWritableObjectMetadataRequest) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BulkWritableObjectMetadataRequest) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *BulkWritableObjectMetadataRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *BulkWritableObjectMetadataRequest) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *BulkWritableObjectMetadataRequest) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetScopedFields

`func (o *BulkWritableObjectMetadataRequest) GetScopedFields() interface{}`

GetScopedFields returns the ScopedFields field if non-nil, zero value otherwise.

### GetScopedFieldsOk

`func (o *BulkWritableObjectMetadataRequest) GetScopedFieldsOk() (*interface{}, bool)`

GetScopedFieldsOk returns a tuple with the ScopedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopedFields

`func (o *BulkWritableObjectMetadataRequest) SetScopedFields(v interface{})`

SetScopedFields sets ScopedFields field to given value.

### HasScopedFields

`func (o *BulkWritableObjectMetadataRequest) HasScopedFields() bool`

HasScopedFields returns a boolean if a field has been set.

### SetScopedFieldsNil

`func (o *BulkWritableObjectMetadataRequest) SetScopedFieldsNil(b bool)`

 SetScopedFieldsNil sets the value for ScopedFields to be an explicit nil

### UnsetScopedFields
`func (o *BulkWritableObjectMetadataRequest) UnsetScopedFields()`

UnsetScopedFields ensures that no value is present for ScopedFields, not even an explicit nil
### GetAssignedObjectId

`func (o *BulkWritableObjectMetadataRequest) GetAssignedObjectId() string`

GetAssignedObjectId returns the AssignedObjectId field if non-nil, zero value otherwise.

### GetAssignedObjectIdOk

`func (o *BulkWritableObjectMetadataRequest) GetAssignedObjectIdOk() (*string, bool)`

GetAssignedObjectIdOk returns a tuple with the AssignedObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObjectId

`func (o *BulkWritableObjectMetadataRequest) SetAssignedObjectId(v string)`

SetAssignedObjectId sets AssignedObjectId field to given value.


### GetMetadataType

`func (o *BulkWritableObjectMetadataRequest) GetMetadataType() BulkWritableCableRequestStatus`

GetMetadataType returns the MetadataType field if non-nil, zero value otherwise.

### GetMetadataTypeOk

`func (o *BulkWritableObjectMetadataRequest) GetMetadataTypeOk() (*BulkWritableCableRequestStatus, bool)`

GetMetadataTypeOk returns a tuple with the MetadataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataType

`func (o *BulkWritableObjectMetadataRequest) SetMetadataType(v BulkWritableCableRequestStatus)`

SetMetadataType sets MetadataType field to given value.


### GetContact

`func (o *BulkWritableObjectMetadataRequest) GetContact() BulkWritableCircuitRequestTenant`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *BulkWritableObjectMetadataRequest) GetContactOk() (*BulkWritableCircuitRequestTenant, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *BulkWritableObjectMetadataRequest) SetContact(v BulkWritableCircuitRequestTenant)`

SetContact sets Contact field to given value.

### HasContact

`func (o *BulkWritableObjectMetadataRequest) HasContact() bool`

HasContact returns a boolean if a field has been set.

### SetContactNil

`func (o *BulkWritableObjectMetadataRequest) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *BulkWritableObjectMetadataRequest) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetTeam

`func (o *BulkWritableObjectMetadataRequest) GetTeam() BulkWritableCircuitRequestTenant`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *BulkWritableObjectMetadataRequest) GetTeamOk() (*BulkWritableCircuitRequestTenant, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *BulkWritableObjectMetadataRequest) SetTeam(v BulkWritableCircuitRequestTenant)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *BulkWritableObjectMetadataRequest) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### SetTeamNil

`func (o *BulkWritableObjectMetadataRequest) SetTeamNil(b bool)`

 SetTeamNil sets the value for Team to be an explicit nil

### UnsetTeam
`func (o *BulkWritableObjectMetadataRequest) UnsetTeam()`

UnsetTeam ensures that no value is present for Team, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


