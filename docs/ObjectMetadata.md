# ObjectMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**AssignedObjectType** | **string** |  | 
**AssignedObject** | [**NullableObjectMetadataAssignedObject**](ObjectMetadataAssignedObject.md) |  | [readonly] 
**Value** | Pointer to **interface{}** |  | [optional] 
**ScopedFields** | Pointer to **interface{}** | List of scoped fields, only direct fields on the model | [optional] 
**AssignedObjectId** | **string** |  | 
**MetadataType** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Contact** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Team** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 

## Methods

### NewObjectMetadata

`func NewObjectMetadata(id string, objectType string, display string, url string, naturalSlug string, assignedObjectType string, assignedObject NullableObjectMetadataAssignedObject, assignedObjectId string, metadataType BulkWritableCableRequestStatus, created NullableTime, lastUpdated NullableTime, ) *ObjectMetadata`

NewObjectMetadata instantiates a new ObjectMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectMetadataWithDefaults

`func NewObjectMetadataWithDefaults() *ObjectMetadata`

NewObjectMetadataWithDefaults instantiates a new ObjectMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ObjectMetadata) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ObjectMetadata) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ObjectMetadata) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *ObjectMetadata) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ObjectMetadata) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ObjectMetadata) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *ObjectMetadata) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *ObjectMetadata) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *ObjectMetadata) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *ObjectMetadata) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ObjectMetadata) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ObjectMetadata) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *ObjectMetadata) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *ObjectMetadata) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *ObjectMetadata) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetAssignedObjectType

`func (o *ObjectMetadata) GetAssignedObjectType() string`

GetAssignedObjectType returns the AssignedObjectType field if non-nil, zero value otherwise.

### GetAssignedObjectTypeOk

`func (o *ObjectMetadata) GetAssignedObjectTypeOk() (*string, bool)`

GetAssignedObjectTypeOk returns a tuple with the AssignedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObjectType

`func (o *ObjectMetadata) SetAssignedObjectType(v string)`

SetAssignedObjectType sets AssignedObjectType field to given value.


### GetAssignedObject

`func (o *ObjectMetadata) GetAssignedObject() ObjectMetadataAssignedObject`

GetAssignedObject returns the AssignedObject field if non-nil, zero value otherwise.

### GetAssignedObjectOk

`func (o *ObjectMetadata) GetAssignedObjectOk() (*ObjectMetadataAssignedObject, bool)`

GetAssignedObjectOk returns a tuple with the AssignedObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObject

`func (o *ObjectMetadata) SetAssignedObject(v ObjectMetadataAssignedObject)`

SetAssignedObject sets AssignedObject field to given value.


### SetAssignedObjectNil

`func (o *ObjectMetadata) SetAssignedObjectNil(b bool)`

 SetAssignedObjectNil sets the value for AssignedObject to be an explicit nil

### UnsetAssignedObject
`func (o *ObjectMetadata) UnsetAssignedObject()`

UnsetAssignedObject ensures that no value is present for AssignedObject, not even an explicit nil
### GetValue

`func (o *ObjectMetadata) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ObjectMetadata) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ObjectMetadata) SetValue(v interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *ObjectMetadata) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ObjectMetadata) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ObjectMetadata) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetScopedFields

`func (o *ObjectMetadata) GetScopedFields() interface{}`

GetScopedFields returns the ScopedFields field if non-nil, zero value otherwise.

### GetScopedFieldsOk

`func (o *ObjectMetadata) GetScopedFieldsOk() (*interface{}, bool)`

GetScopedFieldsOk returns a tuple with the ScopedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopedFields

`func (o *ObjectMetadata) SetScopedFields(v interface{})`

SetScopedFields sets ScopedFields field to given value.

### HasScopedFields

`func (o *ObjectMetadata) HasScopedFields() bool`

HasScopedFields returns a boolean if a field has been set.

### SetScopedFieldsNil

`func (o *ObjectMetadata) SetScopedFieldsNil(b bool)`

 SetScopedFieldsNil sets the value for ScopedFields to be an explicit nil

### UnsetScopedFields
`func (o *ObjectMetadata) UnsetScopedFields()`

UnsetScopedFields ensures that no value is present for ScopedFields, not even an explicit nil
### GetAssignedObjectId

`func (o *ObjectMetadata) GetAssignedObjectId() string`

GetAssignedObjectId returns the AssignedObjectId field if non-nil, zero value otherwise.

### GetAssignedObjectIdOk

`func (o *ObjectMetadata) GetAssignedObjectIdOk() (*string, bool)`

GetAssignedObjectIdOk returns a tuple with the AssignedObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObjectId

`func (o *ObjectMetadata) SetAssignedObjectId(v string)`

SetAssignedObjectId sets AssignedObjectId field to given value.


### GetMetadataType

`func (o *ObjectMetadata) GetMetadataType() BulkWritableCableRequestStatus`

GetMetadataType returns the MetadataType field if non-nil, zero value otherwise.

### GetMetadataTypeOk

`func (o *ObjectMetadata) GetMetadataTypeOk() (*BulkWritableCableRequestStatus, bool)`

GetMetadataTypeOk returns a tuple with the MetadataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataType

`func (o *ObjectMetadata) SetMetadataType(v BulkWritableCableRequestStatus)`

SetMetadataType sets MetadataType field to given value.


### GetContact

`func (o *ObjectMetadata) GetContact() BulkWritableCircuitRequestTenant`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *ObjectMetadata) GetContactOk() (*BulkWritableCircuitRequestTenant, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *ObjectMetadata) SetContact(v BulkWritableCircuitRequestTenant)`

SetContact sets Contact field to given value.

### HasContact

`func (o *ObjectMetadata) HasContact() bool`

HasContact returns a boolean if a field has been set.

### SetContactNil

`func (o *ObjectMetadata) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *ObjectMetadata) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetTeam

`func (o *ObjectMetadata) GetTeam() BulkWritableCircuitRequestTenant`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *ObjectMetadata) GetTeamOk() (*BulkWritableCircuitRequestTenant, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *ObjectMetadata) SetTeam(v BulkWritableCircuitRequestTenant)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *ObjectMetadata) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### SetTeamNil

`func (o *ObjectMetadata) SetTeamNil(b bool)`

 SetTeamNil sets the value for Team to be an explicit nil

### UnsetTeam
`func (o *ObjectMetadata) UnsetTeam()`

UnsetTeam ensures that no value is present for Team, not even an explicit nil
### GetCreated

`func (o *ObjectMetadata) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ObjectMetadata) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ObjectMetadata) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *ObjectMetadata) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *ObjectMetadata) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *ObjectMetadata) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ObjectMetadata) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ObjectMetadata) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *ObjectMetadata) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *ObjectMetadata) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


