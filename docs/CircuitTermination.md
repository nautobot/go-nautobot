# CircuitTermination

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [default to null]
**ObjectType** | **string** |  | [default to null]
**Display** | **string** | Human friendly display value | [default to null]
**Url** | **string** |  | [default to null]
**NaturalSlug** | **string** |  | [default to null]
**CablePeerType** | **string** |  | [default to null]
**CablePeer** | [***AllOfCircuitTerminationCablePeer**](AllOfCircuitTerminationCablePeer.md) |  | [default to null]
**ConnectedEndpointType** | **string** |  | [default to null]
**ConnectedEndpoint** | [***AllOfCircuitTerminationConnectedEndpoint**](AllOfCircuitTerminationConnectedEndpoint.md) |  | [default to null]
**ConnectedEndpointReachable** | **bool** |  | [default to null]
**TermSide** | [***AllOfCircuitTerminationTermSide**](AllOfCircuitTerminationTermSide.md) |  | [default to null]
**PortSpeed** | **int32** |  | [optional] [default to null]
**UpstreamSpeed** | **int32** | Upstream speed, if different from port speed | [optional] [default to null]
**XconnectId** | **string** |  | [optional] [default to null]
**PpInfo** | **string** |  | [optional] [default to null]
**Description** | **string** |  | [optional] [default to null]
**Cable** | [***CircuitCircuitTerminationA**](Circuit_circuit_termination_a.md) |  | [default to null]
**Circuit** | [***BulkWritableCableRequestStatus**](BulkWritableCableRequest_status.md) |  | [default to null]
**Location** | [***BulkWritableCircuitRequestTenant**](BulkWritableCircuitRequest_tenant.md) |  | [optional] [default to null]
**ProviderNetwork** | [***BulkWritableCircuitRequestTenant**](BulkWritableCircuitRequest_tenant.md) |  | [optional] [default to null]
**Created** | [**time.Time**](time.Time.md) |  | [default to null]
**LastUpdated** | [**time.Time**](time.Time.md) |  | [default to null]
**NotesUrl** | **string** |  | [default to null]
**CustomFields** | [**map[string]Object**](.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

