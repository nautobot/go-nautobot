# CableTypeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Mapping** | Pointer to [**[]BulkWritableCableTypeRequestMappingInner**](BulkWritableCableTypeRequestMappingInner.md) | A-to-B lane mapping. Each entry represents one logical lane, giving its label and its (connector, position) coordinates on each side. If empty on write, it is auto-populated from a_connectors / b_connectors / total_lanes. | [optional] 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**PartNumber** | Pointer to **string** | Discrete part number (optional) | [optional] 
**HasEmbeddedTransceivers** | Pointer to **bool** | Indicates that this cable type has transceivers (e.g. SFP) built in. | [optional] 
**AConnectors** | Pointer to **int32** | Number of physical connectors on the A side. | [optional] 
**BConnectors** | Pointer to **int32** | Number of physical connectors on the B side. | [optional] 
**TotalLanes** | Pointer to **int32** | Total number of logical lanes in the breakout, distributed evenly across connectors on each side. | [optional] 
**IsShuffle** | Pointer to **bool** | Indicates non-linear (polarity-shuffled) position mapping. Informational only. | [optional] 
**StrandsPerLane** | Pointer to **int32** | Number of physical strands per logical lane (e.g. 1 for copper, 2 for duplex fiber). | [optional] 
**PolarityMethod** | Pointer to [**BulkWritableCableTypeRequestPolarityMethod**](BulkWritableCableTypeRequestPolarityMethod.md) |  | [optional] 
**Manufacturer** | Pointer to [**NullableApprovalWorkflowUser**](ApprovalWorkflowUser.md) |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** | Custom field data for this object, keyed by each applicable Custom Field&#39;s &#x60;key&#x60;. Value types vary with the custom field&#39;s type (text, integer, boolean, date, URL, JSON, select, multi-select); undefined values are &#x60;null&#x60;. On write, the payload is merged with existing values (PATCH-style: keys omitted from the payload are left untouched), and keys that do not correspond to a defined custom field are ignored. | [optional] 
**Relationships** | Pointer to [**map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue**](ApprovalWorkflowDefinitionRequestRelationshipsValue.md) |  | [optional] 
**Tags** | Pointer to [**[]BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | [optional] 

## Methods

### NewCableTypeRequest

`func NewCableTypeRequest(name string, ) *CableTypeRequest`

NewCableTypeRequest instantiates a new CableTypeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCableTypeRequestWithDefaults

`func NewCableTypeRequestWithDefaults() *CableTypeRequest`

NewCableTypeRequestWithDefaults instantiates a new CableTypeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CableTypeRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CableTypeRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CableTypeRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CableTypeRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMapping

`func (o *CableTypeRequest) GetMapping() []BulkWritableCableTypeRequestMappingInner`

GetMapping returns the Mapping field if non-nil, zero value otherwise.

### GetMappingOk

`func (o *CableTypeRequest) GetMappingOk() (*[]BulkWritableCableTypeRequestMappingInner, bool)`

GetMappingOk returns a tuple with the Mapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapping

`func (o *CableTypeRequest) SetMapping(v []BulkWritableCableTypeRequestMappingInner)`

SetMapping sets Mapping field to given value.

### HasMapping

`func (o *CableTypeRequest) HasMapping() bool`

HasMapping returns a boolean if a field has been set.

### GetName

`func (o *CableTypeRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CableTypeRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CableTypeRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CableTypeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CableTypeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CableTypeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CableTypeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPartNumber

`func (o *CableTypeRequest) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *CableTypeRequest) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *CableTypeRequest) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *CableTypeRequest) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetHasEmbeddedTransceivers

`func (o *CableTypeRequest) GetHasEmbeddedTransceivers() bool`

GetHasEmbeddedTransceivers returns the HasEmbeddedTransceivers field if non-nil, zero value otherwise.

### GetHasEmbeddedTransceiversOk

`func (o *CableTypeRequest) GetHasEmbeddedTransceiversOk() (*bool, bool)`

GetHasEmbeddedTransceiversOk returns a tuple with the HasEmbeddedTransceivers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasEmbeddedTransceivers

`func (o *CableTypeRequest) SetHasEmbeddedTransceivers(v bool)`

SetHasEmbeddedTransceivers sets HasEmbeddedTransceivers field to given value.

### HasHasEmbeddedTransceivers

`func (o *CableTypeRequest) HasHasEmbeddedTransceivers() bool`

HasHasEmbeddedTransceivers returns a boolean if a field has been set.

### GetAConnectors

`func (o *CableTypeRequest) GetAConnectors() int32`

GetAConnectors returns the AConnectors field if non-nil, zero value otherwise.

### GetAConnectorsOk

`func (o *CableTypeRequest) GetAConnectorsOk() (*int32, bool)`

GetAConnectorsOk returns a tuple with the AConnectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAConnectors

