# PatchedWritableFrontPortTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** | Physical label | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**PortTypeChoices**](PortTypeChoices.md) |  | [optional] 
**RearPortPosition** | Pointer to **int32** |  | [optional] [default to 1]
**DeviceType** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**RearPortTemplate** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewPatchedWritableFrontPortTemplateRequest

`func NewPatchedWritableFrontPortTemplateRequest() *PatchedWritableFrontPortTemplateRequest`

NewPatchedWritableFrontPortTemplateRequest instantiates a new PatchedWritableFrontPortTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritableFrontPortTemplateRequestWithDefaults

`func NewPatchedWritableFrontPortTemplateRequestWithDefaults() *PatchedWritableFrontPortTemplateRequest`

NewPatchedWritableFrontPortTemplateRequestWithDefaults instantiates a new PatchedWritableFrontPortTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedWritableFrontPortTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedWritableFrontPortTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedWritableFrontPortTemplateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedWritableFrontPortTemplateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabel

`func (o *PatchedWritableFrontPortTemplateRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PatchedWritableFrontPortTemplateRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PatchedWritableFrontPortTemplateRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PatchedWritableFrontPortTemplateRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedWritableFrontPortTemplateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedWritableFrontPortTemplateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedWritableFrontPortTemplateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedWritableFrontPortTemplateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *PatchedWritableFrontPortTemplateRequest) GetType() PortTypeChoices`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedWritableFrontPortTemplateRequest) GetTypeOk() (*PortTypeChoices, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedWritableFrontPortTemplateRequest) SetType(v PortTypeChoices)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedWritableFrontPortTemplateRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRearPortPosition

`func (o *PatchedWritableFrontPortTemplateRequest) GetRearPortPosition() int32`

GetRearPortPosition returns the RearPortPosition field if non-nil, zero value otherwise.

### GetRearPortPositionOk

`func (o *PatchedWritableFrontPortTemplateRequest) GetRearPortPositionOk() (*int32, bool)`

GetRearPortPositionOk returns a tuple with the RearPortPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRearPortPosition

`func (o *PatchedWritableFrontPortTemplateRequest) SetRearPortPosition(v int32)`

SetRearPortPosition sets RearPortPosition field to given value.

### HasRearPortPosition

`func (o *PatchedWritableFrontPortTemplateRequest) HasRearPortPosition() bool`

HasRearPortPosition returns a boolean if a field has been set.

### GetDeviceType

`func (o *PatchedWritableFrontPortTemplateRequest) GetDeviceType() BulkWritableCableRequestStatus`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *PatchedWritableFrontPortTemplateRequest) GetDeviceTypeOk() (*BulkWritableCableRequestStatus, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *PatchedWritableFrontPortTemplateRequest) SetDeviceType(v BulkWritableCableRequestStatus)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *PatchedWritableFrontPortTemplateRequest) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetRearPortTemplate

`func (o *PatchedWritableFrontPortTemplateRequest) GetRearPortTemplate() BulkWritableCableRequestStatus`

GetRearPortTemplate returns the RearPortTemplate field if non-nil, zero value otherwise.

### GetRearPortTemplateOk

`func (o *PatchedWritableFrontPortTemplateRequest) GetRearPortTemplateOk() (*BulkWritableCableRequestStatus, bool)`

GetRearPortTemplateOk returns a tuple with the RearPortTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRearPortTemplate

`func (o *PatchedWritableFrontPortTemplateRequest) SetRearPortTemplate(v BulkWritableCableRequestStatus)`

SetRearPortTemplate sets RearPortTemplate field to given value.

### HasRearPortTemplate

`func (o *PatchedWritableFrontPortTemplateRequest) HasRearPortTemplate() bool`

HasRearPortTemplate returns a boolean if a field has been set.

### GetCustomFields

`func (o *PatchedWritableFrontPortTemplateRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritableFrontPortTemplateRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritableFrontPortTemplateRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritableFrontPortTemplateRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedWritableFrontPortTemplateRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedWritableFrontPortTemplateRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedWritableFrontPortTemplateRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedWritableFrontPortTemplateRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


