# PatchedExternalIntegrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**RemoteUrl** | Pointer to **string** |  | [optional] 
**VerifySsl** | Pointer to **bool** | Verify SSL certificates when connecting to the remote system | [optional] 
**Timeout** | Pointer to **int32** | Number of seconds to wait for a response | [optional] 
**ExtraConfig** | Pointer to **interface{}** | Optional user-defined JSON data for this integration | [optional] 
**HttpMethod** | Pointer to [**BulkWritableExternalIntegrationRequestHttpMethod**](BulkWritableExternalIntegrationRequestHttpMethod.md) |  | [optional] 
**Headers** | Pointer to **interface{}** | Headers for the HTTP request | [optional] 
**CaFilePath** | Pointer to **string** |  | [optional] 
**SecretsGroup** | Pointer to [**NullableBulkWritableExternalIntegrationRequestSecretsGroup**](BulkWritableExternalIntegrationRequestSecretsGroup.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewPatchedExternalIntegrationRequest

`func NewPatchedExternalIntegrationRequest() *PatchedExternalIntegrationRequest`

NewPatchedExternalIntegrationRequest instantiates a new PatchedExternalIntegrationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedExternalIntegrationRequestWithDefaults

`func NewPatchedExternalIntegrationRequestWithDefaults() *PatchedExternalIntegrationRequest`

NewPatchedExternalIntegrationRequestWithDefaults instantiates a new PatchedExternalIntegrationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedExternalIntegrationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedExternalIntegrationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedExternalIntegrationRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedExternalIntegrationRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRemoteUrl

`func (o *PatchedExternalIntegrationRequest) GetRemoteUrl() string`

GetRemoteUrl returns the RemoteUrl field if non-nil, zero value otherwise.

### GetRemoteUrlOk

`func (o *PatchedExternalIntegrationRequest) GetRemoteUrlOk() (*string, bool)`

GetRemoteUrlOk returns a tuple with the RemoteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteUrl

`func (o *PatchedExternalIntegrationRequest) SetRemoteUrl(v string)`

SetRemoteUrl sets RemoteUrl field to given value.

### HasRemoteUrl

`func (o *PatchedExternalIntegrationRequest) HasRemoteUrl() bool`

HasRemoteUrl returns a boolean if a field has been set.

### GetVerifySsl

`func (o *PatchedExternalIntegrationRequest) GetVerifySsl() bool`

GetVerifySsl returns the VerifySsl field if non-nil, zero value otherwise.

### GetVerifySslOk

`func (o *PatchedExternalIntegrationRequest) GetVerifySslOk() (*bool, bool)`

GetVerifySslOk returns a tuple with the VerifySsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifySsl

`func (o *PatchedExternalIntegrationRequest) SetVerifySsl(v bool)`

SetVerifySsl sets VerifySsl field to given value.

### HasVerifySsl

`func (o *PatchedExternalIntegrationRequest) HasVerifySsl() bool`

HasVerifySsl returns a boolean if a field has been set.

### GetTimeout

`func (o *PatchedExternalIntegrationRequest) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *PatchedExternalIntegrationRequest) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *PatchedExternalIntegrationRequest) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *PatchedExternalIntegrationRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetExtraConfig

`func (o *PatchedExternalIntegrationRequest) GetExtraConfig() interface{}`

GetExtraConfig returns the ExtraConfig field if non-nil, zero value otherwise.

### GetExtraConfigOk

`func (o *PatchedExternalIntegrationRequest) GetExtraConfigOk() (*interface{}, bool)`

GetExtraConfigOk returns a tuple with the ExtraConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraConfig

`func (o *PatchedExternalIntegrationRequest) SetExtraConfig(v interface{})`

SetExtraConfig sets ExtraConfig field to given value.

### HasExtraConfig

`func (o *PatchedExternalIntegrationRequest) HasExtraConfig() bool`

HasExtraConfig returns a boolean if a field has been set.

### SetExtraConfigNil

`func (o *PatchedExternalIntegrationRequest) SetExtraConfigNil(b bool)`

 SetExtraConfigNil sets the value for ExtraConfig to be an explicit nil

### UnsetExtraConfig
`func (o *PatchedExternalIntegrationRequest) UnsetExtraConfig()`

UnsetExtraConfig ensures that no value is present for ExtraConfig, not even an explicit nil
### GetHttpMethod

`func (o *PatchedExternalIntegrationRequest) GetHttpMethod() BulkWritableExternalIntegrationRequestHttpMethod`

GetHttpMethod returns the HttpMethod field if non-nil, zero value otherwise.

### GetHttpMethodOk

`func (o *PatchedExternalIntegrationRequest) GetHttpMethodOk() (*BulkWritableExternalIntegrationRequestHttpMethod, bool)`

GetHttpMethodOk returns a tuple with the HttpMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpMethod

`func (o *PatchedExternalIntegrationRequest) SetHttpMethod(v BulkWritableExternalIntegrationRequestHttpMethod)`

SetHttpMethod sets HttpMethod field to given value.

### HasHttpMethod

`func (o *PatchedExternalIntegrationRequest) HasHttpMethod() bool`

HasHttpMethod returns a boolean if a field has been set.

### GetHeaders

`func (o *PatchedExternalIntegrationRequest) GetHeaders() interface{}`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *PatchedExternalIntegrationRequest) GetHeadersOk() (*interface{}, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *PatchedExternalIntegrationRequest) SetHeaders(v interface{})`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *PatchedExternalIntegrationRequest) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeadersNil

