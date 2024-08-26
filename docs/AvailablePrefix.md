# AvailablePrefix

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpVersion** | **int32** |  | [readonly] 
**Prefix** | **string** |  | [readonly] 

## Methods

### NewAvailablePrefix

`func NewAvailablePrefix(ipVersion int32, prefix string, ) *AvailablePrefix`

NewAvailablePrefix instantiates a new AvailablePrefix object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailablePrefixWithDefaults

`func NewAvailablePrefixWithDefaults() *AvailablePrefix`

NewAvailablePrefixWithDefaults instantiates a new AvailablePrefix object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpVersion

`func (o *AvailablePrefix) GetIpVersion() int32`

GetIpVersion returns the IpVersion field if non-nil, zero value otherwise.

### GetIpVersionOk

`func (o *AvailablePrefix) GetIpVersionOk() (*int32, bool)`

GetIpVersionOk returns a tuple with the IpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpVersion

`func (o *AvailablePrefix) SetIpVersion(v int32)`

SetIpVersion sets IpVersion field to given value.


### GetPrefix

`func (o *AvailablePrefix) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *AvailablePrefix) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *AvailablePrefix) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


