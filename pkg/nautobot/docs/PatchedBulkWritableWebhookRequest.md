# PatchedBulkWritableWebhookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ContentTypes** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**TypeCreate** | Pointer to **bool** | Call this webhook when a matching object is created. | [optional] 
**TypeUpdate** | Pointer to **bool** | Call this webhook when a matching object is updated. | [optional] 
**TypeDelete** | Pointer to **bool** | Call this webhook when a matching object is deleted. | [optional] 
**PayloadUrl** | Pointer to **string** | A POST will be sent to this URL when the webhook is called. | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**HttpMethod** | Pointer to [**HttpMethodEnum**](HttpMethodEnum.md) |  | [optional] 
**HttpContentType** | Pointer to **string** | The complete list of official content types is available &lt;a href&#x3D;\&quot;https://www.iana.org/assignments/media-types/media-types.xhtml\&quot;&gt;here&lt;/a&gt;. | [optional] 
**AdditionalHeaders** | Pointer to **string** | User-supplied HTTP headers to be sent with the request in addition to the HTTP content type. Headers should be defined in the format &lt;code&gt;Name: Value&lt;/code&gt;. Jinja2 template processing is supported with the same context as the request body (below). | [optional] 
**BodyTemplate** | Pointer to **string** | Jinja2 template for a custom request body. If blank, a JSON object representing the change will be included. Available context data includes: &lt;code&gt;event&lt;/code&gt;, &lt;code&gt;model&lt;/code&gt;, &lt;code&gt;timestamp&lt;/code&gt;, &lt;code&gt;username&lt;/code&gt;, &lt;code&gt;request_id&lt;/code&gt;, and &lt;code&gt;data&lt;/code&gt;. | [optional] 
**Secret** | Pointer to **string** | When provided, the request will include a &#39;X-Hook-Signature&#39; header containing a HMAC hex digest of the payload body using the secret as the key. The secret is not transmitted in the request. | [optional] 
**SslVerification** | Pointer to **bool** | Enable SSL certificate verification. Disable with caution! | [optional] 
**CaFilePath** | Pointer to **string** | The specific CA certificate file to use for SSL verification. Leave blank to use the system defaults. | [optional] 

## Methods

### NewPatchedBulkWritableWebhookRequest

`func NewPatchedBulkWritableWebhookRequest(id string, ) *PatchedBulkWritableWebhookRequest`

NewPatchedBulkWritableWebhookRequest instantiates a new PatchedBulkWritableWebhookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableWebhookRequestWithDefaults

`func NewPatchedBulkWritableWebhookRequestWithDefaults() *PatchedBulkWritableWebhookRequest`

