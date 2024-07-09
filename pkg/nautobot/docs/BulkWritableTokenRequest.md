# BulkWritableTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Key** | Pointer to **string** |  | [optional] 
**Expires** | Pointer to **NullableTime** |  | [optional] 
**WriteEnabled** | Pointer to **bool** | Permit create/update/delete operations using this key | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewBulkWritableTokenRequest

`func NewBulkWritableTokenRequest(id string, ) *BulkWritableTokenRequest`

NewBulkWritableTokenRequest instantiates a new BulkWritableTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableTokenRequestWithDefaults

`func NewBulkWritableTokenRequestWithDefaults() *BulkWritableTokenRequest`

NewBulkWritableTokenRequestWithDefaults instantiates a new BulkWritableTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableTokenRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableTokenRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableTokenRequest) SetId(v string)`

SetId sets Id field to given value.


### GetKey

`func (o *BulkWritableTokenRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *BulkWritableTokenRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *BulkWritableTokenRequest) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *BulkWritableTokenRequest) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetExpires

`func (o *BulkWritableTokenRequest) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *BulkWritableTokenRequest) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *BulkWritableTokenRequest) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *BulkWritableTokenRequest) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### SetExpiresNil

`func (o *BulkWritableTokenRequest) SetExpiresNil(b bool)`

 SetExpiresNil sets the value for Expires to be an explicit nil

### UnsetExpires
`func (o *BulkWritableTokenRequest) UnsetExpires()`

UnsetExpires ensures that no value is present for Expires, not even an explicit nil
### GetWriteEnabled

`func (o *BulkWritableTokenRequest) GetWriteEnabled() bool`

GetWriteEnabled returns the WriteEnabled field if non-nil, zero value otherwise.

### GetWriteEnabledOk

`func (o *BulkWritableTokenRequest) GetWriteEnabledOk() (*bool, bool)`

GetWriteEnabledOk returns a tuple with the WriteEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteEnabled

`func (o *BulkWritableTokenRequest) SetWriteEnabled(v bool)`

SetWriteEnabled sets WriteEnabled field to given value.

### HasWriteEnabled

`func (o *BulkWritableTokenRequest) HasWriteEnabled() bool`

HasWriteEnabled returns a boolean if a field has been set.

### GetDescription

`func (o *BulkWritableTokenRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkWritableTokenRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkWritableTokenRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkWritableTokenRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


