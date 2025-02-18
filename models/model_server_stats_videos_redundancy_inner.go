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

// checks if the ServerStatsVideosRedundancyInner type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ServerStatsVideosRedundancyInner{}

// ServerStatsVideosRedundancyInner struct for ServerStatsVideosRedundancyInner
type ServerStatsVideosRedundancyInner struct {
	Strategy        *string  `json:"strategy,omitempty"`
	TotalSize       *float32 `json:"totalSize,omitempty"`
	TotalUsed       *float32 `json:"totalUsed,omitempty"`
	TotalVideoFiles *float32 `json:"totalVideoFiles,omitempty"`
	TotalVideos     *float32 `json:"totalVideos,omitempty"`
}

// NewServerStatsVideosRedundancyInner instantiates a new ServerStatsVideosRedundancyInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerStatsVideosRedundancyInner() *ServerStatsVideosRedundancyInner {
	this := ServerStatsVideosRedundancyInner{}
	return &this
}

// NewServerStatsVideosRedundancyInnerWithDefaults instantiates a new ServerStatsVideosRedundancyInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerStatsVideosRedundancyInnerWithDefaults() *ServerStatsVideosRedundancyInner {
	this := ServerStatsVideosRedundancyInner{}
	return &this
}

// GetStrategy returns the Strategy field value if set, zero value otherwise.
func (o *ServerStatsVideosRedundancyInner) GetStrategy() string {
	if o == nil || utils.IsNil(o.Strategy) {
		var ret string
		return ret
	}
	return *o.Strategy
}

// GetStrategyOk returns a tuple with the Strategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerStatsVideosRedundancyInner) GetStrategyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Strategy) {
		return nil, false
	}
	return o.Strategy, true
}

// HasStrategy returns a boolean if a field has been set.
func (o *ServerStatsVideosRedundancyInner) HasStrategy() bool {
	if o != nil && !utils.IsNil(o.Strategy) {
		return true
	}

	return false
}

// SetStrategy gets a reference to the given string and assigns it to the Strategy field.
func (o *ServerStatsVideosRedundancyInner) SetStrategy(v string) {
	o.Strategy = &v
}

// GetTotalSize returns the TotalSize field value if set, zero value otherwise.
func (o *ServerStatsVideosRedundancyInner) GetTotalSize() float32 {
	if o == nil || utils.IsNil(o.TotalSize) {
		var ret float32
		return ret
	}
	return *o.TotalSize
}

// GetTotalSizeOk returns a tuple with the TotalSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerStatsVideosRedundancyInner) GetTotalSizeOk() (*float32, bool) {
	if o == nil || utils.IsNil(o.TotalSize) {
		return nil, false
	}
	return o.TotalSize, true
}

// HasTotalSize returns a boolean if a field has been set.
func (o *ServerStatsVideosRedundancyInner) HasTotalSize() bool {
	if o != nil && !utils.IsNil(o.TotalSize) {
		return true
	}

	return false
}

// SetTotalSize gets a reference to the given float32 and assigns it to the TotalSize field.
func (o *ServerStatsVideosRedundancyInner) SetTotalSize(v float32) {
	o.TotalSize = &v
}

// GetTotalUsed returns the TotalUsed field value if set, zero value otherwise.
func (o *ServerStatsVideosRedundancyInner) GetTotalUsed() float32 {
	if o == nil || utils.IsNil(o.TotalUsed) {
		var ret float32
		return ret
	}
	return *o.TotalUsed
}

// GetTotalUsedOk returns a tuple with the TotalUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerStatsVideosRedundancyInner) GetTotalUsedOk() (*float32, bool) {
	if o == nil || utils.IsNil(o.TotalUsed) {
		return nil, false
	}
	return o.TotalUsed, true
}

// HasTotalUsed returns a boolean if a field has been set.
func (o *ServerStatsVideosRedundancyInner) HasTotalUsed() bool {
	if o != nil && !utils.IsNil(o.TotalUsed) {
		return true
	}

	return false
}

