# UserWithStats

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

### NewUserWithStats

`func NewUserWithStats() *UserWithStats`

NewUserWithStats instantiates a new UserWithStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithStatsWithDefaults

`func NewUserWithStatsWithDefaults() *UserWithStats`

NewUserWithStatsWithDefaults instantiates a new UserWithStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *UserWithStats) GetAccount() Account`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *UserWithStats) GetAccountOk() (*Account, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *UserWithStats) SetAccount(v Account)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *UserWithStats) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAutoPlayNextVideo

`func (o *UserWithStats) GetAutoPlayNextVideo() bool`

GetAutoPlayNextVideo returns the AutoPlayNextVideo field if non-nil, zero value otherwise.

### GetAutoPlayNextVideoOk

`func (o *UserWithStats) GetAutoPlayNextVideoOk() (*bool, bool)`

GetAutoPlayNextVideoOk returns a tuple with the AutoPlayNextVideo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPlayNextVideo

`func (o *UserWithStats) SetAutoPlayNextVideo(v bool)`

SetAutoPlayNextVideo sets AutoPlayNextVideo field to given value.

### HasAutoPlayNextVideo

`func (o *UserWithStats) HasAutoPlayNextVideo() bool`

HasAutoPlayNextVideo returns a boolean if a field has been set.

### GetAutoPlayNextVideoPlaylist

`func (o *UserWithStats) GetAutoPlayNextVideoPlaylist() bool`

GetAutoPlayNextVideoPlaylist returns the AutoPlayNextVideoPlaylist field if non-nil, zero value otherwise.

### GetAutoPlayNextVideoPlaylistOk

`func (o *UserWithStats) GetAutoPlayNextVideoPlaylistOk() (*bool, bool)`

GetAutoPlayNextVideoPlaylistOk returns a tuple with the AutoPlayNextVideoPlaylist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPlayNextVideoPlaylist

`func (o *UserWithStats) SetAutoPlayNextVideoPlaylist(v bool)`

SetAutoPlayNextVideoPlaylist sets AutoPlayNextVideoPlaylist field to given value.

### HasAutoPlayNextVideoPlaylist

`func (o *UserWithStats) HasAutoPlayNextVideoPlaylist() bool`

HasAutoPlayNextVideoPlaylist returns a boolean if a field has been set.

### GetAutoPlayVideo

`func (o *UserWithStats) GetAutoPlayVideo() bool`

GetAutoPlayVideo returns the AutoPlayVideo field if non-nil, zero value otherwise.

### GetAutoPlayVideoOk

`func (o *UserWithStats) GetAutoPlayVideoOk() (*bool, bool)`

GetAutoPlayVideoOk returns a tuple with the AutoPlayVideo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPlayVideo

`func (o *UserWithStats) SetAutoPlayVideo(v bool)`

SetAutoPlayVideo sets AutoPlayVideo field to given value.

### HasAutoPlayVideo

`func (o *UserWithStats) HasAutoPlayVideo() bool`

HasAutoPlayVideo returns a boolean if a field has been set.

### GetBlocked

`func (o *UserWithStats) GetBlocked() bool`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *UserWithStats) GetBlockedOk() (*bool, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *UserWithStats) SetBlocked(v bool)`

SetBlocked sets Blocked field to given value.

### HasBlocked

`func (o *UserWithStats) HasBlocked() bool`

HasBlocked returns a boolean if a field has been set.

### GetBlockedReason

`func (o *UserWithStats) GetBlockedReason() string`

GetBlockedReason returns the BlockedReason field if non-nil, zero value otherwise.

### GetBlockedReasonOk

`func (o *UserWithStats) GetBlockedReasonOk() (*string, bool)`

GetBlockedReasonOk returns a tuple with the BlockedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedReason

`func (o *UserWithStats) SetBlockedReason(v string)`

SetBlockedReason sets BlockedReason field to given value.

### HasBlockedReason

`func (o *UserWithStats) HasBlockedReason() bool`

HasBlockedReason returns a boolean if a field has been set.

### GetCreatedAt

