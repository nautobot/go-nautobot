# ObjectPermission

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [default to null]
**ObjectType** | **string** |  | [default to null]
**Display** | **string** | Human friendly display value | [default to null]
**Url** | **string** |  | [default to null]
**NaturalSlug** | **string** |  | [default to null]
**ObjectTypes** | **[]string** |  | [default to null]
**Name** | **string** |  | [default to null]
**Description** | **string** |  | [optional] [default to null]
**Enabled** | **bool** |  | [optional] [default to null]
**Actions** | [**map[string]Object**](.md) | The list of actions granted by this permission | [default to null]
**Constraints** | [**map[string]Object**](.md) | Queryset filter matching the applicable objects of the selected type(s) | [optional] [default to null]
**Groups** | [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequest_status.md) |  | [optional] [default to null]
**Users** | [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequest_status.md) |  | [optional] [default to null]
**Created** | [**time.Time**](time.Time.md) |  | [default to null]
**LastUpdated** | [**time.Time**](time.Time.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