// SetTotalUsed gets a reference to the given float32 and assigns it to the TotalUsed field.
func (o *ServerStatsVideosRedundancyInner) SetTotalUsed(v float32) {
	o.TotalUsed = &v
}

// GetTotalVideoFiles returns the TotalVideoFiles field value if set, zero value otherwise.
func (o *ServerStatsVideosRedundancyInner) GetTotalVideoFiles() float32 {
	if o == nil || utils.IsNil(o.TotalVideoFiles) {
		var ret float32
		return ret
	}
	return *o.TotalVideoFiles
}

// GetTotalVideoFilesOk returns a tuple with the TotalVideoFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerStatsVideosRedundancyInner) GetTotalVideoFilesOk() (*float32, bool) {
	if o == nil || utils.IsNil(o.TotalVideoFiles) {
		return nil, false
	}
	return o.TotalVideoFiles, true
}

// HasTotalVideoFiles returns a boolean if a field has been set.
func (o *ServerStatsVideosRedundancyInner) HasTotalVideoFiles() bool {
	if o != nil && !utils.IsNil(o.TotalVideoFiles) {
		return true
	}

	return false
}

// SetTotalVideoFiles gets a reference to the given float32 and assigns it to the TotalVideoFiles field.
func (o *ServerStatsVideosRedundancyInner) SetTotalVideoFiles(v float32) {
	o.TotalVideoFiles = &v
}

// GetTotalVideos returns the TotalVideos field value if set, zero value otherwise.
func (o *ServerStatsVideosRedundancyInner) GetTotalVideos() float32 {
	if o == nil || utils.IsNil(o.TotalVideos) {
		var ret float32
		return ret
	}
	return *o.TotalVideos
}

// GetTotalVideosOk returns a tuple with the TotalVideos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerStatsVideosRedundancyInner) GetTotalVideosOk() (*float32, bool) {
	if o == nil || utils.IsNil(o.TotalVideos) {
		return nil, false
	}
	return o.TotalVideos, true
}

// HasTotalVideos returns a boolean if a field has been set.
func (o *ServerStatsVideosRedundancyInner) HasTotalVideos() bool {
	if o != nil && !utils.IsNil(o.TotalVideos) {
		return true
	}

	return false
}

// SetTotalVideos gets a reference to the given float32 and assigns it to the TotalVideos field.
func (o *ServerStatsVideosRedundancyInner) SetTotalVideos(v float32) {
	o.TotalVideos = &v
}

func (o ServerStatsVideosRedundancyInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerStatsVideosRedundancyInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Strategy) {
		toSerialize["strategy"] = o.Strategy
	}
	if !utils.IsNil(o.TotalSize) {
		toSerialize["totalSize"] = o.TotalSize
	}
	if !utils.IsNil(o.TotalUsed) {
		toSerialize["totalUsed"] = o.TotalUsed
	}
	if !utils.IsNil(o.TotalVideoFiles) {
		toSerialize["totalVideoFiles"] = o.TotalVideoFiles
	}
	if !utils.IsNil(o.TotalVideos) {
		toSerialize["totalVideos"] = o.TotalVideos
	}
	return toSerialize, nil
}

type NullableServerStatsVideosRedundancyInner struct {
	value *ServerStatsVideosRedundancyInner
	isSet bool
}

func (v NullableServerStatsVideosRedundancyInner) Get() *ServerStatsVideosRedundancyInner {
	return v.value
}

func (v *NullableServerStatsVideosRedundancyInner) Set(val *ServerStatsVideosRedundancyInner) {
	v.value = val
	v.isSet = true
}

func (v NullableServerStatsVideosRedundancyInner) IsSet() bool {
	return v.isSet
}

func (v *NullableServerStatsVideosRedundancyInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerStatsVideosRedundancyInner(val *ServerStatsVideosRedundancyInner) *NullableServerStatsVideosRedundancyInner {
	return &NullableServerStatsVideosRedundancyInner{value: val, isSet: true}
}

func (v NullableServerStatsVideosRedundancyInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerStatsVideosRedundancyInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
