# IPAddressToInterfaceRequest

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
**IpAddress** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Interface** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**VmInterface** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 

## Methods

### NewIPAddressToInterfaceRequest

`func NewIPAddressToInterfaceRequest(ipAddress BulkWritableCableRequestStatus, ) *IPAddressToInterfaceRequest`

NewIPAddressToInterfaceRequest instantiates a new IPAddressToInterfaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPAddressToInterfaceRequestWithDefaults

`func NewIPAddressToInterfaceRequestWithDefaults() *IPAddressToInterfaceRequest`

NewIPAddressToInterfaceRequestWithDefaults instantiates a new IPAddressToInterfaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsSource

`func (o *IPAddressToInterfaceRequest) GetIsSource() bool`

GetIsSource returns the IsSource field if non-nil, zero value otherwise.

### GetIsSourceOk

`func (o *IPAddressToInterfaceRequest) GetIsSourceOk() (*bool, bool)`

GetIsSourceOk returns a tuple with the IsSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSource

`func (o *IPAddressToInterfaceRequest) SetIsSource(v bool)`

SetIsSource sets IsSource field to given value.

### HasIsSource

`func (o *IPAddressToInterfaceRequest) HasIsSource() bool`

HasIsSource returns a boolean if a field has been set.

### GetIsDestination

`func (o *IPAddressToInterfaceRequest) GetIsDestination() bool`

GetIsDestination returns the IsDestination field if non-nil, zero value otherwise.

### GetIsDestinationOk

`func (o *IPAddressToInterfaceRequest) GetIsDestinationOk() (*bool, bool)`

GetIsDestinationOk returns a tuple with the IsDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDestination

`func (o *IPAddressToInterfaceRequest) SetIsDestination(v bool)`

SetIsDestination sets IsDestination field to given value.

### HasIsDestination

`func (o *IPAddressToInterfaceRequest) HasIsDestination() bool`

HasIsDestination returns a boolean if a field has been set.

### GetIsDefault

`func (o *IPAddressToInterfaceRequest) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *IPAddressToInterfaceRequest) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *IPAddressToInterfaceRequest) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *IPAddressToInterfaceRequest) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetIsPreferred

`func (o *IPAddressToInterfaceRequest) GetIsPreferred() bool`

GetIsPreferred returns the IsPreferred field if non-nil, zero value otherwise.

### GetIsPreferredOk

`func (o *IPAddressToInterfaceRequest) GetIsPreferredOk() (*bool, bool)`

GetIsPreferredOk returns a tuple with the IsPreferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPreferred

`func (o *IPAddressToInterfaceRequest) SetIsPreferred(v bool)`

SetIsPreferred sets IsPreferred field to given value.

### HasIsPreferred

`func (o *IPAddressToInterfaceRequest) HasIsPreferred() bool`

HasIsPreferred returns a boolean if a field has been set.

### GetIsPrimary

`func (o *IPAddressToInterfaceRequest) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *IPAddressToInterfaceRequest) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *IPAddressToInterfaceRequest) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.

### HasIsPrimary

`func (o *IPAddressToInterfaceRequest) HasIsPrimary() bool`

HasIsPrimary returns a boolean if a field has been set.

### GetIsSecondary

`func (o *IPAddressToInterfaceRequest) GetIsSecondary() bool`

GetIsSecondary returns the IsSecondary field if non-nil, zero value otherwise.

### GetIsSecondaryOk

`func (o *IPAddressToInterfaceRequest) GetIsSecondaryOk() (*bool, bool)`

GetIsSecondaryOk returns a tuple with the IsSecondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecondary

`func (o *IPAddressToInterfaceRequest) SetIsSecondary(v bool)`

SetIsSecondary sets IsSecondary field to given value.

### HasIsSecondary

`func (o *IPAddressToInterfaceRequest) HasIsSecondary() bool`

HasIsSecondary returns a boolean if a field has been set.

### GetIsStandby

`func (o *IPAddressToInterfaceRequest) GetIsStandby() bool`

GetIsStandby returns the IsStandby field if non-nil, zero value otherwise.

### GetIsStandbyOk

`func (o *IPAddressToInterfaceRequest) GetIsStandbyOk() (*bool, bool)`

GetIsStandbyOk returns a tuple with the IsStandby field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStandby

`func (o *IPAddressToInterfaceRequest) SetIsStandby(v bool)`

SetIsStandby sets IsStandby field to given value.

### HasIsStandby

`func (o *IPAddressToInterfaceRequest) HasIsStandby() bool`

HasIsStandby returns a boolean if a field has been set.

### GetIpAddress

`func (o *IPAddressToInterfaceRequest) GetIpAddress() BulkWritableCableRequestStatus`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *IPAddressToInterfaceRequest) GetIpAddressOk() (*BulkWritableCableRequestStatus, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *IPAddressToInterfaceRequest) SetIpAddress(v BulkWritableCableRequestStatus)`

SetIpAddress sets IpAddress field to given value.


### GetInterface

`func (o *IPAddressToInterfaceRequest) GetInterface() BulkWritableCircuitRequestTenant`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *IPAddressToInterfaceRequest) GetInterfaceOk() (*BulkWritableCircuitRequestTenant, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *IPAddressToInterfaceRequest) SetInterface(v BulkWritableCircuitRequestTenant)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *IPAddressToInterfaceRequest) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### SetInterfaceNil

`func (o *IPAddressToInterfaceRequest) SetInterfaceNil(b bool)`

 SetInterfaceNil sets the value for Interface to be an explicit nil

### UnsetInterface
`func (o *IPAddressToInterfaceRequest) UnsetInterface()`

UnsetInterface ensures that no value is present for Interface, not even an explicit nil
### GetVmInterface

`func (o *IPAddressToInterfaceRequest) GetVmInterface() BulkWritableCircuitRequestTenant`

GetVmInterface returns the VmInterface field if non-nil, zero value otherwise.

### GetVmInterfaceOk

`func (o *IPAddressToInterfaceRequest) GetVmInterfaceOk() (*BulkWritableCircuitRequestTenant, bool)`

GetVmInterfaceOk returns a tuple with the VmInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmInterface

`func (o *IPAddressToInterfaceRequest) SetVmInterface(v BulkWritableCircuitRequestTenant)`

SetVmInterface sets VmInterface field to given value.

### HasVmInterface

`func (o *IPAddressToInterfaceRequest) HasVmInterface() bool`

HasVmInterface returns a boolean if a field has been set.

### SetVmInterfaceNil

`func (o *IPAddressToInterfaceRequest) SetVmInterfaceNil(b bool)`

 SetVmInterfaceNil sets the value for VmInterface to be an explicit nil

### UnsetVmInterface
`func (o *IPAddressToInterfaceRequest) UnsetVmInterface()`

UnsetVmInterface ensures that no value is present for VmInterface, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


