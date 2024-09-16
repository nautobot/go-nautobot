# StatusRetrieve200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DjangoVersion** | Pointer to **string** |  | [optional] 
**InstalledApps** | Pointer to **map[string]interface{}** |  | [optional] 
**NautobotVersion** | Pointer to **string** |  | [optional] 
**NautobotApps** | Pointer to **map[string]interface{}** |  | [optional] 
**Plugins** | Pointer to **map[string]interface{}** |  | [optional] 
**PythonVersion** | Pointer to **string** |  | [optional] 
**CeleryWorkersRunning** | Pointer to **int32** |  | [optional] 

## Methods

### NewStatusRetrieve200Response

`func NewStatusRetrieve200Response() *StatusRetrieve200Response`

NewStatusRetrieve200Response instantiates a new StatusRetrieve200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusRetrieve200ResponseWithDefaults

`func NewStatusRetrieve200ResponseWithDefaults() *StatusRetrieve200Response`

NewStatusRetrieve200ResponseWithDefaults instantiates a new StatusRetrieve200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDjangoVersion

`func (o *StatusRetrieve200Response) GetDjangoVersion() string`

GetDjangoVersion returns the DjangoVersion field if non-nil, zero value otherwise.

### GetDjangoVersionOk

`func (o *StatusRetrieve200Response) GetDjangoVersionOk() (*string, bool)`

GetDjangoVersionOk returns a tuple with the DjangoVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDjangoVersion

`func (o *StatusRetrieve200Response) SetDjangoVersion(v string)`

SetDjangoVersion sets DjangoVersion field to given value.

### HasDjangoVersion

`func (o *StatusRetrieve200Response) HasDjangoVersion() bool`

HasDjangoVersion returns a boolean if a field has been set.

### GetInstalledApps

`func (o *StatusRetrieve200Response) GetInstalledApps() map[string]interface{}`

GetInstalledApps returns the InstalledApps field if non-nil, zero value otherwise.

### GetInstalledAppsOk

`func (o *StatusRetrieve200Response) GetInstalledAppsOk() (*map[string]interface{}, bool)`

GetInstalledAppsOk returns a tuple with the InstalledApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledApps

`func (o *StatusRetrieve200Response) SetInstalledApps(v map[string]interface{})`

SetInstalledApps sets InstalledApps field to given value.

### HasInstalledApps

`func (o *StatusRetrieve200Response) HasInstalledApps() bool`

HasInstalledApps returns a boolean if a field has been set.

### GetNautobotVersion

`func (o *StatusRetrieve200Response) GetNautobotVersion() string`

GetNautobotVersion returns the NautobotVersion field if non-nil, zero value otherwise.

### GetNautobotVersionOk

`func (o *StatusRetrieve200Response) GetNautobotVersionOk() (*string, bool)`

GetNautobotVersionOk returns a tuple with the NautobotVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNautobotVersion

`func (o *StatusRetrieve200Response) SetNautobotVersion(v string)`

SetNautobotVersion sets NautobotVersion field to given value.

### HasNautobotVersion

`func (o *StatusRetrieve200Response) HasNautobotVersion() bool`

HasNautobotVersion returns a boolean if a field has been set.

### GetNautobotApps

`func (o *StatusRetrieve200Response) GetNautobotApps() map[string]interface{}`

GetNautobotApps returns the NautobotApps field if non-nil, zero value otherwise.

### GetNautobotAppsOk

`func (o *StatusRetrieve200Response) GetNautobotAppsOk() (*map[string]interface{}, bool)`

GetNautobotAppsOk returns a tuple with the NautobotApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNautobotApps

`func (o *StatusRetrieve200Response) SetNautobotApps(v map[string]interface{})`

SetNautobotApps sets NautobotApps field to given value.

### HasNautobotApps

`func (o *StatusRetrieve200Response) HasNautobotApps() bool`

HasNautobotApps returns a boolean if a field has been set.

### GetPlugins

`func (o *StatusRetrieve200Response) GetPlugins() map[string]interface{}`

GetPlugins returns the Plugins field if non-nil, zero value otherwise.

### GetPluginsOk

`func (o *StatusRetrieve200Response) GetPluginsOk() (*map[string]interface{}, bool)`

GetPluginsOk returns a tuple with the Plugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugins

`func (o *StatusRetrieve200Response) SetPlugins(v map[string]interface{})`

SetPlugins sets Plugins field to given value.

### HasPlugins

`func (o *StatusRetrieve200Response) HasPlugins() bool`

HasPlugins returns a boolean if a field has been set.

### GetPythonVersion

`func (o *StatusRetrieve200Response) GetPythonVersion() string`

GetPythonVersion returns the PythonVersion field if non-nil, zero value otherwise.

### GetPythonVersionOk

`func (o *StatusRetrieve200Response) GetPythonVersionOk() (*string, bool)`

GetPythonVersionOk returns a tuple with the PythonVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPythonVersion

`func (o *StatusRetrieve200Response) SetPythonVersion(v string)`

SetPythonVersion sets PythonVersion field to given value.

### HasPythonVersion

`func (o *StatusRetrieve200Response) HasPythonVersion() bool`

HasPythonVersion returns a boolean if a field has been set.

### GetCeleryWorkersRunning

`func (o *StatusRetrieve200Response) GetCeleryWorkersRunning() int32`

GetCeleryWorkersRunning returns the CeleryWorkersRunning field if non-nil, zero value otherwise.

### GetCeleryWorkersRunningOk

`func (o *StatusRetrieve200Response) GetCeleryWorkersRunningOk() (*int32, bool)`

GetCeleryWorkersRunningOk returns a tuple with the CeleryWorkersRunning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCeleryWorkersRunning

`func (o *StatusRetrieve200Response) SetCeleryWorkersRunning(v int32)`

SetCeleryWorkersRunning sets CeleryWorkersRunning field to given value.

### HasCeleryWorkersRunning

`func (o *StatusRetrieve200Response) HasCeleryWorkersRunning() bool`

HasCeleryWorkersRunning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


