# PatchedComputedFieldRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentType** | **string** |  | [optional] [default to null]
**Key** | **string** | Internal field name. Please use underscores rather than dashes in this key. | [optional] [default to null]
**Label** | **string** | Name of the field as displayed to users | [optional] [default to null]
**Description** | **string** |  | [optional] [default to null]
**Template** | **string** | Jinja2 template code for field value | [optional] [default to null]
**FallbackValue** | **string** | Fallback value (if any) to be output for the field in the case of a template rendering error. | [optional] [default to null]
**Weight** | **int32** |  | [optional] [default to null]
**AdvancedUi** | **bool** | Hide this field from the object&#x27;s primary information tab. It will appear in the \&quot;Advanced\&quot; tab instead. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

