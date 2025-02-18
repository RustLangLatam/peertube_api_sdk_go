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

// checks if the RunnerJobAdmin type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RunnerJobAdmin{}

// RunnerJobAdmin struct for RunnerJobAdmin
type RunnerJobAdmin struct {
	Uuid    *string                 `json:"uuid,omitempty" validate:"regexp=^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$"`
	Type    *RunnerJobType          `json:"type,omitempty"`
	State   *RunnerJobStateConstant `json:"state,omitempty"`
	Payload *RunnerJobPayload       `json:"payload,omitempty"`
	// Number of times a remote runner failed to process this job. After too many failures, the job in \"error\" state
	Failures *int32 `json:"failures,omitempty"`
	// Error message if the job is errored
	Error utils.NullableString `json:"error,omitempty"`
	// Percentage progress
	Progress *int32 `json:"progress,omitempty"`
	// Job priority (less has more priority)
	Priority       *int32                  `json:"priority,omitempty"`
	UpdatedAt      *time.Time              `json:"updatedAt,omitempty"`
	CreatedAt      *time.Time              `json:"createdAt,omitempty"`
	StartedAt      *time.Time              `json:"startedAt,omitempty"`
	FinishedAt     *time.Time              `json:"finishedAt,omitempty"`
	Parent         NullableRunnerJobParent `json:"parent,omitempty"`
	Runner         NullableRunnerJobRunner `json:"runner,omitempty"`
	PrivatePayload map[string]interface{}  `json:"privatePayload,omitempty"`
}

// NewRunnerJobAdmin instantiates a new RunnerJobAdmin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunnerJobAdmin() *RunnerJobAdmin {
	this := RunnerJobAdmin{}
	return &this
}

