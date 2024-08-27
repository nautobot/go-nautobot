# PatchedBulkWritableRelationshipRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**SourceType** | Pointer to **string** |  | [optional] 
**DestinationType** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** | Label of the relationship as displayed to users | [optional] 
**Key** | Pointer to **string** | Internal relationship key. Please use underscores rather than dashes in this key. | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**RelationshipTypeChoices**](RelationshipTypeChoices.md) | Cardinality of this relationship | [optional] 
**RequiredOn** | Pointer to [**BulkWritableRelationshipRequestRequiredOn**](BulkWritableRelationshipRequestRequiredOn.md) |  | [optional] 
**SourceLabel** | Pointer to **string** | Label for related destination objects, as displayed on the source object. | [optional] 
**SourceHidden** | Pointer to **bool** | Hide this relationship on the source object. | [optional] 
**SourceFilter** | Pointer to **interface{}** | Filterset filter matching the applicable source objects of the selected type | [optional] 
**DestinationLabel** | Pointer to **string** | Label for related source objects, as displayed on the destination object. | [optional] 
**DestinationHidden** | Pointer to **bool** | Hide this relationship on the destination object. | [optional] 
**DestinationFilter** | Pointer to **interface{}** | Filterset filter matching the applicable destination objects of the selected type | [optional] 
**AdvancedUi** | Pointer to **bool** | Hide this field from the object&#39;s primary information tab. It will appear in the \&quot;Advanced\&quot; tab instead. | [optional] 

## Methods

### NewPatchedBulkWritableRelationshipRequest

`func NewPatchedBulkWritableRelationshipRequest(id string, ) *PatchedBulkWritableRelationshipRequest`

NewPatchedBulkWritableRelationshipRequest instantiates a new PatchedBulkWritableRelationshipRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedBulkWritableRelationshipRequestWithDefaults

`func NewPatchedBulkWritableRelationshipRequestWithDefaults() *PatchedBulkWritableRelationshipRequest`

NewPatchedBulkWritableRelationshipRequestWithDefaults instantiates a new PatchedBulkWritableRelationshipRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedBulkWritableRelationshipRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedBulkWritableRelationshipRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedBulkWritableRelationshipRequest) SetId(v string)`

SetId sets Id field to given value.


### GetSourceType

`func (o *PatchedBulkWritableRelationshipRequest) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *PatchedBulkWritableRelationshipRequest) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *PatchedBulkWritableRelationshipRequest) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *PatchedBulkWritableRelationshipRequest) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetDestinationType

`func (o *PatchedBulkWritableRelationshipRequest) GetDestinationType() string`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *PatchedBulkWritableRelationshipRequest) GetDestinationTypeOk() (*string, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *PatchedBulkWritableRelationshipRequest) SetDestinationType(v string)`

SetDestinationType sets DestinationType field to given value.

### HasDestinationType

`func (o *PatchedBulkWritableRelationshipRequest) HasDestinationType() bool`

HasDestinationType returns a boolean if a field has been set.

### GetLabel

`func (o *PatchedBulkWritableRelationshipRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PatchedBulkWritableRelationshipRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PatchedBulkWritableRelationshipRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PatchedBulkWritableRelationshipRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetKey

`func (o *PatchedBulkWritableRelationshipRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PatchedBulkWritableRelationshipRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PatchedBulkWritableRelationshipRequest) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *PatchedBulkWritableRelationshipRequest) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedBulkWritableRelationshipRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedBulkWritableRelationshipRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedBulkWritableRelationshipRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedBulkWritableRelationshipRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *PatchedBulkWritableRelationshipRequest) GetType() RelationshipTypeChoices`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedBulkWritableRelationshipRequest) GetTypeOk() (*RelationshipTypeChoices, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedBulkWritableRelationshipRequest) SetType(v RelationshipTypeChoices)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedBulkWritableRelationshipRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRequiredOn

`func (o *PatchedBulkWritableRelationshipRequest) GetRequiredOn() BulkWritableRelationshipRequestRequiredOn`

GetRequiredOn returns the RequiredOn field if non-nil, zero value otherwise.

### GetRequiredOnOk

`func (o *PatchedBulkWritableRelationshipRequest) GetRequiredOnOk() (*BulkWritableRelationshipRequestRequiredOn, bool)`

GetRequiredOnOk returns a tuple with the RequiredOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredOn

`func (o *PatchedBulkWritableRelationshipRequest) SetRequiredOn(v BulkWritableRelationshipRequestRequiredOn)`

SetRequiredOn sets RequiredOn field to given value.

### HasRequiredOn

`func (o *PatchedBulkWritableRelationshipRequest) HasRequiredOn() bool`

HasRequiredOn returns a boolean if a field has been set.

### GetSourceLabel

`func (o *PatchedBulkWritableRelationshipRequest) GetSourceLabel() string`

GetSourceLabel returns the SourceLabel field if non-nil, zero value otherwise.

### GetSourceLabelOk

`func (o *PatchedBulkWritableRelationshipRequest) GetSourceLabelOk() (*string, bool)`

GetSourceLabelOk returns a tuple with the SourceLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLabel

`func (o *PatchedBulkWritableRelationshipRequest) SetSourceLabel(v string)`

SetSourceLabel sets SourceLabel field to given value.

### HasSourceLabel

`func (o *PatchedBulkWritableRelationshipRequest) HasSourceLabel() bool`

HasSourceLabel returns a boolean if a field has been set.

