# VPNProfilePhase2PolicyAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ObjectType** | **string** |  | [readonly] 
**Display** | **string** | Human friendly display value | [readonly] 
**Url** | **string** |  | [readonly] 
**NaturalSlug** | **string** |  | [readonly] 
**NotesUrl** | **string** |  | [readonly] 
**Weight** | Pointer to **int32** | Higher weights appear later in the list | [optional] 
**VpnProfile** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**VpnPhase2Policy** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewVPNProfilePhase2PolicyAssignment

`func NewVPNProfilePhase2PolicyAssignment(objectType string, display string, url string, naturalSlug string, notesUrl string, vpnProfile BulkWritableCableRequestStatus, vpnPhase2Policy BulkWritableCableRequestStatus, ) *VPNProfilePhase2PolicyAssignment`

NewVPNProfilePhase2PolicyAssignment instantiates a new VPNProfilePhase2PolicyAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVPNProfilePhase2PolicyAssignmentWithDefaults

`func NewVPNProfilePhase2PolicyAssignmentWithDefaults() *VPNProfilePhase2PolicyAssignment`

NewVPNProfilePhase2PolicyAssignmentWithDefaults instantiates a new VPNProfilePhase2PolicyAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VPNProfilePhase2PolicyAssignment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VPNProfilePhase2PolicyAssignment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VPNProfilePhase2PolicyAssignment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VPNProfilePhase2PolicyAssignment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectType

`func (o *VPNProfilePhase2PolicyAssignment) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VPNProfilePhase2PolicyAssignment) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VPNProfilePhase2PolicyAssignment) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *VPNProfilePhase2PolicyAssignment) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *VPNProfilePhase2PolicyAssignment) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *VPNProfilePhase2PolicyAssignment) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *VPNProfilePhase2PolicyAssignment) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VPNProfilePhase2PolicyAssignment) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VPNProfilePhase2PolicyAssignment) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *VPNProfilePhase2PolicyAssignment) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *VPNProfilePhase2PolicyAssignment) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *VPNProfilePhase2PolicyAssignment) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetNotesUrl

`func (o *VPNProfilePhase2PolicyAssignment) GetNotesUrl() string`

GetNotesUrl returns the NotesUrl field if non-nil, zero value otherwise.

### GetNotesUrlOk

`func (o *VPNProfilePhase2PolicyAssignment) GetNotesUrlOk() (*string, bool)`

GetNotesUrlOk returns a tuple with the NotesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesUrl

`func (o *VPNProfilePhase2PolicyAssignment) SetNotesUrl(v string)`

SetNotesUrl sets NotesUrl field to given value.


### GetWeight

`func (o *VPNProfilePhase2PolicyAssignment) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *VPNProfilePhase2PolicyAssignment) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *VPNProfilePhase2PolicyAssignment) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *VPNProfilePhase2PolicyAssignment) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetVpnProfile

`func (o *VPNProfilePhase2PolicyAssignment) GetVpnProfile() BulkWritableCableRequestStatus`

GetVpnProfile returns the VpnProfile field if non-nil, zero value otherwise.

### GetVpnProfileOk

`func (o *VPNProfilePhase2PolicyAssignment) GetVpnProfileOk() (*BulkWritableCableRequestStatus, bool)`

GetVpnProfileOk returns a tuple with the VpnProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnProfile

`func (o *VPNProfilePhase2PolicyAssignment) SetVpnProfile(v BulkWritableCableRequestStatus)`

SetVpnProfile sets VpnProfile field to given value.


### GetVpnPhase2Policy

`func (o *VPNProfilePhase2PolicyAssignment) GetVpnPhase2Policy() BulkWritableCableRequestStatus`

GetVpnPhase2Policy returns the VpnPhase2Policy field if non-nil, zero value otherwise.

### GetVpnPhase2PolicyOk

`func (o *VPNProfilePhase2PolicyAssignment) GetVpnPhase2PolicyOk() (*BulkWritableCableRequestStatus, bool)`

GetVpnPhase2PolicyOk returns a tuple with the VpnPhase2Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnPhase2Policy

`func (o *VPNProfilePhase2PolicyAssignment) SetVpnPhase2Policy(v BulkWritableCableRequestStatus)`

SetVpnPhase2Policy sets VpnPhase2Policy field to given value.


### GetCustomFields

`func (o *VPNProfilePhase2PolicyAssignment) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *VPNProfilePhase2PolicyAssignment) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *VPNProfilePhase2PolicyAssignment) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *VPNProfilePhase2PolicyAssignment) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


