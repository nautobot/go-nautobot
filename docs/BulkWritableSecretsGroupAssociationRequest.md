# BulkWritableSecretsGroupAssociationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**AccessType** | [**AccessTypeEnum**](AccessTypeEnum.md) |  | 
**SecretType** | [**SecretTypeEnum**](SecretTypeEnum.md) |  | 
**SecretsGroup** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 
**Secret** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 

## Methods

### NewBulkWritableSecretsGroupAssociationRequest

`func NewBulkWritableSecretsGroupAssociationRequest(id string, accessType AccessTypeEnum, secretType SecretTypeEnum, secretsGroup ApprovalWorkflowStageResponseApprovalWorkflowStage, secret ApprovalWorkflowStageResponseApprovalWorkflowStage, ) *BulkWritableSecretsGroupAssociationRequest`

NewBulkWritableSecretsGroupAssociationRequest instantiates a new BulkWritableSecretsGroupAssociationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritableSecretsGroupAssociationRequestWithDefaults

`func NewBulkWritableSecretsGroupAssociationRequestWithDefaults() *BulkWritableSecretsGroupAssociationRequest`

NewBulkWritableSecretsGroupAssociationRequestWithDefaults instantiates a new BulkWritableSecretsGroupAssociationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritableSecretsGroupAssociationRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritableSecretsGroupAssociationRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritableSecretsGroupAssociationRequest) SetId(v string)`

SetId sets Id field to given value.


### GetAccessType

`func (o *BulkWritableSecretsGroupAssociationRequest) GetAccessType() AccessTypeEnum`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *BulkWritableSecretsGroupAssociationRequest) GetAccessTypeOk() (*AccessTypeEnum, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *BulkWritableSecretsGroupAssociationRequest) SetAccessType(v AccessTypeEnum)`

SetAccessType sets AccessType field to given value.


### GetSecretType

`func (o *BulkWritableSecretsGroupAssociationRequest) GetSecretType() SecretTypeEnum`

GetSecretType returns the SecretType field if non-nil, zero value otherwise.

### GetSecretTypeOk

`func (o *BulkWritableSecretsGroupAssociationRequest) GetSecretTypeOk() (*SecretTypeEnum, bool)`

GetSecretTypeOk returns a tuple with the SecretType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretType

`func (o *BulkWritableSecretsGroupAssociationRequest) SetSecretType(v SecretTypeEnum)`

SetSecretType sets SecretType field to given value.


### GetSecretsGroup

`func (o *BulkWritableSecretsGroupAssociationRequest) GetSecretsGroup() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetSecretsGroup returns the SecretsGroup field if non-nil, zero value otherwise.

### GetSecretsGroupOk

`func (o *BulkWritableSecretsGroupAssociationRequest) GetSecretsGroupOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetSecretsGroupOk returns a tuple with the SecretsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretsGroup

`func (o *BulkWritableSecretsGroupAssociationRequest) SetSecretsGroup(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetSecretsGroup sets SecretsGroup field to given value.


### GetSecret

`func (o *BulkWritableSecretsGroupAssociationRequest) GetSecret() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *BulkWritableSecretsGroupAssociationRequest) GetSecretOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *BulkWritableSecretsGroupAssociationRequest) SetSecret(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetSecret sets Secret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