### GetSourceHidden

`func (o *PatchedBulkWritableRelationshipRequest) GetSourceHidden() bool`

GetSourceHidden returns the SourceHidden field if non-nil, zero value otherwise.

### GetSourceHiddenOk

`func (o *PatchedBulkWritableRelationshipRequest) GetSourceHiddenOk() (*bool, bool)`

GetSourceHiddenOk returns a tuple with the SourceHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceHidden

`func (o *PatchedBulkWritableRelationshipRequest) SetSourceHidden(v bool)`

SetSourceHidden sets SourceHidden field to given value.

### HasSourceHidden

`func (o *PatchedBulkWritableRelationshipRequest) HasSourceHidden() bool`

HasSourceHidden returns a boolean if a field has been set.

### GetSourceFilter

`func (o *PatchedBulkWritableRelationshipRequest) GetSourceFilter() interface{}`

GetSourceFilter returns the SourceFilter field if non-nil, zero value otherwise.

### GetSourceFilterOk

`func (o *PatchedBulkWritableRelationshipRequest) GetSourceFilterOk() (*interface{}, bool)`

GetSourceFilterOk returns a tuple with the SourceFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFilter

`func (o *PatchedBulkWritableRelationshipRequest) SetSourceFilter(v interface{})`

SetSourceFilter sets SourceFilter field to given value.

### HasSourceFilter

`func (o *PatchedBulkWritableRelationshipRequest) HasSourceFilter() bool`

HasSourceFilter returns a boolean if a field has been set.

### SetSourceFilterNil

`func (o *PatchedBulkWritableRelationshipRequest) SetSourceFilterNil(b bool)`

 SetSourceFilterNil sets the value for SourceFilter to be an explicit nil

### UnsetSourceFilter
`func (o *PatchedBulkWritableRelationshipRequest) UnsetSourceFilter()`

UnsetSourceFilter ensures that no value is present for SourceFilter, not even an explicit nil
### GetDestinationLabel

`func (o *PatchedBulkWritableRelationshipRequest) GetDestinationLabel() string`

GetDestinationLabel returns the DestinationLabel field if non-nil, zero value otherwise.

### GetDestinationLabelOk

`func (o *PatchedBulkWritableRelationshipRequest) GetDestinationLabelOk() (*string, bool)`

GetDestinationLabelOk returns a tuple with the DestinationLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationLabel

`func (o *PatchedBulkWritableRelationshipRequest) SetDestinationLabel(v string)`

SetDestinationLabel sets DestinationLabel field to given value.

### HasDestinationLabel

`func (o *PatchedBulkWritableRelationshipRequest) HasDestinationLabel() bool`

HasDestinationLabel returns a boolean if a field has been set.

### GetDestinationHidden

`func (o *PatchedBulkWritableRelationshipRequest) GetDestinationHidden() bool`

GetDestinationHidden returns the DestinationHidden field if non-nil, zero value otherwise.

### GetDestinationHiddenOk

`func (o *PatchedBulkWritableRelationshipRequest) GetDestinationHiddenOk() (*bool, bool)`

GetDestinationHiddenOk returns a tuple with the DestinationHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationHidden

`func (o *PatchedBulkWritableRelationshipRequest) SetDestinationHidden(v bool)`

SetDestinationHidden sets DestinationHidden field to given value.

### HasDestinationHidden

`func (o *PatchedBulkWritableRelationshipRequest) HasDestinationHidden() bool`

HasDestinationHidden returns a boolean if a field has been set.

### GetDestinationFilter

`func (o *PatchedBulkWritableRelationshipRequest) GetDestinationFilter() interface{}`

GetDestinationFilter returns the DestinationFilter field if non-nil, zero value otherwise.

### GetDestinationFilterOk

`func (o *PatchedBulkWritableRelationshipRequest) GetDestinationFilterOk() (*interface{}, bool)`

GetDestinationFilterOk returns a tuple with the DestinationFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationFilter

`func (o *PatchedBulkWritableRelationshipRequest) SetDestinationFilter(v interface{})`

SetDestinationFilter sets DestinationFilter field to given value.

### HasDestinationFilter

`func (o *PatchedBulkWritableRelationshipRequest) HasDestinationFilter() bool`

HasDestinationFilter returns a boolean if a field has been set.

### SetDestinationFilterNil

`func (o *PatchedBulkWritableRelationshipRequest) SetDestinationFilterNil(b bool)`

 SetDestinationFilterNil sets the value for DestinationFilter to be an explicit nil

### UnsetDestinationFilter
`func (o *PatchedBulkWritableRelationshipRequest) UnsetDestinationFilter()`

UnsetDestinationFilter ensures that no value is present for DestinationFilter, not even an explicit nil
### GetAdvancedUi

`func (o *PatchedBulkWritableRelationshipRequest) GetAdvancedUi() bool`

GetAdvancedUi returns the AdvancedUi field if non-nil, zero value otherwise.

### GetAdvancedUiOk

`func (o *PatchedBulkWritableRelationshipRequest) GetAdvancedUiOk() (*bool, bool)`

GetAdvancedUiOk returns a tuple with the AdvancedUi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedUi

`func (o *PatchedBulkWritableRelationshipRequest) SetAdvancedUi(v bool)`

SetAdvancedUi sets AdvancedUi field to given value.

### HasAdvancedUi

`func (o *PatchedBulkWritableRelationshipRequest) HasAdvancedUi() bool`

HasAdvancedUi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


