# ScheduledJobState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to [**ScheduledJobStateValue**](ScheduledJobStateValue.md) |  | [optional] 
**Label** | Pointer to [**ScheduledJobStateLabel**](ScheduledJobStateLabel.md) |  | [optional] 

## Methods

### NewScheduledJobState

`func NewScheduledJobState() *ScheduledJobState`

NewScheduledJobState instantiates a new ScheduledJobState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduledJobStateWithDefaults

`func NewScheduledJobStateWithDefaults() *ScheduledJobState`

NewScheduledJobStateWithDefaults instantiates a new ScheduledJobState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *ScheduledJobState) GetValue() ScheduledJobStateValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ScheduledJobState) GetValueOk() (*ScheduledJobStateValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ScheduledJobState) SetValue(v ScheduledJobStateValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *ScheduledJobState) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetLabel

`func (o *ScheduledJobState) GetLabel() ScheduledJobStateLabel`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ScheduledJobState) GetLabelOk() (*ScheduledJobStateLabel, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ScheduledJobState) SetLabel(v ScheduledJobStateLabel)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ScheduledJobState) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


