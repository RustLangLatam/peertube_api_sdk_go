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

// checks if the ServerConfigCustomTranscodingResolutions type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ServerConfigCustomTranscodingResolutions{}

// ServerConfigCustomTranscodingResolutions Resolutions to transcode _new videos_ to
type ServerConfigCustomTranscodingResolutions struct {
	Var0p    *bool `json:"0p,omitempty"`
	Var144p  *bool `json:"144p,omitempty"`
	Var240p  *bool `json:"240p,omitempty"`
	Var360p  *bool `json:"360p,omitempty"`
	Var480p  *bool `json:"480p,omitempty"`
	Var720p  *bool `json:"720p,omitempty"`
	Var1080p *bool `json:"1080p,omitempty"`
	Var1440p *bool `json:"1440p,omitempty"`
	Var2160p *bool `json:"2160p,omitempty"`
}

// NewServerConfigCustomTranscodingResolutions instantiates a new ServerConfigCustomTranscodingResolutions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerConfigCustomTranscodingResolutions() *ServerConfigCustomTranscodingResolutions {
	this := ServerConfigCustomTranscodingResolutions{}
	return &this
}

// NewServerConfigCustomTranscodingResolutionsWithDefaults instantiates a new ServerConfigCustomTranscodingResolutions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerConfigCustomTranscodingResolutionsWithDefaults() *ServerConfigCustomTranscodingResolutions {
	this := ServerConfigCustomTranscodingResolutions{}
	return &this
}

// GetVar0p returns the Var0p field value if set, zero value otherwise.
func (o *ServerConfigCustomTranscodingResolutions) GetVar0p() bool {
	if o == nil || utils.IsNil(o.Var0p) {
		var ret bool
		return ret
	}
	return *o.Var0p
}

// GetVar0pOk returns a tuple with the Var0p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigCustomTranscodingResolutions) GetVar0pOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.Var0p) {
		return nil, false
	}
	return o.Var0p, true
}

// HasVar0p returns a boolean if a field has been set.
func (o *ServerConfigCustomTranscodingResolutions) HasVar0p() bool {
	if o != nil && !utils.IsNil(o.Var0p) {
		return true
	}

	return false
}

// SetVar0p gets a reference to the given bool and assigns it to the Var0p field.
func (o *ServerConfigCustomTranscodingResolutions) SetVar0p(v bool) {
	o.Var0p = &v
}

// GetVar144p returns the Var144p field value if set, zero value otherwise.
func (o *ServerConfigCustomTranscodingResolutions) GetVar144p() bool {
	if o == nil || utils.IsNil(o.Var144p) {
		var ret bool
		return ret
	}
	return *o.Var144p
}

// GetVar144pOk returns a tuple with the Var144p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigCustomTranscodingResolutions) GetVar144pOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.Var144p) {
		return nil, false
	}
	return o.Var144p, true
}

// HasVar144p returns a boolean if a field has been set.
func (o *ServerConfigCustomTranscodingResolutions) HasVar144p() bool {
	if o != nil && !utils.IsNil(o.Var144p) {
		return true
	}

	return false
}

// SetVar144p gets a reference to the given bool and assigns it to the Var144p field.
func (o *ServerConfigCustomTranscodingResolutions) SetVar144p(v bool) {
	o.Var144p = &v
}

// GetVar240p returns the Var240p field value if set, zero value otherwise.
func (o *ServerConfigCustomTranscodingResolutions) GetVar240p() bool {
	if o == nil || utils.IsNil(o.Var240p) {
		var ret bool
		return ret
	}
	return *o.Var240p
}

// GetVar240pOk returns a tuple with the Var240p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigCustomTranscodingResolutions) GetVar240pOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.Var240p) {
		return nil, false
	}
	return o.Var240p, true
}

// HasVar240p returns a boolean if a field has been set.
func (o *ServerConfigCustomTranscodingResolutions) HasVar240p() bool {
	if o != nil && !utils.IsNil(o.Var240p) {
		return true
	}

	return false
}

