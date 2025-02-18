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

// checks if the UserRegistration type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &UserRegistration{}

// UserRegistration struct for UserRegistration
type UserRegistration struct {
	Id                 *int32                       `json:"id,omitempty"`
	State              *UserRegistrationState       `json:"state,omitempty"`
	RegistrationReason *string                      `json:"registrationReason,omitempty"`
	ModerationResponse utils.NullableString         `json:"moderationResponse,omitempty"`
	Username           *string                      `json:"username,omitempty"`
	Email              *string                      `json:"email,omitempty"`
	EmailVerified      *bool                        `json:"emailVerified,omitempty"`
	AccountDisplayName *string                      `json:"accountDisplayName,omitempty"`
	ChannelHandle      *string                      `json:"channelHandle,omitempty"`
	ChannelDisplayName *string                      `json:"channelDisplayName,omitempty"`
	CreatedAt          *time.Time                   `json:"createdAt,omitempty"`
	UpdatedAt          *time.Time                   `json:"updatedAt,omitempty"`
	User               NullableUserRegistrationUser `json:"user,omitempty"`
}

// NewUserRegistration instantiates a new UserRegistration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserRegistration() *UserRegistration {
	this := UserRegistration{}
	return &this
}

// NewUserRegistrationWithDefaults instantiates a new UserRegistration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserRegistrationWithDefaults() *UserRegistration {
	this := UserRegistration{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserRegistration) GetId() int32 {
	if o == nil || utils.IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRegistration) GetIdOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserRegistration) HasId() bool {
	if o != nil && !utils.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *UserRegistration) SetId(v int32) {
	o.Id = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *UserRegistration) GetState() UserRegistrationState {
	if o == nil || utils.IsNil(o.State) {
		var ret UserRegistrationState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRegistration) GetStateOk() (*UserRegistrationState, bool) {
	if o == nil || utils.IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *UserRegistration) HasState() bool {
	if o != nil && !utils.IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given UserRegistrationState and assigns it to the State field.
func (o *UserRegistration) SetState(v UserRegistrationState) {
	o.State = &v
}

// GetRegistrationReason returns the RegistrationReason field value if set, zero value otherwise.
func (o *UserRegistration) GetRegistrationReason() string {
	if o == nil || utils.IsNil(o.RegistrationReason) {
		var ret string
		return ret
	}
	return *o.RegistrationReason
}

// GetRegistrationReasonOk returns a tuple with the RegistrationReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRegistration) GetRegistrationReasonOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RegistrationReason) {
		return nil, false
	}
	return o.RegistrationReason, true
}

// HasRegistrationReason returns a boolean if a field has been set.
func (o *UserRegistration) HasRegistrationReason() bool {
	if o != nil && !utils.IsNil(o.RegistrationReason) {
		return true
	}

	return false
}

// SetRegistrationReason gets a reference to the given string and assigns it to the RegistrationReason field.
func (o *UserRegistration) SetRegistrationReason(v string) {
	o.RegistrationReason = &v
}

// GetModerationResponse returns the ModerationResponse field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserRegistration) GetModerationResponse() string {
	if o == nil || utils.IsNil(o.ModerationResponse.Get()) {
		var ret string
		return ret
	}
	return *o.ModerationResponse.Get()
}

// GetModerationResponseOk returns a tuple with the ModerationResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserRegistration) GetModerationResponseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ModerationResponse.Get(), o.ModerationResponse.IsSet()
}

// HasModerationResponse returns a boolean if a field has been set.
func (o *UserRegistration) HasModerationResponse() bool {
	if o != nil && o.ModerationResponse.IsSet() {
		return true
	}

	return false
}

// SetModerationResponse gets a reference to the given NullableString and assigns it to the ModerationResponse field.
func (o *UserRegistration) SetModerationResponse(v string) {
	o.ModerationResponse.Set(&v)
}

// SetModerationResponseNil sets the value for ModerationResponse to be an explicit nil
func (o *UserRegistration) SetModerationResponseNil() {
	o.ModerationResponse.Set(nil)
}

// UnsetModerationResponse ensures that no value is present for ModerationResponse, not even an explicit nil
func (o *UserRegistration) UnsetModerationResponse() {
	o.ModerationResponse.Unset()
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *UserRegistration) GetUsername() string {
	if o == nil || utils.IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRegistration) GetUsernameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *UserRegistration) HasUsername() bool {
	if o != nil && !utils.IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *UserRegistration) SetUsername(v string) {
	o.Username = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UserRegistration) GetEmail() string {
	if o == nil || utils.IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRegistration) GetEmailOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UserRegistration) HasEmail() bool {
	if o != nil && !utils.IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UserRegistration) SetEmail(v string) {
	o.Email = &v
}

// GetEmailVerified returns the EmailVerified field value if set, zero value otherwise.
func (o *UserRegistration) GetEmailVerified() bool {
	if o == nil || utils.IsNil(o.EmailVerified) {
		var ret bool
		return ret
	}
	return *o.EmailVerified
}

// GetEmailVerifiedOk returns a tuple with the EmailVerified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRegistration) GetEmailVerifiedOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.EmailVerified) {
		return nil, false
	}
	return o.EmailVerified, true
}

// HasEmailVerified returns a boolean if a field has been set.
func (o *UserRegistration) HasEmailVerified() bool {
	if o != nil && !utils.IsNil(o.EmailVerified) {
		return true
	}

	return false
}

