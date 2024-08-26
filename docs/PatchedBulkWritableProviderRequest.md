# PatchedBulkWritableProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
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

### NewPatchedBulkWritableProviderRequest

`func NewPatchedBulkWritableProviderRequest(id string, ) *PatchedBulkWritableProviderRequest`

NewPatchedBulkWritableProviderRequest instantiates a new PatchedBulkWritableProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableProviderRequestWithDefaults

`func NewPatchedBulkWritableProviderRequestWithDefaults() *PatchedBulkWritableProviderRequest`

NewPatchedBulkWritableProviderRequestWithDefaults instantiates a new PatchedBulkWritableProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableProviderRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableProviderRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableProviderRequest) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PatchedBulkWritableProviderRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkWritableProviderRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkWritableProviderRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkWritableProviderRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAsn

`func (o *PatchedBulkWritableProviderRequest) GetAsn() int64`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *PatchedBulkWritableProviderRequest) GetAsnOk() (*int64, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *PatchedBulkWritableProviderRequest) SetAsn(v int64)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *PatchedBulkWritableProviderRequest) HasAsn() bool`

HasAsn returns a boolean if a field has been set.

### SetAsnNil

`func (o *PatchedBulkWritableProviderRequest) SetAsnNil(b bool)`

 SetAsnNil sets the value for Asn to be an explicit nil

### UnsetAsn
`func (o *PatchedBulkWritableProviderRequest) UnsetAsn()`

UnsetAsn ensures that no value is present for Asn, not even an explicit nil
### GetAccount

`func (o *PatchedBulkWritableProviderRequest) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *PatchedBulkWritableProviderRequest) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *PatchedBulkWritableProviderRequest) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *PatchedBulkWritableProviderRequest) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetPortalUrl

`func (o *PatchedBulkWritableProviderRequest) GetPortalUrl() string`

GetPortalUrl returns the PortalUrl field if non-nil, zero value otherwise.

### GetPortalUrlOk

`func (o *PatchedBulkWritableProviderRequest) GetPortalUrlOk() (*string, bool)`

GetPortalUrlOk returns a tuple with the PortalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalUrl

`func (o *PatchedBulkWritableProviderRequest) SetPortalUrl(v string)`

SetPortalUrl sets PortalUrl field to given value.

### HasPortalUrl

`func (o *PatchedBulkWritableProviderRequest) HasPortalUrl() bool`

HasPortalUrl returns a boolean if a field has been set.

### GetNocContact

`func (o *PatchedBulkWritableProviderRequest) GetNocContact() string`

GetNocContact returns the NocContact field if non-nil, zero value otherwise.

### GetNocContactOk

`func (o *PatchedBulkWritableProviderRequest) GetNocContactOk() (*string, bool)`

GetNocContactOk returns a tuple with the NocContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNocContact

`func (o *PatchedBulkWritableProviderRequest) SetNocContact(v string)`

SetNocContact sets NocContact field to given value.

### HasNocContact

`func (o *PatchedBulkWritableProviderRequest) HasNocContact() bool`

HasNocContact returns a boolean if a field has been set.

### GetAdminContact

`func (o *PatchedBulkWritableProviderRequest) GetAdminContact() string`

GetAdminContact returns the AdminContact field if non-nil, zero value otherwise.

### GetAdminContactOk

`func (o *PatchedBulkWritableProviderRequest) GetAdminContactOk() (*string, bool)`

GetAdminContactOk returns a tuple with the AdminContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminContact

`func (o *PatchedBulkWritableProviderRequest) SetAdminContact(v string)`

SetAdminContact sets AdminContact field to given value.

### HasAdminContact

`func (o *PatchedBulkWritableProviderRequest) HasAdminContact() bool`

HasAdminContact returns a boolean if a field has been set.

### GetComments

`func (o *PatchedBulkWritableProviderRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PatchedBulkWritableProviderRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PatchedBulkWritableProviderRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PatchedBulkWritableProviderRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *PatchedBulkWritableProviderRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedBulkWritableProviderRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedBulkWritableProviderRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedBulkWritableProviderRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedBulkWritableProviderRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkWritableProviderRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkWritableProviderRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkWritableProviderRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedBulkWritableProviderRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedBulkWritableProviderRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedBulkWritableProviderRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedBulkWritableProviderRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


