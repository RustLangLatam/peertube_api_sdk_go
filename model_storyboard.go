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

// checks if the Storyboard type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Storyboard{}

// Storyboard struct for Storyboard
type Storyboard struct {
	StoryboardPath *string `json:"storyboardPath,omitempty"`
	TotalHeight    *int32  `json:"totalHeight,omitempty"`
	TotalWidth     *int32  `json:"totalWidth,omitempty"`
	SpriteHeight   *int32  `json:"spriteHeight,omitempty"`
	SpriteWidth    *int32  `json:"spriteWidth,omitempty"`
	SpriteDuration *int32  `json:"spriteDuration,omitempty"`
}

// NewStoryboard instantiates a new Storyboard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoryboard() *Storyboard {
	this := Storyboard{}
	return &this
}

// NewStoryboardWithDefaults instantiates a new Storyboard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoryboardWithDefaults() *Storyboard {
	this := Storyboard{}
	return &this
}

// GetStoryboardPath returns the StoryboardPath field value if set, zero value otherwise.
func (o *Storyboard) GetStoryboardPath() string {
	if o == nil || IsNil(o.StoryboardPath) {
		var ret string
		return ret
	}
	return *o.StoryboardPath
}

// GetStoryboardPathOk returns a tuple with the StoryboardPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Storyboard) GetStoryboardPathOk() (*string, bool) {
	if o == nil || IsNil(o.StoryboardPath) {
		return nil, false
	}
	return o.StoryboardPath, true
}

// HasStoryboardPath returns a boolean if a field has been set.
func (o *Storyboard) HasStoryboardPath() bool {
	if o != nil && !IsNil(o.StoryboardPath) {
		return true
	}

	return false
}

// SetStoryboardPath gets a reference to the given string and assigns it to the StoryboardPath field.
func (o *Storyboard) SetStoryboardPath(v string) {
	o.StoryboardPath = &v
}

// GetTotalHeight returns the TotalHeight field value if set, zero value otherwise.
func (o *Storyboard) GetTotalHeight() int32 {
	if o == nil || IsNil(o.TotalHeight) {
		var ret int32
		return ret
	}
	return *o.TotalHeight
}

// GetTotalHeightOk returns a tuple with the TotalHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Storyboard) GetTotalHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalHeight) {
		return nil, false
	}
	return o.TotalHeight, true
}

// HasTotalHeight returns a boolean if a field has been set.
func (o *Storyboard) HasTotalHeight() bool {
	if o != nil && !IsNil(o.TotalHeight) {
		return true
	}

	return false
}

// SetTotalHeight gets a reference to the given int32 and assigns it to the TotalHeight field.
func (o *Storyboard) SetTotalHeight(v int32) {
	o.TotalHeight = &v
}

// GetTotalWidth returns the TotalWidth field value if set, zero value otherwise.
func (o *Storyboard) GetTotalWidth() int32 {
	if o == nil || IsNil(o.TotalWidth) {
		var ret int32
		return ret
	}
	return *o.TotalWidth
}

// GetTotalWidthOk returns a tuple with the TotalWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Storyboard) GetTotalWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalWidth) {
		return nil, false
	}
	return o.TotalWidth, true
}

// HasTotalWidth returns a boolean if a field has been set.
func (o *Storyboard) HasTotalWidth() bool {
	if o != nil && !IsNil(o.TotalWidth) {
		return true
	}

	return false
}

// SetTotalWidth gets a reference to the given int32 and assigns it to the TotalWidth field.
func (o *Storyboard) SetTotalWidth(v int32) {
	o.TotalWidth = &v
}

// GetSpriteHeight returns the SpriteHeight field value if set, zero value otherwise.
func (o *Storyboard) GetSpriteHeight() int32 {
	if o == nil || IsNil(o.SpriteHeight) {
		var ret int32
		return ret
	}
	return *o.SpriteHeight
}

// GetSpriteHeightOk returns a tuple with the SpriteHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Storyboard) GetSpriteHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.SpriteHeight) {
		return nil, false
	}
	return o.SpriteHeight, true
}

// HasSpriteHeight returns a boolean if a field has been set.
func (o *Storyboard) HasSpriteHeight() bool {
	if o != nil && !IsNil(o.SpriteHeight) {
		return true
	}

	return false
}

// SetSpriteHeight gets a reference to the given int32 and assigns it to the SpriteHeight field.
func (o *Storyboard) SetSpriteHeight(v int32) {
	o.SpriteHeight = &v
}

// GetSpriteWidth returns the SpriteWidth field value if set, zero value otherwise.
func (o *Storyboard) GetSpriteWidth() int32 {
	if o == nil || IsNil(o.SpriteWidth) {
		var ret int32
		return ret
	}
	return *o.SpriteWidth
}

// GetSpriteWidthOk returns a tuple with the SpriteWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Storyboard) GetSpriteWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.SpriteWidth) {
		return nil, false
	}
	return o.SpriteWidth, true
}

// HasSpriteWidth returns a boolean if a field has been set.
func (o *Storyboard) HasSpriteWidth() bool {
	if o != nil && !IsNil(o.SpriteWidth) {
		return true
	}

	return false
}

// SetSpriteWidth gets a reference to the given int32 and assigns it to the SpriteWidth field.
func (o *Storyboard) SetSpriteWidth(v int32) {
	o.SpriteWidth = &v
}

// GetSpriteDuration returns the SpriteDuration field value if set, zero value otherwise.
func (o *Storyboard) GetSpriteDuration() int32 {
	if o == nil || IsNil(o.SpriteDuration) {
		var ret int32
		return ret
	}
	return *o.SpriteDuration
}

// GetSpriteDurationOk returns a tuple with the SpriteDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Storyboard) GetSpriteDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.SpriteDuration) {
		return nil, false
	}
	return o.SpriteDuration, true
}

// HasSpriteDuration returns a boolean if a field has been set.
func (o *Storyboard) HasSpriteDuration() bool {
	if o != nil && !IsNil(o.SpriteDuration) {
		return true
	}

	return false
}

// SetSpriteDuration gets a reference to the given int32 and assigns it to the SpriteDuration field.
func (o *Storyboard) SetSpriteDuration(v int32) {
	o.SpriteDuration = &v
}

func (o Storyboard) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Storyboard) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StoryboardPath) {
		toSerialize["storyboardPath"] = o.StoryboardPath
	}
	if !IsNil(o.TotalHeight) {
		toSerialize["totalHeight"] = o.TotalHeight
	}
	if !IsNil(o.TotalWidth) {
		toSerialize["totalWidth"] = o.TotalWidth
	}
	if !IsNil(o.SpriteHeight) {
		toSerialize["spriteHeight"] = o.SpriteHeight
	}
	if !IsNil(o.SpriteWidth) {
		toSerialize["spriteWidth"] = o.SpriteWidth
	}
	if !IsNil(o.SpriteDuration) {
		toSerialize["spriteDuration"] = o.SpriteDuration
	}
	return toSerialize, nil
}

type NullableStoryboard struct {
	value *Storyboard
	isSet bool
}

func (v NullableStoryboard) Get() *Storyboard {
	return v.value
}

func (v *NullableStoryboard) Set(val *Storyboard) {
	v.value = val
	v.isSet = true
}

func (v NullableStoryboard) IsSet() bool {
	return v.isSet
}

func (v *NullableStoryboard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoryboard(val *Storyboard) *NullableStoryboard {
	return &NullableStoryboard{value: val, isSet: true}
}

func (v NullableStoryboard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoryboard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
