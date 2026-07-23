# ApprovalWorkflowStageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**User** | [**User**](User.md) |  | [readonly] 
**Comments** | Pointer to **string** | User comments to explain the decision that he/she made | [optional] 
**State** | [**ApprovalWorkflowStateChoices**](ApprovalWorkflowStateChoices.md) | User response to this approval workflow stage instance. Eligible values are: Pending, Comment, Approved, Denied. | [readonly] 
**LastUpdated** | **NullableTime** |  | [readonly] 

## Methods

### NewApprovalWorkflowStageResponse

`func NewApprovalWorkflowStageResponse(objectType string, display string, user User, state ApprovalWorkflowStateChoices, lastUpdated NullableTime, ) *ApprovalWorkflowStageResponse`

NewApprovalWorkflowStageResponse instantiates a new ApprovalWorkflowStageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalWorkflowStageResponseWithDefaults

`func NewApprovalWorkflowStageResponseWithDefaults() *ApprovalWorkflowStageResponse`

NewApprovalWorkflowStageResponseWithDefaults instantiates a new ApprovalWorkflowStageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApprovalWorkflowStageResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApprovalWorkflowStageResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApprovalWorkflowStageResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApprovalWorkflowStageResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectType

`func (o *ApprovalWorkflowStageResponse) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ApprovalWorkflowStageResponse) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ApprovalWorkflowStageResponse) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *ApprovalWorkflowStageResponse) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *ApprovalWorkflowStageResponse) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *ApprovalWorkflowStageResponse) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUser

`func (o *ApprovalWorkflowStageResponse) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ApprovalWorkflowStageResponse) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ApprovalWorkflowStageResponse) SetUser(v User)`

SetUser sets User field to given value.


### GetComments

`func (o *ApprovalWorkflowStageResponse) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *ApprovalWorkflowStageResponse) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *ApprovalWorkflowStageResponse) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *ApprovalWorkflowStageResponse) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetState

`func (o *ApprovalWorkflowStageResponse) GetState() ApprovalWorkflowStateChoices`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ApprovalWorkflowStageResponse) GetStateOk() (*ApprovalWorkflowStateChoices, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ApprovalWorkflowStageResponse) SetState(v ApprovalWorkflowStateChoices)`

SetState sets State field to given value.


### GetLastUpdated

`func (o *ApprovalWorkflowStageResponse) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ApprovalWorkflowStageResponse) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ApprovalWorkflowStageResponse) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### SetLastUpdatedNil

`func (o *ApprovalWorkflowStageResponse) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *ApprovalWorkflowStageResponse) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


