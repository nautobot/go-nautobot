# PatchedProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Asn** | Pointer to **NullableInt64** | 32-bit autonomous system number | [optional] 
**Account** | Pointer to **string** |  | [optional] 
**PortalUrl** | Pointer to **string** |  | [optional] 
**NocContact** | Pointer to **string** |  | [optional] 
**AdminContact** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewPatchedProviderRequest

`func NewPatchedProviderRequest() *PatchedProviderRequest`

NewPatchedProviderRequest instantiates a new PatchedProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedProviderRequestWithDefaults

`func NewPatchedProviderRequestWithDefaults() *PatchedProviderRequest`

NewPatchedProviderRequestWithDefaults instantiates a new PatchedProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedProviderRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedProviderRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedProviderRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedProviderRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAsn

`func (o *PatchedProviderRequest) GetAsn() int64`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *PatchedProviderRequest) GetAsnOk() (*int64, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *PatchedProviderRequest) SetAsn(v int64)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *PatchedProviderRequest) HasAsn() bool`

HasAsn returns a boolean if a field has been set.

### SetAsnNil

`func (o *PatchedProviderRequest) SetAsnNil(b bool)`

 SetAsnNil sets the value for Asn to be an explicit nil

### UnsetAsn
`func (o *PatchedProviderRequest) UnsetAsn()`

UnsetAsn ensures that no value is present for Asn, not even an explicit nil
### GetAccount

`func (o *PatchedProviderRequest) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *PatchedProviderRequest) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *PatchedProviderRequest) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *PatchedProviderRequest) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetPortalUrl

`func (o *PatchedProviderRequest) GetPortalUrl() string`

GetPortalUrl returns the PortalUrl field if non-nil, zero value otherwise.

### GetPortalUrlOk

`func (o *PatchedProviderRequest) GetPortalUrlOk() (*string, bool)`

GetPortalUrlOk returns a tuple with the PortalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalUrl

`func (o *PatchedProviderRequest) SetPortalUrl(v string)`

SetPortalUrl sets PortalUrl field to given value.

### HasPortalUrl

`func (o *PatchedProviderRequest) HasPortalUrl() bool`

HasPortalUrl returns a boolean if a field has been set.

### GetNocContact

`func (o *PatchedProviderRequest) GetNocContact() string`

GetNocContact returns the NocContact field if non-nil, zero value otherwise.

### GetNocContactOk

`func (o *PatchedProviderRequest) GetNocContactOk() (*string, bool)`

GetNocContactOk returns a tuple with the NocContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNocContact

`func (o *PatchedProviderRequest) SetNocContact(v string)`

SetNocContact sets NocContact field to given value.

### HasNocContact

`func (o *PatchedProviderRequest) HasNocContact() bool`

HasNocContact returns a boolean if a field has been set.

### GetAdminContact

`func (o *PatchedProviderRequest) GetAdminContact() string`

GetAdminContact returns the AdminContact field if non-nil, zero value otherwise.

### GetAdminContactOk

`func (o *PatchedProviderRequest) GetAdminContactOk() (*string, bool)`

GetAdminContactOk returns a tuple with the AdminContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminContact

`func (o *PatchedProviderRequest) SetAdminContact(v string)`

SetAdminContact sets AdminContact field to given value.

### HasAdminContact

`func (o *PatchedProviderRequest) HasAdminContact() bool`

HasAdminContact returns a boolean if a field has been set.

### GetComments

`func (o *PatchedProviderRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedProviderRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedProviderRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedProviderRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *PatchedProviderRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedProviderRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedProviderRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedProviderRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedProviderRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedProviderRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedProviderRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedProviderRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedProviderRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedProviderRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedProviderRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedProviderRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


