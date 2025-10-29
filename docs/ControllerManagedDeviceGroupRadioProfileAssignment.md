# ControllerManagedDeviceGroupRadioProfileAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**ControllerManagedDeviceGroup** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**RadioProfile** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 

## Methods

### NewControllerManagedDeviceGroupRadioProfileAssignment

`func NewControllerManagedDeviceGroupRadioProfileAssignment(objectType string, display string, url string, naturalSlug string, controllerManagedDeviceGroup BulkWritableCableRequestStatus, radioProfile BulkWritableCableRequestStatus, ) *ControllerManagedDeviceGroupRadioProfileAssignment`

NewControllerManagedDeviceGroupRadioProfileAssignment instantiates a new ControllerManagedDeviceGroupRadioProfileAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllerManagedDeviceGroupRadioProfileAssignmentWithDefaults

`func NewControllerManagedDeviceGroupRadioProfileAssignmentWithDefaults() *ControllerManagedDeviceGroupRadioProfileAssignment`

NewControllerManagedDeviceGroupRadioProfileAssignmentWithDefaults instantiates a new ControllerManagedDeviceGroupRadioProfileAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ControllerManagedDeviceGroupRadioProfileAssignment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ControllerManagedDeviceGroupRadioProfileAssignment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ControllerManagedDeviceGroupRadioProfileAssignment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ControllerManagedDeviceGroupRadioProfileAssignment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectType

`func (o *ControllerManagedDeviceGroupRadioProfileAssignment) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ControllerManagedDeviceGroupRadioProfileAssignment) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ControllerManagedDeviceGroupRadioProfileAssignment) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *ControllerManagedDeviceGroupRadioProfileAssignment) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *ControllerManagedDeviceGroupRadioProfileAssignment) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *ControllerManagedDeviceGroupRadioProfileAssignment) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *ControllerManagedDeviceGroupRadioProfileAssignment) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ControllerManagedDeviceGroupRadioProfileAssignment) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ControllerManagedDeviceGroupRadioProfileAssignment) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *ControllerManagedDeviceGroupRadioProfileAssignment) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *ControllerManagedDeviceGroupRadioProfileAssignment) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *ControllerManagedDeviceGroupRadioProfileAssignment) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetControllerManagedDeviceGroup

`func (o *ControllerManagedDeviceGroupRadioProfileAssignment) GetControllerManagedDeviceGroup() BulkWritableCableRequestStatus`

GetControllerManagedDeviceGroup returns the ControllerManagedDeviceGroup field if non-nil, zero value otherwise.

### GetControllerManagedDeviceGroupOk

`func (o *ControllerManagedDeviceGroupRadioProfileAssignment) GetControllerManagedDeviceGroupOk() (*BulkWritableCableRequestStatus, bool)`

GetControllerManagedDeviceGroupOk returns a tuple with the ControllerManagedDeviceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerManagedDeviceGroup

`func (o *ControllerManagedDeviceGroupRadioProfileAssignment) SetControllerManagedDeviceGroup(v BulkWritableCableRequestStatus)`

SetControllerManagedDeviceGroup sets ControllerManagedDeviceGroup field to given value.


### GetRadioProfile

`func (o *ControllerManagedDeviceGroupRadioProfileAssignment) GetRadioProfile() BulkWritableCableRequestStatus`

GetRadioProfile returns the RadioProfile field if non-nil, zero value otherwise.

### GetRadioProfileOk

`func (o *ControllerManagedDeviceGroupRadioProfileAssignment) GetRadioProfileOk() (*BulkWritableCableRequestStatus, bool)`

GetRadioProfileOk returns a tuple with the RadioProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadioProfile

`func (o *ControllerManagedDeviceGroupRadioProfileAssignment) SetRadioProfile(v BulkWritableCableRequestStatus)`

SetRadioProfile sets RadioProfile field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


