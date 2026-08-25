# ControllerManagedDeviceGroupRadioProfileAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ControllerManagedDeviceGroup** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**RadioProfile** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 

## Methods

### NewControllerManagedDeviceGroupRadioProfileAssignmentRequest

`func NewControllerManagedDeviceGroupRadioProfileAssignmentRequest(controllerManagedDeviceGroup BulkWritableCableRequestStatus, radioProfile BulkWritableCableRequestStatus, ) *ControllerManagedDeviceGroupRadioProfileAssignmentRequest`

NewControllerManagedDeviceGroupRadioProfileAssignmentRequest instantiates a new ControllerManagedDeviceGroupRadioProfileAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewControllerManagedDeviceGroupRadioProfileAssignmentRequestWithDefaults

`func NewControllerManagedDeviceGroupRadioProfileAssignmentRequestWithDefaults() *ControllerManagedDeviceGroupRadioProfileAssignmentRequest`

NewControllerManagedDeviceGroupRadioProfileAssignmentRequestWithDefaults instantiates a new ControllerManagedDeviceGroupRadioProfileAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ControllerManagedDeviceGroupRadioProfileAssignmentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ControllerManagedDeviceGroupRadioProfileAssignmentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ControllerManagedDeviceGroupRadioProfileAssignmentRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ControllerManagedDeviceGroupRadioProfileAssignmentRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetControllerManagedDeviceGroup

`func (o *ControllerManagedDeviceGroupRadioProfileAssignmentRequest) GetControllerManagedDeviceGroup() BulkWritableCableRequestStatus`

GetControllerManagedDeviceGroup returns the ControllerManagedDeviceGroup field if non-nil, zero value otherwise.

### GetControllerManagedDeviceGroupOk

`func (o *ControllerManagedDeviceGroupRadioProfileAssignmentRequest) GetControllerManagedDeviceGroupOk() (*BulkWritableCableRequestStatus, bool)`

GetControllerManagedDeviceGroupOk returns a tuple with the ControllerManagedDeviceGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerManagedDeviceGroup

`func (o *ControllerManagedDeviceGroupRadioProfileAssignmentRequest) SetControllerManagedDeviceGroup(v BulkWritableCableRequestStatus)`

SetControllerManagedDeviceGroup sets ControllerManagedDeviceGroup field to given value.


### GetRadioProfile

`func (o *ControllerManagedDeviceGroupRadioProfileAssignmentRequest) GetRadioProfile() BulkWritableCableRequestStatus`

GetRadioProfile returns the RadioProfile field if non-nil, zero value otherwise.

### GetRadioProfileOk

`func (o *ControllerManagedDeviceGroupRadioProfileAssignmentRequest) GetRadioProfileOk() (*BulkWritableCableRequestStatus, bool)`

GetRadioProfileOk returns a tuple with the RadioProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadioProfile

`func (o *ControllerManagedDeviceGroupRadioProfileAssignmentRequest) SetRadioProfile(v BulkWritableCableRequestStatus)`

SetRadioProfile sets RadioProfile field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


