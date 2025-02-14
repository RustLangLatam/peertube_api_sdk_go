# UpdateMe

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **string** |  | [optional] 
**CurrentPassword** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** | new email used for login and service communications | [optional] 
**DisplayName** | Pointer to **string** | new name of the user in its representations | [optional] 
**DisplayNSFW** | Pointer to **string** | new NSFW display policy | [optional] 
**P2pEnabled** | Pointer to **bool** | whether to enable P2P in the player or not | [optional] 
**AutoPlayVideo** | Pointer to **bool** | new preference regarding playing videos automatically | [optional] 
**AutoPlayNextVideo** | Pointer to **bool** | new preference regarding playing following videos automatically | [optional] 
**AutoPlayNextVideoPlaylist** | Pointer to **bool** | new preference regarding playing following playlist videos automatically | [optional] 
**VideosHistoryEnabled** | Pointer to **bool** | whether to keep track of watched history or not | [optional] 
**VideoLanguages** | Pointer to **[]string** | list of languages to filter videos down to | [optional] 
**Theme** | Pointer to **string** |  | [optional] 
**NoInstanceConfigWarningModal** | Pointer to **bool** |  | [optional] 
**NoAccountSetupWarningModal** | Pointer to **bool** |  | [optional] 
**NoWelcomeModal** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateMe

`func NewUpdateMe() *UpdateMe`

NewUpdateMe instantiates a new UpdateMe object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMeWithDefaults

`func NewUpdateMeWithDefaults() *UpdateMe`

NewUpdateMeWithDefaults instantiates a new UpdateMe object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *UpdateMe) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateMe) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateMe) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UpdateMe) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetCurrentPassword

`func (o *UpdateMe) GetCurrentPassword() string`

GetCurrentPassword returns the CurrentPassword field if non-nil, zero value otherwise.

### GetCurrentPasswordOk

`func (o *UpdateMe) GetCurrentPasswordOk() (*string, bool)`

GetCurrentPasswordOk returns a tuple with the CurrentPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPassword

`func (o *UpdateMe) SetCurrentPassword(v string)`

SetCurrentPassword sets CurrentPassword field to given value.

### HasCurrentPassword

`func (o *UpdateMe) HasCurrentPassword() bool`

HasCurrentPassword returns a boolean if a field has been set.

### GetEmail

`func (o *UpdateMe) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateMe) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateMe) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdateMe) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetDisplayName

`func (o *UpdateMe) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UpdateMe) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UpdateMe) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UpdateMe) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDisplayNSFW

`func (o *UpdateMe) GetDisplayNSFW() string`

GetDisplayNSFW returns the DisplayNSFW field if non-nil, zero value otherwise.

### GetDisplayNSFWOk

`func (o *UpdateMe) GetDisplayNSFWOk() (*string, bool)`

GetDisplayNSFWOk returns a tuple with the DisplayNSFW field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNSFW

`func (o *UpdateMe) SetDisplayNSFW(v string)`

SetDisplayNSFW sets DisplayNSFW field to given value.

### HasDisplayNSFW

`func (o *UpdateMe) HasDisplayNSFW() bool`

HasDisplayNSFW returns a boolean if a field has been set.

### GetP2pEnabled

`func (o *UpdateMe) GetP2pEnabled() bool`

GetP2pEnabled returns the P2pEnabled field if non-nil, zero value otherwise.

### GetP2pEnabledOk

`func (o *UpdateMe) GetP2pEnabledOk() (*bool, bool)`

GetP2pEnabledOk returns a tuple with the P2pEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP2pEnabled

`func (o *UpdateMe) SetP2pEnabled(v bool)`

SetP2pEnabled sets P2pEnabled field to given value.

### HasP2pEnabled

`func (o *UpdateMe) HasP2pEnabled() bool`

HasP2pEnabled returns a boolean if a field has been set.

### GetAutoPlayVideo

`func (o *UpdateMe) GetAutoPlayVideo() bool`

GetAutoPlayVideo returns the AutoPlayVideo field if non-nil, zero value otherwise.

### GetAutoPlayVideoOk

`func (o *UpdateMe) GetAutoPlayVideoOk() (*bool, bool)`

GetAutoPlayVideoOk returns a tuple with the AutoPlayVideo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPlayVideo

`func (o *UpdateMe) SetAutoPlayVideo(v bool)`

SetAutoPlayVideo sets AutoPlayVideo field to given value.

### HasAutoPlayVideo

`func (o *UpdateMe) HasAutoPlayVideo() bool`

HasAutoPlayVideo returns a boolean if a field has been set.

### GetAutoPlayNextVideo

`func (o *UpdateMe) GetAutoPlayNextVideo() bool`

GetAutoPlayNextVideo returns the AutoPlayNextVideo field if non-nil, zero value otherwise.

### GetAutoPlayNextVideoOk

`func (o *UpdateMe) GetAutoPlayNextVideoOk() (*bool, bool)`

GetAutoPlayNextVideoOk returns a tuple with the AutoPlayNextVideo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPlayNextVideo

`func (o *UpdateMe) SetAutoPlayNextVideo(v bool)`

SetAutoPlayNextVideo sets AutoPlayNextVideo field to given value.

### HasAutoPlayNextVideo

`func (o *UpdateMe) HasAutoPlayNextVideo() bool`

HasAutoPlayNextVideo returns a boolean if a field has been set.

### GetAutoPlayNextVideoPlaylist

