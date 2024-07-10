# ObjectPermissionRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectTypes** | **[]string** |  | [default to null]
**Name** | **string** |  | [default to null]
**Description** | **string** |  | [optional] [default to null]
**Enabled** | **bool** |  | [optional] [default to null]
**Actions** | [**map[string]Object**](.md) | The list of actions granted by this permission | [default to null]
**Constraints** | [**map[string]Object**](.md) | Queryset filter matching the applicable objects of the selected type(s) | [optional] [default to null]
**Groups** | [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequest_status.md) |  | [optional] [default to null]
**Users** | [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequest_status.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

