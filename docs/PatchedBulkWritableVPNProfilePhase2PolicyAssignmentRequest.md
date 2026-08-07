# PatchedBulkWritableVPNProfilePhase2PolicyAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Weight** | Pointer to **int32** | Higher weights appear later in the list | [optional] 
**VpnProfile** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**VpnPhase2Policy** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewPatchedBulkWritableVPNProfilePhase2PolicyAssignmentRequest

`func NewPatchedBulkWritableVPNProfilePhase2PolicyAssignmentRequest(id string, ) *PatchedBulkWritableVPNProfilePhase2PolicyAssignmentRequest`

NewPatchedBulkWritableVPNProfilePhase2PolicyAssignmentRequest instantiates a new PatchedBulkWritableVPNProfilePhase2PolicyAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableVPNProfilePhase2PolicyAssignmentRequestWithDefaults

`func NewPatchedBulkWritableVPNProfilePhase2PolicyAssignmentRequestWithDefaults() *PatchedBulkWritableVPNProfilePhase2PolicyAssignmentRequest`

NewPatchedBulkWritableVPNProfilePhase2PolicyAssignmentRequestWithDefaults instantiates a new PatchedBulkWritableVPNProfilePhase2PolicyAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableVPNProfilePhase2PolicyAssignmentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableVPNProfilePhase2PolicyAssignmentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableVPNProfilePhase2PolicyAssignmentRequest) SetId(v string)`

SetId sets Id field to given value.


### GetWeight

`func (o *PatchedBulkWritableVPNProfilePhase2PolicyAssignmentRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *PatchedBulkWritableVPNProfilePhase2PolicyAssignmentRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *PatchedBulkWritableVPNProfilePhase2PolicyAssignmentRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *PatchedBulkWritableVPNProfilePhase2PolicyAssignmentRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetVpnProfile

`func (o *PatchedBulkWritableVPNProfilePhase2PolicyAssignmentRequest) GetVpnProfile() BulkWritableCableRequestStatus`

GetVpnProfile returns the VpnProfile field if non-nil, zero value otherwise.

### GetVpnProfileOk

`func (o *PatchedBulkWritableVPNProfilePhase2PolicyAssignmentRequest) GetVpnProfileOk() (*BulkWritableCableRequestStatus, bool)`

GetVpnProfileOk returns a tuple with the VpnProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnProfile

`func (o *PatchedBulkWritableVPNProfilePhase2PolicyAssignmentRequest) SetVpnProfile(v BulkWritableCableRequestStatus)`

SetVpnProfile sets VpnProfile field to given value.

### HasVpnProfile

`func (o *PatchedBulkWritableVPNProfilePhase2PolicyAssignmentRequest) HasVpnProfile() bool`

HasVpnProfile returns a boolean if a field has been set.

### GetVpnPhase2Policy

`func (o *PatchedBulkWritableVPNProfilePhase2PolicyAssignmentRequest) GetVpnPhase2Policy() BulkWritableCableRequestStatus`

GetVpnPhase2Policy returns the VpnPhase2Policy field if non-nil, zero value otherwise.

### GetVpnPhase2PolicyOk

`func (o *PatchedBulkWritableVPNProfilePhase2PolicyAssignmentRequest) GetVpnPhase2PolicyOk() (*BulkWritableCableRequestStatus, bool)`

GetVpnPhase2PolicyOk returns a tuple with the VpnPhase2Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnPhase2Policy

`func (o *PatchedBulkWritableVPNProfilePhase2PolicyAssignmentRequest) SetVpnPhase2Policy(v BulkWritableCableRequestStatus)`

SetVpnPhase2Policy sets VpnPhase2Policy field to given value.

### HasVpnPhase2Policy

`func (o *PatchedBulkWritableVPNProfilePhase2PolicyAssignmentRequest) HasVpnPhase2Policy() bool`

HasVpnPhase2Policy returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedBulkWritableVPNProfilePhase2PolicyAssignmentRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedBulkWritableVPNProfilePhase2PolicyAssignmentRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedBulkWritableVPNProfilePhase2PolicyAssignmentRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedBulkWritableVPNProfilePhase2PolicyAssignmentRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedBulkWritableVPNProfilePhase2PolicyAssignmentRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedBulkWritableVPNProfilePhase2PolicyAssignmentRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedBulkWritableVPNProfilePhase2PolicyAssignmentRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedBulkWritableVPNProfilePhase2PolicyAssignmentRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


