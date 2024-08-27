# ContactAssociationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociatedObjectType** | **string** |  | 
**AssociatedObjectId** | **string** |  | 
**Contact** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Team** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Role** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Status** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewContactAssociationRequest

`func NewContactAssociationRequest(associatedObjectType string, associatedObjectId string, role BulkWritableCableRequestStatus, status BulkWritableCableRequestStatus, ) *ContactAssociationRequest`

NewContactAssociationRequest instantiates a new ContactAssociationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactAssociationRequestWithDefaults

`func NewContactAssociationRequestWithDefaults() *ContactAssociationRequest`

NewContactAssociationRequestWithDefaults instantiates a new ContactAssociationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociatedObjectType

`func (o *ContactAssociationRequest) GetAssociatedObjectType() string`

GetAssociatedObjectType returns the AssociatedObjectType field if non-nil, zero value otherwise.

### GetAssociatedObjectTypeOk

`func (o *ContactAssociationRequest) GetAssociatedObjectTypeOk() (*string, bool)`

GetAssociatedObjectTypeOk returns a tuple with the AssociatedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedObjectType

`func (o *ContactAssociationRequest) SetAssociatedObjectType(v string)`

SetAssociatedObjectType sets AssociatedObjectType field to given value.


### GetAssociatedObjectId

`func (o *ContactAssociationRequest) GetAssociatedObjectId() string`

GetAssociatedObjectId returns the AssociatedObjectId field if non-nil, zero value otherwise.

### GetAssociatedObjectIdOk

`func (o *ContactAssociationRequest) GetAssociatedObjectIdOk() (*string, bool)`

GetAssociatedObjectIdOk returns a tuple with the AssociatedObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedObjectId

`func (o *ContactAssociationRequest) SetAssociatedObjectId(v string)`

SetAssociatedObjectId sets AssociatedObjectId field to given value.


### GetContact

`func (o *ContactAssociationRequest) GetContact() BulkWritableCircuitRequestTenant`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *ContactAssociationRequest) GetContactOk() (*BulkWritableCircuitRequestTenant, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *ContactAssociationRequest) SetContact(v BulkWritableCircuitRequestTenant)`

SetContact sets Contact field to given value.

### HasContact

`func (o *ContactAssociationRequest) HasContact() bool`

HasContact returns a boolean if a field has been set.

### SetContactNil

`func (o *ContactAssociationRequest) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *ContactAssociationRequest) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetTeam

`func (o *ContactAssociationRequest) GetTeam() BulkWritableCircuitRequestTenant`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *ContactAssociationRequest) GetTeamOk() (*BulkWritableCircuitRequestTenant, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *ContactAssociationRequest) SetTeam(v BulkWritableCircuitRequestTenant)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *ContactAssociationRequest) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### SetTeamNil

`func (o *ContactAssociationRequest) SetTeamNil(b bool)`

 SetTeamNil sets the value for Team to be an explicit nil

### UnsetTeam
`func (o *ContactAssociationRequest) UnsetTeam()`

UnsetTeam ensures that no value is present for Team, not even an explicit nil
### GetRole

`func (o *ContactAssociationRequest) GetRole() BulkWritableCableRequestStatus`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ContactAssociationRequest) GetRoleOk() (*BulkWritableCableRequestStatus, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ContactAssociationRequest) SetRole(v BulkWritableCableRequestStatus)`

SetRole sets Role field to given value.


### GetStatus

`func (o *ContactAssociationRequest) GetStatus() BulkWritableCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ContactAssociationRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ContactAssociationRequest) SetStatus(v BulkWritableCableRequestStatus)`

SetStatus sets Status field to given value.


### GetCustomFields

`func (o *ContactAssociationRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ContactAssociationRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ContactAssociationRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ContactAssociationRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *ContactAssociationRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ContactAssociationRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ContactAssociationRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ContactAssociationRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


