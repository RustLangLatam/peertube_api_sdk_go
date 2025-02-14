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

// checks if the VODWebVideoTranscoding1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VODWebVideoTranscoding1{}

// VODWebVideoTranscoding1 struct for VODWebVideoTranscoding1
type VODWebVideoTranscoding1 struct {
	Input  *VODWebVideoTranscoding1Input  `json:"input,omitempty"`
	Output *VODWebVideoTranscoding1Output `json:"output,omitempty"`
}

// NewVODWebVideoTranscoding1 instantiates a new VODWebVideoTranscoding1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVODWebVideoTranscoding1() *VODWebVideoTranscoding1 {
	this := VODWebVideoTranscoding1{}
	return &this
}

// NewVODWebVideoTranscoding1WithDefaults instantiates a new VODWebVideoTranscoding1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVODWebVideoTranscoding1WithDefaults() *VODWebVideoTranscoding1 {
	this := VODWebVideoTranscoding1{}
	return &this
}

// GetInput returns the Input field value if set, zero value otherwise.
func (o *VODWebVideoTranscoding1) GetInput() VODWebVideoTranscoding1Input {
	if o == nil || IsNil(o.Input) {
		var ret VODWebVideoTranscoding1Input
		return ret
	}
	return *o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VODWebVideoTranscoding1) GetInputOk() (*VODWebVideoTranscoding1Input, bool) {
	if o == nil || IsNil(o.Input) {
		return nil, false
	}
	return o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *VODWebVideoTranscoding1) HasInput() bool {
	if o != nil && !IsNil(o.Input) {
		return true
	}

	return false
}

// SetInput gets a reference to the given VODWebVideoTranscoding1Input and assigns it to the Input field.
func (o *VODWebVideoTranscoding1) SetInput(v VODWebVideoTranscoding1Input) {
	o.Input = &v
}

// GetOutput returns the Output field value if set, zero value otherwise.
func (o *VODWebVideoTranscoding1) GetOutput() VODWebVideoTranscoding1Output {
	if o == nil || IsNil(o.Output) {
		var ret VODWebVideoTranscoding1Output
		return ret
	}
	return *o.Output
}

// GetOutputOk returns a tuple with the Output field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VODWebVideoTranscoding1) GetOutputOk() (*VODWebVideoTranscoding1Output, bool) {
	if o == nil || IsNil(o.Output) {
		return nil, false
	}
	return o.Output, true
}

// HasOutput returns a boolean if a field has been set.
func (o *VODWebVideoTranscoding1) HasOutput() bool {
	if o != nil && !IsNil(o.Output) {
		return true
	}

	return false
}

// SetOutput gets a reference to the given VODWebVideoTranscoding1Output and assigns it to the Output field.
func (o *VODWebVideoTranscoding1) SetOutput(v VODWebVideoTranscoding1Output) {
	o.Output = &v
}

func (o VODWebVideoTranscoding1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VODWebVideoTranscoding1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Input) {
		toSerialize["input"] = o.Input
	}
	if !IsNil(o.Output) {
		toSerialize["output"] = o.Output
	}
	return toSerialize, nil
}

type NullableVODWebVideoTranscoding1 struct {
	value *VODWebVideoTranscoding1
	isSet bool
}

func (v NullableVODWebVideoTranscoding1) Get() *VODWebVideoTranscoding1 {
	return v.value
}

func (v *NullableVODWebVideoTranscoding1) Set(val *VODWebVideoTranscoding1) {
	v.value = val
	v.isSet = true
}

func (v NullableVODWebVideoTranscoding1) IsSet() bool {
	return v.isSet
}

func (v *NullableVODWebVideoTranscoding1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVODWebVideoTranscoding1(val *VODWebVideoTranscoding1) *NullableVODWebVideoTranscoding1 {
	return &NullableVODWebVideoTranscoding1{value: val, isSet: true}
}

func (v NullableVODWebVideoTranscoding1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVODWebVideoTranscoding1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
