# PaginatedSecretsGroupAssociationList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** |  | 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | [**[]SecretsGroupAssociation**](SecretsGroupAssociation.md) |  | 

## Methods

### NewPaginatedSecretsGroupAssociationList

`func NewPaginatedSecretsGroupAssociationList(count int32, results []SecretsGroupAssociation, ) *PaginatedSecretsGroupAssociationList`

NewPaginatedSecretsGroupAssociationList instantiates a new PaginatedSecretsGroupAssociationList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedSecretsGroupAssociationListWithDefaults

`func NewPaginatedSecretsGroupAssociationListWithDefaults() *PaginatedSecretsGroupAssociationList`

NewPaginatedSecretsGroupAssociationListWithDefaults instantiates a new PaginatedSecretsGroupAssociationList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedSecretsGroupAssociationList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedSecretsGroupAssociationList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedSecretsGroupAssociationList) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *PaginatedSecretsGroupAssociationList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedSecretsGroupAssociationList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedSecretsGroupAssociationList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedSecretsGroupAssociationList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedSecretsGroupAssociationList) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedSecretsGroupAssociationList) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedSecretsGroupAssociationList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedSecretsGroupAssociationList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedSecretsGroupAssociationList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedSecretsGroupAssociationList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedSecretsGroupAssociationList) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedSecretsGroupAssociationList) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginatedSecretsGroupAssociationList) GetResults() []SecretsGroupAssociation`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedSecretsGroupAssociationList) GetResultsOk() (*[]SecretsGroupAssociation, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedSecretsGroupAssociationList) SetResults(v []SecretsGroupAssociation)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


