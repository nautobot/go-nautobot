# PatchedWritableCustomFieldRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentTypes** | **[]string** |  | [optional] [default to null]
**Label** | **string** |  | [optional] [default to null]
**Grouping** | **string** | Human-readable grouping that this custom field belongs to. | [optional] [default to null]
**Type_** | [***AllOfPatchedWritableCustomFieldRequestType_**](AllOfPatchedWritableCustomFieldRequestType_.md) | The type of value(s) allowed for this field. | [optional] [default to null]
**Key** | **string** | Internal field name. Please use underscores rather than dashes in this key. | [optional] [default to null]
**Description** | **string** | A helpful description for this field. | [optional] [default to null]
**Required** | **bool** | If true, this field is required when creating new objects or editing an existing object. | [optional] [default to null]
**FilterLogic** | [***AllOfPatchedWritableCustomFieldRequestFilterLogic**](AllOfPatchedWritableCustomFieldRequestFilterLogic.md) | Loose matches any instance of a given string; Exact matches the entire field. | [optional] [default to null]
**Default_** | [**map[string]Object**](.md) | Default value for the field (must be a JSON value). Encapsulate strings with double quotes (e.g. \&quot;Foo\&quot;). | [optional] [default to null]
**Weight** | **int32** | Fields with higher weights appear lower in a form. | [optional] [default to null]
**ValidationMinimum** | **int64** | Minimum allowed value (for numeric fields) or length (for text fields). | [optional] [default to null]
**ValidationMaximum** | **int64** | Maximum allowed value (for numeric fields) or length (for text fields). | [optional] [default to null]
**ValidationRegex** | **string** | Regular expression to enforce on text field values. Use ^ and $ to force matching of entire string. For example, &lt;code&gt;^[A-Z]{3}$&lt;/code&gt; will limit values to exactly three uppercase letters. Regular expression on select and multi-select will be applied at &lt;code&gt;Custom Field Choices&lt;/code&gt; definition. | [optional] [default to null]
**AdvancedUi** | **bool** | Hide this field from the object&#x27;s primary information tab. It will appear in the \&quot;Advanced\&quot; tab instead. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

