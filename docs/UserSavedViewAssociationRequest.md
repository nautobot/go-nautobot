# UserSavedViewAssociationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ViewName** | **string** |  | 
**SavedView** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**User** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 

## Methods

### NewUserSavedViewAssociationRequest

`func NewUserSavedViewAssociationRequest(viewName string, savedView BulkWritableCableRequestStatus, user BulkWritableCableRequestStatus, ) *UserSavedViewAssociationRequest`

NewUserSavedViewAssociationRequest instantiates a new UserSavedViewAssociationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSavedViewAssociationRequestWithDefaults

`func NewUserSavedViewAssociationRequestWithDefaults() *UserSavedViewAssociationRequest`

NewUserSavedViewAssociationRequestWithDefaults instantiates a new UserSavedViewAssociationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetViewName

`func (o *UserSavedViewAssociationRequest) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *UserSavedViewAssociationRequest) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *UserSavedViewAssociationRequest) SetViewName(v string)`

SetViewName sets ViewName field to given value.


### GetSavedView

`func (o *UserSavedViewAssociationRequest) GetSavedView() BulkWritableCableRequestStatus`

GetSavedView returns the SavedView field if non-nil, zero value otherwise.

### GetSavedViewOk

`func (o *UserSavedViewAssociationRequest) GetSavedViewOk() (*BulkWritableCableRequestStatus, bool)`

GetSavedViewOk returns a tuple with the SavedView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavedView

`func (o *UserSavedViewAssociationRequest) SetSavedView(v BulkWritableCableRequestStatus)`

SetSavedView sets SavedView field to given value.


### GetUser

`func (o *UserSavedViewAssociationRequest) GetUser() BulkWritableCableRequestStatus`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UserSavedViewAssociationRequest) GetUserOk() (*BulkWritableCableRequestStatus, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UserSavedViewAssociationRequest) SetUser(v BulkWritableCableRequestStatus)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


