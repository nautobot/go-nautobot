# RelationshipRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceType** | **string** |  | 
**DestinationType** | **string** |  | 
**Label** | **string** | Label of the relationship as displayed to users | 
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

### NewRelationshipRequest

`func NewRelationshipRequest(sourceType string, destinationType string, label string, ) *RelationshipRequest`

NewRelationshipRequest instantiates a new RelationshipRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationshipRequestWithDefaults

`func NewRelationshipRequestWithDefaults() *RelationshipRequest`

NewRelationshipRequestWithDefaults instantiates a new RelationshipRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceType

`func (o *RelationshipRequest) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *RelationshipRequest) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *RelationshipRequest) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.


### GetDestinationType

`func (o *RelationshipRequest) GetDestinationType() string`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *RelationshipRequest) GetDestinationTypeOk() (*string, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *RelationshipRequest) SetDestinationType(v string)`

SetDestinationType sets DestinationType field to given value.


### GetLabel

`func (o *RelationshipRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *RelationshipRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *RelationshipRequest) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetKey

`func (o *RelationshipRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *RelationshipRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *RelationshipRequest) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *RelationshipRequest) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetDescription

`func (o *RelationshipRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RelationshipRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RelationshipRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RelationshipRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *RelationshipRequest) GetType() RelationshipTypeChoices`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RelationshipRequest) GetTypeOk() (*RelationshipTypeChoices, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RelationshipRequest) SetType(v RelationshipTypeChoices)`

SetType sets Type field to given value.

### HasType

`func (o *RelationshipRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRequiredOn

`func (o *RelationshipRequest) GetRequiredOn() BulkWritableRelationshipRequestRequiredOn`

GetRequiredOn returns the RequiredOn field if non-nil, zero value otherwise.

### GetRequiredOnOk

`func (o *RelationshipRequest) GetRequiredOnOk() (*BulkWritableRelationshipRequestRequiredOn, bool)`

GetRequiredOnOk returns a tuple with the RequiredOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredOn

`func (o *RelationshipRequest) SetRequiredOn(v BulkWritableRelationshipRequestRequiredOn)`

SetRequiredOn sets RequiredOn field to given value.

### HasRequiredOn

`func (o *RelationshipRequest) HasRequiredOn() bool`

HasRequiredOn returns a boolean if a field has been set.

### GetSourceLabel

`func (o *RelationshipRequest) GetSourceLabel() string`

GetSourceLabel returns the SourceLabel field if non-nil, zero value otherwise.

### GetSourceLabelOk

`func (o *RelationshipRequest) GetSourceLabelOk() (*string, bool)`

GetSourceLabelOk returns a tuple with the SourceLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLabel

`func (o *RelationshipRequest) SetSourceLabel(v string)`

SetSourceLabel sets SourceLabel field to given value.

### HasSourceLabel

`func (o *RelationshipRequest) HasSourceLabel() bool`

HasSourceLabel returns a boolean if a field has been set.

### GetSourceHidden

`func (o *RelationshipRequest) GetSourceHidden() bool`

GetSourceHidden returns the SourceHidden field if non-nil, zero value otherwise.

### GetSourceHiddenOk

`func (o *RelationshipRequest) GetSourceHiddenOk() (*bool, bool)`

GetSourceHiddenOk returns a tuple with the SourceHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceHidden

`func (o *RelationshipRequest) SetSourceHidden(v bool)`

SetSourceHidden sets SourceHidden field to given value.

### HasSourceHidden

`func (o *RelationshipRequest) HasSourceHidden() bool`

HasSourceHidden returns a boolean if a field has been set.

### GetSourceFilter

`func (o *RelationshipRequest) GetSourceFilter() map[string]interface{}`

GetSourceFilter returns the SourceFilter field if non-nil, zero value otherwise.

### GetSourceFilterOk

`func (o *RelationshipRequest) GetSourceFilterOk() (*map[string]interface{}, bool)`

GetSourceFilterOk returns a tuple with the SourceFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFilter

`func (o *RelationshipRequest) SetSourceFilter(v map[string]interface{})`

SetSourceFilter sets SourceFilter field to given value.

### HasSourceFilter

`func (o *RelationshipRequest) HasSourceFilter() bool`

HasSourceFilter returns a boolean if a field has been set.

### SetSourceFilterNil

`func (o *RelationshipRequest) SetSourceFilterNil(b bool)`

 SetSourceFilterNil sets the value for SourceFilter to be an explicit nil

### UnsetSourceFilter
`func (o *RelationshipRequest) UnsetSourceFilter()`

UnsetSourceFilter ensures that no value is present for SourceFilter, not even an explicit nil
### GetDestinationLabel

`func (o *RelationshipRequest) GetDestinationLabel() string`

GetDestinationLabel returns the DestinationLabel field if non-nil, zero value otherwise.

### GetDestinationLabelOk

`func (o *RelationshipRequest) GetDestinationLabelOk() (*string, bool)`

GetDestinationLabelOk returns a tuple with the DestinationLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationLabel

`func (o *RelationshipRequest) SetDestinationLabel(v string)`

SetDestinationLabel sets DestinationLabel field to given value.

### HasDestinationLabel

`func (o *RelationshipRequest) HasDestinationLabel() bool`

HasDestinationLabel returns a boolean if a field has been set.

### GetDestinationHidden

`func (o *RelationshipRequest) GetDestinationHidden() bool`

GetDestinationHidden returns the DestinationHidden field if non-nil, zero value otherwise.

### GetDestinationHiddenOk

`func (o *RelationshipRequest) GetDestinationHiddenOk() (*bool, bool)`

GetDestinationHiddenOk returns a tuple with the DestinationHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationHidden

`func (o *RelationshipRequest) SetDestinationHidden(v bool)`

SetDestinationHidden sets DestinationHidden field to given value.

### HasDestinationHidden

`func (o *RelationshipRequest) HasDestinationHidden() bool`

HasDestinationHidden returns a boolean if a field has been set.

### GetDestinationFilter

`func (o *RelationshipRequest) GetDestinationFilter() map[string]interface{}`

GetDestinationFilter returns the DestinationFilter field if non-nil, zero value otherwise.

### GetDestinationFilterOk

`func (o *RelationshipRequest) GetDestinationFilterOk() (*map[string]interface{}, bool)`

GetDestinationFilterOk returns a tuple with the DestinationFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationFilter

`func (o *RelationshipRequest) SetDestinationFilter(v map[string]interface{})`

SetDestinationFilter sets DestinationFilter field to given value.

### HasDestinationFilter

`func (o *RelationshipRequest) HasDestinationFilter() bool`

HasDestinationFilter returns a boolean if a field has been set.

### SetDestinationFilterNil

`func (o *RelationshipRequest) SetDestinationFilterNil(b bool)`

 SetDestinationFilterNil sets the value for DestinationFilter to be an explicit nil

### UnsetDestinationFilter
`func (o *RelationshipRequest) UnsetDestinationFilter()`

UnsetDestinationFilter ensures that no value is present for DestinationFilter, not even an explicit nil
### GetAdvancedUi

`func (o *RelationshipRequest) GetAdvancedUi() bool`

GetAdvancedUi returns the AdvancedUi field if non-nil, zero value otherwise.

### GetAdvancedUiOk

`func (o *RelationshipRequest) GetAdvancedUiOk() (*bool, bool)`

GetAdvancedUiOk returns a tuple with the AdvancedUi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedUi

`func (o *RelationshipRequest) SetAdvancedUi(v bool)`

SetAdvancedUi sets AdvancedUi field to given value.

### HasAdvancedUi

`func (o *RelationshipRequest) HasAdvancedUi() bool`

HasAdvancedUi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