NewPatchedBulkWritableWebhookRequestWithDefaults instantiates a new PatchedBulkWritableWebhookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableWebhookRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableWebhookRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableWebhookRequest) SetId(v string)`

SetId sets Id field to given value.


### GetContentTypes

`func (o *PatchedBulkWritableWebhookRequest) GetContentTypes() []string`

GetContentTypes returns the ContentTypes field if non-nil, zero value otherwise.

### GetContentTypesOk

`func (o *PatchedBulkWritableWebhookRequest) GetContentTypesOk() (*[]string, bool)`

GetContentTypesOk returns a tuple with the ContentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentTypes

`func (o *PatchedBulkWritableWebhookRequest) SetContentTypes(v []string)`

SetContentTypes sets ContentTypes field to given value.

### HasContentTypes

`func (o *PatchedBulkWritableWebhookRequest) HasContentTypes() bool`

HasContentTypes returns a boolean if a field has been set.

### GetName

`func (o *PatchedBulkWritableWebhookRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkWritableWebhookRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkWritableWebhookRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkWritableWebhookRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTypeCreate

`func (o *PatchedBulkWritableWebhookRequest) GetTypeCreate() bool`

GetTypeCreate returns the TypeCreate field if non-nil, zero value otherwise.

### GetTypeCreateOk

`func (o *PatchedBulkWritableWebhookRequest) GetTypeCreateOk() (*bool, bool)`

GetTypeCreateOk returns a tuple with the TypeCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeCreate

`func (o *PatchedBulkWritableWebhookRequest) SetTypeCreate(v bool)`

SetTypeCreate sets TypeCreate field to given value.

### HasTypeCreate

`func (o *PatchedBulkWritableWebhookRequest) HasTypeCreate() bool`

HasTypeCreate returns a boolean if a field has been set.

### GetTypeUpdate

`func (o *PatchedBulkWritableWebhookRequest) GetTypeUpdate() bool`

GetTypeUpdate returns the TypeUpdate field if non-nil, zero value otherwise.

### GetTypeUpdateOk

`func (o *PatchedBulkWritableWebhookRequest) GetTypeUpdateOk() (*bool, bool)`

GetTypeUpdateOk returns a tuple with the TypeUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeUpdate

`func (o *PatchedBulkWritableWebhookRequest) SetTypeUpdate(v bool)`

SetTypeUpdate sets TypeUpdate field to given value.

### HasTypeUpdate

`func (o *PatchedBulkWritableWebhookRequest) HasTypeUpdate() bool`

HasTypeUpdate returns a boolean if a field has been set.

### GetTypeDelete

`func (o *PatchedBulkWritableWebhookRequest) GetTypeDelete() bool`

GetTypeDelete returns the TypeDelete field if non-nil, zero value otherwise.

### GetTypeDeleteOk

`func (o *PatchedBulkWritableWebhookRequest) GetTypeDeleteOk() (*bool, bool)`

GetTypeDeleteOk returns a tuple with the TypeDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeDelete

`func (o *PatchedBulkWritableWebhookRequest) SetTypeDelete(v bool)`

SetTypeDelete sets TypeDelete field to given value.

### HasTypeDelete

`func (o *PatchedBulkWritableWebhookRequest) HasTypeDelete() bool`

HasTypeDelete returns a boolean if a field has been set.

### GetPayloadUrl

`func (o *PatchedBulkWritableWebhookRequest) GetPayloadUrl() string`

GetPayloadUrl returns the PayloadUrl field if non-nil, zero value otherwise.

### GetPayloadUrlOk

`func (o *PatchedBulkWritableWebhookRequest) GetPayloadUrlOk() (*string, bool)`

GetPayloadUrlOk returns a tuple with the PayloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadUrl

`func (o *PatchedBulkWritableWebhookRequest) SetPayloadUrl(v string)`

SetPayloadUrl sets PayloadUrl field to given value.

### HasPayloadUrl

`func (o *PatchedBulkWritableWebhookRequest) HasPayloadUrl() bool`

HasPayloadUrl returns a boolean if a field has been set.

### GetEnabled

`func (o *PatchedBulkWritableWebhookRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PatchedBulkWritableWebhookRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PatchedBulkWritableWebhookRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PatchedBulkWritableWebhookRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHttpMethod

`func (o *PatchedBulkWritableWebhookRequest) GetHttpMethod() HttpMethodEnum`

GetHttpMethod returns the HttpMethod field if non-nil, zero value otherwise.

### GetHttpMethodOk

`func (o *PatchedBulkWritableWebhookRequest) GetHttpMethodOk() (*HttpMethodEnum, bool)`

GetHttpMethodOk returns a tuple with the HttpMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpMethod

`func (o *PatchedBulkWritableWebhookRequest) SetHttpMethod(v HttpMethodEnum)`

SetHttpMethod sets HttpMethod field to given value.

### HasHttpMethod

`func (o *PatchedBulkWritableWebhookRequest) HasHttpMethod() bool`

HasHttpMethod returns a boolean if a field has been set.

### GetHttpContentType

`func (o *PatchedBulkWritableWebhookRequest) GetHttpContentType() string`

GetHttpContentType returns the HttpContentType field if non-nil, zero value otherwise.

### GetHttpContentTypeOk

`func (o *PatchedBulkWritableWebhookRequest) GetHttpContentTypeOk() (*string, bool)`

GetHttpContentTypeOk returns a tuple with the HttpContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpContentType

`func (o *PatchedBulkWritableWebhookRequest) SetHttpContentType(v string)`

SetHttpContentType sets HttpContentType field to given value.

### HasHttpContentType

`func (o *PatchedBulkWritableWebhookRequest) HasHttpContentType() bool`

HasHttpContentType returns a boolean if a field has been set.

### GetAdditionalHeaders

`func (o *PatchedBulkWritableWebhookRequest) GetAdditionalHeaders() string`

GetAdditionalHeaders returns the AdditionalHeaders field if non-nil, zero value otherwise.

### GetAdditionalHeadersOk

`func (o *PatchedBulkWritableWebhookRequest) GetAdditionalHeadersOk() (*string, bool)`

GetAdditionalHeadersOk returns a tuple with the AdditionalHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalHeaders

`func (o *PatchedBulkWritableWebhookRequest) SetAdditionalHeaders(v string)`

SetAdditionalHeaders sets AdditionalHeaders field to given value.

### HasAdditionalHeaders

`func (o *PatchedBulkWritableWebhookRequest) HasAdditionalHeaders() bool`

HasAdditionalHeaders returns a boolean if a field has been set.

### GetBodyTemplate

`func (o *PatchedBulkWritableWebhookRequest) GetBodyTemplate() string`

GetBodyTemplate returns the BodyTemplate field if non-nil, zero value otherwise.

### GetBodyTemplateOk

`func (o *PatchedBulkWritableWebhookRequest) GetBodyTemplateOk() (*string, bool)`

GetBodyTemplateOk returns a tuple with the BodyTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyTemplate

`func (o *PatchedBulkWritableWebhookRequest) SetBodyTemplate(v string)`

SetBodyTemplate sets BodyTemplate field to given value.

### HasBodyTemplate

`func (o *PatchedBulkWritableWebhookRequest) HasBodyTemplate() bool`

HasBodyTemplate returns a boolean if a field has been set.

### GetSecret

`func (o *PatchedBulkWritableWebhookRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *PatchedBulkWritableWebhookRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *PatchedBulkWritableWebhookRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *PatchedBulkWritableWebhookRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetSslVerification

`func (o *PatchedBulkWritableWebhookRequest) GetSslVerification() bool`

GetSslVerification returns the SslVerification field if non-nil, zero value otherwise.

### GetSslVerificationOk

`func (o *PatchedBulkWritableWebhookRequest) GetSslVerificationOk() (*bool, bool)`

GetSslVerificationOk returns a tuple with the SslVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslVerification

`func (o *PatchedBulkWritableWebhookRequest) SetSslVerification(v bool)`

SetSslVerification sets SslVerification field to given value.

### HasSslVerification

`func (o *PatchedBulkWritableWebhookRequest) HasSslVerification() bool`

HasSslVerification returns a boolean if a field has been set.

### GetCaFilePath

`func (o *PatchedBulkWritableWebhookRequest) GetCaFilePath() string`

GetCaFilePath returns the CaFilePath field if non-nil, zero value otherwise.

### GetCaFilePathOk

`func (o *PatchedBulkWritableWebhookRequest) GetCaFilePathOk() (*string, bool)`

GetCaFilePathOk returns a tuple with the CaFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaFilePath

`func (o *PatchedBulkWritableWebhookRequest) SetCaFilePath(v string)`

SetCaFilePath sets CaFilePath field to given value.

### HasCaFilePath

`func (o *PatchedBulkWritableWebhookRequest) HasCaFilePath() bool`

HasCaFilePath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


