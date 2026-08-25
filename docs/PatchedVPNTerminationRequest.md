# PatchedVPNTerminationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Vpn** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**Vlan** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**Interface** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**VmInterface** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewPatchedVPNTerminationRequest

`func NewPatchedVPNTerminationRequest() *PatchedVPNTerminationRequest`

NewPatchedVPNTerminationRequest instantiates a new PatchedVPNTerminationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedVPNTerminationRequestWithDefaults

`func NewPatchedVPNTerminationRequestWithDefaults() *PatchedVPNTerminationRequest`

NewPatchedVPNTerminationRequestWithDefaults instantiates a new PatchedVPNTerminationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedVPNTerminationRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedVPNTerminationRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedVPNTerminationRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedVPNTerminationRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVpn

`func (o *PatchedVPNTerminationRequest) GetVpn() BulkWritableCableRequestStatus`

GetVpn returns the Vpn field if non-nil, zero value otherwise.

### GetVpnOk

`func (o *PatchedVPNTerminationRequest) GetVpnOk() (*BulkWritableCableRequestStatus, bool)`

GetVpnOk returns a tuple with the Vpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpn

`func (o *PatchedVPNTerminationRequest) SetVpn(v BulkWritableCableRequestStatus)`

SetVpn sets Vpn field to given value.

### HasVpn

`func (o *PatchedVPNTerminationRequest) HasVpn() bool`

HasVpn returns a boolean if a field has been set.

### GetVlan

`func (o *PatchedVPNTerminationRequest) GetVlan() ApprovalWorkflowUser`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *PatchedVPNTerminationRequest) GetVlanOk() (*ApprovalWorkflowUser, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *PatchedVPNTerminationRequest) SetVlan(v ApprovalWorkflowUser)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *PatchedVPNTerminationRequest) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### SetVlanNil

`func (o *PatchedVPNTerminationRequest) SetVlanNil(b bool)`

 SetVlanNil sets the value for Vlan to be an explicit nil

### UnsetVlan
`func (o *PatchedVPNTerminationRequest) UnsetVlan()`

UnsetVlan ensures that no value is present for Vlan, not even an explicit nil
### GetInterface

`func (o *PatchedVPNTerminationRequest) GetInterface() ApprovalWorkflowUser`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *PatchedVPNTerminationRequest) GetInterfaceOk() (*ApprovalWorkflowUser, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *PatchedVPNTerminationRequest) SetInterface(v ApprovalWorkflowUser)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *PatchedVPNTerminationRequest) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### SetInterfaceNil

`func (o *PatchedVPNTerminationRequest) SetInterfaceNil(b bool)`

 SetInterfaceNil sets the value for Interface to be an explicit nil

### UnsetInterface
`func (o *PatchedVPNTerminationRequest) UnsetInterface()`

UnsetInterface ensures that no value is present for Interface, not even an explicit nil
### GetVmInterface

`func (o *PatchedVPNTerminationRequest) GetVmInterface() ApprovalWorkflowUser`

GetVmInterface returns the VmInterface field if non-nil, zero value otherwise.

### GetVmInterfaceOk

`func (o *PatchedVPNTerminationRequest) GetVmInterfaceOk() (*ApprovalWorkflowUser, bool)`

GetVmInterfaceOk returns a tuple with the VmInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmInterface

`func (o *PatchedVPNTerminationRequest) SetVmInterface(v ApprovalWorkflowUser)`

SetVmInterface sets VmInterface field to given value.

### HasVmInterface

`func (o *PatchedVPNTerminationRequest) HasVmInterface() bool`

HasVmInterface returns a boolean if a field has been set.

### SetVmInterfaceNil

`func (o *PatchedVPNTerminationRequest) SetVmInterfaceNil(b bool)`

 SetVmInterfaceNil sets the value for VmInterface to be an explicit nil

### UnsetVmInterface
`func (o *PatchedVPNTerminationRequest) UnsetVmInterface()`

UnsetVmInterface ensures that no value is present for VmInterface, not even an explicit nil
### GetCustomFields

`func (o *PatchedVPNTerminationRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedVPNTerminationRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedVPNTerminationRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedVPNTerminationRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedVPNTerminationRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedVPNTerminationRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedVPNTerminationRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedVPNTerminationRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *PatchedVPNTerminationRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedVPNTerminationRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedVPNTerminationRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedVPNTerminationRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


