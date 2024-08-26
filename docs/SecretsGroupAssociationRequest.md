# SecretsGroupAssociationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessType** | [**AccessTypeEnum**](AccessTypeEnum.md) |  | 
**SecretType** | [**SecretTypeEnum**](SecretTypeEnum.md) |  | 
**SecretsGroup** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Secret** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 

## Methods

### NewSecretsGroupAssociationRequest

`func NewSecretsGroupAssociationRequest(accessType AccessTypeEnum, secretType SecretTypeEnum, secretsGroup BulkWritableCableRequestStatus, secret BulkWritableCableRequestStatus, ) *SecretsGroupAssociationRequest`

NewSecretsGroupAssociationRequest instantiates a new SecretsGroupAssociationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretsGroupAssociationRequestWithDefaults

`func NewSecretsGroupAssociationRequestWithDefaults() *SecretsGroupAssociationRequest`

NewSecretsGroupAssociationRequestWithDefaults instantiates a new SecretsGroupAssociationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessType

`func (o *SecretsGroupAssociationRequest) GetAccessType() AccessTypeEnum`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *SecretsGroupAssociationRequest) GetAccessTypeOk() (*AccessTypeEnum, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *SecretsGroupAssociationRequest) SetAccessType(v AccessTypeEnum)`

SetAccessType sets AccessType field to given value.


### GetSecretType

`func (o *SecretsGroupAssociationRequest) GetSecretType() SecretTypeEnum`

GetSecretType returns the SecretType field if non-nil, zero value otherwise.

### GetSecretTypeOk

`func (o *SecretsGroupAssociationRequest) GetSecretTypeOk() (*SecretTypeEnum, bool)`

GetSecretTypeOk returns a tuple with the SecretType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretType

`func (o *SecretsGroupAssociationRequest) SetSecretType(v SecretTypeEnum)`

SetSecretType sets SecretType field to given value.


### GetSecretsGroup

`func (o *SecretsGroupAssociationRequest) GetSecretsGroup() BulkWritableCableRequestStatus`

GetSecretsGroup returns the SecretsGroup field if non-nil, zero value otherwise.

### GetSecretsGroupOk

`func (o *SecretsGroupAssociationRequest) GetSecretsGroupOk() (*BulkWritableCableRequestStatus, bool)`

GetSecretsGroupOk returns a tuple with the SecretsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretsGroup

`func (o *SecretsGroupAssociationRequest) SetSecretsGroup(v BulkWritableCableRequestStatus)`

SetSecretsGroup sets SecretsGroup field to given value.


### GetSecret

`func (o *SecretsGroupAssociationRequest) GetSecret() BulkWritableCableRequestStatus`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *SecretsGroupAssociationRequest) GetSecretOk() (*BulkWritableCableRequestStatus, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *SecretsGroupAssociationRequest) SetSecret(v BulkWritableCableRequestStatus)`

SetSecret sets Secret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


