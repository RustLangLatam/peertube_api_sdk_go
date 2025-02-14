# LiveVideoUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SaveReplay** | Pointer to **bool** |  | [optional] 
**ReplaySettings** | Pointer to [**LiveVideoReplaySettings**](LiveVideoReplaySettings.md) |  | [optional] 
**PermanentLive** | Pointer to **bool** | User can stream multiple times in a permanent live | [optional] 
**LatencyMode** | Pointer to [**LiveVideoLatencyMode**](LiveVideoLatencyMode.md) |  | [optional] 

## Methods

### NewLiveVideoUpdate

`func NewLiveVideoUpdate() *LiveVideoUpdate`

NewLiveVideoUpdate instantiates a new LiveVideoUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveVideoUpdateWithDefaults

`func NewLiveVideoUpdateWithDefaults() *LiveVideoUpdate`

NewLiveVideoUpdateWithDefaults instantiates a new LiveVideoUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSaveReplay

`func (o *LiveVideoUpdate) GetSaveReplay() bool`

GetSaveReplay returns the SaveReplay field if non-nil, zero value otherwise.

### GetSaveReplayOk

`func (o *LiveVideoUpdate) GetSaveReplayOk() (*bool, bool)`

GetSaveReplayOk returns a tuple with the SaveReplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveReplay

`func (o *LiveVideoUpdate) SetSaveReplay(v bool)`

SetSaveReplay sets SaveReplay field to given value.

### HasSaveReplay

`func (o *LiveVideoUpdate) HasSaveReplay() bool`

HasSaveReplay returns a boolean if a field has been set.

### GetReplaySettings

`func (o *LiveVideoUpdate) GetReplaySettings() LiveVideoReplaySettings`

GetReplaySettings returns the ReplaySettings field if non-nil, zero value otherwise.

### GetReplaySettingsOk

`func (o *LiveVideoUpdate) GetReplaySettingsOk() (*LiveVideoReplaySettings, bool)`

GetReplaySettingsOk returns a tuple with the ReplaySettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaySettings

`func (o *LiveVideoUpdate) SetReplaySettings(v LiveVideoReplaySettings)`

SetReplaySettings sets ReplaySettings field to given value.

### HasReplaySettings

`func (o *LiveVideoUpdate) HasReplaySettings() bool`

HasReplaySettings returns a boolean if a field has been set.

### GetPermanentLive

`func (o *LiveVideoUpdate) GetPermanentLive() bool`

GetPermanentLive returns the PermanentLive field if non-nil, zero value otherwise.

### GetPermanentLiveOk

`func (o *LiveVideoUpdate) GetPermanentLiveOk() (*bool, bool)`

GetPermanentLiveOk returns a tuple with the PermanentLive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermanentLive

`func (o *LiveVideoUpdate) SetPermanentLive(v bool)`

SetPermanentLive sets PermanentLive field to given value.

### HasPermanentLive

`func (o *LiveVideoUpdate) HasPermanentLive() bool`

HasPermanentLive returns a boolean if a field has been set.

### GetLatencyMode

`func (o *LiveVideoUpdate) GetLatencyMode() LiveVideoLatencyMode`

GetLatencyMode returns the LatencyMode field if non-nil, zero value otherwise.

### GetLatencyModeOk

`func (o *LiveVideoUpdate) GetLatencyModeOk() (*LiveVideoLatencyMode, bool)`

GetLatencyModeOk returns a tuple with the LatencyMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyMode

`func (o *LiveVideoUpdate) SetLatencyMode(v LiveVideoLatencyMode)`

SetLatencyMode sets LatencyMode field to given value.

### HasLatencyMode

`func (o *LiveVideoUpdate) HasLatencyMode() bool`

HasLatencyMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


