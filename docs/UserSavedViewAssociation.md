# UserSavedViewAssociation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**ViewName** | **string** |  | 
**SavedView** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**User** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 

## Methods

### NewUserSavedViewAssociation

`func NewUserSavedViewAssociation(id string, objectType string, display string, url string, naturalSlug string, viewName string, savedView BulkWritableCableRequestStatus, user BulkWritableCableRequestStatus, ) *UserSavedViewAssociation`

NewUserSavedViewAssociation instantiates a new UserSavedViewAssociation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSavedViewAssociationWithDefaults

`func NewUserSavedViewAssociationWithDefaults() *UserSavedViewAssociation`

NewUserSavedViewAssociationWithDefaults instantiates a new UserSavedViewAssociation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserSavedViewAssociation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserSavedViewAssociation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserSavedViewAssociation) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *UserSavedViewAssociation) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *UserSavedViewAssociation) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *UserSavedViewAssociation) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *UserSavedViewAssociation) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *UserSavedViewAssociation) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *UserSavedViewAssociation) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *UserSavedViewAssociation) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UserSavedViewAssociation) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UserSavedViewAssociation) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *UserSavedViewAssociation) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *UserSavedViewAssociation) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *UserSavedViewAssociation) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetViewName

`func (o *UserSavedViewAssociation) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *UserSavedViewAssociation) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *UserSavedViewAssociation) SetViewName(v string)`

SetViewName sets ViewName field to given value.


### GetSavedView

`func (o *UserSavedViewAssociation) GetSavedView() BulkWritableCableRequestStatus`

GetSavedView returns the SavedView field if non-nil, zero value otherwise.

### GetSavedViewOk

`func (o *UserSavedViewAssociation) GetSavedViewOk() (*BulkWritableCableRequestStatus, bool)`

GetSavedViewOk returns a tuple with the SavedView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavedView

`func (o *UserSavedViewAssociation) SetSavedView(v BulkWritableCableRequestStatus)`

SetSavedView sets SavedView field to given value.


### GetUser

`func (o *UserSavedViewAssociation) GetUser() BulkWritableCableRequestStatus`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UserSavedViewAssociation) GetUserOk() (*BulkWritableCableRequestStatus, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UserSavedViewAssociation) SetUser(v BulkWritableCableRequestStatus)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


