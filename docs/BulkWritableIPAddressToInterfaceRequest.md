# BulkWritableIPAddressToInterfaceRequest

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
**IpAddress** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Interface** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**VmInterface** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 

## Methods

### NewBulkWritableIPAddressToInterfaceRequest

`func NewBulkWritableIPAddressToInterfaceRequest(id string, ipAddress BulkWritableCableRequestStatus, ) *BulkWritableIPAddressToInterfaceRequest`

NewBulkWritableIPAddressToInterfaceRequest instantiates a new BulkWritableIPAddressToInterfaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableIPAddressToInterfaceRequestWithDefaults

`func NewBulkWritableIPAddressToInterfaceRequestWithDefaults() *BulkWritableIPAddressToInterfaceRequest`

NewBulkWritableIPAddressToInterfaceRequestWithDefaults instantiates a new BulkWritableIPAddressToInterfaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableIPAddressToInterfaceRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableIPAddressToInterfaceRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableIPAddressToInterfaceRequest) SetId(v string)`

SetId sets Id field to given value.


### GetIsSource

`func (o *BulkWritableIPAddressToInterfaceRequest) GetIsSource() bool`

GetIsSource returns the IsSource field if non-nil, zero value otherwise.

### GetIsSourceOk

`func (o *BulkWritableIPAddressToInterfaceRequest) GetIsSourceOk() (*bool, bool)`

GetIsSourceOk returns a tuple with the IsSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSource

`func (o *BulkWritableIPAddressToInterfaceRequest) SetIsSource(v bool)`

SetIsSource sets IsSource field to given value.

### HasIsSource

`func (o *BulkWritableIPAddressToInterfaceRequest) HasIsSource() bool`

HasIsSource returns a boolean if a field has been set.

### GetIsDestination

`func (o *BulkWritableIPAddressToInterfaceRequest) GetIsDestination() bool`

GetIsDestination returns the IsDestination field if non-nil, zero value otherwise.

### GetIsDestinationOk

`func (o *BulkWritableIPAddressToInterfaceRequest) GetIsDestinationOk() (*bool, bool)`

GetIsDestinationOk returns a tuple with the IsDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDestination

`func (o *BulkWritableIPAddressToInterfaceRequest) SetIsDestination(v bool)`

SetIsDestination sets IsDestination field to given value.

### HasIsDestination

`func (o *BulkWritableIPAddressToInterfaceRequest) HasIsDestination() bool`

HasIsDestination returns a boolean if a field has been set.

### GetIsDefault

`func (o *BulkWritableIPAddressToInterfaceRequest) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *BulkWritableIPAddressToInterfaceRequest) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *BulkWritableIPAddressToInterfaceRequest) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *BulkWritableIPAddressToInterfaceRequest) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetIsPreferred

`func (o *BulkWritableIPAddressToInterfaceRequest) GetIsPreferred() bool`

GetIsPreferred returns the IsPreferred field if non-nil, zero value otherwise.

### GetIsPreferredOk

`func (o *BulkWritableIPAddressToInterfaceRequest) GetIsPreferredOk() (*bool, bool)`

GetIsPreferredOk returns a tuple with the IsPreferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPreferred

`func (o *BulkWritableIPAddressToInterfaceRequest) SetIsPreferred(v bool)`

SetIsPreferred sets IsPreferred field to given value.

### HasIsPreferred

`func (o *BulkWritableIPAddressToInterfaceRequest) HasIsPreferred() bool`

HasIsPreferred returns a boolean if a field has been set.

### GetIsPrimary

