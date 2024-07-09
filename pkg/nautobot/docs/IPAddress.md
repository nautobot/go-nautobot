# IPAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**Address** | **string** |  | 
**Host** | **string** | IPv4 or IPv6 host address | [readonly] 
**MaskLength** | **int32** | Length of the network mask, in bits. | [readonly] 
**Type** | Pointer to [**IPAddressTypeChoices**](IPAddressTypeChoices.md) |  | [optional] 
**IpVersion** | **int32** |  | [readonly] 
**DnsName** | Pointer to **string** | Hostname or FQDN (not case-sensitive) | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Status** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Role** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Parent** | Pointer to [**NullableBulkWritableIPAddressRequestParent**](BulkWritableIPAddressRequestParent.md) |  | [optional] 
**Tenant** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**NatInside** | Pointer to [**NullableNATInside**](NATInside.md) |  | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**NotesUrl** | **string** |  | [readonly] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**NatOutsideList** | [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [readonly] 
**Interfaces** | [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [readonly] 
**VmInterfaces** | [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [readonly] 

## Methods

### NewIPAddress

`func NewIPAddress(id string, objectType string, display string, url string, naturalSlug string, address string, host string, maskLength int32, ipVersion int32, status BulkWritableCableRequestStatus, created NullableTime, lastUpdated NullableTime, notesUrl string, natOutsideList []BulkWritableCableRequestStatus, interfaces []BulkWritableCableRequestStatus, vmInterfaces []BulkWritableCableRequestStatus, ) *IPAddress`

NewIPAddress instantiates a new IPAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPAddressWithDefaults

`func NewIPAddressWithDefaults() *IPAddress`

NewIPAddressWithDefaults instantiates a new IPAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IPAddress) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IPAddress) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IPAddress) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *IPAddress) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IPAddress) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IPAddress) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *IPAddress) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *IPAddress) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *IPAddress) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *IPAddress) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IPAddress) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IPAddress) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *IPAddress) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *IPAddress) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *IPAddress) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetAddress

`func (o *IPAddress) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *IPAddress) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *IPAddress) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetHost

`func (o *IPAddress) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *IPAddress) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *IPAddress) SetHost(v string)`

SetHost sets Host field to given value.


### GetMaskLength

`func (o *IPAddress) GetMaskLength() int32`

GetMaskLength returns the MaskLength field if non-nil, zero value otherwise.

### GetMaskLengthOk

`func (o *IPAddress) GetMaskLengthOk() (*int32, bool)`

GetMaskLengthOk returns a tuple with the MaskLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskLength

`func (o *IPAddress) SetMaskLength(v int32)`

SetMaskLength sets MaskLength field to given value.


### GetType

`func (o *IPAddress) GetType() IPAddressTypeChoices`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IPAddress) GetTypeOk() (*IPAddressTypeChoices, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IPAddress) SetType(v IPAddressTypeChoices)`

SetType sets Type field to given value.

### HasType

`func (o *IPAddress) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIpVersion

`func (o *IPAddress) GetIpVersion() int32`

GetIpVersion returns the IpVersion field if non-nil, zero value otherwise.

### GetIpVersionOk

`func (o *IPAddress) GetIpVersionOk() (*int32, bool)`

GetIpVersionOk returns a tuple with the IpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpVersion

`func (o *IPAddress) SetIpVersion(v int32)`

SetIpVersion sets IpVersion field to given value.


### GetDnsName

`func (o *IPAddress) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *IPAddress) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *IPAddress) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *IPAddress) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetDescription

`func (o *IPAddress) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IPAddress) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IPAddress) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IPAddress) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *IPAddress) GetStatus() BulkWritableCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IPAddress) GetStatusOk() (*BulkWritableCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IPAddress) SetStatus(v BulkWritableCableRequestStatus)`

SetStatus sets Status field to given value.


### GetRole

`func (o *IPAddress) GetRole() BulkWritableCircuitRequestTenant`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *IPAddress) GetRoleOk() (*BulkWritableCircuitRequestTenant, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *IPAddress) SetRole(v BulkWritableCircuitRequestTenant)`

SetRole sets Role field to given value.

### HasRole

`func (o *IPAddress) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *IPAddress) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *IPAddress) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetParent

