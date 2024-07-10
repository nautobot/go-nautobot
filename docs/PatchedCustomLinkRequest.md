# PatchedCustomLinkRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentType** | **string** |  | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**Text** | **string** | Jinja2 template code for link text. Reference the object as &lt;code&gt;{{ obj }}&lt;/code&gt; such as &lt;code&gt;{{ obj.platform.name }}&lt;/code&gt;. Links which render as empty text will not be displayed. | [optional] [default to null]
**TargetUrl** | **string** | Jinja2 template code for link URL. Reference the object as &lt;code&gt;{{ obj }}&lt;/code&gt; such as &lt;code&gt;{{ obj.platform.name }}&lt;/code&gt;. | [optional] [default to null]
**Weight** | **int32** |  | [optional] [default to null]
**GroupName** | **string** | Links with the same group will appear as a dropdown menu | [optional] [default to null]
**ButtonClass** | [***AllOfPatchedCustomLinkRequestButtonClass**](AllOfPatchedCustomLinkRequestButtonClass.md) | The class of the first link in a group will be used for the dropdown button | [optional] [default to null]
**NewWindow** | **bool** | Force link to open in a new window | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

