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

// checks if the UpdateUser type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &UpdateUser{}

// UpdateUser struct for UpdateUser
type UpdateUser struct {
	// The updated email of the user
	Email *string `json:"email,omitempty"`
	// Set the email as verified
	EmailVerified *bool `json:"emailVerified,omitempty"`
	// The updated video quota of the user in bytes
	VideoQuota *int32 `json:"videoQuota,omitempty"`
	// The updated daily video quota of the user in bytes
	VideoQuotaDaily *int32 `json:"videoQuotaDaily,omitempty"`
	// The auth plugin to use to authenticate the user
	PluginAuth utils.NullableString `json:"pluginAuth,omitempty"`
	Role       *UserRole            `json:"role,omitempty"`
	AdminFlags *UserAdminFlags      `json:"adminFlags,omitempty"`
	Password   *string              `json:"password,omitempty"`
}

// NewUpdateUser instantiates a new UpdateUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateUser() *UpdateUser {
	this := UpdateUser{}
	return &this
}

// NewUpdateUserWithDefaults instantiates a new UpdateUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateUserWithDefaults() *UpdateUser {
	this := UpdateUser{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UpdateUser) GetEmail() string {
	if o == nil || utils.IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUser) GetEmailOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UpdateUser) HasEmail() bool {
	if o != nil && !utils.IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UpdateUser) SetEmail(v string) {
	o.Email = &v
}

// GetEmailVerified returns the EmailVerified field value if set, zero value otherwise.
func (o *UpdateUser) GetEmailVerified() bool {
	if o == nil || utils.IsNil(o.EmailVerified) {
		var ret bool
		return ret
	}
	return *o.EmailVerified
}

// GetEmailVerifiedOk returns a tuple with the EmailVerified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUser) GetEmailVerifiedOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.EmailVerified) {
		return nil, false
	}
	return o.EmailVerified, true
}

// HasEmailVerified returns a boolean if a field has been set.
func (o *UpdateUser) HasEmailVerified() bool {
	if o != nil && !utils.IsNil(o.EmailVerified) {
		return true
	}

	return false
}

// SetEmailVerified gets a reference to the given bool and assigns it to the EmailVerified field.
func (o *UpdateUser) SetEmailVerified(v bool) {
	o.EmailVerified = &v
}

// GetVideoQuota returns the VideoQuota field value if set, zero value otherwise.
func (o *UpdateUser) GetVideoQuota() int32 {
	if o == nil || utils.IsNil(o.VideoQuota) {
		var ret int32
		return ret
	}
	return *o.VideoQuota
}

// GetVideoQuotaOk returns a tuple with the VideoQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUser) GetVideoQuotaOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.VideoQuota) {
		return nil, false
	}
	return o.VideoQuota, true
}

// HasVideoQuota returns a boolean if a field has been set.
func (o *UpdateUser) HasVideoQuota() bool {
	if o != nil && !utils.IsNil(o.VideoQuota) {
		return true
	}

	return false
}

// SetVideoQuota gets a reference to the given int32 and assigns it to the VideoQuota field.
func (o *UpdateUser) SetVideoQuota(v int32) {
	o.VideoQuota = &v
}

// GetVideoQuotaDaily returns the VideoQuotaDaily field value if set, zero value otherwise.
func (o *UpdateUser) GetVideoQuotaDaily() int32 {
	if o == nil || utils.IsNil(o.VideoQuotaDaily) {
		var ret int32
		return ret
	}
	return *o.VideoQuotaDaily
}

// GetVideoQuotaDailyOk returns a tuple with the VideoQuotaDaily field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUser) GetVideoQuotaDailyOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.VideoQuotaDaily) {
		return nil, false
	}
	return o.VideoQuotaDaily, true
}

// HasVideoQuotaDaily returns a boolean if a field has been set.
func (o *UpdateUser) HasVideoQuotaDaily() bool {
	if o != nil && !utils.IsNil(o.VideoQuotaDaily) {
		return true
	}

	return false
}

// SetVideoQuotaDaily gets a reference to the given int32 and assigns it to the VideoQuotaDaily field.
func (o *UpdateUser) SetVideoQuotaDaily(v int32) {
	o.VideoQuotaDaily = &v
}

