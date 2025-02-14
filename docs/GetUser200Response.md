# GetUser200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**Account**](Account.md) |  | [optional] 
**AutoPlayNextVideo** | Pointer to **bool** | Automatically start playing the upcoming video after the currently playing video | [optional] 
**AutoPlayNextVideoPlaylist** | Pointer to **bool** | Automatically start playing the video on the playlist after the currently playing video | [optional] 
**AutoPlayVideo** | Pointer to **bool** | Automatically start playing the video on the watch page | [optional] 
**Blocked** | Pointer to **bool** |  | [optional] 
**BlockedReason** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** | The user email | [optional] 
**EmailVerified** | Pointer to **bool** | Has the user confirmed their email address? | [optional] 
**Id** | Pointer to **int32** |  | [optional] [readonly] 
**PluginAuth** | Pointer to **string** | Auth plugin to use to authenticate the user | [optional] 
**LastLoginDate** | Pointer to **time.Time** |  | [optional] 
**NoInstanceConfigWarningModal** | Pointer to **bool** |  | [optional] 
**NoAccountSetupWarningModal** | Pointer to **bool** |  | [optional] 
**NoWelcomeModal** | Pointer to **bool** |  | [optional] 
**NsfwPolicy** | Pointer to [**NSFWPolicy**](NSFWPolicy.md) |  | [optional] 
**Role** | Pointer to [**UserRole**](UserRole.md) |  | [optional] 
**Theme** | Pointer to **string** | Theme enabled by this user | [optional] 
**Username** | Pointer to **string** | immutable name of the user, used to find or mention its actor | [optional] 
**VideoChannels** | Pointer to [**[]VideoChannel**](VideoChannel.md) |  | [optional] 
**VideoQuota** | Pointer to **int32** | The user video quota in bytes | [optional] 
**VideoQuotaDaily** | Pointer to **int32** | The user daily video quota in bytes | [optional] 
**P2pEnabled** | Pointer to **bool** | Enable P2P in the player | [optional] 
**VideosCount** | Pointer to **int32** | Count of videos published | [optional] 
**AbusesCount** | Pointer to **int32** | Count of reports/abuses of which the user is a target | [optional] 
**AbusesAcceptedCount** | Pointer to **int32** | Count of reports/abuses created by the user and accepted/acted upon by the moderation team | [optional] 
**AbusesCreatedCount** | Pointer to **int32** | Count of reports/abuses created by the user | [optional] 
**VideoCommentsCount** | Pointer to **int32** | Count of comments published | [optional] 

## Methods

### NewGetUser200Response

`func NewGetUser200Response() *GetUser200Response`

NewGetUser200Response instantiates a new GetUser200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUser200ResponseWithDefaults

`func NewGetUser200ResponseWithDefaults() *GetUser200Response`

NewGetUser200ResponseWithDefaults instantiates a new GetUser200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *GetUser200Response) GetAccount() Account`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetUser200Response) GetAccountOk() (*Account, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetUser200Response) SetAccount(v Account)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetUser200Response) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAutoPlayNextVideo

`func (o *GetUser200Response) GetAutoPlayNextVideo() bool`

GetAutoPlayNextVideo returns the AutoPlayNextVideo field if non-nil, zero value otherwise.

### GetAutoPlayNextVideoOk

`func (o *GetUser200Response) GetAutoPlayNextVideoOk() (*bool, bool)`

GetAutoPlayNextVideoOk returns a tuple with the AutoPlayNextVideo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPlayNextVideo

`func (o *GetUser200Response) SetAutoPlayNextVideo(v bool)`

SetAutoPlayNextVideo sets AutoPlayNextVideo field to given value.

### HasAutoPlayNextVideo

`func (o *GetUser200Response) HasAutoPlayNextVideo() bool`

HasAutoPlayNextVideo returns a boolean if a field has been set.

### GetAutoPlayNextVideoPlaylist

`func (o *GetUser200Response) GetAutoPlayNextVideoPlaylist() bool`

GetAutoPlayNextVideoPlaylist returns the AutoPlayNextVideoPlaylist field if non-nil, zero value otherwise.

### GetAutoPlayNextVideoPlaylistOk

`func (o *GetUser200Response) GetAutoPlayNextVideoPlaylistOk() (*bool, bool)`

GetAutoPlayNextVideoPlaylistOk returns a tuple with the AutoPlayNextVideoPlaylist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPlayNextVideoPlaylist

`func (o *GetUser200Response) SetAutoPlayNextVideoPlaylist(v bool)`

SetAutoPlayNextVideoPlaylist sets AutoPlayNextVideoPlaylist field to given value.

### HasAutoPlayNextVideoPlaylist

`func (o *GetUser200Response) HasAutoPlayNextVideoPlaylist() bool`

HasAutoPlayNextVideoPlaylist returns a boolean if a field has been set.

### GetAutoPlayVideo

`func (o *GetUser200Response) GetAutoPlayVideo() bool`

GetAutoPlayVideo returns the AutoPlayVideo field if non-nil, zero value otherwise.

### GetAutoPlayVideoOk

`func (o *GetUser200Response) GetAutoPlayVideoOk() (*bool, bool)`

GetAutoPlayVideoOk returns a tuple with the AutoPlayVideo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPlayVideo

`func (o *GetUser200Response) SetAutoPlayVideo(v bool)`

SetAutoPlayVideo sets AutoPlayVideo field to given value.

### HasAutoPlayVideo

`func (o *GetUser200Response) HasAutoPlayVideo() bool`

HasAutoPlayVideo returns a boolean if a field has been set.

### GetBlocked

`func (o *GetUser200Response) GetBlocked() bool`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *GetUser200Response) GetBlockedOk() (*bool, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *GetUser200Response) SetBlocked(v bool)`

