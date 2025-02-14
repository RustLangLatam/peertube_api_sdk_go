# Video

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
**Account** | Pointer to [**AccountSummary**](AccountSummary.md) |  | [optional] 
**Channel** | Pointer to [**VideoChannelSummary**](VideoChannelSummary.md) |  | [optional] 
**UserHistory** | Pointer to [**NullableVideoUserHistory**](VideoUserHistory.md) |  | [optional] 

## Methods

### NewVideo

`func NewVideo() *Video`

NewVideo instantiates a new Video object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoWithDefaults

`func NewVideoWithDefaults() *Video`

NewVideoWithDefaults instantiates a new Video object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Video) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Video) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Video) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Video) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *Video) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Video) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Video) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Video) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetShortUUID

`func (o *Video) GetShortUUID() string`

GetShortUUID returns the ShortUUID field if non-nil, zero value otherwise.

### GetShortUUIDOk

`func (o *Video) GetShortUUIDOk() (*string, bool)`

GetShortUUIDOk returns a tuple with the ShortUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortUUID

`func (o *Video) SetShortUUID(v string)`

SetShortUUID sets ShortUUID field to given value.

### HasShortUUID

`func (o *Video) HasShortUUID() bool`

HasShortUUID returns a boolean if a field has been set.

### GetIsLive

`func (o *Video) GetIsLive() bool`

GetIsLive returns the IsLive field if non-nil, zero value otherwise.

### GetIsLiveOk

`func (o *Video) GetIsLiveOk() (*bool, bool)`

GetIsLiveOk returns a tuple with the IsLive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLive

`func (o *Video) SetIsLive(v bool)`

SetIsLive sets IsLive field to given value.

### HasIsLive

`func (o *Video) HasIsLive() bool`

HasIsLive returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Video) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Video) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Video) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Video) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetPublishedAt

`func (o *Video) GetPublishedAt() time.Time`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *Video) GetPublishedAtOk() (*time.Time, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *Video) SetPublishedAt(v time.Time)`

SetPublishedAt sets PublishedAt field to given value.

### HasPublishedAt

`func (o *Video) HasPublishedAt() bool`

HasPublishedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Video) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Video) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Video) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Video) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetOriginallyPublishedAt

`func (o *Video) GetOriginallyPublishedAt() time.Time`

GetOriginallyPublishedAt returns the OriginallyPublishedAt field if non-nil, zero value otherwise.

### GetOriginallyPublishedAtOk

`func (o *Video) GetOriginallyPublishedAtOk() (*time.Time, bool)`

GetOriginallyPublishedAtOk returns a tuple with the OriginallyPublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginallyPublishedAt

`func (o *Video) SetOriginallyPublishedAt(v time.Time)`

SetOriginallyPublishedAt sets OriginallyPublishedAt field to given value.

### HasOriginallyPublishedAt

`func (o *Video) HasOriginallyPublishedAt() bool`

HasOriginallyPublishedAt returns a boolean if a field has been set.

### SetOriginallyPublishedAtNil

`func (o *Video) SetOriginallyPublishedAtNil(b bool)`

 SetOriginallyPublishedAtNil sets the value for OriginallyPublishedAt to be an explicit nil

### UnsetOriginallyPublishedAt
`func (o *Video) UnsetOriginallyPublishedAt()`

UnsetOriginallyPublishedAt ensures that no value is present for OriginallyPublishedAt, not even an explicit nil
### GetCategory

`func (o *Video) GetCategory() VideoConstantNumberCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Video) GetCategoryOk() (*VideoConstantNumberCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Video) SetCategory(v VideoConstantNumberCategory)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Video) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetLicence

`func (o *Video) GetLicence() VideoConstantNumberLicence`

GetLicence returns the Licence field if non-nil, zero value otherwise.

### GetLicenceOk

`func (o *Video) GetLicenceOk() (*VideoConstantNumberLicence, bool)`

GetLicenceOk returns a tuple with the Licence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicence

`func (o *Video) SetLicence(v VideoConstantNumberLicence)`

SetLicence sets Licence field to given value.

### HasLicence

`func (o *Video) HasLicence() bool`

HasLicence returns a boolean if a field has been set.

### GetLanguage

