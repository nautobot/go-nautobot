# PatchedWritableRearPortTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** | Physical label | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**PortTypeChoices**](PortTypeChoices.md) |  | [optional] 
**Positions** | Pointer to **int32** |  | [optional] 
**DeviceType** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewPatchedWritableRearPortTemplateRequest

`func NewPatchedWritableRearPortTemplateRequest() *PatchedWritableRearPortTemplateRequest`

NewPatchedWritableRearPortTemplateRequest instantiates a new PatchedWritableRearPortTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableRearPortTemplateRequestWithDefaults

`func NewPatchedWritableRearPortTemplateRequestWithDefaults() *PatchedWritableRearPortTemplateRequest`

NewPatchedWritableRearPortTemplateRequestWithDefaults instantiates a new PatchedWritableRearPortTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedWritableRearPortTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedWritableRearPortTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedWritableRearPortTemplateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedWritableRearPortTemplateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabel

`func (o *PatchedWritableRearPortTemplateRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PatchedWritableRearPortTemplateRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PatchedWritableRearPortTemplateRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PatchedWritableRearPortTemplateRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedWritableRearPortTemplateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedWritableRearPortTemplateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedWritableRearPortTemplateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedWritableRearPortTemplateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *PatchedWritableRearPortTemplateRequest) GetType() PortTypeChoices`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedWritableRearPortTemplateRequest) GetTypeOk() (*PortTypeChoices, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedWritableRearPortTemplateRequest) SetType(v PortTypeChoices)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedWritableRearPortTemplateRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPositions

`func (o *PatchedWritableRearPortTemplateRequest) GetPositions() int32`

GetPositions returns the Positions field if non-nil, zero value otherwise.

### GetPositionsOk

`func (o *PatchedWritableRearPortTemplateRequest) GetPositionsOk() (*int32, bool)`

GetPositionsOk returns a tuple with the Positions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositions

`func (o *PatchedWritableRearPortTemplateRequest) SetPositions(v int32)`

SetPositions sets Positions field to given value.

### HasPositions

`func (o *PatchedWritableRearPortTemplateRequest) HasPositions() bool`

HasPositions returns a boolean if a field has been set.

### GetDeviceType

`func (o *PatchedWritableRearPortTemplateRequest) GetDeviceType() BulkWritableCableRequestStatus`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *PatchedWritableRearPortTemplateRequest) GetDeviceTypeOk() (*BulkWritableCableRequestStatus, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *PatchedWritableRearPortTemplateRequest) SetDeviceType(v BulkWritableCableRequestStatus)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *PatchedWritableRearPortTemplateRequest) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedWritableRearPortTemplateRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritableRearPortTemplateRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritableRearPortTemplateRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritableRearPortTemplateRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedWritableRearPortTemplateRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedWritableRearPortTemplateRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedWritableRearPortTemplateRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedWritableRearPortTemplateRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


