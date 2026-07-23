# BulkWritableHealthCheckMonitorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Interval** | Pointer to **NullableInt32** |  | [optional] 
**Retry** | Pointer to **NullableInt32** | Number of retries before marking as down | [optional] 
**Timeout** | Pointer to **NullableInt32** |  | [optional] 
**Port** | Pointer to **NullableInt32** |  | [optional] 
**HealthCheckType** | Pointer to [**BulkWritableHealthCheckMonitorRequestHealthCheckType**](BulkWritableHealthCheckMonitorRequestHealthCheckType.md) |  | [optional] 
**Tenant** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewBulkWritableHealthCheckMonitorRequest

`func NewBulkWritableHealthCheckMonitorRequest(id string, name string, ) *BulkWritableHealthCheckMonitorRequest`

NewBulkWritableHealthCheckMonitorRequest instantiates a new BulkWritableHealthCheckMonitorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableHealthCheckMonitorRequestWithDefaults

`func NewBulkWritableHealthCheckMonitorRequestWithDefaults() *BulkWritableHealthCheckMonitorRequest`

NewBulkWritableHealthCheckMonitorRequestWithDefaults instantiates a new BulkWritableHealthCheckMonitorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableHealthCheckMonitorRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableHealthCheckMonitorRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableHealthCheckMonitorRequest) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BulkWritableHealthCheckMonitorRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkWritableHealthCheckMonitorRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkWritableHealthCheckMonitorRequest) SetName(v string)`

SetName sets Name field to given value.


### GetInterval

`func (o *BulkWritableHealthCheckMonitorRequest) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *BulkWritableHealthCheckMonitorRequest) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *BulkWritableHealthCheckMonitorRequest) SetInterval(v int32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *BulkWritableHealthCheckMonitorRequest) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### SetIntervalNil

`func (o *BulkWritableHealthCheckMonitorRequest) SetIntervalNil(b bool)`

 SetIntervalNil sets the value for Interval to be an explicit nil

### UnsetInterval
`func (o *BulkWritableHealthCheckMonitorRequest) UnsetInterval()`

UnsetInterval ensures that no value is present for Interval, not even an explicit nil
### GetRetry

`func (o *BulkWritableHealthCheckMonitorRequest) GetRetry() int32`

GetRetry returns the Retry field if non-nil, zero value otherwise.

### GetRetryOk

`func (o *BulkWritableHealthCheckMonitorRequest) GetRetryOk() (*int32, bool)`

GetRetryOk returns a tuple with the Retry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetry

`func (o *BulkWritableHealthCheckMonitorRequest) SetRetry(v int32)`

SetRetry sets Retry field to given value.

### HasRetry

`func (o *BulkWritableHealthCheckMonitorRequest) HasRetry() bool`

HasRetry returns a boolean if a field has been set.

### SetRetryNil

`func (o *BulkWritableHealthCheckMonitorRequest) SetRetryNil(b bool)`

 SetRetryNil sets the value for Retry to be an explicit nil

### UnsetRetry
`func (o *BulkWritableHealthCheckMonitorRequest) UnsetRetry()`

UnsetRetry ensures that no value is present for Retry, not even an explicit nil
### GetTimeout

`func (o *BulkWritableHealthCheckMonitorRequest) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *BulkWritableHealthCheckMonitorRequest) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *BulkWritableHealthCheckMonitorRequest) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *BulkWritableHealthCheckMonitorRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *BulkWritableHealthCheckMonitorRequest) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *BulkWritableHealthCheckMonitorRequest) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetPort

`func (o *BulkWritableHealthCheckMonitorRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *BulkWritableHealthCheckMonitorRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *BulkWritableHealthCheckMonitorRequest) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *BulkWritableHealthCheckMonitorRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *BulkWritableHealthCheckMonitorRequest) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *BulkWritableHealthCheckMonitorRequest) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetHealthCheckType

`func (o *BulkWritableHealthCheckMonitorRequest) GetHealthCheckType() BulkWritableHealthCheckMonitorRequestHealthCheckType`

GetHealthCheckType returns the HealthCheckType field if non-nil, zero value otherwise.

### GetHealthCheckTypeOk

`func (o *BulkWritableHealthCheckMonitorRequest) GetHealthCheckTypeOk() (*BulkWritableHealthCheckMonitorRequestHealthCheckType, bool)`

GetHealthCheckTypeOk returns a tuple with the HealthCheckType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckType

`func (o *BulkWritableHealthCheckMonitorRequest) SetHealthCheckType(v BulkWritableHealthCheckMonitorRequestHealthCheckType)`

SetHealthCheckType sets HealthCheckType field to given value.

### HasHealthCheckType

`func (o *BulkWritableHealthCheckMonitorRequest) HasHealthCheckType() bool`

HasHealthCheckType returns a boolean if a field has been set.

### GetTenant

`func (o *BulkWritableHealthCheckMonitorRequest) GetTenant() ApprovalWorkflowUser`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *BulkWritableHealthCheckMonitorRequest) GetTenantOk() (*ApprovalWorkflowUser, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *BulkWritableHealthCheckMonitorRequest) SetTenant(v ApprovalWorkflowUser)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *BulkWritableHealthCheckMonitorRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *BulkWritableHealthCheckMonitorRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *BulkWritableHealthCheckMonitorRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetCustomFields

`func (o *BulkWritableHealthCheckMonitorRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkWritableHealthCheckMonitorRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkWritableHealthCheckMonitorRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkWritableHealthCheckMonitorRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *BulkWritableHealthCheckMonitorRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BulkWritableHealthCheckMonitorRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BulkWritableHealthCheckMonitorRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BulkWritableHealthCheckMonitorRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *BulkWritableHealthCheckMonitorRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkWritableHealthCheckMonitorRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkWritableHealthCheckMonitorRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkWritableHealthCheckMonitorRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


