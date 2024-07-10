# JobButtonRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentTypes** | **[]string** |  | [default to null]
**Name** | **string** |  | [default to null]
**Text** | **string** | Jinja2 template code for button text. Reference the object as &lt;code&gt;{{ obj }}&lt;/code&gt; such as &lt;code&gt;{{ obj.platform.name }}&lt;/code&gt;. Buttons which render as empty text will not be displayed. | [default to null]
**Weight** | **int32** |  | [optional] [default to null]
**GroupName** | **string** | Buttons with the same group will appear as a dropdown menu. Group dropdown buttons will inherit the button class from the button with the lowest weight in the group. | [optional] [default to null]
**ButtonClass** | [***ButtonClassEnum**](ButtonClassEnum.md) |  | [optional] [default to null]
**Confirmation** | **bool** | Enable confirmation pop-up box. &lt;span class&#x3D;&#x27;text-danger&#x27;&gt;WARNING: unselecting this option will allow the Job to run (and commit changes) with a single click!&lt;/span&gt; | [optional] [default to null]
**Job** | [***BulkWritableJobButtonRequestJob**](BulkWritableJobButtonRequest_job.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

