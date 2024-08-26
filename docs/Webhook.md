# Webhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**ContentTypes** | **[]string** |  | 
**Name** | **string** |  | 
**TypeCreate** | Pointer to **bool** | Call this webhook when a matching object is created. | [optional] 
**TypeUpdate** | Pointer to **bool** | Call this webhook when a matching object is updated. | [optional] 
**TypeDelete** | Pointer to **bool** | Call this webhook when a matching object is deleted. | [optional] 
**PayloadUrl** | **string** | A POST will be sent to this URL when the webhook is called. | 
**Enabled** | Pointer to **bool** |  | [optional] 
**HttpMethod** | Pointer to [**HttpMethodEnum**](HttpMethodEnum.md) |  | [optional] 
**HttpContentType** | Pointer to **string** | The complete list of official content types is available &lt;a href&#x3D;\&quot;https://www.iana.org/assignments/media-types/media-types.xhtml\&quot;&gt;here&lt;/a&gt;. | [optional] 
**AdditionalHeaders** | Pointer to **string** | User-supplied HTTP headers to be sent with the request in addition to the HTTP content type. Headers should be defined in the format &lt;code&gt;Name: Value&lt;/code&gt;. Jinja2 template processing is supported with the same context as the request body (below). | [optional] 
**BodyTemplate** | Pointer to **string** | Jinja2 template for a custom request body. If blank, a JSON object representing the change will be included. Available context data includes: &lt;code&gt;event&lt;/code&gt;, &lt;code&gt;model&lt;/code&gt;, &lt;code&gt;timestamp&lt;/code&gt;, &lt;code&gt;username&lt;/code&gt;, &lt;code&gt;request_id&lt;/code&gt;, and &lt;code&gt;data&lt;/code&gt;. | [optional] 
**Secret** | Pointer to **string** | When provided, the request will include a &#39;X-Hook-Signature&#39; header containing a HMAC hex digest of the payload body using the secret as the key. The secret is not transmitted in the request. | [optional] 
**SslVerification** | Pointer to **bool** | Enable SSL certificate verification. Disable with caution! | [optional] 
**CaFilePath** | Pointer to **string** | The specific CA certificate file to use for SSL verification. Leave blank to use the system defaults. | [optional] 
**Created** | **NullableTime** |  | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 
**NotesUrl** | **string** |  | [readonly] 

## Methods

### NewWebhook

`func NewWebhook(id string, objectType string, display string, url string, naturalSlug string, contentTypes []string, name string, payloadUrl string, created NullableTime, lastUpdated NullableTime, notesUrl string, ) *Webhook`

NewWebhook instantiates a new Webhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookWithDefaults

`func NewWebhookWithDefaults() *Webhook`

NewWebhookWithDefaults instantiates a new Webhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Webhook) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Webhook) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Webhook) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *Webhook) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *Webhook) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *Webhook) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *Webhook) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *Webhook) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *Webhook) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *Webhook) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Webhook) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Webhook) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *Webhook) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *Webhook) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *Webhook) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetContentTypes

`func (o *Webhook) GetContentTypes() []string`

GetContentTypes returns the ContentTypes field if non-nil, zero value otherwise.

### GetContentTypesOk

`func (o *Webhook) GetContentTypesOk() (*[]string, bool)`

GetContentTypesOk returns a tuple with the ContentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentTypes

`func (o *Webhook) SetContentTypes(v []string)`

SetContentTypes sets ContentTypes field to given value.


### GetName

`func (o *Webhook) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Webhook) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Webhook) SetName(v string)`

SetName sets Name field to given value.


### GetTypeCreate

`func (o *Webhook) GetTypeCreate() bool`

GetTypeCreate returns the TypeCreate field if non-nil, zero value otherwise.

### GetTypeCreateOk

`func (o *Webhook) GetTypeCreateOk() (*bool, bool)`

GetTypeCreateOk returns a tuple with the TypeCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeCreate

`func (o *Webhook) SetTypeCreate(v bool)`

SetTypeCreate sets TypeCreate field to given value.

### HasTypeCreate

`func (o *Webhook) HasTypeCreate() bool`

HasTypeCreate returns a boolean if a field has been set.

### GetTypeUpdate

`func (o *Webhook) GetTypeUpdate() bool`

GetTypeUpdate returns the TypeUpdate field if non-nil, zero value otherwise.

### GetTypeUpdateOk

`func (o *Webhook) GetTypeUpdateOk() (*bool, bool)`

GetTypeUpdateOk returns a tuple with the TypeUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeUpdate

`func (o *Webhook) SetTypeUpdate(v bool)`

SetTypeUpdate sets TypeUpdate field to given value.

### HasTypeUpdate

`func (o *Webhook) HasTypeUpdate() bool`

HasTypeUpdate returns a boolean if a field has been set.

### GetTypeDelete

`func (o *Webhook) GetTypeDelete() bool`

GetTypeDelete returns the TypeDelete field if non-nil, zero value otherwise.

### GetTypeDeleteOk

`func (o *Webhook) GetTypeDeleteOk() (*bool, bool)`

GetTypeDeleteOk returns a tuple with the TypeDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeDelete

`func (o *Webhook) SetTypeDelete(v bool)`

