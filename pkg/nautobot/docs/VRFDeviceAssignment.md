# VRFDeviceAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**Rd** | Pointer to **NullableString** | Unique route distinguisher (as defined in RFC 4364) | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Vrf** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Device** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**VirtualMachine** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 

## Methods

### NewVRFDeviceAssignment

`func NewVRFDeviceAssignment(id string, objectType string, display string, url string, naturalSlug string, vrf BulkWritableCableRequestStatus, ) *VRFDeviceAssignment`

NewVRFDeviceAssignment instantiates a new VRFDeviceAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVRFDeviceAssignmentWithDefaults

`func NewVRFDeviceAssignmentWithDefaults() *VRFDeviceAssignment`

NewVRFDeviceAssignmentWithDefaults instantiates a new VRFDeviceAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VRFDeviceAssignment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VRFDeviceAssignment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VRFDeviceAssignment) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *VRFDeviceAssignment) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VRFDeviceAssignment) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VRFDeviceAssignment) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *VRFDeviceAssignment) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *VRFDeviceAssignment) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *VRFDeviceAssignment) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *VRFDeviceAssignment) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VRFDeviceAssignment) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VRFDeviceAssignment) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *VRFDeviceAssignment) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *VRFDeviceAssignment) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *VRFDeviceAssignment) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetRd

`func (o *VRFDeviceAssignment) GetRd() string`

GetRd returns the Rd field if non-nil, zero value otherwise.

### GetRdOk

`func (o *VRFDeviceAssignment) GetRdOk() (*string, bool)`

GetRdOk returns a tuple with the Rd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRd

`func (o *VRFDeviceAssignment) SetRd(v string)`

SetRd sets Rd field to given value.

### HasRd

`func (o *VRFDeviceAssignment) HasRd() bool`

HasRd returns a boolean if a field has been set.

### SetRdNil

`func (o *VRFDeviceAssignment) SetRdNil(b bool)`

 SetRdNil sets the value for Rd to be an explicit nil

### UnsetRd
`func (o *VRFDeviceAssignment) UnsetRd()`

UnsetRd ensures that no value is present for Rd, not even an explicit nil
### GetName

`func (o *VRFDeviceAssignment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VRFDeviceAssignment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VRFDeviceAssignment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VRFDeviceAssignment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVrf

`func (o *VRFDeviceAssignment) GetVrf() BulkWritableCableRequestStatus`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *VRFDeviceAssignment) GetVrfOk() (*BulkWritableCableRequestStatus, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *VRFDeviceAssignment) SetVrf(v BulkWritableCableRequestStatus)`

SetVrf sets Vrf field to given value.


### GetDevice

`func (o *VRFDeviceAssignment) GetDevice() BulkWritableCircuitRequestTenant`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *VRFDeviceAssignment) GetDeviceOk() (*BulkWritableCircuitRequestTenant, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *VRFDeviceAssignment) SetDevice(v BulkWritableCircuitRequestTenant)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *VRFDeviceAssignment) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### SetDeviceNil

`func (o *VRFDeviceAssignment) SetDeviceNil(b bool)`

 SetDeviceNil sets the value for Device to be an explicit nil

### UnsetDevice
`func (o *VRFDeviceAssignment) UnsetDevice()`

UnsetDevice ensures that no value is present for Device, not even an explicit nil
### GetVirtualMachine

`func (o *VRFDeviceAssignment) GetVirtualMachine() BulkWritableCircuitRequestTenant`

GetVirtualMachine returns the VirtualMachine field if non-nil, zero value otherwise.

### GetVirtualMachineOk

`func (o *VRFDeviceAssignment) GetVirtualMachineOk() (*BulkWritableCircuitRequestTenant, bool)`

GetVirtualMachineOk returns a tuple with the VirtualMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachine

`func (o *VRFDeviceAssignment) SetVirtualMachine(v BulkWritableCircuitRequestTenant)`

SetVirtualMachine sets VirtualMachine field to given value.

### HasVirtualMachine

`func (o *VRFDeviceAssignment) HasVirtualMachine() bool`

HasVirtualMachine returns a boolean if a field has been set.

### SetVirtualMachineNil

`func (o *VRFDeviceAssignment) SetVirtualMachineNil(b bool)`

 SetVirtualMachineNil sets the value for VirtualMachine to be an explicit nil

### UnsetVirtualMachine
`func (o *VRFDeviceAssignment) UnsetVirtualMachine()`

UnsetVirtualMachine ensures that no value is present for VirtualMachine, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


