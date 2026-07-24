# VPNProfilePhase1PolicyAssignment

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
**VpnPhase1Policy** | [**BulkWritableCableRequestStatus**](BulkWritableCableRequestStatus.md) |  | 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewVPNProfilePhase1PolicyAssignment

`func NewVPNProfilePhase1PolicyAssignment(objectType string, display string, url string, naturalSlug string, notesUrl string, vpnProfile BulkWritableCableRequestStatus, vpnPhase1Policy BulkWritableCableRequestStatus, ) *VPNProfilePhase1PolicyAssignment`

NewVPNProfilePhase1PolicyAssignment instantiates a new VPNProfilePhase1PolicyAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVPNProfilePhase1PolicyAssignmentWithDefaults

`func NewVPNProfilePhase1PolicyAssignmentWithDefaults() *VPNProfilePhase1PolicyAssignment`

NewVPNProfilePhase1PolicyAssignmentWithDefaults instantiates a new VPNProfilePhase1PolicyAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VPNProfilePhase1PolicyAssignment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VPNProfilePhase1PolicyAssignment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VPNProfilePhase1PolicyAssignment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VPNProfilePhase1PolicyAssignment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObjectType

`func (o *VPNProfilePhase1PolicyAssignment) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *VPNProfilePhase1PolicyAssignment) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *VPNProfilePhase1PolicyAssignment) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetDisplay

`func (o *VPNProfilePhase1PolicyAssignment) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *VPNProfilePhase1PolicyAssignment) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *VPNProfilePhase1PolicyAssignment) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetUrl

`func (o *VPNProfilePhase1PolicyAssignment) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VPNProfilePhase1PolicyAssignment) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VPNProfilePhase1PolicyAssignment) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetNaturalSlug

`func (o *VPNProfilePhase1PolicyAssignment) GetNaturalSlug() string`

GetNaturalSlug returns the NaturalSlug field if non-nil, zero value otherwise.

### GetNaturalSlugOk

`func (o *VPNProfilePhase1PolicyAssignment) GetNaturalSlugOk() (*string, bool)`

GetNaturalSlugOk returns a tuple with the NaturalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaturalSlug

`func (o *VPNProfilePhase1PolicyAssignment) SetNaturalSlug(v string)`

SetNaturalSlug sets NaturalSlug field to given value.


### GetNotesUrl

`func (o *VPNProfilePhase1PolicyAssignment) GetNotesUrl() string`

GetNotesUrl returns the NotesUrl field if non-nil, zero value otherwise.

### GetNotesUrlOk

`func (o *VPNProfilePhase1PolicyAssignment) GetNotesUrlOk() (*string, bool)`

GetNotesUrlOk returns a tuple with the NotesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotesUrl

`func (o *VPNProfilePhase1PolicyAssignment) SetNotesUrl(v string)`

SetNotesUrl sets NotesUrl field to given value.


### GetWeight

`func (o *VPNProfilePhase1PolicyAssignment) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *VPNProfilePhase1PolicyAssignment) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *VPNProfilePhase1PolicyAssignment) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *VPNProfilePhase1PolicyAssignment) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetVpnProfile

`func (o *VPNProfilePhase1PolicyAssignment) GetVpnProfile() BulkWritableCableRequestStatus`

GetVpnProfile returns the VpnProfile field if non-nil, zero value otherwise.

### GetVpnProfileOk

`func (o *VPNProfilePhase1PolicyAssignment) GetVpnProfileOk() (*BulkWritableCableRequestStatus, bool)`

GetVpnProfileOk returns a tuple with the VpnProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnProfile

`func (o *VPNProfilePhase1PolicyAssignment) SetVpnProfile(v BulkWritableCableRequestStatus)`

SetVpnProfile sets VpnProfile field to given value.


### GetVpnPhase1Policy

`func (o *VPNProfilePhase1PolicyAssignment) GetVpnPhase1Policy() BulkWritableCableRequestStatus`

GetVpnPhase1Policy returns the VpnPhase1Policy field if non-nil, zero value otherwise.

### GetVpnPhase1PolicyOk

`func (o *VPNProfilePhase1PolicyAssignment) GetVpnPhase1PolicyOk() (*BulkWritableCableRequestStatus, bool)`

GetVpnPhase1PolicyOk returns a tuple with the VpnPhase1Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnPhase1Policy

`func (o *VPNProfilePhase1PolicyAssignment) SetVpnPhase1Policy(v BulkWritableCableRequestStatus)`

SetVpnPhase1Policy sets VpnPhase1Policy field to given value.


### GetCustomFields

`func (o *VPNProfilePhase1PolicyAssignment) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *VPNProfilePhase1PolicyAssignment) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *VPNProfilePhase1PolicyAssignment) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *VPNProfilePhase1PolicyAssignment) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


