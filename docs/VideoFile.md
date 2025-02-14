# VideoFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**MagnetUri** | Pointer to **string** | magnet URI allowing to resolve the video via BitTorrent without a metainfo file | [optional] 
**Resolution** | Pointer to [**VideoResolutionConstant**](VideoResolutionConstant.md) |  | [optional] 
**Size** | Pointer to **int32** | Video file size in bytes | [optional] 
**TorrentUrl** | Pointer to **string** | Direct URL of the torrent file | [optional] 
**TorrentDownloadUrl** | Pointer to **string** | URL endpoint that transfers the torrent file as an attachment (so that the browser opens a download dialog) | [optional] 
**FileUrl** | Pointer to **string** | Direct URL of the video | [optional] 
**FileDownloadUrl** | Pointer to **string** | URL endpoint that transfers the video file as an attachment (so that the browser opens a download dialog) | [optional] 
**Fps** | Pointer to **float32** | Frames per second of the video file | [optional] 
**Width** | Pointer to **float32** | **PeerTube &gt;&#x3D; 6.1** Video stream width | [optional] 
**Height** | Pointer to **float32** | **PeerTube &gt;&#x3D; 6.1** Video stream height | [optional] 
**MetadataUrl** | Pointer to **string** | URL dereferencing the output of ffprobe on the file | [optional] 
**HasAudio** | Pointer to **bool** | **PeerTube &gt;&#x3D; 6.2** The file container has an audio stream | [optional] 
**HasVideo** | Pointer to **bool** | **PeerTube &gt;&#x3D; 6.2** The file container has a video stream | [optional] 
**Storage** | Pointer to [**FileStorage**](FileStorage.md) |  | [optional] 

## Methods

### NewVideoFile

`func NewVideoFile() *VideoFile`

NewVideoFile instantiates a new VideoFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoFileWithDefaults

`func NewVideoFileWithDefaults() *VideoFile`

NewVideoFileWithDefaults instantiates a new VideoFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VideoFile) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VideoFile) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VideoFile) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *VideoFile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMagnetUri

`func (o *VideoFile) GetMagnetUri() string`

GetMagnetUri returns the MagnetUri field if non-nil, zero value otherwise.

### GetMagnetUriOk

`func (o *VideoFile) GetMagnetUriOk() (*string, bool)`

GetMagnetUriOk returns a tuple with the MagnetUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMagnetUri

`func (o *VideoFile) SetMagnetUri(v string)`

SetMagnetUri sets MagnetUri field to given value.

### HasMagnetUri

`func (o *VideoFile) HasMagnetUri() bool`

HasMagnetUri returns a boolean if a field has been set.

### GetResolution

`func (o *VideoFile) GetResolution() VideoResolutionConstant`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *VideoFile) GetResolutionOk() (*VideoResolutionConstant, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *VideoFile) SetResolution(v VideoResolutionConstant)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *VideoFile) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetSize

