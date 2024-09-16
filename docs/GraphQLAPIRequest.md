# GraphQLAPIRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | **string** | GraphQL query | 
**Variables** | Pointer to **interface{}** | Variables in JSON Format | [optional] 

## Methods

### NewGraphQLAPIRequest

`func NewGraphQLAPIRequest(query string, ) *GraphQLAPIRequest`

NewGraphQLAPIRequest instantiates a new GraphQLAPIRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphQLAPIRequestWithDefaults

`func NewGraphQLAPIRequestWithDefaults() *GraphQLAPIRequest`

NewGraphQLAPIRequestWithDefaults instantiates a new GraphQLAPIRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *GraphQLAPIRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *GraphQLAPIRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *GraphQLAPIRequest) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetVariables

`func (o *GraphQLAPIRequest) GetVariables() interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *GraphQLAPIRequest) GetVariablesOk() (*interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *GraphQLAPIRequest) SetVariables(v interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *GraphQLAPIRequest) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### SetVariablesNil

`func (o *GraphQLAPIRequest) SetVariablesNil(b bool)`

 SetVariablesNil sets the value for Variables to be an explicit nil

### UnsetVariables
`func (o *GraphQLAPIRequest) UnsetVariables()`

UnsetVariables ensures that no value is present for Variables, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


