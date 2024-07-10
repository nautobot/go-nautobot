# PatchedSoftwareImageFileRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageFileName** | **string** |  | [optional] [default to null]
**ImageFileChecksum** | **string** |  | [optional] [default to null]
**HashingAlgorithm** | [***OneOfPatchedSoftwareImageFileRequestHashingAlgorithm**](OneOfPatchedSoftwareImageFileRequestHashingAlgorithm.md) | Hashing algorithm for image file checksum | [optional] [default to null]
**ImageFileSize** | **int64** | Image file size in bytes | [optional] [default to null]
**DownloadUrl** | **string** |  | [optional] [default to null]
**DefaultImage** | **bool** | Is the default image for this software version | [optional] [default to null]
**SoftwareVersion** | [***BulkWritableCableRequestStatus**](BulkWritableCableRequest_status.md) |  | [optional] [default to null]
**Status** | [***BulkWritableCableRequestStatus**](BulkWritableCableRequest_status.md) |  | [optional] [default to null]
**Tags** | [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequest_status.md) |  | [optional] [default to null]
**CustomFields** | [**map[string]Object**](.md) |  | [optional] [default to null]
**Relationships** | [**map[string]BulkWritableCableRequestRelationships**](BulkWritableCableRequest_relationships.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

