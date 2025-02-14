# VideoDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | object id for the video | [optional] 
**Uuid** | Pointer to **string** | universal identifier for the video, that can be used across instances | [optional] 
**ShortUUID** | Pointer to **string** | translation of a uuid v4 with a bigger alphabet to have a shorter uuid | [optional] 
**IsLive** | Pointer to **bool** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | time at which the video object was first drafted | [optional] 
**PublishedAt** | Pointer to **time.Time** | time at which the video was marked as ready for playback (with restrictions depending on &#x60;privacy&#x60;). Usually set after a &#x60;state&#x60; evolution. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | last time the video&#39;s metadata was modified | [optional] 
**OriginallyPublishedAt** | Pointer to **NullableTime** | used to represent a date of first publication, prior to the practical publication date of &#x60;publishedAt&#x60; | [optional] 
**Category** | Pointer to [**VideoConstantNumberCategory**](VideoConstantNumberCategory.md) | category in which the video is classified | [optional] 
**Licence** | Pointer to [**VideoConstantNumberLicence**](VideoConstantNumberLicence.md) | licence under which the video is distributed | [optional] 
**Language** | Pointer to [**VideoConstantStringLanguage**](VideoConstantStringLanguage.md) | main language used in the video | [optional] 
**Privacy** | Pointer to [**VideoPrivacyConstant**](VideoPrivacyConstant.md) | privacy policy used to distribute the video | [optional] 
**TruncatedDescription** | Pointer to **NullableString** | truncated description of the video, written in Markdown.  | [optional] 
**Duration** | Pointer to **int32** | duration of the video in seconds | [optional] 
**AspectRatio** | Pointer to **NullableFloat32** | **PeerTube &gt;&#x3D; 6.1** Aspect ratio of the video stream | [optional] 
**IsLocal** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** | title of the video | [optional] 
**ThumbnailPath** | Pointer to **string** |  | [optional] 
**PreviewPath** | Pointer to **string** |  | [optional] 
**EmbedPath** | Pointer to **string** |  | [optional] 
**Views** | Pointer to **int32** |  | [optional] 
**Likes** | Pointer to **int32** |  | [optional] 
**Dislikes** | Pointer to **int32** |  | [optional] 
**Nsfw** | Pointer to **bool** |  | [optional] 
**WaitTranscoding** | Pointer to **NullableBool** |  | [optional] 
**State** | Pointer to [**VideoStateConstant**](VideoStateConstant.md) | represents the internal state of the video processing within the PeerTube instance | [optional] 
**ScheduledUpdate** | Pointer to [**NullableVideoScheduledUpdate**](VideoScheduledUpdate.md) |  | [optional] 
**Blacklisted** | Pointer to **NullableBool** |  | [optional] 
**BlacklistedReason** | Pointer to **NullableString** |  | [optional] 
**Account** | Pointer to [**Account**](Account.md) |  | [optional] 
**Channel** | Pointer to [**VideoChannel**](VideoChannel.md) |  | [optional] 
**UserHistory** | Pointer to [**NullableVideoUserHistory**](VideoUserHistory.md) |  | [optional] 
**Viewers** | Pointer to **int32** | If the video is a live, you have the amount of current viewers | [optional] 
**Description** | Pointer to **NullableString** | full description of the video, written in Markdown.  | [optional] 
**Support** | Pointer to **NullableString** | A text tell the audience how to support the video creator | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**CommentsEnabled** | Pointer to **bool** | Deprecated in 6.2, use commentsPolicy instead | [optional] 
**CommentsPolicy** | Pointer to [**VideoCommentsPolicyConstant**](VideoCommentsPolicyConstant.md) |  | [optional] 
**DownloadEnabled** | Pointer to **bool** |  | [optional] 
**InputFileUpdatedAt** | Pointer to **NullableTime** | Latest input file update. Null if the file has never been replaced since the original upload | [optional] 
**TrackerUrls** | Pointer to **[]string** |  | [optional] 
**Files** | Pointer to [**[]VideoFile**](VideoFile.md) | Web compatible video files. If Web Video is disabled on the server:  - field will be empty - video files will be found in &#x60;streamingPlaylists[].files&#x60; field  | [optional] 
**StreamingPlaylists** | Pointer to [**[]VideoStreamingPlaylists**](VideoStreamingPlaylists.md) | HLS playlists/manifest files. If HLS is disabled on the server:  - field will be empty - video files will be found in &#x60;files&#x60; field  | [optional] 

## Methods

### NewVideoDetails

`func NewVideoDetails() *VideoDetails`

NewVideoDetails instantiates a new VideoDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoDetailsWithDefaults

`func NewVideoDetailsWithDefaults() *VideoDetails`

NewVideoDetailsWithDefaults instantiates a new VideoDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VideoDetails) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VideoDetails) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VideoDetails) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *VideoDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *VideoDetails) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *VideoDetails) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *VideoDetails) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *VideoDetails) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetShortUUID

