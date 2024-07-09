# PatchedWritablePowerOutletTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** | Physical label | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**PatchedWritablePowerOutletTemplateRequestType**](PatchedWritablePowerOutletTemplateRequestType.md) |  | [optional] 
**FeedLeg** | Pointer to [**PatchedWritablePowerOutletRequestFeedLeg**](PatchedWritablePowerOutletRequestFeedLeg.md) |  | [optional] 
**DeviceType** | Pointer to [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 
**PowerPortTemplate** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewPatchedWritablePowerOutletTemplateRequest

`func NewPatchedWritablePowerOutletTemplateRequest() *PatchedWritablePowerOutletTemplateRequest`

NewPatchedWritablePowerOutletTemplateRequest instantiates a new PatchedWritablePowerOutletTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedWritablePowerOutletTemplateRequestWithDefaults

`func NewPatchedWritablePowerOutletTemplateRequestWithDefaults() *PatchedWritablePowerOutletTemplateRequest`

NewPatchedWritablePowerOutletTemplateRequestWithDefaults instantiates a new PatchedWritablePowerOutletTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchedWritablePowerOutletTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedWritablePowerOutletTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedWritablePowerOutletTemplateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedWritablePowerOutletTemplateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabel

`func (o *PatchedWritablePowerOutletTemplateRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PatchedWritablePowerOutletTemplateRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PatchedWritablePowerOutletTemplateRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PatchedWritablePowerOutletTemplateRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedWritablePowerOutletTemplateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedWritablePowerOutletTemplateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedWritablePowerOutletTemplateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedWritablePowerOutletTemplateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *PatchedWritablePowerOutletTemplateRequest) GetType() PatchedWritablePowerOutletTemplateRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchedWritablePowerOutletTemplateRequest) GetTypeOk() (*PatchedWritablePowerOutletTemplateRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchedWritablePowerOutletTemplateRequest) SetType(v PatchedWritablePowerOutletTemplateRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *PatchedWritablePowerOutletTemplateRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFeedLeg

`func (o *PatchedWritablePowerOutletTemplateRequest) GetFeedLeg() PatchedWritablePowerOutletRequestFeedLeg`

GetFeedLeg returns the FeedLeg field if non-nil, zero value otherwise.

### GetFeedLegOk

`func (o *PatchedWritablePowerOutletTemplateRequest) GetFeedLegOk() (*PatchedWritablePowerOutletRequestFeedLeg, bool)`

GetFeedLegOk returns a tuple with the FeedLeg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedLeg

`func (o *PatchedWritablePowerOutletTemplateRequest) SetFeedLeg(v PatchedWritablePowerOutletRequestFeedLeg)`

SetFeedLeg sets FeedLeg field to given value.

### HasFeedLeg

`func (o *PatchedWritablePowerOutletTemplateRequest) HasFeedLeg() bool`

HasFeedLeg returns a boolean if a field has been set.

### GetDeviceType

`func (o *PatchedWritablePowerOutletTemplateRequest) GetDeviceType() BulkWritableCableRequestStatus`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *PatchedWritablePowerOutletTemplateRequest) GetDeviceTypeOk() (*BulkWritableCableRequestStatus, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *PatchedWritablePowerOutletTemplateRequest) SetDeviceType(v BulkWritableCableRequestStatus)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *PatchedWritablePowerOutletTemplateRequest) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetPowerPortTemplate

`func (o *PatchedWritablePowerOutletTemplateRequest) GetPowerPortTemplate() BulkWritableCircuitRequestTenant`

GetPowerPortTemplate returns the PowerPortTemplate field if non-nil, zero value otherwise.

### GetPowerPortTemplateOk

`func (o *PatchedWritablePowerOutletTemplateRequest) GetPowerPortTemplateOk() (*BulkWritableCircuitRequestTenant, bool)`

GetPowerPortTemplateOk returns a tuple with the PowerPortTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerPortTemplate

`func (o *PatchedWritablePowerOutletTemplateRequest) SetPowerPortTemplate(v BulkWritableCircuitRequestTenant)`

SetPowerPortTemplate sets PowerPortTemplate field to given value.

### HasPowerPortTemplate

`func (o *PatchedWritablePowerOutletTemplateRequest) HasPowerPortTemplate() bool`

HasPowerPortTemplate returns a boolean if a field has been set.

### SetPowerPortTemplateNil

`func (o *PatchedWritablePowerOutletTemplateRequest) SetPowerPortTemplateNil(b bool)`

 SetPowerPortTemplateNil sets the value for PowerPortTemplate to be an explicit nil

### UnsetPowerPortTemplate
`func (o *PatchedWritablePowerOutletTemplateRequest) UnsetPowerPortTemplate()`

UnsetPowerPortTemplate ensures that no value is present for PowerPortTemplate, not even an explicit nil
### GetCustomFields

`func (o *PatchedWritablePowerOutletTemplateRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedWritablePowerOutletTemplateRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedWritablePowerOutletTemplateRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedWritablePowerOutletTemplateRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedWritablePowerOutletTemplateRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedWritablePowerOutletTemplateRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedWritablePowerOutletTemplateRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedWritablePowerOutletTemplateRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


