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

// checks if the ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob{}

// ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob struct for ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob
type ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob struct {
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
	Priority   *int32                  `json:"priority,omitempty"`
	UpdatedAt  *time.Time              `json:"updatedAt,omitempty"`
	CreatedAt  *time.Time              `json:"createdAt,omitempty"`
	StartedAt  *time.Time              `json:"startedAt,omitempty"`
	FinishedAt *time.Time              `json:"finishedAt,omitempty"`
	Parent     NullableRunnerJobParent `json:"parent,omitempty"`
	Runner     NullableRunnerJobRunner `json:"runner,omitempty"`
	JobToken   *string                 `json:"jobToken,omitempty"`
}

// NewApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob instantiates a new ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob() *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob {
	this := ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob{}
	return &this
}

// NewApiV1RunnersJobsJobUUIDAcceptPost200ResponseJobWithDefaults instantiates a new ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiV1RunnersJobsJobUUIDAcceptPost200ResponseJobWithDefaults() *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob {
	this := ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetUuid() string {
	if o == nil || utils.IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetUuidOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) HasUuid() bool {
	if o != nil && !utils.IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) SetUuid(v string) {
	o.Uuid = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetType() RunnerJobType {
	if o == nil || utils.IsNil(o.Type) {
		var ret RunnerJobType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetTypeOk() (*RunnerJobType, bool) {
	if o == nil || utils.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) HasType() bool {
	if o != nil && !utils.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given RunnerJobType and assigns it to the Type field.
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) SetType(v RunnerJobType) {
	o.Type = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetState() RunnerJobStateConstant {
	if o == nil || utils.IsNil(o.State) {
		var ret RunnerJobStateConstant
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetStateOk() (*RunnerJobStateConstant, bool) {
	if o == nil || utils.IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) HasState() bool {
	if o != nil && !utils.IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given RunnerJobStateConstant and assigns it to the State field.
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) SetState(v RunnerJobStateConstant) {
	o.State = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetPayload() RunnerJobPayload {
	if o == nil || utils.IsNil(o.Payload) {
		var ret RunnerJobPayload
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetPayloadOk() (*RunnerJobPayload, bool) {
	if o == nil || utils.IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) HasPayload() bool {
	if o != nil && !utils.IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given RunnerJobPayload and assigns it to the Payload field.
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) SetPayload(v RunnerJobPayload) {
	o.Payload = &v
}

// GetFailures returns the Failures field value if set, zero value otherwise.
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetFailures() int32 {
	if o == nil || utils.IsNil(o.Failures) {
		var ret int32
		return ret
	}
	return *o.Failures
}

// GetFailuresOk returns a tuple with the Failures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetFailuresOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Failures) {
		return nil, false
	}
	return o.Failures, true
}

// HasFailures returns a boolean if a field has been set.
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) HasFailures() bool {
	if o != nil && !utils.IsNil(o.Failures) {
		return true
	}

	return false
}

// SetFailures gets a reference to the given int32 and assigns it to the Failures field.
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) SetFailures(v int32) {
	o.Failures = &v
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetError() string {
	if o == nil || utils.IsNil(o.Error.Get()) {
		var ret string
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) HasError() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullableString and assigns it to the Error field.
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) SetError(v string) {
	o.Error.Set(&v)
}

// SetErrorNil sets the value for Error to be an explicit nil
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) UnsetError() {
	o.Error.Unset()
}

// GetProgress returns the Progress field value if set, zero value otherwise.
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetProgress() int32 {
	if o == nil || utils.IsNil(o.Progress) {
		var ret int32
		return ret
	}
	return *o.Progress
}

// GetProgressOk returns a tuple with the Progress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetProgressOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Progress) {
		return nil, false
	}
	return o.Progress, true
}

// HasProgress returns a boolean if a field has been set.
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) HasProgress() bool {
	if o != nil && !utils.IsNil(o.Progress) {
		return true
	}

	return false
}

// SetProgress gets a reference to the given int32 and assigns it to the Progress field.
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) SetProgress(v int32) {
	o.Progress = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetPriority() int32 {
	if o == nil || utils.IsNil(o.Priority) {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetPriorityOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) HasPriority() bool {
	if o != nil && !utils.IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) SetPriority(v int32) {
	o.Priority = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetUpdatedAt() time.Time {
	if o == nil || utils.IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) HasUpdatedAt() bool {
	if o != nil && !utils.IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetCreatedAt() time.Time {
	if o == nil || utils.IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) HasCreatedAt() bool {
	if o != nil && !utils.IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetStartedAt() time.Time {
	if o == nil || utils.IsNil(o.StartedAt) {
		var ret time.Time
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetStartedAtOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.StartedAt) {
		return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) HasStartedAt() bool {
	if o != nil && !utils.IsNil(o.StartedAt) {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given time.Time and assigns it to the StartedAt field.
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) SetStartedAt(v time.Time) {
	o.StartedAt = &v
}

// GetFinishedAt returns the FinishedAt field value if set, zero value otherwise.
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetFinishedAt() time.Time {
	if o == nil || utils.IsNil(o.FinishedAt) {
		var ret time.Time
		return ret
	}
	return *o.FinishedAt
}

// GetFinishedAtOk returns a tuple with the FinishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetFinishedAtOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.FinishedAt) {
		return nil, false
	}
	return o.FinishedAt, true
}

// HasFinishedAt returns a boolean if a field has been set.
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) HasFinishedAt() bool {
	if o != nil && !utils.IsNil(o.FinishedAt) {
		return true
	}

	return false
}

// SetFinishedAt gets a reference to the given time.Time and assigns it to the FinishedAt field.
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) SetFinishedAt(v time.Time) {
	o.FinishedAt = &v
}

// GetParent returns the Parent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetParent() RunnerJobParent {
	if o == nil || utils.IsNil(o.Parent.Get()) {
		var ret RunnerJobParent
		return ret
	}
	return *o.Parent.Get()
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetParentOk() (*RunnerJobParent, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parent.Get(), o.Parent.IsSet()
}

// HasParent returns a boolean if a field has been set.
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) HasParent() bool {
	if o != nil && o.Parent.IsSet() {
		return true
	}

	return false
}

// SetParent gets a reference to the given NullableRunnerJobParent and assigns it to the Parent field.
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) SetParent(v RunnerJobParent) {
	o.Parent.Set(&v)
}

// SetParentNil sets the value for Parent to be an explicit nil
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) SetParentNil() {
	o.Parent.Set(nil)
}

// UnsetParent ensures that no value is present for Parent, not even an explicit nil
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) UnsetParent() {
	o.Parent.Unset()
}