`func (o *VideoDetails) GetShortUUID() string`

GetShortUUID returns the ShortUUID field if non-nil, zero value otherwise.

### GetShortUUIDOk

`func (o *VideoDetails) GetShortUUIDOk() (*string, bool)`

GetShortUUIDOk returns a tuple with the ShortUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortUUID

`func (o *VideoDetails) SetShortUUID(v string)`

SetShortUUID sets ShortUUID field to given value.

### HasShortUUID

`func (o *VideoDetails) HasShortUUID() bool`

HasShortUUID returns a boolean if a field has been set.

### GetIsLive

`func (o *VideoDetails) GetIsLive() bool`

GetIsLive returns the IsLive field if non-nil, zero value otherwise.

### GetIsLiveOk

`func (o *VideoDetails) GetIsLiveOk() (*bool, bool)`

GetIsLiveOk returns a tuple with the IsLive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLive

`func (o *VideoDetails) SetIsLive(v bool)`

SetIsLive sets IsLive field to given value.

### HasIsLive

`func (o *VideoDetails) HasIsLive() bool`

HasIsLive returns a boolean if a field has been set.

### GetCreatedAt

`func (o *VideoDetails) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VideoDetails) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VideoDetails) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VideoDetails) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetPublishedAt

`func (o *VideoDetails) GetPublishedAt() time.Time`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *VideoDetails) GetPublishedAtOk() (*time.Time, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *VideoDetails) SetPublishedAt(v time.Time)`

SetPublishedAt sets PublishedAt field to given value.

### HasPublishedAt

`func (o *VideoDetails) HasPublishedAt() bool`

HasPublishedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *VideoDetails) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VideoDetails) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VideoDetails) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *VideoDetails) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetOriginallyPublishedAt

`func (o *VideoDetails) GetOriginallyPublishedAt() time.Time`

GetOriginallyPublishedAt returns the OriginallyPublishedAt field if non-nil, zero value otherwise.

### GetOriginallyPublishedAtOk

`func (o *VideoDetails) GetOriginallyPublishedAtOk() (*time.Time, bool)`

GetOriginallyPublishedAtOk returns a tuple with the OriginallyPublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginallyPublishedAt

`func (o *VideoDetails) SetOriginallyPublishedAt(v time.Time)`

SetOriginallyPublishedAt sets OriginallyPublishedAt field to given value.

### HasOriginallyPublishedAt

`func (o *VideoDetails) HasOriginallyPublishedAt() bool`

HasOriginallyPublishedAt returns a boolean if a field has been set.

### SetOriginallyPublishedAtNil

`func (o *VideoDetails) SetOriginallyPublishedAtNil(b bool)`

 SetOriginallyPublishedAtNil sets the value for OriginallyPublishedAt to be an explicit nil

### UnsetOriginallyPublishedAt
`func (o *VideoDetails) UnsetOriginallyPublishedAt()`

UnsetOriginallyPublishedAt ensures that no value is present for OriginallyPublishedAt, not even an explicit nil
### GetCategory

