# PatchedUserSavedViewAssociationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ViewName** | Pointer to **string** |  | [optional] 
**SavedView** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**User** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewPatchedUserSavedViewAssociationRequest

`func NewPatchedUserSavedViewAssociationRequest() *PatchedUserSavedViewAssociationRequest`

NewPatchedUserSavedViewAssociationRequest instantiates a new PatchedUserSavedViewAssociationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedUserSavedViewAssociationRequestWithDefaults

`func NewPatchedUserSavedViewAssociationRequestWithDefaults() *PatchedUserSavedViewAssociationRequest`

NewPatchedUserSavedViewAssociationRequestWithDefaults instantiates a new PatchedUserSavedViewAssociationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetViewName

`func (o *PatchedUserSavedViewAssociationRequest) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *PatchedUserSavedViewAssociationRequest) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *PatchedUserSavedViewAssociationRequest) SetViewName(v string)`

SetViewName sets ViewName field to given value.

### HasViewName

`func (o *PatchedUserSavedViewAssociationRequest) HasViewName() bool`

HasViewName returns a boolean if a field has been set.

### GetSavedView

`func (o *PatchedUserSavedViewAssociationRequest) GetSavedView() BulkWritableCableRequestStatus`

GetSavedView returns the SavedView field if non-nil, zero value otherwise.

### GetSavedViewOk

`func (o *PatchedUserSavedViewAssociationRequest) GetSavedViewOk() (*BulkWritableCableRequestStatus, bool)`

GetSavedViewOk returns a tuple with the SavedView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavedView

`func (o *PatchedUserSavedViewAssociationRequest) SetSavedView(v BulkWritableCableRequestStatus)`

SetSavedView sets SavedView field to given value.

### HasSavedView

`func (o *PatchedUserSavedViewAssociationRequest) HasSavedView() bool`

HasSavedView returns a boolean if a field has been set.

### GetUser

`func (o *PatchedUserSavedViewAssociationRequest) GetUser() BulkWritableCableRequestStatus`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PatchedUserSavedViewAssociationRequest) GetUserOk() (*BulkWritableCableRequestStatus, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PatchedUserSavedViewAssociationRequest) SetUser(v BulkWritableCableRequestStatus)`

SetUser sets User field to given value.

### HasUser

`func (o *PatchedUserSavedViewAssociationRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


