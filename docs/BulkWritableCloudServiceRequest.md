# BulkWritableCloudServiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ExtraConfig** | Pointer to **interface{}** |  | [optional] 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**CloudResourceType** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**CloudAccount** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewBulkWritableCloudServiceRequest

`func NewBulkWritableCloudServiceRequest(id string, name string, cloudResourceType BulkWritableCableRequestStatus, ) *BulkWritableCloudServiceRequest`

NewBulkWritableCloudServiceRequest instantiates a new BulkWritableCloudServiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableCloudServiceRequestWithDefaults

`func NewBulkWritableCloudServiceRequestWithDefaults() *BulkWritableCloudServiceRequest`

NewBulkWritableCloudServiceRequestWithDefaults instantiates a new BulkWritableCloudServiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableCloudServiceRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableCloudServiceRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableCloudServiceRequest) SetId(v string)`

SetId sets Id field to given value.


### GetExtraConfig

`func (o *BulkWritableCloudServiceRequest) GetExtraConfig() interface{}`

GetExtraConfig returns the ExtraConfig field if non-nil, zero value otherwise.

### GetExtraConfigOk

`func (o *BulkWritableCloudServiceRequest) GetExtraConfigOk() (*interface{}, bool)`

GetExtraConfigOk returns a tuple with the ExtraConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraConfig

`func (o *BulkWritableCloudServiceRequest) SetExtraConfig(v interface{})`

SetExtraConfig sets ExtraConfig field to given value.

### HasExtraConfig

`func (o *BulkWritableCloudServiceRequest) HasExtraConfig() bool`

HasExtraConfig returns a boolean if a field has been set.

### SetExtraConfigNil

`func (o *BulkWritableCloudServiceRequest) SetExtraConfigNil(b bool)`

 SetExtraConfigNil sets the value for ExtraConfig to be an explicit nil

### UnsetExtraConfig
`func (o *BulkWritableCloudServiceRequest) UnsetExtraConfig()`

UnsetExtraConfig ensures that no value is present for ExtraConfig, not even an explicit nil
### GetName

`func (o *BulkWritableCloudServiceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkWritableCloudServiceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkWritableCloudServiceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *BulkWritableCloudServiceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkWritableCloudServiceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkWritableCloudServiceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkWritableCloudServiceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCloudResourceType

`func (o *BulkWritableCloudServiceRequest) GetCloudResourceType() BulkWritableCableRequestStatus`

GetCloudResourceType returns the CloudResourceType field if non-nil, zero value otherwise.

### GetCloudResourceTypeOk

`func (o *BulkWritableCloudServiceRequest) GetCloudResourceTypeOk() (*BulkWritableCableRequestStatus, bool)`

GetCloudResourceTypeOk returns a tuple with the CloudResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudResourceType

`func (o *BulkWritableCloudServiceRequest) SetCloudResourceType(v BulkWritableCableRequestStatus)`

SetCloudResourceType sets CloudResourceType field to given value.


### GetCloudAccount

`func (o *BulkWritableCloudServiceRequest) GetCloudAccount() BulkWritableCircuitRequestTenant`

GetCloudAccount returns the CloudAccount field if non-nil, zero value otherwise.

### GetCloudAccountOk

`func (o *BulkWritableCloudServiceRequest) GetCloudAccountOk() (*BulkWritableCircuitRequestTenant, bool)`

GetCloudAccountOk returns a tuple with the CloudAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudAccount

`func (o *BulkWritableCloudServiceRequest) SetCloudAccount(v BulkWritableCircuitRequestTenant)`

SetCloudAccount sets CloudAccount field to given value.

### HasCloudAccount

`func (o *BulkWritableCloudServiceRequest) HasCloudAccount() bool`

HasCloudAccount returns a boolean if a field has been set.

### SetCloudAccountNil

`func (o *BulkWritableCloudServiceRequest) SetCloudAccountNil(b bool)`

 SetCloudAccountNil sets the value for CloudAccount to be an explicit nil

### UnsetCloudAccount
`func (o *BulkWritableCloudServiceRequest) UnsetCloudAccount()`

UnsetCloudAccount ensures that no value is present for CloudAccount, not even an explicit nil
### GetCustomFields

`func (o *BulkWritableCloudServiceRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkWritableCloudServiceRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkWritableCloudServiceRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkWritableCloudServiceRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *BulkWritableCloudServiceRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BulkWritableCloudServiceRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BulkWritableCloudServiceRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BulkWritableCloudServiceRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *BulkWritableCloudServiceRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkWritableCloudServiceRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkWritableCloudServiceRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkWritableCloudServiceRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


