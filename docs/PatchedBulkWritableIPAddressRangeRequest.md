# PatchedBulkWritableIPAddressRangeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**StartAddress** | Pointer to **string** |  | [optional] 
**EndAddress** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to [**BulkWritableIPAddressRangeRequestNamespace**](BulkWritableIPAddressRangeRequestNamespace.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the IP Address Range | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**CountAsUtilized** | Pointer to **bool** | Forces this range to count as fully utilized in prefix utilization calculations. | [optional] 
**IsExclusive** | Pointer to **bool** | Prevent individual IP Address objects from being created within this range. | [optional] 
**Parent** | Pointer to [**BulkWritableIPAddressRangeRequestParent**](BulkWritableIPAddressRangeRequestParent.md) |  | [optional] 
**Status** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**Role** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**Tenant** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** | Custom field data for this object, keyed by each applicable Custom Field&#39;s &#x60;key&#x60;. Value types vary with the custom field&#39;s type (text, integer, boolean, date, URL, JSON, select, multi-select); undefined values are &#x60;null&#x60;. On write, the payload is merged with existing values (PATCH-style: keys omitted from the payload are left untouched), and keys that do not correspond to a defined custom field are ignored. | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewPatchedBulkWritableIPAddressRangeRequest

`func NewPatchedBulkWritableIPAddressRangeRequest(id string, ) *PatchedBulkWritableIPAddressRangeRequest`

NewPatchedBulkWritableIPAddressRangeRequest instantiates a new PatchedBulkWritableIPAddressRangeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableIPAddressRangeRequestWithDefaults

`func NewPatchedBulkWritableIPAddressRangeRequestWithDefaults() *PatchedBulkWritableIPAddressRangeRequest`

NewPatchedBulkWritableIPAddressRangeRequestWithDefaults instantiates a new PatchedBulkWritableIPAddressRangeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableIPAddressRangeRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableIPAddressRangeRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableIPAddressRangeRequest) SetId(v string)`

SetId sets Id field to given value.


### GetStartAddress

`func (o *PatchedBulkWritableIPAddressRangeRequest) GetStartAddress() string`

GetStartAddress returns the StartAddress field if non-nil, zero value otherwise.

### GetStartAddressOk

`func (o *PatchedBulkWritableIPAddressRangeRequest) GetStartAddressOk() (*string, bool)`

GetStartAddressOk returns a tuple with the StartAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAddress

`func (o *PatchedBulkWritableIPAddressRangeRequest) SetStartAddress(v string)`

SetStartAddress sets StartAddress field to given value.

### HasStartAddress

`func (o *PatchedBulkWritableIPAddressRangeRequest) HasStartAddress() bool`

HasStartAddress returns a boolean if a field has been set.

### GetEndAddress

`func (o *PatchedBulkWritableIPAddressRangeRequest) GetEndAddress() string`

GetEndAddress returns the EndAddress field if non-nil, zero value otherwise.

### GetEndAddressOk

`func (o *PatchedBulkWritableIPAddressRangeRequest) GetEndAddressOk() (*string, bool)`

GetEndAddressOk returns a tuple with the EndAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAddress

`func (o *PatchedBulkWritableIPAddressRangeRequest) SetEndAddress(v string)`

SetEndAddress sets EndAddress field to given value.

### HasEndAddress

`func (o *PatchedBulkWritableIPAddressRangeRequest) HasEndAddress() bool`

HasEndAddress returns a boolean if a field has been set.

### GetNamespace

`func (o *PatchedBulkWritableIPAddressRangeRequest) GetNamespace() BulkWritableIPAddressRangeRequestNamespace`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *PatchedBulkWritableIPAddressRangeRequest) GetNamespaceOk() (*BulkWritableIPAddressRangeRequestNamespace, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *PatchedBulkWritableIPAddressRangeRequest) SetNamespace(v BulkWritableIPAddressRangeRequestNamespace)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *PatchedBulkWritableIPAddressRangeRequest) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetName

`func (o *PatchedBulkWritableIPAddressRangeRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkWritableIPAddressRangeRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkWritableIPAddressRangeRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkWritableIPAddressRangeRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedBulkWritableIPAddressRangeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkWritableIPAddressRangeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkWritableIPAddressRangeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkWritableIPAddressRangeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCountAsUtilized

`func (o *PatchedBulkWritableIPAddressRangeRequest) GetCountAsUtilized() bool`

GetCountAsUtilized returns the CountAsUtilized field if non-nil, zero value otherwise.

### GetCountAsUtilizedOk

`func (o *PatchedBulkWritableIPAddressRangeRequest) GetCountAsUtilizedOk() (*bool, bool)`

GetCountAsUtilizedOk returns a tuple with the CountAsUtilized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountAsUtilized

`func (o *PatchedBulkWritableIPAddressRangeRequest) SetCountAsUtilized(v bool)`

SetCountAsUtilized sets CountAsUtilized field to given value.

### HasCountAsUtilized

`func (o *PatchedBulkWritableIPAddressRangeRequest) HasCountAsUtilized() bool`

HasCountAsUtilized returns a boolean if a field has been set.

### GetIsExclusive

`func (o *PatchedBulkWritableIPAddressRangeRequest) GetIsExclusive() bool`

GetIsExclusive returns the IsExclusive field if non-nil, zero value otherwise.

### GetIsExclusiveOk

`func (o *PatchedBulkWritableIPAddressRangeRequest) GetIsExclusiveOk() (*bool, bool)`

GetIsExclusiveOk returns a tuple with the IsExclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExclusive

`func (o *PatchedBulkWritableIPAddressRangeRequest) SetIsExclusive(v bool)`

SetIsExclusive sets IsExclusive field to given value.

### HasIsExclusive

`func (o *PatchedBulkWritableIPAddressRangeRequest) HasIsExclusive() bool`

HasIsExclusive returns a boolean if a field has been set.

### GetParent

`func (o *PatchedBulkWritableIPAddressRangeRequest) GetParent() BulkWritableIPAddressRangeRequestParent`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *PatchedBulkWritableIPAddressRangeRequest) GetParentOk() (*BulkWritableIPAddressRangeRequestParent, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *PatchedBulkWritableIPAddressRangeRequest) SetParent(v BulkWritableIPAddressRangeRequestParent)`

SetParent sets Parent field to given value.

### HasParent

`func (o *PatchedBulkWritableIPAddressRangeRequest) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetStatus

