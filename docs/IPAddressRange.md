# IPAddressRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**StartAddress** | **string** |  | 
**EndAddress** | **string** |  | 
**Size** | **int32** |  | [readonly] 
**Name** | Pointer to **string** | Name of the IP Address Range | [optional] 
**StartHost** | **string** | First IP host address in the range (inclusive) | [readonly] 
**EndHost** | **string** | Last IP host address in the range (inclusive) | [readonly] 
**IpVersion** | **int32** |  | [readonly] 
**Description** | Pointer to **string** |  | [optional] 
**CountAsUtilized** | Pointer to **bool** | Forces this range to count as fully utilized in prefix utilization calculations. | [optional] 
**IsExclusive** | Pointer to **bool** | Prevent individual IP Address objects from being created within this range. | [optional] 
**Parent** | Pointer to [**BulkWritableIPAddressRangeRequestParent**](BulkWritableIPAddressRangeRequestParent.md) |  | [optional] 
**Status** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Role** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**Tenant** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**NotesUrl** | **string** |  | [readonly] 
**CustomFields** | Pointer to **map[string]interface{}** | Custom field data for this object, keyed by each applicable Custom Field&#39;s &#x60;key&#x60;. Value types vary with the custom field&#39;s type (text, integer, boolean, date, URL, JSON, select, multi-select); undefined values are &#x60;null&#x60;. On write, the payload is merged with existing values (PATCH-style: keys omitted from the payload are left untouched), and keys that do not correspond to a defined custom field are ignored. | [optional] 

## Methods

### NewIPAddressRange

`func NewIPAddressRange(objectType string, display string, url string, naturalSlug string, startAddress string, endAddress string, size int32, startHost string, endHost string, ipVersion int32, status BulkWritableCableRequestStatus, created NullableTime, lastUpdated NullableTime, notesUrl string, ) *IPAddressRange`

NewIPAddressRange instantiates a new IPAddressRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPAddressRangeWithDefaults

`func NewIPAddressRangeWithDefaults() *IPAddressRange`

NewIPAddressRangeWithDefaults instantiates a new IPAddressRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IPAddressRange) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IPAddressRange) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IPAddressRange) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IPAddressRange) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectType

`func (o *IPAddressRange) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IPAddressRange) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IPAddressRange) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *IPAddressRange) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *IPAddressRange) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *IPAddressRange) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *IPAddressRange) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IPAddressRange) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IPAddressRange) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *IPAddressRange) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *IPAddressRange) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *IPAddressRange) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetStartAddress

`func (o *IPAddressRange) GetStartAddress() string`

GetStartAddress returns the StartAddress field if non-nil, zero value otherwise.

### GetStartAddressOk

`func (o *IPAddressRange) GetStartAddressOk() (*string, bool)`

GetStartAddressOk returns a tuple with the StartAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAddress

`func (o *IPAddressRange) SetStartAddress(v string)`

SetStartAddress sets StartAddress field to given value.


### GetEndAddress

`func (o *IPAddressRange) GetEndAddress() string`

GetEndAddress returns the EndAddress field if non-nil, zero value otherwise.

### GetEndAddressOk

`func (o *IPAddressRange) GetEndAddressOk() (*string, bool)`

GetEndAddressOk returns a tuple with the EndAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAddress

`func (o *IPAddressRange) SetEndAddress(v string)`

SetEndAddress sets EndAddress field to given value.


### GetSize

`func (o *IPAddressRange) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *IPAddressRange) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *IPAddressRange) SetSize(v int32)`

SetSize sets Size field to given value.


### GetName

`func (o *IPAddressRange) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IPAddressRange) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IPAddressRange) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IPAddressRange) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStartHost

`func (o *IPAddressRange) GetStartHost() string`

GetStartHost returns the StartHost field if non-nil, zero value otherwise.

### GetStartHostOk

`func (o *IPAddressRange) GetStartHostOk() (*string, bool)`

GetStartHostOk returns a tuple with the StartHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartHost

`func (o *IPAddressRange) SetStartHost(v string)`

SetStartHost sets StartHost field to given value.


### GetEndHost

`func (o *IPAddressRange) GetEndHost() string`

GetEndHost returns the EndHost field if non-nil, zero value otherwise.

### GetEndHostOk

`func (o *IPAddressRange) GetEndHostOk() (*string, bool)`

GetEndHostOk returns a tuple with the EndHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndHost

`func (o *IPAddressRange) SetEndHost(v string)`

