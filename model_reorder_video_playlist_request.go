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

// checks if the ReorderVideoPlaylistRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReorderVideoPlaylistRequest{}

// ReorderVideoPlaylistRequest struct for ReorderVideoPlaylistRequest
type ReorderVideoPlaylistRequest struct {
	// Start position of the element to reorder
	StartPosition int32 `json:"startPosition"`
	// New position for the block to reorder, to add the block before the first element
	InsertAfterPosition int32 `json:"insertAfterPosition"`
	// How many element from `startPosition` to reorder
	ReorderLength *int32 `json:"reorderLength,omitempty"`
}

type _ReorderVideoPlaylistRequest ReorderVideoPlaylistRequest

// NewReorderVideoPlaylistRequest instantiates a new ReorderVideoPlaylistRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReorderVideoPlaylistRequest(startPosition int32, insertAfterPosition int32) *ReorderVideoPlaylistRequest {
	this := ReorderVideoPlaylistRequest{}
	this.StartPosition = startPosition
	this.InsertAfterPosition = insertAfterPosition
	return &this
}

// NewReorderVideoPlaylistRequestWithDefaults instantiates a new ReorderVideoPlaylistRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReorderVideoPlaylistRequestWithDefaults() *ReorderVideoPlaylistRequest {
	this := ReorderVideoPlaylistRequest{}
	return &this
}

// GetStartPosition returns the StartPosition field value
func (o *ReorderVideoPlaylistRequest) GetStartPosition() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StartPosition
}

// GetStartPositionOk returns a tuple with the StartPosition field value
// and a boolean to check if the value has been set.
func (o *ReorderVideoPlaylistRequest) GetStartPositionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartPosition, true
}

// SetStartPosition sets field value
func (o *ReorderVideoPlaylistRequest) SetStartPosition(v int32) {
	o.StartPosition = v
}

// GetInsertAfterPosition returns the InsertAfterPosition field value
func (o *ReorderVideoPlaylistRequest) GetInsertAfterPosition() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.InsertAfterPosition
}

// GetInsertAfterPositionOk returns a tuple with the InsertAfterPosition field value
// and a boolean to check if the value has been set.
func (o *ReorderVideoPlaylistRequest) GetInsertAfterPositionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InsertAfterPosition, true
}

// SetInsertAfterPosition sets field value
func (o *ReorderVideoPlaylistRequest) SetInsertAfterPosition(v int32) {
	o.InsertAfterPosition = v
}

// GetReorderLength returns the ReorderLength field value if set, zero value otherwise.
func (o *ReorderVideoPlaylistRequest) GetReorderLength() int32 {
	if o == nil || IsNil(o.ReorderLength) {
		var ret int32
		return ret
	}
	return *o.ReorderLength
}

// GetReorderLengthOk returns a tuple with the ReorderLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReorderVideoPlaylistRequest) GetReorderLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.ReorderLength) {
		return nil, false
	}
	return o.ReorderLength, true
}

// HasReorderLength returns a boolean if a field has been set.
func (o *ReorderVideoPlaylistRequest) HasReorderLength() bool {
	if o != nil && !IsNil(o.ReorderLength) {
		return true
	}

	return false
}

// SetReorderLength gets a reference to the given int32 and assigns it to the ReorderLength field.
func (o *ReorderVideoPlaylistRequest) SetReorderLength(v int32) {
	o.ReorderLength = &v
}

func (o ReorderVideoPlaylistRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReorderVideoPlaylistRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["startPosition"] = o.StartPosition
	toSerialize["insertAfterPosition"] = o.InsertAfterPosition
	if !IsNil(o.ReorderLength) {
		toSerialize["reorderLength"] = o.ReorderLength
	}
	return toSerialize, nil
}

func (o *ReorderVideoPlaylistRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"startPosition",
		"insertAfterPosition",
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

	varReorderVideoPlaylistRequest := _ReorderVideoPlaylistRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReorderVideoPlaylistRequest)

	if err != nil {
		return err
	}

	*o = ReorderVideoPlaylistRequest(varReorderVideoPlaylistRequest)

	return err
}

type NullableReorderVideoPlaylistRequest struct {
	value *ReorderVideoPlaylistRequest
	isSet bool
}

func (v NullableReorderVideoPlaylistRequest) Get() *ReorderVideoPlaylistRequest {
	return v.value
}

func (v *NullableReorderVideoPlaylistRequest) Set(val *ReorderVideoPlaylistRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReorderVideoPlaylistRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReorderVideoPlaylistRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReorderVideoPlaylistRequest(val *ReorderVideoPlaylistRequest) *NullableReorderVideoPlaylistRequest {
	return &NullableReorderVideoPlaylistRequest{value: val, isSet: true}
}

func (v NullableReorderVideoPlaylistRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReorderVideoPlaylistRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
