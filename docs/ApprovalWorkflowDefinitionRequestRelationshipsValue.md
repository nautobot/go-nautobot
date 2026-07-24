# ApprovalWorkflowDefinitionRequestRelationshipsValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Url** | **string** |  | [readonly] 
**Label** | **string** |  | [readonly] 
**Type** | **string** |  | [readonly] 
**Source** | Pointer to [**ApprovalWorkflowDefinitionRequestRelationshipsValueSource**](ApprovalWorkflowDefinitionRequestRelationshipsValueSource.md) |  | [optional] 
**Destination** | Pointer to [**ApprovalWorkflowDefinitionRequestRelationshipsValueSource**](ApprovalWorkflowDefinitionRequestRelationshipsValueSource.md) |  | [optional] 
**Peer** | Pointer to [**ApprovalWorkflowDefinitionRequestRelationshipsValueSource**](ApprovalWorkflowDefinitionRequestRelationshipsValueSource.md) |  | [optional] 

## Methods

### NewApprovalWorkflowDefinitionRequestRelationshipsValue

`func NewApprovalWorkflowDefinitionRequestRelationshipsValue(id string, url string, label string, type_ string, ) *ApprovalWorkflowDefinitionRequestRelationshipsValue`

NewApprovalWorkflowDefinitionRequestRelationshipsValue instantiates a new ApprovalWorkflowDefinitionRequestRelationshipsValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalWorkflowDefinitionRequestRelationshipsValueWithDefaults

`func NewApprovalWorkflowDefinitionRequestRelationshipsValueWithDefaults() *ApprovalWorkflowDefinitionRequestRelationshipsValue`

NewApprovalWorkflowDefinitionRequestRelationshipsValueWithDefaults instantiates a new ApprovalWorkflowDefinitionRequestRelationshipsValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApprovalWorkflowDefinitionRequestRelationshipsValue) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApprovalWorkflowDefinitionRequestRelationshipsValue) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApprovalWorkflowDefinitionRequestRelationshipsValue) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *ApprovalWorkflowDefinitionRequestRelationshipsValue) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ApprovalWorkflowDefinitionRequestRelationshipsValue) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ApprovalWorkflowDefinitionRequestRelationshipsValue) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetLabel

`func (o *ApprovalWorkflowDefinitionRequestRelationshipsValue) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ApprovalWorkflowDefinitionRequestRelationshipsValue) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ApprovalWorkflowDefinitionRequestRelationshipsValue) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetType

`func (o *ApprovalWorkflowDefinitionRequestRelationshipsValue) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApprovalWorkflowDefinitionRequestRelationshipsValue) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApprovalWorkflowDefinitionRequestRelationshipsValue) SetType(v string)`

SetType sets Type field to given value.


### GetSource

`func (o *ApprovalWorkflowDefinitionRequestRelationshipsValue) GetSource() ApprovalWorkflowDefinitionRequestRelationshipsValueSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ApprovalWorkflowDefinitionRequestRelationshipsValue) GetSourceOk() (*ApprovalWorkflowDefinitionRequestRelationshipsValueSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ApprovalWorkflowDefinitionRequestRelationshipsValue) SetSource(v ApprovalWorkflowDefinitionRequestRelationshipsValueSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *ApprovalWorkflowDefinitionRequestRelationshipsValue) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetDestination

`func (o *ApprovalWorkflowDefinitionRequestRelationshipsValue) GetDestination() ApprovalWorkflowDefinitionRequestRelationshipsValueSource`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *ApprovalWorkflowDefinitionRequestRelationshipsValue) GetDestinationOk() (*ApprovalWorkflowDefinitionRequestRelationshipsValueSource, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *ApprovalWorkflowDefinitionRequestRelationshipsValue) SetDestination(v ApprovalWorkflowDefinitionRequestRelationshipsValueSource)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *ApprovalWorkflowDefinitionRequestRelationshipsValue) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetPeer

`func (o *ApprovalWorkflowDefinitionRequestRelationshipsValue) GetPeer() ApprovalWorkflowDefinitionRequestRelationshipsValueSource`

GetPeer returns the Peer field if non-nil, zero value otherwise.

### GetPeerOk

`func (o *ApprovalWorkflowDefinitionRequestRelationshipsValue) GetPeerOk() (*ApprovalWorkflowDefinitionRequestRelationshipsValueSource, bool)`

GetPeerOk returns a tuple with the Peer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeer

`func (o *ApprovalWorkflowDefinitionRequestRelationshipsValue) SetPeer(v ApprovalWorkflowDefinitionRequestRelationshipsValueSource)`

SetPeer sets Peer field to given value.

### HasPeer

`func (o *ApprovalWorkflowDefinitionRequestRelationshipsValue) HasPeer() bool`

HasPeer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


