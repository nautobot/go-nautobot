# BulkWritableInventoryItemRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [default to null]
**Name** | **string** |  | [default to null]
**Label** | **string** | Physical label | [optional] [default to null]
**Description** | **string** |  | [optional] [default to null]
**PartId** | **string** | Manufacturer-assigned part identifier | [optional] [default to null]
**Serial** | **string** |  | [optional] [default to null]
**AssetTag** | **string** | A unique tag used to identify this item | [optional] [default to null]
**Discovered** | **bool** | This item was automatically discovered | [optional] [default to null]
**Parent** | [***BulkWritableCircuitRequestTenant**](BulkWritableCircuitRequest_tenant.md) |  | [optional] [default to null]
**Device** | [***BulkWritableCableRequestStatus**](BulkWritableCableRequest_status.md) |  | [default to null]
**Manufacturer** | [***BulkWritableCircuitRequestTenant**](BulkWritableCircuitRequest_tenant.md) |  | [optional] [default to null]
**SoftwareVersion** | [***TheSoftwareVersionInstalledOnThisInventoryItem**](The software version installed on this inventory item.md) |  | [optional] [default to null]
**SoftwareImageFiles** | [**[]SoftwareImageFiles**](Software Image Files.md) | Override the software image files associated with the software version for this inventory item | [optional] [default to null]
**Tags** | [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequest_status.md) |  | [optional] [default to null]
**CustomFields** | [**map[string]Object**](.md) |  | [optional] [default to null]
**Relationships** | [**map[string]BulkWritableCableRequestRelationships**](BulkWritableCableRequest_relationships.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

