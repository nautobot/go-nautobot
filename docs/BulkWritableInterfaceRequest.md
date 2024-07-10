# BulkWritableInterfaceRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [default to null]
**Type_** | [***InterfaceTypeChoices**](InterfaceTypeChoices.md) |  | [default to null]
**Mode** | [***ModeEnum**](ModeEnum.md) |  | [optional] [default to null]
**MacAddress** | **string** |  | [optional] [default to null]
**Name** | **string** |  | [default to null]
**Label** | **string** | Physical label | [optional] [default to null]
**Description** | **string** |  | [optional] [default to null]
**Enabled** | **bool** |  | [optional] [default to null]
**Mtu** | **int32** |  | [optional] [default to null]
**MgmtOnly** | **bool** | This interface is used only for out-of-band management | [optional] [default to null]
**Device** | [***BulkWritableCableRequestStatus**](BulkWritableCableRequest_status.md) |  | [default to null]
**Status** | [***BulkWritableCableRequestStatus**](BulkWritableCableRequest_status.md) |  | [default to null]
**ParentInterface** | [***BulkWritableInterfaceRequestParentInterface**](BulkWritableInterfaceRequest_parent_interface.md) |  | [optional] [default to null]
**Bridge** | [***BridgeInterface**](Bridge interface.md) |  | [optional] [default to null]
**Lag** | [***ParentLag**](Parent LAG.md) |  | [optional] [default to null]
**UntaggedVlan** | [***BulkWritableCircuitRequestTenant**](BulkWritableCircuitRequest_tenant.md) |  | [optional] [default to null]
**Vrf** | [***BulkWritableCircuitRequestTenant**](BulkWritableCircuitRequest_tenant.md) |  | [optional] [default to null]
**TaggedVlans** | [**[]TaggedVlans**](Tagged VLANs.md) |  | [optional] [default to null]
**Tags** | [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequest_status.md) |  | [optional] [default to null]
**CustomFields** | [**map[string]Object**](.md) |  | [optional] [default to null]
**Relationships** | [**map[string]BulkWritableCableRequestRelationships**](BulkWritableCableRequest_relationships.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

