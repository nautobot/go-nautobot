# PatchedIPAddressToInterfaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSource** | Pointer to **bool** | Is source address on interface | [optional] 
**IsDestination** | Pointer to **bool** | Is destination address on interface | [optional] 
**IsDefault** | Pointer to **bool** | Is default address on interface | [optional] 
**IsPreferred** | Pointer to **bool** | Is preferred address on interface | [optional] 
**IsPrimary** | Pointer to **bool** | Is primary address on interface | [optional] 
**IsSecondary** | Pointer to **bool** | Is secondary address on interface | [optional] 
**IsStandby** | Pointer to **bool** | Is standby address on interface | [optional] 
**IpAddress** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**Interface** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**VmInterface** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 

## Methods

### NewPatchedIPAddressToInterfaceRequest

`func NewPatchedIPAddressToInterfaceRequest() *PatchedIPAddressToInterfaceRequest`

NewPatchedIPAddressToInterfaceRequest instantiates a new PatchedIPAddressToInterfaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedIPAddressToInterfaceRequestWithDefaults

`func NewPatchedIPAddressToInterfaceRequestWithDefaults() *PatchedIPAddressToInterfaceRequest`

NewPatchedIPAddressToInterfaceRequestWithDefaults instantiates a new PatchedIPAddressToInterfaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsSource

`func (o *PatchedIPAddressToInterfaceRequest) GetIsSource() bool`

GetIsSource returns the IsSource field if non-nil, zero value otherwise.

### GetIsSourceOk

`func (o *PatchedIPAddressToInterfaceRequest) GetIsSourceOk() (*bool, bool)`

GetIsSourceOk returns a tuple with the IsSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSource

`func (o *PatchedIPAddressToInterfaceRequest) SetIsSource(v bool)`

SetIsSource sets IsSource field to given value.

### HasIsSource

`func (o *PatchedIPAddressToInterfaceRequest) HasIsSource() bool`

HasIsSource returns a boolean if a field has been set.

### GetIsDestination

`func (o *PatchedIPAddressToInterfaceRequest) GetIsDestination() bool`

GetIsDestination returns the IsDestination field if non-nil, zero value otherwise.

### GetIsDestinationOk

`func (o *PatchedIPAddressToInterfaceRequest) GetIsDestinationOk() (*bool, bool)`

GetIsDestinationOk returns a tuple with the IsDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDestination

`func (o *PatchedIPAddressToInterfaceRequest) SetIsDestination(v bool)`

SetIsDestination sets IsDestination field to given value.

### HasIsDestination

`func (o *PatchedIPAddressToInterfaceRequest) HasIsDestination() bool`

HasIsDestination returns a boolean if a field has been set.

### GetIsDefault

`func (o *PatchedIPAddressToInterfaceRequest) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *PatchedIPAddressToInterfaceRequest) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *PatchedIPAddressToInterfaceRequest) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *PatchedIPAddressToInterfaceRequest) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetIsPreferred

`func (o *PatchedIPAddressToInterfaceRequest) GetIsPreferred() bool`

GetIsPreferred returns the IsPreferred field if non-nil, zero value otherwise.

### GetIsPreferredOk

`func (o *PatchedIPAddressToInterfaceRequest) GetIsPreferredOk() (*bool, bool)`

GetIsPreferredOk returns a tuple with the IsPreferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPreferred

`func (o *PatchedIPAddressToInterfaceRequest) SetIsPreferred(v bool)`

SetIsPreferred sets IsPreferred field to given value.

### HasIsPreferred

`func (o *PatchedIPAddressToInterfaceRequest) HasIsPreferred() bool`

HasIsPreferred returns a boolean if a field has been set.

### GetIsPrimary

`func (o *PatchedIPAddressToInterfaceRequest) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *PatchedIPAddressToInterfaceRequest) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *PatchedIPAddressToInterfaceRequest) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.

### HasIsPrimary

`func (o *PatchedIPAddressToInterfaceRequest) HasIsPrimary() bool`

