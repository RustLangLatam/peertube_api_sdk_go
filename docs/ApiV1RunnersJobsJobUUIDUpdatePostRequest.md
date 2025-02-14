# ApiV1RunnersJobsJobUUIDUpdatePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RunnerToken** | **string** |  | 
**JobToken** | **string** |  | 
**Progress** | Pointer to **int32** | Update job progression percentage (optional) | [optional] 
**Payload** | Pointer to [**ApiV1RunnersJobsJobUUIDUpdatePostRequestPayload**](ApiV1RunnersJobsJobUUIDUpdatePostRequestPayload.md) |  | [optional] 

## Methods

### NewApiV1RunnersJobsJobUUIDUpdatePostRequest

`func NewApiV1RunnersJobsJobUUIDUpdatePostRequest(runnerToken string, jobToken string, ) *ApiV1RunnersJobsJobUUIDUpdatePostRequest`

NewApiV1RunnersJobsJobUUIDUpdatePostRequest instantiates a new ApiV1RunnersJobsJobUUIDUpdatePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1RunnersJobsJobUUIDUpdatePostRequestWithDefaults

`func NewApiV1RunnersJobsJobUUIDUpdatePostRequestWithDefaults() *ApiV1RunnersJobsJobUUIDUpdatePostRequest`

NewApiV1RunnersJobsJobUUIDUpdatePostRequestWithDefaults instantiates a new ApiV1RunnersJobsJobUUIDUpdatePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRunnerToken

`func (o *ApiV1RunnersJobsJobUUIDUpdatePostRequest) GetRunnerToken() string`

GetRunnerToken returns the RunnerToken field if non-nil, zero value otherwise.

### GetRunnerTokenOk

`func (o *ApiV1RunnersJobsJobUUIDUpdatePostRequest) GetRunnerTokenOk() (*string, bool)`

GetRunnerTokenOk returns a tuple with the RunnerToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunnerToken

`func (o *ApiV1RunnersJobsJobUUIDUpdatePostRequest) SetRunnerToken(v string)`

SetRunnerToken sets RunnerToken field to given value.


### GetJobToken

`func (o *ApiV1RunnersJobsJobUUIDUpdatePostRequest) GetJobToken() string`

GetJobToken returns the JobToken field if non-nil, zero value otherwise.

### GetJobTokenOk

`func (o *ApiV1RunnersJobsJobUUIDUpdatePostRequest) GetJobTokenOk() (*string, bool)`

GetJobTokenOk returns a tuple with the JobToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobToken

`func (o *ApiV1RunnersJobsJobUUIDUpdatePostRequest) SetJobToken(v string)`

SetJobToken sets JobToken field to given value.


### GetProgress

`func (o *ApiV1RunnersJobsJobUUIDUpdatePostRequest) GetProgress() int32`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *ApiV1RunnersJobsJobUUIDUpdatePostRequest) GetProgressOk() (*int32, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *ApiV1RunnersJobsJobUUIDUpdatePostRequest) SetProgress(v int32)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *ApiV1RunnersJobsJobUUIDUpdatePostRequest) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetPayload

`func (o *ApiV1RunnersJobsJobUUIDUpdatePostRequest) GetPayload() ApiV1RunnersJobsJobUUIDUpdatePostRequestPayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *ApiV1RunnersJobsJobUUIDUpdatePostRequest) GetPayloadOk() (*ApiV1RunnersJobsJobUUIDUpdatePostRequestPayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *ApiV1RunnersJobsJobUUIDUpdatePostRequest) SetPayload(v ApiV1RunnersJobsJobUUIDUpdatePostRequestPayload)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *ApiV1RunnersJobsJobUUIDUpdatePostRequest) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