`func (o *VideoDetails) GetCategory() VideoConstantNumberCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *VideoDetails) GetCategoryOk() (*VideoConstantNumberCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *VideoDetails) SetCategory(v VideoConstantNumberCategory)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *VideoDetails) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetLicence

`func (o *VideoDetails) GetLicence() VideoConstantNumberLicence`

GetLicence returns the Licence field if non-nil, zero value otherwise.

### GetLicenceOk

`func (o *VideoDetails) GetLicenceOk() (*VideoConstantNumberLicence, bool)`

GetLicenceOk returns a tuple with the Licence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicence

`func (o *VideoDetails) SetLicence(v VideoConstantNumberLicence)`

SetLicence sets Licence field to given value.

### HasLicence

`func (o *VideoDetails) HasLicence() bool`

HasLicence returns a boolean if a field has been set.

### GetLanguage

`func (o *VideoDetails) GetLanguage() VideoConstantStringLanguage`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *VideoDetails) GetLanguageOk() (*VideoConstantStringLanguage, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *VideoDetails) SetLanguage(v VideoConstantStringLanguage)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *VideoDetails) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetPrivacy

`func (o *VideoDetails) GetPrivacy() VideoPrivacyConstant`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *VideoDetails) GetPrivacyOk() (*VideoPrivacyConstant, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *VideoDetails) SetPrivacy(v VideoPrivacyConstant)`

SetPrivacy sets Privacy field to given value.

### HasPrivacy

`func (o *VideoDetails) HasPrivacy() bool`

HasPrivacy returns a boolean if a field has been set.

### GetTruncatedDescription

`func (o *VideoDetails) GetTruncatedDescription() string`

GetTruncatedDescription returns the TruncatedDescription field if non-nil, zero value otherwise.

### GetTruncatedDescriptionOk

`func (o *VideoDetails) GetTruncatedDescriptionOk() (*string, bool)`

GetTruncatedDescriptionOk returns a tuple with the TruncatedDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncatedDescription

`func (o *VideoDetails) SetTruncatedDescription(v string)`

SetTruncatedDescription sets TruncatedDescription field to given value.

### HasTruncatedDescription

`func (o *VideoDetails) HasTruncatedDescription() bool`

HasTruncatedDescription returns a boolean if a field has been set.

### SetTruncatedDescriptionNil

`func (o *VideoDetails) SetTruncatedDescriptionNil(b bool)`

 SetTruncatedDescriptionNil sets the value for TruncatedDescription to be an explicit nil

### UnsetTruncatedDescription
`func (o *VideoDetails) UnsetTruncatedDescription()`

UnsetTruncatedDescription ensures that no value is present for TruncatedDescription, not even an explicit nil
### GetDuration

