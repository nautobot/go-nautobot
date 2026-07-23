# VPNTermination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**AssignedObjectType** | **NullableString** |  | [readonly] 
**Vpn** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Vlan** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**Interface** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**VmInterface** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**NotesUrl** | **string** |  | [readonly] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewVPNTermination

`func NewVPNTermination(objectType string, display string, url string, naturalSlug string, assignedObjectType NullableString, vpn BulkWritableCableRequestStatus, created NullableTime, lastUpdated NullableTime, notesUrl string, ) *VPNTermination`

NewVPNTermination instantiates a new VPNTermination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVPNTerminationWithDefaults

`func NewVPNTerminationWithDefaults() *VPNTermination`

NewVPNTerminationWithDefaults instantiates a new VPNTermination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VPNTermination) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VPNTermination) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VPNTermination) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VPNTermination) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectType

`func (o *VPNTermination) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VPNTermination) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VPNTermination) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *VPNTermination) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *VPNTermination) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *VPNTermination) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *VPNTermination) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VPNTermination) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VPNTermination) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *VPNTermination) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *VPNTermination) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *VPNTermination) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetAssignedObjectType

`func (o *VPNTermination) GetAssignedObjectType() string`

GetAssignedObjectType returns the AssignedObjectType field if non-nil, zero value otherwise.

### GetAssignedObjectTypeOk

`func (o *VPNTermination) GetAssignedObjectTypeOk() (*string, bool)`

GetAssignedObjectTypeOk returns a tuple with the AssignedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedObjectType

`func (o *VPNTermination) SetAssignedObjectType(v string)`

SetAssignedObjectType sets AssignedObjectType field to given value.


### SetAssignedObjectTypeNil

`func (o *VPNTermination) SetAssignedObjectTypeNil(b bool)`

 SetAssignedObjectTypeNil sets the value for AssignedObjectType to be an explicit nil

### UnsetAssignedObjectType
`func (o *VPNTermination) UnsetAssignedObjectType()`

UnsetAssignedObjectType ensures that no value is present for AssignedObjectType, not even an explicit nil
### GetVpn

`func (o *VPNTermination) GetVpn() BulkWritableCableRequestStatus`

GetVpn returns the Vpn field if non-nil, zero value otherwise.

### GetVpnOk

`func (o *VPNTermination) GetVpnOk() (*BulkWritableCableRequestStatus, bool)`

GetVpnOk returns a tuple with the Vpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpn

`func (o *VPNTermination) SetVpn(v BulkWritableCableRequestStatus)`

SetVpn sets Vpn field to given value.


### GetVlan

`func (o *VPNTermination) GetVlan() ApprovalWorkflowUser`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *VPNTermination) GetVlanOk() (*ApprovalWorkflowUser, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *VPNTermination) SetVlan(v ApprovalWorkflowUser)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *VPNTermination) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### SetVlanNil

`func (o *VPNTermination) SetVlanNil(b bool)`

 SetVlanNil sets the value for Vlan to be an explicit nil

### UnsetVlan
`func (o *VPNTermination) UnsetVlan()`

UnsetVlan ensures that no value is present for Vlan, not even an explicit nil
### GetInterface

`func (o *VPNTermination) GetInterface() ApprovalWorkflowUser`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *VPNTermination) GetInterfaceOk() (*ApprovalWorkflowUser, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *VPNTermination) SetInterface(v ApprovalWorkflowUser)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *VPNTermination) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### SetInterfaceNil

`func (o *VPNTermination) SetInterfaceNil(b bool)`

 SetInterfaceNil sets the value for Interface to be an explicit nil

### UnsetInterface
`func (o *VPNTermination) UnsetInterface()`

UnsetInterface ensures that no value is present for Interface, not even an explicit nil
### GetVmInterface

`func (o *VPNTermination) GetVmInterface() ApprovalWorkflowUser`

GetVmInterface returns the VmInterface field if non-nil, zero value otherwise.

### GetVmInterfaceOk

`func (o *VPNTermination) GetVmInterfaceOk() (*ApprovalWorkflowUser, bool)`

GetVmInterfaceOk returns a tuple with the VmInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmInterface

`func (o *VPNTermination) SetVmInterface(v ApprovalWorkflowUser)`

SetVmInterface sets VmInterface field to given value.

### HasVmInterface

`func (o *VPNTermination) HasVmInterface() bool`

HasVmInterface returns a boolean if a field has been set.

### SetVmInterfaceNil

`func (o *VPNTermination) SetVmInterfaceNil(b bool)`

 SetVmInterfaceNil sets the value for VmInterface to be an explicit nil

### UnsetVmInterface
`func (o *VPNTermination) UnsetVmInterface()`

UnsetVmInterface ensures that no value is present for VmInterface, not even an explicit nil
### GetCreated

`func (o *VPNTermination) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *VPNTermination) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *VPNTermination) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *VPNTermination) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *VPNTermination) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *VPNTermination) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *VPNTermination) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *VPNTermination) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *VPNTermination) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *VPNTermination) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetNotesUrl

`func (o *VPNTermination) GetNotesUrl() string`

GetNotesUrl returns the NotesUrl field if non-nil, zero value otherwise.

### GetNotesUrlOk

`func (o *VPNTermination) GetNotesUrlOk() (*string, bool)`

GetNotesUrlOk returns a tuple with the NotesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesUrl

`func (o *VPNTermination) SetNotesUrl(v string)`

SetNotesUrl sets NotesUrl field to given value.


### GetCustomFields

`func (o *VPNTermination) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *VPNTermination) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *VPNTermination) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *VPNTermination) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetTags

`func (o *VPNTermination) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VPNTermination) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VPNTermination) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VPNTermination) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


