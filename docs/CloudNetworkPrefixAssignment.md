# CloudNetworkPrefixAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**CloudNetwork** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 
**Prefix** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 

## Methods

### NewCloudNetworkPrefixAssignment

`func NewCloudNetworkPrefixAssignment(objectType string, display string, url string, naturalSlug string, cloudNetwork ApprovalWorkflowStageResponseApprovalWorkflowStage, prefix ApprovalWorkflowStageResponseApprovalWorkflowStage, ) *CloudNetworkPrefixAssignment`

NewCloudNetworkPrefixAssignment instantiates a new CloudNetworkPrefixAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudNetworkPrefixAssignmentWithDefaults

`func NewCloudNetworkPrefixAssignmentWithDefaults() *CloudNetworkPrefixAssignment`

NewCloudNetworkPrefixAssignmentWithDefaults instantiates a new CloudNetworkPrefixAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CloudNetworkPrefixAssignment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudNetworkPrefixAssignment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudNetworkPrefixAssignment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudNetworkPrefixAssignment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectType

`func (o *CloudNetworkPrefixAssignment) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *CloudNetworkPrefixAssignment) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *CloudNetworkPrefixAssignment) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *CloudNetworkPrefixAssignment) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *CloudNetworkPrefixAssignment) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *CloudNetworkPrefixAssignment) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *CloudNetworkPrefixAssignment) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CloudNetworkPrefixAssignment) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CloudNetworkPrefixAssignment) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *CloudNetworkPrefixAssignment) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *CloudNetworkPrefixAssignment) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *CloudNetworkPrefixAssignment) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetCloudNetwork

`func (o *CloudNetworkPrefixAssignment) GetCloudNetwork() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetCloudNetwork returns the CloudNetwork field if non-nil, zero value otherwise.

### GetCloudNetworkOk

`func (o *CloudNetworkPrefixAssignment) GetCloudNetworkOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetCloudNetworkOk returns a tuple with the CloudNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudNetwork

`func (o *CloudNetworkPrefixAssignment) SetCloudNetwork(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetCloudNetwork sets CloudNetwork field to given value.


### GetPrefix

`func (o *CloudNetworkPrefixAssignment) GetPrefix() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *CloudNetworkPrefixAssignment) GetPrefixOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *CloudNetworkPrefixAssignment) SetPrefix(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetPrefix sets Prefix field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


