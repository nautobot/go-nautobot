# BulkWritableControllerManagedDeviceGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** | Name of the controller device group | 
**Weight** | Pointer to **int32** | Weight of the controller device group, used to sort the groups within its parent group | [optional] 
**Parent** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Controller** | [**BulkWritableControllerManagedDeviceGroupRequestController**](BulkWritableControllerManagedDeviceGroupRequestController.md) |  | 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewBulkWritableControllerManagedDeviceGroupRequest

`func NewBulkWritableControllerManagedDeviceGroupRequest(id string, name string, controller BulkWritableControllerManagedDeviceGroupRequestController, ) *BulkWritableControllerManagedDeviceGroupRequest`

NewBulkWritableControllerManagedDeviceGroupRequest instantiates a new BulkWritableControllerManagedDeviceGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableControllerManagedDeviceGroupRequestWithDefaults

`func NewBulkWritableControllerManagedDeviceGroupRequestWithDefaults() *BulkWritableControllerManagedDeviceGroupRequest`

NewBulkWritableControllerManagedDeviceGroupRequestWithDefaults instantiates a new BulkWritableControllerManagedDeviceGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableControllerManagedDeviceGroupRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableControllerManagedDeviceGroupRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableControllerManagedDeviceGroupRequest) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BulkWritableControllerManagedDeviceGroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkWritableControllerManagedDeviceGroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkWritableControllerManagedDeviceGroupRequest) SetName(v string)`

SetName sets Name field to given value.


### GetWeight

`func (o *BulkWritableControllerManagedDeviceGroupRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *BulkWritableControllerManagedDeviceGroupRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *BulkWritableControllerManagedDeviceGroupRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *BulkWritableControllerManagedDeviceGroupRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetParent

`func (o *BulkWritableControllerManagedDeviceGroupRequest) GetParent() BulkWritableCircuitRequestTenant`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *BulkWritableControllerManagedDeviceGroupRequest) GetParentOk() (*BulkWritableCircuitRequestTenant, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *BulkWritableControllerManagedDeviceGroupRequest) SetParent(v BulkWritableCircuitRequestTenant)`

SetParent sets Parent field to given value.

### HasParent

`func (o *BulkWritableControllerManagedDeviceGroupRequest) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *BulkWritableControllerManagedDeviceGroupRequest) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *BulkWritableControllerManagedDeviceGroupRequest) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetController

`func (o *BulkWritableControllerManagedDeviceGroupRequest) GetController() BulkWritableControllerManagedDeviceGroupRequestController`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *BulkWritableControllerManagedDeviceGroupRequest) GetControllerOk() (*BulkWritableControllerManagedDeviceGroupRequestController, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *BulkWritableControllerManagedDeviceGroupRequest) SetController(v BulkWritableControllerManagedDeviceGroupRequestController)`

SetController sets Controller field to given value.


### GetTags

`func (o *BulkWritableControllerManagedDeviceGroupRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkWritableControllerManagedDeviceGroupRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkWritableControllerManagedDeviceGroupRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkWritableControllerManagedDeviceGroupRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *BulkWritableControllerManagedDeviceGroupRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkWritableControllerManagedDeviceGroupRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkWritableControllerManagedDeviceGroupRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkWritableControllerManagedDeviceGroupRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *BulkWritableControllerManagedDeviceGroupRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BulkWritableControllerManagedDeviceGroupRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BulkWritableControllerManagedDeviceGroupRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BulkWritableControllerManagedDeviceGroupRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


