# PatchedCableTypeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Mapping** | Pointer to [**[]BulkWritableCableTypeRequestMappingInner**](BulkWritableCableTypeRequestMappingInner.md) | A-to-B lane mapping. Each entry represents one logical lane, giving its label and its (connector, position) coordinates on each side. If empty on write, it is auto-populated from a_connectors / b_connectors / total_lanes. | [optional] 
**Name** | Pointer to **string** |  | [optional] 
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

### NewPatchedCableTypeRequest

`func NewPatchedCableTypeRequest() *PatchedCableTypeRequest`

NewPatchedCableTypeRequest instantiates a new PatchedCableTypeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchedCableTypeRequestWithDefaults

`func NewPatchedCableTypeRequestWithDefaults() *PatchedCableTypeRequest`

NewPatchedCableTypeRequestWithDefaults instantiates a new PatchedCableTypeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PatchedCableTypeRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PatchedCableTypeRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PatchedCableTypeRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PatchedCableTypeRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMapping

`func (o *PatchedCableTypeRequest) GetMapping() []BulkWritableCableTypeRequestMappingInner`

GetMapping returns the Mapping field if non-nil, zero value otherwise.

### GetMappingOk

`func (o *PatchedCableTypeRequest) GetMappingOk() (*[]BulkWritableCableTypeRequestMappingInner, bool)`

GetMappingOk returns a tuple with the Mapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapping

`func (o *PatchedCableTypeRequest) SetMapping(v []BulkWritableCableTypeRequestMappingInner)`

SetMapping sets Mapping field to given value.

### HasMapping

`func (o *PatchedCableTypeRequest) HasMapping() bool`

HasMapping returns a boolean if a field has been set.

### GetName

`func (o *PatchedCableTypeRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchedCableTypeRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchedCableTypeRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchedCableTypeRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PatchedCableTypeRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PatchedCableTypeRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PatchedCableTypeRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PatchedCableTypeRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPartNumber

`func (o *PatchedCableTypeRequest) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *PatchedCableTypeRequest) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *PatchedCableTypeRequest) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *PatchedCableTypeRequest) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetHasEmbeddedTransceivers

`func (o *PatchedCableTypeRequest) GetHasEmbeddedTransceivers() bool`

GetHasEmbeddedTransceivers returns the HasEmbeddedTransceivers field if non-nil, zero value otherwise.

### GetHasEmbeddedTransceiversOk

`func (o *PatchedCableTypeRequest) GetHasEmbeddedTransceiversOk() (*bool, bool)`

GetHasEmbeddedTransceiversOk returns a tuple with the HasEmbeddedTransceivers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasEmbeddedTransceivers

`func (o *PatchedCableTypeRequest) SetHasEmbeddedTransceivers(v bool)`

SetHasEmbeddedTransceivers sets HasEmbeddedTransceivers field to given value.

### HasHasEmbeddedTransceivers

`func (o *PatchedCableTypeRequest) HasHasEmbeddedTransceivers() bool`

HasHasEmbeddedTransceivers returns a boolean if a field has been set.

### GetAConnectors

`func (o *PatchedCableTypeRequest) GetAConnectors() int32`

GetAConnectors returns the AConnectors field if non-nil, zero value otherwise.

### GetAConnectorsOk

`func (o *PatchedCableTypeRequest) GetAConnectorsOk() (*int32, bool)`

GetAConnectorsOk returns a tuple with the AConnectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAConnectors

`func (o *PatchedCableTypeRequest) SetAConnectors(v int32)`

SetAConnectors sets AConnectors field to given value.

### HasAConnectors

`func (o *PatchedCableTypeRequest) HasAConnectors() bool`

HasAConnectors returns a boolean if a field has been set.

### GetBConnectors

`func (o *PatchedCableTypeRequest) GetBConnectors() int32`

GetBConnectors returns the BConnectors field if non-nil, zero value otherwise.

### GetBConnectorsOk

`func (o *PatchedCableTypeRequest) GetBConnectorsOk() (*int32, bool)`

GetBConnectorsOk returns a tuple with the BConnectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBConnectors

`func (o *PatchedCableTypeRequest) SetBConnectors(v int32)`

SetBConnectors sets BConnectors field to given value.

### HasBConnectors

`func (o *PatchedCableTypeRequest) HasBConnectors() bool`

HasBConnectors returns a boolean if a field has been set.

### GetTotalLanes

`func (o *PatchedCableTypeRequest) GetTotalLanes() int32`

GetTotalLanes returns the TotalLanes field if non-nil, zero value otherwise.

