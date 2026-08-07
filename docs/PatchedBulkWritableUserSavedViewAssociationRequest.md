# PatchedBulkWritableUserSavedViewAssociationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ViewName** | Pointer to **string** |  | [optional] 
**SavedView** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**User** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewPatchedBulkWritableUserSavedViewAssociationRequest

`func NewPatchedBulkWritableUserSavedViewAssociationRequest(id string, ) *PatchedBulkWritableUserSavedViewAssociationRequest`

NewPatchedBulkWritableUserSavedViewAssociationRequest instantiates a new PatchedBulkWritableUserSavedViewAssociationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableUserSavedViewAssociationRequestWithDefaults

`func NewPatchedBulkWritableUserSavedViewAssociationRequestWithDefaults() *PatchedBulkWritableUserSavedViewAssociationRequest`

NewPatchedBulkWritableUserSavedViewAssociationRequestWithDefaults instantiates a new PatchedBulkWritableUserSavedViewAssociationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableUserSavedViewAssociationRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableUserSavedViewAssociationRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableUserSavedViewAssociationRequest) SetId(v string)`

SetId sets Id field to given value.


### GetViewName

`func (o *PatchedBulkWritableUserSavedViewAssociationRequest) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *PatchedBulkWritableUserSavedViewAssociationRequest) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *PatchedBulkWritableUserSavedViewAssociationRequest) SetViewName(v string)`

SetViewName sets ViewName field to given value.

### HasViewName

`func (o *PatchedBulkWritableUserSavedViewAssociationRequest) HasViewName() bool`

HasViewName returns a boolean if a field has been set.

### GetSavedView

`func (o *PatchedBulkWritableUserSavedViewAssociationRequest) GetSavedView() BulkWritableCableRequestStatus`

GetSavedView returns the SavedView field if non-nil, zero value otherwise.

### GetSavedViewOk

`func (o *PatchedBulkWritableUserSavedViewAssociationRequest) GetSavedViewOk() (*BulkWritableCableRequestStatus, bool)`

GetSavedViewOk returns a tuple with the SavedView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavedView

`func (o *PatchedBulkWritableUserSavedViewAssociationRequest) SetSavedView(v BulkWritableCableRequestStatus)`

SetSavedView sets SavedView field to given value.

### HasSavedView

`func (o *PatchedBulkWritableUserSavedViewAssociationRequest) HasSavedView() bool`

HasSavedView returns a boolean if a field has been set.

### GetUser

`func (o *PatchedBulkWritableUserSavedViewAssociationRequest) GetUser() BulkWritableCableRequestStatus`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PatchedBulkWritableUserSavedViewAssociationRequest) GetUserOk() (*BulkWritableCableRequestStatus, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PatchedBulkWritableUserSavedViewAssociationRequest) SetUser(v BulkWritableCableRequestStatus)`

SetUser sets User field to given value.

### HasUser

`func (o *PatchedBulkWritableUserSavedViewAssociationRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


