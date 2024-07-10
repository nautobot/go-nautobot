# VmInterface

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [default to null]
**ObjectType** | **string** |  | [default to null]
**Display** | **string** | Human friendly display value | [default to null]
**Url** | **string** |  | [default to null]
**NaturalSlug** | **string** |  | [default to null]
**Mode** | [***InterfaceMode**](Interface_mode.md) |  | [optional] [default to null]
**MacAddress** | **string** |  | [optional] [default to null]
**Enabled** | **bool** |  | [optional] [default to null]
**Mtu** | **int32** |  | [optional] [default to null]
**Name** | **string** |  | [default to null]
**Description** | **string** |  | [optional] [default to null]
**Status** | [***BulkWritableCableRequestStatus**](BulkWritableCableRequest_status.md) |  | [default to null]
**ParentInterface** | [***BulkWritableInterfaceRequestParentInterface**](BulkWritableInterfaceRequest_parent_interface.md) |  | [optional] [default to null]
**Bridge** | [***BridgeInterface**](Bridge interface.md) |  | [optional] [default to null]
**VirtualMachine** | [***BulkWritableCableRequestStatus**](BulkWritableCableRequest_status.md) |  | [default to null]
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

