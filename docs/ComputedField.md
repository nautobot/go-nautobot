# ComputedField

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [default to null]
**ObjectType** | **string** |  | [default to null]
**Display** | **string** | Human friendly display value | [default to null]
**Url** | **string** |  | [default to null]
**NaturalSlug** | **string** |  | [default to null]
**ContentType** | **string** |  | [default to null]
**Key** | **string** | Internal field name. Please use underscores rather than dashes in this key. | [optional] [default to null]
**Label** | **string** | Name of the field as displayed to users | [default to null]
**Description** | **string** |  | [optional] [default to null]
**Template** | **string** | Jinja2 template code for field value | [default to null]
**FallbackValue** | **string** | Fallback value (if any) to be output for the field in the case of a template rendering error. | [optional] [default to null]
**Weight** | **int32** |  | [optional] [default to null]
**AdvancedUi** | **bool** | Hide this field from the object&#x27;s primary information tab. It will appear in the \&quot;Advanced\&quot; tab instead. | [optional] [default to null]
**Created** | [**time.Time**](time.Time.md) |  | [default to null]
**LastUpdated** | [**time.Time**](time.Time.md) |  | [default to null]
**NotesUrl** | **string** |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

