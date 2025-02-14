# RunnerJobAdmin

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
**PrivatePayload** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewRunnerJobAdmin

`func NewRunnerJobAdmin() *RunnerJobAdmin`

NewRunnerJobAdmin instantiates a new RunnerJobAdmin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunnerJobAdminWithDefaults

`func NewRunnerJobAdminWithDefaults() *RunnerJobAdmin`

NewRunnerJobAdminWithDefaults instantiates a new RunnerJobAdmin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *RunnerJobAdmin) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RunnerJobAdmin) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RunnerJobAdmin) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *RunnerJobAdmin) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetType

`func (o *RunnerJobAdmin) GetType() RunnerJobType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RunnerJobAdmin) GetTypeOk() (*RunnerJobType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RunnerJobAdmin) SetType(v RunnerJobType)`

SetType sets Type field to given value.

### HasType

`func (o *RunnerJobAdmin) HasType() bool`

HasType returns a boolean if a field has been set.

### GetState

`func (o *RunnerJobAdmin) GetState() RunnerJobStateConstant`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RunnerJobAdmin) GetStateOk() (*RunnerJobStateConstant, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RunnerJobAdmin) SetState(v RunnerJobStateConstant)`

SetState sets State field to given value.

### HasState

`func (o *RunnerJobAdmin) HasState() bool`

HasState returns a boolean if a field has been set.

### GetPayload

`func (o *RunnerJobAdmin) GetPayload() RunnerJobPayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *RunnerJobAdmin) GetPayloadOk() (*RunnerJobPayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *RunnerJobAdmin) SetPayload(v RunnerJobPayload)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *RunnerJobAdmin) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetFailures

`func (o *RunnerJobAdmin) GetFailures() int32`

GetFailures returns the Failures field if non-nil, zero value otherwise.

### GetFailuresOk

`func (o *RunnerJobAdmin) GetFailuresOk() (*int32, bool)`

GetFailuresOk returns a tuple with the Failures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailures

`func (o *RunnerJobAdmin) SetFailures(v int32)`

SetFailures sets Failures field to given value.

### HasFailures

`func (o *RunnerJobAdmin) HasFailures() bool`

HasFailures returns a boolean if a field has been set.

### GetError

`func (o *RunnerJobAdmin) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *RunnerJobAdmin) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *RunnerJobAdmin) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *RunnerJobAdmin) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *RunnerJobAdmin) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *RunnerJobAdmin) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetProgress

`func (o *RunnerJobAdmin) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *RunnerJobAdmin) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *RunnerJobAdmin) SetProgress(v int32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *RunnerJobAdmin) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetPriority

`func (o *RunnerJobAdmin) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *RunnerJobAdmin) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *RunnerJobAdmin) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *RunnerJobAdmin) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RunnerJobAdmin) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RunnerJobAdmin) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RunnerJobAdmin) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RunnerJobAdmin) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RunnerJobAdmin) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RunnerJobAdmin) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RunnerJobAdmin) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RunnerJobAdmin) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetStartedAt

`func (o *RunnerJobAdmin) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *RunnerJobAdmin) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *RunnerJobAdmin) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *RunnerJobAdmin) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetFinishedAt

`func (o *RunnerJobAdmin) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *RunnerJobAdmin) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *RunnerJobAdmin) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *RunnerJobAdmin) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetParent

`func (o *RunnerJobAdmin) GetParent() RunnerJobParent`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *RunnerJobAdmin) GetParentOk() (*RunnerJobParent, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *RunnerJobAdmin) SetParent(v RunnerJobParent)`

SetParent sets Parent field to given value.

### HasParent

`func (o *RunnerJobAdmin) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *RunnerJobAdmin) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *RunnerJobAdmin) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetRunner

`func (o *RunnerJobAdmin) GetRunner() RunnerJobRunner`

GetRunner returns the Runner field if non-nil, zero value otherwise.

### GetRunnerOk

`func (o *RunnerJobAdmin) GetRunnerOk() (*RunnerJobRunner, bool)`

GetRunnerOk returns a tuple with the Runner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunner

`func (o *RunnerJobAdmin) SetRunner(v RunnerJobRunner)`

SetRunner sets Runner field to given value.

### HasRunner

`func (o *RunnerJobAdmin) HasRunner() bool`

HasRunner returns a boolean if a field has been set.

### SetRunnerNil

`func (o *RunnerJobAdmin) SetRunnerNil(b bool)`

 SetRunnerNil sets the value for Runner to be an explicit nil

### UnsetRunner
`func (o *RunnerJobAdmin) UnsetRunner()`

UnsetRunner ensures that no value is present for Runner, not even an explicit nil
### GetPrivatePayload

`func (o *RunnerJobAdmin) GetPrivatePayload() map[string]interface{}`

GetPrivatePayload returns the PrivatePayload field if non-nil, zero value otherwise.

### GetPrivatePayloadOk

`func (o *RunnerJobAdmin) GetPrivatePayloadOk() (*map[string]interface{}, bool)`

GetPrivatePayloadOk returns a tuple with the PrivatePayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivatePayload

`func (o *RunnerJobAdmin) SetPrivatePayload(v map[string]interface{})`

SetPrivatePayload sets PrivatePayload field to given value.

### HasPrivatePayload

`func (o *RunnerJobAdmin) HasPrivatePayload() bool`

HasPrivatePayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