`func (o *VideoDetails) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *VideoDetails) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *VideoDetails) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *VideoDetails) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetAspectRatio

`func (o *VideoDetails) GetAspectRatio() float32`

GetAspectRatio returns the AspectRatio field if non-nil, zero value otherwise.

### GetAspectRatioOk

`func (o *VideoDetails) GetAspectRatioOk() (*float32, bool)`

GetAspectRatioOk returns a tuple with the AspectRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspectRatio

`func (o *VideoDetails) SetAspectRatio(v float32)`

SetAspectRatio sets AspectRatio field to given value.

### HasAspectRatio

`func (o *VideoDetails) HasAspectRatio() bool`

HasAspectRatio returns a boolean if a field has been set.

### SetAspectRatioNil

`func (o *VideoDetails) SetAspectRatioNil(b bool)`

 SetAspectRatioNil sets the value for AspectRatio to be an explicit nil

### UnsetAspectRatio
`func (o *VideoDetails) UnsetAspectRatio()`

UnsetAspectRatio ensures that no value is present for AspectRatio, not even an explicit nil
### GetIsLocal

`func (o *VideoDetails) GetIsLocal() bool`

GetIsLocal returns the IsLocal field if non-nil, zero value otherwise.

### GetIsLocalOk

`func (o *VideoDetails) GetIsLocalOk() (*bool, bool)`

GetIsLocalOk returns a tuple with the IsLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocal

`func (o *VideoDetails) SetIsLocal(v bool)`

SetIsLocal sets IsLocal field to given value.

### HasIsLocal

`func (o *VideoDetails) HasIsLocal() bool`

HasIsLocal returns a boolean if a field has been set.

### GetName

`func (o *VideoDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VideoDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VideoDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VideoDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetThumbnailPath

`func (o *VideoDetails) GetThumbnailPath() string`

GetThumbnailPath returns the ThumbnailPath field if non-nil, zero value otherwise.

### GetThumbnailPathOk

`func (o *VideoDetails) GetThumbnailPathOk() (*string, bool)`

GetThumbnailPathOk returns a tuple with the ThumbnailPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailPath

`func (o *VideoDetails) SetThumbnailPath(v string)`

SetThumbnailPath sets ThumbnailPath field to given value.

### HasThumbnailPath

`func (o *VideoDetails) HasThumbnailPath() bool`

HasThumbnailPath returns a boolean if a field has been set.

### GetPreviewPath

`func (o *VideoDetails) GetPreviewPath() string`

GetPreviewPath returns the PreviewPath field if non-nil, zero value otherwise.

### GetPreviewPathOk

`func (o *VideoDetails) GetPreviewPathOk() (*string, bool)`

GetPreviewPathOk returns a tuple with the PreviewPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewPath

`func (o *VideoDetails) SetPreviewPath(v string)`

SetPreviewPath sets PreviewPath field to given value.

### HasPreviewPath

`func (o *VideoDetails) HasPreviewPath() bool`

HasPreviewPath returns a boolean if a field has been set.

### GetEmbedPath

`func (o *VideoDetails) GetEmbedPath() string`

GetEmbedPath returns the EmbedPath field if non-nil, zero value otherwise.

### GetEmbedPathOk

`func (o *VideoDetails) GetEmbedPathOk() (*string, bool)`

GetEmbedPathOk returns a tuple with the EmbedPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedPath

`func (o *VideoDetails) SetEmbedPath(v string)`

SetEmbedPath sets EmbedPath field to given value.

### HasEmbedPath

`func (o *VideoDetails) HasEmbedPath() bool`

HasEmbedPath returns a boolean if a field has been set.

### GetViews

`func (o *VideoDetails) GetViews() int32`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *VideoDetails) GetViewsOk() (*int32, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *VideoDetails) SetViews(v int32)`

SetViews sets Views field to given value.

### HasViews

`func (o *VideoDetails) HasViews() bool`

HasViews returns a boolean if a field has been set.

### GetLikes

`func (o *VideoDetails) GetLikes() int32`

GetLikes returns the Likes field if non-nil, zero value otherwise.

### GetLikesOk

`func (o *VideoDetails) GetLikesOk() (*int32, bool)`

GetLikesOk returns a tuple with the Likes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikes

`func (o *VideoDetails) SetLikes(v int32)`

SetLikes sets Likes field to given value.

### HasLikes

`func (o *VideoDetails) HasLikes() bool`

HasLikes returns a boolean if a field has been set.

### GetDislikes

`func (o *VideoDetails) GetDislikes() int32`

GetDislikes returns the Dislikes field if non-nil, zero value otherwise.

### GetDislikesOk

`func (o *VideoDetails) GetDislikesOk() (*int32, bool)`

GetDislikesOk returns a tuple with the Dislikes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDislikes

`func (o *VideoDetails) SetDislikes(v int32)`

SetDislikes sets Dislikes field to given value.

### HasDislikes

`func (o *VideoDetails) HasDislikes() bool`

HasDislikes returns a boolean if a field has been set.

### GetNsfw

`func (o *VideoDetails) GetNsfw() bool`

GetNsfw returns the Nsfw field if non-nil, zero value otherwise.

### GetNsfwOk

`func (o *VideoDetails) GetNsfwOk() (*bool, bool)`

GetNsfwOk returns a tuple with the Nsfw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsfw

`func (o *VideoDetails) SetNsfw(v bool)`

SetNsfw sets Nsfw field to given value.

### HasNsfw

`func (o *VideoDetails) HasNsfw() bool`

HasNsfw returns a boolean if a field has been set.

### GetWaitTranscoding

`func (o *VideoDetails) GetWaitTranscoding() bool`

GetWaitTranscoding returns the WaitTranscoding field if non-nil, zero value otherwise.

### GetWaitTranscodingOk

`func (o *VideoDetails) GetWaitTranscodingOk() (*bool, bool)`

GetWaitTranscodingOk returns a tuple with the WaitTranscoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitTranscoding

`func (o *VideoDetails) SetWaitTranscoding(v bool)`

SetWaitTranscoding sets WaitTranscoding field to given value.

### HasWaitTranscoding

`func (o *VideoDetails) HasWaitTranscoding() bool`

HasWaitTranscoding returns a boolean if a field has been set.

### SetWaitTranscodingNil

`func (o *VideoDetails) SetWaitTranscodingNil(b bool)`

 SetWaitTranscodingNil sets the value for WaitTranscoding to be an explicit nil

### UnsetWaitTranscoding
`func (o *VideoDetails) UnsetWaitTranscoding()`

UnsetWaitTranscoding ensures that no value is present for WaitTranscoding, not even an explicit nil
### GetState

`func (o *VideoDetails) GetState() VideoStateConstant`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VideoDetails) GetStateOk() (*VideoStateConstant, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VideoDetails) SetState(v VideoStateConstant)`

