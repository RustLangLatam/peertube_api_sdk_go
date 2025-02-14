# PlaybackMetricCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlayerMode** | **string** |  | 
**Resolution** | Pointer to **float32** | Current player video resolution | [optional] 
**Fps** | Pointer to **float32** | Current player video fps | [optional] 
**P2pEnabled** | **bool** |  | 
**P2pPeers** | Pointer to **float32** | P2P peers connected (doesn&#39;t include WebSeed peers) | [optional] 
**ResolutionChanges** | **float32** | How many resolution changes occurred since the last metric creation | 
**BufferStalled** | Pointer to **float32** | How many times buffer has been stalled since the last metric creation | [optional] 
**Errors** | **float32** | How many errors occurred since the last metric creation | 
**DownloadedBytesP2P** | **float32** | How many bytes were downloaded with P2P since the last metric creation | 
**DownloadedBytesHTTP** | **float32** | How many bytes were downloaded with HTTP since the last metric creation | 
**UploadedBytesP2P** | **float32** | How many bytes were uploaded with P2P since the last metric creation | 
**VideoId** | [**ApiV1VideosOwnershipIdAcceptPostIdParameter**](ApiV1VideosOwnershipIdAcceptPostIdParameter.md) |  | 

## Methods

### NewPlaybackMetricCreate

`func NewPlaybackMetricCreate(playerMode string, p2pEnabled bool, resolutionChanges float32, errors float32, downloadedBytesP2P float32, downloadedBytesHTTP float32, uploadedBytesP2P float32, videoId ApiV1VideosOwnershipIdAcceptPostIdParameter, ) *PlaybackMetricCreate`

NewPlaybackMetricCreate instantiates a new PlaybackMetricCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaybackMetricCreateWithDefaults

`func NewPlaybackMetricCreateWithDefaults() *PlaybackMetricCreate`

NewPlaybackMetricCreateWithDefaults instantiates a new PlaybackMetricCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlayerMode

`func (o *PlaybackMetricCreate) GetPlayerMode() string`

GetPlayerMode returns the PlayerMode field if non-nil, zero value otherwise.

### GetPlayerModeOk

`func (o *PlaybackMetricCreate) GetPlayerModeOk() (*string, bool)`

GetPlayerModeOk returns a tuple with the PlayerMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerMode

`func (o *PlaybackMetricCreate) SetPlayerMode(v string)`

SetPlayerMode sets PlayerMode field to given value.


### GetResolution

`func (o *PlaybackMetricCreate) GetResolution() float32`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *PlaybackMetricCreate) GetResolutionOk() (*float32, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *PlaybackMetricCreate) SetResolution(v float32)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *PlaybackMetricCreate) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetFps

`func (o *PlaybackMetricCreate) GetFps() float32`

GetFps returns the Fps field if non-nil, zero value otherwise.

### GetFpsOk

`func (o *PlaybackMetricCreate) GetFpsOk() (*float32, bool)`

GetFpsOk returns a tuple with the Fps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFps

`func (o *PlaybackMetricCreate) SetFps(v float32)`

SetFps sets Fps field to given value.

### HasFps

`func (o *PlaybackMetricCreate) HasFps() bool`

HasFps returns a boolean if a field has been set.

### GetP2pEnabled

`func (o *PlaybackMetricCreate) GetP2pEnabled() bool`

GetP2pEnabled returns the P2pEnabled field if non-nil, zero value otherwise.

### GetP2pEnabledOk

`func (o *PlaybackMetricCreate) GetP2pEnabledOk() (*bool, bool)`

GetP2pEnabledOk returns a tuple with the P2pEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP2pEnabled

`func (o *PlaybackMetricCreate) SetP2pEnabled(v bool)`

SetP2pEnabled sets P2pEnabled field to given value.


### GetP2pPeers

`func (o *PlaybackMetricCreate) GetP2pPeers() float32`

GetP2pPeers returns the P2pPeers field if non-nil, zero value otherwise.

### GetP2pPeersOk

`func (o *PlaybackMetricCreate) GetP2pPeersOk() (*float32, bool)`

GetP2pPeersOk returns a tuple with the P2pPeers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP2pPeers

`func (o *PlaybackMetricCreate) SetP2pPeers(v float32)`

SetP2pPeers sets P2pPeers field to given value.

