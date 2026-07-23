# ApprovalWorkflowDefinitionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ModelContentType** | **string** |  | 
**Name** | **string** |  | 
**ModelConstraints** | Pointer to **interface{}** | Constraints to filter the objects that can be approved using this workflow. | [optional] 
**Weight** | Pointer to **int32** | Determines workflow relevance when multiple apply. Higher weight wins. | [optional] [default to 0]
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewApprovalWorkflowDefinitionRequest

`func NewApprovalWorkflowDefinitionRequest(modelContentType string, name string, ) *ApprovalWorkflowDefinitionRequest`

NewApprovalWorkflowDefinitionRequest instantiates a new ApprovalWorkflowDefinitionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalWorkflowDefinitionRequestWithDefaults

`func NewApprovalWorkflowDefinitionRequestWithDefaults() *ApprovalWorkflowDefinitionRequest`

NewApprovalWorkflowDefinitionRequestWithDefaults instantiates a new ApprovalWorkflowDefinitionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApprovalWorkflowDefinitionRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApprovalWorkflowDefinitionRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApprovalWorkflowDefinitionRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApprovalWorkflowDefinitionRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModelContentType

`func (o *ApprovalWorkflowDefinitionRequest) GetModelContentType() string`

GetModelContentType returns the ModelContentType field if non-nil, zero value otherwise.

### GetModelContentTypeOk

`func (o *ApprovalWorkflowDefinitionRequest) GetModelContentTypeOk() (*string, bool)`

GetModelContentTypeOk returns a tuple with the ModelContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelContentType

`func (o *ApprovalWorkflowDefinitionRequest) SetModelContentType(v string)`

SetModelContentType sets ModelContentType field to given value.


### GetName

`func (o *ApprovalWorkflowDefinitionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApprovalWorkflowDefinitionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApprovalWorkflowDefinitionRequest) SetName(v string)`

SetName sets Name field to given value.


### GetModelConstraints

`func (o *ApprovalWorkflowDefinitionRequest) GetModelConstraints() interface{}`

GetModelConstraints returns the ModelConstraints field if non-nil, zero value otherwise.

### GetModelConstraintsOk

`func (o *ApprovalWorkflowDefinitionRequest) GetModelConstraintsOk() (*interface{}, bool)`

GetModelConstraintsOk returns a tuple with the ModelConstraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelConstraints

`func (o *ApprovalWorkflowDefinitionRequest) SetModelConstraints(v interface{})`

SetModelConstraints sets ModelConstraints field to given value.

### HasModelConstraints

`func (o *ApprovalWorkflowDefinitionRequest) HasModelConstraints() bool`

HasModelConstraints returns a boolean if a field has been set.

### SetModelConstraintsNil

`func (o *ApprovalWorkflowDefinitionRequest) SetModelConstraintsNil(b bool)`

 SetModelConstraintsNil sets the value for ModelConstraints to be an explicit nil

### UnsetModelConstraints
`func (o *ApprovalWorkflowDefinitionRequest) UnsetModelConstraints()`

UnsetModelConstraints ensures that no value is present for ModelConstraints, not even an explicit nil
### GetWeight

`func (o *ApprovalWorkflowDefinitionRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *ApprovalWorkflowDefinitionRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *ApprovalWorkflowDefinitionRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *ApprovalWorkflowDefinitionRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetCustomFields

`func (o *ApprovalWorkflowDefinitionRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ApprovalWorkflowDefinitionRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ApprovalWorkflowDefinitionRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ApprovalWorkflowDefinitionRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *ApprovalWorkflowDefinitionRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ApprovalWorkflowDefinitionRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ApprovalWorkflowDefinitionRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ApprovalWorkflowDefinitionRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


