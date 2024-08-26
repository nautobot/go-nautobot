# PatchedRelationshipRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceType** | Pointer to **string** |  | [optional] 
**DestinationType** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** | Label of the relationship as displayed to users | [optional] 
**Key** | Pointer to **string** | Internal relationship key. Please use underscores rather than dashes in this key. | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**RelationshipTypeChoices**](RelationshipTypeChoices.md) | Cardinality of this relationship | [optional] 
**RequiredOn** | Pointer to [**BulkWritableRelationshipRequestRequiredOn**](BulkWritableRelationshipRequestRequiredOn.md) |  | [optional] 
**SourceLabel** | Pointer to **string** | Label for related destination objects, as displayed on the source object. | [optional] 
**SourceHidden** | Pointer to **bool** | Hide this relationship on the source object. | [optional] 
**SourceFilter** | Pointer to **map[string]interface{}** | Filterset filter matching the applicable source objects of the selected type | [optional] 
**DestinationLabel** | Pointer to **string** | Label for related source objects, as displayed on the destination object. | [optional] 
**DestinationHidden** | Pointer to **bool** | Hide this relationship on the destination object. | [optional] 
**DestinationFilter** | Pointer to **map[string]interface{}** | Filterset filter matching the applicable destination objects of the selected type | [optional] 
**AdvancedUi** | Pointer to **bool** | Hide this field from the object&#39;s primary information tab. It will appear in the \&quot;Advanced\&quot; tab instead. | [optional] 

## Methods

### NewPatchedRelationshipRequest

`func NewPatchedRelationshipRequest() *PatchedRelationshipRequest`

NewPatchedRelationshipRequest instantiates a new PatchedRelationshipRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedRelationshipRequestWithDefaults

`func NewPatchedRelationshipRequestWithDefaults() *PatchedRelationshipRequest`

NewPatchedRelationshipRequestWithDefaults instantiates a new PatchedRelationshipRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceType

`func (o *PatchedRelationshipRequest) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *PatchedRelationshipRequest) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *PatchedRelationshipRequest) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *PatchedRelationshipRequest) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetDestinationType

`func (o *PatchedRelationshipRequest) GetDestinationType() string`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *PatchedRelationshipRequest) GetDestinationTypeOk() (*string, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *PatchedRelationshipRequest) SetDestinationType(v string)`

SetDestinationType sets DestinationType field to given value.

### HasDestinationType

`func (o *PatchedRelationshipRequest) HasDestinationType() bool`

HasDestinationType returns a boolean if a field has been set.

### GetLabel