### HasP2pPeers

`func (o *PlaybackMetricCreate) HasP2pPeers() bool`

HasP2pPeers returns a boolean if a field has been set.

### GetResolutionChanges

`func (o *PlaybackMetricCreate) GetResolutionChanges() float32`

GetResolutionChanges returns the ResolutionChanges field if non-nil, zero value otherwise.

### GetResolutionChangesOk

`func (o *PlaybackMetricCreate) GetResolutionChangesOk() (*float32, bool)`

GetResolutionChangesOk returns a tuple with the ResolutionChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionChanges

`func (o *PlaybackMetricCreate) SetResolutionChanges(v float32)`

SetResolutionChanges sets ResolutionChanges field to given value.


### GetBufferStalled

`func (o *PlaybackMetricCreate) GetBufferStalled() float32`

GetBufferStalled returns the BufferStalled field if non-nil, zero value otherwise.

### GetBufferStalledOk

`func (o *PlaybackMetricCreate) GetBufferStalledOk() (*float32, bool)`

GetBufferStalledOk returns a tuple with the BufferStalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBufferStalled

`func (o *PlaybackMetricCreate) SetBufferStalled(v float32)`

SetBufferStalled sets BufferStalled field to given value.

### HasBufferStalled

`func (o *PlaybackMetricCreate) HasBufferStalled() bool`

HasBufferStalled returns a boolean if a field has been set.

### GetErrors

`func (o *PlaybackMetricCreate) GetErrors() float32`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *PlaybackMetricCreate) GetErrorsOk() (*float32, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *PlaybackMetricCreate) SetErrors(v float32)`

SetErrors sets Errors field to given value.


### GetDownloadedBytesP2P

`func (o *PlaybackMetricCreate) GetDownloadedBytesP2P() float32`

GetDownloadedBytesP2P returns the DownloadedBytesP2P field if non-nil, zero value otherwise.

### GetDownloadedBytesP2POk

`func (o *PlaybackMetricCreate) GetDownloadedBytesP2POk() (*float32, bool)`

GetDownloadedBytesP2POk returns a tuple with the DownloadedBytesP2P field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadedBytesP2P

`func (o *PlaybackMetricCreate) SetDownloadedBytesP2P(v float32)`

SetDownloadedBytesP2P sets DownloadedBytesP2P field to given value.


### GetDownloadedBytesHTTP

`func (o *PlaybackMetricCreate) GetDownloadedBytesHTTP() float32`

GetDownloadedBytesHTTP returns the DownloadedBytesHTTP field if non-nil, zero value otherwise.

### GetDownloadedBytesHTTPOk

`func (o *PlaybackMetricCreate) GetDownloadedBytesHTTPOk() (*float32, bool)`

GetDownloadedBytesHTTPOk returns a tuple with the DownloadedBytesHTTP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadedBytesHTTP

`func (o *PlaybackMetricCreate) SetDownloadedBytesHTTP(v float32)`

SetDownloadedBytesHTTP sets DownloadedBytesHTTP field to given value.


### GetUploadedBytesP2P

`func (o *PlaybackMetricCreate) GetUploadedBytesP2P() float32`

GetUploadedBytesP2P returns the UploadedBytesP2P field if non-nil, zero value otherwise.

### GetUploadedBytesP2POk

`func (o *PlaybackMetricCreate) GetUploadedBytesP2POk() (*float32, bool)`

GetUploadedBytesP2POk returns a tuple with the UploadedBytesP2P field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadedBytesP2P

`func (o *PlaybackMetricCreate) SetUploadedBytesP2P(v float32)`

SetUploadedBytesP2P sets UploadedBytesP2P field to given value.


### GetVideoId

`func (o *PlaybackMetricCreate) GetVideoId() ApiV1VideosOwnershipIdAcceptPostIdParameter`

GetVideoId returns the VideoId field if non-nil, zero value otherwise.

### GetVideoIdOk

`func (o *PlaybackMetricCreate) GetVideoIdOk() (*ApiV1VideosOwnershipIdAcceptPostIdParameter, bool)`

GetVideoIdOk returns a tuple with the VideoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoId

`func (o *PlaybackMetricCreate) SetVideoId(v ApiV1VideosOwnershipIdAcceptPostIdParameter)`

SetVideoId sets VideoId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


