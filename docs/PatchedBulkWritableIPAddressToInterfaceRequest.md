# PatchedBulkWritableIPAddressToInterfaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
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

### NewPatchedBulkWritableIPAddressToInterfaceRequest

`func NewPatchedBulkWritableIPAddressToInterfaceRequest(id string, ) *PatchedBulkWritableIPAddressToInterfaceRequest`

NewPatchedBulkWritableIPAddressToInterfaceRequest instantiates a new PatchedBulkWritableIPAddressToInterfaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableIPAddressToInterfaceRequestWithDefaults

`func NewPatchedBulkWritableIPAddressToInterfaceRequestWithDefaults() *PatchedBulkWritableIPAddressToInterfaceRequest`

NewPatchedBulkWritableIPAddressToInterfaceRequestWithDefaults instantiates a new PatchedBulkWritableIPAddressToInterfaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableIPAddressToInterfaceRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableIPAddressToInterfaceRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableIPAddressToInterfaceRequest) SetId(v string)`

SetId sets Id field to given value.


### GetIsSource

`func (o *PatchedBulkWritableIPAddressToInterfaceRequest) GetIsSource() bool`

GetIsSource returns the IsSource field if non-nil, zero value otherwise.

### GetIsSourceOk

`func (o *PatchedBulkWritableIPAddressToInterfaceRequest) GetIsSourceOk() (*bool, bool)`

GetIsSourceOk returns a tuple with the IsSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSource

`func (o *PatchedBulkWritableIPAddressToInterfaceRequest) SetIsSource(v bool)`

SetIsSource sets IsSource field to given value.

### HasIsSource

`func (o *PatchedBulkWritableIPAddressToInterfaceRequest) HasIsSource() bool`

HasIsSource returns a boolean if a field has been set.

### GetIsDestination

`func (o *PatchedBulkWritableIPAddressToInterfaceRequest) GetIsDestination() bool`

GetIsDestination returns the IsDestination field if non-nil, zero value otherwise.

### GetIsDestinationOk

`func (o *PatchedBulkWritableIPAddressToInterfaceRequest) GetIsDestinationOk() (*bool, bool)`

GetIsDestinationOk returns a tuple with the IsDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDestination

`func (o *PatchedBulkWritableIPAddressToInterfaceRequest) SetIsDestination(v bool)`

SetIsDestination sets IsDestination field to given value.

### HasIsDestination

`func (o *PatchedBulkWritableIPAddressToInterfaceRequest) HasIsDestination() bool`

HasIsDestination returns a boolean if a field has been set.

### GetIsDefault

`func (o *PatchedBulkWritableIPAddressToInterfaceRequest) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *PatchedBulkWritableIPAddressToInterfaceRequest) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *PatchedBulkWritableIPAddressToInterfaceRequest) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *PatchedBulkWritableIPAddressToInterfaceRequest) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetIsPreferred

`func (o *PatchedBulkWritableIPAddressToInterfaceRequest) GetIsPreferred() bool`

GetIsPreferred returns the IsPreferred field if non-nil, zero value otherwise.

### GetIsPreferredOk

`func (o *PatchedBulkWritableIPAddressToInterfaceRequest) GetIsPreferredOk() (*bool, bool)`

GetIsPreferredOk returns a tuple with the IsPreferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPreferred

`func (o *PatchedBulkWritableIPAddressToInterfaceRequest) SetIsPreferred(v bool)`

SetIsPreferred sets IsPreferred field to given value.

### HasIsPreferred

`func (o *PatchedBulkWritableIPAddressToInterfaceRequest) HasIsPreferred() bool`

HasIsPreferred returns a boolean if a field has been set.

### GetIsPrimary