SetEndHost sets EndHost field to given value.


### GetIpVersion

`func (o *IPAddressRange) GetIpVersion() int32`

GetIpVersion returns the IpVersion field if non-nil, zero value otherwise.

### GetIpVersionOk

`func (o *IPAddressRange) GetIpVersionOk() (*int32, bool)`

GetIpVersionOk returns a tuple with the IpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpVersion

`func (o *IPAddressRange) SetIpVersion(v int32)`

SetIpVersion sets IpVersion field to given value.


### GetDescription

`func (o *IPAddressRange) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IPAddressRange) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IPAddressRange) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IPAddressRange) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCountAsUtilized

`func (o *IPAddressRange) GetCountAsUtilized() bool`

GetCountAsUtilized returns the CountAsUtilized field if non-nil, zero value otherwise.

### GetCountAsUtilizedOk

`func (o *IPAddressRange) GetCountAsUtilizedOk() (*bool, bool)`

GetCountAsUtilizedOk returns a tuple with the CountAsUtilized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountAsUtilized

`func (o *IPAddressRange) SetCountAsUtilized(v bool)`

SetCountAsUtilized sets CountAsUtilized field to given value.

### HasCountAsUtilized

`func (o *IPAddressRange) HasCountAsUtilized() bool`

HasCountAsUtilized returns a boolean if a field has been set.

### GetIsExclusive

`func (o *IPAddressRange) GetIsExclusive() bool`

GetIsExclusive returns the IsExclusive field if non-nil, zero value otherwise.

### GetIsExclusiveOk

`func (o *IPAddressRange) GetIsExclusiveOk() (*bool, bool)`

GetIsExclusiveOk returns a tuple with the IsExclusive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExclusive

`func (o *IPAddressRange) SetIsExclusive(v bool)`

SetIsExclusive sets IsExclusive field to given value.

### HasIsExclusive

`func (o *IPAddressRange) HasIsExclusive() bool`

HasIsExclusive returns a boolean if a field has been set.

### GetParent

`func (o *IPAddressRange) GetParent() BulkWritableIPAddressRangeRequestParent`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *IPAddressRange) GetParentOk() (*BulkWritableIPAddressRangeRequestParent, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *IPAddressRange) SetParent(v BulkWritableIPAddressRangeRequestParent)`

SetParent sets Parent field to given value.

### HasParent

`func (o *IPAddressRange) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetStatus

`func (o *IPAddressRange) GetStatus() BulkWritableCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IPAddressRange) GetStatusOk() (*BulkWritableCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IPAddressRange) SetStatus(v BulkWritableCableRequestStatus)`

SetStatus sets Status field to given value.


### GetRole

`func (o *IPAddressRange) GetRole() ApprovalWorkflowUser`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *IPAddressRange) GetRoleOk() (*ApprovalWorkflowUser, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *IPAddressRange) SetRole(v ApprovalWorkflowUser)`

SetRole sets Role field to given value.

### HasRole

`func (o *IPAddressRange) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *IPAddressRange) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *IPAddressRange) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetTenant

`func (o *IPAddressRange) GetTenant() ApprovalWorkflowUser`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *IPAddressRange) GetTenantOk() (*ApprovalWorkflowUser, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *IPAddressRange) SetTenant(v ApprovalWorkflowUser)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *IPAddressRange) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *IPAddressRange) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *IPAddressRange) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetCreated

`func (o *IPAddressRange) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *IPAddressRange) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *IPAddressRange) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *IPAddressRange) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *IPAddressRange) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *IPAddressRange) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *IPAddressRange) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *IPAddressRange) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *IPAddressRange) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *IPAddressRange) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetTags

`func (o *IPAddressRange) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IPAddressRange) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IPAddressRange) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *IPAddressRange) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetNotesUrl

`func (o *IPAddressRange) GetNotesUrl() string`

GetNotesUrl returns the NotesUrl field if non-nil, zero value otherwise.

### GetNotesUrlOk

`func (o *IPAddressRange) GetNotesUrlOk() (*string, bool)`

GetNotesUrlOk returns a tuple with the NotesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesUrl

`func (o *IPAddressRange) SetNotesUrl(v string)`

SetNotesUrl sets NotesUrl field to given value.


### GetCustomFields

`func (o *IPAddressRange) GetCustomFields() map[string]*interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *IPAddressRange) GetCustomFieldsOk() (*map[string]*interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *IPAddressRange) SetCustomFields(v map[string]*interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *IPAddressRange) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