SetState sets State field to given value.

### HasState

`func (o *VideoDetails) HasState() bool`

HasState returns a boolean if a field has been set.

### GetScheduledUpdate

`func (o *VideoDetails) GetScheduledUpdate() VideoScheduledUpdate`

GetScheduledUpdate returns the ScheduledUpdate field if non-nil, zero value otherwise.

### GetScheduledUpdateOk

`func (o *VideoDetails) GetScheduledUpdateOk() (*VideoScheduledUpdate, bool)`

GetScheduledUpdateOk returns a tuple with the ScheduledUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledUpdate

`func (o *VideoDetails) SetScheduledUpdate(v VideoScheduledUpdate)`

SetScheduledUpdate sets ScheduledUpdate field to given value.

### HasScheduledUpdate

`func (o *VideoDetails) HasScheduledUpdate() bool`

HasScheduledUpdate returns a boolean if a field has been set.

### SetScheduledUpdateNil

`func (o *VideoDetails) SetScheduledUpdateNil(b bool)`

 SetScheduledUpdateNil sets the value for ScheduledUpdate to be an explicit nil

### UnsetScheduledUpdate
`func (o *VideoDetails) UnsetScheduledUpdate()`

UnsetScheduledUpdate ensures that no value is present for ScheduledUpdate, not even an explicit nil
### GetBlacklisted

`func (o *VideoDetails) GetBlacklisted() bool`

GetBlacklisted returns the Blacklisted field if non-nil, zero value otherwise.

### GetBlacklistedOk

`func (o *VideoDetails) GetBlacklistedOk() (*bool, bool)`

GetBlacklistedOk returns a tuple with the Blacklisted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklisted

`func (o *VideoDetails) SetBlacklisted(v bool)`

SetBlacklisted sets Blacklisted field to given value.

### HasBlacklisted

`func (o *VideoDetails) HasBlacklisted() bool`

HasBlacklisted returns a boolean if a field has been set.

### SetBlacklistedNil

`func (o *VideoDetails) SetBlacklistedNil(b bool)`

 SetBlacklistedNil sets the value for Blacklisted to be an explicit nil

### UnsetBlacklisted
`func (o *VideoDetails) UnsetBlacklisted()`

UnsetBlacklisted ensures that no value is present for Blacklisted, not even an explicit nil
### GetBlacklistedReason

`func (o *VideoDetails) GetBlacklistedReason() string`

GetBlacklistedReason returns the BlacklistedReason field if non-nil, zero value otherwise.

### GetBlacklistedReasonOk

`func (o *VideoDetails) GetBlacklistedReasonOk() (*string, bool)`

GetBlacklistedReasonOk returns a tuple with the BlacklistedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklistedReason

`func (o *VideoDetails) SetBlacklistedReason(v string)`

SetBlacklistedReason sets BlacklistedReason field to given value.

### HasBlacklistedReason

`func (o *VideoDetails) HasBlacklistedReason() bool`

HasBlacklistedReason returns a boolean if a field has been set.

### SetBlacklistedReasonNil

`func (o *VideoDetails) SetBlacklistedReasonNil(b bool)`

 SetBlacklistedReasonNil sets the value for BlacklistedReason to be an explicit nil

