# WebhookRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentTypes** | **[]string** |  | [default to null]
**Name** | **string** |  | [default to null]
**TypeCreate** | **bool** | Call this webhook when a matching object is created. | [optional] [default to null]
**TypeUpdate** | **bool** | Call this webhook when a matching object is updated. | [optional] [default to null]
**TypeDelete** | **bool** | Call this webhook when a matching object is deleted. | [optional] [default to null]
**PayloadUrl** | **string** | A POST will be sent to this URL when the webhook is called. | [default to null]
**Enabled** | **bool** |  | [optional] [default to null]
**HttpMethod** | [***HttpMethodEnum**](HttpMethodEnum.md) |  | [optional] [default to null]
**HttpContentType** | **string** | The complete list of official content types is available &lt;a href&#x3D;\&quot;https://www.iana.org/assignments/media-types/media-types.xhtml\&quot;&gt;here&lt;/a&gt;. | [optional] [default to null]
**AdditionalHeaders** | **string** | User-supplied HTTP headers to be sent with the request in addition to the HTTP content type. Headers should be defined in the format &lt;code&gt;Name: Value&lt;/code&gt;. Jinja2 template processing is supported with the same context as the request body (below). | [optional] [default to null]
**BodyTemplate** | **string** | Jinja2 template for a custom request body. If blank, a JSON object representing the change will be included. Available context data includes: &lt;code&gt;event&lt;/code&gt;, &lt;code&gt;model&lt;/code&gt;, &lt;code&gt;timestamp&lt;/code&gt;, &lt;code&gt;username&lt;/code&gt;, &lt;code&gt;request_id&lt;/code&gt;, and &lt;code&gt;data&lt;/code&gt;. | [optional] [default to null]
**Secret** | **string** | When provided, the request will include a &#x27;X-Hook-Signature&#x27; header containing a HMAC hex digest of the payload body using the secret as the key. The secret is not transmitted in the request. | [optional] [default to null]
**SslVerification** | **bool** | Enable SSL certificate verification. Disable with caution! | [optional] [default to null]
**CaFilePath** | **string** | The specific CA certificate file to use for SSL verification. Leave blank to use the system defaults. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

