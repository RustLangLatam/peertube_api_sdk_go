# ServerConfigImportVideos

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Http** | Pointer to [**ServerConfigEmail**](ServerConfigEmail.md) |  | [optional] 
**Torrent** | Pointer to [**ServerConfigEmail**](ServerConfigEmail.md) |  | [optional] 

## Methods

### NewServerConfigImportVideos

`func NewServerConfigImportVideos() *ServerConfigImportVideos`

NewServerConfigImportVideos instantiates a new ServerConfigImportVideos object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerConfigImportVideosWithDefaults

`func NewServerConfigImportVideosWithDefaults() *ServerConfigImportVideos`

NewServerConfigImportVideosWithDefaults instantiates a new ServerConfigImportVideos object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHttp

`func (o *ServerConfigImportVideos) GetHttp() ServerConfigEmail`

GetHttp returns the Http field if non-nil, zero value otherwise.

### GetHttpOk

`func (o *ServerConfigImportVideos) GetHttpOk() (*ServerConfigEmail, bool)`

GetHttpOk returns a tuple with the Http field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttp

`func (o *ServerConfigImportVideos) SetHttp(v ServerConfigEmail)`

SetHttp sets Http field to given value.

### HasHttp

`func (o *ServerConfigImportVideos) HasHttp() bool`

HasHttp returns a boolean if a field has been set.

### GetTorrent

`func (o *ServerConfigImportVideos) GetTorrent() ServerConfigEmail`

GetTorrent returns the Torrent field if non-nil, zero value otherwise.

### GetTorrentOk

`func (o *ServerConfigImportVideos) GetTorrentOk() (*ServerConfigEmail, bool)`

GetTorrentOk returns a tuple with the Torrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTorrent

`func (o *ServerConfigImportVideos) SetTorrent(v ServerConfigEmail)`

SetTorrent sets Torrent field to given value.

### HasTorrent

`func (o *ServerConfigImportVideos) HasTorrent() bool`

HasTorrent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


