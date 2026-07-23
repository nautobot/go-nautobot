# SecretsGroupAssociation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**AccessType** | [**AccessTypeEnum**](AccessTypeEnum.md) |  | 
**SecretType** | [**SecretTypeEnum**](SecretTypeEnum.md) |  | 
**SecretsGroup** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 
**Secret** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 

## Methods

### NewSecretsGroupAssociation

`func NewSecretsGroupAssociation(objectType string, display string, url string, naturalSlug string, accessType AccessTypeEnum, secretType SecretTypeEnum, secretsGroup ApprovalWorkflowStageResponseApprovalWorkflowStage, secret ApprovalWorkflowStageResponseApprovalWorkflowStage, ) *SecretsGroupAssociation`

NewSecretsGroupAssociation instantiates a new SecretsGroupAssociation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretsGroupAssociationWithDefaults

`func NewSecretsGroupAssociationWithDefaults() *SecretsGroupAssociation`

NewSecretsGroupAssociationWithDefaults instantiates a new SecretsGroupAssociation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SecretsGroupAssociation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecretsGroupAssociation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecretsGroupAssociation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SecretsGroupAssociation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectType

`func (o *SecretsGroupAssociation) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SecretsGroupAssociation) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SecretsGroupAssociation) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *SecretsGroupAssociation) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *SecretsGroupAssociation) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *SecretsGroupAssociation) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *SecretsGroupAssociation) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SecretsGroupAssociation) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SecretsGroupAssociation) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *SecretsGroupAssociation) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *SecretsGroupAssociation) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *SecretsGroupAssociation) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetAccessType

`func (o *SecretsGroupAssociation) GetAccessType() AccessTypeEnum`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *SecretsGroupAssociation) GetAccessTypeOk() (*AccessTypeEnum, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *SecretsGroupAssociation) SetAccessType(v AccessTypeEnum)`

SetAccessType sets AccessType field to given value.


### GetSecretType

`func (o *SecretsGroupAssociation) GetSecretType() SecretTypeEnum`

GetSecretType returns the SecretType field if non-nil, zero value otherwise.

### GetSecretTypeOk

`func (o *SecretsGroupAssociation) GetSecretTypeOk() (*SecretTypeEnum, bool)`

GetSecretTypeOk returns a tuple with the SecretType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretType

`func (o *SecretsGroupAssociation) SetSecretType(v SecretTypeEnum)`

SetSecretType sets SecretType field to given value.


### GetSecretsGroup

`func (o *SecretsGroupAssociation) GetSecretsGroup() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetSecretsGroup returns the SecretsGroup field if non-nil, zero value otherwise.

### GetSecretsGroupOk

`func (o *SecretsGroupAssociation) GetSecretsGroupOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetSecretsGroupOk returns a tuple with the SecretsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretsGroup

`func (o *SecretsGroupAssociation) SetSecretsGroup(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetSecretsGroup sets SecretsGroup field to given value.


### GetSecret

`func (o *SecretsGroupAssociation) GetSecret() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *SecretsGroupAssociation) GetSecretOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *SecretsGroupAssociation) SetSecret(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetSecret sets Secret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


