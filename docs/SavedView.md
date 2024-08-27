# SavedView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**Name** | **string** | The name of this view | 
**View** | **string** | The name of the list view that the saved view is derived from, e.g. dcim:device_list | 
**Config** | Pointer to **interface{}** | Saved Configuration on this view | [optional] 
**IsGlobalDefault** | Pointer to **bool** |  | [optional] 
**IsShared** | Pointer to **bool** |  | [optional] 
**Owner** | [**BulkWritableSavedViewRequestOwner**](BulkWritableSavedViewRequestOwner.md) |  | 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 

## Methods

### NewSavedView

`func NewSavedView(id string, objectType string, display string, url string, naturalSlug string, name string, view string, owner BulkWritableSavedViewRequestOwner, created NullableTime, lastUpdated NullableTime, ) *SavedView`

NewSavedView instantiates a new SavedView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSavedViewWithDefaults

`func NewSavedViewWithDefaults() *SavedView`

NewSavedViewWithDefaults instantiates a new SavedView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SavedView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SavedView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SavedView) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *SavedView) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SavedView) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SavedView) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *SavedView) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *SavedView) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *SavedView) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *SavedView) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SavedView) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SavedView) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *SavedView) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *SavedView) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *SavedView) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetName

`func (o *SavedView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SavedView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SavedView) SetName(v string)`

SetName sets Name field to given value.


### GetView

`func (o *SavedView) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *SavedView) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *SavedView) SetView(v string)`

SetView sets View field to given value.


### GetConfig

`func (o *SavedView) GetConfig() interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *SavedView) GetConfigOk() (*interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *SavedView) SetConfig(v interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *SavedView) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *SavedView) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *SavedView) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetIsGlobalDefault

`func (o *SavedView) GetIsGlobalDefault() bool`

GetIsGlobalDefault returns the IsGlobalDefault field if non-nil, zero value otherwise.

### GetIsGlobalDefaultOk

`func (o *SavedView) GetIsGlobalDefaultOk() (*bool, bool)`

GetIsGlobalDefaultOk returns a tuple with the IsGlobalDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGlobalDefault

`func (o *SavedView) SetIsGlobalDefault(v bool)`

SetIsGlobalDefault sets IsGlobalDefault field to given value.

### HasIsGlobalDefault

`func (o *SavedView) HasIsGlobalDefault() bool`

HasIsGlobalDefault returns a boolean if a field has been set.

### GetIsShared

`func (o *SavedView) GetIsShared() bool`

GetIsShared returns the IsShared field if non-nil, zero value otherwise.

### GetIsSharedOk

`func (o *SavedView) GetIsSharedOk() (*bool, bool)`

GetIsSharedOk returns a tuple with the IsShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShared

`func (o *SavedView) SetIsShared(v bool)`

SetIsShared sets IsShared field to given value.

### HasIsShared

`func (o *SavedView) HasIsShared() bool`

HasIsShared returns a boolean if a field has been set.

### GetOwner

`func (o *SavedView) GetOwner() BulkWritableSavedViewRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *SavedView) GetOwnerOk() (*BulkWritableSavedViewRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *SavedView) SetOwner(v BulkWritableSavedViewRequestOwner)`

SetOwner sets Owner field to given value.


### GetCreated

`func (o *SavedView) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SavedView) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SavedView) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *SavedView) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *SavedView) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *SavedView) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *SavedView) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *SavedView) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *SavedView) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *SavedView) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