// NewRunnerJobAdminWithDefaults instantiates a new RunnerJobAdmin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunnerJobAdminWithDefaults() *RunnerJobAdmin {
	this := RunnerJobAdmin{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *RunnerJobAdmin) GetUuid() string {
	if o == nil || utils.IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunnerJobAdmin) GetUuidOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *RunnerJobAdmin) HasUuid() bool {
	if o != nil && !utils.IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *RunnerJobAdmin) SetUuid(v string) {
	o.Uuid = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RunnerJobAdmin) GetType() RunnerJobType {
	if o == nil || utils.IsNil(o.Type) {
		var ret RunnerJobType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunnerJobAdmin) GetTypeOk() (*RunnerJobType, bool) {
	if o == nil || utils.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RunnerJobAdmin) HasType() bool {
	if o != nil && !utils.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given RunnerJobType and assigns it to the Type field.
func (o *RunnerJobAdmin) SetType(v RunnerJobType) {
	o.Type = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *RunnerJobAdmin) GetState() RunnerJobStateConstant {
	if o == nil || utils.IsNil(o.State) {
		var ret RunnerJobStateConstant
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunnerJobAdmin) GetStateOk() (*RunnerJobStateConstant, bool) {
	if o == nil || utils.IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *RunnerJobAdmin) HasState() bool {
	if o != nil && !utils.IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given RunnerJobStateConstant and assigns it to the State field.
func (o *RunnerJobAdmin) SetState(v RunnerJobStateConstant) {
	o.State = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *RunnerJobAdmin) GetPayload() RunnerJobPayload {
	if o == nil || utils.IsNil(o.Payload) {
		var ret RunnerJobPayload
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunnerJobAdmin) GetPayloadOk() (*RunnerJobPayload, bool) {
	if o == nil || utils.IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *RunnerJobAdmin) HasPayload() bool {
	if o != nil && !utils.IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given RunnerJobPayload and assigns it to the Payload field.
func (o *RunnerJobAdmin) SetPayload(v RunnerJobPayload) {
	o.Payload = &v
}

// GetFailures returns the Failures field value if set, zero value otherwise.
func (o *RunnerJobAdmin) GetFailures() int32 {
	if o == nil || utils.IsNil(o.Failures) {
		var ret int32
		return ret
	}
	return *o.Failures
}

// GetFailuresOk returns a tuple with the Failures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunnerJobAdmin) GetFailuresOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Failures) {
		return nil, false
	}
	return o.Failures, true
}

// HasFailures returns a boolean if a field has been set.
func (o *RunnerJobAdmin) HasFailures() bool {
	if o != nil && !utils.IsNil(o.Failures) {
		return true
	}

	return false
}

// SetFailures gets a reference to the given int32 and assigns it to the Failures field.
func (o *RunnerJobAdmin) SetFailures(v int32) {
	o.Failures = &v
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RunnerJobAdmin) GetError() string {
	if o == nil || utils.IsNil(o.Error.Get()) {
		var ret string
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RunnerJobAdmin) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *RunnerJobAdmin) HasError() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullableString and assigns it to the Error field.
func (o *RunnerJobAdmin) SetError(v string) {
	o.Error.Set(&v)
}

// SetErrorNil sets the value for Error to be an explicit nil
func (o *RunnerJobAdmin) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *RunnerJobAdmin) UnsetError() {
	o.Error.Unset()
}

// GetProgress returns the Progress field value if set, zero value otherwise.
func (o *RunnerJobAdmin) GetProgress() int32 {
	if o == nil || utils.IsNil(o.Progress) {
		var ret int32
		return ret
	}
	return *o.Progress
}

// GetProgressOk returns a tuple with the Progress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunnerJobAdmin) GetProgressOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Progress) {
		return nil, false
	}
	return o.Progress, true
}

// HasProgress returns a boolean if a field has been set.
func (o *RunnerJobAdmin) HasProgress() bool {
	if o != nil && !utils.IsNil(o.Progress) {
		return true
	}

	return false
}

// SetProgress gets a reference to the given int32 and assigns it to the Progress field.
func (o *RunnerJobAdmin) SetProgress(v int32) {
	o.Progress = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *RunnerJobAdmin) GetPriority() int32 {
	if o == nil || utils.IsNil(o.Priority) {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunnerJobAdmin) GetPriorityOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *RunnerJobAdmin) HasPriority() bool {
	if o != nil && !utils.IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *RunnerJobAdmin) SetPriority(v int32) {
	o.Priority = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *RunnerJobAdmin) GetUpdatedAt() time.Time {
	if o == nil || utils.IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunnerJobAdmin) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *RunnerJobAdmin) HasUpdatedAt() bool {
	if o != nil && !utils.IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *RunnerJobAdmin) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RunnerJobAdmin) GetCreatedAt() time.Time {
	if o == nil || utils.IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunnerJobAdmin) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RunnerJobAdmin) HasCreatedAt() bool {
	if o != nil && !utils.IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *RunnerJobAdmin) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *RunnerJobAdmin) GetStartedAt() time.Time {
	if o == nil || utils.IsNil(o.StartedAt) {
		var ret time.Time
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunnerJobAdmin) GetStartedAtOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.StartedAt) {
		return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *RunnerJobAdmin) HasStartedAt() bool {
	if o != nil && !utils.IsNil(o.StartedAt) {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given time.Time and assigns it to the StartedAt field.
func (o *RunnerJobAdmin) SetStartedAt(v time.Time) {
	o.StartedAt = &v
}

// GetFinishedAt returns the FinishedAt field value if set, zero value otherwise.
func (o *RunnerJobAdmin) GetFinishedAt() time.Time {
	if o == nil || utils.IsNil(o.FinishedAt) {
		var ret time.Time
		return ret
	}
	return *o.FinishedAt
}

// GetFinishedAtOk returns a tuple with the FinishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunnerJobAdmin) GetFinishedAtOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.FinishedAt) {
		return nil, false
	}
	return o.FinishedAt, true
}

// HasFinishedAt returns a boolean if a field has been set.
func (o *RunnerJobAdmin) HasFinishedAt() bool {
	if o != nil && !utils.IsNil(o.FinishedAt) {
		return true
	}

	return false
}

// SetFinishedAt gets a reference to the given time.Time and assigns it to the FinishedAt field.
func (o *RunnerJobAdmin) SetFinishedAt(v time.Time) {
	o.FinishedAt = &v
}

// GetParent returns the Parent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RunnerJobAdmin) GetParent() RunnerJobParent {
	if o == nil || utils.IsNil(o.Parent.Get()) {
		var ret RunnerJobParent
		return ret
	}
	return *o.Parent.Get()
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RunnerJobAdmin) GetParentOk() (*RunnerJobParent, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parent.Get(), o.Parent.IsSet()
}

// HasParent returns a boolean if a field has been set.
func (o *RunnerJobAdmin) HasParent() bool {
	if o != nil && o.Parent.IsSet() {
		return true
	}

	return false
}

// SetParent gets a reference to the given NullableRunnerJobParent and assigns it to the Parent field.
func (o *RunnerJobAdmin) SetParent(v RunnerJobParent) {
	o.Parent.Set(&v)
}