// GetRunner returns the Runner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetRunner() RunnerJobRunner {
	if o == nil || utils.IsNil(o.Runner.Get()) {
		var ret RunnerJobRunner
		return ret
	}
	return *o.Runner.Get()
}

// GetRunnerOk returns a tuple with the Runner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetRunnerOk() (*RunnerJobRunner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Runner.Get(), o.Runner.IsSet()
}

// HasRunner returns a boolean if a field has been set.
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) HasRunner() bool {
	if o != nil && o.Runner.IsSet() {
		return true
	}

	return false
}

// SetRunner gets a reference to the given NullableRunnerJobRunner and assigns it to the Runner field.
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) SetRunner(v RunnerJobRunner) {
	o.Runner.Set(&v)
}

// SetRunnerNil sets the value for Runner to be an explicit nil
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) SetRunnerNil() {
	o.Runner.Set(nil)
}

// UnsetRunner ensures that no value is present for Runner, not even an explicit nil
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) UnsetRunner() {
	o.Runner.Unset()
}

// GetJobToken returns the JobToken field value if set, zero value otherwise.
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetJobToken() string {
	if o == nil || utils.IsNil(o.JobToken) {
		var ret string
		return ret
	}
	return *o.JobToken
}

// GetJobTokenOk returns a tuple with the JobToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetJobTokenOk() (*string, bool) {
	if o == nil || utils.IsNil(o.JobToken) {
		return nil, false
	}
	return o.JobToken, true
}

// HasJobToken returns a boolean if a field has been set.
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) HasJobToken() bool {
	if o != nil && !utils.IsNil(o.JobToken) {
		return true
	}

	return false
}

// SetJobToken gets a reference to the given string and assigns it to the JobToken field.
func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) SetJobToken(v string) {
	o.JobToken = &v
}

func (o ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) ToMap() (map[string]interface{}, error) {
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
	if !utils.IsNil(o.JobToken) {
		toSerialize["jobToken"] = o.JobToken
	}
	return toSerialize, nil
}

type NullableApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob struct {
	value *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob
	isSet bool
}

func (v NullableApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) Get() *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob {
	return v.value
}

func (v *NullableApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) Set(val *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) {
	v.value = val
	v.isSet = true
}

func (v NullableApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) IsSet() bool {
	return v.isSet
}

func (v *NullableApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob(val *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) *NullableApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob {
	return &NullableApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob{value: val, isSet: true}
}

func (v NullableApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