`func (o *IPAddress) GetParent() BulkWritableIPAddressRequestParent`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *IPAddress) GetParentOk() (*BulkWritableIPAddressRequestParent, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *IPAddress) SetParent(v BulkWritableIPAddressRequestParent)`

SetParent sets Parent field to given value.

### HasParent

`func (o *IPAddress) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *IPAddress) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *IPAddress) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetTenant

`func (o *IPAddress) GetTenant() BulkWritableCircuitRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *IPAddress) GetTenantOk() (*BulkWritableCircuitRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *IPAddress) SetTenant(v BulkWritableCircuitRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *IPAddress) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *IPAddress) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *IPAddress) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetNatInside

`func (o *IPAddress) GetNatInside() NATInside`

GetNatInside returns the NatInside field if non-nil, zero value otherwise.

### GetNatInsideOk

`func (o *IPAddress) GetNatInsideOk() (*NATInside, bool)`

GetNatInsideOk returns a tuple with the NatInside field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatInside

`func (o *IPAddress) SetNatInside(v NATInside)`

SetNatInside sets NatInside field to given value.

### HasNatInside

`func (o *IPAddress) HasNatInside() bool`

HasNatInside returns a boolean if a field has been set.

### SetNatInsideNil

`func (o *IPAddress) SetNatInsideNil(b bool)`

 SetNatInsideNil sets the value for NatInside to be an explicit nil

### UnsetNatInside
`func (o *IPAddress) UnsetNatInside()`

UnsetNatInside ensures that no value is present for NatInside, not even an explicit nil
### GetCreated

`func (o *IPAddress) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *IPAddress) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *IPAddress) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *IPAddress) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *IPAddress) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *IPAddress) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *IPAddress) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *IPAddress) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *IPAddress) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *IPAddress) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetTags

`func (o *IPAddress) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IPAddress) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IPAddress) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *IPAddress) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetNotesUrl

`func (o *IPAddress) GetNotesUrl() string`

GetNotesUrl returns the NotesUrl field if non-nil, zero value otherwise.

### GetNotesUrlOk

`func (o *IPAddress) GetNotesUrlOk() (*string, bool)`

GetNotesUrlOk returns a tuple with the NotesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesUrl

`func (o *IPAddress) SetNotesUrl(v string)`

SetNotesUrl sets NotesUrl field to given value.


### GetCustomFields

`func (o *IPAddress) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *IPAddress) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *IPAddress) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *IPAddress) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetNatOutsideList

`func (o *IPAddress) GetNatOutsideList() []BulkWritableCableRequestStatus`

GetNatOutsideList returns the NatOutsideList field if non-nil, zero value otherwise.

### GetNatOutsideListOk

`func (o *IPAddress) GetNatOutsideListOk() (*[]BulkWritableCableRequestStatus, bool)`

GetNatOutsideListOk returns a tuple with the NatOutsideList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatOutsideList

`func (o *IPAddress) SetNatOutsideList(v []BulkWritableCableRequestStatus)`

SetNatOutsideList sets NatOutsideList field to given value.


### GetInterfaces

`func (o *IPAddress) GetInterfaces() []BulkWritableCableRequestStatus`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *IPAddress) GetInterfacesOk() (*[]BulkWritableCableRequestStatus, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *IPAddress) SetInterfaces(v []BulkWritableCableRequestStatus)`

SetInterfaces sets Interfaces field to given value.


### GetVmInterfaces

`func (o *IPAddress) GetVmInterfaces() []BulkWritableCableRequestStatus`

GetVmInterfaces returns the VmInterfaces field if non-nil, zero value otherwise.

### GetVmInterfacesOk

`func (o *IPAddress) GetVmInterfacesOk() (*[]BulkWritableCableRequestStatus, bool)`

GetVmInterfacesOk returns a tuple with the VmInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmInterfaces

`func (o *IPAddress) SetVmInterfaces(v []BulkWritableCableRequestStatus)`

SetVmInterfaces sets VmInterfaces field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