`func (o *PatchedBulkWritableIPAddressRangeRequest) GetStatus() BulkWritableCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedBulkWritableIPAddressRangeRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedBulkWritableIPAddressRangeRequest) SetStatus(v BulkWritableCableRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedBulkWritableIPAddressRangeRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRole

`func (o *PatchedBulkWritableIPAddressRangeRequest) GetRole() ApprovalWorkflowUser`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PatchedBulkWritableIPAddressRangeRequest) GetRoleOk() (*ApprovalWorkflowUser, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PatchedBulkWritableIPAddressRangeRequest) SetRole(v ApprovalWorkflowUser)`

SetRole sets Role field to given value.

### HasRole

`func (o *PatchedBulkWritableIPAddressRangeRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *PatchedBulkWritableIPAddressRangeRequest) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *PatchedBulkWritableIPAddressRangeRequest) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetTenant

`func (o *PatchedBulkWritableIPAddressRangeRequest) GetTenant() ApprovalWorkflowUser`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedBulkWritableIPAddressRangeRequest) GetTenantOk() (*ApprovalWorkflowUser, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedBulkWritableIPAddressRangeRequest) SetTenant(v ApprovalWorkflowUser)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedBulkWritableIPAddressRangeRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedBulkWritableIPAddressRangeRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedBulkWritableIPAddressRangeRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetTags

`func (o *PatchedBulkWritableIPAddressRangeRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkWritableIPAddressRangeRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkWritableIPAddressRangeRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkWritableIPAddressRangeRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedBulkWritableIPAddressRangeRequest) GetCustomFields() map[string]*interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkWritableIPAddressRangeRequest) GetCustomFieldsOk() (*map[string]*interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkWritableIPAddressRangeRequest) SetCustomFields(v map[string]*interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkWritableIPAddressRangeRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedBulkWritableIPAddressRangeRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedBulkWritableIPAddressRangeRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedBulkWritableIPAddressRangeRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedBulkWritableIPAddressRangeRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


