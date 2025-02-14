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

// checks if the ServerConfigCustomTranscoding type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerConfigCustomTranscoding{}

// ServerConfigCustomTranscoding Settings pertaining to transcoding jobs
type ServerConfigCustomTranscoding struct {
	Enabled      *bool                                      `json:"enabled,omitempty"`
	OriginalFile *ServerConfigCustomTranscodingOriginalFile `json:"originalFile,omitempty"`
	// Allow your users to upload .mkv, .mov, .avi, .wmv, .flv, .f4v, .3g2, .3gp, .mts, m2ts, .mxf, .nut videos
	AllowAdditionalExtensions *bool `json:"allowAdditionalExtensions,omitempty"`
	// If a user uploads an audio file, PeerTube will create a video by merging the preview file and the audio file
	AllowAudioFiles *bool `json:"allowAudioFiles,omitempty"`
	// Amount of threads used by ffmpeg for 1 transcoding job
	Threads *int32 `json:"threads,omitempty"`
	// Amount of transcoding jobs to execute in parallel
	Concurrency *float32 `json:"concurrency,omitempty"`
	// New profiles can be added by plugins ; available in core PeerTube: 'default'.
	Profile     *string                                   `json:"profile,omitempty"`
	Resolutions *ServerConfigCustomTranscodingResolutions `json:"resolutions,omitempty"`
	WebVideos   *ServerConfigCustomTranscodingWebVideos   `json:"web_videos,omitempty"`
	Hls         *ServerConfigCustomTranscodingHls         `json:"hls,omitempty"`
}

// NewServerConfigCustomTranscoding instantiates a new ServerConfigCustomTranscoding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerConfigCustomTranscoding() *ServerConfigCustomTranscoding {
	this := ServerConfigCustomTranscoding{}
	return &this
}

// NewServerConfigCustomTranscodingWithDefaults instantiates a new ServerConfigCustomTranscoding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerConfigCustomTranscodingWithDefaults() *ServerConfigCustomTranscoding {
	this := ServerConfigCustomTranscoding{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ServerConfigCustomTranscoding) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigCustomTranscoding) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ServerConfigCustomTranscoding) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ServerConfigCustomTranscoding) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetOriginalFile returns the OriginalFile field value if set, zero value otherwise.
func (o *ServerConfigCustomTranscoding) GetOriginalFile() ServerConfigCustomTranscodingOriginalFile {
	if o == nil || IsNil(o.OriginalFile) {
		var ret ServerConfigCustomTranscodingOriginalFile
		return ret
	}
	return *o.OriginalFile
}

// GetOriginalFileOk returns a tuple with the OriginalFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigCustomTranscoding) GetOriginalFileOk() (*ServerConfigCustomTranscodingOriginalFile, bool) {
	if o == nil || IsNil(o.OriginalFile) {
		return nil, false
	}
	return o.OriginalFile, true
}

// HasOriginalFile returns a boolean if a field has been set.
func (o *ServerConfigCustomTranscoding) HasOriginalFile() bool {
	if o != nil && !IsNil(o.OriginalFile) {
		return true
	}

	return false
}

// SetOriginalFile gets a reference to the given ServerConfigCustomTranscodingOriginalFile and assigns it to the OriginalFile field.
func (o *ServerConfigCustomTranscoding) SetOriginalFile(v ServerConfigCustomTranscodingOriginalFile) {
	o.OriginalFile = &v
}

// GetAllowAdditionalExtensions returns the AllowAdditionalExtensions field value if set, zero value otherwise.
func (o *ServerConfigCustomTranscoding) GetAllowAdditionalExtensions() bool {
	if o == nil || IsNil(o.AllowAdditionalExtensions) {
		var ret bool
		return ret
	}
	return *o.AllowAdditionalExtensions
}

// GetAllowAdditionalExtensionsOk returns a tuple with the AllowAdditionalExtensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigCustomTranscoding) GetAllowAdditionalExtensionsOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowAdditionalExtensions) {
		return nil, false
	}
	return o.AllowAdditionalExtensions, true
}

// HasAllowAdditionalExtensions returns a boolean if a field has been set.
func (o *ServerConfigCustomTranscoding) HasAllowAdditionalExtensions() bool {
	if o != nil && !IsNil(o.AllowAdditionalExtensions) {
		return true
	}

	return false
}

