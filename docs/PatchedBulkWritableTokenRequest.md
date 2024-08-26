# PatchedBulkWritableTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Key** | Pointer to **string** |  | [optional] 
**Expires** | Pointer to **NullableTime** |  | [optional] 
**WriteEnabled** | Pointer to **bool** | Permit create/update/delete operations using this key | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewPatchedBulkWritableTokenRequest

`func NewPatchedBulkWritableTokenRequest(id string, ) *PatchedBulkWritableTokenRequest`

NewPatchedBulkWritableTokenRequest instantiates a new PatchedBulkWritableTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableTokenRequestWithDefaults

`func NewPatchedBulkWritableTokenRequestWithDefaults() *PatchedBulkWritableTokenRequest`

NewPatchedBulkWritableTokenRequestWithDefaults instantiates a new PatchedBulkWritableTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableTokenRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableTokenRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableTokenRequest) SetId(v string)`

SetId sets Id field to given value.


### GetKey

`func (o *PatchedBulkWritableTokenRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PatchedBulkWritableTokenRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PatchedBulkWritableTokenRequest) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *PatchedBulkWritableTokenRequest) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetExpires

`func (o *PatchedBulkWritableTokenRequest) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *PatchedBulkWritableTokenRequest) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *PatchedBulkWritableTokenRequest) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *PatchedBulkWritableTokenRequest) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### SetExpiresNil

`func (o *PatchedBulkWritableTokenRequest) SetExpiresNil(b bool)`

 SetExpiresNil sets the value for Expires to be an explicit nil

### UnsetExpires
`func (o *PatchedBulkWritableTokenRequest) UnsetExpires()`

UnsetExpires ensures that no value is present for Expires, not even an explicit nil
### GetWriteEnabled

`func (o *PatchedBulkWritableTokenRequest) GetWriteEnabled() bool`

GetWriteEnabled returns the WriteEnabled field if non-nil, zero value otherwise.

### GetWriteEnabledOk

`func (o *PatchedBulkWritableTokenRequest) GetWriteEnabledOk() (*bool, bool)`

GetWriteEnabledOk returns a tuple with the WriteEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteEnabled

`func (o *PatchedBulkWritableTokenRequest) SetWriteEnabled(v bool)`

SetWriteEnabled sets WriteEnabled field to given value.

### HasWriteEnabled

`func (o *PatchedBulkWritableTokenRequest) HasWriteEnabled() bool`

HasWriteEnabled returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedBulkWritableTokenRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkWritableTokenRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkWritableTokenRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkWritableTokenRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


