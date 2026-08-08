# InterfaceVDCAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**VirtualDeviceContext** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 
**Interface** | [**ApprovalWorkflowStageResponseApprovalWorkflowStage**](ApprovalWorkflowStageResponseApprovalWorkflowStage.md) |  | 

## Methods

### NewInterfaceVDCAssignment

`func NewInterfaceVDCAssignment(objectType string, display string, url string, naturalSlug string, virtualDeviceContext ApprovalWorkflowStageResponseApprovalWorkflowStage, interface_ ApprovalWorkflowStageResponseApprovalWorkflowStage, ) *InterfaceVDCAssignment`

NewInterfaceVDCAssignment instantiates a new InterfaceVDCAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterfaceVDCAssignmentWithDefaults

`func NewInterfaceVDCAssignmentWithDefaults() *InterfaceVDCAssignment`

NewInterfaceVDCAssignmentWithDefaults instantiates a new InterfaceVDCAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InterfaceVDCAssignment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InterfaceVDCAssignment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InterfaceVDCAssignment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InterfaceVDCAssignment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectType

`func (o *InterfaceVDCAssignment) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *InterfaceVDCAssignment) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *InterfaceVDCAssignment) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *InterfaceVDCAssignment) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *InterfaceVDCAssignment) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *InterfaceVDCAssignment) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *InterfaceVDCAssignment) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InterfaceVDCAssignment) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InterfaceVDCAssignment) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *InterfaceVDCAssignment) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *InterfaceVDCAssignment) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *InterfaceVDCAssignment) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetVirtualDeviceContext

`func (o *InterfaceVDCAssignment) GetVirtualDeviceContext() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetVirtualDeviceContext returns the VirtualDeviceContext field if non-nil, zero value otherwise.

### GetVirtualDeviceContextOk

`func (o *InterfaceVDCAssignment) GetVirtualDeviceContextOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetVirtualDeviceContextOk returns a tuple with the VirtualDeviceContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDeviceContext

`func (o *InterfaceVDCAssignment) SetVirtualDeviceContext(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetVirtualDeviceContext sets VirtualDeviceContext field to given value.


### GetInterface

`func (o *InterfaceVDCAssignment) GetInterface() ApprovalWorkflowStageResponseApprovalWorkflowStage`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *InterfaceVDCAssignment) GetInterfaceOk() (*ApprovalWorkflowStageResponseApprovalWorkflowStage, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *InterfaceVDCAssignment) SetInterface(v ApprovalWorkflowStageResponseApprovalWorkflowStage)`

SetInterface sets Interface field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


