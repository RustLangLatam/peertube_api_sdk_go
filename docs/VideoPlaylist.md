# VideoPlaylist

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**ShortUUID** | Pointer to **string** | translation of a uuid v4 with a bigger alphabet to have a shorter uuid | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**IsLocal** | Pointer to **bool** |  | [optional] 
**VideoLength** | Pointer to **int32** |  | [optional] 
**ThumbnailPath** | Pointer to **string** |  | [optional] 
**Privacy** | Pointer to [**VideoPlaylistPrivacyConstant**](VideoPlaylistPrivacyConstant.md) |  | [optional] 
**Type** | Pointer to [**VideoPlaylistTypeConstant**](VideoPlaylistTypeConstant.md) |  | [optional] 
**OwnerAccount** | Pointer to [**AccountSummary**](AccountSummary.md) |  | [optional] 
**VideoChannel** | Pointer to [**VideoChannelSummary**](VideoChannelSummary.md) |  | [optional] 

## Methods

### NewVideoPlaylist

`func NewVideoPlaylist() *VideoPlaylist`

NewVideoPlaylist instantiates a new VideoPlaylist object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoPlaylistWithDefaults

`func NewVideoPlaylistWithDefaults() *VideoPlaylist`

NewVideoPlaylistWithDefaults instantiates a new VideoPlaylist object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VideoPlaylist) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VideoPlaylist) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VideoPlaylist) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *VideoPlaylist) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *VideoPlaylist) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *VideoPlaylist) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *VideoPlaylist) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *VideoPlaylist) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetShortUUID

`func (o *VideoPlaylist) GetShortUUID() string`

GetShortUUID returns the ShortUUID field if non-nil, zero value otherwise.

### GetShortUUIDOk

`func (o *VideoPlaylist) GetShortUUIDOk() (*string, bool)`

GetShortUUIDOk returns a tuple with the ShortUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortUUID

`func (o *VideoPlaylist) SetShortUUID(v string)`

SetShortUUID sets ShortUUID field to given value.

### HasShortUUID

`func (o *VideoPlaylist) HasShortUUID() bool`

HasShortUUID returns a boolean if a field has been set.

### GetCreatedAt

`func (o *VideoPlaylist) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VideoPlaylist) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VideoPlaylist) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VideoPlaylist) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *VideoPlaylist) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VideoPlaylist) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VideoPlaylist) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *VideoPlaylist) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *VideoPlaylist) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VideoPlaylist) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VideoPlaylist) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VideoPlaylist) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *VideoPlaylist) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *VideoPlaylist) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *VideoPlaylist) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *VideoPlaylist) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetIsLocal

`func (o *VideoPlaylist) GetIsLocal() bool`

GetIsLocal returns the IsLocal field if non-nil, zero value otherwise.

### GetIsLocalOk

`func (o *VideoPlaylist) GetIsLocalOk() (*bool, bool)`

GetIsLocalOk returns a tuple with the IsLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocal

`func (o *VideoPlaylist) SetIsLocal(v bool)`

SetIsLocal sets IsLocal field to given value.

### HasIsLocal

`func (o *VideoPlaylist) HasIsLocal() bool`

HasIsLocal returns a boolean if a field has been set.

### GetVideoLength

`func (o *VideoPlaylist) GetVideoLength() int32`

GetVideoLength returns the VideoLength field if non-nil, zero value otherwise.

### GetVideoLengthOk

`func (o *VideoPlaylist) GetVideoLengthOk() (*int32, bool)`

GetVideoLengthOk returns a tuple with the VideoLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoLength

`func (o *VideoPlaylist) SetVideoLength(v int32)`

SetVideoLength sets VideoLength field to given value.

### HasVideoLength

`func (o *VideoPlaylist) HasVideoLength() bool`

HasVideoLength returns a boolean if a field has been set.

### GetThumbnailPath

`func (o *VideoPlaylist) GetThumbnailPath() string`

GetThumbnailPath returns the ThumbnailPath field if non-nil, zero value otherwise.

### GetThumbnailPathOk

`func (o *VideoPlaylist) GetThumbnailPathOk() (*string, bool)`

GetThumbnailPathOk returns a tuple with the ThumbnailPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailPath

`func (o *VideoPlaylist) SetThumbnailPath(v string)`

SetThumbnailPath sets ThumbnailPath field to given value.

### HasThumbnailPath

`func (o *VideoPlaylist) HasThumbnailPath() bool`

HasThumbnailPath returns a boolean if a field has been set.

### GetPrivacy

`func (o *VideoPlaylist) GetPrivacy() VideoPlaylistPrivacyConstant`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *VideoPlaylist) GetPrivacyOk() (*VideoPlaylistPrivacyConstant, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *VideoPlaylist) SetPrivacy(v VideoPlaylistPrivacyConstant)`

SetPrivacy sets Privacy field to given value.

### HasPrivacy

`func (o *VideoPlaylist) HasPrivacy() bool`

HasPrivacy returns a boolean if a field has been set.

### GetType

`func (o *VideoPlaylist) GetType() VideoPlaylistTypeConstant`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VideoPlaylist) GetTypeOk() (*VideoPlaylistTypeConstant, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VideoPlaylist) SetType(v VideoPlaylistTypeConstant)`

SetType sets Type field to given value.

### HasType

`func (o *VideoPlaylist) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOwnerAccount

`func (o *VideoPlaylist) GetOwnerAccount() AccountSummary`

GetOwnerAccount returns the OwnerAccount field if non-nil, zero value otherwise.

### GetOwnerAccountOk

`func (o *VideoPlaylist) GetOwnerAccountOk() (*AccountSummary, bool)`

GetOwnerAccountOk returns a tuple with the OwnerAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerAccount

`func (o *VideoPlaylist) SetOwnerAccount(v AccountSummary)`

SetOwnerAccount sets OwnerAccount field to given value.

### HasOwnerAccount

`func (o *VideoPlaylist) HasOwnerAccount() bool`

HasOwnerAccount returns a boolean if a field has been set.

### GetVideoChannel

`func (o *VideoPlaylist) GetVideoChannel() VideoChannelSummary`

GetVideoChannel returns the VideoChannel field if non-nil, zero value otherwise.

### GetVideoChannelOk

`func (o *VideoPlaylist) GetVideoChannelOk() (*VideoChannelSummary, bool)`

GetVideoChannelOk returns a tuple with the VideoChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoChannel

`func (o *VideoPlaylist) SetVideoChannel(v VideoChannelSummary)`

SetVideoChannel sets VideoChannel field to given value.

### HasVideoChannel

`func (o *VideoPlaylist) HasVideoChannel() bool`

HasVideoChannel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


