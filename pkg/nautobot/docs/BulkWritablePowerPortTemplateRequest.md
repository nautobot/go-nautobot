# BulkWritablePowerPortTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | Pointer to [**PowerPortTypeChoices**](PowerPortTypeChoices.md) |  | [optional] 
**Name** | **string** |  | 
**Label** | Pointer to **string** | Physical label | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**MaximumDraw** | Pointer to **NullableInt32** | Maximum power draw (watts) | [optional] 
**AllocatedDraw** | Pointer to **NullableInt32** | Allocated power draw (watts) | [optional] 
**DeviceType** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewBulkWritablePowerPortTemplateRequest

`func NewBulkWritablePowerPortTemplateRequest(id string, name string, deviceType BulkWritableCableRequestStatus, ) *BulkWritablePowerPortTemplateRequest`

NewBulkWritablePowerPortTemplateRequest instantiates a new BulkWritablePowerPortTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritablePowerPortTemplateRequestWithDefaults

`func NewBulkWritablePowerPortTemplateRequestWithDefaults() *BulkWritablePowerPortTemplateRequest`

NewBulkWritablePowerPortTemplateRequestWithDefaults instantiates a new BulkWritablePowerPortTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritablePowerPortTemplateRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritablePowerPortTemplateRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritablePowerPortTemplateRequest) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *BulkWritablePowerPortTemplateRequest) GetType() PowerPortTypeChoices`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BulkWritablePowerPortTemplateRequest) GetTypeOk() (*PowerPortTypeChoices, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BulkWritablePowerPortTemplateRequest) SetType(v PowerPortTypeChoices)`

SetType sets Type field to given value.

### HasType

`func (o *BulkWritablePowerPortTemplateRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *BulkWritablePowerPortTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkWritablePowerPortTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkWritablePowerPortTemplateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *BulkWritablePowerPortTemplateRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *BulkWritablePowerPortTemplateRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *BulkWritablePowerPortTemplateRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *BulkWritablePowerPortTemplateRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDescription

`func (o *BulkWritablePowerPortTemplateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkWritablePowerPortTemplateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkWritablePowerPortTemplateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkWritablePowerPortTemplateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMaximumDraw

`func (o *BulkWritablePowerPortTemplateRequest) GetMaximumDraw() int32`

GetMaximumDraw returns the MaximumDraw field if non-nil, zero value otherwise.

### GetMaximumDrawOk

`func (o *BulkWritablePowerPortTemplateRequest) GetMaximumDrawOk() (*int32, bool)`

GetMaximumDrawOk returns a tuple with the MaximumDraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumDraw

`func (o *BulkWritablePowerPortTemplateRequest) SetMaximumDraw(v int32)`

SetMaximumDraw sets MaximumDraw field to given value.

### HasMaximumDraw

`func (o *BulkWritablePowerPortTemplateRequest) HasMaximumDraw() bool`

HasMaximumDraw returns a boolean if a field has been set.

### SetMaximumDrawNil

`func (o *BulkWritablePowerPortTemplateRequest) SetMaximumDrawNil(b bool)`

 SetMaximumDrawNil sets the value for MaximumDraw to be an explicit nil

### UnsetMaximumDraw
`func (o *BulkWritablePowerPortTemplateRequest) UnsetMaximumDraw()`

UnsetMaximumDraw ensures that no value is present for MaximumDraw, not even an explicit nil
### GetAllocatedDraw

`func (o *BulkWritablePowerPortTemplateRequest) GetAllocatedDraw() int32`

GetAllocatedDraw returns the AllocatedDraw field if non-nil, zero value otherwise.

### GetAllocatedDrawOk

`func (o *BulkWritablePowerPortTemplateRequest) GetAllocatedDrawOk() (*int32, bool)`

GetAllocatedDrawOk returns a tuple with the AllocatedDraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedDraw

`func (o *BulkWritablePowerPortTemplateRequest) SetAllocatedDraw(v int32)`

SetAllocatedDraw sets AllocatedDraw field to given value.

### HasAllocatedDraw

`func (o *BulkWritablePowerPortTemplateRequest) HasAllocatedDraw() bool`

HasAllocatedDraw returns a boolean if a field has been set.

### SetAllocatedDrawNil

`func (o *BulkWritablePowerPortTemplateRequest) SetAllocatedDrawNil(b bool)`

 SetAllocatedDrawNil sets the value for AllocatedDraw to be an explicit nil

### UnsetAllocatedDraw
`func (o *BulkWritablePowerPortTemplateRequest) UnsetAllocatedDraw()`

UnsetAllocatedDraw ensures that no value is present for AllocatedDraw, not even an explicit nil
### GetDeviceType

`func (o *BulkWritablePowerPortTemplateRequest) GetDeviceType() BulkWritableCableRequestStatus`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *BulkWritablePowerPortTemplateRequest) GetDeviceTypeOk() (*BulkWritableCableRequestStatus, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *BulkWritablePowerPortTemplateRequest) SetDeviceType(v BulkWritableCableRequestStatus)`

SetDeviceType sets DeviceType field to given value.


### GetCustomFields

`func (o *BulkWritablePowerPortTemplateRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkWritablePowerPortTemplateRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkWritablePowerPortTemplateRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkWritablePowerPortTemplateRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *BulkWritablePowerPortTemplateRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BulkWritablePowerPortTemplateRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BulkWritablePowerPortTemplateRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BulkWritablePowerPortTemplateRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