// SetVar240p gets a reference to the given bool and assigns it to the Var240p field.
func (o *ServerConfigCustomTranscodingResolutions) SetVar240p(v bool) {
	o.Var240p = &v
}

// GetVar360p returns the Var360p field value if set, zero value otherwise.
func (o *ServerConfigCustomTranscodingResolutions) GetVar360p() bool {
	if o == nil || utils.IsNil(o.Var360p) {
		var ret bool
		return ret
	}
	return *o.Var360p
}

// GetVar360pOk returns a tuple with the Var360p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigCustomTranscodingResolutions) GetVar360pOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.Var360p) {
		return nil, false
	}
	return o.Var360p, true
}

// HasVar360p returns a boolean if a field has been set.
func (o *ServerConfigCustomTranscodingResolutions) HasVar360p() bool {
	if o != nil && !utils.IsNil(o.Var360p) {
		return true
	}

	return false
}

// SetVar360p gets a reference to the given bool and assigns it to the Var360p field.
func (o *ServerConfigCustomTranscodingResolutions) SetVar360p(v bool) {
	o.Var360p = &v
}

// GetVar480p returns the Var480p field value if set, zero value otherwise.
func (o *ServerConfigCustomTranscodingResolutions) GetVar480p() bool {
	if o == nil || utils.IsNil(o.Var480p) {
		var ret bool
		return ret
	}
	return *o.Var480p
}

// GetVar480pOk returns a tuple with the Var480p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigCustomTranscodingResolutions) GetVar480pOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.Var480p) {
		return nil, false
	}
	return o.Var480p, true
}

// HasVar480p returns a boolean if a field has been set.
func (o *ServerConfigCustomTranscodingResolutions) HasVar480p() bool {
	if o != nil && !utils.IsNil(o.Var480p) {
		return true
	}

	return false
}

// SetVar480p gets a reference to the given bool and assigns it to the Var480p field.
func (o *ServerConfigCustomTranscodingResolutions) SetVar480p(v bool) {
	o.Var480p = &v
}

// GetVar720p returns the Var720p field value if set, zero value otherwise.
func (o *ServerConfigCustomTranscodingResolutions) GetVar720p() bool {
	if o == nil || utils.IsNil(o.Var720p) {
		var ret bool
		return ret
	}
	return *o.Var720p
}

// GetVar720pOk returns a tuple with the Var720p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigCustomTranscodingResolutions) GetVar720pOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.Var720p) {
		return nil, false
	}
	return o.Var720p, true
}

// HasVar720p returns a boolean if a field has been set.
func (o *ServerConfigCustomTranscodingResolutions) HasVar720p() bool {
	if o != nil && !utils.IsNil(o.Var720p) {
		return true
	}

	return false
}

// SetVar720p gets a reference to the given bool and assigns it to the Var720p field.
func (o *ServerConfigCustomTranscodingResolutions) SetVar720p(v bool) {
	o.Var720p = &v
}

// GetVar1080p returns the Var1080p field value if set, zero value otherwise.
func (o *ServerConfigCustomTranscodingResolutions) GetVar1080p() bool {
	if o == nil || utils.IsNil(o.Var1080p) {
		var ret bool
		return ret
	}
	return *o.Var1080p
}

// GetVar1080pOk returns a tuple with the Var1080p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigCustomTranscodingResolutions) GetVar1080pOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.Var1080p) {
		return nil, false
	}
	return o.Var1080p, true
}

// HasVar1080p returns a boolean if a field has been set.
func (o *ServerConfigCustomTranscodingResolutions) HasVar1080p() bool {
	if o != nil && !utils.IsNil(o.Var1080p) {
		return true
	}

	return false
}

// SetVar1080p gets a reference to the given bool and assigns it to the Var1080p field.
func (o *ServerConfigCustomTranscodingResolutions) SetVar1080p(v bool) {
	o.Var1080p = &v
}

