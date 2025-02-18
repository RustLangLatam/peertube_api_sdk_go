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

// checks if the ServerConfigCustom type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ServerConfigCustom{}

// ServerConfigCustom struct for ServerConfigCustom
type ServerConfigCustom struct {
	Instance      *ServerConfigCustomInstance    `json:"instance,omitempty"`
	Theme         *ServerConfigCustomTheme       `json:"theme,omitempty"`
	Services      *ServerConfigCustomServices    `json:"services,omitempty"`
	Cache         *ServerConfigCustomCache       `json:"cache,omitempty"`
	Signup        *ServerConfigCustomSignup      `json:"signup,omitempty"`
	Admin         *ServerConfigCustomAdmin       `json:"admin,omitempty"`
	ContactForm   *ServerConfigEmail             `json:"contactForm,omitempty"`
	User          *ServerConfigCustomUser        `json:"user,omitempty"`
	Transcoding   *ServerConfigCustomTranscoding `json:"transcoding,omitempty"`
	Import        *ServerConfigCustomImport      `json:"import,omitempty"`
	AutoBlacklist *ServerConfigAutoBlacklist     `json:"autoBlacklist,omitempty"`
	Followers     *ServerConfigCustomFollowers   `json:"followers,omitempty"`
}

// NewServerConfigCustom instantiates a new ServerConfigCustom object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerConfigCustom() *ServerConfigCustom {
	this := ServerConfigCustom{}
	return &this
}

// NewServerConfigCustomWithDefaults instantiates a new ServerConfigCustom object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerConfigCustomWithDefaults() *ServerConfigCustom {
	this := ServerConfigCustom{}
	return &this
}

// GetInstance returns the Instance field value if set, zero value otherwise.
func (o *ServerConfigCustom) GetInstance() ServerConfigCustomInstance {
	if o == nil || utils.IsNil(o.Instance) {
		var ret ServerConfigCustomInstance
		return ret
	}
	return *o.Instance
}

// GetInstanceOk returns a tuple with the Instance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigCustom) GetInstanceOk() (*ServerConfigCustomInstance, bool) {
	if o == nil || utils.IsNil(o.Instance) {
		return nil, false
	}
	return o.Instance, true
}

// HasInstance returns a boolean if a field has been set.
func (o *ServerConfigCustom) HasInstance() bool {
	if o != nil && !utils.IsNil(o.Instance) {
		return true
	}

	return false
}

// SetInstance gets a reference to the given ServerConfigCustomInstance and assigns it to the Instance field.
func (o *ServerConfigCustom) SetInstance(v ServerConfigCustomInstance) {
	o.Instance = &v
}

// GetTheme returns the Theme field value if set, zero value otherwise.
func (o *ServerConfigCustom) GetTheme() ServerConfigCustomTheme {
	if o == nil || utils.IsNil(o.Theme) {
		var ret ServerConfigCustomTheme
		return ret
	}
	return *o.Theme
}

// GetThemeOk returns a tuple with the Theme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigCustom) GetThemeOk() (*ServerConfigCustomTheme, bool) {
	if o == nil || utils.IsNil(o.Theme) {
		return nil, false
	}
	return o.Theme, true
}

// HasTheme returns a boolean if a field has been set.
func (o *ServerConfigCustom) HasTheme() bool {
	if o != nil && !utils.IsNil(o.Theme) {
		return true
	}

	return false
}

// SetTheme gets a reference to the given ServerConfigCustomTheme and assigns it to the Theme field.
func (o *ServerConfigCustom) SetTheme(v ServerConfigCustomTheme) {
	o.Theme = &v
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *ServerConfigCustom) GetServices() ServerConfigCustomServices {
	if o == nil || utils.IsNil(o.Services) {
		var ret ServerConfigCustomServices
		return ret
	}
	return *o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigCustom) GetServicesOk() (*ServerConfigCustomServices, bool) {
	if o == nil || utils.IsNil(o.Services) {
		return nil, false
	}
	return o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *ServerConfigCustom) HasServices() bool {
	if o != nil && !utils.IsNil(o.Services) {
		return true
	}

	return false
}

// SetServices gets a reference to the given ServerConfigCustomServices and assigns it to the Services field.
func (o *ServerConfigCustom) SetServices(v ServerConfigCustomServices) {
	o.Services = &v
}

// GetCache returns the Cache field value if set, zero value otherwise.
func (o *ServerConfigCustom) GetCache() ServerConfigCustomCache {
	if o == nil || utils.IsNil(o.Cache) {
		var ret ServerConfigCustomCache
		return ret
	}
	return *o.Cache
}

// GetCacheOk returns a tuple with the Cache field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigCustom) GetCacheOk() (*ServerConfigCustomCache, bool) {
	if o == nil || utils.IsNil(o.Cache) {
		return nil, false
	}
	return o.Cache, true
}

// HasCache returns a boolean if a field has been set.
func (o *ServerConfigCustom) HasCache() bool {
	if o != nil && !utils.IsNil(o.Cache) {
		return true
	}

	return false
}