// SetEmailVerified gets a reference to the given bool and assigns it to the EmailVerified field.
func (o *UserRegistration) SetEmailVerified(v bool) {
	o.EmailVerified = &v
}

// GetAccountDisplayName returns the AccountDisplayName field value if set, zero value otherwise.
func (o *UserRegistration) GetAccountDisplayName() string {
	if o == nil || utils.IsNil(o.AccountDisplayName) {
		var ret string
		return ret
	}
	return *o.AccountDisplayName
}

// GetAccountDisplayNameOk returns a tuple with the AccountDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRegistration) GetAccountDisplayNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.AccountDisplayName) {
		return nil, false
	}
	return o.AccountDisplayName, true
}

// HasAccountDisplayName returns a boolean if a field has been set.
func (o *UserRegistration) HasAccountDisplayName() bool {
	if o != nil && !utils.IsNil(o.AccountDisplayName) {
		return true
	}

	return false
}

// SetAccountDisplayName gets a reference to the given string and assigns it to the AccountDisplayName field.
func (o *UserRegistration) SetAccountDisplayName(v string) {
	o.AccountDisplayName = &v
}

// GetChannelHandle returns the ChannelHandle field value if set, zero value otherwise.
func (o *UserRegistration) GetChannelHandle() string {
	if o == nil || utils.IsNil(o.ChannelHandle) {
		var ret string
		return ret
	}
	return *o.ChannelHandle
}

// GetChannelHandleOk returns a tuple with the ChannelHandle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRegistration) GetChannelHandleOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ChannelHandle) {
		return nil, false
	}
	return o.ChannelHandle, true
}

// HasChannelHandle returns a boolean if a field has been set.
func (o *UserRegistration) HasChannelHandle() bool {
	if o != nil && !utils.IsNil(o.ChannelHandle) {
		return true
	}

	return false
}

// SetChannelHandle gets a reference to the given string and assigns it to the ChannelHandle field.
func (o *UserRegistration) SetChannelHandle(v string) {
	o.ChannelHandle = &v
}

// GetChannelDisplayName returns the ChannelDisplayName field value if set, zero value otherwise.
func (o *UserRegistration) GetChannelDisplayName() string {
	if o == nil || utils.IsNil(o.ChannelDisplayName) {
		var ret string
		return ret
	}
	return *o.ChannelDisplayName
}

// GetChannelDisplayNameOk returns a tuple with the ChannelDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRegistration) GetChannelDisplayNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ChannelDisplayName) {
		return nil, false
	}
	return o.ChannelDisplayName, true
}

// HasChannelDisplayName returns a boolean if a field has been set.
func (o *UserRegistration) HasChannelDisplayName() bool {
	if o != nil && !utils.IsNil(o.ChannelDisplayName) {
		return true
	}

	return false
}

// SetChannelDisplayName gets a reference to the given string and assigns it to the ChannelDisplayName field.
func (o *UserRegistration) SetChannelDisplayName(v string) {
	o.ChannelDisplayName = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *UserRegistration) GetCreatedAt() time.Time {
	if o == nil || utils.IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRegistration) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *UserRegistration) HasCreatedAt() bool {
	if o != nil && !utils.IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *UserRegistration) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *UserRegistration) GetUpdatedAt() time.Time {
	if o == nil || utils.IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRegistration) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *UserRegistration) HasUpdatedAt() bool {
	if o != nil && !utils.IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *UserRegistration) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserRegistration) GetUser() UserRegistrationUser {
	if o == nil || utils.IsNil(o.User.Get()) {
		var ret UserRegistrationUser
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserRegistration) GetUserOk() (*UserRegistrationUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *UserRegistration) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableUserRegistrationUser and assigns it to the User field.
func (o *UserRegistration) SetUser(v UserRegistrationUser) {
	o.User.Set(&v)
}

// SetUserNil sets the value for User to be an explicit nil
func (o *UserRegistration) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *UserRegistration) UnsetUser() {
	o.User.Unset()
}

func (o UserRegistration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserRegistration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !utils.IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !utils.IsNil(o.RegistrationReason) {
		toSerialize["registrationReason"] = o.RegistrationReason
	}
	if o.ModerationResponse.IsSet() {
		toSerialize["moderationResponse"] = o.ModerationResponse.Get()
	}
	if !utils.IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !utils.IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !utils.IsNil(o.EmailVerified) {
		toSerialize["emailVerified"] = o.EmailVerified
	}
	if !utils.IsNil(o.AccountDisplayName) {
		toSerialize["accountDisplayName"] = o.AccountDisplayName
	}
	if !utils.IsNil(o.ChannelHandle) {
		toSerialize["channelHandle"] = o.ChannelHandle
	}
	if !utils.IsNil(o.ChannelDisplayName) {
		toSerialize["channelDisplayName"] = o.ChannelDisplayName
	}
	if !utils.IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !utils.IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.User.IsSet() {
		toSerialize["user"] = o.User.Get()
	}
	return toSerialize, nil
}

type NullableUserRegistration struct {
	value *UserRegistration
	isSet bool
}

func (v NullableUserRegistration) Get() *UserRegistration {
	return v.value
}

func (v *NullableUserRegistration) Set(val *UserRegistration) {
	v.value = val
	v.isSet = true
}

func (v NullableUserRegistration) IsSet() bool {
	return v.isSet
}

func (v *NullableUserRegistration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserRegistration(val *UserRegistration) *NullableUserRegistration {
	return &NullableUserRegistration{value: val, isSet: true}
}

func (v NullableUserRegistration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserRegistration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
