# ClusterGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewClusterGroupRequest

`func NewClusterGroupRequest(name string, ) *ClusterGroupRequest`

NewClusterGroupRequest instantiates a new ClusterGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterGroupRequestWithDefaults

`func NewClusterGroupRequestWithDefaults() *ClusterGroupRequest`

NewClusterGroupRequestWithDefaults instantiates a new ClusterGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ClusterGroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterGroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterGroupRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ClusterGroupRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClusterGroupRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClusterGroupRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ClusterGroupRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCustomFields

`func (o *ClusterGroupRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ClusterGroupRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ClusterGroupRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ClusterGroupRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *ClusterGroupRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ClusterGroupRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ClusterGroupRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ClusterGroupRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

