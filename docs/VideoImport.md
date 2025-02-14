# VideoImport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**TargetUrl** | Pointer to **string** | remote URL where to find the import&#39;s source video | [optional] 
**MagnetUri** | Pointer to **string** | magnet URI allowing to resolve the import&#39;s source video | [optional] 
**Torrentfile** | Pointer to ***os.File** | Torrent file containing only the video file | [optional] 
**TorrentName** | Pointer to **string** |  | [optional] [readonly] 
**State** | Pointer to [**VideoImportStateConstant**](VideoImportStateConstant.md) |  | [optional] [readonly] 
**Error** | Pointer to **string** |  | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Video** | Pointer to [**NullableVideo**](Video.md) |  | [optional] [readonly] 

## Methods

### NewVideoImport

`func NewVideoImport() *VideoImport`

NewVideoImport instantiates a new VideoImport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoImportWithDefaults

`func NewVideoImportWithDefaults() *VideoImport`

NewVideoImportWithDefaults instantiates a new VideoImport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VideoImport) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VideoImport) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VideoImport) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *VideoImport) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTargetUrl

`func (o *VideoImport) GetTargetUrl() string`

GetTargetUrl returns the TargetUrl field if non-nil, zero value otherwise.

### GetTargetUrlOk

`func (o *VideoImport) GetTargetUrlOk() (*string, bool)`

GetTargetUrlOk returns a tuple with the TargetUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUrl

`func (o *VideoImport) SetTargetUrl(v string)`

SetTargetUrl sets TargetUrl field to given value.

### HasTargetUrl

`func (o *VideoImport) HasTargetUrl() bool`

HasTargetUrl returns a boolean if a field has been set.

### GetMagnetUri

`func (o *VideoImport) GetMagnetUri() string`

GetMagnetUri returns the MagnetUri field if non-nil, zero value otherwise.

### GetMagnetUriOk

`func (o *VideoImport) GetMagnetUriOk() (*string, bool)`

GetMagnetUriOk returns a tuple with the MagnetUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMagnetUri

`func (o *VideoImport) SetMagnetUri(v string)`

SetMagnetUri sets MagnetUri field to given value.

### HasMagnetUri

`func (o *VideoImport) HasMagnetUri() bool`

HasMagnetUri returns a boolean if a field has been set.

### GetTorrentfile

`func (o *VideoImport) GetTorrentfile() *os.File`

GetTorrentfile returns the Torrentfile field if non-nil, zero value otherwise.

### GetTorrentfileOk

`func (o *VideoImport) GetTorrentfileOk() (**os.File, bool)`

GetTorrentfileOk returns a tuple with the Torrentfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTorrentfile

`func (o *VideoImport) SetTorrentfile(v *os.File)`

SetTorrentfile sets Torrentfile field to given value.

### HasTorrentfile

`func (o *VideoImport) HasTorrentfile() bool`

HasTorrentfile returns a boolean if a field has been set.

### GetTorrentName

`func (o *VideoImport) GetTorrentName() string`

GetTorrentName returns the TorrentName field if non-nil, zero value otherwise.

### GetTorrentNameOk

`func (o *VideoImport) GetTorrentNameOk() (*string, bool)`

GetTorrentNameOk returns a tuple with the TorrentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTorrentName

`func (o *VideoImport) SetTorrentName(v string)`

SetTorrentName sets TorrentName field to given value.

### HasTorrentName

`func (o *VideoImport) HasTorrentName() bool`

HasTorrentName returns a boolean if a field has been set.

### GetState

`func (o *VideoImport) GetState() VideoImportStateConstant`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VideoImport) GetStateOk() (*VideoImportStateConstant, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VideoImport) SetState(v VideoImportStateConstant)`

SetState sets State field to given value.

### HasState

`func (o *VideoImport) HasState() bool`

HasState returns a boolean if a field has been set.

### GetError

`func (o *VideoImport) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *VideoImport) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *VideoImport) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *VideoImport) HasError() bool`

HasError returns a boolean if a field has been set.

### GetCreatedAt

`func (o *VideoImport) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VideoImport) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VideoImport) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VideoImport) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *VideoImport) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VideoImport) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VideoImport) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *VideoImport) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVideo

`func (o *VideoImport) GetVideo() Video`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *VideoImport) GetVideoOk() (*Video, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *VideoImport) SetVideo(v Video)`

SetVideo sets Video field to given value.

### HasVideo

`func (o *VideoImport) HasVideo() bool`

HasVideo returns a boolean if a field has been set.

### SetVideoNil

`func (o *VideoImport) SetVideoNil(b bool)`

 SetVideoNil sets the value for Video to be an explicit nil

### UnsetVideo
`func (o *VideoImport) UnsetVideo()`

UnsetVideo ensures that no value is present for Video, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


