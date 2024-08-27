# StaticGroupAssociationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociatedObjectType** | **string** |  | 
**AssociatedObjectId** | **string** |  | 
**DynamicGroup** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewStaticGroupAssociationRequest

`func NewStaticGroupAssociationRequest(associatedObjectType string, associatedObjectId string, dynamicGroup BulkWritableCableRequestStatus, ) *StaticGroupAssociationRequest`

NewStaticGroupAssociationRequest instantiates a new StaticGroupAssociationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStaticGroupAssociationRequestWithDefaults

`func NewStaticGroupAssociationRequestWithDefaults() *StaticGroupAssociationRequest`

NewStaticGroupAssociationRequestWithDefaults instantiates a new StaticGroupAssociationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociatedObjectType

`func (o *StaticGroupAssociationRequest) GetAssociatedObjectType() string`

GetAssociatedObjectType returns the AssociatedObjectType field if non-nil, zero value otherwise.

### GetAssociatedObjectTypeOk

`func (o *StaticGroupAssociationRequest) GetAssociatedObjectTypeOk() (*string, bool)`

GetAssociatedObjectTypeOk returns a tuple with the AssociatedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedObjectType

`func (o *StaticGroupAssociationRequest) SetAssociatedObjectType(v string)`

SetAssociatedObjectType sets AssociatedObjectType field to given value.


### GetAssociatedObjectId

`func (o *StaticGroupAssociationRequest) GetAssociatedObjectId() string`

GetAssociatedObjectId returns the AssociatedObjectId field if non-nil, zero value otherwise.

### GetAssociatedObjectIdOk

`func (o *StaticGroupAssociationRequest) GetAssociatedObjectIdOk() (*string, bool)`

GetAssociatedObjectIdOk returns a tuple with the AssociatedObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedObjectId

`func (o *StaticGroupAssociationRequest) SetAssociatedObjectId(v string)`

SetAssociatedObjectId sets AssociatedObjectId field to given value.


### GetDynamicGroup

`func (o *StaticGroupAssociationRequest) GetDynamicGroup() BulkWritableCableRequestStatus`

GetDynamicGroup returns the DynamicGroup field if non-nil, zero value otherwise.

### GetDynamicGroupOk

`func (o *StaticGroupAssociationRequest) GetDynamicGroupOk() (*BulkWritableCableRequestStatus, bool)`

GetDynamicGroupOk returns a tuple with the DynamicGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicGroup

`func (o *StaticGroupAssociationRequest) SetDynamicGroup(v BulkWritableCableRequestStatus)`

SetDynamicGroup sets DynamicGroup field to given value.


### GetCustomFields

`func (o *StaticGroupAssociationRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *StaticGroupAssociationRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *StaticGroupAssociationRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *StaticGroupAssociationRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *StaticGroupAssociationRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *StaticGroupAssociationRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *StaticGroupAssociationRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *StaticGroupAssociationRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