`func (o *Video) GetLanguage() VideoConstantStringLanguage`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *Video) GetLanguageOk() (*VideoConstantStringLanguage, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *Video) SetLanguage(v VideoConstantStringLanguage)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *Video) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetPrivacy

`func (o *Video) GetPrivacy() VideoPrivacyConstant`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *Video) GetPrivacyOk() (*VideoPrivacyConstant, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *Video) SetPrivacy(v VideoPrivacyConstant)`

SetPrivacy sets Privacy field to given value.

### HasPrivacy

`func (o *Video) HasPrivacy() bool`

HasPrivacy returns a boolean if a field has been set.

### GetTruncatedDescription

`func (o *Video) GetTruncatedDescription() string`

GetTruncatedDescription returns the TruncatedDescription field if non-nil, zero value otherwise.

### GetTruncatedDescriptionOk

`func (o *Video) GetTruncatedDescriptionOk() (*string, bool)`

GetTruncatedDescriptionOk returns a tuple with the TruncatedDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncatedDescription

`func (o *Video) SetTruncatedDescription(v string)`

SetTruncatedDescription sets TruncatedDescription field to given value.

### HasTruncatedDescription

`func (o *Video) HasTruncatedDescription() bool`

HasTruncatedDescription returns a boolean if a field has been set.

### SetTruncatedDescriptionNil

`func (o *Video) SetTruncatedDescriptionNil(b bool)`

 SetTruncatedDescriptionNil sets the value for TruncatedDescription to be an explicit nil

### UnsetTruncatedDescription
`func (o *Video) UnsetTruncatedDescription()`

UnsetTruncatedDescription ensures that no value is present for TruncatedDescription, not even an explicit nil
### GetDuration

`func (o *Video) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *Video) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *Video) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *Video) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetAspectRatio

`func (o *Video) GetAspectRatio() float32`

GetAspectRatio returns the AspectRatio field if non-nil, zero value otherwise.

### GetAspectRatioOk

`func (o *Video) GetAspectRatioOk() (*float32, bool)`

GetAspectRatioOk returns a tuple with the AspectRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspectRatio

`func (o *Video) SetAspectRatio(v float32)`

SetAspectRatio sets AspectRatio field to given value.

### HasAspectRatio

`func (o *Video) HasAspectRatio() bool`

HasAspectRatio returns a boolean if a field has been set.

### SetAspectRatioNil

`func (o *Video) SetAspectRatioNil(b bool)`

 SetAspectRatioNil sets the value for AspectRatio to be an explicit nil

### UnsetAspectRatio
`func (o *Video) UnsetAspectRatio()`

UnsetAspectRatio ensures that no value is present for AspectRatio, not even an explicit nil
### GetIsLocal

`func (o *Video) GetIsLocal() bool`

GetIsLocal returns the IsLocal field if non-nil, zero value otherwise.

### GetIsLocalOk

`func (o *Video) GetIsLocalOk() (*bool, bool)`

GetIsLocalOk returns a tuple with the IsLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocal

`func (o *Video) SetIsLocal(v bool)`

SetIsLocal sets IsLocal field to given value.

### HasIsLocal

`func (o *Video) HasIsLocal() bool`

HasIsLocal returns a boolean if a field has been set.

### GetName

`func (o *Video) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Video) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Video) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Video) HasName() bool`

HasName returns a boolean if a field has been set.

### GetThumbnailPath

`func (o *Video) GetThumbnailPath() string`

GetThumbnailPath returns the ThumbnailPath field if non-nil, zero value otherwise.

### GetThumbnailPathOk

`func (o *Video) GetThumbnailPathOk() (*string, bool)`

GetThumbnailPathOk returns a tuple with the ThumbnailPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailPath

`func (o *Video) SetThumbnailPath(v string)`

SetThumbnailPath sets ThumbnailPath field to given value.

### HasThumbnailPath

`func (o *Video) HasThumbnailPath() bool`

HasThumbnailPath returns a boolean if a field has been set.

### GetPreviewPath

`func (o *Video) GetPreviewPath() string`

GetPreviewPath returns the PreviewPath field if non-nil, zero value otherwise.

### GetPreviewPathOk

`func (o *Video) GetPreviewPathOk() (*string, bool)`

