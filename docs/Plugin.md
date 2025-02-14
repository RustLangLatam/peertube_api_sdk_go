# Plugin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **int32** | - &#x60;1&#x60;: PLUGIN - &#x60;2&#x60;: THEME  | [optional] 
**LatestVersion** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Uninstalled** | Pointer to **bool** |  | [optional] 
**PeertubeEngine** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Homepage** | Pointer to **string** |  | [optional] 
**Settings** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewPlugin

`func NewPlugin() *Plugin`

NewPlugin instantiates a new Plugin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginWithDefaults

`func NewPluginWithDefaults() *Plugin`

NewPluginWithDefaults instantiates a new Plugin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Plugin) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Plugin) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Plugin) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Plugin) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *Plugin) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Plugin) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Plugin) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *Plugin) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLatestVersion

`func (o *Plugin) GetLatestVersion() string`

GetLatestVersion returns the LatestVersion field if non-nil, zero value otherwise.

### GetLatestVersionOk

`func (o *Plugin) GetLatestVersionOk() (*string, bool)`

GetLatestVersionOk returns a tuple with the LatestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersion

`func (o *Plugin) SetLatestVersion(v string)`

SetLatestVersion sets LatestVersion field to given value.

### HasLatestVersion

`func (o *Plugin) HasLatestVersion() bool`

HasLatestVersion returns a boolean if a field has been set.

### GetVersion

`func (o *Plugin) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Plugin) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Plugin) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Plugin) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetEnabled

`func (o *Plugin) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Plugin) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Plugin) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Plugin) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetUninstalled

`func (o *Plugin) GetUninstalled() bool`

GetUninstalled returns the Uninstalled field if non-nil, zero value otherwise.

### GetUninstalledOk

`func (o *Plugin) GetUninstalledOk() (*bool, bool)`

GetUninstalledOk returns a tuple with the Uninstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUninstalled

`func (o *Plugin) SetUninstalled(v bool)`

SetUninstalled sets Uninstalled field to given value.

### HasUninstalled

`func (o *Plugin) HasUninstalled() bool`

HasUninstalled returns a boolean if a field has been set.

### GetPeertubeEngine

`func (o *Plugin) GetPeertubeEngine() string`

GetPeertubeEngine returns the PeertubeEngine field if non-nil, zero value otherwise.

### GetPeertubeEngineOk

`func (o *Plugin) GetPeertubeEngineOk() (*string, bool)`

GetPeertubeEngineOk returns a tuple with the PeertubeEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeertubeEngine

`func (o *Plugin) SetPeertubeEngine(v string)`

SetPeertubeEngine sets PeertubeEngine field to given value.

### HasPeertubeEngine

`func (o *Plugin) HasPeertubeEngine() bool`

HasPeertubeEngine returns a boolean if a field has been set.

### GetDescription

`func (o *Plugin) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Plugin) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Plugin) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Plugin) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHomepage

`func (o *Plugin) GetHomepage() string`

GetHomepage returns the Homepage field if non-nil, zero value otherwise.

### GetHomepageOk

`func (o *Plugin) GetHomepageOk() (*string, bool)`

GetHomepageOk returns a tuple with the Homepage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomepage

`func (o *Plugin) SetHomepage(v string)`

SetHomepage sets Homepage field to given value.

### HasHomepage

`func (o *Plugin) HasHomepage() bool`

HasHomepage returns a boolean if a field has been set.

### GetSettings

`func (o *Plugin) GetSettings() map[string]interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *Plugin) GetSettingsOk() (*map[string]interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *Plugin) SetSettings(v map[string]interface{})`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *Plugin) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Plugin) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Plugin) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Plugin) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Plugin) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Plugin) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Plugin) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Plugin) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Plugin) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


