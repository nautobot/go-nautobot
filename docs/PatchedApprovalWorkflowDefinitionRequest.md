# PatchedApprovalWorkflowDefinitionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ModelContentType** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ModelConstraints** | Pointer to **interface{}** | Constraints to filter the objects that can be approved using this workflow. | [optional] 
**Weight** | Pointer to **int32** | Determines workflow relevance when multiple apply. Higher weight wins. | [optional] [default to 0]
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewPatchedApprovalWorkflowDefinitionRequest

`func NewPatchedApprovalWorkflowDefinitionRequest() *PatchedApprovalWorkflowDefinitionRequest`

NewPatchedApprovalWorkflowDefinitionRequest instantiates a new PatchedApprovalWorkflowDefinitionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedApprovalWorkflowDefinitionRequestWithDefaults

`func NewPatchedApprovalWorkflowDefinitionRequestWithDefaults() *PatchedApprovalWorkflowDefinitionRequest`

NewPatchedApprovalWorkflowDefinitionRequestWithDefaults instantiates a new PatchedApprovalWorkflowDefinitionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedApprovalWorkflowDefinitionRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedApprovalWorkflowDefinitionRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedApprovalWorkflowDefinitionRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedApprovalWorkflowDefinitionRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModelContentType

`func (o *PatchedApprovalWorkflowDefinitionRequest) GetModelContentType() string`

GetModelContentType returns the ModelContentType field if non-nil, zero value otherwise.

### GetModelContentTypeOk

`func (o *PatchedApprovalWorkflowDefinitionRequest) GetModelContentTypeOk() (*string, bool)`

GetModelContentTypeOk returns a tuple with the ModelContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelContentType

`func (o *PatchedApprovalWorkflowDefinitionRequest) SetModelContentType(v string)`

SetModelContentType sets ModelContentType field to given value.

### HasModelContentType

`func (o *PatchedApprovalWorkflowDefinitionRequest) HasModelContentType() bool`

HasModelContentType returns a boolean if a field has been set.

### GetName

`func (o *PatchedApprovalWorkflowDefinitionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedApprovalWorkflowDefinitionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedApprovalWorkflowDefinitionRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedApprovalWorkflowDefinitionRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetModelConstraints

`func (o *PatchedApprovalWorkflowDefinitionRequest) GetModelConstraints() interface{}`

GetModelConstraints returns the ModelConstraints field if non-nil, zero value otherwise.

### GetModelConstraintsOk

`func (o *PatchedApprovalWorkflowDefinitionRequest) GetModelConstraintsOk() (*interface{}, bool)`

GetModelConstraintsOk returns a tuple with the ModelConstraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelConstraints

`func (o *PatchedApprovalWorkflowDefinitionRequest) SetModelConstraints(v interface{})`

SetModelConstraints sets ModelConstraints field to given value.

### HasModelConstraints

`func (o *PatchedApprovalWorkflowDefinitionRequest) HasModelConstraints() bool`

HasModelConstraints returns a boolean if a field has been set.

### SetModelConstraintsNil

`func (o *PatchedApprovalWorkflowDefinitionRequest) SetModelConstraintsNil(b bool)`

 SetModelConstraintsNil sets the value for ModelConstraints to be an explicit nil

### UnsetModelConstraints
`func (o *PatchedApprovalWorkflowDefinitionRequest) UnsetModelConstraints()`

UnsetModelConstraints ensures that no value is present for ModelConstraints, not even an explicit nil
### GetWeight

`func (o *PatchedApprovalWorkflowDefinitionRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *PatchedApprovalWorkflowDefinitionRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *PatchedApprovalWorkflowDefinitionRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *PatchedApprovalWorkflowDefinitionRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedApprovalWorkflowDefinitionRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedApprovalWorkflowDefinitionRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedApprovalWorkflowDefinitionRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedApprovalWorkflowDefinitionRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedApprovalWorkflowDefinitionRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedApprovalWorkflowDefinitionRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedApprovalWorkflowDefinitionRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedApprovalWorkflowDefinitionRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


