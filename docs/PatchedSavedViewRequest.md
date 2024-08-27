# PatchedSavedViewRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of this view | [optional] 
**View** | Pointer to **string** | The name of the list view that the saved view is derived from, e.g. dcim:device_list | [optional] 
**Config** | Pointer to **interface{}** | Saved Configuration on this view | [optional] 
**IsGlobalDefault** | Pointer to **bool** |  | [optional] 
**IsShared** | Pointer to **bool** |  | [optional] 
**Owner** | Pointer to [**BulkWritableSavedViewRequestOwner**](BulkWritableSavedViewRequestOwner.md) |  | [optional] 

## Methods

### NewPatchedSavedViewRequest

`func NewPatchedSavedViewRequest() *PatchedSavedViewRequest`

NewPatchedSavedViewRequest instantiates a new PatchedSavedViewRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedSavedViewRequestWithDefaults

`func NewPatchedSavedViewRequestWithDefaults() *PatchedSavedViewRequest`

NewPatchedSavedViewRequestWithDefaults instantiates a new PatchedSavedViewRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedSavedViewRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedSavedViewRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedSavedViewRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedSavedViewRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetView

`func (o *PatchedSavedViewRequest) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *PatchedSavedViewRequest) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *PatchedSavedViewRequest) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *PatchedSavedViewRequest) HasView() bool`

HasView returns a boolean if a field has been set.

### GetConfig

`func (o *PatchedSavedViewRequest) GetConfig() interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *PatchedSavedViewRequest) GetConfigOk() (*interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *PatchedSavedViewRequest) SetConfig(v interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *PatchedSavedViewRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *PatchedSavedViewRequest) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *PatchedSavedViewRequest) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetIsGlobalDefault

`func (o *PatchedSavedViewRequest) GetIsGlobalDefault() bool`

GetIsGlobalDefault returns the IsGlobalDefault field if non-nil, zero value otherwise.

### GetIsGlobalDefaultOk

`func (o *PatchedSavedViewRequest) GetIsGlobalDefaultOk() (*bool, bool)`

GetIsGlobalDefaultOk returns a tuple with the IsGlobalDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGlobalDefault

`func (o *PatchedSavedViewRequest) SetIsGlobalDefault(v bool)`

SetIsGlobalDefault sets IsGlobalDefault field to given value.

### HasIsGlobalDefault

`func (o *PatchedSavedViewRequest) HasIsGlobalDefault() bool`

HasIsGlobalDefault returns a boolean if a field has been set.

### GetIsShared

`func (o *PatchedSavedViewRequest) GetIsShared() bool`

GetIsShared returns the IsShared field if non-nil, zero value otherwise.

### GetIsSharedOk

`func (o *PatchedSavedViewRequest) GetIsSharedOk() (*bool, bool)`

GetIsSharedOk returns a tuple with the IsShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShared

`func (o *PatchedSavedViewRequest) SetIsShared(v bool)`

SetIsShared sets IsShared field to given value.

### HasIsShared

`func (o *PatchedSavedViewRequest) HasIsShared() bool`

HasIsShared returns a boolean if a field has been set.

### GetOwner

`func (o *PatchedSavedViewRequest) GetOwner() BulkWritableSavedViewRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PatchedSavedViewRequest) GetOwnerOk() (*BulkWritableSavedViewRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PatchedSavedViewRequest) SetOwner(v BulkWritableSavedViewRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PatchedSavedViewRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