`func (o *UpdateMe) GetAutoPlayNextVideoPlaylist() bool`

GetAutoPlayNextVideoPlaylist returns the AutoPlayNextVideoPlaylist field if non-nil, zero value otherwise.

### GetAutoPlayNextVideoPlaylistOk

`func (o *UpdateMe) GetAutoPlayNextVideoPlaylistOk() (*bool, bool)`

GetAutoPlayNextVideoPlaylistOk returns a tuple with the AutoPlayNextVideoPlaylist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPlayNextVideoPlaylist

`func (o *UpdateMe) SetAutoPlayNextVideoPlaylist(v bool)`

SetAutoPlayNextVideoPlaylist sets AutoPlayNextVideoPlaylist field to given value.

### HasAutoPlayNextVideoPlaylist

`func (o *UpdateMe) HasAutoPlayNextVideoPlaylist() bool`

HasAutoPlayNextVideoPlaylist returns a boolean if a field has been set.

### GetVideosHistoryEnabled

`func (o *UpdateMe) GetVideosHistoryEnabled() bool`

GetVideosHistoryEnabled returns the VideosHistoryEnabled field if non-nil, zero value otherwise.

### GetVideosHistoryEnabledOk

`func (o *UpdateMe) GetVideosHistoryEnabledOk() (*bool, bool)`

GetVideosHistoryEnabledOk returns a tuple with the VideosHistoryEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideosHistoryEnabled

`func (o *UpdateMe) SetVideosHistoryEnabled(v bool)`

SetVideosHistoryEnabled sets VideosHistoryEnabled field to given value.

### HasVideosHistoryEnabled

`func (o *UpdateMe) HasVideosHistoryEnabled() bool`

HasVideosHistoryEnabled returns a boolean if a field has been set.

### GetVideoLanguages

`func (o *UpdateMe) GetVideoLanguages() []string`

GetVideoLanguages returns the VideoLanguages field if non-nil, zero value otherwise.

### GetVideoLanguagesOk

`func (o *UpdateMe) GetVideoLanguagesOk() (*[]string, bool)`

GetVideoLanguagesOk returns a tuple with the VideoLanguages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoLanguages

`func (o *UpdateMe) SetVideoLanguages(v []string)`

SetVideoLanguages sets VideoLanguages field to given value.

### HasVideoLanguages

`func (o *UpdateMe) HasVideoLanguages() bool`

HasVideoLanguages returns a boolean if a field has been set.

### GetTheme

`func (o *UpdateMe) GetTheme() string`

GetTheme returns the Theme field if non-nil, zero value otherwise.

### GetThemeOk

`func (o *UpdateMe) GetThemeOk() (*string, bool)`

GetThemeOk returns a tuple with the Theme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheme

`func (o *UpdateMe) SetTheme(v string)`

SetTheme sets Theme field to given value.

### HasTheme

`func (o *UpdateMe) HasTheme() bool`

HasTheme returns a boolean if a field has been set.

### GetNoInstanceConfigWarningModal

`func (o *UpdateMe) GetNoInstanceConfigWarningModal() bool`

GetNoInstanceConfigWarningModal returns the NoInstanceConfigWarningModal field if non-nil, zero value otherwise.

### GetNoInstanceConfigWarningModalOk

`func (o *UpdateMe) GetNoInstanceConfigWarningModalOk() (*bool, bool)`

GetNoInstanceConfigWarningModalOk returns a tuple with the NoInstanceConfigWarningModal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoInstanceConfigWarningModal

`func (o *UpdateMe) SetNoInstanceConfigWarningModal(v bool)`

SetNoInstanceConfigWarningModal sets NoInstanceConfigWarningModal field to given value.

### HasNoInstanceConfigWarningModal

`func (o *UpdateMe) HasNoInstanceConfigWarningModal() bool`

HasNoInstanceConfigWarningModal returns a boolean if a field has been set.

### GetNoAccountSetupWarningModal

`func (o *UpdateMe) GetNoAccountSetupWarningModal() bool`

GetNoAccountSetupWarningModal returns the NoAccountSetupWarningModal field if non-nil, zero value otherwise.

### GetNoAccountSetupWarningModalOk

`func (o *UpdateMe) GetNoAccountSetupWarningModalOk() (*bool, bool)`

GetNoAccountSetupWarningModalOk returns a tuple with the NoAccountSetupWarningModal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoAccountSetupWarningModal

`func (o *UpdateMe) SetNoAccountSetupWarningModal(v bool)`

SetNoAccountSetupWarningModal sets NoAccountSetupWarningModal field to given value.

### HasNoAccountSetupWarningModal

`func (o *UpdateMe) HasNoAccountSetupWarningModal() bool`

HasNoAccountSetupWarningModal returns a boolean if a field has been set.

### GetNoWelcomeModal

`func (o *UpdateMe) GetNoWelcomeModal() bool`

GetNoWelcomeModal returns the NoWelcomeModal field if non-nil, zero value otherwise.

### GetNoWelcomeModalOk

`func (o *UpdateMe) GetNoWelcomeModalOk() (*bool, bool)`

GetNoWelcomeModalOk returns a tuple with the NoWelcomeModal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoWelcomeModal

`func (o *UpdateMe) SetNoWelcomeModal(v bool)`

SetNoWelcomeModal sets NoWelcomeModal field to given value.

### HasNoWelcomeModal

`func (o *UpdateMe) HasNoWelcomeModal() bool`

HasNoWelcomeModal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


