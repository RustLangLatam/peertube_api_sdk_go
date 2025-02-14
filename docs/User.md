# User

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

## Methods

### NewUser

`func NewUser() *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *User) GetAccount() Account`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *User) GetAccountOk() (*Account, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *User) SetAccount(v Account)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *User) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAutoPlayNextVideo

`func (o *User) GetAutoPlayNextVideo() bool`

GetAutoPlayNextVideo returns the AutoPlayNextVideo field if non-nil, zero value otherwise.

### GetAutoPlayNextVideoOk

`func (o *User) GetAutoPlayNextVideoOk() (*bool, bool)`

GetAutoPlayNextVideoOk returns a tuple with the AutoPlayNextVideo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPlayNextVideo

`func (o *User) SetAutoPlayNextVideo(v bool)`

SetAutoPlayNextVideo sets AutoPlayNextVideo field to given value.

### HasAutoPlayNextVideo

`func (o *User) HasAutoPlayNextVideo() bool`

HasAutoPlayNextVideo returns a boolean if a field has been set.

### GetAutoPlayNextVideoPlaylist

`func (o *User) GetAutoPlayNextVideoPlaylist() bool`

GetAutoPlayNextVideoPlaylist returns the AutoPlayNextVideoPlaylist field if non-nil, zero value otherwise.

### GetAutoPlayNextVideoPlaylistOk

`func (o *User) GetAutoPlayNextVideoPlaylistOk() (*bool, bool)`

GetAutoPlayNextVideoPlaylistOk returns a tuple with the AutoPlayNextVideoPlaylist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPlayNextVideoPlaylist

`func (o *User) SetAutoPlayNextVideoPlaylist(v bool)`

SetAutoPlayNextVideoPlaylist sets AutoPlayNextVideoPlaylist field to given value.

### HasAutoPlayNextVideoPlaylist

`func (o *User) HasAutoPlayNextVideoPlaylist() bool`

HasAutoPlayNextVideoPlaylist returns a boolean if a field has been set.

### GetAutoPlayVideo

`func (o *User) GetAutoPlayVideo() bool`

GetAutoPlayVideo returns the AutoPlayVideo field if non-nil, zero value otherwise.

### GetAutoPlayVideoOk

`func (o *User) GetAutoPlayVideoOk() (*bool, bool)`

GetAutoPlayVideoOk returns a tuple with the AutoPlayVideo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPlayVideo

`func (o *User) SetAutoPlayVideo(v bool)`

SetAutoPlayVideo sets AutoPlayVideo field to given value.

### HasAutoPlayVideo

`func (o *User) HasAutoPlayVideo() bool`

HasAutoPlayVideo returns a boolean if a field has been set.

### GetBlocked

`func (o *User) GetBlocked() bool`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *User) GetBlockedOk() (*bool, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *User) SetBlocked(v bool)`

SetBlocked sets Blocked field to given value.

### HasBlocked

`func (o *User) HasBlocked() bool`

HasBlocked returns a boolean if a field has been set.

### GetBlockedReason

`func (o *User) GetBlockedReason() string`

GetBlockedReason returns the BlockedReason field if non-nil, zero value otherwise.

### GetBlockedReasonOk

`func (o *User) GetBlockedReasonOk() (*string, bool)`

GetBlockedReasonOk returns a tuple with the BlockedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedReason

`func (o *User) SetBlockedReason(v string)`

SetBlockedReason sets BlockedReason field to given value.

### HasBlockedReason

`func (o *User) HasBlockedReason() bool`

HasBlockedReason returns a boolean if a field has been set.

### GetCreatedAt

`func (o *User) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *User) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *User) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *User) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEmail

`func (o *User) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *User) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *User) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *User) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEmailVerified

`func (o *User) GetEmailVerified() bool`

GetEmailVerified returns the EmailVerified field if non-nil, zero value otherwise.

### GetEmailVerifiedOk

`func (o *User) GetEmailVerifiedOk() (*bool, bool)`

GetEmailVerifiedOk returns a tuple with the EmailVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailVerified

`func (o *User) SetEmailVerified(v bool)`

SetEmailVerified sets EmailVerified field to given value.

### HasEmailVerified

`func (o *User) HasEmailVerified() bool`

HasEmailVerified returns a boolean if a field has been set.

### GetId

`func (o *User) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *User) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPluginAuth

`func (o *User) GetPluginAuth() string`

GetPluginAuth returns the PluginAuth field if non-nil, zero value otherwise.

### GetPluginAuthOk

`func (o *User) GetPluginAuthOk() (*string, bool)`

GetPluginAuthOk returns a tuple with the PluginAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginAuth

`func (o *User) SetPluginAuth(v string)`

SetPluginAuth sets PluginAuth field to given value.

### HasPluginAuth

`func (o *User) HasPluginAuth() bool`

HasPluginAuth returns a boolean if a field has been set.

### GetLastLoginDate

`func (o *User) GetLastLoginDate() time.Time`

GetLastLoginDate returns the LastLoginDate field if non-nil, zero value otherwise.

### GetLastLoginDateOk

`func (o *User) GetLastLoginDateOk() (*time.Time, bool)`

GetLastLoginDateOk returns a tuple with the LastLoginDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginDate

`func (o *User) SetLastLoginDate(v time.Time)`

SetLastLoginDate sets LastLoginDate field to given value.

### HasLastLoginDate

`func (o *User) HasLastLoginDate() bool`

HasLastLoginDate returns a boolean if a field has been set.

### GetNoInstanceConfigWarningModal

`func (o *User) GetNoInstanceConfigWarningModal() bool`

GetNoInstanceConfigWarningModal returns the NoInstanceConfigWarningModal field if non-nil, zero value otherwise.

### GetNoInstanceConfigWarningModalOk

`func (o *User) GetNoInstanceConfigWarningModalOk() (*bool, bool)`

GetNoInstanceConfigWarningModalOk returns a tuple with the NoInstanceConfigWarningModal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoInstanceConfigWarningModal

`func (o *User) SetNoInstanceConfigWarningModal(v bool)`

SetNoInstanceConfigWarningModal sets NoInstanceConfigWarningModal field to given value.

### HasNoInstanceConfigWarningModal

`func (o *User) HasNoInstanceConfigWarningModal() bool`

HasNoInstanceConfigWarningModal returns a boolean if a field has been set.

### GetNoAccountSetupWarningModal

`func (o *User) GetNoAccountSetupWarningModal() bool`

GetNoAccountSetupWarningModal returns the NoAccountSetupWarningModal field if non-nil, zero value otherwise.

### GetNoAccountSetupWarningModalOk

`func (o *User) GetNoAccountSetupWarningModalOk() (*bool, bool)`

GetNoAccountSetupWarningModalOk returns a tuple with the NoAccountSetupWarningModal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoAccountSetupWarningModal

`func (o *User) SetNoAccountSetupWarningModal(v bool)`

SetNoAccountSetupWarningModal sets NoAccountSetupWarningModal field to given value.

### HasNoAccountSetupWarningModal

`func (o *User) HasNoAccountSetupWarningModal() bool`

HasNoAccountSetupWarningModal returns a boolean if a field has been set.

### GetNoWelcomeModal

`func (o *User) GetNoWelcomeModal() bool`

GetNoWelcomeModal returns the NoWelcomeModal field if non-nil, zero value otherwise.

### GetNoWelcomeModalOk

`func (o *User) GetNoWelcomeModalOk() (*bool, bool)`

GetNoWelcomeModalOk returns a tuple with the NoWelcomeModal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoWelcomeModal

`func (o *User) SetNoWelcomeModal(v bool)`

SetNoWelcomeModal sets NoWelcomeModal field to given value.

### HasNoWelcomeModal

`func (o *User) HasNoWelcomeModal() bool`