`func (o *CableTypeRequest) SetAConnectors(v int32)`

SetAConnectors sets AConnectors field to given value.

### HasAConnectors

`func (o *CableTypeRequest) HasAConnectors() bool`

HasAConnectors returns a boolean if a field has been set.

### GetBConnectors

`func (o *CableTypeRequest) GetBConnectors() int32`

GetBConnectors returns the BConnectors field if non-nil, zero value otherwise.

### GetBConnectorsOk

`func (o *CableTypeRequest) GetBConnectorsOk() (*int32, bool)`

GetBConnectorsOk returns a tuple with the BConnectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBConnectors

`func (o *CableTypeRequest) SetBConnectors(v int32)`

SetBConnectors sets BConnectors field to given value.

### HasBConnectors

`func (o *CableTypeRequest) HasBConnectors() bool`

HasBConnectors returns a boolean if a field has been set.

### GetTotalLanes

`func (o *CableTypeRequest) GetTotalLanes() int32`

GetTotalLanes returns the TotalLanes field if non-nil, zero value otherwise.

### GetTotalLanesOk

`func (o *CableTypeRequest) GetTotalLanesOk() (*int32, bool)`

GetTotalLanesOk returns a tuple with the TotalLanes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLanes

`func (o *CableTypeRequest) SetTotalLanes(v int32)`

SetTotalLanes sets TotalLanes field to given value.

### HasTotalLanes

`func (o *CableTypeRequest) HasTotalLanes() bool`

HasTotalLanes returns a boolean if a field has been set.

### GetIsShuffle

`func (o *CableTypeRequest) GetIsShuffle() bool`

GetIsShuffle returns the IsShuffle field if non-nil, zero value otherwise.

### GetIsShuffleOk

`func (o *CableTypeRequest) GetIsShuffleOk() (*bool, bool)`

GetIsShuffleOk returns a tuple with the IsShuffle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShuffle

`func (o *CableTypeRequest) SetIsShuffle(v bool)`

SetIsShuffle sets IsShuffle field to given value.

### HasIsShuffle

`func (o *CableTypeRequest) HasIsShuffle() bool`

HasIsShuffle returns a boolean if a field has been set.

### GetStrandsPerLane

`func (o *CableTypeRequest) GetStrandsPerLane() int32`

GetStrandsPerLane returns the StrandsPerLane field if non-nil, zero value otherwise.

### GetStrandsPerLaneOk

`func (o *CableTypeRequest) GetStrandsPerLaneOk() (*int32, bool)`

GetStrandsPerLaneOk returns a tuple with the StrandsPerLane field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrandsPerLane

`func (o *CableTypeRequest) SetStrandsPerLane(v int32)`

SetStrandsPerLane sets StrandsPerLane field to given value.

### HasStrandsPerLane

`func (o *CableTypeRequest) HasStrandsPerLane() bool`

HasStrandsPerLane returns a boolean if a field has been set.

### GetPolarityMethod

`func (o *CableTypeRequest) GetPolarityMethod() BulkWritableCableTypeRequestPolarityMethod`

GetPolarityMethod returns the PolarityMethod field if non-nil, zero value otherwise.

### GetPolarityMethodOk

`func (o *CableTypeRequest) GetPolarityMethodOk() (*BulkWritableCableTypeRequestPolarityMethod, bool)`

GetPolarityMethodOk returns a tuple with the PolarityMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolarityMethod

`func (o *CableTypeRequest) SetPolarityMethod(v BulkWritableCableTypeRequestPolarityMethod)`

SetPolarityMethod sets PolarityMethod field to given value.

### HasPolarityMethod

`func (o *CableTypeRequest) HasPolarityMethod() bool`

HasPolarityMethod returns a boolean if a field has been set.

### GetManufacturer

`func (o *CableTypeRequest) GetManufacturer() ApprovalWorkflowUser`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *CableTypeRequest) GetManufacturerOk() (*ApprovalWorkflowUser, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *CableTypeRequest) SetManufacturer(v ApprovalWorkflowUser)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *CableTypeRequest) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### SetManufacturerNil

`func (o *CableTypeRequest) SetManufacturerNil(b bool)`

 SetManufacturerNil sets the value for Manufacturer to be an explicit nil

### UnsetManufacturer
`func (o *CableTypeRequest) UnsetManufacturer()`

UnsetManufacturer ensures that no value is present for Manufacturer, not even an explicit nil
### GetCustomFields

`func (o *CableTypeRequest) GetCustomFields() map[string]*interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CableTypeRequest) GetCustomFieldsOk() (*map[string]*interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CableTypeRequest) SetCustomFields(v map[string]*interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CableTypeRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *CableTypeRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CableTypeRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CableTypeRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *CableTypeRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *CableTypeRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CableTypeRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CableTypeRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CableTypeRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


