# InventoryItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [default to null]
**ObjectType** | **string** |  | [default to null]
**Display** | **string** | Human friendly display value | [default to null]
**Url** | **string** |  | [default to null]
**NaturalSlug** | **string** |  | [default to null]
**TreeDepth** | **int32** |  | [default to null]
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
**Created** | [**time.Time**](time.Time.md) |  | [default to null]
**LastUpdated** | [**time.Time**](time.Time.md) |  | [default to null]
**Tags** | [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequest_status.md) |  | [optional] [default to null]
**NotesUrl** | **string** |  | [default to null]
**CustomFields** | [**map[string]Object**](.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

