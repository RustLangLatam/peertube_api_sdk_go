# RunnerJob

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

## Methods

### NewRunnerJob

`func NewRunnerJob() *RunnerJob`

NewRunnerJob instantiates a new RunnerJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunnerJobWithDefaults

`func NewRunnerJobWithDefaults() *RunnerJob`

NewRunnerJobWithDefaults instantiates a new RunnerJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *RunnerJob) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RunnerJob) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RunnerJob) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *RunnerJob) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetType

`func (o *RunnerJob) GetType() RunnerJobType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RunnerJob) GetTypeOk() (*RunnerJobType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RunnerJob) SetType(v RunnerJobType)`

SetType sets Type field to given value.

### HasType

`func (o *RunnerJob) HasType() bool`

HasType returns a boolean if a field has been set.

### GetState

`func (o *RunnerJob) GetState() RunnerJobStateConstant`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RunnerJob) GetStateOk() (*RunnerJobStateConstant, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RunnerJob) SetState(v RunnerJobStateConstant)`

SetState sets State field to given value.

### HasState

`func (o *RunnerJob) HasState() bool`

HasState returns a boolean if a field has been set.

### GetPayload

`func (o *RunnerJob) GetPayload() RunnerJobPayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *RunnerJob) GetPayloadOk() (*RunnerJobPayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *RunnerJob) SetPayload(v RunnerJobPayload)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *RunnerJob) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetFailures

`func (o *RunnerJob) GetFailures() int32`

GetFailures returns the Failures field if non-nil, zero value otherwise.

### GetFailuresOk

`func (o *RunnerJob) GetFailuresOk() (*int32, bool)`

GetFailuresOk returns a tuple with the Failures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailures

`func (o *RunnerJob) SetFailures(v int32)`

SetFailures sets Failures field to given value.

### HasFailures

`func (o *RunnerJob) HasFailures() bool`

HasFailures returns a boolean if a field has been set.

### GetError

`func (o *RunnerJob) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RunnerJob) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RunnerJob) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *RunnerJob) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *RunnerJob) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *RunnerJob) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetProgress

`func (o *RunnerJob) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *RunnerJob) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *RunnerJob) SetProgress(v int32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *RunnerJob) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetPriority

`func (o *RunnerJob) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *RunnerJob) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *RunnerJob) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *RunnerJob) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RunnerJob) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RunnerJob) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RunnerJob) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RunnerJob) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RunnerJob) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RunnerJob) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RunnerJob) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RunnerJob) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetStartedAt

`func (o *RunnerJob) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *RunnerJob) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *RunnerJob) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *RunnerJob) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetFinishedAt

`func (o *RunnerJob) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *RunnerJob) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *RunnerJob) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *RunnerJob) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetParent

`func (o *RunnerJob) GetParent() RunnerJobParent`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *RunnerJob) GetParentOk() (*RunnerJobParent, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *RunnerJob) SetParent(v RunnerJobParent)`

SetParent sets Parent field to given value.

### HasParent

`func (o *RunnerJob) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *RunnerJob) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *RunnerJob) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetRunner

`func (o *RunnerJob) GetRunner() RunnerJobRunner`

GetRunner returns the Runner field if non-nil, zero value otherwise.

### GetRunnerOk

`func (o *RunnerJob) GetRunnerOk() (*RunnerJobRunner, bool)`

GetRunnerOk returns a tuple with the Runner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunner

`func (o *RunnerJob) SetRunner(v RunnerJobRunner)`

SetRunner sets Runner field to given value.

### HasRunner

`func (o *RunnerJob) HasRunner() bool`

HasRunner returns a boolean if a field has been set.

### SetRunnerNil

`func (o *RunnerJob) SetRunnerNil(b bool)`

 SetRunnerNil sets the value for Runner to be an explicit nil

### UnsetRunner
`func (o *RunnerJob) UnsetRunner()`

UnsetRunner ensures that no value is present for Runner, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