GetPreviewPathOk returns a tuple with the PreviewPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewPath

`func (o *Video) SetPreviewPath(v string)`

SetPreviewPath sets PreviewPath field to given value.

### HasPreviewPath

`func (o *Video) HasPreviewPath() bool`

HasPreviewPath returns a boolean if a field has been set.

### GetEmbedPath

`func (o *Video) GetEmbedPath() string`

GetEmbedPath returns the EmbedPath field if non-nil, zero value otherwise.

### GetEmbedPathOk

`func (o *Video) GetEmbedPathOk() (*string, bool)`

GetEmbedPathOk returns a tuple with the EmbedPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedPath

`func (o *Video) SetEmbedPath(v string)`

SetEmbedPath sets EmbedPath field to given value.

### HasEmbedPath

`func (o *Video) HasEmbedPath() bool`

HasEmbedPath returns a boolean if a field has been set.

### GetViews

`func (o *Video) GetViews() int32`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *Video) GetViewsOk() (*int32, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *Video) SetViews(v int32)`

SetViews sets Views field to given value.

### HasViews

`func (o *Video) HasViews() bool`

HasViews returns a boolean if a field has been set.

### GetLikes

`func (o *Video) GetLikes() int32`

GetLikes returns the Likes field if non-nil, zero value otherwise.

### GetLikesOk

`func (o *Video) GetLikesOk() (*int32, bool)`

GetLikesOk returns a tuple with the Likes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikes

`func (o *Video) SetLikes(v int32)`

SetLikes sets Likes field to given value.

### HasLikes

`func (o *Video) HasLikes() bool`

HasLikes returns a boolean if a field has been set.

### GetDislikes

`func (o *Video) GetDislikes() int32`

GetDislikes returns the Dislikes field if non-nil, zero value otherwise.

### GetDislikesOk

`func (o *Video) GetDislikesOk() (*int32, bool)`

GetDislikesOk returns a tuple with the Dislikes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDislikes

`func (o *Video) SetDislikes(v int32)`

SetDislikes sets Dislikes field to given value.

### HasDislikes

`func (o *Video) HasDislikes() bool`

HasDislikes returns a boolean if a field has been set.

### GetNsfw

`func (o *Video) GetNsfw() bool`

GetNsfw returns the Nsfw field if non-nil, zero value otherwise.

### GetNsfwOk

`func (o *Video) GetNsfwOk() (*bool, bool)`

GetNsfwOk returns a tuple with the Nsfw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsfw

`func (o *Video) SetNsfw(v bool)`

SetNsfw sets Nsfw field to given value.

### HasNsfw

`func (o *Video) HasNsfw() bool`

HasNsfw returns a boolean if a field has been set.

### GetWaitTranscoding

`func (o *Video) GetWaitTranscoding() bool`

GetWaitTranscoding returns the WaitTranscoding field if non-nil, zero value otherwise.

### GetWaitTranscodingOk

`func (o *Video) GetWaitTranscodingOk() (*bool, bool)`

GetWaitTranscodingOk returns a tuple with the WaitTranscoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitTranscoding

`func (o *Video) SetWaitTranscoding(v bool)`

SetWaitTranscoding sets WaitTranscoding field to given value.

### HasWaitTranscoding

`func (o *Video) HasWaitTranscoding() bool`

HasWaitTranscoding returns a boolean if a field has been set.

### SetWaitTranscodingNil

`func (o *Video) SetWaitTranscodingNil(b bool)`

 SetWaitTranscodingNil sets the value for WaitTranscoding to be an explicit nil

### UnsetWaitTranscoding
`func (o *Video) UnsetWaitTranscoding()`

UnsetWaitTranscoding ensures that no value is present for WaitTranscoding, not even an explicit nil
### GetState

`func (o *Video) GetState() VideoStateConstant`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Video) GetStateOk() (*VideoStateConstant, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Video) SetState(v VideoStateConstant)`

SetState sets State field to given value.

### HasState

`func (o *Video) HasState() bool`

HasState returns a boolean if a field has been set.

### GetScheduledUpdate

`func (o *Video) GetScheduledUpdate() VideoScheduledUpdate`

GetScheduledUpdate returns the ScheduledUpdate field if non-nil, zero value otherwise.

