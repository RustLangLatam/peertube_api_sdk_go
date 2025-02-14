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

// checks if the ServerConfigInstanceSocial type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerConfigInstanceSocial{}

// ServerConfigInstanceSocial struct for ServerConfigInstanceSocial
type ServerConfigInstanceSocial struct {
	ExternalLink *string `json:"externalLink,omitempty"`
	MastodonLink *string `json:"mastodonLink,omitempty"`
	BlueskyLink  *string `json:"blueskyLink,omitempty"`
}

// NewServerConfigInstanceSocial instantiates a new ServerConfigInstanceSocial object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerConfigInstanceSocial() *ServerConfigInstanceSocial {
	this := ServerConfigInstanceSocial{}
	return &this
}

// NewServerConfigInstanceSocialWithDefaults instantiates a new ServerConfigInstanceSocial object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerConfigInstanceSocialWithDefaults() *ServerConfigInstanceSocial {
	this := ServerConfigInstanceSocial{}
	return &this
}

// GetExternalLink returns the ExternalLink field value if set, zero value otherwise.
func (o *ServerConfigInstanceSocial) GetExternalLink() string {
	if o == nil || IsNil(o.ExternalLink) {
		var ret string
		return ret
	}
	return *o.ExternalLink
}

// GetExternalLinkOk returns a tuple with the ExternalLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigInstanceSocial) GetExternalLinkOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalLink) {
		return nil, false
	}
	return o.ExternalLink, true
}

// HasExternalLink returns a boolean if a field has been set.
func (o *ServerConfigInstanceSocial) HasExternalLink() bool {
	if o != nil && !IsNil(o.ExternalLink) {
		return true
	}

	return false
}

// SetExternalLink gets a reference to the given string and assigns it to the ExternalLink field.
func (o *ServerConfigInstanceSocial) SetExternalLink(v string) {
	o.ExternalLink = &v
}

// GetMastodonLink returns the MastodonLink field value if set, zero value otherwise.
func (o *ServerConfigInstanceSocial) GetMastodonLink() string {
	if o == nil || IsNil(o.MastodonLink) {
		var ret string
		return ret
	}
	return *o.MastodonLink
}

// GetMastodonLinkOk returns a tuple with the MastodonLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigInstanceSocial) GetMastodonLinkOk() (*string, bool) {
	if o == nil || IsNil(o.MastodonLink) {
		return nil, false
	}
	return o.MastodonLink, true
}

// HasMastodonLink returns a boolean if a field has been set.
func (o *ServerConfigInstanceSocial) HasMastodonLink() bool {
	if o != nil && !IsNil(o.MastodonLink) {
		return true
	}

	return false
}

// SetMastodonLink gets a reference to the given string and assigns it to the MastodonLink field.
func (o *ServerConfigInstanceSocial) SetMastodonLink(v string) {
	o.MastodonLink = &v
}

// GetBlueskyLink returns the BlueskyLink field value if set, zero value otherwise.
func (o *ServerConfigInstanceSocial) GetBlueskyLink() string {
	if o == nil || IsNil(o.BlueskyLink) {
		var ret string
		return ret
	}
	return *o.BlueskyLink
}

// GetBlueskyLinkOk returns a tuple with the BlueskyLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigInstanceSocial) GetBlueskyLinkOk() (*string, bool) {
	if o == nil || IsNil(o.BlueskyLink) {
		return nil, false
	}
	return o.BlueskyLink, true
}

// HasBlueskyLink returns a boolean if a field has been set.
func (o *ServerConfigInstanceSocial) HasBlueskyLink() bool {
	if o != nil && !IsNil(o.BlueskyLink) {
		return true
	}

	return false
}

// SetBlueskyLink gets a reference to the given string and assigns it to the BlueskyLink field.
func (o *ServerConfigInstanceSocial) SetBlueskyLink(v string) {
	o.BlueskyLink = &v
}

func (o ServerConfigInstanceSocial) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerConfigInstanceSocial) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExternalLink) {
		toSerialize["externalLink"] = o.ExternalLink
	}
	if !IsNil(o.MastodonLink) {
		toSerialize["mastodonLink"] = o.MastodonLink
	}
	if !IsNil(o.BlueskyLink) {
		toSerialize["blueskyLink"] = o.BlueskyLink
	}
	return toSerialize, nil
}

type NullableServerConfigInstanceSocial struct {
	value *ServerConfigInstanceSocial
	isSet bool
}

func (v NullableServerConfigInstanceSocial) Get() *ServerConfigInstanceSocial {
	return v.value
}

func (v *NullableServerConfigInstanceSocial) Set(val *ServerConfigInstanceSocial) {
	v.value = val
	v.isSet = true
}

func (v NullableServerConfigInstanceSocial) IsSet() bool {
	return v.isSet
}

func (v *NullableServerConfigInstanceSocial) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerConfigInstanceSocial(val *ServerConfigInstanceSocial) *NullableServerConfigInstanceSocial {
	return &NullableServerConfigInstanceSocial{value: val, isSet: true}
}

func (v NullableServerConfigInstanceSocial) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerConfigInstanceSocial) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
