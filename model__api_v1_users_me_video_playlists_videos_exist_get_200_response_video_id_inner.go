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

// checks if the ApiV1UsersMeVideoPlaylistsVideosExistGet200ResponseVideoIdInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiV1UsersMeVideoPlaylistsVideosExistGet200ResponseVideoIdInner{}

// ApiV1UsersMeVideoPlaylistsVideosExistGet200ResponseVideoIdInner struct for ApiV1UsersMeVideoPlaylistsVideosExistGet200ResponseVideoIdInner
type ApiV1UsersMeVideoPlaylistsVideosExistGet200ResponseVideoIdInner struct {
	PlaylistElementId *int32 `json:"playlistElementId,omitempty"`
	PlaylistId        *int32 `json:"playlistId,omitempty"`
	StartTimestamp    *int32 `json:"startTimestamp,omitempty"`
	StopTimestamp     *int32 `json:"stopTimestamp,omitempty"`
}

// NewApiV1UsersMeVideoPlaylistsVideosExistGet200ResponseVideoIdInner instantiates a new ApiV1UsersMeVideoPlaylistsVideosExistGet200ResponseVideoIdInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiV1UsersMeVideoPlaylistsVideosExistGet200ResponseVideoIdInner() *ApiV1UsersMeVideoPlaylistsVideosExistGet200ResponseVideoIdInner {
	this := ApiV1UsersMeVideoPlaylistsVideosExistGet200ResponseVideoIdInner{}
	return &this
}

// NewApiV1UsersMeVideoPlaylistsVideosExistGet200ResponseVideoIdInnerWithDefaults instantiates a new ApiV1UsersMeVideoPlaylistsVideosExistGet200ResponseVideoIdInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiV1UsersMeVideoPlaylistsVideosExistGet200ResponseVideoIdInnerWithDefaults() *ApiV1UsersMeVideoPlaylistsVideosExistGet200ResponseVideoIdInner {
	this := ApiV1UsersMeVideoPlaylistsVideosExistGet200ResponseVideoIdInner{}
	return &this
}

// GetPlaylistElementId returns the PlaylistElementId field value if set, zero value otherwise.
func (o *ApiV1UsersMeVideoPlaylistsVideosExistGet200ResponseVideoIdInner) GetPlaylistElementId() int32 {
	if o == nil || IsNil(o.PlaylistElementId) {
		var ret int32
		return ret
	}
	return *o.PlaylistElementId
}

// GetPlaylistElementIdOk returns a tuple with the PlaylistElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1UsersMeVideoPlaylistsVideosExistGet200ResponseVideoIdInner) GetPlaylistElementIdOk() (*int32, bool) {
	if o == nil || IsNil(o.PlaylistElementId) {
		return nil, false
	}
	return o.PlaylistElementId, true
}

// HasPlaylistElementId returns a boolean if a field has been set.
func (o *ApiV1UsersMeVideoPlaylistsVideosExistGet200ResponseVideoIdInner) HasPlaylistElementId() bool {
	if o != nil && !IsNil(o.PlaylistElementId) {
		return true
	}

	return false
}

// SetPlaylistElementId gets a reference to the given int32 and assigns it to the PlaylistElementId field.
func (o *ApiV1UsersMeVideoPlaylistsVideosExistGet200ResponseVideoIdInner) SetPlaylistElementId(v int32) {
	o.PlaylistElementId = &v
}

// GetPlaylistId returns the PlaylistId field value if set, zero value otherwise.
func (o *ApiV1UsersMeVideoPlaylistsVideosExistGet200ResponseVideoIdInner) GetPlaylistId() int32 {
	if o == nil || IsNil(o.PlaylistId) {
		var ret int32
		return ret
	}
	return *o.PlaylistId
}

// GetPlaylistIdOk returns a tuple with the PlaylistId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1UsersMeVideoPlaylistsVideosExistGet200ResponseVideoIdInner) GetPlaylistIdOk() (*int32, bool) {
	if o == nil || IsNil(o.PlaylistId) {
		return nil, false
	}
	return o.PlaylistId, true
}

// HasPlaylistId returns a boolean if a field has been set.
func (o *ApiV1UsersMeVideoPlaylistsVideosExistGet200ResponseVideoIdInner) HasPlaylistId() bool {
	if o != nil && !IsNil(o.PlaylistId) {
		return true
	}

	return false
}

// SetPlaylistId gets a reference to the given int32 and assigns it to the PlaylistId field.
func (o *ApiV1UsersMeVideoPlaylistsVideosExistGet200ResponseVideoIdInner) SetPlaylistId(v int32) {
	o.PlaylistId = &v
}

