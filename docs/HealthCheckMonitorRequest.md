# HealthCheckMonitorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
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

### NewHealthCheckMonitorRequest

`func NewHealthCheckMonitorRequest(name string, ) *HealthCheckMonitorRequest`

NewHealthCheckMonitorRequest instantiates a new HealthCheckMonitorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthCheckMonitorRequestWithDefaults

`func NewHealthCheckMonitorRequestWithDefaults() *HealthCheckMonitorRequest`

NewHealthCheckMonitorRequestWithDefaults instantiates a new HealthCheckMonitorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HealthCheckMonitorRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HealthCheckMonitorRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HealthCheckMonitorRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HealthCheckMonitorRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *HealthCheckMonitorRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HealthCheckMonitorRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HealthCheckMonitorRequest) SetName(v string)`

SetName sets Name field to given value.


### GetInterval

`func (o *HealthCheckMonitorRequest) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *HealthCheckMonitorRequest) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *HealthCheckMonitorRequest) SetInterval(v int32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *HealthCheckMonitorRequest) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### SetIntervalNil

`func (o *HealthCheckMonitorRequest) SetIntervalNil(b bool)`

 SetIntervalNil sets the value for Interval to be an explicit nil

### UnsetInterval
`func (o *HealthCheckMonitorRequest) UnsetInterval()`

UnsetInterval ensures that no value is present for Interval, not even an explicit nil
### GetRetry

`func (o *HealthCheckMonitorRequest) GetRetry() int32`

GetRetry returns the Retry field if non-nil, zero value otherwise.

### GetRetryOk

`func (o *HealthCheckMonitorRequest) GetRetryOk() (*int32, bool)`

GetRetryOk returns a tuple with the Retry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetry

`func (o *HealthCheckMonitorRequest) SetRetry(v int32)`

SetRetry sets Retry field to given value.

### HasRetry

`func (o *HealthCheckMonitorRequest) HasRetry() bool`

HasRetry returns a boolean if a field has been set.

### SetRetryNil

`func (o *HealthCheckMonitorRequest) SetRetryNil(b bool)`

 SetRetryNil sets the value for Retry to be an explicit nil

### UnsetRetry
`func (o *HealthCheckMonitorRequest) UnsetRetry()`

UnsetRetry ensures that no value is present for Retry, not even an explicit nil
### GetTimeout

`func (o *HealthCheckMonitorRequest) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *HealthCheckMonitorRequest) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *HealthCheckMonitorRequest) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *HealthCheckMonitorRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *HealthCheckMonitorRequest) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *HealthCheckMonitorRequest) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetPort

`func (o *HealthCheckMonitorRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *HealthCheckMonitorRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *HealthCheckMonitorRequest) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *HealthCheckMonitorRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *HealthCheckMonitorRequest) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *HealthCheckMonitorRequest) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetHealthCheckType

`func (o *HealthCheckMonitorRequest) GetHealthCheckType() BulkWritableHealthCheckMonitorRequestHealthCheckType`

GetHealthCheckType returns the HealthCheckType field if non-nil, zero value otherwise.

### GetHealthCheckTypeOk

`func (o *HealthCheckMonitorRequest) GetHealthCheckTypeOk() (*BulkWritableHealthCheckMonitorRequestHealthCheckType, bool)`

GetHealthCheckTypeOk returns a tuple with the HealthCheckType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckType

`func (o *HealthCheckMonitorRequest) SetHealthCheckType(v BulkWritableHealthCheckMonitorRequestHealthCheckType)`

SetHealthCheckType sets HealthCheckType field to given value.

### HasHealthCheckType

`func (o *HealthCheckMonitorRequest) HasHealthCheckType() bool`

HasHealthCheckType returns a boolean if a field has been set.

### GetTenant

`func (o *HealthCheckMonitorRequest) GetTenant() ApprovalWorkflowUser`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *HealthCheckMonitorRequest) GetTenantOk() (*ApprovalWorkflowUser, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *HealthCheckMonitorRequest) SetTenant(v ApprovalWorkflowUser)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *HealthCheckMonitorRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *HealthCheckMonitorRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *HealthCheckMonitorRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetCustomFields

`func (o *HealthCheckMonitorRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *HealthCheckMonitorRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *HealthCheckMonitorRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *HealthCheckMonitorRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *HealthCheckMonitorRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *HealthCheckMonitorRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *HealthCheckMonitorRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *HealthCheckMonitorRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *HealthCheckMonitorRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *HealthCheckMonitorRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *HealthCheckMonitorRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *HealthCheckMonitorRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


