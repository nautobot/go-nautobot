# VLANLocationAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**Vlan** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**Location** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 

## Methods

### NewVLANLocationAssignment

`func NewVLANLocationAssignment(id string, objectType string, display string, url string, naturalSlug string, vlan BulkWritableCableRequestStatus, location BulkWritableCableRequestStatus, ) *VLANLocationAssignment`

NewVLANLocationAssignment instantiates a new VLANLocationAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVLANLocationAssignmentWithDefaults

`func NewVLANLocationAssignmentWithDefaults() *VLANLocationAssignment`

NewVLANLocationAssignmentWithDefaults instantiates a new VLANLocationAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VLANLocationAssignment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VLANLocationAssignment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VLANLocationAssignment) SetId(v string)`

SetId sets Id field to given value.


### GetObjectType

`func (o *VLANLocationAssignment) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VLANLocationAssignment) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VLANLocationAssignment) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *VLANLocationAssignment) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *VLANLocationAssignment) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *VLANLocationAssignment) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *VLANLocationAssignment) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VLANLocationAssignment) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VLANLocationAssignment) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *VLANLocationAssignment) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *VLANLocationAssignment) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *VLANLocationAssignment) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetVlan

`func (o *VLANLocationAssignment) GetVlan() BulkWritableCableRequestStatus`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *VLANLocationAssignment) GetVlanOk() (*BulkWritableCableRequestStatus, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *VLANLocationAssignment) SetVlan(v BulkWritableCableRequestStatus)`

SetVlan sets Vlan field to given value.


### GetLocation

`func (o *VLANLocationAssignment) GetLocation() BulkWritableCableRequestStatus`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *VLANLocationAssignment) GetLocationOk() (*BulkWritableCableRequestStatus, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *VLANLocationAssignment) SetLocation(v BulkWritableCableRequestStatus)`

SetLocation sets Location field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