`func (o *PatchedExternalIntegrationRequest) SetHeadersNil(b bool)`

 SetHeadersNil sets the value for Headers to be an explicit nil

### UnsetHeaders
`func (o *PatchedExternalIntegrationRequest) UnsetHeaders()`

UnsetHeaders ensures that no value is present for Headers, not even an explicit nil
### GetCaFilePath

`func (o *PatchedExternalIntegrationRequest) GetCaFilePath() string`

GetCaFilePath returns the CaFilePath field if non-nil, zero value otherwise.

### GetCaFilePathOk

`func (o *PatchedExternalIntegrationRequest) GetCaFilePathOk() (*string, bool)`

GetCaFilePathOk returns a tuple with the CaFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaFilePath

`func (o *PatchedExternalIntegrationRequest) SetCaFilePath(v string)`

SetCaFilePath sets CaFilePath field to given value.

### HasCaFilePath

`func (o *PatchedExternalIntegrationRequest) HasCaFilePath() bool`

HasCaFilePath returns a boolean if a field has been set.

### GetSecretsGroup

`func (o *PatchedExternalIntegrationRequest) GetSecretsGroup() BulkWritableExternalIntegrationRequestSecretsGroup`

GetSecretsGroup returns the SecretsGroup field if non-nil, zero value otherwise.

### GetSecretsGroupOk

`func (o *PatchedExternalIntegrationRequest) GetSecretsGroupOk() (*BulkWritableExternalIntegrationRequestSecretsGroup, bool)`

GetSecretsGroupOk returns a tuple with the SecretsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretsGroup

`func (o *PatchedExternalIntegrationRequest) SetSecretsGroup(v BulkWritableExternalIntegrationRequestSecretsGroup)`

SetSecretsGroup sets SecretsGroup field to given value.

### HasSecretsGroup

`func (o *PatchedExternalIntegrationRequest) HasSecretsGroup() bool`

HasSecretsGroup returns a boolean if a field has been set.

### SetSecretsGroupNil

`func (o *PatchedExternalIntegrationRequest) SetSecretsGroupNil(b bool)`

 SetSecretsGroupNil sets the value for SecretsGroup to be an explicit nil

### UnsetSecretsGroup
`func (o *PatchedExternalIntegrationRequest) UnsetSecretsGroup()`

UnsetSecretsGroup ensures that no value is present for SecretsGroup, not even an explicit nil
### GetCustomFields

`func (o *PatchedExternalIntegrationRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedExternalIntegrationRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedExternalIntegrationRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedExternalIntegrationRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedExternalIntegrationRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedExternalIntegrationRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedExternalIntegrationRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedExternalIntegrationRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