HasIsPrimary returns a boolean if a field has been set.

### GetIsSecondary

`func (o *PatchedIPAddressToInterfaceRequest) GetIsSecondary() bool`

GetIsSecondary returns the IsSecondary field if non-nil, zero value otherwise.

### GetIsSecondaryOk

`func (o *PatchedIPAddressToInterfaceRequest) GetIsSecondaryOk() (*bool, bool)`

GetIsSecondaryOk returns a tuple with the IsSecondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecondary

`func (o *PatchedIPAddressToInterfaceRequest) SetIsSecondary(v bool)`

SetIsSecondary sets IsSecondary field to given value.

### HasIsSecondary

`func (o *PatchedIPAddressToInterfaceRequest) HasIsSecondary() bool`

HasIsSecondary returns a boolean if a field has been set.

### GetIsStandby

`func (o *PatchedIPAddressToInterfaceRequest) GetIsStandby() bool`

GetIsStandby returns the IsStandby field if non-nil, zero value otherwise.

### GetIsStandbyOk

`func (o *PatchedIPAddressToInterfaceRequest) GetIsStandbyOk() (*bool, bool)`

GetIsStandbyOk returns a tuple with the IsStandby field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStandby

`func (o *PatchedIPAddressToInterfaceRequest) SetIsStandby(v bool)`

SetIsStandby sets IsStandby field to given value.

### HasIsStandby

`func (o *PatchedIPAddressToInterfaceRequest) HasIsStandby() bool`

HasIsStandby returns a boolean if a field has been set.

### GetIpAddress

`func (o *PatchedIPAddressToInterfaceRequest) GetIpAddress() BulkWritableCableRequestStatus`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *PatchedIPAddressToInterfaceRequest) GetIpAddressOk() (*BulkWritableCableRequestStatus, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *PatchedIPAddressToInterfaceRequest) SetIpAddress(v BulkWritableCableRequestStatus)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *PatchedIPAddressToInterfaceRequest) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetInterface

`func (o *PatchedIPAddressToInterfaceRequest) GetInterface() BulkWritableCircuitRequestTenant`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *PatchedIPAddressToInterfaceRequest) GetInterfaceOk() (*BulkWritableCircuitRequestTenant, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *PatchedIPAddressToInterfaceRequest) SetInterface(v BulkWritableCircuitRequestTenant)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *PatchedIPAddressToInterfaceRequest) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### SetInterfaceNil

`func (o *PatchedIPAddressToInterfaceRequest) SetInterfaceNil(b bool)`

 SetInterfaceNil sets the value for Interface to be an explicit nil

### UnsetInterface
`func (o *PatchedIPAddressToInterfaceRequest) UnsetInterface()`

UnsetInterface ensures that no value is present for Interface, not even an explicit nil
### GetVmInterface

`func (o *PatchedIPAddressToInterfaceRequest) GetVmInterface() BulkWritableCircuitRequestTenant`

GetVmInterface returns the VmInterface field if non-nil, zero value otherwise.

### GetVmInterfaceOk

`func (o *PatchedIPAddressToInterfaceRequest) GetVmInterfaceOk() (*BulkWritableCircuitRequestTenant, bool)`

GetVmInterfaceOk returns a tuple with the VmInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmInterface

`func (o *PatchedIPAddressToInterfaceRequest) SetVmInterface(v BulkWritableCircuitRequestTenant)`

SetVmInterface sets VmInterface field to given value.

### HasVmInterface

`func (o *PatchedIPAddressToInterfaceRequest) HasVmInterface() bool`

HasVmInterface returns a boolean if a field has been set.

### SetVmInterfaceNil

`func (o *PatchedIPAddressToInterfaceRequest) SetVmInterfaceNil(b bool)`

 SetVmInterfaceNil sets the value for VmInterface to be an explicit nil

### UnsetVmInterface
`func (o *PatchedIPAddressToInterfaceRequest) UnsetVmInterface()`

UnsetVmInterface ensures that no value is present for VmInterface, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