// GetPluginAuth returns the PluginAuth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateUser) GetPluginAuth() string {
	if o == nil || utils.IsNil(o.PluginAuth.Get()) {
		var ret string
		return ret
	}
	return *o.PluginAuth.Get()
}

// GetPluginAuthOk returns a tuple with the PluginAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateUser) GetPluginAuthOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PluginAuth.Get(), o.PluginAuth.IsSet()
}

// HasPluginAuth returns a boolean if a field has been set.
func (o *UpdateUser) HasPluginAuth() bool {
	if o != nil && o.PluginAuth.IsSet() {
		return true
	}

	return false
}

// SetPluginAuth gets a reference to the given NullableString and assigns it to the PluginAuth field.
func (o *UpdateUser) SetPluginAuth(v string) {
	o.PluginAuth.Set(&v)
}

// SetPluginAuthNil sets the value for PluginAuth to be an explicit nil
func (o *UpdateUser) SetPluginAuthNil() {
	o.PluginAuth.Set(nil)
}

// UnsetPluginAuth ensures that no value is present for PluginAuth, not even an explicit nil
func (o *UpdateUser) UnsetPluginAuth() {
	o.PluginAuth.Unset()
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *UpdateUser) GetRole() UserRole {
	if o == nil || utils.IsNil(o.Role) {
		var ret UserRole
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUser) GetRoleOk() (*UserRole, bool) {
	if o == nil || utils.IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *UpdateUser) HasRole() bool {
	if o != nil && !utils.IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given UserRole and assigns it to the Role field.
func (o *UpdateUser) SetRole(v UserRole) {
	o.Role = &v
}

// GetAdminFlags returns the AdminFlags field value if set, zero value otherwise.
func (o *UpdateUser) GetAdminFlags() UserAdminFlags {
	if o == nil || utils.IsNil(o.AdminFlags) {
		var ret UserAdminFlags
		return ret
	}
	return *o.AdminFlags
}

// GetAdminFlagsOk returns a tuple with the AdminFlags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUser) GetAdminFlagsOk() (*UserAdminFlags, bool) {
	if o == nil || utils.IsNil(o.AdminFlags) {
		return nil, false
	}
	return o.AdminFlags, true
}

// HasAdminFlags returns a boolean if a field has been set.
func (o *UpdateUser) HasAdminFlags() bool {
	if o != nil && !utils.IsNil(o.AdminFlags) {
		return true
	}

	return false
}

// SetAdminFlags gets a reference to the given UserAdminFlags and assigns it to the AdminFlags field.
func (o *UpdateUser) SetAdminFlags(v UserAdminFlags) {
	o.AdminFlags = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *UpdateUser) GetPassword() string {
	if o == nil || utils.IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUser) GetPasswordOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *UpdateUser) HasPassword() bool {
	if o != nil && !utils.IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *UpdateUser) SetPassword(v string) {
	o.Password = &v
}

func (o UpdateUser) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !utils.IsNil(o.EmailVerified) {
		toSerialize["emailVerified"] = o.EmailVerified
	}
	if !utils.IsNil(o.VideoQuota) {
		toSerialize["videoQuota"] = o.VideoQuota
	}
	if !utils.IsNil(o.VideoQuotaDaily) {
		toSerialize["videoQuotaDaily"] = o.VideoQuotaDaily
	}
	if o.PluginAuth.IsSet() {
		toSerialize["pluginAuth"] = o.PluginAuth.Get()
	}
	if !utils.IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !utils.IsNil(o.AdminFlags) {
		toSerialize["adminFlags"] = o.AdminFlags
	}
	if !utils.IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	return toSerialize, nil
}

type NullableUpdateUser struct {
	value *UpdateUser
	isSet bool
}

func (v NullableUpdateUser) Get() *UpdateUser {
	return v.value
}

func (v *NullableUpdateUser) Set(val *UpdateUser) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateUser(val *UpdateUser) *NullableUpdateUser {
	return &NullableUpdateUser{value: val, isSet: true}
}

func (v NullableUpdateUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
