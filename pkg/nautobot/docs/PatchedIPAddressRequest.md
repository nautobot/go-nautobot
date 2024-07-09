# PatchedIPAddressRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to [**BulkWritableIPAddressRequestNamespace**](BulkWritableIPAddressRequestNamespace.md) |  | [optional] 
**Type** | Pointer to [**IPAddressTypeChoices**](IPAddressTypeChoices.md) |  | [optional] 
**DnsName** | Pointer to **string** | Hostname or FQDN (not case-sensitive) | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**Role** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Parent** | Pointer to [**NullableBulkWritableIPAddressRequestParent**](BulkWritableIPAddressRequestParent.md) |  | [optional] 
**Tenant** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**NatInside** | Pointer to [**NullableNATInside**](NATInside.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewPatchedIPAddressRequest

`func NewPatchedIPAddressRequest() *PatchedIPAddressRequest`

NewPatchedIPAddressRequest instantiates a new PatchedIPAddressRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedIPAddressRequestWithDefaults

`func NewPatchedIPAddressRequestWithDefaults() *PatchedIPAddressRequest`

NewPatchedIPAddressRequestWithDefaults instantiates a new PatchedIPAddressRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *PatchedIPAddressRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PatchedIPAddressRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PatchedIPAddressRequest) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *PatchedIPAddressRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetNamespace

`func (o *PatchedIPAddressRequest) GetNamespace() BulkWritableIPAddressRequestNamespace`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *PatchedIPAddressRequest) GetNamespaceOk() (*BulkWritableIPAddressRequestNamespace, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *PatchedIPAddressRequest) SetNamespace(v BulkWritableIPAddressRequestNamespace)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *PatchedIPAddressRequest) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetType

`func (o *PatchedIPAddressRequest) GetType() IPAddressTypeChoices`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedIPAddressRequest) GetTypeOk() (*IPAddressTypeChoices, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedIPAddressRequest) SetType(v IPAddressTypeChoices)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedIPAddressRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDnsName

`func (o *PatchedIPAddressRequest) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *PatchedIPAddressRequest) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *PatchedIPAddressRequest) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *PatchedIPAddressRequest) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedIPAddressRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedIPAddressRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedIPAddressRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedIPAddressRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *PatchedIPAddressRequest) GetStatus() BulkWritableCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PatchedIPAddressRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PatchedIPAddressRequest) SetStatus(v BulkWritableCableRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PatchedIPAddressRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRole

`func (o *PatchedIPAddressRequest) GetRole() BulkWritableCircuitRequestTenant`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PatchedIPAddressRequest) GetRoleOk() (*BulkWritableCircuitRequestTenant, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PatchedIPAddressRequest) SetRole(v BulkWritableCircuitRequestTenant)`

SetRole sets Role field to given value.

### HasRole

`func (o *PatchedIPAddressRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *PatchedIPAddressRequest) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *PatchedIPAddressRequest) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetParent

`func (o *PatchedIPAddressRequest) GetParent() BulkWritableIPAddressRequestParent`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *PatchedIPAddressRequest) GetParentOk() (*BulkWritableIPAddressRequestParent, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *PatchedIPAddressRequest) SetParent(v BulkWritableIPAddressRequestParent)`

SetParent sets Parent field to given value.

### HasParent

`func (o *PatchedIPAddressRequest) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *PatchedIPAddressRequest) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *PatchedIPAddressRequest) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetTenant

`func (o *PatchedIPAddressRequest) GetTenant() BulkWritableCircuitRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *PatchedIPAddressRequest) GetTenantOk() (*BulkWritableCircuitRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *PatchedIPAddressRequest) SetTenant(v BulkWritableCircuitRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *PatchedIPAddressRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *PatchedIPAddressRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *PatchedIPAddressRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetNatInside

`func (o *PatchedIPAddressRequest) GetNatInside() NATInside`

GetNatInside returns the NatInside field if non-nil, zero value otherwise.

### GetNatInsideOk

`func (o *PatchedIPAddressRequest) GetNatInsideOk() (*NATInside, bool)`

GetNatInsideOk returns a tuple with the NatInside field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatInside

`func (o *PatchedIPAddressRequest) SetNatInside(v NATInside)`

SetNatInside sets NatInside field to given value.

### HasNatInside

`func (o *PatchedIPAddressRequest) HasNatInside() bool`

HasNatInside returns a boolean if a field has been set.

### SetNatInsideNil

`func (o *PatchedIPAddressRequest) SetNatInsideNil(b bool)`

 SetNatInsideNil sets the value for NatInside to be an explicit nil

### UnsetNatInside
`func (o *PatchedIPAddressRequest) UnsetNatInside()`

UnsetNatInside ensures that no value is present for NatInside, not even an explicit nil
### GetTags

`func (o *PatchedIPAddressRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedIPAddressRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedIPAddressRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedIPAddressRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedIPAddressRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedIPAddressRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedIPAddressRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedIPAddressRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedIPAddressRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedIPAddressRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedIPAddressRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedIPAddressRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


