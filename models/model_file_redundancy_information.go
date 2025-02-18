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

// checks if the FileRedundancyInformation type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &FileRedundancyInformation{}

// FileRedundancyInformation struct for FileRedundancyInformation
type FileRedundancyInformation struct {
	Id        *int32     `json:"id,omitempty"`
	FileUrl   *string    `json:"fileUrl,omitempty"`
	Strategy  *string    `json:"strategy,omitempty"`
	Size      *int32     `json:"size,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	ExpiresOn *time.Time `json:"expiresOn,omitempty"`
}

// NewFileRedundancyInformation instantiates a new FileRedundancyInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileRedundancyInformation() *FileRedundancyInformation {
	this := FileRedundancyInformation{}
	return &this
}

// NewFileRedundancyInformationWithDefaults instantiates a new FileRedundancyInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileRedundancyInformationWithDefaults() *FileRedundancyInformation {
	this := FileRedundancyInformation{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FileRedundancyInformation) GetId() int32 {
	if o == nil || utils.IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileRedundancyInformation) GetIdOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FileRedundancyInformation) HasId() bool {
	if o != nil && !utils.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *FileRedundancyInformation) SetId(v int32) {
	o.Id = &v
}

// GetFileUrl returns the FileUrl field value if set, zero value otherwise.
func (o *FileRedundancyInformation) GetFileUrl() string {
	if o == nil || utils.IsNil(o.FileUrl) {
		var ret string
		return ret
	}
	return *o.FileUrl
}

// GetFileUrlOk returns a tuple with the FileUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileRedundancyInformation) GetFileUrlOk() (*string, bool) {
	if o == nil || utils.IsNil(o.FileUrl) {
		return nil, false
	}
	return o.FileUrl, true
}

// HasFileUrl returns a boolean if a field has been set.
func (o *FileRedundancyInformation) HasFileUrl() bool {
	if o != nil && !utils.IsNil(o.FileUrl) {
		return true
	}

	return false
}

// SetFileUrl gets a reference to the given string and assigns it to the FileUrl field.
func (o *FileRedundancyInformation) SetFileUrl(v string) {
	o.FileUrl = &v
}

// GetStrategy returns the Strategy field value if set, zero value otherwise.
func (o *FileRedundancyInformation) GetStrategy() string {
	if o == nil || utils.IsNil(o.Strategy) {
		var ret string
		return ret
	}
	return *o.Strategy
}

// GetStrategyOk returns a tuple with the Strategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileRedundancyInformation) GetStrategyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Strategy) {
		return nil, false
	}
	return o.Strategy, true
}

// HasStrategy returns a boolean if a field has been set.
func (o *FileRedundancyInformation) HasStrategy() bool {
	if o != nil && !utils.IsNil(o.Strategy) {
		return true
	}

	return false
}

// SetStrategy gets a reference to the given string and assigns it to the Strategy field.
func (o *FileRedundancyInformation) SetStrategy(v string) {
	o.Strategy = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *FileRedundancyInformation) GetSize() int32 {
	if o == nil || utils.IsNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileRedundancyInformation) GetSizeOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *FileRedundancyInformation) HasSize() bool {
	if o != nil && !utils.IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *FileRedundancyInformation) SetSize(v int32) {
	o.Size = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *FileRedundancyInformation) GetCreatedAt() time.Time {
	if o == nil || utils.IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileRedundancyInformation) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *FileRedundancyInformation) HasCreatedAt() bool {
	if o != nil && !utils.IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *FileRedundancyInformation) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *FileRedundancyInformation) GetUpdatedAt() time.Time {
	if o == nil || utils.IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileRedundancyInformation) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *FileRedundancyInformation) HasUpdatedAt() bool {
	if o != nil && !utils.IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *FileRedundancyInformation) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetExpiresOn returns the ExpiresOn field value if set, zero value otherwise.
func (o *FileRedundancyInformation) GetExpiresOn() time.Time {
	if o == nil || utils.IsNil(o.ExpiresOn) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresOn
}

// GetExpiresOnOk returns a tuple with the ExpiresOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileRedundancyInformation) GetExpiresOnOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.ExpiresOn) {
		return nil, false
	}
	return o.ExpiresOn, true
}

// HasExpiresOn returns a boolean if a field has been set.
func (o *FileRedundancyInformation) HasExpiresOn() bool {
	if o != nil && !utils.IsNil(o.ExpiresOn) {
		return true
	}

	return false
}

// SetExpiresOn gets a reference to the given time.Time and assigns it to the ExpiresOn field.
func (o *FileRedundancyInformation) SetExpiresOn(v time.Time) {
	o.ExpiresOn = &v
}

func (o FileRedundancyInformation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileRedundancyInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !utils.IsNil(o.FileUrl) {
		toSerialize["fileUrl"] = o.FileUrl
	}
	if !utils.IsNil(o.Strategy) {
		toSerialize["strategy"] = o.Strategy
	}
	if !utils.IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !utils.IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !utils.IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !utils.IsNil(o.ExpiresOn) {
		toSerialize["expiresOn"] = o.ExpiresOn
	}
	return toSerialize, nil
}

type NullableFileRedundancyInformation struct {
	value *FileRedundancyInformation
	isSet bool
}

func (v NullableFileRedundancyInformation) Get() *FileRedundancyInformation {
	return v.value
}

func (v *NullableFileRedundancyInformation) Set(val *FileRedundancyInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableFileRedundancyInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableFileRedundancyInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileRedundancyInformation(val *FileRedundancyInformation) *NullableFileRedundancyInformation {
	return &NullableFileRedundancyInformation{value: val, isSet: true}
}

func (v NullableFileRedundancyInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileRedundancyInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
