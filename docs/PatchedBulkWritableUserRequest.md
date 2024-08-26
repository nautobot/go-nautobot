# PatchedBulkWritableUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Password** | Pointer to **NullableString** |  | [optional] 
**LastLogin** | Pointer to **NullableTime** |  | [optional] 
**IsSuperuser** | Pointer to **bool** | Designates that this user has all permissions without explicitly assigning them. | [optional] 
**Username** | Pointer to **string** | Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only. | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**IsStaff** | Pointer to **bool** | Designates whether the user can log into this admin site. | [optional] 
**IsActive** | Pointer to **bool** | Designates whether this user should be treated as active. Unselect this instead of deleting accounts. | [optional] 
**DateJoined** | Pointer to **time.Time** |  | [optional] 
**ConfigData** | Pointer to **map[string]interface{}** |  | [optional] 
**Groups** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) | The groups this user belongs to. A user will get all permissions granted to each of their groups. | [optional] 

## Methods

### NewPatchedBulkWritableUserRequest

`func NewPatchedBulkWritableUserRequest(id string, ) *PatchedBulkWritableUserRequest`

NewPatchedBulkWritableUserRequest instantiates a new PatchedBulkWritableUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableUserRequestWithDefaults

`func NewPatchedBulkWritableUserRequestWithDefaults() *PatchedBulkWritableUserRequest`

NewPatchedBulkWritableUserRequestWithDefaults instantiates a new PatchedBulkWritableUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableUserRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableUserRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableUserRequest) SetId(v string)`

SetId sets Id field to given value.


### GetPassword

`func (o *PatchedBulkWritableUserRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PatchedBulkWritableUserRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PatchedBulkWritableUserRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PatchedBulkWritableUserRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *PatchedBulkWritableUserRequest) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *PatchedBulkWritableUserRequest) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetLastLogin

`func (o *PatchedBulkWritableUserRequest) GetLastLogin() time.Time`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *PatchedBulkWritableUserRequest) GetLastLoginOk() (*time.Time, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *PatchedBulkWritableUserRequest) SetLastLogin(v time.Time)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *PatchedBulkWritableUserRequest) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### SetLastLoginNil

`func (o *PatchedBulkWritableUserRequest) SetLastLoginNil(b bool)`

 SetLastLoginNil sets the value for LastLogin to be an explicit nil

### UnsetLastLogin
`func (o *PatchedBulkWritableUserRequest) UnsetLastLogin()`

UnsetLastLogin ensures that no value is present for LastLogin, not even an explicit nil
### GetIsSuperuser

`func (o *PatchedBulkWritableUserRequest) GetIsSuperuser() bool`

GetIsSuperuser returns the IsSuperuser field if non-nil, zero value otherwise.

### GetIsSuperuserOk

`func (o *PatchedBulkWritableUserRequest) GetIsSuperuserOk() (*bool, bool)`

GetIsSuperuserOk returns a tuple with the IsSuperuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuperuser

`func (o *PatchedBulkWritableUserRequest) SetIsSuperuser(v bool)`

SetIsSuperuser sets IsSuperuser field to given value.

### HasIsSuperuser

`func (o *PatchedBulkWritableUserRequest) HasIsSuperuser() bool`

HasIsSuperuser returns a boolean if a field has been set.

### GetUsername

`func (o *PatchedBulkWritableUserRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *PatchedBulkWritableUserRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *PatchedBulkWritableUserRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *PatchedBulkWritableUserRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetFirstName

`func (o *PatchedBulkWritableUserRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *PatchedBulkWritableUserRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *PatchedBulkWritableUserRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *PatchedBulkWritableUserRequest) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *PatchedBulkWritableUserRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *PatchedBulkWritableUserRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *PatchedBulkWritableUserRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *PatchedBulkWritableUserRequest) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEmail

`func (o *PatchedBulkWritableUserRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PatchedBulkWritableUserRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PatchedBulkWritableUserRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PatchedBulkWritableUserRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetIsStaff

`func (o *PatchedBulkWritableUserRequest) GetIsStaff() bool`

GetIsStaff returns the IsStaff field if non-nil, zero value otherwise.

### GetIsStaffOk

`func (o *PatchedBulkWritableUserRequest) GetIsStaffOk() (*bool, bool)`

GetIsStaffOk returns a tuple with the IsStaff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStaff

`func (o *PatchedBulkWritableUserRequest) SetIsStaff(v bool)`

SetIsStaff sets IsStaff field to given value.

### HasIsStaff

`func (o *PatchedBulkWritableUserRequest) HasIsStaff() bool`

HasIsStaff returns a boolean if a field has been set.

### GetIsActive

`func (o *PatchedBulkWritableUserRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *PatchedBulkWritableUserRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *PatchedBulkWritableUserRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *PatchedBulkWritableUserRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetDateJoined

`func (o *PatchedBulkWritableUserRequest) GetDateJoined() time.Time`

GetDateJoined returns the DateJoined field if non-nil, zero value otherwise.

### GetDateJoinedOk

`func (o *PatchedBulkWritableUserRequest) GetDateJoinedOk() (*time.Time, bool)`

GetDateJoinedOk returns a tuple with the DateJoined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateJoined

`func (o *PatchedBulkWritableUserRequest) SetDateJoined(v time.Time)`

SetDateJoined sets DateJoined field to given value.

### HasDateJoined

`func (o *PatchedBulkWritableUserRequest) HasDateJoined() bool`

HasDateJoined returns a boolean if a field has been set.

### GetConfigData

`func (o *PatchedBulkWritableUserRequest) GetConfigData() map[string]interface{}`

GetConfigData returns the ConfigData field if non-nil, zero value otherwise.

### GetConfigDataOk

`func (o *PatchedBulkWritableUserRequest) GetConfigDataOk() (*map[string]interface{}, bool)`

GetConfigDataOk returns a tuple with the ConfigData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigData

`func (o *PatchedBulkWritableUserRequest) SetConfigData(v map[string]interface{})`

SetConfigData sets ConfigData field to given value.

### HasConfigData

`func (o *PatchedBulkWritableUserRequest) HasConfigData() bool`

HasConfigData returns a boolean if a field has been set.

### GetGroups

`func (o *PatchedBulkWritableUserRequest) GetGroups() []BulkWritableCableRequestStatus`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *PatchedBulkWritableUserRequest) GetGroupsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *PatchedBulkWritableUserRequest) SetGroups(v []BulkWritableCableRequestStatus)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *PatchedBulkWritableUserRequest) HasGroups() bool`

HasGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


