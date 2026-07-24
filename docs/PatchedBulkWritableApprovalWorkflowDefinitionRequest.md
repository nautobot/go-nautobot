# PatchedBulkWritableApprovalWorkflowDefinitionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ModelContentType** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ModelConstraints** | Pointer to **interface{}** | Constraints to filter the objects that can be approved using this workflow. | [optional] 
**Weight** | Pointer to **int32** | Determines workflow relevance when multiple apply. Higher weight wins. | [optional] [default to 0]
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewPatchedBulkWritableApprovalWorkflowDefinitionRequest

`func NewPatchedBulkWritableApprovalWorkflowDefinitionRequest(id string, ) *PatchedBulkWritableApprovalWorkflowDefinitionRequest`

NewPatchedBulkWritableApprovalWorkflowDefinitionRequest instantiates a new PatchedBulkWritableApprovalWorkflowDefinitionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableApprovalWorkflowDefinitionRequestWithDefaults

`func NewPatchedBulkWritableApprovalWorkflowDefinitionRequestWithDefaults() *PatchedBulkWritableApprovalWorkflowDefinitionRequest`

NewPatchedBulkWritableApprovalWorkflowDefinitionRequestWithDefaults instantiates a new PatchedBulkWritableApprovalWorkflowDefinitionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableApprovalWorkflowDefinitionRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableApprovalWorkflowDefinitionRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableApprovalWorkflowDefinitionRequest) SetId(v string)`

SetId sets Id field to given value.


### GetModelContentType

`func (o *PatchedBulkWritableApprovalWorkflowDefinitionRequest) GetModelContentType() string`

GetModelContentType returns the ModelContentType field if non-nil, zero value otherwise.

### GetModelContentTypeOk

`func (o *PatchedBulkWritableApprovalWorkflowDefinitionRequest) GetModelContentTypeOk() (*string, bool)`

GetModelContentTypeOk returns a tuple with the ModelContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelContentType

`func (o *PatchedBulkWritableApprovalWorkflowDefinitionRequest) SetModelContentType(v string)`

SetModelContentType sets ModelContentType field to given value.

### HasModelContentType

`func (o *PatchedBulkWritableApprovalWorkflowDefinitionRequest) HasModelContentType() bool`

HasModelContentType returns a boolean if a field has been set.

### GetName

`func (o *PatchedBulkWritableApprovalWorkflowDefinitionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedBulkWritableApprovalWorkflowDefinitionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedBulkWritableApprovalWorkflowDefinitionRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedBulkWritableApprovalWorkflowDefinitionRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetModelConstraints

`func (o *PatchedBulkWritableApprovalWorkflowDefinitionRequest) GetModelConstraints() interface{}`

GetModelConstraints returns the ModelConstraints field if non-nil, zero value otherwise.

### GetModelConstraintsOk

`func (o *PatchedBulkWritableApprovalWorkflowDefinitionRequest) GetModelConstraintsOk() (*interface{}, bool)`

GetModelConstraintsOk returns a tuple with the ModelConstraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelConstraints

`func (o *PatchedBulkWritableApprovalWorkflowDefinitionRequest) SetModelConstraints(v interface{})`

SetModelConstraints sets ModelConstraints field to given value.

### HasModelConstraints

`func (o *PatchedBulkWritableApprovalWorkflowDefinitionRequest) HasModelConstraints() bool`

HasModelConstraints returns a boolean if a field has been set.

### SetModelConstraintsNil

`func (o *PatchedBulkWritableApprovalWorkflowDefinitionRequest) SetModelConstraintsNil(b bool)`

 SetModelConstraintsNil sets the value for ModelConstraints to be an explicit nil

### UnsetModelConstraints
`func (o *PatchedBulkWritableApprovalWorkflowDefinitionRequest) UnsetModelConstraints()`

UnsetModelConstraints ensures that no value is present for ModelConstraints, not even an explicit nil
### GetWeight

`func (o *PatchedBulkWritableApprovalWorkflowDefinitionRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *PatchedBulkWritableApprovalWorkflowDefinitionRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *PatchedBulkWritableApprovalWorkflowDefinitionRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *PatchedBulkWritableApprovalWorkflowDefinitionRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedBulkWritableApprovalWorkflowDefinitionRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkWritableApprovalWorkflowDefinitionRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkWritableApprovalWorkflowDefinitionRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkWritableApprovalWorkflowDefinitionRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedBulkWritableApprovalWorkflowDefinitionRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedBulkWritableApprovalWorkflowDefinitionRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedBulkWritableApprovalWorkflowDefinitionRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedBulkWritableApprovalWorkflowDefinitionRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


