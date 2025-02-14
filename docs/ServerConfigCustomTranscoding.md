# ServerConfigCustomTranscoding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**OriginalFile** | Pointer to [**ServerConfigCustomTranscodingOriginalFile**](ServerConfigCustomTranscodingOriginalFile.md) |  | [optional] 
**AllowAdditionalExtensions** | Pointer to **bool** | Allow your users to upload .mkv, .mov, .avi, .wmv, .flv, .f4v, .3g2, .3gp, .mts, m2ts, .mxf, .nut videos | [optional] 
**AllowAudioFiles** | Pointer to **bool** | If a user uploads an audio file, PeerTube will create a video by merging the preview file and the audio file | [optional] 
**Threads** | Pointer to **int32** | Amount of threads used by ffmpeg for 1 transcoding job | [optional] 
**Concurrency** | Pointer to **float32** | Amount of transcoding jobs to execute in parallel | [optional] 
**Profile** | Pointer to **string** | New profiles can be added by plugins ; available in core PeerTube: &#39;default&#39;.  | [optional] 
**Resolutions** | Pointer to [**ServerConfigCustomTranscodingResolutions**](ServerConfigCustomTranscodingResolutions.md) |  | [optional] 
**WebVideos** | Pointer to [**ServerConfigCustomTranscodingWebVideos**](ServerConfigCustomTranscodingWebVideos.md) |  | [optional] 
**Hls** | Pointer to [**ServerConfigCustomTranscodingHls**](ServerConfigCustomTranscodingHls.md) |  | [optional] 

## Methods

### NewServerConfigCustomTranscoding

`func NewServerConfigCustomTranscoding() *ServerConfigCustomTranscoding`

NewServerConfigCustomTranscoding instantiates a new ServerConfigCustomTranscoding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerConfigCustomTranscodingWithDefaults

`func NewServerConfigCustomTranscodingWithDefaults() *ServerConfigCustomTranscoding`

NewServerConfigCustomTranscodingWithDefaults instantiates a new ServerConfigCustomTranscoding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ServerConfigCustomTranscoding) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ServerConfigCustomTranscoding) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ServerConfigCustomTranscoding) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ServerConfigCustomTranscoding) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetOriginalFile

`func (o *ServerConfigCustomTranscoding) GetOriginalFile() ServerConfigCustomTranscodingOriginalFile`

GetOriginalFile returns the OriginalFile field if non-nil, zero value otherwise.

### GetOriginalFileOk

`func (o *ServerConfigCustomTranscoding) GetOriginalFileOk() (*ServerConfigCustomTranscodingOriginalFile, bool)`

GetOriginalFileOk returns a tuple with the OriginalFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalFile

`func (o *ServerConfigCustomTranscoding) SetOriginalFile(v ServerConfigCustomTranscodingOriginalFile)`

SetOriginalFile sets OriginalFile field to given value.

### HasOriginalFile

`func (o *ServerConfigCustomTranscoding) HasOriginalFile() bool`

HasOriginalFile returns a boolean if a field has been set.

### GetAllowAdditionalExtensions

`func (o *ServerConfigCustomTranscoding) GetAllowAdditionalExtensions() bool`

GetAllowAdditionalExtensions returns the AllowAdditionalExtensions field if non-nil, zero value otherwise.

### GetAllowAdditionalExtensionsOk

`func (o *ServerConfigCustomTranscoding) GetAllowAdditionalExtensionsOk() (*bool, bool)`

GetAllowAdditionalExtensionsOk returns a tuple with the AllowAdditionalExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAdditionalExtensions

`func (o *ServerConfigCustomTranscoding) SetAllowAdditionalExtensions(v bool)`

SetAllowAdditionalExtensions sets AllowAdditionalExtensions field to given value.

### HasAllowAdditionalExtensions

`func (o *ServerConfigCustomTranscoding) HasAllowAdditionalExtensions() bool`

HasAllowAdditionalExtensions returns a boolean if a field has been set.

### GetAllowAudioFiles

`func (o *ServerConfigCustomTranscoding) GetAllowAudioFiles() bool`

GetAllowAudioFiles returns the AllowAudioFiles field if non-nil, zero value otherwise.

### GetAllowAudioFilesOk

