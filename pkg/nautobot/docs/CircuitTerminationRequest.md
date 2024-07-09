# CircuitTerminationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TermSide** | [**TermSideEnum**](TermSideEnum.md) |  | 
**PortSpeed** | Pointer to **NullableInt32** |  | [optional] 
**UpstreamSpeed** | Pointer to **NullableInt32** | Upstream speed, if different from port speed | [optional] 
**XconnectId** | Pointer to **string** |  | [optional] 
**PpInfo** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Circuit** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Location** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**ProviderNetwork** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewCircuitTerminationRequest

`func NewCircuitTerminationRequest(termSide TermSideEnum, circuit BulkWritableCableRequestStatus, ) *CircuitTerminationRequest`

NewCircuitTerminationRequest instantiates a new CircuitTerminationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCircuitTerminationRequestWithDefaults

`func NewCircuitTerminationRequestWithDefaults() *CircuitTerminationRequest`

NewCircuitTerminationRequestWithDefaults instantiates a new CircuitTerminationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTermSide

`func (o *CircuitTerminationRequest) GetTermSide() TermSideEnum`

GetTermSide returns the TermSide field if non-nil, zero value otherwise.

### GetTermSideOk

`func (o *CircuitTerminationRequest) GetTermSideOk() (*TermSideEnum, bool)`

GetTermSideOk returns a tuple with the TermSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermSide

`func (o *CircuitTerminationRequest) SetTermSide(v TermSideEnum)`

SetTermSide sets TermSide field to given value.


### GetPortSpeed

`func (o *CircuitTerminationRequest) GetPortSpeed() int32`

GetPortSpeed returns the PortSpeed field if non-nil, zero value otherwise.

### GetPortSpeedOk

`func (o *CircuitTerminationRequest) GetPortSpeedOk() (*int32, bool)`

GetPortSpeedOk returns a tuple with the PortSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortSpeed

`func (o *CircuitTerminationRequest) SetPortSpeed(v int32)`

SetPortSpeed sets PortSpeed field to given value.

### HasPortSpeed

`func (o *CircuitTerminationRequest) HasPortSpeed() bool`

HasPortSpeed returns a boolean if a field has been set.

### SetPortSpeedNil

`func (o *CircuitTerminationRequest) SetPortSpeedNil(b bool)`

 SetPortSpeedNil sets the value for PortSpeed to be an explicit nil

### UnsetPortSpeed
`func (o *CircuitTerminationRequest) UnsetPortSpeed()`

UnsetPortSpeed ensures that no value is present for PortSpeed, not even an explicit nil
### GetUpstreamSpeed

`func (o *CircuitTerminationRequest) GetUpstreamSpeed() int32`

GetUpstreamSpeed returns the UpstreamSpeed field if non-nil, zero value otherwise.

### GetUpstreamSpeedOk

`func (o *CircuitTerminationRequest) GetUpstreamSpeedOk() (*int32, bool)`

GetUpstreamSpeedOk returns a tuple with the UpstreamSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamSpeed

`func (o *CircuitTerminationRequest) SetUpstreamSpeed(v int32)`

SetUpstreamSpeed sets UpstreamSpeed field to given value.

### HasUpstreamSpeed

`func (o *CircuitTerminationRequest) HasUpstreamSpeed() bool`

HasUpstreamSpeed returns a boolean if a field has been set.

### SetUpstreamSpeedNil

`func (o *CircuitTerminationRequest) SetUpstreamSpeedNil(b bool)`

 SetUpstreamSpeedNil sets the value for UpstreamSpeed to be an explicit nil

### UnsetUpstreamSpeed
`func (o *CircuitTerminationRequest) UnsetUpstreamSpeed()`

UnsetUpstreamSpeed ensures that no value is present for UpstreamSpeed, not even an explicit nil
### GetXconnectId

`func (o *CircuitTerminationRequest) GetXconnectId() string`

GetXconnectId returns the XconnectId field if non-nil, zero value otherwise.

### GetXconnectIdOk

`func (o *CircuitTerminationRequest) GetXconnectIdOk() (*string, bool)`

GetXconnectIdOk returns a tuple with the XconnectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXconnectId

`func (o *CircuitTerminationRequest) SetXconnectId(v string)`

SetXconnectId sets XconnectId field to given value.

### HasXconnectId

`func (o *CircuitTerminationRequest) HasXconnectId() bool`

HasXconnectId returns a boolean if a field has been set.

### GetPpInfo

`func (o *CircuitTerminationRequest) GetPpInfo() string`

GetPpInfo returns the PpInfo field if non-nil, zero value otherwise.

### GetPpInfoOk

`func (o *CircuitTerminationRequest) GetPpInfoOk() (*string, bool)`

GetPpInfoOk returns a tuple with the PpInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpInfo

`func (o *CircuitTerminationRequest) SetPpInfo(v string)`

SetPpInfo sets PpInfo field to given value.

### HasPpInfo

`func (o *CircuitTerminationRequest) HasPpInfo() bool`

HasPpInfo returns a boolean if a field has been set.

### GetDescription

`func (o *CircuitTerminationRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CircuitTerminationRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CircuitTerminationRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CircuitTerminationRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCircuit

`func (o *CircuitTerminationRequest) GetCircuit() BulkWritableCableRequestStatus`

GetCircuit returns the Circuit field if non-nil, zero value otherwise.

### GetCircuitOk

`func (o *CircuitTerminationRequest) GetCircuitOk() (*BulkWritableCableRequestStatus, bool)`

GetCircuitOk returns a tuple with the Circuit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuit

`func (o *CircuitTerminationRequest) SetCircuit(v BulkWritableCableRequestStatus)`

SetCircuit sets Circuit field to given value.


### GetLocation

`func (o *CircuitTerminationRequest) GetLocation() BulkWritableCircuitRequestTenant`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CircuitTerminationRequest) GetLocationOk() (*BulkWritableCircuitRequestTenant, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *CircuitTerminationRequest) SetLocation(v BulkWritableCircuitRequestTenant)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *CircuitTerminationRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *CircuitTerminationRequest) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *CircuitTerminationRequest) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetProviderNetwork

`func (o *CircuitTerminationRequest) GetProviderNetwork() BulkWritableCircuitRequestTenant`

GetProviderNetwork returns the ProviderNetwork field if non-nil, zero value otherwise.

### GetProviderNetworkOk

`func (o *CircuitTerminationRequest) GetProviderNetworkOk() (*BulkWritableCircuitRequestTenant, bool)`

GetProviderNetworkOk returns a tuple with the ProviderNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderNetwork

`func (o *CircuitTerminationRequest) SetProviderNetwork(v BulkWritableCircuitRequestTenant)`

SetProviderNetwork sets ProviderNetwork field to given value.

### HasProviderNetwork

`func (o *CircuitTerminationRequest) HasProviderNetwork() bool`

HasProviderNetwork returns a boolean if a field has been set.

### SetProviderNetworkNil

`func (o *CircuitTerminationRequest) SetProviderNetworkNil(b bool)`

 SetProviderNetworkNil sets the value for ProviderNetwork to be an explicit nil

### UnsetProviderNetwork
`func (o *CircuitTerminationRequest) UnsetProviderNetwork()`

UnsetProviderNetwork ensures that no value is present for ProviderNetwork, not even an explicit nil
### GetCustomFields

`func (o *CircuitTerminationRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CircuitTerminationRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CircuitTerminationRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CircuitTerminationRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *CircuitTerminationRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CircuitTerminationRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CircuitTerminationRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *CircuitTerminationRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


