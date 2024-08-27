# CloudServiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtraConfig** | Pointer to **interface{}** |  | [optional] 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**CloudResourceType** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**CloudAccount** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewCloudServiceRequest

`func NewCloudServiceRequest(name string, cloudResourceType BulkWritableCableRequestStatus, ) *CloudServiceRequest`

NewCloudServiceRequest instantiates a new CloudServiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudServiceRequestWithDefaults

`func NewCloudServiceRequestWithDefaults() *CloudServiceRequest`

NewCloudServiceRequestWithDefaults instantiates a new CloudServiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtraConfig

`func (o *CloudServiceRequest) GetExtraConfig() interface{}`

GetExtraConfig returns the ExtraConfig field if non-nil, zero value otherwise.

### GetExtraConfigOk

`func (o *CloudServiceRequest) GetExtraConfigOk() (*interface{}, bool)`

GetExtraConfigOk returns a tuple with the ExtraConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraConfig

`func (o *CloudServiceRequest) SetExtraConfig(v interface{})`

SetExtraConfig sets ExtraConfig field to given value.

### HasExtraConfig

`func (o *CloudServiceRequest) HasExtraConfig() bool`

HasExtraConfig returns a boolean if a field has been set.

### SetExtraConfigNil

`func (o *CloudServiceRequest) SetExtraConfigNil(b bool)`

 SetExtraConfigNil sets the value for ExtraConfig to be an explicit nil

### UnsetExtraConfig
`func (o *CloudServiceRequest) UnsetExtraConfig()`

UnsetExtraConfig ensures that no value is present for ExtraConfig, not even an explicit nil
### GetName

`func (o *CloudServiceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudServiceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudServiceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CloudServiceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudServiceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudServiceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudServiceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCloudResourceType

`func (o *CloudServiceRequest) GetCloudResourceType() BulkWritableCableRequestStatus`

GetCloudResourceType returns the CloudResourceType field if non-nil, zero value otherwise.

### GetCloudResourceTypeOk

`func (o *CloudServiceRequest) GetCloudResourceTypeOk() (*BulkWritableCableRequestStatus, bool)`

GetCloudResourceTypeOk returns a tuple with the CloudResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudResourceType

`func (o *CloudServiceRequest) SetCloudResourceType(v BulkWritableCableRequestStatus)`

SetCloudResourceType sets CloudResourceType field to given value.


### GetCloudAccount

`func (o *CloudServiceRequest) GetCloudAccount() BulkWritableCircuitRequestTenant`

GetCloudAccount returns the CloudAccount field if non-nil, zero value otherwise.

### GetCloudAccountOk

`func (o *CloudServiceRequest) GetCloudAccountOk() (*BulkWritableCircuitRequestTenant, bool)`

GetCloudAccountOk returns a tuple with the CloudAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudAccount

`func (o *CloudServiceRequest) SetCloudAccount(v BulkWritableCircuitRequestTenant)`

SetCloudAccount sets CloudAccount field to given value.

### HasCloudAccount

`func (o *CloudServiceRequest) HasCloudAccount() bool`

HasCloudAccount returns a boolean if a field has been set.

### SetCloudAccountNil

`func (o *CloudServiceRequest) SetCloudAccountNil(b bool)`

 SetCloudAccountNil sets the value for CloudAccount to be an explicit nil

### UnsetCloudAccount
`func (o *CloudServiceRequest) UnsetCloudAccount()`

UnsetCloudAccount ensures that no value is present for CloudAccount, not even an explicit nil
### GetCustomFields

`func (o *CloudServiceRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CloudServiceRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CloudServiceRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CloudServiceRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *CloudServiceRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CloudServiceRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CloudServiceRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *CloudServiceRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *CloudServiceRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CloudServiceRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CloudServiceRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CloudServiceRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