`func (o *ServerConfigCustomTranscoding) GetAllowAudioFilesOk() (*bool, bool)`

GetAllowAudioFilesOk returns a tuple with the AllowAudioFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAudioFiles

`func (o *ServerConfigCustomTranscoding) SetAllowAudioFiles(v bool)`

SetAllowAudioFiles sets AllowAudioFiles field to given value.

### HasAllowAudioFiles

`func (o *ServerConfigCustomTranscoding) HasAllowAudioFiles() bool`

HasAllowAudioFiles returns a boolean if a field has been set.

### GetThreads

`func (o *ServerConfigCustomTranscoding) GetThreads() int32`

GetThreads returns the Threads field if non-nil, zero value otherwise.

### GetThreadsOk

`func (o *ServerConfigCustomTranscoding) GetThreadsOk() (*int32, bool)`

GetThreadsOk returns a tuple with the Threads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreads

`func (o *ServerConfigCustomTranscoding) SetThreads(v int32)`

SetThreads sets Threads field to given value.

### HasThreads

`func (o *ServerConfigCustomTranscoding) HasThreads() bool`

HasThreads returns a boolean if a field has been set.

### GetConcurrency

`func (o *ServerConfigCustomTranscoding) GetConcurrency() float32`

GetConcurrency returns the Concurrency field if non-nil, zero value otherwise.

### GetConcurrencyOk

`func (o *ServerConfigCustomTranscoding) GetConcurrencyOk() (*float32, bool)`

GetConcurrencyOk returns a tuple with the Concurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrency

`func (o *ServerConfigCustomTranscoding) SetConcurrency(v float32)`

SetConcurrency sets Concurrency field to given value.

### HasConcurrency

`func (o *ServerConfigCustomTranscoding) HasConcurrency() bool`

HasConcurrency returns a boolean if a field has been set.

### GetProfile

`func (o *ServerConfigCustomTranscoding) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ServerConfigCustomTranscoding) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ServerConfigCustomTranscoding) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *ServerConfigCustomTranscoding) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetResolutions

`func (o *ServerConfigCustomTranscoding) GetResolutions() ServerConfigCustomTranscodingResolutions`

GetResolutions returns the Resolutions field if non-nil, zero value otherwise.

### GetResolutionsOk

`func (o *ServerConfigCustomTranscoding) GetResolutionsOk() (*ServerConfigCustomTranscodingResolutions, bool)`

GetResolutionsOk returns a tuple with the Resolutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutions

`func (o *ServerConfigCustomTranscoding) SetResolutions(v ServerConfigCustomTranscodingResolutions)`

SetResolutions sets Resolutions field to given value.

### HasResolutions

`func (o *ServerConfigCustomTranscoding) HasResolutions() bool`

HasResolutions returns a boolean if a field has been set.

### GetWebVideos

`func (o *ServerConfigCustomTranscoding) GetWebVideos() ServerConfigCustomTranscodingWebVideos`

GetWebVideos returns the WebVideos field if non-nil, zero value otherwise.

### GetWebVideosOk

`func (o *ServerConfigCustomTranscoding) GetWebVideosOk() (*ServerConfigCustomTranscodingWebVideos, bool)`

GetWebVideosOk returns a tuple with the WebVideos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebVideos

`func (o *ServerConfigCustomTranscoding) SetWebVideos(v ServerConfigCustomTranscodingWebVideos)`

SetWebVideos sets WebVideos field to given value.

### HasWebVideos

`func (o *ServerConfigCustomTranscoding) HasWebVideos() bool`

HasWebVideos returns a boolean if a field has been set.

### GetHls

`func (o *ServerConfigCustomTranscoding) GetHls() ServerConfigCustomTranscodingHls`

GetHls returns the Hls field if non-nil, zero value otherwise.

### GetHlsOk

`func (o *ServerConfigCustomTranscoding) GetHlsOk() (*ServerConfigCustomTranscodingHls, bool)`

GetHlsOk returns a tuple with the Hls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHls

`func (o *ServerConfigCustomTranscoding) SetHls(v ServerConfigCustomTranscodingHls)`

SetHls sets Hls field to given value.

### HasHls

`func (o *ServerConfigCustomTranscoding) HasHls() bool`

HasHls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


