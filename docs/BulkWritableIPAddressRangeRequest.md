# BulkWritableIPAddressRangeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**StartAddress** | **string** |  | 
**EndAddress** | **string** |  | 
**Namespace** | Pointer to [**BulkWritableIPAddressRangeRequestNamespace**](BulkWritableIPAddressRangeRequestNamespace.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the IP Address Range | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**CountAsUtilized** | Pointer to **bool** | Forces this range to count as fully utilized in prefix utilization calculations. | [optional] 
**IsExclusive** | Pointer to **bool** | Prevent individual IP Address objects from being created within this range. | [optional] 
**Parent** | Pointer to [**BulkWritableIPAddressRangeRequestParent**](BulkWritableIPAddressRangeRequestParent.md) |  | [optional] 
**Status** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Role** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**Tenant** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** | Custom field data for this object, keyed by each applicable Custom Field&#39;s &#x60;key&#x60;. Value types vary with the custom field&#39;s type (text, integer, boolean, date, URL, JSON, select, multi-select); undefined values are &#x60;null&#x60;. On write, the payload is merged with existing values (PATCH-style: keys omitted from the payload are left untouched), and keys that do not correspond to a defined custom field are ignored. | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewBulkWritableIPAddressRangeRequest

`func NewBulkWritableIPAddressRangeRequest(id string, startAddress string, endAddress string, status BulkWritableCableRequestStatus, ) *BulkWritableIPAddressRangeRequest`

NewBulkWritableIPAddressRangeRequest instantiates a new BulkWritableIPAddressRangeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableIPAddressRangeRequestWithDefaults

`func NewBulkWritableIPAddressRangeRequestWithDefaults() *BulkWritableIPAddressRangeRequest`

NewBulkWritableIPAddressRangeRequestWithDefaults instantiates a new BulkWritableIPAddressRangeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableIPAddressRangeRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableIPAddressRangeRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableIPAddressRangeRequest) SetId(v string)`

SetId sets Id field to given value.


### GetStartAddress

`func (o *BulkWritableIPAddressRangeRequest) GetStartAddress() string`

GetStartAddress returns the StartAddress field if non-nil, zero value otherwise.

### GetStartAddressOk

`func (o *BulkWritableIPAddressRangeRequest) GetStartAddressOk() (*string, bool)`

GetStartAddressOk returns a tuple with the StartAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAddress

`func (o *BulkWritableIPAddressRangeRequest) SetStartAddress(v string)`

SetStartAddress sets StartAddress field to given value.


### GetEndAddress

`func (o *BulkWritableIPAddressRangeRequest) GetEndAddress() string`

GetEndAddress returns the EndAddress field if non-nil, zero value otherwise.

### GetEndAddressOk

`func (o *BulkWritableIPAddressRangeRequest) GetEndAddressOk() (*string, bool)`

GetEndAddressOk returns a tuple with the EndAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAddress

`func (o *BulkWritableIPAddressRangeRequest) SetEndAddress(v string)`

SetEndAddress sets EndAddress field to given value.


### GetNamespace

`func (o *BulkWritableIPAddressRangeRequest) GetNamespace() BulkWritableIPAddressRangeRequestNamespace`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *BulkWritableIPAddressRangeRequest) GetNamespaceOk() (*BulkWritableIPAddressRangeRequestNamespace, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *BulkWritableIPAddressRangeRequest) SetNamespace(v BulkWritableIPAddressRangeRequestNamespace)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *BulkWritableIPAddressRangeRequest) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetName

`func (o *BulkWritableIPAddressRangeRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkWritableIPAddressRangeRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkWritableIPAddressRangeRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BulkWritableIPAddressRangeRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *BulkWritableIPAddressRangeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkWritableIPAddressRangeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkWritableIPAddressRangeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkWritableIPAddressRangeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCountAsUtilized

`func (o *BulkWritableIPAddressRangeRequest) GetCountAsUtilized() bool`

GetCountAsUtilized returns the CountAsUtilized field if non-nil, zero value otherwise.

### GetCountAsUtilizedOk

`func (o *BulkWritableIPAddressRangeRequest) GetCountAsUtilizedOk() (*bool, bool)`

GetCountAsUtilizedOk returns a tuple with the CountAsUtilized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountAsUtilized

`func (o *BulkWritableIPAddressRangeRequest) SetCountAsUtilized(v bool)`

SetCountAsUtilized sets CountAsUtilized field to given value.

### HasCountAsUtilized

`func (o *BulkWritableIPAddressRangeRequest) HasCountAsUtilized() bool`

HasCountAsUtilized returns a boolean if a field has been set.

### GetIsExclusive

`func (o *BulkWritableIPAddressRangeRequest) GetIsExclusive() bool`

GetIsExclusive returns the IsExclusive field if non-nil, zero value otherwise.

### GetIsExclusiveOk

`func (o *BulkWritableIPAddressRangeRequest) GetIsExclusiveOk() (*bool, bool)`

GetIsExclusiveOk returns a tuple with the IsExclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExclusive

`func (o *BulkWritableIPAddressRangeRequest) SetIsExclusive(v bool)`

SetIsExclusive sets IsExclusive field to given value.

### HasIsExclusive

`func (o *BulkWritableIPAddressRangeRequest) HasIsExclusive() bool`

HasIsExclusive returns a boolean if a field has been set.

### GetParent

`func (o *BulkWritableIPAddressRangeRequest) GetParent() BulkWritableIPAddressRangeRequestParent`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *BulkWritableIPAddressRangeRequest) GetParentOk() (*BulkWritableIPAddressRangeRequestParent, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *BulkWritableIPAddressRangeRequest) SetParent(v BulkWritableIPAddressRangeRequestParent)`

SetParent sets Parent field to given value.

### HasParent

`func (o *BulkWritableIPAddressRangeRequest) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetStatus

`func (o *BulkWritableIPAddressRangeRequest) GetStatus() BulkWritableCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkWritableIPAddressRangeRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkWritableIPAddressRangeRequest) SetStatus(v BulkWritableCableRequestStatus)`

SetStatus sets Status field to given value.


### GetRole

`func (o *BulkWritableIPAddressRangeRequest) GetRole() ApprovalWorkflowUser`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *BulkWritableIPAddressRangeRequest) GetRoleOk() (*ApprovalWorkflowUser, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *BulkWritableIPAddressRangeRequest) SetRole(v ApprovalWorkflowUser)`

SetRole sets Role field to given value.

### HasRole

`func (o *BulkWritableIPAddressRangeRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *BulkWritableIPAddressRangeRequest) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *BulkWritableIPAddressRangeRequest) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetTenant

`func (o *BulkWritableIPAddressRangeRequest) GetTenant() ApprovalWorkflowUser`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *BulkWritableIPAddressRangeRequest) GetTenantOk() (*ApprovalWorkflowUser, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *BulkWritableIPAddressRangeRequest) SetTenant(v ApprovalWorkflowUser)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *BulkWritableIPAddressRangeRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *BulkWritableIPAddressRangeRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *BulkWritableIPAddressRangeRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetTags

`func (o *BulkWritableIPAddressRangeRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BulkWritableIPAddressRangeRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BulkWritableIPAddressRangeRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BulkWritableIPAddressRangeRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *BulkWritableIPAddressRangeRequest) GetCustomFields() map[string]*interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkWritableIPAddressRangeRequest) GetCustomFieldsOk() (*map[string]*interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkWritableIPAddressRangeRequest) SetCustomFields(v map[string]*interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkWritableIPAddressRangeRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *BulkWritableIPAddressRangeRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BulkWritableIPAddressRangeRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BulkWritableIPAddressRangeRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BulkWritableIPAddressRangeRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