`func (o *PatchedBulkWritableIPAddressToInterfaceRequest) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *PatchedBulkWritableIPAddressToInterfaceRequest) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *PatchedBulkWritableIPAddressToInterfaceRequest) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.

### HasIsPrimary

`func (o *PatchedBulkWritableIPAddressToInterfaceRequest) HasIsPrimary() bool`

HasIsPrimary returns a boolean if a field has been set.

### GetIsSecondary

`func (o *PatchedBulkWritableIPAddressToInterfaceRequest) GetIsSecondary() bool`

GetIsSecondary returns the IsSecondary field if non-nil, zero value otherwise.

### GetIsSecondaryOk

`func (o *PatchedBulkWritableIPAddressToInterfaceRequest) GetIsSecondaryOk() (*bool, bool)`

GetIsSecondaryOk returns a tuple with the IsSecondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecondary

`func (o *PatchedBulkWritableIPAddressToInterfaceRequest) SetIsSecondary(v bool)`

SetIsSecondary sets IsSecondary field to given value.

### HasIsSecondary

`func (o *PatchedBulkWritableIPAddressToInterfaceRequest) HasIsSecondary() bool`

HasIsSecondary returns a boolean if a field has been set.

### GetIsStandby

`func (o *PatchedBulkWritableIPAddressToInterfaceRequest) GetIsStandby() bool`

GetIsStandby returns the IsStandby field if non-nil, zero value otherwise.

### GetIsStandbyOk

`func (o *PatchedBulkWritableIPAddressToInterfaceRequest) GetIsStandbyOk() (*bool, bool)`

GetIsStandbyOk returns a tuple with the IsStandby field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStandby

`func (o *PatchedBulkWritableIPAddressToInterfaceRequest) SetIsStandby(v bool)`

SetIsStandby sets IsStandby field to given value.

### HasIsStandby

`func (o *PatchedBulkWritableIPAddressToInterfaceRequest) HasIsStandby() bool`

HasIsStandby returns a boolean if a field has been set.

### GetIpAddress

`func (o *PatchedBulkWritableIPAddressToInterfaceRequest) GetIpAddress() BulkWritableCableRequestStatus`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *PatchedBulkWritableIPAddressToInterfaceRequest) GetIpAddressOk() (*BulkWritableCableRequestStatus, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *PatchedBulkWritableIPAddressToInterfaceRequest) SetIpAddress(v BulkWritableCableRequestStatus)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *PatchedBulkWritableIPAddressToInterfaceRequest) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetInterface

`func (o *PatchedBulkWritableIPAddressToInterfaceRequest) GetInterface() BulkWritableCircuitRequestTenant`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *PatchedBulkWritableIPAddressToInterfaceRequest) GetInterfaceOk() (*BulkWritableCircuitRequestTenant, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *PatchedBulkWritableIPAddressToInterfaceRequest) SetInterface(v BulkWritableCircuitRequestTenant)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *PatchedBulkWritableIPAddressToInterfaceRequest) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### SetInterfaceNil

`func (o *PatchedBulkWritableIPAddressToInterfaceRequest) SetInterfaceNil(b bool)`

 SetInterfaceNil sets the value for Interface to be an explicit nil

### UnsetInterface
`func (o *PatchedBulkWritableIPAddressToInterfaceRequest) UnsetInterface()`

UnsetInterface ensures that no value is present for Interface, not even an explicit nil
### GetVmInterface

`func (o *PatchedBulkWritableIPAddressToInterfaceRequest) GetVmInterface() BulkWritableCircuitRequestTenant`

GetVmInterface returns the VmInterface field if non-nil, zero value otherwise.

### GetVmInterfaceOk

`func (o *PatchedBulkWritableIPAddressToInterfaceRequest) GetVmInterfaceOk() (*BulkWritableCircuitRequestTenant, bool)`

GetVmInterfaceOk returns a tuple with the VmInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmInterface

`func (o *PatchedBulkWritableIPAddressToInterfaceRequest) SetVmInterface(v BulkWritableCircuitRequestTenant)`

SetVmInterface sets VmInterface field to given value.

### HasVmInterface

`func (o *PatchedBulkWritableIPAddressToInterfaceRequest) HasVmInterface() bool`

HasVmInterface returns a boolean if a field has been set.

### SetVmInterfaceNil

`func (o *PatchedBulkWritableIPAddressToInterfaceRequest) SetVmInterfaceNil(b bool)`

 SetVmInterfaceNil sets the value for VmInterface to be an explicit nil

### UnsetVmInterface
`func (o *PatchedBulkWritableIPAddressToInterfaceRequest) UnsetVmInterface()`

UnsetVmInterface ensures that no value is present for VmInterface, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


