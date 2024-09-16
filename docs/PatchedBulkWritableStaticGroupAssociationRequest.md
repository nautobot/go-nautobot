# PatchedBulkWritableStaticGroupAssociationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**AssociatedObjectType** | Pointer to **string** |  | [optional] 
**AssociatedObjectId** | Pointer to **string** |  | [optional] 
**DynamicGroup** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewPatchedBulkWritableStaticGroupAssociationRequest

`func NewPatchedBulkWritableStaticGroupAssociationRequest(id string, ) *PatchedBulkWritableStaticGroupAssociationRequest`

NewPatchedBulkWritableStaticGroupAssociationRequest instantiates a new PatchedBulkWritableStaticGroupAssociationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableStaticGroupAssociationRequestWithDefaults

`func NewPatchedBulkWritableStaticGroupAssociationRequestWithDefaults() *PatchedBulkWritableStaticGroupAssociationRequest`

NewPatchedBulkWritableStaticGroupAssociationRequestWithDefaults instantiates a new PatchedBulkWritableStaticGroupAssociationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableStaticGroupAssociationRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableStaticGroupAssociationRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableStaticGroupAssociationRequest) SetId(v string)`

SetId sets Id field to given value.


### GetAssociatedObjectType

`func (o *PatchedBulkWritableStaticGroupAssociationRequest) GetAssociatedObjectType() string`

GetAssociatedObjectType returns the AssociatedObjectType field if non-nil, zero value otherwise.

### GetAssociatedObjectTypeOk

`func (o *PatchedBulkWritableStaticGroupAssociationRequest) GetAssociatedObjectTypeOk() (*string, bool)`

GetAssociatedObjectTypeOk returns a tuple with the AssociatedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedObjectType

`func (o *PatchedBulkWritableStaticGroupAssociationRequest) SetAssociatedObjectType(v string)`

SetAssociatedObjectType sets AssociatedObjectType field to given value.

### HasAssociatedObjectType

`func (o *PatchedBulkWritableStaticGroupAssociationRequest) HasAssociatedObjectType() bool`

HasAssociatedObjectType returns a boolean if a field has been set.

### GetAssociatedObjectId

`func (o *PatchedBulkWritableStaticGroupAssociationRequest) GetAssociatedObjectId() string`

GetAssociatedObjectId returns the AssociatedObjectId field if non-nil, zero value otherwise.

### GetAssociatedObjectIdOk

`func (o *PatchedBulkWritableStaticGroupAssociationRequest) GetAssociatedObjectIdOk() (*string, bool)`

GetAssociatedObjectIdOk returns a tuple with the AssociatedObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedObjectId

`func (o *PatchedBulkWritableStaticGroupAssociationRequest) SetAssociatedObjectId(v string)`

SetAssociatedObjectId sets AssociatedObjectId field to given value.

### HasAssociatedObjectId

`func (o *PatchedBulkWritableStaticGroupAssociationRequest) HasAssociatedObjectId() bool`

HasAssociatedObjectId returns a boolean if a field has been set.

### GetDynamicGroup

`func (o *PatchedBulkWritableStaticGroupAssociationRequest) GetDynamicGroup() BulkWritableCableRequestStatus`

GetDynamicGroup returns the DynamicGroup field if non-nil, zero value otherwise.

### GetDynamicGroupOk

`func (o *PatchedBulkWritableStaticGroupAssociationRequest) GetDynamicGroupOk() (*BulkWritableCableRequestStatus, bool)`

GetDynamicGroupOk returns a tuple with the DynamicGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicGroup

`func (o *PatchedBulkWritableStaticGroupAssociationRequest) SetDynamicGroup(v BulkWritableCableRequestStatus)`

SetDynamicGroup sets DynamicGroup field to given value.

### HasDynamicGroup

`func (o *PatchedBulkWritableStaticGroupAssociationRequest) HasDynamicGroup() bool`

HasDynamicGroup returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedBulkWritableStaticGroupAssociationRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkWritableStaticGroupAssociationRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkWritableStaticGroupAssociationRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkWritableStaticGroupAssociationRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedBulkWritableStaticGroupAssociationRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedBulkWritableStaticGroupAssociationRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedBulkWritableStaticGroupAssociationRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedBulkWritableStaticGroupAssociationRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