SetBlocked sets Blocked field to given value.

### HasBlocked

`func (o *GetUser200Response) HasBlocked() bool`

HasBlocked returns a boolean if a field has been set.

### GetBlockedReason

`func (o *GetUser200Response) GetBlockedReason() string`

GetBlockedReason returns the BlockedReason field if non-nil, zero value otherwise.

### GetBlockedReasonOk

`func (o *GetUser200Response) GetBlockedReasonOk() (*string, bool)`

GetBlockedReasonOk returns a tuple with the BlockedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedReason

`func (o *GetUser200Response) SetBlockedReason(v string)`

SetBlockedReason sets BlockedReason field to given value.

### HasBlockedReason

`func (o *GetUser200Response) HasBlockedReason() bool`

HasBlockedReason returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetUser200Response) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetUser200Response) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetUser200Response) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetUser200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEmail

`func (o *GetUser200Response) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetUser200Response) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetUser200Response) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GetUser200Response) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEmailVerified

`func (o *GetUser200Response) GetEmailVerified() bool`

GetEmailVerified returns the EmailVerified field if non-nil, zero value otherwise.

### GetEmailVerifiedOk

`func (o *GetUser200Response) GetEmailVerifiedOk() (*bool, bool)`

GetEmailVerifiedOk returns a tuple with the EmailVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailVerified

`func (o *GetUser200Response) SetEmailVerified(v bool)`

SetEmailVerified sets EmailVerified field to given value.

### HasEmailVerified

`func (o *GetUser200Response) HasEmailVerified() bool`

HasEmailVerified returns a boolean if a field has been set.

### GetId

`func (o *GetUser200Response) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetUser200Response) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetUser200Response) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetUser200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPluginAuth

`func (o *GetUser200Response) GetPluginAuth() string`

GetPluginAuth returns the PluginAuth field if non-nil, zero value otherwise.

### GetPluginAuthOk

`func (o *GetUser200Response) GetPluginAuthOk() (*string, bool)`

GetPluginAuthOk returns a tuple with the PluginAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginAuth

`func (o *GetUser200Response) SetPluginAuth(v string)`

SetPluginAuth sets PluginAuth field to given value.

### HasPluginAuth

`func (o *GetUser200Response) HasPluginAuth() bool`

HasPluginAuth returns a boolean if a field has been set.

### GetLastLoginDate

`func (o *GetUser200Response) GetLastLoginDate() time.Time`

GetLastLoginDate returns the LastLoginDate field if non-nil, zero value otherwise.

### GetLastLoginDateOk

`func (o *GetUser200Response) GetLastLoginDateOk() (*time.Time, bool)`

GetLastLoginDateOk returns a tuple with the LastLoginDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginDate

`func (o *GetUser200Response) SetLastLoginDate(v time.Time)`

SetLastLoginDate sets LastLoginDate field to given value.

### HasLastLoginDate

`func (o *GetUser200Response) HasLastLoginDate() bool`

HasLastLoginDate returns a boolean if a field has been set.

### GetNoInstanceConfigWarningModal

`func (o *GetUser200Response) GetNoInstanceConfigWarningModal() bool`

GetNoInstanceConfigWarningModal returns the NoInstanceConfigWarningModal field if non-nil, zero value otherwise.

### GetNoInstanceConfigWarningModalOk

`func (o *GetUser200Response) GetNoInstanceConfigWarningModalOk() (*bool, bool)`

GetNoInstanceConfigWarningModalOk returns a tuple with the NoInstanceConfigWarningModal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoInstanceConfigWarningModal

`func (o *GetUser200Response) SetNoInstanceConfigWarningModal(v bool)`

SetNoInstanceConfigWarningModal sets NoInstanceConfigWarningModal field to given value.

### HasNoInstanceConfigWarningModal

`func (o *GetUser200Response) HasNoInstanceConfigWarningModal() bool`