// SetParentNil sets the value for Parent to be an explicit nil
func (o *RunnerJobAdmin) SetParentNil() {
	o.Parent.Set(nil)
}

// UnsetParent ensures that no value is present for Parent, not even an explicit nil
func (o *RunnerJobAdmin) UnsetParent() {
	o.Parent.Unset()
}

// GetRunner returns the Runner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RunnerJobAdmin) GetRunner() RunnerJobRunner {
	if o == nil || utils.IsNil(o.Runner.Get()) {
		var ret RunnerJobRunner
		return ret
	}
	return *o.Runner.Get()
}

// GetRunnerOk returns a tuple with the Runner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RunnerJobAdmin) GetRunnerOk() (*RunnerJobRunner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Runner.Get(), o.Runner.IsSet()
}

// HasRunner returns a boolean if a field has been set.
func (o *RunnerJobAdmin) HasRunner() bool {
	if o != nil && o.Runner.IsSet() {
		return true
	}

	return false
}

// SetRunner gets a reference to the given NullableRunnerJobRunner and assigns it to the Runner field.
func (o *RunnerJobAdmin) SetRunner(v RunnerJobRunner) {
	o.Runner.Set(&v)
}

// SetRunnerNil sets the value for Runner to be an explicit nil
func (o *RunnerJobAdmin) SetRunnerNil() {
	o.Runner.Set(nil)
}

// UnsetRunner ensures that no value is present for Runner, not even an explicit nil
func (o *RunnerJobAdmin) UnsetRunner() {
	o.Runner.Unset()
}

// GetPrivatePayload returns the PrivatePayload field value if set, zero value otherwise.
func (o *RunnerJobAdmin) GetPrivatePayload() map[string]interface{} {
	if o == nil || utils.IsNil(o.PrivatePayload) {
		var ret map[string]interface{}
		return ret
	}
	return o.PrivatePayload
}

// GetPrivatePayloadOk returns a tuple with the PrivatePayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunnerJobAdmin) GetPrivatePayloadOk() (map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.PrivatePayload) {
		return map[string]interface{}{}, false
	}
	return o.PrivatePayload, true
}

// HasPrivatePayload returns a boolean if a field has been set.
func (o *RunnerJobAdmin) HasPrivatePayload() bool {
	if o != nil && !utils.IsNil(o.PrivatePayload) {
		return true
	}

	return false
}

// SetPrivatePayload gets a reference to the given map[string]interface{} and assigns it to the PrivatePayload field.
func (o *RunnerJobAdmin) SetPrivatePayload(v map[string]interface{}) {
	o.PrivatePayload = v
}

func (o RunnerJobAdmin) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RunnerJobAdmin) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !utils.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !utils.IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !utils.IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !utils.IsNil(o.Failures) {
		toSerialize["failures"] = o.Failures
	}
	if o.Error.IsSet() {
		toSerialize["error"] = o.Error.Get()
	}
	if !utils.IsNil(o.Progress) {
		toSerialize["progress"] = o.Progress
	}
	if !utils.IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !utils.IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !utils.IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !utils.IsNil(o.StartedAt) {
		toSerialize["startedAt"] = o.StartedAt
	}
	if !utils.IsNil(o.FinishedAt) {
		toSerialize["finishedAt"] = o.FinishedAt
	}
	if o.Parent.IsSet() {
		toSerialize["parent"] = o.Parent.Get()
	}
	if o.Runner.IsSet() {
		toSerialize["runner"] = o.Runner.Get()
	}
	if !utils.IsNil(o.PrivatePayload) {
		toSerialize["privatePayload"] = o.PrivatePayload
	}
	return toSerialize, nil
}

type NullableRunnerJobAdmin struct {
	value *RunnerJobAdmin
	isSet bool
}

func (v NullableRunnerJobAdmin) Get() *RunnerJobAdmin {
	return v.value
}

func (v *NullableRunnerJobAdmin) Set(val *RunnerJobAdmin) {
	v.value = val
	v.isSet = true
}

func (v NullableRunnerJobAdmin) IsSet() bool {
	return v.isSet
}

func (v *NullableRunnerJobAdmin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunnerJobAdmin(val *RunnerJobAdmin) *NullableRunnerJobAdmin {
	return &NullableRunnerJobAdmin{value: val, isSet: true}
}

func (v NullableRunnerJobAdmin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunnerJobAdmin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
