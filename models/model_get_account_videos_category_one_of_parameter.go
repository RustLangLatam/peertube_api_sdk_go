/*
PeerTube

The PeerTube API is built on HTTP(S) and is RESTful. You can use your favorite HTTP/REST library for your programming language to use PeerTube. The spec API is fully compatible with [openapi-generator](https://github.com/OpenAPITools/openapi-generator/wiki/API-client-generator-HOWTO) which generates a client SDK in the language of your choice - we generate some client SDKs automatically:  - [Python](https://framagit.org/framasoft/peertube/clients/python) - [Go](https://framagit.org/framasoft/peertube/clients/go) - [Kotlin](https://framagit.org/framasoft/peertube/clients/kotlin)  See the [REST API quick start](https://docs.joinpeertube.org/api/rest-getting-started) for a few examples of using the PeerTube API.  # Authentication  When you sign up for an account on a PeerTube instance, you are given the possibility to generate sessions on it, and authenticate there using an access token. Only __one access token can currently be used at a time__.  ## Roles  Accounts are given permissions based on their role. There are three roles on PeerTube: Administrator, Moderator, and User. See the [roles guide](https://docs.joinpeertube.org/admin/managing-users#roles) for a detail of their permissions.  # Errors  The API uses standard HTTP status codes to indicate the success or failure of the API call, completed by a [RFC7807-compliant](https://tools.ietf.org/html/rfc7807) response body.  ``` HTTP 1.1 404 Not Found Content-Type: application/problem+json; charset=utf-8  {   \"detail\": \"Video not found\",   \"docs\": \"https://docs.joinpeertube.org/api-rest-reference.html#operation/getVideo\",   \"status\": 404,   \"title\": \"Not Found\",   \"type\": \"about:blank\" } ```  We provide error `type` (following RFC7807) and `code` (internal PeerTube code) values for [a growing number of cases](https://github.com/Chocobozzz/PeerTube/blob/develop/packages/models/src/server/server-error-code.enum.ts), but it is still optional. Types are used to disambiguate errors that bear the same status code and are non-obvious:  ``` HTTP 1.1 403 Forbidden Content-Type: application/problem+json; charset=utf-8  {   \"detail\": \"Cannot get this video regarding follow constraints\",   \"docs\": \"https://docs.joinpeertube.org/api-rest-reference.html#operation/getVideo\",   \"status\": 403,   \"title\": \"Forbidden\",   \"type\": \"https://docs.joinpeertube.org/api-rest-reference.html#section/Errors/does_not_respect_follow_constraints\" } ```  Here a 403 error could otherwise mean that the video is private or blocklisted.  ### Validation errors  Each parameter is evaluated on its own against a set of rules before the route validator proceeds with potential testing involving parameter combinations. Errors coming from validation errors appear earlier and benefit from a more detailed error description:  ``` HTTP 1.1 400 Bad Request Content-Type: application/problem+json; charset=utf-8  {   \"detail\": \"Incorrect request parameters: id\",   \"docs\": \"https://docs.joinpeertube.org/api-rest-reference.html#operation/getVideo\",   \"instance\": \"/api/v1/videos/9c9de5e8-0a1e-484a-b099-e80766180\",   \"invalid-params\": {     \"id\": {       \"location\": \"params\",       \"msg\": \"Invalid value\",       \"param\": \"id\",       \"value\": \"9c9de5e8-0a1e-484a-b099-e80766180\"     }   },   \"status\": 400,   \"title\": \"Bad Request\",   \"type\": \"about:blank\" } ```  Where `id` is the name of the field concerned by the error, within the route definition. `invalid-params.<field>.location` can be either 'params', 'body', 'header', 'query' or 'cookies', and `invalid-params.<field>.value` reports the value that didn't pass validation whose `invalid-params.<field>.msg` is about.  ### Deprecated error fields  Some fields could be included with previous versions. They are still included but their use is deprecated: - `error`: superseded by `detail`  # Rate limits  We are rate-limiting all endpoints of PeerTube's API. Custom values can be set by administrators:  | Endpoint (prefix: `/api/v1`) | Calls         | Time frame   | |------------------------------|---------------|--------------| | `/_*`                         | 50            | 10 seconds   | | `POST /users/token`          | 15            | 5 minutes    | | `POST /users/register`       | 2<sup>*</sup> | 5 minutes    | | `POST /users/ask-send-verify-email` | 3      | 5 minutes    |  Depending on the endpoint, <sup>*</sup>failed requests are not taken into account. A service limit is announced by a `429 Too Many Requests` status code.  You can get details about the current state of your rate limit by reading the following headers:  | Header                  | Description                                                | |-------------------------|------------------------------------------------------------| | `X-RateLimit-Limit`     | Number of max requests allowed in the current time period  | | `X-RateLimit-Remaining` | Number of remaining requests in the current time period    | | `X-RateLimit-Reset`     | Timestamp of end of current time period as UNIX timestamp  | | `Retry-After`           | Seconds to delay after the first `429` is received         |  # CORS  This API features [Cross-Origin Resource Sharing (CORS)](https://fetch.spec.whatwg.org/), allowing cross-domain communication from the browser for some routes:  | Endpoint                    | |------------------------- ---| | `/api/_*`                    | | `/download/_*`               | | `/lazy-static/_*`            | | `/.well-known/webfinger`    |  In addition, all routes serving ActivityPub are CORS-enabled for all origins.

API version: 7.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"fmt"
)

// GetAccountVideosCategoryOneOfParameter - struct for GetAccountVideosCategoryOneOfParameter
type GetAccountVideosCategoryOneOfParameter struct {
	ArrayOfInt32 *[]int32
	Int32        *int32
}

// []int32AsGetAccountVideosCategoryOneOfParameter is a convenience function that returns []int32 wrapped in GetAccountVideosCategoryOneOfParameter
func ArrayOfInt32AsGetAccountVideosCategoryOneOfParameter(v *[]int32) GetAccountVideosCategoryOneOfParameter {
	return GetAccountVideosCategoryOneOfParameter{
		ArrayOfInt32: v,
	}
}

// int32AsGetAccountVideosCategoryOneOfParameter is a convenience function that returns int32 wrapped in GetAccountVideosCategoryOneOfParameter
func Int32AsGetAccountVideosCategoryOneOfParameter(v *int32) GetAccountVideosCategoryOneOfParameter {
	return GetAccountVideosCategoryOneOfParameter{
		Int32: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetAccountVideosCategoryOneOfParameter) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ArrayOfInt32
	err = json.Unmarshal(data, &dst.ArrayOfInt32)
	if err == nil {
		jsonArrayOfInt32, _ := json.Marshal(dst.ArrayOfInt32)
		if string(jsonArrayOfInt32) == "{}" { // empty struct
			dst.ArrayOfInt32 = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfInt32 = nil
	}

	// try to unmarshal data into Int32
	err = json.Unmarshal(data, &dst.Int32)
	if err == nil {
		jsonInt32, _ := json.Marshal(dst.Int32)
		if string(jsonInt32) == "{}" { // empty struct
			dst.Int32 = nil
		} else {
			match++
		}
	} else {
		dst.Int32 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ArrayOfInt32 = nil
		dst.Int32 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetAccountVideosCategoryOneOfParameter)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetAccountVideosCategoryOneOfParameter)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetAccountVideosCategoryOneOfParameter) MarshalJSON() ([]byte, error) {
	if src.ArrayOfInt32 != nil {
		return json.Marshal(&src.ArrayOfInt32)
	}

	if src.Int32 != nil {
		return json.Marshal(&src.Int32)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetAccountVideosCategoryOneOfParameter) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ArrayOfInt32 != nil {
		return obj.ArrayOfInt32
	}

	if obj.Int32 != nil {
		return obj.Int32
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj GetAccountVideosCategoryOneOfParameter) GetActualInstanceValue() interface{} {
	if obj.ArrayOfInt32 != nil {
		return *obj.ArrayOfInt32
	}

	if obj.Int32 != nil {
		return *obj.Int32
	}

	// all schemas are nil
	return nil
}

type NullableGetAccountVideosCategoryOneOfParameter struct {
	value *GetAccountVideosCategoryOneOfParameter
	isSet bool
}

func (v NullableGetAccountVideosCategoryOneOfParameter) Get() *GetAccountVideosCategoryOneOfParameter {
	return v.value
}

func (v *NullableGetAccountVideosCategoryOneOfParameter) Set(val *GetAccountVideosCategoryOneOfParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAccountVideosCategoryOneOfParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAccountVideosCategoryOneOfParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAccountVideosCategoryOneOfParameter(val *GetAccountVideosCategoryOneOfParameter) *NullableGetAccountVideosCategoryOneOfParameter {
	return &NullableGetAccountVideosCategoryOneOfParameter{value: val, isSet: true}
}

func (v NullableGetAccountVideosCategoryOneOfParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAccountVideosCategoryOneOfParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
