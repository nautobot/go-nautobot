# PatchedHealthCheckMonitorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
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

### NewPatchedHealthCheckMonitorRequest

`func NewPatchedHealthCheckMonitorRequest() *PatchedHealthCheckMonitorRequest`

NewPatchedHealthCheckMonitorRequest instantiates a new PatchedHealthCheckMonitorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedHealthCheckMonitorRequestWithDefaults

`func NewPatchedHealthCheckMonitorRequestWithDefaults() *PatchedHealthCheckMonitorRequest`

NewPatchedHealthCheckMonitorRequestWithDefaults instantiates a new PatchedHealthCheckMonitorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedHealthCheckMonitorRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedHealthCheckMonitorRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedHealthCheckMonitorRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedHealthCheckMonitorRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PatchedHealthCheckMonitorRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedHealthCheckMonitorRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedHealthCheckMonitorRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedHealthCheckMonitorRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetInterval

`func (o *PatchedHealthCheckMonitorRequest) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *PatchedHealthCheckMonitorRequest) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *PatchedHealthCheckMonitorRequest) SetInterval(v int32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *PatchedHealthCheckMonitorRequest) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### SetIntervalNil

`func (o *PatchedHealthCheckMonitorRequest) SetIntervalNil(b bool)`

 SetIntervalNil sets the value for Interval to be an explicit nil

### UnsetInterval
`func (o *PatchedHealthCheckMonitorRequest) UnsetInterval()`

UnsetInterval ensures that no value is present for Interval, not even an explicit nil
### GetRetry

`func (o *PatchedHealthCheckMonitorRequest) GetRetry() int32`

GetRetry returns the Retry field if non-nil, zero value otherwise.

### GetRetryOk

`func (o *PatchedHealthCheckMonitorRequest) GetRetryOk() (*int32, bool)`

GetRetryOk returns a tuple with the Retry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetry

`func (o *PatchedHealthCheckMonitorRequest) SetRetry(v int32)`

SetRetry sets Retry field to given value.

### HasRetry

`func (o *PatchedHealthCheckMonitorRequest) HasRetry() bool`

HasRetry returns a boolean if a field has been set.

### SetRetryNil

`func (o *PatchedHealthCheckMonitorRequest) SetRetryNil(b bool)`

 SetRetryNil sets the value for Retry to be an explicit nil

### UnsetRetry
`func (o *PatchedHealthCheckMonitorRequest) UnsetRetry()`

UnsetRetry ensures that no value is present for Retry, not even an explicit nil
### GetTimeout

`func (o *PatchedHealthCheckMonitorRequest) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *PatchedHealthCheckMonitorRequest) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *PatchedHealthCheckMonitorRequest) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *PatchedHealthCheckMonitorRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *PatchedHealthCheckMonitorRequest) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *PatchedHealthCheckMonitorRequest) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil
### GetPort

`func (o *PatchedHealthCheckMonitorRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PatchedHealthCheckMonitorRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PatchedHealthCheckMonitorRequest) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *PatchedHealthCheckMonitorRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *PatchedHealthCheckMonitorRequest) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *PatchedHealthCheckMonitorRequest) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetHealthCheckType

`func (o *PatchedHealthCheckMonitorRequest) GetHealthCheckType() BulkWritableHealthCheckMonitorRequestHealthCheckType`

GetHealthCheckType returns the HealthCheckType field if non-nil, zero value otherwise.

### GetHealthCheckTypeOk

`func (o *PatchedHealthCheckMonitorRequest) GetHealthCheckTypeOk() (*BulkWritableHealthCheckMonitorRequestHealthCheckType, bool)`

GetHealthCheckTypeOk returns a tuple with the HealthCheckType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckType

`func (o *PatchedHealthCheckMonitorRequest) SetHealthCheckType(v BulkWritableHealthCheckMonitorRequestHealthCheckType)`

SetHealthCheckType sets HealthCheckType field to given value.

### HasHealthCheckType

`func (o *PatchedHealthCheckMonitorRequest) HasHealthCheckType() bool`

HasHealthCheckType returns a boolean if a field has been set.

### GetTenant

`func (o *PatchedHealthCheckMonitorRequest) GetTenant() ApprovalWorkflowUser`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedHealthCheckMonitorRequest) GetTenantOk() (*ApprovalWorkflowUser, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedHealthCheckMonitorRequest) SetTenant(v ApprovalWorkflowUser)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedHealthCheckMonitorRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedHealthCheckMonitorRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedHealthCheckMonitorRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetCustomFields

`func (o *PatchedHealthCheckMonitorRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedHealthCheckMonitorRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedHealthCheckMonitorRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedHealthCheckMonitorRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedHealthCheckMonitorRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedHealthCheckMonitorRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedHealthCheckMonitorRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedHealthCheckMonitorRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *PatchedHealthCheckMonitorRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedHealthCheckMonitorRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedHealthCheckMonitorRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedHealthCheckMonitorRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