### UnsetBlacklistedReason
`func (o *VideoDetails) UnsetBlacklistedReason()`

UnsetBlacklistedReason ensures that no value is present for BlacklistedReason, not even an explicit nil
### GetAccount

`func (o *VideoDetails) GetAccount() Account`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *VideoDetails) GetAccountOk() (*Account, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *VideoDetails) SetAccount(v Account)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *VideoDetails) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetChannel

`func (o *VideoDetails) GetChannel() VideoChannel`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *VideoDetails) GetChannelOk() (*VideoChannel, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *VideoDetails) SetChannel(v VideoChannel)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *VideoDetails) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetUserHistory

`func (o *VideoDetails) GetUserHistory() VideoUserHistory`

GetUserHistory returns the UserHistory field if non-nil, zero value otherwise.

### GetUserHistoryOk

`func (o *VideoDetails) GetUserHistoryOk() (*VideoUserHistory, bool)`

GetUserHistoryOk returns a tuple with the UserHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserHistory

`func (o *VideoDetails) SetUserHistory(v VideoUserHistory)`

SetUserHistory sets UserHistory field to given value.

### HasUserHistory

`func (o *VideoDetails) HasUserHistory() bool`

HasUserHistory returns a boolean if a field has been set.

### SetUserHistoryNil

`func (o *VideoDetails) SetUserHistoryNil(b bool)`

 SetUserHistoryNil sets the value for UserHistory to be an explicit nil

### UnsetUserHistory
`func (o *VideoDetails) UnsetUserHistory()`

UnsetUserHistory ensures that no value is present for UserHistory, not even an explicit nil
### GetViewers

`func (o *VideoDetails) GetViewers() int32`

GetViewers returns the Viewers field if non-nil, zero value otherwise.

### GetViewersOk

`func (o *VideoDetails) GetViewersOk() (*int32, bool)`

GetViewersOk returns a tuple with the Viewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewers

`func (o *VideoDetails) SetViewers(v int32)`

SetViewers sets Viewers field to given value.

### HasViewers

`func (o *VideoDetails) HasViewers() bool`

HasViewers returns a boolean if a field has been set.

### GetDescription

`func (o *VideoDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VideoDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VideoDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VideoDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *VideoDetails) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *VideoDetails) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSupport

`func (o *VideoDetails) GetSupport() string`

GetSupport returns the Support field if non-nil, zero value otherwise.

### GetSupportOk

`func (o *VideoDetails) GetSupportOk() (*string, bool)`

GetSupportOk returns a tuple with the Support field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport

`func (o *VideoDetails) SetSupport(v string)`

SetSupport sets Support field to given value.

### HasSupport

`func (o *VideoDetails) HasSupport() bool`

HasSupport returns a boolean if a field has been set.

### SetSupportNil

`func (o *VideoDetails) SetSupportNil(b bool)`

 SetSupportNil sets the value for Support to be an explicit nil

### UnsetSupport
`func (o *VideoDetails) UnsetSupport()`

UnsetSupport ensures that no value is present for Support, not even an explicit nil
### GetTags

`func (o *VideoDetails) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VideoDetails) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VideoDetails) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VideoDetails) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCommentsEnabled

`func (o *VideoDetails) GetCommentsEnabled() bool`

GetCommentsEnabled returns the CommentsEnabled field if non-nil, zero value otherwise.

### GetCommentsEnabledOk

`func (o *VideoDetails) GetCommentsEnabledOk() (*bool, bool)`

GetCommentsEnabledOk returns a tuple with the CommentsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentsEnabled

`func (o *VideoDetails) SetCommentsEnabled(v bool)`

SetCommentsEnabled sets CommentsEnabled field to given value.

### HasCommentsEnabled

`func (o *VideoDetails) HasCommentsEnabled() bool`

HasCommentsEnabled returns a boolean if a field has been set.

### GetCommentsPolicy

`func (o *VideoDetails) GetCommentsPolicy() VideoCommentsPolicyConstant`

GetCommentsPolicy returns the CommentsPolicy field if non-nil, zero value otherwise.

### GetCommentsPolicyOk

`func (o *VideoDetails) GetCommentsPolicyOk() (*VideoCommentsPolicyConstant, bool)`

