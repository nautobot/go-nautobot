# BulkWritableSavedViewRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** | The name of this view | 
**View** | **string** | The name of the list view that the saved view is derived from, e.g. dcim:device_list | 
**Config** | Pointer to **interface{}** | Saved Configuration on this view | [optional] 
**IsGlobalDefault** | Pointer to **bool** |  | [optional] 
**IsShared** | Pointer to **bool** |  | [optional] 
**Owner** | [**BulkWritableSavedViewRequestOwner**](BulkWritableSavedViewRequestOwner.md) |  | 

## Methods

### NewBulkWritableSavedViewRequest

`func NewBulkWritableSavedViewRequest(id string, name string, view string, owner BulkWritableSavedViewRequestOwner, ) *BulkWritableSavedViewRequest`

NewBulkWritableSavedViewRequest instantiates a new BulkWritableSavedViewRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableSavedViewRequestWithDefaults

`func NewBulkWritableSavedViewRequestWithDefaults() *BulkWritableSavedViewRequest`

NewBulkWritableSavedViewRequestWithDefaults instantiates a new BulkWritableSavedViewRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableSavedViewRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableSavedViewRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableSavedViewRequest) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BulkWritableSavedViewRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkWritableSavedViewRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkWritableSavedViewRequest) SetName(v string)`

SetName sets Name field to given value.


### GetView

`func (o *BulkWritableSavedViewRequest) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *BulkWritableSavedViewRequest) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *BulkWritableSavedViewRequest) SetView(v string)`

SetView sets View field to given value.


### GetConfig

`func (o *BulkWritableSavedViewRequest) GetConfig() interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *BulkWritableSavedViewRequest) GetConfigOk() (*interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *BulkWritableSavedViewRequest) SetConfig(v interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *BulkWritableSavedViewRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *BulkWritableSavedViewRequest) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *BulkWritableSavedViewRequest) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetIsGlobalDefault

`func (o *BulkWritableSavedViewRequest) GetIsGlobalDefault() bool`

GetIsGlobalDefault returns the IsGlobalDefault field if non-nil, zero value otherwise.

### GetIsGlobalDefaultOk

`func (o *BulkWritableSavedViewRequest) GetIsGlobalDefaultOk() (*bool, bool)`

GetIsGlobalDefaultOk returns a tuple with the IsGlobalDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGlobalDefault

`func (o *BulkWritableSavedViewRequest) SetIsGlobalDefault(v bool)`

SetIsGlobalDefault sets IsGlobalDefault field to given value.

### HasIsGlobalDefault

`func (o *BulkWritableSavedViewRequest) HasIsGlobalDefault() bool`

HasIsGlobalDefault returns a boolean if a field has been set.

### GetIsShared

`func (o *BulkWritableSavedViewRequest) GetIsShared() bool`

GetIsShared returns the IsShared field if non-nil, zero value otherwise.

### GetIsSharedOk

`func (o *BulkWritableSavedViewRequest) GetIsSharedOk() (*bool, bool)`

GetIsSharedOk returns a tuple with the IsShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShared

`func (o *BulkWritableSavedViewRequest) SetIsShared(v bool)`

SetIsShared sets IsShared field to given value.

### HasIsShared

`func (o *BulkWritableSavedViewRequest) HasIsShared() bool`

HasIsShared returns a boolean if a field has been set.

### GetOwner

`func (o *BulkWritableSavedViewRequest) GetOwner() BulkWritableSavedViewRequestOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BulkWritableSavedViewRequest) GetOwnerOk() (*BulkWritableSavedViewRequestOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BulkWritableSavedViewRequest) SetOwner(v BulkWritableSavedViewRequestOwner)`

SetOwner sets Owner field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


