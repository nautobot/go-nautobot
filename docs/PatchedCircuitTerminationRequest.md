# PatchedCircuitTerminationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TermSide** | Pointer to [**TermSideEnum**](TermSideEnum.md) |  | [optional] 
**PortSpeed** | Pointer to **NullableInt32** |  | [optional] 
**UpstreamSpeed** | Pointer to **NullableInt32** | Upstream speed, if different from port speed | [optional] 
**XconnectId** | Pointer to **string** |  | [optional] 
**PpInfo** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Circuit** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**Location** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**ProviderNetwork** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**CloudNetwork** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewPatchedCircuitTerminationRequest

`func NewPatchedCircuitTerminationRequest() *PatchedCircuitTerminationRequest`

NewPatchedCircuitTerminationRequest instantiates a new PatchedCircuitTerminationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedCircuitTerminationRequestWithDefaults

`func NewPatchedCircuitTerminationRequestWithDefaults() *PatchedCircuitTerminationRequest`

NewPatchedCircuitTerminationRequestWithDefaults instantiates a new PatchedCircuitTerminationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTermSide

`func (o *PatchedCircuitTerminationRequest) GetTermSide() TermSideEnum`

GetTermSide returns the TermSide field if non-nil, zero value otherwise.

### GetTermSideOk

`func (o *PatchedCircuitTerminationRequest) GetTermSideOk() (*TermSideEnum, bool)`

GetTermSideOk returns a tuple with the TermSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermSide

`func (o *PatchedCircuitTerminationRequest) SetTermSide(v TermSideEnum)`

SetTermSide sets TermSide field to given value.

### HasTermSide

`func (o *PatchedCircuitTerminationRequest) HasTermSide() bool`

HasTermSide returns a boolean if a field has been set.

### GetPortSpeed

`func (o *PatchedCircuitTerminationRequest) GetPortSpeed() int32`

GetPortSpeed returns the PortSpeed field if non-nil, zero value otherwise.

### GetPortSpeedOk

`func (o *PatchedCircuitTerminationRequest) GetPortSpeedOk() (*int32, bool)`

GetPortSpeedOk returns a tuple with the PortSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortSpeed

`func (o *PatchedCircuitTerminationRequest) SetPortSpeed(v int32)`

SetPortSpeed sets PortSpeed field to given value.

### HasPortSpeed

`func (o *PatchedCircuitTerminationRequest) HasPortSpeed() bool`

HasPortSpeed returns a boolean if a field has been set.

### SetPortSpeedNil

`func (o *PatchedCircuitTerminationRequest) SetPortSpeedNil(b bool)`

 SetPortSpeedNil sets the value for PortSpeed to be an explicit nil

### UnsetPortSpeed
`func (o *PatchedCircuitTerminationRequest) UnsetPortSpeed()`

UnsetPortSpeed ensures that no value is present for PortSpeed, not even an explicit nil
### GetUpstreamSpeed

`func (o *PatchedCircuitTerminationRequest) GetUpstreamSpeed() int32`

GetUpstreamSpeed returns the UpstreamSpeed field if non-nil, zero value otherwise.

### GetUpstreamSpeedOk

`func (o *PatchedCircuitTerminationRequest) GetUpstreamSpeedOk() (*int32, bool)`

GetUpstreamSpeedOk returns a tuple with the UpstreamSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamSpeed

`func (o *PatchedCircuitTerminationRequest) SetUpstreamSpeed(v int32)`

SetUpstreamSpeed sets UpstreamSpeed field to given value.

### HasUpstreamSpeed

`func (o *PatchedCircuitTerminationRequest) HasUpstreamSpeed() bool`

HasUpstreamSpeed returns a boolean if a field has been set.

### SetUpstreamSpeedNil

`func (o *PatchedCircuitTerminationRequest) SetUpstreamSpeedNil(b bool)`

 SetUpstreamSpeedNil sets the value for UpstreamSpeed to be an explicit nil

### UnsetUpstreamSpeed
`func (o *PatchedCircuitTerminationRequest) UnsetUpstreamSpeed()`

UnsetUpstreamSpeed ensures that no value is present for UpstreamSpeed, not even an explicit nil
### GetXconnectId

`func (o *PatchedCircuitTerminationRequest) GetXconnectId() string`

GetXconnectId returns the XconnectId field if non-nil, zero value otherwise.

### GetXconnectIdOk

`func (o *PatchedCircuitTerminationRequest) GetXconnectIdOk() (*string, bool)`

GetXconnectIdOk returns a tuple with the XconnectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXconnectId

`func (o *PatchedCircuitTerminationRequest) SetXconnectId(v string)`

SetXconnectId sets XconnectId field to given value.

### HasXconnectId

`func (o *PatchedCircuitTerminationRequest) HasXconnectId() bool`