`func (o *BulkWritableIPAddressToInterfaceRequest) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *BulkWritableIPAddressToInterfaceRequest) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *BulkWritableIPAddressToInterfaceRequest) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.

### HasIsPrimary

`func (o *BulkWritableIPAddressToInterfaceRequest) HasIsPrimary() bool`

HasIsPrimary returns a boolean if a field has been set.

### GetIsSecondary

`func (o *BulkWritableIPAddressToInterfaceRequest) GetIsSecondary() bool`

GetIsSecondary returns the IsSecondary field if non-nil, zero value otherwise.

### GetIsSecondaryOk

`func (o *BulkWritableIPAddressToInterfaceRequest) GetIsSecondaryOk() (*bool, bool)`

GetIsSecondaryOk returns a tuple with the IsSecondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecondary

`func (o *BulkWritableIPAddressToInterfaceRequest) SetIsSecondary(v bool)`

SetIsSecondary sets IsSecondary field to given value.

### HasIsSecondary

`func (o *BulkWritableIPAddressToInterfaceRequest) HasIsSecondary() bool`

HasIsSecondary returns a boolean if a field has been set.

### GetIsStandby

`func (o *BulkWritableIPAddressToInterfaceRequest) GetIsStandby() bool`

GetIsStandby returns the IsStandby field if non-nil, zero value otherwise.

### GetIsStandbyOk

`func (o *BulkWritableIPAddressToInterfaceRequest) GetIsStandbyOk() (*bool, bool)`

GetIsStandbyOk returns a tuple with the IsStandby field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStandby

`func (o *BulkWritableIPAddressToInterfaceRequest) SetIsStandby(v bool)`

SetIsStandby sets IsStandby field to given value.

### HasIsStandby

`func (o *BulkWritableIPAddressToInterfaceRequest) HasIsStandby() bool`

HasIsStandby returns a boolean if a field has been set.

### GetIpAddress

`func (o *BulkWritableIPAddressToInterfaceRequest) GetIpAddress() BulkWritableCableRequestStatus`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *BulkWritableIPAddressToInterfaceRequest) GetIpAddressOk() (*BulkWritableCableRequestStatus, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *BulkWritableIPAddressToInterfaceRequest) SetIpAddress(v BulkWritableCableRequestStatus)`

SetIpAddress sets IpAddress field to given value.


### GetInterface

`func (o *BulkWritableIPAddressToInterfaceRequest) GetInterface() BulkWritableCircuitRequestTenant`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *BulkWritableIPAddressToInterfaceRequest) GetInterfaceOk() (*BulkWritableCircuitRequestTenant, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *BulkWritableIPAddressToInterfaceRequest) SetInterface(v BulkWritableCircuitRequestTenant)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *BulkWritableIPAddressToInterfaceRequest) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### SetInterfaceNil

`func (o *BulkWritableIPAddressToInterfaceRequest) SetInterfaceNil(b bool)`

 SetInterfaceNil sets the value for Interface to be an explicit nil

### UnsetInterface
`func (o *BulkWritableIPAddressToInterfaceRequest) UnsetInterface()`

UnsetInterface ensures that no value is present for Interface, not even an explicit nil
### GetVmInterface

`func (o *BulkWritableIPAddressToInterfaceRequest) GetVmInterface() BulkWritableCircuitRequestTenant`

GetVmInterface returns the VmInterface field if non-nil, zero value otherwise.

### GetVmInterfaceOk

`func (o *BulkWritableIPAddressToInterfaceRequest) GetVmInterfaceOk() (*BulkWritableCircuitRequestTenant, bool)`

GetVmInterfaceOk returns a tuple with the VmInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmInterface

`func (o *BulkWritableIPAddressToInterfaceRequest) SetVmInterface(v BulkWritableCircuitRequestTenant)`

SetVmInterface sets VmInterface field to given value.

### HasVmInterface

`func (o *BulkWritableIPAddressToInterfaceRequest) HasVmInterface() bool`

HasVmInterface returns a boolean if a field has been set.

### SetVmInterfaceNil

`func (o *BulkWritableIPAddressToInterfaceRequest) SetVmInterfaceNil(b bool)`

 SetVmInterfaceNil sets the value for VmInterface to be an explicit nil

### UnsetVmInterface
`func (o *BulkWritableIPAddressToInterfaceRequest) UnsetVmInterface()`

UnsetVmInterface ensures that no value is present for VmInterface, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


