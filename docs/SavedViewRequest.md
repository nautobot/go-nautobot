# SavedViewRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of this view | 
**View** | **string** | The name of the list view that the saved view is derived from, e.g. dcim:device_list | 
**Config** | Pointer to **interface{}** | Saved Configuration on this view | [optional] 
**IsGlobalDefault** | Pointer to **bool** |  | [optional] 
**IsShared** | Pointer to **bool** |  | [optional] 
**Owner** | [**BulkWritableSavedViewRequestOwner**](BulkWritableSavedViewRequestOwner.md) |  | 

## Methods

### NewSavedViewRequest

`func NewSavedViewRequest(name string, view string, owner BulkWritableSavedViewRequestOwner, ) *SavedViewRequest`

NewSavedViewRequest instantiates a new SavedViewRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSavedViewRequestWithDefaults

`func NewSavedViewRequestWithDefaults() *SavedViewRequest`

NewSavedViewRequestWithDefaults instantiates a new SavedViewRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SavedViewRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SavedViewRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SavedViewRequest) SetName(v string)`

SetName sets Name field to given value.


### GetView

`func (o *SavedViewRequest) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *SavedViewRequest) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *SavedViewRequest) SetView(v string)`

SetView sets View field to given value.


### GetConfig

`func (o *SavedViewRequest) GetConfig() interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *SavedViewRequest) GetConfigOk() (*interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *SavedViewRequest) SetConfig(v interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *SavedViewRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *SavedViewRequest) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *SavedViewRequest) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetIsGlobalDefault

`func (o *SavedViewRequest) GetIsGlobalDefault() bool`

GetIsGlobalDefault returns the IsGlobalDefault field if non-nil, zero value otherwise.

### GetIsGlobalDefaultOk

`func (o *SavedViewRequest) GetIsGlobalDefaultOk() (*bool, bool)`

GetIsGlobalDefaultOk returns a tuple with the IsGlobalDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGlobalDefault

`func (o *SavedViewRequest) SetIsGlobalDefault(v bool)`

SetIsGlobalDefault sets IsGlobalDefault field to given value.

### HasIsGlobalDefault

`func (o *SavedViewRequest) HasIsGlobalDefault() bool`

HasIsGlobalDefault returns a boolean if a field has been set.

### GetIsShared

`func (o *SavedViewRequest) GetIsShared() bool`

GetIsShared returns the IsShared field if non-nil, zero value otherwise.

### GetIsSharedOk

`func (o *SavedViewRequest) GetIsSharedOk() (*bool, bool)`

GetIsSharedOk returns a tuple with the IsShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShared

`func (o *SavedViewRequest) SetIsShared(v bool)`

SetIsShared sets IsShared field to given value.

### HasIsShared

`func (o *SavedViewRequest) HasIsShared() bool`

HasIsShared returns a boolean if a field has been set.

### GetOwner

`func (o *SavedViewRequest) GetOwner() BulkWritableSavedViewRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *SavedViewRequest) GetOwnerOk() (*BulkWritableSavedViewRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *SavedViewRequest) SetOwner(v BulkWritableSavedViewRequestOwner)`

SetOwner sets Owner field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


