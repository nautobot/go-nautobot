# LoadBalancerPoolMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**Label** | Pointer to **string** | Optional label for the load balancer pool member. | [optional] 
**Port** | **int32** |  | 
**SslOffload** | Pointer to **bool** |  | [optional] 
**IpAddress** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 
**LoadBalancerPool** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 
**HealthCheckMonitor** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**Tenant** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**Status** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**NotesUrl** | **string** |  | [readonly] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Tags** | Pointer to [**[]ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | [optional] 

## Methods

### NewLoadBalancerPoolMember

`func NewLoadBalancerPoolMember(objectType string, display string, url string, naturalSlug string, port int32, ipAddress ApprovalWorkflowStageResponseApprovalWorkflowStage, loadBalancerPool ApprovalWorkflowStageResponseApprovalWorkflowStage, status ApprovalWorkflowStageResponseApprovalWorkflowStage, created NullableTime, lastUpdated NullableTime, notesUrl string, ) *LoadBalancerPoolMember`

NewLoadBalancerPoolMember instantiates a new LoadBalancerPoolMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadBalancerPoolMemberWithDefaults

`func NewLoadBalancerPoolMemberWithDefaults() *LoadBalancerPoolMember`

NewLoadBalancerPoolMemberWithDefaults instantiates a new LoadBalancerPoolMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LoadBalancerPoolMember) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LoadBalancerPoolMember) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LoadBalancerPoolMember) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LoadBalancerPoolMember) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectType

`func (o *LoadBalancerPoolMember) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *LoadBalancerPoolMember) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *LoadBalancerPoolMember) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *LoadBalancerPoolMember) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *LoadBalancerPoolMember) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *LoadBalancerPoolMember) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *LoadBalancerPoolMember) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *LoadBalancerPoolMember) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *LoadBalancerPoolMember) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *LoadBalancerPoolMember) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *LoadBalancerPoolMember) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *LoadBalancerPoolMember) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetLabel

`func (o *LoadBalancerPoolMember) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *LoadBalancerPoolMember) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *LoadBalancerPoolMember) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *LoadBalancerPoolMember) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetPort

`func (o *LoadBalancerPoolMember) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *LoadBalancerPoolMember) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *LoadBalancerPoolMember) SetPort(v int32)`

SetPort sets Port field to given value.


### GetSslOffload

`func (o *LoadBalancerPoolMember) GetSslOffload() bool`

GetSslOffload returns the SslOffload field if non-nil, zero value otherwise.

### GetSslOffloadOk

`func (o *LoadBalancerPoolMember) GetSslOffloadOk() (*bool, bool)`

GetSslOffloadOk returns a tuple with the SslOffload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslOffload

`func (o *LoadBalancerPoolMember) SetSslOffload(v bool)`

SetSslOffload sets SslOffload field to given value.

### HasSslOffload

`func (o *LoadBalancerPoolMember) HasSslOffload() bool`

HasSslOffload returns a boolean if a field has been set.

### GetIpAddress

`func (o *LoadBalancerPoolMember) GetIpAddress() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *LoadBalancerPoolMember) GetIpAddressOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *LoadBalancerPoolMember) SetIpAddress(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetIpAddress sets IpAddress field to given value.


### GetLoadBalancerPool

`func (o *LoadBalancerPoolMember) GetLoadBalancerPool() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetLoadBalancerPool returns the LoadBalancerPool field if non-nil, zero value otherwise.

### GetLoadBalancerPoolOk

`func (o *LoadBalancerPoolMember) GetLoadBalancerPoolOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetLoadBalancerPoolOk returns a tuple with the LoadBalancerPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerPool

`func (o *LoadBalancerPoolMember) SetLoadBalancerPool(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetLoadBalancerPool sets LoadBalancerPool field to given value.


### GetHealthCheckMonitor

`func (o *LoadBalancerPoolMember) GetHealthCheckMonitor() ApprovalWorkflowUser`

GetHealthCheckMonitor returns the HealthCheckMonitor field if non-nil, zero value otherwise.

### GetHealthCheckMonitorOk

`func (o *LoadBalancerPoolMember) GetHealthCheckMonitorOk() (*ApprovalWorkflowUser, bool)`

GetHealthCheckMonitorOk returns a tuple with the HealthCheckMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckMonitor

`func (o *LoadBalancerPoolMember) SetHealthCheckMonitor(v ApprovalWorkflowUser)`

SetHealthCheckMonitor sets HealthCheckMonitor field to given value.

### HasHealthCheckMonitor

`func (o *LoadBalancerPoolMember) HasHealthCheckMonitor() bool`

HasHealthCheckMonitor returns a boolean if a field has been set.

### SetHealthCheckMonitorNil

`func (o *LoadBalancerPoolMember) SetHealthCheckMonitorNil(b bool)`

 SetHealthCheckMonitorNil sets the value for HealthCheckMonitor to be an explicit nil

### UnsetHealthCheckMonitor
`func (o *LoadBalancerPoolMember) UnsetHealthCheckMonitor()`

UnsetHealthCheckMonitor ensures that no value is present for HealthCheckMonitor, not even an explicit nil
### GetTenant

`func (o *LoadBalancerPoolMember) GetTenant() ApprovalWorkflowUser`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *LoadBalancerPoolMember) GetTenantOk() (*ApprovalWorkflowUser, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *LoadBalancerPoolMember) SetTenant(v ApprovalWorkflowUser)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *LoadBalancerPoolMember) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *LoadBalancerPoolMember) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *LoadBalancerPoolMember) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetStatus

`func (o *LoadBalancerPoolMember) GetStatus() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LoadBalancerPoolMember) GetStatusOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LoadBalancerPoolMember) SetStatus(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetStatus sets Status field to given value.


### GetCreated

`func (o *LoadBalancerPoolMember) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *LoadBalancerPoolMember) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *LoadBalancerPoolMember) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *LoadBalancerPoolMember) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *LoadBalancerPoolMember) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *LoadBalancerPoolMember) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *LoadBalancerPoolMember) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *LoadBalancerPoolMember) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *LoadBalancerPoolMember) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *LoadBalancerPoolMember) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetNotesUrl

`func (o *LoadBalancerPoolMember) GetNotesUrl() string`

GetNotesUrl returns the NotesUrl field if non-nil, zero value otherwise.

### GetNotesUrlOk

`func (o *LoadBalancerPoolMember) GetNotesUrlOk() (*string, bool)`

GetNotesUrlOk returns a tuple with the NotesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesUrl

`func (o *LoadBalancerPoolMember) SetNotesUrl(v string)`

SetNotesUrl sets NotesUrl field to given value.


### GetCustomFields

`func (o *LoadBalancerPoolMember) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *LoadBalancerPoolMember) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *LoadBalancerPoolMember) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *LoadBalancerPoolMember) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetTags

`func (o *LoadBalancerPoolMember) GetTags() []ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *LoadBalancerPoolMember) GetTagsOk() (*[]ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *LoadBalancerPoolMember) SetTags(v []ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetTags sets Tags field to given value.

### HasTags

`func (o *LoadBalancerPoolMember) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


