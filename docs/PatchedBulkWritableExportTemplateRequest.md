# PatchedBulkWritableExportTemplateRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [default to null]
**ContentType** | **string** |  | [optional] [default to null]
**OwnerContentType** | **string** |  | [optional] [default to null]
**OwnerObjectId** | **string** |  | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**Description** | **string** |  | [optional] [default to null]
**TemplateCode** | **string** | The list of objects being exported is passed as a context variable named &lt;code&gt;queryset&lt;/code&gt;. | [optional] [default to null]
**MimeType** | **string** | Defaults to &lt;code&gt;text/plain&lt;/code&gt; | [optional] [default to null]
**FileExtension** | **string** | Extension to append to the rendered filename | [optional] [default to null]
**Relationships** | [**map[string]BulkWritableCableRequestRelationships**](BulkWritableCableRequest_relationships.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

