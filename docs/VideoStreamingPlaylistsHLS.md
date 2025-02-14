# VideoStreamingPlaylistsHLS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlaylistUrl** | Pointer to **string** |  | [optional] 
**SegmentsSha256Url** | Pointer to **string** |  | [optional] 
**Files** | Pointer to [**[]VideoFile**](VideoFile.md) | Video files associated to this playlist.  The difference with the root &#x60;files&#x60; property is that these files are fragmented, so they can be used in this streaming playlist (HLS, etc.)  | [optional] 
**Redundancies** | Pointer to [**[]VideoStreamingPlaylistsHLSRedundanciesInner**](VideoStreamingPlaylistsHLSRedundanciesInner.md) |  | [optional] 

## Methods

### NewVideoStreamingPlaylistsHLS

`func NewVideoStreamingPlaylistsHLS() *VideoStreamingPlaylistsHLS`

NewVideoStreamingPlaylistsHLS instantiates a new VideoStreamingPlaylistsHLS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoStreamingPlaylistsHLSWithDefaults

`func NewVideoStreamingPlaylistsHLSWithDefaults() *VideoStreamingPlaylistsHLS`

NewVideoStreamingPlaylistsHLSWithDefaults instantiates a new VideoStreamingPlaylistsHLS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlaylistUrl

`func (o *VideoStreamingPlaylistsHLS) GetPlaylistUrl() string`

GetPlaylistUrl returns the PlaylistUrl field if non-nil, zero value otherwise.

### GetPlaylistUrlOk

`func (o *VideoStreamingPlaylistsHLS) GetPlaylistUrlOk() (*string, bool)`

GetPlaylistUrlOk returns a tuple with the PlaylistUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylistUrl

`func (o *VideoStreamingPlaylistsHLS) SetPlaylistUrl(v string)`

SetPlaylistUrl sets PlaylistUrl field to given value.

### HasPlaylistUrl

`func (o *VideoStreamingPlaylistsHLS) HasPlaylistUrl() bool`

HasPlaylistUrl returns a boolean if a field has been set.

### GetSegmentsSha256Url

`func (o *VideoStreamingPlaylistsHLS) GetSegmentsSha256Url() string`

GetSegmentsSha256Url returns the SegmentsSha256Url field if non-nil, zero value otherwise.

### GetSegmentsSha256UrlOk

`func (o *VideoStreamingPlaylistsHLS) GetSegmentsSha256UrlOk() (*string, bool)`

GetSegmentsSha256UrlOk returns a tuple with the SegmentsSha256Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentsSha256Url

`func (o *VideoStreamingPlaylistsHLS) SetSegmentsSha256Url(v string)`

SetSegmentsSha256Url sets SegmentsSha256Url field to given value.

### HasSegmentsSha256Url

`func (o *VideoStreamingPlaylistsHLS) HasSegmentsSha256Url() bool`

HasSegmentsSha256Url returns a boolean if a field has been set.

### GetFiles

`func (o *VideoStreamingPlaylistsHLS) GetFiles() []VideoFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *VideoStreamingPlaylistsHLS) GetFilesOk() (*[]VideoFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *VideoStreamingPlaylistsHLS) SetFiles(v []VideoFile)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *VideoStreamingPlaylistsHLS) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetRedundancies

`func (o *VideoStreamingPlaylistsHLS) GetRedundancies() []VideoStreamingPlaylistsHLSRedundanciesInner`

GetRedundancies returns the Redundancies field if non-nil, zero value otherwise.

### GetRedundanciesOk

`func (o *VideoStreamingPlaylistsHLS) GetRedundanciesOk() (*[]VideoStreamingPlaylistsHLSRedundanciesInner, bool)`

GetRedundanciesOk returns a tuple with the Redundancies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancies

`func (o *VideoStreamingPlaylistsHLS) SetRedundancies(v []VideoStreamingPlaylistsHLSRedundanciesInner)`

SetRedundancies sets Redundancies field to given value.

### HasRedundancies

`func (o *VideoStreamingPlaylistsHLS) HasRedundancies() bool`

HasRedundancies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


