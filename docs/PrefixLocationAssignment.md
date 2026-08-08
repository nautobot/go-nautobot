# PrefixLocationAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**Prefix** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 
**Location** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 

## Methods

### NewPrefixLocationAssignment

`func NewPrefixLocationAssignment(objectType string, display string, url string, naturalSlug string, prefix ApprovalWorkflowStageResponseApprovalWorkflowStage, location ApprovalWorkflowStageResponseApprovalWorkflowStage, ) *PrefixLocationAssignment`

NewPrefixLocationAssignment instantiates a new PrefixLocationAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrefixLocationAssignmentWithDefaults

`func NewPrefixLocationAssignmentWithDefaults() *PrefixLocationAssignment`

NewPrefixLocationAssignmentWithDefaults instantiates a new PrefixLocationAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PrefixLocationAssignment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrefixLocationAssignment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrefixLocationAssignment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PrefixLocationAssignment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectType

`func (o *PrefixLocationAssignment) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *PrefixLocationAssignment) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *PrefixLocationAssignment) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *PrefixLocationAssignment) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *PrefixLocationAssignment) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *PrefixLocationAssignment) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *PrefixLocationAssignment) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PrefixLocationAssignment) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PrefixLocationAssignment) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *PrefixLocationAssignment) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *PrefixLocationAssignment) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *PrefixLocationAssignment) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetPrefix

`func (o *PrefixLocationAssignment) GetPrefix() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *PrefixLocationAssignment) GetPrefixOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *PrefixLocationAssignment) SetPrefix(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetPrefix sets Prefix field to given value.


### GetLocation

`func (o *PrefixLocationAssignment) GetLocation() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PrefixLocationAssignment) GetLocationOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PrefixLocationAssignment) SetLocation(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetLocation sets Location field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