// SetCache gets a reference to the given ServerConfigCustomCache and assigns it to the Cache field.
func (o *ServerConfigCustom) SetCache(v ServerConfigCustomCache) {
	o.Cache = &v
}

// GetSignup returns the Signup field value if set, zero value otherwise.
func (o *ServerConfigCustom) GetSignup() ServerConfigCustomSignup {
	if o == nil || utils.IsNil(o.Signup) {
		var ret ServerConfigCustomSignup
		return ret
	}
	return *o.Signup
}

// GetSignupOk returns a tuple with the Signup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigCustom) GetSignupOk() (*ServerConfigCustomSignup, bool) {
	if o == nil || utils.IsNil(o.Signup) {
		return nil, false
	}
	return o.Signup, true
}

// HasSignup returns a boolean if a field has been set.
func (o *ServerConfigCustom) HasSignup() bool {
	if o != nil && !utils.IsNil(o.Signup) {
		return true
	}

	return false
}

// SetSignup gets a reference to the given ServerConfigCustomSignup and assigns it to the Signup field.
func (o *ServerConfigCustom) SetSignup(v ServerConfigCustomSignup) {
	o.Signup = &v
}

// GetAdmin returns the Admin field value if set, zero value otherwise.
func (o *ServerConfigCustom) GetAdmin() ServerConfigCustomAdmin {
	if o == nil || utils.IsNil(o.Admin) {
		var ret ServerConfigCustomAdmin
		return ret
	}
	return *o.Admin
}

// GetAdminOk returns a tuple with the Admin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigCustom) GetAdminOk() (*ServerConfigCustomAdmin, bool) {
	if o == nil || utils.IsNil(o.Admin) {
		return nil, false
	}
	return o.Admin, true
}

// HasAdmin returns a boolean if a field has been set.
func (o *ServerConfigCustom) HasAdmin() bool {
	if o != nil && !utils.IsNil(o.Admin) {
		return true
	}

	return false
}

// SetAdmin gets a reference to the given ServerConfigCustomAdmin and assigns it to the Admin field.
func (o *ServerConfigCustom) SetAdmin(v ServerConfigCustomAdmin) {
	o.Admin = &v
}

// GetContactForm returns the ContactForm field value if set, zero value otherwise.
func (o *ServerConfigCustom) GetContactForm() ServerConfigEmail {
	if o == nil || utils.IsNil(o.ContactForm) {
		var ret ServerConfigEmail
		return ret
	}
	return *o.ContactForm
}

// GetContactFormOk returns a tuple with the ContactForm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigCustom) GetContactFormOk() (*ServerConfigEmail, bool) {
	if o == nil || utils.IsNil(o.ContactForm) {
		return nil, false
	}
	return o.ContactForm, true
}

// HasContactForm returns a boolean if a field has been set.
func (o *ServerConfigCustom) HasContactForm() bool {
	if o != nil && !utils.IsNil(o.ContactForm) {
		return true
	}

	return false
}

// SetContactForm gets a reference to the given ServerConfigEmail and assigns it to the ContactForm field.
func (o *ServerConfigCustom) SetContactForm(v ServerConfigEmail) {
	o.ContactForm = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *ServerConfigCustom) GetUser() ServerConfigCustomUser {
	if o == nil || utils.IsNil(o.User) {
		var ret ServerConfigCustomUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigCustom) GetUserOk() (*ServerConfigCustomUser, bool) {
	if o == nil || utils.IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *ServerConfigCustom) HasUser() bool {
	if o != nil && !utils.IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given ServerConfigCustomUser and assigns it to the User field.
func (o *ServerConfigCustom) SetUser(v ServerConfigCustomUser) {
	o.User = &v
}

// GetTranscoding returns the Transcoding field value if set, zero value otherwise.
func (o *ServerConfigCustom) GetTranscoding() ServerConfigCustomTranscoding {
	if o == nil || utils.IsNil(o.Transcoding) {
		var ret ServerConfigCustomTranscoding
		return ret
	}
	return *o.Transcoding
}

// GetTranscodingOk returns a tuple with the Transcoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigCustom) GetTranscodingOk() (*ServerConfigCustomTranscoding, bool) {
	if o == nil || utils.IsNil(o.Transcoding) {
		return nil, false
	}
	return o.Transcoding, true
}

// HasTranscoding returns a boolean if a field has been set.
func (o *ServerConfigCustom) HasTranscoding() bool {
	if o != nil && !utils.IsNil(o.Transcoding) {
		return true
	}

	return false
}

// SetTranscoding gets a reference to the given ServerConfigCustomTranscoding and assigns it to the Transcoding field.
func (o *ServerConfigCustom) SetTranscoding(v ServerConfigCustomTranscoding) {
	o.Transcoding = &v
}

// GetImport returns the Import field value if set, zero value otherwise.
func (o *ServerConfigCustom) GetImport() ServerConfigCustomImport {
	if o == nil || utils.IsNil(o.Import) {
		var ret ServerConfigCustomImport
		return ret
	}
	return *o.Import
}

// GetImportOk returns a tuple with the Import field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigCustom) GetImportOk() (*ServerConfigCustomImport, bool) {
	if o == nil || utils.IsNil(o.Import) {
		return nil, false
	}
	return o.Import, true
}

