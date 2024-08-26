# BulkWritablePowerOutletTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | Pointer to [**PowerOutletTypeChoices**](PowerOutletTypeChoices.md) |  | [optional] 
**FeedLeg** | Pointer to [**FeedLegEnum**](FeedLegEnum.md) |  | [optional] 
**Name** | **string** |  | 
**Label** | Pointer to **string** | Physical label | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DeviceType** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**PowerPortTemplate** | Pointer to [**NullableBulkWritableCircuitRequestTenant**](BulkWritableCircuitRequestTenant.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Relationships** | Pointer to [**map[string]BulkWritableCableRequestRelationshipsValue**](BulkWritableCableRequestRelationshipsValue.md) |  | [optional] 

## Methods

### NewBulkWritablePowerOutletTemplateRequest

`func NewBulkWritablePowerOutletTemplateRequest(id string, name string, deviceType BulkWritableCableRequestStatus, ) *BulkWritablePowerOutletTemplateRequest`

NewBulkWritablePowerOutletTemplateRequest instantiates a new BulkWritablePowerOutletTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkWritablePowerOutletTemplateRequestWithDefaults

`func NewBulkWritablePowerOutletTemplateRequestWithDefaults() *BulkWritablePowerOutletTemplateRequest`

NewBulkWritablePowerOutletTemplateRequestWithDefaults instantiates a new BulkWritablePowerOutletTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkWritablePowerOutletTemplateRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkWritablePowerOutletTemplateRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkWritablePowerOutletTemplateRequest) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *BulkWritablePowerOutletTemplateRequest) GetType() PowerOutletTypeChoices`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BulkWritablePowerOutletTemplateRequest) GetTypeOk() (*PowerOutletTypeChoices, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BulkWritablePowerOutletTemplateRequest) SetType(v PowerOutletTypeChoices)`

SetType sets Type field to given value.

### HasType

`func (o *BulkWritablePowerOutletTemplateRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFeedLeg

`func (o *BulkWritablePowerOutletTemplateRequest) GetFeedLeg() FeedLegEnum`

GetFeedLeg returns the FeedLeg field if non-nil, zero value otherwise.

### GetFeedLegOk

`func (o *BulkWritablePowerOutletTemplateRequest) GetFeedLegOk() (*FeedLegEnum, bool)`

GetFeedLegOk returns a tuple with the FeedLeg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedLeg

`func (o *BulkWritablePowerOutletTemplateRequest) SetFeedLeg(v FeedLegEnum)`

SetFeedLeg sets FeedLeg field to given value.

### HasFeedLeg

`func (o *BulkWritablePowerOutletTemplateRequest) HasFeedLeg() bool`

HasFeedLeg returns a boolean if a field has been set.

### GetName

`func (o *BulkWritablePowerOutletTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BulkWritablePowerOutletTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BulkWritablePowerOutletTemplateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *BulkWritablePowerOutletTemplateRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *BulkWritablePowerOutletTemplateRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *BulkWritablePowerOutletTemplateRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *BulkWritablePowerOutletTemplateRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDescription

`func (o *BulkWritablePowerOutletTemplateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BulkWritablePowerOutletTemplateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BulkWritablePowerOutletTemplateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BulkWritablePowerOutletTemplateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDeviceType

`func (o *BulkWritablePowerOutletTemplateRequest) GetDeviceType() BulkWritableCableRequestStatus`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *BulkWritablePowerOutletTemplateRequest) GetDeviceTypeOk() (*BulkWritableCableRequestStatus, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *BulkWritablePowerOutletTemplateRequest) SetDeviceType(v BulkWritableCableRequestStatus)`

SetDeviceType sets DeviceType field to given value.


### GetPowerPortTemplate

`func (o *BulkWritablePowerOutletTemplateRequest) GetPowerPortTemplate() BulkWritableCircuitRequestTenant`

GetPowerPortTemplate returns the PowerPortTemplate field if non-nil, zero value otherwise.

### GetPowerPortTemplateOk

`func (o *BulkWritablePowerOutletTemplateRequest) GetPowerPortTemplateOk() (*BulkWritableCircuitRequestTenant, bool)`

GetPowerPortTemplateOk returns a tuple with the PowerPortTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerPortTemplate

`func (o *BulkWritablePowerOutletTemplateRequest) SetPowerPortTemplate(v BulkWritableCircuitRequestTenant)`

SetPowerPortTemplate sets PowerPortTemplate field to given value.

### HasPowerPortTemplate

`func (o *BulkWritablePowerOutletTemplateRequest) HasPowerPortTemplate() bool`

HasPowerPortTemplate returns a boolean if a field has been set.

### SetPowerPortTemplateNil

`func (o *BulkWritablePowerOutletTemplateRequest) SetPowerPortTemplateNil(b bool)`

 SetPowerPortTemplateNil sets the value for PowerPortTemplate to be an explicit nil

### UnsetPowerPortTemplate
`func (o *BulkWritablePowerOutletTemplateRequest) UnsetPowerPortTemplate()`

UnsetPowerPortTemplate ensures that no value is present for PowerPortTemplate, not even an explicit nil
### GetCustomFields

`func (o *BulkWritablePowerOutletTemplateRequest) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *BulkWritablePowerOutletTemplateRequest) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *BulkWritablePowerOutletTemplateRequest) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *BulkWritablePowerOutletTemplateRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *BulkWritablePowerOutletTemplateRequest) GetRelationships() map[string]BulkWritableCableRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BulkWritablePowerOutletTemplateRequest) GetRelationshipsOk() (*map[string]BulkWritableCableRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BulkWritablePowerOutletTemplateRequest) SetRelationships(v map[string]BulkWritableCableRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BulkWritablePowerOutletTemplateRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