HasNoInstanceConfigWarningModal returns a boolean if a field has been set.

### GetNoAccountSetupWarningModal

`func (o *GetUser200Response) GetNoAccountSetupWarningModal() bool`

GetNoAccountSetupWarningModal returns the NoAccountSetupWarningModal field if non-nil, zero value otherwise.

### GetNoAccountSetupWarningModalOk

`func (o *GetUser200Response) GetNoAccountSetupWarningModalOk() (*bool, bool)`

GetNoAccountSetupWarningModalOk returns a tuple with the NoAccountSetupWarningModal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoAccountSetupWarningModal

`func (o *GetUser200Response) SetNoAccountSetupWarningModal(v bool)`

SetNoAccountSetupWarningModal sets NoAccountSetupWarningModal field to given value.

### HasNoAccountSetupWarningModal

`func (o *GetUser200Response) HasNoAccountSetupWarningModal() bool`

HasNoAccountSetupWarningModal returns a boolean if a field has been set.

### GetNoWelcomeModal

`func (o *GetUser200Response) GetNoWelcomeModal() bool`

GetNoWelcomeModal returns the NoWelcomeModal field if non-nil, zero value otherwise.

### GetNoWelcomeModalOk

`func (o *GetUser200Response) GetNoWelcomeModalOk() (*bool, bool)`

GetNoWelcomeModalOk returns a tuple with the NoWelcomeModal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoWelcomeModal

`func (o *GetUser200Response) SetNoWelcomeModal(v bool)`

SetNoWelcomeModal sets NoWelcomeModal field to given value.

### HasNoWelcomeModal

`func (o *GetUser200Response) HasNoWelcomeModal() bool`

HasNoWelcomeModal returns a boolean if a field has been set.

### GetNsfwPolicy

`func (o *GetUser200Response) GetNsfwPolicy() NSFWPolicy`

GetNsfwPolicy returns the NsfwPolicy field if non-nil, zero value otherwise.

### GetNsfwPolicyOk

`func (o *GetUser200Response) GetNsfwPolicyOk() (*NSFWPolicy, bool)`

GetNsfwPolicyOk returns a tuple with the NsfwPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsfwPolicy

`func (o *GetUser200Response) SetNsfwPolicy(v NSFWPolicy)`

SetNsfwPolicy sets NsfwPolicy field to given value.

### HasNsfwPolicy

`func (o *GetUser200Response) HasNsfwPolicy() bool`

HasNsfwPolicy returns a boolean if a field has been set.

### GetRole

`func (o *GetUser200Response) GetRole() UserRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *GetUser200Response) GetRoleOk() (*UserRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *GetUser200Response) SetRole(v UserRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *GetUser200Response) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetTheme

`func (o *GetUser200Response) GetTheme() string`

GetTheme returns the Theme field if non-nil, zero value otherwise.

### GetThemeOk

`func (o *GetUser200Response) GetThemeOk() (*string, bool)`

GetThemeOk returns a tuple with the Theme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheme

`func (o *GetUser200Response) SetTheme(v string)`

SetTheme sets Theme field to given value.

### HasTheme

`func (o *GetUser200Response) HasTheme() bool`

HasTheme returns a boolean if a field has been set.

### GetUsername

`func (o *GetUser200Response) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GetUser200Response) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GetUser200Response) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GetUser200Response) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetVideoChannels

`func (o *GetUser200Response) GetVideoChannels() []VideoChannel`

GetVideoChannels returns the VideoChannels field if non-nil, zero value otherwise.

### GetVideoChannelsOk

`func (o *GetUser200Response) GetVideoChannelsOk() (*[]VideoChannel, bool)`

GetVideoChannelsOk returns a tuple with the VideoChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoChannels

`func (o *GetUser200Response) SetVideoChannels(v []VideoChannel)`

SetVideoChannels sets VideoChannels field to given value.

### HasVideoChannels

`func (o *GetUser200Response) HasVideoChannels() bool`

HasVideoChannels returns a boolean if a field has been set.

### GetVideoQuota

`func (o *GetUser200Response) GetVideoQuota() int32`

GetVideoQuota returns the VideoQuota field if non-nil, zero value otherwise.

### GetVideoQuotaOk

`func (o *GetUser200Response) GetVideoQuotaOk() (*int32, bool)`

GetVideoQuotaOk returns a tuple with the VideoQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoQuota

`func (o *GetUser200Response) SetVideoQuota(v int32)`

SetVideoQuota sets VideoQuota field to given value.

### HasVideoQuota

`func (o *GetUser200Response) HasVideoQuota() bool`

HasVideoQuota returns a boolean if a field has been set.

### GetVideoQuotaDaily

`func (o *GetUser200Response) GetVideoQuotaDaily() int32`

GetVideoQuotaDaily returns the VideoQuotaDaily field if non-nil, zero value otherwise.

