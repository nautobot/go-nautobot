# BulkWritableStaticGroupAssociationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**AssociatedObjectType** | **string** |  | 
**AssociatedObjectId** | **string** |  | 
**DynamicGroup** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewBulkWritableStaticGroupAssociationRequest

`func NewBulkWritableStaticGroupAssociationRequest(id string, associatedObjectType string, associatedObjectId string, dynamicGroup BulkWritableCableRequestStatus, ) *BulkWritableStaticGroupAssociationRequest`

NewBulkWritableStaticGroupAssociationRequest instantiates a new BulkWritableStaticGroupAssociationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableStaticGroupAssociationRequestWithDefaults

`func NewBulkWritableStaticGroupAssociationRequestWithDefaults() *BulkWritableStaticGroupAssociationRequest`

NewBulkWritableStaticGroupAssociationRequestWithDefaults instantiates a new BulkWritableStaticGroupAssociationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableStaticGroupAssociationRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableStaticGroupAssociationRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableStaticGroupAssociationRequest) SetId(v string)`

SetId sets Id field to given value.


### GetAssociatedObjectType

`func (o *BulkWritableStaticGroupAssociationRequest) GetAssociatedObjectType() string`

GetAssociatedObjectType returns the AssociatedObjectType field if non-nil, zero value otherwise.

### GetAssociatedObjectTypeOk

`func (o *BulkWritableStaticGroupAssociationRequest) GetAssociatedObjectTypeOk() (*string, bool)`

GetAssociatedObjectTypeOk returns a tuple with the AssociatedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedObjectType

`func (o *BulkWritableStaticGroupAssociationRequest) SetAssociatedObjectType(v string)`

SetAssociatedObjectType sets AssociatedObjectType field to given value.


### GetAssociatedObjectId

`func (o *BulkWritableStaticGroupAssociationRequest) GetAssociatedObjectId() string`

GetAssociatedObjectId returns the AssociatedObjectId field if non-nil, zero value otherwise.

### GetAssociatedObjectIdOk

`func (o *BulkWritableStaticGroupAssociationRequest) GetAssociatedObjectIdOk() (*string, bool)`

GetAssociatedObjectIdOk returns a tuple with the AssociatedObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedObjectId

`func (o *BulkWritableStaticGroupAssociationRequest) SetAssociatedObjectId(v string)`

SetAssociatedObjectId sets AssociatedObjectId field to given value.


### GetDynamicGroup

`func (o *BulkWritableStaticGroupAssociationRequest) GetDynamicGroup() BulkWritableCableRequestStatus`

GetDynamicGroup returns the DynamicGroup field if non-nil, zero value otherwise.

### GetDynamicGroupOk

`func (o *BulkWritableStaticGroupAssociationRequest) GetDynamicGroupOk() (*BulkWritableCableRequestStatus, bool)`

GetDynamicGroupOk returns a tuple with the DynamicGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicGroup

`func (o *BulkWritableStaticGroupAssociationRequest) SetDynamicGroup(v BulkWritableCableRequestStatus)`

SetDynamicGroup sets DynamicGroup field to given value.


### GetCustomFields

`func (o *BulkWritableStaticGroupAssociationRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkWritableStaticGroupAssociationRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkWritableStaticGroupAssociationRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkWritableStaticGroupAssociationRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *BulkWritableStaticGroupAssociationRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BulkWritableStaticGroupAssociationRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BulkWritableStaticGroupAssociationRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BulkWritableStaticGroupAssociationRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


