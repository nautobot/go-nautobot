# PatchedBulkWritableHealthCheckMonitorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
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

### NewPatchedBulkWritableHealthCheckMonitorRequest

`func NewPatchedBulkWritableHealthCheckMonitorRequest(id string, ) *PatchedBulkWritableHealthCheckMonitorRequest`

NewPatchedBulkWritableHealthCheckMonitorRequest instantiates a new PatchedBulkWritableHealthCheckMonitorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableHealthCheckMonitorRequestWithDefaults

`func NewPatchedBulkWritableHealthCheckMonitorRequestWithDefaults() *PatchedBulkWritableHealthCheckMonitorRequest`

NewPatchedBulkWritableHealthCheckMonitorRequestWithDefaults instantiates a new PatchedBulkWritableHealthCheckMonitorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableHealthCheckMonitorRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableHealthCheckMonitorRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableHealthCheckMonitorRequest) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PatchedBulkWritableHealthCheckMonitorRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkWritableHealthCheckMonitorRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkWritableHealthCheckMonitorRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkWritableHealthCheckMonitorRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetInterval

`func (o *PatchedBulkWritableHealthCheckMonitorRequest) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *PatchedBulkWritableHealthCheckMonitorRequest) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *PatchedBulkWritableHealthCheckMonitorRequest) SetInterval(v int32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *PatchedBulkWritableHealthCheckMonitorRequest) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### SetIntervalNil

`func (o *PatchedBulkWritableHealthCheckMonitorRequest) SetIntervalNil(b bool)`

 SetIntervalNil sets the value for Interval to be an explicit nil

### UnsetInterval
`func (o *PatchedBulkWritableHealthCheckMonitorRequest) UnsetInterval()`

UnsetInterval ensures that no value is present for Interval, not even an explicit nil
### GetRetry

`func (o *PatchedBulkWritableHealthCheckMonitorRequest) GetRetry() int32`

GetRetry returns the Retry field if non-nil, zero value otherwise.

### GetRetryOk

`func (o *PatchedBulkWritableHealthCheckMonitorRequest) GetRetryOk() (*int32, bool)`

GetRetryOk returns a tuple with the Retry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetry

`func (o *PatchedBulkWritableHealthCheckMonitorRequest) SetRetry(v int32)`

SetRetry sets Retry field to given value.

### HasRetry

`func (o *PatchedBulkWritableHealthCheckMonitorRequest) HasRetry() bool`

HasRetry returns a boolean if a field has been set.

### SetRetryNil

`func (o *PatchedBulkWritableHealthCheckMonitorRequest) SetRetryNil(b bool)`

 SetRetryNil sets the value for Retry to be an explicit nil

### UnsetRetry
`func (o *PatchedBulkWritableHealthCheckMonitorRequest) UnsetRetry()`

UnsetRetry ensures that no value is present for Retry, not even an explicit nil
### GetTimeout

`func (o *PatchedBulkWritableHealthCheckMonitorRequest) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *PatchedBulkWritableHealthCheckMonitorRequest) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *PatchedBulkWritableHealthCheckMonitorRequest) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *PatchedBulkWritableHealthCheckMonitorRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *PatchedBulkWritableHealthCheckMonitorRequest) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *PatchedBulkWritableHealthCheckMonitorRequest) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetPort

`func (o *PatchedBulkWritableHealthCheckMonitorRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PatchedBulkWritableHealthCheckMonitorRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PatchedBulkWritableHealthCheckMonitorRequest) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *PatchedBulkWritableHealthCheckMonitorRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *PatchedBulkWritableHealthCheckMonitorRequest) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *PatchedBulkWritableHealthCheckMonitorRequest) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetHealthCheckType

`func (o *PatchedBulkWritableHealthCheckMonitorRequest) GetHealthCheckType() BulkWritableHealthCheckMonitorRequestHealthCheckType`

GetHealthCheckType returns the HealthCheckType field if non-nil, zero value otherwise.

### GetHealthCheckTypeOk

`func (o *PatchedBulkWritableHealthCheckMonitorRequest) GetHealthCheckTypeOk() (*BulkWritableHealthCheckMonitorRequestHealthCheckType, bool)`

GetHealthCheckTypeOk returns a tuple with the HealthCheckType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckType

`func (o *PatchedBulkWritableHealthCheckMonitorRequest) SetHealthCheckType(v BulkWritableHealthCheckMonitorRequestHealthCheckType)`

SetHealthCheckType sets HealthCheckType field to given value.

### HasHealthCheckType

`func (o *PatchedBulkWritableHealthCheckMonitorRequest) HasHealthCheckType() bool`

HasHealthCheckType returns a boolean if a field has been set.

### GetTenant

`func (o *PatchedBulkWritableHealthCheckMonitorRequest) GetTenant() ApprovalWorkflowUser`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedBulkWritableHealthCheckMonitorRequest) GetTenantOk() (*ApprovalWorkflowUser, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedBulkWritableHealthCheckMonitorRequest) SetTenant(v ApprovalWorkflowUser)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedBulkWritableHealthCheckMonitorRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedBulkWritableHealthCheckMonitorRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedBulkWritableHealthCheckMonitorRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetCustomFields

`func (o *PatchedBulkWritableHealthCheckMonitorRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkWritableHealthCheckMonitorRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkWritableHealthCheckMonitorRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkWritableHealthCheckMonitorRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedBulkWritableHealthCheckMonitorRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedBulkWritableHealthCheckMonitorRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedBulkWritableHealthCheckMonitorRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedBulkWritableHealthCheckMonitorRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *PatchedBulkWritableHealthCheckMonitorRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkWritableHealthCheckMonitorRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkWritableHealthCheckMonitorRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkWritableHealthCheckMonitorRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