HasNoWelcomeModal returns a boolean if a field has been set.

### GetNsfwPolicy

`func (o *User) GetNsfwPolicy() NSFWPolicy`

GetNsfwPolicy returns the NsfwPolicy field if non-nil, zero value otherwise.

### GetNsfwPolicyOk

`func (o *User) GetNsfwPolicyOk() (*NSFWPolicy, bool)`

GetNsfwPolicyOk returns a tuple with the NsfwPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsfwPolicy

`func (o *User) SetNsfwPolicy(v NSFWPolicy)`

SetNsfwPolicy sets NsfwPolicy field to given value.

### HasNsfwPolicy

`func (o *User) HasNsfwPolicy() bool`

HasNsfwPolicy returns a boolean if a field has been set.

### GetRole

`func (o *User) GetRole() UserRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *User) GetRoleOk() (*UserRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *User) SetRole(v UserRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *User) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetTheme

`func (o *User) GetTheme() string`

GetTheme returns the Theme field if non-nil, zero value otherwise.

### GetThemeOk

`func (o *User) GetThemeOk() (*string, bool)`

GetThemeOk returns a tuple with the Theme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheme

`func (o *User) SetTheme(v string)`

SetTheme sets Theme field to given value.

### HasTheme

`func (o *User) HasTheme() bool`

HasTheme returns a boolean if a field has been set.

### GetUsername

`func (o *User) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *User) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *User) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *User) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetVideoChannels

`func (o *User) GetVideoChannels() []VideoChannel`

GetVideoChannels returns the VideoChannels field if non-nil, zero value otherwise.

### GetVideoChannelsOk

`func (o *User) GetVideoChannelsOk() (*[]VideoChannel, bool)`

GetVideoChannelsOk returns a tuple with the VideoChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoChannels

`func (o *User) SetVideoChannels(v []VideoChannel)`

SetVideoChannels sets VideoChannels field to given value.

### HasVideoChannels

`func (o *User) HasVideoChannels() bool`

HasVideoChannels returns a boolean if a field has been set.

### GetVideoQuota

`func (o *User) GetVideoQuota() int32`

GetVideoQuota returns the VideoQuota field if non-nil, zero value otherwise.

### GetVideoQuotaOk

`func (o *User) GetVideoQuotaOk() (*int32, bool)`

GetVideoQuotaOk returns a tuple with the VideoQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoQuota

`func (o *User) SetVideoQuota(v int32)`

SetVideoQuota sets VideoQuota field to given value.

### HasVideoQuota

`func (o *User) HasVideoQuota() bool`

HasVideoQuota returns a boolean if a field has been set.

### GetVideoQuotaDaily

`func (o *User) GetVideoQuotaDaily() int32`

GetVideoQuotaDaily returns the VideoQuotaDaily field if non-nil, zero value otherwise.

### GetVideoQuotaDailyOk

`func (o *User) GetVideoQuotaDailyOk() (*int32, bool)`

GetVideoQuotaDailyOk returns a tuple with the VideoQuotaDaily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoQuotaDaily

`func (o *User) SetVideoQuotaDaily(v int32)`

SetVideoQuotaDaily sets VideoQuotaDaily field to given value.

### HasVideoQuotaDaily

`func (o *User) HasVideoQuotaDaily() bool`

HasVideoQuotaDaily returns a boolean if a field has been set.

### GetP2pEnabled

`func (o *User) GetP2pEnabled() bool`

GetP2pEnabled returns the P2pEnabled field if non-nil, zero value otherwise.

### GetP2pEnabledOk

`func (o *User) GetP2pEnabledOk() (*bool, bool)`

GetP2pEnabledOk returns a tuple with the P2pEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP2pEnabled

`func (o *User) SetP2pEnabled(v bool)`

SetP2pEnabled sets P2pEnabled field to given value.

### HasP2pEnabled

`func (o *User) HasP2pEnabled() bool`

HasP2pEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