HasXconnectId returns a boolean if a field has been set.

### GetPpInfo

`func (o *PatchedCircuitTerminationRequest) GetPpInfo() string`

GetPpInfo returns the PpInfo field if non-nil, zero value otherwise.

### GetPpInfoOk

`func (o *PatchedCircuitTerminationRequest) GetPpInfoOk() (*string, bool)`

GetPpInfoOk returns a tuple with the PpInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpInfo

`func (o *PatchedCircuitTerminationRequest) SetPpInfo(v string)`

SetPpInfo sets PpInfo field to given value.

### HasPpInfo

`func (o *PatchedCircuitTerminationRequest) HasPpInfo() bool`

HasPpInfo returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedCircuitTerminationRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedCircuitTerminationRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedCircuitTerminationRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedCircuitTerminationRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCircuit

`func (o *PatchedCircuitTerminationRequest) GetCircuit() BulkWritableCableRequestStatus`

GetCircuit returns the Circuit field if non-nil, zero value otherwise.

### GetCircuitOk

`func (o *PatchedCircuitTerminationRequest) GetCircuitOk() (*BulkWritableCableRequestStatus, bool)`

GetCircuitOk returns a tuple with the Circuit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuit

`func (o *PatchedCircuitTerminationRequest) SetCircuit(v BulkWritableCableRequestStatus)`

SetCircuit sets Circuit field to given value.

### HasCircuit

`func (o *PatchedCircuitTerminationRequest) HasCircuit() bool`

HasCircuit returns a boolean if a field has been set.

### GetLocation

`func (o *PatchedCircuitTerminationRequest) GetLocation() BulkWritableCircuitRequestTenant`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PatchedCircuitTerminationRequest) GetLocationOk() (*BulkWritableCircuitRequestTenant, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PatchedCircuitTerminationRequest) SetLocation(v BulkWritableCircuitRequestTenant)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PatchedCircuitTerminationRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *PatchedCircuitTerminationRequest) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *PatchedCircuitTerminationRequest) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetProviderNetwork

`func (o *PatchedCircuitTerminationRequest) GetProviderNetwork() BulkWritableCircuitRequestTenant`

GetProviderNetwork returns the ProviderNetwork field if non-nil, zero value otherwise.

### GetProviderNetworkOk

`func (o *PatchedCircuitTerminationRequest) GetProviderNetworkOk() (*BulkWritableCircuitRequestTenant, bool)`

GetProviderNetworkOk returns a tuple with the ProviderNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderNetwork

`func (o *PatchedCircuitTerminationRequest) SetProviderNetwork(v BulkWritableCircuitRequestTenant)`

SetProviderNetwork sets ProviderNetwork field to given value.

### HasProviderNetwork

`func (o *PatchedCircuitTerminationRequest) HasProviderNetwork() bool`

HasProviderNetwork returns a boolean if a field has been set.

### SetProviderNetworkNil

`func (o *PatchedCircuitTerminationRequest) SetProviderNetworkNil(b bool)`

 SetProviderNetworkNil sets the value for ProviderNetwork to be an explicit nil

### UnsetProviderNetwork
`func (o *PatchedCircuitTerminationRequest) UnsetProviderNetwork()`

UnsetProviderNetwork ensures that no value is present for ProviderNetwork, not even an explicit nil
### GetCloudNetwork

`func (o *PatchedCircuitTerminationRequest) GetCloudNetwork() BulkWritableCircuitRequestTenant`

GetCloudNetwork returns the CloudNetwork field if non-nil, zero value otherwise.

### GetCloudNetworkOk

`func (o *PatchedCircuitTerminationRequest) GetCloudNetworkOk() (*BulkWritableCircuitRequestTenant, bool)`

GetCloudNetworkOk returns a tuple with the CloudNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudNetwork

`func (o *PatchedCircuitTerminationRequest) SetCloudNetwork(v BulkWritableCircuitRequestTenant)`

SetCloudNetwork sets CloudNetwork field to given value.

### HasCloudNetwork

`func (o *PatchedCircuitTerminationRequest) HasCloudNetwork() bool`

HasCloudNetwork returns a boolean if a field has been set.

### SetCloudNetworkNil

`func (o *PatchedCircuitTerminationRequest) SetCloudNetworkNil(b bool)`

 SetCloudNetworkNil sets the value for CloudNetwork to be an explicit nil

### UnsetCloudNetwork
`func (o *PatchedCircuitTerminationRequest) UnsetCloudNetwork()`

UnsetCloudNetwork ensures that no value is present for CloudNetwork, not even an explicit nil
### GetCustomFields

`func (o *PatchedCircuitTerminationRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedCircuitTerminationRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedCircuitTerminationRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedCircuitTerminationRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedCircuitTerminationRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedCircuitTerminationRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedCircuitTerminationRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedCircuitTerminationRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

