# ProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
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

### NewProviderRequest

`func NewProviderRequest(name string, ) *ProviderRequest`

NewProviderRequest instantiates a new ProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderRequestWithDefaults

`func NewProviderRequestWithDefaults() *ProviderRequest`

NewProviderRequestWithDefaults instantiates a new ProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProviderRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProviderRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProviderRequest) SetName(v string)`

SetName sets Name field to given value.


### GetAsn

`func (o *ProviderRequest) GetAsn() int64`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *ProviderRequest) GetAsnOk() (*int64, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *ProviderRequest) SetAsn(v int64)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *ProviderRequest) HasAsn() bool`

HasAsn returns a boolean if a field has been set.

### SetAsnNil

`func (o *ProviderRequest) SetAsnNil(b bool)`

 SetAsnNil sets the value for Asn to be an explicit nil

### UnsetAsn
`func (o *ProviderRequest) UnsetAsn()`

UnsetAsn ensures that no value is present for Asn, not even an explicit nil
### GetAccount

`func (o *ProviderRequest) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ProviderRequest) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ProviderRequest) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ProviderRequest) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetPortalUrl

`func (o *ProviderRequest) GetPortalUrl() string`

GetPortalUrl returns the PortalUrl field if non-nil, zero value otherwise.

### GetPortalUrlOk

`func (o *ProviderRequest) GetPortalUrlOk() (*string, bool)`

GetPortalUrlOk returns a tuple with the PortalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalUrl

`func (o *ProviderRequest) SetPortalUrl(v string)`

SetPortalUrl sets PortalUrl field to given value.

### HasPortalUrl

`func (o *ProviderRequest) HasPortalUrl() bool`

HasPortalUrl returns a boolean if a field has been set.

### GetNocContact

`func (o *ProviderRequest) GetNocContact() string`

GetNocContact returns the NocContact field if non-nil, zero value otherwise.

### GetNocContactOk

`func (o *ProviderRequest) GetNocContactOk() (*string, bool)`

GetNocContactOk returns a tuple with the NocContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNocContact

`func (o *ProviderRequest) SetNocContact(v string)`

SetNocContact sets NocContact field to given value.

### HasNocContact

`func (o *ProviderRequest) HasNocContact() bool`

HasNocContact returns a boolean if a field has been set.

### GetAdminContact

`func (o *ProviderRequest) GetAdminContact() string`

GetAdminContact returns the AdminContact field if non-nil, zero value otherwise.

### GetAdminContactOk

`func (o *ProviderRequest) GetAdminContactOk() (*string, bool)`

GetAdminContactOk returns a tuple with the AdminContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminContact

`func (o *ProviderRequest) SetAdminContact(v string)`

SetAdminContact sets AdminContact field to given value.

### HasAdminContact

`func (o *ProviderRequest) HasAdminContact() bool`

HasAdminContact returns a boolean if a field has been set.

### GetComments

`func (o *ProviderRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *ProviderRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *ProviderRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *ProviderRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetTags

`func (o *ProviderRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ProviderRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ProviderRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ProviderRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCustomFields

`func (o *ProviderRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ProviderRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ProviderRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ProviderRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *ProviderRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ProviderRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ProviderRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ProviderRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