`func (o *VideoFile) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *VideoFile) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *VideoFile) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *VideoFile) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetTorrentUrl

`func (o *VideoFile) GetTorrentUrl() string`

GetTorrentUrl returns the TorrentUrl field if non-nil, zero value otherwise.

### GetTorrentUrlOk

`func (o *VideoFile) GetTorrentUrlOk() (*string, bool)`

GetTorrentUrlOk returns a tuple with the TorrentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTorrentUrl

`func (o *VideoFile) SetTorrentUrl(v string)`

SetTorrentUrl sets TorrentUrl field to given value.

### HasTorrentUrl

`func (o *VideoFile) HasTorrentUrl() bool`

HasTorrentUrl returns a boolean if a field has been set.

### GetTorrentDownloadUrl

`func (o *VideoFile) GetTorrentDownloadUrl() string`

GetTorrentDownloadUrl returns the TorrentDownloadUrl field if non-nil, zero value otherwise.

### GetTorrentDownloadUrlOk

`func (o *VideoFile) GetTorrentDownloadUrlOk() (*string, bool)`

GetTorrentDownloadUrlOk returns a tuple with the TorrentDownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTorrentDownloadUrl

`func (o *VideoFile) SetTorrentDownloadUrl(v string)`

SetTorrentDownloadUrl sets TorrentDownloadUrl field to given value.

### HasTorrentDownloadUrl

`func (o *VideoFile) HasTorrentDownloadUrl() bool`

HasTorrentDownloadUrl returns a boolean if a field has been set.

### GetFileUrl

`func (o *VideoFile) GetFileUrl() string`

GetFileUrl returns the FileUrl field if non-nil, zero value otherwise.

### GetFileUrlOk

`func (o *VideoFile) GetFileUrlOk() (*string, bool)`

GetFileUrlOk returns a tuple with the FileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUrl

`func (o *VideoFile) SetFileUrl(v string)`

SetFileUrl sets FileUrl field to given value.

### HasFileUrl

`func (o *VideoFile) HasFileUrl() bool`

HasFileUrl returns a boolean if a field has been set.

### GetFileDownloadUrl

`func (o *VideoFile) GetFileDownloadUrl() string`

GetFileDownloadUrl returns the FileDownloadUrl field if non-nil, zero value otherwise.

### GetFileDownloadUrlOk

`func (o *VideoFile) GetFileDownloadUrlOk() (*string, bool)`

GetFileDownloadUrlOk returns a tuple with the FileDownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileDownloadUrl

`func (o *VideoFile) SetFileDownloadUrl(v string)`

SetFileDownloadUrl sets FileDownloadUrl field to given value.

### HasFileDownloadUrl

`func (o *VideoFile) HasFileDownloadUrl() bool`

HasFileDownloadUrl returns a boolean if a field has been set.

### GetFps

`func (o *VideoFile) GetFps() float32`

GetFps returns the Fps field if non-nil, zero value otherwise.

### GetFpsOk

`func (o *VideoFile) GetFpsOk() (*float32, bool)`

GetFpsOk returns a tuple with the Fps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFps

`func (o *VideoFile) SetFps(v float32)`

SetFps sets Fps field to given value.

### HasFps

`func (o *VideoFile) HasFps() bool`

HasFps returns a boolean if a field has been set.

### GetWidth

`func (o *VideoFile) GetWidth() float32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *VideoFile) GetWidthOk() (*float32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *VideoFile) SetWidth(v float32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *VideoFile) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetHeight

`func (o *VideoFile) GetHeight() float32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *VideoFile) GetHeightOk() (*float32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *VideoFile) SetHeight(v float32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *VideoFile) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetMetadataUrl

`func (o *VideoFile) GetMetadataUrl() string`

GetMetadataUrl returns the MetadataUrl field if non-nil, zero value otherwise.

### GetMetadataUrlOk

`func (o *VideoFile) GetMetadataUrlOk() (*string, bool)`

GetMetadataUrlOk returns a tuple with the MetadataUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataUrl

`func (o *VideoFile) SetMetadataUrl(v string)`

SetMetadataUrl sets MetadataUrl field to given value.

### HasMetadataUrl

`func (o *VideoFile) HasMetadataUrl() bool`

HasMetadataUrl returns a boolean if a field has been set.

### GetHasAudio

`func (o *VideoFile) GetHasAudio() bool`

GetHasAudio returns the HasAudio field if non-nil, zero value otherwise.

### GetHasAudioOk

`func (o *VideoFile) GetHasAudioOk() (*bool, bool)`

GetHasAudioOk returns a tuple with the HasAudio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAudio

`func (o *VideoFile) SetHasAudio(v bool)`

SetHasAudio sets HasAudio field to given value.

### HasHasAudio

`func (o *VideoFile) HasHasAudio() bool`

HasHasAudio returns a boolean if a field has been set.

### GetHasVideo

`func (o *VideoFile) GetHasVideo() bool`

GetHasVideo returns the HasVideo field if non-nil, zero value otherwise.

### GetHasVideoOk

`func (o *VideoFile) GetHasVideoOk() (*bool, bool)`

GetHasVideoOk returns a tuple with the HasVideo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasVideo

`func (o *VideoFile) SetHasVideo(v bool)`

SetHasVideo sets HasVideo field to given value.

### HasHasVideo

`func (o *VideoFile) HasHasVideo() bool`

HasHasVideo returns a boolean if a field has been set.

### GetStorage

`func (o *VideoFile) GetStorage() FileStorage`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *VideoFile) GetStorageOk() (*FileStorage, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *VideoFile) SetStorage(v FileStorage)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *VideoFile) HasStorage() bool`

HasStorage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


