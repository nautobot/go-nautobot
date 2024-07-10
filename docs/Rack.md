# Rack

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [default to null]
**ObjectType** | **string** |  | [default to null]
**Display** | **string** | Human friendly display value | [default to null]
**Url** | **string** |  | [default to null]
**NaturalSlug** | **string** |  | [default to null]
**Type_** | [***RackType**](Rack_type.md) |  | [optional] [default to null]
**Width** | [***RackWidth**](Rack_width.md) |  | [optional] [default to null]
**OuterUnit** | [***RackOuterUnit**](Rack_outer_unit.md) |  | [optional] [default to null]
**DeviceCount** | **int32** |  | [default to null]
**PowerFeedCount** | **int32** |  | [default to null]
**Name** | **string** |  | [default to null]
**FacilityId** | **string** | Locally-assigned identifier | [optional] [default to null]
**Serial** | **string** |  | [optional] [default to null]
**AssetTag** | **string** | A unique tag used to identify this rack | [optional] [default to null]
**UHeight** | **int32** | Height in rack units | [optional] [default to null]
**DescUnits** | **bool** | Units are numbered top-to-bottom | [optional] [default to null]
**OuterWidth** | **int32** | Outer dimension of rack (width) | [optional] [default to null]
**OuterDepth** | **int32** | Outer dimension of rack (depth) | [optional] [default to null]
**Comments** | **string** |  | [optional] [default to null]
**Status** | [***BulkWritableCableRequestStatus**](BulkWritableCableRequest_status.md) |  | [default to null]
**Role** | [***BulkWritableCircuitRequestTenant**](BulkWritableCircuitRequest_tenant.md) |  | [optional] [default to null]
**Location** | [***BulkWritableCableRequestStatus**](BulkWritableCableRequest_status.md) |  | [default to null]
**RackGroup** | [***BulkWritableRackRequestRackGroup**](BulkWritableRackRequest_rack_group.md) |  | [optional] [default to null]
**Tenant** | [***BulkWritableCircuitRequestTenant**](BulkWritableCircuitRequest_tenant.md) |  | [optional] [default to null]
**Created** | [**time.Time**](time.Time.md) |  | [default to null]
**LastUpdated** | [**time.Time**](time.Time.md) |  | [default to null]
**Tags** | [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequest_status.md) |  | [optional] [default to null]
**NotesUrl** | **string** |  | [default to null]
**CustomFields** | [**map[string]Object**](.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