`func (o *UserWithStats) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserWithStats) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserWithStats) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *UserWithStats) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEmail

`func (o *UserWithStats) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserWithStats) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserWithStats) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserWithStats) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEmailVerified

`func (o *UserWithStats) GetEmailVerified() bool`

GetEmailVerified returns the EmailVerified field if non-nil, zero value otherwise.

### GetEmailVerifiedOk

`func (o *UserWithStats) GetEmailVerifiedOk() (*bool, bool)`

GetEmailVerifiedOk returns a tuple with the EmailVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailVerified

`func (o *UserWithStats) SetEmailVerified(v bool)`

SetEmailVerified sets EmailVerified field to given value.

### HasEmailVerified

`func (o *UserWithStats) HasEmailVerified() bool`

HasEmailVerified returns a boolean if a field has been set.

### GetId

`func (o *UserWithStats) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserWithStats) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserWithStats) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UserWithStats) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPluginAuth

`func (o *UserWithStats) GetPluginAuth() string`

GetPluginAuth returns the PluginAuth field if non-nil, zero value otherwise.

### GetPluginAuthOk

`func (o *UserWithStats) GetPluginAuthOk() (*string, bool)`

GetPluginAuthOk returns a tuple with the PluginAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginAuth

`func (o *UserWithStats) SetPluginAuth(v string)`

SetPluginAuth sets PluginAuth field to given value.

### HasPluginAuth

`func (o *UserWithStats) HasPluginAuth() bool`

HasPluginAuth returns a boolean if a field has been set.

### GetLastLoginDate

`func (o *UserWithStats) GetLastLoginDate() time.Time`

GetLastLoginDate returns the LastLoginDate field if non-nil, zero value otherwise.

### GetLastLoginDateOk

`func (o *UserWithStats) GetLastLoginDateOk() (*time.Time, bool)`

GetLastLoginDateOk returns a tuple with the LastLoginDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginDate

`func (o *UserWithStats) SetLastLoginDate(v time.Time)`

SetLastLoginDate sets LastLoginDate field to given value.

### HasLastLoginDate

`func (o *UserWithStats) HasLastLoginDate() bool`

HasLastLoginDate returns a boolean if a field has been set.

### GetNoInstanceConfigWarningModal

`func (o *UserWithStats) GetNoInstanceConfigWarningModal() bool`

GetNoInstanceConfigWarningModal returns the NoInstanceConfigWarningModal field if non-nil, zero value otherwise.

### GetNoInstanceConfigWarningModalOk

`func (o *UserWithStats) GetNoInstanceConfigWarningModalOk() (*bool, bool)`

GetNoInstanceConfigWarningModalOk returns a tuple with the NoInstanceConfigWarningModal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoInstanceConfigWarningModal

`func (o *UserWithStats) SetNoInstanceConfigWarningModal(v bool)`

SetNoInstanceConfigWarningModal sets NoInstanceConfigWarningModal field to given value.

### HasNoInstanceConfigWarningModal

`func (o *UserWithStats) HasNoInstanceConfigWarningModal() bool`

HasNoInstanceConfigWarningModal returns a boolean if a field has been set.

### GetNoAccountSetupWarningModal

`func (o *UserWithStats) GetNoAccountSetupWarningModal() bool`

GetNoAccountSetupWarningModal returns the NoAccountSetupWarningModal field if non-nil, zero value otherwise.

### GetNoAccountSetupWarningModalOk

`func (o *UserWithStats) GetNoAccountSetupWarningModalOk() (*bool, bool)`

GetNoAccountSetupWarningModalOk returns a tuple with the NoAccountSetupWarningModal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoAccountSetupWarningModal

`func (o *UserWithStats) SetNoAccountSetupWarningModal(v bool)`

SetNoAccountSetupWarningModal sets NoAccountSetupWarningModal field to given value.

### HasNoAccountSetupWarningModal

`func (o *UserWithStats) HasNoAccountSetupWarningModal() bool`

HasNoAccountSetupWarningModal returns a boolean if a field has been set.

### GetNoWelcomeModal

`func (o *UserWithStats) GetNoWelcomeModal() bool`

GetNoWelcomeModal returns the NoWelcomeModal field if non-nil, zero value otherwise.

### GetNoWelcomeModalOk

`func (o *UserWithStats) GetNoWelcomeModalOk() (*bool, bool)`

GetNoWelcomeModalOk returns a tuple with the NoWelcomeModal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoWelcomeModal

`func (o *UserWithStats) SetNoWelcomeModal(v bool)`

SetNoWelcomeModal sets NoWelcomeModal field to given value.

### HasNoWelcomeModal

`func (o *UserWithStats) HasNoWelcomeModal() bool`

HasNoWelcomeModal returns a boolean if a field has been set.

### GetNsfwPolicy

`func (o *UserWithStats) GetNsfwPolicy() NSFWPolicy`

GetNsfwPolicy returns the NsfwPolicy field if non-nil, zero value otherwise.

### GetNsfwPolicyOk

`func (o *UserWithStats) GetNsfwPolicyOk() (*NSFWPolicy, bool)`

GetNsfwPolicyOk returns a tuple with the NsfwPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsfwPolicy

`func (o *UserWithStats) SetNsfwPolicy(v NSFWPolicy)`

SetNsfwPolicy sets NsfwPolicy field to given value.

### HasNsfwPolicy

`func (o *UserWithStats) HasNsfwPolicy() bool`

HasNsfwPolicy returns a boolean if a field has been set.

### GetRole

`func (o *UserWithStats) GetRole() UserRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UserWithStats) GetRoleOk() (*UserRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UserWithStats) SetRole(v UserRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *UserWithStats) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetTheme

`func (o *UserWithStats) GetTheme() string`

GetTheme returns the Theme field if non-nil, zero value otherwise.

### GetThemeOk

`func (o *UserWithStats) GetThemeOk() (*string, bool)`

GetThemeOk returns a tuple with the Theme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheme

`func (o *UserWithStats) SetTheme(v string)`

SetTheme sets Theme field to given value.

### HasTheme

`func (o *UserWithStats) HasTheme() bool`

HasTheme returns a boolean if a field has been set.

### GetUsername

`func (o *UserWithStats) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserWithStats) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserWithStats) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UserWithStats) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetVideoChannels

`func (o *UserWithStats) GetVideoChannels() []VideoChannel`

GetVideoChannels returns the VideoChannels field if non-nil, zero value otherwise.

### GetVideoChannelsOk

`func (o *UserWithStats) GetVideoChannelsOk() (*[]VideoChannel, bool)`

GetVideoChannelsOk returns a tuple with the VideoChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoChannels

`func (o *UserWithStats) SetVideoChannels(v []VideoChannel)`

SetVideoChannels sets VideoChannels field to given value.

### HasVideoChannels

`func (o *UserWithStats) HasVideoChannels() bool`

HasVideoChannels returns a boolean if a field has been set.

### GetVideoQuota

`func (o *UserWithStats) GetVideoQuota() int32`

GetVideoQuota returns the VideoQuota field if non-nil, zero value otherwise.

### GetVideoQuotaOk

`func (o *UserWithStats) GetVideoQuotaOk() (*int32, bool)`

GetVideoQuotaOk returns a tuple with the VideoQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoQuota

`func (o *UserWithStats) SetVideoQuota(v int32)`

SetVideoQuota sets VideoQuota field to given value.

### HasVideoQuota

`func (o *UserWithStats) HasVideoQuota() bool`

HasVideoQuota returns a boolean if a field has been set.

### GetVideoQuotaDaily

`func (o *UserWithStats) GetVideoQuotaDaily() int32`

GetVideoQuotaDaily returns the VideoQuotaDaily field if non-nil, zero value otherwise.

### GetVideoQuotaDailyOk

`func (o *UserWithStats) GetVideoQuotaDailyOk() (*int32, bool)`

GetVideoQuotaDailyOk returns a tuple with the VideoQuotaDaily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoQuotaDaily

`func (o *UserWithStats) SetVideoQuotaDaily(v int32)`

SetVideoQuotaDaily sets VideoQuotaDaily field to given value.

### HasVideoQuotaDaily

`func (o *UserWithStats) HasVideoQuotaDaily() bool`

HasVideoQuotaDaily returns a boolean if a field has been set.

### GetP2pEnabled

`func (o *UserWithStats) GetP2pEnabled() bool`

GetP2pEnabled returns the P2pEnabled field if non-nil, zero value otherwise.

### GetP2pEnabledOk

`func (o *UserWithStats) GetP2pEnabledOk() (*bool, bool)`

GetP2pEnabledOk returns a tuple with the P2pEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP2pEnabled

`func (o *UserWithStats) SetP2pEnabled(v bool)`

SetP2pEnabled sets P2pEnabled field to given value.

### HasP2pEnabled

`func (o *UserWithStats) HasP2pEnabled() bool`

HasP2pEnabled returns a boolean if a field has been set.

### GetVideosCount

`func (o *UserWithStats) GetVideosCount() int32`

GetVideosCount returns the VideosCount field if non-nil, zero value otherwise.

### GetVideosCountOk

`func (o *UserWithStats) GetVideosCountOk() (*int32, bool)`

GetVideosCountOk returns a tuple with the VideosCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideosCount

`func (o *UserWithStats) SetVideosCount(v int32)`

SetVideosCount sets VideosCount field to given value.

### HasVideosCount

`func (o *UserWithStats) HasVideosCount() bool`

HasVideosCount returns a boolean if a field has been set.

### GetAbusesCount

`func (o *UserWithStats) GetAbusesCount() int32`

GetAbusesCount returns the AbusesCount field if non-nil, zero value otherwise.

### GetAbusesCountOk

`func (o *UserWithStats) GetAbusesCountOk() (*int32, bool)`

GetAbusesCountOk returns a tuple with the AbusesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbusesCount

`func (o *UserWithStats) SetAbusesCount(v int32)`

SetAbusesCount sets AbusesCount field to given value.

### HasAbusesCount

`func (o *UserWithStats) HasAbusesCount() bool`

HasAbusesCount returns a boolean if a field has been set.

### GetAbusesAcceptedCount

`func (o *UserWithStats) GetAbusesAcceptedCount() int32`

GetAbusesAcceptedCount returns the AbusesAcceptedCount field if non-nil, zero value otherwise.

### GetAbusesAcceptedCountOk

`func (o *UserWithStats) GetAbusesAcceptedCountOk() (*int32, bool)`

GetAbusesAcceptedCountOk returns a tuple with the AbusesAcceptedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbusesAcceptedCount

`func (o *UserWithStats) SetAbusesAcceptedCount(v int32)`

SetAbusesAcceptedCount sets AbusesAcceptedCount field to given value.

### HasAbusesAcceptedCount

`func (o *UserWithStats) HasAbusesAcceptedCount() bool`

HasAbusesAcceptedCount returns a boolean if a field has been set.

### GetAbusesCreatedCount

`func (o *UserWithStats) GetAbusesCreatedCount() int32`

GetAbusesCreatedCount returns the AbusesCreatedCount field if non-nil, zero value otherwise.

### GetAbusesCreatedCountOk

`func (o *UserWithStats) GetAbusesCreatedCountOk() (*int32, bool)`

GetAbusesCreatedCountOk returns a tuple with the AbusesCreatedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbusesCreatedCount

`func (o *UserWithStats) SetAbusesCreatedCount(v int32)`

SetAbusesCreatedCount sets AbusesCreatedCount field to given value.

### HasAbusesCreatedCount

`func (o *UserWithStats) HasAbusesCreatedCount() bool`

HasAbusesCreatedCount returns a boolean if a field has been set.

### GetVideoCommentsCount

`func (o *UserWithStats) GetVideoCommentsCount() int32`

GetVideoCommentsCount returns the VideoCommentsCount field if non-nil, zero value otherwise.

### GetVideoCommentsCountOk

`func (o *UserWithStats) GetVideoCommentsCountOk() (*int32, bool)`

GetVideoCommentsCountOk returns a tuple with the VideoCommentsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoCommentsCount

`func (o *UserWithStats) SetVideoCommentsCount(v int32)`

SetVideoCommentsCount sets VideoCommentsCount field to given value.

### HasVideoCommentsCount

`func (o *UserWithStats) HasVideoCommentsCount() bool`

HasVideoCommentsCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