// HasImport returns a boolean if a field has been set.
func (o *ServerConfigCustom) HasImport() bool {
	if o != nil && !utils.IsNil(o.Import) {
		return true
	}

	return false
}

// SetImport gets a reference to the given ServerConfigCustomImport and assigns it to the Import field.
func (o *ServerConfigCustom) SetImport(v ServerConfigCustomImport) {
	o.Import = &v
}

// GetAutoBlacklist returns the AutoBlacklist field value if set, zero value otherwise.
func (o *ServerConfigCustom) GetAutoBlacklist() ServerConfigAutoBlacklist {
	if o == nil || utils.IsNil(o.AutoBlacklist) {
		var ret ServerConfigAutoBlacklist
		return ret
	}
	return *o.AutoBlacklist
}

// GetAutoBlacklistOk returns a tuple with the AutoBlacklist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigCustom) GetAutoBlacklistOk() (*ServerConfigAutoBlacklist, bool) {
	if o == nil || utils.IsNil(o.AutoBlacklist) {
		return nil, false
	}
	return o.AutoBlacklist, true
}

// HasAutoBlacklist returns a boolean if a field has been set.
func (o *ServerConfigCustom) HasAutoBlacklist() bool {
	if o != nil && !utils.IsNil(o.AutoBlacklist) {
		return true
	}

	return false
}

// SetAutoBlacklist gets a reference to the given ServerConfigAutoBlacklist and assigns it to the AutoBlacklist field.
func (o *ServerConfigCustom) SetAutoBlacklist(v ServerConfigAutoBlacklist) {
	o.AutoBlacklist = &v
}

// GetFollowers returns the Followers field value if set, zero value otherwise.
func (o *ServerConfigCustom) GetFollowers() ServerConfigCustomFollowers {
	if o == nil || utils.IsNil(o.Followers) {
		var ret ServerConfigCustomFollowers
		return ret
	}
	return *o.Followers
}

// GetFollowersOk returns a tuple with the Followers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerConfigCustom) GetFollowersOk() (*ServerConfigCustomFollowers, bool) {
	if o == nil || utils.IsNil(o.Followers) {
		return nil, false
	}
	return o.Followers, true
}

// HasFollowers returns a boolean if a field has been set.
func (o *ServerConfigCustom) HasFollowers() bool {
	if o != nil && !utils.IsNil(o.Followers) {
		return true
	}

	return false
}

// SetFollowers gets a reference to the given ServerConfigCustomFollowers and assigns it to the Followers field.
func (o *ServerConfigCustom) SetFollowers(v ServerConfigCustomFollowers) {
	o.Followers = &v
}

func (o ServerConfigCustom) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerConfigCustom) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Instance) {
		toSerialize["instance"] = o.Instance
	}
	if !utils.IsNil(o.Theme) {
		toSerialize["theme"] = o.Theme
	}
	if !utils.IsNil(o.Services) {
		toSerialize["services"] = o.Services
	}
	if !utils.IsNil(o.Cache) {
		toSerialize["cache"] = o.Cache
	}
	if !utils.IsNil(o.Signup) {
		toSerialize["signup"] = o.Signup
	}
	if !utils.IsNil(o.Admin) {
		toSerialize["admin"] = o.Admin
	}
	if !utils.IsNil(o.ContactForm) {
		toSerialize["contactForm"] = o.ContactForm
	}
	if !utils.IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !utils.IsNil(o.Transcoding) {
		toSerialize["transcoding"] = o.Transcoding
	}
	if !utils.IsNil(o.Import) {
		toSerialize["import"] = o.Import
	}
	if !utils.IsNil(o.AutoBlacklist) {
		toSerialize["autoBlacklist"] = o.AutoBlacklist
	}
	if !utils.IsNil(o.Followers) {
		toSerialize["followers"] = o.Followers
	}
	return toSerialize, nil
}

type NullableServerConfigCustom struct {
	value *ServerConfigCustom
	isSet bool
}

func (v NullableServerConfigCustom) Get() *ServerConfigCustom {
	return v.value
}

func (v *NullableServerConfigCustom) Set(val *ServerConfigCustom) {
	v.value = val
	v.isSet = true
}

func (v NullableServerConfigCustom) IsSet() bool {
	return v.isSet
}

func (v *NullableServerConfigCustom) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerConfigCustom(val *ServerConfigCustom) *NullableServerConfigCustom {
	return &NullableServerConfigCustom{value: val, isSet: true}
}

func (v NullableServerConfigCustom) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerConfigCustom) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
