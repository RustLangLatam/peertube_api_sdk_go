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

// checks if the VideoStatsRetentionDataInner type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &VideoStatsRetentionDataInner{}

// VideoStatsRetentionDataInner struct for VideoStatsRetentionDataInner
type VideoStatsRetentionDataInner struct {
	Second           *float32 `json:"second,omitempty"`
	RetentionPercent *float32 `json:"retentionPercent,omitempty"`
}

// NewVideoStatsRetentionDataInner instantiates a new VideoStatsRetentionDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideoStatsRetentionDataInner() *VideoStatsRetentionDataInner {
	this := VideoStatsRetentionDataInner{}
	return &this
}

// NewVideoStatsRetentionDataInnerWithDefaults instantiates a new VideoStatsRetentionDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoStatsRetentionDataInnerWithDefaults() *VideoStatsRetentionDataInner {
	this := VideoStatsRetentionDataInner{}
	return &this
}

// GetSecond returns the Second field value if set, zero value otherwise.
func (o *VideoStatsRetentionDataInner) GetSecond() float32 {
	if o == nil || utils.IsNil(o.Second) {
		var ret float32
		return ret
	}
	return *o.Second
}

// GetSecondOk returns a tuple with the Second field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoStatsRetentionDataInner) GetSecondOk() (*float32, bool) {
	if o == nil || utils.IsNil(o.Second) {
		return nil, false
	}
	return o.Second, true
}

// HasSecond returns a boolean if a field has been set.
func (o *VideoStatsRetentionDataInner) HasSecond() bool {
	if o != nil && !utils.IsNil(o.Second) {
		return true
	}

	return false
}

// SetSecond gets a reference to the given float32 and assigns it to the Second field.
func (o *VideoStatsRetentionDataInner) SetSecond(v float32) {
	o.Second = &v
}

// GetRetentionPercent returns the RetentionPercent field value if set, zero value otherwise.
func (o *VideoStatsRetentionDataInner) GetRetentionPercent() float32 {
	if o == nil || utils.IsNil(o.RetentionPercent) {
		var ret float32
		return ret
	}
	return *o.RetentionPercent
}

// GetRetentionPercentOk returns a tuple with the RetentionPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoStatsRetentionDataInner) GetRetentionPercentOk() (*float32, bool) {
	if o == nil || utils.IsNil(o.RetentionPercent) {
		return nil, false
	}
	return o.RetentionPercent, true
}

// HasRetentionPercent returns a boolean if a field has been set.
func (o *VideoStatsRetentionDataInner) HasRetentionPercent() bool {
	if o != nil && !utils.IsNil(o.RetentionPercent) {
		return true
	}

	return false
}

// SetRetentionPercent gets a reference to the given float32 and assigns it to the RetentionPercent field.
func (o *VideoStatsRetentionDataInner) SetRetentionPercent(v float32) {
	o.RetentionPercent = &v
}

func (o VideoStatsRetentionDataInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VideoStatsRetentionDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Second) {
		toSerialize["second"] = o.Second
	}
	if !utils.IsNil(o.RetentionPercent) {
		toSerialize["retentionPercent"] = o.RetentionPercent
	}
	return toSerialize, nil
}

type NullableVideoStatsRetentionDataInner struct {
	value *VideoStatsRetentionDataInner
	isSet bool
}

func (v NullableVideoStatsRetentionDataInner) Get() *VideoStatsRetentionDataInner {
	return v.value
}

func (v *NullableVideoStatsRetentionDataInner) Set(val *VideoStatsRetentionDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoStatsRetentionDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoStatsRetentionDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoStatsRetentionDataInner(val *VideoStatsRetentionDataInner) *NullableVideoStatsRetentionDataInner {
	return &NullableVideoStatsRetentionDataInner{value: val, isSet: true}
}

func (v NullableVideoStatsRetentionDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVideoStatsRetentionDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
