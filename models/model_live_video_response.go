/*
PeerTube

The PeerTube API is built on HTTP(S) and is RESTful. You can use your favorite HTTP/REST library for your programming language to use PeerTube. The spec API is fully compatible with [openapi-generator](https://github.com/OpenAPITools/openapi-generator/wiki/API-client-generator-HOWTO) which generates a client SDK in the language of your choice - we generate some client SDKs automatically:  - [Python](https://framagit.org/framasoft/peertube/clients/python) - [Go](https://framagit.org/framasoft/peertube/clients/go) - [Kotlin](https://framagit.org/framasoft/peertube/clients/kotlin)  See the [REST API quick start](https://docs.joinpeertube.org/api/rest-getting-started) for a few examples of using the PeerTube API.  # Authentication  When you sign up for an account on a PeerTube instance, you are given the possibility to generate sessions on it, and authenticate there using an access token. Only __one access token can currently be used at a time__.  ## Roles  Accounts are given permissions based on their role. There are three roles on PeerTube: Administrator, Moderator, and User. See the [roles guide](https://docs.joinpeertube.org/admin/managing-users#roles) for a detail of their permissions.  # Errors  The API uses standard HTTP status codes to indicate the success or failure of the API call, completed by a [RFC7807-compliant](https://tools.ietf.org/html/rfc7807) response body.  ``` HTTP 1.1 404 Not Found Content-Type: application/problem+json; charset=utf-8  {   \"detail\": \"Video not found\",   \"docs\": \"https://docs.joinpeertube.org/api-rest-reference.html#operation/getVideo\",   \"status\": 404,   \"title\": \"Not Found\",   \"type\": \"about:blank\" } ```  We provide error `type` (following RFC7807) and `code` (internal PeerTube code) values for [a growing number of cases](https://github.com/Chocobozzz/PeerTube/blob/develop/packages/models/src/server/server-error-code.enum.ts), but it is still optional. Types are used to disambiguate errors that bear the same status code and are non-obvious:  ``` HTTP 1.1 403 Forbidden Content-Type: application/problem+json; charset=utf-8  {   \"detail\": \"Cannot get this video regarding follow constraints\",   \"docs\": \"https://docs.joinpeertube.org/api-rest-reference.html#operation/getVideo\",   \"status\": 403,   \"title\": \"Forbidden\",   \"type\": \"https://docs.joinpeertube.org/api-rest-reference.html#section/Errors/does_not_respect_follow_constraints\" } ```  Here a 403 error could otherwise mean that the video is private or blocklisted.  ### Validation errors  Each parameter is evaluated on its own against a set of rules before the route validator proceeds with potential testing involving parameter combinations. Errors coming from validation errors appear earlier and benefit from a more detailed error description:  ``` HTTP 1.1 400 Bad Request Content-Type: application/problem+json; charset=utf-8  {   \"detail\": \"Incorrect request parameters: id\",   \"docs\": \"https://docs.joinpeertube.org/api-rest-reference.html#operation/getVideo\",   \"instance\": \"/api/v1/videos/9c9de5e8-0a1e-484a-b099-e80766180\",   \"invalid-params\": {     \"id\": {       \"location\": \"params\",       \"msg\": \"Invalid value\",       \"param\": \"id\",       \"value\": \"9c9de5e8-0a1e-484a-b099-e80766180\"     }   },   \"status\": 400,   \"title\": \"Bad Request\",   \"type\": \"about:blank\" } ```  Where `id` is the name of the field concerned by the error, within the route definition. `invalid-params.<field>.location` can be either 'params', 'body', 'header', 'query' or 'cookies', and `invalid-params.<field>.value` reports the value that didn't pass validation whose `invalid-params.<field>.msg` is about.  ### Deprecated error fields  Some fields could be included with previous versions. They are still included but their use is deprecated: - `error`: superseded by `detail`  # Rate limits  We are rate-limiting all endpoints of PeerTube's API. Custom values can be set by administrators:  | Endpoint (prefix: `/api/v1`) | Calls         | Time frame   | |------------------------------|---------------|--------------| | `/_*`                         | 50            | 10 seconds   | | `POST /users/token`          | 15            | 5 minutes    | | `POST /users/register`       | 2<sup>*</sup> | 5 minutes    | | `POST /users/ask-send-verify-email` | 3      | 5 minutes    |  Depending on the endpoint, <sup>*</sup>failed requests are not taken into account. A service limit is announced by a `429 Too Many Requests` status code.  You can get details about the current state of your rate limit by reading the following headers:  | Header                  | Description                                                | |-------------------------|------------------------------------------------------------| | `X-RateLimit-Limit`     | Number of max requests allowed in the current time period  | | `X-RateLimit-Remaining` | Number of remaining requests in the current time period    | | `X-RateLimit-Reset`     | Timestamp of end of current time period as UNIX timestamp  | | `Retry-After`           | Seconds to delay after the first `429` is received         |  # CORS  This API features [Cross-Origin Resource Sharing (CORS)](https://fetch.spec.whatwg.org/), allowing cross-domain communication from the browser for some routes:  | Endpoint                    | |------------------------- ---| | `/api/_*`                    | | `/download/_*`               | | `/lazy-static/_*`            | | `/.well-known/webfinger`    |  In addition, all routes serving ActivityPub are CORS-enabled for all origins.

API version: 7.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"github.com/RustLangLatam/peertube_api_sdk_go/utils"
)

// checks if the LiveVideoResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &LiveVideoResponse{}

// LiveVideoResponse struct for LiveVideoResponse
type LiveVideoResponse struct {
	// Included in the response if an appropriate token is provided
	RtmpUrl *string `json:"rtmpUrl,omitempty"`
	// Included in the response if an appropriate token is provided
	RtmpsUrl *string `json:"rtmpsUrl,omitempty"`
	// RTMP stream key to use to stream into this live video. Included in the response if an appropriate token is provided
	StreamKey      *string                  `json:"streamKey,omitempty"`
	SaveReplay     *bool                    `json:"saveReplay,omitempty"`
	ReplaySettings *LiveVideoReplaySettings `json:"replaySettings,omitempty"`
	// User can stream multiple times in a permanent live
	PermanentLive *bool                 `json:"permanentLive,omitempty"`
	LatencyMode   *LiveVideoLatencyMode `json:"latencyMode,omitempty"`
}

// NewLiveVideoResponse instantiates a new LiveVideoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLiveVideoResponse() *LiveVideoResponse {
	this := LiveVideoResponse{}
	return &this
}

// NewLiveVideoResponseWithDefaults instantiates a new LiveVideoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLiveVideoResponseWithDefaults() *LiveVideoResponse {
	this := LiveVideoResponse{}
	return &this
}

// GetRtmpUrl returns the RtmpUrl field value if set, zero value otherwise.
func (o *LiveVideoResponse) GetRtmpUrl() string {
	if o == nil || utils.IsNil(o.RtmpUrl) {
		var ret string
		return ret
	}
	return *o.RtmpUrl
}

// GetRtmpUrlOk returns a tuple with the RtmpUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveVideoResponse) GetRtmpUrlOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RtmpUrl) {
		return nil, false
	}
	return o.RtmpUrl, true
}

// HasRtmpUrl returns a boolean if a field has been set.
func (o *LiveVideoResponse) HasRtmpUrl() bool {
	if o != nil && !utils.IsNil(o.RtmpUrl) {
		return true
	}

	return false
}

// SetRtmpUrl gets a reference to the given string and assigns it to the RtmpUrl field.
func (o *LiveVideoResponse) SetRtmpUrl(v string) {
	o.RtmpUrl = &v
}

// GetRtmpsUrl returns the RtmpsUrl field value if set, zero value otherwise.
func (o *LiveVideoResponse) GetRtmpsUrl() string {
	if o == nil || utils.IsNil(o.RtmpsUrl) {
		var ret string
		return ret
	}
	return *o.RtmpsUrl
}

// GetRtmpsUrlOk returns a tuple with the RtmpsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveVideoResponse) GetRtmpsUrlOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RtmpsUrl) {
		return nil, false
	}
	return o.RtmpsUrl, true
}

// HasRtmpsUrl returns a boolean if a field has been set.
func (o *LiveVideoResponse) HasRtmpsUrl() bool {
	if o != nil && !utils.IsNil(o.RtmpsUrl) {
		return true
	}

	return false
}

// SetRtmpsUrl gets a reference to the given string and assigns it to the RtmpsUrl field.
func (o *LiveVideoResponse) SetRtmpsUrl(v string) {
	o.RtmpsUrl = &v
}

// GetStreamKey returns the StreamKey field value if set, zero value otherwise.
func (o *LiveVideoResponse) GetStreamKey() string {
	if o == nil || utils.IsNil(o.StreamKey) {
		var ret string
		return ret
	}
	return *o.StreamKey
}

// GetStreamKeyOk returns a tuple with the StreamKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveVideoResponse) GetStreamKeyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.StreamKey) {
		return nil, false
	}
	return o.StreamKey, true
}

// HasStreamKey returns a boolean if a field has been set.
func (o *LiveVideoResponse) HasStreamKey() bool {
	if o != nil && !utils.IsNil(o.StreamKey) {
		return true
	}

	return false
}

// SetStreamKey gets a reference to the given string and assigns it to the StreamKey field.
func (o *LiveVideoResponse) SetStreamKey(v string) {
	o.StreamKey = &v
}

// GetSaveReplay returns the SaveReplay field value if set, zero value otherwise.
func (o *LiveVideoResponse) GetSaveReplay() bool {
	if o == nil || utils.IsNil(o.SaveReplay) {
		var ret bool
		return ret
	}
	return *o.SaveReplay
}

// GetSaveReplayOk returns a tuple with the SaveReplay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveVideoResponse) GetSaveReplayOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.SaveReplay) {
		return nil, false
	}
	return o.SaveReplay, true
}

// HasSaveReplay returns a boolean if a field has been set.
func (o *LiveVideoResponse) HasSaveReplay() bool {
	if o != nil && !utils.IsNil(o.SaveReplay) {
		return true
	}

	return false
}

// SetSaveReplay gets a reference to the given bool and assigns it to the SaveReplay field.
func (o *LiveVideoResponse) SetSaveReplay(v bool) {
	o.SaveReplay = &v
}

// GetReplaySettings returns the ReplaySettings field value if set, zero value otherwise.
func (o *LiveVideoResponse) GetReplaySettings() LiveVideoReplaySettings {
	if o == nil || utils.IsNil(o.ReplaySettings) {
		var ret LiveVideoReplaySettings
		return ret
	}
	return *o.ReplaySettings
}

// GetReplaySettingsOk returns a tuple with the ReplaySettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveVideoResponse) GetReplaySettingsOk() (*LiveVideoReplaySettings, bool) {
	if o == nil || utils.IsNil(o.ReplaySettings) {
		return nil, false
	}
	return o.ReplaySettings, true
}

// HasReplaySettings returns a boolean if a field has been set.
func (o *LiveVideoResponse) HasReplaySettings() bool {
	if o != nil && !utils.IsNil(o.ReplaySettings) {
		return true
	}

	return false
}

// SetReplaySettings gets a reference to the given LiveVideoReplaySettings and assigns it to the ReplaySettings field.
func (o *LiveVideoResponse) SetReplaySettings(v LiveVideoReplaySettings) {
	o.ReplaySettings = &v
}

// GetPermanentLive returns the PermanentLive field value if set, zero value otherwise.
func (o *LiveVideoResponse) GetPermanentLive() bool {
	if o == nil || utils.IsNil(o.PermanentLive) {
		var ret bool
		return ret
	}
	return *o.PermanentLive
}

// GetPermanentLiveOk returns a tuple with the PermanentLive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveVideoResponse) GetPermanentLiveOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.PermanentLive) {
		return nil, false
	}
	return o.PermanentLive, true
}

// HasPermanentLive returns a boolean if a field has been set.
func (o *LiveVideoResponse) HasPermanentLive() bool {
	if o != nil && !utils.IsNil(o.PermanentLive) {
		return true
	}

	return false
}

// SetPermanentLive gets a reference to the given bool and assigns it to the PermanentLive field.
func (o *LiveVideoResponse) SetPermanentLive(v bool) {
	o.PermanentLive = &v
}

// GetLatencyMode returns the LatencyMode field value if set, zero value otherwise.
func (o *LiveVideoResponse) GetLatencyMode() LiveVideoLatencyMode {
	if o == nil || utils.IsNil(o.LatencyMode) {
		var ret LiveVideoLatencyMode
		return ret
	}
	return *o.LatencyMode
}

// GetLatencyModeOk returns a tuple with the LatencyMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveVideoResponse) GetLatencyModeOk() (*LiveVideoLatencyMode, bool) {
	if o == nil || utils.IsNil(o.LatencyMode) {
		return nil, false
	}
	return o.LatencyMode, true
}

// HasLatencyMode returns a boolean if a field has been set.
func (o *LiveVideoResponse) HasLatencyMode() bool {
	if o != nil && !utils.IsNil(o.LatencyMode) {
		return true
	}

	return false
}

// SetLatencyMode gets a reference to the given LiveVideoLatencyMode and assigns it to the LatencyMode field.
func (o *LiveVideoResponse) SetLatencyMode(v LiveVideoLatencyMode) {
	o.LatencyMode = &v
}

func (o LiveVideoResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LiveVideoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.RtmpUrl) {
		toSerialize["rtmpUrl"] = o.RtmpUrl
	}
	if !utils.IsNil(o.RtmpsUrl) {
		toSerialize["rtmpsUrl"] = o.RtmpsUrl
	}
	if !utils.IsNil(o.StreamKey) {
		toSerialize["streamKey"] = o.StreamKey
	}
	if !utils.IsNil(o.SaveReplay) {
		toSerialize["saveReplay"] = o.SaveReplay
	}
	if !utils.IsNil(o.ReplaySettings) {
		toSerialize["replaySettings"] = o.ReplaySettings
	}
	if !utils.IsNil(o.PermanentLive) {
		toSerialize["permanentLive"] = o.PermanentLive
	}
	if !utils.IsNil(o.LatencyMode) {
		toSerialize["latencyMode"] = o.LatencyMode
	}
	return toSerialize, nil
}

type NullableLiveVideoResponse struct {
	value *LiveVideoResponse
	isSet bool
}

func (v NullableLiveVideoResponse) Get() *LiveVideoResponse {
	return v.value
}

func (v *NullableLiveVideoResponse) Set(val *LiveVideoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLiveVideoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLiveVideoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLiveVideoResponse(val *LiveVideoResponse) *NullableLiveVideoResponse {
	return &NullableLiveVideoResponse{value: val, isSet: true}
}

func (v NullableLiveVideoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLiveVideoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
