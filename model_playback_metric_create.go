/*
PeerTube

The PeerTube API is built on HTTP(S) and is RESTful. You can use your favorite HTTP/REST library for your programming language to use PeerTube. The spec API is fully compatible with [openapi-generator](https://github.com/OpenAPITools/openapi-generator/wiki/API-client-generator-HOWTO) which generates a client SDK in the language of your choice - we generate some client SDKs automatically:  - [Python](https://framagit.org/framasoft/peertube/clients/python) - [Go](https://framagit.org/framasoft/peertube/clients/go) - [Kotlin](https://framagit.org/framasoft/peertube/clients/kotlin)  See the [REST API quick start](https://docs.joinpeertube.org/api/rest-getting-started) for a few examples of using the PeerTube API.  # Authentication  When you sign up for an account on a PeerTube instance, you are given the possibility to generate sessions on it, and authenticate there using an access token. Only __one access token can currently be used at a time__.  ## Roles  Accounts are given permissions based on their role. There are three roles on PeerTube: Administrator, Moderator, and User. See the [roles guide](https://docs.joinpeertube.org/admin/managing-users#roles) for a detail of their permissions.  # Errors  The API uses standard HTTP status codes to indicate the success or failure of the API call, completed by a [RFC7807-compliant](https://tools.ietf.org/html/rfc7807) response body.  ``` HTTP 1.1 404 Not Found Content-Type: application/problem+json; charset=utf-8  {   \"detail\": \"Video not found\",   \"docs\": \"https://docs.joinpeertube.org/api-rest-reference.html#operation/getVideo\",   \"status\": 404,   \"title\": \"Not Found\",   \"type\": \"about:blank\" } ```  We provide error `type` (following RFC7807) and `code` (internal PeerTube code) values for [a growing number of cases](https://github.com/Chocobozzz/PeerTube/blob/develop/packages/models/src/server/server-error-code.enum.ts), but it is still optional. Types are used to disambiguate errors that bear the same status code and are non-obvious:  ``` HTTP 1.1 403 Forbidden Content-Type: application/problem+json; charset=utf-8  {   \"detail\": \"Cannot get this video regarding follow constraints\",   \"docs\": \"https://docs.joinpeertube.org/api-rest-reference.html#operation/getVideo\",   \"status\": 403,   \"title\": \"Forbidden\",   \"type\": \"https://docs.joinpeertube.org/api-rest-reference.html#section/Errors/does_not_respect_follow_constraints\" } ```  Here a 403 error could otherwise mean that the video is private or blocklisted.  ### Validation errors  Each parameter is evaluated on its own against a set of rules before the route validator proceeds with potential testing involving parameter combinations. Errors coming from validation errors appear earlier and benefit from a more detailed error description:  ``` HTTP 1.1 400 Bad Request Content-Type: application/problem+json; charset=utf-8  {   \"detail\": \"Incorrect request parameters: id\",   \"docs\": \"https://docs.joinpeertube.org/api-rest-reference.html#operation/getVideo\",   \"instance\": \"/api/v1/videos/9c9de5e8-0a1e-484a-b099-e80766180\",   \"invalid-params\": {     \"id\": {       \"location\": \"params\",       \"msg\": \"Invalid value\",       \"param\": \"id\",       \"value\": \"9c9de5e8-0a1e-484a-b099-e80766180\"     }   },   \"status\": 400,   \"title\": \"Bad Request\",   \"type\": \"about:blank\" } ```  Where `id` is the name of the field concerned by the error, within the route definition. `invalid-params.<field>.location` can be either 'params', 'body', 'header', 'query' or 'cookies', and `invalid-params.<field>.value` reports the value that didn't pass validation whose `invalid-params.<field>.msg` is about.  ### Deprecated error fields  Some fields could be included with previous versions. They are still included but their use is deprecated: - `error`: superseded by `detail`  # Rate limits  We are rate-limiting all endpoints of PeerTube's API. Custom values can be set by administrators:  | Endpoint (prefix: `/api/v1`) | Calls         | Time frame   | |------------------------------|---------------|--------------| | `/_*`                         | 50            | 10 seconds   | | `POST /users/token`          | 15            | 5 minutes    | | `POST /users/register`       | 2<sup>*</sup> | 5 minutes    | | `POST /users/ask-send-verify-email` | 3      | 5 minutes    |  Depending on the endpoint, <sup>*</sup>failed requests are not taken into account. A service limit is announced by a `429 Too Many Requests` status code.  You can get details about the current state of your rate limit by reading the following headers:  | Header                  | Description                                                | |-------------------------|------------------------------------------------------------| | `X-RateLimit-Limit`     | Number of max requests allowed in the current time period  | | `X-RateLimit-Remaining` | Number of remaining requests in the current time period    | | `X-RateLimit-Reset`     | Timestamp of end of current time period as UNIX timestamp  | | `Retry-After`           | Seconds to delay after the first `429` is received         |  # CORS  This API features [Cross-Origin Resource Sharing (CORS)](https://fetch.spec.whatwg.org/), allowing cross-domain communication from the browser for some routes:  | Endpoint                    | |------------------------- ---| | `/api/_*`                    | | `/download/_*`               | | `/lazy-static/_*`            | | `/.well-known/webfinger`    |  In addition, all routes serving ActivityPub are CORS-enabled for all origins.

API version: 7.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package peertube_api_sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the PlaybackMetricCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlaybackMetricCreate{}

// PlaybackMetricCreate struct for PlaybackMetricCreate
type PlaybackMetricCreate struct {
	PlayerMode string `json:"playerMode"`
	// Current player video resolution
	Resolution *float32 `json:"resolution,omitempty"`
	// Current player video fps
	Fps        *float32 `json:"fps,omitempty"`
	P2pEnabled bool     `json:"p2pEnabled"`
	// P2P peers connected (doesn't include WebSeed peers)
	P2pPeers *float32 `json:"p2pPeers,omitempty"`
	// How many resolution changes occurred since the last metric creation
	ResolutionChanges float32 `json:"resolutionChanges"`
	// How many times buffer has been stalled since the last metric creation
	BufferStalled *float32 `json:"bufferStalled,omitempty"`
	// How many errors occurred since the last metric creation
	Errors float32 `json:"errors"`
	// How many bytes were downloaded with P2P since the last metric creation
	DownloadedBytesP2P float32 `json:"downloadedBytesP2P"`
	// How many bytes were downloaded with HTTP since the last metric creation
	DownloadedBytesHTTP float32 `json:"downloadedBytesHTTP"`
	// How many bytes were uploaded with P2P since the last metric creation
	UploadedBytesP2P float32                                     `json:"uploadedBytesP2P"`
	VideoId          ApiV1VideosOwnershipIdAcceptPostIdParameter `json:"videoId"`
}

type _PlaybackMetricCreate PlaybackMetricCreate

// NewPlaybackMetricCreate instantiates a new PlaybackMetricCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlaybackMetricCreate(playerMode string, p2pEnabled bool, resolutionChanges float32, errors float32, downloadedBytesP2P float32, downloadedBytesHTTP float32, uploadedBytesP2P float32, videoId ApiV1VideosOwnershipIdAcceptPostIdParameter) *PlaybackMetricCreate {
	this := PlaybackMetricCreate{}
	this.PlayerMode = playerMode
	this.P2pEnabled = p2pEnabled
	this.ResolutionChanges = resolutionChanges
	this.Errors = errors
	this.DownloadedBytesP2P = downloadedBytesP2P
	this.DownloadedBytesHTTP = downloadedBytesHTTP
	this.UploadedBytesP2P = uploadedBytesP2P
	this.VideoId = videoId
	return &this
}

// NewPlaybackMetricCreateWithDefaults instantiates a new PlaybackMetricCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlaybackMetricCreateWithDefaults() *PlaybackMetricCreate {
	this := PlaybackMetricCreate{}
	return &this
}

// GetPlayerMode returns the PlayerMode field value
func (o *PlaybackMetricCreate) GetPlayerMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlayerMode
}

// GetPlayerModeOk returns a tuple with the PlayerMode field value
// and a boolean to check if the value has been set.
func (o *PlaybackMetricCreate) GetPlayerModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlayerMode, true
}

// SetPlayerMode sets field value
func (o *PlaybackMetricCreate) SetPlayerMode(v string) {
	o.PlayerMode = v
}

// GetResolution returns the Resolution field value if set, zero value otherwise.
func (o *PlaybackMetricCreate) GetResolution() float32 {
	if o == nil || IsNil(o.Resolution) {
		var ret float32
		return ret
	}
	return *o.Resolution
}

// GetResolutionOk returns a tuple with the Resolution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaybackMetricCreate) GetResolutionOk() (*float32, bool) {
	if o == nil || IsNil(o.Resolution) {
		return nil, false
	}
	return o.Resolution, true
}

// HasResolution returns a boolean if a field has been set.
func (o *PlaybackMetricCreate) HasResolution() bool {
	if o != nil && !IsNil(o.Resolution) {
		return true
	}

	return false
}

// SetResolution gets a reference to the given float32 and assigns it to the Resolution field.
func (o *PlaybackMetricCreate) SetResolution(v float32) {
	o.Resolution = &v
}

// GetFps returns the Fps field value if set, zero value otherwise.
func (o *PlaybackMetricCreate) GetFps() float32 {
	if o == nil || IsNil(o.Fps) {
		var ret float32
		return ret
	}
	return *o.Fps
}

// GetFpsOk returns a tuple with the Fps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaybackMetricCreate) GetFpsOk() (*float32, bool) {
	if o == nil || IsNil(o.Fps) {
		return nil, false
	}
	return o.Fps, true
}

// HasFps returns a boolean if a field has been set.
func (o *PlaybackMetricCreate) HasFps() bool {
	if o != nil && !IsNil(o.Fps) {
		return true
	}

	return false
}

// SetFps gets a reference to the given float32 and assigns it to the Fps field.
func (o *PlaybackMetricCreate) SetFps(v float32) {
	o.Fps = &v
}

// GetP2pEnabled returns the P2pEnabled field value
func (o *PlaybackMetricCreate) GetP2pEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.P2pEnabled
}

// GetP2pEnabledOk returns a tuple with the P2pEnabled field value
// and a boolean to check if the value has been set.
func (o *PlaybackMetricCreate) GetP2pEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.P2pEnabled, true
}

// SetP2pEnabled sets field value
func (o *PlaybackMetricCreate) SetP2pEnabled(v bool) {
	o.P2pEnabled = v
}

// GetP2pPeers returns the P2pPeers field value if set, zero value otherwise.
func (o *PlaybackMetricCreate) GetP2pPeers() float32 {
	if o == nil || IsNil(o.P2pPeers) {
		var ret float32
		return ret
	}
	return *o.P2pPeers
}

// GetP2pPeersOk returns a tuple with the P2pPeers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaybackMetricCreate) GetP2pPeersOk() (*float32, bool) {
	if o == nil || IsNil(o.P2pPeers) {
		return nil, false
	}
	return o.P2pPeers, true
}

// HasP2pPeers returns a boolean if a field has been set.
func (o *PlaybackMetricCreate) HasP2pPeers() bool {
	if o != nil && !IsNil(o.P2pPeers) {
		return true
	}

	return false
}

// SetP2pPeers gets a reference to the given float32 and assigns it to the P2pPeers field.
func (o *PlaybackMetricCreate) SetP2pPeers(v float32) {
	o.P2pPeers = &v
}

// GetResolutionChanges returns the ResolutionChanges field value
func (o *PlaybackMetricCreate) GetResolutionChanges() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ResolutionChanges
}

// GetResolutionChangesOk returns a tuple with the ResolutionChanges field value
// and a boolean to check if the value has been set.
func (o *PlaybackMetricCreate) GetResolutionChangesOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResolutionChanges, true
}

// SetResolutionChanges sets field value
func (o *PlaybackMetricCreate) SetResolutionChanges(v float32) {
	o.ResolutionChanges = v
}

// GetBufferStalled returns the BufferStalled field value if set, zero value otherwise.
func (o *PlaybackMetricCreate) GetBufferStalled() float32 {
	if o == nil || IsNil(o.BufferStalled) {
		var ret float32
		return ret
	}
	return *o.BufferStalled
}

// GetBufferStalledOk returns a tuple with the BufferStalled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlaybackMetricCreate) GetBufferStalledOk() (*float32, bool) {
	if o == nil || IsNil(o.BufferStalled) {
		return nil, false
	}
	return o.BufferStalled, true
}

// HasBufferStalled returns a boolean if a field has been set.
func (o *PlaybackMetricCreate) HasBufferStalled() bool {
	if o != nil && !IsNil(o.BufferStalled) {
		return true
	}

	return false
}

// SetBufferStalled gets a reference to the given float32 and assigns it to the BufferStalled field.
func (o *PlaybackMetricCreate) SetBufferStalled(v float32) {
	o.BufferStalled = &v
}

// GetErrors returns the Errors field value
func (o *PlaybackMetricCreate) GetErrors() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *PlaybackMetricCreate) GetErrorsOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Errors, true
}

// SetErrors sets field value
func (o *PlaybackMetricCreate) SetErrors(v float32) {
	o.Errors = v
}

// GetDownloadedBytesP2P returns the DownloadedBytesP2P field value
func (o *PlaybackMetricCreate) GetDownloadedBytesP2P() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.DownloadedBytesP2P
}

// GetDownloadedBytesP2POk returns a tuple with the DownloadedBytesP2P field value
// and a boolean to check if the value has been set.
func (o *PlaybackMetricCreate) GetDownloadedBytesP2POk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DownloadedBytesP2P, true
}

// SetDownloadedBytesP2P sets field value
func (o *PlaybackMetricCreate) SetDownloadedBytesP2P(v float32) {
	o.DownloadedBytesP2P = v
}

// GetDownloadedBytesHTTP returns the DownloadedBytesHTTP field value
func (o *PlaybackMetricCreate) GetDownloadedBytesHTTP() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.DownloadedBytesHTTP
}

// GetDownloadedBytesHTTPOk returns a tuple with the DownloadedBytesHTTP field value
// and a boolean to check if the value has been set.
func (o *PlaybackMetricCreate) GetDownloadedBytesHTTPOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DownloadedBytesHTTP, true
}

// SetDownloadedBytesHTTP sets field value
func (o *PlaybackMetricCreate) SetDownloadedBytesHTTP(v float32) {
	o.DownloadedBytesHTTP = v
}

// GetUploadedBytesP2P returns the UploadedBytesP2P field value
func (o *PlaybackMetricCreate) GetUploadedBytesP2P() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.UploadedBytesP2P
}

// GetUploadedBytesP2POk returns a tuple with the UploadedBytesP2P field value
// and a boolean to check if the value has been set.
func (o *PlaybackMetricCreate) GetUploadedBytesP2POk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UploadedBytesP2P, true
}

// SetUploadedBytesP2P sets field value
func (o *PlaybackMetricCreate) SetUploadedBytesP2P(v float32) {
	o.UploadedBytesP2P = v
}

// GetVideoId returns the VideoId field value
func (o *PlaybackMetricCreate) GetVideoId() ApiV1VideosOwnershipIdAcceptPostIdParameter {
	if o == nil {
		var ret ApiV1VideosOwnershipIdAcceptPostIdParameter
		return ret
	}

	return o.VideoId
}

// GetVideoIdOk returns a tuple with the VideoId field value
// and a boolean to check if the value has been set.
func (o *PlaybackMetricCreate) GetVideoIdOk() (*ApiV1VideosOwnershipIdAcceptPostIdParameter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VideoId, true
}

// SetVideoId sets field value
func (o *PlaybackMetricCreate) SetVideoId(v ApiV1VideosOwnershipIdAcceptPostIdParameter) {
	o.VideoId = v
}

func (o PlaybackMetricCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlaybackMetricCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["playerMode"] = o.PlayerMode
	if !IsNil(o.Resolution) {
		toSerialize["resolution"] = o.Resolution
	}
	if !IsNil(o.Fps) {
		toSerialize["fps"] = o.Fps
	}
	toSerialize["p2pEnabled"] = o.P2pEnabled
	if !IsNil(o.P2pPeers) {
		toSerialize["p2pPeers"] = o.P2pPeers
	}
	toSerialize["resolutionChanges"] = o.ResolutionChanges
	if !IsNil(o.BufferStalled) {
		toSerialize["bufferStalled"] = o.BufferStalled
	}
	toSerialize["errors"] = o.Errors
	toSerialize["downloadedBytesP2P"] = o.DownloadedBytesP2P
	toSerialize["downloadedBytesHTTP"] = o.DownloadedBytesHTTP
	toSerialize["uploadedBytesP2P"] = o.UploadedBytesP2P
	toSerialize["videoId"] = o.VideoId
	return toSerialize, nil
}

func (o *PlaybackMetricCreate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"playerMode",
		"p2pEnabled",
		"resolutionChanges",
		"errors",
		"downloadedBytesP2P",
		"downloadedBytesHTTP",
		"uploadedBytesP2P",
		"videoId",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPlaybackMetricCreate := _PlaybackMetricCreate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPlaybackMetricCreate)

	if err != nil {
		return err
	}

	*o = PlaybackMetricCreate(varPlaybackMetricCreate)

	return err
}

type NullablePlaybackMetricCreate struct {
	value *PlaybackMetricCreate
	isSet bool
}

func (v NullablePlaybackMetricCreate) Get() *PlaybackMetricCreate {
	return v.value
}

func (v *NullablePlaybackMetricCreate) Set(val *PlaybackMetricCreate) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaybackMetricCreate) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaybackMetricCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaybackMetricCreate(val *PlaybackMetricCreate) *NullablePlaybackMetricCreate {
	return &NullablePlaybackMetricCreate{value: val, isSet: true}
}

func (v NullablePlaybackMetricCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaybackMetricCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
