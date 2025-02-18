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
	"time"
)

// checks if the VideoStatsOverall type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &VideoStatsOverall{}

// VideoStatsOverall struct for VideoStatsOverall
type VideoStatsOverall struct {
	AverageWatchTime *float32                          `json:"averageWatchTime,omitempty"`
	TotalWatchTime   *float32                          `json:"totalWatchTime,omitempty"`
	ViewersPeak      *float32                          `json:"viewersPeak,omitempty"`
	ViewersPeakDate  *time.Time                        `json:"viewersPeakDate,omitempty"`
	Countries        []VideoStatsOverallCountriesInner `json:"countries,omitempty"`
}

// NewVideoStatsOverall instantiates a new VideoStatsOverall object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideoStatsOverall() *VideoStatsOverall {
	this := VideoStatsOverall{}
	return &this
}

// NewVideoStatsOverallWithDefaults instantiates a new VideoStatsOverall object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoStatsOverallWithDefaults() *VideoStatsOverall {
	this := VideoStatsOverall{}
	return &this
}

// GetAverageWatchTime returns the AverageWatchTime field value if set, zero value otherwise.
func (o *VideoStatsOverall) GetAverageWatchTime() float32 {
	if o == nil || utils.IsNil(o.AverageWatchTime) {
		var ret float32
		return ret
	}
	return *o.AverageWatchTime
}

// GetAverageWatchTimeOk returns a tuple with the AverageWatchTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoStatsOverall) GetAverageWatchTimeOk() (*float32, bool) {
	if o == nil || utils.IsNil(o.AverageWatchTime) {
		return nil, false
	}
	return o.AverageWatchTime, true
}

// HasAverageWatchTime returns a boolean if a field has been set.
func (o *VideoStatsOverall) HasAverageWatchTime() bool {
	if o != nil && !utils.IsNil(o.AverageWatchTime) {
		return true
	}

	return false
}

// SetAverageWatchTime gets a reference to the given float32 and assigns it to the AverageWatchTime field.
func (o *VideoStatsOverall) SetAverageWatchTime(v float32) {
	o.AverageWatchTime = &v
}

// GetTotalWatchTime returns the TotalWatchTime field value if set, zero value otherwise.
func (o *VideoStatsOverall) GetTotalWatchTime() float32 {
	if o == nil || utils.IsNil(o.TotalWatchTime) {
		var ret float32
		return ret
	}
	return *o.TotalWatchTime
}

// GetTotalWatchTimeOk returns a tuple with the TotalWatchTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoStatsOverall) GetTotalWatchTimeOk() (*float32, bool) {
	if o == nil || utils.IsNil(o.TotalWatchTime) {
		return nil, false
	}
	return o.TotalWatchTime, true
}

// HasTotalWatchTime returns a boolean if a field has been set.
func (o *VideoStatsOverall) HasTotalWatchTime() bool {
	if o != nil && !utils.IsNil(o.TotalWatchTime) {
		return true
	}

	return false
}

// SetTotalWatchTime gets a reference to the given float32 and assigns it to the TotalWatchTime field.
func (o *VideoStatsOverall) SetTotalWatchTime(v float32) {
	o.TotalWatchTime = &v
}

// GetViewersPeak returns the ViewersPeak field value if set, zero value otherwise.
func (o *VideoStatsOverall) GetViewersPeak() float32 {
	if o == nil || utils.IsNil(o.ViewersPeak) {
		var ret float32
		return ret
	}
	return *o.ViewersPeak
}

// GetViewersPeakOk returns a tuple with the ViewersPeak field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoStatsOverall) GetViewersPeakOk() (*float32, bool) {
	if o == nil || utils.IsNil(o.ViewersPeak) {
		return nil, false
	}
	return o.ViewersPeak, true
}

// HasViewersPeak returns a boolean if a field has been set.
func (o *VideoStatsOverall) HasViewersPeak() bool {
	if o != nil && !utils.IsNil(o.ViewersPeak) {
		return true
	}

	return false
}

// SetViewersPeak gets a reference to the given float32 and assigns it to the ViewersPeak field.
func (o *VideoStatsOverall) SetViewersPeak(v float32) {
	o.ViewersPeak = &v
}

// GetViewersPeakDate returns the ViewersPeakDate field value if set, zero value otherwise.
func (o *VideoStatsOverall) GetViewersPeakDate() time.Time {
	if o == nil || utils.IsNil(o.ViewersPeakDate) {
		var ret time.Time
		return ret
	}
	return *o.ViewersPeakDate
}

// GetViewersPeakDateOk returns a tuple with the ViewersPeakDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoStatsOverall) GetViewersPeakDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.ViewersPeakDate) {
		return nil, false
	}
	return o.ViewersPeakDate, true
}

// HasViewersPeakDate returns a boolean if a field has been set.
func (o *VideoStatsOverall) HasViewersPeakDate() bool {
	if o != nil && !utils.IsNil(o.ViewersPeakDate) {
		return true
	}

	return false
}

// SetViewersPeakDate gets a reference to the given time.Time and assigns it to the ViewersPeakDate field.
func (o *VideoStatsOverall) SetViewersPeakDate(v time.Time) {
	o.ViewersPeakDate = &v
}

// GetCountries returns the Countries field value if set, zero value otherwise.
func (o *VideoStatsOverall) GetCountries() []VideoStatsOverallCountriesInner {
	if o == nil || utils.IsNil(o.Countries) {
		var ret []VideoStatsOverallCountriesInner
		return ret
	}
	return o.Countries
}

// GetCountriesOk returns a tuple with the Countries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoStatsOverall) GetCountriesOk() ([]VideoStatsOverallCountriesInner, bool) {
	if o == nil || utils.IsNil(o.Countries) {
		return nil, false
	}
	return o.Countries, true
}

// HasCountries returns a boolean if a field has been set.
func (o *VideoStatsOverall) HasCountries() bool {
	if o != nil && !utils.IsNil(o.Countries) {
		return true
	}

	return false
}

// SetCountries gets a reference to the given []VideoStatsOverallCountriesInner and assigns it to the Countries field.
func (o *VideoStatsOverall) SetCountries(v []VideoStatsOverallCountriesInner) {
	o.Countries = v
}

func (o VideoStatsOverall) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VideoStatsOverall) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.AverageWatchTime) {
		toSerialize["averageWatchTime"] = o.AverageWatchTime
	}
	if !utils.IsNil(o.TotalWatchTime) {
		toSerialize["totalWatchTime"] = o.TotalWatchTime
	}
	if !utils.IsNil(o.ViewersPeak) {
		toSerialize["viewersPeak"] = o.ViewersPeak
	}
	if !utils.IsNil(o.ViewersPeakDate) {
		toSerialize["viewersPeakDate"] = o.ViewersPeakDate
	}
	if !utils.IsNil(o.Countries) {
		toSerialize["countries"] = o.Countries
	}
	return toSerialize, nil
}

type NullableVideoStatsOverall struct {
	value *VideoStatsOverall
	isSet bool
}

func (v NullableVideoStatsOverall) Get() *VideoStatsOverall {
	return v.value
}

func (v *NullableVideoStatsOverall) Set(val *VideoStatsOverall) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoStatsOverall) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoStatsOverall) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoStatsOverall(val *VideoStatsOverall) *NullableVideoStatsOverall {
	return &NullableVideoStatsOverall{value: val, isSet: true}
}

func (v NullableVideoStatsOverall) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVideoStatsOverall) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