SetTypeDelete sets TypeDelete field to given value.

### HasTypeDelete

`func (o *Webhook) HasTypeDelete() bool`

HasTypeDelete returns a boolean if a field has been set.

### GetPayloadUrl

`func (o *Webhook) GetPayloadUrl() string`

GetPayloadUrl returns the PayloadUrl field if non-nil, zero value otherwise.

### GetPayloadUrlOk

`func (o *Webhook) GetPayloadUrlOk() (*string, bool)`

GetPayloadUrlOk returns a tuple with the PayloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadUrl

`func (o *Webhook) SetPayloadUrl(v string)`

SetPayloadUrl sets PayloadUrl field to given value.


### GetEnabled

`func (o *Webhook) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Webhook) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Webhook) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Webhook) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHttpMethod

`func (o *Webhook) GetHttpMethod() HttpMethodEnum`

GetHttpMethod returns the HttpMethod field if non-nil, zero value otherwise.

### GetHttpMethodOk

`func (o *Webhook) GetHttpMethodOk() (*HttpMethodEnum, bool)`

GetHttpMethodOk returns a tuple with the HttpMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpMethod

`func (o *Webhook) SetHttpMethod(v HttpMethodEnum)`

SetHttpMethod sets HttpMethod field to given value.

### HasHttpMethod

`func (o *Webhook) HasHttpMethod() bool`

HasHttpMethod returns a boolean if a field has been set.

### GetHttpContentType

`func (o *Webhook) GetHttpContentType() string`

GetHttpContentType returns the HttpContentType field if non-nil, zero value otherwise.

### GetHttpContentTypeOk

`func (o *Webhook) GetHttpContentTypeOk() (*string, bool)`

GetHttpContentTypeOk returns a tuple with the HttpContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpContentType

`func (o *Webhook) SetHttpContentType(v string)`

SetHttpContentType sets HttpContentType field to given value.

### HasHttpContentType

`func (o *Webhook) HasHttpContentType() bool`

HasHttpContentType returns a boolean if a field has been set.

### GetAdditionalHeaders

`func (o *Webhook) GetAdditionalHeaders() string`

GetAdditionalHeaders returns the AdditionalHeaders field if non-nil, zero value otherwise.

### GetAdditionalHeadersOk

`func (o *Webhook) GetAdditionalHeadersOk() (*string, bool)`

GetAdditionalHeadersOk returns a tuple with the AdditionalHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalHeaders

`func (o *Webhook) SetAdditionalHeaders(v string)`

SetAdditionalHeaders sets AdditionalHeaders field to given value.

### HasAdditionalHeaders

`func (o *Webhook) HasAdditionalHeaders() bool`

HasAdditionalHeaders returns a boolean if a field has been set.

### GetBodyTemplate

`func (o *Webhook) GetBodyTemplate() string`

GetBodyTemplate returns the BodyTemplate field if non-nil, zero value otherwise.

### GetBodyTemplateOk

`func (o *Webhook) GetBodyTemplateOk() (*string, bool)`

GetBodyTemplateOk returns a tuple with the BodyTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyTemplate

`func (o *Webhook) SetBodyTemplate(v string)`

SetBodyTemplate sets BodyTemplate field to given value.

### HasBodyTemplate

`func (o *Webhook) HasBodyTemplate() bool`

HasBodyTemplate returns a boolean if a field has been set.

### GetSecret

`func (o *Webhook) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *Webhook) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *Webhook) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *Webhook) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetSslVerification

`func (o *Webhook) GetSslVerification() bool`

GetSslVerification returns the SslVerification field if non-nil, zero value otherwise.

### GetSslVerificationOk

`func (o *Webhook) GetSslVerificationOk() (*bool, bool)`

GetSslVerificationOk returns a tuple with the SslVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslVerification

`func (o *Webhook) SetSslVerification(v bool)`

SetSslVerification sets SslVerification field to given value.

### HasSslVerification

`func (o *Webhook) HasSslVerification() bool`

HasSslVerification returns a boolean if a field has been set.

### GetCaFilePath

`func (o *Webhook) GetCaFilePath() string`

GetCaFilePath returns the CaFilePath field if non-nil, zero value otherwise.

### GetCaFilePathOk

`func (o *Webhook) GetCaFilePathOk() (*string, bool)`

GetCaFilePathOk returns a tuple with the CaFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaFilePath

`func (o *Webhook) SetCaFilePath(v string)`

SetCaFilePath sets CaFilePath field to given value.

### HasCaFilePath

`func (o *Webhook) HasCaFilePath() bool`

HasCaFilePath returns a boolean if a field has been set.

### GetCreated

`func (o *Webhook) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Webhook) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Webhook) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *Webhook) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *Webhook) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetLastUpdated

`func (o *Webhook) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Webhook) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Webhook) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *Webhook) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *Webhook) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetNotesUrl

`func (o *Webhook) GetNotesUrl() string`

GetNotesUrl returns the NotesUrl field if non-nil, zero value otherwise.

### GetNotesUrlOk

`func (o *Webhook) GetNotesUrlOk() (*string, bool)`

GetNotesUrlOk returns a tuple with the NotesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesUrl

`func (o *Webhook) SetNotesUrl(v string)`

SetNotesUrl sets NotesUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


