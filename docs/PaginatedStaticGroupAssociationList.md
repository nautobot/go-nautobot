# PaginatedStaticGroupAssociationList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** |  | 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 
**Results** | [**[]StaticGroupAssociation**](StaticGroupAssociation.md) |  | 

## Methods

### NewPaginatedStaticGroupAssociationList

`func NewPaginatedStaticGroupAssociationList(count int32, results []StaticGroupAssociation, ) *PaginatedStaticGroupAssociationList`

NewPaginatedStaticGroupAssociationList instantiates a new PaginatedStaticGroupAssociationList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedStaticGroupAssociationListWithDefaults

`func NewPaginatedStaticGroupAssociationListWithDefaults() *PaginatedStaticGroupAssociationList`

NewPaginatedStaticGroupAssociationListWithDefaults instantiates a new PaginatedStaticGroupAssociationList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PaginatedStaticGroupAssociationList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedStaticGroupAssociationList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedStaticGroupAssociationList) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *PaginatedStaticGroupAssociationList) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedStaticGroupAssociationList) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedStaticGroupAssociationList) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedStaticGroupAssociationList) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedStaticGroupAssociationList) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedStaticGroupAssociationList) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedStaticGroupAssociationList) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedStaticGroupAssociationList) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedStaticGroupAssociationList) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedStaticGroupAssociationList) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedStaticGroupAssociationList) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedStaticGroupAssociationList) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
### GetResults

`func (o *PaginatedStaticGroupAssociationList) GetResults() []StaticGroupAssociation`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedStaticGroupAssociationList) GetResultsOk() (*[]StaticGroupAssociation, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedStaticGroupAssociationList) SetResults(v []StaticGroupAssociation)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


