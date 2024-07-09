# IPAddressRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** |  | 
**Namespace** | Pointer to [**BulkWritableIPAddressRequestNamespace**](BulkWritableIPAddressRequestNamespace.md) |  | [optional] 
**Type** | Pointer to [**IPAddressTypeChoices**](IPAddressTypeChoices.md) |  | [optional] 
**DnsName** | Pointer to **string** | Hostname or FQDN (not case-sensitive) | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Status** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Role** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**Parent** | Pointer to [**NullableBulkWritableIPAddressRequestParent**](BulkWritableIPAddressRequestParent.md) |  | [optional] 
**Tenant** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**NatInside** | Pointer to [**NullableNATInside**](NATInside.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewIPAddressRequest

`func NewIPAddressRequest(address string, status BulkWritableCableRequestStatus, ) *IPAddressRequest`

NewIPAddressRequest instantiates a new IPAddressRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPAddressRequestWithDefaults

`func NewIPAddressRequestWithDefaults() *IPAddressRequest`

NewIPAddressRequestWithDefaults instantiates a new IPAddressRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *IPAddressRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *IPAddressRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *IPAddressRequest) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetNamespace

`func (o *IPAddressRequest) GetNamespace() BulkWritableIPAddressRequestNamespace`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *IPAddressRequest) GetNamespaceOk() (*BulkWritableIPAddressRequestNamespace, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *IPAddressRequest) SetNamespace(v BulkWritableIPAddressRequestNamespace)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *IPAddressRequest) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetType

`func (o *IPAddressRequest) GetType() IPAddressTypeChoices`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IPAddressRequest) GetTypeOk() (*IPAddressTypeChoices, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IPAddressRequest) SetType(v IPAddressTypeChoices)`

SetType sets Type field to given value.

### HasType

`func (o *IPAddressRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDnsName

`func (o *IPAddressRequest) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *IPAddressRequest) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *IPAddressRequest) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *IPAddressRequest) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetDescription

`func (o *IPAddressRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IPAddressRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IPAddressRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IPAddressRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *IPAddressRequest) GetStatus() BulkWritableCableRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IPAddressRequest) GetStatusOk() (*BulkWritableCableRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IPAddressRequest) SetStatus(v BulkWritableCableRequestStatus)`

SetStatus sets Status field to given value.


### GetRole

`func (o *IPAddressRequest) GetRole() BulkWritableCircuitRequestTenant`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *IPAddressRequest) GetRoleOk() (*BulkWritableCircuitRequestTenant, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *IPAddressRequest) SetRole(v BulkWritableCircuitRequestTenant)`

SetRole sets Role field to given value.

### HasRole

`func (o *IPAddressRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *IPAddressRequest) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *IPAddressRequest) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetParent

`func (o *IPAddressRequest) GetParent() BulkWritableIPAddressRequestParent`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *IPAddressRequest) GetParentOk() (*BulkWritableIPAddressRequestParent, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *IPAddressRequest) SetParent(v BulkWritableIPAddressRequestParent)`

SetParent sets Parent field to given value.

### HasParent

`func (o *IPAddressRequest) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *IPAddressRequest) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *IPAddressRequest) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetTenant

`func (o *IPAddressRequest) GetTenant() BulkWritableCircuitRequestTenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *IPAddressRequest) GetTenantOk() (*BulkWritableCircuitRequestTenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *IPAddressRequest) SetTenant(v BulkWritableCircuitRequestTenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *IPAddressRequest) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### SetTenantNil

`func (o *IPAddressRequest) SetTenantNil(b bool)`

 SetTenantNil sets the value for Tenant to be an explicit nil

### UnsetTenant
`func (o *IPAddressRequest) UnsetTenant()`

UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
### GetNatInside

`func (o *IPAddressRequest) GetNatInside() NATInside`

GetNatInside returns the NatInside field if non-nil, zero value otherwise.

### GetNatInsideOk

`func (o *IPAddressRequest) GetNatInsideOk() (*NATInside, bool)`

GetNatInsideOk returns a tuple with the NatInside field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatInside

`func (o *IPAddressRequest) SetNatInside(v NATInside)`

SetNatInside sets NatInside field to given value.

### HasNatInside

`func (o *IPAddressRequest) HasNatInside() bool`

HasNatInside returns a boolean if a field has been set.

### SetNatInsideNil

`func (o *IPAddressRequest) SetNatInsideNil(b bool)`

 SetNatInsideNil sets the value for NatInside to be an explicit nil

### UnsetNatInside
`func (o *IPAddressRequest) UnsetNatInside()`

UnsetNatInside ensures that no value is present for NatInside, not even an explicit nil
### GetTags

`func (o *IPAddressRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IPAddressRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IPAddressRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *IPAddressRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *IPAddressRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *IPAddressRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *IPAddressRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *IPAddressRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *IPAddressRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *IPAddressRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *IPAddressRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *IPAddressRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


