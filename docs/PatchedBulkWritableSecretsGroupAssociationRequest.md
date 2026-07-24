# PatchedBulkWritableSecretsGroupAssociationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**AccessType** | Pointer to [**AccessTypeEnum**](AccessTypeEnum.md) |  | [optional] 
**SecretType** | Pointer to [**SecretTypeEnum**](SecretTypeEnum.md) |  | [optional] 
**SecretsGroup** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**Secret** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewPatchedBulkWritableSecretsGroupAssociationRequest

`func NewPatchedBulkWritableSecretsGroupAssociationRequest(id string, ) *PatchedBulkWritableSecretsGroupAssociationRequest`

NewPatchedBulkWritableSecretsGroupAssociationRequest instantiates a new PatchedBulkWritableSecretsGroupAssociationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableSecretsGroupAssociationRequestWithDefaults

`func NewPatchedBulkWritableSecretsGroupAssociationRequestWithDefaults() *PatchedBulkWritableSecretsGroupAssociationRequest`

NewPatchedBulkWritableSecretsGroupAssociationRequestWithDefaults instantiates a new PatchedBulkWritableSecretsGroupAssociationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableSecretsGroupAssociationRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableSecretsGroupAssociationRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableSecretsGroupAssociationRequest) SetId(v string)`

SetId sets Id field to given value.


### GetAccessType

`func (o *PatchedBulkWritableSecretsGroupAssociationRequest) GetAccessType() AccessTypeEnum`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *PatchedBulkWritableSecretsGroupAssociationRequest) GetAccessTypeOk() (*AccessTypeEnum, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *PatchedBulkWritableSecretsGroupAssociationRequest) SetAccessType(v AccessTypeEnum)`

SetAccessType sets AccessType field to given value.

### HasAccessType

`func (o *PatchedBulkWritableSecretsGroupAssociationRequest) HasAccessType() bool`

HasAccessType returns a boolean if a field has been set.

### GetSecretType

`func (o *PatchedBulkWritableSecretsGroupAssociationRequest) GetSecretType() SecretTypeEnum`

GetSecretType returns the SecretType field if non-nil, zero value otherwise.

### GetSecretTypeOk

`func (o *PatchedBulkWritableSecretsGroupAssociationRequest) GetSecretTypeOk() (*SecretTypeEnum, bool)`

GetSecretTypeOk returns a tuple with the SecretType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretType

`func (o *PatchedBulkWritableSecretsGroupAssociationRequest) SetSecretType(v SecretTypeEnum)`

SetSecretType sets SecretType field to given value.

### HasSecretType

`func (o *PatchedBulkWritableSecretsGroupAssociationRequest) HasSecretType() bool`

HasSecretType returns a boolean if a field has been set.

### GetSecretsGroup

`func (o *PatchedBulkWritableSecretsGroupAssociationRequest) GetSecretsGroup() BulkWritableCableRequestStatus`

GetSecretsGroup returns the SecretsGroup field if non-nil, zero value otherwise.

### GetSecretsGroupOk

`func (o *PatchedBulkWritableSecretsGroupAssociationRequest) GetSecretsGroupOk() (*BulkWritableCableRequestStatus, bool)`

GetSecretsGroupOk returns a tuple with the SecretsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretsGroup

`func (o *PatchedBulkWritableSecretsGroupAssociationRequest) SetSecretsGroup(v BulkWritableCableRequestStatus)`

SetSecretsGroup sets SecretsGroup field to given value.

### HasSecretsGroup

`func (o *PatchedBulkWritableSecretsGroupAssociationRequest) HasSecretsGroup() bool`

HasSecretsGroup returns a boolean if a field has been set.

### GetSecret

`func (o *PatchedBulkWritableSecretsGroupAssociationRequest) GetSecret() BulkWritableCableRequestStatus`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *PatchedBulkWritableSecretsGroupAssociationRequest) GetSecretOk() (*BulkWritableCableRequestStatus, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *PatchedBulkWritableSecretsGroupAssociationRequest) SetSecret(v BulkWritableCableRequestStatus)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *PatchedBulkWritableSecretsGroupAssociationRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