// GetStartTimestamp returns the StartTimestamp field value if set, zero value otherwise.
func (o *ApiV1UsersMeVideoPlaylistsVideosExistGet200ResponseVideoIdInner) GetStartTimestamp() int32 {
	if o == nil || IsNil(o.StartTimestamp) {
		var ret int32
		return ret
	}
	return *o.StartTimestamp
}

// GetStartTimestampOk returns a tuple with the StartTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1UsersMeVideoPlaylistsVideosExistGet200ResponseVideoIdInner) GetStartTimestampOk() (*int32, bool) {
	if o == nil || IsNil(o.StartTimestamp) {
		return nil, false
	}
	return o.StartTimestamp, true
}

// HasStartTimestamp returns a boolean if a field has been set.
func (o *ApiV1UsersMeVideoPlaylistsVideosExistGet200ResponseVideoIdInner) HasStartTimestamp() bool {
	if o != nil && !IsNil(o.StartTimestamp) {
		return true
	}

	return false
}

// SetStartTimestamp gets a reference to the given int32 and assigns it to the StartTimestamp field.
func (o *ApiV1UsersMeVideoPlaylistsVideosExistGet200ResponseVideoIdInner) SetStartTimestamp(v int32) {
	o.StartTimestamp = &v
}

// GetStopTimestamp returns the StopTimestamp field value if set, zero value otherwise.
func (o *ApiV1UsersMeVideoPlaylistsVideosExistGet200ResponseVideoIdInner) GetStopTimestamp() int32 {
	if o == nil || IsNil(o.StopTimestamp) {
		var ret int32
		return ret
	}
	return *o.StopTimestamp
}

// GetStopTimestampOk returns a tuple with the StopTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1UsersMeVideoPlaylistsVideosExistGet200ResponseVideoIdInner) GetStopTimestampOk() (*int32, bool) {
	if o == nil || IsNil(o.StopTimestamp) {
		return nil, false
	}
	return o.StopTimestamp, true
}

// HasStopTimestamp returns a boolean if a field has been set.
func (o *ApiV1UsersMeVideoPlaylistsVideosExistGet200ResponseVideoIdInner) HasStopTimestamp() bool {
	if o != nil && !IsNil(o.StopTimestamp) {
		return true
	}

	return false
}

// SetStopTimestamp gets a reference to the given int32 and assigns it to the StopTimestamp field.
func (o *ApiV1UsersMeVideoPlaylistsVideosExistGet200ResponseVideoIdInner) SetStopTimestamp(v int32) {
	o.StopTimestamp = &v
}

func (o ApiV1UsersMeVideoPlaylistsVideosExistGet200ResponseVideoIdInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiV1UsersMeVideoPlaylistsVideosExistGet200ResponseVideoIdInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PlaylistElementId) {
		toSerialize["playlistElementId"] = o.PlaylistElementId
	}
	if !IsNil(o.PlaylistId) {
		toSerialize["playlistId"] = o.PlaylistId
	}
	if !IsNil(o.StartTimestamp) {
		toSerialize["startTimestamp"] = o.StartTimestamp
	}
	if !IsNil(o.StopTimestamp) {
		toSerialize["stopTimestamp"] = o.StopTimestamp
	}
	return toSerialize, nil
}

type NullableApiV1UsersMeVideoPlaylistsVideosExistGet200ResponseVideoIdInner struct {
	value *ApiV1UsersMeVideoPlaylistsVideosExistGet200ResponseVideoIdInner
	isSet bool
}

func (v NullableApiV1UsersMeVideoPlaylistsVideosExistGet200ResponseVideoIdInner) Get() *ApiV1UsersMeVideoPlaylistsVideosExistGet200ResponseVideoIdInner {
	return v.value
}

func (v *NullableApiV1UsersMeVideoPlaylistsVideosExistGet200ResponseVideoIdInner) Set(val *ApiV1UsersMeVideoPlaylistsVideosExistGet200ResponseVideoIdInner) {
	v.value = val
	v.isSet = true
}

func (v NullableApiV1UsersMeVideoPlaylistsVideosExistGet200ResponseVideoIdInner) IsSet() bool {
	return v.isSet
}

func (v *NullableApiV1UsersMeVideoPlaylistsVideosExistGet200ResponseVideoIdInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiV1UsersMeVideoPlaylistsVideosExistGet200ResponseVideoIdInner(val *ApiV1UsersMeVideoPlaylistsVideosExistGet200ResponseVideoIdInner) *NullableApiV1UsersMeVideoPlaylistsVideosExistGet200ResponseVideoIdInner {
	return &NullableApiV1UsersMeVideoPlaylistsVideosExistGet200ResponseVideoIdInner{value: val, isSet: true}
}

func (v NullableApiV1UsersMeVideoPlaylistsVideosExistGet200ResponseVideoIdInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiV1UsersMeVideoPlaylistsVideosExistGet200ResponseVideoIdInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