// SetAllowAdditionalExtensions gets a reference to the given bool and assigns it to the AllowAdditionalExtensions field.
func (o *ServerConfigCustomTranscoding) SetAllowAdditionalExtensions(v bool) {
	o.AllowAdditionalExtensions = &v
}

// GetAllowAudioFiles returns the AllowAudioFiles field value if set, zero value otherwise.
func (o *ServerConfigCustomTranscoding) GetAllowAudioFiles() bool {
	if o == nil || IsNil(o.AllowAudioFiles) {
		var ret bool
		return ret
	}
	return *o.AllowAudioFiles
}

// GetAllowAudioFilesOk returns a tuple with the AllowAudioFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigCustomTranscoding) GetAllowAudioFilesOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowAudioFiles) {
		return nil, false
	}
	return o.AllowAudioFiles, true
}

// HasAllowAudioFiles returns a boolean if a field has been set.
func (o *ServerConfigCustomTranscoding) HasAllowAudioFiles() bool {
	if o != nil && !IsNil(o.AllowAudioFiles) {
		return true
	}

	return false
}

// SetAllowAudioFiles gets a reference to the given bool and assigns it to the AllowAudioFiles field.
func (o *ServerConfigCustomTranscoding) SetAllowAudioFiles(v bool) {
	o.AllowAudioFiles = &v
}

// GetThreads returns the Threads field value if set, zero value otherwise.
func (o *ServerConfigCustomTranscoding) GetThreads() int32 {
	if o == nil || IsNil(o.Threads) {
		var ret int32
		return ret
	}
	return *o.Threads
}

// GetThreadsOk returns a tuple with the Threads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigCustomTranscoding) GetThreadsOk() (*int32, bool) {
	if o == nil || IsNil(o.Threads) {
		return nil, false
	}
	return o.Threads, true
}

// HasThreads returns a boolean if a field has been set.
func (o *ServerConfigCustomTranscoding) HasThreads() bool {
	if o != nil && !IsNil(o.Threads) {
		return true
	}

	return false
}

// SetThreads gets a reference to the given int32 and assigns it to the Threads field.
func (o *ServerConfigCustomTranscoding) SetThreads(v int32) {
	o.Threads = &v
}

// GetConcurrency returns the Concurrency field value if set, zero value otherwise.
func (o *ServerConfigCustomTranscoding) GetConcurrency() float32 {
	if o == nil || IsNil(o.Concurrency) {
		var ret float32
		return ret
	}
	return *o.Concurrency
}

// GetConcurrencyOk returns a tuple with the Concurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigCustomTranscoding) GetConcurrencyOk() (*float32, bool) {
	if o == nil || IsNil(o.Concurrency) {
		return nil, false
	}
	return o.Concurrency, true
}

// HasConcurrency returns a boolean if a field has been set.
func (o *ServerConfigCustomTranscoding) HasConcurrency() bool {
	if o != nil && !IsNil(o.Concurrency) {
		return true
	}

	return false
}

// SetConcurrency gets a reference to the given float32 and assigns it to the Concurrency field.
func (o *ServerConfigCustomTranscoding) SetConcurrency(v float32) {
	o.Concurrency = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *ServerConfigCustomTranscoding) GetProfile() string {
	if o == nil || IsNil(o.Profile) {
		var ret string
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigCustomTranscoding) GetProfileOk() (*string, bool) {
	if o == nil || IsNil(o.Profile) {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *ServerConfigCustomTranscoding) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given string and assigns it to the Profile field.
func (o *ServerConfigCustomTranscoding) SetProfile(v string) {
	o.Profile = &v
}

// GetResolutions returns the Resolutions field value if set, zero value otherwise.
func (o *ServerConfigCustomTranscoding) GetResolutions() ServerConfigCustomTranscodingResolutions {
	if o == nil || IsNil(o.Resolutions) {
		var ret ServerConfigCustomTranscodingResolutions
		return ret
	}
	return *o.Resolutions
}

// GetResolutionsOk returns a tuple with the Resolutions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigCustomTranscoding) GetResolutionsOk() (*ServerConfigCustomTranscodingResolutions, bool) {
	if o == nil || IsNil(o.Resolutions) {
		return nil, false
	}
	return o.Resolutions, true
}

// HasResolutions returns a boolean if a field has been set.
func (o *ServerConfigCustomTranscoding) HasResolutions() bool {
	if o != nil && !IsNil(o.Resolutions) {
		return true
	}

	return false
}

// SetResolutions gets a reference to the given ServerConfigCustomTranscodingResolutions and assigns it to the Resolutions field.
func (o *ServerConfigCustomTranscoding) SetResolutions(v ServerConfigCustomTranscodingResolutions) {
	o.Resolutions = &v
}

// GetWebVideos returns the WebVideos field value if set, zero value otherwise.
func (o *ServerConfigCustomTranscoding) GetWebVideos() ServerConfigCustomTranscodingWebVideos {
	if o == nil || IsNil(o.WebVideos) {
		var ret ServerConfigCustomTranscodingWebVideos
		return ret
	}
	return *o.WebVideos
}

// GetWebVideosOk returns a tuple with the WebVideos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigCustomTranscoding) GetWebVideosOk() (*ServerConfigCustomTranscodingWebVideos, bool) {
	if o == nil || IsNil(o.WebVideos) {
		return nil, false
	}
	return o.WebVideos, true
}