### GetTotalLanesOk

`func (o *PatchedCableTypeRequest) GetTotalLanesOk() (*int32, bool)`

GetTotalLanesOk returns a tuple with the TotalLanes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLanes

`func (o *PatchedCableTypeRequest) SetTotalLanes(v int32)`

SetTotalLanes sets TotalLanes field to given value.

### HasTotalLanes

`func (o *PatchedCableTypeRequest) HasTotalLanes() bool`

HasTotalLanes returns a boolean if a field has been set.

### GetIsShuffle

`func (o *PatchedCableTypeRequest) GetIsShuffle() bool`

GetIsShuffle returns the IsShuffle field if non-nil, zero value otherwise.

### GetIsShuffleOk

`func (o *PatchedCableTypeRequest) GetIsShuffleOk() (*bool, bool)`

GetIsShuffleOk returns a tuple with the IsShuffle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsShuffle

`func (o *PatchedCableTypeRequest) SetIsShuffle(v bool)`

SetIsShuffle sets IsShuffle field to given value.

### HasIsShuffle

`func (o *PatchedCableTypeRequest) HasIsShuffle() bool`

HasIsShuffle returns a boolean if a field has been set.

### GetStrandsPerLane

`func (o *PatchedCableTypeRequest) GetStrandsPerLane() int32`

GetStrandsPerLane returns the StrandsPerLane field if non-nil, zero value otherwise.

### GetStrandsPerLaneOk

`func (o *PatchedCableTypeRequest) GetStrandsPerLaneOk() (*int32, bool)`

GetStrandsPerLaneOk returns a tuple with the StrandsPerLane field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrandsPerLane

`func (o *PatchedCableTypeRequest) SetStrandsPerLane(v int32)`

SetStrandsPerLane sets StrandsPerLane field to given value.

### HasStrandsPerLane

`func (o *PatchedCableTypeRequest) HasStrandsPerLane() bool`

HasStrandsPerLane returns a boolean if a field has been set.

### GetPolarityMethod

`func (o *PatchedCableTypeRequest) GetPolarityMethod() BulkWritableCableTypeRequestPolarityMethod`

GetPolarityMethod returns the PolarityMethod field if non-nil, zero value otherwise.

### GetPolarityMethodOk

`func (o *PatchedCableTypeRequest) GetPolarityMethodOk() (*BulkWritableCableTypeRequestPolarityMethod, bool)`

GetPolarityMethodOk returns a tuple with the PolarityMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolarityMethod

`func (o *PatchedCableTypeRequest) SetPolarityMethod(v BulkWritableCableTypeRequestPolarityMethod)`

SetPolarityMethod sets PolarityMethod field to given value.

### HasPolarityMethod

`func (o *PatchedCableTypeRequest) HasPolarityMethod() bool`

HasPolarityMethod returns a boolean if a field has been set.

### GetManufacturer

`func (o *PatchedCableTypeRequest) GetManufacturer() ApprovalWorkflowUser`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *PatchedCableTypeRequest) GetManufacturerOk() (*ApprovalWorkflowUser, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *PatchedCableTypeRequest) SetManufacturer(v ApprovalWorkflowUser)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *PatchedCableTypeRequest) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### SetManufacturerNil

`func (o *PatchedCableTypeRequest) SetManufacturerNil(b bool)`

 SetManufacturerNil sets the value for Manufacturer to be an explicit nil

### UnsetManufacturer
`func (o *PatchedCableTypeRequest) UnsetManufacturer()`

UnsetManufacturer ensures that no value is present for Manufacturer, not even an explicit nil
### GetCustomFields

`func (o *PatchedCableTypeRequest) GetCustomFields() map[string]*interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *PatchedCableTypeRequest) GetCustomFieldsOk() (*map[string]*interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *PatchedCableTypeRequest) SetCustomFields(v map[string]*interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *PatchedCableTypeRequest) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetRelationships

`func (o *PatchedCableTypeRequest) GetRelationships() map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *PatchedCableTypeRequest) GetRelationshipsOk() (*map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *PatchedCableTypeRequest) SetRelationships(v map[string]ApprovalWorkflowDefinitionRequestRelationshipsValue)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *PatchedCableTypeRequest) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetTags

`func (o *PatchedCableTypeRequest) GetTags() []BulkWritableCableRequestStatus`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PatchedCableTypeRequest) GetTagsOk() (*[]BulkWritableCableRequestStatus, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PatchedCableTypeRequest) SetTags(v []BulkWritableCableRequestStatus)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PatchedCableTypeRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


