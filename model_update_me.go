/*
PeerTube

The PeerTube API is built on HTTP(S) and is RESTful. You can use your favorite HTTP/REST library for your programming language to use PeerTube. The spec API is fully compatible with [openapi-generator](https://github.com/OpenAPITools/openapi-generator/wiki/API-client-generator-HOWTO) which generates a client SDK in the language of your choice - we generate some client SDKs automatically:  - [Python](https://framagit.org/framasoft/peertube/clients/python) - [Go](https://framagit.org/framasoft/peertube/clients/go) - [Kotlin](https://framagit.org/framasoft/peertube/clients/kotlin)  See the [REST API quick start](https://docs.joinpeertube.org/api/rest-getting-started) for a few examples of using the PeerTube API.  # Authentication  When you sign up for an account on a PeerTube instance, you are given the possibility to generate sessions on it, and authenticate there using an access token. Only __one access token can currently be used at a time__.  ## Roles  Accounts are given permissions based on their role. There are three roles on PeerTube: Administrator, Moderator, and User. See the [roles guide](https://docs.joinpeertube.org/admin/managing-users#roles) for a detail of their permissions.  # Errors  The API uses standard HTTP status codes to indicate the success or failure of the API call, completed by a [RFC7807-compliant](https://tools.ietf.org/html/rfc7807) response body.  ``` HTTP 1.1 404 Not Found Content-Type: application/problem+json; charset=utf-8  {   \"detail\": \"Video not found\",   \"docs\": \"https://docs.joinpeertube.org/api-rest-reference.html#operation/getVideo\",   \"status\": 404,   \"title\": \"Not Found\",   \"type\": \"about:blank\" } ```  We provide error `type` (following RFC7807) and `code` (internal PeerTube code) values for [a growing number of cases](https://github.com/Chocobozzz/PeerTube/blob/develop/packages/models/src/server/server-error-code.enum.ts), but it is still optional. Types are used to disambiguate errors that bear the same status code and are non-obvious:  ``` HTTP 1.1 403 Forbidden Content-Type: application/problem+json; charset=utf-8  {   \"detail\": \"Cannot get this video regarding follow constraints\",   \"docs\": \"https://docs.joinpeertube.org/api-rest-reference.html#operation/getVideo\",   \"status\": 403,   \"title\": \"Forbidden\",   \"type\": \"https://docs.joinpeertube.org/api-rest-reference.html#section/Errors/does_not_respect_follow_constraints\" } ```  Here a 403 error could otherwise mean that the video is private or blocklisted.  ### Validation errors  Each parameter is evaluated on its own against a set of rules before the route validator proceeds with potential testing involving parameter combinations. Errors coming from validation errors appear earlier and benefit from a more detailed error description:  ``` HTTP 1.1 400 Bad Request Content-Type: application/problem+json; charset=utf-8  {   \"detail\": \"Incorrect request parameters: id\",   \"docs\": \"https://docs.joinpeertube.org/api-rest-reference.html#operation/getVideo\",   \"instance\": \"/api/v1/videos/9c9de5e8-0a1e-484a-b099-e80766180\",   \"invalid-params\": {     \"id\": {       \"location\": \"params\",       \"msg\": \"Invalid value\",       \"param\": \"id\",       \"value\": \"9c9de5e8-0a1e-484a-b099-e80766180\"     }   },   \"status\": 400,   \"title\": \"Bad Request\",   \"type\": \"about:blank\" } ```  Where `id` is the name of the field concerned by the error, within the route definition. `invalid-params.<field>.location` can be either 'params', 'body', 'header', 'query' or 'cookies', and `invalid-params.<field>.value` reports the value that didn't pass validation whose `invalid-params.<field>.msg` is about.  ### Deprecated error fields  Some fields could be included with previous versions. They are still included but their use is deprecated: - `error`: superseded by `detail`  # Rate limits  We are rate-limiting all endpoints of PeerTube's API. Custom values can be set by administrators:  | Endpoint (prefix: `/api/v1`) | Calls         | Time frame   | |------------------------------|---------------|--------------| | `/_*`                         | 50            | 10 seconds   | | `POST /users/token`          | 15            | 5 minutes    | | `POST /users/register`       | 2<sup>*</sup> | 5 minutes    | | `POST /users/ask-send-verify-email` | 3      | 5 minutes    |  Depending on the endpoint, <sup>*</sup>failed requests are not taken into account. A service limit is announced by a `429 Too Many Requests` status code.  You can get details about the current state of your rate limit by reading the following headers:  | Header                  | Description                                                | |-------------------------|------------------------------------------------------------| | `X-RateLimit-Limit`     | Number of max requests allowed in the current time period  | | `X-RateLimit-Remaining` | Number of remaining requests in the current time period    | | `X-RateLimit-Reset`     | Timestamp of end of current time period as UNIX timestamp  | | `Retry-After`           | Seconds to delay after the first `429` is received         |  # CORS  This API features [Cross-Origin Resource Sharing (CORS)](https://fetch.spec.whatwg.org/), allowing cross-domain communication from the browser for some routes:  | Endpoint                    | |------------------------- ---| | `/api/_*`                    | | `/download/_*`               | | `/lazy-static/_*`            | | `/.well-known/webfinger`    |  In addition, all routes serving ActivityPub are CORS-enabled for all origins.

API version: 7.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package peertube_api_sdk

import (
	"encoding/json"
)

// checks if the UpdateMe type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateMe{}

// UpdateMe struct for UpdateMe
type UpdateMe struct {
	Password        *string `json:"password,omitempty"`
	CurrentPassword *string `json:"currentPassword,omitempty"`
	// new email used for login and service communications
	Email *string `json:"email,omitempty"`
	// new name of the user in its representations
	DisplayName *string `json:"displayName,omitempty"`
	// new NSFW display policy
	DisplayNSFW *string `json:"displayNSFW,omitempty"`
	// whether to enable P2P in the player or not
	P2pEnabled *bool `json:"p2pEnabled,omitempty"`
	// new preference regarding playing videos automatically
	AutoPlayVideo *bool `json:"autoPlayVideo,omitempty"`
	// new preference regarding playing following videos automatically
	AutoPlayNextVideo *bool `json:"autoPlayNextVideo,omitempty"`
	// new preference regarding playing following playlist videos automatically
	AutoPlayNextVideoPlaylist *bool `json:"autoPlayNextVideoPlaylist,omitempty"`
	// whether to keep track of watched history or not
	VideosHistoryEnabled *bool `json:"videosHistoryEnabled,omitempty"`
	// list of languages to filter videos down to
	VideoLanguages               []string `json:"videoLanguages,omitempty"`
	Theme                        *string  `json:"theme,omitempty"`
	NoInstanceConfigWarningModal *bool    `json:"noInstanceConfigWarningModal,omitempty"`
	NoAccountSetupWarningModal   *bool    `json:"noAccountSetupWarningModal,omitempty"`
	NoWelcomeModal               *bool    `json:"noWelcomeModal,omitempty"`
}

// NewUpdateMe instantiates a new UpdateMe object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateMe() *UpdateMe {
	this := UpdateMe{}
	return &this
}

// NewUpdateMeWithDefaults instantiates a new UpdateMe object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateMeWithDefaults() *UpdateMe {
	this := UpdateMe{}
	return &this
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *UpdateMe) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMe) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *UpdateMe) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *UpdateMe) SetPassword(v string) {
	o.Password = &v
}

// GetCurrentPassword returns the CurrentPassword field value if set, zero value otherwise.
func (o *UpdateMe) GetCurrentPassword() string {
	if o == nil || IsNil(o.CurrentPassword) {
		var ret string
		return ret
	}
	return *o.CurrentPassword
}

// GetCurrentPasswordOk returns a tuple with the CurrentPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMe) GetCurrentPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.CurrentPassword) {
		return nil, false
	}
	return o.CurrentPassword, true
}

// HasCurrentPassword returns a boolean if a field has been set.
func (o *UpdateMe) HasCurrentPassword() bool {
	if o != nil && !IsNil(o.CurrentPassword) {
		return true
	}

	return false
}

// SetCurrentPassword gets a reference to the given string and assigns it to the CurrentPassword field.
func (o *UpdateMe) SetCurrentPassword(v string) {
	o.CurrentPassword = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UpdateMe) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMe) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UpdateMe) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UpdateMe) SetEmail(v string) {
	o.Email = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *UpdateMe) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMe) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *UpdateMe) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *UpdateMe) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDisplayNSFW returns the DisplayNSFW field value if set, zero value otherwise.
func (o *UpdateMe) GetDisplayNSFW() string {
	if o == nil || IsNil(o.DisplayNSFW) {
		var ret string
		return ret
	}
	return *o.DisplayNSFW
}

// GetDisplayNSFWOk returns a tuple with the DisplayNSFW field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMe) GetDisplayNSFWOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayNSFW) {
		return nil, false
	}
	return o.DisplayNSFW, true
}

// HasDisplayNSFW returns a boolean if a field has been set.
func (o *UpdateMe) HasDisplayNSFW() bool {
	if o != nil && !IsNil(o.DisplayNSFW) {
		return true
	}

	return false
}

// SetDisplayNSFW gets a reference to the given string and assigns it to the DisplayNSFW field.
func (o *UpdateMe) SetDisplayNSFW(v string) {
	o.DisplayNSFW = &v
}

// GetP2pEnabled returns the P2pEnabled field value if set, zero value otherwise.
func (o *UpdateMe) GetP2pEnabled() bool {
	if o == nil || IsNil(o.P2pEnabled) {
		var ret bool
		return ret
	}
	return *o.P2pEnabled
}

// GetP2pEnabledOk returns a tuple with the P2pEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMe) GetP2pEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.P2pEnabled) {
		return nil, false
	}
	return o.P2pEnabled, true
}

// HasP2pEnabled returns a boolean if a field has been set.
func (o *UpdateMe) HasP2pEnabled() bool {
	if o != nil && !IsNil(o.P2pEnabled) {
		return true
	}

	return false
}

// SetP2pEnabled gets a reference to the given bool and assigns it to the P2pEnabled field.
func (o *UpdateMe) SetP2pEnabled(v bool) {
	o.P2pEnabled = &v
}

// GetAutoPlayVideo returns the AutoPlayVideo field value if set, zero value otherwise.
func (o *UpdateMe) GetAutoPlayVideo() bool {
	if o == nil || IsNil(o.AutoPlayVideo) {
		var ret bool
		return ret
	}
	return *o.AutoPlayVideo
}

// GetAutoPlayVideoOk returns a tuple with the AutoPlayVideo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMe) GetAutoPlayVideoOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoPlayVideo) {
		return nil, false
	}
	return o.AutoPlayVideo, true
}

// HasAutoPlayVideo returns a boolean if a field has been set.
func (o *UpdateMe) HasAutoPlayVideo() bool {
	if o != nil && !IsNil(o.AutoPlayVideo) {
		return true
	}

	return false
}

// SetAutoPlayVideo gets a reference to the given bool and assigns it to the AutoPlayVideo field.
func (o *UpdateMe) SetAutoPlayVideo(v bool) {
	o.AutoPlayVideo = &v
}

// GetAutoPlayNextVideo returns the AutoPlayNextVideo field value if set, zero value otherwise.
func (o *UpdateMe) GetAutoPlayNextVideo() bool {
	if o == nil || IsNil(o.AutoPlayNextVideo) {
		var ret bool
		return ret
	}
	return *o.AutoPlayNextVideo
}

// GetAutoPlayNextVideoOk returns a tuple with the AutoPlayNextVideo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMe) GetAutoPlayNextVideoOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoPlayNextVideo) {
		return nil, false
	}
	return o.AutoPlayNextVideo, true
}

// HasAutoPlayNextVideo returns a boolean if a field has been set.
func (o *UpdateMe) HasAutoPlayNextVideo() bool {
	if o != nil && !IsNil(o.AutoPlayNextVideo) {
		return true
	}

	return false
}

// SetAutoPlayNextVideo gets a reference to the given bool and assigns it to the AutoPlayNextVideo field.
func (o *UpdateMe) SetAutoPlayNextVideo(v bool) {
	o.AutoPlayNextVideo = &v
}

// GetAutoPlayNextVideoPlaylist returns the AutoPlayNextVideoPlaylist field value if set, zero value otherwise.
func (o *UpdateMe) GetAutoPlayNextVideoPlaylist() bool {
	if o == nil || IsNil(o.AutoPlayNextVideoPlaylist) {
		var ret bool
		return ret
	}
	return *o.AutoPlayNextVideoPlaylist
}

// GetAutoPlayNextVideoPlaylistOk returns a tuple with the AutoPlayNextVideoPlaylist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMe) GetAutoPlayNextVideoPlaylistOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoPlayNextVideoPlaylist) {
		return nil, false
	}
	return o.AutoPlayNextVideoPlaylist, true
}

// HasAutoPlayNextVideoPlaylist returns a boolean if a field has been set.
func (o *UpdateMe) HasAutoPlayNextVideoPlaylist() bool {
	if o != nil && !IsNil(o.AutoPlayNextVideoPlaylist) {
		return true
	}

	return false
}

// SetAutoPlayNextVideoPlaylist gets a reference to the given bool and assigns it to the AutoPlayNextVideoPlaylist field.
func (o *UpdateMe) SetAutoPlayNextVideoPlaylist(v bool) {
	o.AutoPlayNextVideoPlaylist = &v
}

// GetVideosHistoryEnabled returns the VideosHistoryEnabled field value if set, zero value otherwise.
func (o *UpdateMe) GetVideosHistoryEnabled() bool {
	if o == nil || IsNil(o.VideosHistoryEnabled) {
		var ret bool
		return ret
	}
	return *o.VideosHistoryEnabled
}

// GetVideosHistoryEnabledOk returns a tuple with the VideosHistoryEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMe) GetVideosHistoryEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.VideosHistoryEnabled) {
		return nil, false
	}
	return o.VideosHistoryEnabled, true
}

// HasVideosHistoryEnabled returns a boolean if a field has been set.
func (o *UpdateMe) HasVideosHistoryEnabled() bool {
	if o != nil && !IsNil(o.VideosHistoryEnabled) {
		return true
	}

	return false
}

// SetVideosHistoryEnabled gets a reference to the given bool and assigns it to the VideosHistoryEnabled field.
func (o *UpdateMe) SetVideosHistoryEnabled(v bool) {
	o.VideosHistoryEnabled = &v
}

// GetVideoLanguages returns the VideoLanguages field value if set, zero value otherwise.
func (o *UpdateMe) GetVideoLanguages() []string {
	if o == nil || IsNil(o.VideoLanguages) {
		var ret []string
		return ret
	}
	return o.VideoLanguages
}

// GetVideoLanguagesOk returns a tuple with the VideoLanguages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMe) GetVideoLanguagesOk() ([]string, bool) {
	if o == nil || IsNil(o.VideoLanguages) {
		return nil, false
	}
	return o.VideoLanguages, true
}

// HasVideoLanguages returns a boolean if a field has been set.
func (o *UpdateMe) HasVideoLanguages() bool {
	if o != nil && !IsNil(o.VideoLanguages) {
		return true
	}

	return false
}

// SetVideoLanguages gets a reference to the given []string and assigns it to the VideoLanguages field.
func (o *UpdateMe) SetVideoLanguages(v []string) {
	o.VideoLanguages = v
}

// GetTheme returns the Theme field value if set, zero value otherwise.
func (o *UpdateMe) GetTheme() string {
	if o == nil || IsNil(o.Theme) {
		var ret string
		return ret
	}
	return *o.Theme
}

// GetThemeOk returns a tuple with the Theme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMe) GetThemeOk() (*string, bool) {
	if o == nil || IsNil(o.Theme) {
		return nil, false
	}
	return o.Theme, true
}

// HasTheme returns a boolean if a field has been set.
func (o *UpdateMe) HasTheme() bool {
	if o != nil && !IsNil(o.Theme) {
		return true
	}

	return false
}

// SetTheme gets a reference to the given string and assigns it to the Theme field.
func (o *UpdateMe) SetTheme(v string) {
	o.Theme = &v
}

// GetNoInstanceConfigWarningModal returns the NoInstanceConfigWarningModal field value if set, zero value otherwise.
func (o *UpdateMe) GetNoInstanceConfigWarningModal() bool {
	if o == nil || IsNil(o.NoInstanceConfigWarningModal) {
		var ret bool
		return ret
	}
	return *o.NoInstanceConfigWarningModal
}

// GetNoInstanceConfigWarningModalOk returns a tuple with the NoInstanceConfigWarningModal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMe) GetNoInstanceConfigWarningModalOk() (*bool, bool) {
	if o == nil || IsNil(o.NoInstanceConfigWarningModal) {
		return nil, false
	}
	return o.NoInstanceConfigWarningModal, true
}

// HasNoInstanceConfigWarningModal returns a boolean if a field has been set.
func (o *UpdateMe) HasNoInstanceConfigWarningModal() bool {
	if o != nil && !IsNil(o.NoInstanceConfigWarningModal) {
		return true
	}

	return false
}

// SetNoInstanceConfigWarningModal gets a reference to the given bool and assigns it to the NoInstanceConfigWarningModal field.
func (o *UpdateMe) SetNoInstanceConfigWarningModal(v bool) {
	o.NoInstanceConfigWarningModal = &v
}

// GetNoAccountSetupWarningModal returns the NoAccountSetupWarningModal field value if set, zero value otherwise.
func (o *UpdateMe) GetNoAccountSetupWarningModal() bool {
	if o == nil || IsNil(o.NoAccountSetupWarningModal) {
		var ret bool
		return ret
	}
	return *o.NoAccountSetupWarningModal
}

// GetNoAccountSetupWarningModalOk returns a tuple with the NoAccountSetupWarningModal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMe) GetNoAccountSetupWarningModalOk() (*bool, bool) {
	if o == nil || IsNil(o.NoAccountSetupWarningModal) {
		return nil, false
	}
	return o.NoAccountSetupWarningModal, true
}

// HasNoAccountSetupWarningModal returns a boolean if a field has been set.
func (o *UpdateMe) HasNoAccountSetupWarningModal() bool {
	if o != nil && !IsNil(o.NoAccountSetupWarningModal) {
		return true
	}

	return false
}

// SetNoAccountSetupWarningModal gets a reference to the given bool and assigns it to the NoAccountSetupWarningModal field.
func (o *UpdateMe) SetNoAccountSetupWarningModal(v bool) {
	o.NoAccountSetupWarningModal = &v
}

// GetNoWelcomeModal returns the NoWelcomeModal field value if set, zero value otherwise.
func (o *UpdateMe) GetNoWelcomeModal() bool {
	if o == nil || IsNil(o.NoWelcomeModal) {
		var ret bool
		return ret
	}
	return *o.NoWelcomeModal
}

// GetNoWelcomeModalOk returns a tuple with the NoWelcomeModal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMe) GetNoWelcomeModalOk() (*bool, bool) {
	if o == nil || IsNil(o.NoWelcomeModal) {
		return nil, false
	}
	return o.NoWelcomeModal, true
}

// HasNoWelcomeModal returns a boolean if a field has been set.
func (o *UpdateMe) HasNoWelcomeModal() bool {
	if o != nil && !IsNil(o.NoWelcomeModal) {
		return true
	}

	return false
}

// SetNoWelcomeModal gets a reference to the given bool and assigns it to the NoWelcomeModal field.
func (o *UpdateMe) SetNoWelcomeModal(v bool) {
	o.NoWelcomeModal = &v
}

func (o UpdateMe) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateMe) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.CurrentPassword) {
		toSerialize["currentPassword"] = o.CurrentPassword
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.DisplayNSFW) {
		toSerialize["displayNSFW"] = o.DisplayNSFW
	}
	if !IsNil(o.P2pEnabled) {
		toSerialize["p2pEnabled"] = o.P2pEnabled
	}
	if !IsNil(o.AutoPlayVideo) {
		toSerialize["autoPlayVideo"] = o.AutoPlayVideo
	}
	if !IsNil(o.AutoPlayNextVideo) {
		toSerialize["autoPlayNextVideo"] = o.AutoPlayNextVideo
	}
	if !IsNil(o.AutoPlayNextVideoPlaylist) {
		toSerialize["autoPlayNextVideoPlaylist"] = o.AutoPlayNextVideoPlaylist
	}
	if !IsNil(o.VideosHistoryEnabled) {
		toSerialize["videosHistoryEnabled"] = o.VideosHistoryEnabled
	}
	if !IsNil(o.VideoLanguages) {
		toSerialize["videoLanguages"] = o.VideoLanguages
	}
	if !IsNil(o.Theme) {
		toSerialize["theme"] = o.Theme
	}
	if !IsNil(o.NoInstanceConfigWarningModal) {
		toSerialize["noInstanceConfigWarningModal"] = o.NoInstanceConfigWarningModal
	}
	if !IsNil(o.NoAccountSetupWarningModal) {
		toSerialize["noAccountSetupWarningModal"] = o.NoAccountSetupWarningModal
	}
	if !IsNil(o.NoWelcomeModal) {
		toSerialize["noWelcomeModal"] = o.NoWelcomeModal
	}
	return toSerialize, nil
}

type NullableUpdateMe struct {
	value *UpdateMe
	isSet bool
}

func (v NullableUpdateMe) Get() *UpdateMe {
	return v.value
}

func (v *NullableUpdateMe) Set(val *UpdateMe) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateMe) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateMe) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateMe(val *UpdateMe) *NullableUpdateMe {
	return &NullableUpdateMe{value: val, isSet: true}
}

func (v NullableUpdateMe) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateMe) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
