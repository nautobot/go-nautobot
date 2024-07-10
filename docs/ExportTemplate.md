# ExportTemplate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [default to null]
**ObjectType** | **string** |  | [default to null]
**Display** | **string** | Human friendly display value | [default to null]
**Url** | **string** |  | [default to null]
**NaturalSlug** | **string** |  | [default to null]
**ContentType** | **string** |  | [default to null]
**OwnerContentType** | **string** |  | [optional] [default to null]
**Owner** | [***AllOfExportTemplateOwner**](AllOfExportTemplateOwner.md) |  | [default to null]
**OwnerObjectId** | **string** |  | [optional] [default to null]
**Name** | **string** |  | [default to null]
**Description** | **string** |  | [optional] [default to null]
**TemplateCode** | **string** | The list of objects being exported is passed as a context variable named &lt;code&gt;queryset&lt;/code&gt;. | [default to null]
**MimeType** | **string** | Defaults to &lt;code&gt;text/plain&lt;/code&gt; | [optional] [default to null]
**FileExtension** | **string** | Extension to append to the rendered filename | [optional] [default to null]
**Created** | [**time.Time**](time.Time.md) |  | [default to null]
**LastUpdated** | [**time.Time**](time.Time.md) |  | [default to null]
**NotesUrl** | **string** |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

