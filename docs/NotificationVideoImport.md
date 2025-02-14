# NotificationVideoImport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Video** | Pointer to [**VideoInfo**](VideoInfo.md) |  | [optional] 
**TorrentName** | Pointer to **NullableString** |  | [optional] 
**MagnetUri** | Pointer to **string** | magnet URI allowing to resolve the import&#39;s source video | [optional] 
**TargetUri** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewNotificationVideoImport

`func NewNotificationVideoImport() *NotificationVideoImport`

NewNotificationVideoImport instantiates a new NotificationVideoImport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationVideoImportWithDefaults

`func NewNotificationVideoImportWithDefaults() *NotificationVideoImport`

NewNotificationVideoImportWithDefaults instantiates a new NotificationVideoImport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NotificationVideoImport) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationVideoImport) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationVideoImport) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *NotificationVideoImport) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVideo

`func (o *NotificationVideoImport) GetVideo() VideoInfo`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *NotificationVideoImport) GetVideoOk() (*VideoInfo, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *NotificationVideoImport) SetVideo(v VideoInfo)`

SetVideo sets Video field to given value.

### HasVideo

`func (o *NotificationVideoImport) HasVideo() bool`

HasVideo returns a boolean if a field has been set.

### GetTorrentName

`func (o *NotificationVideoImport) GetTorrentName() string`

GetTorrentName returns the TorrentName field if non-nil, zero value otherwise.

### GetTorrentNameOk

`func (o *NotificationVideoImport) GetTorrentNameOk() (*string, bool)`

GetTorrentNameOk returns a tuple with the TorrentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTorrentName

`func (o *NotificationVideoImport) SetTorrentName(v string)`

SetTorrentName sets TorrentName field to given value.

### HasTorrentName

`func (o *NotificationVideoImport) HasTorrentName() bool`

HasTorrentName returns a boolean if a field has been set.

### SetTorrentNameNil

`func (o *NotificationVideoImport) SetTorrentNameNil(b bool)`

 SetTorrentNameNil sets the value for TorrentName to be an explicit nil

### UnsetTorrentName
`func (o *NotificationVideoImport) UnsetTorrentName()`

UnsetTorrentName ensures that no value is present for TorrentName, not even an explicit nil
### GetMagnetUri

`func (o *NotificationVideoImport) GetMagnetUri() string`

GetMagnetUri returns the MagnetUri field if non-nil, zero value otherwise.

### GetMagnetUriOk

`func (o *NotificationVideoImport) GetMagnetUriOk() (*string, bool)`

GetMagnetUriOk returns a tuple with the MagnetUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMagnetUri

`func (o *NotificationVideoImport) SetMagnetUri(v string)`

SetMagnetUri sets MagnetUri field to given value.

### HasMagnetUri

`func (o *NotificationVideoImport) HasMagnetUri() bool`

HasMagnetUri returns a boolean if a field has been set.

### GetTargetUri

`func (o *NotificationVideoImport) GetTargetUri() string`

GetTargetUri returns the TargetUri field if non-nil, zero value otherwise.

### GetTargetUriOk

`func (o *NotificationVideoImport) GetTargetUriOk() (*string, bool)`

GetTargetUriOk returns a tuple with the TargetUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUri

`func (o *NotificationVideoImport) SetTargetUri(v string)`

SetTargetUri sets TargetUri field to given value.

### HasTargetUri

`func (o *NotificationVideoImport) HasTargetUri() bool`

HasTargetUri returns a boolean if a field has been set.

### SetTargetUriNil

`func (o *NotificationVideoImport) SetTargetUriNil(b bool)`

 SetTargetUriNil sets the value for TargetUri to be an explicit nil

### UnsetTargetUri
`func (o *NotificationVideoImport) UnsetTargetUri()`

UnsetTargetUri ensures that no value is present for TargetUri, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


