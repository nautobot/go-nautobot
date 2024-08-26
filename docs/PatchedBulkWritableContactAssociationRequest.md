# PatchedBulkWritableContactAssociationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**AssociatedObjectType** | Pointer to **string** |  | [optional] 
**AssociatedObjectId** | Pointer to **string** |  | [optional] 
**Contact** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Team** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Role** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**Status** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewPatchedBulkWritableContactAssociationRequest

`func NewPatchedBulkWritableContactAssociationRequest(id string, ) *PatchedBulkWritableContactAssociationRequest`

NewPatchedBulkWritableContactAssociationRequest instantiates a new PatchedBulkWritableContactAssociationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableContactAssociationRequestWithDefaults

`func NewPatchedBulkWritableContactAssociationRequestWithDefaults() *PatchedBulkWritableContactAssociationRequest`

NewPatchedBulkWritableContactAssociationRequestWithDefaults instantiates a new PatchedBulkWritableContactAssociationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableContactAssociationRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableContactAssociationRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableContactAssociationRequest) SetId(v string)`

SetId sets Id field to given value.


### GetAssociatedObjectType

`func (o *PatchedBulkWritableContactAssociationRequest) GetAssociatedObjectType() string`

GetAssociatedObjectType returns the AssociatedObjectType field if non-nil, zero value otherwise.

### GetAssociatedObjectTypeOk

`func (o *PatchedBulkWritableContactAssociationRequest) GetAssociatedObjectTypeOk() (*string, bool)`

GetAssociatedObjectTypeOk returns a tuple with the AssociatedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedObjectType

`func (o *PatchedBulkWritableContactAssociationRequest) SetAssociatedObjectType(v string)`

SetAssociatedObjectType sets AssociatedObjectType field to given value.

### HasAssociatedObjectType

`func (o *PatchedBulkWritableContactAssociationRequest) HasAssociatedObjectType() bool`

HasAssociatedObjectType returns a boolean if a field has been set.

### GetAssociatedObjectId

`func (o *PatchedBulkWritableContactAssociationRequest) GetAssociatedObjectId() string`

GetAssociatedObjectId returns the AssociatedObjectId field if non-nil, zero value otherwise.

### GetAssociatedObjectIdOk

`func (o *PatchedBulkWritableContactAssociationRequest) GetAssociatedObjectIdOk() (*string, bool)`

GetAssociatedObjectIdOk returns a tuple with the AssociatedObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedObjectId

`func (o *PatchedBulkWritableContactAssociationRequest) SetAssociatedObjectId(v string)`

SetAssociatedObjectId sets AssociatedObjectId field to given value.

### HasAssociatedObjectId

`func (o *PatchedBulkWritableContactAssociationRequest) HasAssociatedObjectId() bool`

HasAssociatedObjectId returns a boolean if a field has been set.

### GetContact

`func (o *PatchedBulkWritableContactAssociationRequest) GetContact() BulkWritableCircuitRequestTenant`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *PatchedBulkWritableContactAssociationRequest) GetContactOk() (*BulkWritableCircuitRequestTenant, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *PatchedBulkWritableContactAssociationRequest) SetContact(v BulkWritableCircuitRequestTenant)`

SetContact sets Contact field to given value.

### HasContact

`func (o *PatchedBulkWritableContactAssociationRequest) HasContact() bool`

HasContact returns a boolean if a field has been set.

### SetContactNil

`func (o *PatchedBulkWritableContactAssociationRequest) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *PatchedBulkWritableContactAssociationRequest) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetTeam

`func (o *PatchedBulkWritableContactAssociationRequest) GetTeam() BulkWritableCircuitRequestTenant`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *PatchedBulkWritableContactAssociationRequest) GetTeamOk() (*BulkWritableCircuitRequestTenant, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *PatchedBulkWritableContactAssociationRequest) SetTeam(v BulkWritableCircuitRequestTenant)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *PatchedBulkWritableContactAssociationRequest) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### SetTeamNil

`func (o *PatchedBulkWritableContactAssociationRequest) SetTeamNil(b bool)`

 SetTeamNil sets the value for Team to be an explicit nil

### UnsetTeam
`func (o *PatchedBulkWritableContactAssociationRequest) UnsetTeam()`

UnsetTeam ensures that no value is present for Team, not even an explicit nil
### GetRole

`func (o *PatchedBulkWritableContactAssociationRequest) GetRole() BulkWritableCableRequestStatus`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PatchedBulkWritableContactAssociationRequest) GetRoleOk() (*BulkWritableCableRequestStatus, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PatchedBulkWritableContactAssociationRequest) SetRole(v BulkWritableCableRequestStatus)`

SetRole sets Role field to given value.

### HasRole

`func (o *PatchedBulkWritableContactAssociationRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetStatus

`func (o *PatchedBulkWritableContactAssociationRequest) GetStatus() BulkWritableCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedBulkWritableContactAssociationRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedBulkWritableContactAssociationRequest) SetStatus(v BulkWritableCableRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedBulkWritableContactAssociationRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedBulkWritableContactAssociationRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkWritableContactAssociationRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkWritableContactAssociationRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkWritableContactAssociationRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedBulkWritableContactAssociationRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedBulkWritableContactAssociationRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedBulkWritableContactAssociationRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedBulkWritableContactAssociationRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


