# PatchedSecretsGroupAssociationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessType** | Pointer to [**AccessTypeEnum**](AccessTypeEnum.md) |  | [optional] 
**SecretType** | Pointer to [**SecretTypeEnum**](SecretTypeEnum.md) |  | [optional] 
**SecretsGroup** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**Secret** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewPatchedSecretsGroupAssociationRequest

`func NewPatchedSecretsGroupAssociationRequest() *PatchedSecretsGroupAssociationRequest`

NewPatchedSecretsGroupAssociationRequest instantiates a new PatchedSecretsGroupAssociationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedSecretsGroupAssociationRequestWithDefaults

`func NewPatchedSecretsGroupAssociationRequestWithDefaults() *PatchedSecretsGroupAssociationRequest`

NewPatchedSecretsGroupAssociationRequestWithDefaults instantiates a new PatchedSecretsGroupAssociationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessType

`func (o *PatchedSecretsGroupAssociationRequest) GetAccessType() AccessTypeEnum`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *PatchedSecretsGroupAssociationRequest) GetAccessTypeOk() (*AccessTypeEnum, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *PatchedSecretsGroupAssociationRequest) SetAccessType(v AccessTypeEnum)`

SetAccessType sets AccessType field to given value.

### HasAccessType

`func (o *PatchedSecretsGroupAssociationRequest) HasAccessType() bool`

HasAccessType returns a boolean if a field has been set.

### GetSecretType

`func (o *PatchedSecretsGroupAssociationRequest) GetSecretType() SecretTypeEnum`

GetSecretType returns the SecretType field if non-nil, zero value otherwise.

### GetSecretTypeOk

`func (o *PatchedSecretsGroupAssociationRequest) GetSecretTypeOk() (*SecretTypeEnum, bool)`

GetSecretTypeOk returns a tuple with the SecretType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretType

`func (o *PatchedSecretsGroupAssociationRequest) SetSecretType(v SecretTypeEnum)`

SetSecretType sets SecretType field to given value.

### HasSecretType

`func (o *PatchedSecretsGroupAssociationRequest) HasSecretType() bool`

HasSecretType returns a boolean if a field has been set.

### GetSecretsGroup

`func (o *PatchedSecretsGroupAssociationRequest) GetSecretsGroup() BulkWritableCableRequestStatus`

GetSecretsGroup returns the SecretsGroup field if non-nil, zero value otherwise.

### GetSecretsGroupOk

`func (o *PatchedSecretsGroupAssociationRequest) GetSecretsGroupOk() (*BulkWritableCableRequestStatus, bool)`

GetSecretsGroupOk returns a tuple with the SecretsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretsGroup

`func (o *PatchedSecretsGroupAssociationRequest) SetSecretsGroup(v BulkWritableCableRequestStatus)`

SetSecretsGroup sets SecretsGroup field to given value.

### HasSecretsGroup

`func (o *PatchedSecretsGroupAssociationRequest) HasSecretsGroup() bool`

HasSecretsGroup returns a boolean if a field has been set.

### GetSecret

`func (o *PatchedSecretsGroupAssociationRequest) GetSecret() BulkWritableCableRequestStatus`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *PatchedSecretsGroupAssociationRequest) GetSecretOk() (*BulkWritableCableRequestStatus, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *PatchedSecretsGroupAssociationRequest) SetSecret(v BulkWritableCableRequestStatus)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *PatchedSecretsGroupAssociationRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


