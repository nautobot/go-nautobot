# SoftwareImageFile

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [default to null]
**ObjectType** | **string** |  | [default to null]
**Display** | **string** | Human friendly display value | [default to null]
**Url** | **string** |  | [default to null]
**NaturalSlug** | **string** |  | [default to null]
**ImageFileName** | **string** |  | [default to null]
**ImageFileChecksum** | **string** |  | [optional] [default to null]
**HashingAlgorithm** | [***OneOfSoftwareImageFileHashingAlgorithm**](OneOfSoftwareImageFileHashingAlgorithm.md) | Hashing algorithm for image file checksum | [optional] [default to null]
**ImageFileSize** | **int64** | Image file size in bytes | [optional] [default to null]
**DownloadUrl** | **string** |  | [optional] [default to null]
**DefaultImage** | **bool** | Is the default image for this software version | [optional] [default to null]
**SoftwareVersion** | [***BulkWritableCableRequestStatus**](BulkWritableCableRequest_status.md) |  | [default to null]
**Status** | [***BulkWritableCableRequestStatus**](BulkWritableCableRequest_status.md) |  | [default to null]
**Created** | [**time.Time**](time.Time.md) |  | [default to null]
**LastUpdated** | [**time.Time**](time.Time.md) |  | [default to null]
**Tags** | [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequest_status.md) |  | [optional] [default to null]
**NotesUrl** | **string** |  | [default to null]
**CustomFields** | [**map[string]Object**](.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