GetCommentsPolicyOk returns a tuple with the CommentsPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentsPolicy

`func (o *VideoDetails) SetCommentsPolicy(v VideoCommentsPolicyConstant)`

SetCommentsPolicy sets CommentsPolicy field to given value.

### HasCommentsPolicy

`func (o *VideoDetails) HasCommentsPolicy() bool`

HasCommentsPolicy returns a boolean if a field has been set.

### GetDownloadEnabled

`func (o *VideoDetails) GetDownloadEnabled() bool`

GetDownloadEnabled returns the DownloadEnabled field if non-nil, zero value otherwise.

### GetDownloadEnabledOk

`func (o *VideoDetails) GetDownloadEnabledOk() (*bool, bool)`

GetDownloadEnabledOk returns a tuple with the DownloadEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadEnabled

`func (o *VideoDetails) SetDownloadEnabled(v bool)`

SetDownloadEnabled sets DownloadEnabled field to given value.

### HasDownloadEnabled

`func (o *VideoDetails) HasDownloadEnabled() bool`

HasDownloadEnabled returns a boolean if a field has been set.

### GetInputFileUpdatedAt

`func (o *VideoDetails) GetInputFileUpdatedAt() time.Time`

GetInputFileUpdatedAt returns the InputFileUpdatedAt field if non-nil, zero value otherwise.

### GetInputFileUpdatedAtOk

`func (o *VideoDetails) GetInputFileUpdatedAtOk() (*time.Time, bool)`

GetInputFileUpdatedAtOk returns a tuple with the InputFileUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputFileUpdatedAt

`func (o *VideoDetails) SetInputFileUpdatedAt(v time.Time)`

SetInputFileUpdatedAt sets InputFileUpdatedAt field to given value.

### HasInputFileUpdatedAt

`func (o *VideoDetails) HasInputFileUpdatedAt() bool`

HasInputFileUpdatedAt returns a boolean if a field has been set.

### SetInputFileUpdatedAtNil

`func (o *VideoDetails) SetInputFileUpdatedAtNil(b bool)`

 SetInputFileUpdatedAtNil sets the value for InputFileUpdatedAt to be an explicit nil

### UnsetInputFileUpdatedAt
`func (o *VideoDetails) UnsetInputFileUpdatedAt()`

UnsetInputFileUpdatedAt ensures that no value is present for InputFileUpdatedAt, not even an explicit nil
### GetTrackerUrls

`func (o *VideoDetails) GetTrackerUrls() []string`

GetTrackerUrls returns the TrackerUrls field if non-nil, zero value otherwise.

### GetTrackerUrlsOk

`func (o *VideoDetails) GetTrackerUrlsOk() (*[]string, bool)`

GetTrackerUrlsOk returns a tuple with the TrackerUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackerUrls

`func (o *VideoDetails) SetTrackerUrls(v []string)`

SetTrackerUrls sets TrackerUrls field to given value.

### HasTrackerUrls

`func (o *VideoDetails) HasTrackerUrls() bool`

HasTrackerUrls returns a boolean if a field has been set.

### GetFiles

`func (o *VideoDetails) GetFiles() []VideoFile`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *VideoDetails) GetFilesOk() (*[]VideoFile, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *VideoDetails) SetFiles(v []VideoFile)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *VideoDetails) HasFiles() bool`

HasFiles returns a boolean if a field has been set.

### GetStreamingPlaylists

`func (o *VideoDetails) GetStreamingPlaylists() []VideoStreamingPlaylists`

GetStreamingPlaylists returns the StreamingPlaylists field if non-nil, zero value otherwise.

### GetStreamingPlaylistsOk

`func (o *VideoDetails) GetStreamingPlaylistsOk() (*[]VideoStreamingPlaylists, bool)`

GetStreamingPlaylistsOk returns a tuple with the StreamingPlaylists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamingPlaylists

`func (o *VideoDetails) SetStreamingPlaylists(v []VideoStreamingPlaylists)`

SetStreamingPlaylists sets StreamingPlaylists field to given value.

### HasStreamingPlaylists

`func (o *VideoDetails) HasStreamingPlaylists() bool`

HasStreamingPlaylists returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