### GetVideoQuotaDailyOk

`func (o *GetUser200Response) GetVideoQuotaDailyOk() (*int32, bool)`

GetVideoQuotaDailyOk returns a tuple with the VideoQuotaDaily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoQuotaDaily

`func (o *GetUser200Response) SetVideoQuotaDaily(v int32)`

SetVideoQuotaDaily sets VideoQuotaDaily field to given value.

### HasVideoQuotaDaily

`func (o *GetUser200Response) HasVideoQuotaDaily() bool`

HasVideoQuotaDaily returns a boolean if a field has been set.

### GetP2pEnabled

`func (o *GetUser200Response) GetP2pEnabled() bool`

GetP2pEnabled returns the P2pEnabled field if non-nil, zero value otherwise.

### GetP2pEnabledOk

`func (o *GetUser200Response) GetP2pEnabledOk() (*bool, bool)`

GetP2pEnabledOk returns a tuple with the P2pEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP2pEnabled

`func (o *GetUser200Response) SetP2pEnabled(v bool)`

SetP2pEnabled sets P2pEnabled field to given value.

### HasP2pEnabled

`func (o *GetUser200Response) HasP2pEnabled() bool`

HasP2pEnabled returns a boolean if a field has been set.

### GetVideosCount

`func (o *GetUser200Response) GetVideosCount() int32`

GetVideosCount returns the VideosCount field if non-nil, zero value otherwise.

### GetVideosCountOk

`func (o *GetUser200Response) GetVideosCountOk() (*int32, bool)`

GetVideosCountOk returns a tuple with the VideosCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideosCount

`func (o *GetUser200Response) SetVideosCount(v int32)`

SetVideosCount sets VideosCount field to given value.

### HasVideosCount

`func (o *GetUser200Response) HasVideosCount() bool`

HasVideosCount returns a boolean if a field has been set.

### GetAbusesCount

`func (o *GetUser200Response) GetAbusesCount() int32`

GetAbusesCount returns the AbusesCount field if non-nil, zero value otherwise.

### GetAbusesCountOk

`func (o *GetUser200Response) GetAbusesCountOk() (*int32, bool)`

GetAbusesCountOk returns a tuple with the AbusesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbusesCount

`func (o *GetUser200Response) SetAbusesCount(v int32)`

SetAbusesCount sets AbusesCount field to given value.

### HasAbusesCount

`func (o *GetUser200Response) HasAbusesCount() bool`

HasAbusesCount returns a boolean if a field has been set.

### GetAbusesAcceptedCount

`func (o *GetUser200Response) GetAbusesAcceptedCount() int32`

GetAbusesAcceptedCount returns the AbusesAcceptedCount field if non-nil, zero value otherwise.

### GetAbusesAcceptedCountOk

`func (o *GetUser200Response) GetAbusesAcceptedCountOk() (*int32, bool)`

GetAbusesAcceptedCountOk returns a tuple with the AbusesAcceptedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbusesAcceptedCount

`func (o *GetUser200Response) SetAbusesAcceptedCount(v int32)`

SetAbusesAcceptedCount sets AbusesAcceptedCount field to given value.

### HasAbusesAcceptedCount

`func (o *GetUser200Response) HasAbusesAcceptedCount() bool`

HasAbusesAcceptedCount returns a boolean if a field has been set.

### GetAbusesCreatedCount

`func (o *GetUser200Response) GetAbusesCreatedCount() int32`

GetAbusesCreatedCount returns the AbusesCreatedCount field if non-nil, zero value otherwise.

### GetAbusesCreatedCountOk

`func (o *GetUser200Response) GetAbusesCreatedCountOk() (*int32, bool)`

GetAbusesCreatedCountOk returns a tuple with the AbusesCreatedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbusesCreatedCount

`func (o *GetUser200Response) SetAbusesCreatedCount(v int32)`

SetAbusesCreatedCount sets AbusesCreatedCount field to given value.

### HasAbusesCreatedCount

`func (o *GetUser200Response) HasAbusesCreatedCount() bool`

HasAbusesCreatedCount returns a boolean if a field has been set.

### GetVideoCommentsCount

`func (o *GetUser200Response) GetVideoCommentsCount() int32`

GetVideoCommentsCount returns the VideoCommentsCount field if non-nil, zero value otherwise.

### GetVideoCommentsCountOk

`func (o *GetUser200Response) GetVideoCommentsCountOk() (*int32, bool)`

GetVideoCommentsCountOk returns a tuple with the VideoCommentsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoCommentsCount

`func (o *GetUser200Response) SetVideoCommentsCount(v int32)`

SetVideoCommentsCount sets VideoCommentsCount field to given value.

### HasVideoCommentsCount

`func (o *GetUser200Response) HasVideoCommentsCount() bool`

HasVideoCommentsCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


