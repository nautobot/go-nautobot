# ModelInterface

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [default to null]
**ObjectType** | **string** |  | [default to null]
**Display** | **string** | Human friendly display value | [default to null]
**Url** | **string** |  | [default to null]
**NaturalSlug** | **string** |  | [default to null]
**CablePeerType** | **string** |  | [default to null]
**CablePeer** | [***AllOfInterfaceCablePeer**](AllOfInterfaceCablePeer.md) |  | [default to null]
**ConnectedEndpointType** | **string** |  | [default to null]
**ConnectedEndpoint** | [***AllOfInterfaceConnectedEndpoint**](AllOfInterfaceConnectedEndpoint.md) |  | [default to null]
**ConnectedEndpointReachable** | **bool** |  | [default to null]
**Type_** | [***InterfaceType**](Interface_type.md) |  | [default to null]
**Mode** | [***InterfaceMode**](Interface_mode.md) |  | [optional] [default to null]
**MacAddress** | **string** |  | [optional] [default to null]
**IpAddressCount** | **int32** |  | [default to null]
**Name** | **string** |  | [default to null]
**Label** | **string** | Physical label | [optional] [default to null]
**Description** | **string** |  | [optional] [default to null]
**Enabled** | **bool** |  | [optional] [default to null]
**Mtu** | **int32** |  | [optional] [default to null]
**MgmtOnly** | **bool** | This interface is used only for out-of-band management | [optional] [default to null]
**Device** | [***BulkWritableCableRequestStatus**](BulkWritableCableRequest_status.md) |  | [default to null]
**Cable** | [***CircuitCircuitTerminationA**](Circuit_circuit_termination_a.md) |  | [default to null]
**Status** | [***BulkWritableCableRequestStatus**](BulkWritableCableRequest_status.md) |  | [default to null]
**ParentInterface** | [***BulkWritableInterfaceRequestParentInterface**](BulkWritableInterfaceRequest_parent_interface.md) |  | [optional] [default to null]
**Bridge** | [***BridgeInterface**](Bridge interface.md) |  | [optional] [default to null]
**Lag** | [***ParentLag**](Parent LAG.md) |  | [optional] [default to null]
**UntaggedVlan** | [***BulkWritableCircuitRequestTenant**](BulkWritableCircuitRequest_tenant.md) |  | [optional] [default to null]
**Vrf** | [***BulkWritableCircuitRequestTenant**](BulkWritableCircuitRequest_tenant.md) |  | [optional] [default to null]
**TaggedVlans** | [**[]TaggedVlans**](Tagged VLANs.md) |  | [optional] [default to null]
**IpAddresses** | [**[]IpAddresses**](IP Addresses.md) |  | [default to null]
**Created** | [**time.Time**](time.Time.md) |  | [default to null]
**LastUpdated** | [**time.Time**](time.Time.md) |  | [default to null]
**Tags** | [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequest_status.md) |  | [optional] [default to null]
**NotesUrl** | **string** |  | [default to null]
**CustomFields** | [**map[string]Object**](.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