// HasWebVideos returns a boolean if a field has been set.
func (o *ServerConfigCustomTranscoding) HasWebVideos() bool {
	if o != nil && !IsNil(o.WebVideos) {
		return true
	}

	return false
}

// SetWebVideos gets a reference to the given ServerConfigCustomTranscodingWebVideos and assigns it to the WebVideos field.
func (o *ServerConfigCustomTranscoding) SetWebVideos(v ServerConfigCustomTranscodingWebVideos) {
	o.WebVideos = &v
}

// GetHls returns the Hls field value if set, zero value otherwise.
func (o *ServerConfigCustomTranscoding) GetHls() ServerConfigCustomTranscodingHls {
	if o == nil || IsNil(o.Hls) {
		var ret ServerConfigCustomTranscodingHls
		return ret
	}
	return *o.Hls
}

// GetHlsOk returns a tuple with the Hls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigCustomTranscoding) GetHlsOk() (*ServerConfigCustomTranscodingHls, bool) {
	if o == nil || IsNil(o.Hls) {
		return nil, false
	}
	return o.Hls, true
}

// HasHls returns a boolean if a field has been set.
func (o *ServerConfigCustomTranscoding) HasHls() bool {
	if o != nil && !IsNil(o.Hls) {
		return true
	}

	return false
}

// SetHls gets a reference to the given ServerConfigCustomTranscodingHls and assigns it to the Hls field.
func (o *ServerConfigCustomTranscoding) SetHls(v ServerConfigCustomTranscodingHls) {
	o.Hls = &v
}

func (o ServerConfigCustomTranscoding) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerConfigCustomTranscoding) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.OriginalFile) {
		toSerialize["originalFile"] = o.OriginalFile
	}
	if !IsNil(o.AllowAdditionalExtensions) {
		toSerialize["allowAdditionalExtensions"] = o.AllowAdditionalExtensions
	}
	if !IsNil(o.AllowAudioFiles) {
		toSerialize["allowAudioFiles"] = o.AllowAudioFiles
	}
	if !IsNil(o.Threads) {
		toSerialize["threads"] = o.Threads
	}
	if !IsNil(o.Concurrency) {
		toSerialize["concurrency"] = o.Concurrency
	}
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}
	if !IsNil(o.Resolutions) {
		toSerialize["resolutions"] = o.Resolutions
	}
	if !IsNil(o.WebVideos) {
		toSerialize["web_videos"] = o.WebVideos
	}
	if !IsNil(o.Hls) {
		toSerialize["hls"] = o.Hls
	}
	return toSerialize, nil
}

type NullableServerConfigCustomTranscoding struct {
	value *ServerConfigCustomTranscoding
	isSet bool
}

func (v NullableServerConfigCustomTranscoding) Get() *ServerConfigCustomTranscoding {
	return v.value
}

func (v *NullableServerConfigCustomTranscoding) Set(val *ServerConfigCustomTranscoding) {
	v.value = val
	v.isSet = true
}

func (v NullableServerConfigCustomTranscoding) IsSet() bool {
	return v.isSet
}

func (v *NullableServerConfigCustomTranscoding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerConfigCustomTranscoding(val *ServerConfigCustomTranscoding) *NullableServerConfigCustomTranscoding {
	return &NullableServerConfigCustomTranscoding{value: val, isSet: true}
}

func (v NullableServerConfigCustomTranscoding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerConfigCustomTranscoding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
