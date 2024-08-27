# BulkWritableExternalIntegrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**RemoteUrl** | **string** |  | 
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

### NewBulkWritableExternalIntegrationRequest

`func NewBulkWritableExternalIntegrationRequest(id string, name string, remoteUrl string, ) *BulkWritableExternalIntegrationRequest`

NewBulkWritableExternalIntegrationRequest instantiates a new BulkWritableExternalIntegrationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableExternalIntegrationRequestWithDefaults

`func NewBulkWritableExternalIntegrationRequestWithDefaults() *BulkWritableExternalIntegrationRequest`

NewBulkWritableExternalIntegrationRequestWithDefaults instantiates a new BulkWritableExternalIntegrationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableExternalIntegrationRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableExternalIntegrationRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableExternalIntegrationRequest) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BulkWritableExternalIntegrationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkWritableExternalIntegrationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkWritableExternalIntegrationRequest) SetName(v string)`

SetName sets Name field to given value.


### GetRemoteUrl

`func (o *BulkWritableExternalIntegrationRequest) GetRemoteUrl() string`

GetRemoteUrl returns the RemoteUrl field if non-nil, zero value otherwise.

### GetRemoteUrlOk

`func (o *BulkWritableExternalIntegrationRequest) GetRemoteUrlOk() (*string, bool)`

GetRemoteUrlOk returns a tuple with the RemoteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteUrl

`func (o *BulkWritableExternalIntegrationRequest) SetRemoteUrl(v string)`

SetRemoteUrl sets RemoteUrl field to given value.


### GetVerifySsl

`func (o *BulkWritableExternalIntegrationRequest) GetVerifySsl() bool`

GetVerifySsl returns the VerifySsl field if non-nil, zero value otherwise.

### GetVerifySslOk

`func (o *BulkWritableExternalIntegrationRequest) GetVerifySslOk() (*bool, bool)`

GetVerifySslOk returns a tuple with the VerifySsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifySsl

`func (o *BulkWritableExternalIntegrationRequest) SetVerifySsl(v bool)`

SetVerifySsl sets VerifySsl field to given value.

### HasVerifySsl

`func (o *BulkWritableExternalIntegrationRequest) HasVerifySsl() bool`

HasVerifySsl returns a boolean if a field has been set.

### GetTimeout

`func (o *BulkWritableExternalIntegrationRequest) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *BulkWritableExternalIntegrationRequest) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *BulkWritableExternalIntegrationRequest) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *BulkWritableExternalIntegrationRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetExtraConfig

`func (o *BulkWritableExternalIntegrationRequest) GetExtraConfig() interface{}`

GetExtraConfig returns the ExtraConfig field if non-nil, zero value otherwise.

### GetExtraConfigOk

`func (o *BulkWritableExternalIntegrationRequest) GetExtraConfigOk() (*interface{}, bool)`

GetExtraConfigOk returns a tuple with the ExtraConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraConfig

`func (o *BulkWritableExternalIntegrationRequest) SetExtraConfig(v interface{})`

SetExtraConfig sets ExtraConfig field to given value.

### HasExtraConfig

`func (o *BulkWritableExternalIntegrationRequest) HasExtraConfig() bool`

HasExtraConfig returns a boolean if a field has been set.

### SetExtraConfigNil

`func (o *BulkWritableExternalIntegrationRequest) SetExtraConfigNil(b bool)`

 SetExtraConfigNil sets the value for ExtraConfig to be an explicit nil

### UnsetExtraConfig
`func (o *BulkWritableExternalIntegrationRequest) UnsetExtraConfig()`

UnsetExtraConfig ensures that no value is present for ExtraConfig, not even an explicit nil
### GetHttpMethod

`func (o *BulkWritableExternalIntegrationRequest) GetHttpMethod() BulkWritableExternalIntegrationRequestHttpMethod`

GetHttpMethod returns the HttpMethod field if non-nil, zero value otherwise.

### GetHttpMethodOk

`func (o *BulkWritableExternalIntegrationRequest) GetHttpMethodOk() (*BulkWritableExternalIntegrationRequestHttpMethod, bool)`

GetHttpMethodOk returns a tuple with the HttpMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpMethod

`func (o *BulkWritableExternalIntegrationRequest) SetHttpMethod(v BulkWritableExternalIntegrationRequestHttpMethod)`

SetHttpMethod sets HttpMethod field to given value.

### HasHttpMethod

`func (o *BulkWritableExternalIntegrationRequest) HasHttpMethod() bool`

HasHttpMethod returns a boolean if a field has been set.

### GetHeaders

`func (o *BulkWritableExternalIntegrationRequest) GetHeaders() interface{}`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *BulkWritableExternalIntegrationRequest) GetHeadersOk() (*interface{}, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *BulkWritableExternalIntegrationRequest) SetHeaders(v interface{})`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *BulkWritableExternalIntegrationRequest) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeadersNil

`func (o *BulkWritableExternalIntegrationRequest) SetHeadersNil(b bool)`

 SetHeadersNil sets the value for Headers to be an explicit nil

### UnsetHeaders
`func (o *BulkWritableExternalIntegrationRequest) UnsetHeaders()`

UnsetHeaders ensures that no value is present for Headers, not even an explicit nil
### GetCaFilePath

`func (o *BulkWritableExternalIntegrationRequest) GetCaFilePath() string`

GetCaFilePath returns the CaFilePath field if non-nil, zero value otherwise.

### GetCaFilePathOk

`func (o *BulkWritableExternalIntegrationRequest) GetCaFilePathOk() (*string, bool)`

GetCaFilePathOk returns a tuple with the CaFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaFilePath

`func (o *BulkWritableExternalIntegrationRequest) SetCaFilePath(v string)`

SetCaFilePath sets CaFilePath field to given value.

### HasCaFilePath

`func (o *BulkWritableExternalIntegrationRequest) HasCaFilePath() bool`

HasCaFilePath returns a boolean if a field has been set.

### GetSecretsGroup

`func (o *BulkWritableExternalIntegrationRequest) GetSecretsGroup() BulkWritableExternalIntegrationRequestSecretsGroup`

GetSecretsGroup returns the SecretsGroup field if non-nil, zero value otherwise.

### GetSecretsGroupOk

`func (o *BulkWritableExternalIntegrationRequest) GetSecretsGroupOk() (*BulkWritableExternalIntegrationRequestSecretsGroup, bool)`

GetSecretsGroupOk returns a tuple with the SecretsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretsGroup

`func (o *BulkWritableExternalIntegrationRequest) SetSecretsGroup(v BulkWritableExternalIntegrationRequestSecretsGroup)`

SetSecretsGroup sets SecretsGroup field to given value.

### HasSecretsGroup

`func (o *BulkWritableExternalIntegrationRequest) HasSecretsGroup() bool`

HasSecretsGroup returns a boolean if a field has been set.

### SetSecretsGroupNil

`func (o *BulkWritableExternalIntegrationRequest) SetSecretsGroupNil(b bool)`

 SetSecretsGroupNil sets the value for SecretsGroup to be an explicit nil

### UnsetSecretsGroup
`func (o *BulkWritableExternalIntegrationRequest) UnsetSecretsGroup()`

UnsetSecretsGroup ensures that no value is present for SecretsGroup, not even an explicit nil
### GetCustomFields

`func (o *BulkWritableExternalIntegrationRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkWritableExternalIntegrationRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkWritableExternalIntegrationRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkWritableExternalIntegrationRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *BulkWritableExternalIntegrationRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BulkWritableExternalIntegrationRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BulkWritableExternalIntegrationRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BulkWritableExternalIntegrationRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


