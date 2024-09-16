# AvailableIP

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpVersion** | **int32** |  | [readonly] 
**Address** | **string** |  | [readonly] 

## Methods

### NewAvailableIP

`func NewAvailableIP(ipVersion int32, address string, ) *AvailableIP`

NewAvailableIP instantiates a new AvailableIP object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailableIPWithDefaults

`func NewAvailableIPWithDefaults() *AvailableIP`

NewAvailableIPWithDefaults instantiates a new AvailableIP object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpVersion

`func (o *AvailableIP) GetIpVersion() int32`

GetIpVersion returns the IpVersion field if non-nil, zero value otherwise.

### GetIpVersionOk

`func (o *AvailableIP) GetIpVersionOk() (*int32, bool)`

GetIpVersionOk returns a tuple with the IpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpVersion

`func (o *AvailableIP) SetIpVersion(v int32)`

SetIpVersion sets IpVersion field to given value.


### GetAddress

`func (o *AvailableIP) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AvailableIP) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AvailableIP) SetAddress(v string)`

SetAddress sets Address field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


