# PatchedBulkWritableSavedViewRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | Pointer to **string** | The name of this view | [optional] 
**View** | Pointer to **string** | The name of the list view that the saved view is derived from, e.g. dcim:device_list | [optional] 
**Config** | Pointer to **interface{}** | Saved Configuration on this view | [optional] 
**IsGlobalDefault** | Pointer to **bool** |  | [optional] 
**IsShared** | Pointer to **bool** |  | [optional] 
**Owner** | Pointer to [**BulkWritableSavedViewRequestOwner**](BulkWritableSavedViewRequestOwner.md) |  | [optional] 

## Methods

### NewPatchedBulkWritableSavedViewRequest

`func NewPatchedBulkWritableSavedViewRequest(id string, ) *PatchedBulkWritableSavedViewRequest`

NewPatchedBulkWritableSavedViewRequest instantiates a new PatchedBulkWritableSavedViewRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableSavedViewRequestWithDefaults

`func NewPatchedBulkWritableSavedViewRequestWithDefaults() *PatchedBulkWritableSavedViewRequest`

NewPatchedBulkWritableSavedViewRequestWithDefaults instantiates a new PatchedBulkWritableSavedViewRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableSavedViewRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableSavedViewRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableSavedViewRequest) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PatchedBulkWritableSavedViewRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkWritableSavedViewRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkWritableSavedViewRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkWritableSavedViewRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetView

`func (o *PatchedBulkWritableSavedViewRequest) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *PatchedBulkWritableSavedViewRequest) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *PatchedBulkWritableSavedViewRequest) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *PatchedBulkWritableSavedViewRequest) HasView() bool`

HasView returns a boolean if a field has been set.

### GetConfig

`func (o *PatchedBulkWritableSavedViewRequest) GetConfig() interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *PatchedBulkWritableSavedViewRequest) GetConfigOk() (*interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *PatchedBulkWritableSavedViewRequest) SetConfig(v interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *PatchedBulkWritableSavedViewRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *PatchedBulkWritableSavedViewRequest) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *PatchedBulkWritableSavedViewRequest) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetIsGlobalDefault

`func (o *PatchedBulkWritableSavedViewRequest) GetIsGlobalDefault() bool`

GetIsGlobalDefault returns the IsGlobalDefault field if non-nil, zero value otherwise.

### GetIsGlobalDefaultOk

`func (o *PatchedBulkWritableSavedViewRequest) GetIsGlobalDefaultOk() (*bool, bool)`

GetIsGlobalDefaultOk returns a tuple with the IsGlobalDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGlobalDefault

`func (o *PatchedBulkWritableSavedViewRequest) SetIsGlobalDefault(v bool)`

SetIsGlobalDefault sets IsGlobalDefault field to given value.

### HasIsGlobalDefault

`func (o *PatchedBulkWritableSavedViewRequest) HasIsGlobalDefault() bool`

HasIsGlobalDefault returns a boolean if a field has been set.

### GetIsShared

`func (o *PatchedBulkWritableSavedViewRequest) GetIsShared() bool`

GetIsShared returns the IsShared field if non-nil, zero value otherwise.

### GetIsSharedOk

`func (o *PatchedBulkWritableSavedViewRequest) GetIsSharedOk() (*bool, bool)`

GetIsSharedOk returns a tuple with the IsShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShared

`func (o *PatchedBulkWritableSavedViewRequest) SetIsShared(v bool)`

SetIsShared sets IsShared field to given value.

### HasIsShared

`func (o *PatchedBulkWritableSavedViewRequest) HasIsShared() bool`

HasIsShared returns a boolean if a field has been set.

### GetOwner

`func (o *PatchedBulkWritableSavedViewRequest) GetOwner() BulkWritableSavedViewRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PatchedBulkWritableSavedViewRequest) GetOwnerOk() (*BulkWritableSavedViewRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PatchedBulkWritableSavedViewRequest) SetOwner(v BulkWritableSavedViewRequestOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PatchedBulkWritableSavedViewRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