### GetScheduledUpdateOk

`func (o *Video) GetScheduledUpdateOk() (*VideoScheduledUpdate, bool)`

GetScheduledUpdateOk returns a tuple with the ScheduledUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledUpdate

`func (o *Video) SetScheduledUpdate(v VideoScheduledUpdate)`

SetScheduledUpdate sets ScheduledUpdate field to given value.

### HasScheduledUpdate

`func (o *Video) HasScheduledUpdate() bool`

HasScheduledUpdate returns a boolean if a field has been set.

### SetScheduledUpdateNil

`func (o *Video) SetScheduledUpdateNil(b bool)`

 SetScheduledUpdateNil sets the value for ScheduledUpdate to be an explicit nil

### UnsetScheduledUpdate
`func (o *Video) UnsetScheduledUpdate()`

UnsetScheduledUpdate ensures that no value is present for ScheduledUpdate, not even an explicit nil
### GetBlacklisted

`func (o *Video) GetBlacklisted() bool`

GetBlacklisted returns the Blacklisted field if non-nil, zero value otherwise.

### GetBlacklistedOk

`func (o *Video) GetBlacklistedOk() (*bool, bool)`

GetBlacklistedOk returns a tuple with the Blacklisted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklisted

`func (o *Video) SetBlacklisted(v bool)`

SetBlacklisted sets Blacklisted field to given value.

### HasBlacklisted

`func (o *Video) HasBlacklisted() bool`

HasBlacklisted returns a boolean if a field has been set.

### SetBlacklistedNil

`func (o *Video) SetBlacklistedNil(b bool)`

 SetBlacklistedNil sets the value for Blacklisted to be an explicit nil

### UnsetBlacklisted
`func (o *Video) UnsetBlacklisted()`

UnsetBlacklisted ensures that no value is present for Blacklisted, not even an explicit nil
### GetBlacklistedReason

`func (o *Video) GetBlacklistedReason() string`

GetBlacklistedReason returns the BlacklistedReason field if non-nil, zero value otherwise.

### GetBlacklistedReasonOk

`func (o *Video) GetBlacklistedReasonOk() (*string, bool)`

GetBlacklistedReasonOk returns a tuple with the BlacklistedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklistedReason

`func (o *Video) SetBlacklistedReason(v string)`

SetBlacklistedReason sets BlacklistedReason field to given value.

### HasBlacklistedReason

`func (o *Video) HasBlacklistedReason() bool`

HasBlacklistedReason returns a boolean if a field has been set.

### SetBlacklistedReasonNil

`func (o *Video) SetBlacklistedReasonNil(b bool)`

 SetBlacklistedReasonNil sets the value for BlacklistedReason to be an explicit nil

### UnsetBlacklistedReason
`func (o *Video) UnsetBlacklistedReason()`

UnsetBlacklistedReason ensures that no value is present for BlacklistedReason, not even an explicit nil
### GetAccount

`func (o *Video) GetAccount() AccountSummary`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Video) GetAccountOk() (*AccountSummary, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Video) SetAccount(v AccountSummary)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *Video) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetChannel

`func (o *Video) GetChannel() VideoChannelSummary`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *Video) GetChannelOk() (*VideoChannelSummary, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *Video) SetChannel(v VideoChannelSummary)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *Video) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetUserHistory

`func (o *Video) GetUserHistory() VideoUserHistory`

GetUserHistory returns the UserHistory field if non-nil, zero value otherwise.

### GetUserHistoryOk

`func (o *Video) GetUserHistoryOk() (*VideoUserHistory, bool)`

GetUserHistoryOk returns a tuple with the UserHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserHistory

`func (o *Video) SetUserHistory(v VideoUserHistory)`

SetUserHistory sets UserHistory field to given value.

### HasUserHistory

`func (o *Video) HasUserHistory() bool`

HasUserHistory returns a boolean if a field has been set.

### SetUserHistoryNil

`func (o *Video) SetUserHistoryNil(b bool)`

 SetUserHistoryNil sets the value for UserHistory to be an explicit nil

### UnsetUserHistory
`func (o *Video) UnsetUserHistory()`

UnsetUserHistory ensures that no value is present for UserHistory, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


