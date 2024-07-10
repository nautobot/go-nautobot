# BulkWritableDeviceTypeRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [default to null]
**SubdeviceRole** | [***SubdeviceRoleEnum**](SubdeviceRoleEnum.md) |  | [optional] [default to null]
**FrontImage** | [****os.File**](*os.File.md) |  | [optional] [default to null]
**RearImage** | [****os.File**](*os.File.md) |  | [optional] [default to null]
**Model** | **string** |  | [default to null]
**PartNumber** | **string** | Discrete part number (optional) | [optional] [default to null]
**UHeight** | **int32** |  | [optional] [default to null]
**IsFullDepth** | **bool** | Device consumes both front and rear rack faces | [optional] [default to null]
**Comments** | **string** |  | [optional] [default to null]
**Manufacturer** | [***BulkWritableCableRequestStatus**](BulkWritableCableRequest_status.md) |  | [default to null]
**DeviceFamily** | [***BulkWritableCircuitRequestTenant**](BulkWritableCircuitRequest_tenant.md) |  | [optional] [default to null]
**Tags** | [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequest_status.md) |  | [optional] [default to null]
**CustomFields** | [**map[string]Object**](.md) |  | [optional] [default to null]
**Relationships** | [**map[string]BulkWritableCableRequestRelationships**](BulkWritableCableRequest_relationships.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

