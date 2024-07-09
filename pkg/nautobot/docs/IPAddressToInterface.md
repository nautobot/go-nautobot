# IPAddressToInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
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

### NewIPAddressToInterface

`func NewIPAddressToInterface(id string, objectType string, display string, url string, naturalSlug string, ipAddress BulkWritableCableRequestStatus, ) *IPAddressToInterface`

NewIPAddressToInterface instantiates a new IPAddressToInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPAddressToInterfaceWithDefaults

`func NewIPAddressToInterfaceWithDefaults() *IPAddressToInterface`

NewIPAddressToInterfaceWithDefaults instantiates a new IPAddressToInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IPAddressToInterface) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IPAddressToInterface) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IPAddressToInterface) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *IPAddressToInterface) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IPAddressToInterface) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IPAddressToInterface) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *IPAddressToInterface) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *IPAddressToInterface) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *IPAddressToInterface) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *IPAddressToInterface) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IPAddressToInterface) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IPAddressToInterface) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *IPAddressToInterface) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *IPAddressToInterface) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *IPAddressToInterface) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetIsSource

`func (o *IPAddressToInterface) GetIsSource() bool`

GetIsSource returns the IsSource field if non-nil, zero value otherwise.

### GetIsSourceOk

`func (o *IPAddressToInterface) GetIsSourceOk() (*bool, bool)`

GetIsSourceOk returns a tuple with the IsSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSource

`func (o *IPAddressToInterface) SetIsSource(v bool)`

SetIsSource sets IsSource field to given value.

### HasIsSource

`func (o *IPAddressToInterface) HasIsSource() bool`

HasIsSource returns a boolean if a field has been set.

### GetIsDestination

`func (o *IPAddressToInterface) GetIsDestination() bool`

GetIsDestination returns the IsDestination field if non-nil, zero value otherwise.

### GetIsDestinationOk

`func (o *IPAddressToInterface) GetIsDestinationOk() (*bool, bool)`

GetIsDestinationOk returns a tuple with the IsDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDestination

`func (o *IPAddressToInterface) SetIsDestination(v bool)`

SetIsDestination sets IsDestination field to given value.

### HasIsDestination

`func (o *IPAddressToInterface) HasIsDestination() bool`

HasIsDestination returns a boolean if a field has been set.

### GetIsDefault

`func (o *IPAddressToInterface) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *IPAddressToInterface) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *IPAddressToInterface) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *IPAddressToInterface) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetIsPreferred

`func (o *IPAddressToInterface) GetIsPreferred() bool`

GetIsPreferred returns the IsPreferred field if non-nil, zero value otherwise.

### GetIsPreferredOk

`func (o *IPAddressToInterface) GetIsPreferredOk() (*bool, bool)`

GetIsPreferredOk returns a tuple with the IsPreferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPreferred

`func (o *IPAddressToInterface) SetIsPreferred(v bool)`

SetIsPreferred sets IsPreferred field to given value.

### HasIsPreferred

`func (o *IPAddressToInterface) HasIsPreferred() bool`

HasIsPreferred returns a boolean if a field has been set.

### GetIsPrimary

`func (o *IPAddressToInterface) GetIsPrimary() bool`

GetIsPrimary returns the IsPrimary field if non-nil, zero value otherwise.

### GetIsPrimaryOk

`func (o *IPAddressToInterface) GetIsPrimaryOk() (*bool, bool)`

GetIsPrimaryOk returns a tuple with the IsPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrimary

`func (o *IPAddressToInterface) SetIsPrimary(v bool)`

SetIsPrimary sets IsPrimary field to given value.

### HasIsPrimary

`func (o *IPAddressToInterface) HasIsPrimary() bool`

HasIsPrimary returns a boolean if a field has been set.

### GetIsSecondary

`func (o *IPAddressToInterface) GetIsSecondary() bool`

GetIsSecondary returns the IsSecondary field if non-nil, zero value otherwise.

### GetIsSecondaryOk

`func (o *IPAddressToInterface) GetIsSecondaryOk() (*bool, bool)`

GetIsSecondaryOk returns a tuple with the IsSecondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecondary

`func (o *IPAddressToInterface) SetIsSecondary(v bool)`

SetIsSecondary sets IsSecondary field to given value.

### HasIsSecondary

`func (o *IPAddressToInterface) HasIsSecondary() bool`

HasIsSecondary returns a boolean if a field has been set.

### GetIsStandby

`func (o *IPAddressToInterface) GetIsStandby() bool`

GetIsStandby returns the IsStandby field if non-nil, zero value otherwise.

### GetIsStandbyOk

`func (o *IPAddressToInterface) GetIsStandbyOk() (*bool, bool)`

GetIsStandbyOk returns a tuple with the IsStandby field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStandby

`func (o *IPAddressToInterface) SetIsStandby(v bool)`

SetIsStandby sets IsStandby field to given value.

### HasIsStandby

`func (o *IPAddressToInterface) HasIsStandby() bool`

HasIsStandby returns a boolean if a field has been set.

### GetIpAddress

`func (o *IPAddressToInterface) GetIpAddress() BulkWritableCableRequestStatus`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *IPAddressToInterface) GetIpAddressOk() (*BulkWritableCableRequestStatus, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *IPAddressToInterface) SetIpAddress(v BulkWritableCableRequestStatus)`

SetIpAddress sets IpAddress field to given value.


### GetInterface

`func (o *IPAddressToInterface) GetInterface() BulkWritableCircuitRequestTenant`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *IPAddressToInterface) GetInterfaceOk() (*BulkWritableCircuitRequestTenant, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *IPAddressToInterface) SetInterface(v BulkWritableCircuitRequestTenant)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *IPAddressToInterface) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### SetInterfaceNil

`func (o *IPAddressToInterface) SetInterfaceNil(b bool)`

 SetInterfaceNil sets the value for Interface to be an explicit nil

### UnsetInterface
`func (o *IPAddressToInterface) UnsetInterface()`

UnsetInterface ensures that no value is present for Interface, not even an explicit nil
### GetVmInterface

`func (o *IPAddressToInterface) GetVmInterface() BulkWritableCircuitRequestTenant`

GetVmInterface returns the VmInterface field if non-nil, zero value otherwise.

### GetVmInterfaceOk

`func (o *IPAddressToInterface) GetVmInterfaceOk() (*BulkWritableCircuitRequestTenant, bool)`

GetVmInterfaceOk returns a tuple with the VmInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmInterface

`func (o *IPAddressToInterface) SetVmInterface(v BulkWritableCircuitRequestTenant)`

SetVmInterface sets VmInterface field to given value.

### HasVmInterface

`func (o *IPAddressToInterface) HasVmInterface() bool`

HasVmInterface returns a boolean if a field has been set.

### SetVmInterfaceNil

`func (o *IPAddressToInterface) SetVmInterfaceNil(b bool)`

 SetVmInterfaceNil sets the value for VmInterface to be an explicit nil

### UnsetVmInterface
`func (o *IPAddressToInterface) UnsetVmInterface()`

UnsetVmInterface ensures that no value is present for VmInterface, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


