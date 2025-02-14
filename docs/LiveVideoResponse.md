# LiveVideoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RtmpUrl** | Pointer to **string** | Included in the response if an appropriate token is provided | [optional] 
**RtmpsUrl** | Pointer to **string** | Included in the response if an appropriate token is provided | [optional] 
**StreamKey** | Pointer to **string** | RTMP stream key to use to stream into this live video. Included in the response if an appropriate token is provided | [optional] 
**SaveReplay** | Pointer to **bool** |  | [optional] 
**ReplaySettings** | Pointer to [**LiveVideoReplaySettings**](LiveVideoReplaySettings.md) |  | [optional] 
**PermanentLive** | Pointer to **bool** | User can stream multiple times in a permanent live | [optional] 
**LatencyMode** | Pointer to [**LiveVideoLatencyMode**](LiveVideoLatencyMode.md) |  | [optional] 

## Methods

### NewLiveVideoResponse

`func NewLiveVideoResponse() *LiveVideoResponse`

NewLiveVideoResponse instantiates a new LiveVideoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveVideoResponseWithDefaults

`func NewLiveVideoResponseWithDefaults() *LiveVideoResponse`

NewLiveVideoResponseWithDefaults instantiates a new LiveVideoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRtmpUrl

`func (o *LiveVideoResponse) GetRtmpUrl() string`

GetRtmpUrl returns the RtmpUrl field if non-nil, zero value otherwise.

### GetRtmpUrlOk

`func (o *LiveVideoResponse) GetRtmpUrlOk() (*string, bool)`

GetRtmpUrlOk returns a tuple with the RtmpUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRtmpUrl

`func (o *LiveVideoResponse) SetRtmpUrl(v string)`

SetRtmpUrl sets RtmpUrl field to given value.

### HasRtmpUrl

`func (o *LiveVideoResponse) HasRtmpUrl() bool`

HasRtmpUrl returns a boolean if a field has been set.

### GetRtmpsUrl

`func (o *LiveVideoResponse) GetRtmpsUrl() string`

GetRtmpsUrl returns the RtmpsUrl field if non-nil, zero value otherwise.

### GetRtmpsUrlOk

`func (o *LiveVideoResponse) GetRtmpsUrlOk() (*string, bool)`

GetRtmpsUrlOk returns a tuple with the RtmpsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRtmpsUrl

`func (o *LiveVideoResponse) SetRtmpsUrl(v string)`

SetRtmpsUrl sets RtmpsUrl field to given value.

### HasRtmpsUrl

`func (o *LiveVideoResponse) HasRtmpsUrl() bool`

HasRtmpsUrl returns a boolean if a field has been set.

### GetStreamKey

`func (o *LiveVideoResponse) GetStreamKey() string`

GetStreamKey returns the StreamKey field if non-nil, zero value otherwise.

### GetStreamKeyOk

`func (o *LiveVideoResponse) GetStreamKeyOk() (*string, bool)`

GetStreamKeyOk returns a tuple with the StreamKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamKey

`func (o *LiveVideoResponse) SetStreamKey(v string)`

SetStreamKey sets StreamKey field to given value.

### HasStreamKey

`func (o *LiveVideoResponse) HasStreamKey() bool`

HasStreamKey returns a boolean if a field has been set.

### GetSaveReplay

`func (o *LiveVideoResponse) GetSaveReplay() bool`

GetSaveReplay returns the SaveReplay field if non-nil, zero value otherwise.

### GetSaveReplayOk

`func (o *LiveVideoResponse) GetSaveReplayOk() (*bool, bool)`

GetSaveReplayOk returns a tuple with the SaveReplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveReplay

`func (o *LiveVideoResponse) SetSaveReplay(v bool)`

SetSaveReplay sets SaveReplay field to given value.

### HasSaveReplay

`func (o *LiveVideoResponse) HasSaveReplay() bool`

HasSaveReplay returns a boolean if a field has been set.

### GetReplaySettings

`func (o *LiveVideoResponse) GetReplaySettings() LiveVideoReplaySettings`

GetReplaySettings returns the ReplaySettings field if non-nil, zero value otherwise.

### GetReplaySettingsOk

`func (o *LiveVideoResponse) GetReplaySettingsOk() (*LiveVideoReplaySettings, bool)`

GetReplaySettingsOk returns a tuple with the ReplaySettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplaySettings

`func (o *LiveVideoResponse) SetReplaySettings(v LiveVideoReplaySettings)`

SetReplaySettings sets ReplaySettings field to given value.

### HasReplaySettings

`func (o *LiveVideoResponse) HasReplaySettings() bool`

HasReplaySettings returns a boolean if a field has been set.

### GetPermanentLive

`func (o *LiveVideoResponse) GetPermanentLive() bool`

GetPermanentLive returns the PermanentLive field if non-nil, zero value otherwise.

### GetPermanentLiveOk

`func (o *LiveVideoResponse) GetPermanentLiveOk() (*bool, bool)`

GetPermanentLiveOk returns a tuple with the PermanentLive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermanentLive

`func (o *LiveVideoResponse) SetPermanentLive(v bool)`

SetPermanentLive sets PermanentLive field to given value.

### HasPermanentLive

`func (o *LiveVideoResponse) HasPermanentLive() bool`

HasPermanentLive returns a boolean if a field has been set.

### GetLatencyMode

`func (o *LiveVideoResponse) GetLatencyMode() LiveVideoLatencyMode`

GetLatencyMode returns the LatencyMode field if non-nil, zero value otherwise.

### GetLatencyModeOk

`func (o *LiveVideoResponse) GetLatencyModeOk() (*LiveVideoLatencyMode, bool)`

GetLatencyModeOk returns a tuple with the LatencyMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyMode

`func (o *LiveVideoResponse) SetLatencyMode(v LiveVideoLatencyMode)`

SetLatencyMode sets LatencyMode field to given value.

### HasLatencyMode

`func (o *LiveVideoResponse) HasLatencyMode() bool`

HasLatencyMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