`func (o *PatchedRelationshipRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PatchedRelationshipRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PatchedRelationshipRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PatchedRelationshipRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetKey

`func (o *PatchedRelationshipRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PatchedRelationshipRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PatchedRelationshipRequest) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *PatchedRelationshipRequest) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedRelationshipRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedRelationshipRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedRelationshipRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedRelationshipRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *PatchedRelationshipRequest) GetType() RelationshipTypeChoices`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedRelationshipRequest) GetTypeOk() (*RelationshipTypeChoices, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedRelationshipRequest) SetType(v RelationshipTypeChoices)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedRelationshipRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRequiredOn

`func (o *PatchedRelationshipRequest) GetRequiredOn() BulkWritableRelationshipRequestRequiredOn`

GetRequiredOn returns the RequiredOn field if non-nil, zero value otherwise.

### GetRequiredOnOk

`func (o *PatchedRelationshipRequest) GetRequiredOnOk() (*BulkWritableRelationshipRequestRequiredOn, bool)`

GetRequiredOnOk returns a tuple with the RequiredOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredOn

`func (o *PatchedRelationshipRequest) SetRequiredOn(v BulkWritableRelationshipRequestRequiredOn)`

SetRequiredOn sets RequiredOn field to given value.

### HasRequiredOn

`func (o *PatchedRelationshipRequest) HasRequiredOn() bool`

HasRequiredOn returns a boolean if a field has been set.

### GetSourceLabel

`func (o *PatchedRelationshipRequest) GetSourceLabel() string`

GetSourceLabel returns the SourceLabel field if non-nil, zero value otherwise.

### GetSourceLabelOk

`func (o *PatchedRelationshipRequest) GetSourceLabelOk() (*string, bool)`

GetSourceLabelOk returns a tuple with the SourceLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLabel

`func (o *PatchedRelationshipRequest) SetSourceLabel(v string)`

SetSourceLabel sets SourceLabel field to given value.

### HasSourceLabel

`func (o *PatchedRelationshipRequest) HasSourceLabel() bool`

HasSourceLabel returns a boolean if a field has been set.

### GetSourceHidden

`func (o *PatchedRelationshipRequest) GetSourceHidden() bool`

GetSourceHidden returns the SourceHidden field if non-nil, zero value otherwise.

### GetSourceHiddenOk

`func (o *PatchedRelationshipRequest) GetSourceHiddenOk() (*bool, bool)`

GetSourceHiddenOk returns a tuple with the SourceHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceHidden

`func (o *PatchedRelationshipRequest) SetSourceHidden(v bool)`

SetSourceHidden sets SourceHidden field to given value.

### HasSourceHidden

`func (o *PatchedRelationshipRequest) HasSourceHidden() bool`

HasSourceHidden returns a boolean if a field has been set.

### GetSourceFilter

`func (o *PatchedRelationshipRequest) GetSourceFilter() map[string]interface{}`

GetSourceFilter returns the SourceFilter field if non-nil, zero value otherwise.

### GetSourceFilterOk

`func (o *PatchedRelationshipRequest) GetSourceFilterOk() (*map[string]interface{}, bool)`

GetSourceFilterOk returns a tuple with the SourceFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFilter

`func (o *PatchedRelationshipRequest) SetSourceFilter(v map[string]interface{})`

SetSourceFilter sets SourceFilter field to given value.

### HasSourceFilter

`func (o *PatchedRelationshipRequest) HasSourceFilter() bool`

HasSourceFilter returns a boolean if a field has been set.

### SetSourceFilterNil

`func (o *PatchedRelationshipRequest) SetSourceFilterNil(b bool)`

 SetSourceFilterNil sets the value for SourceFilter to be an explicit nil

### UnsetSourceFilter
`func (o *PatchedRelationshipRequest) UnsetSourceFilter()`

UnsetSourceFilter ensures that no value is present for SourceFilter, not even an explicit nil
### GetDestinationLabel

`func (o *PatchedRelationshipRequest) GetDestinationLabel() string`

GetDestinationLabel returns the DestinationLabel field if non-nil, zero value otherwise.

### GetDestinationLabelOk

`func (o *PatchedRelationshipRequest) GetDestinationLabelOk() (*string, bool)`

GetDestinationLabelOk returns a tuple with the DestinationLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationLabel

`func (o *PatchedRelationshipRequest) SetDestinationLabel(v string)`

SetDestinationLabel sets DestinationLabel field to given value.

### HasDestinationLabel

`func (o *PatchedRelationshipRequest) HasDestinationLabel() bool`

HasDestinationLabel returns a boolean if a field has been set.

### GetDestinationHidden

`func (o *PatchedRelationshipRequest) GetDestinationHidden() bool`

GetDestinationHidden returns the DestinationHidden field if non-nil, zero value otherwise.

### GetDestinationHiddenOk

`func (o *PatchedRelationshipRequest) GetDestinationHiddenOk() (*bool, bool)`

GetDestinationHiddenOk returns a tuple with the DestinationHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationHidden

`func (o *PatchedRelationshipRequest) SetDestinationHidden(v bool)`

SetDestinationHidden sets DestinationHidden field to given value.

### HasDestinationHidden

`func (o *PatchedRelationshipRequest) HasDestinationHidden() bool`

HasDestinationHidden returns a boolean if a field has been set.

### GetDestinationFilter

`func (o *PatchedRelationshipRequest) GetDestinationFilter() map[string]interface{}`

GetDestinationFilter returns the DestinationFilter field if non-nil, zero value otherwise.

### GetDestinationFilterOk

`func (o *PatchedRelationshipRequest) GetDestinationFilterOk() (*map[string]interface{}, bool)`

GetDestinationFilterOk returns a tuple with the DestinationFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationFilter

`func (o *PatchedRelationshipRequest) SetDestinationFilter(v map[string]interface{})`

SetDestinationFilter sets DestinationFilter field to given value.

### HasDestinationFilter

`func (o *PatchedRelationshipRequest) HasDestinationFilter() bool`

HasDestinationFilter returns a boolean if a field has been set.

### SetDestinationFilterNil

`func (o *PatchedRelationshipRequest) SetDestinationFilterNil(b bool)`

 SetDestinationFilterNil sets the value for DestinationFilter to be an explicit nil

### UnsetDestinationFilter
`func (o *PatchedRelationshipRequest) UnsetDestinationFilter()`

UnsetDestinationFilter ensures that no value is present for DestinationFilter, not even an explicit nil
### GetAdvancedUi

`func (o *PatchedRelationshipRequest) GetAdvancedUi() bool`

GetAdvancedUi returns the AdvancedUi field if non-nil, zero value otherwise.

### GetAdvancedUiOk

`func (o *PatchedRelationshipRequest) GetAdvancedUiOk() (*bool, bool)`

GetAdvancedUiOk returns a tuple with the AdvancedUi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedUi

`func (o *PatchedRelationshipRequest) SetAdvancedUi(v bool)`

SetAdvancedUi sets AdvancedUi field to given value.

### HasAdvancedUi

`func (o *PatchedRelationshipRequest) HasAdvancedUi() bool`

HasAdvancedUi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


