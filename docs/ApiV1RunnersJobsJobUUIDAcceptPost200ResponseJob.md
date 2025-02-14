# ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**RunnerJobType**](RunnerJobType.md) |  | [optional] 
**State** | Pointer to [**RunnerJobStateConstant**](RunnerJobStateConstant.md) |  | [optional] 
**Payload** | Pointer to [**RunnerJobPayload**](RunnerJobPayload.md) |  | [optional] 
**Failures** | Pointer to **int32** | Number of times a remote runner failed to process this job. After too many failures, the job in \&quot;error\&quot; state | [optional] 
**Error** | Pointer to **NullableString** | Error message if the job is errored | [optional] 
**Progress** | Pointer to **int32** | Percentage progress | [optional] 
**Priority** | Pointer to **int32** | Job priority (less has more priority) | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**FinishedAt** | Pointer to **time.Time** |  | [optional] 
**Parent** | Pointer to [**NullableRunnerJobParent**](RunnerJobParent.md) |  | [optional] 
**Runner** | Pointer to [**NullableRunnerJobRunner**](RunnerJobRunner.md) |  | [optional] 
**JobToken** | Pointer to **string** |  | [optional] 

## Methods

### NewApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob

`func NewApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob() *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob`

NewApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob instantiates a new ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1RunnersJobsJobUUIDAcceptPost200ResponseJobWithDefaults

`func NewApiV1RunnersJobsJobUUIDAcceptPost200ResponseJobWithDefaults() *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob`

NewApiV1RunnersJobsJobUUIDAcceptPost200ResponseJobWithDefaults instantiates a new ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetType

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetType() RunnerJobType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetTypeOk() (*RunnerJobType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) SetType(v RunnerJobType)`

SetType sets Type field to given value.

### HasType

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) HasType() bool`

HasType returns a boolean if a field has been set.

### GetState

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetState() RunnerJobStateConstant`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetStateOk() (*RunnerJobStateConstant, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) SetState(v RunnerJobStateConstant)`

SetState sets State field to given value.

### HasState

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) HasState() bool`

HasState returns a boolean if a field has been set.

### GetPayload

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetPayload() RunnerJobPayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetPayloadOk() (*RunnerJobPayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) SetPayload(v RunnerJobPayload)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetFailures

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetFailures() int32`

GetFailures returns the Failures field if non-nil, zero value otherwise.

### GetFailuresOk

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetFailuresOk() (*int32, bool)`

GetFailuresOk returns a tuple with the Failures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailures

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) SetFailures(v int32)`

SetFailures sets Failures field to given value.

### HasFailures

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) HasFailures() bool`

HasFailures returns a boolean if a field has been set.

### GetError

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetProgress

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) SetProgress(v int32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetPriority

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetStartedAt

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetFinishedAt

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetParent

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetParent() RunnerJobParent`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetParentOk() (*RunnerJobParent, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) SetParent(v RunnerJobParent)`

SetParent sets Parent field to given value.

### HasParent

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetRunner

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetRunner() RunnerJobRunner`

GetRunner returns the Runner field if non-nil, zero value otherwise.

### GetRunnerOk

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetRunnerOk() (*RunnerJobRunner, bool)`

GetRunnerOk returns a tuple with the Runner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunner

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) SetRunner(v RunnerJobRunner)`

SetRunner sets Runner field to given value.

### HasRunner

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) HasRunner() bool`

HasRunner returns a boolean if a field has been set.

### SetRunnerNil

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) SetRunnerNil(b bool)`

 SetRunnerNil sets the value for Runner to be an explicit nil

### UnsetRunner
`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) UnsetRunner()`

UnsetRunner ensures that no value is present for Runner, not even an explicit nil
### GetJobToken

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetJobToken() string`

GetJobToken returns the JobToken field if non-nil, zero value otherwise.

### GetJobTokenOk

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) GetJobTokenOk() (*string, bool)`

GetJobTokenOk returns a tuple with the JobToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobToken

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) SetJobToken(v string)`

SetJobToken sets JobToken field to given value.

### HasJobToken

`func (o *ApiV1RunnersJobsJobUUIDAcceptPost200ResponseJob) HasJobToken() bool`

HasJobToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