// GetVar1440p returns the Var1440p field value if set, zero value otherwise.
func (o *ServerConfigCustomTranscodingResolutions) GetVar1440p() bool {
	if o == nil || utils.IsNil(o.Var1440p) {
		var ret bool
		return ret
	}
	return *o.Var1440p
}

// GetVar1440pOk returns a tuple with the Var1440p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigCustomTranscodingResolutions) GetVar1440pOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.Var1440p) {
		return nil, false
	}
	return o.Var1440p, true
}

// HasVar1440p returns a boolean if a field has been set.
func (o *ServerConfigCustomTranscodingResolutions) HasVar1440p() bool {
	if o != nil && !utils.IsNil(o.Var1440p) {
		return true
	}

	return false
}

// SetVar1440p gets a reference to the given bool and assigns it to the Var1440p field.
func (o *ServerConfigCustomTranscodingResolutions) SetVar1440p(v bool) {
	o.Var1440p = &v
}

// GetVar2160p returns the Var2160p field value if set, zero value otherwise.
func (o *ServerConfigCustomTranscodingResolutions) GetVar2160p() bool {
	if o == nil || utils.IsNil(o.Var2160p) {
		var ret bool
		return ret
	}
	return *o.Var2160p
}

// GetVar2160pOk returns a tuple with the Var2160p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigCustomTranscodingResolutions) GetVar2160pOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.Var2160p) {
		return nil, false
	}
	return o.Var2160p, true
}

// HasVar2160p returns a boolean if a field has been set.
func (o *ServerConfigCustomTranscodingResolutions) HasVar2160p() bool {
	if o != nil && !utils.IsNil(o.Var2160p) {
		return true
	}

	return false
}

// SetVar2160p gets a reference to the given bool and assigns it to the Var2160p field.
func (o *ServerConfigCustomTranscodingResolutions) SetVar2160p(v bool) {
	o.Var2160p = &v
}

func (o ServerConfigCustomTranscodingResolutions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerConfigCustomTranscodingResolutions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Var0p) {
		toSerialize["0p"] = o.Var0p
	}
	if !utils.IsNil(o.Var144p) {
		toSerialize["144p"] = o.Var144p
	}
	if !utils.IsNil(o.Var240p) {
		toSerialize["240p"] = o.Var240p
	}
	if !utils.IsNil(o.Var360p) {
		toSerialize["360p"] = o.Var360p
	}
	if !utils.IsNil(o.Var480p) {
		toSerialize["480p"] = o.Var480p
	}
	if !utils.IsNil(o.Var720p) {
		toSerialize["720p"] = o.Var720p
	}
	if !utils.IsNil(o.Var1080p) {
		toSerialize["1080p"] = o.Var1080p
	}
	if !utils.IsNil(o.Var1440p) {
		toSerialize["1440p"] = o.Var1440p
	}
	if !utils.IsNil(o.Var2160p) {
		toSerialize["2160p"] = o.Var2160p
	}
	return toSerialize, nil
}

type NullableServerConfigCustomTranscodingResolutions struct {
	value *ServerConfigCustomTranscodingResolutions
	isSet bool
}

func (v NullableServerConfigCustomTranscodingResolutions) Get() *ServerConfigCustomTranscodingResolutions {
	return v.value
}

func (v *NullableServerConfigCustomTranscodingResolutions) Set(val *ServerConfigCustomTranscodingResolutions) {
	v.value = val
	v.isSet = true
}

func (v NullableServerConfigCustomTranscodingResolutions) IsSet() bool {
	return v.isSet
}

func (v *NullableServerConfigCustomTranscodingResolutions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerConfigCustomTranscodingResolutions(val *ServerConfigCustomTranscodingResolutions) *NullableServerConfigCustomTranscodingResolutions {
	return &NullableServerConfigCustomTranscodingResolutions{value: val, isSet: true}
}

func (v NullableServerConfigCustomTranscodingResolutions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerConfigCustomTranscodingResolutions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
