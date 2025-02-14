# VideoStreamingPlaylists

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlaylistUrl** | Pointer to **string** |  | [optional] 
**SegmentsSha256Url** | Pointer to **string** |  | [optional] 
**Files** | Pointer to [**[]VideoFile**](VideoFile.md) | Video files associated to this playlist.  The difference with the root &#x60;files&#x60; property is that these files are fragmented, so they can be used in this streaming playlist (HLS, etc.)  | [optional] 
**Redundancies** | Pointer to [**[]VideoStreamingPlaylistsHLSRedundanciesInner**](VideoStreamingPlaylistsHLSRedundanciesInner.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **int32** | Playlist type: - &#x60;1&#x60;: HLS  | [optional] 

## Methods

### NewVideoStreamingPlaylists

`func NewVideoStreamingPlaylists() *VideoStreamingPlaylists`

NewVideoStreamingPlaylists instantiates a new VideoStreamingPlaylists object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoStreamingPlaylistsWithDefaults

`func NewVideoStreamingPlaylistsWithDefaults() *VideoStreamingPlaylists`

NewVideoStreamingPlaylistsWithDefaults instantiates a new VideoStreamingPlaylists object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlaylistUrl

`func (o *VideoStreamingPlaylists) GetPlaylistUrl() string`

GetPlaylistUrl returns the PlaylistUrl field if non-nil, zero value otherwise.

### GetPlaylistUrlOk

`func (o *VideoStreamingPlaylists) GetPlaylistUrlOk() (*string, bool)`

GetPlaylistUrlOk returns a tuple with the PlaylistUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylistUrl

`func (o *VideoStreamingPlaylists) SetPlaylistUrl(v string)`

SetPlaylistUrl sets PlaylistUrl field to given value.

### HasPlaylistUrl

`func (o *VideoStreamingPlaylists) HasPlaylistUrl() bool`

HasPlaylistUrl returns a boolean if a field has been set.

### GetSegmentsSha256Url

`func (o *VideoStreamingPlaylists) GetSegmentsSha256Url() string`

GetSegmentsSha256Url returns the SegmentsSha256Url field if non-nil, zero value otherwise.

### GetSegmentsSha256UrlOk

`func (o *VideoStreamingPlaylists) GetSegmentsSha256UrlOk() (*string, bool)`

GetSegmentsSha256UrlOk returns a tuple with the SegmentsSha256Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentsSha256Url

`func (o *VideoStreamingPlaylists) SetSegmentsSha256Url(v string)`

SetSegmentsSha256Url sets SegmentsSha256Url field to given value.

### HasSegmentsSha256Url

`func (o *VideoStreamingPlaylists) HasSegmentsSha256Url() bool`

HasSegmentsSha256Url returns a boolean if a field has been set.

### GetFiles

`func (o *VideoStreamingPlaylists) GetFiles() []VideoFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *VideoStreamingPlaylists) GetFilesOk() (*[]VideoFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *VideoStreamingPlaylists) SetFiles(v []VideoFile)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *VideoStreamingPlaylists) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetRedundancies

`func (o *VideoStreamingPlaylists) GetRedundancies() []VideoStreamingPlaylistsHLSRedundanciesInner`

GetRedundancies returns the Redundancies field if non-nil, zero value otherwise.

### GetRedundanciesOk

`func (o *VideoStreamingPlaylists) GetRedundanciesOk() (*[]VideoStreamingPlaylistsHLSRedundanciesInner, bool)`

GetRedundanciesOk returns a tuple with the Redundancies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancies

`func (o *VideoStreamingPlaylists) SetRedundancies(v []VideoStreamingPlaylistsHLSRedundanciesInner)`

SetRedundancies sets Redundancies field to given value.

### HasRedundancies

`func (o *VideoStreamingPlaylists) HasRedundancies() bool`

HasRedundancies returns a boolean if a field has been set.

### GetId

`func (o *VideoStreamingPlaylists) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VideoStreamingPlaylists) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VideoStreamingPlaylists) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *VideoStreamingPlaylists) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *VideoStreamingPlaylists) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VideoStreamingPlaylists) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VideoStreamingPlaylists) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *VideoStreamingPlaylists) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


